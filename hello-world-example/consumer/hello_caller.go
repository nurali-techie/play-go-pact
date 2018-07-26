package consumer

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HelloCaller(host, name string) (string, error) {
	url := fmt.Sprintf("%s?name=%s", host, name)
	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	msg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(msg), nil
}
