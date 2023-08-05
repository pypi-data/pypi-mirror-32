import os
import uuid
from datetime import datetime

from selenium.common.exceptions import WebDriverException
from selenium.webdriver.remote.webdriver import WebDriver as RemoteWebDriver

# noinspection PyProtectedMember
from applitools import logger, _viewport_size
from applitools.common import StitchMode
from applitools.errors import DiffsFoundError
from applitools.geometry import Region
from applitools.test_results import TestResultsStatus
from ._agent_connector import AgentConnector
from ._webdriver import EyesWebDriver, build_position_provider_for, ElementPositionProvider
from ._match_window_task import MatchWindowTask
from ._triggers import TextTrigger, MouseTrigger
from ._webdriver import EyesFrame
from .errors import EyesError, NewTestError, TestFailedError
from .test_results import TestResults
from .utils import general_utils
from applitools import VERSION


class FailureReports(object):
    """
    Failures are either reported immediately when they are detected, or when the test is closed.
    """
    IMMEDIATE = "Immediate"
    ON_CLOSE = "OnClose"


class MatchLevel(object):
    """
    The extent in which two images match (or are expected to match).
    """
    NONE = "None"
    LEGACY_LAYOUT = "Layout1"
    LAYOUT = "Layout2"
    LAYOUT2 = "Layout2"
    CONTENT = "Content"
    STRICT = "Strict"
    EXACT = "Exact"


class ExactMatchSettings(object):
    """
    Encapsulates settings for the "Exact" match level.
    """

    def __init__(self, min_diff_intensity=0, min_diff_width=0, min_diff_height=0, match_threshold=0.0):
        """
        Ctor.

        :param min_diff_intensity: Minimal non-ignorable pixel intensity difference.
        :param min_diff_width: Minimal non-ignorable diff region width.
        :param min_diff_height: Minimal non-ignorable diff region height.
        :param match_threshold: The ratio of differing pixels above which images are considered mismatching.
        """
        self.min_diff_intensity = min_diff_intensity
        self.min_diff_width = min_diff_width
        self.min_diff_height = min_diff_height
        self.match_threshold = match_threshold

    def __getstate__(self):
        return dict(minDiffIntensity=self.min_diff_intensity,
                    minDiffWidth=self.min_diff_width,
                    minDiffHeight=self.min_diff_height,
                    matchThreshold=self.match_threshold)

    # This is required in order for jsonpickle to work on this object.
    # noinspection PyMethodMayBeStatic
    def __setstate__(self, state):
        raise EyesError('Cannot create ExactMatchSettings instance from dict!')

    def __str__(self):
        return "[min diff intensity: %d, min diff width: %d, min diff height: %d, match threshold: %f]" % (
            self.min_diff_intensity, self.min_diff_width, self.min_diff_height, self.match_threshold)


class ImageMatchSettings(object):
    """
    Encapsulates match settings for the a session.
    """

    def __init__(self, match_level=MatchLevel.STRICT, exact_settings=None):
        """

        :param match_level: The "strictness" level of the match.
        :param exact_settings: Parameter for fine tuning the match when "Exact" match level is used.
        """
        self.match_level = match_level
        self.exact_settings = exact_settings

    def __getstate__(self):
        return dict(matchLevel=self.match_level, exact=self.exact_settings)

    # This is required in order for jsonpickle to work on this object.
    # noinspection PyMethodMayBeStatic
    def __setstate__(self, state):
        raise EyesError('Cannot create ImageMatchSettings instance from dict!')

    def __str__(self):
        return "[Match level: %s, Exact match settings: %s]" % (self.match_level, self.exact_settings)


class BatchInfo(object):
    """
    A batch of tests.
    """

    def __init__(self, name=None, started_at=datetime.now(general_utils.UTC)):
        self.name = name if name else os.environ.get('APPLITOOLS_BATCH_NAME', None)
        self.started_at = started_at
        self.id_ = os.environ.get('APPLITOOLS_BATCH_ID', str(uuid.uuid4()))

    def __getstate__(self):
        return dict(name=self.name, startedAt=self.started_at.isoformat(), id=self.id_)

    # Required is required in order for jsonpickle to work on this object.
    # noinspection PyMethodMayBeStatic
    def __setstate__(self, state):
        raise EyesError('Cannot create BatchInfo instance from dict!')

    def __str__(self):
        return "%s - %s - %s" % (self.name, self.started_at, self.id_)


