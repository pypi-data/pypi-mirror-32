# python3 only
#: Okay
VAR1, *VAR2, VAR3 = 1, 2, 3
#: Okay
[VAR1, *VAR2, VAR3] = (1, 2, 3)
#: N806
def extended_unpacking_not_ok():
    Var1, *Var2, Var3 = 1, 2, 3
#: N806
def extended_unpacking_not_ok():
    [Var1, *Var2, Var3] = (1, 2, 3)
#: Okay
def assing_to_unpack_ok():
    a, *[b] = 1, 2
#: N806
def assing_to_unpack_not_ok():
    a, *[bB] = 1, 2
