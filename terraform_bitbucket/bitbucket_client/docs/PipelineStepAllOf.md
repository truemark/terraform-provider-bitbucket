# PipelineStepAllOf

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

### NewPipelineStepAllOf

`func NewPipelineStepAllOf() *PipelineStepAllOf`

NewPipelineStepAllOf instantiates a new PipelineStepAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineStepAllOfWithDefaults

`func NewPipelineStepAllOfWithDefaults() *PipelineStepAllOf`

NewPipelineStepAllOfWithDefaults instantiates a new PipelineStepAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PipelineStepAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PipelineStepAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PipelineStepAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PipelineStepAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStartedOn

`func (o *PipelineStepAllOf) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *PipelineStepAllOf) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *PipelineStepAllOf) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *PipelineStepAllOf) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### GetCompletedOn

`func (o *PipelineStepAllOf) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *PipelineStepAllOf) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *PipelineStepAllOf) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *PipelineStepAllOf) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### GetState

`func (o *PipelineStepAllOf) GetState() PipelineStepState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PipelineStepAllOf) GetStateOk() (*PipelineStepState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PipelineStepAllOf) SetState(v PipelineStepState)`

SetState sets State field to given value.

### HasState

`func (o *PipelineStepAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetImage

`func (o *PipelineStepAllOf) GetImage() PipelineImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PipelineStepAllOf) GetImageOk() (*PipelineImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PipelineStepAllOf) SetImage(v PipelineImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *PipelineStepAllOf) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetSetupCommands

`func (o *PipelineStepAllOf) GetSetupCommands() []PipelineCommand`

GetSetupCommands returns the SetupCommands field if non-nil, zero value otherwise.

### GetSetupCommandsOk

`func (o *PipelineStepAllOf) GetSetupCommandsOk() (*[]PipelineCommand, bool)`

GetSetupCommandsOk returns a tuple with the SetupCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupCommands

`func (o *PipelineStepAllOf) SetSetupCommands(v []PipelineCommand)`

SetSetupCommands sets SetupCommands field to given value.

### HasSetupCommands

`func (o *PipelineStepAllOf) HasSetupCommands() bool`

HasSetupCommands returns a boolean if a field has been set.

### GetScriptCommands

`func (o *PipelineStepAllOf) GetScriptCommands() []PipelineCommand`

GetScriptCommands returns the ScriptCommands field if non-nil, zero value otherwise.

### GetScriptCommandsOk

`func (o *PipelineStepAllOf) GetScriptCommandsOk() (*[]PipelineCommand, bool)`

GetScriptCommandsOk returns a tuple with the ScriptCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptCommands

`func (o *PipelineStepAllOf) SetScriptCommands(v []PipelineCommand)`

SetScriptCommands sets ScriptCommands field to given value.

### HasScriptCommands

`func (o *PipelineStepAllOf) HasScriptCommands() bool`

HasScriptCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


