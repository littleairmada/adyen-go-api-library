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

// checks if the StarInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StarInfo{}

// StarInfo struct for StarInfo
type StarInfo struct {
	// The type of transactions processed over this payment method.  Allowed values: - **pos** for in-person payments.  - **billpay** for subscription payments, both the initial payment and the later recurring payments. These transactions have `recurringProcessingModel` **Subscription**.  - **ecom** for all other card not present transactions. This includes non-recurring transactions and transactions with `recurringProcessingModel` **CardOnFile** or **UnscheduledCardOnFile**.
	ProcessingType         string                      `json:"processingType"`
	TransactionDescription *TransactionDescriptionInfo `json:"transactionDescription,omitempty"`
}

// NewStarInfo instantiates a new StarInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStarInfo(processingType string) *StarInfo {
	this := StarInfo{}
	this.ProcessingType = processingType
	return &this
}

// NewStarInfoWithDefaults instantiates a new StarInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStarInfoWithDefaults() *StarInfo {
	this := StarInfo{}
	return &this
}

// GetProcessingType returns the ProcessingType field value
func (o *StarInfo) GetProcessingType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessingType
}

// GetProcessingTypeOk returns a tuple with the ProcessingType field value
// and a boolean to check if the value has been set.
func (o *StarInfo) GetProcessingTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessingType, true
}

// SetProcessingType sets field value
func (o *StarInfo) SetProcessingType(v string) {
	o.ProcessingType = v
}

// GetTransactionDescription returns the TransactionDescription field value if set, zero value otherwise.
func (o *StarInfo) GetTransactionDescription() TransactionDescriptionInfo {
	if o == nil || common.IsNil(o.TransactionDescription) {
		var ret TransactionDescriptionInfo
		return ret
	}
	return *o.TransactionDescription
}

// GetTransactionDescriptionOk returns a tuple with the TransactionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StarInfo) GetTransactionDescriptionOk() (*TransactionDescriptionInfo, bool) {
	if o == nil || common.IsNil(o.TransactionDescription) {
		return nil, false
	}
	return o.TransactionDescription, true
}

// HasTransactionDescription returns a boolean if a field has been set.
func (o *StarInfo) HasTransactionDescription() bool {
	if o != nil && !common.IsNil(o.TransactionDescription) {
		return true
	}

	return false
}

// SetTransactionDescription gets a reference to the given TransactionDescriptionInfo and assigns it to the TransactionDescription field.
func (o *StarInfo) SetTransactionDescription(v TransactionDescriptionInfo) {
	o.TransactionDescription = &v
}

func (o StarInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StarInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["processingType"] = o.ProcessingType
	if !common.IsNil(o.TransactionDescription) {
		toSerialize["transactionDescription"] = o.TransactionDescription
	}
	return toSerialize, nil
}

type NullableStarInfo struct {
	value *StarInfo
	isSet bool
}

func (v NullableStarInfo) Get() *StarInfo {
	return v.value
}

func (v *NullableStarInfo) Set(val *StarInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStarInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStarInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStarInfo(val *StarInfo) *NullableStarInfo {
	return &NullableStarInfo{value: val, isSet: true}
}

func (v NullableStarInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStarInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *StarInfo) isValidProcessingType() bool {
	var allowedEnumValues = []string{"billpay", "ecom", "pos"}
	for _, allowed := range allowedEnumValues {
		if o.GetProcessingType() == allowed {
			return true
		}
	}
	return false
}
