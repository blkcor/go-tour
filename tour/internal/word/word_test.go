package word

import (
	"fmt"
	"testing"
)

func TestUnderscoreToCamelCase(t *testing.T) {
	fmtS := UnderscoreToCamelCase("a_bdddd_c")
	fmt.Println(fmtS)
}
