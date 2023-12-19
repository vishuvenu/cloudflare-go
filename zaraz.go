package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type ZarazConfig = map[string]interface{}

type UpdateZarazConfigParams = map[string]interface{}

type ZarazConfigRow struct {
	ID          int64     `json:"id,omitempty"`
	UserId      string    `json:"usedId,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

type GetZarazConfigHistoryResponse struct {
	Data  []ZarazConfigRow `json:"data"`
	Count int              `json:"count"`
}

type GetZarazConfigHistoryDiffResponse = map[string]interface{}

func getEndpointVersion(version string) string {
	if version == "v2" {
		return "settings/zaraz/v2"
	} else {
		return "settings/zaraz"
	}
}

func (api *API) GetZarazConfig(ctx context.Context, rc *ResourceContainer) (ZarazConfig, error) {
	if rc.Identifier == "" {
		return ZarazConfig{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/%s/config", rc.Identifier, getEndpointVersion("v1"))
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZarazConfig{}, err
	}

	var recordResp ZarazConfig
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return ZarazConfig{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

func (api *API) UpdateZarazConfig(ctx context.Context, rc *ResourceContainer, params UpdateZarazConfigParams) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/%s/config", rc.Identifier, getEndpointVersion("v1"))
	_, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) GetZarazWorkflow(ctx context.Context, rc *ResourceContainer) (string, error) {
	if rc.Identifier == "" {
		return "", ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/%s/workflow", rc.Identifier, getEndpointVersion("v1"))
	workflow, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return "", err
	}

	return string(workflow), nil
}

func (api *API) UpdateZarazWorkflow(ctx context.Context, rc *ResourceContainer, workflowToUpdate string) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/%s/workflow", rc.Identifier, getEndpointVersion("v1"))
	_, err := api.makeRequestContext(ctx, http.MethodPut, uri, workflowToUpdate)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) PublishZarazConfig(ctx context.Context, rc *ResourceContainer, description string) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/%s/publish", rc.Identifier, getEndpointVersion("v1"))
	_, err := api.makeRequestContext(ctx, http.MethodPost, uri, description)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) GetZarazConfigHistory(ctx context.Context, rc *ResourceContainer, limit int64, offset int64) (GetZarazConfigHistoryResponse, error) {
	if rc.Identifier == "" {
		return GetZarazConfigHistoryResponse{
			Data:  []ZarazConfigRow{},
			Count: 0,
		}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/%s/history?limit=%d&offset=%d", rc.Identifier, getEndpointVersion("v1"), limit, offset)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return GetZarazConfigHistoryResponse{
			Data:  []ZarazConfigRow{},
			Count: 0,
		}, err
	}

	var recordResp GetZarazConfigHistoryResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return GetZarazConfigHistoryResponse{
			Data:  []ZarazConfigRow{},
			Count: 0,
		}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

func (api *API) GetZarazConfigHistoryDiff(ctx context.Context, rc *ResourceContainer, limit int64, offset int64, configIds string) (GetZarazConfigHistoryDiffResponse, error) {
	if rc.Identifier == "" {
		return GetZarazConfigHistoryDiffResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/%s/history/configs?ids=%s", rc.Identifier, getEndpointVersion("v1"), configIds)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return GetZarazConfigHistoryDiffResponse{}, err
	}

	var recordResp GetZarazConfigHistoryDiffResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return GetZarazConfigHistoryDiffResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

func (api *API) GetDefaultZarazConfig(ctx context.Context, rc *ResourceContainer) (ZarazConfig, error) {
	if rc.Identifier == "" {
		return ZarazConfig{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/%s/default", rc.Identifier, getEndpointVersion("v1"))
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZarazConfig{}, err
	}

	var recordResp ZarazConfig
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return ZarazConfig{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

func (api *API) ExportZarazConfig(ctx context.Context, rc *ResourceContainer) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/%s/export", rc.Identifier, getEndpointVersion("v1"))
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return err
	}

	var recordResp ZarazConfig
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}
