# coding=utf-8
"""Helpers for config validation using voluptuous."""
from __future__ import print_function

import logging
import re

import voluptuous as vol

from esphomeyaml import core
from esphomeyaml.const import CONF_AVAILABILITY, CONF_COMMAND_TOPIC, CONF_DISCOVERY, CONF_ID, \
    CONF_NAME, CONF_PAYLOAD_AVAILABLE, \
    CONF_PAYLOAD_NOT_AVAILABLE, CONF_PLATFORM, CONF_RETAIN, CONF_STATE_TOPIC, CONF_TOPIC, \
    ESP_PLATFORM_ESP32, ESP_PLATFORM_ESP8266
from esphomeyaml.core import HexInt, IPAddress, TimePeriod, TimePeriodMilliseconds, \
    TimePeriodMicroseconds, TimePeriodSeconds
from esphomeyaml.helpers import ensure_unique_string

_LOGGER = logging.getLogger(__name__)

# pylint: disable=invalid-name

port = vol.All(vol.Coerce(int), vol.Range(min=1, max=65535))
positive_float = vol.All(vol.Coerce(float), vol.Range(min=0))
zero_to_one_float = vol.All(vol.Coerce(float), vol.Range(min=0, max=1))
positive_int = vol.All(vol.Coerce(int), vol.Range(min=0))
positive_not_null_int = vol.All(vol.Coerce(int), vol.Range(min=0, min_included=False))

ALLOWED_NAME_CHARS = u'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_'

RESERVED_IDS = [
    # C++ keywords http://en.cppreference.com/w/cpp/keyword
    'alignas', 'alignof', 'and', 'and_eq', 'asm', 'auto', 'bitand', 'bitor', 'bool', 'break',
    'case', 'catch', 'char', 'char16_t', 'char32_t', 'class', 'compl', 'concept', 'const',
    'constexpr', 'const_cast', 'continue', 'decltype', 'default', 'delete', 'do', 'double',
    'dynamic_cast', 'else', 'enum', 'explicit', 'export', 'export', 'extern', 'false', 'float',
    'for', 'friend', 'goto', 'if', 'inline', 'int', 'long', 'mutable', 'namespace', 'new',
    'noexcept', 'not', 'not_eq', 'nullptr', 'operator', 'or', 'or_eq', 'private', 'protected',
    'public', 'register', 'reinterpret_cast', 'requires', 'return', 'short', 'signed', 'sizeof',
    'static', 'static_assert', 'static_cast', 'struct', 'switch', 'template', 'this',
    'thread_local', 'throw', 'true', 'try', 'typedef', 'typeid', 'typename', 'union', 'unsigned',
    'using', 'virtual', 'void', 'volatile', 'wchar_t', 'while', 'xor', 'xor_eq',

    'App', 'pinMode', 'delay', 'delayMicroseconds', 'digitalRead', 'digitalWrite', 'INPUT',
    'OUTPUT',
]


def alphanumeric(value):
    if value is None:
        raise vol.Invalid("string value is None")
    value = unicode(value)
    if not value.isalnum():
        raise vol.Invalid("string value is not alphanumeric")
    return value


def valid_name(value):
    value = string_strict(value)
    if not all(c in ALLOWED_NAME_CHARS for c in value):
        raise vol.Invalid(u"Valid characters for name are {}".format(ALLOWED_NAME_CHARS))
    return value


def string(value):
    if isinstance(value, (dict, list)):
        raise vol.Invalid("string value cannot be dictionary or list.")
    if value is not None:
        return unicode(value)
    raise vol.Invalid("string value is None")


def string_strict(value):
    """Strictly only allow strings."""
    if isinstance(value, (str, unicode)):
        return value
    raise vol.Invalid("Must be string, did you forget putting quotes "
                      "around the value?")


def icon(value):
    """Validate icon."""
    value = string_strict(value)
    if value.startswith('mdi:'):
        return value
    raise vol.Invalid('Icons should start with prefix "mdi:"')


def boolean(value):
    """Validate and coerce a boolean value."""
    if isinstance(value, str):
        value = value.lower()
        if value in ('1', 'true', 'yes', 'on', 'enable'):
            return True
        if value in ('0', 'false', 'no', 'off', 'disable'):
            return False
        raise vol.Invalid('invalid boolean value {}'.format(value))
    return bool(value)


