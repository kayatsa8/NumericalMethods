from Simplex import simplex


from typing import List


objective: List[float] = [5, 4, 1]
constraints: List[List[float]] = [
    [3, 5, 78],
    [4, 1, 36]
]

res: List[float] = simplex(objective=objective, constraints=constraints)

print(res)










