# DeploymentStateInProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of deployment state (IN_PROGRESS). | [optional] 
**Url** | Pointer to **string** | Link to the deployment result. | [optional] 
**Deployer** | Pointer to [**Account**](Account.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** | The timestamp when the deployment was started. | [optional] 

## Methods

### NewDeploymentStateInProgress

`func NewDeploymentStateInProgress() *DeploymentStateInProgress`

NewDeploymentStateInProgress instantiates a new DeploymentStateInProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStateInProgressWithDefaults

`func NewDeploymentStateInProgressWithDefaults() *DeploymentStateInProgress`

NewDeploymentStateInProgressWithDefaults instantiates a new DeploymentStateInProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentStateInProgress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentStateInProgress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentStateInProgress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentStateInProgress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *DeploymentStateInProgress) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeploymentStateInProgress) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeploymentStateInProgress) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DeploymentStateInProgress) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDeployer

`func (o *DeploymentStateInProgress) GetDeployer() Account`

GetDeployer returns the Deployer field if non-nil, zero value otherwise.

### GetDeployerOk

`func (o *DeploymentStateInProgress) GetDeployerOk() (*Account, bool)`

GetDeployerOk returns a tuple with the Deployer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployer

`func (o *DeploymentStateInProgress) SetDeployer(v Account)`

SetDeployer sets Deployer field to given value.

### HasDeployer

`func (o *DeploymentStateInProgress) HasDeployer() bool`

HasDeployer returns a boolean if a field has been set.

### GetStartDate

`func (o *DeploymentStateInProgress) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DeploymentStateInProgress) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DeploymentStateInProgress) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DeploymentStateInProgress) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