def ensure_list(value):
    """Wrap value in list if it is not one."""
    if value is None:
        return []
    if isinstance(value, list):
        return value
    return [value]


def ensure_dict(value):
    if value is None:
        return {}
    if not isinstance(value, dict):
        raise vol.Invalid("Expected a dictionary")
    return value


def hex_int_(value):
    if isinstance(value, (int, long)):
        return HexInt(value)
    value = string_strict(value).lower()
    if value.startswith('0x'):
        return HexInt(int(value, 16))
    return HexInt(int(value))


def int_(value):
    if isinstance(value, (int, long)):
        return value
    value = string_strict(value).lower()
    if value.startswith('0x'):
        return int(value, 16)
    return int(value)


hex_int = vol.Coerce(hex_int_)
match_cpp_var_ = vol.Match(r'^[a-zA-Z_][a-zA-Z0-9_]+$', msg=u"Must be a valid C++ variable name")


def variable_id(value):
    value = match_cpp_var_(value)
    if value in RESERVED_IDS:
        raise vol.Invalid(u"ID {} is reserved internally and cannot be used".format(value))
    return value


def only_on(platforms):
    if not isinstance(platforms, list):
        platforms = [platforms]

    def validator_(obj):
        if core.ESP_PLATFORM not in platforms:
            raise vol.Invalid(u"This feature is only available on {}".format(platforms))
        return obj

    return validator_


only_on_esp32 = only_on(ESP_PLATFORM_ESP32)
only_on_esp8266 = only_on(ESP_PLATFORM_ESP8266)


# Adapted from:
# https://github.com/alecthomas/voluptuous/issues/115#issuecomment-144464666
def has_at_least_one_key(*keys):
    """Validate that at least one key exists."""

    def validate(obj):
        """Test keys exist in dict."""
        if not isinstance(obj, dict):
            raise vol.Invalid('expected dictionary')

        for k in obj.keys():
            if k in keys:
                return obj
        raise vol.Invalid('must contain one of {}.'.format(', '.join(keys)))

    return validate


TIME_PERIOD_ERROR = "Time period {} should be format number + unit, for example 5ms, 5s, 5min, 5h"

time_period_dict = vol.All(
    dict, vol.Schema({
        'days': vol.Coerce(float),
        'hours': vol.Coerce(float),
        'minutes': vol.Coerce(float),
        'seconds': vol.Coerce(float),
        'milliseconds': vol.Coerce(float),
        'microseconds': vol.Coerce(float),
    }),
    has_at_least_one_key('days', 'hours', 'minutes',
                         'seconds', 'milliseconds', 'microseconds'),
    lambda value: TimePeriod(**value))


TIME_PERIOD_EXPLICIT_MESSAGE = ("The old way of being able to write time values without a "
                                "time unit (like \"1000\" for 1000 milliseconds) has been "
                                "removed in 1.5.0 as it was ambiguous in some places. Please "
                                "now explicitly specify the time unit (like \"1000ms\"). See "
                                "https://esphomelib.com/esphomeyaml/configuration-types.html#time "
                                "for more information.")


def time_period_str_colon(value):
    """Validate and transform time offset with format HH:MM[:SS]."""
    if isinstance(value, int):
        raise vol.Invalid('Make sure you wrap time values in quotes')
    elif not isinstance(value, str):
        raise vol.Invalid(TIME_PERIOD_ERROR.format(value))

    negative_offset = False
    if value.startswith('-'):
        negative_offset = True
        value = value[1:]
    elif value.startswith('+'):
        value = value[1:]

    try:
        parsed = [int(x) for x in value.split(':')]
    except ValueError:
        raise vol.Invalid(TIME_PERIOD_ERROR.format(value))

    if len(parsed) == 2:
        hour, minute = parsed
        second = 0
    elif len(parsed) == 3:
        hour, minute, second = parsed
    else:
        raise vol.Invalid(TIME_PERIOD_ERROR.format(value))

    offset = TimePeriod(hours=hour, minutes=minute, seconds=second)

    if negative_offset:
        offset *= -1

    return offset


