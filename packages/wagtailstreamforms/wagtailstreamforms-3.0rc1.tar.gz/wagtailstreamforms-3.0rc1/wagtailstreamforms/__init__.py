from wagtailstreamforms.versioning import get_version

# major.minor.patch.release.number
# release must be one of alpha, beta, rc, or final
VERSION = (3, 0, 0, 'rc', 1)

__version__ = get_version(VERSION)
