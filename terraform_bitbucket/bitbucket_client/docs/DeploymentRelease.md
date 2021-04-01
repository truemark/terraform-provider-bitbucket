# DeploymentRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID identifying the release. | [optional] 
**Name** | Pointer to **string** | The name of the release. | [optional] 
**Url** | Pointer to **string** | Link to the pipeline that produced the release. | [optional] 
**Commit** | Pointer to [**Commit**](Commit.md) |  | [optional] 
**CreatedOn** | Pointer to **time.Time** | The timestamp when the release was created. | [optional] 

## Methods

### NewDeploymentRelease

`func NewDeploymentRelease() *DeploymentRelease`

NewDeploymentRelease instantiates a new DeploymentRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentReleaseWithDefaults

`func NewDeploymentReleaseWithDefaults() *DeploymentRelease`

NewDeploymentReleaseWithDefaults instantiates a new DeploymentRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *DeploymentRelease) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeploymentRelease) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeploymentRelease) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeploymentRelease) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *DeploymentRelease) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentRelease) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentRelease) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentRelease) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *DeploymentRelease) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeploymentRelease) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeploymentRelease) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DeploymentRelease) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCommit

`func (o *DeploymentRelease) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *DeploymentRelease) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *DeploymentRelease) SetCommit(v Commit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *DeploymentRelease) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetCreatedOn

`func (o *DeploymentRelease) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *DeploymentRelease) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *DeploymentRelease) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *DeploymentRelease) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


