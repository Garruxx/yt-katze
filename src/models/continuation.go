package models

type Continuation struct {
	Query  string  `json:"query"`
	Params *string `json:"params,omitempty"`
}
