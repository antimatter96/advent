import re

inp = input()

split = inp.split(",")
ranges = [s.split("-") for s in split]
ranges = [[int(s[0]), int(s[1])] for s in ranges]

regexp = re.compile(r"^(\d+)(\1){1,}$")

sum = 0
for r in ranges:
    for i in range(r[0], r[1] + 1):
        if regexp.match(str(i)):
            sum += i

print(sum)
