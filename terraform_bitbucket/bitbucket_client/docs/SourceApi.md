# \SourceApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet**](SourceApi.md#RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet) | **Get** /repositories/{workspace}/{repo_slug}/filehistory/{commit}/{path} | 
[**RepositoriesWorkspaceRepoSlugSrcCommitPathGet**](SourceApi.md#RepositoriesWorkspaceRepoSlugSrcCommitPathGet) | **Get** /repositories/{workspace}/{repo_slug}/src/{commit}/{path} | 
[**RepositoriesWorkspaceRepoSlugSrcGet**](SourceApi.md#RepositoriesWorkspaceRepoSlugSrcGet) | **Get** /repositories/{workspace}/{repo_slug}/src | 
[**RepositoriesWorkspaceRepoSlugSrcPost**](SourceApi.md#RepositoriesWorkspaceRepoSlugSrcPost) | **Post** /repositories/{workspace}/{repo_slug}/src | 



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
    resp, r, err := api_client.SourceApi.RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet(context.Background(), commit, path, repoSlug, workspace).Renames(renames).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet`: PaginatedFiles
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet`: %v\n", resp)
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
    resp, r, err := api_client.SourceApi.RepositoriesWorkspaceRepoSlugSrcCommitPathGet(context.Background(), commit, path, repoSlug, workspace).Format(format).Q(q).Sort(sort).MaxDepth(maxDepth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.RepositoriesWorkspaceRepoSlugSrcCommitPathGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugSrcCommitPathGet`: PaginatedTreeentries
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.RepositoriesWorkspaceRepoSlugSrcCommitPathGet`: %v\n", resp)
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
    resp, r, err := api_client.SourceApi.RepositoriesWorkspaceRepoSlugSrcGet(context.Background(), repoSlug, workspace).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.RepositoriesWorkspaceRepoSlugSrcGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugSrcGet`: PaginatedTreeentries
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.RepositoriesWorkspaceRepoSlugSrcGet`: %v\n", resp)
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
    resp, r, err := api_client.SourceApi.RepositoriesWorkspaceRepoSlugSrcPost(context.Background(), repoSlug, workspace).Message(message).Author(author).Parents(parents).Files(files).Branch(branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.RepositoriesWorkspaceRepoSlugSrcPost``: %v\n", err)
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

