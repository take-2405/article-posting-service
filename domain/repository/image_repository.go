package repository

type ImageRepository interface {
	CreateImage([]string) error
	DeleteImage() error
	SearchImages() error
	SendImages() error
}
