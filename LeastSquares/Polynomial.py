from typing import Dict


class Polynomial:
    def __init__(self, coefficients: Dict[int, float]) -> None:
        self.coefficients: Dict[int, float] = coefficients