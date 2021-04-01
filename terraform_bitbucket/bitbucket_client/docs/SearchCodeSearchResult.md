# SearchCodeSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [readonly] 
**ContentMatchCount** | Pointer to **int64** |  | [optional] [readonly] 
**ContentMatches** | Pointer to [**[]SearchContentMatch**](SearchContentMatch.md) |  | [optional] [readonly] 
**PathMatches** | Pointer to [**[]SearchSegment**](SearchSegment.md) |  | [optional] [readonly] 
**File** | Pointer to [**CommitFile**](CommitFile.md) |  | [optional] 

## Methods

### NewSearchCodeSearchResult

`func NewSearchCodeSearchResult() *SearchCodeSearchResult`

NewSearchCodeSearchResult instantiates a new SearchCodeSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchCodeSearchResultWithDefaults

`func NewSearchCodeSearchResultWithDefaults() *SearchCodeSearchResult`

NewSearchCodeSearchResultWithDefaults instantiates a new SearchCodeSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SearchCodeSearchResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchCodeSearchResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchCodeSearchResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchCodeSearchResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContentMatchCount

`func (o *SearchCodeSearchResult) GetContentMatchCount() int64`

GetContentMatchCount returns the ContentMatchCount field if non-nil, zero value otherwise.

### GetContentMatchCountOk

`func (o *SearchCodeSearchResult) GetContentMatchCountOk() (*int64, bool)`

GetContentMatchCountOk returns a tuple with the ContentMatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentMatchCount

`func (o *SearchCodeSearchResult) SetContentMatchCount(v int64)`

SetContentMatchCount sets ContentMatchCount field to given value.

### HasContentMatchCount

`func (o *SearchCodeSearchResult) HasContentMatchCount() bool`

HasContentMatchCount returns a boolean if a field has been set.

### GetContentMatches

`func (o *SearchCodeSearchResult) GetContentMatches() []SearchContentMatch`

GetContentMatches returns the ContentMatches field if non-nil, zero value otherwise.

### GetContentMatchesOk

`func (o *SearchCodeSearchResult) GetContentMatchesOk() (*[]SearchContentMatch, bool)`

GetContentMatchesOk returns a tuple with the ContentMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentMatches

`func (o *SearchCodeSearchResult) SetContentMatches(v []SearchContentMatch)`

SetContentMatches sets ContentMatches field to given value.

### HasContentMatches

`func (o *SearchCodeSearchResult) HasContentMatches() bool`

HasContentMatches returns a boolean if a field has been set.

### GetPathMatches

`func (o *SearchCodeSearchResult) GetPathMatches() []SearchSegment`

GetPathMatches returns the PathMatches field if non-nil, zero value otherwise.

### GetPathMatchesOk

`func (o *SearchCodeSearchResult) GetPathMatchesOk() (*[]SearchSegment, bool)`

GetPathMatchesOk returns a tuple with the PathMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathMatches

`func (o *SearchCodeSearchResult) SetPathMatches(v []SearchSegment)`

SetPathMatches sets PathMatches field to given value.

### HasPathMatches

`func (o *SearchCodeSearchResult) HasPathMatches() bool`

HasPathMatches returns a boolean if a field has been set.

### GetFile

`func (o *SearchCodeSearchResult) GetFile() CommitFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SearchCodeSearchResult) GetFileOk() (*CommitFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SearchCodeSearchResult) SetFile(v CommitFile)`

SetFile sets File field to given value.

### HasFile

`func (o *SearchCodeSearchResult) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


