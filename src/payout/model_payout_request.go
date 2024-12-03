/*
Adyen Payout API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payout

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the PayoutRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PayoutRequest{}

// PayoutRequest struct for PayoutRequest
type PayoutRequest struct {
	Amount Amount `json:"amount"`
	BillingAddress *Address `json:"billingAddress,omitempty"`
	Card *Card `json:"card,omitempty"`
	// An integer value that is added to the normal fraud score. The value can be either positive or negative.
	FraudOffset *int32 `json:"fraudOffset,omitempty"`
	FundSource *FundSource `json:"fundSource,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	Recurring *Recurring `json:"recurring,omitempty"`
	// The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\"-\"). Maximum length: 80 characters.
	Reference string `json:"reference"`
	// The `recurringDetailReference` you want to use for this payment. The value `LATEST` can be used to select the most recently stored recurring detail.
	SelectedRecurringDetailReference *string `json:"selectedRecurringDetailReference,omitempty"`
	// The shopper's email address. We recommend that you provide this data, as it is used in velocity fraud checks. > For 3D Secure 2 transactions, schemes require `shopperEmail` for all browser-based and mobile implementations.
	ShopperEmail *string `json:"shopperEmail,omitempty"`
	// Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * `Ecommerce` - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * `ContAuth` - Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * `Moto` - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * `POS` - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal.
	ShopperInteraction *string `json:"shopperInteraction,omitempty"`
	ShopperName *Name `json:"shopperName,omitempty"`
	// Required for recurring payments.  Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. > Your reference must not include personally identifiable information (PII), for example name or email address.
	ShopperReference *string `json:"shopperReference,omitempty"`
	// The shopper's telephone number.
	TelephoneNumber *string `json:"telephoneNumber,omitempty"`
}

// NewPayoutRequest instantiates a new PayoutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayoutRequest(amount Amount, merchantAccount string, reference string) *PayoutRequest {
	this := PayoutRequest{}
	this.Amount = amount
	this.MerchantAccount = merchantAccount
	this.Reference = reference
	return &this
}

// NewPayoutRequestWithDefaults instantiates a new PayoutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutRequestWithDefaults() *PayoutRequest {
	this := PayoutRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PayoutRequest) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PayoutRequest) SetAmount(v Amount) {
	o.Amount = v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *PayoutRequest) GetBillingAddress() Address {
	if o == nil || common.IsNil(o.BillingAddress) {
		var ret Address
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetBillingAddressOk() (*Address, bool) {
	if o == nil || common.IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *PayoutRequest) HasBillingAddress() bool {
	if o != nil && !common.IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given Address and assigns it to the BillingAddress field.
func (o *PayoutRequest) SetBillingAddress(v Address) {
	o.BillingAddress = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *PayoutRequest) GetCard() Card {
	if o == nil || common.IsNil(o.Card) {
		var ret Card
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetCardOk() (*Card, bool) {
	if o == nil || common.IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *PayoutRequest) HasCard() bool {
	if o != nil && !common.IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given Card and assigns it to the Card field.
func (o *PayoutRequest) SetCard(v Card) {
	o.Card = &v
}

// GetFraudOffset returns the FraudOffset field value if set, zero value otherwise.
func (o *PayoutRequest) GetFraudOffset() int32 {
	if o == nil || common.IsNil(o.FraudOffset) {
		var ret int32
		return ret
	}
	return *o.FraudOffset
}

// GetFraudOffsetOk returns a tuple with the FraudOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetFraudOffsetOk() (*int32, bool) {
	if o == nil || common.IsNil(o.FraudOffset) {
		return nil, false
	}
	return o.FraudOffset, true
}

// HasFraudOffset returns a boolean if a field has been set.
func (o *PayoutRequest) HasFraudOffset() bool {
	if o != nil && !common.IsNil(o.FraudOffset) {
		return true
	}

	return false
}

// SetFraudOffset gets a reference to the given int32 and assigns it to the FraudOffset field.
func (o *PayoutRequest) SetFraudOffset(v int32) {
	o.FraudOffset = &v
}

// GetFundSource returns the FundSource field value if set, zero value otherwise.
func (o *PayoutRequest) GetFundSource() FundSource {
	if o == nil || common.IsNil(o.FundSource) {
		var ret FundSource
		return ret
	}
	return *o.FundSource
}

// GetFundSourceOk returns a tuple with the FundSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetFundSourceOk() (*FundSource, bool) {
	if o == nil || common.IsNil(o.FundSource) {
		return nil, false
	}
	return o.FundSource, true
}

// HasFundSource returns a boolean if a field has been set.
func (o *PayoutRequest) HasFundSource() bool {
	if o != nil && !common.IsNil(o.FundSource) {
		return true
	}

	return false
}

// SetFundSource gets a reference to the given FundSource and assigns it to the FundSource field.
func (o *PayoutRequest) SetFundSource(v FundSource) {
	o.FundSource = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *PayoutRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *PayoutRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetRecurring returns the Recurring field value if set, zero value otherwise.
func (o *PayoutRequest) GetRecurring() Recurring {
	if o == nil || common.IsNil(o.Recurring) {
		var ret Recurring
		return ret
	}
	return *o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetRecurringOk() (*Recurring, bool) {
	if o == nil || common.IsNil(o.Recurring) {
		return nil, false
	}
	return o.Recurring, true
}

// HasRecurring returns a boolean if a field has been set.
func (o *PayoutRequest) HasRecurring() bool {
	if o != nil && !common.IsNil(o.Recurring) {
		return true
	}

	return false
}

// SetRecurring gets a reference to the given Recurring and assigns it to the Recurring field.
func (o *PayoutRequest) SetRecurring(v Recurring) {
	o.Recurring = &v
}

// GetReference returns the Reference field value
func (o *PayoutRequest) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *PayoutRequest) SetReference(v string) {
	o.Reference = v
}

// GetSelectedRecurringDetailReference returns the SelectedRecurringDetailReference field value if set, zero value otherwise.
func (o *PayoutRequest) GetSelectedRecurringDetailReference() string {
	if o == nil || common.IsNil(o.SelectedRecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.SelectedRecurringDetailReference
}

// GetSelectedRecurringDetailReferenceOk returns a tuple with the SelectedRecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetSelectedRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelectedRecurringDetailReference) {
		return nil, false
	}
	return o.SelectedRecurringDetailReference, true
}

// HasSelectedRecurringDetailReference returns a boolean if a field has been set.
func (o *PayoutRequest) HasSelectedRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.SelectedRecurringDetailReference) {
		return true
	}

	return false
}

// SetSelectedRecurringDetailReference gets a reference to the given string and assigns it to the SelectedRecurringDetailReference field.
func (o *PayoutRequest) SetSelectedRecurringDetailReference(v string) {
	o.SelectedRecurringDetailReference = &v
}

// GetShopperEmail returns the ShopperEmail field value if set, zero value otherwise.
func (o *PayoutRequest) GetShopperEmail() string {
	if o == nil || common.IsNil(o.ShopperEmail) {
		var ret string
		return ret
	}
	return *o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetShopperEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperEmail) {
		return nil, false
	}
	return o.ShopperEmail, true
}

// HasShopperEmail returns a boolean if a field has been set.
func (o *PayoutRequest) HasShopperEmail() bool {
	if o != nil && !common.IsNil(o.ShopperEmail) {
		return true
	}

	return false
}

// SetShopperEmail gets a reference to the given string and assigns it to the ShopperEmail field.
func (o *PayoutRequest) SetShopperEmail(v string) {
	o.ShopperEmail = &v
}

// GetShopperInteraction returns the ShopperInteraction field value if set, zero value otherwise.
func (o *PayoutRequest) GetShopperInteraction() string {
	if o == nil || common.IsNil(o.ShopperInteraction) {
		var ret string
		return ret
	}
	return *o.ShopperInteraction
}

// GetShopperInteractionOk returns a tuple with the ShopperInteraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetShopperInteractionOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperInteraction) {
		return nil, false
	}
	return o.ShopperInteraction, true
}

// HasShopperInteraction returns a boolean if a field has been set.
func (o *PayoutRequest) HasShopperInteraction() bool {
	if o != nil && !common.IsNil(o.ShopperInteraction) {
		return true
	}

	return false
}

// SetShopperInteraction gets a reference to the given string and assigns it to the ShopperInteraction field.
func (o *PayoutRequest) SetShopperInteraction(v string) {
	o.ShopperInteraction = &v
}

// GetShopperName returns the ShopperName field value if set, zero value otherwise.
func (o *PayoutRequest) GetShopperName() Name {
	if o == nil || common.IsNil(o.ShopperName) {
		var ret Name
		return ret
	}
	return *o.ShopperName
}

// GetShopperNameOk returns a tuple with the ShopperName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetShopperNameOk() (*Name, bool) {
	if o == nil || common.IsNil(o.ShopperName) {
		return nil, false
	}
	return o.ShopperName, true
}

// HasShopperName returns a boolean if a field has been set.
func (o *PayoutRequest) HasShopperName() bool {
	if o != nil && !common.IsNil(o.ShopperName) {
		return true
	}

	return false
}

// SetShopperName gets a reference to the given Name and assigns it to the ShopperName field.
func (o *PayoutRequest) SetShopperName(v Name) {
	o.ShopperName = &v
}

// GetShopperReference returns the ShopperReference field value if set, zero value otherwise.
func (o *PayoutRequest) GetShopperReference() string {
	if o == nil || common.IsNil(o.ShopperReference) {
		var ret string
		return ret
	}
	return *o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetShopperReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperReference) {
		return nil, false
	}
	return o.ShopperReference, true
}

// HasShopperReference returns a boolean if a field has been set.
func (o *PayoutRequest) HasShopperReference() bool {
	if o != nil && !common.IsNil(o.ShopperReference) {
		return true
	}

	return false
}

// SetShopperReference gets a reference to the given string and assigns it to the ShopperReference field.
func (o *PayoutRequest) SetShopperReference(v string) {
	o.ShopperReference = &v
}

// GetTelephoneNumber returns the TelephoneNumber field value if set, zero value otherwise.
func (o *PayoutRequest) GetTelephoneNumber() string {
	if o == nil || common.IsNil(o.TelephoneNumber) {
		var ret string
		return ret
	}
	return *o.TelephoneNumber
}

// GetTelephoneNumberOk returns a tuple with the TelephoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetTelephoneNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.TelephoneNumber) {
		return nil, false
	}
	return o.TelephoneNumber, true
}

// HasTelephoneNumber returns a boolean if a field has been set.
func (o *PayoutRequest) HasTelephoneNumber() bool {
	if o != nil && !common.IsNil(o.TelephoneNumber) {
		return true
	}

	return false
}

// SetTelephoneNumber gets a reference to the given string and assigns it to the TelephoneNumber field.
func (o *PayoutRequest) SetTelephoneNumber(v string) {
	o.TelephoneNumber = &v
}

func (o PayoutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayoutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !common.IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	if !common.IsNil(o.FraudOffset) {
		toSerialize["fraudOffset"] = o.FraudOffset
	}
	if !common.IsNil(o.FundSource) {
		toSerialize["fundSource"] = o.FundSource
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if !common.IsNil(o.Recurring) {
		toSerialize["recurring"] = o.Recurring
	}
	toSerialize["reference"] = o.Reference
	if !common.IsNil(o.SelectedRecurringDetailReference) {
		toSerialize["selectedRecurringDetailReference"] = o.SelectedRecurringDetailReference
	}
	if !common.IsNil(o.ShopperEmail) {
		toSerialize["shopperEmail"] = o.ShopperEmail
	}
	if !common.IsNil(o.ShopperInteraction) {
		toSerialize["shopperInteraction"] = o.ShopperInteraction
	}
	if !common.IsNil(o.ShopperName) {
		toSerialize["shopperName"] = o.ShopperName
	}
	if !common.IsNil(o.ShopperReference) {
		toSerialize["shopperReference"] = o.ShopperReference
	}
	if !common.IsNil(o.TelephoneNumber) {
		toSerialize["telephoneNumber"] = o.TelephoneNumber
	}
	return toSerialize, nil
}

type NullablePayoutRequest struct {
	value *PayoutRequest
	isSet bool
}

func (v NullablePayoutRequest) Get() *PayoutRequest {
	return v.value
}

func (v *NullablePayoutRequest) Set(val *PayoutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePayoutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePayoutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayoutRequest(val *PayoutRequest) *NullablePayoutRequest {
	return &NullablePayoutRequest{value: val, isSet: true}
}

func (v NullablePayoutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayoutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *PayoutRequest) isValidShopperInteraction() bool {
    var allowedEnumValues = []string{ "Ecommerce", "ContAuth", "Moto", "POS" }
    for _, allowed := range allowedEnumValues {
        if o.GetShopperInteraction() == allowed {
            return true
        }
    }
    return false
}

