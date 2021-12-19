
import statistics
stack = []
matches = {}
matches[')'] = '('
matches[']'] = '['
matches['}'] = '{'
matches['>'] = '<'

points = {}
points[')'] = 3
points[']'] = 57
points['}'] = 1197
points['>'] = 25137


def solveI():
    res = 0
    invalids = []
    for line in open("data.in"):
        line = line.strip()
        for c in line:
            if c in matches.values():
                stack.append(c)
            else:
                last = stack.pop()
                if last != matches[c]:
                    invalids.append(c)
                    break

    for i in invalids:
        res += points[i]

    print(res)

def solveII():
    res = 0
    idx = 0
    invalids = []
    for line in open("data.in"):
        stack = []
        line = line.strip()
        for c in line:
            if c in matches.values():
                stack.append(c)
            else:
                last = stack.pop()
                if last != matches[c]:
                    invalids.append(idx)
                    break
        idx += 1

    incompletes = []
    for idx, line in enumerate(open("data.in")):
        if idx in invalids:
            continue
        line = line.strip()
        stack = []
        for c in line:
            if c in matches.values():
                stack.append(c)
            else:
                last = stack.pop()
        incompletes.append(stack)
        idx += 1

    points = {}
    points["("] = 1
    points["["] = 2
    points["{"] = 3
    points["<"] = 4

    scores = []
    for i in incompletes:
        score = 0
        while i:
            last = i.pop()
            score *= 5
            score += points[last]
        scores.append(score)

    print(statistics.median(scores))


solveII()