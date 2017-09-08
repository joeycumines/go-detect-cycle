# floyds
--
    import "github.com/joeycumines/go-detect-cycle/floyds"

Package floyds provides means of detecting cycles, currently it is specifically
designed to be used within recursive functions, and implements Floyd's Tortoise
and Hare algorithm.

## Usage

#### type BranchingDetector

```go
type BranchingDetector struct {
}
```

BranchingDetector uses the same logic as Detector (which implements the tortoise
and the hare), but with the addition of the ability to support branching logic,
at the cost of something like O(n) memory usage, but can be used with a simple
stepper, that simply gets passed each step sequentially.

#### func  NewBranchingDetector

```go
func NewBranchingDetector(start interface{}, compare func(tortoise, hare interface{}) bool) BranchingDetector
```
NewBranchingDetector constructs a new BranchingDetector with the given start
value, and optionally a custom comparison func, to determine if there was a
cycle.

#### func (BranchingDetector) Clear

```go
func (f BranchingDetector) Clear()
```
Clear will ensure that any step references, to the step IMMEDIATELY PRECEDING
the BranchingDetector f, will be cleared from it's internal structure. This
method should ideally be deferred on every level of recursion, since where the
recursive function (which received f as an argument) is the highest level that
will need to access that step.

#### func (BranchingDetector) Hare

```go
func (f BranchingDetector) Hare(step interface{}) BranchingDetector
```
Hare takes a step for the hare, automatically taking a step for the tortoise if
necessary, by storing the step internally, for later use, you should ensure the
return value's `Clear` method is called after it is no longer necessary.

#### func (BranchingDetector) HareCount

```go
func (f BranchingDetector) HareCount() int
```
HareCount gets the number of steps that hare has taken, since the start.

#### func (BranchingDetector) Ok

```go
func (f BranchingDetector) Ok() bool
```
Ok will return true only if there has been no cycle detected so far.

#### func (BranchingDetector) TortoiseCount

```go
func (f BranchingDetector) TortoiseCount() int
```
TortoiseCount gets the number of steps that tortoise has taken, since the start.

#### type Detector

```go
type Detector struct {
}
```

The Detector struct is a cycle detector using Floyd's tortoise and hare
algorithm. It is designed to be used in recursive algorithms, but will work just
as well in simple loops (it is very easy to manually implement in simple loops,
however). This implementation requires each step to be determined by the
previous value, like `f(old) = new`. The `BranchingDetector` struct only
requires a single value to be passed (where the main worker is iterating over
the graph as the hare), however it is unlikely to be as performant, and does not
provide the same level of memory efficiency, since it will, necessarily, store
the previous values of the hare in a stack, for use by the tortoise. If you call
any methods on something that was not constructed using either of the
constructors, a panic will occur.

It's worth mentioning that this implementation treats the current state as
immutable, which results in safer code, at the expense of garbage collection.
That aside, the memory cost of the Detector struct itself is O(1), one of the
advantages of the algorithm used.

I may implement the rest of the algorithm at a later date, which will probably
include other types of cycle detection, if it happens.

Usage:

    - Create with `NewDetector`, providing start and next, optionally providing compare (defaults to equality).
    - Increment with either `Hare` OR `Tortoise`, using `Ok` check for cycles, and passing down the new structs.
    - If you are using `Tortoise` method, you will want to indicate done (out of bounds), by make it return false
    	from the next method, or you may not get the results you expect.

Branching Logic:

If your algorithm has branching logic, where it forms a directed graph (and you
only care about cycles from the current leaf to the root), please use the
`BranchingDetector` struct, by calling it's constructor `NewBranchingDetector`.

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

#### func  NewDetector

```go
func NewDetector(start interface{}, next func(v interface{}) (interface{}, bool), compare func(tortoise, hare interface{}) bool) Detector
```
NewDetector constructs a new Detector struct, and must provide the start step,
and function to resolve the next step (from the previous step each time), and
may optionally include a custom comparison method.

#### func (Detector) Done

```go
func (f Detector) Done() bool
```
Done will return true if any calls to next have returned a false ok value.

#### func (Detector) Hare

```go
func (f Detector) Hare(step interface{}) Detector
```
Hare returns a new Detector with the hare moved forward one step (which must be
provided), incrementing the tortoise one step if required.

#### func (Detector) HareCount

```go
func (f Detector) HareCount() int
```
HareCount gets the number of steps that hare has taken, since the start.

#### func (Detector) Ok

```go
func (f Detector) Ok() bool
```
Ok will return true only if there has been no cycle detected so far.

#### func (Detector) SetCompare

```go
func (f Detector) SetCompare(compare func(tortoise, hare interface{}) bool) Detector
```
SetCompare returns a new Detector that is the same as the receiver, but with the
provided compare function.

#### func (Detector) SetNext

```go
func (f Detector) SetNext(next func(v interface{}) (step interface{}, ok bool)) Detector
```
SetNext returns a new Detector that is the same as the receiver, but with the
provided next function.

#### func (Detector) Tortoise

```go
func (f Detector) Tortoise(step interface{}) Detector
```
Tortoise returns a new Detector with the tortoise moved forward one step (which
must be provided), incrementing the hare at least two steps, using next (it will
always be two if Hare has not been called at all).

#### func (Detector) TortoiseCount

```go
func (f Detector) TortoiseCount() int
```
TortoiseCount gets the number of steps that tortoise has taken, since the start.
