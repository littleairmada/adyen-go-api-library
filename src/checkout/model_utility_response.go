/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the UtilityResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UtilityResponse{}

// UtilityResponse struct for UtilityResponse
type UtilityResponse struct {
	// The list of origin keys for all requested domains. For each list item, the key is the domain and the value is the origin key.
	OriginKeys *map[string]string `json:"originKeys,omitempty"`
}

// NewUtilityResponse instantiates a new UtilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilityResponse() *UtilityResponse {
	this := UtilityResponse{}
	return &this
}

// NewUtilityResponseWithDefaults instantiates a new UtilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilityResponseWithDefaults() *UtilityResponse {
	this := UtilityResponse{}
	return &this
}

// GetOriginKeys returns the OriginKeys field value if set, zero value otherwise.
func (o *UtilityResponse) GetOriginKeys() map[string]string {
	if o == nil || common.IsNil(o.OriginKeys) {
		var ret map[string]string
		return ret
	}
	return *o.OriginKeys
}

// GetOriginKeysOk returns a tuple with the OriginKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilityResponse) GetOriginKeysOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.OriginKeys) {
		return nil, false
	}
	return o.OriginKeys, true
}

// HasOriginKeys returns a boolean if a field has been set.
func (o *UtilityResponse) HasOriginKeys() bool {
	if o != nil && !common.IsNil(o.OriginKeys) {
		return true
	}

	return false
}

// SetOriginKeys gets a reference to the given map[string]string and assigns it to the OriginKeys field.
func (o *UtilityResponse) SetOriginKeys(v map[string]string) {
	o.OriginKeys = &v
}

func (o UtilityResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OriginKeys) {
		toSerialize["originKeys"] = o.OriginKeys
	}
	return toSerialize, nil
}

type NullableUtilityResponse struct {
	value *UtilityResponse
	isSet bool
}

func (v NullableUtilityResponse) Get() *UtilityResponse {
	return v.value
}

func (v *NullableUtilityResponse) Set(val *UtilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilityResponse(val *UtilityResponse) *NullableUtilityResponse {
	return &NullableUtilityResponse{value: val, isSet: true}
}

func (v NullableUtilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
