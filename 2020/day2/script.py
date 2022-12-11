with open("./input", "r") as f:
    valid_password = 0

    line = f.readline()
    while line:
        # 1. Get letter to check and password
        tmp = line.split(":")
        letter = tmp[0][-1]
        password = tmp[1].strip()

        # 2. Get lower limit
        tmp2 = tmp[0].split("-")
        a = eval(tmp2[0])   # eval() to convert limit to int

        # 3. Get upper limit
        tmp3 = tmp2[1].split(" ")
        b = eval(tmp3[0])   # eval() to convert limit to int

        ## Part1
        # 3. Check if number of occurrences of letter to check is inside de interval
        # if a <= password.count(letter) <= b:
            # valid_password += 1

        ## Part2
        # 3. Check if letter is either at position a or b. Not at both positoins. Index = Position - 1
        if password[a - 1] == letter:
            if password[b - 1] != letter:
                valid_password += 1
        else:
            if password[b - 1] == letter:
                valid_password += 1
        line = f.readline()
    print(valid_password)
