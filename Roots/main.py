from typing import Callable
from RootMethods import newtonMethod
import math





if __name__ == "__main__":
    f: Callable[[float], float] = lambda x: x**4 + 2*x**2 - x - 3
    fTag: Callable[[float], float] = lambda x: 4*x**3 + 4*x - 1
    delta: float = 0.0001

    root = newtonMethod(guess=2, f=f, fTag=fTag, delta=delta)

    print(root)