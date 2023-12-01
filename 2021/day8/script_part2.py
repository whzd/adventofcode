NUMBER0_FREQ = 42
NUMBER1_FREQ = 17
NUMBER2_FREQ = 34
NUMBER3_FREQ = 39
NUMBER4_FREQ = 30
NUMBER5_FREQ = 37
NUMBER6_FREQ = 41
NUMBER7_FREQ = 25
NUMBER8_FREQ = 49
NUMBER9_FREQ = 45

FREQ_TO_NUMBER = {NUMBER0_FREQ: 0, NUMBER1_FREQ: 1, NUMBER2_FREQ: 2, NUMBER3_FREQ: 3, NUMBER4_FREQ: 4,
                  NUMBER5_FREQ: 5, NUMBER6_FREQ: 6, NUMBER7_FREQ: 7, NUMBER8_FREQ: 8, NUMBER9_FREQ: 9}


def import_data():
    with open('input') as f:
        return  f.readlines()


def decode_value(numbers, outputs):
    freqA = 0
    freqB = 0
    freqC = 0
    freqD = 0
    freqE = 0
    freqF = 0
    freqG = 0

    for number in numbers:
        for digit in number:
            if digit == 'a':
                freqA += 1
            if digit == 'b':
                freqB += 1
            if digit == 'c':
                freqC += 1
            if digit == 'd':
                freqD += 1
            if digit == 'e':
                freqE += 1
            if digit == 'f':
                freqF += 1
            if digit == 'g':
                freqG += 1

    # USED TO CALCULATE EACH NUMBER FREQ, TO BUILD THE FREQ_TO_NUMBER DICTIONARY
    # for number in numbers:
    #     number_freq = 0
    #     for digit in number:
    #         if digit == 'a':
    #             number_freq += freqA
    #         if digit == 'b':
    #             number_freq += freqB
    #         if digit == 'c':
    #             number_freq += freqC
    #         if digit == 'd':
    #             number_freq += freqD
    #         if digit == 'e':
    #             number_freq += freqE
    #         if digit == 'f':
    #             number_freq += freqF
    #         if digit == 'g':
    #             number_freq += freqG
    #     print(f"{number} -> {number_freq}")
    # print()

    composed_value = ''
    for output in outputs:
        output_freq = 0
        for out in output:
            if out == 'a':
                output_freq += freqA
            if out == 'b':
                output_freq += freqB
            if out == 'c':
                output_freq += freqC
            if out == 'd':
                output_freq += freqD
            if out == 'e':
                output_freq += freqE
            if out == 'f':
                output_freq += freqF
            if out == 'g':
                output_freq += freqG
        # print(f"{output} -> {output_freq}")
        composed_value += str(FREQ_TO_NUMBER[output_freq])
    return composed_value

def main():
    total = 0
    content = import_data()
    for line in content:
        numbers, output = line.split(' | ')
        output_value = decode_value(numbers.split(' '), output.strip().split(' '))
        # print(output_value)
        # print()
        total += int(output_value)
    print(total)


if __name__ == "__main__":
    main()
