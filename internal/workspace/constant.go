package workspace

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/ohanan/bpg/internal/panics"
)

var WorkDir = filepath.Join(panics.If1(os.UserHomeDir()), ".bpg/")

func RegisterFlags() {
	flag.StringVar(&WorkDir, "work-dir", WorkDir, "working directory")
}
