package repository

import "e-money-svc/domain/model"

// client
type BillerServiceClient interface {
	GetListBiller() (*model.ListBillerResponse, error)
	GetDetailBiller(id string) (*model.DetailBillerResponse, error)
}
