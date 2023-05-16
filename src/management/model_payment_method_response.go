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

// checks if the PaymentMethodResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentMethodResponse{}

// PaymentMethodResponse struct for PaymentMethodResponse
type PaymentMethodResponse struct {
	Links *PaginationLinks `json:"_links,omitempty"`
	// Payment methods details.
	Data []PaymentMethod `json:"data,omitempty"`
	// Total number of items.
	ItemsTotal int32 `json:"itemsTotal"`
	// Total number of pages.
	PagesTotal int32 `json:"pagesTotal"`
	// Payment method types with errors.
	TypesWithErrors []string `json:"typesWithErrors,omitempty"`
}

// NewPaymentMethodResponse instantiates a new PaymentMethodResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodResponse(itemsTotal int32, pagesTotal int32) *PaymentMethodResponse {
	this := PaymentMethodResponse{}
	this.ItemsTotal = itemsTotal
	this.PagesTotal = pagesTotal
	return &this
}

// NewPaymentMethodResponseWithDefaults instantiates a new PaymentMethodResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodResponseWithDefaults() *PaymentMethodResponse {
	this := PaymentMethodResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaymentMethodResponse) GetLinks() PaginationLinks {
	if o == nil || common.IsNil(o.Links) {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || common.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaymentMethodResponse) HasLinks() bool {
	if o != nil && !common.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *PaymentMethodResponse) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PaymentMethodResponse) GetData() []PaymentMethod {
	if o == nil || common.IsNil(o.Data) {
		var ret []PaymentMethod
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetDataOk() ([]PaymentMethod, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PaymentMethodResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PaymentMethod and assigns it to the Data field.
func (o *PaymentMethodResponse) SetData(v []PaymentMethod) {
	o.Data = v
}

// GetItemsTotal returns the ItemsTotal field value
func (o *PaymentMethodResponse) GetItemsTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ItemsTotal
}

// GetItemsTotalOk returns a tuple with the ItemsTotal field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetItemsTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsTotal, true
}

// SetItemsTotal sets field value
func (o *PaymentMethodResponse) SetItemsTotal(v int32) {
	o.ItemsTotal = v
}

// GetPagesTotal returns the PagesTotal field value
func (o *PaymentMethodResponse) GetPagesTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PagesTotal
}

// GetPagesTotalOk returns a tuple with the PagesTotal field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetPagesTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PagesTotal, true
}

// SetPagesTotal sets field value
func (o *PaymentMethodResponse) SetPagesTotal(v int32) {
	o.PagesTotal = v
}

// GetTypesWithErrors returns the TypesWithErrors field value if set, zero value otherwise.
func (o *PaymentMethodResponse) GetTypesWithErrors() []string {
	if o == nil || common.IsNil(o.TypesWithErrors) {
		var ret []string
		return ret
	}
	return o.TypesWithErrors
}

// GetTypesWithErrorsOk returns a tuple with the TypesWithErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetTypesWithErrorsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.TypesWithErrors) {
		return nil, false
	}
	return o.TypesWithErrors, true
}

// HasTypesWithErrors returns a boolean if a field has been set.
func (o *PaymentMethodResponse) HasTypesWithErrors() bool {
	if o != nil && !common.IsNil(o.TypesWithErrors) {
		return true
	}

	return false
}

// SetTypesWithErrors gets a reference to the given []string and assigns it to the TypesWithErrors field.
func (o *PaymentMethodResponse) SetTypesWithErrors(v []string) {
	o.TypesWithErrors = v
}

func (o PaymentMethodResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["itemsTotal"] = o.ItemsTotal
	toSerialize["pagesTotal"] = o.PagesTotal
	if !common.IsNil(o.TypesWithErrors) {
		toSerialize["typesWithErrors"] = o.TypesWithErrors
	}
	return toSerialize, nil
}

type NullablePaymentMethodResponse struct {
	value *PaymentMethodResponse
	isSet bool
}

func (v NullablePaymentMethodResponse) Get() *PaymentMethodResponse {
	return v.value
}

func (v *NullablePaymentMethodResponse) Set(val *PaymentMethodResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodResponse(val *PaymentMethodResponse) *NullablePaymentMethodResponse {
	return &NullablePaymentMethodResponse{value: val, isSet: true}
}

func (v NullablePaymentMethodResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
