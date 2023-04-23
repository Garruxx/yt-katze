package tests

import (
	"katze/src/services/models"
	"katze/src/services/utils"
	"testing"
)

func TestReqBodyConfig(t *testing.T) {
	// Test case 1: valid input values
	config := models.BodyReqConfig{
		Params:   "params_value",
		Query:    "query_value",
		BrowseID: "browse_id_value",
		VideoID:  "video_id_value",
	}
	expectedJSON := `{"context":{"client":{"clientName":"WEB_REMIX","clientVersion":"1.20230306.01.00","hl":"en"}},"params":"params_value","query":"query_value","browseId":"browse_id_value","videoId":"video_id_value"}`
	jsonConfig, err := utils.ReqBodyConfig(config)
	if err != nil {
		t.Errorf("Error producing JSON string: %v", err)
	}
	if string(jsonConfig) != expectedJSON {
		t.Errorf("Unexpected result. Expected: %s, Got: %s", expectedJSON, jsonConfig)
	}

	// Test case 2: empty input values
	config = models.BodyReqConfig{}
	expectedJSON = `{"context":{"client":{"clientName":"WEB_REMIX","clientVersion":"1.20230306.01.00","hl":"en"}}}`
	jsonConfig, err = utils.ReqBodyConfig(config)
	if err != nil {
		t.Errorf("Error producing JSON string: %v", err)
	}
	if string(jsonConfig) != expectedJSON {
		t.Errorf("Unexpected result. Expected: %s, Got: %s", expectedJSON, jsonConfig)
	}

}
