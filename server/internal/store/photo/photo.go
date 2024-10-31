package photo

import (
	"io"
)

type Store interface {
	Put(photoReader io.Reader, photoName string, photoSize int64) (string, error)
	Get(photoName string) (string, error)
}
