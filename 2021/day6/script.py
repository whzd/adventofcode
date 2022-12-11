def get_input_data():
    with open('test_input') as f:
        return list(map(int, f.readline().strip().split(',')))


def main():
    fishes = get_input_data()
    for _ in range(0, 80):
        for i in range(0, len(fishes)):
            if fishes[i] == 0:
                fishes.append(8)
                fishes[i] = 6
            else:
                fishes[i] -= 1
    print(len(fishes))


if __name__ == "__main__":
    main()
