[![https://github.com/zzwx/current](./doc/gobadge.svg)](https://pkg.go.dev/github.com/zzwx/current)

An attempt to make modules with resource files act as if they are embedded instead of actually embedding resources. Works in environments when it's known that the module will be built from source.

A module uses `current.NewPath()` to declare a variable or simply `current.Path` as an embedded type to make `.Join(...)` method work from the physical location on the drive of the `.go` source file even when it becomes a downloaded module. The `.Path()` method returns the exact position of the `.go` file on the drive.

The returned paths are always absolute.