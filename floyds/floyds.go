// Package floyds provides means of detecting cycles, currently it is specifically designed to be used within
// recursive functions, and implements Floyd's Tortoise and Hare algorithm.
package floyds

import "errors"

/*
The Detector struct is a cycle detector using Floyd's tortoise and hare algorithm. It is designed to be used in
recursive algorithms, but will work just as well in simple loops (it is very easy to manually implement in simple
loops, however). This implementation requires each step to be determined by the previous value, like `f(old) = new`.
The `BranchingDetector` struct only requires a single value to be passed (where the main worker is iterating over the
graph as the hare), however it is unlikely to be as performant, and does not provide the same level of memory
efficiency, since it will, necessarily, store the previous values of the hare in a stack, for use by the tortoise.
If you call any methods on something that was not constructed using either of the constructors, a panic will occur.

It's worth mentioning that this implementation treats the current state as immutable, which results in safer code,
at the expense of garbage collection. That aside, the memory cost of the Detector struct itself is O(1), one of the
advantages of the algorithm used.

I may implement the rest of the algorithm at a later date, which will probably include other types of cycle detection,
if it happens.

Usage:

	- Create with `NewDetector`, providing start and next, optionally providing compare (defaults to equality).
	- Increment with either `Hare` OR `Tortoise`, using `Ok` check for cycles, and passing down the new structs.
	- If you are using `Tortoise` method, you will want to indicate done (out of bounds), by make it return false
		from the next method, or you may not get the results you expect.

Branching Logic:

If your algorithm has branching logic, where it forms a directed graph (and you only care about cycles from the
current leaf to the root), please use the `BranchingDetector` struct, by calling it's constructor
`NewBranchingDetector`.

Floyd's Tortoise and Hare algorithm, for reference. Only the first segment is currently implemented.

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
 */
type Detector struct {
	next          func(v interface{}) (step interface{}, ok bool)
	compare       func(tortoise, hare interface{}) bool
	tortoise      interface{}
	hare          interface{}
	ok            bool
	done          bool
	hareCount     int
	tortoiseCount int
}

// The default compare function simply compares equality.
func compareEquality(a, b interface{}) bool {
	return a == b
}

// NewDetector constructs a new Detector struct, and must provide the start step, and function to resolve the next
// step (from the previous step each time), and may optionally include a custom comparison method.
func NewDetector(start interface{}, next func(v interface{}) (interface{}, bool), compare func(tortoise, hare interface{}) bool) Detector {
	if nil == next {
		panic(errors.New("[NewDetector] next must be non-nil"))
	}
	if nil == compare {
		compare = compareEquality
	}
	return Detector{next, compare, start, start, true, false, 0, 0}
}

func (f Detector) validate() {
	if nil == f.next || nil == f.compare {
		panic(errors.New("[Detector.validate] nil property encountered, use the constructor NewDetector"))
	}
}

// Ok will return true only if there has been no cycle detected so far.
func (f Detector) Ok() bool {
	f.validate()
	return f.ok
}

// Checks if hare caught up to tortoise, which would indicate a cycle.
func (f Detector) check() bool {
	f.validate()
	if false == f.ok {
		return false
	}
	return false == f.compare(f.tortoise, f.hare)
}

// SetNext returns a new Detector that is the same as the receiver, but with the provided next function.
func (f Detector) SetNext(next func(v interface{}) (step interface{}, ok bool)) Detector {
	f.validate()
	if nil == next {
		panic(errors.New("[Detector.SetNext] you cannot set a nil next"))
	}
	f.next = next
	return f
}

// SetCompare returns a new Detector that is the same as the receiver, but with the provided compare function.
func (f Detector) SetCompare(compare func(tortoise, hare interface{}) bool) Detector {
	f.validate()
	if nil == compare {
		panic(errors.New("[Detector.SetCompare] you cannot set a nil compare"))
	}
	f.compare = compare
	return f
}

// Hare returns a new Detector with the hare moved forward one step (which must be provided), incrementing the tortoise
// one step if required.
func (f Detector) Hare(step interface{}) Detector {
	f.validate()
	if false == f.ok || true == f.done {
		return f
	}

	// on even counts tortoise is incremented
	if 0 == (f.hareCount % 2) {
		next, ok := f.next(f.tortoise)
		if false == ok {
			// no change, exit immediately - there was no cycle
			f.done = true
			return f
		}
		f.tortoise = next
		f.tortoiseCount++
	}

	// one step for hare - we only provided one value
	f.hare = step
	f.hareCount++

	// we can only check for equality after one tortoise and two hares
	if 0 == (f.hareCount % 2) {
		f.ok = f.check()
	}

	return f
}

