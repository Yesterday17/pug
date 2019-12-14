package temp

import (
	"encoding/hex"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
)

type Dir string

func NewDir() (Dir, error) {
	dir, err := ioutil.TempDir("", "pug")
	if err != nil {
		return "", err
	}
	return Dir(dir), nil
}

func (d Dir) Clear() {
	_ = os.RemoveAll(string(d))
}

func (d Dir) NewFile(ext string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join(string(d), hex.EncodeToString(randBytes)+ext)
}
