# \CommitsApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateOrUpdateAnnotations**](CommitsApi.md#BulkCreateOrUpdateAnnotations) | **Post** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations | 
[**CreateOrUpdateAnnotation**](CommitsApi.md#CreateOrUpdateAnnotation) | **Put** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations/{annotationId} | 
[**CreateOrUpdateReport**](CommitsApi.md#CreateOrUpdateReport) | **Put** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId} | 
[**DeleteAnnotation**](CommitsApi.md#DeleteAnnotation) | **Delete** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations/{annotationId} | 
[**DeleteReport**](CommitsApi.md#DeleteReport) | **Delete** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId} | 
[**GetAnnotation**](CommitsApi.md#GetAnnotation) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations/{annotationId} | 
[**GetAnnotationsForReport**](CommitsApi.md#GetAnnotationsForReport) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations | 
[**GetReport**](CommitsApi.md#GetReport) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId} | 
[**GetReportsForCommit**](CommitsApi.md#GetReportsForCommit) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports | 
[**RepositoriesWorkspaceRepoSlugCommitCommitApproveDelete**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitCommitApproveDelete) | **Delete** /repositories/{workspace}/{repo_slug}/commit/{commit}/approve | 
[**RepositoriesWorkspaceRepoSlugCommitCommitApprovePost**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitCommitApprovePost) | **Post** /repositories/{workspace}/{repo_slug}/commit/{commit}/approve | 
[**RepositoriesWorkspaceRepoSlugCommitCommitCommentsCommentIdGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitCommitCommentsCommentIdGet) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/comments/{comment_id} | 
[**RepositoriesWorkspaceRepoSlugCommitCommitCommentsGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitCommitCommentsGet) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/comments | 
[**RepositoriesWorkspaceRepoSlugCommitCommitCommentsPost**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitCommitCommentsPost) | **Post** /repositories/{workspace}/{repo_slug}/commit/{commit}/comments | 
[**RepositoriesWorkspaceRepoSlugCommitCommitGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitCommitGet) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit} | 
[**RepositoriesWorkspaceRepoSlugCommitsGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitsGet) | **Get** /repositories/{workspace}/{repo_slug}/commits | 
[**RepositoriesWorkspaceRepoSlugCommitsPost**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitsPost) | **Post** /repositories/{workspace}/{repo_slug}/commits | 
[**RepositoriesWorkspaceRepoSlugCommitsRevisionGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitsRevisionGet) | **Get** /repositories/{workspace}/{repo_slug}/commits/{revision} | 
[**RepositoriesWorkspaceRepoSlugCommitsRevisionPost**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitsRevisionPost) | **Post** /repositories/{workspace}/{repo_slug}/commits/{revision} | 
[**RepositoriesWorkspaceRepoSlugDiffSpecGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugDiffSpecGet) | **Get** /repositories/{workspace}/{repo_slug}/diff/{spec} | 
[**RepositoriesWorkspaceRepoSlugDiffstatSpecGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugDiffstatSpecGet) | **Get** /repositories/{workspace}/{repo_slug}/diffstat/{spec} | 
[**RepositoriesWorkspaceRepoSlugMergeBaseRevspecGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugMergeBaseRevspecGet) | **Get** /repositories/{workspace}/{repo_slug}/merge-base/{revspec} | 
[**RepositoriesWorkspaceRepoSlugPatchSpecGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugPatchSpecGet) | **Get** /repositories/{workspace}/{repo_slug}/patch/{spec} | 



## BulkCreateOrUpdateAnnotations

> []ReportAnnotation BulkCreateOrUpdateAnnotations(ctx, workspace, repoSlug, commit, reportId).Body(body).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.
    commit := "commit_example" // string | The commit for which to retrieve reports.
    reportId := "reportId_example" // string | Uuid or external-if of the report for which to get annotations for.
    body := []openapiclient.ReportAnnotation{*openapiclient.NewReportAnnotation("Type_example")} // []ReportAnnotation | The annotations to create or update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.BulkCreateOrUpdateAnnotations(context.Background(), workspace, repoSlug, commit, reportId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.BulkCreateOrUpdateAnnotations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateOrUpdateAnnotations`: []ReportAnnotation
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.BulkCreateOrUpdateAnnotations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**commit** | **string** | The commit for which to retrieve reports. | 
**reportId** | **string** | Uuid or external-if of the report for which to get annotations for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateOrUpdateAnnotationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**[]ReportAnnotation**](ReportAnnotation.md) | The annotations to create or update | 

### Return type

[**[]ReportAnnotation**](ReportAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateAnnotation

> ReportAnnotation CreateOrUpdateAnnotation(ctx, workspace, repoSlug, commit, reportId, annotationId).Body(body).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.
    commit := "commit_example" // string | The commit the report belongs to.
    reportId := "reportId_example" // string | Either the uuid or external-id of the report.
    annotationId := "annotationId_example" // string | Either the uuid or external-id of the annotation.
    body := *openapiclient.NewReportAnnotation("Type_example") // ReportAnnotation | The annotation to create or update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.CreateOrUpdateAnnotation(context.Background(), workspace, repoSlug, commit, reportId, annotationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.CreateOrUpdateAnnotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateAnnotation`: ReportAnnotation
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.CreateOrUpdateAnnotation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**commit** | **string** | The commit the report belongs to. | 
**reportId** | **string** | Either the uuid or external-id of the report. | 
**annotationId** | **string** | Either the uuid or external-id of the annotation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateAnnotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **body** | [**ReportAnnotation**](ReportAnnotation.md) | The annotation to create or update | 

### Return type

[**ReportAnnotation**](ReportAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateReport

> Report CreateOrUpdateReport(ctx, workspace, repoSlug, commit, reportId).Body(body).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.
    commit := "commit_example" // string | The commit the report belongs to.
    reportId := "reportId_example" // string | Either the uuid or external-id of the report.
    body := *openapiclient.NewReport("Type_example") // Report | The report to create or update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.CreateOrUpdateReport(context.Background(), workspace, repoSlug, commit, reportId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.CreateOrUpdateReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateReport`: Report
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.CreateOrUpdateReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**commit** | **string** | The commit the report belongs to. | 
**reportId** | **string** | Either the uuid or external-id of the report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**Report**](Report.md) | The report to create or update | 

### Return type

[**Report**](Report.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnnotation

> DeleteAnnotation(ctx, workspace, repoSlug, commit, reportId, annotationId).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.
    commit := "commit_example" // string | The commit the annotation belongs to.
    reportId := "reportId_example" // string | Either the uuid or external-id of the annotation.
    annotationId := "annotationId_example" // string | Either the uuid or external-id of the annotation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.DeleteAnnotation(context.Background(), workspace, repoSlug, commit, reportId, annotationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.DeleteAnnotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**commit** | **string** | The commit the annotation belongs to. | 
**reportId** | **string** | Either the uuid or external-id of the annotation. | 
**annotationId** | **string** | Either the uuid or external-id of the annotation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnnotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReport

> DeleteReport(ctx, workspace, repoSlug, commit, reportId).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.
    commit := "commit_example" // string | The commit the report belongs to.
    reportId := "reportId_example" // string | Either the uuid or external-id of the report.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.DeleteReport(context.Background(), workspace, repoSlug, commit, reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.DeleteReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**commit** | **string** | The commit the report belongs to. | 
**reportId** | **string** | Either the uuid or external-id of the report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnnotation

> ReportAnnotation GetAnnotation(ctx, workspace, repoSlug, commit, reportId, annotationId).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.
    commit := "commit_example" // string | The commit the report belongs to.
    reportId := "reportId_example" // string | Either the uuid or external-id of the report.
    annotationId := "annotationId_example" // string | Either the uuid or external-id of the annotation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.GetAnnotation(context.Background(), workspace, repoSlug, commit, reportId, annotationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.GetAnnotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnnotation`: ReportAnnotation
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.GetAnnotation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**commit** | **string** | The commit the report belongs to. | 
**reportId** | **string** | Either the uuid or external-id of the report. | 
**annotationId** | **string** | Either the uuid or external-id of the annotation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnnotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**ReportAnnotation**](ReportAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnnotationsForReport

> PaginatedAnnotations GetAnnotationsForReport(ctx, workspace, repoSlug, commit, reportId).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.
    commit := "commit_example" // string | The commit for which to retrieve reports.
    reportId := "reportId_example" // string | Uuid or external-if of the report for which to get annotations for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.GetAnnotationsForReport(context.Background(), workspace, repoSlug, commit, reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.GetAnnotationsForReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnnotationsForReport`: PaginatedAnnotations
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.GetAnnotationsForReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**commit** | **string** | The commit for which to retrieve reports. | 
**reportId** | **string** | Uuid or external-if of the report for which to get annotations for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnnotationsForReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PaginatedAnnotations**](PaginatedAnnotations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReport

> Report GetReport(ctx, workspace, repoSlug, commit, reportId).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.
    commit := "commit_example" // string | The commit the report belongs to.
    reportId := "reportId_example" // string | Either the uuid or external-id of the report.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.GetReport(context.Background(), workspace, repoSlug, commit, reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.GetReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReport`: Report
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.GetReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**commit** | **string** | The commit the report belongs to. | 
**reportId** | **string** | Either the uuid or external-id of the report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Report**](Report.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsForCommit

> PaginatedReports GetReportsForCommit(ctx, workspace, repoSlug, commit).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.
    commit := "commit_example" // string | The commit for which to retrieve reports.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.GetReportsForCommit(context.Background(), workspace, repoSlug, commit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.GetReportsForCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportsForCommit`: PaginatedReports
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.GetReportsForCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**commit** | **string** | The commit for which to retrieve reports. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsForCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PaginatedReports**](PaginatedReports.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugCommitCommitApproveDelete

> RepositoriesWorkspaceRepoSlugCommitCommitApproveDelete(ctx, commit, repoSlug, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitApproveDelete(context.Background(), commit, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitApproveDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugCommitCommitApproveDeleteRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugCommitCommitApprovePost

> Participant RepositoriesWorkspaceRepoSlugCommitCommitApprovePost(ctx, commit, repoSlug, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitApprovePost(context.Background(), commit, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitApprovePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitCommitApprovePost`: Participant
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitApprovePost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugCommitCommitApprovePostRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugCommitCommitCommentsCommentIdGet

> CommitComment RepositoriesWorkspaceRepoSlugCommitCommitCommentsCommentIdGet(ctx, commentId, commit, repoSlug, workspace).Execute()





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
    commit := "commit_example" // string | The commit's SHA1.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitCommentsCommentIdGet(context.Background(), commentId, commit, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitCommentsCommentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitCommitCommentsCommentIdGet`: CommitComment
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitCommentsCommentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** | The id of the comment. | 
**commit** | **string** | The commit&#39;s SHA1. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugCommitCommitCommentsCommentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**CommitComment**](CommitComment.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugCommitCommitCommentsGet

> PaginatedCommitComments RepositoriesWorkspaceRepoSlugCommitCommitCommentsGet(ctx, commit, repoSlug, workspace).Q(q).Sort(sort).Execute()





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
    sort := "sort_example" // string | Field by which the results should be sorted as per [filtering and sorting](../../../../../../meta/filtering).  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitCommentsGet(context.Background(), commit, repoSlug, workspace).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitCommentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitCommitCommentsGet`: PaginatedCommitComments
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitCommentsGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugCommitCommitCommentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **q** | **string** | Query string to narrow down the response as per [filtering and sorting](../../../../../../meta/filtering).  | 
 **sort** | **string** | Field by which the results should be sorted as per [filtering and sorting](../../../../../../meta/filtering).  | 

### Return type

[**PaginatedCommitComments**](PaginatedCommitComments.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugCommitCommitCommentsPost

> RepositoriesWorkspaceRepoSlugCommitCommitCommentsPost(ctx, commit, repoSlug, workspace).Body(body).Execute()





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
    body := *openapiclient.NewCommitComment() // CommitComment | The specified comment.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitCommentsPost(context.Background(), commit, repoSlug, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitCommentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugCommitCommitCommentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**CommitComment**](CommitComment.md) | The specified comment. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugCommitCommitGet

> Commit RepositoriesWorkspaceRepoSlugCommitCommitGet(ctx, commit, repoSlug, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitGet(context.Background(), commit, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitCommitGet`: Commit
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.RepositoriesWorkspaceRepoSlugCommitCommitGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugCommitCommitGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Commit**](Commit.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugCommitsGet

> PaginatedChangeset RepositoriesWorkspaceRepoSlugCommitsGet(ctx, repoSlug, revision, workspace).Execute()





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
    revision := "revision_example" // string | (optional) The commit's SHA1.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.RepositoriesWorkspaceRepoSlugCommitsGet(context.Background(), repoSlug, revision, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.RepositoriesWorkspaceRepoSlugCommitsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitsGet`: PaginatedChangeset
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.RepositoriesWorkspaceRepoSlugCommitsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**revision** | **string** | (optional) The commit&#39;s SHA1. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugCommitsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PaginatedChangeset**](PaginatedChangeset.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugCommitsPost

> PaginatedChangeset RepositoriesWorkspaceRepoSlugCommitsPost(ctx, repoSlug, revision, workspace).Execute()





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
    revision := "revision_example" // string | (optional) The commit's SHA1.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.RepositoriesWorkspaceRepoSlugCommitsPost(context.Background(), repoSlug, revision, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.RepositoriesWorkspaceRepoSlugCommitsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitsPost`: PaginatedChangeset
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.RepositoriesWorkspaceRepoSlugCommitsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**revision** | **string** | (optional) The commit&#39;s SHA1. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugCommitsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PaginatedChangeset**](PaginatedChangeset.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugCommitsRevisionGet

> PaginatedChangeset RepositoriesWorkspaceRepoSlugCommitsRevisionGet(ctx, repoSlug, revision, workspace).Execute()





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
    revision := "revision_example" // string | (optional) The commit's SHA1.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.RepositoriesWorkspaceRepoSlugCommitsRevisionGet(context.Background(), repoSlug, revision, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.RepositoriesWorkspaceRepoSlugCommitsRevisionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitsRevisionGet`: PaginatedChangeset
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.RepositoriesWorkspaceRepoSlugCommitsRevisionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**revision** | **string** | (optional) The commit&#39;s SHA1. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugCommitsRevisionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PaginatedChangeset**](PaginatedChangeset.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugCommitsRevisionPost

> PaginatedChangeset RepositoriesWorkspaceRepoSlugCommitsRevisionPost(ctx, repoSlug, revision, workspace).Execute()





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
    revision := "revision_example" // string | (optional) The commit's SHA1.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.RepositoriesWorkspaceRepoSlugCommitsRevisionPost(context.Background(), repoSlug, revision, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.RepositoriesWorkspaceRepoSlugCommitsRevisionPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugCommitsRevisionPost`: PaginatedChangeset
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.RepositoriesWorkspaceRepoSlugCommitsRevisionPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**revision** | **string** | (optional) The commit&#39;s SHA1. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugCommitsRevisionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PaginatedChangeset**](PaginatedChangeset.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugDiffSpecGet

> RepositoriesWorkspaceRepoSlugDiffSpecGet(ctx, repoSlug, spec, workspace).Context(context).Path(path).IgnoreWhitespace(ignoreWhitespace).Binary(binary).Renames(renames).Merge(merge).Execute()





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
    spec := "spec_example" // string | A commit SHA (e.g. `3a8b42`) or a commit range using double dot notation (e.g. `3a8b42..9ff173`). 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
    context := int32(56) // int32 | Generate diffs with <n> lines of context instead of the usual three. (optional)
    path := "path_example" // string | Limit the diff to a particular file (this parameter can be repeated for multiple paths). (optional)
    ignoreWhitespace := true // bool | Generate diffs that ignore whitespace. (optional)
    binary := true // bool | Generate diffs that include binary files, true if omitted. (optional)
    renames := true // bool | Whether to perform rename detection, true if omitted. (optional)
    merge := true // bool | If true, the source commit is merged into the destination commit, and then a diff from the destination to the merge result is returned. If false, a simple 'two dot' diff between the source and destination is returned. True if omitted. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.RepositoriesWorkspaceRepoSlugDiffSpecGet(context.Background(), repoSlug, spec, workspace).Context(context).Path(path).IgnoreWhitespace(ignoreWhitespace).Binary(binary).Renames(renames).Merge(merge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.RepositoriesWorkspaceRepoSlugDiffSpecGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**spec** | **string** | A commit SHA (e.g. &#x60;3a8b42&#x60;) or a commit range using double dot notation (e.g. &#x60;3a8b42..9ff173&#x60;).  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDiffSpecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **int32** | Generate diffs with &lt;n&gt; lines of context instead of the usual three. | 
 **path** | **string** | Limit the diff to a particular file (this parameter can be repeated for multiple paths). | 
 **ignoreWhitespace** | **bool** | Generate diffs that ignore whitespace. | 
 **binary** | **bool** | Generate diffs that include binary files, true if omitted. | 
 **renames** | **bool** | Whether to perform rename detection, true if omitted. | 
 **merge** | **bool** | If true, the source commit is merged into the destination commit, and then a diff from the destination to the merge result is returned. If false, a simple &#39;two dot&#39; diff between the source and destination is returned. True if omitted. | 

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


## RepositoriesWorkspaceRepoSlugDiffstatSpecGet

> PaginatedDiffstats RepositoriesWorkspaceRepoSlugDiffstatSpecGet(ctx, repoSlug, spec, workspace).IgnoreWhitespace(ignoreWhitespace).Merge(merge).Path(path).Renames(renames).Execute()





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
    spec := "spec_example" // string | A commit SHA (e.g. `3a8b42`) or a commit range using double dot notation (e.g. `3a8b42..9ff173`). 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
    ignoreWhitespace := true // bool | Generate diffs that ignore whitespace (optional)
    merge := true // bool | If true, the source commit is merged into the destination commit, and then a diffstat from the destination to the merge result is returned. If false, a simple 'two dot' diffstat between the source and destination is returned. True if omitted. (optional)
    path := "path_example" // string | Limit the diffstat to a particular file (this parameter can be repeated for multiple paths). (optional)
    renames := true // bool | Whether to perform rename detection, true if omitted. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.RepositoriesWorkspaceRepoSlugDiffstatSpecGet(context.Background(), repoSlug, spec, workspace).IgnoreWhitespace(ignoreWhitespace).Merge(merge).Path(path).Renames(renames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.RepositoriesWorkspaceRepoSlugDiffstatSpecGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugDiffstatSpecGet`: PaginatedDiffstats
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.RepositoriesWorkspaceRepoSlugDiffstatSpecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**spec** | **string** | A commit SHA (e.g. &#x60;3a8b42&#x60;) or a commit range using double dot notation (e.g. &#x60;3a8b42..9ff173&#x60;).  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDiffstatSpecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ignoreWhitespace** | **bool** | Generate diffs that ignore whitespace | 
 **merge** | **bool** | If true, the source commit is merged into the destination commit, and then a diffstat from the destination to the merge result is returned. If false, a simple &#39;two dot&#39; diffstat between the source and destination is returned. True if omitted. | 
 **path** | **string** | Limit the diffstat to a particular file (this parameter can be repeated for multiple paths). | 
 **renames** | **bool** | Whether to perform rename detection, true if omitted. | 

### Return type

[**PaginatedDiffstats**](PaginatedDiffstats.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugMergeBaseRevspecGet

> Commit RepositoriesWorkspaceRepoSlugMergeBaseRevspecGet(ctx, repoSlug, revspec, workspace).Execute()





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
    revspec := "revspec_example" // string | A commit range using double dot notation (e.g. `3a8b42..9ff173`). 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.RepositoriesWorkspaceRepoSlugMergeBaseRevspecGet(context.Background(), repoSlug, revspec, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.RepositoriesWorkspaceRepoSlugMergeBaseRevspecGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugMergeBaseRevspecGet`: Commit
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.RepositoriesWorkspaceRepoSlugMergeBaseRevspecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**revspec** | **string** | A commit range using double dot notation (e.g. &#x60;3a8b42..9ff173&#x60;).  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugMergeBaseRevspecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Commit**](Commit.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugPatchSpecGet

> RepositoriesWorkspaceRepoSlugPatchSpecGet(ctx, repoSlug, spec, workspace).Execute()





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
    spec := "spec_example" // string | A commit SHA (e.g. `3a8b42`) or a commit range using double dot notation (e.g. `3a8b42..9ff173`). 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.RepositoriesWorkspaceRepoSlugPatchSpecGet(context.Background(), repoSlug, spec, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.RepositoriesWorkspaceRepoSlugPatchSpecGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**spec** | **string** | A commit SHA (e.g. &#x60;3a8b42&#x60;) or a commit range using double dot notation (e.g. &#x60;3a8b42..9ff173&#x60;).  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPatchSpecGetRequest struct via the builder pattern


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

