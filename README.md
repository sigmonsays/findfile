findfile is a simple tool that finds matching files using positional arguments

common prefix is found and ignored during matching to allow unique matches

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
