package GoModule

import "testing"

func TestGreet(t *testing.T) {
    got := Greet("Leon", "Palermo")
    want := "Hi, Leon Palermo!"
    if got != want {
        t.Errorf("Greet() = %q; want %q", got, want)
    }
}
