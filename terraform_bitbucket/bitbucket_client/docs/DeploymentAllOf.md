# DeploymentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID identifying the deployment. | [optional] 
**State** | Pointer to [**DeploymentState**](DeploymentState.md) |  | [optional] 
**Environment** | Pointer to [**DeploymentEnvironment**](DeploymentEnvironment.md) |  | [optional] 
**Release** | Pointer to [**DeploymentRelease**](DeploymentRelease.md) |  | [optional] 

## Methods

### NewDeploymentAllOf

`func NewDeploymentAllOf() *DeploymentAllOf`

NewDeploymentAllOf instantiates a new DeploymentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentAllOfWithDefaults

`func NewDeploymentAllOfWithDefaults() *DeploymentAllOf`

NewDeploymentAllOfWithDefaults instantiates a new DeploymentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *DeploymentAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeploymentAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeploymentAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeploymentAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *DeploymentAllOf) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeploymentAllOf) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeploymentAllOf) SetState(v DeploymentState)`

SetState sets State field to given value.

### HasState

`func (o *DeploymentAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEnvironment

`func (o *DeploymentAllOf) GetEnvironment() DeploymentEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DeploymentAllOf) GetEnvironmentOk() (*DeploymentEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DeploymentAllOf) SetEnvironment(v DeploymentEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DeploymentAllOf) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRelease

`func (o *DeploymentAllOf) GetRelease() DeploymentRelease`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *DeploymentAllOf) GetReleaseOk() (*DeploymentRelease, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *DeploymentAllOf) SetRelease(v DeploymentRelease)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *DeploymentAllOf) HasRelease() bool`

HasRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


