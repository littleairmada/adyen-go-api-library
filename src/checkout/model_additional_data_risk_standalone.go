/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the AdditionalDataRiskStandalone type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalDataRiskStandalone{}

// AdditionalDataRiskStandalone struct for AdditionalDataRiskStandalone
type AdditionalDataRiskStandalone struct {
	// Shopper's country of residence in the form of ISO standard 3166 2-character country codes.
	PayPalCountryCode *string `json:"PayPal.CountryCode,omitempty"`
	// Shopper's email.
	PayPalEmailId *string `json:"PayPal.EmailId,omitempty"`
	// Shopper's first name.
	PayPalFirstName *string `json:"PayPal.FirstName,omitempty"`
	// Shopper's last name.
	PayPalLastName *string `json:"PayPal.LastName,omitempty"`
	// Unique PayPal Customer Account identification number. Character length and limitations: 13 single-byte alphanumeric characters.
	PayPalPayerId *string `json:"PayPal.PayerId,omitempty"`
	// Shopper's phone number.
	PayPalPhone *string `json:"PayPal.Phone,omitempty"`
	// Allowed values: * **Eligible** — Merchant is protected by PayPal's Seller Protection Policy for Unauthorized Payments and Item Not Received.  * **PartiallyEligible** — Merchant is protected by PayPal's Seller Protection Policy for Item Not Received.  * **Ineligible** — Merchant is not protected under the Seller Protection Policy.
	PayPalProtectionEligibility *string `json:"PayPal.ProtectionEligibility,omitempty"`
	// Unique transaction ID of the payment.
	PayPalTransactionId *string `json:"PayPal.TransactionId,omitempty"`
	// Raw AVS result received from the acquirer, where available. Example: D
	AvsResultRaw *string `json:"avsResultRaw,omitempty"`
	// The Bank Identification Number of a credit card, which is the first six digits of a card number. Required for [tokenized card request](https://docs.adyen.com/risk-management/standalone-risk#tokenised-pan-request).
	Bin *string `json:"bin,omitempty"`
	// Raw CVC result received from the acquirer, where available. Example: 1
	CvcResultRaw *string `json:"cvcResultRaw,omitempty"`
	// Unique identifier or token for the shopper's card details.
	RiskToken *string `json:"riskToken,omitempty"`
	// A Boolean value indicating whether 3DS authentication was completed on this payment. Example: true
	ThreeDAuthenticated *string `json:"threeDAuthenticated,omitempty"`
	// A Boolean value indicating whether 3DS was offered for this payment. Example: true
	ThreeDOffered *string `json:"threeDOffered,omitempty"`
	// Required for PayPal payments only. The only supported value is: **paypal**.
	TokenDataType *string `json:"tokenDataType,omitempty"`
}

// NewAdditionalDataRiskStandalone instantiates a new AdditionalDataRiskStandalone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalDataRiskStandalone() *AdditionalDataRiskStandalone {
	this := AdditionalDataRiskStandalone{}
	return &this
}

// NewAdditionalDataRiskStandaloneWithDefaults instantiates a new AdditionalDataRiskStandalone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataRiskStandaloneWithDefaults() *AdditionalDataRiskStandalone {
	this := AdditionalDataRiskStandalone{}
	return &this
}

// GetPayPalCountryCode returns the PayPalCountryCode field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetPayPalCountryCode() string {
	if o == nil || common.IsNil(o.PayPalCountryCode) {
		var ret string
		return ret
	}
	return *o.PayPalCountryCode
}

// GetPayPalCountryCodeOk returns a tuple with the PayPalCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetPayPalCountryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayPalCountryCode) {
		return nil, false
	}
	return o.PayPalCountryCode, true
}

// HasPayPalCountryCode returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasPayPalCountryCode() bool {
	if o != nil && !common.IsNil(o.PayPalCountryCode) {
		return true
	}

	return false
}

// SetPayPalCountryCode gets a reference to the given string and assigns it to the PayPalCountryCode field.
func (o *AdditionalDataRiskStandalone) SetPayPalCountryCode(v string) {
	o.PayPalCountryCode = &v
}

// GetPayPalEmailId returns the PayPalEmailId field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetPayPalEmailId() string {
	if o == nil || common.IsNil(o.PayPalEmailId) {
		var ret string
		return ret
	}
	return *o.PayPalEmailId
}

// GetPayPalEmailIdOk returns a tuple with the PayPalEmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetPayPalEmailIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayPalEmailId) {
		return nil, false
	}
	return o.PayPalEmailId, true
}

// HasPayPalEmailId returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasPayPalEmailId() bool {
	if o != nil && !common.IsNil(o.PayPalEmailId) {
		return true
	}

	return false
}

