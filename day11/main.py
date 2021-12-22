import collections
from typing import List, Set, Tuple

M: List[List[int]] = []
for line in open("data.in"):
    M.append([int(x) for x in line.strip()])


def bfs(grid: List[List[int]], r: int, c: int, flashed: Set[Tuple[int, int]]) -> None:
    directions = [[1, 0], [-1, 0], [0, 1], [0, -1], [1, 1], [1, -1], [-1, 1], [-1, -1]]
    q = collections.deque()
    q.append((r, c))

    while q:
        row, col = q.popleft()
        for dr, dc in directions:
            r, c = row + dr, col + dc
            if r in range(len(grid)) and c in range(len(grid[0])):
                grid[r][c] += 1
                if grid[r][c] > 9 and (r, c) not in flashed:
                    flashed.add((r, c))
                    q.append((r, c))


def print_grid(grid):
    for r in range(len(grid)):
        print(" ".join([str(x) for x in grid[r]]))


def all_flash(grid) -> bool:
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] != 0:
                return False
    return True


def run_steps(num_steps: int) -> None:
    total_flashes = 0
    for i in range(num_steps):
        flashed = set()
        for r in range(len(M)):
            for c in range(len(M[0])):
                M[r][c] += 1

        for r in range(len(M)):
            for c in range(len(M[0])):
                if M[r][c] > 9 and (r, c) not in flashed:
                    flashed.add((r, c))
                    bfs(M, r, c, flashed)

        for row, col in flashed:
            M[row][col] = 0
        total_flashes += len(flashed)
        if all_flash(M):
            print("flashed at ", i+1)


if __name__ == '__main__':
    run_steps(300)
