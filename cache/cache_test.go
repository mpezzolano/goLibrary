package cache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := NewCache()

	// Test Set and Get
	c.Set("key1", "value1", 1*time.Minute)
	value, found := c.Get("key1")
	if !found || value != "value1" {
		t.Errorf("Expected to find 'value1' for 'key1', but got '%v'", value)
	}

	// Test Expiration
	c.Set("key2", "value2", 1*time.Second)
	time.Sleep(2 * time.Second)
	_, found = c.Get("key2")
	if found {
		t.Errorf("Expected 'key2' to be expired, but it was found")
	}

	// Test Delete
	c.Set("key3", "value3", 1*time.Minute)
	c.Delete("key3")
	_, found = c.Get("key3")
	if found {
		t.Errorf("Expected 'key3' to be deleted, but it was found")
	}

	// Test Clear
	c.Set("key4", "value4", 1*time.Minute)
	c.Clear()
	_, found = c.Get("key4")
	if found {
		t.Errorf("Expected cache to be cleared, but 'key4' was found")
	}
}
