"""Functions to interact with a webdriver browser object."""
import time
import types
import traceback

from selenium import webdriver
from selenium.webdriver.remote.webelement import WebElement
from selenium.webdriver.common.desired_capabilities import DesiredCapabilities
from selenium.webdriver.chrome.options import Options
from selenium.common.exceptions import WebDriverException

from golem.core.exceptions import IncorrectSelectorType, ElementNotFound
from golem.core import utils
from golem import execution


def _find_webelement(root, selector_type, selector_value,
                     element_name, remaining_time):
    """Finds a web element."""
    webelement = None
    start_time = time.time()
    try:
        if selector_type == 'id':
            webelement = root.find_element_by_id(selector_value)
        elif selector_type == 'css':
            webelement = root.find_element_by_css_selector(selector_value)
        elif selector_type == 'text':
            webelement = root.find_element_by_css_selector("text[{}]".format(selector_value))
        elif selector_type == 'link_text':
            webelement = root.find_element_by_link_text(selector_value)
        elif selector_type == 'partial_link_text':
            webelement = root.find_element_by_partial_link_text(selector_value)
        elif selector_type == 'name':
            webelement = root.find_element_by_name(selector_value)
        elif selector_type == 'xpath':
            webelement = root.find_element_by_xpath(selector_value)
        elif selector_type == 'tag_name':
            webelement = root.find_element_by_tag_name(selector_value)
        else:
            msg = 'Selector {0} is not a valid option'.format(selector_type)
            raise IncorrectSelectorType(msg)
        execution.logger.debug('Element found')
    except:
        time.sleep(0.5)
        end_time = time.time()
        remaining_time -= end_time - start_time
        if remaining_time > 0:
            execution.logger.debug('Element not found yet, remaining time: {}'.format(remaining_time))
            webelement = _find_webelement(root, selector_type, selector_value,
                                          element_name, remaining_time)
        else:
            raise ElementNotFound('Element {0} not found using selector {1}:\'{2}\''
                                  .format(element_name, selector_type, selector_value))
    
    # Use remaining time to wait until element is visible (is_displayed)
    # TODO add this as a setting
    remaining_time = remaining_time - (time.time() - start_time)
    while not webelement.is_displayed() and remaining_time > 0:
        # Element is not visible yet
        execution.logger.debug('Element still not visible, waiting')
        time.sleep(0.5)
        remaining_time = remaining_time - (time.time() - start_time)

    if not webelement.is_displayed():
        execution.logger.debug('Element not visible, continuing..')
        
    return webelement


def _find(self, element=None, id=None, name=None, text=None,
          link_text=None, partial_link_text=None, css=None,
          xpath=None, tag_name=None, timeout=None):
    """Find a webelement.

    This function is added as a method to Webdriver and Webelement objects.

    `element` can be:
      a web element
      a tuple with the format (<selector_type>, <selector_value>, <display_nane>)
      a string: css selector is used
    """
    webelement = None

    selector_type = None
    selector_value = None
    element_name = None

    if timeout is None:
        timeout = execution.settings['implicit_wait']

    if isinstance(element, WebElement):
        webelement = element
    elif isinstance(element, tuple):
        selector_type = element[0]
        selector_value = element[1]
        element_name = element[2] if len(element) == 3 else element[1]
    elif isinstance(element, str):
        selector_type = 'css'
        selector_value = element
        element_name = element
    elif id:
        selector_type = 'id'
        selector_value = element_name = id
    elif name:
        selector_type = 'name'
        selector_value = element_name = name
    elif text:
        selector_type = 'text'
        selector_value = element_name = text
    elif link_text:
        selector_type = 'link_text'
        selector_value = element_name = link_text
    elif partial_link_text:
        selector_type = 'partial_link_text'
        selector_value = element_name = partial_link_text
    elif css:
        selector_type = 'css'
        selector_value = element_name = css
    elif xpath:
        selector_type = 'xpath'
        selector_value = element_name = xpath
    elif tag_name:
        selector_type = 'tag_name'
        selector_value = element_name = tag_name
    else:
        raise IncorrectSelectorType('Selector is not a valid option')
    
    if not webelement:
        webelement = _find_webelement(self, selector_type, selector_value,
                                      element_name, timeout)
        webelement.selector_type = selector_type
        webelement.selector_value = selector_value
        webelement.name = element_name
        # add find and find_all as methods to the webelement object
        webelement.find = types.MethodType(_find, webelement)
        webelement.find_all = types.MethodType(_find_all, webelement)
    return webelement


