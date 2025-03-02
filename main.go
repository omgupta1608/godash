package godash

import (
	"fmt"
	"runtime/debug"
	"strings"
)

// Returns the current module version or nil otherwise
func Version() *string {
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		return &buildInfo.Main.Version
	}
	return nil
}

func main() {
	a := []int{1, 2, 3, 5}
	b := []int{4, 6, 7, 10, 1, 2, 3, 5, 1}
	c := []string{}
	// chunk := 1
	// // fmt.Println(Chunk(a, &chunk))
	// fmt.Println(Concat(a, 2, 2, 5, 2, 3, 4, 2, 3), a)
	// fmt.Println(DropRight(a, &chunk))

	fmt.Println(Intersection(a, b))
	fmt.Print(Join(c, "|"), "\n", strings.Join(c, "|"))

	DropRight([]string{}, Int(3))
}
