# \RepositoriesApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesGet**](RepositoriesApi.md#RepositoriesGet) | **Get** /repositories | 
[**RepositoriesWorkspaceGet**](RepositoriesApi.md#RepositoriesWorkspaceGet) | **Get** /repositories/{workspace} | 
[**RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/statuses/build/{key} | 
[**RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut) | **Put** /repositories/{workspace}/{repo_slug}/commit/{commit}/statuses/build/{key} | 
[**RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost) | **Post** /repositories/{workspace}/{repo_slug}/commit/{commit}/statuses/build | 
[**RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/statuses | 
[**RepositoriesWorkspaceRepoSlugDelete**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugDelete) | **Delete** /repositories/{workspace}/{repo_slug} | 
[**RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet) | **Get** /repositories/{workspace}/{repo_slug}/filehistory/{commit}/{path} | 
[**RepositoriesWorkspaceRepoSlugForksGet**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugForksGet) | **Get** /repositories/{workspace}/{repo_slug}/forks | 
[**RepositoriesWorkspaceRepoSlugForksPost**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugForksPost) | **Post** /repositories/{workspace}/{repo_slug}/forks | 
[**RepositoriesWorkspaceRepoSlugGet**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugGet) | **Get** /repositories/{workspace}/{repo_slug} | 
[**RepositoriesWorkspaceRepoSlugHooksGet**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugHooksGet) | **Get** /repositories/{workspace}/{repo_slug}/hooks | 
[**RepositoriesWorkspaceRepoSlugHooksPost**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugHooksPost) | **Post** /repositories/{workspace}/{repo_slug}/hooks | 
[**RepositoriesWorkspaceRepoSlugHooksUidDelete**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugHooksUidDelete) | **Delete** /repositories/{workspace}/{repo_slug}/hooks/{uid} | 
[**RepositoriesWorkspaceRepoSlugHooksUidGet**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugHooksUidGet) | **Get** /repositories/{workspace}/{repo_slug}/hooks/{uid} | 
[**RepositoriesWorkspaceRepoSlugHooksUidPut**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugHooksUidPut) | **Put** /repositories/{workspace}/{repo_slug}/hooks/{uid} | 
[**RepositoriesWorkspaceRepoSlugPost**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugPost) | **Post** /repositories/{workspace}/{repo_slug} | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/statuses | 
[**RepositoriesWorkspaceRepoSlugPut**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugPut) | **Put** /repositories/{workspace}/{repo_slug} | 
[**RepositoriesWorkspaceRepoSlugSrcCommitPathGet**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugSrcCommitPathGet) | **Get** /repositories/{workspace}/{repo_slug}/src/{commit}/{path} | 
[**RepositoriesWorkspaceRepoSlugSrcGet**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugSrcGet) | **Get** /repositories/{workspace}/{repo_slug}/src | 
[**RepositoriesWorkspaceRepoSlugSrcPost**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugSrcPost) | **Post** /repositories/{workspace}/{repo_slug}/src | 
[**RepositoriesWorkspaceRepoSlugWatchersGet**](RepositoriesApi.md#RepositoriesWorkspaceRepoSlugWatchersGet) | **Get** /repositories/{workspace}/{repo_slug}/watchers | 
[**UserPermissionsRepositoriesGet**](RepositoriesApi.md#UserPermissionsRepositoriesGet) | **Get** /user/permissions/repositories | 



## RepositoriesGet

> PaginatedRepositories RepositoriesGet(ctx).After(after).Role(role).Q(q).Sort(sort).Execute()





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
    after := "after_example" // string | Filter the results to include only repositories created on or after this [ISO-8601](https://en.wikipedia.org/wiki/ISO_8601)  timestamp. Example: `YYYY-MM-DDTHH:mm:ss.sssZ` (optional)
    role := "role_example" // string | Filters the result based on the authenticated user's role on each repository.  * **member**: returns repositories to which the user has explicit read access * **contributor**: returns repositories to which the user has explicit write access * **admin**: returns repositories to which the user has explicit administrator access * **owner**: returns all repositories owned by the current user  (optional)
    q := "q_example" // string | Query string to narrow down the response as per [filtering and sorting](../meta/filtering). `role` parameter must also be specified.  (optional)
    sort := "sort_example" // string | Field by which the results should be sorted as per [filtering and sorting](../meta/filtering).  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesGet(context.Background()).After(after).Role(role).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesGet`: PaginatedRepositories
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | Filter the results to include only repositories created on or after this [ISO-8601](https://en.wikipedia.org/wiki/ISO_8601)  timestamp. Example: &#x60;YYYY-MM-DDTHH:mm:ss.sssZ&#x60; | 
 **role** | **string** | Filters the result based on the authenticated user&#39;s role on each repository.  * **member**: returns repositories to which the user has explicit read access * **contributor**: returns repositories to which the user has explicit write access * **admin**: returns repositories to which the user has explicit administrator access * **owner**: returns all repositories owned by the current user  | 
 **q** | **string** | Query string to narrow down the response as per [filtering and sorting](../meta/filtering). &#x60;role&#x60; parameter must also be specified.  | 
 **sort** | **string** | Field by which the results should be sorted as per [filtering and sorting](../meta/filtering).  | 

### Return type

[**PaginatedRepositories**](PaginatedRepositories.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceGet

> PaginatedRepositories RepositoriesWorkspaceGet(ctx, workspace).Role(role).Q(q).Sort(sort).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
    role := "role_example" // string |  Filters the result based on the authenticated user's role on each repository.  * **member**: returns repositories to which the user has explicit read access * **contributor**: returns repositories to which the user has explicit write access * **admin**: returns repositories to which the user has explicit administrator access * **owner**: returns all repositories owned by the current user  (optional)
    q := "q_example" // string |  Query string to narrow down the response as per [filtering and sorting](../../meta/filtering).  (optional)
    sort := "sort_example" // string |  Field by which the results should be sorted as per [filtering and sorting](../../meta/filtering).          (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceGet(context.Background(), workspace).Role(role).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceGet`: PaginatedRepositories
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | **string** |  Filters the result based on the authenticated user&#39;s role on each repository.  * **member**: returns repositories to which the user has explicit read access * **contributor**: returns repositories to which the user has explicit write access * **admin**: returns repositories to which the user has explicit administrator access * **owner**: returns all repositories owned by the current user  | 
 **q** | **string** |  Query string to narrow down the response as per [filtering and sorting](../../meta/filtering).  | 
 **sort** | **string** |  Field by which the results should be sorted as per [filtering and sorting](../../meta/filtering).          | 

### Return type

[**PaginatedRepositories**](PaginatedRepositories.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet(context.Background(), commit, key, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet`: Commitstatus
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet`: %v\n", resp)
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
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut(context.Background(), commit, key, repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut`: Commitstatus
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut`: %v\n", resp)
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
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost(context.Background(), commit, repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost`: Commitstatus
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost`: %v\n", resp)
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
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet(context.Background(), commit, repoSlug, workspace).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet`: PaginatedCommitstatuses
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet`: %v\n", resp)
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


## RepositoriesWorkspaceRepoSlugDelete

> RepositoriesWorkspaceRepoSlugDelete(ctx, repoSlug, workspace).RedirectTo(redirectTo).Execute()





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
    redirectTo := "redirectTo_example" // string | If a repository has been moved to a new location, use this parameter to show users a friendly message in the Bitbucket UI that the repository has moved to a new location. However, a GET to this endpoint will still return a 404.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugDelete(context.Background(), repoSlug, workspace).RedirectTo(redirectTo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **redirectTo** | **string** | If a repository has been moved to a new location, use this parameter to show users a friendly message in the Bitbucket UI that the repository has moved to a new location. However, a GET to this endpoint will still return a 404.  | 

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


## RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet

> PaginatedFiles RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet(ctx, commit, path, repoSlug, workspace).Renames(renames).Q(q).Sort(sort).Execute()





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
    path := "path_example" // string | Path to the file.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
    renames := "renames_example" // string |  When `true`, Bitbucket will follow the history of the file across renames (this is the default behavior). This can be turned off by specifying `false`. (optional)
    q := "q_example" // string |  Query string to narrow down the response as per [filtering and sorting](../../../../../../meta/filtering). (optional)
    sort := "sort_example" // string |  Name of a response property sort the result by as per [filtering and sorting](../../../../../../meta/filtering#query-sort).  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet(context.Background(), commit, path, repoSlug, workspace).Renames(renames).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet`: PaginatedFiles
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commit** | **string** | The commit&#39;s SHA1. | 
**path** | **string** | Path to the file. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugFilehistoryCommitPathGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **renames** | **string** |  When &#x60;true&#x60;, Bitbucket will follow the history of the file across renames (this is the default behavior). This can be turned off by specifying &#x60;false&#x60;. | 
 **q** | **string** |  Query string to narrow down the response as per [filtering and sorting](../../../../../../meta/filtering). | 
 **sort** | **string** |  Name of a response property sort the result by as per [filtering and sorting](../../../../../../meta/filtering#query-sort).  | 

### Return type

[**PaginatedFiles**](PaginatedFiles.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugForksGet

> PaginatedRepositories RepositoriesWorkspaceRepoSlugForksGet(ctx, repoSlug, workspace).Role(role).Q(q).Sort(sort).Execute()





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
    role := "role_example" // string | Filters the result based on the authenticated user's role on each repository.  * **member**: returns repositories to which the user has explicit read access * **contributor**: returns repositories to which the user has explicit write access * **admin**: returns repositories to which the user has explicit administrator access * **owner**: returns all repositories owned by the current user  (optional)
    q := "q_example" // string | Query string to narrow down the response as per [filtering and sorting](../../../../meta/filtering).  (optional)
    sort := "sort_example" // string | Field by which the results should be sorted as per [filtering and sorting](../../../../meta/filtering).  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugForksGet(context.Background(), repoSlug, workspace).Role(role).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugForksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugForksGet`: PaginatedRepositories
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugForksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugForksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **role** | **string** | Filters the result based on the authenticated user&#39;s role on each repository.  * **member**: returns repositories to which the user has explicit read access * **contributor**: returns repositories to which the user has explicit write access * **admin**: returns repositories to which the user has explicit administrator access * **owner**: returns all repositories owned by the current user  | 
 **q** | **string** | Query string to narrow down the response as per [filtering and sorting](../../../../meta/filtering).  | 
 **sort** | **string** | Field by which the results should be sorted as per [filtering and sorting](../../../../meta/filtering).  | 

### Return type

[**PaginatedRepositories**](PaginatedRepositories.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugForksPost

> Repository RepositoriesWorkspaceRepoSlugForksPost(ctx, repoSlug, workspace).Body(body).Execute()





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
    body := *openapiclient.NewRepository() // Repository | A repository object. This can be left blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugForksPost(context.Background(), repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugForksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugForksPost`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugForksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugForksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Repository**](Repository.md) | A repository object. This can be left blank. | 

### Return type

[**Repository**](Repository.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugGet

> Repository RepositoriesWorkspaceRepoSlugGet(ctx, repoSlug, workspace).Execute()





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
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugGet(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugGet`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Repository**](Repository.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugHooksGet

> PaginatedWebhookSubscriptions RepositoriesWorkspaceRepoSlugHooksGet(ctx, repoSlug, workspace).Execute()





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
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugHooksGet(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugHooksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugHooksGet`: PaginatedWebhookSubscriptions
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugHooksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugHooksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedWebhookSubscriptions**](PaginatedWebhookSubscriptions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugHooksPost

> WebhookSubscription RepositoriesWorkspaceRepoSlugHooksPost(ctx, repoSlug, workspace).Execute()





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
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugHooksPost(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugHooksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugHooksPost`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugHooksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugHooksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WebhookSubscription**](WebhookSubscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugHooksUidDelete

> RepositoriesWorkspaceRepoSlugHooksUidDelete(ctx, repoSlug, uid, workspace).Execute()





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
    uid := "uid_example" // string | Installed webhook's ID
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugHooksUidDelete(context.Background(), repoSlug, uid, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugHooksUidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**uid** | **string** | Installed webhook&#39;s ID | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugHooksUidDeleteRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugHooksUidGet

> WebhookSubscription RepositoriesWorkspaceRepoSlugHooksUidGet(ctx, repoSlug, uid, workspace).Execute()





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
    uid := "uid_example" // string | Installed webhook's ID
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugHooksUidGet(context.Background(), repoSlug, uid, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugHooksUidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugHooksUidGet`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugHooksUidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**uid** | **string** | Installed webhook&#39;s ID | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugHooksUidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WebhookSubscription**](WebhookSubscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugHooksUidPut

> WebhookSubscription RepositoriesWorkspaceRepoSlugHooksUidPut(ctx, repoSlug, uid, workspace).Execute()





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
    uid := "uid_example" // string | Installed webhook's ID
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugHooksUidPut(context.Background(), repoSlug, uid, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugHooksUidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugHooksUidPut`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugHooksUidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**uid** | **string** | Installed webhook&#39;s ID | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugHooksUidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WebhookSubscription**](WebhookSubscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugPost

> Repository RepositoriesWorkspaceRepoSlugPost(ctx, repoSlug, workspace).Body(body).Execute()





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
    body := *openapiclient.NewRepository() // Repository | The repository that is to be created. Note that most object elements are optional. Elements \"owner\" and \"full_name\" are ignored as the URL implies them. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugPost(context.Background(), repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPost`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Repository**](Repository.md) | The repository that is to be created. Note that most object elements are optional. Elements \&quot;owner\&quot; and \&quot;full_name\&quot; are ignored as the URL implies them. | 

### Return type

[**Repository**](Repository.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
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
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet(context.Background(), pullRequestId, repoSlug, workspace).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet`: PaginatedCommitstatuses
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet`: %v\n", resp)
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


## RepositoriesWorkspaceRepoSlugPut

> Repository RepositoriesWorkspaceRepoSlugPut(ctx, repoSlug, workspace).Body(body).Execute()





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
    body := *openapiclient.NewRepository() // Repository | The repository that is to be updated.  Note that the elements \"owner\" and \"full_name\" are ignored since the URL implies them.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugPut(context.Background(), repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPut`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Repository**](Repository.md) | The repository that is to be updated.  Note that the elements \&quot;owner\&quot; and \&quot;full_name\&quot; are ignored since the URL implies them.  | 

### Return type

[**Repository**](Repository.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugSrcCommitPathGet

> PaginatedTreeentries RepositoriesWorkspaceRepoSlugSrcCommitPathGet(ctx, commit, path, repoSlug, workspace).Format(format).Q(q).Sort(sort).MaxDepth(maxDepth).Execute()





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
    path := "path_example" // string | Path to the file.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
    format := "format_example" // string | If 'meta' is provided, returns the (json) meta data for the contents of the file.  If 'rendered' is provided, returns the contents of a non-binary file in HTML-formatted rendered markup. Since Git does not generally track what text encoding scheme is used, this endpoint attempts to detect the most appropriate character encoding. While usually correct, determining the character encoding can be ambiguous which in exceptional cases can lead to misinterpretation of the characters. As such, the raw element in the response object should not be treated as equivalent to the file's actual contents. (optional)
    q := "q_example" // string | Optional filter expression as per [filtering and sorting](../../../../../../meta/filtering). (optional)
    sort := "sort_example" // string | Optional sorting parameter as per [filtering and sorting](../../../../../../meta/filtering#query-sort). (optional)
    maxDepth := int32(56) // int32 | If provided, returns the contents of the repository and its subdirectories recursively until the specified max_depth of nested directories. When omitted, this defaults to 1. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugSrcCommitPathGet(context.Background(), commit, path, repoSlug, workspace).Format(format).Q(q).Sort(sort).MaxDepth(maxDepth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugSrcCommitPathGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugSrcCommitPathGet`: PaginatedTreeentries
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugSrcCommitPathGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commit** | **string** | The commit&#39;s SHA1. | 
**path** | **string** | Path to the file. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugSrcCommitPathGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **format** | **string** | If &#39;meta&#39; is provided, returns the (json) meta data for the contents of the file.  If &#39;rendered&#39; is provided, returns the contents of a non-binary file in HTML-formatted rendered markup. Since Git does not generally track what text encoding scheme is used, this endpoint attempts to detect the most appropriate character encoding. While usually correct, determining the character encoding can be ambiguous which in exceptional cases can lead to misinterpretation of the characters. As such, the raw element in the response object should not be treated as equivalent to the file&#39;s actual contents. | 
 **q** | **string** | Optional filter expression as per [filtering and sorting](../../../../../../meta/filtering). | 
 **sort** | **string** | Optional sorting parameter as per [filtering and sorting](../../../../../../meta/filtering#query-sort). | 
 **maxDepth** | **int32** | If provided, returns the contents of the repository and its subdirectories recursively until the specified max_depth of nested directories. When omitted, this defaults to 1. | 

### Return type

[**PaginatedTreeentries**](PaginatedTreeentries.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugSrcGet

> PaginatedTreeentries RepositoriesWorkspaceRepoSlugSrcGet(ctx, repoSlug, workspace).Format(format).Execute()





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
    format := "format_example" // string | Instead of returning the file's contents, return the (json) meta data for it. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugSrcGet(context.Background(), repoSlug, workspace).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugSrcGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugSrcGet`: PaginatedTreeentries
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesWorkspaceRepoSlugSrcGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugSrcGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** | Instead of returning the file&#39;s contents, return the (json) meta data for it. | 

### Return type

[**PaginatedTreeentries**](PaginatedTreeentries.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugSrcPost

> RepositoriesWorkspaceRepoSlugSrcPost(ctx, repoSlug, workspace).Message(message).Author(author).Parents(parents).Files(files).Branch(branch).Execute()





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
    message := "message_example" // string | The commit message. When omitted, Bitbucket uses a canned string. (optional)
    author := "author_example" // string |  The raw string to be used as the new commit's author. This string follows the format `Erik van Zijst <evzijst@atlassian.com>`.  When omitted, Bitbucket uses the authenticated user's full/display name and primary email address. Commits cannot be created anonymously. (optional)
    parents := "parents_example" // string |  A comma-separated list of SHA1s of the commits that should be the parents of the newly created commit.  When omitted, the new commit will inherit from and become a child of the main branch's tip/HEAD commit.  When more than one SHA1 is provided, the first SHA1 identifies the commit from which the content will be inherited.\". (optional)
    files := "files_example" // string |  Optional field that declares the files that the request is manipulating. When adding a new file to a repo, or when overwriting an existing file, the client can just upload the full contents of the file in a normal form field and the use of this `files` meta data field is redundant. However, when the `files` field contains a file path that does not have a corresponding, identically-named form field, then Bitbucket interprets that as the client wanting to replace the named file with the null set and the file is deleted instead.  Paths in the repo that are referenced in neither files nor an individual file field, remain unchanged and carry over from the parent to the new commit.  This API does not support renaming as an explicit feature. To rename a file, simply delete it and recreate it under the new name in the same commit.  (optional)
    branch := "branch_example" // string |  The name of the branch that the new commit should be created on. When omitted, the commit will be created on top of the main branch and will become the main branch's new head.  When a branch name is provided that already exists in the repo, then the commit will be created on top of that branch. In this case, *if* a parent SHA1 was also provided, then it is asserted that the parent is the branch's tip/HEAD at the time the request is made. When this is not the case, a 409 is returned.  When a new branch name is specified (that does not already exist in the repo), and no parent SHA1s are provided, then the new commit will inherit from the current main branch's tip/HEAD commit, but not advance the main branch. The new commit will be the new branch. When the request *also* specifies a parent SHA1, then the new commit and branch are created directly on top of the parent commit, regardless of the state of the main branch.  When a branch name is not specified, but a parent SHA1 is provided, then Bitbucket asserts that it represents the main branch's current HEAD/tip, or a 409 is returned.  When a branch name is not specified and the repo is empty, the new commit will become the repo's root commit and will be on the main branch.  When a branch name is specified and the repo is empty, the new commit will become the repo's root commit and also define the repo's main branch going forward.  This API cannot be used to create additional root commits in non-empty repos.  The branch field cannot be repeated.  As a side effect, this API can be used to create a new branch without modifying any files, by specifying a new branch name in this field, together with `parents`, but omitting the `files` fields, while not sending any files. This will create a new commit and branch with the same contents as the first parent. The diff of this commit against its first parent will be empty.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugSrcPost(context.Background(), repoSlug, workspace).Message(message).Author(author).Parents(parents).Files(files).Branch(branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugSrcPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugSrcPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **message** | **string** | The commit message. When omitted, Bitbucket uses a canned string. | 
 **author** | **string** |  The raw string to be used as the new commit&#39;s author. This string follows the format &#x60;Erik van Zijst &lt;evzijst@atlassian.com&gt;&#x60;.  When omitted, Bitbucket uses the authenticated user&#39;s full/display name and primary email address. Commits cannot be created anonymously. | 
 **parents** | **string** |  A comma-separated list of SHA1s of the commits that should be the parents of the newly created commit.  When omitted, the new commit will inherit from and become a child of the main branch&#39;s tip/HEAD commit.  When more than one SHA1 is provided, the first SHA1 identifies the commit from which the content will be inherited.\&quot;. | 
 **files** | **string** |  Optional field that declares the files that the request is manipulating. When adding a new file to a repo, or when overwriting an existing file, the client can just upload the full contents of the file in a normal form field and the use of this &#x60;files&#x60; meta data field is redundant. However, when the &#x60;files&#x60; field contains a file path that does not have a corresponding, identically-named form field, then Bitbucket interprets that as the client wanting to replace the named file with the null set and the file is deleted instead.  Paths in the repo that are referenced in neither files nor an individual file field, remain unchanged and carry over from the parent to the new commit.  This API does not support renaming as an explicit feature. To rename a file, simply delete it and recreate it under the new name in the same commit.  | 
 **branch** | **string** |  The name of the branch that the new commit should be created on. When omitted, the commit will be created on top of the main branch and will become the main branch&#39;s new head.  When a branch name is provided that already exists in the repo, then the commit will be created on top of that branch. In this case, *if* a parent SHA1 was also provided, then it is asserted that the parent is the branch&#39;s tip/HEAD at the time the request is made. When this is not the case, a 409 is returned.  When a new branch name is specified (that does not already exist in the repo), and no parent SHA1s are provided, then the new commit will inherit from the current main branch&#39;s tip/HEAD commit, but not advance the main branch. The new commit will be the new branch. When the request *also* specifies a parent SHA1, then the new commit and branch are created directly on top of the parent commit, regardless of the state of the main branch.  When a branch name is not specified, but a parent SHA1 is provided, then Bitbucket asserts that it represents the main branch&#39;s current HEAD/tip, or a 409 is returned.  When a branch name is not specified and the repo is empty, the new commit will become the repo&#39;s root commit and will be on the main branch.  When a branch name is specified and the repo is empty, the new commit will become the repo&#39;s root commit and also define the repo&#39;s main branch going forward.  This API cannot be used to create additional root commits in non-empty repos.  The branch field cannot be repeated.  As a side effect, this API can be used to create a new branch without modifying any files, by specifying a new branch name in this field, together with &#x60;parents&#x60;, but omitting the &#x60;files&#x60; fields, while not sending any files. This will create a new commit and branch with the same contents as the first parent. The diff of this commit against its first parent will be empty.  | 

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


## RepositoriesWorkspaceRepoSlugWatchersGet

> RepositoriesWorkspaceRepoSlugWatchersGet(ctx, repoSlug, workspace).Execute()





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
    resp, r, err := api_client.RepositoriesApi.RepositoriesWorkspaceRepoSlugWatchersGet(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesWorkspaceRepoSlugWatchersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugWatchersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserPermissionsRepositoriesGet

> PaginatedRepositoryPermissions UserPermissionsRepositoriesGet(ctx).Q(q).Sort(sort).Execute()





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
    q := "q_example" // string |  Query string to narrow down the response as per [filtering and sorting](../../../meta/filtering). (optional)
    sort := "sort_example" // string |  Name of a response property sort the result by as per [filtering and sorting](../../../meta/filtering#query-sort).  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.UserPermissionsRepositoriesGet(context.Background()).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.UserPermissionsRepositoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserPermissionsRepositoriesGet`: PaginatedRepositoryPermissions
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.UserPermissionsRepositoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserPermissionsRepositoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  Query string to narrow down the response as per [filtering and sorting](../../../meta/filtering). | 
 **sort** | **string** |  Name of a response property sort the result by as per [filtering and sorting](../../../meta/filtering#query-sort).  | 

### Return type

[**PaginatedRepositoryPermissions**](PaginatedRepositoryPermissions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