def _find_all(self, element=None, id=None, name=None, text=None, link_text=None,
              partial_link_text=None, css=None, xpath=None, tag_name=None):
    """Find all webelements."""
    webelements = []
    selector_type = None
    selector_value = None
    element_name = None
    if isinstance(element, tuple):
        selector_type = element[0]
        selector_value = element_name = element[1]
        if selector_type == 'id':
            id = selector_value
        elif selector_type == 'css':
            css = selector_value
        elif selector_type == 'text':
            text = selector_value
        elif selector_type == 'link_text':
            link_text = selector_value
        elif selector_type == 'partial_link_text':
            partial_link_text = selector_value
        elif selector_type == 'name':
            name = selector_value
        elif selector_type == 'xpath':
            xpath = selector_value
        elif selector_type == 'tag_name':
            tag_name = selector_value
        else:
            raise Exception('Incorrect element {}'.format(element))
    elif isinstance(element, str):
        css = element
        selector_type = 'css'
        selector_value = element_name = element
    
    if id:
        selector_type = 'id'
        selector_value = element_name = id
        webelements = self.find_elements_by_id(id)
    elif css:
        selector_type = 'css'
        selector_value = element_name = css
        webelements = self.find_elements_by_css_selector(css)
    elif text:
        selector_type = 'text'
        selector_value = element_name = text
        webelements = self.find_elements_by_css_selector("text[{}]".format(text))
    elif link_text:
        selector_type = 'link_text'
        selector_value = element_name = link_text
        webelements = self.find_elements_by_link_text(link_text)
    elif partial_link_text:
        selector_type = 'partial_link_text'
        selector_value = element_name = partial_link_text
        webelements = self.find_elements_by_partial_link_text(partial_link_text)
    elif name:
        selector_type = 'name'
        selector_value = element_name = name
        webelements = self.find_elements_by_name(name)
    elif xpath:
        selector_type = 'xpath'
        selector_value = element_name = xpath
        webelements = self.find_elements_by_xpath(xpath)
    elif tag_name:
        selector_type = 'tag_name'
        selector_value = element_name = tag_name
        webelements = self.find_elements_by_tag_name(tag_name)
    else:
        raise IncorrectSelectorType('Incorrect selector provided')

    for elem in webelements:
        elem.selector_type = selector_type
        elem.selector_value = selector_value
        elem.name = element_name
        # add find and find_all methods to the webelement object
        elem.find = types.MethodType(_find, elem)
        elem.find_all = types.MethodType(_find_all, elem)

    return webelements


def element(*args, **kwargs):
    """Shortcut to golem.browser.get_browser().find()"""
    if len(args) == 1:
        kwargs['element'] = args[0]
    webelement = get_browser().find(**kwargs)
    return webelement


def elements(*args, **kwargs):
    """Shortcut to golem.browser.get_browser().find_all()"""
    if len(args) == 1:
        kwargs['element'] = args[0]
    webelement = get_browser().find_all(**kwargs)
    return webelement


