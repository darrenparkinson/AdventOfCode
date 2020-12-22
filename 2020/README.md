# 2020

Notes from 2020 AOC

## Day12

* I was able to reuse the manhattan distance calculation I'd already created for 2019 Day 3.
* I prefer [this implementation](https://github.com/AnthonyNixon/AdventOfCode/blob/main/2020/day12/main.go)
* I should probably make use of standard library features, like image.Point [like this did](https://github.com/mnml/aoc/blob/master/2020/12/2.go)

## Day14

* Discovered the use of [`fmt.Sscanf`](https://golang.org/pkg/fmt/#Sscanf) which allows you to store values into variables given a specific format
    * e.g. `fmt.Sscanf(s, "mem[%d] = %d", &loc, &val)` puts the values 8 and 11 into loc and val respectively given an input of `mem[8] = 11`