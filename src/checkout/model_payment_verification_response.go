/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the PaymentVerificationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentVerificationResponse{}

// PaymentVerificationResponse struct for PaymentVerificationResponse
type PaymentVerificationResponse struct {
	// Contains additional information about the payment. Some data fields are included only if you select them first: Go to **Customer Area** > **Developers** > **Additional data**.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	FraudResult    *FraudResult       `json:"fraudResult,omitempty"`
	// A unique value that you provided in the initial `/paymentSession` request as a `reference` field.
	MerchantReference string                 `json:"merchantReference"`
	Order             *CheckoutOrderResponse `json:"order,omitempty"`
	// Adyen's 16-character reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request.
	PspReference *string `json:"pspReference,omitempty"`
	// If the payment's authorisation is refused or an error occurs during authorisation, this field holds Adyen's mapped reason for the refusal or a description of the error. When a transaction fails, the authorisation response includes `resultCode` and `refusalReason` values.  For more information, see [Refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons).
	RefusalReason *string `json:"refusalReason,omitempty"`
	// Code that specifies the refusal reason. For more information, see [Authorisation refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons).
	RefusalReasonCode *string `json:"refusalReasonCode,omitempty"`
	// The result of the payment. For more information, see [Result codes](https://docs.adyen.com/online-payments/payment-result-codes).  Possible values:  * **AuthenticationFinished** – The payment has been successfully authenticated with 3D Secure 2. Returned for 3D Secure 2 authentication-only transactions. * **AuthenticationNotRequired** – The transaction does not require 3D Secure authentication. Returned for [standalone authentication-only integrations](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only). * **Authorised** – The payment was successfully authorised. This state serves as an indicator to proceed with the delivery of goods and services. This is a final state. * **Cancelled** – Indicates the payment has been cancelled (either by the shopper or the merchant) before processing was completed. This is a final state. * **ChallengeShopper** – The issuer requires further shopper interaction before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **Error** – There was an error when the payment was being processed. The reason is given in the `refusalReason` field. This is a final state. * **IdentifyShopper** – The issuer requires the shopper's device fingerprint before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **PartiallyAuthorised** – The payment has been authorised for a partial amount. This happens for card payments when the merchant supports Partial Authorisations and the cardholder has insufficient funds. * **Pending** – Indicates that it is not possible to obtain the final status of the payment. This can happen if the systems providing final status information for the payment are unavailable, or if the shopper needs to take further action to complete the payment. * **PresentToShopper** – Indicates that the response contains additional information that you need to present to a shopper, so that they can use it to complete a payment. * **Received** – Indicates the payment has successfully been received by Adyen, and will be processed. This is the initial state for all payments. * **RedirectShopper** – Indicates the shopper should be redirected to an external web page or app to complete the authorisation. * **Refused** – Indicates the payment was refused. The reason is given in the `refusalReason` field. This is a final state.
	ResultCode   *string              `json:"resultCode,omitempty"`
	ServiceError *ServiceErrorDetails `json:"serviceError,omitempty"`
	// The shopperLocale value provided in the payment request.
	ShopperLocale string `json:"shopperLocale"`
}

// NewPaymentVerificationResponse instantiates a new PaymentVerificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentVerificationResponse(merchantReference string, shopperLocale string) *PaymentVerificationResponse {
	this := PaymentVerificationResponse{}
	this.MerchantReference = merchantReference
	this.ShopperLocale = shopperLocale
	return &this
}

// NewPaymentVerificationResponseWithDefaults instantiates a new PaymentVerificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentVerificationResponseWithDefaults() *PaymentVerificationResponse {
	this := PaymentVerificationResponse{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *PaymentVerificationResponse) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVerificationResponse) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *PaymentVerificationResponse) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *PaymentVerificationResponse) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetFraudResult returns the FraudResult field value if set, zero value otherwise.
