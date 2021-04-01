# PipelineScheduleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID identifying the schedule. | [optional] 
**Enabled** | Pointer to **bool** | Whether the schedule is enabled. | [optional] 
**Target** | Pointer to [**PipelineTarget**](PipelineTarget.md) |  | [optional] 
**Selector** | Pointer to [**PipelineSelector**](PipelineSelector.md) |  | [optional] 
**CronPattern** | Pointer to **string** | The cron expression that the schedule applies. | [optional] 
**CreatedOn** | Pointer to **time.Time** | The timestamp when the schedule was created. | [optional] 
**UpdatedOn** | Pointer to **time.Time** | The timestamp when the schedule was updated. | [optional] 

## Methods

### NewPipelineScheduleAllOf

`func NewPipelineScheduleAllOf() *PipelineScheduleAllOf`

NewPipelineScheduleAllOf instantiates a new PipelineScheduleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineScheduleAllOfWithDefaults

`func NewPipelineScheduleAllOfWithDefaults() *PipelineScheduleAllOf`

NewPipelineScheduleAllOfWithDefaults instantiates a new PipelineScheduleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PipelineScheduleAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PipelineScheduleAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PipelineScheduleAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PipelineScheduleAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetEnabled

`func (o *PipelineScheduleAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PipelineScheduleAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PipelineScheduleAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PipelineScheduleAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTarget

`func (o *PipelineScheduleAllOf) GetTarget() PipelineTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PipelineScheduleAllOf) GetTargetOk() (*PipelineTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PipelineScheduleAllOf) SetTarget(v PipelineTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PipelineScheduleAllOf) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetSelector

`func (o *PipelineScheduleAllOf) GetSelector() PipelineSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *PipelineScheduleAllOf) GetSelectorOk() (*PipelineSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *PipelineScheduleAllOf) SetSelector(v PipelineSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *PipelineScheduleAllOf) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetCronPattern

`func (o *PipelineScheduleAllOf) GetCronPattern() string`

GetCronPattern returns the CronPattern field if non-nil, zero value otherwise.

### GetCronPatternOk

`func (o *PipelineScheduleAllOf) GetCronPatternOk() (*string, bool)`

GetCronPatternOk returns a tuple with the CronPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronPattern

`func (o *PipelineScheduleAllOf) SetCronPattern(v string)`

SetCronPattern sets CronPattern field to given value.

### HasCronPattern

`func (o *PipelineScheduleAllOf) HasCronPattern() bool`

HasCronPattern returns a boolean if a field has been set.

### GetCreatedOn

`func (o *PipelineScheduleAllOf) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *PipelineScheduleAllOf) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *PipelineScheduleAllOf) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *PipelineScheduleAllOf) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *PipelineScheduleAllOf) GetUpdatedOn() time.Time`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *PipelineScheduleAllOf) GetUpdatedOnOk() (*time.Time, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *PipelineScheduleAllOf) SetUpdatedOn(v time.Time)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *PipelineScheduleAllOf) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


