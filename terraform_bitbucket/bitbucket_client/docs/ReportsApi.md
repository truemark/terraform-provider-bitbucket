# \ReportsApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateOrUpdateAnnotations**](ReportsApi.md#BulkCreateOrUpdateAnnotations) | **Post** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations | 
[**CreateOrUpdateAnnotation**](ReportsApi.md#CreateOrUpdateAnnotation) | **Put** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations/{annotationId} | 
[**CreateOrUpdateReport**](ReportsApi.md#CreateOrUpdateReport) | **Put** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId} | 
[**DeleteAnnotation**](ReportsApi.md#DeleteAnnotation) | **Delete** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations/{annotationId} | 
[**DeleteReport**](ReportsApi.md#DeleteReport) | **Delete** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId} | 
[**GetAnnotation**](ReportsApi.md#GetAnnotation) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations/{annotationId} | 
[**GetAnnotationsForReport**](ReportsApi.md#GetAnnotationsForReport) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations | 
[**GetReport**](ReportsApi.md#GetReport) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId} | 
[**GetReportsForCommit**](ReportsApi.md#GetReportsForCommit) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports | 



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
    resp, r, err := api_client.ReportsApi.BulkCreateOrUpdateAnnotations(context.Background(), workspace, repoSlug, commit, reportId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.BulkCreateOrUpdateAnnotations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateOrUpdateAnnotations`: []ReportAnnotation
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.BulkCreateOrUpdateAnnotations`: %v\n", resp)
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
    resp, r, err := api_client.ReportsApi.CreateOrUpdateAnnotation(context.Background(), workspace, repoSlug, commit, reportId, annotationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.CreateOrUpdateAnnotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateAnnotation`: ReportAnnotation
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.CreateOrUpdateAnnotation`: %v\n", resp)
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
    resp, r, err := api_client.ReportsApi.CreateOrUpdateReport(context.Background(), workspace, repoSlug, commit, reportId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.CreateOrUpdateReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateReport`: Report
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.CreateOrUpdateReport`: %v\n", resp)
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
    resp, r, err := api_client.ReportsApi.DeleteAnnotation(context.Background(), workspace, repoSlug, commit, reportId, annotationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.DeleteAnnotation``: %v\n", err)
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
    resp, r, err := api_client.ReportsApi.DeleteReport(context.Background(), workspace, repoSlug, commit, reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.DeleteReport``: %v\n", err)
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
    resp, r, err := api_client.ReportsApi.GetAnnotation(context.Background(), workspace, repoSlug, commit, reportId, annotationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetAnnotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnnotation`: ReportAnnotation
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetAnnotation`: %v\n", resp)
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
    resp, r, err := api_client.ReportsApi.GetAnnotationsForReport(context.Background(), workspace, repoSlug, commit, reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetAnnotationsForReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnnotationsForReport`: PaginatedAnnotations
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetAnnotationsForReport`: %v\n", resp)
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
    resp, r, err := api_client.ReportsApi.GetReport(context.Background(), workspace, repoSlug, commit, reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReport`: Report
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetReport`: %v\n", resp)
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
    resp, r, err := api_client.ReportsApi.GetReportsForCommit(context.Background(), workspace, repoSlug, commit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetReportsForCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportsForCommit`: PaginatedReports
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetReportsForCommit`: %v\n", resp)
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

