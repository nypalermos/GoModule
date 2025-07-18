package GoModule

import "testing"

func TestGreet(t *testing.T) {
    got := Greet("Leon")
    want := "Hello, Leon!"
    if got != want {
        t.Errorf("Greet() = %q; want %q", got, want)
    }
}
