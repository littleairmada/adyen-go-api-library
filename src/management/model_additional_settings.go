/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the AdditionalSettings type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalSettings{}

// AdditionalSettings struct for AdditionalSettings
type AdditionalSettings struct {
	// Object containing list of event codes for which the notification will be sent. 
	IncludeEventCodes []string `json:"includeEventCodes,omitempty"`
	// Object containing boolean key-value pairs. The key can be any [standard webhook additional setting](https://docs.adyen.com/development-resources/webhooks/additional-settings), and the value indicates if the setting is enabled. For example, `captureDelayHours`: **true** means the standard notifications you get will contain the number of hours remaining until the payment will be captured.
	Properties *map[string]bool `json:"properties,omitempty"`
}

// NewAdditionalSettings instantiates a new AdditionalSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalSettings() *AdditionalSettings {
	this := AdditionalSettings{}
	return &this
}

// NewAdditionalSettingsWithDefaults instantiates a new AdditionalSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalSettingsWithDefaults() *AdditionalSettings {
	this := AdditionalSettings{}
	return &this
}

// GetIncludeEventCodes returns the IncludeEventCodes field value if set, zero value otherwise.
func (o *AdditionalSettings) GetIncludeEventCodes() []string {
	if o == nil || common.IsNil(o.IncludeEventCodes) {
		var ret []string
		return ret
	}
	return o.IncludeEventCodes
}

// GetIncludeEventCodesOk returns a tuple with the IncludeEventCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSettings) GetIncludeEventCodesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.IncludeEventCodes) {
		return nil, false
	}
	return o.IncludeEventCodes, true
}

// HasIncludeEventCodes returns a boolean if a field has been set.
func (o *AdditionalSettings) HasIncludeEventCodes() bool {
	if o != nil && !common.IsNil(o.IncludeEventCodes) {
		return true
	}

	return false
}

// SetIncludeEventCodes gets a reference to the given []string and assigns it to the IncludeEventCodes field.
func (o *AdditionalSettings) SetIncludeEventCodes(v []string) {
	o.IncludeEventCodes = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *AdditionalSettings) GetProperties() map[string]bool {
	if o == nil || common.IsNil(o.Properties) {
		var ret map[string]bool
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSettings) GetPropertiesOk() (*map[string]bool, bool) {
	if o == nil || common.IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *AdditionalSettings) HasProperties() bool {
	if o != nil && !common.IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]bool and assigns it to the Properties field.
func (o *AdditionalSettings) SetProperties(v map[string]bool) {
	o.Properties = &v
}

func (o AdditionalSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.IncludeEventCodes) {
		toSerialize["includeEventCodes"] = o.IncludeEventCodes
	}
	if !common.IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableAdditionalSettings struct {
	value *AdditionalSettings
	isSet bool
}

func (v NullableAdditionalSettings) Get() *AdditionalSettings {
	return v.value
}

func (v *NullableAdditionalSettings) Set(val *AdditionalSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalSettings(val *AdditionalSettings) *NullableAdditionalSettings {
	return &NullableAdditionalSettings{value: val, isSet: true}
}

func (v NullableAdditionalSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



