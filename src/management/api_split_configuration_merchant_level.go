/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// SplitConfigurationMerchantLevelApi service
type SplitConfigurationMerchantLevelApi common.Service

// All parameters accepted by SplitConfigurationMerchantLevelApi.CreateRule
type SplitConfigurationMerchantLevelApiCreateRuleInput struct {
	merchantId             string
	splitConfigurationId   string
	splitConfigurationRule *SplitConfigurationRule
}

func (r SplitConfigurationMerchantLevelApiCreateRuleInput) SplitConfigurationRule(splitConfigurationRule SplitConfigurationRule) SplitConfigurationMerchantLevelApiCreateRuleInput {
	r.splitConfigurationRule = &splitConfigurationRule
	return r
}

/*
Prepare a request for CreateRule
@param merchantId The unique identifier of the merchant account.@param splitConfigurationId The unique identifier of the split configuration.
@return SplitConfigurationMerchantLevelApiCreateRuleInput
*/
func (a *SplitConfigurationMerchantLevelApi) CreateRuleInput(merchantId string, splitConfigurationId string) SplitConfigurationMerchantLevelApiCreateRuleInput {
	return SplitConfigurationMerchantLevelApiCreateRuleInput{
		merchantId:           merchantId,
		splitConfigurationId: splitConfigurationId,
	}
}

/*
CreateRule Create a rule

Creates a rule in the split configuration specified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API - SplitConfiguration read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r SplitConfigurationMerchantLevelApiCreateRuleInput - Request parameters, see CreateRuleInput
@return SplitConfiguration, *http.Response, error
*/
func (a *SplitConfigurationMerchantLevelApi) CreateRule(ctx context.Context, r SplitConfigurationMerchantLevelApiCreateRuleInput) (SplitConfiguration, *http.Response, error) {
	res := &SplitConfiguration{}
	path := "/merchants/{merchantId}/splitConfigurations/{splitConfigurationId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"splitConfigurationId"+"}", url.PathEscape(common.ParameterValueToString(r.splitConfigurationId, "splitConfigurationId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.splitConfigurationRule,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by SplitConfigurationMerchantLevelApi.CreateSplitConfiguration
type SplitConfigurationMerchantLevelApiCreateSplitConfigurationInput struct {
	merchantId         string
	splitConfiguration *SplitConfiguration
}

func (r SplitConfigurationMerchantLevelApiCreateSplitConfigurationInput) SplitConfiguration(splitConfiguration SplitConfiguration) SplitConfigurationMerchantLevelApiCreateSplitConfigurationInput {
	r.splitConfiguration = &splitConfiguration
	return r
}

/*
Prepare a request for CreateSplitConfiguration
@param merchantId The unique identifier of the merchant account.
@return SplitConfigurationMerchantLevelApiCreateSplitConfigurationInput
*/
func (a *SplitConfigurationMerchantLevelApi) CreateSplitConfigurationInput(merchantId string) SplitConfigurationMerchantLevelApiCreateSplitConfigurationInput {
	return SplitConfigurationMerchantLevelApiCreateSplitConfigurationInput{
		merchantId: merchantId,
	}
}

