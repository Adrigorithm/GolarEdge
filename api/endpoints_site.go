package api

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// /sites/list
func GetSiteList(params SiteListParams) string {
	endpointUrl := url.URL{
		Scheme: "https",
		Host:   "monitoringapi.solaredge.com",
		Path:   "sites/list",
	}
	endpointValues := url.Values{}

	if params.size != nil {
		size := *params.size

		if size > -1 && size < 101 {
			endpointValues.Add("size", strconv.Itoa(size))
		}
	}

	if params.startIndex != nil {
		startIndex := *params.startIndex

		if startIndex > -1 {
			endpointValues.Add("startIndex", strconv.Itoa(startIndex))
		}
	}

	if params.searchText != "" {
		endpointValues.Add("searchText", params.searchText)
	}

	if params.sortProperty != "" {
		sortProperty := strings.ToLower(params.sortProperty)

		switch sortProperty {
		case "name":
			endpointValues.Add("sortProperty", "Name")
		case "country":
			endpointValues.Add("sortProperty", "Country")
		case "state":
			endpointValues.Add("sortProperty", "State")
		case "city":
			endpointValues.Add("sortProperty", "City")
		case "address":
			endpointValues.Add("sortProperty", "Address")
		case "zip":
			endpointValues.Add("sortProperty", "Zip")
		case "status":
			endpointValues.Add("sortProperty", "Status")
		case "peakpower":
			endpointValues.Add("sortProperty", "PeakPower")
		case "installationdate":
			endpointValues.Add("sortProperty", "InstallationDate")
		case "amount":
			endpointValues.Add("sortProperty", "Amount")
		case "maxseverity":
			endpointValues.Add("sortProperty", "MaxSeverity")
		case "creationtime":
			endpointValues.Add("sortProperty", "CreationTime")
		}
	}

	if params.sortOrder != "" {
		sortOrder := strings.ToLower(params.sortOrder)

		switch sortOrder {
		case "asc":
			endpointValues.Add("sortOrder", "ASC")
		case "desc":
			endpointValues.Add("sortOrder", "DESC")
		}
	}

	if len(params.status) > 0 {
		statuses := map[string]bool{
			"Active":   false,
			"Pending":  false,
			"Disabled": false,
			"All":      false,
		}
		statusString := ""

		for i := range params.status {
			status := strings.ToLower(params.status[i])

			switch status {
			case "active":
				statuses["Active"] = true
			case "pending":
				statuses["Pending"] = true
			case "disabled":
				statuses["Disabled"] = true
			case "all":
				statuses["All"] = true
			}
		}

		for key, value := range statuses {
			if value {
				statusString = fmt.Sprint(statusString, key, ',')
			}
		}

		if len(statusString) > 0 {
			statusString = statusString[:len(statusString)-1]
			endpointValues.Add("status", statusString)
		}
	}

	endpointUrl.RawQuery = endpointValues.Encode()

	return endpointUrl.String()
}
