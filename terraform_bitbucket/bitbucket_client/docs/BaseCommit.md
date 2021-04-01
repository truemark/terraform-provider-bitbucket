# BaseCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**Author** | Pointer to [**Author**](Author.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to [**IssueContent**](IssueContent.md) |  | [optional] 
**Parents** | Pointer to [**[]BaseCommit**](BaseCommit.md) |  | [optional] 

## Methods

### NewBaseCommit

`func NewBaseCommit() *BaseCommit`

NewBaseCommit instantiates a new BaseCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCommitWithDefaults

`func NewBaseCommitWithDefaults() *BaseCommit`

NewBaseCommitWithDefaults instantiates a new BaseCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *BaseCommit) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *BaseCommit) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *BaseCommit) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *BaseCommit) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetDate

`func (o *BaseCommit) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BaseCommit) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BaseCommit) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *BaseCommit) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetAuthor

`func (o *BaseCommit) GetAuthor() Author`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *BaseCommit) GetAuthorOk() (*Author, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *BaseCommit) SetAuthor(v Author)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *BaseCommit) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetMessage

`func (o *BaseCommit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BaseCommit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BaseCommit) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BaseCommit) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSummary

`func (o *BaseCommit) GetSummary() IssueContent`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *BaseCommit) GetSummaryOk() (*IssueContent, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *BaseCommit) SetSummary(v IssueContent)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *BaseCommit) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetParents

`func (o *BaseCommit) GetParents() []BaseCommit`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *BaseCommit) GetParentsOk() (*[]BaseCommit, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *BaseCommit) SetParents(v []BaseCommit)`

SetParents sets Parents field to given value.

### HasParents

`func (o *BaseCommit) HasParents() bool`

HasParents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


