package pokecache

import ("testing"
		"fmt"
		"time")

func TestCache (t *testing.T) {
	cases := []struct {
		input    string
		expected []byte
	}{
		{
			input: "https://example.com",
			expected: []byte("testdata"),
		},
		{
			input: "https://example.com/path",
			expected: []byte("moretestdata"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			interval := 5* time.Second
			cache := NewCache(interval)
			cache.Add(c.input, c.expected)
			val, ok := cache.Get(c.input)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.expected) {
				t.Errorf("expected to find value")
				return
			}
		})
	}	
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(5 * time.Millisecond)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)
	
	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}