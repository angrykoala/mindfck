0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765 10946

Running fib(21) example:

Original:
Code Size: 10980 characters
Memory: [0 22 0 22 69 47 111 241 69 47 0 22 0 22 0 0 1 0 0 0 0 0 0 0 9 0 0 0 0 214 0 0]
Memory Size: 32 bytes
Executed Instrunctions: 702338727

With Dirty variables
Code Size: 10946 characters
Memory: [0 22 0 22 69 47 111 241 69 47 0 22 0 22 0 0 1 0 0 0 0 0 0 0 9 0 0 0 0 214 0 0]
Memory Size: 32 bytes
Executed Instructions: 702338579

// Single pass on ><
Code Size: 10920 characters
Memory: [0 22 0 22 69 47 111 241 69 47 0 22 0 22 0 0 1 0 0 0 0 0 0 0 9 0 0 0 0 214 0 0]
Memory Size: 32 bytes
Executed Instructions: 702338091

// Multipass
Code Size: 10622 characters
Memory: [0 22 0 22 69 47 111 241 69 47 0 22 0 22 0 0 1 0 0 0 0 0 0 0 9 0 0 0 0 214 0 0]
Memory Size: 32 bytes
Executed Instructions: 702329873

// [-]][-] -> [-]]
Code Size: 10568 characters
Memory: [0 22 0 22 69 47 111 241 69 47 0 22 0 22 0 0 1 0 0 0 0 0 0 0 9 0 0 0 0 214 0 0]
Memory Size: 32 bytes
Executed Instructions: 702238959

// All optimised, no dirty
Code Size: 10592 characters
Memory: [0 22 0 22 69 47 111 241 69 47 0 22 0 22 0 0 1 0 0 0 0 0 0 0 9 0 0 0 0 214 0 0]
Memory Size: 32 bytes
Executed Instructions: 702239055

// Optimise incInt decInt by changing equals zero with not
Code Size: 10345 characters
Memory: [0 22 0 22 69 47 111 241 69 47 0 22 0 22 0 0 1 0 0 0 0 0 0 0 9 0 0 0 0 214 0 0]
Memory Size: 32 bytes
Executed Instructions: 622235940
