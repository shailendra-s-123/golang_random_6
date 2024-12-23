package main

import (
	"testing"
)

func TestDashboardManager(t *testing.T) {
	dm := newDashboardManager()

	// Test append
	err := dm.Append("Widget A")
	if err != nil {
		t.Errorf("Error appending: %v", err)
	}

	err = dm.Append("Widget B")
	if err != nil {
		t.Errorf("Error appending: %v", err)
	}

	// Test get
	item, err := dm.Get(0)
	if err != nil || item != "Widget A" {
		t.Errorf("Error getting item at index 0: %v, item: %s", err, item)
	}

	item, err = dm.Get(1)
	if err != nil || item != "Widget B" {
		t.Errorf("Error getting item at index 1: %v, item: %s", err, item)
	}

	// Test update
	err = dm.Update(0, "Updated Widget A")
	if err != nil {
		t.Errorf("Error updating: %v", err)
	}

	item, err := dm.Get(0)
	if err != nil || item != "Updated Widget A" {
		t.Errorf("Error getting item at index 0 after update: %v, item: %s", err, item)
	}

	// Test remove
	err = dm.Remove(0)
	if err != nil {
		t.Errorf("Error removing: %v", err)
	}

	item, err := dm.Get(0)
	if err != nil || item != "Widget B" {
		t.Errorf("Error getting item at index 0 after remove: %v, item: %s", err, item)
	}

	// Test remove on empty dashboard
	err = dm.Remove(0)
	if err == nil {
		t.Errorf("Expected error removing from empty dashboard")
	}

	// Test append on nil dashboard
	dm = nil
	err = dm.Append("Widget C")
	if err == nil {
		t.Errorf("Expected error appending to nil dashboard")
	}
}