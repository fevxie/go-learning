package methods

import "fmt"

type triangle struct {
    size int
}

type square struct {
    size int
}

func (t triangle) perimeter() int {
    return t.size * 3
}
func (t *triangle) doubleSize() {
    t.size *= 2
}

func (s square) perimeter() int {
    return s.size * 4
}
func main() {
    t := triangle{3}
    s := square{4}
    t.doubleSize()
    fmt.Println("Perimeter(square):", s.perimeter())
    fmt.Println("Perimeter:", t.perimeter())
}
