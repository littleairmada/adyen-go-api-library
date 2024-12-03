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

// checks if the GetTermsOfServiceDocumentRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetTermsOfServiceDocumentRequest{}

// GetTermsOfServiceDocumentRequest struct for GetTermsOfServiceDocumentRequest
type GetTermsOfServiceDocumentRequest struct {
	// The language to be used for the Terms of Service document, specified by the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code. Possible value: **en** for English.
	Language string `json:"language"`
	// The requested format for the Terms of Service document. Default value: JSON. Possible values: **JSON**, **PDF**, or **TXT**.
	TermsOfServiceDocumentFormat *string `json:"termsOfServiceDocumentFormat,omitempty"`
	// The type of Terms of Service.  Possible values: *  **adyenForPlatformsManage** *  **adyenIssuing** *  **adyenForPlatformsAdvanced** *  **adyenCapital** *  **adyenAccount** *  **adyenCard** *  **adyenFranchisee** *  **adyenPccr** *  **adyenChargeCard**  
	Type string `json:"type"`
}

// NewGetTermsOfServiceDocumentRequest instantiates a new GetTermsOfServiceDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTermsOfServiceDocumentRequest(language string, type_ string) *GetTermsOfServiceDocumentRequest {
	this := GetTermsOfServiceDocumentRequest{}
	this.Language = language
	this.Type = type_
	return &this
}

// NewGetTermsOfServiceDocumentRequestWithDefaults instantiates a new GetTermsOfServiceDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTermsOfServiceDocumentRequestWithDefaults() *GetTermsOfServiceDocumentRequest {
	this := GetTermsOfServiceDocumentRequest{}
	return &this
}

// GetLanguage returns the Language field value
func (o *GetTermsOfServiceDocumentRequest) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *GetTermsOfServiceDocumentRequest) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *GetTermsOfServiceDocumentRequest) SetLanguage(v string) {
	o.Language = v
}

// GetTermsOfServiceDocumentFormat returns the TermsOfServiceDocumentFormat field value if set, zero value otherwise.
func (o *GetTermsOfServiceDocumentRequest) GetTermsOfServiceDocumentFormat() string {
	if o == nil || common.IsNil(o.TermsOfServiceDocumentFormat) {
		var ret string
		return ret
	}
	return *o.TermsOfServiceDocumentFormat
}

// GetTermsOfServiceDocumentFormatOk returns a tuple with the TermsOfServiceDocumentFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTermsOfServiceDocumentRequest) GetTermsOfServiceDocumentFormatOk() (*string, bool) {
	if o == nil || common.IsNil(o.TermsOfServiceDocumentFormat) {
		return nil, false
	}
	return o.TermsOfServiceDocumentFormat, true
}

// HasTermsOfServiceDocumentFormat returns a boolean if a field has been set.
func (o *GetTermsOfServiceDocumentRequest) HasTermsOfServiceDocumentFormat() bool {
	if o != nil && !common.IsNil(o.TermsOfServiceDocumentFormat) {
		return true
	}

	return false
}

// SetTermsOfServiceDocumentFormat gets a reference to the given string and assigns it to the TermsOfServiceDocumentFormat field.
func (o *GetTermsOfServiceDocumentRequest) SetTermsOfServiceDocumentFormat(v string) {
	o.TermsOfServiceDocumentFormat = &v
}

// GetType returns the Type field value
func (o *GetTermsOfServiceDocumentRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetTermsOfServiceDocumentRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetTermsOfServiceDocumentRequest) SetType(v string) {
	o.Type = v
}

func (o GetTermsOfServiceDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTermsOfServiceDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["language"] = o.Language
	if !common.IsNil(o.TermsOfServiceDocumentFormat) {
		toSerialize["termsOfServiceDocumentFormat"] = o.TermsOfServiceDocumentFormat
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableGetTermsOfServiceDocumentRequest struct {
	value *GetTermsOfServiceDocumentRequest
	isSet bool
}

func (v NullableGetTermsOfServiceDocumentRequest) Get() *GetTermsOfServiceDocumentRequest {
	return v.value
}

func (v *NullableGetTermsOfServiceDocumentRequest) Set(val *GetTermsOfServiceDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTermsOfServiceDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTermsOfServiceDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTermsOfServiceDocumentRequest(val *GetTermsOfServiceDocumentRequest) *NullableGetTermsOfServiceDocumentRequest {
	return &NullableGetTermsOfServiceDocumentRequest{value: val, isSet: true}
}

func (v NullableGetTermsOfServiceDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTermsOfServiceDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *GetTermsOfServiceDocumentRequest) isValidType() bool {
    var allowedEnumValues = []string{ "adyenAccount", "adyenCapital", "adyenCard", "adyenChargeCard", "adyenForPlatformsAdvanced", "adyenForPlatformsManage", "adyenFranchisee", "adyenIssuing", "adyenPccr" }
    for _, allowed := range allowedEnumValues {
        if o.GetType() == allowed {
            return true
        }
    }
    return false
}

