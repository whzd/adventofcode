with open('input') as f:
    content = f.read().splitlines()

player_points = {'X': 1, 'Y': 2, 'Z': 3}
elf_points = {'A': 1, 'B': 2, 'C': 3}

total = 0
for line in content:
    elf, player = line.split(' ')

    total += player_points[player]
    if elf_points[elf] == player_points[player]:
        total += 3
    elif elf_points[elf] == 3 and player_points[player] == 1:
            total += 6
    elif elf_points[elf] == 2 and player_points[player] == 3:
            total += 6
    elif elf_points[elf] == 1 and player_points[player] == 2:
            total += 6

    # print("elf " + elf)
    # print("player " + player)
    # print(f"total {total}")
    # print()
print(total)
