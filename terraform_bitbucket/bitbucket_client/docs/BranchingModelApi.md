# \BranchingModelApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesWorkspaceRepoSlugBranchingModelGet**](BranchingModelApi.md#RepositoriesWorkspaceRepoSlugBranchingModelGet) | **Get** /repositories/{workspace}/{repo_slug}/branching-model | 
[**RepositoriesWorkspaceRepoSlugBranchingModelSettingsGet**](BranchingModelApi.md#RepositoriesWorkspaceRepoSlugBranchingModelSettingsGet) | **Get** /repositories/{workspace}/{repo_slug}/branching-model/settings | 
[**RepositoriesWorkspaceRepoSlugBranchingModelSettingsPut**](BranchingModelApi.md#RepositoriesWorkspaceRepoSlugBranchingModelSettingsPut) | **Put** /repositories/{workspace}/{repo_slug}/branching-model/settings | 



## RepositoriesWorkspaceRepoSlugBranchingModelGet

> BranchingModel RepositoriesWorkspaceRepoSlugBranchingModelGet(ctx, repoSlug, workspace).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BranchingModelApi.RepositoriesWorkspaceRepoSlugBranchingModelGet(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchingModelApi.RepositoriesWorkspaceRepoSlugBranchingModelGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugBranchingModelGet`: BranchingModel
    fmt.Fprintf(os.Stdout, "Response from `BranchingModelApi.RepositoriesWorkspaceRepoSlugBranchingModelGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugBranchingModelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BranchingModel**](BranchingModel.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugBranchingModelSettingsGet

> BranchingModelSettings RepositoriesWorkspaceRepoSlugBranchingModelSettingsGet(ctx, repoSlug, workspace).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BranchingModelApi.RepositoriesWorkspaceRepoSlugBranchingModelSettingsGet(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchingModelApi.RepositoriesWorkspaceRepoSlugBranchingModelSettingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugBranchingModelSettingsGet`: BranchingModelSettings
    fmt.Fprintf(os.Stdout, "Response from `BranchingModelApi.RepositoriesWorkspaceRepoSlugBranchingModelSettingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugBranchingModelSettingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BranchingModelSettings**](BranchingModelSettings.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugBranchingModelSettingsPut

> BranchingModelSettings RepositoriesWorkspaceRepoSlugBranchingModelSettingsPut(ctx, repoSlug, workspace).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BranchingModelApi.RepositoriesWorkspaceRepoSlugBranchingModelSettingsPut(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchingModelApi.RepositoriesWorkspaceRepoSlugBranchingModelSettingsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugBranchingModelSettingsPut`: BranchingModelSettings
    fmt.Fprintf(os.Stdout, "Response from `BranchingModelApi.RepositoriesWorkspaceRepoSlugBranchingModelSettingsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugBranchingModelSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BranchingModelSettings**](BranchingModelSettings.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