class Eyes(object):
    """
    Applitools Selenium Eyes API for python.
    """
    _DEFAULT_MATCH_TIMEOUT = 2000  # Milliseconds
    _DEFAULT_WAIT_BEFORE_SCREENSHOTS = 100  # ms
    BASE_AGENT_ID = "eyes.selenium.python/%s" % VERSION
    DEFAULT_EYES_SERVER = 'https://eyessdk.applitools.com'

    @staticmethod
    def set_viewport_size(driver, viewport_size):
        _viewport_size.set_viewport_size(driver, viewport_size)

    def __init__(self, server_url=DEFAULT_EYES_SERVER):
        """
        Creates a new (possibly disabled) Eyes instance that interacts with the Eyes server.

        :param server_url: The URL of the Eyes server
        """
        self.is_disabled = False
        """(Boolean) Disables Applitools Eyes and uses the webdriver directly."""
        self._user_inputs = []
        self._running_session = None
        self._agent_connector = AgentConnector(server_url)
        self._should_get_title = False
        self._driver = None
        self._match_window_task = None
        self._is_open = False
        self._app_name = None
        self._last_screenshot = None
        self._should_match_once_on_timeout = False
        self._start_info = None
        self._test_name = None
        self._viewport_size = None
        self._match_timeout = Eyes._DEFAULT_MATCH_TIMEOUT
        self._stitch_mode = StitchMode.Scroll
        self._properties = []
        """(list) key-value pairs to be associated with the test. Can be used for filtering later."""
        self.agent_id = None
        """(String) An optional string identifying the current library using the SDK."""
        self.failure_reports = FailureReports.ON_CLOSE
        """(String) Should the test report mismatches immediately or when it is finished. See FailureReports."""
        self.default_match_settings = ImageMatchSettings()
        """(ImageMatchSettings) The default match settings for the session. See ImageMatchSettings"""
        self.batch = None
        """(String) The batch to which the tests belong to. See BatchInfo. None means no batch."""
        self.host_os = None
        """(String) A string identifying the OS running the AUT. Use this to override Eyes automatic inference."""
        self.host_app = None
        """(String) A string identifying the app running the AUT. Use this to override Eyes automatic inference."""
        self.baseline_name = None
        """(String) A string that, if specified, determines the baseline to compare with and disables automatic baseline
        inference."""
        self.save_new_tests = True
        """(Boolean) A boolean denoting whether new tests should be automatically accepted."""
        self.save_failed_tests = False
        """(Boolean) Whether failed tests should be automatically saved with all new output accepted."""
        self.branch_name = None
        """(String) A string identifying the branch in which tests are run."""
        self.parent_branch_name = None
        """(String) A string identifying the parent branch of the branch set by "branch_name"."""
        self.force_full_page_screenshot = False
        """(Boolean) if true, Eyes will create a full page screenshot (by using stitching) for browsers which only
        returns the viewport screenshot."""
        self.hide_scrollbars = False
        """(Boolean) if true, Eyes will remove the scrollbars from the pages before taking the screenshot."""
        self.fail_on_new_test = False
        """(Boolean) If true, Eyes will treat new tests the same as failed tests."""
        self.wait_before_screenshots = Eyes._DEFAULT_WAIT_BEFORE_SCREENSHOTS
        """(int) The number of milliseconds to wait before each time a screenshot is taken."""

    def add_property(self, name, value):
        """
        Associates a key/value pair with the test. This can be used later for filtering.
        :param name: (string) The property name.
        :param value: (string) The property value
        :return: None
        """
        self._properties.append({'name': name, 'value': value})

    @property
    def match_level(self):
        """
        Gets the default match level for the entire session. See ImageMatchSettings.
        """
        return self.default_match_settings.match_level

    @match_level.setter
    def match_level(self, match_level):
        """
        Sets the default match level for the entire session. See ImageMatchSettings.

        :param match_level: The match level to set. Should be one of the values defined by MatchLevel
        """
        self.default_match_settings.match_level = match_level

    @property
    def stitch_mode(self):
        """
        Gets the stitch mode.

        :return: The stitch mode.
        """
        return self._stitch_mode

    @stitch_mode.setter
    def stitch_mode(self, stitch_mode):
        """
        Sets the stitch property - default is by scrolling.

        :param stitch_mode: The stitch mode to set - either scrolling or css.
        """
        self._stitch_mode = stitch_mode
        if stitch_mode == StitchMode.CSS:
            self.hide_scrollbars = True

    @property
    def match_timeout(self):
        """
        Gets the default timeout for check_XXXX operations. (milliseconds)

        :return: The match timeout (milliseconds)
        """
        return self._match_timeout

    @match_timeout.setter
    def match_timeout(self, match_timeout):
        """
        Sets the default timeout for check_XXXX operations. (milliseconds)
        """
        if 0 < match_timeout < MatchWindowTask.MINIMUM_MATCH_TIMEOUT:
            raise ValueError("Match timeout must be at least 60ms.")
        self._match_timeout = match_timeout

    @property
    def api_key(self):
        """
        Gets the Api key used for authenticating the user with Eyes.

        :return: The Api key used for authenticating the user with Eyes.
        """
        return self._agent_connector.api_key

    @api_key.setter
    def api_key(self, api_key):
        """
        Sets the api key used for authenticating the user with Eyes.

        :param api_key: The api key used for authenticating the user with Eyes.
        """
        self._agent_connector.api_key = api_key

    @property
    def server_url(self):
        """
        Gets the URL of the Eyes server.

        :return: The URL of the Eyes server, or None to use the default server.
        """
        return self._agent_connector.server_url

    @server_url.setter
    def server_url(self, server_url):
        """
        Sets the URL of the Eyes server.

        :param server_url: The URL of the Eyes server, or None to use the default server.
        :return: None
        """
        if server_url is None:
            self._agent_connector.server_url = Eyes.DEFAULT_EYES_SERVER
        else:
            self._agent_connector.server_url = server_url

    @property
    def _full_agent_id(self):
        """
        Gets the agent id, which identifies the current library using the SDK.

        :return: The agent id.
        """
        if self.agent_id is None:
            return self.BASE_AGENT_ID
        return "%s [%s]" % (self.agent_id, self.BASE_AGENT_ID)

    def is_open(self):
        """
        Returns whether the session is currently running.

        :return: True is a session is running, False otherwise.
        """
        return self._is_open

    def get_driver(self):
        """
        Returns the current web driver.

        :return: (selenium.webdriver.remote.webdriver) The web driver currently used by the Eyes
                                                    instance.
        """
        return self._driver

    def get_viewport_size(self):
        """
        Returns the size of the viewport.

        :return: ({width, height}) The size of the viewport of the application under test (e.g,
                                the browser).
        """
        return self._driver.get_viewport_size()

    def abort_if_not_closed(self):
        """
        If a test is running, aborts it. Otherwise, does nothing.
        """
        if self.is_disabled:
            logger.debug('abort_if_not_closed(): ignored (disabled)')
            return
        try:
            self._reset_last_screenshot()

            if self._running_session:
                logger.debug('abort_if_not_closed(): Aborting session...')
                try:
                    self._agent_connector.stop_session(self._running_session, True, False)
                    logger.info('--- Test aborted.')
                except EyesError as e:
                    logger.info("Failed to abort server session: %s " % e)
                    pass
                finally:
                    self._running_session = None
        finally:
            logger.close()

    def open(self, driver, app_name, test_name, viewport_size=None):
        """
        Starts a test.

        :param driver: The webdriver to use.
        :param app_name: The name of the application under test.
        :param test_name: The test name.
        :param viewport_size: The client's viewport size (i.e., the visible part of the document's body) or None to
                                allow any viewport size.
        :return: An updated web driver
        :raise EyesError: If the session was already open.
        """
        logger.open_()
        if self.is_disabled:
            logger.debug('open(): ignored (disabled)')
            return driver

        if self.api_key is None:
            try:
                self.api_key = os.environ['APPLITOOLS_API_KEY']
            except KeyError:
                raise EyesError("API key not set! Log in to https://applitools.com to obtain your"
                                " API Key and use 'api_key' to set it.")

        if isinstance(driver, EyesWebDriver):
            # If the driver is an EyesWebDriver (as might be the case when tests are ran
            # consecutively using the same driver object)
            self._driver = driver
        else:
            if not isinstance(driver, RemoteWebDriver):
                logger.info("WARNING: driver is not a RemoteWebDriver (class: {0})".format(driver.__class__))
            self._driver = EyesWebDriver(driver, self, self._stitch_mode)

        logger.info("open(%s, %s, %s, %s)" % (app_name, test_name, viewport_size, self.failure_reports))

        if self.is_open():
            self.abort_if_not_closed()
            raise EyesError('a test is already running')
        self._app_name = app_name
        self._test_name = test_name
        self._viewport_size = viewport_size
        self._is_open = True
        return self._driver

    def _assign_viewport_size(self):
        """
        Assign the viewport size we need to be in the default content frame.
        """
        original_frame_chain = self._driver.get_frame_chain()
        self._driver.switch_to.default_content()
        try:
            if self._viewport_size:
                logger.debug("Assigning viewport size {0}".format(self._viewport_size))
                _viewport_size.set_viewport_size(self._driver, self._viewport_size)
            else:
                logger.debug("No viewport size given. Extracting the viewport size from the driver...")
                self._viewport_size = _viewport_size.get_viewport_size(self._driver)
                logger.debug("Viewport size {0}".format(self._viewport_size))
        except EyesError as e:
            # Going back to the frame we started at
            self._driver.switch_to.frames(original_frame_chain)
            raise TestFailedError('Failed to assign viewport size!')

    def _get_environment(self):
        """
        Application environment is the environment (e.g., the host OS) which runs the application under test.

        :return: The current application environment.
        """
        os = self.host_os
        # If no host OS was set, check for mobile OS.
        if os is None:
            logger.info('No OS set, checking for mobile OS...')
            # Since in Python Appium driver is the same for Android and iOS, we need to use the desired
            # capabilities to figure this out.
            if self._driver.is_mobile_device():
                platform_name = self._driver.get_platform_name()
                logger.info(platform_name + ' detected')
                platform_version = self._driver.get_platform_version()
                if platform_version is not None:
                    # Notice that Python's "split" function's +limit+ is the the maximum splits performed
                    # whereas in Ruby it is the maximum number of elements in the result (which is why they are set
                    # differently).
                    major_version = platform_version.split('.', 1)[0]
                    os = platform_name + ' ' + major_version
                else:
                    os = platform_name
                logger.info("Setting OS: " + os)
            else:
                logger.info('No mobile OS detected.')
        app_env = {'os': os, 'hostingApp': self.host_app,
                   'displaySize': self._viewport_size,
                   'inferred': self._get_inferred_environment()}
        return app_env

    def _create_start_info(self):
        app_env = self._get_environment()
        self._start_info = {'agentId': self._full_agent_id, 'appIdOrName': self._app_name,
                            'scenarioIdOrName': self._test_name, 'batchInfo': self.batch,
                            'envName': self.baseline_name, 'environment': app_env,
                            'defaultMatchSettings': self.default_match_settings, 'verId': None,
                            'branchName': self.branch_name, 'parentBranchName': self.parent_branch_name,
                            'properties': self._properties}

    def _start_session(self):
        logger.debug("_start_session()")
        self._assign_viewport_size()

        # initialization of Eyes parameters if empty from ENV variables
        if not self.branch_name:
            self.branch_name = os.environ.get('APPLITOOLS_BRANCH', None)
        if not self.baseline_name:
            self.baseline_name = os.environ.get('APPLITOOLS_BASELINE_BRANCH', None)
        if not self.parent_branch_name:
            self.parent_branch_name = os.environ.get('APPLITOOLS_PARENT_BRANCH', None)
        if not self.batch:
            self.batch = BatchInfo()

        self._create_start_info()
        # Actually start the session.
        self._running_session = self._agent_connector.start_session(self._start_info)
        self._should_match_once_on_timeout = self._running_session['is_new_session']

    def get_title(self):
        """
        Returns the title of the window.

        :return: The title of the window of the AUT, or empty string if the title is not
                    available.
        """
        if self._should_get_title:
            # noinspection PyBroadException
            try:
                return self._driver.title
            except:
                self._should_get_title = False
                # Couldn't get title, return empty string.
        return ''

    def _get_inferred_environment(self):
        try:
            user_agent = self._driver.execute_script('return navigator.userAgent')
        except WebDriverException:
            user_agent = None
        if user_agent:
            return "useragent:%s" % user_agent
        return None

    def _prepare_to_check(self):
        logger.debug("_prepare_to_check()")
        if not self.is_open():
            raise EyesError('Eyes not open!')

        if not self._running_session:
            self._start_session()
            self._match_window_task = MatchWindowTask(self, self._agent_connector,
                                                      self._running_session, self._driver,
                                                      self.match_timeout)

    def _handle_match_result(self, result, tag):
        self._last_screenshot = result['screenshot']
        as_expected = result['as_expected']
        self._user_inputs = []
        if not as_expected:
            self._should_match_once_on_timeout = True
            if not self._running_session['is_new_session']:
                logger.info("Window mismatch %s" % tag)
                if self.failure_reports == FailureReports.IMMEDIATE:
                    raise TestFailedError("Mismatch found in '%s' of '%s'" %
                                          (self._start_info['scenarioIdOrName'],
                                           self._start_info['appIdOrName']))

    def check_window(self, tag=None, match_timeout=-1, target=None):
        """
        Takes a snapshot from the browser using the web driver and matches it with the expected
        output.

        :param tag: (str) Description of the visual validation checkpoint.
        :param match_timeout: (int) Timeout for the visual validation checkpoint (milliseconds).
        :param target: (Target) The target for the check_window call
        :return: None
        """
        if self.is_disabled:
            logger.info("check_window(%s): ignored (disabled)" % tag)
            return
        logger.info("check_window('%s')" % tag)
        if self.hide_scrollbars:
            original_overflow = self._driver.hide_scrollbars()
        self._prepare_to_check()
        result = self._match_window_task.match_window(match_timeout, tag,
                                                      self.force_full_page_screenshot,
                                                      self._user_inputs,
                                                      self.wait_before_screenshots,
                                                      self.default_match_settings,
                                                      target,
                                                      self._should_match_once_on_timeout)
        if self.hide_scrollbars:
            # noinspection PyUnboundLocalVariable
            self._driver.set_overflow(original_overflow)
        self._handle_match_result(result, tag)

    def check_region(self, region, tag=None, match_timeout=-1, target=None):
        """
        Takes a snapshot of the given region from the browser using the web driver and matches it
        with the expected output. If the current context is a frame, the region is offsetted
        relative to the frame.

        :param region: (Region) The region which will be visually validated. The coordinates are
                         relative to the viewport of the current frame.
        :param tag: (str) Description of the visual validation checkpoint.
        :param match_timeout: (int) Timeout for the visual validation checkpoint (milliseconds).
        :param target: (Target) The target for the check_window call
        :return: None
        """
        if self.is_disabled:
            logger.info('check_region(): ignored (disabled)')
            return
        logger.info("check_region([%s], '%s')" % (region, tag))
        if region.is_empty():
            raise EyesError("region cannot be empty!")
        if self.hide_scrollbars:
            original_overflow = self._driver.hide_scrollbars()
        self._prepare_to_check()
        result = self._match_window_task.match_region(region, match_timeout, tag,
                                                      self.force_full_page_screenshot,
                                                      self._user_inputs,
                                                      self.wait_before_screenshots,
                                                      self.default_match_settings,
                                                      target,
                                                      self._should_match_once_on_timeout)
        if self.hide_scrollbars:
            # noinspection PyUnboundLocalVariable
            self._driver.set_overflow(original_overflow)
        self._handle_match_result(result, tag)

    def check_region_by_element(self, element, tag=None, match_timeout=-1, target=None, stitch_content=False):
        """
        Takes a snapshot of the region of the given element from the browser using the web driver
        and matches it with the expected output.

        :param element: (WebElement)  The element which region will be visually validated.
        :param tag: (str) Description of the visual validation checkpoint.
        :param match_timeout: (int) Timeout for the visual validation checkpoint (milliseconds).
        :param target: (Target) The target for the check_window call
        :return: None
        """
        if self.is_disabled:
            logger.info('check_region_by_element(): ignored (disabled)')
            return
        logger.info("check_region_by_element('%s')" % tag)
        if self.hide_scrollbars:
            original_overflow = self._driver.hide_scrollbars()
        self._prepare_to_check()
        result = self._match_window_task.match_element(element, match_timeout, tag,
                                                       self.force_full_page_screenshot,
                                                       self._user_inputs,
                                                       self.wait_before_screenshots,
                                                       self.default_match_settings,
                                                       target,
                                                       self._should_match_once_on_timeout,
                                                       stitch_content=stitch_content)

        if self.hide_scrollbars:
            # noinspection PyUnboundLocalVariable
            self._driver.set_overflow(original_overflow)
        self._handle_match_result(result, tag)

    def check_region_by_selector(self, by, value, tag=None, match_timeout=-1, target=None, stitch_content=False):
        """
        Takes a snapshot of the region of the element found by calling find_element(by, value)
        and matches it with the expected output.

        :param by: (By) The way by which an element to be validated should be found (e.g., By.ID).
        :param value: (str) The value identifying the element using the "by" type.
        :param tag: (str) Description of the visual validation checkpoint.
        :param match_timeout: (int) Timeout for the visual validation checkpoint (milliseconds).
        :param target: (Target) The target for the check_window call
        :return: None
        """
        if self.is_disabled:
            logger.info('check_region_by_selector(): ignored (disabled)')
            return
        logger.debug("calling 'check_region_by_element'...")
        self.check_region_by_element(self._driver.find_element(by, value), tag,
                                     match_timeout, target, stitch_content)

    def check_region_in_frame_by_selector(self, frame_reference, by, value, tag=None,
                                          match_timeout=-1, target=None, stitch_content=False):
        """
        Checks a region within a frame, and returns to the current frame.

        :param frame_reference: (int/str/WebElement) A reference to the frame in which the region should be checked.
        :param by: (By) The way by which an element to be validated should be found (e.g., By.ID).
        :param value: (str) The value identifying the element using the "by" type.
        :param tag: (str) Description of the visual validation checkpoint.
        :param match_timeout: (int) Timeout for the visual validation checkpoint (milliseconds).
        :param target: (Target) The target for the check_window call
        :return: None
        """
        if self.is_disabled:
            logger.info('check_region_in_frame_by_selector(): ignored (disabled)')
            return
        logger.info("check_region_in_frame_by_selector('%s')" % tag)
        # We need to temporarily save the hide_scrollbars value, since we'll change it to make sure that hide_scrollbars
        # will NOT be called twice (once outside the frame and once inside the frame).
        original_hide_scrollbars_value = self.hide_scrollbars
        if self.hide_scrollbars:
            original_overflow = self._driver.hide_scrollbars()
            self.hide_scrollbars = False
        # Switching to the relevant frame
        self._driver.switch_to.frame(frame_reference)
        logger.debug("calling 'check_region_by_selector'...")
        self.check_region_by_selector(by, value, tag, match_timeout, target, stitch_content)
        # Switching back to our original frame
        self._driver.switch_to.parent_frame()
        if original_hide_scrollbars_value:
            # noinspection PyUnboundLocalVariable
            self._driver.set_overflow(original_overflow)
            self.hide_scrollbars = original_hide_scrollbars_value

    def _reset_last_screenshot(self):
        self._last_screenshot = None
        self._user_inputs = []

    def close(self, raise_ex=True):
        """
        Ends the test.

        :param raise_ex: If true, an exception will be raised for failed/new tests.
        :return: The test results.
        """
        if self.is_disabled:
            logger.debug('close(): ignored (disabled)')
            return
        try:
            logger.debug('close({})'.format(raise_ex))
            if not self._is_open:
                raise ValueError("Eyes not open")

            self._is_open = False

            self._reset_last_screenshot()

            # If there's no running session, we simply return the default test results.
            if not self._running_session:
                logger.debug('close(): Server session was not started')
                logger.info('close(): --- Empty test ended.')
                return TestResults()

            is_new_session = self._running_session['is_new_session']
            results_url = self._running_session['session_url']

            logger.info("close(): Ending server session...")
            should_save = (is_new_session and self.save_new_tests) or \
                          ((not is_new_session) and self.save_failed_tests)
            logger.debug("close(): automatically save session? %s" % should_save)
            results = self._agent_connector.stop_session(self._running_session, False, should_save)
            results.is_new = is_new_session
            results.url = results_url
            logger.info("close(): %s" % results)

            if results.status == TestResultsStatus.Unresolved:
                if results.is_new:
                    instructions = "Please approve the new baseline at " + results_url
                    logger.info("--- New test ended. " + instructions)
                    if raise_ex:
                        message = "'%s' of '%s'. %s" % (self._start_info['scenarioIdOrName'],
                                                        self._start_info['appIdOrName'],
                                                        instructions)
                        raise NewTestError(message, results)
                else:
                    logger.info("--- Failed test ended. See details at {}".format(results_url))
                    if raise_ex:
                        raise DiffsFoundError("Test '{}' of '{}' detected differences! See details at: {}".format(
                            self._start_info['scenarioIdOrName'],
                            self._start_info['appIdOrName'],
                            results_url), results)
            elif results.status == TestResultsStatus.Failed:
                logger.info("--- Failed test ended. See details at {}".format(results_url))
                if raise_ex:
                    raise TestFailedError("Test '{}' of '{}'. See details at: {}".format(
                        self._start_info['scenarioIdOrName'],
                        self._start_info['appIdOrName'],
                        results_url), results)
            # Test passed
            logger.info("--- Test passed. See details at {}".format(results_url))

            return results
        finally:
            self._running_session = None
            logger.close()

    def add_mouse_trigger_by_element(self, action, element):
        """
        Adds a mouse trigger.

        :param action: Mouse action (click, double click etc.)
        :param element: The element on which the action was performed.
        """
        if self.is_disabled:
            logger.debug("add_mouse_trigger: Ignoring %s (disabled)" % action)
            return
        # Triggers are activated on the last checked window.
        if self._last_screenshot is None:
            logger.debug("add_mouse_trigger: Ignoring %s (no screenshot)" % action)
            return
        if not EyesFrame.is_same_frame_chain(self._driver.get_frame_chain(),
                                             self._last_screenshot.get_frame_chain()):
            logger.debug("add_mouse_trigger: Ignoring %s (different frame)" % action)
            return
        control = self._last_screenshot.get_intersected_region_by_element(element)
        # Making sure the trigger is within the last screenshot bounds
        if control.is_empty():
            logger.debug("add_mouse_trigger: Ignoring %s (out of bounds)" % action)
            return
        cursor = control.middle_offset
        trigger = MouseTrigger(action, control, cursor)
        self._user_inputs.append(trigger)
        logger.info("add_mouse_trigger: Added %s" % trigger)

    def add_text_trigger_by_element(self, element, text):
        """
        Adds a text trigger.

        :param element: The element to which the text was sent.
        :param text: The trigger's text.
        """
        if self.is_disabled:
            logger.debug("add_text_trigger: Ignoring '%s' (disabled)" % text)
            return
        # Triggers are activated on the last checked window.
        if self._last_screenshot is None:
            logger.debug("add_text_trigger: Ignoring '%s' (no screenshot)" % text)
            return
        if not EyesFrame.is_same_frame_chain(self._driver.get_frame_chain(),
                                             self._last_screenshot.get_frame_chain()):
            logger.debug("add_text_trigger: Ignoring %s (different frame)" % text)
            return
        control = self._last_screenshot.get_intersected_region_by_element(element)
        # Making sure the trigger is within the last screenshot bounds
        if control.is_empty():
            logger.debug("add_text_trigger: Ignoring %s (out of bounds)" % text)
            return
        trigger = TextTrigger(control, text)
        self._user_inputs.append(trigger)
        logger.info("add_text_trigger: Added %s" % trigger)
