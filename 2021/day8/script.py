def import_data():
    with open('input') as f:
        content = f.readlines()
        digits = []
        for line in content:
            for digit in line.split(' | ')[1].strip().split(' '):
                digits.append(digit)
        return digits


def main():
    count = 0
    digits = import_data()
    for digit in digits:
        if len(digit) == 2 or len(digit) == 4 or len(digit) == 3 or len(digit) == 7:
            count += 1
    print(count)


if __name__ == "__main__":
    main()
