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

// checks if the PayMeInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PayMeInfo{}

// PayMeInfo struct for PayMeInfo
type PayMeInfo struct {
	// Merchant display name
	DisplayName string `json:"displayName"`
	// Merchant logo. Format: Base64-encoded string.
	Logo string `json:"logo"`
	// The email address of merchant support.
	SupportEmail string `json:"supportEmail"`
}

// NewPayMeInfo instantiates a new PayMeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayMeInfo(displayName string, logo string, supportEmail string) *PayMeInfo {
	this := PayMeInfo{}
	this.DisplayName = displayName
	this.Logo = logo
	this.SupportEmail = supportEmail
	return &this
}

// NewPayMeInfoWithDefaults instantiates a new PayMeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayMeInfoWithDefaults() *PayMeInfo {
	this := PayMeInfo{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *PayMeInfo) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PayMeInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *PayMeInfo) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetLogo returns the Logo field value
func (o *PayMeInfo) GetLogo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *PayMeInfo) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value
func (o *PayMeInfo) SetLogo(v string) {
	o.Logo = v
}

// GetSupportEmail returns the SupportEmail field value
func (o *PayMeInfo) GetSupportEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportEmail
}

// GetSupportEmailOk returns a tuple with the SupportEmail field value
// and a boolean to check if the value has been set.
func (o *PayMeInfo) GetSupportEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportEmail, true
}

// SetSupportEmail sets field value
func (o *PayMeInfo) SetSupportEmail(v string) {
	o.SupportEmail = v
}

func (o PayMeInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayMeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["logo"] = o.Logo
	toSerialize["supportEmail"] = o.SupportEmail
	return toSerialize, nil
}

type NullablePayMeInfo struct {
	value *PayMeInfo
	isSet bool
}

func (v NullablePayMeInfo) Get() *PayMeInfo {
	return v.value
}

func (v *NullablePayMeInfo) Set(val *PayMeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePayMeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePayMeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayMeInfo(val *PayMeInfo) *NullablePayMeInfo {
	return &NullablePayMeInfo{value: val, isSet: true}
}

func (v NullablePayMeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayMeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