/*
CreateSplitConfiguration Create a split configuration

Creates a split configuration for the merchant account specified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API - SplitConfiguration read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r SplitConfigurationMerchantLevelApiCreateSplitConfigurationInput - Request parameters, see CreateSplitConfigurationInput
@return SplitConfiguration, *http.Response, error
*/
func (a *SplitConfigurationMerchantLevelApi) CreateSplitConfiguration(ctx context.Context, r SplitConfigurationMerchantLevelApiCreateSplitConfigurationInput) (SplitConfiguration, *http.Response, error) {
	res := &SplitConfiguration{}
	path := "/merchants/{merchantId}/splitConfigurations"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.splitConfiguration,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by SplitConfigurationMerchantLevelApi.DeleteSplitConfiguration
type SplitConfigurationMerchantLevelApiDeleteSplitConfigurationInput struct {
	merchantId           string
	splitConfigurationId string
}

/*
Prepare a request for DeleteSplitConfiguration
@param merchantId The unique identifier of the merchant account.@param splitConfigurationId The unique identifier of the split configuration.
@return SplitConfigurationMerchantLevelApiDeleteSplitConfigurationInput
*/
func (a *SplitConfigurationMerchantLevelApi) DeleteSplitConfigurationInput(merchantId string, splitConfigurationId string) SplitConfigurationMerchantLevelApiDeleteSplitConfigurationInput {
	return SplitConfigurationMerchantLevelApiDeleteSplitConfigurationInput{
		merchantId:           merchantId,
		splitConfigurationId: splitConfigurationId,
	}
}

/*
DeleteSplitConfiguration Delete a split configuration

Deletes the split configuration specified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API - SplitConfiguration read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r SplitConfigurationMerchantLevelApiDeleteSplitConfigurationInput - Request parameters, see DeleteSplitConfigurationInput
@return SplitConfiguration, *http.Response, error
*/
func (a *SplitConfigurationMerchantLevelApi) DeleteSplitConfiguration(ctx context.Context, r SplitConfigurationMerchantLevelApiDeleteSplitConfigurationInput) (SplitConfiguration, *http.Response, error) {
	res := &SplitConfiguration{}
	path := "/merchants/{merchantId}/splitConfigurations/{splitConfigurationId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"splitConfigurationId"+"}", url.PathEscape(common.ParameterValueToString(r.splitConfigurationId, "splitConfigurationId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodDelete,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by SplitConfigurationMerchantLevelApi.DeleteSplitConfigurationRule
type SplitConfigurationMerchantLevelApiDeleteSplitConfigurationRuleInput struct {
	merchantId           string
	splitConfigurationId string
	ruleId               string
}

/*
Prepare a request for DeleteSplitConfigurationRule
@param merchantId The unique identifier of the merchant account.@param splitConfigurationId The unique identifier of the split configuration.@param ruleId
@return SplitConfigurationMerchantLevelApiDeleteSplitConfigurationRuleInput
*/
func (a *SplitConfigurationMerchantLevelApi) DeleteSplitConfigurationRuleInput(merchantId string, splitConfigurationId string, ruleId string) SplitConfigurationMerchantLevelApiDeleteSplitConfigurationRuleInput {
	return SplitConfigurationMerchantLevelApiDeleteSplitConfigurationRuleInput{
		merchantId:           merchantId,
		splitConfigurationId: splitConfigurationId,
		ruleId:               ruleId,
	}
}

/*
DeleteSplitConfigurationRule Delete a split configuration rule

Deletes the split configuration rule specified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API - SplitConfiguration read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r SplitConfigurationMerchantLevelApiDeleteSplitConfigurationRuleInput - Request parameters, see DeleteSplitConfigurationRuleInput
@return SplitConfiguration, *http.Response, error
*/
func (a *SplitConfigurationMerchantLevelApi) DeleteSplitConfigurationRule(ctx context.Context, r SplitConfigurationMerchantLevelApiDeleteSplitConfigurationRuleInput) (SplitConfiguration, *http.Response, error) {
	res := &SplitConfiguration{}
	path := "/merchants/{merchantId}/splitConfigurations/{splitConfigurationId}/rules/{ruleId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"splitConfigurationId"+"}", url.PathEscape(common.ParameterValueToString(r.splitConfigurationId, "splitConfigurationId")), -1)
	path = strings.Replace(path, "{"+"ruleId"+"}", url.PathEscape(common.ParameterValueToString(r.ruleId, "ruleId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodDelete,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by SplitConfigurationMerchantLevelApi.GetSplitConfiguration
type SplitConfigurationMerchantLevelApiGetSplitConfigurationInput struct {
	merchantId           string
	splitConfigurationId string
}

/*
Prepare a request for GetSplitConfiguration
@param merchantId The unique identifier of the merchant account.@param splitConfigurationId The unique identifier of the split configuration.
@return SplitConfigurationMerchantLevelApiGetSplitConfigurationInput
*/
func (a *SplitConfigurationMerchantLevelApi) GetSplitConfigurationInput(merchantId string, splitConfigurationId string) SplitConfigurationMerchantLevelApiGetSplitConfigurationInput {
	return SplitConfigurationMerchantLevelApiGetSplitConfigurationInput{
		merchantId:           merchantId,
		splitConfigurationId: splitConfigurationId,
	}
}

/*
GetSplitConfiguration Get a split configuration

Returns the split configuration specified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API - SplitConfiguration read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r SplitConfigurationMerchantLevelApiGetSplitConfigurationInput - Request parameters, see GetSplitConfigurationInput
@return SplitConfiguration, *http.Response, error
*/
func (a *SplitConfigurationMerchantLevelApi) GetSplitConfiguration(ctx context.Context, r SplitConfigurationMerchantLevelApiGetSplitConfigurationInput) (SplitConfiguration, *http.Response, error) {
	res := &SplitConfiguration{}
	path := "/merchants/{merchantId}/splitConfigurations/{splitConfigurationId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"splitConfigurationId"+"}", url.PathEscape(common.ParameterValueToString(r.splitConfigurationId, "splitConfigurationId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by SplitConfigurationMerchantLevelApi.ListSplitConfigurations
type SplitConfigurationMerchantLevelApiListSplitConfigurationsInput struct {
	merchantId string
}

/*
Prepare a request for ListSplitConfigurations
@param merchantId The unique identifier of the merchant account.
@return SplitConfigurationMerchantLevelApiListSplitConfigurationsInput
*/
func (a *SplitConfigurationMerchantLevelApi) ListSplitConfigurationsInput(merchantId string) SplitConfigurationMerchantLevelApiListSplitConfigurationsInput {
	return SplitConfigurationMerchantLevelApiListSplitConfigurationsInput{
		merchantId: merchantId,
	}
}

/*
ListSplitConfigurations Get a list of split configurations

Returns the list of split configurations for the merchant account.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API - SplitConfiguration read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r SplitConfigurationMerchantLevelApiListSplitConfigurationsInput - Request parameters, see ListSplitConfigurationsInput
@return SplitConfigurationList, *http.Response, error
*/
func (a *SplitConfigurationMerchantLevelApi) ListSplitConfigurations(ctx context.Context, r SplitConfigurationMerchantLevelApiListSplitConfigurationsInput) (SplitConfigurationList, *http.Response, error) {
	res := &SplitConfigurationList{}
	path := "/merchants/{merchantId}/splitConfigurations"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by SplitConfigurationMerchantLevelApi.UpdateSplitConditions
type SplitConfigurationMerchantLevelApiUpdateSplitConditionsInput struct {
	merchantId                          string
	splitConfigurationId                string
	ruleId                              string
	updateSplitConfigurationRuleRequest *UpdateSplitConfigurationRuleRequest
}

func (r SplitConfigurationMerchantLevelApiUpdateSplitConditionsInput) UpdateSplitConfigurationRuleRequest(updateSplitConfigurationRuleRequest UpdateSplitConfigurationRuleRequest) SplitConfigurationMerchantLevelApiUpdateSplitConditionsInput {
	r.updateSplitConfigurationRuleRequest = &updateSplitConfigurationRuleRequest
	return r
}

/*
Prepare a request for UpdateSplitConditions
@param merchantId The unique identifier of the merchant account.@param splitConfigurationId The identifier of the split configuration.@param ruleId The unique identifier of the split configuration rule.
@return SplitConfigurationMerchantLevelApiUpdateSplitConditionsInput
*/
func (a *SplitConfigurationMerchantLevelApi) UpdateSplitConditionsInput(merchantId string, splitConfigurationId string, ruleId string) SplitConfigurationMerchantLevelApiUpdateSplitConditionsInput {
	return SplitConfigurationMerchantLevelApiUpdateSplitConditionsInput{
		merchantId:           merchantId,
		splitConfigurationId: splitConfigurationId,
		ruleId:               ruleId,
	}
}

/*
UpdateSplitConditions Update split conditions

Changes the conditions of the split configuration rule specified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API - SplitConfiguration read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r SplitConfigurationMerchantLevelApiUpdateSplitConditionsInput - Request parameters, see UpdateSplitConditionsInput
@return SplitConfiguration, *http.Response, error
*/
func (a *SplitConfigurationMerchantLevelApi) UpdateSplitConditions(ctx context.Context, r SplitConfigurationMerchantLevelApiUpdateSplitConditionsInput) (SplitConfiguration, *http.Response, error) {
	res := &SplitConfiguration{}
	path := "/merchants/{merchantId}/splitConfigurations/{splitConfigurationId}/rules/{ruleId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"splitConfigurationId"+"}", url.PathEscape(common.ParameterValueToString(r.splitConfigurationId, "splitConfigurationId")), -1)
	path = strings.Replace(path, "{"+"ruleId"+"}", url.PathEscape(common.ParameterValueToString(r.ruleId, "ruleId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.updateSplitConfigurationRuleRequest,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by SplitConfigurationMerchantLevelApi.UpdateSplitConfigurationDescription
type SplitConfigurationMerchantLevelApiUpdateSplitConfigurationDescriptionInput struct {
	merchantId                      string
	splitConfigurationId            string
	updateSplitConfigurationRequest *UpdateSplitConfigurationRequest
}

func (r SplitConfigurationMerchantLevelApiUpdateSplitConfigurationDescriptionInput) UpdateSplitConfigurationRequest(updateSplitConfigurationRequest UpdateSplitConfigurationRequest) SplitConfigurationMerchantLevelApiUpdateSplitConfigurationDescriptionInput {
	r.updateSplitConfigurationRequest = &updateSplitConfigurationRequest
	return r
}

/*
Prepare a request for UpdateSplitConfigurationDescription
@param merchantId The unique identifier of the merchant account.@param splitConfigurationId The unique identifier of the split configuration.
@return SplitConfigurationMerchantLevelApiUpdateSplitConfigurationDescriptionInput
*/
func (a *SplitConfigurationMerchantLevelApi) UpdateSplitConfigurationDescriptionInput(merchantId string, splitConfigurationId string) SplitConfigurationMerchantLevelApiUpdateSplitConfigurationDescriptionInput {
	return SplitConfigurationMerchantLevelApiUpdateSplitConfigurationDescriptionInput{
		merchantId:           merchantId,
		splitConfigurationId: splitConfigurationId,
	}
}

/*
UpdateSplitConfigurationDescription Update split configuration description

Changes the description of the split configuration specified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API - SplitConfiguration read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r SplitConfigurationMerchantLevelApiUpdateSplitConfigurationDescriptionInput - Request parameters, see UpdateSplitConfigurationDescriptionInput
@return SplitConfiguration, *http.Response, error
*/
func (a *SplitConfigurationMerchantLevelApi) UpdateSplitConfigurationDescription(ctx context.Context, r SplitConfigurationMerchantLevelApiUpdateSplitConfigurationDescriptionInput) (SplitConfiguration, *http.Response, error) {
	res := &SplitConfiguration{}
	path := "/merchants/{merchantId}/splitConfigurations/{splitConfigurationId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"splitConfigurationId"+"}", url.PathEscape(common.ParameterValueToString(r.splitConfigurationId, "splitConfigurationId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.updateSplitConfigurationRequest,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by SplitConfigurationMerchantLevelApi.UpdateSplitLogic
type SplitConfigurationMerchantLevelApiUpdateSplitLogicInput struct {
	merchantId                           string
	splitConfigurationId                 string
	ruleId                               string
	splitLogicId                         string
	updateSplitConfigurationLogicRequest *UpdateSplitConfigurationLogicRequest
}

func (r SplitConfigurationMerchantLevelApiUpdateSplitLogicInput) UpdateSplitConfigurationLogicRequest(updateSplitConfigurationLogicRequest UpdateSplitConfigurationLogicRequest) SplitConfigurationMerchantLevelApiUpdateSplitLogicInput {
	r.updateSplitConfigurationLogicRequest = &updateSplitConfigurationLogicRequest
	return r
}

/*
Prepare a request for UpdateSplitLogic
@param merchantId The unique identifier of the merchant account.@param splitConfigurationId The unique identifier of the split configuration.@param ruleId The unique identifier of the split configuration rule.@param splitLogicId The unique identifier of the split configuration split.
@return SplitConfigurationMerchantLevelApiUpdateSplitLogicInput
*/
func (a *SplitConfigurationMerchantLevelApi) UpdateSplitLogicInput(merchantId string, splitConfigurationId string, ruleId string, splitLogicId string) SplitConfigurationMerchantLevelApiUpdateSplitLogicInput {
	return SplitConfigurationMerchantLevelApiUpdateSplitLogicInput{
		merchantId:           merchantId,
		splitConfigurationId: splitConfigurationId,
		ruleId:               ruleId,
		splitLogicId:         splitLogicId,
	}
}

/*
UpdateSplitLogic Update the split logic

Changes the split logic specified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API - SplitConfiguration read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r SplitConfigurationMerchantLevelApiUpdateSplitLogicInput - Request parameters, see UpdateSplitLogicInput
@return SplitConfiguration, *http.Response, error
*/
func (a *SplitConfigurationMerchantLevelApi) UpdateSplitLogic(ctx context.Context, r SplitConfigurationMerchantLevelApiUpdateSplitLogicInput) (SplitConfiguration, *http.Response, error) {
	res := &SplitConfiguration{}
	path := "/merchants/{merchantId}/splitConfigurations/{splitConfigurationId}/rules/{ruleId}/splitLogic/{splitLogicId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"splitConfigurationId"+"}", url.PathEscape(common.ParameterValueToString(r.splitConfigurationId, "splitConfigurationId")), -1)
	path = strings.Replace(path, "{"+"ruleId"+"}", url.PathEscape(common.ParameterValueToString(r.ruleId, "ruleId")), -1)
	path = strings.Replace(path, "{"+"splitLogicId"+"}", url.PathEscape(common.ParameterValueToString(r.splitLogicId, "splitLogicId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.updateSplitConfigurationLogicRequest,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}
