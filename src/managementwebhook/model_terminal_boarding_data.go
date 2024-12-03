/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the TerminalBoardingData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalBoardingData{}

// TerminalBoardingData struct for TerminalBoardingData
type TerminalBoardingData struct {
	// The unique identifier of the company account.
	CompanyId string `json:"companyId"`
	// The unique identifier of the merchant account.
	MerchantId *string `json:"merchantId,omitempty"`
	// The unique identifier of the store.
	StoreId *string `json:"storeId,omitempty"`
	// The unique identifier of the terminal.
	UniqueTerminalId string `json:"uniqueTerminalId"`
}

// NewTerminalBoardingData instantiates a new TerminalBoardingData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalBoardingData(companyId string, uniqueTerminalId string) *TerminalBoardingData {
	this := TerminalBoardingData{}
	this.CompanyId = companyId
	this.UniqueTerminalId = uniqueTerminalId
	return &this
}

// NewTerminalBoardingDataWithDefaults instantiates a new TerminalBoardingData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalBoardingDataWithDefaults() *TerminalBoardingData {
	this := TerminalBoardingData{}
	return &this
}

// GetCompanyId returns the CompanyId field value
func (o *TerminalBoardingData) GetCompanyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *TerminalBoardingData) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *TerminalBoardingData) SetCompanyId(v string) {
	o.CompanyId = v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *TerminalBoardingData) GetMerchantId() string {
	if o == nil || common.IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalBoardingData) GetMerchantIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *TerminalBoardingData) HasMerchantId() bool {
	if o != nil && !common.IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *TerminalBoardingData) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *TerminalBoardingData) GetStoreId() string {
	if o == nil || common.IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalBoardingData) GetStoreIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *TerminalBoardingData) HasStoreId() bool {
	if o != nil && !common.IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *TerminalBoardingData) SetStoreId(v string) {
	o.StoreId = &v
}

// GetUniqueTerminalId returns the UniqueTerminalId field value
func (o *TerminalBoardingData) GetUniqueTerminalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniqueTerminalId
}

// GetUniqueTerminalIdOk returns a tuple with the UniqueTerminalId field value
// and a boolean to check if the value has been set.
func (o *TerminalBoardingData) GetUniqueTerminalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniqueTerminalId, true
}

// SetUniqueTerminalId sets field value
func (o *TerminalBoardingData) SetUniqueTerminalId(v string) {
	o.UniqueTerminalId = v
}

func (o TerminalBoardingData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalBoardingData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["companyId"] = o.CompanyId
	if !common.IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !common.IsNil(o.StoreId) {
		toSerialize["storeId"] = o.StoreId
	}
	toSerialize["uniqueTerminalId"] = o.UniqueTerminalId
	return toSerialize, nil
}

type NullableTerminalBoardingData struct {
	value *TerminalBoardingData
	isSet bool
}

func (v NullableTerminalBoardingData) Get() *TerminalBoardingData {
	return v.value
}

func (v *NullableTerminalBoardingData) Set(val *TerminalBoardingData) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalBoardingData) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalBoardingData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalBoardingData(val *TerminalBoardingData) *NullableTerminalBoardingData {
	return &NullableTerminalBoardingData{value: val, isSet: true}
}

func (v NullableTerminalBoardingData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalBoardingData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



