package utils

import (
	"fmt"
	"io"
	"katze/src/services/models"
	"net/http"
)

func Request(req models.Request) ([]byte, error) {

	var body_config = models.BodyReqConfig{
		Params:   req.Params,
		Query:    req.Query,
		BrowseID: req.BrowseID,
		VideoID:  req.VideoID,
	}

	bodyReq, err := ReqBodyConfig(body_config)
	if err != nil {
		err := fmt.Errorf("error creating body config: %v", err)
		return []byte{}, err
	}

	url := "https://music.youtube.com" + req.UrlPath + "?" + req.UrlQueries
	request, err := CreateReqWithHeaders(bodyReq, url, req.GoogVisitorID)
	if err != nil {
		err := fmt.Errorf("error creating request: %v", err)
		return []byte{}, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		err := fmt.Errorf("error sending request: %v", err)
		return []byte{}, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		err := fmt.Errorf("error reading response body: %v", err)
		return []byte{}, err
	}

	return body, nil
}
