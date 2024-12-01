# Day 1: Historian Hysteria

Challenge Page: [Link](https://adventofcode.com/2024/day/1)

## Part 1

For the first part of this challenge we are required to read two columns of numbers, separated by white space.
We then need to calculate the absolute difference between the numbers from column A and column B in a crescent order.

### Approach

1. Read the file and, for each line, split the number using space as delimiter and store index 0 on a A list and index 1 on a B list.
2. Sort each list.
3. Loop through the lists and increment a variable with the absolute difference between the numbers on the same index on each list.

For this approach the time complexity is defined by the sorting algorithm used.
In this case, the sorting algorithm used was the built in sort of the slices package, which is based on a pattern-defeating quicksort(pdqsort).
This sort has a worst-case runtime of O(nk) for n inputs with k distinct elements, as stated on the [paper](https://arxiv.org/pdf/2106.05123.pdf).

## Part 2

For the second part of this challenge we are required to identify how many times a number from the column A appears on the column B, and use those values to calculate the total similarity score.
A similarity score is calculated by multiplying each number from a list by the number of times it appears on another list.

### Approach

1. Read the file and, for each line, split the number using space as delimiter and store index 0 on a A list and index 1 on a B list.
2. Create a map, where the key is a number and the value is the number of times that number it appeared, by looping through the B list.
3. Loop through the A list and increment a variable with the similarity score of each value. The similarity score here is calculated by multiplying the number from the A list with the the value, of the matching key, from the created map, or 0 if not found.

For this approach the time complexity is difined by the for loops used to iterate over the created lists.
Since there are no nested loops, this time complexity is O(n).
