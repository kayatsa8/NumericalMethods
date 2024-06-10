from typing import Callable, List
import ErrorCalculator as err





if __name__ == "__main__":
    original: List[float] = [5, 2]
    estimated: List[float] = [5.11, 1.93]

    f: Callable[[List[float]], float] = lambda args: args[0] ** args[1]

    effect: float = err.functionEffectOnResult(original=original, estimated=estimated, f=f)

    print(effect)