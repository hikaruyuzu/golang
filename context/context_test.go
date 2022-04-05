package context

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateContext(t *testing.T) {
	contextBackground := context.Background()
	fmt.Println(contextBackground)

	contextTodo := context.TODO()
	fmt.Println(contextTodo)

}
