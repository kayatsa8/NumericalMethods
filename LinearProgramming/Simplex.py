from typing import List

def simplex(objective: List[float], constraints: List[List[float]]) -> List[float]:
    if not validate(objective, constraints):
        raise Exception("input is not valid")
    
    table: List[List[float]] = prepareTable(objective, constraints)

    sol: List[float] = solve(table)

    return sol


def validate(objective: List[float], constraints: List[List[float]]) -> bool:
    for c in constraints:
        if len(objective) != len(c):
            return False
        
    return True


def prepareTable(objective: List[float], constraints: List[List[float]]) -> List[List[float]]:
    X: int = len(objective) - 1
    S: int = len(constraints)
    table: List[List[float]] = []

    for i, c in enumerate(constraints):
        E = c[:X] + [0] * (S + 1)
        E[X + i] = 1
        E.append(c[-1])

        table.append(E)

    E = -1 * objective[:X] + [0] * S + objective[-1] + [0]

    table.append(E)

    return table




def solve(table: List[List[float]]) -> List[float]:
    pass
