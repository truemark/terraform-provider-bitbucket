# \CommitStatusesApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet**](CommitStatusesApi.md#RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/statuses/build/{key} | 
[**RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut**](CommitStatusesApi.md#RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut) | **Put** /repositories/{workspace}/{repo_slug}/commit/{commit}/statuses/build/{key} | 
[**RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost**](CommitStatusesApi.md#RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost) | **Post** /repositories/{workspace}/{repo_slug}/commit/{commit}/statuses/build | 
[**RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet**](CommitStatusesApi.md#RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/statuses | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet**](CommitStatusesApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/statuses | 



## RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet

> Commitstatus RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet(ctx, commit, key, repoSlug, workspace).Execute()





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
    commit := "commit_example" // string | The commit's SHA1.
    key := "key_example" // string | The build status' unique key
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitStatusesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet(context.Background(), commit, key, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitStatusesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet`: Commitstatus
    fmt.Fprintf(os.Stdout, "Response from `CommitStatusesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commit** | **string** | The commit&#39;s SHA1. | 
**key** | **string** | The build status&#39; unique key | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Commitstatus**](Commitstatus.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut

> Commitstatus RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut(ctx, commit, key, repoSlug, workspace).Body(body).Execute()





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
    commit := "commit_example" // string | The commit's SHA1.
    key := "key_example" // string | The build status' unique key
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
    body := *openapiclient.NewCommitstatus("Type_example") // Commitstatus | The updated build status object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitStatusesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut(context.Background(), commit, key, repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitStatusesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut`: Commitstatus
    fmt.Fprintf(os.Stdout, "Response from `CommitStatusesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commit** | **string** | The commit&#39;s SHA1. | 
**key** | **string** | The build status&#39; unique key | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**Commitstatus**](Commitstatus.md) | The updated build status object | 

### Return type

[**Commitstatus**](Commitstatus.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost

> Commitstatus RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost(ctx, commit, repoSlug, workspace).Body(body).Execute()





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
    commit := "commit_example" // string | The commit's SHA1.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
    body := *openapiclient.NewCommitstatus("Type_example") // Commitstatus | The new commit status object. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitStatusesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost(context.Background(), commit, repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitStatusesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost`: Commitstatus
    fmt.Fprintf(os.Stdout, "Response from `CommitStatusesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commit** | **string** | The commit&#39;s SHA1. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**Commitstatus**](Commitstatus.md) | The new commit status object. | 

### Return type

[**Commitstatus**](Commitstatus.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet

> PaginatedCommitstatuses RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet(ctx, commit, repoSlug, workspace).Q(q).Sort(sort).Execute()





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
    commit := "commit_example" // string | The commit's SHA1.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
    q := "q_example" // string | Query string to narrow down the response as per [filtering and sorting](../../../../../../meta/filtering).  (optional)
    sort := "sort_example" // string | Field by which the results should be sorted as per [filtering and sorting](../../../../../../meta/filtering). Defaults to `created_on`.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitStatusesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet(context.Background(), commit, repoSlug, workspace).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitStatusesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet`: PaginatedCommitstatuses
    fmt.Fprintf(os.Stdout, "Response from `CommitStatusesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commit** | **string** | The commit&#39;s SHA1. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugCommitCommitStatusesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **q** | **string** | Query string to narrow down the response as per [filtering and sorting](../../../../../../meta/filtering).  | 
 **sort** | **string** | Field by which the results should be sorted as per [filtering and sorting](../../../../../../meta/filtering). Defaults to &#x60;created_on&#x60;.  | 

### Return type

[**PaginatedCommitstatuses**](PaginatedCommitstatuses.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet

> PaginatedCommitstatuses RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet(ctx, pullRequestId, repoSlug, workspace).Q(q).Sort(sort).Execute()





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
    pullRequestId := int32(56) // int32 | The id of the pull request.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
    q := "q_example" // string | Query string to narrow down the response as per [filtering and sorting](../../../../../../meta/filtering).  (optional)
    sort := "sort_example" // string | Field by which the results should be sorted as per [filtering and sorting](../../../../../../meta/filtering). Defaults to `created_on`.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitStatusesApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet(context.Background(), pullRequestId, repoSlug, workspace).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitStatusesApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet`: PaginatedCommitstatuses
    fmt.Fprintf(os.Stdout, "Response from `CommitStatusesApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pullRequestId** | **int32** | The id of the pull request. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **q** | **string** | Query string to narrow down the response as per [filtering and sorting](../../../../../../meta/filtering).  | 
 **sort** | **string** | Field by which the results should be sorted as per [filtering and sorting](../../../../../../meta/filtering). Defaults to &#x60;created_on&#x60;.  | 

### Return type

[**PaginatedCommitstatuses**](PaginatedCommitstatuses.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

