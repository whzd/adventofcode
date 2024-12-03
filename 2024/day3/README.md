# Day 3: Mull It Over

Challenge Page: [Link](https://adventofcode.com/2024/day/3)

## Part 1

...

### Approach

1. Use a regular expression to extract all the matching patterns.
2. On the found patterns use a regular expression to extract the numbers.
3. Convert the numbers from string to int.
4. Calculate the sum of all the converted numbers.

## Part 2

...

### Approach

1. Use regular expression to extract all the mul matching patterns and their starting index.
2. Use regular expression to extract all the instructions and their starting index.
3. Convert the numbers of the mul patterns that have index greater than the "do" instruction or lesser than the first "don't" instruction.
4. Calculate the sum of all the converted numbers.
