def get_bingo_data():
    with open('input') as f:
        numbers = f.readline().split(',')

        boards = {}
        board_count = 1
        line = f.readline()
        while line:
            board = []
            for _ in range(0, 5):
                line = f.readline()
                board.append(line.rsplit())
            boards[board_count] = board
            board_count += 1
            line = f.readline()
        return numbers, boards

def check_row(seen_numbers, row_numbers):
    result = [number for number in row_numbers if number not in seen_numbers]
    if result:
        return False
    else:
        return True

def check_colum(seen_numbers, colum_number, board):
    result = [row[colum_number] for row in board if row[colum_number] not in seen_numbers]
    if result:
        return False
    else:
        return True

def check_winner(seen_numbers, boards):
    for board_number in boards:
        for row in boards[board_number]:
            if(check_row(seen_numbers=seen_numbers, row_numbers=row)):
                return board_number
        for i in range(0, 5):
            if(check_colum(seen_numbers=seen_numbers, colum_number=i, board=boards[board_number])):
                return board_number

def calculate_sum(seen_numbers, board):
    result = 0
    for row in board:
        for number in row:
            if number not in seen_numbers:
                result += int(number)
    return result

if __name__ == "__main__":
    numbers, boards = get_bingo_data()
    seen_numbers = []
    for number in numbers:
        seen_numbers.append(number)
        winner_board = check_winner(seen_numbers, boards)
        if winner_board:
            unmarked_sum = calculate_sum(seen_numbers=seen_numbers, board=boards[winner_board])
            print(unmarked_sum)
            print(seen_numbers[-1])
            print(unmarked_sum * int(seen_numbers[-1]))
            break
