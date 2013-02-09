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

This rpn calculator is implemented with no explicit stack. It uses recursion and continuation-passing to simulate a stack and acheive communication between the various layers of the stack. This explanation goes into some depth, and doesn't expect the reader to necessarily know about things like continuation passing or currying. For a more brief explanation, see the comments in the code itself.

The basic outline of the program is that each step of the recursion represents a value on the stack. Thus, pushing a number is akin to a recursive call. If we're only dealing with numbers, this is relatively straightforward. The stack "1 2 3" would be represented by a function call for 1, which then calls itself to store 2, which in turn calls itself to store 3 (don't worry yet about what "it" is - just think of it as some abstract function which represents the stack).

However, what happens when we push an operator onto the stack? Think about the following stack in terms of standard rpn lingo:

1 2 +

What happens with this stack in a normal rpn calculator is that the value 1 is pushed onto the stack, then the value 2, and then the operator "+". When "+" is pushed, it pops itself off, then pops 2 and 1 off the stack, adds them, and pushes the result, 3.

So how do we pop off the stack when the stack is actually a sequence of recursive calls? We return. So if we have the stack "1 2 3" and want to pop 3, then the function which stores 3 simply returns, and now 3 is no longer on the stack.

But what about the operator? Let's imagine the scenario we've been using with the stack, "1 2 3". When the function storing 3 performs a recursive call, control will not return to it until that recursive call has returned, which is equivalent to all of the values above it on the stack being popped off. The only scenario in which values can be popped off is if an operator is pushed onto the stack. Thus, if the recursive call returns, this means an operator has been pushed onto the stack. So what does the call return, then? It returns the operator itself.

What do we mean by "it returns the operator"? The recursive call actually returns a function which performs the desired operation. So if the recursive call wants to return the "+" operator, it simply returns a function which performs addition.

So back to our example stack. Let's say, like before, we have the stack "1 2 3", but we then push the "c" operator (which performs negation). The function storing 3 has recursed, and this recursive call now returns "c". So should the function storing 3 do with it? It should call it on itself, since it is the top value on the stack. So it calls "c" on the value 3, which returns -3. Now what should it do? Well, -3 has effectively replaced 3 as the value on the top of the stack, so now it should recurse again, and the process continues happily.

But what if we want to handle an operator which consumes more than one value, like "+"? In this case, we perform something called "currying". To see what we mean, let's take our previous example, except instead of "c", we'll use "+". So the function storing 3 recurses, and that recursion returns the "+" operator. Just like before, we call the "+" operator on 3. Except, this time, it can't return a value, since "3 +" doesn't mean anything. Three plus what? Well, it may not be able to return a value, but it can return a function. Let's call this function "3+". The "3+" function again takes an argument - another number. When "3+" is called on a number, say, n, it adds 3 to n and returns it. So "3+"(n) = 3 + n. What we've done here is transform a function which would normally accept two parameters into a function which accepts the first parameter, and returns a function which accepts the second parameter and <i>then</i> returns the final answer. This is known as "currying".

So how does this help us? Well, when we call "+" on the value 3, getting "3+", we can again return that down the stack. The function storing 2 receives this operator from its recursive call, and it just treats it like a normal operator, applying it to 2, and getting back a value - 5 - which it stores as its new value, and recurses.

So this is all nice, but we've now got a problem. We now have <i>two different</i> types of operators - one which return a number, and one which return a function. Specifically, we've got <b>binary</b> operators and <b>unary</b> operators. Here are their function signatures in Go:

type unary_operator func(int) int
type binary_operator func(int) unary_operator

We could simply have the function return two values - a binary operator and a unary operator - and let the caller decide which one it wants to call. In fact, that's is exactly what was done in a previous version of this code. However, we can alter the function signature of an operator so that we only need one type of function. This may appear a little odd, so just stay with me for a moment:

type operator func(int) operator

What we've made operators into are functions which return other operators. This actually makes pretty good sense for binary operators. After all, binary operators used to return unary operators anyway. But what about for unary operators. They're supposed to return numbers, and now they just return operators. Does that mean that numbers are now considered a type of operator? Actually, yes! In order to see how this works, however, we're going to have to change our structure a little bit.

