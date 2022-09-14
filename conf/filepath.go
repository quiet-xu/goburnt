package conf

import (
	"os"
	"os/exec"
	"path/filepath"
)

func LocalFileAuto(file string) string {
	if filepath.IsAbs(file) {
		return file
	}
	return LocalFile(file)
}
func LocalFile(file string) string {
	app, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(app)
	return filepath.Join(filepath.Dir(path), file)
}
