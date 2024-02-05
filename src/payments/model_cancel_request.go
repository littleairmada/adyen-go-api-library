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

// checks if the CancelRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelRequest{}

// CancelRequest struct for CancelRequest
type CancelRequest struct {
	// This field contains additional data, which may be required for a particular modification request.  The additionalData object consists of entries, each of which includes the key and value.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	// The merchant account that is used to process the payment.
	MerchantAccount string            `json:"merchantAccount"`
	MpiData         *ThreeDSecureData `json:"mpiData,omitempty"`
	// The original merchant reference to cancel.
	OriginalMerchantReference *string `json:"originalMerchantReference,omitempty"`
	// The original pspReference of the payment to modify. This reference is returned in: * authorisation response * authorisation notification
	OriginalReference       string                   `json:"originalReference"`
	PlatformChargebackLogic *PlatformChargebackLogic `json:"platformChargebackLogic,omitempty"`
	// Your reference for the payment modification. This reference is visible in Customer Area and in reports. Maximum length: 80 characters.
	Reference *string `json:"reference,omitempty"`
	// An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For details, refer to [Providing split information](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information).
	Splits []Split `json:"splits,omitempty"`
	// The transaction reference provided by the PED. For point-of-sale integrations only.
	TenderReference *string `json:"tenderReference,omitempty"`
	// Unique terminal ID for the PED that originally processed the request. For point-of-sale integrations only.
	UniqueTerminalId *string `json:"uniqueTerminalId,omitempty"`
}

// NewCancelRequest instantiates a new CancelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelRequest(merchantAccount string, originalReference string) *CancelRequest {
	this := CancelRequest{}
	this.MerchantAccount = merchantAccount
	this.OriginalReference = originalReference
	return &this
}

// NewCancelRequestWithDefaults instantiates a new CancelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelRequestWithDefaults() *CancelRequest {
	this := CancelRequest{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *CancelRequest) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelRequest) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *CancelRequest) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *CancelRequest) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *CancelRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *CancelRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *CancelRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetMpiData returns the MpiData field value if set, zero value otherwise.
func (o *CancelRequest) GetMpiData() ThreeDSecureData {
	if o == nil || common.IsNil(o.MpiData) {
		var ret ThreeDSecureData
		return ret
	}
	return *o.MpiData
}

// GetMpiDataOk returns a tuple with the MpiData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelRequest) GetMpiDataOk() (*ThreeDSecureData, bool) {
	if o == nil || common.IsNil(o.MpiData) {
		return nil, false
	}
	return o.MpiData, true
}

// HasMpiData returns a boolean if a field has been set.
func (o *CancelRequest) HasMpiData() bool {
	if o != nil && !common.IsNil(o.MpiData) {
		return true
	}

	return false
}

// SetMpiData gets a reference to the given ThreeDSecureData and assigns it to the MpiData field.
func (o *CancelRequest) SetMpiData(v ThreeDSecureData) {
	o.MpiData = &v
}

// GetOriginalMerchantReference returns the OriginalMerchantReference field value if set, zero value otherwise.
func (o *CancelRequest) GetOriginalMerchantReference() string {
	if o == nil || common.IsNil(o.OriginalMerchantReference) {
		var ret string
		return ret
	}
	return *o.OriginalMerchantReference
}

// GetOriginalMerchantReferenceOk returns a tuple with the OriginalMerchantReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelRequest) GetOriginalMerchantReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.OriginalMerchantReference) {
		return nil, false
	}
	return o.OriginalMerchantReference, true
}

// HasOriginalMerchantReference returns a boolean if a field has been set.
func (o *CancelRequest) HasOriginalMerchantReference() bool {
	if o != nil && !common.IsNil(o.OriginalMerchantReference) {
		return true
	}

	return false
}

// SetOriginalMerchantReference gets a reference to the given string and assigns it to the OriginalMerchantReference field.
func (o *CancelRequest) SetOriginalMerchantReference(v string) {
	o.OriginalMerchantReference = &v
}

// GetOriginalReference returns the OriginalReference field value
func (o *CancelRequest) GetOriginalReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalReference
}

// GetOriginalReferenceOk returns a tuple with the OriginalReference field value
// and a boolean to check if the value has been set.
func (o *CancelRequest) GetOriginalReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalReference, true
}

// SetOriginalReference sets field value
func (o *CancelRequest) SetOriginalReference(v string) {
	o.OriginalReference = v
}

// GetPlatformChargebackLogic returns the PlatformChargebackLogic field value if set, zero value otherwise.
func (o *CancelRequest) GetPlatformChargebackLogic() PlatformChargebackLogic {
	if o == nil || common.IsNil(o.PlatformChargebackLogic) {
		var ret PlatformChargebackLogic
		return ret
	}
	return *o.PlatformChargebackLogic
}

