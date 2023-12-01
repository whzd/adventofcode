def top_elf(list, cal):
    # print(list)
    # print(cal)
    # print()
    if cal > list[2]:
        if cal > list[1]:
            if cal > list[0]:
                list.pop(2)
                list.insert(0, cal)
            else:
                list.pop(1)
                list.insert(1, cal)
        else:
            list.pop(2)
            list.append(cal)


if __name__ == "__main__":

    with open('input') as f:
        content = f.readlines()

    top_3 = [0,0,0]
    cal = 0
    for line in content:
        if len(line) == 1:
            top_elf(top_3, cal)
            cal = 0
        else:
            cal += int(line)

    print(top_3)
    print(sum(top_3))

