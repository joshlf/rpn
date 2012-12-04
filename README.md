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
d duplicate the value on the top of the stack  
(ie, 1 d ==> 1 1)  
p print the top value on the stack  

<b>other</b>  
q quit  

theory
======

This rpn calculator is implemented with no explicit stack; it uses only recursion. 