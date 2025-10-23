package iteration

import (
	"fmt"
	"strings"
)

func Contains(s, substr string) bool {
	fmt.Println(s, substr)
	return strings.Contains(s, substr)
}
