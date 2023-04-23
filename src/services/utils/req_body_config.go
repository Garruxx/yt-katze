package utils

import (
	"encoding/json"
	"fmt"
	"katze/src/services/models"
)

type client struct {
	ClientName    string `json:"clientName"`
	ClientVersion string `json:"clientVersion"`
	Hl            string `json:"hl,omitempty"`
	GL            string `json:"gl,omitempty"`
}

type context struct {
	Client client `json:"client"`
}

type requestConfig struct {
	Context  context `json:"context"`
	Params   string  `json:"params,omitempty"`
	Query    string  `json:"query,omitempty"`
	BrowseID string  `json:"browseId,omitempty"`
	VideoID  string  `json:"videoId,omitempty"`
}

func ReqBodyConfig(config models.BodyReqConfig) ([]byte, error) {

	configuration := requestConfig{
		Context: context{
			Client: client{
				ClientName:    "WEB_REMIX",
				ClientVersion: "1.20230306.01.00",
				Hl:            "en",
			},
		},
		Params:   config.Params,
		Query:    config.Query,
		BrowseID: config.BrowseID,
		VideoID:  config.VideoID,
	}

	jsonConfig, err := json.Marshal(configuration)
	if err != nil {
		err := fmt.Errorf("error marshalling body config: %v", err)
		return nil, err
	}
	return jsonConfig, nil
}
