/*
Adyen Recurring API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the CreatePermitRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreatePermitRequest{}

// CreatePermitRequest struct for CreatePermitRequest
type CreatePermitRequest struct {
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// The permits to create for this recurring contract.
	Permits []Permit `json:"permits"`
	// The recurring contract the new permits will use.
	RecurringDetailReference string `json:"recurringDetailReference"`
	// The shopper's reference to uniquely identify this shopper (e.g. user ID or account ID).
	ShopperReference string `json:"shopperReference"`
}

// NewCreatePermitRequest instantiates a new CreatePermitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePermitRequest(merchantAccount string, permits []Permit, recurringDetailReference string, shopperReference string) *CreatePermitRequest {
	this := CreatePermitRequest{}
	this.MerchantAccount = merchantAccount
	this.Permits = permits
	this.RecurringDetailReference = recurringDetailReference
	this.ShopperReference = shopperReference
	return &this
}

// NewCreatePermitRequestWithDefaults instantiates a new CreatePermitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePermitRequestWithDefaults() *CreatePermitRequest {
	this := CreatePermitRequest{}
	return &this
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *CreatePermitRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *CreatePermitRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *CreatePermitRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetPermits returns the Permits field value
func (o *CreatePermitRequest) GetPermits() []Permit {
	if o == nil {
		var ret []Permit
		return ret
	}

	return o.Permits
}

// GetPermitsOk returns a tuple with the Permits field value
// and a boolean to check if the value has been set.
func (o *CreatePermitRequest) GetPermitsOk() ([]Permit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permits, true
}

// SetPermits sets field value
func (o *CreatePermitRequest) SetPermits(v []Permit) {
	o.Permits = v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value
func (o *CreatePermitRequest) GetRecurringDetailReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value
// and a boolean to check if the value has been set.
func (o *CreatePermitRequest) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecurringDetailReference, true
}

// SetRecurringDetailReference sets field value
func (o *CreatePermitRequest) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = v
}

// GetShopperReference returns the ShopperReference field value
func (o *CreatePermitRequest) GetShopperReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value
// and a boolean to check if the value has been set.
func (o *CreatePermitRequest) GetShopperReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperReference, true
}

// SetShopperReference sets field value
func (o *CreatePermitRequest) SetShopperReference(v string) {
	o.ShopperReference = v
}

func (o CreatePermitRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePermitRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["permits"] = o.Permits
	toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	toSerialize["shopperReference"] = o.ShopperReference
	return toSerialize, nil
}

type NullableCreatePermitRequest struct {
	value *CreatePermitRequest
	isSet bool
}

func (v NullableCreatePermitRequest) Get() *CreatePermitRequest {
	return v.value
}

func (v *NullableCreatePermitRequest) Set(val *CreatePermitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePermitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePermitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePermitRequest(val *CreatePermitRequest) *NullableCreatePermitRequest {
	return &NullableCreatePermitRequest{value: val, isSet: true}
}

func (v NullableCreatePermitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePermitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
