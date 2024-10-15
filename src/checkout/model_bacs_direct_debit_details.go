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

// checks if the BacsDirectDebitDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BacsDirectDebitDetails{}

// BacsDirectDebitDetails struct for BacsDirectDebitDetails
type BacsDirectDebitDetails struct {
	// The bank account number (without separators).
	BankAccountNumber *string `json:"bankAccountNumber,omitempty"`
	// The bank routing number of the account.
	BankLocationId *string `json:"bankLocationId,omitempty"`
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The name of the bank account holder.
	HolderName *string `json:"holderName,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated since Adyen Checkout API v49
	// Use `storedPaymentMethodId` instead.
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// The unique identifier of your user's verified transfer instrument, which you can use to top up their balance accounts.
	TransferInstrumentId *string `json:"transferInstrumentId,omitempty"`
	// **directdebit_GB**
	Type *string `json:"type,omitempty"`
}

// NewBacsDirectDebitDetails instantiates a new BacsDirectDebitDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBacsDirectDebitDetails() *BacsDirectDebitDetails {
	this := BacsDirectDebitDetails{}
	var type_ string = "directdebit_GB"
	this.Type = &type_
	return &this
}

// NewBacsDirectDebitDetailsWithDefaults instantiates a new BacsDirectDebitDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBacsDirectDebitDetailsWithDefaults() *BacsDirectDebitDetails {
	this := BacsDirectDebitDetails{}
	var type_ string = "directdebit_GB"
	this.Type = &type_
	return &this
}

// GetBankAccountNumber returns the BankAccountNumber field value if set, zero value otherwise.
func (o *BacsDirectDebitDetails) GetBankAccountNumber() string {
	if o == nil || common.IsNil(o.BankAccountNumber) {
		var ret string
		return ret
	}
	return *o.BankAccountNumber
}

// GetBankAccountNumberOk returns a tuple with the BankAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BacsDirectDebitDetails) GetBankAccountNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankAccountNumber) {
		return nil, false
	}
	return o.BankAccountNumber, true
}

// HasBankAccountNumber returns a boolean if a field has been set.
func (o *BacsDirectDebitDetails) HasBankAccountNumber() bool {
	if o != nil && !common.IsNil(o.BankAccountNumber) {
		return true
	}

	return false
}

// SetBankAccountNumber gets a reference to the given string and assigns it to the BankAccountNumber field.
func (o *BacsDirectDebitDetails) SetBankAccountNumber(v string) {
	o.BankAccountNumber = &v
}

// GetBankLocationId returns the BankLocationId field value if set, zero value otherwise.
func (o *BacsDirectDebitDetails) GetBankLocationId() string {
	if o == nil || common.IsNil(o.BankLocationId) {
		var ret string
		return ret
	}
	return *o.BankLocationId
}

// GetBankLocationIdOk returns a tuple with the BankLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BacsDirectDebitDetails) GetBankLocationIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankLocationId) {
		return nil, false
	}
	return o.BankLocationId, true
}

// HasBankLocationId returns a boolean if a field has been set.
func (o *BacsDirectDebitDetails) HasBankLocationId() bool {
	if o != nil && !common.IsNil(o.BankLocationId) {
		return true
	}

	return false
}

// SetBankLocationId gets a reference to the given string and assigns it to the BankLocationId field.
func (o *BacsDirectDebitDetails) SetBankLocationId(v string) {
	o.BankLocationId = &v
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *BacsDirectDebitDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BacsDirectDebitDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *BacsDirectDebitDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *BacsDirectDebitDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetHolderName returns the HolderName field value if set, zero value otherwise.
func (o *BacsDirectDebitDetails) GetHolderName() string {
	if o == nil || common.IsNil(o.HolderName) {
		var ret string
		return ret
	}
	return *o.HolderName
}

// GetHolderNameOk returns a tuple with the HolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BacsDirectDebitDetails) GetHolderNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.HolderName) {
		return nil, false
	}
	return o.HolderName, true
}

// HasHolderName returns a boolean if a field has been set.
func (o *BacsDirectDebitDetails) HasHolderName() bool {
	if o != nil && !common.IsNil(o.HolderName) {
		return true
	}

	return false
}

// SetHolderName gets a reference to the given string and assigns it to the HolderName field.
func (o *BacsDirectDebitDetails) SetHolderName(v string) {
	o.HolderName = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *BacsDirectDebitDetails) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *BacsDirectDebitDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *BacsDirectDebitDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *BacsDirectDebitDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *BacsDirectDebitDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BacsDirectDebitDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *BacsDirectDebitDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *BacsDirectDebitDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetTransferInstrumentId returns the TransferInstrumentId field value if set, zero value otherwise.
func (o *BacsDirectDebitDetails) GetTransferInstrumentId() string {
	if o == nil || common.IsNil(o.TransferInstrumentId) {
		var ret string
		return ret
	}
	return *o.TransferInstrumentId
}

// GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BacsDirectDebitDetails) GetTransferInstrumentIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferInstrumentId) {
		return nil, false
	}
	return o.TransferInstrumentId, true
}

// HasTransferInstrumentId returns a boolean if a field has been set.
func (o *BacsDirectDebitDetails) HasTransferInstrumentId() bool {
	if o != nil && !common.IsNil(o.TransferInstrumentId) {
		return true
	}

	return false
}

// SetTransferInstrumentId gets a reference to the given string and assigns it to the TransferInstrumentId field.
func (o *BacsDirectDebitDetails) SetTransferInstrumentId(v string) {
	o.TransferInstrumentId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BacsDirectDebitDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BacsDirectDebitDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BacsDirectDebitDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BacsDirectDebitDetails) SetType(v string) {
	o.Type = &v
}

func (o BacsDirectDebitDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BacsDirectDebitDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BankAccountNumber) {
		toSerialize["bankAccountNumber"] = o.BankAccountNumber
	}
	if !common.IsNil(o.BankLocationId) {
		toSerialize["bankLocationId"] = o.BankLocationId
	}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.HolderName) {
		toSerialize["holderName"] = o.HolderName
	}
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	if !common.IsNil(o.TransferInstrumentId) {
		toSerialize["transferInstrumentId"] = o.TransferInstrumentId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableBacsDirectDebitDetails struct {
	value *BacsDirectDebitDetails
	isSet bool
}

func (v NullableBacsDirectDebitDetails) Get() *BacsDirectDebitDetails {
	return v.value
}

func (v *NullableBacsDirectDebitDetails) Set(val *BacsDirectDebitDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableBacsDirectDebitDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableBacsDirectDebitDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBacsDirectDebitDetails(val *BacsDirectDebitDetails) *NullableBacsDirectDebitDetails {
	return &NullableBacsDirectDebitDetails{value: val, isSet: true}
}

func (v NullableBacsDirectDebitDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBacsDirectDebitDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *BacsDirectDebitDetails) isValidType() bool {
	var allowedEnumValues = []string{"directdebit_GB"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
