# BuilderUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuMilliMax** | Pointer to **int32** | Highest CPU use seen, in millicores. | [optional] 
**CpuMilliAvg** | Pointer to **int32** | Average CPU use over the samples, in millicores. | [optional] 
**MemoryMibMax** | Pointer to **int32** | Highest memory use seen, in MiB. | [optional] 
**Samples** | Pointer to **int32** | Number of samples the values above come from. Usage is sampled periodically, so a short build has few samples and its peaks can be missed.  | [optional] 
**OomKilled** | Pointer to **bool** | Present, and true, only when the builder was killed for running out of memory. Absent otherwise: never test it against false. Raising the build memory setting fixes it.  | [optional] 

## Methods

### NewBuilderUsage

`func NewBuilderUsage() *BuilderUsage`

NewBuilderUsage instantiates a new BuilderUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuilderUsageWithDefaults

`func NewBuilderUsageWithDefaults() *BuilderUsage`

NewBuilderUsageWithDefaults instantiates a new BuilderUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuMilliMax

`func (o *BuilderUsage) GetCpuMilliMax() int32`

GetCpuMilliMax returns the CpuMilliMax field if non-nil, zero value otherwise.

### GetCpuMilliMaxOk

`func (o *BuilderUsage) GetCpuMilliMaxOk() (*int32, bool)`

GetCpuMilliMaxOk returns a tuple with the CpuMilliMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMilliMax

`func (o *BuilderUsage) SetCpuMilliMax(v int32)`

SetCpuMilliMax sets CpuMilliMax field to given value.

### HasCpuMilliMax

`func (o *BuilderUsage) HasCpuMilliMax() bool`

HasCpuMilliMax returns a boolean if a field has been set.

### GetCpuMilliAvg

`func (o *BuilderUsage) GetCpuMilliAvg() int32`

GetCpuMilliAvg returns the CpuMilliAvg field if non-nil, zero value otherwise.

### GetCpuMilliAvgOk

`func (o *BuilderUsage) GetCpuMilliAvgOk() (*int32, bool)`

GetCpuMilliAvgOk returns a tuple with the CpuMilliAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMilliAvg

`func (o *BuilderUsage) SetCpuMilliAvg(v int32)`

SetCpuMilliAvg sets CpuMilliAvg field to given value.

### HasCpuMilliAvg

`func (o *BuilderUsage) HasCpuMilliAvg() bool`

HasCpuMilliAvg returns a boolean if a field has been set.

### GetMemoryMibMax

`func (o *BuilderUsage) GetMemoryMibMax() int32`

GetMemoryMibMax returns the MemoryMibMax field if non-nil, zero value otherwise.

### GetMemoryMibMaxOk

`func (o *BuilderUsage) GetMemoryMibMaxOk() (*int32, bool)`

GetMemoryMibMaxOk returns a tuple with the MemoryMibMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMibMax

`func (o *BuilderUsage) SetMemoryMibMax(v int32)`

SetMemoryMibMax sets MemoryMibMax field to given value.

### HasMemoryMibMax

`func (o *BuilderUsage) HasMemoryMibMax() bool`

HasMemoryMibMax returns a boolean if a field has been set.

### GetSamples

`func (o *BuilderUsage) GetSamples() int32`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *BuilderUsage) GetSamplesOk() (*int32, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *BuilderUsage) SetSamples(v int32)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *BuilderUsage) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetOomKilled

`func (o *BuilderUsage) GetOomKilled() bool`

GetOomKilled returns the OomKilled field if non-nil, zero value otherwise.

### GetOomKilledOk

`func (o *BuilderUsage) GetOomKilledOk() (*bool, bool)`

GetOomKilledOk returns a tuple with the OomKilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOomKilled

`func (o *BuilderUsage) SetOomKilled(v bool)`

SetOomKilled sets OomKilled field to given value.

### HasOomKilled

`func (o *BuilderUsage) HasOomKilled() bool`

HasOomKilled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


