import re

_modifier_regex = (
    '[._-]?'
    '(?:(stable|beta|b|RC|c|pre|alpha|a|patch|pl|p|post|[a-z])'
    '((?:[.-]?\d+)*)?)?'
    '([.-]?dev)?'
)


def normalize_version(version):
    """
    Normalizes a version string to be able to perform comparisons on it.
    """
    version = version.strip()

    # strip off build metadata
    m = re.match('^([^,\s+]+)\+[^\s]+$', version)
    if m:
        version = m.group(1)

    index = None
    # Match classic versioning
    m = re.match(
        '(?i)^v?(\d{{1,5}})(\.\d+)?(\.\d+)?(\.\d+)?{}$'.format(
            _modifier_regex
        ),
        version
    )
    if m:
        version = '{}{}{}{}'.format(
            m.group(1),
            m.group(2) if m.group(2) else '.0',
            m.group(3) if m.group(3) else '.0',
            m.group(4) if m.group(4) else '.0',
        )
        index = 5
    else:
        # Some versions have the form M.m.p-\d+
        # which means M.m.p-post\d+
        m = re.match(
            '(?i)^v?(\d{1,5})(\.\d+)?(\.\d+)?(\.\d+)?-(\d+)$',
            version
        )
        if m:
            version = '{}{}{}{}'.format(
                m.group(1),
                m.group(2) if m.group(2) else '.0',
                m.group(3) if m.group(3) else '.0',
                m.group(4) if m.group(4) else '.0',
            )
            if m.group(5):
                version += '-post.' + m.group(5)

            m = re.match(
                '(?i)^v?(\d{{1,5}})(\.\d+)?(\.\d+)?(\.\d+)?{}$'.format(
                    _modifier_regex
                ),
                version
            )

            index = 5
        else:
            # Match date(time) based versioning
            m = re.match(
                '(?i)^v?(\d{{4}}(?:[.:-]?\d{{2}}){{1,6}}(?:[.:-]?\d{{1,3}})?){}$'.format(
                    _modifier_regex
                ),
                version
            )
            if m:
                version = re.sub('\D', '.', m.group(1))

                index = 2

    # add version modifiers if a version was matched
    if index is not None:
        if len(m.groups()) - 1 >= index and m.group(index):
            if m.group(index) == 'post':
                # Post releases should be considered
                # stable releases
                if '-post' in version:
                    return version

                version = '{}-post'.format(version)
            else:
                version = '{}-{}'.format(
                    version, _expand_stability(m.group(index))
                )

            if m.group(index + 1):
                version = '{}.{}'.format(
                    version, m.group(index + 1).lstrip('.-')
                )

        return version

    raise ValueError('Invalid version string "{}"'.format(version))


def normalize_stability(stability):  # type: (str) -> str
    stability = stability.lower()

    if stability == 'rc':
        return 'RC'

    return stability


def parse_stability(version):  # type: (str) -> str
    """
    Returns the stability of a version.
    """
    version = re.sub('(?i)#.+$', '', version)

    if 'dev-' == version[:4] or '-dev' == version[-4:]:
        return 'dev'

    m = re.search('(?i){}(?:\+.*)?$'.format(_modifier_regex), version.lower())
    if m:
        if m.group(3):
            return 'dev'

        if m.group(1):
            if m.group(1) in ['beta', 'b']:
                return 'beta'
            elif m.group(1) in ['alpha', 'a']:
                return 'alpha'
            elif m.group(1) in ['rc', 'c']:
                return 'RC'
            elif m.group(1) == 'post':
                return 'stable'
            else:
                return 'dev'

    return 'stable'


def _expand_stability(stability):  # type: (str) -> str
    stability = stability.lower()

    if stability == 'a':
        return 'alpha'
    elif stability == 'b':
        return 'beta'
    elif stability in ['c', 'pre']:
        return 'rc'
    elif stability in ['p', 'pl']:
        return 'patch'
    elif stability in ['post']:
        return 'post'

    return stability
