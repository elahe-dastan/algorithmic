https://www.hackerrank.com/certificates/f788bdf183c4
## Introduction
In one of my interviews I was asked to answer to a question from [hacker rank](https://www.hackerrank.com/) since then <br/>
I decided to solve more and more algorithmic questions cause it's so important in order to join a good company.<br/>
In this repository I solve algorithmic questions from hackerrank and save the here<Br/>

## FizzBuzz
Given a number n, for each integer i in the range from 1 to n inclusive, print one value per line as follows<br/>
if i is a multiple of both 3 and 5, print FizzBuzz<br/>
if i is a multiple of 3(but not 5), print Fizz<br/>
if i is a multiple of 5(but not 3), print Buzz<Br/>
if i is not a multiple of 3 or 5,print the value of i.

## VowelSubstring
Given a string and a number 'k', we should find the substring with length k that has the most number of vowels if there<br/>
are no vowels in the string we should write 'Not found!'

## StringAnagram
Given a list of queries and a list of words, for each query we should return how many anagrams it has in the list of<br/>
words

[basic certificate](https://www.hackerrank.com/certificates/f788bdf183c4)

## ValidSudoku
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:<br/>

Each row must contain the digits 1-9 without repetition.<br/>
Each column must contain the digits 1-9 without repetition.<br/>
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

## 3Sum
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets<br/>
in the array which gives the sum of zero.

## Leader Board
Alice is playing an arcade game and wants to climb to the top of the leaderboard and wants to track her ranking. The game
uses Dense Ranking.<br/>
It should return an integer array where each element res[j] represents Alice's rank after the jth game.<br/>

climbingLeaderboard has the following parameter(s):<br/>

scores: an array of integers that represent leaderboard scores<br/>
alice: an array of integers that represent Alice's scores
## Magic Square
You will be given a 3*3 matrix (s) of integers in the inclusive range [1, 9]. We can convert any digit a to any other digit
 b in the range [1, 9] at cost of |a - b|. Given s, convert it into a magic square at minimal cost. Print this cost on a new line.
 
## Encryption
An English text needs to be encrypted using the following encryption scheme.
First, the spaces are removed from the text. Let L be the length of this text.
Then, characters are written into a grid, whose rows and columns have the following constraints:<br/>

floor(sqrt(L)) <= row <= column <= ceil(sqrt(L))<br/>

Ensure that rows * columns <= L
If multiple grids satisfy the above conditions, choose the one with the minimum area, i.e. rows * columns.
The encoded message is obtained by displaying the characters in a column, inserting a space, and then displaying the next 
column and inserting a space, and so on.

## Bomberman
Each bomb can be planted in any cell of the grid but once planted, it will detonate after exactly 3 seconds. Once a bomb 
detonates, it's destroyed — along with anything in its four neighboring cells. 
If there is a bomb in a neighboring cell, the neighboring bomb is destroyed without detonating, so there's no chain reaction.

Bomberman is immune to bombs, so he can move freely throughout the grid. Here's what he does:

Initially, Bomberman arbitrarily plants bombs in some of the cells, the initial state.<br/>
After one second, Bomberman does nothing.<br/>
After one more second, Bomberman plants bombs in all cells without bombs, thus filling the whole grid with bombs. No bombs detonate at this point.<br/>
After one more second, any bombs planted exactly three seconds ago will detonate. Here, Bomberman stands back and observes.<br/>
Bomberman then repeats steps 3 and 4 indefinitely.