package files

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	err := Create("./output.db3")
	fmt.Println(err)
}
