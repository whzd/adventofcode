with open('input') as f:
    content = f.readlines()

# content = ['forward 5', 'down 5', 'forward 8', 'up 3', 'down 8', 'forward 2']

pos = 0
depth = 0
aim = 0
for move in content:
    tmp_move = move.split(' ')
    if tmp_move[0] == 'forward':
        pos += eval(tmp_move[1])
        depth += aim * eval(tmp_move[1])
    if tmp_move[0] == 'down':
        aim += eval(tmp_move[1])
    if tmp_move[0] == 'up':
        aim -= eval(tmp_move[1])

print(pos * depth)
