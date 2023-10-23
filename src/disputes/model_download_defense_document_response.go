/*
Disputes API

API version: 30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputes

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the DownloadDefenseDocumentResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DownloadDefenseDocumentResponse{}

// DownloadDefenseDocumentResponse struct for DownloadDefenseDocumentResponse
type DownloadDefenseDocumentResponse struct {
	// The content of the defense document in Base64 binary format. Must be encoded in the format that is specified in the `contentType` field.
	Content *string `json:"content,omitempty"`
	// The content type of the dispute defense document.  Possible values:   * **image/jpg**  * **image/jpeg**  * **image/tiff**  * **application/pdf**
	ContentType          *string              `json:"contentType,omitempty"`
	DisputeServiceResult DisputeServiceResult `json:"disputeServiceResult"`
}

// NewDownloadDefenseDocumentResponse instantiates a new DownloadDefenseDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadDefenseDocumentResponse(disputeServiceResult DisputeServiceResult) *DownloadDefenseDocumentResponse {
	this := DownloadDefenseDocumentResponse{}
	this.DisputeServiceResult = disputeServiceResult
	return &this
}

// NewDownloadDefenseDocumentResponseWithDefaults instantiates a new DownloadDefenseDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadDefenseDocumentResponseWithDefaults() *DownloadDefenseDocumentResponse {
	this := DownloadDefenseDocumentResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *DownloadDefenseDocumentResponse) GetContent() string {
	if o == nil || common.IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadDefenseDocumentResponse) GetContentOk() (*string, bool) {
	if o == nil || common.IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *DownloadDefenseDocumentResponse) HasContent() bool {
	if o != nil && !common.IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *DownloadDefenseDocumentResponse) SetContent(v string) {
	o.Content = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *DownloadDefenseDocumentResponse) GetContentType() string {
	if o == nil || common.IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadDefenseDocumentResponse) GetContentTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *DownloadDefenseDocumentResponse) HasContentType() bool {
	if o != nil && !common.IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *DownloadDefenseDocumentResponse) SetContentType(v string) {
	o.ContentType = &v
}

// GetDisputeServiceResult returns the DisputeServiceResult field value
func (o *DownloadDefenseDocumentResponse) GetDisputeServiceResult() DisputeServiceResult {
	if o == nil {
		var ret DisputeServiceResult
		return ret
	}

	return o.DisputeServiceResult
}

// GetDisputeServiceResultOk returns a tuple with the DisputeServiceResult field value
// and a boolean to check if the value has been set.
func (o *DownloadDefenseDocumentResponse) GetDisputeServiceResultOk() (*DisputeServiceResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisputeServiceResult, true
}

// SetDisputeServiceResult sets field value
func (o *DownloadDefenseDocumentResponse) SetDisputeServiceResult(v DisputeServiceResult) {
	o.DisputeServiceResult = v
}

func (o DownloadDefenseDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadDefenseDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !common.IsNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	toSerialize["disputeServiceResult"] = o.DisputeServiceResult
	return toSerialize, nil
}

type NullableDownloadDefenseDocumentResponse struct {
	value *DownloadDefenseDocumentResponse
	isSet bool
}

func (v NullableDownloadDefenseDocumentResponse) Get() *DownloadDefenseDocumentResponse {
	return v.value
}

func (v *NullableDownloadDefenseDocumentResponse) Set(val *DownloadDefenseDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadDefenseDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadDefenseDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadDefenseDocumentResponse(val *DownloadDefenseDocumentResponse) *NullableDownloadDefenseDocumentResponse {
	return &NullableDownloadDefenseDocumentResponse{value: val, isSet: true}
}

func (v NullableDownloadDefenseDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadDefenseDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
