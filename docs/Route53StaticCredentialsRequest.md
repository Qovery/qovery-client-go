# Route53StaticCredentialsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AwsAccessKeyId** | **string** | AWS access key identifier. It is returned by the GET endpoint and must be sent back when editing Route53. | 
**AwsSecretAccessKey** | **string** | AWS secret access key. Must be provided when editing Route53 DNS provider. | 

## Methods

### NewRoute53StaticCredentialsRequest

`func NewRoute53StaticCredentialsRequest(type_ string, awsAccessKeyId string, awsSecretAccessKey string, ) *Route53StaticCredentialsRequest`

NewRoute53StaticCredentialsRequest instantiates a new Route53StaticCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoute53StaticCredentialsRequestWithDefaults

`func NewRoute53StaticCredentialsRequestWithDefaults() *Route53StaticCredentialsRequest`

NewRoute53StaticCredentialsRequestWithDefaults instantiates a new Route53StaticCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Route53StaticCredentialsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Route53StaticCredentialsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Route53StaticCredentialsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAwsAccessKeyId

`func (o *Route53StaticCredentialsRequest) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *Route53StaticCredentialsRequest) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *Route53StaticCredentialsRequest) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.


### GetAwsSecretAccessKey

`func (o *Route53StaticCredentialsRequest) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *Route53StaticCredentialsRequest) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *Route53StaticCredentialsRequest) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


