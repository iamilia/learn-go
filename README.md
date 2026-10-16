# first_pr — Learning Go 🐹

My first steps with Go. I already work with TypeScript, Node.js and Lua, and this repo tracks what I learn as I go, one lesson at a time.

## Lessons

| # | Topic | What I practiced |
|---|-------|------------------|
| 01 | [Variables & types](01-variables/main.go) | `:=`, `var`, `const`, basic types, zero values, `fmt.Printf` |
| 02 | [Control flow](02-control-flow/main.go) | `if`, `switch` without a condition, `for` (Go's only loop), `&&` / `\|\|` — FizzBuzz |
| 03 | [Functions](03-functions/main.go) | parameters, multiple return values, named returns, `strconv.Itoa` |
| 04 | [Slices & maps](04-slices-maps/main.go) | `[]int`, `append`, `range`, `map[string]int`, the zero-value counting trick |
| 05 | [Structs & methods](05-structs/main.go) | `type ... struct`, methods, pointer vs value receivers, exported vs unexported names, early return |
| 06 | [Pointers](06-pointers/main.go) | `&` / `*`, passing by address, `swap`, constructor functions, `[]*Player`, `nil` |

## Run a lesson

```bash
go run ./01-variables
go run ./02-control-flow
go run ./03-functions
go run ./04-slices-maps
go run ./05-structs
go run ./06-pointers
```

## Things that surprised me coming from TS / Lua

- No implicit type conversion: `int + float64` is a compile error, so you write `float64(x)`.
- Unused variables and imports don't compile.
- `for` is the only loop; there's no `while`.
- `switch` cases don't fall through, so there's no `break`.
- `&&` and `||` always return `bool`, so the Lua `x or "default"` trick doesn't work.
- Reading a missing map key returns the zero value; use `v, ok := m[k]` to check it.
- `string(65)` is `"A"`, not `"65"`. Use `strconv.Itoa`.
- Go passes everything by value, so methods that modify a struct need a pointer receiver (`*Player`).
- No `public` / `private` keywords: a capital first letter means exported.
- `for _, p := range []Player` gives you copies; use `[]*Player` when the loop must change each item.
