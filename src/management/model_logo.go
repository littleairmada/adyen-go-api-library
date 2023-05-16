/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the Logo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Logo{}

// Logo struct for Logo
type Logo struct {
	// The image file, converted to a Base64-encoded string, of the logo to be shown on the terminal.
	Data *string `json:"data,omitempty"`
}

// NewLogo instantiates a new Logo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogo() *Logo {
	this := Logo{}
	return &this
}

// NewLogoWithDefaults instantiates a new Logo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogoWithDefaults() *Logo {
	this := Logo{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Logo) GetData() string {
	if o == nil || common.IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Logo) GetDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Logo) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *Logo) SetData(v string) {
	o.Data = &v
}

func (o Logo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Logo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLogo struct {
	value *Logo
	isSet bool
}

func (v NullableLogo) Get() *Logo {
	return v.value
}

func (v *NullableLogo) Set(val *Logo) {
	v.value = val
	v.isSet = true
}

func (v NullableLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogo(val *Logo) *NullableLogo {
	return &NullableLogo{value: val, isSet: true}
}

func (v NullableLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
