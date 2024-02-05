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

// checks if the KlarnaDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlarnaDetails{}

// KlarnaDetails struct for KlarnaDetails
type KlarnaDetails struct {
	// The address where to send the invoice.
	BillingAddress *string `json:"billingAddress,omitempty"`
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The address where the goods should be delivered.
	DeliveryAddress *string `json:"deliveryAddress,omitempty"`
	// Shopper name, date of birth, phone number, and email address.
	PersonalDetails *string `json:"personalDetails,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// The type of flow to initiate.
	Subtype *string `json:"subtype,omitempty"`
	// **klarna**
	Type string `json:"type"`
}

// NewKlarnaDetails instantiates a new KlarnaDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaDetails(type_ string) *KlarnaDetails {
	this := KlarnaDetails{}
	this.Type = type_
	return &this
}

// NewKlarnaDetailsWithDefaults instantiates a new KlarnaDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaDetailsWithDefaults() *KlarnaDetails {
	this := KlarnaDetails{}
	var type_ string = "klarna"
	this.Type = type_
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *KlarnaDetails) GetBillingAddress() string {
	if o == nil || common.IsNil(o.BillingAddress) {
		var ret string
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaDetails) GetBillingAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *KlarnaDetails) HasBillingAddress() bool {
	if o != nil && !common.IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given string and assigns it to the BillingAddress field.
func (o *KlarnaDetails) SetBillingAddress(v string) {
	o.BillingAddress = &v
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *KlarnaDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *KlarnaDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *KlarnaDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetDeliveryAddress returns the DeliveryAddress field value if set, zero value otherwise.
func (o *KlarnaDetails) GetDeliveryAddress() string {
	if o == nil || common.IsNil(o.DeliveryAddress) {
		var ret string
		return ret
	}
	return *o.DeliveryAddress
}

// GetDeliveryAddressOk returns a tuple with the DeliveryAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaDetails) GetDeliveryAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeliveryAddress) {
		return nil, false
	}
	return o.DeliveryAddress, true
}

// HasDeliveryAddress returns a boolean if a field has been set.
func (o *KlarnaDetails) HasDeliveryAddress() bool {
	if o != nil && !common.IsNil(o.DeliveryAddress) {
		return true
	}

	return false
}

// SetDeliveryAddress gets a reference to the given string and assigns it to the DeliveryAddress field.
func (o *KlarnaDetails) SetDeliveryAddress(v string) {
	o.DeliveryAddress = &v
}

// GetPersonalDetails returns the PersonalDetails field value if set, zero value otherwise.
func (o *KlarnaDetails) GetPersonalDetails() string {
	if o == nil || common.IsNil(o.PersonalDetails) {
		var ret string
		return ret
	}
	return *o.PersonalDetails
}

// GetPersonalDetailsOk returns a tuple with the PersonalDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaDetails) GetPersonalDetailsOk() (*string, bool) {
	if o == nil || common.IsNil(o.PersonalDetails) {
		return nil, false
	}
	return o.PersonalDetails, true
}

// HasPersonalDetails returns a boolean if a field has been set.
func (o *KlarnaDetails) HasPersonalDetails() bool {
	if o != nil && !common.IsNil(o.PersonalDetails) {
		return true
	}

	return false
}

// SetPersonalDetails gets a reference to the given string and assigns it to the PersonalDetails field.
func (o *KlarnaDetails) SetPersonalDetails(v string) {
	o.PersonalDetails = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated
func (o *KlarnaDetails) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *KlarnaDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *KlarnaDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated
func (o *KlarnaDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *KlarnaDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *KlarnaDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *KlarnaDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *KlarnaDetails) GetSubtype() string {
	if o == nil || common.IsNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaDetails) GetSubtypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *KlarnaDetails) HasSubtype() bool {
	if o != nil && !common.IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *KlarnaDetails) SetSubtype(v string) {
	o.Subtype = &v
}

// GetType returns the Type field value
func (o *KlarnaDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *KlarnaDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *KlarnaDetails) SetType(v string) {
	o.Type = v
}

func (o KlarnaDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlarnaDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.DeliveryAddress) {
		toSerialize["deliveryAddress"] = o.DeliveryAddress
	}
	if !common.IsNil(o.PersonalDetails) {
		toSerialize["personalDetails"] = o.PersonalDetails
	}
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	if !common.IsNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableKlarnaDetails struct {
	value *KlarnaDetails
	isSet bool
}

func (v NullableKlarnaDetails) Get() *KlarnaDetails {
	return v.value
}

func (v *NullableKlarnaDetails) Set(val *KlarnaDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaDetails(val *KlarnaDetails) *NullableKlarnaDetails {
	return &NullableKlarnaDetails{value: val, isSet: true}
}

func (v NullableKlarnaDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *KlarnaDetails) isValidType() bool {
	var allowedEnumValues = []string{"klarna", "klarnapayments", "klarnapayments_account", "klarnapayments_b2b", "klarna_paynow", "klarna_account", "klarna_b2b"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
