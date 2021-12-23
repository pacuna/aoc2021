from collections import defaultdict
from typing import Any, Dict, List


def build_adj_list(filename: str) -> Dict[str, List]:
    graph: defaultdict[str, list] = defaultdict(list)
    for line in open(filename):
        n1, n2 = line.strip().split('-')
        graph[n1].append(n2)
        if n1 != 'start':
            graph[n2].append(n1)

    return graph


def find_all_paths(graph, start, end, path=None):
    if path is None:
        path = []
    path = path + [start]
    if start == end:
        return [path]
    if start not in graph:
        return []
    paths = []
    for node in graph[start]:
        if (node.islower() and path.count(node) == 0) or (node.isupper()):
            new_paths = find_all_paths(graph, node, end, path)
            for new_path in new_paths:
                paths.append(new_path)
    return paths


if __name__ == '__main__':
    g = build_adj_list("data.in")
    all_paths = find_all_paths(g, 'start', 'end')
    print(len(all_paths))