// SetPayPalEmailId gets a reference to the given string and assigns it to the PayPalEmailId field.
func (o *AdditionalDataRiskStandalone) SetPayPalEmailId(v string) {
	o.PayPalEmailId = &v
}

// GetPayPalFirstName returns the PayPalFirstName field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetPayPalFirstName() string {
	if o == nil || common.IsNil(o.PayPalFirstName) {
		var ret string
		return ret
	}
	return *o.PayPalFirstName
}

// GetPayPalFirstNameOk returns a tuple with the PayPalFirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetPayPalFirstNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayPalFirstName) {
		return nil, false
	}
	return o.PayPalFirstName, true
}

// HasPayPalFirstName returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasPayPalFirstName() bool {
	if o != nil && !common.IsNil(o.PayPalFirstName) {
		return true
	}

	return false
}

// SetPayPalFirstName gets a reference to the given string and assigns it to the PayPalFirstName field.
func (o *AdditionalDataRiskStandalone) SetPayPalFirstName(v string) {
	o.PayPalFirstName = &v
}

// GetPayPalLastName returns the PayPalLastName field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetPayPalLastName() string {
	if o == nil || common.IsNil(o.PayPalLastName) {
		var ret string
		return ret
	}
	return *o.PayPalLastName
}

// GetPayPalLastNameOk returns a tuple with the PayPalLastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetPayPalLastNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayPalLastName) {
		return nil, false
	}
	return o.PayPalLastName, true
}

// HasPayPalLastName returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasPayPalLastName() bool {
	if o != nil && !common.IsNil(o.PayPalLastName) {
		return true
	}

	return false
}

// SetPayPalLastName gets a reference to the given string and assigns it to the PayPalLastName field.
func (o *AdditionalDataRiskStandalone) SetPayPalLastName(v string) {
	o.PayPalLastName = &v
}

// GetPayPalPayerId returns the PayPalPayerId field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetPayPalPayerId() string {
	if o == nil || common.IsNil(o.PayPalPayerId) {
		var ret string
		return ret
	}
	return *o.PayPalPayerId
}

// GetPayPalPayerIdOk returns a tuple with the PayPalPayerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetPayPalPayerIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayPalPayerId) {
		return nil, false
	}
	return o.PayPalPayerId, true
}

// HasPayPalPayerId returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasPayPalPayerId() bool {
	if o != nil && !common.IsNil(o.PayPalPayerId) {
		return true
	}

	return false
}

// SetPayPalPayerId gets a reference to the given string and assigns it to the PayPalPayerId field.
func (o *AdditionalDataRiskStandalone) SetPayPalPayerId(v string) {
	o.PayPalPayerId = &v
}

// GetPayPalPhone returns the PayPalPhone field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetPayPalPhone() string {
	if o == nil || common.IsNil(o.PayPalPhone) {
		var ret string
		return ret
	}
	return *o.PayPalPhone
}

// GetPayPalPhoneOk returns a tuple with the PayPalPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetPayPalPhoneOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayPalPhone) {
		return nil, false
	}
	return o.PayPalPhone, true
}

// HasPayPalPhone returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasPayPalPhone() bool {
	if o != nil && !common.IsNil(o.PayPalPhone) {
		return true
	}

	return false
}

// SetPayPalPhone gets a reference to the given string and assigns it to the PayPalPhone field.
func (o *AdditionalDataRiskStandalone) SetPayPalPhone(v string) {
	o.PayPalPhone = &v
}

// GetPayPalProtectionEligibility returns the PayPalProtectionEligibility field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetPayPalProtectionEligibility() string {
	if o == nil || common.IsNil(o.PayPalProtectionEligibility) {
		var ret string
		return ret
	}
	return *o.PayPalProtectionEligibility
}

// GetPayPalProtectionEligibilityOk returns a tuple with the PayPalProtectionEligibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetPayPalProtectionEligibilityOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayPalProtectionEligibility) {
		return nil, false
	}
	return o.PayPalProtectionEligibility, true
}

// HasPayPalProtectionEligibility returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasPayPalProtectionEligibility() bool {
	if o != nil && !common.IsNil(o.PayPalProtectionEligibility) {
		return true
	}

	return false
}

// SetPayPalProtectionEligibility gets a reference to the given string and assigns it to the PayPalProtectionEligibility field.
func (o *AdditionalDataRiskStandalone) SetPayPalProtectionEligibility(v string) {
	o.PayPalProtectionEligibility = &v
}

