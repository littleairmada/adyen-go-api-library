/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"context"
	_context "context"
	_nethttp "net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// TerminalSettingsStoreLevelApi TerminalSettingsStoreLevelApi service
type TerminalSettingsStoreLevelApi common.Service

type TerminalSettingsStoreLevelApiGetTerminalLogoConfig struct {
	ctx        context.Context
	merchantId string
	reference  string
	model      *string
}

// The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T.
func (r TerminalSettingsStoreLevelApiGetTerminalLogoConfig) Model(model string) TerminalSettingsStoreLevelApiGetTerminalLogoConfig {
	r.model = &model
	return r
}

/*
GetTerminalLogo Get the terminal logo

Returns the logo that is configured for a specific payment terminal model at the store identified in the path.
The logo is returned as a Base64-encoded string. You need to Base64-decode the string to get the actual image file.
This logo applies to all terminals of the specified model under the store, unless a different logo is configured for an individual terminal.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read
* Management API—Terminal settings read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param reference The reference that identifies the store.
 @return TerminalSettingsStoreLevelApiGetTerminalLogoConfig
*/
func (a *TerminalSettingsStoreLevelApi) GetTerminalLogoConfig(ctx context.Context, merchantId string, reference string) TerminalSettingsStoreLevelApiGetTerminalLogoConfig {
	return TerminalSettingsStoreLevelApiGetTerminalLogoConfig{
		ctx:        ctx,
		merchantId: merchantId,
		reference:  reference,
	}
}

/*
Get the terminal logo
Returns the logo that is configured for a specific payment terminal model at the store identified in the path.  The logo is returned as a Base64-encoded string. You need to Base64-decode the string to get the actual image file.  This logo applies to all terminals of the specified model under the store, unless a different logo is configured for an individual terminal.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal settings read * Management API—Terminal settings read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param reference The reference that identifies the store.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Logo
*/

func (a *TerminalSettingsStoreLevelApi) GetTerminalLogo(r TerminalSettingsStoreLevelApiGetTerminalLogoConfig) (Logo, *_nethttp.Response, error) {
	res := &Logo{}
	path := "/merchants/{merchantId}/stores/{reference}/terminalLogos"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"reference"+"}", url.PathEscape(common.ParameterValueToString(r.reference, "reference")), -1)
	queryString := url.Values{}
	if r.model != nil {
		common.ParameterAddToQuery(queryString, "model", r.model, "")
	}
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path+"?"+queryString.Encode(), []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalSettingsStoreLevelApiGetTerminalLogoByStoreIdConfig struct {
	ctx     context.Context
	storeId string
	model   *string
}

// The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T.
func (r TerminalSettingsStoreLevelApiGetTerminalLogoByStoreIdConfig) Model(model string) TerminalSettingsStoreLevelApiGetTerminalLogoByStoreIdConfig {
	r.model = &model
	return r
}

/*
GetTerminalLogoByStoreId Get the terminal logo

Returns the logo that is configured for a specific payment terminal model at the store identified in the path.
The logo is returned as a Base64-encoded string. You need to Base64-decode the string to get the actual image file.
This logo applies to all terminals of that model under the store unless a different logo is configured for an individual terminal.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read
* Management API—Terminal settings read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeId The unique identifier of the store.
 @return TerminalSettingsStoreLevelApiGetTerminalLogoByStoreIdConfig
*/
func (a *TerminalSettingsStoreLevelApi) GetTerminalLogoByStoreIdConfig(ctx context.Context, storeId string) TerminalSettingsStoreLevelApiGetTerminalLogoByStoreIdConfig {
	return TerminalSettingsStoreLevelApiGetTerminalLogoByStoreIdConfig{
		ctx:     ctx,
		storeId: storeId,
	}
}

/*
Get the terminal logo
Returns the logo that is configured for a specific payment terminal model at the store identified in the path.  The logo is returned as a Base64-encoded string. You need to Base64-decode the string to get the actual image file.  This logo applies to all terminals of that model under the store unless a different logo is configured for an individual terminal.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal settings read * Management API—Terminal settings read and write
 * @param storeId The unique identifier of the store.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Logo
*/

