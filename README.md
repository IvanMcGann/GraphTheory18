Graph Theory Project 2018

Student Name: Ivan McGann
Student No: G00340138

Intro: 
You must write a program in the Go programming language that can
build a non-deterministic finite automaton (NFA) from a regular expression,
and can use the NFA to check if the regular expression matches any given
string of text. You must write the program from scratch and cannot use the
regexp package from the Go standard library nor any other external library.
A regular expression is a string containing a series of characters, some
of which may have a special meaning. 

You are expected to be able to break this project into a number of smaller
tasks that are easier to solve, and to plug these together after they have been
completed. You might do that for this project as follows:
1. Parse the regular expression from infix to postfix notation.
2. Build a series of small NFA’s for parts of the regular expression.
3. Use the smaller NFA’s to create the overall NFA.
4. Implement the matching algorithm using the NFA.


User Manual:


Links used so far:

http://www.rexegg.com/regex-quickstart.html

https://automatetheboringstuff.com/chapter7/

https://golang.org/pkg/regexp/

https://shapeshed.com/golang-regexp/#how-to-find-a-string-submatch-index