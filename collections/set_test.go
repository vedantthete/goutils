package collections

import "testing"

func TestSetLength(t *testing.T) {
	set := NewSet[string]()
	set.Add("a")
	result := set.Length()
	if result != 1 {
		t.Errorf("Expected 1 received %v", result)
	}
	set.Remove("a")
	result = set.Length()
	if result != 0 {
		t.Errorf("Expected 0 received %v", result)
	}
	set.Add("a")
	set.Add("a")
	result = set.Length()
	if result != 1 {
		t.Errorf("Expected 1 received %v", result)
	}
}

func TestSetHas(t *testing.T) {
	set := NewSet[string]()
	set.Add("a")
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"Should contain", "a", true},
		{"Should not contain", "b", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ans := set.Has(test.input)
			if ans != test.want {
				t.Errorf("got %v, want %v", ans, test.want)
			}
		})
	}
}

func TestSetRemove(t *testing.T) {
	set := NewSet[string]()
	set.Add("a")
	set.Add("b")
	set.Remove("a")
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"Should contain", "b", true},
		{"Should not contain", "a", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ans := set.Has(test.input)
			if ans != test.want {
				t.Errorf("got %v, want %v", ans, test.want)
			}
		})
	}
}

func TestSetDiscard(t *testing.T) {
	set := NewSet[string]()
	set.Add("a")
	ans := set.Has("a")
	if ans != true {
		t.Errorf("Adding to set failed")
	}
	set.Discard("a")
	ans = set.Has("a")
	if ans == true {
		t.Errorf("Discard failed")
	}
}
