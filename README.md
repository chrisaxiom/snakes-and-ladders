# snakes-and-ladders
A solution to the snakes and ladders problem


## Problem

Given a grid for snakes and ladders, what is the minimum number of throws to reach the end? Die is 6-sided. Player starts on tile 0 (not displayed in picture). Landing on the bottom of the a ladder results in the position being at the top of the ladder after the throw. Landing at the head of a snake results in the position being at the bottom of the snake after the throw.

For a grid like this:

![snakes-and-ladders](https://cloud.githubusercontent.com/assets/3770894/25108830/6845b684-238d-11e7-80dc-cf3dfa240c68.jpg)

The input will be a json:
```
{
  "ladders": [
    [1,38],
    [4,14],
    [9,31],
    [28,84],
    [21,42],
    [51,67],
    [72,91],
    [80,99]
  ],
  "snakes": [
    [17,7],
    [54,34],
    [62,19],
    [64,60],
    [87,36],
    [92,73],
    [95,75],
    [98,79]
  ]
}
```

The output will be a single integer denoting the number of rolls required to reach tile 100.