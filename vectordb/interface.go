package vectordb

import "context"

type Fragment struct {
	Id           string `json:"id,omitempty"`
	Text         string `json:"text,omitempty"`
	UserId       string `json:"user_id,omitempty"`
	DocumentId   string `json:"document_id,omitempty"`
	CollectionId string `json:"collection_id,omitempty"`
	Position     uint32 `json:"position,omitempty"`
}

type SearchQuery struct {
	UserId       string  `json:"user_id,omitempty"`
	CollectionId string  `json:"collection_id,omitempty"`
	Query        string  `json:"query,omitempty"`
	Limit        uint32  `json:"limit,omitempty"`
	Threshold    float32 `json:"threshold,omitempty"`
}

type SearchResults struct {
	Fragments []*Fragment `json:"fragments,omitempty"`
	Scores    []float32   `json:"scores,omitempty"`
}

type DB interface {
	Search(context.Context, SearchQuery) (*SearchResults, error)
	Upsert(context.Context, []*Fragment) error
	Delete(ids []string) error
	Close() error
}
