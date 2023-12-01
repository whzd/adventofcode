with open('input') as f:
    depths = [eval(line.rstrip()) for line in f]

# depths = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]

cont = 0
sums = 0
prev_sums = 0
for i in range(2,len(depths)):
    print("==============")
    sums = depths[i] + depths[i-1] + depths[i-2]
    print(sums)
    print(prev_sums)
    if prev_sums == 0:
        prev_sums = sums
    elif sums > prev_sums:
        cont += 1
    prev_sums = sums

    print(cont)
    print()

print(cont)
