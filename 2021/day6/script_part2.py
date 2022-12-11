def get_input_data():
    with open('input') as f:
        return list(map(int, f.readline().strip().split(',')))


def main():
    fishes = get_input_data()
    day0 = 0
    day1 = 0
    day2 = 0
    day3 = 0
    day4 = 0
    day5 = 0
    day6 = 0
    day7 = 0
    day8 = 0

    for fish in fishes:
        if fish == 0:
            day0 += 1
        elif fish == 1:
            day1 += 1
        elif fish == 2:
            day2 += 1
        elif fish == 3:
            day3 += 1
        elif fish == 4:
            day4 += 1
        elif fish == 5:
            day5 += 1
        elif fish == 6:
            day6 += 1
        elif fish == 7:
            day7 += 1
        else:
            day8 += 1

    for _ in range(0,256):
        tmp = day0
        day0 = day1
        day1 = day2
        day2 = day3
        day3 = day4
        day4 = day5
        day5 = day6
        day6 = day7
        day7 = day8
        day8 = tmp
        day6 += tmp

    # print(day0)
    # print(day1)
    # print(day2)
    # print(day3)
    # print(day4)
    # print(day5)
    # print(day6)
    # print(day7)
    # print(day8)
    # print()

    print(day0 + day1 + day2 + day3 + day4 + day5 + day6 + day7 + day8)


if __name__ == "__main__":
    main()
