/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the DokuDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DokuDetails{}

// DokuDetails struct for DokuDetails
type DokuDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The shopper's first name.
	FirstName string `json:"firstName"`
	// The shopper's last name.
	LastName string `json:"lastName"`
	// The shopper's email.
	ShopperEmail string `json:"shopperEmail"`
	// **doku**
	Type string `json:"type"`
}

// NewDokuDetails instantiates a new DokuDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDokuDetails(firstName string, lastName string, shopperEmail string, type_ string) *DokuDetails {
	this := DokuDetails{}
	this.FirstName = firstName
	this.LastName = lastName
	this.ShopperEmail = shopperEmail
	this.Type = type_
	return &this
}

// NewDokuDetailsWithDefaults instantiates a new DokuDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDokuDetailsWithDefaults() *DokuDetails {
	this := DokuDetails{}
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *DokuDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DokuDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *DokuDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *DokuDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetFirstName returns the FirstName field value
func (o *DokuDetails) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *DokuDetails) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *DokuDetails) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *DokuDetails) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *DokuDetails) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *DokuDetails) SetLastName(v string) {
	o.LastName = v
}

// GetShopperEmail returns the ShopperEmail field value
func (o *DokuDetails) GetShopperEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value
// and a boolean to check if the value has been set.
func (o *DokuDetails) GetShopperEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperEmail, true
}

// SetShopperEmail sets field value
func (o *DokuDetails) SetShopperEmail(v string) {
	o.ShopperEmail = v
}

// GetType returns the Type field value
func (o *DokuDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DokuDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DokuDetails) SetType(v string) {
	o.Type = v
}

func (o DokuDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DokuDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	toSerialize["firstName"] = o.FirstName
	toSerialize["lastName"] = o.LastName
	toSerialize["shopperEmail"] = o.ShopperEmail
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableDokuDetails struct {
	value *DokuDetails
	isSet bool
}

func (v NullableDokuDetails) Get() *DokuDetails {
	return v.value
}

func (v *NullableDokuDetails) Set(val *DokuDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDokuDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDokuDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDokuDetails(val *DokuDetails) *NullableDokuDetails {
	return &NullableDokuDetails{value: val, isSet: true}
}

func (v NullableDokuDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDokuDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *DokuDetails) isValidType() bool {
	var allowedEnumValues = []string{"doku_mandiri_va", "doku_cimb_va", "doku_danamon_va", "doku_bni_va", "doku_permata_lite_atm", "doku_bri_va", "doku_bca_va", "doku_alfamart", "doku_indomaret", "doku_wallet", "doku_ovo"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
