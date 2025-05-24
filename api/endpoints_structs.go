package api

import "time"

// SiteListParams represents the parameters for the GetSiteList endpoint function.
// if size is nil, the default value of 100 will be used.
// if startIndex is nil, the default value of 0 will be used.
// if sortText is "", it will be ignored.
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
