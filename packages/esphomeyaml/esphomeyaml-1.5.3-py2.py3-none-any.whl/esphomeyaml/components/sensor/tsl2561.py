import voluptuous as vol

import esphomeyaml.config_validation as cv
from esphomeyaml.components import sensor
from esphomeyaml.const import CONF_ADDRESS, CONF_GAIN, CONF_ID, CONF_INTEGRATION_TIME, CONF_NAME, \
    CONF_UPDATE_INTERVAL
from esphomeyaml.helpers import App, RawExpression, add, variable

DEPENDENCIES = ['i2c']

INTEGRATION_TIMES = {
    14: 'sensor::TSL2561_INTEGRATION_14MS',
    101: 'sensor::TSL2561_INTEGRATION_101MS',
    402: 'sensor::TSL2561_INTEGRATION_402MS',
}
GAINS = {
    '1X': 'sensor::TSL2561_GAIN_1X',
    '16X': 'sensor::TSL2561_GAIN_16X',
}

CONF_IS_CS_PACKAGE = 'is_cs_package'


def validate_integration_time(value):
    value = cv.positive_time_period_milliseconds(value).total_milliseconds
    if value not in INTEGRATION_TIMES:
        raise vol.Invalid(u"Unsupported integration time {}.".format(value))
    return value


PLATFORM_SCHEMA = sensor.PLATFORM_SCHEMA.extend({
    cv.GenerateID('tsl2561_sensor'): cv.register_variable_id,
    vol.Optional(CONF_ADDRESS, default=0x39): cv.i2c_address,
    vol.Optional(CONF_INTEGRATION_TIME): validate_integration_time,
    vol.Optional(CONF_GAIN): vol.All(vol.Upper, vol.Any(*GAINS)),
    vol.Optional(CONF_IS_CS_PACKAGE): cv.boolean,
    vol.Optional(CONF_UPDATE_INTERVAL): cv.positive_time_period_milliseconds,
}).extend(sensor.MQTT_SENSOR_SCHEMA.schema)


def to_code(config):
    rhs = App.make_tsl2561_sensor(config[CONF_NAME], config[CONF_ADDRESS],
                                  config.get(CONF_UPDATE_INTERVAL))
    make_tsl = variable('Application::MakeTSL2561Sensor', config[CONF_ID], rhs)
    tsl2561 = make_tsl.Ptsl2561
    if CONF_INTEGRATION_TIME in config:
        constant = INTEGRATION_TIMES[config[CONF_INTEGRATION_TIME]]
        add(tsl2561.set_integration_time(RawExpression(constant)))
    if CONF_GAIN in config:
        constant = GAINS[config[CONF_GAIN]]
        add(tsl2561.set_gain(RawExpression(constant)))
    if CONF_IS_CS_PACKAGE in config:
        add(tsl2561.set_is_cs_package(config[CONF_IS_CS_PACKAGE]))
    sensor.setup_sensor(tsl2561, config)
    sensor.setup_mqtt_sensor_component(make_tsl.Pmqtt, config)


BUILD_FLAGS = '-DUSE_TSL2561'
