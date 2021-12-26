import collections

rules = {}
template = ""
for line in open("data.in"):
    line = line.strip()
    if '->' in line:
        rule, _, target = line.split(' ')
        rules[rule] = target
    else:
        if line:
            template = line

# Part I
for i in range(10):
    new_template = ""
    for i in range(len(template)):
        if template[i:i+2] in rules:
            new_template += template[i]
            new_template += rules[template[i:i+2]]
        else:
            new_template += template[i:i+2]
    template = new_template
    c = collections.Counter(template)
    print(c)

# Part II

pair_freq = collections.defaultdict(int)
char_freq = collections.defaultdict(int)

for i in range(len(template)-1):
    pair_freq[template[i:i+2]] += 1

for i in range(len(template)):
    char_freq[template[i]] += 1

tmp = pair_freq.copy()
for i in range(40):
    for pair, cnt in pair_freq.items():
        if pair in rules:
            tmp[pair[0]+rules[pair]] += cnt
            tmp[rules[pair]+pair[1]] += cnt
            char_freq[rules[pair]] += cnt
            tmp[pair] -= cnt
    pair_freq = tmp.copy()

print(char_freq)

sorted_vals = sorted(char_freq.values())
print(sorted_vals[-1] - sorted_vals[0])