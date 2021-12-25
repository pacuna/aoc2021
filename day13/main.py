from typing import List


def fold_y(grid, mag):
    g1 = grid[:mag]
    g2 = grid[mag + 1:]

    g3 = [[0 for x in range(len(g1[0]))] for y in range(len(g1))]

    for i in range(len(g1)):
        for j in range(len(g1[0])):
            if g1[i][j] == 1 or g2[len(g2) - i - 1][j] == 1:
                g3[i][j] = 1
    return g3


def fold_x(grid, mag):
    g4 = [[0 for x in range(mag)] for y in range(len(grid))]
    g5 = [[0 for x in range(mag)] for y in range(len(grid))]

    for i in range(len(grid)):
        for j in range(len(grid[i])):
            if j < mag and grid[i][j] == 1:
                g4[i][j] = 1
            elif j > mag and grid[i][j] == 1:
                g5[i][j - (mag + 1)] = 1

    g6 = [[0 for x in range(len(g4[0]))] for y in range(len(g4))]

    for i in range(len(g6)):
        for j in range(len(g6[i])):
            if g4[i][j] == 1 or g5[i][len(g5[j]) - j - 1] == 1:
                g6[i][j] = 1

    return g6


def print_grid(grid: List[List[int]]) -> None:
    for r in range(len(grid)):
        for c in range(len(grid[r])):
            if grid[r][c] == 0:
                grid[r][c] = '.'
            else:
                grid[r][c] = '#'

    for r in range(len(grid)):
        print(''.join(str(grid[r])))

    for r in range(len(grid)):
        print(''.join(str(grid[r])))


def count_dots(grid: List[List[int]]) -> int:
    res = 0
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            if grid[i][j] == 1:
                res += 1

    return res


if __name__ == '__main__':
    coords = []
    folds = []
    f = open('data.in', 'r')
    content = f.readlines()

    for i in range(1125):
        x, y = [int(c) for c in content[i].strip().split(',')]
        coords.append((x, y))

    for i in range(1126, 1138):
        _, _, fold = content[i].strip().split(' ')
        axis, mag = fold.split('=')
        folds.append((axis, mag))

    max_x = max(list(map(lambda x: x[0], coords)))
    max_y = max(list(map(lambda y: y[1], coords)))

    grid = []
    for i in range(max_y + 1):
        grid.append([0 for x in range(max_x + 1)])

    for c in coords:
        grid[c[1]][c[0]] = 1

    tmp_grid = grid
    for fold in folds:
        print("folding: ", fold[0], fold[1])
        if fold[0] == 'y':
            tmp_grid = fold_y(tmp_grid, int(fold[1]))
        elif fold[0] == 'x':
            tmp_grid = fold_x(tmp_grid, int(fold[1]))

    print_grid(tmp_grid)