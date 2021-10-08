package current

import (
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
)

// WhereAmI returns an absolute (in rare cases relative) path for the
// caller source location of the function.
//
// If no number provided, "1" is assumed for the depth of the caller stack.
// "2" will go upper, "0" will return the WhereAmI's location itself.
// If unable to retrieve, returns "".
func WhereAmI(depths ...int) string {
	var depth = 1
	if len(depths) > 0 {
		depth = depths[0]
	}
	if function, file, _, ok := runtime.Caller(depth); ok { // function, file, line, _ := runtime.Caller(depth)
		callerMainModule := CallerMainModulePath()
		callerFileDir := filepath.ToSlash(filepath.Dir(file))
		funcMain := runtime.FuncForPC(function).Name()
		// Executable may be in Temp directory, not referring to actual source folders
		// ex, err := os.Executable()
		_ = callerMainModule
		_ = callerFileDir
		_ = funcMain
		if callerFileDir != "" {
			// Seems to be the most reliable absolute path
			return callerFileDir
		}
		if funcMain == "main.main" { // When called from main()
			panic("current path: failure to determine")
		}
		// Getwd() may not actually work properly
		//wd, err := os.Getwd()
		//if err != nil {
		//	panic("can't read working directory: " + err.Error())
		//}
		//return filepath.ToSlash(wd)
		funcMain = strings.TrimPrefix(funcMain, callerMainModule)
		if strings.HasPrefix(funcMain, "/") {
			return "." + strings.Split(funcMain, ".")[0]
		} else {
			return "./" + strings.Split(funcMain, ".")[0]
		}
	}
	return ""
}

// CallerMainModulePath returns the main module path of the program that
// uses this module, or "" if not unable to retrieve.
func CallerMainModulePath() string {
	if bi, ok := debug.ReadBuildInfo(); ok {
		return bi.Main.Path
	}
	return ""
}

// JoinPath join paths with forward slash and  retains the front ./ dot
// represented originally with either ./ or .\
//
// Additionally simply . and ./ return .
func JoinPath(with ...string) string {
	prependDot := ""
	if len(with) > 0 {
		switch {
		case with[0] == "./" || with[0] == ".\\":
			prependDot = ""
		case strings.HasPrefix(with[0], "./") || strings.HasPrefix(with[0], ".\\"):
			prependDot = "." + string(os.PathSeparator)
		}
	}
	return filepath.ToSlash(prependDot + filepath.Join(with...))
}
