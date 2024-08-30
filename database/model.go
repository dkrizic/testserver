package database

type TagValue struct {
	ID      string `json:"id"`
	TagID   string `json:"tagID"`
	AssetID string `json:"assetID"`
	Value   string `json:"value"`
}
