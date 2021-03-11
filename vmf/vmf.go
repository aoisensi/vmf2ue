package vmf

import (
	"io"

	"github.com/galaco/vmf"
)

func Read(file io.Reader) (*Level, error) {
	reader := vmf.NewReader(file)
	level, err := reader.Read()
	if err != nil {
		return nil, err
	}
	return newLevel(level), nil
}
