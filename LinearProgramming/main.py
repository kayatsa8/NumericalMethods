from Simplex import simplex


from typing import List


objective: List[float] = [4, 5, 1]
constraints: List[List[float]] = [
    [14, 11, 154],
    [7, 16, 112]
]

res: List[float] = simplex(objective=objective, constraints=constraints)

print(res)










