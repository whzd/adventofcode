# Day 2: Red-Nosed Reports

Challenge Page: [Link](https://adventofcode.com/2024/day/2)

## Part 1

For the first part of the challenge we are required to read a sequence of numbers, separated by spaces, from each line of the input file.
We then need to classify each sequence of numbers as safe or unsafe.
In order for a sequence to be safe, all their number need to be increasing or decreasing and two adjacent numbers can only vary by at least one and at most 3.
All sequence that don't meet this standard are considered unsafe.

### Approach

1. Read the file, and for each line, create a list of numbers. Each of this lists will be added to a main list.
2. Loop trough each list on the main list and classify if that sequence is safe or unsafe.
3. Increment a variable for each safe sequence found.

For this approach the time complexity is defined by looping through the main list and all the numbers of each of the lists inside.
Therefore the time complexity is 0(n^2).

## Part 2

For the second part we are required to add a toleration level of 1 to the previous algorithm.
This means that it requires two infringements on a sequence in order to be considered unsafe.

### Approach

1. Take the previous algorithm and create a variable as fault tol

