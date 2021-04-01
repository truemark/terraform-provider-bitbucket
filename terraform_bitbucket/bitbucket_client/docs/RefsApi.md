# \RefsApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesWorkspaceRepoSlugRefsBranchesGet**](RefsApi.md#RepositoriesWorkspaceRepoSlugRefsBranchesGet) | **Get** /repositories/{workspace}/{repo_slug}/refs/branches | 
[**RepositoriesWorkspaceRepoSlugRefsBranchesNameDelete**](RefsApi.md#RepositoriesWorkspaceRepoSlugRefsBranchesNameDelete) | **Delete** /repositories/{workspace}/{repo_slug}/refs/branches/{name} | 
[**RepositoriesWorkspaceRepoSlugRefsBranchesNameGet**](RefsApi.md#RepositoriesWorkspaceRepoSlugRefsBranchesNameGet) | **Get** /repositories/{workspace}/{repo_slug}/refs/branches/{name} | 
[**RepositoriesWorkspaceRepoSlugRefsBranchesPost**](RefsApi.md#RepositoriesWorkspaceRepoSlugRefsBranchesPost) | **Post** /repositories/{workspace}/{repo_slug}/refs/branches | 
[**RepositoriesWorkspaceRepoSlugRefsGet**](RefsApi.md#RepositoriesWorkspaceRepoSlugRefsGet) | **Get** /repositories/{workspace}/{repo_slug}/refs | 
[**RepositoriesWorkspaceRepoSlugRefsTagsGet**](RefsApi.md#RepositoriesWorkspaceRepoSlugRefsTagsGet) | **Get** /repositories/{workspace}/{repo_slug}/refs/tags | 
[**RepositoriesWorkspaceRepoSlugRefsTagsNameDelete**](RefsApi.md#RepositoriesWorkspaceRepoSlugRefsTagsNameDelete) | **Delete** /repositories/{workspace}/{repo_slug}/refs/tags/{name} | 
[**RepositoriesWorkspaceRepoSlugRefsTagsNameGet**](RefsApi.md#RepositoriesWorkspaceRepoSlugRefsTagsNameGet) | **Get** /repositories/{workspace}/{repo_slug}/refs/tags/{name} | 
[**RepositoriesWorkspaceRepoSlugRefsTagsPost**](RefsApi.md#RepositoriesWorkspaceRepoSlugRefsTagsPost) | **Post** /repositories/{workspace}/{repo_slug}/refs/tags | 



## RepositoriesWorkspaceRepoSlugRefsBranchesGet

> PaginatedBranches RepositoriesWorkspaceRepoSlugRefsBranchesGet(ctx, repoSlug, workspace).Q(q).Sort(sort).Execute()





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
    q := "q_example" // string |  Query string to narrow down the response as per [filtering and sorting](../../../../../meta/filtering). (optional)
    sort := "sort_example" // string |  Field by which the results should be sorted as per [filtering and sorting](../../../../../meta/filtering). The `name` field is handled specially for branches in that, if specified as the sort field, it uses a natural sort order instead of the default lexicographical sort order. For example, it will return ['branch1', 'branch2', 'branch10'] instead of ['branch1', 'branch10', 'branch2']. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RefsApi.RepositoriesWorkspaceRepoSlugRefsBranchesGet(context.Background(), repoSlug, workspace).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefsApi.RepositoriesWorkspaceRepoSlugRefsBranchesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugRefsBranchesGet`: PaginatedBranches
    fmt.Fprintf(os.Stdout, "Response from `RefsApi.RepositoriesWorkspaceRepoSlugRefsBranchesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugRefsBranchesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **string** |  Query string to narrow down the response as per [filtering and sorting](../../../../../meta/filtering). | 
 **sort** | **string** |  Field by which the results should be sorted as per [filtering and sorting](../../../../../meta/filtering). The &#x60;name&#x60; field is handled specially for branches in that, if specified as the sort field, it uses a natural sort order instead of the default lexicographical sort order. For example, it will return [&#39;branch1&#39;, &#39;branch2&#39;, &#39;branch10&#39;] instead of [&#39;branch1&#39;, &#39;branch10&#39;, &#39;branch2&#39;]. | 

### Return type

[**PaginatedBranches**](PaginatedBranches.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugRefsBranchesNameDelete

> RepositoriesWorkspaceRepoSlugRefsBranchesNameDelete(ctx, name, repoSlug, workspace).Execute()





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
    name := "name_example" // string | The name of the branch.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RefsApi.RepositoriesWorkspaceRepoSlugRefsBranchesNameDelete(context.Background(), name, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefsApi.RepositoriesWorkspaceRepoSlugRefsBranchesNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the branch. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugRefsBranchesNameDeleteRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugRefsBranchesNameGet

> Branch RepositoriesWorkspaceRepoSlugRefsBranchesNameGet(ctx, name, repoSlug, workspace).Execute()





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
    name := "name_example" // string | The name of the branch.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RefsApi.RepositoriesWorkspaceRepoSlugRefsBranchesNameGet(context.Background(), name, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefsApi.RepositoriesWorkspaceRepoSlugRefsBranchesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugRefsBranchesNameGet`: Branch
    fmt.Fprintf(os.Stdout, "Response from `RefsApi.RepositoriesWorkspaceRepoSlugRefsBranchesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the branch. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugRefsBranchesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Branch**](Branch.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugRefsBranchesPost

> Branch RepositoriesWorkspaceRepoSlugRefsBranchesPost(ctx, repoSlug, workspace).Execute()





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
    resp, r, err := api_client.RefsApi.RepositoriesWorkspaceRepoSlugRefsBranchesPost(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefsApi.RepositoriesWorkspaceRepoSlugRefsBranchesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugRefsBranchesPost`: Branch
    fmt.Fprintf(os.Stdout, "Response from `RefsApi.RepositoriesWorkspaceRepoSlugRefsBranchesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugRefsBranchesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Branch**](Branch.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugRefsGet

> PaginatedRefs RepositoriesWorkspaceRepoSlugRefsGet(ctx, repoSlug, workspace).Q(q).Sort(sort).Execute()





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
    q := "q_example" // string |  Query string to narrow down the response as per [filtering and sorting](../../../../meta/filtering). (optional)
    sort := "sort_example" // string |  Field by which the results should be sorted as per [filtering and sorting](../../../../meta/filtering). The `name` field is handled specially for refs in that, if specified as the sort field, it uses a natural sort order instead of the default lexicographical sort order. For example, it will return ['1.1', '1.2', '1.10'] instead of ['1.1', '1.10', '1.2']. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RefsApi.RepositoriesWorkspaceRepoSlugRefsGet(context.Background(), repoSlug, workspace).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefsApi.RepositoriesWorkspaceRepoSlugRefsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugRefsGet`: PaginatedRefs
    fmt.Fprintf(os.Stdout, "Response from `RefsApi.RepositoriesWorkspaceRepoSlugRefsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugRefsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **string** |  Query string to narrow down the response as per [filtering and sorting](../../../../meta/filtering). | 
 **sort** | **string** |  Field by which the results should be sorted as per [filtering and sorting](../../../../meta/filtering). The &#x60;name&#x60; field is handled specially for refs in that, if specified as the sort field, it uses a natural sort order instead of the default lexicographical sort order. For example, it will return [&#39;1.1&#39;, &#39;1.2&#39;, &#39;1.10&#39;] instead of [&#39;1.1&#39;, &#39;1.10&#39;, &#39;1.2&#39;]. | 

### Return type

[**PaginatedRefs**](PaginatedRefs.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugRefsTagsGet

> PaginatedTags RepositoriesWorkspaceRepoSlugRefsTagsGet(ctx, repoSlug, workspace).Q(q).Sort(sort).Execute()





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
    q := "q_example" // string |  Query string to narrow down the response as per [filtering and sorting](../../../../../meta/filtering). (optional)
    sort := "sort_example" // string |  Field by which the results should be sorted as per [filtering and sorting](../../../../../meta/filtering). The `name` field is handled specially for tags in that, if specified as the sort field, it uses a natural sort order instead of the default lexicographical sort order. For example, it will return ['1.1', '1.2', '1.10'] instead of ['1.1', '1.10', '1.2']. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RefsApi.RepositoriesWorkspaceRepoSlugRefsTagsGet(context.Background(), repoSlug, workspace).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefsApi.RepositoriesWorkspaceRepoSlugRefsTagsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugRefsTagsGet`: PaginatedTags
    fmt.Fprintf(os.Stdout, "Response from `RefsApi.RepositoriesWorkspaceRepoSlugRefsTagsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugRefsTagsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **string** |  Query string to narrow down the response as per [filtering and sorting](../../../../../meta/filtering). | 
 **sort** | **string** |  Field by which the results should be sorted as per [filtering and sorting](../../../../../meta/filtering). The &#x60;name&#x60; field is handled specially for tags in that, if specified as the sort field, it uses a natural sort order instead of the default lexicographical sort order. For example, it will return [&#39;1.1&#39;, &#39;1.2&#39;, &#39;1.10&#39;] instead of [&#39;1.1&#39;, &#39;1.10&#39;, &#39;1.2&#39;]. | 

### Return type

[**PaginatedTags**](PaginatedTags.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugRefsTagsNameDelete

> RepositoriesWorkspaceRepoSlugRefsTagsNameDelete(ctx, name, repoSlug, workspace).Execute()





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
    name := "name_example" // string | The name of the tag.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RefsApi.RepositoriesWorkspaceRepoSlugRefsTagsNameDelete(context.Background(), name, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefsApi.RepositoriesWorkspaceRepoSlugRefsTagsNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the tag. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugRefsTagsNameDeleteRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugRefsTagsNameGet

> Tag RepositoriesWorkspaceRepoSlugRefsTagsNameGet(ctx, name, repoSlug, workspace).Execute()





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
    name := "name_example" // string | The name of the tag.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RefsApi.RepositoriesWorkspaceRepoSlugRefsTagsNameGet(context.Background(), name, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefsApi.RepositoriesWorkspaceRepoSlugRefsTagsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugRefsTagsNameGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `RefsApi.RepositoriesWorkspaceRepoSlugRefsTagsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the tag. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugRefsTagsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Tag**](Tag.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugRefsTagsPost

> Tag RepositoriesWorkspaceRepoSlugRefsTagsPost(ctx, repoSlug, workspace).Body(body).Execute()





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
    body := *openapiclient.NewTag() // Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RefsApi.RepositoriesWorkspaceRepoSlugRefsTagsPost(context.Background(), repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefsApi.RepositoriesWorkspaceRepoSlugRefsTagsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugRefsTagsPost`: Tag
    fmt.Fprintf(os.Stdout, "Response from `RefsApi.RepositoriesWorkspaceRepoSlugRefsTagsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugRefsTagsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Tag**](Tag.md) |  | 

### Return type

[**Tag**](Tag.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

