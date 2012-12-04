rpn
===

An rpn calculator using no explicit stack, only recursion (golang)

usage
=====

Values are entered in reverse polish notation (rpn) order. Values entered are stored on a stack. When an operator is entered, it reaches backwards into the stack to find its arguments, and pushes its result onto the stack. For example, to add 1 and 2, first push 1 onto the stack, then push 2, then push the "+" operator, which takes the previous two values on the stack, 1 and 2, as its arguments:  
1 2 +  
will result in the stack:  
3  

syntax
======
<b>binary operators</b>  
\+ addition  
\- subtraction  
\* multiplication  
/ division  
| bitwise or  
& bitwise and  

<b>unary operators</b>  
c negation  
~ bitwise not  
dup duplicate the value on the top of the stack  
(ie, 1 d ==> 1 1)  
pop pop the top value off of the stack and discard it
print print the top value on the stack  

<b>other</b>  
quit

theory
======

This rpn calculator is implemented with no explicit stack; it uses only recursion. Creating a stack of numeric literals is not so challenging: simply store a pushed value as a local variable and then recurse. However, dealing with operators is slightly trickier. The structure of the program is as follows:  

There is a single "run" function which is called recursively. Each instance of this function represents a single value or operator on the stack. Each instance keeps a local variable which stores the value of its corresponding item on the stack.  

When run is called, it first waits for input. If the input is a numeric literal, it simply stores the literal in its local variable (ie, pushes it onto the stack), and recurses. If its input is an operator, it <b>returns the operator</b>. The instance of the function to which it returns takes this operator and applies it to its local variable. If the operator was a unary operator, this results in another value which can be saved to a local variable (ie, pushed onto the stack). If the operator was a binary operator, its application is curried, and results in a unary operator which, again, is returned to a previous instance of the function.  

Once a recursive call to run has returned, the current instance performs whatever action is dictated (ie, apply a binary operator to the current stack value and return, or apply a unary operator to the current stack value and store it locally), if the instance doesn't return, it loops. Since it has just pushed a value onto the stack, it is now in the same effective state as it would be had it just pushed a user-entered value on the stack. Thus, it loops, and, at the beginning of the loop, recurses.


