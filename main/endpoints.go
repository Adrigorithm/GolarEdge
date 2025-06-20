package main

import (
	"errors"
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"
)

func getUrlNoAuth (path string, values url.Values) string {
	uriBuilder := url.URL{
		Scheme: "https",
		Host:   "monitoringapi.solaredge.com",
		Path:   path,
	}

	if values == nil {
		values = url.Values{}
	}

	uriBuilder.RawQuery = values.Encode()

	return uriBuilder.String()
}

func getUrl(apiKey string, path string, values url.Values) (string, error) {
	if apiKey == "" {
		return "", errors.New("please specify an api key")
	}

	uriBuilder := url.URL{
		Scheme: "https",
		Host:   "monitoringapi.solaredge.com",
		Path:   path,
	}

	if values == nil {
		values = url.Values{}
	}

	values.Add("api_key", apiKey)
	uriBuilder.RawQuery = values.Encode()

	return uriBuilder.String(), nil
}

// Site Data API

func GetSiteListRequest(params SiteListParams, apiKey string) (string, error) {
	values := url.Values{}
	path := "sites/list"

	if params.size != nil {
		size := *params.size

		if size > -1 && size < 101 {
			values.Add("size", strconv.Itoa(size))
		}
	}

	if params.startIndex != nil {
		startIndex := *params.startIndex

		if startIndex > -1 {
			values.Add("startIndex", strconv.Itoa(startIndex))
		}
	}

	if params.searchText != "" {
		values.Add("searchText", params.searchText)
	}

	if params.sortProperty != "" {
		sortProperty := strings.ToLower(params.sortProperty)

		switch sortProperty {
		case "name":
			values.Add("sortProperty", "Name")
		case "country":
			values.Add("sortProperty", "Country")
		case "state":
			values.Add("sortProperty", "State")
		case "city":
			values.Add("sortProperty", "City")
		case "address":
			values.Add("sortProperty", "Address")
		case "zip":
			values.Add("sortProperty", "Zip")
		case "status":
			values.Add("sortProperty", "Status")
		case "peakpower":
			values.Add("sortProperty", "PeakPower")
		case "installationdate":
			values.Add("sortProperty", "InstallationDate")
		case "amount":
			values.Add("sortProperty", "Amount")
		case "maxseverity":
			values.Add("sortProperty", "MaxSeverity")
		case "creationtime":
			values.Add("sortProperty", "CreationTime")
		}
	}

	if params.sortOrder != "" {
		sortOrder := strings.ToLower(params.sortOrder)

		switch sortOrder {
		case "asc":
			values.Add("sortOrder", "ASC")
		case "desc":
			values.Add("sortOrder", "DESC")
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
			values.Add("status", statusString)
		}
	}

	return getUrl(apiKey, path, values)
}

func GetSiteRequest(params SiteParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("site/%d/details", params.siteId)

	return getUrl(apiKey, path, nil)
}

func GetSiteDataStartAndEndDatesRequest(params SiteDataStartAndEndDatesParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("site/%d/dataPeriod", params.siteId)

	return getUrl(apiKey, path, nil)
}

func GetSiteDataStartAndEndDatesBulkRequest(params SiteDataStartAndEndDatesBulkParams, apiKey string) (string, error) {
	if len(params.siteIds) == 0 {
		return "", errors.New("you must at least specify one site id")
	}

	siteIdsFiltered := []int{}
	siteIdsString := ""

	for i := range params.siteIds {
		siteId := params.siteIds[i]

		if siteId < 0 || slices.Contains(siteIdsFiltered, siteId) {
			continue
		}

		siteIdsFiltered = append(siteIdsFiltered, siteId)
		siteIdsString = fmt.Sprintf("%s%d,", siteIdsString, siteId)
	}

	if siteIdsString == "" {
		return "", errors.New("no valid site ids found. site ids must be positive integers")
	}

	siteIdsString = siteIdsString[:len(siteIdsString)-1]

	path := fmt.Sprintf("sites/%s/dataPeriod", siteIdsString)

	return getUrl(apiKey, path, nil)

}

