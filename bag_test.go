package parameterbag

import (
	"testing"
)

func TestBag(t *testing.T) {

	//Call it like this to test that we are actually satisfying our interface
	var _ Bag = NewParameterBag()

	bag := NewParameterBag()
	err := bag.Set("test", "testval")
	if err != nil {
		t.Error("could not set value in new bag", err.Error())
	}

	tv := bag.Get("test")
	if tv != "testval" {
		t.Error("key test should equal string testval")
	}

	tp := bag.Get("cornolio")
	if tp != "" {
		t.Error("expected blank string got", tp)
	}

	if bag.Has("tp") {
		t.Error("bag should not have key tp")
	}

	if !bag.Has("test") {
		t.Error("bag should have key test")
	}

	plist := bag.ListParameters()

	if len(plist) != 1 {
		t.Error("bag should have 1 item")
	}

}

func TestBagFromMap(t *testing.T) {
	params := make(map[string]string)
	params["banana"] = "mango"
	params["flying_in"] = "a blue dream"

	var _ Bag = NewParameterBagFromMap(params)

	var bag = NewParameterBagFromMap(params)

	if !bag.Has("banana") {
		t.Error("bag should contain key:", "banana")
	}

	if !bag.Has("flying_in") {
		t.Error("bag should contain key:", "flying_in")
	}

	err := bag.Set("take me down", "to funky town")
	if err != nil {
		t.Error("could not set value in new bag", err.Error())
	}

	if !bag.Has("take me down") {
		t.Error("bag should contain key:", "take me down")
	}

	town := bag.Get("take me down")

	if town != "to funky town" {
		t.Error("could not take you down to funky town")
	}

	plist := bag.ListParameters()

	if len(plist) != 3 {
		t.Error("bag should have 3 items")
	}
}

func TestMutableCopy(t *testing.T) {

	bag := NewParameterBag()
	err := bag.Set("test", "testval")
	if err != nil {
		t.Error("could not set value in new bag", err.Error())
	}
	bag.Freeze()
	if !bag.IsFrozen() {
		t.Error("bag should report as frozen")
	}
	err = bag.Set("tom_goes", "to the mayor")
	if err == nil {
		t.Error("bag should be immutable")
	}

	mutable := bag.GetMutableCopy()
	if !mutable.Has("test") {
		t.Error("mutable copy does not contain key from origin")
	}

	if mutable.Get("test") != "testval" {
		t.Error("mutable copy value was not set correctly.    ", "mutable:", mutable.Get("test"), "original:", bag.Get("test"))
	}

	err = mutable.Set("tom_goes", "to the mayor")
	if err != nil {
		t.Error("bag should be mutable")
	}

	err = mutable.Set("test", "new value")
	if err != nil {
		t.Error("bag should allow key test to be changed")
	}

	//check to see if the original bag changed when we set a key in the mutable bag.
	//If so a shallow copy took place
	if bag.Get("test") != "testval" {
		t.Error("deep copy of original bag was not taken!   ", "mutable:", mutable.Get("test"), "original:", bag.Get("test"))
	}

	mutable.Freeze()
	if !mutable.IsFrozen() {
		t.Error("bag should report as frozen")
	}
	err = mutable.Set("tom_goes", "to the mayor")
	if err == nil {
		t.Error("mutable bag should have been made immutable")
	}

}
