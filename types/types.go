package types

type ImageInfo struct {
	OriginalName string `json:"original_name"`
	Architecture string `json:"architecture"`
	Tag          string `json:"tag"`
}
