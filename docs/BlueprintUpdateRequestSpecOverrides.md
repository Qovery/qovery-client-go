# BlueprintUpdateRequestSpecOverrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineVersion** | Pointer to **string** | Terraform or OpenTofu version to use for the apply job. Required when the blueprint engine type is &#x60;terraform&#x60; or &#x60;opentofu&#x60;. Must be one of the versions in the manifest&#39;s &#x60;allowedValues&#x60; list. | [optional] 
**Credentials** | Pointer to **string** | How the apply job authenticates against the cloud provider. &#x60;cluster&#x60; reuses the cluster&#39;s cloud credentials (default). &#x60;env&#x60; expects the user to supply provider credentials as environment variables on the service. | [optional] 
**Backend** | Pointer to **string** | Where the Terraform state is stored. &#x60;qovery&#x60; stores state in a Kubernetes secret managed by Qovery (default). &#x60;user_provided&#x60; delegates to a user-controlled remote backend declared in the manifest&#39;s &#x60;backend.user_provided&#x60; block. | [optional] 
**Timeout** | Pointer to **int32** | Maximum duration in seconds for a single apply job before it is marked as timed out. Overrides the manifest&#39;s &#x60;spec.engine.timeout&#x60;. | [optional] 
**Cpu** | Pointer to **string** | CPU request/limit for the apply job pod (Kubernetes-style, e.g. &#x60;500m&#x60;). Overrides &#x60;spec.engine.resources.cpu&#x60; in the manifest. | [optional] 
**Ram** | Pointer to **string** | Memory request/limit for the apply job pod (e.g. &#x60;512Mi&#x60;, &#x60;1Gi&#x60;). Overrides &#x60;spec.engine.resources.ram&#x60; in the manifest. | [optional] 
**Storage** | Pointer to **string** | Ephemeral storage for the apply job pod — used for state files and provider plugins (e.g. &#x60;1Gi&#x60;). Overrides &#x60;spec.engine.resources.storage&#x60; in the manifest. | [optional] 

## Methods

### NewBlueprintUpdateRequestSpecOverrides

`func NewBlueprintUpdateRequestSpecOverrides() *BlueprintUpdateRequestSpecOverrides`

NewBlueprintUpdateRequestSpecOverrides instantiates a new BlueprintUpdateRequestSpecOverrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintUpdateRequestSpecOverridesWithDefaults

`func NewBlueprintUpdateRequestSpecOverridesWithDefaults() *BlueprintUpdateRequestSpecOverrides`

NewBlueprintUpdateRequestSpecOverridesWithDefaults instantiates a new BlueprintUpdateRequestSpecOverrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineVersion

`func (o *BlueprintUpdateRequestSpecOverrides) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *BlueprintUpdateRequestSpecOverrides) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *BlueprintUpdateRequestSpecOverrides) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.

### HasEngineVersion

`func (o *BlueprintUpdateRequestSpecOverrides) HasEngineVersion() bool`

HasEngineVersion returns a boolean if a field has been set.

### GetCredentials

`func (o *BlueprintUpdateRequestSpecOverrides) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *BlueprintUpdateRequestSpecOverrides) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *BlueprintUpdateRequestSpecOverrides) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *BlueprintUpdateRequestSpecOverrides) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetBackend

`func (o *BlueprintUpdateRequestSpecOverrides) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *BlueprintUpdateRequestSpecOverrides) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *BlueprintUpdateRequestSpecOverrides) SetBackend(v string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *BlueprintUpdateRequestSpecOverrides) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetTimeout

`func (o *BlueprintUpdateRequestSpecOverrides) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BlueprintUpdateRequestSpecOverrides) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BlueprintUpdateRequestSpecOverrides) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *BlueprintUpdateRequestSpecOverrides) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetCpu

`func (o *BlueprintUpdateRequestSpecOverrides) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *BlueprintUpdateRequestSpecOverrides) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *BlueprintUpdateRequestSpecOverrides) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *BlueprintUpdateRequestSpecOverrides) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRam

`func (o *BlueprintUpdateRequestSpecOverrides) GetRam() string`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *BlueprintUpdateRequestSpecOverrides) GetRamOk() (*string, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *BlueprintUpdateRequestSpecOverrides) SetRam(v string)`

SetRam sets Ram field to given value.

### HasRam

`func (o *BlueprintUpdateRequestSpecOverrides) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetStorage

`func (o *BlueprintUpdateRequestSpecOverrides) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *BlueprintUpdateRequestSpecOverrides) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *BlueprintUpdateRequestSpecOverrides) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *BlueprintUpdateRequestSpecOverrides) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


