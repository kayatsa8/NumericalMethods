from typing import Dict, List
from Polynomial import Polynomial

import requests
import json 


def LeastSquares(points: Dict[float, float], polynomialDegree: int) -> Polynomial:

    xPowI: List[float] = [sum([x ** i for x in points]) for i in range(polynomialDegree * 2 + 1)]

    xPowIY: List[float] = [sum([(x ** i) * y for x, y in points.items()]) for i in range(polynomialDegree * 2 + 1)]

    matrix: List[List[float]] = makeMatrix(xPowI=xPowI, xPowIY=xPowIY, degree=polynomialDegree)

    result: List[float] = solveMatrix(matrix=matrix)

    coefficients: Dict[int, float] = {degree: coefficient for degree, coefficient in enumerate(result)}

    polynomial: Polynomial = Polynomial(coefficients=coefficients)

    return polynomial



def makeMatrix(xPowI: List[float], xPowIY: List[float], degree: int) -> List[List[float]]:
    matrix: List[List[float]] = []

    for i in range(degree + 1):
        row: List[float] = [xPowI[i + j] for j in range(degree + 1)]

        row.append(xPowIY[i])

        matrix.append(row)

    return matrix


def solveMatrix(matrix: List[List[float]]) -> List[float]:
    req = requests.post("http://localhost:3004/api/solve_equations/gaussian", json=matrix)
    responseDict = json.loads(req.text)

    if responseDict["error"] is not None:
        raise Exception("fetching error: " + responseDict["error"])
    
    result: List[float] = responseDict["result"]

    return result


