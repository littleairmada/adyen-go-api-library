/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the AuthenticationResultRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AuthenticationResultRequest{}

// AuthenticationResultRequest struct for AuthenticationResultRequest
type AuthenticationResultRequest struct {
	// The merchant account identifier, with which the authentication was processed.
	MerchantAccount string `json:"merchantAccount"`
	// The pspReference identifier for the transaction.
	PspReference string `json:"pspReference"`
}

// NewAuthenticationResultRequest instantiates a new AuthenticationResultRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationResultRequest(merchantAccount string, pspReference string) *AuthenticationResultRequest {
	this := AuthenticationResultRequest{}
	this.MerchantAccount = merchantAccount
	this.PspReference = pspReference
	return &this
}

// NewAuthenticationResultRequestWithDefaults instantiates a new AuthenticationResultRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationResultRequestWithDefaults() *AuthenticationResultRequest {
	this := AuthenticationResultRequest{}
	return &this
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *AuthenticationResultRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *AuthenticationResultRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *AuthenticationResultRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetPspReference returns the PspReference field value
func (o *AuthenticationResultRequest) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *AuthenticationResultRequest) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *AuthenticationResultRequest) SetPspReference(v string) {
	o.PspReference = v
}

func (o AuthenticationResultRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationResultRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["pspReference"] = o.PspReference
	return toSerialize, nil
}

type NullableAuthenticationResultRequest struct {
	value *AuthenticationResultRequest
	isSet bool
}

func (v NullableAuthenticationResultRequest) Get() *AuthenticationResultRequest {
	return v.value
}

func (v *NullableAuthenticationResultRequest) Set(val *AuthenticationResultRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationResultRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationResultRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationResultRequest(val *AuthenticationResultRequest) *NullableAuthenticationResultRequest {
	return &NullableAuthenticationResultRequest{value: val, isSet: true}
}

func (v NullableAuthenticationResultRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationResultRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
