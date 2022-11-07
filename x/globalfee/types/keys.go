package types

const (
	// ModuleName is the name of the this module
	ModuleName = "globalfee"

	// StoreKey is the default store key for globalfee .
	StoreKey = ModuleName

	QuerierRoute = ModuleName

	// Query endpoints supported by the globalfee querier .
	QueryParameters = "parameters"
	QueryWhiteList  = "whiteList"
)
