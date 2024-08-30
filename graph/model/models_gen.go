// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Asset struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	TagValues []*TagValue `json:"tagValues"`
}

type DeleteTagValue struct {
	ID string `json:"id"`
}

type Mutation struct {
}

type NewAsset struct {
	Name string `json:"name"`
}

type NewTag struct {
	Name string `json:"name"`
}

type NewTagValue struct {
	ID      string `json:"id"`
	TagID   string `json:"tagID"`
	AssetID string `json:"assetID"`
	Value   string `json:"value"`
}

type Query struct {
}

type Tag struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type TagValue struct {
	ID    string `json:"id"`
	Tag   *Tag   `json:"tag"`
	Asset *Asset `json:"asset"`
	Value string `json:"value"`
}

type UpdateTagValue struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}