func (o *PaymentVerificationResponse) GetFraudResult() FraudResult {
	if o == nil || common.IsNil(o.FraudResult) {
		var ret FraudResult
		return ret
	}
	return *o.FraudResult
}

// GetFraudResultOk returns a tuple with the FraudResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVerificationResponse) GetFraudResultOk() (*FraudResult, bool) {
	if o == nil || common.IsNil(o.FraudResult) {
		return nil, false
	}
	return o.FraudResult, true
}

// HasFraudResult returns a boolean if a field has been set.
func (o *PaymentVerificationResponse) HasFraudResult() bool {
	if o != nil && !common.IsNil(o.FraudResult) {
		return true
	}

	return false
}

// SetFraudResult gets a reference to the given FraudResult and assigns it to the FraudResult field.
func (o *PaymentVerificationResponse) SetFraudResult(v FraudResult) {
	o.FraudResult = &v
}

// GetMerchantReference returns the MerchantReference field value
func (o *PaymentVerificationResponse) GetMerchantReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantReference
}

// GetMerchantReferenceOk returns a tuple with the MerchantReference field value
// and a boolean to check if the value has been set.
func (o *PaymentVerificationResponse) GetMerchantReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantReference, true
}

// SetMerchantReference sets field value
func (o *PaymentVerificationResponse) SetMerchantReference(v string) {
	o.MerchantReference = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *PaymentVerificationResponse) GetOrder() CheckoutOrderResponse {
	if o == nil || common.IsNil(o.Order) {
		var ret CheckoutOrderResponse
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVerificationResponse) GetOrderOk() (*CheckoutOrderResponse, bool) {
	if o == nil || common.IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *PaymentVerificationResponse) HasOrder() bool {
	if o != nil && !common.IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given CheckoutOrderResponse and assigns it to the Order field.
func (o *PaymentVerificationResponse) SetOrder(v CheckoutOrderResponse) {
	o.Order = &v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *PaymentVerificationResponse) GetPspReference() string {
	if o == nil || common.IsNil(o.PspReference) {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVerificationResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PspReference) {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *PaymentVerificationResponse) HasPspReference() bool {
	if o != nil && !common.IsNil(o.PspReference) {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *PaymentVerificationResponse) SetPspReference(v string) {
	o.PspReference = &v
}

// GetRefusalReason returns the RefusalReason field value if set, zero value otherwise.
func (o *PaymentVerificationResponse) GetRefusalReason() string {
	if o == nil || common.IsNil(o.RefusalReason) {
		var ret string
		return ret
	}
	return *o.RefusalReason
}

// GetRefusalReasonOk returns a tuple with the RefusalReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVerificationResponse) GetRefusalReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.RefusalReason) {
		return nil, false
	}
	return o.RefusalReason, true
}

// HasRefusalReason returns a boolean if a field has been set.
func (o *PaymentVerificationResponse) HasRefusalReason() bool {
	if o != nil && !common.IsNil(o.RefusalReason) {
		return true
	}

	return false
}

// SetRefusalReason gets a reference to the given string and assigns it to the RefusalReason field.
func (o *PaymentVerificationResponse) SetRefusalReason(v string) {
	o.RefusalReason = &v
}

// GetRefusalReasonCode returns the RefusalReasonCode field value if set, zero value otherwise.
func (o *PaymentVerificationResponse) GetRefusalReasonCode() string {
	if o == nil || common.IsNil(o.RefusalReasonCode) {
		var ret string
		return ret
	}
	return *o.RefusalReasonCode
}

// GetRefusalReasonCodeOk returns a tuple with the RefusalReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVerificationResponse) GetRefusalReasonCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RefusalReasonCode) {
		return nil, false
	}
	return o.RefusalReasonCode, true
}

// HasRefusalReasonCode returns a boolean if a field has been set.
func (o *PaymentVerificationResponse) HasRefusalReasonCode() bool {
	if o != nil && !common.IsNil(o.RefusalReasonCode) {
		return true
	}

	return false
}

