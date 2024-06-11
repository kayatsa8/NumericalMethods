from typing import Callable


def bisection(a: float, b: float, f: Callable[[float], float], delta: float) -> float:
    if f(a) * f(b) > 0:
        raise Exception("both bounds have the same sign")

    start: float = a
    end: float = b
    mid: float = 0

    while end - start > delta:
        mid = (start + end) / 2

        f_mid: float = f(mid)

        if f_mid == 0:
            return mid

        if f(mid) * f(start) > 0:
            start = mid
        else:
            end = mid

    return mid


def regulaFalsi(a: float, b: float, f: Callable[[float], float], delta: float) -> float:
    if f(a) * f(b) > 0:
        raise Exception("both bounds have the same sign")
    
    start: float = a
    end: float = b
    mid: float = 0

    while end - start > delta:
        f_start: float = f(start)

        m: float = (f(end) - f(start)) / (end - start)

        mid = start - f_start / m

        f_mid: float = f(mid)

        if f_mid == 0:
            return mid
        
        if f_mid * f_start > 0:
            start = mid
        else:
            end = mid

    return mid
    
