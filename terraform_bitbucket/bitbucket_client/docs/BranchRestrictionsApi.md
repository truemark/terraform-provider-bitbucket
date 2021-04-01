# \BranchRestrictionsApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesWorkspaceRepoSlugBranchRestrictionsGet**](BranchRestrictionsApi.md#RepositoriesWorkspaceRepoSlugBranchRestrictionsGet) | **Get** /repositories/{workspace}/{repo_slug}/branch-restrictions | 
[**RepositoriesWorkspaceRepoSlugBranchRestrictionsIdDelete**](BranchRestrictionsApi.md#RepositoriesWorkspaceRepoSlugBranchRestrictionsIdDelete) | **Delete** /repositories/{workspace}/{repo_slug}/branch-restrictions/{id} | 
[**RepositoriesWorkspaceRepoSlugBranchRestrictionsIdGet**](BranchRestrictionsApi.md#RepositoriesWorkspaceRepoSlugBranchRestrictionsIdGet) | **Get** /repositories/{workspace}/{repo_slug}/branch-restrictions/{id} | 
[**RepositoriesWorkspaceRepoSlugBranchRestrictionsIdPut**](BranchRestrictionsApi.md#RepositoriesWorkspaceRepoSlugBranchRestrictionsIdPut) | **Put** /repositories/{workspace}/{repo_slug}/branch-restrictions/{id} | 
[**RepositoriesWorkspaceRepoSlugBranchRestrictionsPost**](BranchRestrictionsApi.md#RepositoriesWorkspaceRepoSlugBranchRestrictionsPost) | **Post** /repositories/{workspace}/{repo_slug}/branch-restrictions | 



## RepositoriesWorkspaceRepoSlugBranchRestrictionsGet

> PaginatedBranchrestrictions RepositoriesWorkspaceRepoSlugBranchRestrictionsGet(ctx, repoSlug, workspace).Kind(kind).Pattern(pattern).Execute()





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
    kind := "kind_example" // string | Branch restrictions of this type (optional)
    pattern := "pattern_example" // string | Branch restrictions applied to branches of this pattern (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BranchRestrictionsApi.RepositoriesWorkspaceRepoSlugBranchRestrictionsGet(context.Background(), repoSlug, workspace).Kind(kind).Pattern(pattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchRestrictionsApi.RepositoriesWorkspaceRepoSlugBranchRestrictionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugBranchRestrictionsGet`: PaginatedBranchrestrictions
    fmt.Fprintf(os.Stdout, "Response from `BranchRestrictionsApi.RepositoriesWorkspaceRepoSlugBranchRestrictionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugBranchRestrictionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kind** | **string** | Branch restrictions of this type | 
 **pattern** | **string** | Branch restrictions applied to branches of this pattern | 

### Return type

[**PaginatedBranchrestrictions**](PaginatedBranchrestrictions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugBranchRestrictionsIdDelete

> RepositoriesWorkspaceRepoSlugBranchRestrictionsIdDelete(ctx, id, repoSlug, workspace).Execute()





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
    id := "id_example" // string | The restriction rule's id
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BranchRestrictionsApi.RepositoriesWorkspaceRepoSlugBranchRestrictionsIdDelete(context.Background(), id, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchRestrictionsApi.RepositoriesWorkspaceRepoSlugBranchRestrictionsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The restriction rule&#39;s id | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugBranchRestrictionsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugBranchRestrictionsIdGet

> Branchrestriction RepositoriesWorkspaceRepoSlugBranchRestrictionsIdGet(ctx, id, repoSlug, workspace).Execute()





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
    id := "id_example" // string | The restriction rule's id
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BranchRestrictionsApi.RepositoriesWorkspaceRepoSlugBranchRestrictionsIdGet(context.Background(), id, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchRestrictionsApi.RepositoriesWorkspaceRepoSlugBranchRestrictionsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugBranchRestrictionsIdGet`: Branchrestriction
    fmt.Fprintf(os.Stdout, "Response from `BranchRestrictionsApi.RepositoriesWorkspaceRepoSlugBranchRestrictionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The restriction rule&#39;s id | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugBranchRestrictionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Branchrestriction**](Branchrestriction.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugBranchRestrictionsIdPut

> Branchrestriction RepositoriesWorkspaceRepoSlugBranchRestrictionsIdPut(ctx, id, repoSlug, workspace).Body(body).Execute()





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
    id := "id_example" // string | The restriction rule's id
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
    body := *openapiclient.NewBranchrestriction("Kind_example", "BranchMatchKind_example", "Pattern_example", "Type_example") // Branchrestriction | The new version of the existing rule

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BranchRestrictionsApi.RepositoriesWorkspaceRepoSlugBranchRestrictionsIdPut(context.Background(), id, repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchRestrictionsApi.RepositoriesWorkspaceRepoSlugBranchRestrictionsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugBranchRestrictionsIdPut`: Branchrestriction
    fmt.Fprintf(os.Stdout, "Response from `BranchRestrictionsApi.RepositoriesWorkspaceRepoSlugBranchRestrictionsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The restriction rule&#39;s id | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugBranchRestrictionsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**Branchrestriction**](Branchrestriction.md) | The new version of the existing rule | 

### Return type

[**Branchrestriction**](Branchrestriction.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugBranchRestrictionsPost

> Branchrestriction RepositoriesWorkspaceRepoSlugBranchRestrictionsPost(ctx, repoSlug, workspace).Body(body).Execute()





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
    body := *openapiclient.NewBranchrestriction("Kind_example", "BranchMatchKind_example", "Pattern_example", "Type_example") // Branchrestriction | The new rule

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BranchRestrictionsApi.RepositoriesWorkspaceRepoSlugBranchRestrictionsPost(context.Background(), repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchRestrictionsApi.RepositoriesWorkspaceRepoSlugBranchRestrictionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugBranchRestrictionsPost`: Branchrestriction
    fmt.Fprintf(os.Stdout, "Response from `BranchRestrictionsApi.RepositoriesWorkspaceRepoSlugBranchRestrictionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugBranchRestrictionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Branchrestriction**](Branchrestriction.md) | The new rule | 

### Return type

[**Branchrestriction**](Branchrestriction.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

