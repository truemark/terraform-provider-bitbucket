# PipelineKnownHostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID identifying the known host. | [optional] 
**Hostname** | Pointer to **string** | The hostname of the known host. | [optional] 
**PublicKey** | Pointer to [**PipelineSshPublicKey**](PipelineSshPublicKey.md) |  | [optional] 

## Methods

### NewPipelineKnownHostAllOf

`func NewPipelineKnownHostAllOf() *PipelineKnownHostAllOf`

NewPipelineKnownHostAllOf instantiates a new PipelineKnownHostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineKnownHostAllOfWithDefaults

`func NewPipelineKnownHostAllOfWithDefaults() *PipelineKnownHostAllOf`

NewPipelineKnownHostAllOfWithDefaults instantiates a new PipelineKnownHostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PipelineKnownHostAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PipelineKnownHostAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PipelineKnownHostAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PipelineKnownHostAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetHostname

`func (o *PipelineKnownHostAllOf) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PipelineKnownHostAllOf) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PipelineKnownHostAllOf) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *PipelineKnownHostAllOf) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPublicKey

`func (o *PipelineKnownHostAllOf) GetPublicKey() PipelineSshPublicKey`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *PipelineKnownHostAllOf) GetPublicKeyOk() (*PipelineSshPublicKey, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *PipelineKnownHostAllOf) SetPublicKey(v PipelineSshPublicKey)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *PipelineKnownHostAllOf) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


