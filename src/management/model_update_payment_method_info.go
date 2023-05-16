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

// checks if the UpdatePaymentMethodInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdatePaymentMethodInfo{}

// UpdatePaymentMethodInfo struct for UpdatePaymentMethodInfo
type UpdatePaymentMethodInfo struct {
	// The list of countries where a payment method is available. By default, all countries supported by the payment method.
	Countries []string `json:"countries,omitempty"`
	// The list of currencies that a payment method supports. By default, all currencies supported by the payment method.
	Currencies []string `json:"currencies,omitempty"`
	// Custom routing flags for acquirer routing.
	CustomRoutingFlags []string `json:"customRoutingFlags,omitempty"`
	// Indicates whether the payment method is enabled (**true**) or disabled (**false**).
	Enabled          *bool             `json:"enabled,omitempty"`
	ShopperStatement *ShopperStatement `json:"shopperStatement,omitempty"`
	// The list of stores for this payment method
	StoreIds []string `json:"storeIds,omitempty"`
}

// NewUpdatePaymentMethodInfo instantiates a new UpdatePaymentMethodInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePaymentMethodInfo() *UpdatePaymentMethodInfo {
	this := UpdatePaymentMethodInfo{}
	return &this
}

// NewUpdatePaymentMethodInfoWithDefaults instantiates a new UpdatePaymentMethodInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePaymentMethodInfoWithDefaults() *UpdatePaymentMethodInfo {
	this := UpdatePaymentMethodInfo{}
	return &this
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetCountries() []string {
	if o == nil || common.IsNil(o.Countries) {
		var ret []string
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetCountriesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasCountries() bool {
	if o != nil && !common.IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []string and assigns it to the Countries field.
func (o *UpdatePaymentMethodInfo) SetCountries(v []string) {
	o.Countries = v
}

// GetCurrencies returns the Currencies field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetCurrencies() []string {
	if o == nil || common.IsNil(o.Currencies) {
		var ret []string
		return ret
	}
	return o.Currencies
}

// GetCurrenciesOk returns a tuple with the Currencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetCurrenciesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Currencies) {
		return nil, false
	}
	return o.Currencies, true
}

// HasCurrencies returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasCurrencies() bool {
	if o != nil && !common.IsNil(o.Currencies) {
		return true
	}

	return false
}

// SetCurrencies gets a reference to the given []string and assigns it to the Currencies field.
func (o *UpdatePaymentMethodInfo) SetCurrencies(v []string) {
	o.Currencies = v
}

// GetCustomRoutingFlags returns the CustomRoutingFlags field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetCustomRoutingFlags() []string {
	if o == nil || common.IsNil(o.CustomRoutingFlags) {
		var ret []string
		return ret
	}
	return o.CustomRoutingFlags
}

// GetCustomRoutingFlagsOk returns a tuple with the CustomRoutingFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetCustomRoutingFlagsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.CustomRoutingFlags) {
		return nil, false
	}
	return o.CustomRoutingFlags, true
}

// HasCustomRoutingFlags returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasCustomRoutingFlags() bool {
	if o != nil && !common.IsNil(o.CustomRoutingFlags) {
		return true
	}

	return false
}

// SetCustomRoutingFlags gets a reference to the given []string and assigns it to the CustomRoutingFlags field.
func (o *UpdatePaymentMethodInfo) SetCustomRoutingFlags(v []string) {
	o.CustomRoutingFlags = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetEnabled() bool {
	if o == nil || common.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasEnabled() bool {
	if o != nil && !common.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdatePaymentMethodInfo) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetShopperStatement returns the ShopperStatement field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetShopperStatement() ShopperStatement {
	if o == nil || common.IsNil(o.ShopperStatement) {
		var ret ShopperStatement
		return ret
	}
	return *o.ShopperStatement
}

// GetShopperStatementOk returns a tuple with the ShopperStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetShopperStatementOk() (*ShopperStatement, bool) {
	if o == nil || common.IsNil(o.ShopperStatement) {
		return nil, false
	}
	return o.ShopperStatement, true
}

// HasShopperStatement returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasShopperStatement() bool {
	if o != nil && !common.IsNil(o.ShopperStatement) {
		return true
	}

	return false
}

// SetShopperStatement gets a reference to the given ShopperStatement and assigns it to the ShopperStatement field.
func (o *UpdatePaymentMethodInfo) SetShopperStatement(v ShopperStatement) {
	o.ShopperStatement = &v
}

// GetStoreIds returns the StoreIds field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetStoreIds() []string {
	if o == nil || common.IsNil(o.StoreIds) {
		var ret []string
		return ret
	}
	return o.StoreIds
}

// GetStoreIdsOk returns a tuple with the StoreIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetStoreIdsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.StoreIds) {
		return nil, false
	}
	return o.StoreIds, true
}

// HasStoreIds returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasStoreIds() bool {
	if o != nil && !common.IsNil(o.StoreIds) {
		return true
	}

	return false
}

// SetStoreIds gets a reference to the given []string and assigns it to the StoreIds field.
func (o *UpdatePaymentMethodInfo) SetStoreIds(v []string) {
	o.StoreIds = v
}

func (o UpdatePaymentMethodInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePaymentMethodInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Countries) {
		toSerialize["countries"] = o.Countries
	}
	if !common.IsNil(o.Currencies) {
		toSerialize["currencies"] = o.Currencies
	}
	if !common.IsNil(o.CustomRoutingFlags) {
		toSerialize["customRoutingFlags"] = o.CustomRoutingFlags
	}
	if !common.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !common.IsNil(o.ShopperStatement) {
		toSerialize["shopperStatement"] = o.ShopperStatement
	}
	if !common.IsNil(o.StoreIds) {
		toSerialize["storeIds"] = o.StoreIds
	}
	return toSerialize, nil
}

type NullableUpdatePaymentMethodInfo struct {
	value *UpdatePaymentMethodInfo
	isSet bool
}

func (v NullableUpdatePaymentMethodInfo) Get() *UpdatePaymentMethodInfo {
	return v.value
}

func (v *NullableUpdatePaymentMethodInfo) Set(val *UpdatePaymentMethodInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePaymentMethodInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePaymentMethodInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePaymentMethodInfo(val *UpdatePaymentMethodInfo) *NullableUpdatePaymentMethodInfo {
	return &NullableUpdatePaymentMethodInfo{value: val, isSet: true}
}

func (v NullableUpdatePaymentMethodInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePaymentMethodInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
