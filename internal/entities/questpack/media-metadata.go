package questpack

type MediaType int

const (
	MediaTypeImage MediaType = 0
	MediaTypeAudio MediaType = 1
	MediaTypeVideo MediaType = 2
)

type MediaMetadata struct {
	Path string    `json:"path"`
	Type MediaType `json:"type"`
	Size int64     `json:"size"`
}

func NewMediaMetadata(path string, size int64) *MediaMetadata {
	return &MediaMetadata{
		Path: path,
		Type: MediaTypeImage,
		Size: size,
	}
}
