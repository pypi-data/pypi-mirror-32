import logging

import voluptuous as vol

import esphomeyaml.config_validation as cv
from esphomeyaml import core
from esphomeyaml.const import CONF_PORT, CONF_JS_URL, CONF_CSS_URL, CONF_ID, ESP_PLATFORM_ESP32
from esphomeyaml.helpers import App, add, Pvariable

_LOGGER = logging.getLogger(__name__)

CONFIG_SCHEMA = vol.Schema({
    cv.GenerateID('web_server'): cv.register_variable_id,
    vol.Optional(CONF_PORT): cv.port,
    vol.Optional(CONF_CSS_URL): vol.Url,
    vol.Optional(CONF_JS_URL): vol.Url,
})


def to_code(config):
    rhs = App.init_web_server(config.get(CONF_PORT))
    web_server = Pvariable('WebServer', config[CONF_ID], rhs)
    if CONF_CSS_URL in config:
        add(web_server.set_css_url(config[CONF_CSS_URL]))
    if CONF_JS_URL in config:
        add(web_server.set_js_url(config[CONF_JS_URL]))


BUILD_FLAGS = '-DUSE_WEB_SERVER'


def lib_deps(config):
    if core.ESP_PLATFORM == ESP_PLATFORM_ESP32:
        return 'FS'
    return ''
