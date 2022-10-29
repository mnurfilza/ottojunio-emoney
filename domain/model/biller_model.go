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
	ID       int    `json:"id,omitempty"`
	Category string `json:"category,omitempty"`
	Product  string `json:"product,omitempty"`
	Desc     string `json:"description,omitempty"`
	Price    int    `json:"price,omitempty"`
	Fee      int    `json:"fee,omitempty"`
}
