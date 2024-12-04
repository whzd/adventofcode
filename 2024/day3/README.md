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

1. Use regular expression to extract all the mul matching patterns and their index.
2. Use regular expression to extract all the instructions and their index.
3. Create a ordered map with all the patterns and their starting index.
4. The "do" and "dont" instructions control a flag that will enable or disable operations on the mul matching patterns.
5. Calculate the multiplications when the flag is true.
