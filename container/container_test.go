package container_test

import (
	"citest/container"
	"testing"
)

func TestAppend(t *testing.T) {
	ct := container.New()

	cases := []struct {
		name    string
		content string
	}{
		{
			name:    "test1",
			content: "test-content1",
		},
		{
			name:    "test2",
			content: "test-content2",
		},
		{
			name:    "test3",
			content: "test-content3",
		},
	}

	for _, c := range cases {
		ct.Append(container.Baggage{Name: c.name, Content: c.content})
	}

	if len(ct.Baggages) != len(cases) {
		t.Errorf("Failed Append() expected: %d, get: %d", len(cases), len(ct.Baggages))
	}
}

func TestFind(t *testing.T) {
	ct := container.New()

	cases := []struct {
		name    string
		content string
	}{
		{
			name:    "test1",
			content: "test-content1",
		},
		{
			name:    "test2",
			content: "test-content2",
		},
		{
			name:    "test3",
			content: "test-content3",
		},
	}

	for _, c := range cases {
		expected := container.Baggage{Name: c.name, Content: c.content}
		ct.Append(expected)
		res := ct.Find(c.name)
		if res != expected {
			t.Errorf("Failed Find() expected: %v, get: %v", expected, res)
		}
	}

	res2 := ct.Find("unknown")
	expected2 := container.Baggage{}
	if res2 != expected2 {
		t.Errorf("Failed Find() expected: %v, get: %v", expected2, res2)
	}
}
