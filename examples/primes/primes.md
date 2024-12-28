# Primes

## Sieve of Eratosthenes

Sieve of Eratosthenes will be the resulting algorithm. This exercise is an example of applications of rules 4 and 5 repeatedly.

The flow will be:

- test factors of 1, generalize to return empty list
- test factors of 2, generalize by putting list into variable, add if 
- test factors of 3, generalize by putting n into list instead of 2
- test factors of 4, generalizing won't work, that's why the rule contains "...where possible.", use modulo, second if to make test pass, refactor and move second if out
- test factors of 5, 6, 7 and they pass
- test factors of 8, generalize by making the if to a while/for
- test factors of 9, generalizing won't work initially, add second while/for covering 3, now you can refactor by generalizing, extract divisor, remove if, clean up loops

### Test Cases

Start with 1..9, then a big one last: 2x2x2x3x11x11x13.

## Runs

1. 16m - 28.12.24, *[x] retained*