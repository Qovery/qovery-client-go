# ServiceStepMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalDurationSec** | Pointer to **NullableInt32** | The total duration in seconds of the service deployment or null if the deployment is not completed. | [optional] 
**TotalComputingDurationSec** | Pointer to **int32** | The total duration in seconds of the service deployment without queuing steps. | [optional] 
**Details** | Pointer to [**[]ServiceStepMetric**](ServiceStepMetric.md) | A list of metrics for deployment steps of the service. | [optional] 

## Methods

### NewServiceStepMetrics

`func NewServiceStepMetrics() *ServiceStepMetrics`

NewServiceStepMetrics instantiates a new ServiceStepMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStepMetricsWithDefaults

`func NewServiceStepMetricsWithDefaults() *ServiceStepMetrics`

NewServiceStepMetricsWithDefaults instantiates a new ServiceStepMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalDurationSec

`func (o *ServiceStepMetrics) GetTotalDurationSec() int32`

GetTotalDurationSec returns the TotalDurationSec field if non-nil, zero value otherwise.

### GetTotalDurationSecOk

`func (o *ServiceStepMetrics) GetTotalDurationSecOk() (*int32, bool)`

GetTotalDurationSecOk returns a tuple with the TotalDurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationSec

`func (o *ServiceStepMetrics) SetTotalDurationSec(v int32)`

SetTotalDurationSec sets TotalDurationSec field to given value.

### HasTotalDurationSec

`func (o *ServiceStepMetrics) HasTotalDurationSec() bool`

HasTotalDurationSec returns a boolean if a field has been set.

### SetTotalDurationSecNil

`func (o *ServiceStepMetrics) SetTotalDurationSecNil(b bool)`

 SetTotalDurationSecNil sets the value for TotalDurationSec to be an explicit nil

### UnsetTotalDurationSec
`func (o *ServiceStepMetrics) UnsetTotalDurationSec()`

UnsetTotalDurationSec ensures that no value is present for TotalDurationSec, not even an explicit nil
### GetTotalComputingDurationSec

`func (o *ServiceStepMetrics) GetTotalComputingDurationSec() int32`

GetTotalComputingDurationSec returns the TotalComputingDurationSec field if non-nil, zero value otherwise.

### GetTotalComputingDurationSecOk

`func (o *ServiceStepMetrics) GetTotalComputingDurationSecOk() (*int32, bool)`

GetTotalComputingDurationSecOk returns a tuple with the TotalComputingDurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalComputingDurationSec

`func (o *ServiceStepMetrics) SetTotalComputingDurationSec(v int32)`

SetTotalComputingDurationSec sets TotalComputingDurationSec field to given value.

### HasTotalComputingDurationSec

`func (o *ServiceStepMetrics) HasTotalComputingDurationSec() bool`

HasTotalComputingDurationSec returns a boolean if a field has been set.

### GetDetails

`func (o *ServiceStepMetrics) GetDetails() []ServiceStepMetric`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ServiceStepMetrics) GetDetailsOk() (*[]ServiceStepMetric, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ServiceStepMetrics) SetDetails(v []ServiceStepMetric)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ServiceStepMetrics) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


