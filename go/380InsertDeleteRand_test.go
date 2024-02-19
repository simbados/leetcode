package main

import (
	"testing"
)

func TestInsertDeleteRand(t *testing.T) {
	rand := Constructor()
	insert := rand.Insert(1)
	if !insert {
		t.Errorf("Entry should not exists and insert should return true")
	}
	rand.Insert(2)
	rand.Insert(3)
	remove := rand.Remove(3)
	if !remove {
		t.Errorf("Should be ablte to remove from set and return true")
	}
	if rand := rand.GetRandom(); rand != 2 && rand != 1 {
		t.Errorf("Rand did not return a value, instead: %v", rand)
	}
}