// GetPayPalTransactionId returns the PayPalTransactionId field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetPayPalTransactionId() string {
	if o == nil || common.IsNil(o.PayPalTransactionId) {
		var ret string
		return ret
	}
	return *o.PayPalTransactionId
}

// GetPayPalTransactionIdOk returns a tuple with the PayPalTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetPayPalTransactionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayPalTransactionId) {
		return nil, false
	}
	return o.PayPalTransactionId, true
}

// HasPayPalTransactionId returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasPayPalTransactionId() bool {
	if o != nil && !common.IsNil(o.PayPalTransactionId) {
		return true
	}

	return false
}

// SetPayPalTransactionId gets a reference to the given string and assigns it to the PayPalTransactionId field.
func (o *AdditionalDataRiskStandalone) SetPayPalTransactionId(v string) {
	o.PayPalTransactionId = &v
}

// GetAvsResultRaw returns the AvsResultRaw field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetAvsResultRaw() string {
	if o == nil || common.IsNil(o.AvsResultRaw) {
		var ret string
		return ret
	}
	return *o.AvsResultRaw
}

// GetAvsResultRawOk returns a tuple with the AvsResultRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetAvsResultRawOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvsResultRaw) {
		return nil, false
	}
	return o.AvsResultRaw, true
}

// HasAvsResultRaw returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasAvsResultRaw() bool {
	if o != nil && !common.IsNil(o.AvsResultRaw) {
		return true
	}

	return false
}

// SetAvsResultRaw gets a reference to the given string and assigns it to the AvsResultRaw field.
func (o *AdditionalDataRiskStandalone) SetAvsResultRaw(v string) {
	o.AvsResultRaw = &v
}

// GetBin returns the Bin field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetBin() string {
	if o == nil || common.IsNil(o.Bin) {
		var ret string
		return ret
	}
	return *o.Bin
}

// GetBinOk returns a tuple with the Bin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetBinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Bin) {
		return nil, false
	}
	return o.Bin, true
}

// HasBin returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasBin() bool {
	if o != nil && !common.IsNil(o.Bin) {
		return true
	}

	return false
}

// SetBin gets a reference to the given string and assigns it to the Bin field.
func (o *AdditionalDataRiskStandalone) SetBin(v string) {
	o.Bin = &v
}

// GetCvcResultRaw returns the CvcResultRaw field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetCvcResultRaw() string {
	if o == nil || common.IsNil(o.CvcResultRaw) {
		var ret string
		return ret
	}
	return *o.CvcResultRaw
}

// GetCvcResultRawOk returns a tuple with the CvcResultRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetCvcResultRawOk() (*string, bool) {
	if o == nil || common.IsNil(o.CvcResultRaw) {
		return nil, false
	}
	return o.CvcResultRaw, true
}

// HasCvcResultRaw returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasCvcResultRaw() bool {
	if o != nil && !common.IsNil(o.CvcResultRaw) {
		return true
	}

	return false
}

// SetCvcResultRaw gets a reference to the given string and assigns it to the CvcResultRaw field.
func (o *AdditionalDataRiskStandalone) SetCvcResultRaw(v string) {
	o.CvcResultRaw = &v
}

// GetRiskToken returns the RiskToken field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetRiskToken() string {
	if o == nil || common.IsNil(o.RiskToken) {
		var ret string
		return ret
	}
	return *o.RiskToken
}

// GetRiskTokenOk returns a tuple with the RiskToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetRiskTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskToken) {
		return nil, false
	}
	return o.RiskToken, true
}

// HasRiskToken returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasRiskToken() bool {
	if o != nil && !common.IsNil(o.RiskToken) {
		return true
	}

	return false
}

// SetRiskToken gets a reference to the given string and assigns it to the RiskToken field.
func (o *AdditionalDataRiskStandalone) SetRiskToken(v string) {
	o.RiskToken = &v
}

// GetThreeDAuthenticated returns the ThreeDAuthenticated field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetThreeDAuthenticated() string {
	if o == nil || common.IsNil(o.ThreeDAuthenticated) {
		var ret string
		return ret
	}
	return *o.ThreeDAuthenticated
}

// GetThreeDAuthenticatedOk returns a tuple with the ThreeDAuthenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetThreeDAuthenticatedOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDAuthenticated) {
		return nil, false
	}
	return o.ThreeDAuthenticated, true
}

// HasThreeDAuthenticated returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasThreeDAuthenticated() bool {
	if o != nil && !common.IsNil(o.ThreeDAuthenticated) {
		return true
	}

	return false
}

// SetThreeDAuthenticated gets a reference to the given string and assigns it to the ThreeDAuthenticated field.
func (o *AdditionalDataRiskStandalone) SetThreeDAuthenticated(v string) {
	o.ThreeDAuthenticated = &v
}

