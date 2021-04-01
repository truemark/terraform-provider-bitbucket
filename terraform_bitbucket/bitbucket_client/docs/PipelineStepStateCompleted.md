# PipelineStepStateCompleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of pipeline step state (COMPLETED). | [optional] 
**Result** | Pointer to [**PipelineStepStateCompletedResult**](PipelineStepStateCompletedResult.md) |  | [optional] 

## Methods

### NewPipelineStepStateCompleted

`func NewPipelineStepStateCompleted() *PipelineStepStateCompleted`

NewPipelineStepStateCompleted instantiates a new PipelineStepStateCompleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineStepStateCompletedWithDefaults

`func NewPipelineStepStateCompletedWithDefaults() *PipelineStepStateCompleted`

NewPipelineStepStateCompletedWithDefaults instantiates a new PipelineStepStateCompleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PipelineStepStateCompleted) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineStepStateCompleted) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineStepStateCompleted) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PipelineStepStateCompleted) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *PipelineStepStateCompleted) GetResult() PipelineStepStateCompletedResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PipelineStepStateCompleted) GetResultOk() (*PipelineStepStateCompletedResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PipelineStepStateCompleted) SetResult(v PipelineStepStateCompletedResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *PipelineStepStateCompleted) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


