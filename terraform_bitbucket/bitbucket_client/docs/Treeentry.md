# Treeentry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Path** | Pointer to **string** | The path in the repository | [optional] 
**Commit** | Pointer to [**Commit**](Commit.md) |  | [optional] 

## Methods

### NewTreeentry

`func NewTreeentry(type_ string, ) *Treeentry`

NewTreeentry instantiates a new Treeentry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreeentryWithDefaults

`func NewTreeentryWithDefaults() *Treeentry`

NewTreeentryWithDefaults instantiates a new Treeentry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Treeentry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Treeentry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Treeentry) SetType(v string)`

SetType sets Type field to given value.


### GetPath

`func (o *Treeentry) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Treeentry) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Treeentry) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Treeentry) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetCommit

`func (o *Treeentry) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *Treeentry) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *Treeentry) SetCommit(v Commit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *Treeentry) HasCommit() bool`

HasCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


