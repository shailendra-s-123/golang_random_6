package main  
import (  
    "fmt"
    "errors"
)

type dashboard []string

// Custom error type for slice operations
type sliceError struct {
    msg string
}

func (e *sliceError) Error() string {
    return fmt.Sprintf("slice error: %s", e.msg)
}

// Function to append an element to the dashboard
func (d *dashboard) Append(item string) error {
    if d == nil {
        return &sliceError{"cannot append to nil dashboard"}
    }
    *d = append(*d, item)
    return nil
}

// Function to get an element from the dashboard by index
func (d dashboard) Get(index int) (string, error) {
    if d == nil {
        return "", &sliceError{"cannot get from nil dashboard"}
    }
    if index < 0 || index >= len(d) {
        return "", &sliceError{"index out of bounds"}
    }
    return d[index], nil
}

// Function to update an element in the dashboard by index
func (d *dashboard) Update(index int, item string) error {
    if d == nil {
        return &sliceError{"cannot update nil dashboard"}
    }
    if index < 0 || index >= len(*d) {
        return &sliceError{"index out of bounds"}
    }
    (*d)[index] = item
    return nil
}

func main() {
    var d dashboard  // Initialize a nil dashboard
    
    // Test append
    err := d.Append("Widget A")
    if err != nil {
        fmt.Println("Error appending to dashboard:", err)
    }
    
    // Test get (from nil dashboard)
    item, err := d.Get(0)
    if err != nil {
        fmt.Println("Error getting from dashboard:", err)
    }
    fmt.Println("Item at index 0:", item)
    
    // Test update (invalid index)
    err = d.Update(-1, "Widget B")
    if err != nil {
        fmt.Println("Error updating dashboard:", err)
    }
    
    // Create a valid dashboard and test operations
    d = dashboard{"Widget C", "Widget D"}
    
    // Test get
    item, err = d.Get(0)
    if err != nil {
        fmt.Println("Error getting from dashboard:", err)
    }
    fmt.Println("Item at index 0:", item)
    
    // Test update
    err = d.Update(1, "Widget E")
    if err != nil {
        fmt.Println("Error updating dashboard:", err)
    }
    fmt.Println("Dashboard after update:", d)
}