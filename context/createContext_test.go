package context

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println("Context bg", background)

	todo := context.TODO()

	fmt.Println("Context Todo", todo)

}

func TestChildParent(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "INi C")
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextC, "e", "E")
	contextF := context.WithValue(contextC, "f", "ini context F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("apa ajas"))
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("b"))

}
