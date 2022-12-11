from itertools import combinations

with open("./input", "r") as f:
    content = [eval(line.rstrip()) for line in f]   # eval() to convert the values to int to make it easier to process

## Part 1: Find the product of the two values that added together are equal to 2020
# for value in content:
#     b = 2020 - value
#     if b in content:
#         print(value*b)
#         break

## Part 2: Find the product of the three values that added together are equal to 2020
three_values = [x for x in combinations(content, 3) if sum(x) == 2020 ]
res = 1
for value in three_values[0]:
    res *= value
print(res)
