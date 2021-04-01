# PipelineStepStateCompletedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the result (ERROR) | [optional] 
**Error** | Pointer to [**PipelineStepError**](PipelineStepError.md) |  | [optional] 

## Methods

### NewPipelineStepStateCompletedError

`func NewPipelineStepStateCompletedError() *PipelineStepStateCompletedError`

NewPipelineStepStateCompletedError instantiates a new PipelineStepStateCompletedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineStepStateCompletedErrorWithDefaults

`func NewPipelineStepStateCompletedErrorWithDefaults() *PipelineStepStateCompletedError`

NewPipelineStepStateCompletedErrorWithDefaults instantiates a new PipelineStepStateCompletedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PipelineStepStateCompletedError) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineStepStateCompletedError) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineStepStateCompletedError) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PipelineStepStateCompletedError) HasName() bool`

HasName returns a boolean if a field has been set.

### GetError

`func (o *PipelineStepStateCompletedError) GetError() PipelineStepError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PipelineStepStateCompletedError) GetErrorOk() (*PipelineStepError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PipelineStepStateCompletedError) SetError(v PipelineStepError)`

SetError sets Error field to given value.

### HasError

`func (o *PipelineStepStateCompletedError) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


