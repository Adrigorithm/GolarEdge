package api

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