func GetSiteEnergyWithParsedSitesRequest(idsString string, startDate time.Time, endDate time.Time, timeUnit string, apiKey string) (string, error) {
	path := fmt.Sprintf("site/%s/energy", idsString)
	values := url.Values{}

	if startDate.IsZero() || endDate.IsZero() {
		return "", errors.New("both start and end date are required")
	}

	if endDate.Before(startDate) {
		return "", errors.New("end date must be after the start date")
	}

	timeUnitUpper := strings.ToUpper(timeUnit)

	switch timeUnitUpper {
	case "QUARTER_OF_AN_HOUR":
	case "HOUR":
		if startDate.AddDate(0, 1, 0).Compare(endDate) < 1 {
			return "", errors.New("specified time unit limits difference in start and end date to one month")
		}

		values.Add("timeUnit", timeUnitUpper)
	case "WEEK":
	case "MONTH":
	case "YEAR":
		values.Add("timeUnit", timeUnitUpper)
	default:
		if startDate.AddDate(1, 0, 0).Compare(endDate) > 1 {
			return "", errors.New("specified time unit (day) limits difference in start and end date to one year")
		}

		values.Add("timeUnit", "DAY")
	}

	values.Add("startDate", startDate.String())
	values.Add("endDate", endDate.String())

	return getUrl(apiKey, path, values)
}

func GetSiteEnergyRequest(params SiteEnergyParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	return GetSiteEnergyWithParsedSitesRequest(strconv.Itoa(params.siteId), params.startDate, params.endDate, params.timeUnit, apiKey)
}

func GetSiteEnergyBulkRequest(params SiteEnergyBulkParams, apiKey string) (string, error) {
	if len(params.siteIds) == 0 {
		return "", errors.New("you must at least specify one site id")
	}

	siteIdsFiltered := []int{}
	siteIdsString := ""

	for i := range params.siteIds {
		siteId := params.siteIds[i]

		if siteId < 0 || slices.Contains(siteIdsFiltered, siteId) {
			continue
		}

		siteIdsFiltered = append(siteIdsFiltered, siteId)
		siteIdsString = fmt.Sprintf("%s%d,", siteIdsString, siteId)
	}

	if siteIdsString == "" {
		return "", errors.New("no valid site ids found. site ids must be positive integers")
	}

	return GetSiteEnergyWithParsedSitesRequest(siteIdsString, params.startDate, params.endDate, params.timeUnit, apiKey)
}

func GetSiteEnergyTimePeriodWithParsedSitesRequest(idsString string, startDate time.Time, endDate time.Time, apiKey string) (string, error) {
	path := fmt.Sprintf("site/%s/energy?timeFrameEnergy", idsString)
	values := url.Values{}

	if startDate.IsZero() || endDate.IsZero() {
		return "", errors.New("both start and end date are required")
	}

	if endDate.Before(startDate) {
		return "", errors.New("end date must be after the start date")
	}

	if startDate.AddDate(1, 0, 0).Compare(endDate) > 1 {
		return "", errors.New("this endpoint limits difference in start and end date to one year")
	}

	values.Add("startDate", startDate.String())
	values.Add("endDate", endDate.String())

	return getUrl(apiKey, path, values)
}

func GetSiteEnergyTimePeriodRequest(params SiteEnergyTimePeriodParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	return GetSiteEnergyTimePeriodWithParsedSitesRequest(strconv.Itoa(params.siteId), params.startDate, params.endDate, apiKey)
}

