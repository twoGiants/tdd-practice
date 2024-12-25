# Stack

## FILO

Implement a stack of integers. It should support:

- pushing integers to it,
- popping integers from it in the FILO style,
- checking the size,
- checking if its empty,
- dealing with an underflow,
- dealing with an overflow.

### Test Cases

We can derive the following test cases from the logic above:

- can create stack
- after one push is not empty
- after one push and pop is empty
- after two pushes size is two
- popping empty stack panics with underflow
- pushing full stack panics with overflow
- pushing X will pop X
- pushing X and Y will pop Y and then X

## FIFO

Same as above but with a firs-in-first-out queue. Use a fixed size array. Use two pointers to track where the elements are to be added and removed.

## Runs

1. 32m - 30.11.24
1. 64m - 25.12.24, *[x] retained; used different test names*
1. 29m - 25.12.24, *[x] retained; used same test names*