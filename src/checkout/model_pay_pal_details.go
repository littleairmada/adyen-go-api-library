/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the PayPalDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PayPalDetails{}

// PayPalDetails struct for PayPalDetails
type PayPalDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The unique ID associated with the order.
	OrderID *string `json:"orderID,omitempty"`
	// IMMEDIATE_PAYMENT_REQUIRED or UNRESTRICTED
	PayeePreferred *string `json:"payeePreferred,omitempty"`
	// The unique ID associated with the payer.
	PayerID *string `json:"payerID,omitempty"`
	// PAYPAL or PAYPAL_CREDIT
	PayerSelected *string `json:"payerSelected,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// The type of flow to initiate.
	Subtype *string `json:"subtype,omitempty"`
	// **paypal**
	Type string `json:"type"`
}

// NewPayPalDetails instantiates a new PayPalDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayPalDetails(type_ string) *PayPalDetails {
	this := PayPalDetails{}
	this.Type = type_
	return &this
}

// NewPayPalDetailsWithDefaults instantiates a new PayPalDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayPalDetailsWithDefaults() *PayPalDetails {
	this := PayPalDetails{}
	var type_ string = "paypal"
	this.Type = type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *PayPalDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPalDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *PayPalDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *PayPalDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetOrderID returns the OrderID field value if set, zero value otherwise.
func (o *PayPalDetails) GetOrderID() string {
	if o == nil || common.IsNil(o.OrderID) {
		var ret string
		return ret
	}
	return *o.OrderID
}

// GetOrderIDOk returns a tuple with the OrderID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPalDetails) GetOrderIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderID) {
		return nil, false
	}
	return o.OrderID, true
}

// HasOrderID returns a boolean if a field has been set.
func (o *PayPalDetails) HasOrderID() bool {
	if o != nil && !common.IsNil(o.OrderID) {
		return true
	}

	return false
}

// SetOrderID gets a reference to the given string and assigns it to the OrderID field.
func (o *PayPalDetails) SetOrderID(v string) {
	o.OrderID = &v
}

// GetPayeePreferred returns the PayeePreferred field value if set, zero value otherwise.
func (o *PayPalDetails) GetPayeePreferred() string {
	if o == nil || common.IsNil(o.PayeePreferred) {
		var ret string
		return ret
	}
	return *o.PayeePreferred
}

// GetPayeePreferredOk returns a tuple with the PayeePreferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPalDetails) GetPayeePreferredOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayeePreferred) {
		return nil, false
	}
	return o.PayeePreferred, true
}

// HasPayeePreferred returns a boolean if a field has been set.
func (o *PayPalDetails) HasPayeePreferred() bool {
	if o != nil && !common.IsNil(o.PayeePreferred) {
		return true
	}

	return false
}

// SetPayeePreferred gets a reference to the given string and assigns it to the PayeePreferred field.
func (o *PayPalDetails) SetPayeePreferred(v string) {
	o.PayeePreferred = &v
}

// GetPayerID returns the PayerID field value if set, zero value otherwise.
func (o *PayPalDetails) GetPayerID() string {
	if o == nil || common.IsNil(o.PayerID) {
		var ret string
		return ret
	}
	return *o.PayerID
}

// GetPayerIDOk returns a tuple with the PayerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPalDetails) GetPayerIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayerID) {
		return nil, false
	}
	return o.PayerID, true
}

// HasPayerID returns a boolean if a field has been set.
func (o *PayPalDetails) HasPayerID() bool {
	if o != nil && !common.IsNil(o.PayerID) {
		return true
	}

	return false
}

// SetPayerID gets a reference to the given string and assigns it to the PayerID field.
func (o *PayPalDetails) SetPayerID(v string) {
	o.PayerID = &v
}

// GetPayerSelected returns the PayerSelected field value if set, zero value otherwise.
func (o *PayPalDetails) GetPayerSelected() string {
	if o == nil || common.IsNil(o.PayerSelected) {
		var ret string
		return ret
	}
	return *o.PayerSelected
}

// GetPayerSelectedOk returns a tuple with the PayerSelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPalDetails) GetPayerSelectedOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayerSelected) {
		return nil, false
	}
	return o.PayerSelected, true
}

// HasPayerSelected returns a boolean if a field has been set.
func (o *PayPalDetails) HasPayerSelected() bool {
	if o != nil && !common.IsNil(o.PayerSelected) {
		return true
	}

	return false
}

// SetPayerSelected gets a reference to the given string and assigns it to the PayerSelected field.
func (o *PayPalDetails) SetPayerSelected(v string) {
	o.PayerSelected = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated
func (o *PayPalDetails) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PayPalDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *PayPalDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated
func (o *PayPalDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *PayPalDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPalDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *PayPalDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *PayPalDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *PayPalDetails) GetSubtype() string {
	if o == nil || common.IsNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPalDetails) GetSubtypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *PayPalDetails) HasSubtype() bool {
	if o != nil && !common.IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *PayPalDetails) SetSubtype(v string) {
	o.Subtype = &v
}

// GetType returns the Type field value
func (o *PayPalDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PayPalDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PayPalDetails) SetType(v string) {
	o.Type = v
}

func (o PayPalDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayPalDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.OrderID) {
		toSerialize["orderID"] = o.OrderID
	}
	if !common.IsNil(o.PayeePreferred) {
		toSerialize["payeePreferred"] = o.PayeePreferred
	}
	if !common.IsNil(o.PayerID) {
		toSerialize["payerID"] = o.PayerID
	}
	if !common.IsNil(o.PayerSelected) {
		toSerialize["payerSelected"] = o.PayerSelected
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

type NullablePayPalDetails struct {
	value *PayPalDetails
	isSet bool
}

func (v NullablePayPalDetails) Get() *PayPalDetails {
	return v.value
}

func (v *NullablePayPalDetails) Set(val *PayPalDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePayPalDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePayPalDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayPalDetails(val *PayPalDetails) *NullablePayPalDetails {
	return &NullablePayPalDetails{value: val, isSet: true}
}

func (v NullablePayPalDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayPalDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PayPalDetails) isValidSubtype() bool {
	var allowedEnumValues = []string{"express", "redirect", "sdk"}
	for _, allowed := range allowedEnumValues {
		if o.GetSubtype() == allowed {
			return true
		}
	}
	return false
}
func (o *PayPalDetails) isValidType() bool {
	var allowedEnumValues = []string{"paypal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
