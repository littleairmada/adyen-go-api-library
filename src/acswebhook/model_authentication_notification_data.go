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

// checks if the AuthenticationNotificationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AuthenticationNotificationData{}

// AuthenticationNotificationData struct for AuthenticationNotificationData
type AuthenticationNotificationData struct {
	Authentication AuthenticationInfo `json:"authentication"`
	// The unique identifier of the balance platform.
	BalancePlatform *string `json:"balancePlatform,omitempty"`
	// Unique identifier of the authentication.
	Id string `json:"id"`
	// Unique identifier of the payment instrument that was used for the authentication.
	PaymentInstrumentId string `json:"paymentInstrumentId"`
	Purchase PurchaseInfo `json:"purchase"`
	// Outcome of the authentication. Allowed values: * authenticated * rejected * error
	Status string `json:"status"`
}

// NewAuthenticationNotificationData instantiates a new AuthenticationNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationNotificationData(authentication AuthenticationInfo, id string, paymentInstrumentId string, purchase PurchaseInfo, status string) *AuthenticationNotificationData {
	this := AuthenticationNotificationData{}
	this.Authentication = authentication
	this.Id = id
	this.PaymentInstrumentId = paymentInstrumentId
	this.Purchase = purchase
	this.Status = status
	return &this
}

// NewAuthenticationNotificationDataWithDefaults instantiates a new AuthenticationNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationNotificationDataWithDefaults() *AuthenticationNotificationData {
	this := AuthenticationNotificationData{}
	return &this
}

// GetAuthentication returns the Authentication field value
func (o *AuthenticationNotificationData) GetAuthentication() AuthenticationInfo {
	if o == nil {
		var ret AuthenticationInfo
		return ret
	}

	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *AuthenticationNotificationData) GetAuthenticationOk() (*AuthenticationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value
func (o *AuthenticationNotificationData) SetAuthentication(v AuthenticationInfo) {
	o.Authentication = v
}

// GetBalancePlatform returns the BalancePlatform field value if set, zero value otherwise.
func (o *AuthenticationNotificationData) GetBalancePlatform() string {
	if o == nil || common.IsNil(o.BalancePlatform) {
		var ret string
		return ret
	}
	return *o.BalancePlatform
}

// GetBalancePlatformOk returns a tuple with the BalancePlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationNotificationData) GetBalancePlatformOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalancePlatform) {
		return nil, false
	}
	return o.BalancePlatform, true
}

// HasBalancePlatform returns a boolean if a field has been set.
func (o *AuthenticationNotificationData) HasBalancePlatform() bool {
	if o != nil && !common.IsNil(o.BalancePlatform) {
		return true
	}

	return false
}

// SetBalancePlatform gets a reference to the given string and assigns it to the BalancePlatform field.
func (o *AuthenticationNotificationData) SetBalancePlatform(v string) {
	o.BalancePlatform = &v
}

// GetId returns the Id field value
func (o *AuthenticationNotificationData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuthenticationNotificationData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuthenticationNotificationData) SetId(v string) {
	o.Id = v
}

// GetPaymentInstrumentId returns the PaymentInstrumentId field value
func (o *AuthenticationNotificationData) GetPaymentInstrumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentInstrumentId
}

// GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field value
// and a boolean to check if the value has been set.
func (o *AuthenticationNotificationData) GetPaymentInstrumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentInstrumentId, true
}

// SetPaymentInstrumentId sets field value
func (o *AuthenticationNotificationData) SetPaymentInstrumentId(v string) {
	o.PaymentInstrumentId = v
}

// GetPurchase returns the Purchase field value
func (o *AuthenticationNotificationData) GetPurchase() PurchaseInfo {
	if o == nil {
		var ret PurchaseInfo
		return ret
	}

	return o.Purchase
}

// GetPurchaseOk returns a tuple with the Purchase field value
// and a boolean to check if the value has been set.
func (o *AuthenticationNotificationData) GetPurchaseOk() (*PurchaseInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purchase, true
}

// SetPurchase sets field value
func (o *AuthenticationNotificationData) SetPurchase(v PurchaseInfo) {
	o.Purchase = v
}

// GetStatus returns the Status field value
func (o *AuthenticationNotificationData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AuthenticationNotificationData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AuthenticationNotificationData) SetStatus(v string) {
	o.Status = v
}

func (o AuthenticationNotificationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationNotificationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authentication"] = o.Authentication
	if !common.IsNil(o.BalancePlatform) {
		toSerialize["balancePlatform"] = o.BalancePlatform
	}
	toSerialize["id"] = o.Id
	toSerialize["paymentInstrumentId"] = o.PaymentInstrumentId
	toSerialize["purchase"] = o.Purchase
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableAuthenticationNotificationData struct {
	value *AuthenticationNotificationData
	isSet bool
}

func (v NullableAuthenticationNotificationData) Get() *AuthenticationNotificationData {
	return v.value
}

func (v *NullableAuthenticationNotificationData) Set(val *AuthenticationNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationNotificationData(val *AuthenticationNotificationData) *NullableAuthenticationNotificationData {
	return &NullableAuthenticationNotificationData{value: val, isSet: true}
}

func (v NullableAuthenticationNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *AuthenticationNotificationData) isValidStatus() bool {
    var allowedEnumValues = []string{ "authenticated", "rejected", "error" }
    for _, allowed := range allowedEnumValues {
        if o.GetStatus() == allowed {
            return true
        }
    }
    return false
}

