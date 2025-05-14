package api

import (
	"fmt"
	"strings"
)

// /sites/list
func GetSiteList(params SiteListParams) string {
	endpointString := "/sittes/list?"

	if params.size != nil {
		size := *params.size

		if size > -1 && size < 101 {
			endpointString = fmt.Sprintf("%ssize=%d&", endpointString, size)
		}
	}

	if params.startIndex != nil {
		startIndex := *params.startIndex

		if startIndex > -1 {
			endpointString = fmt.Sprintf("%sstartIndex=%d&", endpointString, startIndex)
		}
	}

	if params.searchText != "" {
		endpointString = fmt.Sprintf("%ssearchText=%s&", endpointString, params.searchText)
	}

	if params.sortProperty != "" {
		sortProperty := strings.ToLower(params.sortProperty)

		switch sortProperty {
		case "name":
		case "country":
		case "state":
		case "city":
		case "address":
		case "zip":
		case "status":
		case "peakpower":
		case "installationdate":
		case "amount":
		case "maxseverity":
		case "creationtime":
			endpointString = fmt.Sprintf("%ssortProperty=%s&", endpointString, sortProperty)
		}
	}

	if params.sortOrder != "" {
		sortOrder := strings.ToLower(params.sortOrder)

		switch sortOrder {
		case "asc":
		case "desc":
			endpointString = fmt.Sprintf("%ssortProperty=%s&", endpointString, sortOrder)
		}
	}

	if len(params.status) > 0 {
		statusString := ""

		for i := range params.status {
			status := strings.ToLower(params.status[i])

			switch status {
			case "active":
			case "pending":
			case "disabled":
			case "all":
				statusString = fmt.Sprintf("%s%s,", statusString, status)
			}
		}

		if len(statusString) > 0 {
			statusString = statusString[:len(statusString)-1]
			endpointString = fmt.Sprintf("%status=%s&", endpointString, statusString)
		}
	}

	return endpointString
}
