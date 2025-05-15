package api

import (
	"errors"
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"strings"
)

func getBaseUrl(apiKey string) (url.URL, url.Values) {
	baseUrl := url.URL {
		Scheme: "https",
		Host: "monitoringapi.solaredge.com"
	}

	values := url.Values {}

	if apiKey != "" {
		values.Add("api_key", apiKey)
	}

	return baseUrl, values
}

func getBasicParameterisedUrl(path string, apiKey string) (string, error) {
	if apiKey == "" {
		return "", errors.New("Please specify an API key.")
	}

	endpointUrl, endpointValues := getBaseUrl()
	endpointUrl.Path = path
	endpointUrl.RawQuery = endpointValues.Encode()

	return endpointUrl.String(), nil
}

func GetSiteList(params SiteListParams, apiKey string) (string, error) {
	if apiKey == "" {
		return "", errors.New("Please specify an API key.")
	}

	endpointUrl, endpointValues := getBaseUrl()
	endpointUrl.Path: "sites/list"

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

	return endpointUrl.String(), nil
}

func GetSite(params SiteParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("Site ID must be an int >= 0")
	}

	path := fmt.Sprintf("site/%d/details", params.siteId)

	return getBasicParameterisedUrl(path, apiKey)
}

func GetSiteDataStartAndEndDates(params SiteDataStartAndEndDates, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("Site ID must be an int >= 0")
	}

	path := fmt.Sprintf("site/%d/dataPeriod", params.siteId)

	return getBasicParameterisedUrl(path, apiKey)
}

func GetSiteDataStartAndEndDatesBulk(params SiteDataStartAndEndDatesBulk, apiKey string) (string, error) {
	if len(params.siteIds) == 0 {
		return "", errors.New("You must at least specify one Site ID")
	}

	siteIdsFiltered := []int {}
	siteIdsString := ""

	for i := range params.siteIds {
		siteId := params.siteIds[i]

		if siteId < 0 {
			continue
		}

		isDuplicate := false

		for j := range siteIdsFiltered {
			if siteIdsFiltered[j] == siteId {
				isDuplicate = true
				break
			}
		}

		if !isDuplicate {
			siteIdsFiltered = append(siteIdsFiltered, siteId)
			siteIdsString = fmt.Sprintf("%s%s,", siteIdsString, strconv.Itoa(siteId))
		}
	}

	if siteIdsString == "" {
		return "", errors.New("No valid Site IDs found. Site IDs must be positive integers")
	}

	siteIdsString = siteIdsString[:len(siteIdsString)-1]

	path := fmt.Sprintf("sites/%s/dataPeriod", siteIdsString)

	return getBasicParameterisedUrl(path, apiKey)


}
