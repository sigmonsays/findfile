findfile is a simple tool that finds matching files using positional arguments as filters.

a common prefix is found and stripped during each positional argument processing.
This style of matching allows filters to be iterively extended.

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [Examples](#examples)
- [Usage](#usage)

<!-- markdown-toc end -->


# Examples

# Usage

    findfile -h
    Usage: findfile [--verbose] [--dirs-only] [--no-case] [--dir DIR] [--loglevel LOGLEVEL] [ARGS [ARGS ...]]

    Positional arguments:
    ARGS

    Options:
    --verbose, -v          be verbose
    --dirs-only, -D        match only directories
    --no-case, -i          case insensitive match
    --dir DIR, -d DIR      directory to look in [default: /home/sig/code/findfile]
    --loglevel LOGLEVEL, -l LOGLEVEL
                            log level [default: INFO]
    --help, -h             display this help and exit
