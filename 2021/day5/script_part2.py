def calculate_overlapping(bank):
    count = 0
    for key, value in bank.items():
        if value > 1:
            count +=1
    return count


def save_lines(bank, key):
    if key in bank:
        bank[key] += 1
    else:
        bank[key] = 1


def get_horizontal_line_points(x1, x2, y1, bank):
    if x1 > x2:
        for i in range(x2, x1+1):
            key = f"{i},{y1}"
            save_lines(bank=bank, key=key)
    else:
        for i in range(x1, x2+1):
            key = f"{i},{y1}"
            save_lines(bank=bank, key=key)


def get_vertical_line_points(x1, y1, y2, bank):
    if y1 > y2:
        for i in range(y2, y1+1):
            key = f"{x1},{i}"
            save_lines(bank=bank, key=key)
    else:
        for i in range(y1, y2+1):
            key = f"{x1},{i}"
            save_lines(bank=bank, key=key)


def get_diagonal_line_points(x1, y1, x2, y2, bank):
    if x1 > x2:
        for i in range(0, (x1-x2)+1):
            key = f"{x2+i},{y2+i}"
            save_lines(bank=bank, key=key)
    else:
        for i in range(0, (x2-x1)+1):
            key = f"{x1+i},{y1+i}"
            save_lines(bank=bank, key=key)


def get_antidiagonal_line_points(x1, y1, x2, y2, bank):
    if x1 > x2:
        for i in range(0, (x1-x2)+1):
            key = f"{x2+i},{y2-i}"
            save_lines(bank=bank, key=key)
    else:
        for i in range(0, (x2-x1)+1):
            key = f"{x1+i},{y1-i}"
            save_lines(bank=bank, key=key)


def get_input_data():
    with open('input') as f:
        return f.readlines()


def main():
    bank = {}
    lines = get_input_data()
    for line in lines:
        line = line.strip().split(' -> ')
        x1, y1 = list(map(int, line[0].split(',')))
        x2, y2 = list(map(int, line[1].split(',')))

        # vertical
        if x1 == x2:
            get_vertical_line_points(x1, y1, y2, bank)
        # horizontal
        if y1 == y2:
            get_horizontal_line_points(x1, x2, y1, bank)
        # diagonal
        if x1 - y1 == x2 - y2:
            # print("diagonal")
            # print(line)
            # print()
            get_diagonal_line_points(x1, y1, x2, y2, bank)
        # antidiagonal
        if x1 + y1 == x2 + y2:
            # print("antidiagonal")
            # print(line)
            # print()
            get_antidiagonal_line_points(x1, y1, x2, y2, bank)

    print(calculate_overlapping(bank=bank))

if __name__ == "__main__":
    main()
