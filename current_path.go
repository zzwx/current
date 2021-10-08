package current

import (
	"github.com/zzwx/iffound"
	"io"
	"path"
	"path/filepath"
)

// Path is a type can be directly included into another type,
// to merge Join, Reader, Path methods, where the first call to any of these methods
// will initialize the relative path (or absolute for main module and tests),
// based on where it was called from.
//
// Path can also be used as a type for a variable, in which case
// NewPath will initialize the path and return a newly constructed value.
type Path struct {
	path string
}

// NewPath may be used to initialize the relative path
// as a variable to control its actual value, based on
// where it was called from.
//
// Instead of calling NewPath, Path can be included in the struct.
func NewPath() *Path {
	rp := new(Path)
	rp.init()
	return rp
}

func (rp *Path) init() {
	if rp.path == "" {
		rp.path = WhereAmI(3) // 3 is to address that init() adds to the call stack level
	}
}

// Path returns the original Path's path, cleaned and converted to
// forward slash. Call to Join will not modify it.
func (rp *Path) Path() string {
	rp.init() // Don't remove
	return filepath.ToSlash(path.Clean(rp.path))
}

// Join returns a joined absolute (sometimes relative) path from Path and does not modify Path.
// It accepts any amount of paths. The result will be turned into filepath.ToSlash to
// make sure all paths are using forward slash.
//
// ".." and "." in the path for joining are supported.
func (rp *Path) Join(with ...string) string {
	rp.init() // Don't remove
	args := append([]string{rp.path}, with...)
	return filepath.ToSlash(path.Join(args...))
}

// Reader acts as Join and then opens the file, and returns an io.Reader from it,
// registering a delayed close on file upon garbage collection. It also returns a
// file.ZeroReader if the file is not found. This allows to assert to file.ZeroReader
// to find out that that happened and unwrap the error:
//
//	if zr, ok := reader.(file.ZeroReader); ok {
//
// ".." and "." in the path for joining are supported.
func (rp *Path) Reader(with ...string) io.Reader {
	rp.init() // Don't remove
	joined := rp.Join(with...)
	return iffound.IfFound(joined).Reader()
}
