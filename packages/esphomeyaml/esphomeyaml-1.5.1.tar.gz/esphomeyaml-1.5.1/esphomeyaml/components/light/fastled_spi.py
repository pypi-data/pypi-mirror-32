import voluptuous as vol

import esphomeyaml.config_validation as cv
from esphomeyaml import pins
from esphomeyaml.components import light
from esphomeyaml.const import CONF_CHIPSET, CONF_CLOCK_PIN, CONF_DATA_PIN, \
    CONF_DEFAULT_TRANSITION_LENGTH, CONF_GAMMA_CORRECT, CONF_ID, CONF_MAX_REFRESH_RATE, CONF_NAME, \
    CONF_NUM_LEDS, CONF_RGB_ORDER
from esphomeyaml.helpers import App, TemplateArguments, add, setup_mqtt_component, variable, \
    RawExpression

CHIPSETS = [
    'LPD8806',
    'WS2801',
    'WS2803',
    'SM16716',
    'P9813',
    'APA102',
    'SK9822',
    'DOTSTAR',
]

RGB_ORDERS = [
    'RGB',
    'RBG',
    'GRB',
    'GBR',
    'BRG',
    'BGR',
]

PLATFORM_SCHEMA = light.PLATFORM_SCHEMA.extend({
    cv.GenerateID('fast_led_spi_light'): cv.register_variable_id,

    vol.Required(CONF_CHIPSET): vol.All(vol.Upper, vol.Any(*CHIPSETS)),
    vol.Required(CONF_DATA_PIN): pins.output_pin,
    vol.Required(CONF_CLOCK_PIN): pins.output_pin,

    vol.Required(CONF_NUM_LEDS): cv.positive_not_null_int,
    vol.Optional(CONF_RGB_ORDER): vol.All(vol.Upper, vol.Any(*RGB_ORDERS)),
    vol.Optional(CONF_MAX_REFRESH_RATE): cv.positive_time_period_microseconds,

    vol.Optional(CONF_GAMMA_CORRECT): cv.positive_float,
    vol.Optional(CONF_DEFAULT_TRANSITION_LENGTH): cv.positive_time_period_milliseconds,
})


def to_code(config):
    rhs = App.make_fast_led_light(config[CONF_NAME])
    make = variable('Application::MakeFastLEDLight', config[CONF_ID], rhs)
    fast_led = make.Pfast_led

    rgb_order = None
    if CONF_RGB_ORDER in config:
        rgb_order = RawExpression(config[CONF_RGB_ORDER])
    template_args = TemplateArguments(RawExpression(config[CONF_CHIPSET]),
                                      config[CONF_DATA_PIN],
                                      config[CONF_CLOCK_PIN],
                                      rgb_order)
    add(fast_led.add_leds(template_args, config[CONF_NUM_LEDS]))

    if CONF_MAX_REFRESH_RATE in config:
        add(fast_led.set_max_refresh_rate(config[CONF_MAX_REFRESH_RATE]))

    setup_mqtt_component(make.Pmqtt, config)
    light.setup_light_component(make.Pstate, config)


BUILD_FLAGS = '-DUSE_FAST_LED_LIGHT'
