

def absoluteError(original: float, estimated: float) -> float:
    return abs(original - estimated)

def relativeError(original: float, estimated: float) -> float:
    return absoluteError(original, estimated) / abs(original)
