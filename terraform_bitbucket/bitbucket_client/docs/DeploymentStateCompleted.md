# DeploymentStateCompleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of deployment state (COMPLETED). | [optional] 
**Url** | Pointer to **string** | Link to the deployment result. | [optional] 
**Deployer** | Pointer to [**Account**](Account.md) |  | [optional] 
**Status** | Pointer to [**DeploymentStateCompletedStatus**](DeploymentStateCompletedStatus.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** | The timestamp when the deployment was started. | [optional] 
**CompletionDate** | Pointer to **time.Time** | The timestamp when the deployment completed. | [optional] 

## Methods

### NewDeploymentStateCompleted

`func NewDeploymentStateCompleted() *DeploymentStateCompleted`

NewDeploymentStateCompleted instantiates a new DeploymentStateCompleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStateCompletedWithDefaults

`func NewDeploymentStateCompletedWithDefaults() *DeploymentStateCompleted`

NewDeploymentStateCompletedWithDefaults instantiates a new DeploymentStateCompleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentStateCompleted) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentStateCompleted) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentStateCompleted) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentStateCompleted) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *DeploymentStateCompleted) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeploymentStateCompleted) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeploymentStateCompleted) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DeploymentStateCompleted) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDeployer

`func (o *DeploymentStateCompleted) GetDeployer() Account`

GetDeployer returns the Deployer field if non-nil, zero value otherwise.

### GetDeployerOk

`func (o *DeploymentStateCompleted) GetDeployerOk() (*Account, bool)`

GetDeployerOk returns a tuple with the Deployer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployer

`func (o *DeploymentStateCompleted) SetDeployer(v Account)`

SetDeployer sets Deployer field to given value.

### HasDeployer

`func (o *DeploymentStateCompleted) HasDeployer() bool`

HasDeployer returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentStateCompleted) GetStatus() DeploymentStateCompletedStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentStateCompleted) GetStatusOk() (*DeploymentStateCompletedStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentStateCompleted) SetStatus(v DeploymentStateCompletedStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentStateCompleted) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *DeploymentStateCompleted) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DeploymentStateCompleted) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DeploymentStateCompleted) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DeploymentStateCompleted) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetCompletionDate

`func (o *DeploymentStateCompleted) GetCompletionDate() time.Time`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *DeploymentStateCompleted) GetCompletionDateOk() (*time.Time, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *DeploymentStateCompleted) SetCompletionDate(v time.Time)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *DeploymentStateCompleted) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


