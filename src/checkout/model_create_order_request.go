/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the CreateOrderRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateOrderRequest{}

// CreateOrderRequest struct for CreateOrderRequest
type CreateOrderRequest struct {
	Amount Amount `json:"amount"`
	// The date when the order should expire. If not provided, the default expiry duration is 1 day.  [ISO 8601](https://www.w3.org/TR/NOTE-datetime) format: YYYY-MM-DDThh:mm:ss+TZD, for example, **2020-12-18T10:15:30+01:00**.
	ExpiresAt *string `json:"expiresAt,omitempty"`
	// The merchant account identifier, with which you want to process the order.
	MerchantAccount string `json:"merchantAccount"`
	// A custom reference identifying the order.
	Reference string `json:"reference"`
}

// NewCreateOrderRequest instantiates a new CreateOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrderRequest(amount Amount, merchantAccount string, reference string) *CreateOrderRequest {
	this := CreateOrderRequest{}
	this.Amount = amount
	this.MerchantAccount = merchantAccount
	this.Reference = reference
	return &this
}

// NewCreateOrderRequestWithDefaults instantiates a new CreateOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrderRequestWithDefaults() *CreateOrderRequest {
	this := CreateOrderRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *CreateOrderRequest) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateOrderRequest) SetAmount(v Amount) {
	o.Amount = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetExpiresAt() string {
	if o == nil || common.IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetExpiresAtOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasExpiresAt() bool {
	if o != nil && !common.IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *CreateOrderRequest) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *CreateOrderRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *CreateOrderRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetReference returns the Reference field value
func (o *CreateOrderRequest) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *CreateOrderRequest) SetReference(v string) {
	o.Reference = v
}

func (o CreateOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["reference"] = o.Reference
	return toSerialize, nil
}

type NullableCreateOrderRequest struct {
	value *CreateOrderRequest
	isSet bool
}

func (v NullableCreateOrderRequest) Get() *CreateOrderRequest {
	return v.value
}

func (v *NullableCreateOrderRequest) Set(val *CreateOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrderRequest(val *CreateOrderRequest) *NullableCreateOrderRequest {
	return &NullableCreateOrderRequest{value: val, isSet: true}
}

func (v NullableCreateOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
