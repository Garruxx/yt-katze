package models

type Continuation struct {
	Query  string  `json:"query,omitempty"`
	Params *string `json:"params,omitempty"`
}
