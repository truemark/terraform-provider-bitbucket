# Snippet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Scm** | Pointer to **string** | The DVCS used to store the snippet. | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**UpdatedOn** | Pointer to **time.Time** |  | [optional] 
**Owner** | Pointer to [**Account**](Account.md) |  | [optional] 
**Creator** | Pointer to [**Account**](Account.md) |  | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 

## Methods

### NewSnippet

`func NewSnippet() *Snippet`

NewSnippet instantiates a new Snippet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnippetWithDefaults

`func NewSnippetWithDefaults() *Snippet`

NewSnippetWithDefaults instantiates a new Snippet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Snippet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Snippet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Snippet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Snippet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Snippet) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Snippet) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Snippet) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Snippet) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetScm

`func (o *Snippet) GetScm() string`

GetScm returns the Scm field if non-nil, zero value otherwise.

### GetScmOk

`func (o *Snippet) GetScmOk() (*string, bool)`

GetScmOk returns a tuple with the Scm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScm

`func (o *Snippet) SetScm(v string)`

SetScm sets Scm field to given value.

### HasScm

`func (o *Snippet) HasScm() bool`

HasScm returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Snippet) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Snippet) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Snippet) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Snippet) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *Snippet) GetUpdatedOn() time.Time`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *Snippet) GetUpdatedOnOk() (*time.Time, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *Snippet) SetUpdatedOn(v time.Time)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *Snippet) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetOwner

`func (o *Snippet) GetOwner() Account`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Snippet) GetOwnerOk() (*Account, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Snippet) SetOwner(v Account)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Snippet) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCreator

`func (o *Snippet) GetCreator() Account`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Snippet) GetCreatorOk() (*Account, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Snippet) SetCreator(v Account)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Snippet) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetIsPrivate

`func (o *Snippet) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *Snippet) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *Snippet) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *Snippet) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