func (a *TerminalSettingsStoreLevelApi) GetTerminalLogoByStoreId(r TerminalSettingsStoreLevelApiGetTerminalLogoByStoreIdConfig) (Logo, *_nethttp.Response, error) {
	res := &Logo{}
	path := "/stores/{storeId}/terminalLogos"
	path = strings.Replace(path, "{"+"storeId"+"}", url.PathEscape(common.ParameterValueToString(r.storeId, "storeId")), -1)
	queryString := url.Values{}
	if r.model != nil {
		common.ParameterAddToQuery(queryString, "model", r.model, "")
	}
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path+"?"+queryString.Encode(), []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalSettingsStoreLevelApiGetTerminalSettingsConfig struct {
	ctx        context.Context
	merchantId string
	reference  string
}

/*
GetTerminalSettings Get terminal settings

Returns the payment terminal settings that are configured for the store identified in the path. These settings apply to all terminals under the store unless different values are configured for an individual terminal.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read
* Management API—Terminal settings read and write

For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role:
* Management API—Terminal settings Advanced read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param reference The reference that identifies the store.
 @return TerminalSettingsStoreLevelApiGetTerminalSettingsConfig
*/
func (a *TerminalSettingsStoreLevelApi) GetTerminalSettingsConfig(ctx context.Context, merchantId string, reference string) TerminalSettingsStoreLevelApiGetTerminalSettingsConfig {
	return TerminalSettingsStoreLevelApiGetTerminalSettingsConfig{
		ctx:        ctx,
		merchantId: merchantId,
		reference:  reference,
	}
}

/*
Get terminal settings
Returns the payment terminal settings that are configured for the store identified in the path. These settings apply to all terminals under the store unless different values are configured for an individual terminal.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal settings read * Management API—Terminal settings read and write  For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role: * Management API—Terminal settings Advanced read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param reference The reference that identifies the store.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TerminalSettings
*/

func (a *TerminalSettingsStoreLevelApi) GetTerminalSettings(r TerminalSettingsStoreLevelApiGetTerminalSettingsConfig) (TerminalSettings, *_nethttp.Response, error) {
	res := &TerminalSettings{}
	path := "/merchants/{merchantId}/stores/{reference}/terminalSettings"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"reference"+"}", url.PathEscape(common.ParameterValueToString(r.reference, "reference")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalSettingsStoreLevelApiGetTerminalSettingsByStoreIdConfig struct {
	ctx     context.Context
	storeId string
}

/*
GetTerminalSettingsByStoreId Get terminal settings

Returns the payment terminal settings that are configured for the store identified in the path. These settings apply to all terminals under the store unless different values are configured for an individual terminal.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read
* Management API—Terminal settings read and write

For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role:
* Management API—Terminal settings Advanced read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeId The unique identifier of the store.
 @return TerminalSettingsStoreLevelApiGetTerminalSettingsByStoreIdConfig
*/
func (a *TerminalSettingsStoreLevelApi) GetTerminalSettingsByStoreIdConfig(ctx context.Context, storeId string) TerminalSettingsStoreLevelApiGetTerminalSettingsByStoreIdConfig {
	return TerminalSettingsStoreLevelApiGetTerminalSettingsByStoreIdConfig{
		ctx:     ctx,
		storeId: storeId,
	}
}

/*
Get terminal settings
Returns the payment terminal settings that are configured for the store identified in the path. These settings apply to all terminals under the store unless different values are configured for an individual terminal.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal settings read * Management API—Terminal settings read and write  For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role: * Management API—Terminal settings Advanced read and write
 * @param storeId The unique identifier of the store.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TerminalSettings
*/

func (a *TerminalSettingsStoreLevelApi) GetTerminalSettingsByStoreId(r TerminalSettingsStoreLevelApiGetTerminalSettingsByStoreIdConfig) (TerminalSettings, *_nethttp.Response, error) {
	res := &TerminalSettings{}
	path := "/stores/{storeId}/terminalSettings"
	path = strings.Replace(path, "{"+"storeId"+"}", url.PathEscape(common.ParameterValueToString(r.storeId, "storeId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalSettingsStoreLevelApiUpdateTerminalLogoConfig struct {
	ctx        context.Context
	merchantId string
	reference  string
	model      *string
	logo       *Logo
}

// The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T
func (r TerminalSettingsStoreLevelApiUpdateTerminalLogoConfig) Model(model string) TerminalSettingsStoreLevelApiUpdateTerminalLogoConfig {
	r.model = &model
	return r
}

func (r TerminalSettingsStoreLevelApiUpdateTerminalLogoConfig) Logo(logo Logo) TerminalSettingsStoreLevelApiUpdateTerminalLogoConfig {
	r.logo = &logo
	return r
}

/*
UpdateTerminalLogo Update the terminal logo

Updates the logo that is configured for a specific payment terminal model at the store identified in the path. You can update the logo for only one terminal model at a time.
This logo applies to all terminals of the specified model under the store, unless a different logo is configured for an individual terminal.

* To change the logo, specify the image file as a Base64-encoded string.
* To restore the logo inherited from a higher level (merchant or company account), specify an empty logo value.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param reference The reference that identifies the store.
 @return TerminalSettingsStoreLevelApiUpdateTerminalLogoConfig
*/
func (a *TerminalSettingsStoreLevelApi) UpdateTerminalLogoConfig(ctx context.Context, merchantId string, reference string) TerminalSettingsStoreLevelApiUpdateTerminalLogoConfig {
	return TerminalSettingsStoreLevelApiUpdateTerminalLogoConfig{
		ctx:        ctx,
		merchantId: merchantId,
		reference:  reference,
	}
}

/*
Update the terminal logo
Updates the logo that is configured for a specific payment terminal model at the store identified in the path. You can update the logo for only one terminal model at a time. This logo applies to all terminals of the specified model under the store, unless a different logo is configured for an individual terminal.   * To change the logo, specify the image file as a Base64-encoded string. * To restore the logo inherited from a higher level (merchant or company account), specify an empty logo value.  To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal settings read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param reference The reference that identifies the store.
 * @param req Logo - reference of Logo).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Logo
*/

func (a *TerminalSettingsStoreLevelApi) UpdateTerminalLogo(r TerminalSettingsStoreLevelApiUpdateTerminalLogoConfig) (Logo, *_nethttp.Response, error) {
	res := &Logo{}
	path := "/merchants/{merchantId}/stores/{reference}/terminalLogos"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"reference"+"}", url.PathEscape(common.ParameterValueToString(r.reference, "reference")), -1)
	queryString := url.Values{}
	if r.model != nil {
		common.ParameterAddToQuery(queryString, "model", r.model, "")
	}
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPatch, r.logo, res, a.BasePath()+path+"?"+queryString.Encode(), []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalSettingsStoreLevelApiUpdateTerminalLogoByStoreIdConfig struct {
	ctx     context.Context
	storeId string
	model   *string
	logo    *Logo
}

// The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T.
func (r TerminalSettingsStoreLevelApiUpdateTerminalLogoByStoreIdConfig) Model(model string) TerminalSettingsStoreLevelApiUpdateTerminalLogoByStoreIdConfig {
	r.model = &model
	return r
}

func (r TerminalSettingsStoreLevelApiUpdateTerminalLogoByStoreIdConfig) Logo(logo Logo) TerminalSettingsStoreLevelApiUpdateTerminalLogoByStoreIdConfig {
	r.logo = &logo
	return r
}

/*
UpdateTerminalLogoByStoreId Update the terminal logo

Updates the logo that is configured for a specific payment terminal model at the store identified in the path. You can update the logo for only one terminal model at a time.
This logo applies to all terminals of the specified model under the store, unless a different logo is configured for an individual terminal.

* To change the logo, specify the image file as a Base64-encoded string.
* To restore the logo inherited from a higher level (merchant or company account), specify an empty logo value.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeId The unique identifier of the store.
 @return TerminalSettingsStoreLevelApiUpdateTerminalLogoByStoreIdConfig
*/
func (a *TerminalSettingsStoreLevelApi) UpdateTerminalLogoByStoreIdConfig(ctx context.Context, storeId string) TerminalSettingsStoreLevelApiUpdateTerminalLogoByStoreIdConfig {
	return TerminalSettingsStoreLevelApiUpdateTerminalLogoByStoreIdConfig{
		ctx:     ctx,
		storeId: storeId,
	}
}

/*
Update the terminal logo
Updates the logo that is configured for a specific payment terminal model at the store identified in the path. You can update the logo for only one terminal model at a time. This logo applies to all terminals of the specified model under the store, unless a different logo is configured for an individual terminal.   * To change the logo, specify the image file as a Base64-encoded string. * To restore the logo inherited from a higher level (merchant or company account), specify an empty logo value.  To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal settings read and write
 * @param storeId The unique identifier of the store.
 * @param req Logo - reference of Logo).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Logo
*/

func (a *TerminalSettingsStoreLevelApi) UpdateTerminalLogoByStoreId(r TerminalSettingsStoreLevelApiUpdateTerminalLogoByStoreIdConfig) (Logo, *_nethttp.Response, error) {
	res := &Logo{}
	path := "/stores/{storeId}/terminalLogos"
	path = strings.Replace(path, "{"+"storeId"+"}", url.PathEscape(common.ParameterValueToString(r.storeId, "storeId")), -1)
	queryString := url.Values{}
	if r.model != nil {
		common.ParameterAddToQuery(queryString, "model", r.model, "")
	}
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPatch, r.logo, res, a.BasePath()+path+"?"+queryString.Encode(), []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalSettingsStoreLevelApiUpdateTerminalSettingsConfig struct {
	ctx              context.Context
	merchantId       string
	reference        string
	terminalSettings *TerminalSettings
}

func (r TerminalSettingsStoreLevelApiUpdateTerminalSettingsConfig) TerminalSettings(terminalSettings TerminalSettings) TerminalSettingsStoreLevelApiUpdateTerminalSettingsConfig {
	r.terminalSettings = &terminalSettings
	return r
}

/*
UpdateTerminalSettings Update terminal settings

Updates payment terminal settings for the store identified in the path. These settings apply to all terminals under the store, unless different values are configured for an individual terminal.

* To change a parameter value, include the full object that contains the parameter, even if you don't want to change all parameters in the object.
* To restore a parameter value inherited from a higher level, include the full object that contains the parameter, and specify an empty value for the parameter or omit the parameter.
* Objects that are not included in the request are not updated.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read and write

For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role:
* Management API—Terminal settings Advanced read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param reference The reference that identifies the store.
 @return TerminalSettingsStoreLevelApiUpdateTerminalSettingsConfig
*/
func (a *TerminalSettingsStoreLevelApi) UpdateTerminalSettingsConfig(ctx context.Context, merchantId string, reference string) TerminalSettingsStoreLevelApiUpdateTerminalSettingsConfig {
	return TerminalSettingsStoreLevelApiUpdateTerminalSettingsConfig{
		ctx:        ctx,
		merchantId: merchantId,
		reference:  reference,
	}
}

/*
Update terminal settings
Updates payment terminal settings for the store identified in the path. These settings apply to all terminals under the store, unless different values are configured for an individual terminal.  * To change a parameter value, include the full object that contains the parameter, even if you don&#39;t want to change all parameters in the object. * To restore a parameter value inherited from a higher level, include the full object that contains the parameter, and specify an empty value for the parameter or omit the parameter. * Objects that are not included in the request are not updated.  To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal settings read and write  For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role: * Management API—Terminal settings Advanced read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param reference The reference that identifies the store.
 * @param req TerminalSettings - reference of TerminalSettings).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TerminalSettings
*/

func (a *TerminalSettingsStoreLevelApi) UpdateTerminalSettings(r TerminalSettingsStoreLevelApiUpdateTerminalSettingsConfig) (TerminalSettings, *_nethttp.Response, error) {
	res := &TerminalSettings{}
	path := "/merchants/{merchantId}/stores/{reference}/terminalSettings"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"reference"+"}", url.PathEscape(common.ParameterValueToString(r.reference, "reference")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPatch, r.terminalSettings, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalSettingsStoreLevelApiUpdateTerminalSettingsByStoreIdConfig struct {
	ctx              context.Context
	storeId          string
	terminalSettings *TerminalSettings
}

func (r TerminalSettingsStoreLevelApiUpdateTerminalSettingsByStoreIdConfig) TerminalSettings(terminalSettings TerminalSettings) TerminalSettingsStoreLevelApiUpdateTerminalSettingsByStoreIdConfig {
	r.terminalSettings = &terminalSettings
	return r
}

/*
UpdateTerminalSettingsByStoreId Update terminal settings

Updates payment terminal settings for the store identified in the path. These settings apply to all terminals under the store, unless different values are configured for an individual terminal.

* To change a parameter value, include the full object that contains the parameter, even if you don't want to change all parameters in the object.
* To restore a parameter value inherited from a higher level, include the full object that contains the parameter, and specify an empty value for the parameter or omit the parameter.
* Objects that are not included in the request are not updated.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read and write

For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role:
* Management API—Terminal settings Advanced read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeId The unique identifier of the store.
 @return TerminalSettingsStoreLevelApiUpdateTerminalSettingsByStoreIdConfig
*/
func (a *TerminalSettingsStoreLevelApi) UpdateTerminalSettingsByStoreIdConfig(ctx context.Context, storeId string) TerminalSettingsStoreLevelApiUpdateTerminalSettingsByStoreIdConfig {
	return TerminalSettingsStoreLevelApiUpdateTerminalSettingsByStoreIdConfig{
		ctx:     ctx,
		storeId: storeId,
	}
}

/*
Update terminal settings
Updates payment terminal settings for the store identified in the path. These settings apply to all terminals under the store, unless different values are configured for an individual terminal.  * To change a parameter value, include the full object that contains the parameter, even if you don&#39;t want to change all parameters in the object. * To restore a parameter value inherited from a higher level, include the full object that contains the parameter, and specify an empty value for the parameter or omit the parameter. * Objects that are not included in the request are not updated.  To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal settings read and write  For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role: * Management API—Terminal settings Advanced read and write
 * @param storeId The unique identifier of the store.
 * @param req TerminalSettings - reference of TerminalSettings).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TerminalSettings
*/

func (a *TerminalSettingsStoreLevelApi) UpdateTerminalSettingsByStoreId(r TerminalSettingsStoreLevelApiUpdateTerminalSettingsByStoreIdConfig) (TerminalSettings, *_nethttp.Response, error) {
	res := &TerminalSettings{}
	path := "/stores/{storeId}/terminalSettings"
	path = strings.Replace(path, "{"+"storeId"+"}", url.PathEscape(common.ParameterValueToString(r.storeId, "storeId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPatch, r.terminalSettings, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}
