def calculate_overlapping(bank):
    count = 0
    for _, value in bank.items():
        if value > 1:
            count +=1
    return count

def save_lines(bank, key):
    if key in bank:
        bank[key] += 1
    else:
        bank[key] = 1


def test_lines(x1, x2, y1, y2, horizontal, bank):
    if horizontal:
        if x1 > x2:
            for i in range(x2, x1+1):
                key = f"{i},{y1}"
                save_lines(bank=bank, key=key)
        else:
            for i in range(x1, x2+1):
                key = f"{i},{y1}"
                save_lines(bank=bank, key=key)
    else:
        if y1 > y2:
            for i in range(y2, y1+1):
                key = f"{x1},{i}"
                save_lines(bank=bank, key=key)
        else:
            for i in range(y1, y2+1):
                key = f"{x1},{i}"
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
        if x1 == x2:
            test_lines(x1, x2, y1, y2, False, bank)
        if y1 == y2:
            test_lines(x1, x2, y1, y2, True, bank)
    print(calculate_overlapping(bank=bank))

if __name__ == "__main__":
    main()