// GetPlatformChargebackLogicOk returns a tuple with the PlatformChargebackLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelRequest) GetPlatformChargebackLogicOk() (*PlatformChargebackLogic, bool) {
	if o == nil || common.IsNil(o.PlatformChargebackLogic) {
		return nil, false
	}
	return o.PlatformChargebackLogic, true
}

// HasPlatformChargebackLogic returns a boolean if a field has been set.
func (o *CancelRequest) HasPlatformChargebackLogic() bool {
	if o != nil && !common.IsNil(o.PlatformChargebackLogic) {
		return true
	}

	return false
}

// SetPlatformChargebackLogic gets a reference to the given PlatformChargebackLogic and assigns it to the PlatformChargebackLogic field.
func (o *CancelRequest) SetPlatformChargebackLogic(v PlatformChargebackLogic) {
	o.PlatformChargebackLogic = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *CancelRequest) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelRequest) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *CancelRequest) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *CancelRequest) SetReference(v string) {
	o.Reference = &v
}

// GetSplits returns the Splits field value if set, zero value otherwise.
func (o *CancelRequest) GetSplits() []Split {
	if o == nil || common.IsNil(o.Splits) {
		var ret []Split
		return ret
	}
	return o.Splits
}

// GetSplitsOk returns a tuple with the Splits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelRequest) GetSplitsOk() ([]Split, bool) {
	if o == nil || common.IsNil(o.Splits) {
		return nil, false
	}
	return o.Splits, true
}

// HasSplits returns a boolean if a field has been set.
func (o *CancelRequest) HasSplits() bool {
	if o != nil && !common.IsNil(o.Splits) {
		return true
	}

	return false
}

// SetSplits gets a reference to the given []Split and assigns it to the Splits field.
func (o *CancelRequest) SetSplits(v []Split) {
	o.Splits = v
}

// GetTenderReference returns the TenderReference field value if set, zero value otherwise.
func (o *CancelRequest) GetTenderReference() string {
	if o == nil || common.IsNil(o.TenderReference) {
		var ret string
		return ret
	}
	return *o.TenderReference
}

// GetTenderReferenceOk returns a tuple with the TenderReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelRequest) GetTenderReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TenderReference) {
		return nil, false
	}
	return o.TenderReference, true
}

// HasTenderReference returns a boolean if a field has been set.
func (o *CancelRequest) HasTenderReference() bool {
	if o != nil && !common.IsNil(o.TenderReference) {
		return true
	}

	return false
}

// SetTenderReference gets a reference to the given string and assigns it to the TenderReference field.
func (o *CancelRequest) SetTenderReference(v string) {
	o.TenderReference = &v
}

// GetUniqueTerminalId returns the UniqueTerminalId field value if set, zero value otherwise.
func (o *CancelRequest) GetUniqueTerminalId() string {
	if o == nil || common.IsNil(o.UniqueTerminalId) {
		var ret string
		return ret
	}
	return *o.UniqueTerminalId
}

// GetUniqueTerminalIdOk returns a tuple with the UniqueTerminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelRequest) GetUniqueTerminalIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.UniqueTerminalId) {
		return nil, false
	}
	return o.UniqueTerminalId, true
}

// HasUniqueTerminalId returns a boolean if a field has been set.
func (o *CancelRequest) HasUniqueTerminalId() bool {
	if o != nil && !common.IsNil(o.UniqueTerminalId) {
		return true
	}

	return false
}

// SetUniqueTerminalId gets a reference to the given string and assigns it to the UniqueTerminalId field.
func (o *CancelRequest) SetUniqueTerminalId(v string) {
	o.UniqueTerminalId = &v
}

func (o CancelRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if !common.IsNil(o.MpiData) {
		toSerialize["mpiData"] = o.MpiData
	}
	if !common.IsNil(o.OriginalMerchantReference) {
		toSerialize["originalMerchantReference"] = o.OriginalMerchantReference
	}
	toSerialize["originalReference"] = o.OriginalReference
	if !common.IsNil(o.PlatformChargebackLogic) {
		toSerialize["platformChargebackLogic"] = o.PlatformChargebackLogic
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Splits) {
		toSerialize["splits"] = o.Splits
	}
	if !common.IsNil(o.TenderReference) {
		toSerialize["tenderReference"] = o.TenderReference
	}
	if !common.IsNil(o.UniqueTerminalId) {
		toSerialize["uniqueTerminalId"] = o.UniqueTerminalId
	}
	return toSerialize, nil
}

type NullableCancelRequest struct {
	value *CancelRequest
	isSet bool
}

func (v NullableCancelRequest) Get() *CancelRequest {
	return v.value
}

func (v *NullableCancelRequest) Set(val *CancelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelRequest(val *CancelRequest) *NullableCancelRequest {
	return &NullableCancelRequest{value: val, isSet: true}
}

func (v NullableCancelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
