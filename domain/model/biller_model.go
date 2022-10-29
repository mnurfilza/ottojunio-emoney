package model

type ListBillerResponse struct {
	Code   int       `json:"code"`
	Status string    `json:"status"`
	Data   []*Biller `json:"data"`
}

type DetailBillerResponse struct {
	Code   int     `json:"code"`
	Status string  `json:"status"`
	Data   *Biller `json:"data"`
}

type Biller struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
	Product  string `json:"product"`
	Desc     string `json:"description"`
	Price    int    `json:"price"`
	Fee      int    `json:"fee"`
}
