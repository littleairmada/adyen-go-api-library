/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the IdentificationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IdentificationData{}

// IdentificationData struct for IdentificationData
type IdentificationData struct {
	// The card number of the document that was issued (AU only).
	CardNumber *string `json:"cardNumber,omitempty"`
	// The expiry date of the document, in YYYY-MM-DD format.
	ExpiryDate *string `json:"expiryDate,omitempty"`
	// The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the document was issued. For example, **US**.
    // Deprecated since Legal Entity Management API v1
	IssuerCountry *string `json:"issuerCountry,omitempty"`
	// The state or province where the document was issued (AU only).
	IssuerState *string `json:"issuerState,omitempty"`
	// Applies only to individuals in the US. Set to **true** if the individual does not have an SSN. To verify their identity, Adyen will require them to upload an ID document.
	NationalIdExempt *bool `json:"nationalIdExempt,omitempty"`
	// The number in the document.
	Number *string `json:"number,omitempty"`
	// Type of identity data. For individuals, the following types are supported. See our [onboarding guide](https://docs.adyen.com/platforms/onboard-users/onboarding-steps/?onboarding_type=custom) for other supported countries.  - Australia: **driversLicense**, **passport**  - Hong Kong: **driversLicense**, **nationalIdNumber**, **passport**  - New Zealand: **driversLicense**, **passport**  - Singapore: **driversLicense**, **nationalIdNumber**, **passport**   - All other supported countries: **nationalIdNumber**
	Type string `json:"type"`
}

// NewIdentificationData instantiates a new IdentificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentificationData(type_ string) *IdentificationData {
	this := IdentificationData{}
	this.Type = type_
	return &this
}

// NewIdentificationDataWithDefaults instantiates a new IdentificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentificationDataWithDefaults() *IdentificationData {
	this := IdentificationData{}
	return &this
}

// GetCardNumber returns the CardNumber field value if set, zero value otherwise.
func (o *IdentificationData) GetCardNumber() string {
	if o == nil || common.IsNil(o.CardNumber) {
		var ret string
		return ret
	}
	return *o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationData) GetCardNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardNumber) {
		return nil, false
	}
	return o.CardNumber, true
}

// HasCardNumber returns a boolean if a field has been set.
func (o *IdentificationData) HasCardNumber() bool {
	if o != nil && !common.IsNil(o.CardNumber) {
		return true
	}

	return false
}

// SetCardNumber gets a reference to the given string and assigns it to the CardNumber field.
func (o *IdentificationData) SetCardNumber(v string) {
	o.CardNumber = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *IdentificationData) GetExpiryDate() string {
	if o == nil || common.IsNil(o.ExpiryDate) {
		var ret string
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationData) GetExpiryDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *IdentificationData) HasExpiryDate() bool {
	if o != nil && !common.IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.
func (o *IdentificationData) SetExpiryDate(v string) {
	o.ExpiryDate = &v
}

// GetIssuerCountry returns the IssuerCountry field value if set, zero value otherwise.
// Deprecated since Legal Entity Management API v1
func (o *IdentificationData) GetIssuerCountry() string {
	if o == nil || common.IsNil(o.IssuerCountry) {
		var ret string
		return ret
	}
	return *o.IssuerCountry
}

// GetIssuerCountryOk returns a tuple with the IssuerCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Legal Entity Management API v1
func (o *IdentificationData) GetIssuerCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuerCountry) {
		return nil, false
	}
	return o.IssuerCountry, true
}

// HasIssuerCountry returns a boolean if a field has been set.
func (o *IdentificationData) HasIssuerCountry() bool {
	if o != nil && !common.IsNil(o.IssuerCountry) {
		return true
	}

	return false
}

// SetIssuerCountry gets a reference to the given string and assigns it to the IssuerCountry field.
// Deprecated since Legal Entity Management API v1
func (o *IdentificationData) SetIssuerCountry(v string) {
	o.IssuerCountry = &v
}

// GetIssuerState returns the IssuerState field value if set, zero value otherwise.
func (o *IdentificationData) GetIssuerState() string {
	if o == nil || common.IsNil(o.IssuerState) {
		var ret string
		return ret
	}
	return *o.IssuerState
}

// GetIssuerStateOk returns a tuple with the IssuerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationData) GetIssuerStateOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuerState) {
		return nil, false
	}
	return o.IssuerState, true
}

// HasIssuerState returns a boolean if a field has been set.
func (o *IdentificationData) HasIssuerState() bool {
	if o != nil && !common.IsNil(o.IssuerState) {
		return true
	}

	return false
}

// SetIssuerState gets a reference to the given string and assigns it to the IssuerState field.
func (o *IdentificationData) SetIssuerState(v string) {
	o.IssuerState = &v
}

// GetNationalIdExempt returns the NationalIdExempt field value if set, zero value otherwise.
func (o *IdentificationData) GetNationalIdExempt() bool {
	if o == nil || common.IsNil(o.NationalIdExempt) {
		var ret bool
		return ret
	}
	return *o.NationalIdExempt
}

// GetNationalIdExemptOk returns a tuple with the NationalIdExempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationData) GetNationalIdExemptOk() (*bool, bool) {
	if o == nil || common.IsNil(o.NationalIdExempt) {
		return nil, false
	}
	return o.NationalIdExempt, true
}

// HasNationalIdExempt returns a boolean if a field has been set.
func (o *IdentificationData) HasNationalIdExempt() bool {
	if o != nil && !common.IsNil(o.NationalIdExempt) {
		return true
	}

	return false
}

// SetNationalIdExempt gets a reference to the given bool and assigns it to the NationalIdExempt field.
func (o *IdentificationData) SetNationalIdExempt(v bool) {
	o.NationalIdExempt = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *IdentificationData) GetNumber() string {
	if o == nil || common.IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationData) GetNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *IdentificationData) HasNumber() bool {
	if o != nil && !common.IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *IdentificationData) SetNumber(v string) {
	o.Number = &v
}

// GetType returns the Type field value
func (o *IdentificationData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IdentificationData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IdentificationData) SetType(v string) {
	o.Type = v
}

func (o IdentificationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentificationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CardNumber) {
		toSerialize["cardNumber"] = o.CardNumber
	}
	if !common.IsNil(o.ExpiryDate) {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if !common.IsNil(o.IssuerCountry) {
		toSerialize["issuerCountry"] = o.IssuerCountry
	}
	if !common.IsNil(o.IssuerState) {
		toSerialize["issuerState"] = o.IssuerState
	}
	if !common.IsNil(o.NationalIdExempt) {
		toSerialize["nationalIdExempt"] = o.NationalIdExempt
	}
	if !common.IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableIdentificationData struct {
	value *IdentificationData
	isSet bool
}

func (v NullableIdentificationData) Get() *IdentificationData {
	return v.value
}

func (v *NullableIdentificationData) Set(val *IdentificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentificationData(val *IdentificationData) *NullableIdentificationData {
	return &NullableIdentificationData{value: val, isSet: true}
}

func (v NullableIdentificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *IdentificationData) isValidType() bool {
    var allowedEnumValues = []string{ "nationalIdNumber", "passport", "driversLicense", "identityCard" }
    for _, allowed := range allowedEnumValues {
        if o.GetType() == allowed {
            return true
        }
    }
    return false
}

