# Conceptual Interview Problems in Golang
This list was taken from a post [found on reddit](https://www.reddit.com/r/cscareerquestions/comments/20ahfq/heres_a_pretty_big_list_of_programming_interview/?st=j8dc7k2e&sh=11599122)

Each section (General/Strings/Trees...) is under its own subdirectory.

Each problem can be run with `go run <filename>`. The problem is always solved
as a function, and the test data is inside the problem's `main()` function.

### General
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/mostfreqint.go) - Find the most frequent integer in an array
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/sum10.go) - Find pairs in an integer array whose sum is equal to 10 (bonus: do it in linear time)
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/rotatedarrays.go) - Given 2 integer arrays, determine of the 2nd array is a rotated version of the 1st array. Ex. Original Array A={1,2,3,5,6,7,8} Rotated Array B={5,6,7,8,1,2,3}
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/fibonacci.go) - Write fibonacci iteratively and recursively (bonus: use dynamic programming)
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/singlenum.go) - Find the only element in an array that only occurs once.
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/commonelements.go) - Find the common elements of 2 int arrays
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/binarysearch.go) - Implement binary search of a sorted array of integers
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/binarysearch-rotated.go) - Implement binary search in a rotated array (ex. {5,6,7,8,1,2,3})
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/primesievedynamic.go) - Use dynamic programming to find the first X prime numbers
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/printbinaryint.go) - Write a function that prints out the binary form of an int
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/parseint.go) - Implement parseInt
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/sqrt.go) - Implement squareroot function
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/exponent.go) - Implement an exponent function (bonus: now try in log(n) time)
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/multiply.go) - Write a multiply function that multiples 2 integers without using *
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/rand7.go) - HARD: Given a function rand5() that returns a random int between 0 and 5, implement rand7()
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/General/islandclusters.go) - HARD: Given a 2D array of 1s and 0s, count the number of "islands of 1s" (e.g. groups of connecting 1s)

### Strings
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/Strings/firstnon-repeated.go) - Find the first non-repeated character in a String
 - [DONE](https://github.com/devedge/GolangPractice/blob/master/ConceptualInterviewProblems/Strings/reverse-string.go) - Reverse a String iteratively and recursively
 - DONE - Determine if 2 Strings are anagrams
 - DONE - Check if String is a palindrome
 - DONE - Check if a String is composed of all unique characters
 - Determine if a String is an int or a double
 - HARD: Find the shortest palindrome in a String
 - HARD: Print all permutations of a String
 - HARD: Given a single-line text String and a maximum width value, write the function 'String justify(String text, int maxWidth)' that formats the input text using full-justification, i.e., extra spaces on each line are equally distributed between the words; the first word on each line is flushed left and the last word on each line is flushed right
