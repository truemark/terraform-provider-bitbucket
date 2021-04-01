# PipelineKnownHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID identifying the known host. | [optional] 
**Hostname** | Pointer to **string** | The hostname of the known host. | [optional] 
**PublicKey** | Pointer to [**PipelineSshPublicKey**](PipelineSshPublicKey.md) |  | [optional] 

## Methods

### NewPipelineKnownHost

`func NewPipelineKnownHost() *PipelineKnownHost`

NewPipelineKnownHost instantiates a new PipelineKnownHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineKnownHostWithDefaults

`func NewPipelineKnownHostWithDefaults() *PipelineKnownHost`

NewPipelineKnownHostWithDefaults instantiates a new PipelineKnownHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PipelineKnownHost) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PipelineKnownHost) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PipelineKnownHost) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PipelineKnownHost) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetHostname

`func (o *PipelineKnownHost) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PipelineKnownHost) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PipelineKnownHost) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *PipelineKnownHost) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPublicKey

`func (o *PipelineKnownHost) GetPublicKey() PipelineSshPublicKey`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *PipelineKnownHost) GetPublicKeyOk() (*PipelineSshPublicKey, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *PipelineKnownHost) SetPublicKey(v PipelineSshPublicKey)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *PipelineKnownHost) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


