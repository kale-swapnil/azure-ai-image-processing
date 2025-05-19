package functions

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func ClassifyImage(imageUrl string) ([]string, error) {
	body, _ := json.Marshal(map[string]string{"url": imageUrl})
	req, _ := http.NewRequest("POST", os.Getenv("CV_API_ENDPOINT")+"/vision/v3.2/tag", bytes.NewReader(body))
	req.Header.Add("Ocp-Apim_Subscription-Key", os.Getenv("CV_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Tags []struct {
			Name string `json:"name"`
		} `json:"tags"`
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &result)

	tags := make([]string, 0)
	for _, tag := range result.Tags {
		tags = append(tags, tag.Name)
	}

	return tags, nil

}