func GetSiteEnergyTimePeriodBulkRequest(params SiteEnergyTimePeriodBulkParams, apiKey string) (string, error) {
	if len(params.siteIds) == 0 {
		return "", errors.New("you must at least specify one site id")
	}

	siteIdsFiltered := []int{}
	siteIdsString := ""

	for i := range params.siteIds {
		siteId := params.siteIds[i]

		if siteId < 0 || slices.Contains(siteIdsFiltered, siteId) {
			continue
		}

		siteIdsFiltered = append(siteIdsFiltered, siteId)
		siteIdsString = fmt.Sprintf("%s%d,", siteIdsString, siteId)
	}

	if siteIdsString == "" {
		return "", errors.New("no valid site ids found. site ids must be positive integers")
	}

	return GetSiteEnergyTimePeriodWithParsedSitesRequest(siteIdsString, params.startDate, params.endDate, apiKey)
}

func GetSitePowerWithParsedSitesRequest(idsString string, startTime time.Time, endTime time.Time, apiKey string) (string, error) {
	path := fmt.Sprintf("site/%s/power", idsString)
	values := url.Values{}

	if startTime.IsZero() || endTime.IsZero() {
		return "", errors.New("both start and end time are required")
	}

	if endTime.Before(startTime) {
		return "", errors.New("end time must be after the start time")
	}

	if startTime.AddDate(0, 1, 0).Compare(endTime) > 1 {
		return "", errors.New("this endpoint limits difference in start and end time to one month")
	}

	values.Add("startTime", startTime.String())
	values.Add("endTime", endTime.String())

	return getUrl(apiKey, path, values)
}

func GetSitePowerRequest(params SitePowerParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	return GetSitePowerWithParsedSitesRequest(strconv.Itoa(params.siteId), params.startTime, params.endTime, apiKey)
}

func GetSitePowerBulkRequest(params SitePowerBulkParams, apiKey string) (string, error) {
	if len(params.siteIds) == 0 {
		return "", errors.New("you must at least specify one site id")
	}

	siteIdsFiltered := []int{}
	siteIdsString := ""

	for i := range params.siteIds {
		siteId := params.siteIds[i]

		if siteId < 0 || slices.Contains(siteIdsFiltered, siteId) {
			continue
		}

		siteIdsFiltered = append(siteIdsFiltered, siteId)
		siteIdsString = fmt.Sprintf("%s%d,", siteIdsString, siteId)
	}

	if siteIdsString == "" {
		return "", errors.New("no valid site ids found. site ids must be positive integers")
	}

	return GetSitePowerWithParsedSitesRequest(siteIdsString, params.startTime, params.endTime, apiKey)
}

func GetSiteOverviewRequest(params SiteOverviewParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("sites/%d/ovewview", params.siteId)

	return getUrl(apiKey, path, nil)
}

func GetSiteOverviewBulkRequest(params SiteOverviewBulkParams, apiKey string) (string, error) {
	if len(params.siteIds) == 0 {
		return "", errors.New("you must at least specify one site id")
	}

	siteIdsFiltered := []int{}
	siteIdsString := ""

	for i := range params.siteIds {
		siteId := params.siteIds[i]

		if siteId < 0 || slices.Contains(siteIdsFiltered, siteId) {
			continue
		}

		siteIdsFiltered = append(siteIdsFiltered, siteId)
		siteIdsString = fmt.Sprintf("%s%d,", siteIdsString, siteId)
	}

	if siteIdsString == "" {
		return "", errors.New("no valid site ids found. site ids must be positive integers")
	}

	siteIdsString = siteIdsString[:len(siteIdsString)-1]

	path := fmt.Sprintf("sites/%s/overview", siteIdsString)

	return getUrl(apiKey, path, nil)
}

