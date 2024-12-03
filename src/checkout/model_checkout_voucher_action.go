/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the CheckoutVoucherAction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutVoucherAction{}

// CheckoutVoucherAction struct for CheckoutVoucherAction
type CheckoutVoucherAction struct {
	// The voucher alternative reference code.
	AlternativeReference *string `json:"alternativeReference,omitempty"`
	// A collection institution number (store number) for Econtext Pay-Easy ATM.
	CollectionInstitutionNumber *string `json:"collectionInstitutionNumber,omitempty"`
	// The URL to download the voucher.
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	// An entity number of Multibanco.
	Entity *string `json:"entity,omitempty"`
	// The date time of the voucher expiry.
	ExpiresAt     *string `json:"expiresAt,omitempty"`
	InitialAmount *Amount `json:"initialAmount,omitempty"`
	// The URL to the detailed instructions to make payment using the voucher.
	InstructionsUrl *string `json:"instructionsUrl,omitempty"`
	// The issuer of the voucher.
	Issuer *string `json:"issuer,omitempty"`
	// The shopper telephone number (partially masked).
	MaskedTelephoneNumber *string `json:"maskedTelephoneNumber,omitempty"`
	// The merchant name.
	MerchantName *string `json:"merchantName,omitempty"`
	// The merchant reference.
	MerchantReference *string `json:"merchantReference,omitempty"`
	// A Base64-encoded token containing all properties of the voucher. For iOS, you can use this to pass a voucher to Apple Wallet.
	PassCreationToken *string `json:"passCreationToken,omitempty"`
	// Encoded payment data.
	PaymentData *string `json:"paymentData,omitempty"`
	// Specifies the payment method.
	PaymentMethodType *string `json:"paymentMethodType,omitempty"`
	// The voucher reference code.
	Reference *string `json:"reference,omitempty"`
	// The shopper email.
	ShopperEmail *string `json:"shopperEmail,omitempty"`
	// The shopper name.
	ShopperName *string `json:"shopperName,omitempty"`
	Surcharge   *Amount `json:"surcharge,omitempty"`
	TotalAmount *Amount `json:"totalAmount,omitempty"`
	// **voucher**
	Type string `json:"type"`
	// Specifies the URL to redirect to.
	Url *string `json:"url,omitempty"`
}

// NewCheckoutVoucherAction instantiates a new CheckoutVoucherAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutVoucherAction(type_ string) *CheckoutVoucherAction {
	this := CheckoutVoucherAction{}
	this.Type = type_
	return &this
}

// NewCheckoutVoucherActionWithDefaults instantiates a new CheckoutVoucherAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutVoucherActionWithDefaults() *CheckoutVoucherAction {
	this := CheckoutVoucherAction{}
	return &this
}

// GetAlternativeReference returns the AlternativeReference field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetAlternativeReference() string {
	if o == nil || common.IsNil(o.AlternativeReference) {
		var ret string
		return ret
	}
	return *o.AlternativeReference
}

// GetAlternativeReferenceOk returns a tuple with the AlternativeReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetAlternativeReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AlternativeReference) {
		return nil, false
	}
	return o.AlternativeReference, true
}

// HasAlternativeReference returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasAlternativeReference() bool {
	if o != nil && !common.IsNil(o.AlternativeReference) {
		return true
	}

	return false
}

// SetAlternativeReference gets a reference to the given string and assigns it to the AlternativeReference field.
func (o *CheckoutVoucherAction) SetAlternativeReference(v string) {
	o.AlternativeReference = &v
}

// GetCollectionInstitutionNumber returns the CollectionInstitutionNumber field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetCollectionInstitutionNumber() string {
	if o == nil || common.IsNil(o.CollectionInstitutionNumber) {
		var ret string
		return ret
	}
	return *o.CollectionInstitutionNumber
}

// GetCollectionInstitutionNumberOk returns a tuple with the CollectionInstitutionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetCollectionInstitutionNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollectionInstitutionNumber) {
		return nil, false
	}
	return o.CollectionInstitutionNumber, true
}