Instead of one central function which recurses and stores value (what we were calling "it" before), we're now going to break that functionality up so that different functions - operators included - can perform the task of storing one value on the stack and recursively calling a function to grow the stack.

This starts off with the "input" function. When the calculator is first started, "input" is the first function which is called, and it does exactly what you'd expect - it accepts user input. User input comes in two main flavors - values and operators. When "input" receives an operator, it behaves exactly as our "it" function behaved before - it just returns the function which represents that operator. However, when it receives a number, instead of making a recursive call to <i>itself</i>, it makes a recursive call to the "number" function.

The "number" function is the number-as-an-operator function that we were talking about earlier. Just like all operators, it accepts a value as input. However, it does something special with this input - it stores it as its own value. Sound familiar? It's just the behavior of the old "it" function, except instead of asking the user for the number, it gets its number as an argument. Then what it does is something very familiar - it recursively calls input. And what does it do with the result of this recursive call? Well, like before, the recursive call returns an operator, so it simply applies the operator to its value. However, unlike before, since that operator returns another operator rather than a number, instead of storing the result of calling the operator, it returns it.

This is all very confusing, so let's look at an example to see how this works.

Let's look at the stack "1 2 + pop". This stack should result in the empty stack. We'll use "func1 -> func2 -> func3" to mean that func1 calls func2, which calls func3, and "func1 -> func2 <- 3" to mean that func3 then returned the value 3 to func2.

1:

	"input" is called. The user inputs the number 1. "input" calls "number(1)":

		input -> number(1)


2:

	"number(1)" calls "input":

		input -> number(1) -> input


3:

	The user inputs the number 2. The current call to "input" calls "number(2)":

		input -> number(1) -> input -> number(2)


4:

	"number(2)" calls "input":

		input -> number(1) -> input -> number(2) -> input


5:

	The user inputs "+". The current call to "input" returns "+":

		input -> number(1) -> input -> number(2) <- "+"

		
6:
	"number(2)" applies the "+" operator to its value, 2:
		input -> number(1) -> input -> number(2) -> +(2)
7:
	"+(2)" returns the operator "2+":
		input -> number(1) -> input -> number(2) <- "2+"
8:
	"number(2)" also returns "2+":
		input -> number(1) -> input <- "2+"
9:
	"input" also returns "2+":
		input -> number(1) <- "2+"
10:
	"number(1)" applies the "2+" operator to its value, 1:
		input -> number(1) -> 2+(1)
11:
	"2+(1)" returns the number 3 (which is an operator!). In doing this, though, it has to first call "number(3)"
		input -> number(1) -> 2+(1) -> number(3)
12:
	"number(3)" calls input:
		input -> number(1) -> 2+(1) -> number(3) -> input
13:
	The user inputs "pop", which the current call to "input" returns:
		input -> number(1) -> 2+(1) -> number(3) <- "pop"
14:
	"number(3)" calls "pop" on its value, 3:
		input -> number(1) -> 2+(1) -> number(3) -> pop(3)
15:
	"pop(3)" returns another operator. This operator is the identity operator, "id" (the idea here is to take two arguments and only return the second one. That is, popping a value off the stack is equivalent to popping two values and then pushing back the second one):
		input -> number(1) -> 2+(1) -> number(3) <- "id"
16:
	"number(3)" also returns "id":
		input -> number(1) -> 2+(1) <- "id"
17:
	"2+(1)" also returns "id":
		input -> number(1) <- "id"
18:
	"number(1)" also returns "id":
		input <- "id"
19:
	"input" also returns "id":
		<- "id"
20:
	The bottom of the stack recieves "id", doesn't know what to do with it, and so it just starts over again by calling input:
		input

Hopefully that gave you a sense of how this whole business of numbers as operators works. Reading the code should help your understanding, too. And if you have any questions, the internet is your friend. Look up "currying" and "continuation passing", since the internet is bound to have far better explanations of these than I am able to deliver.
