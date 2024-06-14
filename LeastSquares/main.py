from LeastSquares import LeastSquares
from Polynomial import Polynomial

points = {
    1: 3,
    2: 2,
    4: 3,
    5: 6
}
degree = 2


pol: Polynomial = LeastSquares(points, degree)

print(pol.coefficients)