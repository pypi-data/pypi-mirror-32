
def use_ipdb():
    # TODO: make `from kpa import use_ipdb` set `sys.excepthook` without `use_ipdb()`?
    import sys
    def excepthook(*args, **kwargs):
        from IPython.core import ultratb
        sys.excepthook = ultratb.FormattedTB(mode='Verbose', color_scheme='Linux', call_pdb=1)
        return sys.excepthook(*args, **kwargs)
    sys.excepthook = excepthook


def termcolor(text, fg=None, bg=None):
    from subprocess import check_output
    def get_code(args):
        if not hasattr(termcol,'_cache'): termcol._cache = {}
        if termcol._cache.get(args, None) is None: termcol._cache[args] = check_output(f'tput {args}'.split()).decode()
        return termcol._cache[args]
    return (get_code(f'setaf {fg}') if fg else '') + (get_code(f'setab {bg}') if bg else '') + text + get_code('sgr0')


class TerminalLineWrapper:
    # TODO: make these classmethods
    '.write_text(text) wraps long lines using an eight-space indent and returns the result'
    def __init__(self):
        import shutil
        self.cols = shutil.get_terminal_size().columns
    def write_text(self, text):
        for line in text.split('\n'):
            self._write_line(line)
    def _write_line(self, text):
        ret = ''
        x = self._how_many_chars_fit_in_width(text, self.cols)
        ret += (text[:x])
        while x < len(text):
            text = text[x:]
            x = self._how_many_chars_fit_in_width(text, self.cols - 8)
            ret += '\n' + ' '*8 + text[:x]
        return ret
    def _how_many_chars_fit_in_width(self, s, width):
        for i, c in enumerate(s):
            width -= self._width_of_char(c)
            if width < 0: return i
        return len(s)
    def _width_of_char(self, c):
        import unicodedata
        ## A: ambiguous (1 in iterm)
        ## F: full width (2)
        ## H: halfwidth (1)
        ## N : not asian (1)
        ## Na: narrow (1)
        ## W: wide (2)
        u_eaw = unicodedata.east_asian_width(c)
        if u_eaw in ('H','N','Na','A'): return 1
        elif u_eaw in ('F','W'): return 2
        else: raise Exception(ord(c))

