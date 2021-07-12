# BillingInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Email** | **string** | email used for billing, and to receive all invoices by email | 
**Address** | **string** |  | 
**City** | **string** |  | 
**Zip** | **string** |  | 
**State** | Pointer to **string** | only for US | [optional] 
**CountryCode** | **string** | ISO code of the country | 
**Company** | Pointer to **string** | name of the company to bill | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingInfoResponse

`func NewBillingInfoResponse(firstName string, lastName string, email string, address string, city string, zip string, countryCode string, ) *BillingInfoResponse`

NewBillingInfoResponse instantiates a new BillingInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInfoResponseWithDefaults

`func NewBillingInfoResponseWithDefaults() *BillingInfoResponse`

NewBillingInfoResponseWithDefaults instantiates a new BillingInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *BillingInfoResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BillingInfoResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BillingInfoResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *BillingInfoResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BillingInfoResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BillingInfoResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *BillingInfoResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BillingInfoResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BillingInfoResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAddress

`func (o *BillingInfoResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BillingInfoResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BillingInfoResponse) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCity

`func (o *BillingInfoResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BillingInfoResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BillingInfoResponse) SetCity(v string)`

SetCity sets City field to given value.


### GetZip

`func (o *BillingInfoResponse) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *BillingInfoResponse) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *BillingInfoResponse) SetZip(v string)`

SetZip sets Zip field to given value.


### GetState

`func (o *BillingInfoResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BillingInfoResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BillingInfoResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BillingInfoResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountryCode

`func (o *BillingInfoResponse) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *BillingInfoResponse) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *BillingInfoResponse) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetCompany

`func (o *BillingInfoResponse) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *BillingInfoResponse) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *BillingInfoResponse) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *BillingInfoResponse) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetVatNumber

`func (o *BillingInfoResponse) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *BillingInfoResponse) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *BillingInfoResponse) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *BillingInfoResponse) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


