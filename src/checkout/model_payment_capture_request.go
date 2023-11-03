/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the PaymentCaptureRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentCaptureRequest{}

// PaymentCaptureRequest struct for PaymentCaptureRequest
type PaymentCaptureRequest struct {
	Amount          Amount           `json:"amount"`
	ApplicationInfo *ApplicationInfo `json:"applicationInfo,omitempty"`
	// Price and product information of the refunded items, required for [partial refunds](https://docs.adyen.com/online-payments/refund#refund-a-payment). > This field is required for partial refunds with 3x 4x Oney, Affirm, Afterpay, Atome, Clearpay, Klarna, Ratepay, Walley, and Zip.
	LineItems []LineItem `json:"lineItems,omitempty"`
	// The merchant account that is used to process the payment.
	MerchantAccount         string                   `json:"merchantAccount"`
	PlatformChargebackLogic *PlatformChargebackLogic `json:"platformChargebackLogic,omitempty"`
	// Your reference for the capture request. Maximum length: 80 characters.
	Reference *string `json:"reference,omitempty"`
	// An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For details, refer to [Providing split information](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information).
	Splits []Split `json:"splits,omitempty"`
	// A List of sub-merchants.
	SubMerchants []SubMerchantInfo `json:"subMerchants,omitempty"`
}

// NewPaymentCaptureRequest instantiates a new PaymentCaptureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentCaptureRequest(amount Amount, merchantAccount string) *PaymentCaptureRequest {
	this := PaymentCaptureRequest{}
	this.Amount = amount
	this.MerchantAccount = merchantAccount
	return &this
}

// NewPaymentCaptureRequestWithDefaults instantiates a new PaymentCaptureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentCaptureRequestWithDefaults() *PaymentCaptureRequest {
	this := PaymentCaptureRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PaymentCaptureRequest) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentCaptureRequest) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentCaptureRequest) SetAmount(v Amount) {
	o.Amount = v
}

// GetApplicationInfo returns the ApplicationInfo field value if set, zero value otherwise.
func (o *PaymentCaptureRequest) GetApplicationInfo() ApplicationInfo {
	if o == nil || common.IsNil(o.ApplicationInfo) {
		var ret ApplicationInfo
		return ret
	}
	return *o.ApplicationInfo
}

// GetApplicationInfoOk returns a tuple with the ApplicationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCaptureRequest) GetApplicationInfoOk() (*ApplicationInfo, bool) {
	if o == nil || common.IsNil(o.ApplicationInfo) {
		return nil, false
	}
	return o.ApplicationInfo, true
}

// HasApplicationInfo returns a boolean if a field has been set.
func (o *PaymentCaptureRequest) HasApplicationInfo() bool {
	if o != nil && !common.IsNil(o.ApplicationInfo) {
		return true
	}

	return false
}

// SetApplicationInfo gets a reference to the given ApplicationInfo and assigns it to the ApplicationInfo field.
func (o *PaymentCaptureRequest) SetApplicationInfo(v ApplicationInfo) {
	o.ApplicationInfo = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *PaymentCaptureRequest) GetLineItems() []LineItem {
	if o == nil || common.IsNil(o.LineItems) {
		var ret []LineItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCaptureRequest) GetLineItemsOk() ([]LineItem, bool) {
	if o == nil || common.IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *PaymentCaptureRequest) HasLineItems() bool {
	if o != nil && !common.IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []LineItem and assigns it to the LineItems field.
func (o *PaymentCaptureRequest) SetLineItems(v []LineItem) {
	o.LineItems = v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *PaymentCaptureRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *PaymentCaptureRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *PaymentCaptureRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetPlatformChargebackLogic returns the PlatformChargebackLogic field value if set, zero value otherwise.
func (o *PaymentCaptureRequest) GetPlatformChargebackLogic() PlatformChargebackLogic {
	if o == nil || common.IsNil(o.PlatformChargebackLogic) {
		var ret PlatformChargebackLogic
		return ret
	}
	return *o.PlatformChargebackLogic
}

// GetPlatformChargebackLogicOk returns a tuple with the PlatformChargebackLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCaptureRequest) GetPlatformChargebackLogicOk() (*PlatformChargebackLogic, bool) {
	if o == nil || common.IsNil(o.PlatformChargebackLogic) {
		return nil, false
	}
	return o.PlatformChargebackLogic, true
}

// HasPlatformChargebackLogic returns a boolean if a field has been set.
func (o *PaymentCaptureRequest) HasPlatformChargebackLogic() bool {
	if o != nil && !common.IsNil(o.PlatformChargebackLogic) {
		return true
	}

	return false
}

// SetPlatformChargebackLogic gets a reference to the given PlatformChargebackLogic and assigns it to the PlatformChargebackLogic field.
func (o *PaymentCaptureRequest) SetPlatformChargebackLogic(v PlatformChargebackLogic) {
	o.PlatformChargebackLogic = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentCaptureRequest) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCaptureRequest) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentCaptureRequest) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentCaptureRequest) SetReference(v string) {
	o.Reference = &v
}