func GetSitePowerDetailedRequest(params SitePowerDetailedParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("site/%d/powerDetails", params.siteId)
	values := url.Values{}

	if params.startTime.IsZero() || params.endTime.IsZero() {
		return "", errors.New("both start and end time are required")
	}

	if params.endTime.Before(params.startTime) {
		return "", errors.New("end time must be after the start time")
	}

	if params.startTime.AddDate(0, 1, 0).Compare(params.endTime) > 1 {
		return "", errors.New("this endpoint limits difference in start and end time to one month")
	}

	if len(params.meters) > 0 {
		meters := map[string]bool{
			"Production":   false,
			"Consumption":  false,
			"SelfConsumption": false,
			"FeedIn":      false,
			"Purchased": false,
		}
		metersString := ""

		for i := range params.meters {
			meter := strings.ToLower(params.meters[i])

			switch meter {
				case "production":
					meters["Production"] = true
				case "consumption":
					meters["Consumption"] = true
				case "selfconsumption":
					meters["SelfConsumption"] = true
				case "feedin":
					meters["FeedIn"] = true
				case "purchased":
					meters["Purchased"] = true
			}
		}

		for key, value := range meters {
			if value {
				metersString = fmt.Sprint(metersString, key, ',')
			}
		}

		if len(metersString) > 0 {
			metersString = metersString[:len(metersString)-1]
			values.Add("meters", metersString)
		}
	}

	values.Add("startTime", params.startTime.String())
	values.Add("endTime", params.endTime.String())

	return getUrl(apiKey, path, values)
}

func GetSiteEnergyDetailedRequest(params SiteEnergyDetailedParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("site/%d/energyDetails", params.siteId)
	values := url.Values{}

	if params.startTime.IsZero() || params.endTime.IsZero() {
		return "", errors.New("both start and end time are required")
	}

	if params.endTime.Before(params.startTime) {
		return "", errors.New("end time must be after the start time")
	}

	if params.startTime.AddDate(0, 1, 0).Compare(params.endTime) > 1 {
		return "", errors.New("this endpoint limits difference in start and end time to one month")
	}

	timeUnitUpper := strings.ToUpper(params.timeUnit)

	switch timeUnitUpper {
		case "QUARTER_OF_AN_HOUR":
		case "HOUR":
			if params.startTime.AddDate(0, 1, 0).Compare(params.endTime) < 1 {
				return "", errors.New("specified time unit limits difference in start and end date to one month")
			}

			values.Add("timeUnit", timeUnitUpper)
		case "WEEK":
		case "MONTH":
		case "YEAR":
			values.Add("timeUnit", timeUnitUpper)
		default:
			if params.startTime.AddDate(1, 0, 0).Compare(params.endTime) > 1 {
				return "", errors.New("specified time unit (day) limits difference in start and end date to one year")
			}

			values.Add("timeUnit", "DAY")
	}

	if len(params.meters) > 0 {
		meters := map[string]bool{
			"Production":   false,
			"Consumption":  false,
			"SelfConsumption": false,
			"FeedIn":      false,
			"Purchased": false,
		}
		metersString := ""

		for i := range params.meters {
			meter := strings.ToLower(params.meters[i])

			switch meter {
				case "production":
					meters["Production"] = true
				case "consumption":
					meters["Consumption"] = true
				case "selfconsumption":
					meters["SelfConsumption"] = true
				case "feedin":
					meters["FeedIn"] = true
				case "purchased":
					meters["Purchased"] = true
			}
		}

		for key, value := range meters {
			if value {
				metersString = fmt.Sprint(metersString, key, ',')
			}
		}

		if len(metersString) > 0 {
			metersString = metersString[:len(metersString)-1]
			values.Add("meters", metersString)
		}
	}

	values.Add("startTime", params.startTime.String())
	values.Add("endTime", params.endTime.String())

	return getUrl(apiKey, path, values)
}

func GetSitePowerFlowRequest(params SitePowerFlowParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("site/%d/currentPowerFlow", params.siteId)

	return getUrl(apiKey, path, nil)
}

func GetStorageInformationRequest(params StorageInformationParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("site/%d/storageData", params.siteId)
	values := url.Values{}

	if params.startTime.IsZero() || params.endTime.IsZero() {
		return "", errors.New("both start and end time are required")
	}

	if params.endTime.Before(params.startTime) {
		return "", errors.New("end time must be after the start time")
	}

	if params.startTime.AddDate(0, 0, 7).Compare(params.endTime) > 1 {
		return "", errors.New("this endpoint limits difference in start and end time to one week")
	}

	if (len(params.serials) > 0) {
		serialsString := ""
		serials := []string{}

		for i := range params.serials {
			serial := params.serials[i]
			isDuplicate := false

			for j := range serials {
				if (serials[j] == serial) {
					isDuplicate = true
				}
			}

			if (!isDuplicate) {
				serials = append(serials, serial)
				serialsString = fmt.Sprint(serialsString, serial , ',')
			}
		}

		values.Add("serials", serialsString[:len(serialsString)-1])
	}

	values.Add("startTime", params.startTime.String())
	values.Add("endTime", params.endTime.String())

	return getUrl(apiKey, path, values)
}

