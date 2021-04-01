# PipelineCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the command. | [optional] 
**Command** | Pointer to **string** | The executable command. | [optional] 

## Methods

### NewPipelineCommand

`func NewPipelineCommand() *PipelineCommand`

NewPipelineCommand instantiates a new PipelineCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineCommandWithDefaults

`func NewPipelineCommandWithDefaults() *PipelineCommand`

NewPipelineCommandWithDefaults instantiates a new PipelineCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PipelineCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PipelineCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCommand

`func (o *PipelineCommand) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *PipelineCommand) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *PipelineCommand) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *PipelineCommand) HasCommand() bool`

HasCommand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


