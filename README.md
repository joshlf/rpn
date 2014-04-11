<!--
Copyright 2013 The Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
-->

rpn [![Build Status](https://travis-ci.org/joshlf13/rpn.svg?branch=master)](https://travis-ci.org/joshlf13/rpn)
===

An rpn calculator using no explicit stack, only recursion and function-passing (golang)

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
(ie, 1 dup ==> 1 1)  
<b>pop</b> pop the top value off of the stack and discard it  
<b>swap</b> swap the top two values on the stack  
<b>zero</b> pop and discard all values on the stack    
<b>print</b> print the top value on the stack  

###other
<b>quit</b> quit
