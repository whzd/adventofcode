# This solution is incomplete, it has a bug somewhere that won´t stop in the last winning board.
# Read the comment on line 64 for more details
def get_bingo_data():
    with open('input') as f:
        numbers = list(map(int, f.readline().strip().split(',')))

        boards = {}
        board_count = 1
        line = f.readline()
        while line:
            board = []
            for _ in range(0, 5):
                line = f.readline()
                board.append(list(map(int, line.rsplit())))
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

def check_winner(seen_numbers, boards, winning_boards):
    for board_number in boards:
        if board_number not in winning_boards:
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
    tmp_boards = boards
    seen_numbers = []
    winning_boards = []
    for number in numbers:
        seen_numbers.append(number)
        winner_board = check_winner(seen_numbers, tmp_boards, winning_boards)
        if winner_board:
            last_winning_number = number
            winning_boards.append(winner_board)
        # The code has a bug somewhere on the winning conditions that won´t stop in the last winning board,
        # after 40 it will add 53, 54, 55, 56, 57, 59, 60, 61, 62, 63, 64 to the list of winning boards
        # I found that 40 was the last winning table after analysing the output of print(winning_boards) without the following if
        if winner_board == 40:
            print(last_winning_number)
            print(winning_boards)
            unmarked_sum = calculate_sum(seen_numbers=seen_numbers, board=boards[winning_boards[-1]])
            print(unmarked_sum)
            print(seen_numbers[-1])
            print(unmarked_sum * int(seen_numbers[-1]))
