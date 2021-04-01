# Branchrestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**BranchingModelSettingsLinks**](BranchingModelSettingsLinks.md) |  | [optional] 
**Id** | Pointer to **int32** | The branch restriction status&#39; id. | [optional] 
**Kind** | **string** | The type of restriction that is being applied. | 
**BranchMatchKind** | **string** | Indicates how the restriction is matched against a branch. The default is &#x60;glob&#x60;. | 
**BranchType** | Pointer to **string** | Apply the restriction to branches of this type. Active when &#x60;branch_match_kind&#x60; is &#x60;branching_model&#x60;. The branch type will be calculated using the branching model configured for the repository. | [optional] 
**Pattern** | **string** | Apply the restriction to branches that match this pattern. Active when &#x60;branch_match_kind&#x60; is &#x60;glob&#x60;. Will be empty when &#x60;branch_match_kind&#x60; is &#x60;branching_model&#x60;. | 
**Users** | Pointer to [**[]Account**](Account.md) |  | [optional] 
**Groups** | Pointer to [**[]Group**](Group.md) |  | [optional] 
**Value** | Pointer to **int32** | Value with kind-specific semantics: \&quot;require_approvals_to_merge\&quot; uses it to require a minimum number of approvals on a PR; \&quot;require_default_reviewer_approvals_to_merge\&quot; uses it to require a minimum number of approvals from default reviewers on a PR; \&quot;require_passing_builds_to_merge\&quot; uses it to require a minimum number of passing builds. | [optional] 

## Methods

### NewBranchrestriction

`func NewBranchrestriction(kind string, branchMatchKind string, pattern string, ) *Branchrestriction`

NewBranchrestriction instantiates a new Branchrestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchrestrictionWithDefaults

`func NewBranchrestrictionWithDefaults() *Branchrestriction`

NewBranchrestrictionWithDefaults instantiates a new Branchrestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Branchrestriction) GetLinks() BranchingModelSettingsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Branchrestriction) GetLinksOk() (*BranchingModelSettingsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Branchrestriction) SetLinks(v BranchingModelSettingsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Branchrestriction) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *Branchrestriction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Branchrestriction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Branchrestriction) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Branchrestriction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Branchrestriction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Branchrestriction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Branchrestriction) SetKind(v string)`

SetKind sets Kind field to given value.


### GetBranchMatchKind

`func (o *Branchrestriction) GetBranchMatchKind() string`

GetBranchMatchKind returns the BranchMatchKind field if non-nil, zero value otherwise.

### GetBranchMatchKindOk

`func (o *Branchrestriction) GetBranchMatchKindOk() (*string, bool)`

GetBranchMatchKindOk returns a tuple with the BranchMatchKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchMatchKind

`func (o *Branchrestriction) SetBranchMatchKind(v string)`

SetBranchMatchKind sets BranchMatchKind field to given value.


### GetBranchType

`func (o *Branchrestriction) GetBranchType() string`

GetBranchType returns the BranchType field if non-nil, zero value otherwise.

### GetBranchTypeOk

`func (o *Branchrestriction) GetBranchTypeOk() (*string, bool)`

GetBranchTypeOk returns a tuple with the BranchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchType

`func (o *Branchrestriction) SetBranchType(v string)`

SetBranchType sets BranchType field to given value.

### HasBranchType

`func (o *Branchrestriction) HasBranchType() bool`

HasBranchType returns a boolean if a field has been set.

### GetPattern

`func (o *Branchrestriction) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *Branchrestriction) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *Branchrestriction) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetUsers

`func (o *Branchrestriction) GetUsers() []Account`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Branchrestriction) GetUsersOk() (*[]Account, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Branchrestriction) SetUsers(v []Account)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Branchrestriction) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetGroups

`func (o *Branchrestriction) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Branchrestriction) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Branchrestriction) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Branchrestriction) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetValue

`func (o *Branchrestriction) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Branchrestriction) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Branchrestriction) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *Branchrestriction) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


