# PipelineCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID identifying the pipeline cache. | [optional] 
**PipelineUuid** | Pointer to **string** | The UUID of the pipeline that created the cache. | [optional] 
**StepUuid** | Pointer to **string** | The uuid of the step that created the cache. | [optional] 
**Name** | Pointer to **string** | The name of the cache. | [optional] 
**Path** | Pointer to **string** | The path where the cache contents were retrieved from. | [optional] 
**FileSizeBytes** | Pointer to **int32** | The size of the file containing the archive of the cache. | [optional] 
**CreatedOn** | Pointer to **time.Time** | The timestamp when the cache was created. | [optional] 

## Methods

### NewPipelineCache

`func NewPipelineCache() *PipelineCache`

NewPipelineCache instantiates a new PipelineCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineCacheWithDefaults

`func NewPipelineCacheWithDefaults() *PipelineCache`

NewPipelineCacheWithDefaults instantiates a new PipelineCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PipelineCache) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PipelineCache) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PipelineCache) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PipelineCache) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetPipelineUuid

`func (o *PipelineCache) GetPipelineUuid() string`

GetPipelineUuid returns the PipelineUuid field if non-nil, zero value otherwise.

### GetPipelineUuidOk

`func (o *PipelineCache) GetPipelineUuidOk() (*string, bool)`

GetPipelineUuidOk returns a tuple with the PipelineUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineUuid

`func (o *PipelineCache) SetPipelineUuid(v string)`

SetPipelineUuid sets PipelineUuid field to given value.

### HasPipelineUuid

`func (o *PipelineCache) HasPipelineUuid() bool`

HasPipelineUuid returns a boolean if a field has been set.

### GetStepUuid

`func (o *PipelineCache) GetStepUuid() string`

GetStepUuid returns the StepUuid field if non-nil, zero value otherwise.

### GetStepUuidOk

`func (o *PipelineCache) GetStepUuidOk() (*string, bool)`

GetStepUuidOk returns a tuple with the StepUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepUuid

`func (o *PipelineCache) SetStepUuid(v string)`

SetStepUuid sets StepUuid field to given value.

### HasStepUuid

`func (o *PipelineCache) HasStepUuid() bool`

HasStepUuid returns a boolean if a field has been set.

### GetName

`func (o *PipelineCache) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineCache) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineCache) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PipelineCache) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *PipelineCache) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PipelineCache) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PipelineCache) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PipelineCache) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetFileSizeBytes

`func (o *PipelineCache) GetFileSizeBytes() int32`

GetFileSizeBytes returns the FileSizeBytes field if non-nil, zero value otherwise.

### GetFileSizeBytesOk

`func (o *PipelineCache) GetFileSizeBytesOk() (*int32, bool)`

GetFileSizeBytesOk returns a tuple with the FileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeBytes

`func (o *PipelineCache) SetFileSizeBytes(v int32)`

SetFileSizeBytes sets FileSizeBytes field to given value.

### HasFileSizeBytes

`func (o *PipelineCache) HasFileSizeBytes() bool`

HasFileSizeBytes returns a boolean if a field has been set.

### GetCreatedOn

`func (o *PipelineCache) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *PipelineCache) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *PipelineCache) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *PipelineCache) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


