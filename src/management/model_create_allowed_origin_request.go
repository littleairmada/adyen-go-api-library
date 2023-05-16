/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the CreateAllowedOriginRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateAllowedOriginRequest{}

// CreateAllowedOriginRequest struct for CreateAllowedOriginRequest
type CreateAllowedOriginRequest struct {
	Links *Links `json:"_links,omitempty"`
	// Domain of the allowed origin.
	Domain string `json:"domain"`
	// Unique identifier of the allowed origin.
	Id *string `json:"id,omitempty"`
}

// NewCreateAllowedOriginRequest instantiates a new CreateAllowedOriginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAllowedOriginRequest(domain string) *CreateAllowedOriginRequest {
	this := CreateAllowedOriginRequest{}
	this.Domain = domain
	return &this
}

// NewCreateAllowedOriginRequestWithDefaults instantiates a new CreateAllowedOriginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAllowedOriginRequestWithDefaults() *CreateAllowedOriginRequest {
	this := CreateAllowedOriginRequest{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreateAllowedOriginRequest) GetLinks() Links {
	if o == nil || common.IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAllowedOriginRequest) GetLinksOk() (*Links, bool) {
	if o == nil || common.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreateAllowedOriginRequest) HasLinks() bool {
	if o != nil && !common.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *CreateAllowedOriginRequest) SetLinks(v Links) {
	o.Links = &v
}

// GetDomain returns the Domain field value
func (o *CreateAllowedOriginRequest) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *CreateAllowedOriginRequest) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *CreateAllowedOriginRequest) SetDomain(v string) {
	o.Domain = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateAllowedOriginRequest) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAllowedOriginRequest) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateAllowedOriginRequest) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateAllowedOriginRequest) SetId(v string) {
	o.Id = &v
}

func (o CreateAllowedOriginRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAllowedOriginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	toSerialize["domain"] = o.Domain
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCreateAllowedOriginRequest struct {
	value *CreateAllowedOriginRequest
	isSet bool
}

func (v NullableCreateAllowedOriginRequest) Get() *CreateAllowedOriginRequest {
	return v.value
}

func (v *NullableCreateAllowedOriginRequest) Set(val *CreateAllowedOriginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAllowedOriginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAllowedOriginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAllowedOriginRequest(val *CreateAllowedOriginRequest) *NullableCreateAllowedOriginRequest {
	return &NullableCreateAllowedOriginRequest{value: val, isSet: true}
}

func (v NullableCreateAllowedOriginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAllowedOriginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
