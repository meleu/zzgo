# Learning Notes

My notes about new things I'm learning while build this project.

Although public, it's very personal (which means "I'm not taking notes of things I already know by heart").

## `zz random`

### Go questions

I need to research about this:

> Error wrapping with `%w`: `fmt.Errorf("...: %w", err)` preserves the original error for `errors.Is`/`errors.As`. It's acceptable for user facing messages, but loses introspection.

### Cobra basics

- A command is a `&cobra.Command{...}` struct literal.
- Relevant fields:
  - `Use: "random [number1 [number2]]` - the usage line. We must specify the (non-flag) args explicitly (it doesn't handle it automatically, like Bashly does)
  - `Short` / `Long`: help text
  - `Args`: a **validator function** that runs before `Run`/`RunE`. Cobra aborts if validation fails. Examples:
    - `cobra.NoArgs`
    - `.ExactArgs(n)`
    - `.MinimumNArgs(n)` / `.MaximumNArgs(n)`
    - `.RangeArgs(min, max)`
    - `.MatchAll(v1, v2, ...)` to compose multiple validations
  - `Run` vs. `RunE`:
    - `RunE` returns an error; Cobra prints it and exit with non-zero (prefer `RunE`).
    - `Run` the developer is responsible to explicitly handle error cases/messages.
  - `SilenceUsage: true` - suppress usage message on `RunE` error.

Cobra testing trick: use `fmt.Fprintln(cmd.OutOrStdout(), "something")` to print something to stdout.
The reason for this is to allow `rootCmd.SetOut(io.Discard)` in the tests.

### Testing Cobra commands

Run the tests with `go test ./cmd/...`

Test files lives in the same package so it can reference `rootCmd`, then drive the tests like this:

```go
func TestZZRandom_ArgsValidation(t *testing.T) {
 argTests := []struct {
  name    string
  args    []string
  wantErr bool
 }{
  { "no args ok", []string{"random"}, false, },
  { "one arg ok", []string{"random", "10"}, false, },
  { "two args ok", []string{"random", "10", "1"}, false, },
  { "three args fails", []string{"random", "1", "5", "10"}, true, },
  // ...
 }

 for _, tt := range argTests {
  t.Run(tt.name, func(t *testing.T) {
   rootCmd.SetArgs(tt.args)
   rootCmd.SetOut(io.Discard) // keep test output clean
   rootCmd.SetErr(io.Discard) // idem
   gotErr := rootCmd.Execute()

   if (gotErr != nil) != tt.wantErr {
    t.Errorf( "test: %q, args: %v, wantErr: %v, gotErr: %v", tt.name, tt.args, tt.wantErr, gotErr)
   }
  })
 }
}
```

## `pkg/random`

New things I learned while building `pkg/random`.

### Testing

#### `go test -race -count=N`

My [first implementation](https://github.com/meleu/zzgo/blob/5603319a5d24275bdcb762fa65f125823e65059d/pkg/random/random.go) uses a global variable, which is a shared mutable state.
This approach is not thread-safe, which can be confirmed by flagging all tests
with `t.Parallel()` and then running `go test -race -count=100`

### Go idioms

- **Swap without a temp**: `a, b = b, a`.
- **Integer range loop** (Go 1.22+): `for range 50 { ... }`
- Go has no `constructor` keyword; convention is a `New...` factory function returning a pointer.
- Two constructors for a default vs. custom path: `New()` (time-seeded) and `NewWithCustomSeed(seed int64)` (seed provided by caller).

#### The `init()` function

**NOTE**: this is not used in the final version but was an interesting learning.

- Runs automatically once when the package is loaded, before any caller uses the package.
- Useful for one-time setup of package-level state.
- Equivalent alternative for simple cases: initialize directly in the `var` declaration.
- Used in [this version of the code](https://github.com/meleu/zzgo/blob/5603319a5d24275bdcb762fa65f125823e65059d/pkg/random/random.go)

### `math/rand` basics

- `rand.New(src Source) *rand.Rand` — wraps a `rand.Source into a generator.
- `rand.NewSource(seed int64) rand.Source` — creates a `rand.Source` with a seed (useful for deterministic results in testing).
- `(*rand.Rand).Intn(n int) int` — returns a value in `[0, n)`.
  - `r.Intn(max - min + 1) + min`: map to `[min, max]` inclusive.

### Package-level state vs encapsulation

- Exported mutable vars (`var Seed int64`) are a design smell: any caller can mutate them; order dependencies become implicit.
- Solution: callers own the state explicitly
  - A `Generator` type with a constructor and methods (clean design — no globals)
  - This what we see [currently implemented](https://github.com/meleu/zzgo/blob/main/pkg/random/random.go)
