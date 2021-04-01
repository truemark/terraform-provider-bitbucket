# Ref

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Links** | Pointer to [**RefLinks**](RefLinks.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the ref. | [optional] 
**Target** | Pointer to [**Commit**](Commit.md) |  | [optional] 

## Methods

### NewRef

`func NewRef(type_ string, ) *Ref`

NewRef instantiates a new Ref object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefWithDefaults

`func NewRefWithDefaults() *Ref`

NewRefWithDefaults instantiates a new Ref object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Ref) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Ref) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Ref) SetType(v string)`

SetType sets Type field to given value.


### GetLinks

`func (o *Ref) GetLinks() RefLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Ref) GetLinksOk() (*RefLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Ref) SetLinks(v RefLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Ref) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *Ref) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ref) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ref) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ref) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTarget

`func (o *Ref) GetTarget() Commit`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Ref) GetTargetOk() (*Commit, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Ref) SetTarget(v Commit)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Ref) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


