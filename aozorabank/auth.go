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
