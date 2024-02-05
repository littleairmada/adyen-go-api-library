/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the PciSigningRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PciSigningRequest{}

// PciSigningRequest struct for PciSigningRequest
type PciSigningRequest struct {
	// The array of Adyen-generated unique identifiers for the questionnaires.
	PciTemplateReferences []string `json:"pciTemplateReferences"`
	// The [legal entity ID](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id) of the individual who signs the PCI questionnaire.
	SignedBy string `json:"signedBy"`
}

// NewPciSigningRequest instantiates a new PciSigningRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPciSigningRequest(pciTemplateReferences []string, signedBy string) *PciSigningRequest {
	this := PciSigningRequest{}
	this.PciTemplateReferences = pciTemplateReferences
	this.SignedBy = signedBy
	return &this
}

// NewPciSigningRequestWithDefaults instantiates a new PciSigningRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPciSigningRequestWithDefaults() *PciSigningRequest {
	this := PciSigningRequest{}
	return &this
}

// GetPciTemplateReferences returns the PciTemplateReferences field value
func (o *PciSigningRequest) GetPciTemplateReferences() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PciTemplateReferences
}

// GetPciTemplateReferencesOk returns a tuple with the PciTemplateReferences field value
// and a boolean to check if the value has been set.
func (o *PciSigningRequest) GetPciTemplateReferencesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PciTemplateReferences, true
}

// SetPciTemplateReferences sets field value
func (o *PciSigningRequest) SetPciTemplateReferences(v []string) {
	o.PciTemplateReferences = v
}

// GetSignedBy returns the SignedBy field value
func (o *PciSigningRequest) GetSignedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignedBy
}

// GetSignedByOk returns a tuple with the SignedBy field value
// and a boolean to check if the value has been set.
func (o *PciSigningRequest) GetSignedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignedBy, true
}

// SetSignedBy sets field value
func (o *PciSigningRequest) SetSignedBy(v string) {
	o.SignedBy = v
}

func (o PciSigningRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PciSigningRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pciTemplateReferences"] = o.PciTemplateReferences
	toSerialize["signedBy"] = o.SignedBy
	return toSerialize, nil
}

type NullablePciSigningRequest struct {
	value *PciSigningRequest
	isSet bool
}

func (v NullablePciSigningRequest) Get() *PciSigningRequest {
	return v.value
}

func (v *NullablePciSigningRequest) Set(val *PciSigningRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePciSigningRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePciSigningRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePciSigningRequest(val *PciSigningRequest) *NullablePciSigningRequest {
	return &NullablePciSigningRequest{value: val, isSet: true}
}

func (v NullablePciSigningRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePciSigningRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
