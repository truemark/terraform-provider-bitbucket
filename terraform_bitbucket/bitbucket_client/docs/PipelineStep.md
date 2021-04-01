# PipelineStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID identifying the step. | [optional] 
**StartedOn** | Pointer to **time.Time** | The timestamp when the step execution was started. This is not set when the step hasn&#39;t executed yet. | [optional] 
**CompletedOn** | Pointer to **time.Time** | The timestamp when the step execution was completed. This is not set if the step is still in progress. | [optional] 
**State** | Pointer to [**PipelineStepState**](PipelineStepState.md) |  | [optional] 
**Image** | Pointer to [**PipelineImage**](PipelineImage.md) |  | [optional] 
**SetupCommands** | Pointer to [**[]PipelineCommand**](PipelineCommand.md) | The list of commands that are executed as part of the setup phase of the build. These commands are executed outside the build container. | [optional] 
**ScriptCommands** | Pointer to [**[]PipelineCommand**](PipelineCommand.md) | The list of build commands. These commands are executed in the build container. | [optional] 

## Methods

### NewPipelineStep

`func NewPipelineStep() *PipelineStep`

NewPipelineStep instantiates a new PipelineStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineStepWithDefaults

`func NewPipelineStepWithDefaults() *PipelineStep`

NewPipelineStepWithDefaults instantiates a new PipelineStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PipelineStep) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PipelineStep) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PipelineStep) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PipelineStep) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStartedOn

`func (o *PipelineStep) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *PipelineStep) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *PipelineStep) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *PipelineStep) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### GetCompletedOn

`func (o *PipelineStep) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *PipelineStep) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *PipelineStep) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *PipelineStep) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### GetState

`func (o *PipelineStep) GetState() PipelineStepState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PipelineStep) GetStateOk() (*PipelineStepState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PipelineStep) SetState(v PipelineStepState)`

SetState sets State field to given value.

### HasState

`func (o *PipelineStep) HasState() bool`

HasState returns a boolean if a field has been set.

### GetImage

`func (o *PipelineStep) GetImage() PipelineImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PipelineStep) GetImageOk() (*PipelineImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PipelineStep) SetImage(v PipelineImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *PipelineStep) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetSetupCommands

`func (o *PipelineStep) GetSetupCommands() []PipelineCommand`

GetSetupCommands returns the SetupCommands field if non-nil, zero value otherwise.

### GetSetupCommandsOk

`func (o *PipelineStep) GetSetupCommandsOk() (*[]PipelineCommand, bool)`

GetSetupCommandsOk returns a tuple with the SetupCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupCommands

`func (o *PipelineStep) SetSetupCommands(v []PipelineCommand)`

SetSetupCommands sets SetupCommands field to given value.

### HasSetupCommands

`func (o *PipelineStep) HasSetupCommands() bool`

HasSetupCommands returns a boolean if a field has been set.

### GetScriptCommands

`func (o *PipelineStep) GetScriptCommands() []PipelineCommand`

GetScriptCommands returns the ScriptCommands field if non-nil, zero value otherwise.

### GetScriptCommandsOk

`func (o *PipelineStep) GetScriptCommandsOk() (*[]PipelineCommand, bool)`

GetScriptCommandsOk returns a tuple with the ScriptCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptCommands

`func (o *PipelineStep) SetScriptCommands(v []PipelineCommand)`

SetScriptCommands sets ScriptCommands field to given value.

### HasScriptCommands

`func (o *PipelineStep) HasScriptCommands() bool`

HasScriptCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


