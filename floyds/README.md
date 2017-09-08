# floyds
--
    import "github.com/joeycumines/go-detect-cycle/floyds"

Package floyds provides means of detecting cycles, currently it is specifically
designed to be used within recursive functions, and implements Floyd's Tortoise
and Hare algorithm.

## Usage

#### type Floyds

```go
type Floyds struct {
}
```

The Floyds struct is cycle detector using Floyd's tortoise and hare algorithm,
it is designed to be used in recursive algorithms, but will work just as well in
simple loops (it is very easy to implement in simple loops however). This
implementation requires each step to be determined by the previous value, like
`f(old) = new`. The `FloydsBranch` struct only requires a single value to be
passed (where the main worker is iterating over the graph as the hare), however
it is unlikely to be as performant, and does not provide the same level of
memory efficiency, since it will, necessarily, store the previous values of the
hare in a stack, for use by the tortoise. If you call any methods on something
that was not constructed using `NewFloyds`, a panic will occur.

It's worth mentioning that this implementation treats the current state as
immutable, which results in safer code, at the expense of garbage collection.
That aside, the memory cost of the structs themselves is constant, one of the
advantages of the algorithm used.

I may implement the rest of the algorithm at a later date, which will probably
include other types of cycle detection, if it happens.

Usage:

    - Create with `NewFloyds`, providing start and next, optionally providing compare (defaults to equality).
    - Increment with either `Hare` OR `Tortoise`, using `Ok` check for cycles, and passing down the new structs.
    - If you are using `Tortoise` method, you will want to indicate done (out of bounds), by make it return false
    	from the next method, or you may not get the results you expect.

Branching Logic:

If your algorithm has branching logic, where it forms a directed graph (and you
only care about cycles from the current leaf to the root), please use the
`FloydsBranch` struct, by calling it's constructor `NewFloydsBranch`.

Floyd's Tortoise and Hare algorithm, for reference. Only the first segment is
currently implemented.

https://en.wikipedia.org/wiki/Cycle_detection

    def floyd(f, x0):
    	# Main phase of algorithm: finding a repetition x_i = x_2i.
    	# The hare moves twice as quickly as the tortoise and
    	# the distance between them increases by 1 at each step.
    	# Eventually they will both be inside the cycle and then,
    	# at some point, the distance between them will be
    	# divisible by the period λ.
    	tortoise = f(x0) # f(x0) is the element/node next to x0.
    	hare = f(f(x0))
    	while tortoise != hare:
    		tortoise = f(tortoise)
    		hare = f(f(hare))

    	# At this point the tortoise position, ν, which is also equal
    	# to the distance between hare and tortoise, is divisible by
    	# the period λ. So hare moving in circle one step at a time,
    	# and tortoise (reset to x0) moving towards the circle, will
    	# intersect at the beginning of the circle. Because the
    	# distance between them is constant at 2ν, a multiple of λ,
    	# they will agree as soon as the tortoise reaches index μ.

    	# Find the position μ of first repetition.
    	mu = 0
    	tortoise = x0
    	while tortoise != hare:
    		tortoise = f(tortoise)
    		hare = f(hare)   # Hare and tortoise move at same speed
    		mu += 1

    	# Find the length of the shortest cycle starting from x_μ
    	# The hare moves one step at a time while tortoise is still.
    	# lam is incremented until λ is found.
    	lam = 1
    	hare = f(tortoise)
    	while tortoise != hare:
    		hare = f(hare)
    		lam += 1

    	return lam, mu

#### func  NewFloyds

```go
func NewFloyds(start interface{}, next func(v interface{}) (interface{}, bool), compare func(tortoise, hare interface{}) bool) Floyds
```
NewFloyds is the constructor for the Floyds struct, and must provide the start
step, and function to resolve the next step (from the previous step each time),
and may optionally include a custom comparison method.

#### func (Floyds) Done

```go
func (f Floyds) Done() bool
```
Done will return true if any calls to next have returned a false ok value.

#### func (Floyds) Hare

```go
func (f Floyds) Hare(step interface{}) Floyds
```
Hare returns a new Floyds with the hare moved forward one step (which must be
provided), incrementing the tortoise one step if required.

#### func (Floyds) HareCount

```go
func (f Floyds) HareCount() int
```
HareCount gets the number of steps that hare has taken, since the start.

#### func (Floyds) Ok

```go
func (f Floyds) Ok() bool
```
Ok will return true only if there has been no cycle detected so far.

#### func (Floyds) SetCompare

```go
func (f Floyds) SetCompare(compare func(tortoise, hare interface{}) bool) Floyds
```
SetCompare returns a new Floyds that is the same as the receiver, but with the
provided compare function.

#### func (Floyds) SetNext

```go
func (f Floyds) SetNext(next func(v interface{}) (step interface{}, ok bool)) Floyds
```
SetNext returns a new Floyds that is the same as the receiver, but with the
provided next function.

#### func (Floyds) Tortoise

```go
func (f Floyds) Tortoise(step interface{}) Floyds
```
Tortoise returns a new Floyds with the tortoise moved forward one step (which
must be provided), incrementing the hare at least two steps, using next (it will
always be two if Hare has not been called at all).

#### func (Floyds) TortoiseCount

```go
func (f Floyds) TortoiseCount() int
```
TortoiseCount gets the number of steps that tortoise has taken, since the start.

#### type FloydsBranch

```go
type FloydsBranch struct {
}
```

FloydsBranch uses all the same logic as Floyds (which implements the tortoise
and the hare), but with the addition of the ability to support branching logic,
at the cost of efficiency, and the sacrifice of the `Tortoise` method.

#### func  NewFloydsBranch

```go
func NewFloydsBranch(start interface{}, compare func(tortoise, hare interface{}) bool) FloydsBranch
```
NewFloydsBranch constructs a new FloydsBranch with the given start value, and
optionally a custom comparison func, to determine if there was a cycle.

#### func (FloydsBranch) Clear

```go
func (f FloydsBranch) Clear()
```
Clear will ensure that any step references, to the step IMMEDIATELY PRECEDING
the FloydsBranch f, will be cleared from it's internal structure. This method
should ideally be deferred on every level of recursion, since where the
recursive function (which received f as an argument) is the highest level that
will need to access that step.

#### func (FloydsBranch) Hare

```go
func (f FloydsBranch) Hare(step interface{}) FloydsBranch
```
Hare takes a step for the hare, automatically taking a step for the tortoise if
necessary, by storing the step internally, for later use, you should ensure the
return value's `Clear` method is called after it is no longer necessary.

#### func (FloydsBranch) HareCount

```go
func (f FloydsBranch) HareCount() int
```
HareCount gets the number of steps that hare has taken, since the start.

#### func (FloydsBranch) Ok

```go
func (f FloydsBranch) Ok() bool
```
Ok will return true only if there has been no cycle detected so far.

#### func (FloydsBranch) TortoiseCount

```go
func (f FloydsBranch) TortoiseCount() int
```
TortoiseCount gets the number of steps that tortoise has taken, since the start.
