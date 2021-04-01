# \PullrequestsApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPullrequestsForCommit**](PullrequestsApi.md#GetPullrequestsForCommit) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/pullrequests | Returns a paginated list of all pull requests as part of which this commit was reviewed. Pull Request Commit Links app must be installed first before using this API; installation automatically occurs when &#39;Go to pull request&#39; is clicked from the web interface for a commit&#39;s details.
[**PullrequestsSelectedUserGet**](PullrequestsApi.md#PullrequestsSelectedUserGet) | **Get** /pullrequests/{selected_user} | 
[**RepositoriesWorkspaceRepoSlugDefaultReviewersGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugDefaultReviewersGet) | **Get** /repositories/{workspace}/{repo_slug}/default-reviewers | 
[**RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameDelete**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameDelete) | **Delete** /repositories/{workspace}/{repo_slug}/default-reviewers/{target_username} | 
[**RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameGet) | **Get** /repositories/{workspace}/{repo_slug}/default-reviewers/{target_username} | 
[**RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernamePut**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernamePut) | **Put** /repositories/{workspace}/{repo_slug}/default-reviewers/{target_username} | 
[**RepositoriesWorkspaceRepoSlugPullrequestsActivityGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsActivityGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/activity | 
[**RepositoriesWorkspaceRepoSlugPullrequestsGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPost**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPost) | **Post** /repositories/{workspace}/{repo_slug}/pullrequests | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdActivityGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdActivityGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/activity | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApproveDelete**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApproveDelete) | **Delete** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/approve | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApprovePost**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApprovePost) | **Post** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/approve | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdDelete**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdDelete) | **Delete** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/comments/{comment_id} | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/comments/{comment_id} | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdPut**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdPut) | **Put** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/comments/{comment_id} | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/comments | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsPost**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsPost) | **Post** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/comments | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommitsGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommitsGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/commits | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDeclinePost**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDeclinePost) | **Post** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/decline | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/diff | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffstatGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffstatGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/diffstat | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id} | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePost**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePost) | **Post** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/merge | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeTaskStatusTaskIdGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeTaskStatusTaskIdGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/merge/task-status/{task_id} | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPatchGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPatchGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/patch | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPut**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPut) | **Put** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id} | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesDelete**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesDelete) | **Delete** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/request-changes | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesPost**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesPost) | **Post** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/request-changes | 
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/statuses | 



## GetPullrequestsForCommit

> PaginatedPullrequests GetPullrequestsForCommit(ctx, workspace, repoSlug, commit).Page(page).Pagelen(pagelen).Execute()

Returns a paginated list of all pull requests as part of which this commit was reviewed. Pull Request Commit Links app must be installed first before using this API; installation automatically occurs when 'Go to pull request' is clicked from the web interface for a commit's details.

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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces
    repoSlug := "repoSlug_example" // string | The repository; either the UUID in curly braces, or the slug
    commit := "commit_example" // string | The SHA1 of the commit
    page := int32(56) // int32 | Which page to retrieve (optional) (default to 1)
    pagelen := int32(56) // int32 | How many pull requests to retrieve per page (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.GetPullrequestsForCommit(context.Background(), workspace, repoSlug, commit).Page(page).Pagelen(pagelen).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.GetPullrequestsForCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPullrequestsForCommit`: PaginatedPullrequests
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.GetPullrequestsForCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces | 
**repoSlug** | **string** | The repository; either the UUID in curly braces, or the slug | 
**commit** | **string** | The SHA1 of the commit | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPullrequestsForCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Which page to retrieve | [default to 1]
 **pagelen** | **int32** | How many pull requests to retrieve per page | [default to 30]

### Return type

