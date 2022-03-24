Interesting about Go:

- standard built-in syntax checking and style format: no endless complains about tab versus spaces (no unused variables allowed, using camelCase and warn about snakecase)
- priority is clarity of intent and readability, so there are explicit type conversions instead of automatic
- Go is a call by value language: if you pass a variable for a parameter to a function, Go will always make a copy, even with struct
- len, append, make, copy are standard build-in functions
- no standard set
- intuitive use of Map -> JSON like, very similar like a slice
- try avoiding shadowing at all times!
- avoiding indentation and deep nesting. Try {} finally {} will be replaced by defered in the middle of the function.
- the stack in Go can dynamically grow. Optimize the programs to allocate as much as possible on the stack.