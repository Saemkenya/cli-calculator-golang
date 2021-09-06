# Command Line Interface calculator implementation with the Go-lang flag package

## The following commands should work outta the box:

### To show available subcommands
./cli-calculator-golang



### To add two numbers
./cli-calculator-golang add -num1=100 -num2=87

### To subtract two  numbers
./cli-calculator-golang sub -num1=100 -num2=87

### To multiply two  numbers
./cli-calculator-golang mult -num1=100 -num2=87

### To divide two  numbers
./cli-calculator-golang div -num1=100 -num2=87

### To get the square root of a number
./cli-calculator-golang sqr -num1=100 -num2=87



### To add two numbers and show formula info
./cli-calculator-golang add -num1=100 -num2=87 -show

### To subtract two  numbers and show formula info
./cli-calculator-golang sub -num1=100 -num2=87 -show

### To multiply two  numbers and show formula info
./cli-calculator-golang mult -num1=100 -num2=87 -show

### To divide two  numbers and show formula info
./cli-calculator-golang div -num1=100 -num2=87 -show

### To get the square root of a number and show formula info
./cli-calculator-golang sqr -num1=100 -num2=87 -show


## ToDo:
* Implement edge test cases for all formulars
* A regex router to offer calc as REST API