// HasCollectionInstitutionNumber returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasCollectionInstitutionNumber() bool {
	if o != nil && !common.IsNil(o.CollectionInstitutionNumber) {
		return true
	}

	return false
}

// SetCollectionInstitutionNumber gets a reference to the given string and assigns it to the CollectionInstitutionNumber field.
func (o *CheckoutVoucherAction) SetCollectionInstitutionNumber(v string) {
	o.CollectionInstitutionNumber = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetDownloadUrl() string {
	if o == nil || common.IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetDownloadUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.DownloadUrl) {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasDownloadUrl() bool {
	if o != nil && !common.IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *CheckoutVoucherAction) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetEntity() string {
	if o == nil || common.IsNil(o.Entity) {
		var ret string
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetEntityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasEntity() bool {
	if o != nil && !common.IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given string and assigns it to the Entity field.
func (o *CheckoutVoucherAction) SetEntity(v string) {
	o.Entity = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetExpiresAt() string {
	if o == nil || common.IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetExpiresAtOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasExpiresAt() bool {
	if o != nil && !common.IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *CheckoutVoucherAction) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetInitialAmount returns the InitialAmount field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetInitialAmount() Amount {
	if o == nil || common.IsNil(o.InitialAmount) {
		var ret Amount
		return ret
	}
	return *o.InitialAmount
}

// GetInitialAmountOk returns a tuple with the InitialAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetInitialAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.InitialAmount) {
		return nil, false
	}
	return o.InitialAmount, true
}

// HasInitialAmount returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasInitialAmount() bool {
	if o != nil && !common.IsNil(o.InitialAmount) {
		return true
	}

	return false
}

// SetInitialAmount gets a reference to the given Amount and assigns it to the InitialAmount field.
func (o *CheckoutVoucherAction) SetInitialAmount(v Amount) {
	o.InitialAmount = &v
}

// GetInstructionsUrl returns the InstructionsUrl field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetInstructionsUrl() string {
	if o == nil || common.IsNil(o.InstructionsUrl) {
		var ret string
		return ret
	}
	return *o.InstructionsUrl
}

// GetInstructionsUrlOk returns a tuple with the InstructionsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetInstructionsUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstructionsUrl) {
		return nil, false
	}
	return o.InstructionsUrl, true
}

// HasInstructionsUrl returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasInstructionsUrl() bool {
	if o != nil && !common.IsNil(o.InstructionsUrl) {
		return true
	}

	return false
}

// SetInstructionsUrl gets a reference to the given string and assigns it to the InstructionsUrl field.
func (o *CheckoutVoucherAction) SetInstructionsUrl(v string) {
	o.InstructionsUrl = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetIssuer() string {
	if o == nil || common.IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetIssuerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasIssuer() bool {
	if o != nil && !common.IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *CheckoutVoucherAction) SetIssuer(v string) {
	o.Issuer = &v
}

// GetMaskedTelephoneNumber returns the MaskedTelephoneNumber field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetMaskedTelephoneNumber() string {
	if o == nil || common.IsNil(o.MaskedTelephoneNumber) {
		var ret string
		return ret
	}
	return *o.MaskedTelephoneNumber
}

// GetMaskedTelephoneNumberOk returns a tuple with the MaskedTelephoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetMaskedTelephoneNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaskedTelephoneNumber) {
		return nil, false
	}
	return o.MaskedTelephoneNumber, true
}

// HasMaskedTelephoneNumber returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasMaskedTelephoneNumber() bool {
	if o != nil && !common.IsNil(o.MaskedTelephoneNumber) {
		return true
	}

	return false
}

// SetMaskedTelephoneNumber gets a reference to the given string and assigns it to the MaskedTelephoneNumber field.
func (o *CheckoutVoucherAction) SetMaskedTelephoneNumber(v string) {
	o.MaskedTelephoneNumber = &v
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetMerchantName() string {
	if o == nil || common.IsNil(o.MerchantName) {
		var ret string
		return ret
	}
	return *o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetMerchantNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantName) {
		return nil, false
	}
	return o.MerchantName, true
}

