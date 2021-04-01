# DeployKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The deploy key value. | [optional] 
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**Comment** | Pointer to **string** | The comment parsed from the deploy key (if present) | [optional] 
**Label** | Pointer to **string** | The user-defined label for the deploy key | [optional] 
**AddedOn** | Pointer to **time.Time** |  | [optional] 
**LastUsed** | Pointer to **time.Time** |  | [optional] 
**Links** | Pointer to [**BranchingModelSettingsLinks**](BranchingModelSettingsLinks.md) |  | [optional] 
**Owner** | Pointer to [**Account**](Account.md) |  | [optional] 

## Methods

### NewDeployKey

`func NewDeployKey() *DeployKey`

NewDeployKey instantiates a new DeployKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployKeyWithDefaults

`func NewDeployKeyWithDefaults() *DeployKey`

NewDeployKeyWithDefaults instantiates a new DeployKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DeployKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DeployKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DeployKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DeployKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRepository

`func (o *DeployKey) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DeployKey) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DeployKey) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DeployKey) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetComment

`func (o *DeployKey) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DeployKey) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DeployKey) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DeployKey) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetLabel

`func (o *DeployKey) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DeployKey) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DeployKey) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DeployKey) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetAddedOn

`func (o *DeployKey) GetAddedOn() time.Time`

GetAddedOn returns the AddedOn field if non-nil, zero value otherwise.

### GetAddedOnOk

`func (o *DeployKey) GetAddedOnOk() (*time.Time, bool)`

GetAddedOnOk returns a tuple with the AddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOn

`func (o *DeployKey) SetAddedOn(v time.Time)`

SetAddedOn sets AddedOn field to given value.

### HasAddedOn

`func (o *DeployKey) HasAddedOn() bool`

HasAddedOn returns a boolean if a field has been set.

### GetLastUsed

`func (o *DeployKey) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *DeployKey) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *DeployKey) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *DeployKey) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetLinks

`func (o *DeployKey) GetLinks() BranchingModelSettingsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeployKey) GetLinksOk() (*BranchingModelSettingsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeployKey) SetLinks(v BranchingModelSettingsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeployKey) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetOwner

`func (o *DeployKey) GetOwner() Account`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DeployKey) GetOwnerOk() (*Account, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DeployKey) SetOwner(v Account)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DeployKey) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


