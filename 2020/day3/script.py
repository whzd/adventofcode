## Part 1
def part1():
    with open("./input", "r") as f:
        x = 0
        tree_count = 0
        line = f.readline().rstrip()
        line = f.readline().rstrip() # skip first line
        while line:
            x = (x+3) % (len(line))
            # First draft
            # if y + 3 >= len(line):
            #     y = y + 3 - len(line)
            # else:
            #     y = y+ 3
            if line[x] == "#":
                tree_count += 1
            line = f.readline().rstrip()
        print(tree_count)

## Part 2
def read_file():
    with open("./input", "r") as f:
        return [line.rstrip() for line in f]

def calculate_path(lines, right, down):
    tree_count = 0
    x = 0
    y = 0
    while y + down < len(lines):
        y = y + down
        line = lines[y]
        x = (x + right) % (len(line))
        if line[x] == "#":
            tree_count += 1
    return tree_count

if __name__ == "__main__":
    content = read_file()
    res1 = calculate_path(content, 1, 1)
    print(res1)
    res2 = calculate_path(content, 3, 1)
    print(res2)
    res3 = calculate_path(content, 5, 1)
    print(res3)
    res4 = calculate_path(content, 7, 1)
    print(res4)
    res5 = calculate_path(content, 1, 2)
    print(res5)
    print(f"Final res: {res1*res2*res3*res4*res5}")


