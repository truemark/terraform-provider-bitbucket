# CommitFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Path** | Pointer to **string** | The path in the repository | [optional] 
**Commit** | Pointer to [**Commit**](Commit.md) |  | [optional] 
**Attributes** | Pointer to **string** |  | [optional] 
**EscapedPath** | Pointer to **string** | The escaped version of the path as it appears in a diff. If the path does not require escaping this will be the same as path. | [optional] 

## Methods

### NewCommitFile

`func NewCommitFile(type_ string, ) *CommitFile`

NewCommitFile instantiates a new CommitFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitFileWithDefaults

`func NewCommitFileWithDefaults() *CommitFile`

NewCommitFileWithDefaults instantiates a new CommitFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CommitFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommitFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommitFile) SetType(v string)`

SetType sets Type field to given value.


### GetPath

`func (o *CommitFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CommitFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CommitFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CommitFile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetCommit

`func (o *CommitFile) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *CommitFile) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *CommitFile) SetCommit(v Commit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *CommitFile) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetAttributes

`func (o *CommitFile) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CommitFile) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CommitFile) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CommitFile) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEscapedPath

`func (o *CommitFile) GetEscapedPath() string`

GetEscapedPath returns the EscapedPath field if non-nil, zero value otherwise.

### GetEscapedPathOk

`func (o *CommitFile) GetEscapedPathOk() (*string, bool)`

GetEscapedPathOk returns a tuple with the EscapedPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscapedPath

`func (o *CommitFile) SetEscapedPath(v string)`

SetEscapedPath sets EscapedPath field to given value.

### HasEscapedPath

`func (o *CommitFile) HasEscapedPath() bool`

HasEscapedPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


