/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the PayByBankDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PayByBankDetails{}

// PayByBankDetails struct for PayByBankDetails
type PayByBankDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The PayByBank issuer value of the shopper's selected bank.
	Issuer *string `json:"issuer,omitempty"`
	// **paybybank**
	Type string `json:"type"`
}

// NewPayByBankDetails instantiates a new PayByBankDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayByBankDetails(type_ string) *PayByBankDetails {
	this := PayByBankDetails{}
	this.Type = type_
	return &this
}

// NewPayByBankDetailsWithDefaults instantiates a new PayByBankDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayByBankDetailsWithDefaults() *PayByBankDetails {
	this := PayByBankDetails{}
	var type_ string = "paybybank"
	this.Type = type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *PayByBankDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayByBankDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *PayByBankDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *PayByBankDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *PayByBankDetails) GetIssuer() string {
	if o == nil || common.IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayByBankDetails) GetIssuerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *PayByBankDetails) HasIssuer() bool {
	if o != nil && !common.IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *PayByBankDetails) SetIssuer(v string) {
	o.Issuer = &v
}

// GetType returns the Type field value
func (o *PayByBankDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PayByBankDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PayByBankDetails) SetType(v string) {
	o.Type = v
}

func (o PayByBankDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayByBankDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullablePayByBankDetails struct {
	value *PayByBankDetails
	isSet bool
}

func (v NullablePayByBankDetails) Get() *PayByBankDetails {
	return v.value
}

func (v *NullablePayByBankDetails) Set(val *PayByBankDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePayByBankDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePayByBankDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayByBankDetails(val *PayByBankDetails) *NullablePayByBankDetails {
	return &NullablePayByBankDetails{value: val, isSet: true}
}

func (v NullablePayByBankDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayByBankDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PayByBankDetails) isValidType() bool {
	var allowedEnumValues = []string{"paybybank"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
