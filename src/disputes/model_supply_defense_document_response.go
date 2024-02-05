/*
Disputes API

API version: 30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputes

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the SupplyDefenseDocumentResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SupplyDefenseDocumentResponse{}

// SupplyDefenseDocumentResponse struct for SupplyDefenseDocumentResponse
type SupplyDefenseDocumentResponse struct {
	DisputeServiceResult DisputeServiceResult `json:"disputeServiceResult"`
}

// NewSupplyDefenseDocumentResponse instantiates a new SupplyDefenseDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplyDefenseDocumentResponse(disputeServiceResult DisputeServiceResult) *SupplyDefenseDocumentResponse {
	this := SupplyDefenseDocumentResponse{}
	this.DisputeServiceResult = disputeServiceResult
	return &this
}

// NewSupplyDefenseDocumentResponseWithDefaults instantiates a new SupplyDefenseDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplyDefenseDocumentResponseWithDefaults() *SupplyDefenseDocumentResponse {
	this := SupplyDefenseDocumentResponse{}
	return &this
}

// GetDisputeServiceResult returns the DisputeServiceResult field value
func (o *SupplyDefenseDocumentResponse) GetDisputeServiceResult() DisputeServiceResult {
	if o == nil {
		var ret DisputeServiceResult
		return ret
	}

	return o.DisputeServiceResult
}

// GetDisputeServiceResultOk returns a tuple with the DisputeServiceResult field value
// and a boolean to check if the value has been set.
func (o *SupplyDefenseDocumentResponse) GetDisputeServiceResultOk() (*DisputeServiceResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisputeServiceResult, true
}

// SetDisputeServiceResult sets field value
func (o *SupplyDefenseDocumentResponse) SetDisputeServiceResult(v DisputeServiceResult) {
	o.DisputeServiceResult = v
}

func (o SupplyDefenseDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupplyDefenseDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["disputeServiceResult"] = o.DisputeServiceResult
	return toSerialize, nil
}

type NullableSupplyDefenseDocumentResponse struct {
	value *SupplyDefenseDocumentResponse
	isSet bool
}

func (v NullableSupplyDefenseDocumentResponse) Get() *SupplyDefenseDocumentResponse {
	return v.value
}

func (v *NullableSupplyDefenseDocumentResponse) Set(val *SupplyDefenseDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplyDefenseDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplyDefenseDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplyDefenseDocumentResponse(val *SupplyDefenseDocumentResponse) *NullableSupplyDefenseDocumentResponse {
	return &NullableSupplyDefenseDocumentResponse{value: val, isSet: true}
}

func (v NullableSupplyDefenseDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplyDefenseDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
