# content = [
# '00100',
# '11110',
# '10110',
# '10111',
# '10101',
# '01111',
# '00111',
# '11100',
# '10000',
# '11001',
# '00010',
# '01010'
# ]

content = []
reader = open('input', 'r')
for line in reader:
    line = line.strip()
    if len(line) > 0:
        content.append(line)

oxygen_list = content
oxygen = ''
for i in range(0,len(oxygen_list[0])):
    zero = 0
    one = 0
    for number in oxygen_list:
        if number[i] == '0':
            zero += 1
        else:
            one += 1
    if zero == one:
        oxygen += '1'
    elif zero > one:
        oxygen += '0'
    else:
        oxygen += '1'
    oxygen_list = [line for line in oxygen_list if line.startswith(oxygen)]
    if len(oxygen_list) == 1:
        print(oxygen_list[0])
        break

co2_list = content
co2 = ''
for i in range(0,len(co2_list[0])):
    zero = 0
    one = 0
    for number in co2_list:
        if number[i] == '0':
            zero += 1
        else:
            one += 1
    if zero == one:
        co2 += '0'
    elif zero > one:
        co2 += '1'
    else:
        co2 += '0'
    co2_list = [line for line in co2_list if line.startswith(co2)]
    if len(co2_list) == 1:
        print(co2_list[0])
        break

# print(gamma_rate)
# print(int(gamma_rate,2))
# print(epsilon_rate)
# print(int(epsilon_rate,2))
print(int(oxygen_list[0],2) * int(co2_list[0],2))