def time_period_str_unit(value):
    """Validate and transform time period with time unit and integer value."""
    if isinstance(value, int):
        value = str(value)
    elif not isinstance(value, (str, unicode)):
        raise vol.Invalid("Expected string for time period with unit.")

    try:
        float(value)
    except ValueError:
        pass
    else:
        raise vol.Invalid(TIME_PERIOD_EXPLICIT_MESSAGE)

    unit_to_kwarg = {
        'us': 'microseconds',
        'microseconds': 'microseconds',
        'ms': 'milliseconds',
        'milliseconds': 'milliseconds',
        's': 'seconds',
        'sec': 'seconds',
        'seconds': 'seconds',
        'min': 'minutes',
        'minutes': 'minutes',
        'h': 'hours',
        'hours': 'hours',
        'd': 'days',
        'days': 'days',
    }

    match = re.match(r"^([-+]?[0-9]*\.?[0-9]*)\s*(\w*)$", value)

    if match is None or match.group(2) not in unit_to_kwarg:
        raise vol.Invalid(u"Expected time period with unit, "
                          u"got {}".format(value))

    kwarg = unit_to_kwarg[match.group(2)]
    return TimePeriod(**{kwarg: float(match.group(1))})


def time_period_in_milliseconds(value):
    if value.microseconds is not None and value.microseconds != 0:
        raise vol.Invalid("Maximum precision is milliseconds")
    return TimePeriodMilliseconds(**value.as_dict())


def time_period_in_microseconds(value):
    return TimePeriodMicroseconds(**value.as_dict())


def time_period_in_seconds(value):
    if value.microseconds is not None and value.microseconds != 0:
        raise vol.Invalid("Maximum precision is seconds")
    if value.milliseconds is not None and value.milliseconds != 0:
        raise vol.Invalid("Maximum precision is seconds")
    return TimePeriodSeconds(**value.as_dict())


time_period = vol.Any(time_period_str_unit, time_period_str_colon, time_period_dict)
positive_time_period = vol.All(time_period, vol.Range(min=TimePeriod()))
positive_time_period_milliseconds = vol.All(positive_time_period, time_period_in_milliseconds)
positive_time_period_seconds = vol.All(positive_time_period, time_period_in_seconds)
positive_time_period_microseconds = vol.All(positive_time_period, time_period_in_microseconds)
positive_not_null_time_period = vol.All(time_period,
                                        vol.Range(min=TimePeriod(), min_included=False))


METRIC_SUFFIXES = {
    'E': 1e18, 'P': 1e15, 'T': 1e12, 'G': 1e9, 'M': 1e6, 'k': 1e3, 'da': 10, 'd': 1e-1,
    'c': 1e-2, 'm': 0.001, u'µ': 1e-6, 'u': 1e-6, 'n': 1e-9, 'p': 1e-12, 'f': 1e-15, 'a': 1e-18,
    '': 1
}


def frequency(value):
    value = string(value)
    match = re.match(r"^([-+]?[0-9]*\.?[0-9]*)\s*(\w*?)(?:Hz|HZ|hz)?$", value)

    if match is None:
        raise vol.Invalid(u"Expected frequency with unit, "
                          u"got {}".format(value))

    mantissa = float(match.group(1))
    if match.group(2) not in METRIC_SUFFIXES:
        raise vol.Invalid(u"Invalid frequency suffix {}".format(match.group(2)))

    multiplier = METRIC_SUFFIXES[match.group(2)]
    return mantissa * multiplier


def hostname(value):
    value = string(value)
    if len(value) > 63:
        raise vol.Invalid("Hostnames can only be 63 characters long")
    for c in value:
        if not (c.isalnum() or c in '_-'):
            raise vol.Invalid("Hostname can only have alphanumeric characters and _ or -")
    return value


def ssid(value):
    if value is None:
        raise vol.Invalid("SSID can not be None")
    if not isinstance(value, str):
        raise vol.Invalid("SSID must be a string. Did you wrap it in quotes?")
    if not value:
        raise vol.Invalid("SSID can't be empty.")
    if len(value) > 31:
        raise vol.Invalid("SSID can't be longer than 31 characters")
    return value


def ipv4(value):
    if isinstance(value, list):
        parts = value
    elif isinstance(value, str):
        parts = value.split('.')
    else:
        raise vol.Invalid("IPv4 address must consist of either string or "
                          "integer list")
    if len(parts) != 4:
        raise vol.Invalid("IPv4 address must consist of four point-separated "
                          "integers")
    parts_ = list(map(int, parts))
    if not all(0 <= x < 256 for x in parts_):
        raise vol.Invalid("IPv4 address parts must be in range from 0 to 255")
    return IPAddress(*parts_)


