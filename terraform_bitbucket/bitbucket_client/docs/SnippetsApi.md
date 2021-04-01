# \SnippetsApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SnippetsGet**](SnippetsApi.md#SnippetsGet) | **Get** /snippets | 
[**SnippetsPost**](SnippetsApi.md#SnippetsPost) | **Post** /snippets | 
[**SnippetsWorkspaceEncodedIdCommentsCommentIdDelete**](SnippetsApi.md#SnippetsWorkspaceEncodedIdCommentsCommentIdDelete) | **Delete** /snippets/{workspace}/{encoded_id}/comments/{comment_id} | 
[**SnippetsWorkspaceEncodedIdCommentsCommentIdGet**](SnippetsApi.md#SnippetsWorkspaceEncodedIdCommentsCommentIdGet) | **Get** /snippets/{workspace}/{encoded_id}/comments/{comment_id} | 
[**SnippetsWorkspaceEncodedIdCommentsCommentIdPut**](SnippetsApi.md#SnippetsWorkspaceEncodedIdCommentsCommentIdPut) | **Put** /snippets/{workspace}/{encoded_id}/comments/{comment_id} | 
[**SnippetsWorkspaceEncodedIdCommentsGet**](SnippetsApi.md#SnippetsWorkspaceEncodedIdCommentsGet) | **Get** /snippets/{workspace}/{encoded_id}/comments | 
[**SnippetsWorkspaceEncodedIdCommentsPost**](SnippetsApi.md#SnippetsWorkspaceEncodedIdCommentsPost) | **Post** /snippets/{workspace}/{encoded_id}/comments | 
[**SnippetsWorkspaceEncodedIdCommitsGet**](SnippetsApi.md#SnippetsWorkspaceEncodedIdCommitsGet) | **Get** /snippets/{workspace}/{encoded_id}/commits | 
[**SnippetsWorkspaceEncodedIdCommitsRevisionGet**](SnippetsApi.md#SnippetsWorkspaceEncodedIdCommitsRevisionGet) | **Get** /snippets/{workspace}/{encoded_id}/commits/{revision} | 
[**SnippetsWorkspaceEncodedIdDelete**](SnippetsApi.md#SnippetsWorkspaceEncodedIdDelete) | **Delete** /snippets/{workspace}/{encoded_id} | 
[**SnippetsWorkspaceEncodedIdFilesPathGet**](SnippetsApi.md#SnippetsWorkspaceEncodedIdFilesPathGet) | **Get** /snippets/{workspace}/{encoded_id}/files/{path} | 
[**SnippetsWorkspaceEncodedIdGet**](SnippetsApi.md#SnippetsWorkspaceEncodedIdGet) | **Get** /snippets/{workspace}/{encoded_id} | 
[**SnippetsWorkspaceEncodedIdNodeIdDelete**](SnippetsApi.md#SnippetsWorkspaceEncodedIdNodeIdDelete) | **Delete** /snippets/{workspace}/{encoded_id}/{node_id} | 
[**SnippetsWorkspaceEncodedIdNodeIdFilesPathGet**](SnippetsApi.md#SnippetsWorkspaceEncodedIdNodeIdFilesPathGet) | **Get** /snippets/{workspace}/{encoded_id}/{node_id}/files/{path} | 
[**SnippetsWorkspaceEncodedIdNodeIdGet**](SnippetsApi.md#SnippetsWorkspaceEncodedIdNodeIdGet) | **Get** /snippets/{workspace}/{encoded_id}/{node_id} | 
[**SnippetsWorkspaceEncodedIdNodeIdPut**](SnippetsApi.md#SnippetsWorkspaceEncodedIdNodeIdPut) | **Put** /snippets/{workspace}/{encoded_id}/{node_id} | 
[**SnippetsWorkspaceEncodedIdPut**](SnippetsApi.md#SnippetsWorkspaceEncodedIdPut) | **Put** /snippets/{workspace}/{encoded_id} | 
[**SnippetsWorkspaceEncodedIdRevisionDiffGet**](SnippetsApi.md#SnippetsWorkspaceEncodedIdRevisionDiffGet) | **Get** /snippets/{workspace}/{encoded_id}/{revision}/diff | 
[**SnippetsWorkspaceEncodedIdRevisionPatchGet**](SnippetsApi.md#SnippetsWorkspaceEncodedIdRevisionPatchGet) | **Get** /snippets/{workspace}/{encoded_id}/{revision}/patch | 
[**SnippetsWorkspaceEncodedIdWatchDelete**](SnippetsApi.md#SnippetsWorkspaceEncodedIdWatchDelete) | **Delete** /snippets/{workspace}/{encoded_id}/watch | 
[**SnippetsWorkspaceEncodedIdWatchGet**](SnippetsApi.md#SnippetsWorkspaceEncodedIdWatchGet) | **Get** /snippets/{workspace}/{encoded_id}/watch | 
[**SnippetsWorkspaceEncodedIdWatchPut**](SnippetsApi.md#SnippetsWorkspaceEncodedIdWatchPut) | **Put** /snippets/{workspace}/{encoded_id}/watch | 
[**SnippetsWorkspaceEncodedIdWatchersGet**](SnippetsApi.md#SnippetsWorkspaceEncodedIdWatchersGet) | **Get** /snippets/{workspace}/{encoded_id}/watchers | 
[**SnippetsWorkspaceGet**](SnippetsApi.md#SnippetsWorkspaceGet) | **Get** /snippets/{workspace} | 
[**SnippetsWorkspacePost**](SnippetsApi.md#SnippetsWorkspacePost) | **Post** /snippets/{workspace} | 



## SnippetsGet

> PaginatedSnippets SnippetsGet(ctx).Role(role).Execute()





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
    role := "role_example" // string | Filter down the result based on the authenticated user's role (`owner`, `contributor`, or `member`). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsGet(context.Background()).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsGet`: PaginatedSnippets
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **string** | Filter down the result based on the authenticated user&#39;s role (&#x60;owner&#x60;, &#x60;contributor&#x60;, or &#x60;member&#x60;). | 

### Return type

[**PaginatedSnippets**](PaginatedSnippets.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsPost

> Snippet SnippetsPost(ctx).Body(body).Execute()





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
    body := *openapiclient.NewSnippet() // Snippet | The new snippet object.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsPost`: Snippet
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Snippet**](Snippet.md) | The new snippet object. | 

### Return type

[**Snippet**](Snippet.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspaceEncodedIdCommentsCommentIdDelete

> SnippetsWorkspaceEncodedIdCommentsCommentIdDelete(ctx, commentId, encodedId, workspace).Execute()





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
    commentId := int32(56) // int32 | The id of the comment.
    encodedId := "encodedId_example" // string | The snippet id.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdCommentsCommentIdDelete(context.Background(), commentId, encodedId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdCommentsCommentIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** | The id of the comment. | 
**encodedId** | **string** | The snippet id. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdCommentsCommentIdDeleteRequest struct via the builder pattern


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


## SnippetsWorkspaceEncodedIdCommentsCommentIdGet

> SnippetComment SnippetsWorkspaceEncodedIdCommentsCommentIdGet(ctx, commentId, encodedId, workspace).Execute()





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
    commentId := int32(56) // int32 | The id of the comment.
    encodedId := "encodedId_example" // string | The snippet id.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdCommentsCommentIdGet(context.Background(), commentId, encodedId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdCommentsCommentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspaceEncodedIdCommentsCommentIdGet`: SnippetComment
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspaceEncodedIdCommentsCommentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** | The id of the comment. | 
**encodedId** | **string** | The snippet id. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdCommentsCommentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SnippetComment**](SnippetComment.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspaceEncodedIdCommentsCommentIdPut

> SnippetsWorkspaceEncodedIdCommentsCommentIdPut(ctx, commentId, encodedId, workspace).Execute()





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
    commentId := int32(56) // int32 | The id of the comment.
    encodedId := "encodedId_example" // string | The snippet id.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdCommentsCommentIdPut(context.Background(), commentId, encodedId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdCommentsCommentIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** | The id of the comment. | 
**encodedId** | **string** | The snippet id. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdCommentsCommentIdPutRequest struct via the builder pattern


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


## SnippetsWorkspaceEncodedIdCommentsGet

> PaginatedSnippetComments SnippetsWorkspaceEncodedIdCommentsGet(ctx, encodedId, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdCommentsGet(context.Background(), encodedId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdCommentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspaceEncodedIdCommentsGet`: PaginatedSnippetComments
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspaceEncodedIdCommentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdCommentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedSnippetComments**](PaginatedSnippetComments.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspaceEncodedIdCommentsPost

> Snippet SnippetsWorkspaceEncodedIdCommentsPost(ctx, encodedId, workspace).Body(body).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
    body := *openapiclient.NewSnippet() // Snippet | The contents of the new comment.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdCommentsPost(context.Background(), encodedId, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdCommentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspaceEncodedIdCommentsPost`: Snippet
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspaceEncodedIdCommentsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdCommentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Snippet**](Snippet.md) | The contents of the new comment. | 

### Return type

[**Snippet**](Snippet.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspaceEncodedIdCommitsGet

> PaginatedSnippetCommit SnippetsWorkspaceEncodedIdCommitsGet(ctx, encodedId, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdCommitsGet(context.Background(), encodedId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdCommitsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspaceEncodedIdCommitsGet`: PaginatedSnippetCommit
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspaceEncodedIdCommitsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdCommitsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedSnippetCommit**](PaginatedSnippetCommit.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspaceEncodedIdCommitsRevisionGet

> SnippetCommit SnippetsWorkspaceEncodedIdCommitsRevisionGet(ctx, encodedId, revision, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    revision := "revision_example" // string | The commit's SHA1.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdCommitsRevisionGet(context.Background(), encodedId, revision, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdCommitsRevisionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspaceEncodedIdCommitsRevisionGet`: SnippetCommit
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspaceEncodedIdCommitsRevisionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**revision** | **string** | The commit&#39;s SHA1. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdCommitsRevisionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SnippetCommit**](SnippetCommit.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspaceEncodedIdDelete

> SnippetsWorkspaceEncodedIdDelete(ctx, encodedId, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdDelete(context.Background(), encodedId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdDeleteRequest struct via the builder pattern


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


## SnippetsWorkspaceEncodedIdFilesPathGet

> SnippetsWorkspaceEncodedIdFilesPathGet(ctx, encodedId, path, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    path := "path_example" // string | Path to the file.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdFilesPathGet(context.Background(), encodedId, path, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdFilesPathGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**path** | **string** | Path to the file. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdFilesPathGetRequest struct via the builder pattern


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


## SnippetsWorkspaceEncodedIdGet

> Snippet SnippetsWorkspaceEncodedIdGet(ctx, encodedId, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdGet(context.Background(), encodedId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspaceEncodedIdGet`: Snippet
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspaceEncodedIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Snippet**](Snippet.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/related, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspaceEncodedIdNodeIdDelete

> SnippetsWorkspaceEncodedIdNodeIdDelete(ctx, encodedId, nodeId, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    nodeId := "nodeId_example" // string | A commit revision (SHA1).
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdNodeIdDelete(context.Background(), encodedId, nodeId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdNodeIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**nodeId** | **string** | A commit revision (SHA1). | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdNodeIdDeleteRequest struct via the builder pattern


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


## SnippetsWorkspaceEncodedIdNodeIdFilesPathGet

> SnippetsWorkspaceEncodedIdNodeIdFilesPathGet(ctx, encodedId, nodeId, path, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    nodeId := "nodeId_example" // string | A commit revision (SHA1).
    path := "path_example" // string | Path to the file.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdNodeIdFilesPathGet(context.Background(), encodedId, nodeId, path, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdNodeIdFilesPathGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**nodeId** | **string** | A commit revision (SHA1). | 
**path** | **string** | Path to the file. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdNodeIdFilesPathGetRequest struct via the builder pattern


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


## SnippetsWorkspaceEncodedIdNodeIdGet

> Snippet SnippetsWorkspaceEncodedIdNodeIdGet(ctx, encodedId, nodeId, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    nodeId := "nodeId_example" // string | A commit revision (SHA1).
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdNodeIdGet(context.Background(), encodedId, nodeId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdNodeIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspaceEncodedIdNodeIdGet`: Snippet
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspaceEncodedIdNodeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**nodeId** | **string** | A commit revision (SHA1). | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdNodeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Snippet**](Snippet.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/related, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspaceEncodedIdNodeIdPut

> Snippet SnippetsWorkspaceEncodedIdNodeIdPut(ctx, encodedId, nodeId, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    nodeId := "nodeId_example" // string | A commit revision (SHA1).
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdNodeIdPut(context.Background(), encodedId, nodeId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdNodeIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspaceEncodedIdNodeIdPut`: Snippet
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspaceEncodedIdNodeIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**nodeId** | **string** | A commit revision (SHA1). | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdNodeIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Snippet**](Snippet.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/related, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspaceEncodedIdPut

> Snippet SnippetsWorkspaceEncodedIdPut(ctx, encodedId, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdPut(context.Background(), encodedId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspaceEncodedIdPut`: Snippet
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspaceEncodedIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Snippet**](Snippet.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/related, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspaceEncodedIdRevisionDiffGet

> SnippetsWorkspaceEncodedIdRevisionDiffGet(ctx, encodedId, revision, workspace).Path(path).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    revision := "revision_example" // string | A revspec expression. This can simply be a commit SHA1, a ref name, or a compare expression like `staging..production`.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
    path := "path_example" // string | When used, only one the diff of the specified file will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdRevisionDiffGet(context.Background(), encodedId, revision, workspace).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdRevisionDiffGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**revision** | **string** | A revspec expression. This can simply be a commit SHA1, a ref name, or a compare expression like &#x60;staging..production&#x60;. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdRevisionDiffGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **path** | **string** | When used, only one the diff of the specified file will be returned. | 

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


## SnippetsWorkspaceEncodedIdRevisionPatchGet

> SnippetsWorkspaceEncodedIdRevisionPatchGet(ctx, encodedId, revision, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    revision := "revision_example" // string | A revspec expression. This can simply be a commit SHA1, a ref name, or a compare expression like `staging..production`.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdRevisionPatchGet(context.Background(), encodedId, revision, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdRevisionPatchGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**revision** | **string** | A revspec expression. This can simply be a commit SHA1, a ref name, or a compare expression like &#x60;staging..production&#x60;. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdRevisionPatchGetRequest struct via the builder pattern


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


## SnippetsWorkspaceEncodedIdWatchDelete

> PaginatedUsers SnippetsWorkspaceEncodedIdWatchDelete(ctx, encodedId, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdWatchDelete(context.Background(), encodedId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdWatchDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspaceEncodedIdWatchDelete`: PaginatedUsers
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspaceEncodedIdWatchDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdWatchDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedUsers**](PaginatedUsers.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspaceEncodedIdWatchGet

> PaginatedUsers SnippetsWorkspaceEncodedIdWatchGet(ctx, encodedId, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdWatchGet(context.Background(), encodedId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdWatchGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspaceEncodedIdWatchGet`: PaginatedUsers
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspaceEncodedIdWatchGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdWatchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedUsers**](PaginatedUsers.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspaceEncodedIdWatchPut

> PaginatedUsers SnippetsWorkspaceEncodedIdWatchPut(ctx, encodedId, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdWatchPut(context.Background(), encodedId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdWatchPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspaceEncodedIdWatchPut`: PaginatedUsers
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspaceEncodedIdWatchPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdWatchPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedUsers**](PaginatedUsers.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspaceEncodedIdWatchersGet

> PaginatedUsers SnippetsWorkspaceEncodedIdWatchersGet(ctx, encodedId, workspace).Execute()





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
    encodedId := "encodedId_example" // string | The snippet id.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceEncodedIdWatchersGet(context.Background(), encodedId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceEncodedIdWatchersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspaceEncodedIdWatchersGet`: PaginatedUsers
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspaceEncodedIdWatchersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedId** | **string** | The snippet id. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceEncodedIdWatchersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedUsers**](PaginatedUsers.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspaceGet

> PaginatedSnippets SnippetsWorkspaceGet(ctx, workspace).Role(role).Execute()





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
    role := "role_example" // string | Filter down the result based on the authenticated user's role (`owner`, `contributor`, or `member`). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspaceGet(context.Background(), workspace).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspaceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspaceGet`: PaginatedSnippets
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspaceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspaceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | **string** | Filter down the result based on the authenticated user&#39;s role (&#x60;owner&#x60;, &#x60;contributor&#x60;, or &#x60;member&#x60;). | 

### Return type

[**PaginatedSnippets**](PaginatedSnippets.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnippetsWorkspacePost

> Snippet SnippetsWorkspacePost(ctx, workspace).Body(body).Execute()





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
    body := *openapiclient.NewSnippet() // Snippet | The new snippet object.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnippetsApi.SnippetsWorkspacePost(context.Background(), workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetsApi.SnippetsWorkspacePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnippetsWorkspacePost`: Snippet
    fmt.Fprintf(os.Stdout, "Response from `SnippetsApi.SnippetsWorkspacePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnippetsWorkspacePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Snippet**](Snippet.md) | The new snippet object. | 

### Return type

[**Snippet**](Snippet.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

