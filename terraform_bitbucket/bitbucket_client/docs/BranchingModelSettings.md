# BranchingModelSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**BranchingModelSettingsLinks**](BranchingModelSettingsLinks.md) |  | [optional] 
**BranchTypes** | Pointer to [**[]BranchingModelSettingsBranchTypes**](BranchingModelSettingsBranchTypes.md) |  | [optional] 
**Development** | Pointer to [**BranchingModelSettingsDevelopment**](BranchingModelSettingsDevelopment.md) |  | [optional] 
**Production** | Pointer to [**BranchingModelSettingsProduction**](BranchingModelSettingsProduction.md) |  | [optional] 

## Methods

### NewBranchingModelSettings

`func NewBranchingModelSettings() *BranchingModelSettings`

NewBranchingModelSettings instantiates a new BranchingModelSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchingModelSettingsWithDefaults

`func NewBranchingModelSettingsWithDefaults() *BranchingModelSettings`

NewBranchingModelSettingsWithDefaults instantiates a new BranchingModelSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *BranchingModelSettings) GetLinks() BranchingModelSettingsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BranchingModelSettings) GetLinksOk() (*BranchingModelSettingsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BranchingModelSettings) SetLinks(v BranchingModelSettingsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BranchingModelSettings) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetBranchTypes

`func (o *BranchingModelSettings) GetBranchTypes() []BranchingModelSettingsBranchTypes`

GetBranchTypes returns the BranchTypes field if non-nil, zero value otherwise.

### GetBranchTypesOk

`func (o *BranchingModelSettings) GetBranchTypesOk() (*[]BranchingModelSettingsBranchTypes, bool)`

GetBranchTypesOk returns a tuple with the BranchTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchTypes

`func (o *BranchingModelSettings) SetBranchTypes(v []BranchingModelSettingsBranchTypes)`

SetBranchTypes sets BranchTypes field to given value.

### HasBranchTypes

`func (o *BranchingModelSettings) HasBranchTypes() bool`

HasBranchTypes returns a boolean if a field has been set.

### GetDevelopment

`func (o *BranchingModelSettings) GetDevelopment() BranchingModelSettingsDevelopment`

GetDevelopment returns the Development field if non-nil, zero value otherwise.

### GetDevelopmentOk

`func (o *BranchingModelSettings) GetDevelopmentOk() (*BranchingModelSettingsDevelopment, bool)`

GetDevelopmentOk returns a tuple with the Development field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevelopment

`func (o *BranchingModelSettings) SetDevelopment(v BranchingModelSettingsDevelopment)`

SetDevelopment sets Development field to given value.

### HasDevelopment

`func (o *BranchingModelSettings) HasDevelopment() bool`

HasDevelopment returns a boolean if a field has been set.

### GetProduction

`func (o *BranchingModelSettings) GetProduction() BranchingModelSettingsProduction`

GetProduction returns the Production field if non-nil, zero value otherwise.

### GetProductionOk

`func (o *BranchingModelSettings) GetProductionOk() (*BranchingModelSettingsProduction, bool)`

GetProductionOk returns a tuple with the Production field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduction

`func (o *BranchingModelSettings) SetProduction(v BranchingModelSettingsProduction)`

SetProduction sets Production field to given value.

### HasProduction

`func (o *BranchingModelSettings) HasProduction() bool`

HasProduction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


