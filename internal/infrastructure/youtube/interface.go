package youtube

type Interface interface {
	GetPreview(url string) ([]byte, error)
}