func GetSiteImageRequest(params SiteImageParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("site/%d/siteImage", params.siteId)
	values := url.Values{}

	if params.name != "" {
		path = fmt.Sprint(path, '/', params.name)
	}

	if params.maxHeight != nil {
		if *params.maxHeight <= 0 {
			return "", errors.New("image max height (if specified) must be a positive integer > 0")
		}

		values.Add("maxHeight", strconv.Itoa(*params.maxHeight))
	}

	if params.maxWidth != nil {
		if *params.maxWidth <= 0 {
			return "", errors.New("image max width (if specified) must be a positive integer > 0")
		}

		values.Add("maxWidth", strconv.Itoa(*params.maxWidth))
	}

	if params.hash != nil {
		values.Add("hash", strconv.Itoa(*params.hash))
	}

	return getUrl(apiKey, path, values)
}

func GetSiteEnvironmentalBenefitsRequest(params SiteEnvironmentalBenefitsParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("site/%d/envBenefits", params.siteId)
	values := url.Values{}

	systemUnitsUpper := strings.ToUpper(params.systemUnits)

	switch systemUnitsUpper {
		case "METRICS":
			values.Add("systemUnits", "Metrics")
		case "IMPERIAL":
			values.Add("systemUnits", "Imperial")
	}

	return getUrl(apiKey, path, values)
}

func GetInstallerImageRequest(params SiteImageParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("site/%d/installerImage", params.siteId)
	values := url.Values{}

	if params.name != "" {
		path = fmt.Sprint(path, '/', params.name)
	}

	return getUrl(apiKey, path, values)
}

// Site Equipment API

func GetComponentsListRequest(params ComponentsListParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("equipment/%d/list", params.siteId)

	return getUrl(apiKey, path, nil)
}

func GetInventoryRequest(params InventoryParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("site/%d/inventory", params.siteId)

	return getUrl(apiKey, path, nil)
}

func GetInverterTechnicalDataRequest(params InverterTechnicalDataParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	if params.serialNumber == "" {
		return "", errors.New("serialNumber must be valid (not null or empty string)")
	}

	path := fmt.Sprintf("equipment/%d/%s/data", params.siteId, params.serialNumber)
	values := url.Values{}

	if params.startTime.IsZero() || params.endTime.IsZero() {
		return "", errors.New("both start and end time are required")
	}

	if params.endTime.Before(params.startTime) {
		return "", errors.New("end time must be after the start time")
	}

	if params.startTime.AddDate(0, 0, 7).Compare(params.endTime) > 1 {
		return "", errors.New("this endpoint limits difference in start and end time to one week")
	}

	values.Add("startTime", params.startTime.String())
	values.Add("endTime", params.endTime.String())

	return getUrl(apiKey, path, values)
}

func GetEquipmentChangeLogRequest(params EquipmentChangeLogParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	if params.serialNumber == "" {
		return "", errors.New("serialNumber must be valid (not null or empty string)")
	}

	path := fmt.Sprintf("equipment/%d/%s/changeLog", params.siteId, params.serialNumber)

	return getUrl(apiKey, path, nil)
}

// Account List API

