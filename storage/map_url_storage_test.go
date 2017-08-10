package storage

import "testing"

func TestAddURLToMap(t *testing.T) {
	storage := MakeMapURLStorage()
	values := map[int]string{
		0: "one",
		1: "two",
		2: "three",
	}
	for _, v := range values {
		if _, err := storage.Add(v); err != nil {
			t.Fatalf("Failed during adding values. %s", err.Error())
		}
	}
	for k, v := range values {
		storedValue, ok := storage.storage[k]
		if !ok {
			t.Fatalf("Key %s was not found", k)
		}
		if storedValue != v {
			t.Fatalf("Expected value %s for key %s, Was: %s", v, k, storedValue)
		}
	}
}

func TestDeleteURLFromMap(t *testing.T) {
	storage := MakeMapURLStorage()
	values := map[int]string{
		0: "one",
		1: "two",
	}
	for _, v := range values {
		if _, err := storage.Add(v); err != nil {
			t.Fatalf("Failed during adding values. %s", err.Error())
		}
	}
	storage.Delete(0)

	if _, ok := storage.storage[0]; ok {
		t.Fatal("The key 'one' was not removed")
	}
	if _, ok := storage.storage[1]; !ok {
		t.Fatal("The key 'two' should not have been removed")
	}
}

func TestGetValuesFromMapStorage(t *testing.T) {
	storage := MakeMapURLStorage()
	values := map[int]string{
		0: "one",
		1: "two",
		2: "three",
	}
	for _, v := range values {
		if _, err := storage.Add(v); err != nil {
			t.Fatalf("Failed during adding values. %s", err.Error())
		}
	}
	for k, v := range values {
		storedValue, err := storage.Get(k)
		if err != nil {
			t.Fatalf("Error getting value %s", err.Error())
		}
		if storedValue != v {
			t.Fatalf("Expected value %s from key %s. Got %s", v, k, storedValue)
		}
	}
}
