package storage

type Storage interface {
	PutPublicPhoto(photoContent string) (string, error)
	GetPublicPhoto(name string) (string, error)
}
