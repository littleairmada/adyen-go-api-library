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

// checks if the VippsInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VippsInfo{}

// VippsInfo struct for VippsInfo
type VippsInfo struct {
	// Vipps logo. Format: Base64-encoded string.
	Logo string `json:"logo"`
	// Vipps subscription cancel url (required in case of [recurring payments](https://docs.adyen.com/online-payments/tokenization))
	SubscriptionCancelUrl *string `json:"subscriptionCancelUrl,omitempty"`
}

// NewVippsInfo instantiates a new VippsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVippsInfo(logo string) *VippsInfo {
	this := VippsInfo{}
	this.Logo = logo
	return &this
}

// NewVippsInfoWithDefaults instantiates a new VippsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVippsInfoWithDefaults() *VippsInfo {
	this := VippsInfo{}
	return &this
}

// GetLogo returns the Logo field value
func (o *VippsInfo) GetLogo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *VippsInfo) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value
func (o *VippsInfo) SetLogo(v string) {
	o.Logo = v
}

// GetSubscriptionCancelUrl returns the SubscriptionCancelUrl field value if set, zero value otherwise.
func (o *VippsInfo) GetSubscriptionCancelUrl() string {
	if o == nil || common.IsNil(o.SubscriptionCancelUrl) {
		var ret string
		return ret
	}
	return *o.SubscriptionCancelUrl
}

// GetSubscriptionCancelUrlOk returns a tuple with the SubscriptionCancelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VippsInfo) GetSubscriptionCancelUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubscriptionCancelUrl) {
		return nil, false
	}
	return o.SubscriptionCancelUrl, true
}

// HasSubscriptionCancelUrl returns a boolean if a field has been set.
func (o *VippsInfo) HasSubscriptionCancelUrl() bool {
	if o != nil && !common.IsNil(o.SubscriptionCancelUrl) {
		return true
	}

	return false
}

// SetSubscriptionCancelUrl gets a reference to the given string and assigns it to the SubscriptionCancelUrl field.
func (o *VippsInfo) SetSubscriptionCancelUrl(v string) {
	o.SubscriptionCancelUrl = &v
}

func (o VippsInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VippsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["logo"] = o.Logo
	if !common.IsNil(o.SubscriptionCancelUrl) {
		toSerialize["subscriptionCancelUrl"] = o.SubscriptionCancelUrl
	}
	return toSerialize, nil
}

type NullableVippsInfo struct {
	value *VippsInfo
	isSet bool
}

func (v NullableVippsInfo) Get() *VippsInfo {
	return v.value
}

func (v *NullableVippsInfo) Set(val *VippsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVippsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVippsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVippsInfo(val *VippsInfo) *NullableVippsInfo {
	return &NullableVippsInfo{value: val, isSet: true}
}

func (v NullableVippsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVippsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
