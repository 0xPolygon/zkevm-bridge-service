package zkevmbridgeservice

import (
	"fmt"
	"io"
	"runtime"
)

// Populated during build, don't touch!
var (
	Version   = "v0.1.0"
	GitRev    = "undefined"
	GitBranch = "undefined"
	BuildDate = "Fri, 17 Jun 1988 01:58:00 +0200"
)

// PrintVersion prints version info into the provided io.Writer.
func PrintVersion(w io.Writer) {
	_, _ = fmt.Fprintf(w, "Version:      %s\n", Version)
	_, _ = fmt.Fprintf(w, "Git revision: %s\n", GitRev)
	_, _ = fmt.Fprintf(w, "Git branch:   %s\n", GitBranch)
	_, _ = fmt.Fprintf(w, "Go version:   %s\n", runtime.Version())
	_, _ = fmt.Fprintf(w, "Built:        %s\n", BuildDate)
	_, _ = fmt.Fprintf(w, "OS/Arch:      %s/%s\n", runtime.GOOS, runtime.GOARCH)
}
