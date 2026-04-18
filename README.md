# Jedreks — interpreter written in Go

An interpreter for the **Jedreks** programming language, written in Go. Jedreks is a dynamically-typed language with Polish keywords, inspired by the book *"Writing an Interpreter in Go"*. It implements a full pipeline: Lexer → Parser → Evaluator.

---

## Getting started

```bash
go run main.go
```

Opens an interactive REPL (`>>`) that accepts code line by line.

---

## Architecture

```
Source code
    │
    ▼  lexer/lexer.go
Tokens
    │
    ▼  parser/parser.go  (Pratt parser)
AST (Abstract Syntax Tree)
    │
    ▼  evaluator/evaluator.go  (tree-walking)
Result
```

### Lexer
Converts source text into a stream of tokens. Uses two pointers (`position`, `readPosition`) for single-character lookahead. Supports ASCII — identifiers, integers, strings, and single- and double-character operators (`==`, `!=`).

### Parser
Pratt (top-down operator precedence) parser. Handles operator precedence by registering `prefixParseFn` and `infixParseFn` callbacks per token type. Collects errors without stopping.

Precedence order (lowest to highest): `LOWEST → EQUALS → LESSGREATER → SUM → PRODUCT → PREFIX → CALL → INDEX`

### Evaluator
Tree-walking interpreter — walks AST nodes and executes instructions directly. Uses singleton objects for `TRUE`, `FALSE`, and `NULL`. Errors are first-class objects that propagate up the call tree.

---

## Language syntax

### Variables

```
zmienna x = 5;
zmienna text = "hello";
```

### Functions and closures

```
zmienna add = metoda(a, b) { zwracam a + b; };
add(3, 4);  // 7

zmienna newAdder = metoda(x) { metoda(n) { x + n } };
zmienna addTwo = newAdder(2);
addTwo(3);  // 5
```

### Conditionals

```
gdyby (x > 3) { pokaz(prawda); } inaczej { pokaz(potwarz); }
```

`gdyby` is an expression — it can appear on the right-hand side of an assignment.

### Arrays

```
zmienna arr = [1, "two", metoda(x) { x * x }];
arr[0];            // 1
pierwszy(arr);     // 1
ostatni(arr);      // <function>
reszta(arr);       // ["two", <function>]
wepchnij(arr, 4);  // [1, "two", <function>, 4]  — returns a new array, original unchanged
```

### Hashes

```
zmienna h = {"name": "Jimmy", "age": 16, 99: "birth year"};
h["name"];      // Jimmy
h[100 - 1];     // birth year
```

Keys can be integers, strings, or booleans.

---

## Keywords

| Jedreks    | English equivalent |
|------------|--------------------|
| `zmienna`  | `let`              |
| `metoda`   | `fn`               |
| `zwracam`  | `return`           |
| `gdyby`    | `if`               |
| `inaczej`  | `else`             |
| `prawda`   | `true`             |
| `potwarz`  | `false`            |

---

## Built-in functions

| Function           | Description                                      |
|--------------------|--------------------------------------------------|
| `pokaz(...)`       | Prints arguments to stdout                       |
| `dlug(x)`          | Length of a string or array                      |
| `pierwszy(arr)`    | First element of an array, or null if empty      |
| `ostatni(arr)`     | Last element of an array, or null if empty       |
| `reszta(arr)`      | New array without the first element              |
| `wepchnij(arr, x)` | New array with `x` appended                      |

---

## Data types

| Type     | Example                         |
|----------|---------------------------------|
| Integer  | `42`, `-7`                      |
| String   | `"hello"`                       |
| Boolean  | `prawda`, `potwarz`             |
| Null     | result of an unmatched `gdyby`  |
| Array    | `[1, 2, 3]`                     |
| Hash     | `{"key": "value"}`              |
| Function | `metoda(x) { zwracam x * 2; }` |

---

## Object system and environment

All values implement the `Object` interface (`Type()`, `Inspect()`). Variables are stored in an `Environment` — a `string → Object` map with an `outer` pointer that enables closures. When a function is called, a new environment is created that extends the function's defining scope.

---

## Project structure

```
main.go                   — entry point, starts REPL
token/token.go            — token type definitions and keyword map
lexer/lexer.go            — tokenizer
ast/ast.go                — AST node types
parser/parser.go          — Pratt parser
object/object.go          — object type system
object/environment.go     — environment / closures
evaluator/evaluator.go    — evaluator
evaluator/builtins.go     — built-in functions
repl/repl.go              — REPL loop
```

Implementation notes: [knowledge.md](knowledge.md)
