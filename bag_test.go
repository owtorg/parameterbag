package parameterbag

import (
	"testing"
)

func TestBag(t *testing.T) {

	//Call it like this to test that we are actually satisfying our interface
	var _ Bag = NewParameterBag()

	bag := NewParameterBag()
	bag.Set("test", "testval")

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

	bag.Set("take me down", "to funky town")
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
