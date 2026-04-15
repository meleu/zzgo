# Learning Notes

My notes about new things I'm learning while build this project.

Although public, it's very personal (which means "I'm not taking notes of things I already know by heart").

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