def _valid_topic(value):
    """Validate that this is a valid topic name/filter."""
    if isinstance(value, dict):
        raise vol.Invalid("Can't use dictionary with topic")
    value = string(value)
    try:
        raw_value = value.encode('utf-8')
    except UnicodeError:
        raise vol.Invalid("MQTT topic name/filter must be valid UTF-8 string.")
    if not raw_value:
        raise vol.Invalid("MQTT topic name/filter must not be empty.")
    if len(raw_value) > 65535:
        raise vol.Invalid("MQTT topic name/filter must not be longer than "
                          "65535 encoded bytes.")
    if '\0' in value:
        raise vol.Invalid("MQTT topic name/filter must not contain null "
                          "character.")
    return value


def subscribe_topic(value):
    """Validate that we can subscribe using this MQTT topic."""
    value = _valid_topic(value)
    for i in (i for i, c in enumerate(value) if c == '+'):
        if (i > 0 and value[i - 1] != '/') or \
                (i < len(value) - 1 and value[i + 1] != '/'):
            raise vol.Invalid("Single-level wildcard must occupy an entire "
                              "level of the filter")

    index = value.find('#')
    if index != -1:
        if index != len(value) - 1:
            # If there are multiple wildcards, this will also trigger
            raise vol.Invalid("Multi-level wildcard must be the last "
                              "character in the topic filter.")
        if len(value) > 1 and value[index - 1] != '/':
            raise vol.Invalid("Multi-level wildcard must be after a topic "
                              "level separator.")

    return value


def publish_topic(value):
    """Validate that we can publish using this MQTT topic."""
    value = _valid_topic(value)
    if '+' in value or '#' in value:
        raise vol.Invalid("Wildcards can not be used in topic names")
    return value


def mqtt_payload(value):
    if value is None:
        return ''
    return string(value)


uint8_t = vol.All(int_, vol.Range(min=0, max=255))
uint16_t = vol.All(int_, vol.Range(min=0, max=65535))
uint32_t = vol.All(int_, vol.Range(min=0, max=4294967295))
hex_uint8_t = vol.All(hex_int, vol.Range(min=0, max=255))
hex_uint16_t = vol.All(hex_int, vol.Range(min=0, max=65535))
hex_uint32_t = vol.All(hex_int, vol.Range(min=0, max=4294967295))
i2c_address = hex_uint8_t


def invalid(message):
    def validator(value):
        raise vol.Invalid(message)
    return validator


def valid(value):
    return value


REGISTERED_IDS = set()


def register_variable_id(value):
    s = variable_id(value)
    if s in REGISTERED_IDS:
        raise vol.Invalid("This ID has already been used")
    REGISTERED_IDS.add(s)
    return s


class GenerateID(vol.Optional):
    def __init__(self, basename, key=CONF_ID):
        self._basename = basename
        super(GenerateID, self).__init__(key, default=self.default_variable_id)

    def default_variable_id(self):
        return ensure_unique_string(self._basename, REGISTERED_IDS)


REQUIRED_ID_SCHEMA = vol.Schema({
    vol.Required(CONF_ID): register_variable_id,
})

PLATFORM_SCHEMA = vol.Schema({
    vol.Required(CONF_PLATFORM): valid,
})

MQTT_COMPONENT_AVAILABILITY_SCHEMA = vol.Schema({
    vol.Required(CONF_TOPIC): subscribe_topic,
    vol.Optional(CONF_PAYLOAD_AVAILABLE, default='online'): mqtt_payload,
    vol.Optional(CONF_PAYLOAD_NOT_AVAILABLE, default='offline'): mqtt_payload,
})

MQTT_COMPONENT_SCHEMA = vol.Schema({
    vol.Required(CONF_NAME): string,
    vol.Optional(CONF_RETAIN): boolean,
    vol.Optional(CONF_DISCOVERY): boolean,
    vol.Optional(CONF_STATE_TOPIC): publish_topic,
    vol.Optional(CONF_AVAILABILITY): MQTT_COMPONENT_AVAILABILITY_SCHEMA,
})

MQTT_COMMAND_COMPONENT_SCHEMA = MQTT_COMPONENT_SCHEMA.extend({
    vol.Optional(CONF_COMMAND_TOPIC): subscribe_topic,
})
