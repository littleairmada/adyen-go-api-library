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

// checks if the ResponseAdditionalDataCard type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ResponseAdditionalDataCard{}

// ResponseAdditionalDataCard struct for ResponseAdditionalDataCard
type ResponseAdditionalDataCard struct {
	// The first six digits of the card number.  This is the [Bank Identification Number (BIN)](https://docs.adyen.com/get-started-with-adyen/payment-glossary#bank-identification-number-bin) for card numbers with a six-digit BIN.  Example: 521234
	CardBin *string `json:"cardBin,omitempty"`
	// The cardholder name passed in the payment request.
	CardHolderName *string `json:"cardHolderName,omitempty"`
	// The bank or the financial institution granting lines of credit through card association branded payment cards. This information can be included when available.
	CardIssuingBank *string `json:"cardIssuingBank,omitempty"`
	// The country where the card was issued.  Example: US
	CardIssuingCountry *string `json:"cardIssuingCountry,omitempty"`
	// The currency in which the card is issued, if this information is available. Provided as the currency code or currency number from the ISO-4217 standard.   Example: USD
	CardIssuingCurrency *string `json:"cardIssuingCurrency,omitempty"`
	// The card payment method used for the transaction.  Example: amex
	CardPaymentMethod *string `json:"cardPaymentMethod,omitempty"`
	// The last four digits of a card number.  > Returned only in case of a card payment.
	CardSummary *string `json:"cardSummary,omitempty"`
	// The first eight digits of the card number. Only returned if the card number is 16 digits or more.  This is the [Bank Identification Number (BIN)](https://docs.adyen.com/get-started-with-adyen/payment-glossary#bank-identification-number-bin) for card numbers with an eight-digit BIN.  Example: 52123423
	IssuerBin *string `json:"issuerBin,omitempty"`
}

// NewResponseAdditionalDataCard instantiates a new ResponseAdditionalDataCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAdditionalDataCard() *ResponseAdditionalDataCard {
	this := ResponseAdditionalDataCard{}
	return &this
}

// NewResponseAdditionalDataCardWithDefaults instantiates a new ResponseAdditionalDataCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAdditionalDataCardWithDefaults() *ResponseAdditionalDataCard {
	this := ResponseAdditionalDataCard{}
	return &this
}

// GetCardBin returns the CardBin field value if set, zero value otherwise.
func (o *ResponseAdditionalDataCard) GetCardBin() string {
	if o == nil || common.IsNil(o.CardBin) {
		var ret string
		return ret
	}
	return *o.CardBin
}

// GetCardBinOk returns a tuple with the CardBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataCard) GetCardBinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardBin) {
		return nil, false
	}
	return o.CardBin, true
}

// HasCardBin returns a boolean if a field has been set.
func (o *ResponseAdditionalDataCard) HasCardBin() bool {
	if o != nil && !common.IsNil(o.CardBin) {
		return true
	}

	return false
}

// SetCardBin gets a reference to the given string and assigns it to the CardBin field.
func (o *ResponseAdditionalDataCard) SetCardBin(v string) {
	o.CardBin = &v
}

// GetCardHolderName returns the CardHolderName field value if set, zero value otherwise.
func (o *ResponseAdditionalDataCard) GetCardHolderName() string {
	if o == nil || common.IsNil(o.CardHolderName) {
		var ret string
		return ret
	}
	return *o.CardHolderName
}

// GetCardHolderNameOk returns a tuple with the CardHolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataCard) GetCardHolderNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardHolderName) {
		return nil, false
	}
	return o.CardHolderName, true
}

// HasCardHolderName returns a boolean if a field has been set.
func (o *ResponseAdditionalDataCard) HasCardHolderName() bool {
	if o != nil && !common.IsNil(o.CardHolderName) {
		return true
	}

	return false
}

