# indentutlis
## How to use
    >>> import indentutils
    >>> indentutils.process('''
    ... for i in cats:
    ...     print('i')
    ...     if i == '':
    ...         break
    ... ''')
    ['for i in cats:', ['print(i)', "if i == '':", ['break']]]

