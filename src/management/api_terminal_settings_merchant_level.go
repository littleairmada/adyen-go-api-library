/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"context"
    "net/http"
    "net/url"
    "strings"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// TerminalSettingsMerchantLevelApi service
type TerminalSettingsMerchantLevelApi common.Service

// All parameters accepted by TerminalSettingsMerchantLevelApi.GetTerminalLogo
type TerminalSettingsMerchantLevelApiGetTerminalLogoInput struct {
	merchantId string
	model *string
}

// The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T.
func (r TerminalSettingsMerchantLevelApiGetTerminalLogoInput) Model(model string) TerminalSettingsMerchantLevelApiGetTerminalLogoInput {
	r.model = &model
	return r
}


/*
Prepare a request for GetTerminalLogo
@param merchantId The unique identifier of the merchant account.
@return TerminalSettingsMerchantLevelApiGetTerminalLogoInput
*/
func (a *TerminalSettingsMerchantLevelApi) GetTerminalLogoInput(merchantId string) TerminalSettingsMerchantLevelApiGetTerminalLogoInput {
	return TerminalSettingsMerchantLevelApiGetTerminalLogoInput{
		merchantId: merchantId,
	}
}

/*
GetTerminalLogo Get the terminal logo

Returns the logo that is configured for a specific payment terminal model at the merchant account identified in the path. 
The logo is returned as a Base64-encoded string. You need to Base64-decode the string to get the actual image file. 
This logo applies to all terminals of the specified model under the merchant account, unless a different logo is configured at a lower level (store or individual terminal).

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read
* Management API—Terminal settings read and write

In the live environment, requests to this endpoint are subject to [rate limits](https://docs.adyen.com/point-of-sale/automating-terminal-management#rate-limits-in-the-live-environment).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalSettingsMerchantLevelApiGetTerminalLogoInput - Request parameters, see GetTerminalLogoInput
@return Logo, *http.Response, error
*/
func (a *TerminalSettingsMerchantLevelApi) GetTerminalLogo(ctx context.Context, r TerminalSettingsMerchantLevelApiGetTerminalLogoInput) (Logo, *http.Response, error) {
    res := &Logo{}
	path := "/merchants/{merchantId}/terminalLogos"
    path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
    queryParams := url.Values{}
    headerParams := make(map[string]string)
    if r.model != nil {
        common.ParameterAddToQuery(queryParams, "model", r.model, "")
    }
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


// All parameters accepted by TerminalSettingsMerchantLevelApi.GetTerminalSettings
type TerminalSettingsMerchantLevelApiGetTerminalSettingsInput struct {
	merchantId string
}


/*
Prepare a request for GetTerminalSettings
@param merchantId The unique identifier of the merchant account.
@return TerminalSettingsMerchantLevelApiGetTerminalSettingsInput
*/
func (a *TerminalSettingsMerchantLevelApi) GetTerminalSettingsInput(merchantId string) TerminalSettingsMerchantLevelApiGetTerminalSettingsInput {
	return TerminalSettingsMerchantLevelApiGetTerminalSettingsInput{
		merchantId: merchantId,
	}
}

/*
GetTerminalSettings Get terminal settings

Returns the payment terminal settings that are configured for the merchant account identified in the path. These settings apply to all terminals under the merchant account unless different values are configured at a lower level (store or individual terminal).

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read
* Management API—Terminal settings read and write

For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role:
* Management API—Terminal settings Advanced read and write

In the live environment, requests to this endpoint are subject to [rate limits](https://docs.adyen.com/point-of-sale/automating-terminal-management#rate-limits-in-the-live-environment).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalSettingsMerchantLevelApiGetTerminalSettingsInput - Request parameters, see GetTerminalSettingsInput
@return TerminalSettings, *http.Response, error
*/
func (a *TerminalSettingsMerchantLevelApi) GetTerminalSettings(ctx context.Context, r TerminalSettingsMerchantLevelApiGetTerminalSettingsInput) (TerminalSettings, *http.Response, error) {
    res := &TerminalSettings{}
	path := "/merchants/{merchantId}/terminalSettings"
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


// All parameters accepted by TerminalSettingsMerchantLevelApi.UpdateTerminalLogo
type TerminalSettingsMerchantLevelApiUpdateTerminalLogoInput struct {
	merchantId string
	model *string
	logo *Logo
}

// The terminal model. Allowed values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T.
func (r TerminalSettingsMerchantLevelApiUpdateTerminalLogoInput) Model(model string) TerminalSettingsMerchantLevelApiUpdateTerminalLogoInput {
	r.model = &model
	return r
}

func (r TerminalSettingsMerchantLevelApiUpdateTerminalLogoInput) Logo(logo Logo) TerminalSettingsMerchantLevelApiUpdateTerminalLogoInput {
	r.logo = &logo
	return r
}


/*
Prepare a request for UpdateTerminalLogo
@param merchantId The unique identifier of the merchant account.
@return TerminalSettingsMerchantLevelApiUpdateTerminalLogoInput
*/
func (a *TerminalSettingsMerchantLevelApi) UpdateTerminalLogoInput(merchantId string) TerminalSettingsMerchantLevelApiUpdateTerminalLogoInput {
	return TerminalSettingsMerchantLevelApiUpdateTerminalLogoInput{
		merchantId: merchantId,
	}
}

/*
UpdateTerminalLogo Update the terminal logo

Updates the logo for a specific payment terminal model at the merchant account identified in the path. You can update the logo for only one terminal model at a time. 
This logo applies to all terminals of the specified model under the merchant account, unless a different logo is configured at a lower level (store or individual terminal).

* To change the logo, specify the image file as a Base64-encoded string.
* To restore the logo inherited from the company account, specify an empty logo value.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read and write

In the live environment, requests to this endpoint are subject to [rate limits](https://docs.adyen.com/point-of-sale/automating-terminal-management#rate-limits-in-the-live-environment).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalSettingsMerchantLevelApiUpdateTerminalLogoInput - Request parameters, see UpdateTerminalLogoInput
@return Logo, *http.Response, error
*/
func (a *TerminalSettingsMerchantLevelApi) UpdateTerminalLogo(ctx context.Context, r TerminalSettingsMerchantLevelApiUpdateTerminalLogoInput) (Logo, *http.Response, error) {
    res := &Logo{}
	path := "/merchants/{merchantId}/terminalLogos"
    path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
    queryParams := url.Values{}
    headerParams := make(map[string]string)
    if r.model != nil {
        common.ParameterAddToQuery(queryParams, "model", r.model, "")
    }
    httpRes, err := common.SendAPIRequest(
        ctx,
        a.Client,
        r.logo,
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


// All parameters accepted by TerminalSettingsMerchantLevelApi.UpdateTerminalSettings
type TerminalSettingsMerchantLevelApiUpdateTerminalSettingsInput struct {
	merchantId string
	terminalSettings *TerminalSettings
}

func (r TerminalSettingsMerchantLevelApiUpdateTerminalSettingsInput) TerminalSettings(terminalSettings TerminalSettings) TerminalSettingsMerchantLevelApiUpdateTerminalSettingsInput {
	r.terminalSettings = &terminalSettings
	return r
}


/*
Prepare a request for UpdateTerminalSettings
@param merchantId The unique identifier of the merchant account.
@return TerminalSettingsMerchantLevelApiUpdateTerminalSettingsInput
*/
func (a *TerminalSettingsMerchantLevelApi) UpdateTerminalSettingsInput(merchantId string) TerminalSettingsMerchantLevelApiUpdateTerminalSettingsInput {
	return TerminalSettingsMerchantLevelApiUpdateTerminalSettingsInput{
		merchantId: merchantId,
	}
}

/*
UpdateTerminalSettings Update terminal settings

Updates payment terminal settings for the merchant account identified in the path.
These settings apply to all terminals under the merchant account, unless different values are configured at a lower level (store or individual terminal).

* To change a parameter value, include the full object that contains the parameter, even if you don't want to change all parameters in the object.
* To restore a parameter value inherited from a higher level, include the full object that contains the parameter, and specify an empty value for the parameter or omit the parameter.
* Objects that are not included in the request are not updated.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read and write



For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role:
* Management API—Terminal settings Advanced read and write

In the live environment, requests to this endpoint are subject to [rate limits](https://docs.adyen.com/point-of-sale/automating-terminal-management#rate-limits-in-the-live-environment).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalSettingsMerchantLevelApiUpdateTerminalSettingsInput - Request parameters, see UpdateTerminalSettingsInput
@return TerminalSettings, *http.Response, error
*/
func (a *TerminalSettingsMerchantLevelApi) UpdateTerminalSettings(ctx context.Context, r TerminalSettingsMerchantLevelApiUpdateTerminalSettingsInput) (TerminalSettings, *http.Response, error) {
    res := &TerminalSettings{}
	path := "/merchants/{merchantId}/terminalSettings"
    path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
    queryParams := url.Values{}
    headerParams := make(map[string]string)
    httpRes, err := common.SendAPIRequest(
        ctx,
        a.Client,
        r.terminalSettings,
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

