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

func TestLoadDictCheckAndSuggest(t *testing.T) {
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

	// test Check
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

	// test suggest
	s := "wowzers"
	got := e.Suggest(s)
	if len(got) != 6 {
		t.Errorf("Expected %v suggestions for \"%v\", but got %v. Value is: %v", 6, s, len(got), got)
	}

	s = "lawn"
	got = e.Suggest(s)
	if len(got) != 14 {
		t.Errorf("Expected %v suggestions for \"%v\", but got %v. Value is: %v", 6, s, len(got), got)
	}
}
