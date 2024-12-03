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

// TerminalActionsCompanyLevelApi service
type TerminalActionsCompanyLevelApi common.Service

// All parameters accepted by TerminalActionsCompanyLevelApi.GetTerminalAction
type TerminalActionsCompanyLevelApiGetTerminalActionInput struct {
	companyId string
	actionId string
}


/*
Prepare a request for GetTerminalAction
@param companyId The unique identifier of the company account.@param actionId The unique identifier of the terminal action.
@return TerminalActionsCompanyLevelApiGetTerminalActionInput
*/
func (a *TerminalActionsCompanyLevelApi) GetTerminalActionInput(companyId string, actionId string) TerminalActionsCompanyLevelApiGetTerminalActionInput {
	return TerminalActionsCompanyLevelApiGetTerminalActionInput{
		companyId: companyId,
		actionId: actionId,
	}
}

/*
GetTerminalAction Get terminal action

Returns the details of the [terminal action](https://docs.adyen.com/point-of-sale/automating-terminal-management/terminal-actions-api) identified in the path.
To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal actions read
* Management API—Terminal actions read and write

In the live environment, requests to this endpoint are subject to [rate limits](https://docs.adyen.com/point-of-sale/automating-terminal-management#rate-limits-in-the-live-environment).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalActionsCompanyLevelApiGetTerminalActionInput - Request parameters, see GetTerminalActionInput
@return ExternalTerminalAction, *http.Response, error
*/
func (a *TerminalActionsCompanyLevelApi) GetTerminalAction(ctx context.Context, r TerminalActionsCompanyLevelApiGetTerminalActionInput) (ExternalTerminalAction, *http.Response, error) {
    res := &ExternalTerminalAction{}
	path := "/companies/{companyId}/terminalActions/{actionId}"
    path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
    path = strings.Replace(path, "{"+"actionId"+"}", url.PathEscape(common.ParameterValueToString(r.actionId, "actionId")), -1)
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


// All parameters accepted by TerminalActionsCompanyLevelApi.ListTerminalActions
type TerminalActionsCompanyLevelApiListTerminalActionsInput struct {
	companyId string
	pageNumber *int32
	pageSize *int32
	status *string
	type_ *string
}

// The number of the page to fetch.
func (r TerminalActionsCompanyLevelApiListTerminalActionsInput) PageNumber(pageNumber int32) TerminalActionsCompanyLevelApiListTerminalActionsInput {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page, maximum 100. The default is 20 items on a page.
func (r TerminalActionsCompanyLevelApiListTerminalActionsInput) PageSize(pageSize int32) TerminalActionsCompanyLevelApiListTerminalActionsInput {
	r.pageSize = &pageSize
	return r
}

// Returns terminal actions with the specified status.  Allowed values: **pending**, **successful**, **failed**, **cancelled**, **tryLater**.
func (r TerminalActionsCompanyLevelApiListTerminalActionsInput) Status(status string) TerminalActionsCompanyLevelApiListTerminalActionsInput {
	r.status = &status
	return r
}

// Returns terminal actions of the specified type.  Allowed values: **InstallAndroidApp**, **UninstallAndroidApp**, **InstallAndroidCertificate**, **UninstallAndroidCertificate**.
func (r TerminalActionsCompanyLevelApiListTerminalActionsInput) Type_(type_ string) TerminalActionsCompanyLevelApiListTerminalActionsInput {
	r.type_ = &type_
	return r
}


/*
Prepare a request for ListTerminalActions
@param companyId The unique identifier of the company account.
@return TerminalActionsCompanyLevelApiListTerminalActionsInput
*/
func (a *TerminalActionsCompanyLevelApi) ListTerminalActionsInput(companyId string) TerminalActionsCompanyLevelApiListTerminalActionsInput {
	return TerminalActionsCompanyLevelApiListTerminalActionsInput{
		companyId: companyId,
	}
}

/*
ListTerminalActions Get a list of terminal actions

Returns the [terminal actions](https://docs.adyen.com/point-of-sale/automating-terminal-management/terminal-actions-api) that have been scheduled for the company identified in the path.The response doesn't include actions that are scheduled by Adyen.
To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal actions read
* Management API—Terminal actions read and write

In the live environment, requests to this endpoint are subject to [rate limits](https://docs.adyen.com/point-of-sale/automating-terminal-management#rate-limits-in-the-live-environment).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalActionsCompanyLevelApiListTerminalActionsInput - Request parameters, see ListTerminalActionsInput
@return ListExternalTerminalActionsResponse, *http.Response, error
*/
func (a *TerminalActionsCompanyLevelApi) ListTerminalActions(ctx context.Context, r TerminalActionsCompanyLevelApiListTerminalActionsInput) (ListExternalTerminalActionsResponse, *http.Response, error) {
    res := &ListExternalTerminalActionsResponse{}
	path := "/companies/{companyId}/terminalActions"
    path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
    queryParams := url.Values{}
    headerParams := make(map[string]string)
    if r.pageNumber != nil {
        common.ParameterAddToQuery(queryParams, "pageNumber", r.pageNumber, "")
    }
    if r.pageSize != nil {
        common.ParameterAddToQuery(queryParams, "pageSize", r.pageSize, "")
    }
    if r.status != nil {
        common.ParameterAddToQuery(queryParams, "status", r.status, "")
    }
    if r.type_ != nil {
        common.ParameterAddToQuery(queryParams, "type", r.type_, "")
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