// Tortoise returns a new Detector with the tortoise moved forward one step (which must be provided), incrementing the
// hare at least two steps, using next (it will always be two if Hare has not been called at all).
func (f Detector) Tortoise(step interface{}) Detector {
	f.validate()
	if false == f.ok || true == f.done {
		return f
	}

	// tortoise can only be taken on even steps, this check is just a safeguard for any random Hare calls
	if 0 != (f.hareCount % 2) {
		next, ok := f.next(f.hare)
		if false == ok {
			// no change, exit immediately - there was no cycle
			f.done = true
			return f
		}
		f.hare = next
		f.hareCount++
		f.ok = f.check()
		if false == f.ok {
			return f
		}
	}

	// a single step has been provided for tortoise
	f.tortoise = step
	f.tortoiseCount++

	// step #1 for hare
	next, ok := f.next(f.hare)
	if false == ok {
		// no change, exit immediately - there was no cycle
		f.done = true
		return f
	}
	f.hare = next
	f.hareCount++

	// step #2 for hare
	next, ok = f.next(f.hare)
	if false == ok {
		// no change, exit immediately - there was no cycle
		f.done = true
		return f
	}
	f.hare = next
	f.hareCount++

	// the final state is always at a check point
	f.ok = f.check()
	return f
}

// HareCount gets the number of steps that hare has taken, since the start.
func (f Detector) HareCount() int {
	f.validate()
	return f.hareCount
}

// TortoiseCount gets the number of steps that tortoise has taken, since the start.
func (f Detector) TortoiseCount() int {
	f.validate()
	return f.tortoiseCount
}

// Done will return true if any calls to next have returned a false ok value.
func (f Detector) Done() bool {
	f.validate()
	return f.done
}

// BranchingDetector uses the same logic as Detector (which implements the tortoise and the hare), but with the
// addition of the ability to support branching logic, at the cost of something like O(n) memory usage, but can be
// used with a simple stepper, that simply gets passed each step sequentially.
type BranchingDetector struct {
	f     Detector
	next  []interface{}
	clear func()
}

// The emptyNext function is a placeholder to avoid triggering a panic.
func emptyNext(v interface{}) (interface{}, bool) {
	return nil, false
}

// NewBranchingDetector constructs a new BranchingDetector with the given start value, and optionally a custom
// comparison func, to determine if there was a cycle.
func NewBranchingDetector(start interface{}, compare func(tortoise, hare interface{}) bool) BranchingDetector {
	return BranchingDetector{
		f: NewDetector(
			start,
			emptyNext,
			compare,
		),
	}
}

// Clear will ensure that any step references, to the step IMMEDIATELY PRECEDING the BranchingDetector f, will be
// cleared from it's internal structure. This method should ideally be deferred on every level of recursion, since
// where the recursive function (which received f as an argument) is the highest level that will need to access that
// step.
func (f BranchingDetector) Clear() {
	if nil == f.clear {
		return
	}
	f.clear()
}

// The nextUpdater struct encapsulates the logic required to manage the step queue for the internal tortoise stepper.
type nextUpdater struct {
	next []interface{}
}

// The logNext method provides a next function that will automatically clear each step taken from the internal queue.
func (u *nextUpdater) logNext(v interface{}) (step interface{}, ok bool) {
	if nil == u {
		return nil, false
	}
	for _, n := range u.next {
		u.next = u.next[1:]
		return n, true
	}
	return nil, false
}

// The genClear function returns a function that will clear the last element of the next slice to nil, or the whole
// slice if the all flag is provided.
func genClear(next []interface{}, all bool) func() {
	return func() {
		if nil == next {
			return
		}
		if true == all {
			for i := range next {
				next[i] = nil
			}
			next = nil
			return
		}
		end := len(next) - 1
		if 0 > end {
			next = nil
			return
		}
		next[end] = nil
		next = nil
	}
}

// The updateNext method will return a new BranchingDetector with the (potentially updated) next slice, from the receiver,
// as well as clearing the next method (which should be from the receiver).
// It also clears the reference to the next slice from the receiver, since it's expected to go out of scope.
func (u *nextUpdater) updateNext(f BranchingDetector) BranchingDetector {
	if nil == u {
		return f
	}
	f.next = u.next
	u.next = nil
	f.f.next = emptyNext
	return f
}

// Hare takes a step for the hare, automatically taking a step for the tortoise if necessary, by storing the step
// internally, for later use, you should ensure the return value's `Clear` method is called after it is no longer
// necessary.
func (f BranchingDetector) Hare(step interface{}) BranchingDetector {
	f.f.validate()
	if false == f.f.ok || true == f.f.done {
		return f
	}
	updater := new(nextUpdater)
	// At this point, updater.next has step as it's last value, and f.next might point to a different array.
	updater.next = append(f.next, step)
	f.f.next = updater.logNext
	return updater.updateNext(BranchingDetector{
		f: f.f.Hare(step),
		// When clear is called, it will clear step from the array that holds the reference to it (which is potentially
		// shared between multiple child branches, and, in the event that it was actually a COPY (the capacity had to
		// increase on `append`), it will clear ALL PREVIOUS indexes (in that copy array).
		clear: genClear(
			updater.next,
			cap(updater.next) > cap(f.next),
		),
	})
}

// Ok will return true only if there has been no cycle detected so far.
func (f BranchingDetector) Ok() bool {
	return f.f.Ok()
}

// Done is kept unexported - it's tested since it's part of the logic of the core implementation, but it doesn't
// actually do anything useful in this case.
func (f BranchingDetector) done() bool {
	return f.f.Done()
}

// HareCount gets the number of steps that hare has taken, since the start.
func (f BranchingDetector) HareCount() int {
	return f.f.HareCount()
}

// TortoiseCount gets the number of steps that tortoise has taken, since the start.
func (f BranchingDetector) TortoiseCount() int {
	return f.f.TortoiseCount()
}
