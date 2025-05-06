package weebsocket

import (
	"io"
	"net/http"
)

const baseUri string = "https://monitoringapi.solaredge.com/"

func get(endpoint string) string {
	response, error := http.Get(baseUri + endpoint)

	if error != nil {
		return ""
	}

	body := response.Body

	defer body.Close()

	bytes, error := io.ReadAll(body)

	if error != nil {
		return ""
	}

	return string(bytes)
}
