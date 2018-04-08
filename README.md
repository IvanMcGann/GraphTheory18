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
1. Compile the go file (finalProject.go) by opening the command line of choice and once in the root folder type go build. Enter go build to do this.
2. Once completed run the .exe file here named GraphTheory18.exe 
3. Next you will find the menu enter options 1 or 2 to run project, option 3 to exit.
4. Once you choose option (1.) enter an infix expression then enter a string to test against it to see if it matches.
5. Once entered you will get a true or false statement. 
6. If you choose option (2.) enter a postfix expression to to test, then enter a string to test it against.
7. Once entered you will get a true or false statement.
8. If you enter option (3.) it will close the project.



2
Option 2 Was entered.
Enter postfix expression: abc
Enter a string to test if it matches the nfa:



Links used:

https://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa

http://www.rexegg.com/regex-quickstart.html

https://automatetheboringstuff.com/chapter7/

https://golang.org/pkg/regexp/

https://shapeshed.com/golang-regexp/#how-to-find-a-string-submatch-index