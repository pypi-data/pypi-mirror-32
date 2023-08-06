from collections import OrderedDict
from kooki.tools import write_file

from .metadata import Metadata, parse_metadata, data_merge
from .template import apply_template, apply_interpreter, check_template
from .grammar import check_grammar


def apply_custom(parts, custom):
    print('apply custom')
    result = []
    for index, part in enumerate(parts):
        result.append(custom(part, index))
    return result


def apply_merge(data):
    print('apply merge')
    result = None

    if isinstance(data, list):
        result = ''
        for content in data:
            result += content

    elif isinstance(data, OrderedDict):
        result = ''
        for key, content in OrderedDict(data).items():
            result += content

    else:
        raise Exception('merging bad data type {}'.format(type(data)))

    return result


def apply_split(data, separator):
    print('apply split')
    parts = data.split(separator)
    return parts


def export_to(name, extension, content):
    file_name = '{0}.{1}'.format(name, extension)
    print('export to {}'.format(file_name))
    write_file(file_name, content)


__all__ = [Metadata, apply_template, parse_metadata, apply_interpreter, data_merge, export_to]
