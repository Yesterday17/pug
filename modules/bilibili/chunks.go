package bilibili

import "os"

type chunk struct {
	index bigInt
	total bigInt

	size      bigInt
	start     bigInt
	end       bigInt
	totalSize bigInt
}

func (m *Module) SplitChunks(file *os.File) {

}
