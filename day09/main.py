from typing import List, Tuple
import collections


M: List[List[int]] = []
for line in open('data.in'):
    row: List[int] = []
    line = line.strip()
    for c in line:
        row.append(int(c))
    M.append(row)

def is_low_point(M: List[List[int]], row: int, col: int) -> bool:
    directions: List[List[int]] = [[1, 0], [-1, 0], [0, 1], [0, -1]]
    for dr, dc in directions:
        r, c = row + dr, col + dc
        if r in range(len(M)) and c in range (len(M[0])) and M[r][c] <= M[row][col]:
            return False
    return True

low_points: List[Tuple[int, int]] = []
for r in range(len(M)):
    for c in range(len(M[r])):
        if is_low_point(M, r, c):
            low_points.append((r, c))


# part I:
risk_level: int = 0
for lp in low_points:
    risk_level += M[lp[0]][lp[1]] + 1

print(risk_level)

# part II:

visit = set()
def bfs(M: List[List[int]], row: int, col: int) -> None:
    q = collections.deque()
    visit.add((row, col))
    q.append((row, col))
    basin = []
    basin.append(M[row][col])
    while q:
        R, C = q.popleft()
        directions: List[List[int]] = [[1, 0], [-1, 0], [0, 1], [0, -1]]
        for dr, dc in directions:
            r, c = R + dr, C + dc
            if r in range(len(M)) and c in range(len(M[0])) and M[r][c] != 9 and (r, c) not in visit:
                q.append((r, c))
                basin.append(M[r][c])
                visit.add((r, c))
    return basin

basins = []
for r in range(len(M)):
    for c in range(len(M[r])):
        if is_low_point(M, r, c) and (r, c) not in visit:
            basins.append(bfs(M, r, c))

basins.sort(key=len, reverse=True)
print(len(basins[0])*len(basins[1])*len(basins[2]))