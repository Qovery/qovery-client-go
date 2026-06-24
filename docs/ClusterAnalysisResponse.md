# ClusterAnalysisResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ClusterId** | **string** |  | 
**OrganizationId** | **string** |  | 
**Kind** | [**ClusterAnalysisKind**](ClusterAnalysisKind.md) |  | 
**Status** | [**ClusterAnalysisStatus**](ClusterAnalysisStatus.md) |  | 
**OutputFormat** | [**ClusterAnalysisOutputFormat**](ClusterAnalysisOutputFormat.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**TriggeredBy** | **string** |  | 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewClusterAnalysisResponse

`func NewClusterAnalysisResponse(id string, clusterId string, organizationId string, kind ClusterAnalysisKind, status ClusterAnalysisStatus, outputFormat ClusterAnalysisOutputFormat, createdAt time.Time, updatedAt time.Time, triggeredBy string, ) *ClusterAnalysisResponse`

NewClusterAnalysisResponse instantiates a new ClusterAnalysisResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAnalysisResponseWithDefaults

`func NewClusterAnalysisResponseWithDefaults() *ClusterAnalysisResponse`

NewClusterAnalysisResponseWithDefaults instantiates a new ClusterAnalysisResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterAnalysisResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterAnalysisResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterAnalysisResponse) SetId(v string)`

SetId sets Id field to given value.


### GetClusterId

`func (o *ClusterAnalysisResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterAnalysisResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterAnalysisResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetOrganizationId

`func (o *ClusterAnalysisResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ClusterAnalysisResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ClusterAnalysisResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetKind

`func (o *ClusterAnalysisResponse) GetKind() ClusterAnalysisKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterAnalysisResponse) GetKindOk() (*ClusterAnalysisKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterAnalysisResponse) SetKind(v ClusterAnalysisKind)`

SetKind sets Kind field to given value.


### GetStatus

`func (o *ClusterAnalysisResponse) GetStatus() ClusterAnalysisStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterAnalysisResponse) GetStatusOk() (*ClusterAnalysisStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterAnalysisResponse) SetStatus(v ClusterAnalysisStatus)`

SetStatus sets Status field to given value.


### GetOutputFormat

`func (o *ClusterAnalysisResponse) GetOutputFormat() ClusterAnalysisOutputFormat`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *ClusterAnalysisResponse) GetOutputFormatOk() (*ClusterAnalysisOutputFormat, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *ClusterAnalysisResponse) SetOutputFormat(v ClusterAnalysisOutputFormat)`

SetOutputFormat sets OutputFormat field to given value.


### GetCreatedAt

`func (o *ClusterAnalysisResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterAnalysisResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterAnalysisResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ClusterAnalysisResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClusterAnalysisResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClusterAnalysisResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTriggeredBy

`func (o *ClusterAnalysisResponse) GetTriggeredBy() string`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *ClusterAnalysisResponse) GetTriggeredByOk() (*string, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *ClusterAnalysisResponse) SetTriggeredBy(v string)`

SetTriggeredBy sets TriggeredBy field to given value.


### GetErrorMessage

`func (o *ClusterAnalysisResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ClusterAnalysisResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ClusterAnalysisResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ClusterAnalysisResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ClusterAnalysisResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ClusterAnalysisResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


