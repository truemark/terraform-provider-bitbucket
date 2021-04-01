# PullRequestBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**MergeStrategies** | Pointer to **[]string** | Available merge strategies, when this endpoint is the destination of the pull request. | [optional] 
**DefaultMergeStrategy** | Pointer to **string** | The default merge strategy, when this endpoint is the destination of the pull request. | [optional] 

## Methods

### NewPullRequestBranch

`func NewPullRequestBranch() *PullRequestBranch`

NewPullRequestBranch instantiates a new PullRequestBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestBranchWithDefaults

`func NewPullRequestBranchWithDefaults() *PullRequestBranch`

NewPullRequestBranchWithDefaults instantiates a new PullRequestBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PullRequestBranch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PullRequestBranch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PullRequestBranch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PullRequestBranch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMergeStrategies

`func (o *PullRequestBranch) GetMergeStrategies() []string`

GetMergeStrategies returns the MergeStrategies field if non-nil, zero value otherwise.

### GetMergeStrategiesOk

`func (o *PullRequestBranch) GetMergeStrategiesOk() (*[]string, bool)`

GetMergeStrategiesOk returns a tuple with the MergeStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeStrategies

`func (o *PullRequestBranch) SetMergeStrategies(v []string)`

SetMergeStrategies sets MergeStrategies field to given value.

### HasMergeStrategies

`func (o *PullRequestBranch) HasMergeStrategies() bool`

HasMergeStrategies returns a boolean if a field has been set.

### GetDefaultMergeStrategy

`func (o *PullRequestBranch) GetDefaultMergeStrategy() string`

GetDefaultMergeStrategy returns the DefaultMergeStrategy field if non-nil, zero value otherwise.

### GetDefaultMergeStrategyOk

`func (o *PullRequestBranch) GetDefaultMergeStrategyOk() (*string, bool)`

GetDefaultMergeStrategyOk returns a tuple with the DefaultMergeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMergeStrategy

`func (o *PullRequestBranch) SetDefaultMergeStrategy(v string)`

SetDefaultMergeStrategy sets DefaultMergeStrategy field to given value.

### HasDefaultMergeStrategy

`func (o *PullRequestBranch) HasDefaultMergeStrategy() bool`

HasDefaultMergeStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