// SetCardHolderName gets a reference to the given string and assigns it to the CardHolderName field.
func (o *ResponseAdditionalDataCard) SetCardHolderName(v string) {
	o.CardHolderName = &v
}

// GetCardIssuingBank returns the CardIssuingBank field value if set, zero value otherwise.
func (o *ResponseAdditionalDataCard) GetCardIssuingBank() string {
	if o == nil || common.IsNil(o.CardIssuingBank) {
		var ret string
		return ret
	}
	return *o.CardIssuingBank
}

// GetCardIssuingBankOk returns a tuple with the CardIssuingBank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataCard) GetCardIssuingBankOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardIssuingBank) {
		return nil, false
	}
	return o.CardIssuingBank, true
}

// HasCardIssuingBank returns a boolean if a field has been set.
func (o *ResponseAdditionalDataCard) HasCardIssuingBank() bool {
	if o != nil && !common.IsNil(o.CardIssuingBank) {
		return true
	}

	return false
}

// SetCardIssuingBank gets a reference to the given string and assigns it to the CardIssuingBank field.
func (o *ResponseAdditionalDataCard) SetCardIssuingBank(v string) {
	o.CardIssuingBank = &v
}

// GetCardIssuingCountry returns the CardIssuingCountry field value if set, zero value otherwise.
func (o *ResponseAdditionalDataCard) GetCardIssuingCountry() string {
	if o == nil || common.IsNil(o.CardIssuingCountry) {
		var ret string
		return ret
	}
	return *o.CardIssuingCountry
}

// GetCardIssuingCountryOk returns a tuple with the CardIssuingCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataCard) GetCardIssuingCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardIssuingCountry) {
		return nil, false
	}
	return o.CardIssuingCountry, true
}

// HasCardIssuingCountry returns a boolean if a field has been set.
func (o *ResponseAdditionalDataCard) HasCardIssuingCountry() bool {
	if o != nil && !common.IsNil(o.CardIssuingCountry) {
		return true
	}

	return false
}

// SetCardIssuingCountry gets a reference to the given string and assigns it to the CardIssuingCountry field.
func (o *ResponseAdditionalDataCard) SetCardIssuingCountry(v string) {
	o.CardIssuingCountry = &v
}

// GetCardIssuingCurrency returns the CardIssuingCurrency field value if set, zero value otherwise.
func (o *ResponseAdditionalDataCard) GetCardIssuingCurrency() string {
	if o == nil || common.IsNil(o.CardIssuingCurrency) {
		var ret string
		return ret
	}
	return *o.CardIssuingCurrency
}

// GetCardIssuingCurrencyOk returns a tuple with the CardIssuingCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataCard) GetCardIssuingCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardIssuingCurrency) {
		return nil, false
	}
	return o.CardIssuingCurrency, true
}

// HasCardIssuingCurrency returns a boolean if a field has been set.
func (o *ResponseAdditionalDataCard) HasCardIssuingCurrency() bool {
	if o != nil && !common.IsNil(o.CardIssuingCurrency) {
		return true
	}

	return false
}

// SetCardIssuingCurrency gets a reference to the given string and assigns it to the CardIssuingCurrency field.
func (o *ResponseAdditionalDataCard) SetCardIssuingCurrency(v string) {
	o.CardIssuingCurrency = &v
}

// GetCardPaymentMethod returns the CardPaymentMethod field value if set, zero value otherwise.
func (o *ResponseAdditionalDataCard) GetCardPaymentMethod() string {
	if o == nil || common.IsNil(o.CardPaymentMethod) {
		var ret string
		return ret
	}
	return *o.CardPaymentMethod
}

// GetCardPaymentMethodOk returns a tuple with the CardPaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataCard) GetCardPaymentMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardPaymentMethod) {
		return nil, false
	}
	return o.CardPaymentMethod, true
}

// HasCardPaymentMethod returns a boolean if a field has been set.
func (o *ResponseAdditionalDataCard) HasCardPaymentMethod() bool {
	if o != nil && !common.IsNil(o.CardPaymentMethod) {
		return true
	}

	return false
}

