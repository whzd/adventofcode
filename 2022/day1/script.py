with open('input') as f:
    content = f.readlines()

most_cal = 0
cal = 0
for line in content:
    if len(line) == 1:
        if cal > most_cal:
            most_cal = cal
        cal = 0
    else:
        cal += int(line)

print(most_cal)

