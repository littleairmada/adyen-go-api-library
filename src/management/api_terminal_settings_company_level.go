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

// TerminalSettingsCompanyLevelApi service
type TerminalSettingsCompanyLevelApi common.Service

// All parameters accepted by TerminalSettingsCompanyLevelApi.GetTerminalLogo
type TerminalSettingsCompanyLevelApiGetTerminalLogoInput struct {
	companyId string
	model     *string
}

// The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T.
func (r TerminalSettingsCompanyLevelApiGetTerminalLogoInput) Model(model string) TerminalSettingsCompanyLevelApiGetTerminalLogoInput {
	r.model = &model
	return r
}

/*
Prepare a request for GetTerminalLogo
@param companyId The unique identifier of the company account.
@return TerminalSettingsCompanyLevelApiGetTerminalLogoInput
*/
func (a *TerminalSettingsCompanyLevelApi) GetTerminalLogoInput(companyId string) TerminalSettingsCompanyLevelApiGetTerminalLogoInput {
	return TerminalSettingsCompanyLevelApiGetTerminalLogoInput{
		companyId: companyId,
	}
}

/*
GetTerminalLogo Get the terminal logo

Returns the logo that is configured for a specific payment terminal model at the company identified in the path.

The logo is returned as a Base64-encoded string. You need to Base64-decode the string to get the actual image file.
This logo applies to all terminals of the specified model under the company, unless a different logo is configured at a lower level (merchant account, store, or individual terminal).

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read
* Management API—Terminal settings read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalSettingsCompanyLevelApiGetTerminalLogoInput - Request parameters, see GetTerminalLogoInput
@return Logo, *http.Response, error
*/
func (a *TerminalSettingsCompanyLevelApi) GetTerminalLogo(ctx context.Context, r TerminalSettingsCompanyLevelApiGetTerminalLogoInput) (Logo, *http.Response, error) {
	res := &Logo{}
	path := "/companies/{companyId}/terminalLogos"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
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

// All parameters accepted by TerminalSettingsCompanyLevelApi.GetTerminalSettings
type TerminalSettingsCompanyLevelApiGetTerminalSettingsInput struct {
	companyId string
}

/*
Prepare a request for GetTerminalSettings
@param companyId The unique identifier of the company account.
@return TerminalSettingsCompanyLevelApiGetTerminalSettingsInput
*/
func (a *TerminalSettingsCompanyLevelApi) GetTerminalSettingsInput(companyId string) TerminalSettingsCompanyLevelApiGetTerminalSettingsInput {
	return TerminalSettingsCompanyLevelApiGetTerminalSettingsInput{
		companyId: companyId,
	}
}

/*
GetTerminalSettings Get terminal settings

Returns the payment terminal settings that are configured for the company identified in the path. These settings apply to all terminals under the company, unless different values are configured at a lower level (merchant account, store, or individual terminal).

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read
* Management API—Terminal settings read and write

For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role:
* Management API—Terminal settings Advanced read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalSettingsCompanyLevelApiGetTerminalSettingsInput - Request parameters, see GetTerminalSettingsInput
@return TerminalSettings, *http.Response, error
*/
func (a *TerminalSettingsCompanyLevelApi) GetTerminalSettings(ctx context.Context, r TerminalSettingsCompanyLevelApiGetTerminalSettingsInput) (TerminalSettings, *http.Response, error) {
	res := &TerminalSettings{}
	path := "/companies/{companyId}/terminalSettings"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
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

// All parameters accepted by TerminalSettingsCompanyLevelApi.UpdateTerminalLogo
type TerminalSettingsCompanyLevelApiUpdateTerminalLogoInput struct {
	companyId string
	model     *string
	logo      *Logo
}

// The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T.
func (r TerminalSettingsCompanyLevelApiUpdateTerminalLogoInput) Model(model string) TerminalSettingsCompanyLevelApiUpdateTerminalLogoInput {
	r.model = &model
	return r
}

func (r TerminalSettingsCompanyLevelApiUpdateTerminalLogoInput) Logo(logo Logo) TerminalSettingsCompanyLevelApiUpdateTerminalLogoInput {
	r.logo = &logo
	return r
}

/*
Prepare a request for UpdateTerminalLogo
@param companyId The unique identifier of the company account.
@return TerminalSettingsCompanyLevelApiUpdateTerminalLogoInput
*/
func (a *TerminalSettingsCompanyLevelApi) UpdateTerminalLogoInput(companyId string) TerminalSettingsCompanyLevelApiUpdateTerminalLogoInput {
	return TerminalSettingsCompanyLevelApiUpdateTerminalLogoInput{
		companyId: companyId,
	}
}

/*
UpdateTerminalLogo Update the terminal logo

Updates the logo that is configured for a specific payment terminal model at the company identified in the path. You can update the logo for only one terminal model at a time.
This logo applies to all terminals of the specified model under the company, unless a different logo is configured at a lower level (merchant account, store, or individual terminal).
* To change the logo, specify the image file as a Base64-encoded string.
* To restore the logo inherited from the Adyen PSP level, specify an empty logo value.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalSettingsCompanyLevelApiUpdateTerminalLogoInput - Request parameters, see UpdateTerminalLogoInput
@return Logo, *http.Response, error
*/
func (a *TerminalSettingsCompanyLevelApi) UpdateTerminalLogo(ctx context.Context, r TerminalSettingsCompanyLevelApiUpdateTerminalLogoInput) (Logo, *http.Response, error) {
	res := &Logo{}
	path := "/companies/{companyId}/terminalLogos"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
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

// All parameters accepted by TerminalSettingsCompanyLevelApi.UpdateTerminalSettings
type TerminalSettingsCompanyLevelApiUpdateTerminalSettingsInput struct {
	companyId        string
	terminalSettings *TerminalSettings
}

func (r TerminalSettingsCompanyLevelApiUpdateTerminalSettingsInput) TerminalSettings(terminalSettings TerminalSettings) TerminalSettingsCompanyLevelApiUpdateTerminalSettingsInput {
	r.terminalSettings = &terminalSettings
	return r
}

/*
Prepare a request for UpdateTerminalSettings
@param companyId The unique identifier of the company account.
@return TerminalSettingsCompanyLevelApiUpdateTerminalSettingsInput
*/
func (a *TerminalSettingsCompanyLevelApi) UpdateTerminalSettingsInput(companyId string) TerminalSettingsCompanyLevelApiUpdateTerminalSettingsInput {
	return TerminalSettingsCompanyLevelApiUpdateTerminalSettingsInput{
		companyId: companyId,
	}
}

/*
UpdateTerminalSettings Update terminal settings

Updates payment terminal settings for the company identified in the path. These settings apply to all terminals under the company, unless different values are configured at a lower level (merchant account, store, or individual terminal).

* To change a parameter value, include the full object that contains the parameter, even if you don't want to change all parameters in the object.
* To restore a parameter value inherited from the Adyen PSP level, include the full object that contains the parameter, and specify an empty value for the parameter or omit the parameter.
* Objects that are not included in the request are not updated.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read and write

For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role:
* Management API—Terminal settings Advanced read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalSettingsCompanyLevelApiUpdateTerminalSettingsInput - Request parameters, see UpdateTerminalSettingsInput
@return TerminalSettings, *http.Response, error
*/
func (a *TerminalSettingsCompanyLevelApi) UpdateTerminalSettings(ctx context.Context, r TerminalSettingsCompanyLevelApiUpdateTerminalSettingsInput) (TerminalSettings, *http.Response, error) {
	res := &TerminalSettings{}
	path := "/companies/{companyId}/terminalSettings"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
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
