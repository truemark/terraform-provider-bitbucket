# Repository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**RepositoryLinks**](RepositoryLinks.md) |  | [optional] 
**Uuid** | Pointer to **string** | The repository&#39;s immutable id. This can be used as a substitute for the slug segment in URLs. Doing this guarantees your URLs will survive renaming of the repository by its owner, or even transfer of the repository to a different user. | [optional] 
**FullName** | Pointer to **string** | The concatenation of the repository owner&#39;s username and the slugified name, e.g. \&quot;evzijst/interruptingcow\&quot;. This is the same string used in Bitbucket URLs. | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 
**Parent** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**Scm** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**Account**](Account.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**UpdatedOn** | Pointer to **time.Time** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**HasIssues** | Pointer to **bool** |  | [optional] 
**HasWiki** | Pointer to **bool** |  | [optional] 
**ForkPolicy** | Pointer to **string** |  Controls the rules for forking this repository.  * **allow_forks**: unrestricted forking * **no_public_forks**: restrict forking to private forks (forks cannot   be made public later) * **no_forks**: deny all forking  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Mainbranch** | Pointer to [**Branch**](Branch.md) |  | [optional] 

## Methods

### NewRepository

`func NewRepository() *Repository`

NewRepository instantiates a new Repository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryWithDefaults

`func NewRepositoryWithDefaults() *Repository`

NewRepositoryWithDefaults instantiates a new Repository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Repository) GetLinks() RepositoryLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Repository) GetLinksOk() (*RepositoryLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Repository) SetLinks(v RepositoryLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Repository) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetUuid

`func (o *Repository) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Repository) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Repository) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Repository) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetFullName

`func (o *Repository) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Repository) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Repository) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *Repository) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetIsPrivate

`func (o *Repository) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *Repository) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *Repository) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *Repository) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetParent

`func (o *Repository) GetParent() Repository`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Repository) GetParentOk() (*Repository, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Repository) SetParent(v Repository)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Repository) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetScm

`func (o *Repository) GetScm() string`

GetScm returns the Scm field if non-nil, zero value otherwise.

### GetScmOk

`func (o *Repository) GetScmOk() (*string, bool)`

GetScmOk returns a tuple with the Scm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScm

`func (o *Repository) SetScm(v string)`

SetScm sets Scm field to given value.

### HasScm

`func (o *Repository) HasScm() bool`

HasScm returns a boolean if a field has been set.

### GetOwner

`func (o *Repository) GetOwner() Account`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Repository) GetOwnerOk() (*Account, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Repository) SetOwner(v Account)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Repository) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetName

`func (o *Repository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Repository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Repository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Repository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Repository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Repository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Repository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Repository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Repository) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Repository) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Repository) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Repository) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *Repository) GetUpdatedOn() time.Time`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *Repository) GetUpdatedOnOk() (*time.Time, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *Repository) SetUpdatedOn(v time.Time)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *Repository) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetSize

`func (o *Repository) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Repository) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Repository) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Repository) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLanguage

`func (o *Repository) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Repository) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Repository) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Repository) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetHasIssues

`func (o *Repository) GetHasIssues() bool`

GetHasIssues returns the HasIssues field if non-nil, zero value otherwise.

### GetHasIssuesOk

`func (o *Repository) GetHasIssuesOk() (*bool, bool)`

GetHasIssuesOk returns a tuple with the HasIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIssues

`func (o *Repository) SetHasIssues(v bool)`

SetHasIssues sets HasIssues field to given value.

### HasHasIssues

`func (o *Repository) HasHasIssues() bool`

HasHasIssues returns a boolean if a field has been set.

### GetHasWiki

`func (o *Repository) GetHasWiki() bool`

GetHasWiki returns the HasWiki field if non-nil, zero value otherwise.

### GetHasWikiOk

`func (o *Repository) GetHasWikiOk() (*bool, bool)`

GetHasWikiOk returns a tuple with the HasWiki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWiki

`func (o *Repository) SetHasWiki(v bool)`

SetHasWiki sets HasWiki field to given value.

### HasHasWiki

`func (o *Repository) HasHasWiki() bool`

HasHasWiki returns a boolean if a field has been set.

### GetForkPolicy

`func (o *Repository) GetForkPolicy() string`

GetForkPolicy returns the ForkPolicy field if non-nil, zero value otherwise.

### GetForkPolicyOk

`func (o *Repository) GetForkPolicyOk() (*string, bool)`

GetForkPolicyOk returns a tuple with the ForkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkPolicy

`func (o *Repository) SetForkPolicy(v string)`

SetForkPolicy sets ForkPolicy field to given value.

### HasForkPolicy

`func (o *Repository) HasForkPolicy() bool`

HasForkPolicy returns a boolean if a field has been set.

### GetProject

`func (o *Repository) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Repository) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Repository) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *Repository) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetMainbranch

`func (o *Repository) GetMainbranch() Branch`

GetMainbranch returns the Mainbranch field if non-nil, zero value otherwise.

### GetMainbranchOk

`func (o *Repository) GetMainbranchOk() (*Branch, bool)`

GetMainbranchOk returns a tuple with the Mainbranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainbranch

`func (o *Repository) SetMainbranch(v Branch)`

SetMainbranch sets Mainbranch field to given value.

### HasMainbranch

`func (o *Repository) HasMainbranch() bool`

HasMainbranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


