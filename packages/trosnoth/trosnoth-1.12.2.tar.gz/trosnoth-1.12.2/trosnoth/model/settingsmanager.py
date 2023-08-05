import json
import logging

log = logging.getLogger(__name__)


class SettingsSynchroniser(object):
    RESERVED_ATTRIBUTES = {'set', 'applyMessage', 'dumpState', 'restoreState'}

    def __init__(self, dispatcher, attributes):
        self._dispatch = dispatcher
        self._attributes = set(attributes.keys())

        for k, v in attributes.items():
            if k.startswith('_') or k in self.RESERVED_ATTRIBUTES:
                raise KeyError('{} is not a valid attribute name'.format(k))
            setattr(self, k, v)

    def dumpState(self):
        return {
            attrName: getattr(self, attrName)
            for attrName in self._attributes
        }

    def restoreState(self, data):
        for attrName in self._attributes:
            setattr(self, attrName, data[attrName])

    def set(self, **kwargs):
        for k, v in kwargs.items():
            if k not in self._attributes:
                raise KeyError('Unknown attribute: {!r}'.format(k))
        self._dispatch(json.dumps(kwargs))

    def applyMessage(self, data):
        for k, v in json.loads(data).items():
            if k not in self._attributes:
                log.error('%s: received unknown attribute %s', self, k)
            else:
                setattr(self, k, v)