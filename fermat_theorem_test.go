package fermat_theorem;

import (
    "testing"
    "math"
)

func TestN2(t *testing.T) {
    a := 3.0
    b := 4.0
    c := 5.0

    a2 := math.Pow(a, 2)
    b2 := math.Pow(b, 2)
    c2 := math.Pow(c, 2)
    if a2 + b2 != c2 {
        t.Errorf("3^2 + 4^2 = %f; want 5^2", c2)
    }
}

func TestN3(t *testing.T) {
    a := 3.0
    b := 4.0
    c := 5.0

    a3 := math.Pow(a, 3)
    b3 := math.Pow(b, 3)
    c3 := math.Pow(c, 3)
    if a3 + b3 == c3 {
        t.Errorf("3^3 + 4^3 = %f; don't want", c3)
    }
}

func TestN1(t *testing.T) {
    a := 1.0
    b := 1.0
    c := 2.0

    a1 := math.Pow(a, 1)
    b1 := math.Pow(b, 1)
    c1 := math.Pow(c, 1)
    if a1 + b1 != c1 {
        t.Errorf("1^1 + 1^1 = %f; want 2^1", c1);
    }
}