[**PaginatedPullrequests**](PaginatedPullrequests.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullrequestsSelectedUserGet

> PaginatedPullrequests PullrequestsSelectedUserGet(ctx, selectedUser).State(state).Execute()





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
    selectedUser := "selectedUser_example" // string | This can either be the username of the pull request author, the author's UUID surrounded by curly-braces, for example: `{account UUID}`, or the author's Atlassian ID. 
    state := "state_example" // string | Only return pull requests that are in this state. This parameter can be repeated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.PullrequestsSelectedUserGet(context.Background(), selectedUser).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.PullrequestsSelectedUserGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullrequestsSelectedUserGet`: PaginatedPullrequests
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.PullrequestsSelectedUserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | This can either be the username of the pull request author, the author&#39;s UUID surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, or the author&#39;s Atlassian ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullrequestsSelectedUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | Only return pull requests that are in this state. This parameter can be repeated. | 

### Return type

[**PaginatedPullrequests**](PaginatedPullrequests.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugDefaultReviewersGet

> RepositoriesWorkspaceRepoSlugDefaultReviewersGet(ctx, repoSlug, workspace).Execute()





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
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugDefaultReviewersGet(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugDefaultReviewersGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDefaultReviewersGetRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameDelete

> ModelError RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameDelete(ctx, repoSlug, targetUsername, workspace).Execute()





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
    targetUsername := "targetUsername_example" // string | This can either be the username or the UUID of the default reviewer, surrounded by curly-braces, for example: `{account UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameDelete(context.Background(), repoSlug, targetUsername, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameDelete`: ModelError
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**targetUsername** | **string** | This can either be the username or the UUID of the default reviewer, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameDeleteRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameGet

> ModelError RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameGet(ctx, repoSlug, targetUsername, workspace).Execute()





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
    targetUsername := "targetUsername_example" // string | This can either be the username or the UUID of the default reviewer, surrounded by curly-braces, for example: `{account UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameGet(context.Background(), repoSlug, targetUsername, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameGet`: ModelError
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**targetUsername** | **string** | This can either be the username or the UUID of the default reviewer, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameGetRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernamePut

> ModelError RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernamePut(ctx, repoSlug, targetUsername, workspace).Execute()





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
    targetUsername := "targetUsername_example" // string | This can either be the username or the UUID of the default reviewer, surrounded by curly-braces, for example: `{account UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernamePut(context.Background(), repoSlug, targetUsername, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernamePut`: ModelError
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**targetUsername** | **string** | This can either be the username or the UUID of the default reviewer, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernamePutRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugPullrequestsActivityGet

> RepositoriesWorkspaceRepoSlugPullrequestsActivityGet(ctx, repoSlug, workspace).Execute()





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
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsActivityGet(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsActivityGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsActivityGetRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugPullrequestsGet

> PaginatedPullrequests RepositoriesWorkspaceRepoSlugPullrequestsGet(ctx, repoSlug, workspace).State(state).Execute()





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
    state := "state_example" // string | Only return pull requests that are in this state. This parameter can be repeated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsGet(context.Background(), repoSlug, workspace).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsGet`: PaginatedPullrequests
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **state** | **string** | Only return pull requests that are in this state. This parameter can be repeated. | 

### Return type

[**PaginatedPullrequests**](PaginatedPullrequests.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugPullrequestsPost

> Pullrequest RepositoriesWorkspaceRepoSlugPullrequestsPost(ctx, repoSlug, workspace).Body(body).Execute()





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
    body := *openapiclient.NewPullrequest("Type_example") // Pullrequest | The new pull request.  The request URL you POST to becomes the destination repository URL. For this reason, you must specify an explicit source repository in the request object if you want to pull from a different repository (fork).  Since not all elements are required or even mutable, you only need to include the elements you want to initialize, such as the source branch and the title. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPost(context.Background(), repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPost`: Pullrequest
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Pullrequest**](Pullrequest.md) | The new pull request.  The request URL you POST to becomes the destination repository URL. For this reason, you must specify an explicit source repository in the request object if you want to pull from a different repository (fork).  Since not all elements are required or even mutable, you only need to include the elements you want to initialize, such as the source branch and the title. | 

### Return type

[**Pullrequest**](Pullrequest.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdActivityGet

> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdActivityGet(ctx, pullRequestId, repoSlug, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdActivityGet(context.Background(), pullRequestId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdActivityGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdActivityGetRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApproveDelete

> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApproveDelete(ctx, pullRequestId, repoSlug, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApproveDelete(context.Background(), pullRequestId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApproveDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApproveDeleteRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApprovePost

> Participant RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApprovePost(ctx, pullRequestId, repoSlug, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApprovePost(context.Background(), pullRequestId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApprovePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApprovePost`: Participant
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApprovePost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApprovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Participant**](Participant.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdDelete

> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdDelete(ctx, commentId, pullRequestId, repoSlug, workspace).Execute()





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
    pullRequestId := int32(56) // int32 | The id of the pull request.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdDelete(context.Background(), commentId, pullRequestId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** | The id of the comment. | 
**pullRequestId** | **int32** | The id of the pull request. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdDeleteRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdGet

> PullrequestComment RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdGet(ctx, commentId, pullRequestId, repoSlug, workspace).Execute()





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
    pullRequestId := int32(56) // int32 | The id of the pull request.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdGet(context.Background(), commentId, pullRequestId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdGet`: PullrequestComment
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** | The id of the comment. | 
**pullRequestId** | **int32** | The id of the pull request. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PullrequestComment**](PullrequestComment.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdPut

> PullrequestComment RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdPut(ctx, commentId, pullRequestId, repoSlug, workspace).Body(body).Execute()





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
    pullRequestId := int32(56) // int32 | The id of the pull request.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
    body := *openapiclient.NewPullrequestComment("Type_example") // PullrequestComment | The contents of the updated comment.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdPut(context.Background(), commentId, pullRequestId, repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdPut`: PullrequestComment
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** | The id of the comment. | 
**pullRequestId** | **int32** | The id of the pull request. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**PullrequestComment**](PullrequestComment.md) | The contents of the updated comment. | 

### Return type

[**PullrequestComment**](PullrequestComment.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsGet

> PaginatedPullrequestComments RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsGet(ctx, pullRequestId, repoSlug, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsGet(context.Background(), pullRequestId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsGet`: PaginatedPullrequestComments
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PaginatedPullrequestComments**](PaginatedPullrequestComments.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsPost

> PullrequestComment RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsPost(ctx, pullRequestId, repoSlug, workspace).Body(body).Execute()





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
    body := *openapiclient.NewPullrequestComment("Type_example") // PullrequestComment | The comment object.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsPost(context.Background(), pullRequestId, repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsPost`: PullrequestComment
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsPost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**PullrequestComment**](PullrequestComment.md) | The comment object. | 

### Return type

[**PullrequestComment**](PullrequestComment.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommitsGet

> ModelError RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommitsGet(ctx, pullRequestId, repoSlug, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommitsGet(context.Background(), pullRequestId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommitsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommitsGet`: ModelError
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommitsGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommitsGetRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDeclinePost

> Pullrequest RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDeclinePost(ctx, pullRequestId, repoSlug, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDeclinePost(context.Background(), pullRequestId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDeclinePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDeclinePost`: Pullrequest
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDeclinePost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDeclinePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Pullrequest**](Pullrequest.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffGet

> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffGet(ctx, pullRequestId, repoSlug, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffGet(context.Background(), pullRequestId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffGetRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffstatGet

> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffstatGet(ctx, pullRequestId, repoSlug, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffstatGet(context.Background(), pullRequestId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffstatGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffstatGetRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet

> Pullrequest RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet(ctx, pullRequestId, repoSlug, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet(context.Background(), pullRequestId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet`: Pullrequest
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Pullrequest**](Pullrequest.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePost

> Pullrequest RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePost(ctx, pullRequestId, repoSlug, workspace).Body(body).Execute()





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
    body := *openapiclient.NewPullrequestMergeParameters("Type_example") // PullrequestMergeParameters |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePost(context.Background(), pullRequestId, repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePost`: Pullrequest
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**PullrequestMergeParameters**](PullrequestMergeParameters.md) |  | 

### Return type

[**Pullrequest**](Pullrequest.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeTaskStatusTaskIdGet

> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeTaskStatusTaskIdGet(ctx, pullRequestId, repoSlug, taskId, workspace).Execute()





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
    taskId := "taskId_example" // string | ID of the merge task
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeTaskStatusTaskIdGet(context.Background(), pullRequestId, repoSlug, taskId, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeTaskStatusTaskIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pullRequestId** | **int32** | The id of the pull request. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**taskId** | **string** | ID of the merge task | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeTaskStatusTaskIdGetRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPatchGet

> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPatchGet(ctx, pullRequestId, repoSlug, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPatchGet(context.Background(), pullRequestId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPatchGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPatchGetRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPut

> Pullrequest RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPut(ctx, pullRequestId, repoSlug, workspace).Body(body).Execute()





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
    body := *openapiclient.NewPullrequest("Type_example") // Pullrequest | The pull request that is to be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPut(context.Background(), pullRequestId, repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPut`: Pullrequest
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPut`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**Pullrequest**](Pullrequest.md) | The pull request that is to be updated. | 

### Return type

[**Pullrequest**](Pullrequest.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesDelete

> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesDelete(ctx, pullRequestId, repoSlug, workspace).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesDelete(context.Background(), pullRequestId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesDeleteRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesPost

> Participant RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesPost(ctx, pullRequestId, repoSlug, workspace).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesPost(context.Background(), pullRequestId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesPost`: Participant
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesPost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Participant**](Participant.md)

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
    resp, r, err := api_client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet(context.Background(), pullRequestId, repoSlug, workspace).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet`: PaginatedCommitstatuses
    fmt.Fprintf(os.Stdout, "Response from `PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet`: %v\n", resp)
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

