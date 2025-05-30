package main

import "time"

// Site Data API

// SiteListParams represents the parameters for the GetSiteList endpoint function.
// if size is nil, the default value of 100 will be used.
// if startIndex is nil, the default value of 0 will be used.
// if searchText is "", it will be ignored.
// if sortProperty is not in ["Name", "Country", "State", "City", "Address", "Zip", "Status", "PeakPower", "InstallationDate", "Amount", "MaxSeverity", "CreationTime"], it is ignored.
// if sortOrder is not in ["ASC", "DESC"], it will be ignored.
// if status has values not in ["Active","Pending","Disabled","All"], they will be ignored, while keeping the valid ones.
type SiteListParams struct {
	size         *int
	startIndex   *int
	searchText   string
	sortProperty string
	sortOrder    string
	status       []string
}

type SiteParams struct {
	siteId int
}

type SiteDataStartAndEndDatesParams struct {
	siteId int
}

type SiteDataStartAndEndDatesBulkParams struct {
	siteIds []int
}

// SiteEnergyParams represents the parameters for the GetSiteEnergy endpoint function.
// if timeUnit is not in ["QUARTER_OF_AN_HOUR", "HOUR", "DAY", "WEEK", "MONTH", "YEAR"], it will default to "DAY".
type SiteEnergyParams struct {
	siteId int

	// Precision: 2006-01-02
	startDate time.Time

	// Precision: 2006-01-02
	endDate time.Time

	timeUnit string
}

// SiteEnergyBulkParams represents the parameters for the GetSiteBulkEnergy endpoint function.
// if timeUnit is not in ["QUARTER_OF_AN_HOUR", "HOUR", "DAY", "WEEK", "MONTH", "YEAR"], it will default to "DAY".
type SiteEnergyBulkParams struct {
	siteIds []int

	// Precision: 2006-01-02
	startDate time.Time

	// Precision: 2006-01-02
	endDate time.Time

	timeUnit string
}

// SiteEnergyTimePeriodParams returns total energy on generated at two specific points in time (the startDate and endDate). This metric is pretty useless in my opinion (but it exists).
type SiteEnergyTimePeriodParams struct {
	siteId int

	// Precision: 2006-01-02
	startDate time.Time

	// Precision: 2006-01-02
	endDate time.Time
}

type SiteEnergyTimePeriodBulkParams struct {
	siteIds []int

	// Precision: 2006-01-02
	startDate time.Time

	// Precision: 2006-01-02
	endDate time.Time
}

// SitePowerParams return the site power measurements in 15 minutes resolution
type SitePowerParams struct {
	siteId int

	// Precision: 2006-01-02 11:00:00
	startTime time.Time

	// Precision: 2006-01-02 11:00:00
	endTime time.Time
}

type SitePowerBulkParams struct {
	siteIds []int

	// Precision: 2006-01-02 11:00:00
	startTime time.Time

	// Precision: 2006-01-02 11:00:00
	endTime time.Time
}

type SiteOverviewParams struct {
	siteId int
}

type SiteOverviewBulkParams struct {
	siteIds []int
}

// SitePowerDetailedParams returns detailed site power measurements from meters such as consumption, export (feed-in), import (purchase), etc.
// meters is optional, an array of strings representing meters. If not specified, all meter readings are returned.
type SitePowerDetailedParams struct {
	siteId int

	// Precision: 2006-01-02 11:00:00
	startTime time.Time

	// Precision: 2006-01-02 11:00:00
	endTime time.Time

	meters []string
}

type SiteEnergyDetailedParams struct {
	siteId int

	// Precision: 2006-01-02 11:00:00
	startTime time.Time

	// Precision: 2006-01-02 11:00:00
	endTime time.Time

	timeUnit string
	meters []string
}

type SitePowerFlowParams struct {
	siteId int
}

// StorageInformationParams returns storage information about the batteries
type StorageInformationParams struct {
	siteId int

	// Precision: 2006-01-02 11:00:00
	startTime time.Time

	// Precision: 2006-01-02 11:00:00
	endTime time.Time

	serials []string
}

// SiteImageParams returns the site image (uploaded by the user).
// The name parameter will be used as the name for the downloaded file (optional).
// Specifying maxWidth or maxHeight will rescale the image using the original upload size, but maintaining the ratio. Setting either will ignore the hash parameter.
// When the hash parameter is set the API will return a 304 error if the uploaded file has a different hash than the one provided.
type SiteImageParams struct {
	siteId int
	name string
	maxWidth *int
	maxHeight *int
	hash *int
}

// SiteEnvironmentalBenefitsParams returns the list of environmental benefits associated with the site energy production.
// SystemUnits is either Metrics or Imerial (If not specified logged in user system units are used)
type SiteEnvironmentalBenefitsParams struct {
	siteId int
	systemUnits string
}

type InstallerLogoParams struct {
	siteId int
	name string
}

// Site Equipment API

type ComponentsListParams struct {
	siteId int
}

type InventoryParams struct {
	siteId int
}

// InverterTechnicalDataParams returns specific inverter data for a given timeframe.
// All parameters are required. The difference between startTime and endTime should not exceed one week.
type InverterTechnicalDataParams struct {
	siteId int
	serialNumber string

	// Precision: 2006-01-02 11:00:00
	startTime time.Time

	// Precision: 2006-01-02 11:00:00
	endTime time.Time
}

type EquipmentChangeLogParams struct {
	siteId int
	serialNumber string
}

// Account List API

// AccountListParams the account and list of sub-accounts related to the given token.
// if size is nil, the default value of 100 will be used.
// if startIndex is nil, the default value of 0 will be used.
// if searchText is "", it will be ignored.
// if sortProperty is not in ["Name", "country", "city", "address", "zip", "fax", "phone", "notes"], it is ignored.
// if sortOrder is not in ["ASC", "DESC"], it will be ignored.
type AccountListParams struct {
	size         *int
	startIndex   *int
	searchText   string
	sortProperty string
	sortOrder    string
}

// Meters API

// MetersDataParams returns for each meter on site its lifetime energy reading, metadata and the device to which it’s connected to
// if timeUnit is not in ["QUARTER_OF_AN_HOUR", "HOUR", "DAY", "WEEK", "MONTH", "YEAR"], it will default to "DAY".
type MetersDataParams struct {
	siteId int

	timeUnit string

	// Precision: 2006-01-02 11:00:00
	startTime time.Time

	// Precision: 2006-01-02 11:00:00
	endTime time.Time

	meters []string
}

// Sensors API

type SensorsListparams struct {
	siteId int
}

// SensorDataparams returns the data of all the sensors in the site, by the gateway they are connected to.
// startDate and endDate should not exceed one week as this API is limited as such.
type SensorDataparams struct {
	siteId int

	// Precision: 2006-01-02 11:00:00
	startDate time.Time

	// Precision: 2006-01-02 11:00:00
	endDate time.Time
}