// HasMerchantName returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasMerchantName() bool {
	if o != nil && !common.IsNil(o.MerchantName) {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given string and assigns it to the MerchantName field.
func (o *CheckoutVoucherAction) SetMerchantName(v string) {
	o.MerchantName = &v
}

// GetMerchantReference returns the MerchantReference field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetMerchantReference() string {
	if o == nil || common.IsNil(o.MerchantReference) {
		var ret string
		return ret
	}
	return *o.MerchantReference
}

// GetMerchantReferenceOk returns a tuple with the MerchantReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetMerchantReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantReference) {
		return nil, false
	}
	return o.MerchantReference, true
}

// HasMerchantReference returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasMerchantReference() bool {
	if o != nil && !common.IsNil(o.MerchantReference) {
		return true
	}

	return false
}

// SetMerchantReference gets a reference to the given string and assigns it to the MerchantReference field.
func (o *CheckoutVoucherAction) SetMerchantReference(v string) {
	o.MerchantReference = &v
}

// GetPassCreationToken returns the PassCreationToken field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetPassCreationToken() string {
	if o == nil || common.IsNil(o.PassCreationToken) {
		var ret string
		return ret
	}
	return *o.PassCreationToken
}

// GetPassCreationTokenOk returns a tuple with the PassCreationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetPassCreationTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.PassCreationToken) {
		return nil, false
	}
	return o.PassCreationToken, true
}

// HasPassCreationToken returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasPassCreationToken() bool {
	if o != nil && !common.IsNil(o.PassCreationToken) {
		return true
	}

	return false
}

// SetPassCreationToken gets a reference to the given string and assigns it to the PassCreationToken field.
func (o *CheckoutVoucherAction) SetPassCreationToken(v string) {
	o.PassCreationToken = &v
}

// GetPaymentData returns the PaymentData field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetPaymentData() string {
	if o == nil || common.IsNil(o.PaymentData) {
		var ret string
		return ret
	}
	return *o.PaymentData
}

// GetPaymentDataOk returns a tuple with the PaymentData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetPaymentDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentData) {
		return nil, false
	}
	return o.PaymentData, true
}

// HasPaymentData returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasPaymentData() bool {
	if o != nil && !common.IsNil(o.PaymentData) {
		return true
	}

	return false
}

// SetPaymentData gets a reference to the given string and assigns it to the PaymentData field.
func (o *CheckoutVoucherAction) SetPaymentData(v string) {
	o.PaymentData = &v
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetPaymentMethodType() string {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		var ret string
		return ret
	}
	return *o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		return nil, false
	}
	return o.PaymentMethodType, true
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasPaymentMethodType() bool {
	if o != nil && !common.IsNil(o.PaymentMethodType) {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given string and assigns it to the PaymentMethodType field.
func (o *CheckoutVoucherAction) SetPaymentMethodType(v string) {
	o.PaymentMethodType = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *CheckoutVoucherAction) SetReference(v string) {
	o.Reference = &v
}

// GetShopperEmail returns the ShopperEmail field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetShopperEmail() string {
	if o == nil || common.IsNil(o.ShopperEmail) {
		var ret string
		return ret
	}
	return *o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetShopperEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperEmail) {
		return nil, false
	}
	return o.ShopperEmail, true
}

// HasShopperEmail returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasShopperEmail() bool {
	if o != nil && !common.IsNil(o.ShopperEmail) {
		return true
	}

	return false
}

// SetShopperEmail gets a reference to the given string and assigns it to the ShopperEmail field.
func (o *CheckoutVoucherAction) SetShopperEmail(v string) {
	o.ShopperEmail = &v
}

// GetShopperName returns the ShopperName field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetShopperName() string {
	if o == nil || common.IsNil(o.ShopperName) {
		var ret string
		return ret
	}
	return *o.ShopperName
}

// GetShopperNameOk returns a tuple with the ShopperName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetShopperNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperName) {
		return nil, false
	}
	return o.ShopperName, true
}

// HasShopperName returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasShopperName() bool {
	if o != nil && !common.IsNil(o.ShopperName) {
		return true
	}

	return false
}

