# GetClusterTokenByClusterId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Spec** | **map[string]interface{}** |  | 
**Status** | [**GetClusterTokenByClusterId200ResponseStatus**](GetClusterTokenByClusterId200ResponseStatus.md) |  | 

## Methods

### NewGetClusterTokenByClusterId200Response

`func NewGetClusterTokenByClusterId200Response(apiVersion string, kind string, spec map[string]interface{}, status GetClusterTokenByClusterId200ResponseStatus, ) *GetClusterTokenByClusterId200Response`

NewGetClusterTokenByClusterId200Response instantiates a new GetClusterTokenByClusterId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterTokenByClusterId200ResponseWithDefaults

`func NewGetClusterTokenByClusterId200ResponseWithDefaults() *GetClusterTokenByClusterId200Response`

NewGetClusterTokenByClusterId200ResponseWithDefaults instantiates a new GetClusterTokenByClusterId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetClusterTokenByClusterId200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetClusterTokenByClusterId200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetClusterTokenByClusterId200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetClusterTokenByClusterId200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetClusterTokenByClusterId200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetClusterTokenByClusterId200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSpec

`func (o *GetClusterTokenByClusterId200Response) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *GetClusterTokenByClusterId200Response) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *GetClusterTokenByClusterId200Response) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *GetClusterTokenByClusterId200Response) GetStatus() GetClusterTokenByClusterId200ResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetClusterTokenByClusterId200Response) GetStatusOk() (*GetClusterTokenByClusterId200ResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetClusterTokenByClusterId200Response) SetStatus(v GetClusterTokenByClusterId200ResponseStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


