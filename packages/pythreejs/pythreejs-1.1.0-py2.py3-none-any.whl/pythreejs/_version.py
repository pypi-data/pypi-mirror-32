version_info = (1, 1, 0, 'final')

_specifier_ = {'alpha': 'a', 'beta': 'b', 'candidate': 'rc', 'final': ''}

__version__ = '%s.%s.%s%s'%(version_info[0], version_info[1], version_info[2],
          '' if version_info[3]=='final' else _specifier_[version_info[3]]+str(version_info[4]))

# The version of the attribute spec that this package
# implements. This is the value used in
# _model_module_version/_view_module_version.
#
# Update this value when attributes are added/removed from
# the widget models, or if the serialized format changes.
EXTENSION_SPEC_VERSION = '~1.1.0'
