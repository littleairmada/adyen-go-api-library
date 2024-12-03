/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the UploadAndroidAppResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UploadAndroidAppResponse{}

// UploadAndroidAppResponse struct for UploadAndroidAppResponse
type UploadAndroidAppResponse struct {
	// The unique identifier of the uploaded Android app.
	Id *string `json:"id,omitempty"`
}

// NewUploadAndroidAppResponse instantiates a new UploadAndroidAppResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadAndroidAppResponse() *UploadAndroidAppResponse {
	this := UploadAndroidAppResponse{}
	return &this
}

// NewUploadAndroidAppResponseWithDefaults instantiates a new UploadAndroidAppResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadAndroidAppResponseWithDefaults() *UploadAndroidAppResponse {
	this := UploadAndroidAppResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UploadAndroidAppResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadAndroidAppResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UploadAndroidAppResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UploadAndroidAppResponse) SetId(v string) {
	o.Id = &v
}

func (o UploadAndroidAppResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadAndroidAppResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableUploadAndroidAppResponse struct {
	value *UploadAndroidAppResponse
	isSet bool
}

func (v NullableUploadAndroidAppResponse) Get() *UploadAndroidAppResponse {
	return v.value
}

func (v *NullableUploadAndroidAppResponse) Set(val *UploadAndroidAppResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadAndroidAppResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadAndroidAppResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadAndroidAppResponse(val *UploadAndroidAppResponse) *NullableUploadAndroidAppResponse {
	return &NullableUploadAndroidAppResponse{value: val, isSet: true}
}

func (v NullableUploadAndroidAppResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadAndroidAppResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
