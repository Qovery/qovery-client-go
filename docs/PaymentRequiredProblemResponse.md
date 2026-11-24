# PaymentRequiredProblemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Always sent, &#x60;about:blank&#x60;. | [optional] 
**Title** | **string** | Always sent, &#x60;Payment Required&#x60;. | 
**Status** | Pointer to **int32** | Always sent, 402. | [optional] 
**Detail** | **string** | Always sent. Plain text to show to the user, saying why the action is refused and what to do: the pricing explanation, or a generic text when there is none. | 
**Instance** | Pointer to **string** | Always sent. The path of the request. | [optional] 
**Message** | Pointer to **string** | The same text as &#x60;detail&#x60;. Absent when the refusal carries no pricing explanation. | [optional] 
**Sku** | Pointer to **string** | The plan item that refused the action. Absent when unknown. New items can appear: do not reject an unknown value. | [optional] 
**Action** | Pointer to **string** | The checked action: &#x60;member.invite&#x60;, &#x60;cluster.create&#x60;, &#x60;cluster.create_self_managed&#x60; or &#x60;environment.create&#x60; (also used for a clone). Absent when unknown. New actions can appear: do not reject an unknown value. | [optional] 
**Limit** | Pointer to [**PaymentRequiredProblemResponseLimit**](PaymentRequiredProblemResponseLimit.md) |  | [optional] 
**Used** | Pointer to **int32** | How many of &#x60;sku&#x60; the organization uses now. Absent when unknown. | [optional] 
**Requested** | Pointer to **int32** | How many of &#x60;sku&#x60; the action adds. Absent when unknown. | [optional] 
**UpgradeUrl** | Pointer to **string** | The console page where a billing role changes the plan, the spend limit or the payment method. Always an &#x60;https&#x60; URL. Absent when there is no page to send the user to. | [optional] 
**Offers** | Pointer to **[]map[string]interface{}** | Ways past the refusal, produced by the pricing service, in the order to show them. Each offer has a &#x60;type&#x60; (&#x60;plan_change&#x60;, &#x60;addon&#x60;, &#x60;contact_sales&#x60;, &#x60;free_up&#x60; or &#x60;raise_spend_limit&#x60;; more can appear) and a &#x60;name&#x60;; its other properties depend on the type and can change. Absent unless the caller can manage the organization&#39;s billing. | [optional] 
**Paywall** | Pointer to **map[string]interface{}** | What the refused feature does and the ways to get it, produced by the pricing service for a billing role; its properties can change. Absent unless the caller can manage the organization&#39;s billing, and absent when the refusal comes from a spend limit. | [optional] 

## Methods

### NewPaymentRequiredProblemResponse

`func NewPaymentRequiredProblemResponse(title string, detail string, ) *PaymentRequiredProblemResponse`

NewPaymentRequiredProblemResponse instantiates a new PaymentRequiredProblemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequiredProblemResponseWithDefaults

`func NewPaymentRequiredProblemResponseWithDefaults() *PaymentRequiredProblemResponse`

NewPaymentRequiredProblemResponseWithDefaults instantiates a new PaymentRequiredProblemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentRequiredProblemResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentRequiredProblemResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentRequiredProblemResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentRequiredProblemResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *PaymentRequiredProblemResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PaymentRequiredProblemResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PaymentRequiredProblemResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *PaymentRequiredProblemResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentRequiredProblemResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentRequiredProblemResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentRequiredProblemResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *PaymentRequiredProblemResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *PaymentRequiredProblemResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *PaymentRequiredProblemResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetInstance

`func (o *PaymentRequiredProblemResponse) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *PaymentRequiredProblemResponse) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *PaymentRequiredProblemResponse) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *PaymentRequiredProblemResponse) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetMessage

`func (o *PaymentRequiredProblemResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentRequiredProblemResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentRequiredProblemResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PaymentRequiredProblemResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSku

`func (o *PaymentRequiredProblemResponse) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *PaymentRequiredProblemResponse) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *PaymentRequiredProblemResponse) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *PaymentRequiredProblemResponse) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetAction

`func (o *PaymentRequiredProblemResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PaymentRequiredProblemResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PaymentRequiredProblemResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PaymentRequiredProblemResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetLimit

`func (o *PaymentRequiredProblemResponse) GetLimit() PaymentRequiredProblemResponseLimit`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaymentRequiredProblemResponse) GetLimitOk() (*PaymentRequiredProblemResponseLimit, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaymentRequiredProblemResponse) SetLimit(v PaymentRequiredProblemResponseLimit)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaymentRequiredProblemResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetUsed

`func (o *PaymentRequiredProblemResponse) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *PaymentRequiredProblemResponse) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *PaymentRequiredProblemResponse) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *PaymentRequiredProblemResponse) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetRequested

`func (o *PaymentRequiredProblemResponse) GetRequested() int32`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *PaymentRequiredProblemResponse) GetRequestedOk() (*int32, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *PaymentRequiredProblemResponse) SetRequested(v int32)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *PaymentRequiredProblemResponse) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetUpgradeUrl

`func (o *PaymentRequiredProblemResponse) GetUpgradeUrl() string`

GetUpgradeUrl returns the UpgradeUrl field if non-nil, zero value otherwise.

### GetUpgradeUrlOk

`func (o *PaymentRequiredProblemResponse) GetUpgradeUrlOk() (*string, bool)`

GetUpgradeUrlOk returns a tuple with the UpgradeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeUrl

`func (o *PaymentRequiredProblemResponse) SetUpgradeUrl(v string)`

SetUpgradeUrl sets UpgradeUrl field to given value.

### HasUpgradeUrl

`func (o *PaymentRequiredProblemResponse) HasUpgradeUrl() bool`

HasUpgradeUrl returns a boolean if a field has been set.

### GetOffers

`func (o *PaymentRequiredProblemResponse) GetOffers() []map[string]interface{}`

GetOffers returns the Offers field if non-nil, zero value otherwise.

### GetOffersOk

`func (o *PaymentRequiredProblemResponse) GetOffersOk() (*[]map[string]interface{}, bool)`

GetOffersOk returns a tuple with the Offers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffers

`func (o *PaymentRequiredProblemResponse) SetOffers(v []map[string]interface{})`

SetOffers sets Offers field to given value.

### HasOffers

`func (o *PaymentRequiredProblemResponse) HasOffers() bool`

HasOffers returns a boolean if a field has been set.

### GetPaywall

`func (o *PaymentRequiredProblemResponse) GetPaywall() map[string]interface{}`

GetPaywall returns the Paywall field if non-nil, zero value otherwise.

### GetPaywallOk

`func (o *PaymentRequiredProblemResponse) GetPaywallOk() (*map[string]interface{}, bool)`

GetPaywallOk returns a tuple with the Paywall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaywall

`func (o *PaymentRequiredProblemResponse) SetPaywall(v map[string]interface{})`

SetPaywall sets Paywall field to given value.

### HasPaywall

`func (o *PaymentRequiredProblemResponse) HasPaywall() bool`

HasPaywall returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