// SetShopperName gets a reference to the given string and assigns it to the ShopperName field.
func (o *CheckoutVoucherAction) SetShopperName(v string) {
	o.ShopperName = &v
}

// GetSurcharge returns the Surcharge field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetSurcharge() Amount {
	if o == nil || common.IsNil(o.Surcharge) {
		var ret Amount
		return ret
	}
	return *o.Surcharge
}

// GetSurchargeOk returns a tuple with the Surcharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetSurchargeOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Surcharge) {
		return nil, false
	}
	return o.Surcharge, true
}

// HasSurcharge returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasSurcharge() bool {
	if o != nil && !common.IsNil(o.Surcharge) {
		return true
	}

	return false
}

// SetSurcharge gets a reference to the given Amount and assigns it to the Surcharge field.
func (o *CheckoutVoucherAction) SetSurcharge(v Amount) {
	o.Surcharge = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetTotalAmount() Amount {
	if o == nil || common.IsNil(o.TotalAmount) {
		var ret Amount
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetTotalAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasTotalAmount() bool {
	if o != nil && !common.IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given Amount and assigns it to the TotalAmount field.
func (o *CheckoutVoucherAction) SetTotalAmount(v Amount) {
	o.TotalAmount = &v
}

// GetType returns the Type field value
func (o *CheckoutVoucherAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutVoucherAction) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CheckoutVoucherAction) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutVoucherAction) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CheckoutVoucherAction) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CheckoutVoucherAction) SetUrl(v string) {
	o.Url = &v
}

func (o CheckoutVoucherAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutVoucherAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AlternativeReference) {
		toSerialize["alternativeReference"] = o.AlternativeReference
	}
	if !common.IsNil(o.CollectionInstitutionNumber) {
		toSerialize["collectionInstitutionNumber"] = o.CollectionInstitutionNumber
	}
	if !common.IsNil(o.DownloadUrl) {
		toSerialize["downloadUrl"] = o.DownloadUrl
	}
	if !common.IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !common.IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !common.IsNil(o.InitialAmount) {
		toSerialize["initialAmount"] = o.InitialAmount
	}
	if !common.IsNil(o.InstructionsUrl) {
		toSerialize["instructionsUrl"] = o.InstructionsUrl
	}
	if !common.IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !common.IsNil(o.MaskedTelephoneNumber) {
		toSerialize["maskedTelephoneNumber"] = o.MaskedTelephoneNumber
	}
	if !common.IsNil(o.MerchantName) {
		toSerialize["merchantName"] = o.MerchantName
	}
	if !common.IsNil(o.MerchantReference) {
		toSerialize["merchantReference"] = o.MerchantReference
	}
	if !common.IsNil(o.PassCreationToken) {
		toSerialize["passCreationToken"] = o.PassCreationToken
	}
	if !common.IsNil(o.PaymentData) {
		toSerialize["paymentData"] = o.PaymentData
	}
	if !common.IsNil(o.PaymentMethodType) {
		toSerialize["paymentMethodType"] = o.PaymentMethodType
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.ShopperEmail) {
		toSerialize["shopperEmail"] = o.ShopperEmail
	}
	if !common.IsNil(o.ShopperName) {
		toSerialize["shopperName"] = o.ShopperName
	}
	if !common.IsNil(o.Surcharge) {
		toSerialize["surcharge"] = o.Surcharge
	}
	if !common.IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	toSerialize["type"] = o.Type
	if !common.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableCheckoutVoucherAction struct {
	value *CheckoutVoucherAction
	isSet bool
}

func (v NullableCheckoutVoucherAction) Get() *CheckoutVoucherAction {
	return v.value
}

func (v *NullableCheckoutVoucherAction) Set(val *CheckoutVoucherAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutVoucherAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutVoucherAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutVoucherAction(val *CheckoutVoucherAction) *NullableCheckoutVoucherAction {
	return &NullableCheckoutVoucherAction{value: val, isSet: true}
}

func (v NullableCheckoutVoucherAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutVoucherAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CheckoutVoucherAction) isValidType() bool {
	var allowedEnumValues = []string{"voucher"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
