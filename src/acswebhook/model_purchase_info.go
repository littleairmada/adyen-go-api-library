/*
Authentication webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package acswebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the PurchaseInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PurchaseInfo{}

// PurchaseInfo struct for PurchaseInfo
type PurchaseInfo struct {
	// Date of the purchase.
	Date string `json:"date"`
	// Name of the merchant.
	MerchantName   string `json:"merchantName"`
	OriginalAmount Amount `json:"originalAmount"`
}

// NewPurchaseInfo instantiates a new PurchaseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseInfo(date string, merchantName string, originalAmount Amount) *PurchaseInfo {
	this := PurchaseInfo{}
	this.Date = date
	this.MerchantName = merchantName
	this.OriginalAmount = originalAmount
	return &this
}

// NewPurchaseInfoWithDefaults instantiates a new PurchaseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseInfoWithDefaults() *PurchaseInfo {
	this := PurchaseInfo{}
	return &this
}

// GetDate returns the Date field value
func (o *PurchaseInfo) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *PurchaseInfo) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *PurchaseInfo) SetDate(v string) {
	o.Date = v
}

// GetMerchantName returns the MerchantName field value
func (o *PurchaseInfo) GetMerchantName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value
// and a boolean to check if the value has been set.
func (o *PurchaseInfo) GetMerchantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantName, true
}

// SetMerchantName sets field value
func (o *PurchaseInfo) SetMerchantName(v string) {
	o.MerchantName = v
}

// GetOriginalAmount returns the OriginalAmount field value
func (o *PurchaseInfo) GetOriginalAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.OriginalAmount
}

// GetOriginalAmountOk returns a tuple with the OriginalAmount field value
// and a boolean to check if the value has been set.
func (o *PurchaseInfo) GetOriginalAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalAmount, true
}

// SetOriginalAmount sets field value
func (o *PurchaseInfo) SetOriginalAmount(v Amount) {
	o.OriginalAmount = v
}

func (o PurchaseInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PurchaseInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["merchantName"] = o.MerchantName
	toSerialize["originalAmount"] = o.OriginalAmount
	return toSerialize, nil
}

type NullablePurchaseInfo struct {
	value *PurchaseInfo
	isSet bool
}

func (v NullablePurchaseInfo) Get() *PurchaseInfo {
	return v.value
}

func (v *NullablePurchaseInfo) Set(val *PurchaseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseInfo(val *PurchaseInfo) *NullablePurchaseInfo {
	return &NullablePurchaseInfo{value: val, isSet: true}
}

func (v NullablePurchaseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
