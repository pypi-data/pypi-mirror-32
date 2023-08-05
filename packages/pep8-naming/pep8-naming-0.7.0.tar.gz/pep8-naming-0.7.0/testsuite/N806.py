#: Okay
def test():
    good = 1
#: Okay
def test():
    def test2():
        good = 1
#: Okay
GOOD = 1
#: Okay
class Test(object):
    GOOD = 1
#: N806
def test():
    Bad = 1
#: N806
def test():
    VERY = 2
#: N806
def test():
    def test2():
        class Foo(object):
            def test3(self):
                Bad = 3
#: Okay
def good():
    global Bad
    Bad = 1
#: N806
def bad():
    global Bad

    def foo():
        Bad = 1

#: Okay
def test():
    # namedtuples are often CamelCased since we treat them a bit like classes
    import collections
    Thing = collections.namedtuple('Thing', 'a b c')
    from collections import namedtuple
    ThingTwo = namedtuple('ThingTwo', 'a b c')

#: N806
def bad():
    # Currently don't support aliased imports of namedtuple
    from collections import namedtuple as nt
    Thing = nt('Thing', 'a b c')

#: N806
def unpacking_into_tuple():
    Var1, Var2 = range(2)
#: Okay
def unpacking_into_tuple():
    var1, var2 = range(2)
#: N806
def unpacking_into_list():
    [Var1, Var2] = range(2)
#: Okay
def unpacking_into_list():
    [var1, var2] = range(2)
#: Okay
a, [b, c] = [1, [2, 3]]
#: N806
def recursive_unpack():
    a, [bB, c] = [1, [2, 3]]
#: Okay
def assingnment_to_attribute():
    a.b = 1
#: N806
def f():
    with Foo(), Bar() as Bad:
        pass
#: Okay
def f():
    with FOO() as foo, bar() as bar:
        pass
#: Okay
def f():
    with suppress(E):
        pass
    with contextlib.suppress(E):
        pass
#: Okay
with Test() as bar:
    pass
#: N806
def f():
    with Test() as BAD:
        pass
#: Okay
def f():
    with C() as [a, b, c]:
        pass
#: N806
def f():
    with C() as [a, Bad, c]:
        pass
#: N806
def f():
    with C() as (a, b, baD):
        pass
