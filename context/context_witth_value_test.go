package context

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "test b")
	contextC := context.WithValue(contextA, "c", "test c")

	contextD := context.WithValue(contextB, "d", "test d")
	contextE := context.WithValue(contextB, "e", "test e")

	contextF := context.WithValue(contextC, "f", "test f")

	contextG := context.WithValue(contextF, "g", "test g")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	// get data context
	// hanya bisa akses bawah ke atas, chield ke parent tidak bisa parent ke child
	fmt.Println(contextE.Value("b")) // bisa
	fmt.Println(contextA.Value("e")) // gabisa karena dari parent

}
