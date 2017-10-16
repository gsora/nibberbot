package breath

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL          = "http://gorthaur.biancalana.me/vaporfont"
	maxBreathingTime = 3
)

type BreatRequest struct {
	Message string `json:"message"`
}

func Breath(str string) (string, error) {
	payload := url.Values{}
	payload.Set("payload", str)
	timeout := time.Duration(maxBreathingTime * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.PostForm(baseURL, payload)
	if err != nil {
		return "", err
	}

	var br BreatRequest
	jb := json.NewDecoder(resp.Body)
	err = jb.Decode(&br)
	if err != nil {
		return "", err
	}
	return br.Message, nil
}
