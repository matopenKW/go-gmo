package aozorabank

import (
	"context"
	"fmt"
	"github.com/abyssparanoia/go-gmo/internal/pkg/converter"
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

// https://stg-api.gmo-aozora.com/ganb/api/auth/v1/authorization?

const (
	authPathV1 = "auth/v1"
)

type (
	AuthorizationRequest struct {
		ClientID     string `json:"client_id" validate:"required,min=1,max=128"`
		RedirectURI  string `json:"redirect_uri" validate:"required,uri,min=1,max=256"`
		ResponseType string `json:"response_type" validate:"required"`
		Scope        string `json:"scope" validate:"required,min=1,max=256"`
		State        string `json:"state" validate:"required,min=1,max=128"`
	}

	AuthorizationResponse struct{}
)

func (r *AuthorizationRequest) Validate() error {
	return validate.Struct(r)
}

func (cli *Client) Authorization(
	ctx context.Context,
	req *AuthorizationRequest,
) (*AuthorizationResponse, error) {

	//response_type=code&scope=%7B%E3%82%B9%E3%82%B3%E3%83%BC%E3%83%95%E3%82%9A%7D&client_id=%7B%E3%82%AF%E3%83%A9%E3%82%A4%E3%82%A2%E3%83%B3%E3%83%88ID%7D&state=%7B%E3%82%B9%E3%83%86%E3%83%BC%E3%83%88%E5%80%A4%7D&redirect_uri=%7B%E3%83%AA%E3%82%BF%E3%82%99%E3%82%A4%E3%83%AC%E3%82%AF%E3%83%88%E7%94%A8URL%7D
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	res := &AuthorizationResponse{}
	if _, err := cli.doGet(fmt.Sprintf("%s/authorization", authPathV1), reqMap, res); err != nil {
		return nil, err
	}
	return res, nil
}
