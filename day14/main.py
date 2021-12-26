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
