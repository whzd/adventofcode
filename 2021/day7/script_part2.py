def import_data():
    with open('input') as f:
        return list(map(int, f.readline().strip().split(',')))


def main():
    total = 0
    positions = import_data()
    positions.sort()
    for i in range(positions[0], positions[-1]):
        part_sum = 0
        for pos in positions:
            # The nth partial sum formula is n*(n+1)/2
            part_sum += abs(i-pos)*(abs(i-pos)+1)/2
        # print(part_sum)
        if i == positions[0]:
            total = part_sum
        elif part_sum < total:
            total = part_sum
    print(int(total))


if __name__ == "__main__":
    main()
