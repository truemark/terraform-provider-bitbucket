# HookEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | The event identifier. | [optional] 
**Category** | Pointer to **string** | The category this event belongs to. | [optional] 
**Label** | Pointer to **string** | Summary of the webhook event type. | [optional] 
**Description** | Pointer to **string** | More detailed description of the webhook event type. | [optional] 

## Methods

### NewHookEvent

`func NewHookEvent() *HookEvent`

NewHookEvent instantiates a new HookEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookEventWithDefaults

`func NewHookEventWithDefaults() *HookEvent`

NewHookEventWithDefaults instantiates a new HookEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *HookEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *HookEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *HookEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *HookEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetCategory

`func (o *HookEvent) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HookEvent) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HookEvent) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *HookEvent) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLabel

`func (o *HookEvent) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *HookEvent) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *HookEvent) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *HookEvent) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *HookEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HookEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HookEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HookEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


