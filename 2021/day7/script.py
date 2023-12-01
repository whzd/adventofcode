def import_data():
    with open('input') as f:
        return list(map(int, f.readline().strip().split(',')))


def main():
    total = 0
    positions = import_data()
    positions.sort()
    for i in range(positions[0], positions[-1]):
        abs_sum = 0
        for pos in positions:
            abs_sum += abs(i-pos)
        if i == positions[0]:
            total = abs_sum
        elif abs_sum < total:
            total = abs_sum
    print(total)


if __name__ == "__main__":
    main()