func GetAccountListRequest(params AccountListParams, apiKey string) (string, error) {
	values := url.Values{}
	path := "accounts/list"

	if params.size != nil {
		size := *params.size

		if size > -1 && size < 101 {
			values.Add("size", strconv.Itoa(size))
		}
	}

	if params.startIndex != nil {
		startIndex := *params.startIndex

		if startIndex > -1 {
			values.Add("startIndex", strconv.Itoa(startIndex))
		}
	}

	if params.searchText != "" {
		values.Add("searchText", params.searchText)
	}

	if params.sortProperty != "" {
		sortProperty := strings.ToLower(params.sortProperty)

		switch sortProperty {
			case "name":
				values.Add("sortProperty", "Name")
			case "country":
				values.Add("sortProperty", "country")
			case "city":
				values.Add("sortProperty", "city")
			case "address":
				values.Add("sortProperty", "address")
			case "zip":
				values.Add("sortProperty", "zip")
			case "fax":
				values.Add("sortProperty", "fax")
			case "phone":
				values.Add("sortProperty", "phone")
			case "notes":
				values.Add("sortProperty", "notes")
		}
	}

	if params.sortOrder != "" {
		sortOrder := strings.ToLower(params.sortOrder)

		switch sortOrder {
			case "asc":
				values.Add("sortOrder", "ASC")
			case "desc":
				values.Add("sortOrder", "DESC")
		}
	}

	return getUrl(apiKey, path, values)
}

// Meters API

func GetMetersDataRequest(params MetersDataParams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("site/%d/meters", params.siteId)
	values := url.Values{}

	if params.startTime.IsZero() || params.endTime.IsZero() {
		return "", errors.New("both start and end time are required")
	}

	if params.endTime.Before(params.startTime) {
		return "", errors.New("end time must be after the start time")
	}

	timeUnitUpper := strings.ToUpper(params.timeUnit)

	switch timeUnitUpper {
		case "QUARTER_OF_AN_HOUR":
		case "HOUR":
		case "WEEK":
		case "MONTH":
		case "YEAR":
			values.Add("timeUnit", timeUnitUpper)
		default:
			values.Add("timeUnit", "DAY")
	}

	if len(params.meters) > 0 {
		meters := map[string]bool{
			"Production":   false,
			"Consumption":  false,
			"FeedIn":      false,
			"Purchased": false,
		}
		metersString := ""

		for i := range params.meters {
			meter := strings.ToLower(params.meters[i])

			switch meter {
				case "production":
					meters["Production"] = true
				case "consumption":
					meters["Consumption"] = true
				case "feedin":
					meters["FeedIn"] = true
				case "purchased":
					meters["Purchased"] = true
			}
		}

		for key, value := range meters {
			if value {
				metersString = fmt.Sprint(metersString, key, ',')
			}
		}

		if len(metersString) > 0 {
			metersString = metersString[:len(metersString)-1]
			values.Add("meters", metersString)
		}
	}

	values.Add("startTime", params.startTime.String())
	values.Add("endTime", params.endTime.String())

	return getUrl(apiKey, path, values)
}

// Sensors API

func GetSensorsListRequest(params SensorsListparams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("equipment/%d/sensors", params.siteId)

	return getUrl(apiKey, path, nil)
}

func GetSensorDataRequest(params SensorDataparams, apiKey string) (string, error) {
	if params.siteId < 0 {
		return "", errors.New("site id must be an int >= 0")
	}

	path := fmt.Sprintf("site/%d/sensors", params.siteId)
	values := url.Values{}

	if params.startDate.IsZero() || params.endDate.IsZero() {
		return "", errors.New("both start and end time are required")
	}

	if params.endDate.Before(params.startDate) {
		return "", errors.New("end time must be after the start time")
	}

	if params.startDate.AddDate(0, 0, 7).Compare(params.endDate) > 1 {
		return "", errors.New("this endpoint limits difference in start and end time to one week")
	}

	values.Add("startDate", params.startDate.String())
	values.Add("endDate", params.endDate.String())

	return getUrl(apiKey, path, values)
}

// API Versions

func GetCurrentVersionRequest() (string) {
	return getUrlNoAuth("version/current", nil)
}

func GetSupportedVersionRequest() (string) {
	return getUrlNoAuth( "version/supported", nil)
}
