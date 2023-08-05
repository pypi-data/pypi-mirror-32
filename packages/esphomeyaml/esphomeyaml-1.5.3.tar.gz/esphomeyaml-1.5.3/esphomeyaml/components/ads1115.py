import voluptuous as vol

import esphomeyaml.config_validation as cv
from esphomeyaml.const import CONF_ADDRESS, CONF_ID, CONF_RATE
from esphomeyaml.helpers import App, Pvariable

DEPENDENCIES = ['i2c']

ADS1115_COMPONENT_CLASS = 'sensor::ADS1115Component'

RATE_REMOVE_MESSAGE = """The rate option has been removed in 1.5.0 and is no longer required."""

ADS1115_SCHEMA = vol.Schema({
    cv.GenerateID('ads1115'): cv.register_variable_id,
    vol.Required(CONF_ADDRESS): cv.i2c_address,

    vol.Optional(CONF_RATE): cv.invalid(RATE_REMOVE_MESSAGE)
})

CONFIG_SCHEMA = vol.All(cv.ensure_list, [ADS1115_SCHEMA])


def to_code(config):
    for conf in config:
        rhs = App.make_ads1115_component(conf[CONF_ADDRESS])
        Pvariable(ADS1115_COMPONENT_CLASS, conf[CONF_ID], rhs)


BUILD_FLAGS = '-DUSE_ADS1115_SENSOR'
