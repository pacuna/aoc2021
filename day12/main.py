from collections import defaultdict, Counter
from typing import Dict, List


def build_adj_list(filename: str) -> Dict[str, List]:
    graph: defaultdict[str, list] = defaultdict(list)
    for line in open(filename):
        n1, n2 = line.strip().split('-')
        graph[n1].append(n2)
        graph[n2].append(n1)

    return graph


# Part I
def find_all_paths(graph):
    stack = [['start']]
    paths = []

    while stack:
        path = stack.pop()
        node = path[-1]
        if node != 'end':
            for neighbor in graph[node]:
                print(neighbor)
                if (neighbor.islower() and neighbor not in path) or (neighbor.isupper()):
                    stack.append(path + [neighbor])
        else:
            paths += [path]
    return paths

# Part II
def contains_dup(path: List[str]):
    c = Counter(path)
    for k, v in c.items():
        if k.islower() and v > 1:
            return True
    return False


def find_all_paths2(graph):
    stack = [['start']]
    paths = []

    while stack:
        path = stack.pop()
        node = path[-1]
        if node != 'end':
            for neighbor in graph[node]:
                if neighbor == 'start':
                    continue
                if (neighbor in path) and not contains_dup(path):
                    stack.append(path + [neighbor])
                elif (neighbor.islower() and neighbor not in path) or (neighbor.isupper()):
                    stack.append(path + [neighbor])
        else:
            paths += [path]
    return paths


if __name__ == '__main__':
    g = build_adj_list("data.in")
    p = find_all_paths2(g)
    for idx, path in enumerate(p):
        print(idx, path)
