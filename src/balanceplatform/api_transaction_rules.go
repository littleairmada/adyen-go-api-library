/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"context"
    "net/http"
    "net/url"
    "strings"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// TransactionRulesApi service
type TransactionRulesApi common.Service

// All parameters accepted by TransactionRulesApi.CreateTransactionRule
type TransactionRulesApiCreateTransactionRuleInput struct {
	transactionRuleInfo *TransactionRuleInfo
}

func (r TransactionRulesApiCreateTransactionRuleInput) TransactionRuleInfo(transactionRuleInfo TransactionRuleInfo) TransactionRulesApiCreateTransactionRuleInput {
	r.transactionRuleInfo = &transactionRuleInfo
	return r
}


/*
Prepare a request for CreateTransactionRule

@return TransactionRulesApiCreateTransactionRuleInput
*/
func (a *TransactionRulesApi) CreateTransactionRuleInput() TransactionRulesApiCreateTransactionRuleInput {
	return TransactionRulesApiCreateTransactionRuleInput{
	}
}

/*
CreateTransactionRule Create a transaction rule

Creates a [transaction rule](https://docs.adyen.com/issuing/transaction-rules). When your user makes a transaction with their Adyen-issued card, the transaction is allowed or declined based on the conditions and outcome defined in the transaction rule. You can apply the transaction rule to several cards, such as all the cards in your platform, or to a specific card. For use cases, see [examples](https://docs.adyen.com/issuing/transaction-rules/examples).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransactionRulesApiCreateTransactionRuleInput - Request parameters, see CreateTransactionRuleInput
@return TransactionRule, *http.Response, error
*/
func (a *TransactionRulesApi) CreateTransactionRule(ctx context.Context, r TransactionRulesApiCreateTransactionRuleInput) (TransactionRule, *http.Response, error) {
    res := &TransactionRule{}
	path := "/transactionRules"
    queryParams := url.Values{}
    headerParams := make(map[string]string)
    httpRes, err := common.SendAPIRequest(
        ctx,
        a.Client,
        r.transactionRuleInfo,
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


// All parameters accepted by TransactionRulesApi.DeleteTransactionRule
type TransactionRulesApiDeleteTransactionRuleInput struct {
	transactionRuleId string
}


/*
Prepare a request for DeleteTransactionRule
@param transactionRuleId The unique identifier of the transaction rule.
@return TransactionRulesApiDeleteTransactionRuleInput
*/
func (a *TransactionRulesApi) DeleteTransactionRuleInput(transactionRuleId string) TransactionRulesApiDeleteTransactionRuleInput {
	return TransactionRulesApiDeleteTransactionRuleInput{
		transactionRuleId: transactionRuleId,
	}
}

/*
DeleteTransactionRule Delete a transaction rule

Deletes a transaction rule.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransactionRulesApiDeleteTransactionRuleInput - Request parameters, see DeleteTransactionRuleInput
@return TransactionRule, *http.Response, error
*/
func (a *TransactionRulesApi) DeleteTransactionRule(ctx context.Context, r TransactionRulesApiDeleteTransactionRuleInput) (TransactionRule, *http.Response, error) {
    res := &TransactionRule{}
	path := "/transactionRules/{transactionRuleId}"
    path = strings.Replace(path, "{"+"transactionRuleId"+"}", url.PathEscape(common.ParameterValueToString(r.transactionRuleId, "transactionRuleId")), -1)
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


// All parameters accepted by TransactionRulesApi.GetTransactionRule
type TransactionRulesApiGetTransactionRuleInput struct {
	transactionRuleId string
}


/*
Prepare a request for GetTransactionRule
@param transactionRuleId The unique identifier of the transaction rule.
@return TransactionRulesApiGetTransactionRuleInput
*/
func (a *TransactionRulesApi) GetTransactionRuleInput(transactionRuleId string) TransactionRulesApiGetTransactionRuleInput {
	return TransactionRulesApiGetTransactionRuleInput{
		transactionRuleId: transactionRuleId,
	}
}

/*
GetTransactionRule Get a transaction rule

Returns the details of a transaction rule.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransactionRulesApiGetTransactionRuleInput - Request parameters, see GetTransactionRuleInput
@return TransactionRuleResponse, *http.Response, error
*/
func (a *TransactionRulesApi) GetTransactionRule(ctx context.Context, r TransactionRulesApiGetTransactionRuleInput) (TransactionRuleResponse, *http.Response, error) {
    res := &TransactionRuleResponse{}
	path := "/transactionRules/{transactionRuleId}"
    path = strings.Replace(path, "{"+"transactionRuleId"+"}", url.PathEscape(common.ParameterValueToString(r.transactionRuleId, "transactionRuleId")), -1)
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


// All parameters accepted by TransactionRulesApi.UpdateTransactionRule
type TransactionRulesApiUpdateTransactionRuleInput struct {
	transactionRuleId string
	transactionRuleInfo *TransactionRuleInfo
}

func (r TransactionRulesApiUpdateTransactionRuleInput) TransactionRuleInfo(transactionRuleInfo TransactionRuleInfo) TransactionRulesApiUpdateTransactionRuleInput {
	r.transactionRuleInfo = &transactionRuleInfo
	return r
}


/*
Prepare a request for UpdateTransactionRule
@param transactionRuleId The unique identifier of the transaction rule.
@return TransactionRulesApiUpdateTransactionRuleInput
*/
func (a *TransactionRulesApi) UpdateTransactionRuleInput(transactionRuleId string) TransactionRulesApiUpdateTransactionRuleInput {
	return TransactionRulesApiUpdateTransactionRuleInput{
		transactionRuleId: transactionRuleId,
	}
}

/*
UpdateTransactionRule Update a transaction rule

Updates a transaction rule. 

* To update only the status of a transaction rule, send only the `status` parameter. All other parameters not provided in the request are left unchanged.

* When updating any other parameter, you need to send all existing resource parameters. If you omit a parameter in the request, that parameter is removed from the resource.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransactionRulesApiUpdateTransactionRuleInput - Request parameters, see UpdateTransactionRuleInput
@return TransactionRule, *http.Response, error
*/
func (a *TransactionRulesApi) UpdateTransactionRule(ctx context.Context, r TransactionRulesApiUpdateTransactionRuleInput) (TransactionRule, *http.Response, error) {
    res := &TransactionRule{}
	path := "/transactionRules/{transactionRuleId}"
    path = strings.Replace(path, "{"+"transactionRuleId"+"}", url.PathEscape(common.ParameterValueToString(r.transactionRuleId, "transactionRuleId")), -1)
    queryParams := url.Values{}
    headerParams := make(map[string]string)
    httpRes, err := common.SendAPIRequest(
        ctx,
        a.Client,
        r.transactionRuleInfo,
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

