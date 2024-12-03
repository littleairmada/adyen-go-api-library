/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the DataReviewConfirmationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DataReviewConfirmationResponse{}

// DataReviewConfirmationResponse struct for DataReviewConfirmationResponse
type DataReviewConfirmationResponse struct {
	// Date when data review was confirmed.
	DataReviewedAt *string `json:"dataReviewedAt,omitempty"`
}

// NewDataReviewConfirmationResponse instantiates a new DataReviewConfirmationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataReviewConfirmationResponse() *DataReviewConfirmationResponse {
	this := DataReviewConfirmationResponse{}
	return &this
}

// NewDataReviewConfirmationResponseWithDefaults instantiates a new DataReviewConfirmationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataReviewConfirmationResponseWithDefaults() *DataReviewConfirmationResponse {
	this := DataReviewConfirmationResponse{}
	return &this
}

// GetDataReviewedAt returns the DataReviewedAt field value if set, zero value otherwise.
func (o *DataReviewConfirmationResponse) GetDataReviewedAt() string {
	if o == nil || common.IsNil(o.DataReviewedAt) {
		var ret string
		return ret
	}
	return *o.DataReviewedAt
}

// GetDataReviewedAtOk returns a tuple with the DataReviewedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReviewConfirmationResponse) GetDataReviewedAtOk() (*string, bool) {
	if o == nil || common.IsNil(o.DataReviewedAt) {
		return nil, false
	}
	return o.DataReviewedAt, true
}

// HasDataReviewedAt returns a boolean if a field has been set.
func (o *DataReviewConfirmationResponse) HasDataReviewedAt() bool {
	if o != nil && !common.IsNil(o.DataReviewedAt) {
		return true
	}

	return false
}

// SetDataReviewedAt gets a reference to the given string and assigns it to the DataReviewedAt field.
func (o *DataReviewConfirmationResponse) SetDataReviewedAt(v string) {
	o.DataReviewedAt = &v
}

func (o DataReviewConfirmationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataReviewConfirmationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DataReviewedAt) {
		toSerialize["dataReviewedAt"] = o.DataReviewedAt
	}
	return toSerialize, nil
}

type NullableDataReviewConfirmationResponse struct {
	value *DataReviewConfirmationResponse
	isSet bool
}

func (v NullableDataReviewConfirmationResponse) Get() *DataReviewConfirmationResponse {
	return v.value
}

func (v *NullableDataReviewConfirmationResponse) Set(val *DataReviewConfirmationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDataReviewConfirmationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDataReviewConfirmationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataReviewConfirmationResponse(val *DataReviewConfirmationResponse) *NullableDataReviewConfirmationResponse {
	return &NullableDataReviewConfirmationResponse{value: val, isSet: true}
}

func (v NullableDataReviewConfirmationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataReviewConfirmationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



