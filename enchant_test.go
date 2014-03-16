package enchant

import "testing"

func TestDictExists(t *testing.T) {
	e, err := NewEnchant()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	defer e.Free()

	type Expect struct {
		give string
		want bool
	}

	cases := []Expect{
		Expect{"en", true},
		Expect{"en_GB", true},
		Expect{"none", false},
		Expect{"a", false},
	}

	for _, c := range cases {
		got := e.DictExists(c.give)
		if got != c.want {
			t.Errorf("Wanted DictExists to return %v for \"%v\", but got %v", c.want, c.give, got)
		}
	}
}

func TestLoadDictAndCheck(t *testing.T) {
	e, err := NewEnchant()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	defer e.Free()

	// load the english dictionary
	e.LoadDict("en")

	// check some words
	type Expect struct {
		give string
		want bool
	}

	cases := []Expect{
		Expect{"test", true},
		Expect{"yes", true},
		Expect{"amazing", true},
		Expect{"yoyoyoyo", false},
		Expect{"nosuchword", false},
	}

	for _, c := range cases {
		got := e.Check(c.give)
		if got != c.want {
			t.Errorf("Wanted Check to return %v for \"%v\", but got %v", c.want, c.give, got)
		}
	}
}
