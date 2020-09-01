package tboxs

import (
	"fmt"
	"testing"
)

func TestDistribute(t *testing.T) {
	test1 := Distribute(100, 3)
	fmt.Printf("test1=%v\n", test1)

	test2 := Distribute(40, 7)
	fmt.Printf("test2=%v\n", test2)

	test3 := Distribute(34.34, 8)
	fmt.Printf("test3=%v\n", test3)
}
