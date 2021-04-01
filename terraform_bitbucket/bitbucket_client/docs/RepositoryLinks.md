# RepositoryLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**Link**](Link.md) |  | [optional] 
**Html** | Pointer to [**Link**](Link.md) |  | [optional] 
**Avatar** | Pointer to [**Link**](Link.md) |  | [optional] 
**Pullrequests** | Pointer to [**Link**](Link.md) |  | [optional] 
**Commits** | Pointer to [**Link**](Link.md) |  | [optional] 
**Forks** | Pointer to [**Link**](Link.md) |  | [optional] 
**Watchers** | Pointer to [**Link**](Link.md) |  | [optional] 
**Downloads** | Pointer to [**Link**](Link.md) |  | [optional] 
**Clone** | Pointer to [**[]Link**](Link.md) |  | [optional] 
**Hooks** | Pointer to [**Link**](Link.md) |  | [optional] 

## Methods

### NewRepositoryLinks

`func NewRepositoryLinks() *RepositoryLinks`

NewRepositoryLinks instantiates a new RepositoryLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryLinksWithDefaults

`func NewRepositoryLinksWithDefaults() *RepositoryLinks`

NewRepositoryLinksWithDefaults instantiates a new RepositoryLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *RepositoryLinks) GetSelf() Link`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *RepositoryLinks) GetSelfOk() (*Link, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *RepositoryLinks) SetSelf(v Link)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *RepositoryLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetHtml

`func (o *RepositoryLinks) GetHtml() Link`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *RepositoryLinks) GetHtmlOk() (*Link, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *RepositoryLinks) SetHtml(v Link)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *RepositoryLinks) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetAvatar

`func (o *RepositoryLinks) GetAvatar() Link`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *RepositoryLinks) GetAvatarOk() (*Link, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *RepositoryLinks) SetAvatar(v Link)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *RepositoryLinks) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetPullrequests

`func (o *RepositoryLinks) GetPullrequests() Link`

GetPullrequests returns the Pullrequests field if non-nil, zero value otherwise.

### GetPullrequestsOk

`func (o *RepositoryLinks) GetPullrequestsOk() (*Link, bool)`

GetPullrequestsOk returns a tuple with the Pullrequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullrequests

`func (o *RepositoryLinks) SetPullrequests(v Link)`

SetPullrequests sets Pullrequests field to given value.

### HasPullrequests

`func (o *RepositoryLinks) HasPullrequests() bool`

HasPullrequests returns a boolean if a field has been set.

### GetCommits

`func (o *RepositoryLinks) GetCommits() Link`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *RepositoryLinks) GetCommitsOk() (*Link, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *RepositoryLinks) SetCommits(v Link)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *RepositoryLinks) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetForks

`func (o *RepositoryLinks) GetForks() Link`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *RepositoryLinks) GetForksOk() (*Link, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *RepositoryLinks) SetForks(v Link)`

SetForks sets Forks field to given value.

### HasForks

`func (o *RepositoryLinks) HasForks() bool`

HasForks returns a boolean if a field has been set.

### GetWatchers

`func (o *RepositoryLinks) GetWatchers() Link`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *RepositoryLinks) GetWatchersOk() (*Link, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *RepositoryLinks) SetWatchers(v Link)`

SetWatchers sets Watchers field to given value.

### HasWatchers

`func (o *RepositoryLinks) HasWatchers() bool`

HasWatchers returns a boolean if a field has been set.

### GetDownloads

`func (o *RepositoryLinks) GetDownloads() Link`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *RepositoryLinks) GetDownloadsOk() (*Link, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *RepositoryLinks) SetDownloads(v Link)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *RepositoryLinks) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetClone

`func (o *RepositoryLinks) GetClone() []Link`

GetClone returns the Clone field if non-nil, zero value otherwise.

### GetCloneOk

`func (o *RepositoryLinks) GetCloneOk() (*[]Link, bool)`

GetCloneOk returns a tuple with the Clone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClone

`func (o *RepositoryLinks) SetClone(v []Link)`

SetClone sets Clone field to given value.

### HasClone

`func (o *RepositoryLinks) HasClone() bool`

HasClone returns a boolean if a field has been set.

### GetHooks

`func (o *RepositoryLinks) GetHooks() Link`

GetHooks returns the Hooks field if non-nil, zero value otherwise.

### GetHooksOk

`func (o *RepositoryLinks) GetHooksOk() (*Link, bool)`

GetHooksOk returns a tuple with the Hooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooks

`func (o *RepositoryLinks) SetHooks(v Link)`

SetHooks sets Hooks field to given value.

### HasHooks

`func (o *RepositoryLinks) HasHooks() bool`

HasHooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


