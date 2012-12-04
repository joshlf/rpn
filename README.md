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
###binary operators
<b>\+</b> addition  
<b>\-</b> subtraction  
<b>\*</b> multiplication  
<b>/</b> division  
<b>|</b> bitwise or  
<b>&</b> bitwise and  

###unary operators
<b>c</b> negation  
<b>~</b> bitwise not  
<b>dup</b> duplicate the value on the top of the stack  
(ie, 1 d ==> 1 1)  
<b>pop</b> pop the top value off of the stack and discard it  
<b>swap</b> swap the top two values on the stack
<b>zero</b> pop and discard all values on the stack    
<b>print</b> print the top value on the stack  

###other
<b>quit</b> quit

theory
======

This rpn calculator is implemented with no explicit stack; it uses only recursion. Creating a stack of numeric literals is not so challenging: simply store a pushed value as a local variable and then recurse. However, dealing with operators is slightly trickier. The structure of the program is as follows:  

There is a single "run" function which is called recursively. Each instance of this function represents a single value or operator on the stack. Each instance keeps a local variable which stores the value of its corresponding item on the stack.  

When run is called, it first waits for input. If the input is a numeric literal, it simply stores the literal in its local variable (ie, pushes it onto the stack), and recurses. If its input is an operator, it <b>returns the operator</b>. The instance of the function to which it returns takes this operator and applies it to its local variable. If the operator was a unary operator, this results in another value which can be saved to a local variable (ie, pushed onto the stack). If the operator was a binary operator, its application is curried, and results in a unary operator which, again, is returned to a previous instance of the function.  

Once a recursive call to run has returned, the current instance performs whatever action is dictated (ie, apply a binary operator to the current stack value and return, or apply a unary operator to the current stack value and store it locally), if the instance doesn't return, it loops. Since it has just pushed a value onto the stack, it is now in the same effective state as it would be had it just pushed a user-entered value on the stack. Thus, it loops, and, at the beginning of the loop, recurses.  

Certain operators require a slightly more powerful framework. Specifically, duplicate and swap require that more than one value be pushed onto the stack. Thus, the run function also takes two inputs - a boolean and an integer. The boolean represents whether or not to treat the integer as its input (instead of waiting for input from the user). If this is the case, it pushes that value that it's been passed instead of waiting for user input, and then continues as it normally would.


