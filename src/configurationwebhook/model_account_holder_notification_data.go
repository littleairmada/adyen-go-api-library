/*
Configuration webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the AccountHolderNotificationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountHolderNotificationData{}

// AccountHolderNotificationData struct for AccountHolderNotificationData
type AccountHolderNotificationData struct {
	AccountHolder *AccountHolder `json:"accountHolder,omitempty"`
	// The unique identifier of the balance platform.
	BalancePlatform *string `json:"balancePlatform,omitempty"`
}

// NewAccountHolderNotificationData instantiates a new AccountHolderNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountHolderNotificationData() *AccountHolderNotificationData {
	this := AccountHolderNotificationData{}
	return &this
}

// NewAccountHolderNotificationDataWithDefaults instantiates a new AccountHolderNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountHolderNotificationDataWithDefaults() *AccountHolderNotificationData {
	this := AccountHolderNotificationData{}
	return &this
}

// GetAccountHolder returns the AccountHolder field value if set, zero value otherwise.
func (o *AccountHolderNotificationData) GetAccountHolder() AccountHolder {
	if o == nil || common.IsNil(o.AccountHolder) {
		var ret AccountHolder
		return ret
	}
	return *o.AccountHolder
}

// GetAccountHolderOk returns a tuple with the AccountHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderNotificationData) GetAccountHolderOk() (*AccountHolder, bool) {
	if o == nil || common.IsNil(o.AccountHolder) {
		return nil, false
	}
	return o.AccountHolder, true
}

// HasAccountHolder returns a boolean if a field has been set.
func (o *AccountHolderNotificationData) HasAccountHolder() bool {
	if o != nil && !common.IsNil(o.AccountHolder) {
		return true
	}

	return false
}

// SetAccountHolder gets a reference to the given AccountHolder and assigns it to the AccountHolder field.
func (o *AccountHolderNotificationData) SetAccountHolder(v AccountHolder) {
	o.AccountHolder = &v
}

// GetBalancePlatform returns the BalancePlatform field value if set, zero value otherwise.
func (o *AccountHolderNotificationData) GetBalancePlatform() string {
	if o == nil || common.IsNil(o.BalancePlatform) {
		var ret string
		return ret
	}
	return *o.BalancePlatform
}

// GetBalancePlatformOk returns a tuple with the BalancePlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderNotificationData) GetBalancePlatformOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalancePlatform) {
		return nil, false
	}
	return o.BalancePlatform, true
}

// HasBalancePlatform returns a boolean if a field has been set.
func (o *AccountHolderNotificationData) HasBalancePlatform() bool {
	if o != nil && !common.IsNil(o.BalancePlatform) {
		return true
	}

	return false
}

// SetBalancePlatform gets a reference to the given string and assigns it to the BalancePlatform field.
func (o *AccountHolderNotificationData) SetBalancePlatform(v string) {
	o.BalancePlatform = &v
}

func (o AccountHolderNotificationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountHolderNotificationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountHolder) {
		toSerialize["accountHolder"] = o.AccountHolder
	}
	if !common.IsNil(o.BalancePlatform) {
		toSerialize["balancePlatform"] = o.BalancePlatform
	}
	return toSerialize, nil
}

type NullableAccountHolderNotificationData struct {
	value *AccountHolderNotificationData
	isSet bool
}

func (v NullableAccountHolderNotificationData) Get() *AccountHolderNotificationData {
	return v.value
}

func (v *NullableAccountHolderNotificationData) Set(val *AccountHolderNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountHolderNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountHolderNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountHolderNotificationData(val *AccountHolderNotificationData) *NullableAccountHolderNotificationData {
	return &NullableAccountHolderNotificationData{value: val, isSet: true}
}

func (v NullableAccountHolderNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountHolderNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
