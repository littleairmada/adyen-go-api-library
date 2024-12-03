/*
Adyen Stored Value API

API version: 46
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storedvalue

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the StoredValueLoadRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoredValueLoadRequest{}

// StoredValueLoadRequest struct for StoredValueLoadRequest
type StoredValueLoadRequest struct {
	Amount Amount `json:"amount"`
	// The type of load you are trying to do, when absent we default to 'Load'
	LoadType *string `json:"loadType,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// The collection that contains the type of the payment method and its specific information if available
	PaymentMethod            map[string]string `json:"paymentMethod"`
	RecurringDetailReference *string           `json:"recurringDetailReference,omitempty"`
	// The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\"-\"). Maximum length: 80 characters.
	Reference string `json:"reference"`
	// Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * `Ecommerce` - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * `ContAuth` - Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * `Moto` - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * `POS` - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal.
	ShopperInteraction *string `json:"shopperInteraction,omitempty"`
	ShopperReference   *string `json:"shopperReference,omitempty"`
	// The physical store, for which this payment is processed.
	Store *string `json:"store,omitempty"`
}

// NewStoredValueLoadRequest instantiates a new StoredValueLoadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoredValueLoadRequest(amount Amount, merchantAccount string, paymentMethod map[string]string, reference string) *StoredValueLoadRequest {
	this := StoredValueLoadRequest{}
	this.Amount = amount
	this.MerchantAccount = merchantAccount
	this.PaymentMethod = paymentMethod
	this.Reference = reference
	return &this
}

// NewStoredValueLoadRequestWithDefaults instantiates a new StoredValueLoadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoredValueLoadRequestWithDefaults() *StoredValueLoadRequest {
	this := StoredValueLoadRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *StoredValueLoadRequest) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *StoredValueLoadRequest) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *StoredValueLoadRequest) SetAmount(v Amount) {
	o.Amount = v
}

// GetLoadType returns the LoadType field value if set, zero value otherwise.
func (o *StoredValueLoadRequest) GetLoadType() string {
	if o == nil || common.IsNil(o.LoadType) {
		var ret string
		return ret
	}
	return *o.LoadType
}

// GetLoadTypeOk returns a tuple with the LoadType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueLoadRequest) GetLoadTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoadType) {
		return nil, false
	}
	return o.LoadType, true
}

// HasLoadType returns a boolean if a field has been set.
func (o *StoredValueLoadRequest) HasLoadType() bool {
	if o != nil && !common.IsNil(o.LoadType) {
		return true
	}

	return false
}

// SetLoadType gets a reference to the given string and assigns it to the LoadType field.
func (o *StoredValueLoadRequest) SetLoadType(v string) {
	o.LoadType = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *StoredValueLoadRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *StoredValueLoadRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *StoredValueLoadRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *StoredValueLoadRequest) GetPaymentMethod() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *StoredValueLoadRequest) GetPaymentMethodOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *StoredValueLoadRequest) SetPaymentMethod(v map[string]string) {
	o.PaymentMethod = v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
func (o *StoredValueLoadRequest) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueLoadRequest) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *StoredValueLoadRequest) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
func (o *StoredValueLoadRequest) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetReference returns the Reference field value
func (o *StoredValueLoadRequest) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *StoredValueLoadRequest) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *StoredValueLoadRequest) SetReference(v string) {
	o.Reference = v
}

// GetShopperInteraction returns the ShopperInteraction field value if set, zero value otherwise.
func (o *StoredValueLoadRequest) GetShopperInteraction() string {
	if o == nil || common.IsNil(o.ShopperInteraction) {
		var ret string
		return ret
	}
	return *o.ShopperInteraction
}

// GetShopperInteractionOk returns a tuple with the ShopperInteraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueLoadRequest) GetShopperInteractionOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperInteraction) {
		return nil, false
	}
	return o.ShopperInteraction, true
}

// HasShopperInteraction returns a boolean if a field has been set.
func (o *StoredValueLoadRequest) HasShopperInteraction() bool {
	if o != nil && !common.IsNil(o.ShopperInteraction) {
		return true
	}

	return false
}

// SetShopperInteraction gets a reference to the given string and assigns it to the ShopperInteraction field.
func (o *StoredValueLoadRequest) SetShopperInteraction(v string) {
	o.ShopperInteraction = &v
}

// GetShopperReference returns the ShopperReference field value if set, zero value otherwise.
func (o *StoredValueLoadRequest) GetShopperReference() string {
	if o == nil || common.IsNil(o.ShopperReference) {
		var ret string
		return ret
	}
	return *o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueLoadRequest) GetShopperReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperReference) {
		return nil, false
	}
	return o.ShopperReference, true
}

// HasShopperReference returns a boolean if a field has been set.
func (o *StoredValueLoadRequest) HasShopperReference() bool {
	if o != nil && !common.IsNil(o.ShopperReference) {
		return true
	}

	return false
}

// SetShopperReference gets a reference to the given string and assigns it to the ShopperReference field.
func (o *StoredValueLoadRequest) SetShopperReference(v string) {
	o.ShopperReference = &v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *StoredValueLoadRequest) GetStore() string {
	if o == nil || common.IsNil(o.Store) {
		var ret string
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueLoadRequest) GetStoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *StoredValueLoadRequest) HasStore() bool {
	if o != nil && !common.IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given string and assigns it to the Store field.
func (o *StoredValueLoadRequest) SetStore(v string) {
	o.Store = &v
}

func (o StoredValueLoadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoredValueLoadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.LoadType) {
		toSerialize["loadType"] = o.LoadType
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["paymentMethod"] = o.PaymentMethod
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	toSerialize["reference"] = o.Reference
	if !common.IsNil(o.ShopperInteraction) {
		toSerialize["shopperInteraction"] = o.ShopperInteraction
	}
	if !common.IsNil(o.ShopperReference) {
		toSerialize["shopperReference"] = o.ShopperReference
	}
	if !common.IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	return toSerialize, nil
}

type NullableStoredValueLoadRequest struct {
	value *StoredValueLoadRequest
	isSet bool
}

func (v NullableStoredValueLoadRequest) Get() *StoredValueLoadRequest {
	return v.value
}

func (v *NullableStoredValueLoadRequest) Set(val *StoredValueLoadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStoredValueLoadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStoredValueLoadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoredValueLoadRequest(val *StoredValueLoadRequest) *NullableStoredValueLoadRequest {
	return &NullableStoredValueLoadRequest{value: val, isSet: true}
}

func (v NullableStoredValueLoadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoredValueLoadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *StoredValueLoadRequest) isValidLoadType() bool {
	var allowedEnumValues = []string{"merchandiseReturn", "load"}
	for _, allowed := range allowedEnumValues {
		if o.GetLoadType() == allowed {
			return true
		}
	}
	return false
}
func (o *StoredValueLoadRequest) isValidShopperInteraction() bool {
	var allowedEnumValues = []string{"Ecommerce", "ContAuth", "Moto", "POS"}
	for _, allowed := range allowedEnumValues {
		if o.GetShopperInteraction() == allowed {
			return true
		}
	}
	return false
}
