from typing import Callable, List
import math


def absoluteError(original: float, estimated: float) -> float:
    return abs(original - estimated)


def relativeError(original: float, estimated: float) -> float:
    return absoluteError(original, estimated) / abs(original)


# result < 1 ---> f reduce the error
# result = 1 ---> f preserve the error
# result > 1 ---> f increase the error
def functionEffectOnResult(original: List[float], estimated: List[float],
                           f: Callable[[List[float]], float]) -> float:
    
    if len(original) != len(estimated):
        raise Exception("the inputs are not in the same size")
    
    inputRelativeError: float = math.fsum([relativeError(org, est) for (org, est) in zip(original, estimated)])

    outputRelativeError: float = relativeError(f(original), f(estimated))

    return outputRelativeError / inputRelativeError



