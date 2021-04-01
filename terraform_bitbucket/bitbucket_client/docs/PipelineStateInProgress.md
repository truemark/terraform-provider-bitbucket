# PipelineStateInProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of pipeline state (IN_PROGRESS). | [optional] 
**Stage** | Pointer to [**PipelineStateInProgressStage**](PipelineStateInProgressStage.md) |  | [optional] 

## Methods

### NewPipelineStateInProgress

`func NewPipelineStateInProgress() *PipelineStateInProgress`

NewPipelineStateInProgress instantiates a new PipelineStateInProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineStateInProgressWithDefaults

`func NewPipelineStateInProgressWithDefaults() *PipelineStateInProgress`

NewPipelineStateInProgressWithDefaults instantiates a new PipelineStateInProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PipelineStateInProgress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineStateInProgress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineStateInProgress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PipelineStateInProgress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStage

`func (o *PipelineStateInProgress) GetStage() PipelineStateInProgressStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *PipelineStateInProgress) GetStageOk() (*PipelineStateInProgressStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *PipelineStateInProgress) SetStage(v PipelineStateInProgressStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *PipelineStateInProgress) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


