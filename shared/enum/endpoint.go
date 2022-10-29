package enum

type ClientEndpoint string

const (
	BillerServiceGetList ClientEndpoint = "list"
	BillerServiceDetail  ClientEndpoint = "detail"
)
