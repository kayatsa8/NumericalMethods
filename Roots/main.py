from typing import Callable
from RootMethods import regulaFalsi
import math





if __name__ == "__main__":
    a: float = 2
    b: float = 3
    f: Callable[[float], float] = lambda x: math.exp(math.sin(x)) - 2
    delta: float = 0.000001

    root = regulaFalsi(a, b, f, delta)

    print(root)