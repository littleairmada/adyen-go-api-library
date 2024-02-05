/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the UpdatePayoutSettingsRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdatePayoutSettingsRequest{}

// UpdatePayoutSettingsRequest struct for UpdatePayoutSettingsRequest
type UpdatePayoutSettingsRequest struct {
	// Indicates if payouts to this bank account are enabled. Default: **true**.  To receive payouts into this bank account, both `enabled` and `allowed` must be **true**.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdatePayoutSettingsRequest instantiates a new UpdatePayoutSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePayoutSettingsRequest() *UpdatePayoutSettingsRequest {
	this := UpdatePayoutSettingsRequest{}
	return &this
}

// NewUpdatePayoutSettingsRequestWithDefaults instantiates a new UpdatePayoutSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePayoutSettingsRequestWithDefaults() *UpdatePayoutSettingsRequest {
	this := UpdatePayoutSettingsRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdatePayoutSettingsRequest) GetEnabled() bool {
	if o == nil || common.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePayoutSettingsRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdatePayoutSettingsRequest) HasEnabled() bool {
	if o != nil && !common.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdatePayoutSettingsRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdatePayoutSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePayoutSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdatePayoutSettingsRequest struct {
	value *UpdatePayoutSettingsRequest
	isSet bool
}

func (v NullableUpdatePayoutSettingsRequest) Get() *UpdatePayoutSettingsRequest {
	return v.value
}

func (v *NullableUpdatePayoutSettingsRequest) Set(val *UpdatePayoutSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePayoutSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePayoutSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePayoutSettingsRequest(val *UpdatePayoutSettingsRequest) *NullableUpdatePayoutSettingsRequest {
	return &NullableUpdatePayoutSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdatePayoutSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePayoutSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
