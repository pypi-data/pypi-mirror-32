__title__        = "castero"
__description__  = "Command line podcast client"
__keywords__     = "podcast commandline terminal curses"
__author__       = "Jake Robertson"
__author_email__ = "jake@faltro.com"
__version__      = "0.2.5"
__date__         = "2018-06-15"
__copyright__    = "Copyright (c) 2018 Jake Robertson"
__license__      = "MIT License"
__url__          = "https://github.com/xgi/castero"

__help__ = """\
%s help
============

Version: %s
Updated: %s
Maintainer: %s <%s>
License: %s
URL: %s

Commands
    h            - show this help screen
    q            - exit the client
    a            - add a feed
    d            - delete the selected feed
    r            - reload/refresh feeds
    s            - save episode for offline playback
    arrows       - navigate menus
    page up/down - scroll menus
    enter        - play selected feed/episode
    space        - add selected feed/episode to queue
    c            - clear the queue
    n            - go to the next episode in the queue
    i            - invert the order of the menu
    p or k       - pause/play the current episode
    f or l       - seek forward
    b or j       - seek backward
""" % (
    __title__,
    __version__,
    __date__,
    __author__,
    __author_email__,
    __license__,
    __url__
)
