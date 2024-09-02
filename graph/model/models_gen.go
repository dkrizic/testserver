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

type NewTagValue struct {
	TagID   string `json:"tagID"`
	AssetID string `json:"assetID"`
	Value   string `json:"value"`
}

type Query struct {
}

type Search struct {
	Text            string `json:"text"`
	SearchTagName   bool   `json:"searchTagName"`
	SearchTagValue  bool   `json:"searchTagValue"`
	SearchAssetName bool   `json:"searchAssetName"`
}

type SearchResult struct {
	Tag      []*Tag      `json:"tag"`
	Asset    []*Asset    `json:"asset"`
	TagValue []*TagValue `json:"tagValue"`
}

type Tag struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Assets []*Asset `json:"assets"`
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