// SetCardPaymentMethod gets a reference to the given string and assigns it to the CardPaymentMethod field.
func (o *ResponseAdditionalDataCard) SetCardPaymentMethod(v string) {
	o.CardPaymentMethod = &v
}

// GetCardSummary returns the CardSummary field value if set, zero value otherwise.
func (o *ResponseAdditionalDataCard) GetCardSummary() string {
	if o == nil || common.IsNil(o.CardSummary) {
		var ret string
		return ret
	}
	return *o.CardSummary
}

// GetCardSummaryOk returns a tuple with the CardSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataCard) GetCardSummaryOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardSummary) {
		return nil, false
	}
	return o.CardSummary, true
}

// HasCardSummary returns a boolean if a field has been set.
func (o *ResponseAdditionalDataCard) HasCardSummary() bool {
	if o != nil && !common.IsNil(o.CardSummary) {
		return true
	}

	return false
}

// SetCardSummary gets a reference to the given string and assigns it to the CardSummary field.
func (o *ResponseAdditionalDataCard) SetCardSummary(v string) {
	o.CardSummary = &v
}

// GetIssuerBin returns the IssuerBin field value if set, zero value otherwise.
func (o *ResponseAdditionalDataCard) GetIssuerBin() string {
	if o == nil || common.IsNil(o.IssuerBin) {
		var ret string
		return ret
	}
	return *o.IssuerBin
}

// GetIssuerBinOk returns a tuple with the IssuerBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataCard) GetIssuerBinOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuerBin) {
		return nil, false
	}
	return o.IssuerBin, true
}

// HasIssuerBin returns a boolean if a field has been set.
func (o *ResponseAdditionalDataCard) HasIssuerBin() bool {
	if o != nil && !common.IsNil(o.IssuerBin) {
		return true
	}

	return false
}

// SetIssuerBin gets a reference to the given string and assigns it to the IssuerBin field.
func (o *ResponseAdditionalDataCard) SetIssuerBin(v string) {
	o.IssuerBin = &v
}

func (o ResponseAdditionalDataCard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAdditionalDataCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CardBin) {
		toSerialize["cardBin"] = o.CardBin
	}
	if !common.IsNil(o.CardHolderName) {
		toSerialize["cardHolderName"] = o.CardHolderName
	}
	if !common.IsNil(o.CardIssuingBank) {
		toSerialize["cardIssuingBank"] = o.CardIssuingBank
	}
	if !common.IsNil(o.CardIssuingCountry) {
		toSerialize["cardIssuingCountry"] = o.CardIssuingCountry
	}
	if !common.IsNil(o.CardIssuingCurrency) {
		toSerialize["cardIssuingCurrency"] = o.CardIssuingCurrency
	}
	if !common.IsNil(o.CardPaymentMethod) {
		toSerialize["cardPaymentMethod"] = o.CardPaymentMethod
	}
	if !common.IsNil(o.CardSummary) {
		toSerialize["cardSummary"] = o.CardSummary
	}
	if !common.IsNil(o.IssuerBin) {
		toSerialize["issuerBin"] = o.IssuerBin
	}
	return toSerialize, nil
}

type NullableResponseAdditionalDataCard struct {
	value *ResponseAdditionalDataCard
	isSet bool
}

func (v NullableResponseAdditionalDataCard) Get() *ResponseAdditionalDataCard {
	return v.value
}

func (v *NullableResponseAdditionalDataCard) Set(val *ResponseAdditionalDataCard) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAdditionalDataCard) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAdditionalDataCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAdditionalDataCard(val *ResponseAdditionalDataCard) *NullableResponseAdditionalDataCard {
	return &NullableResponseAdditionalDataCard{value: val, isSet: true}
}

func (v NullableResponseAdditionalDataCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAdditionalDataCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
