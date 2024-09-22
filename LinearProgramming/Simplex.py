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
    
    E = objective[:X]
    E = [x * -1 for x in E] + [0] * S + [objective[-1]] + [0]

    table.append(E)

    return table


def solve(table: List[List[float]]) -> List[float]:
    res: List[float] = []

    while True:
        col: int = getColumn(table)

        if table[-1][col] >= 0:
            break
        
        row: int = getRow(table=table, col=col)

        table[row] *= 1 / table[row][col]

        for i in range(len(table)):
            if i != row:
                table[i] -= table[i][col] * table[row][col]

    res = prepareResult(table)

    return res


def getColumn(table: List[List[float]]) -> int:
    min: int = table[-1][0]
    minIndex: int = 0

    for i in range(len(table[-1]) - 2):
        if min > table[-1][i]:
            min = table[-1][i]
            minIndex = i

    return minIndex


def prepareResult(table: List[List[float]]) -> List[float]:
    # TODO: complete
    print(table)
    return []


def getRow(table: List[List[float]], col: int) -> int:
    minTr: float = float("inf")
    minIndex: int = -1

    for i in range(len(table) - 1):
        if table[i][col] != 0:
            tr = table[i][-1] / table[i][col]

            if tr < minTr:
                minTr = tr
                minIndex = i

    return minIndex




