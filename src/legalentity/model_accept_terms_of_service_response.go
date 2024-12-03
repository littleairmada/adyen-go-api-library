/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the AcceptTermsOfServiceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AcceptTermsOfServiceResponse{}

// AcceptTermsOfServiceResponse struct for AcceptTermsOfServiceResponse
type AcceptTermsOfServiceResponse struct {
	// The unique identifier of the user that accepted the Terms of Service.
	AcceptedBy *string `json:"acceptedBy,omitempty"`
	// The unique identifier of the Terms of Service acceptance.
	Id *string `json:"id,omitempty"`
	// The IP address of the user that accepted the Terms of Service.
	IpAddress *string `json:"ipAddress,omitempty"`
	// The language used for the Terms of Service document, specified by the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code. Possible value: **en** for English.
	Language *string `json:"language,omitempty"`
	// The unique identifier of the Terms of Service document.
	TermsOfServiceDocumentId *string `json:"termsOfServiceDocumentId,omitempty"`
	// The type of Terms of Service.  Possible values: *  **adyenForPlatformsManage** *  **adyenIssuing** *  **adyenForPlatformsAdvanced** *  **adyenCapital** *  **adyenAccount** *  **adyenCard** *  **adyenFranchisee** *  **adyenPccr** *  **adyenChargeCard**
	Type *string `json:"type,omitempty"`
}

// NewAcceptTermsOfServiceResponse instantiates a new AcceptTermsOfServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptTermsOfServiceResponse() *AcceptTermsOfServiceResponse {
	this := AcceptTermsOfServiceResponse{}
	return &this
}

// NewAcceptTermsOfServiceResponseWithDefaults instantiates a new AcceptTermsOfServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptTermsOfServiceResponseWithDefaults() *AcceptTermsOfServiceResponse {
	this := AcceptTermsOfServiceResponse{}
	return &this
}

// GetAcceptedBy returns the AcceptedBy field value if set, zero value otherwise.
func (o *AcceptTermsOfServiceResponse) GetAcceptedBy() string {
	if o == nil || common.IsNil(o.AcceptedBy) {
		var ret string
		return ret
	}
	return *o.AcceptedBy
}

// GetAcceptedByOk returns a tuple with the AcceptedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptTermsOfServiceResponse) GetAcceptedByOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcceptedBy) {
		return nil, false
	}
	return o.AcceptedBy, true
}

// HasAcceptedBy returns a boolean if a field has been set.
func (o *AcceptTermsOfServiceResponse) HasAcceptedBy() bool {
	if o != nil && !common.IsNil(o.AcceptedBy) {
		return true
	}

	return false
}

// SetAcceptedBy gets a reference to the given string and assigns it to the AcceptedBy field.
func (o *AcceptTermsOfServiceResponse) SetAcceptedBy(v string) {
	o.AcceptedBy = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AcceptTermsOfServiceResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptTermsOfServiceResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AcceptTermsOfServiceResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AcceptTermsOfServiceResponse) SetId(v string) {
	o.Id = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *AcceptTermsOfServiceResponse) GetIpAddress() string {
	if o == nil || common.IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptTermsOfServiceResponse) GetIpAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *AcceptTermsOfServiceResponse) HasIpAddress() bool {
	if o != nil && !common.IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *AcceptTermsOfServiceResponse) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AcceptTermsOfServiceResponse) GetLanguage() string {
	if o == nil || common.IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptTermsOfServiceResponse) GetLanguageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AcceptTermsOfServiceResponse) HasLanguage() bool {
	if o != nil && !common.IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *AcceptTermsOfServiceResponse) SetLanguage(v string) {
	o.Language = &v
}

// GetTermsOfServiceDocumentId returns the TermsOfServiceDocumentId field value if set, zero value otherwise.
func (o *AcceptTermsOfServiceResponse) GetTermsOfServiceDocumentId() string {
	if o == nil || common.IsNil(o.TermsOfServiceDocumentId) {
		var ret string
		return ret
	}
	return *o.TermsOfServiceDocumentId
}

// GetTermsOfServiceDocumentIdOk returns a tuple with the TermsOfServiceDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptTermsOfServiceResponse) GetTermsOfServiceDocumentIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TermsOfServiceDocumentId) {
		return nil, false
	}
	return o.TermsOfServiceDocumentId, true
}

// HasTermsOfServiceDocumentId returns a boolean if a field has been set.
func (o *AcceptTermsOfServiceResponse) HasTermsOfServiceDocumentId() bool {
	if o != nil && !common.IsNil(o.TermsOfServiceDocumentId) {
		return true
	}

	return false
}

// SetTermsOfServiceDocumentId gets a reference to the given string and assigns it to the TermsOfServiceDocumentId field.
func (o *AcceptTermsOfServiceResponse) SetTermsOfServiceDocumentId(v string) {
	o.TermsOfServiceDocumentId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AcceptTermsOfServiceResponse) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptTermsOfServiceResponse) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AcceptTermsOfServiceResponse) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AcceptTermsOfServiceResponse) SetType(v string) {
	o.Type = &v
}

func (o AcceptTermsOfServiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptTermsOfServiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AcceptedBy) {
		toSerialize["acceptedBy"] = o.AcceptedBy
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !common.IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !common.IsNil(o.TermsOfServiceDocumentId) {
		toSerialize["termsOfServiceDocumentId"] = o.TermsOfServiceDocumentId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAcceptTermsOfServiceResponse struct {
	value *AcceptTermsOfServiceResponse
	isSet bool
}

func (v NullableAcceptTermsOfServiceResponse) Get() *AcceptTermsOfServiceResponse {
	return v.value
}

func (v *NullableAcceptTermsOfServiceResponse) Set(val *AcceptTermsOfServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptTermsOfServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptTermsOfServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptTermsOfServiceResponse(val *AcceptTermsOfServiceResponse) *NullableAcceptTermsOfServiceResponse {
	return &NullableAcceptTermsOfServiceResponse{value: val, isSet: true}
}

func (v NullableAcceptTermsOfServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptTermsOfServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AcceptTermsOfServiceResponse) isValidType() bool {
	var allowedEnumValues = []string{"adyenAccount", "adyenCapital", "adyenCard", "adyenChargeCard", "adyenForPlatformsAdvanced", "adyenForPlatformsManage", "adyenFranchisee", "adyenIssuing", "adyenPccr"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
