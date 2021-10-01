package version

import "runtime/debug"

// String returns a version of this package.
func String() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	} else {
		return "undefined"
	}
}
