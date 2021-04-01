# Component

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**BranchingModelSettingsLinks**](BranchingModelSettingsLinks.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 

## Methods

### NewComponent

`func NewComponent() *Component`

NewComponent instantiates a new Component object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentWithDefaults

`func NewComponentWithDefaults() *Component`

NewComponentWithDefaults instantiates a new Component object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Component) GetLinks() BranchingModelSettingsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Component) GetLinksOk() (*BranchingModelSettingsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Component) SetLinks(v BranchingModelSettingsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Component) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *Component) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Component) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Component) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Component) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *Component) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Component) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Component) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Component) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