// GetThreeDOffered returns the ThreeDOffered field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetThreeDOffered() string {
	if o == nil || common.IsNil(o.ThreeDOffered) {
		var ret string
		return ret
	}
	return *o.ThreeDOffered
}

// GetThreeDOfferedOk returns a tuple with the ThreeDOffered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetThreeDOfferedOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDOffered) {
		return nil, false
	}
	return o.ThreeDOffered, true
}

// HasThreeDOffered returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasThreeDOffered() bool {
	if o != nil && !common.IsNil(o.ThreeDOffered) {
		return true
	}

	return false
}

// SetThreeDOffered gets a reference to the given string and assigns it to the ThreeDOffered field.
func (o *AdditionalDataRiskStandalone) SetThreeDOffered(v string) {
	o.ThreeDOffered = &v
}

// GetTokenDataType returns the TokenDataType field value if set, zero value otherwise.
func (o *AdditionalDataRiskStandalone) GetTokenDataType() string {
	if o == nil || common.IsNil(o.TokenDataType) {
		var ret string
		return ret
	}
	return *o.TokenDataType
}

// GetTokenDataTypeOk returns a tuple with the TokenDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRiskStandalone) GetTokenDataTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenDataType) {
		return nil, false
	}
	return o.TokenDataType, true
}

// HasTokenDataType returns a boolean if a field has been set.
func (o *AdditionalDataRiskStandalone) HasTokenDataType() bool {
	if o != nil && !common.IsNil(o.TokenDataType) {
		return true
	}

	return false
}

// SetTokenDataType gets a reference to the given string and assigns it to the TokenDataType field.
func (o *AdditionalDataRiskStandalone) SetTokenDataType(v string) {
	o.TokenDataType = &v
}

func (o AdditionalDataRiskStandalone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalDataRiskStandalone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PayPalCountryCode) {
		toSerialize["PayPal.CountryCode"] = o.PayPalCountryCode
	}
	if !common.IsNil(o.PayPalEmailId) {
		toSerialize["PayPal.EmailId"] = o.PayPalEmailId
	}
	if !common.IsNil(o.PayPalFirstName) {
		toSerialize["PayPal.FirstName"] = o.PayPalFirstName
	}
	if !common.IsNil(o.PayPalLastName) {
		toSerialize["PayPal.LastName"] = o.PayPalLastName
	}
	if !common.IsNil(o.PayPalPayerId) {
		toSerialize["PayPal.PayerId"] = o.PayPalPayerId
	}
	if !common.IsNil(o.PayPalPhone) {
		toSerialize["PayPal.Phone"] = o.PayPalPhone
	}
	if !common.IsNil(o.PayPalProtectionEligibility) {
		toSerialize["PayPal.ProtectionEligibility"] = o.PayPalProtectionEligibility
	}
	if !common.IsNil(o.PayPalTransactionId) {
		toSerialize["PayPal.TransactionId"] = o.PayPalTransactionId
	}
	if !common.IsNil(o.AvsResultRaw) {
		toSerialize["avsResultRaw"] = o.AvsResultRaw
	}
	if !common.IsNil(o.Bin) {
		toSerialize["bin"] = o.Bin
	}
	if !common.IsNil(o.CvcResultRaw) {
		toSerialize["cvcResultRaw"] = o.CvcResultRaw
	}
	if !common.IsNil(o.RiskToken) {
		toSerialize["riskToken"] = o.RiskToken
	}
	if !common.IsNil(o.ThreeDAuthenticated) {
		toSerialize["threeDAuthenticated"] = o.ThreeDAuthenticated
	}
	if !common.IsNil(o.ThreeDOffered) {
		toSerialize["threeDOffered"] = o.ThreeDOffered
	}
	if !common.IsNil(o.TokenDataType) {
		toSerialize["tokenDataType"] = o.TokenDataType
	}
	return toSerialize, nil
}

type NullableAdditionalDataRiskStandalone struct {
	value *AdditionalDataRiskStandalone
	isSet bool
}

func (v NullableAdditionalDataRiskStandalone) Get() *AdditionalDataRiskStandalone {
	return v.value
}

func (v *NullableAdditionalDataRiskStandalone) Set(val *AdditionalDataRiskStandalone) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalDataRiskStandalone) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalDataRiskStandalone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalDataRiskStandalone(val *AdditionalDataRiskStandalone) *NullableAdditionalDataRiskStandalone {
	return &NullableAdditionalDataRiskStandalone{value: val, isSet: true}
}

func (v NullableAdditionalDataRiskStandalone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalDataRiskStandalone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
