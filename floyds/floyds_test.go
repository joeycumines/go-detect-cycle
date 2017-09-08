package floyds

import (
	"testing"
	"math/rand"
	"fmt"
	"encoding/json"
)

func TestDetector_Hare1(t *testing.T) {
	list := []string{"one", "two", "three", "one", "two", "three", "one", "two", "three", "one", "two", "three", "one", "two", "three"}
	// create a cycle of list
	next := func(v interface{}) (interface{}, bool) {
		return v.(int) + 1, true
	}
	// compare the string values
	compare := func(a, b interface{}) bool {
		iA, iB := a.(int), b.(int)
		if iA >= len(list) || iB >= len(list) {
			t.Fatal()
		}
		sA, sB := list[iA], list[iB]
		return sA == sB
	}
	f := NewDetector(0, next, compare)
	log := []int{f.hare.(int)}
	x := 0
	for f.Ok() {
		if x != f.hare {
			t.Fatal()
		}
		x++
		f = f.Hare(f.hare.(int) + 1)
		if x != f.hare {
			t.Fatal()
		}

		expectedTortoise := x
		if 1 == (expectedTortoise % 2) {
			expectedTortoise++
		}
		expectedTortoise = expectedTortoise / 2
		if expectedTortoise != f.tortoise {
			t.Fatal()
		}

		log = append(log, f.hare.(int))
	}
}

func TestDetector_Hare2(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5}
	// create a cycle of list
	next := func(v interface{}) (interface{}, bool) {
		return v.(int) + 1, true
	}
	// compare the string values
	compare := func(a, b interface{}) bool {
		iA, iB := a.(int), b.(int)
		if iA >= len(list) || iB >= len(list) {
			t.Fatal()
		}
		sA, sB := list[iA], list[iB]
		return sA == sB
	}
	f := NewDetector(0, next, compare)
	log := []int{f.hare.(int)}
	x := 0
	for f.Ok() {
		if x != f.hare {
			t.Fatal()
		}
		x++
		f = f.Hare(f.hare.(int) + 1)
		if x != f.hare {
			t.Fatalf("%v %v", x, f.hare)
		}

		expectedTortoise := x
		if 1 == (expectedTortoise % 2) {
			expectedTortoise++
		}
		expectedTortoise = expectedTortoise / 2
		if expectedTortoise != f.tortoise {
			t.Fatal()
		}

		log = append(log, f.hare.(int))
	}

	//t.Fatalf("%v", log)
}

func TestDetector_Hare3(t *testing.T) {
	list := []int{0, 1, 2, 3, 1, 2, 3, 2, 123, 0, 1, 2, 3, 1, 2, 3, 2, 123, 0, 1, 2, 3, 1, 2, 3, 2, 123, 0, 1, 2, 3, 1, 2, 3, 2, 123}
	// create a cycle of list
	next := func(v interface{}) (interface{}, bool) {
		return v.(int) + 1, true
	}
	// compare the string values
	compare := func(a, b interface{}) bool {
		iA, iB := a.(int), b.(int)
		if iA >= len(list) || iB >= len(list) {
			t.Fatal()
		}
		sA, sB := list[iA], list[iB]
		return sA == sB
	}
	f := NewDetector(0, next, compare)
	log := []int{f.hare.(int)}
	x := 0
	for f.Ok() {
		if x != f.hare {
			t.Fatal()
		}
		x++
		f = f.Hare(f.hare.(int) + 1)
		if x != f.hare {
			t.Fatal()
		}

		expectedTortoise := x
		if 1 == (expectedTortoise % 2) {
			expectedTortoise++
		}
		expectedTortoise = expectedTortoise / 2
		if expectedTortoise != f.tortoise {
			t.Fatal()
		}

		log = append(log, f.hare.(int))
	}

	//t.Fatalf("%v", log)
}

