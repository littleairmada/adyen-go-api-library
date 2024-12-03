/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the ReceiptOptions type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReceiptOptions{}

// ReceiptOptions struct for ReceiptOptions
type ReceiptOptions struct {
	// The receipt logo converted to a Base64-encoded string. The image must be a .bmp file of < 256 KB, dimensions 240 (H) x 384 (W) px.
	Logo *string `json:"logo,omitempty"`
	// Indicates whether a screen appears asking if you want to print the shopper receipt.
	PromptBeforePrinting *bool `json:"promptBeforePrinting,omitempty"`
	// Data to print on the receipt as a QR code. This can include static text and the following variables:  - `${merchantreference}`: the merchant reference of the transaction. - `${pspreference}`: the PSP reference of the transaction.   For example, **http://www.example.com/order/${pspreference}/${merchantreference}**.
	QrCodeData *string `json:"qrCodeData,omitempty"`
}

// NewReceiptOptions instantiates a new ReceiptOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiptOptions() *ReceiptOptions {
	this := ReceiptOptions{}
	return &this
}

// NewReceiptOptionsWithDefaults instantiates a new ReceiptOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiptOptionsWithDefaults() *ReceiptOptions {
	this := ReceiptOptions{}
	return &this
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ReceiptOptions) GetLogo() string {
	if o == nil || common.IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptOptions) GetLogoOk() (*string, bool) {
	if o == nil || common.IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ReceiptOptions) HasLogo() bool {
	if o != nil && !common.IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *ReceiptOptions) SetLogo(v string) {
	o.Logo = &v
}

// GetPromptBeforePrinting returns the PromptBeforePrinting field value if set, zero value otherwise.
func (o *ReceiptOptions) GetPromptBeforePrinting() bool {
	if o == nil || common.IsNil(o.PromptBeforePrinting) {
		var ret bool
		return ret
	}
	return *o.PromptBeforePrinting
}

// GetPromptBeforePrintingOk returns a tuple with the PromptBeforePrinting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptOptions) GetPromptBeforePrintingOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PromptBeforePrinting) {
		return nil, false
	}
	return o.PromptBeforePrinting, true
}

// HasPromptBeforePrinting returns a boolean if a field has been set.
func (o *ReceiptOptions) HasPromptBeforePrinting() bool {
	if o != nil && !common.IsNil(o.PromptBeforePrinting) {
		return true
	}

	return false
}

// SetPromptBeforePrinting gets a reference to the given bool and assigns it to the PromptBeforePrinting field.
func (o *ReceiptOptions) SetPromptBeforePrinting(v bool) {
	o.PromptBeforePrinting = &v
}

// GetQrCodeData returns the QrCodeData field value if set, zero value otherwise.
func (o *ReceiptOptions) GetQrCodeData() string {
	if o == nil || common.IsNil(o.QrCodeData) {
		var ret string
		return ret
	}
	return *o.QrCodeData
}

// GetQrCodeDataOk returns a tuple with the QrCodeData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptOptions) GetQrCodeDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.QrCodeData) {
		return nil, false
	}
	return o.QrCodeData, true
}

// HasQrCodeData returns a boolean if a field has been set.
func (o *ReceiptOptions) HasQrCodeData() bool {
	if o != nil && !common.IsNil(o.QrCodeData) {
		return true
	}

	return false
}

// SetQrCodeData gets a reference to the given string and assigns it to the QrCodeData field.
func (o *ReceiptOptions) SetQrCodeData(v string) {
	o.QrCodeData = &v
}

func (o ReceiptOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceiptOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !common.IsNil(o.PromptBeforePrinting) {
		toSerialize["promptBeforePrinting"] = o.PromptBeforePrinting
	}
	if !common.IsNil(o.QrCodeData) {
		toSerialize["qrCodeData"] = o.QrCodeData
	}
	return toSerialize, nil
}

type NullableReceiptOptions struct {
	value *ReceiptOptions
	isSet bool
}

func (v NullableReceiptOptions) Get() *ReceiptOptions {
	return v.value
}

func (v *NullableReceiptOptions) Set(val *ReceiptOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiptOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiptOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiptOptions(val *ReceiptOptions) *NullableReceiptOptions {
	return &NullableReceiptOptions{value: val, isSet: true}
}

func (v NullableReceiptOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiptOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



