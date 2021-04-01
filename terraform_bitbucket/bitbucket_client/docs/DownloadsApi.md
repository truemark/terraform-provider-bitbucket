# \DownloadsApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesWorkspaceRepoSlugDownloadsFilenameDelete**](DownloadsApi.md#RepositoriesWorkspaceRepoSlugDownloadsFilenameDelete) | **Delete** /repositories/{workspace}/{repo_slug}/downloads/{filename} | 
[**RepositoriesWorkspaceRepoSlugDownloadsFilenameGet**](DownloadsApi.md#RepositoriesWorkspaceRepoSlugDownloadsFilenameGet) | **Get** /repositories/{workspace}/{repo_slug}/downloads/{filename} | 
[**RepositoriesWorkspaceRepoSlugDownloadsGet**](DownloadsApi.md#RepositoriesWorkspaceRepoSlugDownloadsGet) | **Get** /repositories/{workspace}/{repo_slug}/downloads | 
[**RepositoriesWorkspaceRepoSlugDownloadsPost**](DownloadsApi.md#RepositoriesWorkspaceRepoSlugDownloadsPost) | **Post** /repositories/{workspace}/{repo_slug}/downloads | 



## RepositoriesWorkspaceRepoSlugDownloadsFilenameDelete

> ModelError RepositoriesWorkspaceRepoSlugDownloadsFilenameDelete(ctx, filename, repoSlug, workspace).Execute()





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
    filename := "filename_example" // string | Name of the file.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DownloadsApi.RepositoriesWorkspaceRepoSlugDownloadsFilenameDelete(context.Background(), filename, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadsApi.RepositoriesWorkspaceRepoSlugDownloadsFilenameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugDownloadsFilenameDelete`: ModelError
    fmt.Fprintf(os.Stdout, "Response from `DownloadsApi.RepositoriesWorkspaceRepoSlugDownloadsFilenameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Name of the file. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDownloadsFilenameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ModelError**](ModelError.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugDownloadsFilenameGet

> ModelError RepositoriesWorkspaceRepoSlugDownloadsFilenameGet(ctx, filename, repoSlug, workspace).Execute()





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
    filename := "filename_example" // string | Name of the file.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DownloadsApi.RepositoriesWorkspaceRepoSlugDownloadsFilenameGet(context.Background(), filename, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadsApi.RepositoriesWorkspaceRepoSlugDownloadsFilenameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugDownloadsFilenameGet`: ModelError
    fmt.Fprintf(os.Stdout, "Response from `DownloadsApi.RepositoriesWorkspaceRepoSlugDownloadsFilenameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Name of the file. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDownloadsFilenameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ModelError**](ModelError.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugDownloadsGet

> ModelError RepositoriesWorkspaceRepoSlugDownloadsGet(ctx, repoSlug, workspace).Execute()





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
    resp, r, err := api_client.DownloadsApi.RepositoriesWorkspaceRepoSlugDownloadsGet(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadsApi.RepositoriesWorkspaceRepoSlugDownloadsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugDownloadsGet`: ModelError
    fmt.Fprintf(os.Stdout, "Response from `DownloadsApi.RepositoriesWorkspaceRepoSlugDownloadsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDownloadsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelError**](ModelError.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugDownloadsPost

> ModelError RepositoriesWorkspaceRepoSlugDownloadsPost(ctx, repoSlug, workspace).Execute()





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
    resp, r, err := api_client.DownloadsApi.RepositoriesWorkspaceRepoSlugDownloadsPost(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadsApi.RepositoriesWorkspaceRepoSlugDownloadsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugDownloadsPost`: ModelError
    fmt.Fprintf(os.Stdout, "Response from `DownloadsApi.RepositoriesWorkspaceRepoSlugDownloadsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDownloadsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelError**](ModelError.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