// GetSplits returns the Splits field value if set, zero value otherwise.
func (o *PaymentCaptureRequest) GetSplits() []Split {
	if o == nil || common.IsNil(o.Splits) {
		var ret []Split
		return ret
	}
	return o.Splits
}

// GetSplitsOk returns a tuple with the Splits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCaptureRequest) GetSplitsOk() ([]Split, bool) {
	if o == nil || common.IsNil(o.Splits) {
		return nil, false
	}
	return o.Splits, true
}

// HasSplits returns a boolean if a field has been set.
func (o *PaymentCaptureRequest) HasSplits() bool {
	if o != nil && !common.IsNil(o.Splits) {
		return true
	}

	return false
}

// SetSplits gets a reference to the given []Split and assigns it to the Splits field.
func (o *PaymentCaptureRequest) SetSplits(v []Split) {
	o.Splits = v
}

// GetSubMerchants returns the SubMerchants field value if set, zero value otherwise.
func (o *PaymentCaptureRequest) GetSubMerchants() []SubMerchantInfo {
	if o == nil || common.IsNil(o.SubMerchants) {
		var ret []SubMerchantInfo
		return ret
	}
	return o.SubMerchants
}

// GetSubMerchantsOk returns a tuple with the SubMerchants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCaptureRequest) GetSubMerchantsOk() ([]SubMerchantInfo, bool) {
	if o == nil || common.IsNil(o.SubMerchants) {
		return nil, false
	}
	return o.SubMerchants, true
}

// HasSubMerchants returns a boolean if a field has been set.
func (o *PaymentCaptureRequest) HasSubMerchants() bool {
	if o != nil && !common.IsNil(o.SubMerchants) {
		return true
	}

	return false
}

// SetSubMerchants gets a reference to the given []SubMerchantInfo and assigns it to the SubMerchants field.
func (o *PaymentCaptureRequest) SetSubMerchants(v []SubMerchantInfo) {
	o.SubMerchants = v
}

func (o PaymentCaptureRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentCaptureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.ApplicationInfo) {
		toSerialize["applicationInfo"] = o.ApplicationInfo
	}
	if !common.IsNil(o.LineItems) {
		toSerialize["lineItems"] = o.LineItems
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if !common.IsNil(o.PlatformChargebackLogic) {
		toSerialize["platformChargebackLogic"] = o.PlatformChargebackLogic
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Splits) {
		toSerialize["splits"] = o.Splits
	}
	if !common.IsNil(o.SubMerchants) {
		toSerialize["subMerchants"] = o.SubMerchants
	}
	return toSerialize, nil
}

type NullablePaymentCaptureRequest struct {
	value *PaymentCaptureRequest
	isSet bool
}

func (v NullablePaymentCaptureRequest) Get() *PaymentCaptureRequest {
	return v.value
}

func (v *NullablePaymentCaptureRequest) Set(val *PaymentCaptureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentCaptureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentCaptureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentCaptureRequest(val *PaymentCaptureRequest) *NullablePaymentCaptureRequest {
	return &NullablePaymentCaptureRequest{value: val, isSet: true}
}

func (v NullablePaymentCaptureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentCaptureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
