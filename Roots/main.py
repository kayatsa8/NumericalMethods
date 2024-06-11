from typing import Callable
from RootMethods import fixedPoint
import math





if __name__ == "__main__":
    a: float = 2
    b: float = 3
    f: Callable[[float], float] = lambda x: x**4 + 2*x**2 - x - 3
    g: Callable[[float], float] = lambda x: (3 + x - 2*x**2) ** 0.25
    delta: float = 0.0001

    root = fixedPoint(guess=1, g=g, delta=delta)

    print(root)