func TestDetector_Hare4_noCycle(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// create a cycle of list
	next := func(v interface{}) (interface{}, bool) {
		return v.(int) + 1, true
	}
	// compare the string values
	compare := func(a, b interface{}) bool {
		iA, iB := a.(int), b.(int)
		if iA >= len(list) || iB >= len(list) {
			t.Fatal()
		}
		sA, sB := list[iA], list[iB]
		return sA == sB
	}
	f := NewDetector(0, next, compare)
	log := []int{f.hare.(int)}
	x := 0
	for f.Ok() && x < len(list) {
		expectedHareCount := f.HareCount() + 1
		expectedTortoiseCount := f.TortoiseCount()
		if 1 == (expectedHareCount % 2) {
			expectedTortoiseCount++
		}

		if x != f.hare {
			t.Fatal()
		}
		x++
		f = f.Hare(f.hare.(int) + 1)
		if x != f.hare {
			t.Fatal()
		}

		expectedTortoise := x
		if 1 == (expectedTortoise % 2) {
			expectedTortoise++
		}
		expectedTortoise = expectedTortoise / 2
		if expectedTortoise != f.tortoise {
			t.Fatal()
		}

		if f.HareCount() != expectedHareCount || f.TortoiseCount() != expectedTortoiseCount {
			t.Fatal()
		}

		log = append(log, f.hare.(int))
	}

	if x != len(list) || false == f.Ok() || f.hareCount != len(log)-1 || true == f.Done() {
		t.Fatal()
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func TestDetector_Hare4_noCycleDone(t *testing.T) {
	hareList := makeRange(0, 50)
	tortoiseList := makeRange(0, 10)
	// create a cycle of hareList
	next := func(v interface{}) (interface{}, bool) {
		next := v.(int) + 1
		if next >= len(tortoiseList) {
			return nil, false
		}
		return next, true
	}
	// compare the string values
	compare := func(tortoise, hare interface{}) bool {
		iT, iHare := tortoise.(int), hare.(int)
		if iT >= len(tortoiseList) || iHare >= len(hareList) {

		}
		sT, sH := tortoiseList[iT], hareList[iHare]
		return sT == sH
	}
	f := NewDetector(0, next, compare)
	log := []int{f.hare.(int)}
	x := 0
	for f.Ok() && !f.Done() {
		expectedHareCount := f.HareCount() + 1
		expectedTortoiseCount := f.TortoiseCount()
		if 1 == (expectedHareCount % 2) {
			expectedTortoiseCount++
		}

		if x != f.hare {
			t.Fatal()
		}
		x++

		fOld := f
		f = f.Hare(f.hare.(int) + 1)

		// check the bail out when done
		if true == f.Done() {
			if f.hare != fOld.hare || f.tortoise != fOld.tortoise || f.hareCount != fOld.hareCount || f.tortoiseCount != fOld.tortoiseCount {
				t.Fatal()
			}
			continue
		}

		if x != f.hare {
			t.Fatal()
		}

		expectedTortoise := x
		if 1 == (expectedTortoise % 2) {
			expectedTortoise++
		}
		expectedTortoise = expectedTortoise / 2
		if expectedTortoise != f.tortoise {
			t.Fatal()
		}

		if f.HareCount() != expectedHareCount || f.TortoiseCount() != expectedTortoiseCount {
			t.Fatal()
		}

		log = append(log, f.hare.(int))
	}

	if x < len(tortoiseList) || false == f.Ok() || f.hareCount != len(log)-1 || false == f.Done() {
		t.Fatal()
	}
}

func TestDetector_TortoiseCount(t *testing.T) {
	next := func(v interface{}) (interface{}, bool) {
		return v.(int) + 1, true
	}
	f := NewDetector(0, next, nil)
	if 0 != f.TortoiseCount() {
		t.Fatal()
	}
	f.tortoiseCount = 44
	if 44 != f.TortoiseCount() {
		t.Fatal()
	}
}

func TestDetector_HareCount(t *testing.T) {
	next := func(v interface{}) (interface{}, bool) {
		return v.(int) + 1, true
	}
	f := NewDetector(0, next, nil)
	if 0 != f.HareCount() {
		t.Fatal()
	}
	f.hareCount = 44
	if 44 != f.HareCount() {
		t.Fatal()
	}
}

func TestDetector_Hare_panic(t *testing.T) {
	f := Detector{}
	func() {
		defer func() {
			if err, ok := recover().(error); false == ok || nil == err ||
				"[Detector.validate] nil property encountered, use the constructor NewDetector" != err.Error() {
				t.Fatal()
			}
		}()
		f.Hare(nil)
		t.Fatal()
	}()
}

func TestDetector_Tortoise_panic(t *testing.T) {
	f := Detector{}
	func() {
		defer func() {
			if err, ok := recover().(error); false == ok || nil == err ||
				"[Detector.validate] nil property encountered, use the constructor NewDetector" != err.Error() {
				t.Fatal()
			}
		}()
		f.Tortoise(nil)
		t.Fatal()
	}()
}

func TestDetector_HareCount_panic(t *testing.T) {
	f := Detector{}
	func() {
		defer func() {
			if err, ok := recover().(error); false == ok || nil == err ||
				"[Detector.validate] nil property encountered, use the constructor NewDetector" != err.Error() {
				t.Fatal()
			}
		}()
		f.HareCount()
		t.Fatal()
	}()
}

func TestDetector_TortoiseCount_panic(t *testing.T) {
	f := Detector{}
	func() {
		defer func() {
			if err, ok := recover().(error); false == ok || nil == err ||
				"[Detector.validate] nil property encountered, use the constructor NewDetector" != err.Error() {
				t.Fatal()
			}
		}()
		f.TortoiseCount()
		t.Fatal()
	}()
}

func TestDetector_Ok_panic(t *testing.T) {
	f := Detector{}
	func() {
		defer func() {
			if err, ok := recover().(error); false == ok || nil == err ||
				"[Detector.validate] nil property encountered, use the constructor NewDetector" != err.Error() {
				t.Fatal()
			}
		}()
		f.Ok()
		t.Fatal()
	}()
}

func TestDetector_Done_panic(t *testing.T) {
	f := Detector{}
	func() {
		defer func() {
			if err, ok := recover().(error); false == ok || nil == err ||
				"[Detector.validate] nil property encountered, use the constructor NewDetector" != err.Error() {
				t.Fatal()
			}
		}()
		f.Ok()
		t.Fatal()
	}()
}

func TestDetector_Done(t *testing.T) {
	next := func(v interface{}) (interface{}, bool) {
		return "next", false
	}
	f := NewDetector(23, next, nil)
	if true == f.Done() || true == f.done {
		t.Fatal()
	}
	f.done = true
	if false == f.Done() || false == f.done {
		t.Fatal()
	}
	f.done = false
	f = f.Hare(241)
	if false == f.Done() || false == f.done {
		t.Fatal()
	}
}

func TestNewDetector_panic(t *testing.T) {
	func() {
		defer func() {
			if err, ok := recover().(error); false == ok || nil == err ||
				"[NewDetector] next must be non-nil" != err.Error() {
				t.Fatal()
			}
		}()
		NewDetector(nil, nil, nil)
		t.Fatal()
	}()
}

func TestNewDetector(t *testing.T) {
	next := func(v interface{}) (interface{}, bool) {
		return "next", true
	}
	f := NewDetector(23, next, nil)
	if nil == f.next || nil == f.compare {
		t.Fatal()
	}
	if v, _ := f.next(1); "next" != v {
		t.Fatal()
	}
	if false == f.compare(2323, 2323) || true == f.compare(2323, 2324) {
		t.Fatal()
	}
	if true == f.compare("1", 1) {
		t.Fatal()
	}
	compare := func(a, b interface{}) bool {
		return a == "left" && b == "right"
	}
	f = NewDetector(23, next, compare)
	if true == f.compare("right", "left") {
		t.Fatal()
	}
	if false == f.compare("left", "right") {
		t.Fatal()
	}
}

func TestDetector_check(t *testing.T) {
	next := func(v interface{}) (interface{}, bool) {
		return "next", true
	}
	compared := false
	compare := func(a, b interface{}) bool {
		compared = true
		return false
	}
	f := NewDetector(23, next, compare)
	if false == f.check() {
		t.Fatal()
	}
	if false == compared {
		t.Fatal()
	}
	f.ok = false
	if true == f.check() {
		t.Fatal()
	}
}

func TestDetector_checkNoCompare(t *testing.T) {
	next := func(v interface{}) (interface{}, bool) {
		return "next", true
	}
	compare := func(a, b interface{}) bool {
		t.Fatal()
		return false
	}
	f := NewDetector(23, next, compare)
	f.ok = false
	if true == f.check() {
		t.Fatal()
	}
}

func TestDetector_Hare_done(t *testing.T) {
	next := func(v interface{}) (interface{}, bool) {
		t.Fatal()
		return "next", true
	}
	compare := func(a, b interface{}) bool {
		t.Fatal()
		return false
	}
	f := NewDetector(23, next, compare)
	f.done = true
	f = f.Hare(nil)
	if false == f.done || false == f.ok {
		t.Fatal()
	}
}

func TestDetector_Tortoise_done(t *testing.T) {
	next := func(v interface{}) (interface{}, bool) {
		t.Fatal()
		return "next", true
	}
	compare := func(a, b interface{}) bool {
		t.Fatal()
		return false
	}
	f := NewDetector(23, next, compare)
	f.done = true
	f = f.Tortoise(nil)
	if false == f.done || false == f.ok {
		t.Fatal()
	}
}

func TestDetector_Hare_notOk(t *testing.T) {
	next := func(v interface{}) (interface{}, bool) {
		t.Fatal()
		return "next", true
	}
	compare := func(a, b interface{}) bool {
		t.Fatal()
		return false
	}
	f := NewDetector(23, next, compare)
	f.ok = false
	f = f.Hare(nil)
	if true == f.ok || true == f.done {
		t.Fatal()
	}
}

func TestDetector_Tortoise_notOk(t *testing.T) {
	next := func(v interface{}) (interface{}, bool) {
		t.Fatal()
		return "next", true
	}
	compare := func(a, b interface{}) bool {
		t.Fatal()
		return false
	}
	f := NewDetector(23, next, compare)
	f.ok = false
	f = f.Tortoise(nil)
	if true == f.ok || true == f.done {
		t.Fatal()
	}
}

func TestDetector_Tortoise_noCycle(t *testing.T) {
	list := makeRange(0, 20)
	// create a cycle of list
	next := func(v interface{}) (interface{}, bool) {
		n := v.(int) + 1
		if n >= len(list) {
			return nil, false
		}
		return n, true
	}
	// compare the string values
	compare := func(a, b interface{}) bool {
		iA, iB := a.(int), b.(int)
		if iA >= len(list) || iB >= len(list) {
			t.Fatal()
		}
		sA, sB := list[iA], list[iB]
		return sA == sB
	}
	f := NewDetector(0, next, compare)
	for f.Ok() && !f.Done() {
		fOld := f
		n, ok := next(f.tortoise)
		if false == ok {
			break
		}
		f = f.Tortoise(n)
		if f.tortoiseCount != fOld.tortoiseCount+1 {
			t.Fatalf("%v", f)
		}
		if f.tortoise.(int) != fOld.tortoise.(int)+1 {
			t.Fatalf("%v", f)
		}
		if true == f.Done() {
			if f.hare != fOld.hare || f.hareCount != fOld.hareCount {
				t.Fatalf("%v", f)
			}
			break
		}
		if f.hareCount != fOld.hareCount+2 {
			t.Fatalf("%v", f)
		}
		if f.hare.(int) != fOld.hare.(int)+2 {
			t.Fatalf("%v", f)
		}
	}
	if f.hareCount != len(list)-1 || false == f.Done() || false == f.Ok() {
		t.Fatalf("%v", f)
	}
}

func TestDetector_Tortoise_noCycle_oddNumber(t *testing.T) {
	list := makeRange(0, 21)
	// create a cycle of list
	next := func(v interface{}) (interface{}, bool) {
		n := v.(int) + 1
		if n >= len(list) {
			return nil, false
		}
		return n, true
	}
	// compare the string values
	compare := func(a, b interface{}) bool {
		iA, iB := a.(int), b.(int)
		if iA >= len(list) || iB >= len(list) {
			t.Fatal()
		}
		sA, sB := list[iA], list[iB]
		return sA == sB
	}
	f := NewDetector(0, next, compare)
	for f.Ok() && !f.Done() {
		fOld := f
		n, ok := next(f.tortoise)
		if false == ok {
			t.Fatal()
		}
		f = f.Tortoise(n)
		if f.tortoiseCount != fOld.tortoiseCount+1 {
			t.Fatalf("%v", f)
		}
		if f.tortoise.(int) != fOld.tortoise.(int)+1 {
			t.Fatalf("%v", f)
		}
		if true == f.Done() {
			if f.hareCount != fOld.hareCount+1 {
				t.Fatalf("%v", f)
			}
			if f.hare.(int) != fOld.hare.(int)+1 {
				t.Fatalf("%v", f)
			}
			break
		}
		if f.hareCount != fOld.hareCount+2 {
			t.Fatalf("%v", f)
		}
		if f.hare.(int) != fOld.hare.(int)+2 {
			t.Fatalf("%v", f)
		}
	}
	if f.hareCount != len(list)-1 || false == f.Done() || false == f.Ok() {
		t.Fatalf("%v", f)
	}
}

func TestDetector_Tortoise(t *testing.T) {
	list := []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3}
	// create a cycle of list
	next := func(v interface{}) (interface{}, bool) {
		n := v.(int) + 1
		if n >= len(list) {
			return nil, false
		}
		return n, true
	}
	// compare the string values
	compare := func(a, b interface{}) bool {
		iA, iB := a.(int), b.(int)
		if iA >= len(list) || iB >= len(list) {
			t.Fatal()
		}
		sA, sB := list[iA], list[iB]
		return sA == sB
	}
	f := NewDetector(0, next, compare)
	for f.Ok() && !f.Done() {
		fOld := f
		n, ok := next(f.tortoise)
		if false == ok {
			t.Fatal()
		}
		f = f.Tortoise(n)
		if false == f.Ok() {
			continue
		}
		if f.tortoiseCount != fOld.tortoiseCount+1 {
			t.Fatalf("%v", f)
		}
		if f.tortoise.(int) != fOld.tortoise.(int)+1 {
			t.Fatalf("%v", f)
		}
		if f.hareCount != fOld.hareCount+2 {
			t.Fatalf("%v", f)
		}
		if f.hare.(int) != fOld.hare.(int)+2 {
			t.Fatalf("%v", f)
		}
	}
	if f.tortoiseCount > f.hareCount || true == f.Ok() || true == f.Done() {
		t.Fatalf("%v", f)
	}
}

func TestDetector_Tortoise_badLogic(t *testing.T) {
	list := []int{0, 1, 2, 3, 0, 1, 2, 3, 0, 1, 2, 3, 0, 1, 2, 3, 0, 1, 2, 3, 0, 1, 2, 3, 0, 1, 2, 3}
	// create a cycle of list
	next := func(v interface{}) (interface{}, bool) {
		n := v.(int) + 1
		if n >= len(list) {
			return nil, false
		}
		return n, true
	}
	// compare the string values
	compare := func(a, b interface{}) bool {
		iA, iB := a.(int), b.(int)
		if iA >= len(list) || iB >= len(list) {
			t.Fatal()
		}
		sA, sB := list[iA], list[iB]
		return sA == sB
	}
	f := NewDetector(0, next, compare)

	// take a hare step
	n, _ := next(f.hare)
	f = f.Hare(n)

	// but then take a tortoise step - this will have to be our second tortoise step
	n, _ = next(f.tortoise)
	f = f.Tortoise(n)

	if 2 != f.TortoiseCount() || 4 != f.HareCount() || 4 != f.hare || 2 != f.tortoise {
		t.Fatalf("%v", f)
	}

	// another hare, one step for each tortoise and hare
	n, _ = next(f.hare)
	f = f.Hare(n)
	if 3 != f.TortoiseCount() || 5 != f.HareCount() || 5 != f.hare || 3 != f.tortoise {
		t.Fatalf("%v", f)
	}

	// and another tortoise, again takes one tortoise, 3 hare
	n, _ = next(f.tortoise)
	f = f.Tortoise(n)
	if 4 != f.TortoiseCount() || 8 != f.HareCount() || 8 != f.hare || 4 != f.tortoise {
		t.Fatalf("%v", f)
	}

	// at this point, they are all == 0, which means it would have exited here

	if f.tortoiseCount >= f.hareCount || true == f.Ok() || true == f.Done() {
		t.Fatalf("%v", f)
	}
}

func TestDetector_Bad_noCycle(t *testing.T) {
	for _, length := range makeRange(1, 100) {
		list := []int{}
		for x := 0; x < length; x++ {
			list = append(list, x)
		}
		next := func(v interface{}) (interface{}, bool) {
			n := v.(int) + 1
			if n >= len(list) {
				return nil, false
			}
			return n, true
		}
		// compare the string values
		compare := func(a, b interface{}) bool {
			iA, iB := a.(int), b.(int)
			if iA >= len(list) || iB >= len(list) {
				t.Fatal()
			}
			sA, sB := list[iA], list[iB]
			return sA == sB
		}
		f := NewDetector(0, next, compare)
		for f.Ok() && !f.Done() {
			oldF := f
			// take one hare
			n, ok := next(f.hare)
			if false == ok {
				n, ok = next(f.tortoise)
				f = f.Tortoise(f.tortoise)
				break
			}
			f = f.Hare(n)
			if true == f.Done() {
				t.Fatalf("%v", f)
			}
			// take one tortoise - have to exit on 1 case
			n, ok = next(f.tortoise)
			if false == ok {
				if f.tortoiseCount != f.hareCount {
					t.Fatalf("%v", f)
				}
				f.done = true
				break
			}
			// if we are about to reach the end - check
			_, atEnd := next(f.hare)
			atEnd = !atEnd
			f = f.Tortoise(n)
			if true == f.Done() {
				if true == atEnd {
					if oldF.TortoiseCount()+1 != f.TortoiseCount() || f.tortoise != f.TortoiseCount() ||
						oldF.HareCount()+1 != f.HareCount() || f.hare != f.HareCount() {
						t.Fatalf("%v %v", f, oldF)
					}
				}
				break
			}
			if oldF.TortoiseCount()+2 != f.TortoiseCount() || f.tortoise != f.TortoiseCount() ||
				oldF.HareCount()+4 != f.HareCount() || f.hare != f.HareCount() {
				t.Fatalf("%v", f)
			}
			if 0 != (f.TortoiseCount()%2) ||
				0 != (f.HareCount()%2) ||
				f.tortoiseCount != (f.HareCount()/2) {
				t.Fatalf("%v", f)
			}
		}
		if false == f.Ok() || false == f.Done() {
			t.Fatalf("%v", f)
		}
	}
}

func TestDetector_Bad_cycle(t *testing.T) {
	for _, length := range makeRange(1, 100) {
		list := []int{}
		for y := 0; y < 5; y++ {
			for x := 0; x < length; x++ {
				list = append(list, x)
			}
		}
		next := func(v interface{}) (interface{}, bool) {
			n := v.(int) + 1
			if n >= len(list) {
				return nil, false
			}
			return n, true
		}
		// compare the string values
		compare := func(a, b interface{}) bool {
			iA, iB := a.(int), b.(int)
			if iA >= len(list) || iB >= len(list) {
				t.Fatal()
			}
			sA, sB := list[iA], list[iB]
			return sA == sB
		}
		f := NewDetector(0, next, compare)
		for f.Ok() && !f.Done() {
			oldF := f
			// take one hare
			n, ok := next(f.hare)
			if false == ok {
				t.Fatalf("%v", f)
			}
			f = f.Hare(n)
			if true == f.Done() {
				t.Fatalf("%v", f)
			}
			// take one tortoise - have to exit on 1 case
			n, ok = next(f.tortoise)
			if false == ok {
				t.Fatalf("%v", f)
			}
			// if we are about to reach the end - check
			f = f.Tortoise(n)
			if true == f.Done() {
				t.Fatalf("%v", f)
			}
			if false == f.Ok() {
				break
			}
			if oldF.TortoiseCount()+2 != f.TortoiseCount() || f.tortoise != f.TortoiseCount() ||
				oldF.HareCount()+4 != f.HareCount() || f.hare != f.HareCount() {
				t.Fatalf("%v", f)
			}
			if 0 != (f.TortoiseCount()%2) ||
				0 != (f.HareCount()%2) ||
				f.tortoiseCount != (f.HareCount()/2) {
				t.Fatalf("%v", f)
			}
		}
		if true == f.Ok() || true == f.Done() {
			t.Fatalf("%v", f)
		}
	}
}

func TestDetector_Tortoise_short(t *testing.T) {
	next := func(v interface{}) (interface{}, bool) {
		return 0, true
	}
	compare := func(a, b interface{}) bool {
		return true
	}
	f := NewDetector(0, next, compare)
	f = f.Tortoise(0)
	if 0 != f.tortoise || 0 != f.hare || 1 != f.tortoiseCount || 2 != f.hareCount {
		t.Fatalf("%v", f)
	}
	if true == f.Done() || true == f.Ok() {
		t.Fatalf("%v", f)
	}
}

func TestDetector_hare_short(t *testing.T) {
	next := func(v interface{}) (interface{}, bool) {
		return 0, true
	}
	compare := func(a, b interface{}) bool {
		return true
	}
	f := NewDetector(0, next, compare)
	f = f.Hare(0)
	if 0 != f.tortoise || 0 != f.hare || 1 != f.tortoiseCount || 1 != f.hareCount {
		t.Fatalf("%v", f)
	}
	if true == f.Done() || false == f.Ok() {
		t.Fatalf("%v", f)
	}
	f = f.Hare(0)
	if 0 != f.tortoise || 0 != f.hare || 1 != f.tortoiseCount || 2 != f.hareCount {
		t.Fatalf("%v", f)
	}
	if true == f.Done() || true == f.Ok() {
		t.Fatalf("%v", f)
	}
}

func Test_emptyNext(t *testing.T) {
	n, ok := emptyNext(12)
	if nil != n || false != ok {
		t.Fatal()
	}
}

func TestNewBranchingDetector(t *testing.T) {
	f := NewBranchingDetector(22, nil)
	if nil == f.f.compare || nil == f.f.next || false == f.f.ok || true == f.f.done || 0 != f.f.tortoiseCount || 0 != f.f.hareCount ||
		22 != f.f.hare || 22 != f.f.tortoise || true == f.f.compare(1, 2) || false == f.f.compare(1, 1) ||
		0 != len(f.next) {
		t.Fatal()
	}
}

func TestBranchingDetector_Done(t *testing.T) {
	f := NewBranchingDetector(nil, nil)
	if true == f.done() {
		t.Fatal()
	}
	f.f.done = true
	if false == f.done() {
		t.Fatal()
	}
}

func TestBranchingDetector_Ok(t *testing.T) {
	f := NewBranchingDetector(nil, nil)
	if false == f.Ok() {
		t.Fatal()
	}
	f.f.ok = false
	if true == f.done() {
		t.Fatal()
	}
}

func TestBranchingDetector_HareCount(t *testing.T) {
	f := NewBranchingDetector(nil, nil)
	f.f.hareCount = 999
	if 999 != f.HareCount() {
		t.Fatal()
	}
}

func TestBranchingDetector_TortoiseCount(t *testing.T) {
	f := NewBranchingDetector(nil, nil)
	f.f.tortoiseCount = 999
	if 999 != f.TortoiseCount() {
		t.Fatal()
	}
}

func TestBranchingDetector_Hare_noCycle(t *testing.T) {
	input := map[int]map[int]map[int][]int{
		1: {
			2: {
				3: {4},
				4: {5},
				5: {4},
			},
		},
		2: {
			3: {
				1: {4},
			},
		},
		3: {
			5: {
				1: {2},
			},
		},
	}

	f := NewBranchingDetector(0, nil)

	for x, xMap := range input {
		f := f.Hare(x)
		if false == f.Ok() || true == f.done() {
			t.Fatalf("%v", f)
		}
		for y, yMap := range xMap {
			f := f.Hare(y)
			if false == f.Ok() || true == f.done() {
				t.Fatal()
			}
			for z, list := range yMap {
				f := f.Hare(z)
				if false == f.Ok() || true == f.done() {
					t.Fatal()
				}
				for _, v := range list {
					f := f.Hare(v)
					if false == f.Ok() || true == f.done() {
						t.Fatal()
					}
				}
			}
		}
	}
}

func TestBranchingDetector_Hare_singleCycle(t *testing.T) {
	input := map[int]map[int]map[int][]int{
		1: {
			2: {
				3: {4},
				4: {5},
				5: {4},
			},
		},
		2: {
			10: {
				3: {4},
				4: {5},
				5: {4},
			},
			1: {
				2: {1},
			},
			5: {
				1: {3},
			},
		},
		3: {
			5: {
				1: {2},
			},
		},
	}

	f := NewBranchingDetector(0, nil)

	count := 0

	for x, xMap := range input {
		f := f.Hare(x)
		if false == f.Ok() || true == f.done() {
			t.Fatal()
		}
		for y, yMap := range xMap {
			f := f.Hare(y)
			if false == f.Ok() || true == f.done() {
				t.Fatal()
			}
			for z, list := range yMap {
				f := f.Hare(z)
				if false == f.Ok() || true == f.done() {
					t.Fatal()
				}
				for _, v := range list {
					f := f.Hare(v)
					if true == f.done() {
						t.Fatal()
					}
					if false == f.Ok() {
						count++
						if x != 2 || 1 != y || z != 2 || v != 1 {
							t.Fatal()
						}
					}
				}
			}
		}
	}

	if 1 != count {
		t.Fatal()
	}
}

func TestBranchingDetector_Hare_done(t *testing.T) {
	f := NewBranchingDetector(22, nil)
	f.f.done = true
	f = f.Hare("NO")
	n, ok := f.f.next("wat")
	if true == ok || nil != n {
		t.Fatal()
	}
}

func TestBranchingDetector_Hare_ok(t *testing.T) {
	f := NewBranchingDetector(22, nil)
	f.f.ok = true
	f = f.Hare("NO")
	n, ok := f.f.next("wat")
	if true == ok || nil != n {
		t.Fatal()
	}
}

func TestDetector_SetCompare1(t *testing.T) {
	f := Detector{}
	func() {
		defer func() {
			if err, ok := recover().(error); false == ok || nil == err ||
				"[Detector.validate] nil property encountered, use the constructor NewDetector" != err.Error() {
				t.Fatal()
			}
		}()
		f.SetCompare(nil)
		t.Fatal()
	}()
}

func TestDetector_SetNext1(t *testing.T) {
	f := Detector{}
	func() {
		defer func() {
			if err, ok := recover().(error); false == ok || nil == err ||
				"[Detector.validate] nil property encountered, use the constructor NewDetector" != err.Error() {
				t.Fatal()
			}
		}()
		f.SetNext(nil)
		t.Fatal()
	}()
}

func TestDetector_SetCompare2(t *testing.T) {
	f := NewDetector(1, emptyNext, nil)
	func() {
		defer func() {
			if err, ok := recover().(error); false == ok || nil == err ||
				"[Detector.SetCompare] you cannot set a nil compare" != err.Error() {
				t.Fatal()
			}
		}()
		f.SetCompare(nil)
		t.Fatal()
	}()
}

func TestDetector_SetNext2(t *testing.T) {
	f := NewDetector(1, emptyNext, nil)
	func() {
		defer func() {
			if err, ok := recover().(error); false == ok || nil == err ||
				"[Detector.SetNext] you cannot set a nil next" != err.Error() {
				t.Fatal()
			}
		}()
		f.SetNext(nil)
		t.Fatal()
	}()
}

func TestDetector_SetNext(t *testing.T) {
	f := NewDetector(1, emptyNext, nil)
	f = f.SetNext(func(a interface{}) (interface{}, bool) {
		if 12 != a {
			t.Fatal()
		}
		return "aa", true
	})
	if nil == f.next {
		t.Fatal()
	}
	n, ok := f.next(12)
	if false == ok || "aa" != n {
		t.Fatal()
	}
}

func TestDetector_SetCompare(t *testing.T) {
	f := NewDetector(1, emptyNext, nil)
	f = f.SetCompare(func(a, b interface{}) bool {
		if 1 != a || 2 != b {
			t.Fatal()
		}
		return true
	})
	if nil == f.compare || nil == f.next {
		t.Fatal()
	}
	if true != f.compare(1, 2) {
		t.Fatal()
	}
}

func TestBranchingDetector_branchingNext(t *testing.T) {
	input := map[int]map[int]map[int][]int{
		1: {
			2: {
				3: {4},
				4: {5},
				5: {4},
			},
		},
		2: {
			10: {
				3: {4},
				4: {5},
				5: {4},
			},
			1: {
				-1: {-12, 22, 2321, 2323, 322, 3, 41},
				2:  {1, 2},
			},
			5: {
				1: {3},
			},
		},
		3: {
			5: {
				1: {2},
			},
		},
	}

	f := NewBranchingDetector(0, nil)
	f = f.Hare("1")
	f = f.Hare("2")
	if "1" != f.f.tortoise || "2" != f.f.hare {
		t.Fatal()
	}
	count := 0
	for x := 0; x < 10; x++ {
		for x, xMap := range input {
			f := f.Hare(x)
			if false == f.Ok() || true == f.done() {
				t.Fatal()
			}
			for y, yMap := range xMap {
				f := f.Hare(y)
				if false == f.Ok() || true == f.done() {
					t.Fatal()
				}
				for z, list := range yMap {
					f := f.Hare(z)
					if false == f.Ok() || true == f.done() {
						t.Fatal()
					}
					for _, v := range list {
						f := f.Hare(v)
						if true == f.done() {
							t.Fatal()
						}
						if false == f.Ok() {
							count++
							if x != 2 || 1 != y || z != 2 {
								t.Fatalf("%v %v %v %v", x, y, z, f)
							}
						}
					}
				}
			}
		}
	}

	if 10 != count {
		t.Fatalf("%v", count)
	}
}

func TestNextUpdater_logNext(t *testing.T) {
	var u *nextUpdater
	n, ok := u.logNext(22)
	if false != ok || nil != n {
		t.Fatal()
	}
	u = new(nextUpdater)
	n, ok = u.logNext(22)
	if false != ok || nil != n {
		t.Fatal()
	}
}

func TestNextUpdater_updateNext(t *testing.T) {
	f := NewBranchingDetector(1, nil)
	f.next = []interface{}{2, 3}
	var u *nextUpdater = nil
	f = u.updateNext(f)
	if 2 != len(f.next) {
		t.Fatal()
	}
}

func TestNewBranchingDetector_setsCompare(t *testing.T) {
	f := NewBranchingDetector(1, func(tortoise, hare interface{}) bool {
		return true
	})
	f = f.Hare(2)
	if false == f.Ok() || true == f.done() {
		t.Fatal()
	}
	f = f.Hare(3)
	if true == f.Ok() || true == f.done() {
		t.Fatal()
	}
}

// Generate a map[string]interface{} tree will be formed from countMap maps, (plus one for the origin) with (at most)
// countCycles cycles, and will have countLeaves leaf nodes of type integer, starting at 0, but placed in random
// locations.
// TODO: fix countCycles
// You should probs call rand.Seed(time.Now().Unix()) beforehand.
func generateCycleMap(countMap, countCycles, countLeaves int) map[string]interface{} {
	// generate a list of all the maps, including the origin
	mapList := make([]map[string]interface{}, countMap+1)
	for i := range mapList {
		mapList[i] = make(map[string]interface{})
	}
	// set the leaves at random on the maps
	for x := 0; x < countLeaves; x++ {
		mapList[rand.Intn(len(mapList))][fmt.Sprintf("leaf_%v", x)] = x
	}
	genName := func(ind int) string {
		return fmt.Sprintf("branch_%v", ind)
	}
	// add some cycles
	for x := 0; x < countCycles; x++ {
		mapList[1]["branch_2"] = mapList[2]
		mapList[2]["branch_3"] = mapList[3]
		mapList[3]["branch_1"] = mapList[1]
	}
	originList := []map[string]interface{}{mapList[0]}
	// add node to the origin directly or indirectly
	leftMap := make(map[int]map[string]interface{})
	for k, v := range mapList {
		if 0 == k {
			continue
		}
		leftMap[k] = v
	}
	for 0 != len(leftMap) {
		target := rand.Intn(len(leftMap))
		x := 0
		for k, v := range leftMap {
			if x == target {
				target := originList[rand.Intn(len(originList))]
				name := genName(k)
				if _, ok := target[name]; true == ok {
					break
				}
				target[name] = v
				delete(leftMap, k)
				originList = append(originList, v)
				break
			}
			x++
		}
	}
	return mapList[0]
}

func mapHasCycle(m map[string]interface{}, f BranchingDetector, callClear bool) (bool, func() bool) {
	checkAllNil := func() bool {
		for _, v := range f.next {
			if nil == v {
				continue
			}
			//fmt.Printf("%v", f)
			return false
		}
		return true
	}

	for k, v := range m {
		nf := f.Hare(k)
		if true == callClear {
			defer nf.Clear()
		}
		if false == f.Ok() {
			return true, checkAllNil
		}
		nm, ok := v.(map[string]interface{})
		if false == ok {
			continue
		}
		ok, c := mapHasCycle(nm, nf, callClear)
		oldc := checkAllNil
		checkAllNil = func() bool {
			return oldc() && c()
		}
		if true == ok {
			return true, checkAllNil
		}
	}
	return false, checkAllNil
}

func TestReallocationOfSliceChecking(t *testing.T) {
	slice := []int{0}
	for x := 1; x < 500; x++ {
		next := append(slice, x)
		// if the capacity of next is GREATER than the previous slice's - 1, then it re-allocated
		reallocated := cap(next) > cap(slice)
		slice[0] = -1
		actualReallocated := next[0] != -1
		//fmt.Printf("%v %v %v %v %v %v\n", reallocated, actualReallocated, cap(slice), cap(next), slice, next)
		if actualReallocated != reallocated {
			t.Fatalf("%v %v %v %v %v %v", reallocated, actualReallocated, cap(slice), cap(next), slice, next)
		}
		slice[0] = 0
		next[0] = 0
		slice = next
	}
}

func TestNewBranchingDetector2(t *testing.T) {
	rand.Seed(41212399)
	// verify no cycles detected in ones without cycles
	for x := 0; x < 10; x++ {
		m := generateCycleMap(50, 0, 120)
		if a, b := mapHasCycle(m, NewBranchingDetector(nil, nil), true); false != a || false == b() {
			s, _ := json.MarshalIndent(m, "", "    ")
			t.Fatal(string(s))
		}
	}
	// test detecting cycles
	for x := 0; x < 200; x++ {
		m := generateCycleMap(99, 1, 120)
		if a, b := mapHasCycle(m, NewBranchingDetector(nil, nil), true); true != a || false == b() {
			s, _ := json.MarshalIndent(m, "", "    ")
			t.Fatal(string(s))
		}
	}
	// test detecting cycles, but without clearing the internal array
	for x := 0; x < 10; x++ {
		m := generateCycleMap(50, 0, 120)
		if a, b := mapHasCycle(m, NewBranchingDetector(nil, nil), false); false != a || true == b() {
			s, _ := json.MarshalIndent(m, "", "    ")
			t.Fatal(string(s))
		}
	}
	// test detecting cycles
	for x := 0; x < 200; x++ {
		m := generateCycleMap(99, 1, 120)
		if a, b := mapHasCycle(m, NewBranchingDetector(nil, nil), false); true != a || true == b() {
			s, _ := json.MarshalIndent(m, "", "    ")
			t.Fatal(string(s))
		}
	}
}
