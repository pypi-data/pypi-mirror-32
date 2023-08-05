import voluptuous as vol

import esphomeyaml.config_validation as cv
from esphomeyaml.components import sensor
from esphomeyaml.components.sensor import MQTT_SENSOR_SCHEMA
from esphomeyaml.const import CONF_ADDRESS, CONF_HUMIDITY, CONF_ID, CONF_IIR_FILTER, CONF_NAME, \
    CONF_OVERSAMPLING, CONF_PRESSURE, CONF_TEMPERATURE, CONF_UPDATE_INTERVAL
from esphomeyaml.helpers import App, RawExpression, add, variable

DEPENDENCIES = ['i2c']

OVERSAMPLING_OPTIONS = {
    'NONE': 'sensor::BME280_OVERSAMPLING_NONE',
    '1X': 'sensor::BME280_OVERSAMPLING_1X',
    '2X': 'sensor::BME280_OVERSAMPLING_2X',
    '4X': 'sensor::BME280_OVERSAMPLING_4X',
    '8X': 'sensor::BME280_OVERSAMPLING_8X',
    '16X': 'sensor::BME280_OVERSAMPLING_16X',
}

IIR_FILTER_OPTIONS = {
    'OFF': 'sensor::BME280_IIR_FILTER_OFF',
    '2X': 'sensor::BME280_IIR_FILTER_2X',
    '4X': 'sensor::BME280_IIR_FILTER_4X',
    '8X': 'sensor::BME280_IIR_FILTER_8X',
    '16X': 'sensor::BME280_IIR_FILTER_16X',
}

BME280_OVERSAMPLING_SENSOR_SCHEMA = MQTT_SENSOR_SCHEMA.extend({
    vol.Optional(CONF_OVERSAMPLING): vol.All(vol.Upper, vol.Any(*OVERSAMPLING_OPTIONS)),
})

PLATFORM_SCHEMA = sensor.PLATFORM_SCHEMA.extend({
    cv.GenerateID('bme280'): cv.register_variable_id,
    vol.Optional(CONF_ADDRESS, default=0x77): cv.i2c_address,
    vol.Required(CONF_TEMPERATURE): BME280_OVERSAMPLING_SENSOR_SCHEMA,
    vol.Required(CONF_PRESSURE): BME280_OVERSAMPLING_SENSOR_SCHEMA,
    vol.Required(CONF_HUMIDITY): BME280_OVERSAMPLING_SENSOR_SCHEMA,
    vol.Optional(CONF_IIR_FILTER): vol.All(vol.Upper, vol.Any(*IIR_FILTER_OPTIONS)),
    vol.Optional(CONF_UPDATE_INTERVAL): cv.positive_time_period_milliseconds,
})


def to_code(config):
    rhs = App.make_bme280_sensor(config[CONF_TEMPERATURE][CONF_NAME],
                                 config[CONF_PRESSURE][CONF_NAME],
                                 config[CONF_HUMIDITY][CONF_NAME],
                                 config[CONF_ADDRESS],
                                 config.get(CONF_UPDATE_INTERVAL))
    make = variable('Application::MakeBME280Sensor', config[CONF_ID], rhs)
    bme280 = make.Pbme280
    if CONF_OVERSAMPLING in config[CONF_TEMPERATURE]:
        constant = OVERSAMPLING_OPTIONS[config[CONF_TEMPERATURE][CONF_OVERSAMPLING]]
        add(bme280.set_temperature_oversampling(RawExpression(constant)))
    if CONF_OVERSAMPLING in config[CONF_PRESSURE]:
        constant = OVERSAMPLING_OPTIONS[config[CONF_PRESSURE][CONF_OVERSAMPLING]]
        add(bme280.set_pressure_oversampling(RawExpression(constant)))
    if CONF_OVERSAMPLING in config[CONF_HUMIDITY]:
        constant = OVERSAMPLING_OPTIONS[config[CONF_HUMIDITY][CONF_OVERSAMPLING]]
        add(bme280.set_humidity_oversampling(RawExpression(constant)))
    if CONF_IIR_FILTER in config:
        constant = IIR_FILTER_OPTIONS[config[CONF_IIR_FILTER]]
        add(bme280.set_iir_filter(RawExpression(constant)))

    sensor.setup_sensor(bme280.Pget_temperature_sensor(), config[CONF_TEMPERATURE])
    sensor.setup_mqtt_sensor_component(make.Pmqtt_temperature, config[CONF_TEMPERATURE])

    sensor.setup_sensor(bme280.Pget_pressure_sensor(), config[CONF_PRESSURE])
    sensor.setup_mqtt_sensor_component(make.Pmqtt_pressure, config[CONF_PRESSURE])

    sensor.setup_sensor(bme280.Pget_humidity_sensor(), config[CONF_HUMIDITY])
    sensor.setup_mqtt_sensor_component(make.Pmqtt_humidity, config[CONF_HUMIDITY])


BUILD_FLAGS = '-DUSE_BME280'
