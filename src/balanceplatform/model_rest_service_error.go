/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the RestServiceError type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RestServiceError{}

// RestServiceError struct for RestServiceError
type RestServiceError struct {
	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail"`
	// A code that identifies the problem type.
	ErrorCode string `json:"errorCode"`
	// A unique URI that identifies the specific occurrence of the problem.
	Instance *string `json:"instance,omitempty"`
	// Detailed explanation of each validation error, when applicable.
	InvalidFields []InvalidField `json:"invalidFields,omitempty"`
	// A unique reference for the request, essentially the same as `pspReference`.
	RequestId *string `json:"requestId,omitempty"`
	Response map[string]interface{} `json:"response,omitempty"`
	// The HTTP status code.
	Status int32 `json:"status"`
	// A short, human-readable summary of the problem type.
	Title string `json:"title"`
	// A URI that identifies the problem type, pointing to human-readable documentation on this problem type.
	Type string `json:"type"`
}

// NewRestServiceError instantiates a new RestServiceError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestServiceError(detail string, errorCode string, status int32, title string, type_ string) *RestServiceError {
	this := RestServiceError{}
	this.Detail = detail
	this.ErrorCode = errorCode
	this.Status = status
	this.Title = title
	this.Type = type_
	return &this
}

// NewRestServiceErrorWithDefaults instantiates a new RestServiceError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestServiceErrorWithDefaults() *RestServiceError {
	this := RestServiceError{}
	return &this
}

// GetDetail returns the Detail field value
func (o *RestServiceError) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *RestServiceError) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *RestServiceError) SetDetail(v string) {
	o.Detail = v
}

// GetErrorCode returns the ErrorCode field value
func (o *RestServiceError) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *RestServiceError) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *RestServiceError) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *RestServiceError) GetInstance() string {
	if o == nil || common.IsNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestServiceError) GetInstanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *RestServiceError) HasInstance() bool {
	if o != nil && !common.IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *RestServiceError) SetInstance(v string) {
	o.Instance = &v
}

// GetInvalidFields returns the InvalidFields field value if set, zero value otherwise.
func (o *RestServiceError) GetInvalidFields() []InvalidField {
	if o == nil || common.IsNil(o.InvalidFields) {
		var ret []InvalidField
		return ret
	}
	return o.InvalidFields
}

// GetInvalidFieldsOk returns a tuple with the InvalidFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestServiceError) GetInvalidFieldsOk() ([]InvalidField, bool) {
	if o == nil || common.IsNil(o.InvalidFields) {
		return nil, false
	}
	return o.InvalidFields, true
}

// HasInvalidFields returns a boolean if a field has been set.
func (o *RestServiceError) HasInvalidFields() bool {
	if o != nil && !common.IsNil(o.InvalidFields) {
		return true
	}

	return false
}

// SetInvalidFields gets a reference to the given []InvalidField and assigns it to the InvalidFields field.
func (o *RestServiceError) SetInvalidFields(v []InvalidField) {
	o.InvalidFields = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *RestServiceError) GetRequestId() string {
	if o == nil || common.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestServiceError) GetRequestIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *RestServiceError) HasRequestId() bool {
	if o != nil && !common.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *RestServiceError) SetRequestId(v string) {
	o.RequestId = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *RestServiceError) GetResponse() map[string]interface{} {
	if o == nil || common.IsNil(o.Response) {
		var ret map[string]interface{}
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestServiceError) GetResponseOk() (map[string]interface{}, bool) {
	if o == nil || common.IsNil(o.Response) {
		return map[string]interface{}{}, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *RestServiceError) HasResponse() bool {
	if o != nil && !common.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given map[string]interface{} and assigns it to the Response field.
func (o *RestServiceError) SetResponse(v map[string]interface{}) {
	o.Response = v
}

// GetStatus returns the Status field value
func (o *RestServiceError) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RestServiceError) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RestServiceError) SetStatus(v int32) {
	o.Status = v
}

// GetTitle returns the Title field value
func (o *RestServiceError) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *RestServiceError) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *RestServiceError) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value
func (o *RestServiceError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RestServiceError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RestServiceError) SetType(v string) {
	o.Type = v
}

func (o RestServiceError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestServiceError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["detail"] = o.Detail
	toSerialize["errorCode"] = o.ErrorCode
	if !common.IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !common.IsNil(o.InvalidFields) {
		toSerialize["invalidFields"] = o.InvalidFields
	}
	if !common.IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !common.IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableRestServiceError struct {
	value *RestServiceError
	isSet bool
}

func (v NullableRestServiceError) Get() *RestServiceError {
	return v.value
}

func (v *NullableRestServiceError) Set(val *RestServiceError) {
	v.value = val
	v.isSet = true
}

func (v NullableRestServiceError) IsSet() bool {
	return v.isSet
}

func (v *NullableRestServiceError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestServiceError(val *RestServiceError) *NullableRestServiceError {
	return &NullableRestServiceError{value: val, isSet: true}
}

func (v NullableRestServiceError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestServiceError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



