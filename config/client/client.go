package client

type ClientList struct {
	Biller BillerService
}

type BillerService struct {
	Host string `yaml:"host"`
}
