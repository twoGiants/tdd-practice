# TDD Practice

[Go Refresher.](/docs/go-refresher.md)

## Rules

Taken from *Clean Craftsmanship by Robert C. Martin*.

**`Rule 1:`** *Write the test that forces you to write the code you already know you want to write. (p. 37)*

**`Rule 2:`** *Make it fail. Make it pass. Clean it up. (p. 40)*

**`Rule 3:`** *Don't go for gold. (p. 50)*
  
**`Rule 4:`** *Write the simplest, most specific, most degenerate test that will fail. (p. 53)*

**`Rule 5:`** *Generalize where possible. (p. 55)*

**Generalization Mantra:** *As the tests get more specific, the code gets more generic.*
- go from specific like `[]int{}` to general like `factors := []int{}`
- `if` is the degenerate form of `while/for` OR `while/for` is the general form of `if`

## Examples

1. [Stack](examples/storage/stack.md).
2. [Primes](examples/primes/primes.md).