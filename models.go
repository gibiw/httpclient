package httpclient

import "time"

type ParameterResponse struct {
	CreatedDate    time.Time `json:"createdDate"`
	ModifiedDate   time.Time `json:"modifiedDate"`
	CreatedById    string    `json:"createdById"`
	ModifiedById   string    `json:"modifiedById"`
	IsDeleted      bool      `json:"isDeleted"`
	ParameterKeyId string    `json:"parameterKeyId"`
	Id             string    `json:"id"`
	Value          string    `json:"value"`
	Name           string    `json:"name"`
}

type ParameterRequest struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}
