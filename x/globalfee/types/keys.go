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

var (
	// KeyWhiteList is the key to use for the keeper store .
	KeyWhiteList = []byte{0x00}

	// The actual fees themselfs when for use within the keeper store .
	OverrideFees = []byte{0x01}
)
