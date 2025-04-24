package weebsocket

import "net/http"

const base_uri string = "https://monitoringapi.solaredge.com/"

func get(endpoint string) {
	response, err := http.Get(base_uri + endpoint)

	response.Body.
}