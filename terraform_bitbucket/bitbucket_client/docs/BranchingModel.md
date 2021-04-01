# BranchingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchTypes** | Pointer to [**[]BranchingModelBranchTypes**](BranchingModelBranchTypes.md) | The active branch types. | [optional] 
**Development** | Pointer to [**BranchingModelDevelopment**](BranchingModelDevelopment.md) |  | [optional] 
**Production** | Pointer to [**BranchingModelDevelopment**](BranchingModelDevelopment.md) |  | [optional] 

## Methods

### NewBranchingModel

`func NewBranchingModel() *BranchingModel`

NewBranchingModel instantiates a new BranchingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchingModelWithDefaults

`func NewBranchingModelWithDefaults() *BranchingModel`

NewBranchingModelWithDefaults instantiates a new BranchingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchTypes

`func (o *BranchingModel) GetBranchTypes() []BranchingModelBranchTypes`

GetBranchTypes returns the BranchTypes field if non-nil, zero value otherwise.

### GetBranchTypesOk

`func (o *BranchingModel) GetBranchTypesOk() (*[]BranchingModelBranchTypes, bool)`

GetBranchTypesOk returns a tuple with the BranchTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchTypes

`func (o *BranchingModel) SetBranchTypes(v []BranchingModelBranchTypes)`

SetBranchTypes sets BranchTypes field to given value.

### HasBranchTypes

`func (o *BranchingModel) HasBranchTypes() bool`

HasBranchTypes returns a boolean if a field has been set.

### GetDevelopment

`func (o *BranchingModel) GetDevelopment() BranchingModelDevelopment`

GetDevelopment returns the Development field if non-nil, zero value otherwise.

### GetDevelopmentOk

`func (o *BranchingModel) GetDevelopmentOk() (*BranchingModelDevelopment, bool)`

GetDevelopmentOk returns a tuple with the Development field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevelopment

`func (o *BranchingModel) SetDevelopment(v BranchingModelDevelopment)`

SetDevelopment sets Development field to given value.

### HasDevelopment

`func (o *BranchingModel) HasDevelopment() bool`

HasDevelopment returns a boolean if a field has been set.

### GetProduction

`func (o *BranchingModel) GetProduction() BranchingModelDevelopment`

GetProduction returns the Production field if non-nil, zero value otherwise.

### GetProductionOk

`func (o *BranchingModel) GetProductionOk() (*BranchingModelDevelopment, bool)`

GetProductionOk returns a tuple with the Production field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduction

`func (o *BranchingModel) SetProduction(v BranchingModelDevelopment)`

SetProduction sets Production field to given value.

### HasProduction

`func (o *BranchingModel) HasProduction() bool`

HasProduction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


