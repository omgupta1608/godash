package godash

import (
	"runtime/debug"
)

// Returns the current module version or nil otherwise
func Version() *string {
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		return &buildInfo.Main.Version
	}
	return nil
}

func main() {

}
