# -*- coding: utf-8 -*-

from __future__ import unicode_literals
from __future__ import print_function

from .ylsys import py2, pyi
import os, sys
from functools import reduce, wraps

printf = print


if py2:
    unicode = __builtins__['unicode']
else:
    unicode = str

def isstr(s):
    '''
    `isinstance(s, str)` for compatibility both python 2/3
    '''
    return isinstance(s, (str, unicode))


class Classobj:
    ''' 兼容 python 2 的 classobj'''
    pass
classobj = type(Classobj)
def istype(objOrType):
    ''' 
    `isinstance(objOrType, type)` for compatibility both python 2/3
    '''
    return isinstance(objOrType, (type, classobj))


def setFuncation2ForAutoList():
    '''
    add funcation 'range2', 'map2', 'reduce2', 'zip2' that retrun a list like python2
    '''
    funForAutoList = ['range', 'map', 'reduce', 'filter', 'zip', 'enumerate']
    if py2:
        addListAfter = lambda x:x
    else:    
        def addListAfter(f):
            @wraps(f)
            def innerF(*l, **kv):
                r = f(*l, **kv)
                return list(r)
            return innerF
    reduce
    glob = globals()
    for name in funForAutoList:
        rawf = glob[name] if name in glob else  __builtins__[name]
        glob[name+'2'] = addListAfter(rawf)

setFuncation2ForAutoList()


def __setDisplayEnv():
    msg = '''%s
        os.environ["DISPLAY"] not found, may cuse error like this 
        [QXcbConnection Error](https://github.com/ipython/ipython/issues/10627)
        so, we auto set os.environ["QT_QPA_PLATFORM"] = "offscreen"    '''%'\x1b[36m%s\x1b[0m'% 'warning from boxx'
    msg = '''os.environ["DISPLAY"] is not found
    plt.show() are redirect to plt.savefig(tmp)
    funcation: show, loga, crun, heatmap, plot will be affected'''
    
    from .tool import warn
    warn(msg)
    os.environ["QT_QPA_PLATFORM"] = "offscreen"
    os.environ['DISPLAY'] = ':0'
    
def __noDisplayEnv():
    __setDisplayEnv()
    import matplotlib.pyplot as plt
    def savefig(*l, **kv):
        from .ylsys import tmpboxx
        from .tool import increase, warn, OffScreenWarning
        showtmp = os.path.join(tmpboxx(), 'showtmp')
        if not os.path.isdir(showtmp):
            os.makedirs(showtmp)
        png = os.path.join(showtmp,'%d_show.png'%increase('showtmp'))
        warn('''os.environ["DISPLAY"] is not found
    plt.show() are redirect to plt.savefig("%s")'''%png, OffScreenWarning)
        plt.savefig(png)
        plt.cla()
        plt.clf()
#    pltshow = plt.show
    plt.show = savefig
def beforImportPlt():
    '''
    if os.environ["DISPLAY"] is not found, then  we auto set os.environ["QT_QPA_PLATFORM"] = "offscreen"
    and plt.show() are redirect to plt.savefig('/tmp/boxxTmp/showtmp')
    
    this would avoid exit error.
    '''
    if not pyi.plt and not pyi.reloadplt:
        __noDisplayEnv()
        import matplotlib.pyplot as plt
        pyi.reloadplt = plt

class SetPltActivateInWith():
    '''
    set plt to interactivate mode while use plt.show() inner boxx
    '''
    def __enter__(self):
        if pyi.gui or not pyi.plt:
            return 
        import matplotlib.pyplot as plt
        self.interact = plt.rcParams['interactive']
        if not self.interact and pyi.ipython:
            plt.ion()
    def __exit__(self, *l):
        if pyi.gui or not pyi.plt:
            return 
        import matplotlib.pyplot as plt
        if not self.interact and pyi.ipython:
            plt.ioff()
    
def lazyplt(fun):
    '''
    cause `import matplotlib.pyplot` often lead to Exit Error
    I do `import plt` when funcation need 
    Farther more, if pyi.gui is False, the plt.show() are replace to plt.savefig('/tmp/boxxTmp/showtmp')
    '''
    @wraps(fun)
    def innerf(*l, **kv):
        beforImportPlt()
        with SetPltActivateInWith():
            r = fun(*l, **kv)
        return r
    return innerf
if not py2 and 0:
    __rawOpen__ = open
    open = lambda *l:__rawOpen__(l[0],'r',-1,'utf8') if len(l) == 1 else __rawOpen__(l[0],l[1],-1,'utf8')

