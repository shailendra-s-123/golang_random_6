package main
import (
	"fmt"
	"sync"
	"testing"
)
type Counter struct {
	value int
	mutex sync.Mutex
}
func (c *Counter) Get() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}
func (c *Counter) Set(value int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value = value
}
func (c *Counter) Delete() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value = 0
}
func TestCounter_Get(t *testing.T) {
	t.Run("Get_DefaultValue", func(t *testing.T) {
		c := &Counter{}
		got := c.Get()
		want := 0
		if got != want {
			t.Errorf("Counter.Get() = %v, want %v", got, want)
		}
	})
	t.Run("Get_AfterSet", func(t *testing.T) {
		c := &Counter{}
		c.Set(42)
		got := c.Get()
		want := 42
		if got != want {
			t.Errorf("Counter.Get() = %v, want %v", got, want)
		}
	})
}
func TestCounter_Set(t *testing.T) {
	t.Run("Set_SingleValue", func(t *testing.T) {
		c := &Counter{}
		c.Set(13)
		got := c.Get()
		want := 13
		if got != want {
			t.Errorf("Counter.Get() = %v, want %v", got, want)
		}
	})
	t.Run("Set_MultipleValues", func(t *testing.T) {
		c := &Counter{}
		c.Set(17)
		c.Set(41)
		got := c.Get()
		want := 41
		if got != want {
			t.Errorf("Counter.Get() = %v, want %v", got, want)
		}
	})
}
func TestCounter_Delete(t *testing.T) {
	t.Run("Delete_AfterSet", func(t *testing.T) {
		c := &Counter{}
		c.Set(23)
		c.Delete()
		got := c.Get()
		want := 0
		if got != want {
			t.Errorf("Counter.Get() = %v, want %v", got, want)
		}
	})
}
func ExampleCounter_Get() {
	c := &Counter{}
	c.Set(100)
	fmt.Println(c.Get()) // Output: 100
}
func ExampleCounter_Set() {
	c := &Counter{}
	c.Set(200)
	fmt.Println(c.Get()) // Output: 200
}
func ExampleCounter_Delete() {
	c := &Counter{}
	c.Set(300)
	c.Delete()
	fmt.Println(c.Get()) // Output: 0
}