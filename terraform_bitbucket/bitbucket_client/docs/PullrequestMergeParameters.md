# PullrequestMergeParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Message** | Pointer to **string** | The commit message that will be used on the resulting commit. | [optional] 
**CloseSourceBranch** | Pointer to **bool** | Whether the source branch should be deleted. If this is not provided, we fallback to the value used when the pull request was created, which defaults to False | [optional] 
**MergeStrategy** | Pointer to **string** | The merge strategy that will be used to merge the pull request. | [optional] [default to "merge_commit"]

## Methods

### NewPullrequestMergeParameters

`func NewPullrequestMergeParameters(type_ string, ) *PullrequestMergeParameters`

NewPullrequestMergeParameters instantiates a new PullrequestMergeParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullrequestMergeParametersWithDefaults

`func NewPullrequestMergeParametersWithDefaults() *PullrequestMergeParameters`

NewPullrequestMergeParametersWithDefaults instantiates a new PullrequestMergeParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PullrequestMergeParameters) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PullrequestMergeParameters) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PullrequestMergeParameters) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *PullrequestMergeParameters) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PullrequestMergeParameters) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PullrequestMergeParameters) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PullrequestMergeParameters) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCloseSourceBranch

`func (o *PullrequestMergeParameters) GetCloseSourceBranch() bool`

GetCloseSourceBranch returns the CloseSourceBranch field if non-nil, zero value otherwise.

### GetCloseSourceBranchOk

`func (o *PullrequestMergeParameters) GetCloseSourceBranchOk() (*bool, bool)`

GetCloseSourceBranchOk returns a tuple with the CloseSourceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseSourceBranch

`func (o *PullrequestMergeParameters) SetCloseSourceBranch(v bool)`

SetCloseSourceBranch sets CloseSourceBranch field to given value.

### HasCloseSourceBranch

`func (o *PullrequestMergeParameters) HasCloseSourceBranch() bool`

HasCloseSourceBranch returns a boolean if a field has been set.

### GetMergeStrategy

`func (o *PullrequestMergeParameters) GetMergeStrategy() string`

GetMergeStrategy returns the MergeStrategy field if non-nil, zero value otherwise.

### GetMergeStrategyOk

`func (o *PullrequestMergeParameters) GetMergeStrategyOk() (*string, bool)`

GetMergeStrategyOk returns a tuple with the MergeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeStrategy

`func (o *PullrequestMergeParameters) SetMergeStrategy(v string)`

SetMergeStrategy sets MergeStrategy field to given value.

### HasMergeStrategy

`func (o *PullrequestMergeParameters) HasMergeStrategy() bool`

HasMergeStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


