# Pipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID identifying the pipeline. | [optional] 
**BuildNumber** | Pointer to **int32** | The build number of the pipeline. | [optional] 
**Creator** | Pointer to [**Account**](Account.md) |  | [optional] 
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**Target** | Pointer to [**PipelineTarget**](PipelineTarget.md) |  | [optional] 
**Trigger** | Pointer to [**PipelineTrigger**](PipelineTrigger.md) |  | [optional] 
**State** | Pointer to [**PipelineState**](PipelineState.md) |  | [optional] 
**CreatedOn** | Pointer to **time.Time** | The timestamp when the pipeline was created. | [optional] 
**CompletedOn** | Pointer to **time.Time** | The timestamp when the Pipeline was completed. This is not set if the pipeline is still in progress. | [optional] 
**BuildSecondsUsed** | Pointer to **int32** | The number of build seconds used by this pipeline. | [optional] 

## Methods

### NewPipeline

`func NewPipeline() *Pipeline`

NewPipeline instantiates a new Pipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineWithDefaults

`func NewPipelineWithDefaults() *Pipeline`

NewPipelineWithDefaults instantiates a new Pipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *Pipeline) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Pipeline) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Pipeline) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Pipeline) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetBuildNumber

`func (o *Pipeline) GetBuildNumber() int32`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *Pipeline) GetBuildNumberOk() (*int32, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *Pipeline) SetBuildNumber(v int32)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *Pipeline) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetCreator

`func (o *Pipeline) GetCreator() Account`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Pipeline) GetCreatorOk() (*Account, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Pipeline) SetCreator(v Account)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Pipeline) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetRepository

`func (o *Pipeline) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *Pipeline) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *Pipeline) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *Pipeline) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetTarget

`func (o *Pipeline) GetTarget() PipelineTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Pipeline) GetTargetOk() (*PipelineTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Pipeline) SetTarget(v PipelineTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Pipeline) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTrigger

`func (o *Pipeline) GetTrigger() PipelineTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *Pipeline) GetTriggerOk() (*PipelineTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *Pipeline) SetTrigger(v PipelineTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *Pipeline) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetState

`func (o *Pipeline) GetState() PipelineState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Pipeline) GetStateOk() (*PipelineState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Pipeline) SetState(v PipelineState)`

SetState sets State field to given value.

### HasState

`func (o *Pipeline) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Pipeline) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Pipeline) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Pipeline) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Pipeline) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCompletedOn

`func (o *Pipeline) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *Pipeline) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *Pipeline) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *Pipeline) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### GetBuildSecondsUsed

`func (o *Pipeline) GetBuildSecondsUsed() int32`

GetBuildSecondsUsed returns the BuildSecondsUsed field if non-nil, zero value otherwise.

### GetBuildSecondsUsedOk

`func (o *Pipeline) GetBuildSecondsUsedOk() (*int32, bool)`

GetBuildSecondsUsedOk returns a tuple with the BuildSecondsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSecondsUsed

`func (o *Pipeline) SetBuildSecondsUsed(v int32)`

SetBuildSecondsUsed sets BuildSecondsUsed field to given value.

### HasBuildSecondsUsed

`func (o *Pipeline) HasBuildSecondsUsed() bool`

HasBuildSecondsUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


