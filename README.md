#  Interpreter In GO

* This project reflects the study of the book "Interpreter Go" written by Thorsten Ball. 

* Using the above as a reference, we have created an interpreted toy programming language called Kisumu using GO and incorporated data structures such as (Integer and float), array and boolean. 

# Language Used - Monkey 

"Without an interpreter or a compiler, a programming language is nothing but an idea"

## Features of monkey language 

 * C - like syntax
 * variable bindings 
 * Integers and booleans 
 * arithmetics expressions 
 * built-in functions extended to Arrays and Strings 
 * first-class and higher-order functions 
 * closures 
 * a string data structures 
 * an array  of data structures 
 * hash map data structures

## Here is how  a monkey  looks  like 

 ```
   // variables 
   let version = 1 + (50 / 2) - (8 * 3);
   
   let name = "zone01kisumu";

   let isMonkeySimple = true; 
```

```
//Arrays and hash maps 

Let  people = [{"name": "Anna", "age":24}, {"name": "Bob", "age": 29}]
```

```
 // or something more complex  
 let twice  = fn(f,x) {
  return f(f(x));
 } 
```
## Parts of an interpreter 
* the lexer 
* the  parser 
* Abstract Syntax Tree 
* the internal object system 
* the evaluator 

## the lexer 
- converts a sequence of characters into a sequence of tokens.
```
let five = 5 ;
let ten  = 10;
let add = fn(x, y) {
  x + y;
};
let result = add(five,ten)
```
 From the above,  we can identify tokens like integers, a keyword or even variable names. We'll distinguish from types, keywords and identifiers 

We can specify this in our code by using constants. "ILLEGAL" will denote something we are not expecting and "EOF" will denote the end of our reading process.

The lexer will ignore the spaces since the monkey does not consider them. Also,  _  is supported as part of variable names.

## the Parser 
The parser's role is to ensure that the source code is syntactically correct, meaning it adheres to the rules and structure of the language in which it is written. If the source code does not follow these rules it produces an error message and compilation stops.


## Abstract Syntax Tree 
This represents the syntactic structure of a program. Nodes on the AST represent elements such as statements or expressions. 
For this project, this data structure will be composed of nodes composed of each other.
Represents the syntax almost exactly and omits irrelevant details like whitespaces or so.

![alt text](<Screenshot from 2024-09-23 20-47-59.png>)

The image above is a  representation of an expression ((1 + 2) + 3) 

## Parsing Expressions 
Parsing statements is fairly straightforward. Reading from left to right identify which keyword and from there parse the rest of the statement.
Parsing expressions are more bit complex. One of the problems in building a parser is the operator precedence. The fact that expressions can be found in many different situations is also a problem that we need to take care of by applying correct parsing procedures that are understandable and extensible from the beginning.

### Parser structure insight 

The following struct represents the concept of a parser inside this interpreter.

```
type  Parser struct  {

  l       *lexer.Lexer
  errors   []string 
  curToken  token.Token 
  peekToken  token.Token  

  [...]      
}
```
The presence of a lexer is pretty clear to convert a sequence of source code into a sequence of tokens.
The presence of token fields allows our parser to operate like an iterator. The first points to the current token being looked at and the second allows us to make decisions based on what comes next.

![alt text](<Screenshot from 2024-09-23 22-01-47.png>)

Not only but also because of this feature we can understand what kind of statements are we reading and parse in the best way possible.

## Usage 
To use this interpreted kisumu programming language, clone the repository 

```bash 
cd kisumu_kisumu_programming_language 

go run kisumu/main.go 
```
You can input a statement or an expression on the prompt and it will have interpreted.

## Contributors 
* [somotto](https://github.com/somotto)      
* [a-j-sheilla](https://github.com/a-j-sheilla) 
* [Hilary505](https://github.com/Hilary505)