def open_browser(browser_id=None):
    """Open a browser.

    When opening more than one browser instance per test
    provide a browser_id to switch between browsers later on
    """
    msg_could_not_start_driver = ('Could not start {} driver using the path \'{}\', '
                                  'check the settings file.')
    driver = None
    browser_definition = execution.browser_definition
    settings = execution.settings
    if browser_definition['remote'] is True:
        driver = webdriver.Remote(command_executor=settings['remote_url'],
                                  desired_capabilities=browser_definition['capabilities'])
    elif browser_definition['name'] == 'firefox':
        if settings['geckodriver_path']:
            executable_path = utils.match_latest_executable_path(settings['geckodriver_path'])
            try:
                driver = webdriver.Firefox(executable_path=executable_path)
            except:
                msg = msg_could_not_start_driver.format('firefox', settings['geckodriver_path'])
                execution.logger.error(msg)
                execution.logger.info(traceback.format_exc())
                raise Exception(msg)
        else:
            raise Exception('geckodriver_path setting is not defined')
    elif browser_definition['name'] == 'chrome':
        if settings['chromedriver_path']:
            try:
                executable_path = utils.match_latest_executable_path(settings['chromedriver_path'])
                chrome_options = Options()
                chrome_options.add_argument('--start-maximized')
                driver = webdriver.Chrome(executable_path=executable_path,
                                          chrome_options=chrome_options)
            except:
                msg = msg_could_not_start_driver.format('chrome',
                                                        settings['chromedriver_path'])
                execution.logger.error(msg)
                execution.logger.info(traceback.format_exc())
                raise Exception(msg)
        else:
            raise Exception('chromedriver_path setting is not defined')
    elif browser_definition['name'] == 'chrome-headless':
        if settings['chromedriver_path']:
            try:
                executable_path = utils.match_latest_executable_path(settings['chromedriver_path'])
                options = webdriver.ChromeOptions()
                options.add_argument('headless')
                options.add_argument('--window-size=1600,1600')
                driver = webdriver.Chrome(executable_path=executable_path,
                                          chrome_options=options)
            except:
                msg = msg_could_not_start_driver.format('chrome',
                                                        settings['chromedriver_path'])
                execution.logger.error(msg)
                execution.logger.info(traceback.format_exc())
                raise Exception(msg)
        else:
            raise Exception('chromedriver_path setting is not defined')
    elif browser_definition['name'] == 'ie':
        if settings['iedriver_path']:
            try:
                executable_path = utils.match_latest_executable_path(settings['iedriver_path'])
                driver = webdriver.Ie(executable_path=executable_path)
            except:
                msg = msg_could_not_start_driver.format('internet explorer',
                                                        settings['iedriver_path'])
                execution.logger.error(msg)
                execution.logger.info(traceback.format_exc())
                raise Exception(msg)
        else:
            raise Exception('iedriver_path setting is not defined')
    elif browser_definition['name'] == 'ie-remote':
        driver = webdriver.Remote(command_executor=settings['remote_url'],
                                  desired_capabilities=DesiredCapabilities.IE)
    elif browser_definition['name'] == 'chrome-remote-headless':
        options = webdriver.ChromeOptions()
        options.add_argument('headless')
        desired_capabilities = options.to_capabilities()
        driver = webdriver.Chrome(command_executor=settings['remote_url'],
                                  desired_capabilities=desired_capabilities)
    elif browser_definition['name'] == 'chrome-remote':
        driver = webdriver.Remote(command_executor=settings['remote_url'],
                                  desired_capabilities=DesiredCapabilities.CHROME)
    elif browser_definition['name'] == 'firefox-remote':
        driver = webdriver.Remote(command_executor=settings['remote_url'],
                                  desired_capabilities=DesiredCapabilities.FIREFOX)
    else:
        raise Exception('Error: {} is not a valid driver'.format(browser_definition['name']))

    driver.maximize_window()

    # bind _find and _find_all methods to driver instance
    driver.find = types.MethodType(_find, driver)
    driver.find_all = types.MethodType(_find_all, driver)

    if not browser_id:
        if len(execution.browsers) == 0:
            execution.browsers['main'] = driver
        else:
            browser_id = 'browser{}'.format(len(execution.browsers)+1)
            execution.browsers[browser_id] = driver
    else:
        execution.browsers[browser_id] = driver
    if not execution.browser:
        execution.browser = driver


def get_browser():
    """Returns the active browser. Starts a new one if there is none."""
    if not execution.browser:
        open_browser()
    return execution.browser


def activate_browser(browser_id):
    """Activate a browser.
    Only needed when the test starts more than one browser instance."""
    if not browser_id in execution.browsers:
        # TODO, use error() function
        raise Exception("Error: {} is not a valid browser id. Current browsers "
                        "are: {}".format(browser_id, ', '.join(execution.browsers.keys())))
    else:
        execution.browser = execution.browsers[browser_id]
