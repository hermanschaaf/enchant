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
			t.Errorf("Wanted DictExists to return %v for %v, but got %v", c.want, c.give, got)
		}
	}
}
