# ClusterOverviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**CloudProvider** | [**CloudVendorEnum**](CloudVendorEnum.md) |  | 
**IsDemo** | **bool** |  | 

## Methods

### NewClusterOverviewResponse

`func NewClusterOverviewResponse(id string, name string, cloudProvider CloudVendorEnum, isDemo bool, ) *ClusterOverviewResponse`

NewClusterOverviewResponse instantiates a new ClusterOverviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterOverviewResponseWithDefaults

`func NewClusterOverviewResponseWithDefaults() *ClusterOverviewResponse`

NewClusterOverviewResponseWithDefaults instantiates a new ClusterOverviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterOverviewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterOverviewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterOverviewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ClusterOverviewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterOverviewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterOverviewResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCloudProvider

`func (o *ClusterOverviewResponse) GetCloudProvider() CloudVendorEnum`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ClusterOverviewResponse) GetCloudProviderOk() (*CloudVendorEnum, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ClusterOverviewResponse) SetCloudProvider(v CloudVendorEnum)`

SetCloudProvider sets CloudProvider field to given value.


### GetIsDemo

`func (o *ClusterOverviewResponse) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *ClusterOverviewResponse) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *ClusterOverviewResponse) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