// SetRefusalReasonCode gets a reference to the given string and assigns it to the RefusalReasonCode field.
func (o *PaymentVerificationResponse) SetRefusalReasonCode(v string) {
	o.RefusalReasonCode = &v
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *PaymentVerificationResponse) GetResultCode() string {
	if o == nil || common.IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVerificationResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *PaymentVerificationResponse) HasResultCode() bool {
	if o != nil && !common.IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *PaymentVerificationResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetServiceError returns the ServiceError field value if set, zero value otherwise.
func (o *PaymentVerificationResponse) GetServiceError() ServiceErrorDetails {
	if o == nil || common.IsNil(o.ServiceError) {
		var ret ServiceErrorDetails
		return ret
	}
	return *o.ServiceError
}

// GetServiceErrorOk returns a tuple with the ServiceError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVerificationResponse) GetServiceErrorOk() (*ServiceErrorDetails, bool) {
	if o == nil || common.IsNil(o.ServiceError) {
		return nil, false
	}
	return o.ServiceError, true
}

// HasServiceError returns a boolean if a field has been set.
func (o *PaymentVerificationResponse) HasServiceError() bool {
	if o != nil && !common.IsNil(o.ServiceError) {
		return true
	}

	return false
}

// SetServiceError gets a reference to the given ServiceErrorDetails and assigns it to the ServiceError field.
func (o *PaymentVerificationResponse) SetServiceError(v ServiceErrorDetails) {
	o.ServiceError = &v
}

// GetShopperLocale returns the ShopperLocale field value
func (o *PaymentVerificationResponse) GetShopperLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperLocale
}

// GetShopperLocaleOk returns a tuple with the ShopperLocale field value
// and a boolean to check if the value has been set.
func (o *PaymentVerificationResponse) GetShopperLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperLocale, true
}

// SetShopperLocale sets field value
func (o *PaymentVerificationResponse) SetShopperLocale(v string) {
	o.ShopperLocale = v
}

func (o PaymentVerificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentVerificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if !common.IsNil(o.FraudResult) {
		toSerialize["fraudResult"] = o.FraudResult
	}
	toSerialize["merchantReference"] = o.MerchantReference
	if !common.IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !common.IsNil(o.PspReference) {
		toSerialize["pspReference"] = o.PspReference
	}
	if !common.IsNil(o.RefusalReason) {
		toSerialize["refusalReason"] = o.RefusalReason
	}
	if !common.IsNil(o.RefusalReasonCode) {
		toSerialize["refusalReasonCode"] = o.RefusalReasonCode
	}
	if !common.IsNil(o.ResultCode) {
		toSerialize["resultCode"] = o.ResultCode
	}
	if !common.IsNil(o.ServiceError) {
		toSerialize["serviceError"] = o.ServiceError
	}
	toSerialize["shopperLocale"] = o.ShopperLocale
	return toSerialize, nil
}

type NullablePaymentVerificationResponse struct {
	value *PaymentVerificationResponse
	isSet bool
}

func (v NullablePaymentVerificationResponse) Get() *PaymentVerificationResponse {
	return v.value
}

func (v *NullablePaymentVerificationResponse) Set(val *PaymentVerificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentVerificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentVerificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentVerificationResponse(val *PaymentVerificationResponse) *NullablePaymentVerificationResponse {
	return &NullablePaymentVerificationResponse{value: val, isSet: true}
}

func (v NullablePaymentVerificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentVerificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentVerificationResponse) isValidResultCode() bool {
	var allowedEnumValues = []string{"AuthenticationFinished", "AuthenticationNotRequired", "Authorised", "Cancelled", "ChallengeShopper", "Error", "IdentifyShopper", "PartiallyAuthorised", "Pending", "PresentToShopper", "Received", "RedirectShopper", "Refused", "Success"}
	for _, allowed := range allowedEnumValues {
		if o.GetResultCode() == allowed {
			return true
		}
	}
	return false
}
