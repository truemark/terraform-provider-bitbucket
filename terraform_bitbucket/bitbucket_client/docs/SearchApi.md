# \SearchApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchAccount**](SearchApi.md#SearchAccount) | **Get** /teams/{username}/search/code | Search for code in the repositories of the specified team.  Searching across all repositories:  &#x60;&#x60;&#x60; curl &#39;https://api.bitbucket.org/2.0/teams/team_name/search/code?search_query&#x3D;foo&#39; {   \&quot;size\&quot;: 1,   \&quot;page\&quot;: 1,   \&quot;pagelen\&quot;: 10,   \&quot;query_substituted\&quot;: false,   \&quot;values\&quot;: [     {       \&quot;type\&quot;: \&quot;code_search_result\&quot;,       \&quot;content_match_count\&quot;: 2,       \&quot;content_matches\&quot;: [         {           \&quot;lines\&quot;: [             {               \&quot;line\&quot;: 2,               \&quot;segments\&quot;: []             },             {               \&quot;line\&quot;: 3,               \&quot;segments\&quot;: [                 {                   \&quot;text\&quot;: \&quot;def \&quot;                 },                 {                   \&quot;text\&quot;: \&quot;foo\&quot;,                   \&quot;match\&quot;: true                 },                 {                   \&quot;text\&quot;: \&quot;():\&quot;                 }               ]             },             {               \&quot;line\&quot;: 4,               \&quot;segments\&quot;: [                 {                   \&quot;text\&quot;: \&quot;    print(\\\&quot;snek\\\&quot;)\&quot;                 }               ]             },             {               \&quot;line\&quot;: 5,               \&quot;segments\&quot;: []             }           ]         }       ],       \&quot;path_matches\&quot;: [         {           \&quot;text\&quot;: \&quot;src/\&quot;         },         {           \&quot;text\&quot;: \&quot;foo\&quot;,           \&quot;match\&quot;: true         },         {           \&quot;text\&quot;: \&quot;.py\&quot;         }       ],       \&quot;file\&quot;: {         \&quot;path\&quot;: \&quot;src/foo.py\&quot;,         \&quot;type\&quot;: \&quot;commit_file\&quot;,         \&quot;links\&quot;: {           \&quot;self\&quot;: {             \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\&quot;           }         }       }     }   ] } &#x60;&#x60;&#x60;  Note that searches can match in the file&#39;s text (&#x60;content_matches&#x60;), the path (&#x60;path_matches&#x60;), or both as in the example above.  You can use the same syntax for the search query as in the UI, e.g. to only search within a specific repository:  &#x60;&#x60;&#x60; curl &#39;https://api.bitbucket.org/2.0/teams/team_name/search/code?search_query&#x3D;foo+repo:demo&#39; # results from the \&quot;demo\&quot; repository &#x60;&#x60;&#x60;  Similar to other APIs, you can request more fields using a &#x60;fields&#x60; query parameter. E.g. to get some more information about the repository of matched files (the &#x60;%2B&#x60; is a URL-encoded &#x60;+&#x60;):  &#x60;&#x60;&#x60; curl &#39;https://api.bitbucket.org/2.0/teams/team_name/search/code&#39;\\      &#39;?search_query&#x3D;foo&amp;fields&#x3D;%2Bvalues.file.commit.repository&#39; {   \&quot;size\&quot;: 1,   \&quot;page\&quot;: 1,   \&quot;pagelen\&quot;: 10,   \&quot;query_substituted\&quot;: false,   \&quot;values\&quot;: [     {       \&quot;type\&quot;: \&quot;code_search_result\&quot;,       \&quot;content_match_count\&quot;: 1,       \&quot;content_matches\&quot;: [...],       \&quot;path_matches\&quot;: [...],       \&quot;file\&quot;: {         \&quot;commit\&quot;: {           \&quot;type\&quot;: \&quot;commit\&quot;,           \&quot;hash\&quot;: \&quot;ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\&quot;,           \&quot;links\&quot;: {             \&quot;self\&quot;: {               \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/repositories/my-workspace/demo/commit/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\&quot;             },             \&quot;html\&quot;: {               \&quot;href\&quot;: \&quot;https://bitbucket.org/my-workspace/demo/commits/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\&quot;             }           },           \&quot;repository\&quot;: {             \&quot;name\&quot;: \&quot;demo\&quot;,             \&quot;type\&quot;: \&quot;repository\&quot;,             \&quot;full_name\&quot;: \&quot;my-workspace/demo\&quot;,             \&quot;links\&quot;: {               \&quot;self\&quot;: {                 \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/repositories/my-workspace/demo\&quot;               },               \&quot;html\&quot;: {                 \&quot;href\&quot;: \&quot;https://bitbucket.org/my-workspace/demo\&quot;               },               \&quot;avatar\&quot;: {                 \&quot;href\&quot;: \&quot;https://bytebucket.org/ravatar/%7B850e1749-781a-4115-9316-df39d0600e7a%7D?ts&#x3D;default\&quot;               }             },             \&quot;uuid\&quot;: \&quot;{850e1749-781a-4115-9316-df39d0600e7a}\&quot;           }         },         \&quot;type\&quot;: \&quot;commit_file\&quot;,         \&quot;links\&quot;: {           \&quot;self\&quot;: {             \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\&quot;           }         },         \&quot;path\&quot;: \&quot;src/foo.py\&quot;       }     }   ] } &#x60;&#x60;&#x60;  Try &#x60;fields&#x3D;%2Bvalues.*.*.*.*&#x60; to get an idea what&#39;s possible. 
[**SearchAccount_0**](SearchApi.md#SearchAccount_0) | **Get** /users/{selected_user}/search/code | Search for code in the repositories of the specified user.  Searching across all repositories:  &#x60;&#x60;&#x60; curl &#39;https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/search/code?search_query&#x3D;foo&#39; {   \&quot;size\&quot;: 1,   \&quot;page\&quot;: 1,   \&quot;pagelen\&quot;: 10,   \&quot;query_substituted\&quot;: false,   \&quot;values\&quot;: [     {       \&quot;type\&quot;: \&quot;code_search_result\&quot;,       \&quot;content_match_count\&quot;: 2,       \&quot;content_matches\&quot;: [         {           \&quot;lines\&quot;: [             {               \&quot;line\&quot;: 2,               \&quot;segments\&quot;: []             },             {               \&quot;line\&quot;: 3,               \&quot;segments\&quot;: [                 {                   \&quot;text\&quot;: \&quot;def \&quot;                 },                 {                   \&quot;text\&quot;: \&quot;foo\&quot;,                   \&quot;match\&quot;: true                 },                 {                   \&quot;text\&quot;: \&quot;():\&quot;                 }               ]             },             {               \&quot;line\&quot;: 4,               \&quot;segments\&quot;: [                 {                   \&quot;text\&quot;: \&quot;    print(\\\&quot;snek\\\&quot;)\&quot;                 }               ]             },             {               \&quot;line\&quot;: 5,               \&quot;segments\&quot;: []             }           ]         }       ],       \&quot;path_matches\&quot;: [         {           \&quot;text\&quot;: \&quot;src/\&quot;         },         {           \&quot;text\&quot;: \&quot;foo\&quot;,           \&quot;match\&quot;: true         },         {           \&quot;text\&quot;: \&quot;.py\&quot;         }       ],       \&quot;file\&quot;: {         \&quot;path\&quot;: \&quot;src/foo.py\&quot;,         \&quot;type\&quot;: \&quot;commit_file\&quot;,         \&quot;links\&quot;: {           \&quot;self\&quot;: {             \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\&quot;           }         }       }     }   ] } &#x60;&#x60;&#x60;  Note that searches can match in the file&#39;s text (&#x60;content_matches&#x60;), the path (&#x60;path_matches&#x60;), or both as in the example above.  You can use the same syntax for the search query as in the UI, e.g. to only search within a specific repository:  &#x60;&#x60;&#x60; curl &#39;https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/search/code?search_query&#x3D;foo+repo:demo&#39; # results from the \&quot;demo\&quot; repository &#x60;&#x60;&#x60;  Similar to other APIs, you can request more fields using a &#x60;fields&#x60; query parameter. E.g. to get some more information about the repository of matched files (the &#x60;%2B&#x60; is a URL-encoded &#x60;+&#x60;):  &#x60;&#x60;&#x60; curl &#39;https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/search/code&#39;\\      &#39;?search_query&#x3D;foo&amp;fields&#x3D;%2Bvalues.file.commit.repository&#39; {   \&quot;size\&quot;: 1,   \&quot;page\&quot;: 1,   \&quot;pagelen\&quot;: 10,   \&quot;query_substituted\&quot;: false,   \&quot;values\&quot;: [     {       \&quot;type\&quot;: \&quot;code_search_result\&quot;,       \&quot;content_match_count\&quot;: 1,       \&quot;content_matches\&quot;: [...],       \&quot;path_matches\&quot;: [...],       \&quot;file\&quot;: {         \&quot;commit\&quot;: {           \&quot;type\&quot;: \&quot;commit\&quot;,           \&quot;hash\&quot;: \&quot;ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\&quot;,           \&quot;links\&quot;: {             \&quot;self\&quot;: {               \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/repositories/my-workspace/demo/commit/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\&quot;             },             \&quot;html\&quot;: {               \&quot;href\&quot;: \&quot;https://bitbucket.org/my-workspace/demo/commits/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\&quot;             }           },           \&quot;repository\&quot;: {             \&quot;name\&quot;: \&quot;demo\&quot;,             \&quot;type\&quot;: \&quot;repository\&quot;,             \&quot;full_name\&quot;: \&quot;my-workspace/demo\&quot;,             \&quot;links\&quot;: {               \&quot;self\&quot;: {                 \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/repositories/my-workspace/demo\&quot;               },               \&quot;html\&quot;: {                 \&quot;href\&quot;: \&quot;https://bitbucket.org/my-workspace/demo\&quot;               },               \&quot;avatar\&quot;: {                 \&quot;href\&quot;: \&quot;https://bytebucket.org/ravatar/%7B850e1749-781a-4115-9316-df39d0600e7a%7D?ts&#x3D;default\&quot;               }             },             \&quot;uuid\&quot;: \&quot;{850e1749-781a-4115-9316-df39d0600e7a}\&quot;           }         },         \&quot;type\&quot;: \&quot;commit_file\&quot;,         \&quot;links\&quot;: {           \&quot;self\&quot;: {             \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\&quot;           }         },         \&quot;path\&quot;: \&quot;src/foo.py\&quot;       }     }   ] } &#x60;&#x60;&#x60;  Try &#x60;fields&#x3D;%2Bvalues.*.*.*.*&#x60; to get an idea what&#39;s possible. 
[**SearchAccount_1**](SearchApi.md#SearchAccount_1) | **Get** /workspaces/{workspace}/search/code | Search for code in the repositories of the specified workspace.  Searching across all repositories:  &#x60;&#x60;&#x60; curl &#39;https://api.bitbucket.org/2.0/workspaces/workspace_slug_or_uuid/search/code?search_query&#x3D;foo&#39; {   \&quot;size\&quot;: 1,   \&quot;page\&quot;: 1,   \&quot;pagelen\&quot;: 10,   \&quot;query_substituted\&quot;: false,   \&quot;values\&quot;: [     {       \&quot;type\&quot;: \&quot;code_search_result\&quot;,       \&quot;content_match_count\&quot;: 2,       \&quot;content_matches\&quot;: [         {           \&quot;lines\&quot;: [             {               \&quot;line\&quot;: 2,               \&quot;segments\&quot;: []             },             {               \&quot;line\&quot;: 3,               \&quot;segments\&quot;: [                 {                   \&quot;text\&quot;: \&quot;def \&quot;                 },                 {                   \&quot;text\&quot;: \&quot;foo\&quot;,                   \&quot;match\&quot;: true                 },                 {                   \&quot;text\&quot;: \&quot;():\&quot;                 }               ]             },             {               \&quot;line\&quot;: 4,               \&quot;segments\&quot;: [                 {                   \&quot;text\&quot;: \&quot;    print(\\\&quot;snek\\\&quot;)\&quot;                 }               ]             },             {               \&quot;line\&quot;: 5,               \&quot;segments\&quot;: []             }           ]         }       ],       \&quot;path_matches\&quot;: [         {           \&quot;text\&quot;: \&quot;src/\&quot;         },         {           \&quot;text\&quot;: \&quot;foo\&quot;,           \&quot;match\&quot;: true         },         {           \&quot;text\&quot;: \&quot;.py\&quot;         }       ],       \&quot;file\&quot;: {         \&quot;path\&quot;: \&quot;src/foo.py\&quot;,         \&quot;type\&quot;: \&quot;commit_file\&quot;,         \&quot;links\&quot;: {           \&quot;self\&quot;: {             \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\&quot;           }         }       }     }   ] } &#x60;&#x60;&#x60;  Note that searches can match in the file&#39;s text (&#x60;content_matches&#x60;), the path (&#x60;path_matches&#x60;), or both as in the example above.  You can use the same syntax for the search query as in the UI, e.g. to only search within a specific repository:  &#x60;&#x60;&#x60; curl &#39;https://api.bitbucket.org/2.0/workspaces/my-workspace/search/code?search_query&#x3D;foo+repo:demo&#39; # results from the \&quot;demo\&quot; repository &#x60;&#x60;&#x60;  Similar to other APIs, you can request more fields using a &#x60;fields&#x60; query parameter. E.g. to get some more information about the repository of matched files (the &#x60;%2B&#x60; is a URL-encoded &#x60;+&#x60;):  &#x60;&#x60;&#x60; curl &#39;https://api.bitbucket.org/2.0/workspaces/my-workspace/search/code&#39;\\      &#39;?search_query&#x3D;foo&amp;fields&#x3D;%2Bvalues.file.commit.repository&#39; {   \&quot;size\&quot;: 1,   \&quot;page\&quot;: 1,   \&quot;pagelen\&quot;: 10,   \&quot;query_substituted\&quot;: false,   \&quot;values\&quot;: [     {       \&quot;type\&quot;: \&quot;code_search_result\&quot;,       \&quot;content_match_count\&quot;: 1,       \&quot;content_matches\&quot;: [...],       \&quot;path_matches\&quot;: [...],       \&quot;file\&quot;: {         \&quot;commit\&quot;: {           \&quot;type\&quot;: \&quot;commit\&quot;,           \&quot;hash\&quot;: \&quot;ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\&quot;,           \&quot;links\&quot;: {             \&quot;self\&quot;: {               \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/repositories/my-workspace/demo/commit/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\&quot;             },             \&quot;html\&quot;: {               \&quot;href\&quot;: \&quot;https://bitbucket.org/my-workspace/demo/commits/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\&quot;             }           },           \&quot;repository\&quot;: {             \&quot;name\&quot;: \&quot;demo\&quot;,             \&quot;type\&quot;: \&quot;repository\&quot;,             \&quot;full_name\&quot;: \&quot;my-workspace/demo\&quot;,             \&quot;links\&quot;: {               \&quot;self\&quot;: {                 \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/repositories/my-workspace/demo\&quot;               },               \&quot;html\&quot;: {                 \&quot;href\&quot;: \&quot;https://bitbucket.org/my-workspace/demo\&quot;               },               \&quot;avatar\&quot;: {                 \&quot;href\&quot;: \&quot;https://bytebucket.org/ravatar/%7B850e1749-781a-4115-9316-df39d0600e7a%7D?ts&#x3D;default\&quot;               }             },             \&quot;uuid\&quot;: \&quot;{850e1749-781a-4115-9316-df39d0600e7a}\&quot;           }         },         \&quot;type\&quot;: \&quot;commit_file\&quot;,         \&quot;links\&quot;: {           \&quot;self\&quot;: {             \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\&quot;           }         },         \&quot;path\&quot;: \&quot;src/foo.py\&quot;       }     }   ] } &#x60;&#x60;&#x60;  Try &#x60;fields&#x3D;%2Bvalues.*.*.*.*&#x60; to get an idea what&#39;s possible. 



## SearchAccount

> SearchResultPage SearchAccount(ctx, username).SearchQuery(searchQuery).Page(page).Pagelen(pagelen).Execute()

Search for code in the repositories of the specified team.  Searching across all repositories:  ``` curl 'https://api.bitbucket.org/2.0/teams/team_name/search/code?search_query=foo' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 2,       \"content_matches\": [         {           \"lines\": [             {               \"line\": 2,               \"segments\": []             },             {               \"line\": 3,               \"segments\": [                 {                   \"text\": \"def \"                 },                 {                   \"text\": \"foo\",                   \"match\": true                 },                 {                   \"text\": \"():\"                 }               ]             },             {               \"line\": 4,               \"segments\": [                 {                   \"text\": \"    print(\\\"snek\\\")\"                 }               ]             },             {               \"line\": 5,               \"segments\": []             }           ]         }       ],       \"path_matches\": [         {           \"text\": \"src/\"         },         {           \"text\": \"foo\",           \"match\": true         },         {           \"text\": \".py\"         }       ],       \"file\": {         \"path\": \"src/foo.py\",         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         }       }     }   ] } ```  Note that searches can match in the file's text (`content_matches`), the path (`path_matches`), or both as in the example above.  You can use the same syntax for the search query as in the UI, e.g. to only search within a specific repository:  ``` curl 'https://api.bitbucket.org/2.0/teams/team_name/search/code?search_query=foo+repo:demo' # results from the \"demo\" repository ```  Similar to other APIs, you can request more fields using a `fields` query parameter. E.g. to get some more information about the repository of matched files (the `%2B` is a URL-encoded `+`):  ``` curl 'https://api.bitbucket.org/2.0/teams/team_name/search/code'\\      '?search_query=foo&fields=%2Bvalues.file.commit.repository' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 1,       \"content_matches\": [...],       \"path_matches\": [...],       \"file\": {         \"commit\": {           \"type\": \"commit\",           \"hash\": \"ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\",           \"links\": {             \"self\": {               \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/commit/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             },             \"html\": {               \"href\": \"https://bitbucket.org/my-workspace/demo/commits/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             }           },           \"repository\": {             \"name\": \"demo\",             \"type\": \"repository\",             \"full_name\": \"my-workspace/demo\",             \"links\": {               \"self\": {                 \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo\"               },               \"html\": {                 \"href\": \"https://bitbucket.org/my-workspace/demo\"               },               \"avatar\": {                 \"href\": \"https://bytebucket.org/ravatar/%7B850e1749-781a-4115-9316-df39d0600e7a%7D?ts=default\"               }             },             \"uuid\": \"{850e1749-781a-4115-9316-df39d0600e7a}\"           }         },         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         },         \"path\": \"src/foo.py\"       }     }   ] } ```  Try `fields=%2Bvalues.*.*.*.*` to get an idea what's possible. 

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
    username := "username_example" // string | The account to search in; either the username or the UUID in curly braces
    searchQuery := "searchQuery_example" // string | The search query
    page := int32(56) // int32 | Which page of the search results to retrieve (optional) (default to 1)
    pagelen := int32(56) // int32 | How many search results to retrieve per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.SearchAccount(context.Background(), username).SearchQuery(searchQuery).Page(page).Pagelen(pagelen).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAccount`: SearchResultPage
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The account to search in; either the username or the UUID in curly braces | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchQuery** | **string** | The search query | 
 **page** | **int32** | Which page of the search results to retrieve | [default to 1]
 **pagelen** | **int32** | How many search results to retrieve per page | [default to 10]

### Return type

[**SearchResultPage**](SearchResultPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAccount_0

> SearchResultPage SearchAccount_0(ctx, selectedUser).SearchQuery(searchQuery).Page(page).Pagelen(pagelen).Execute()

Search for code in the repositories of the specified user.  Searching across all repositories:  ``` curl 'https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/search/code?search_query=foo' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 2,       \"content_matches\": [         {           \"lines\": [             {               \"line\": 2,               \"segments\": []             },             {               \"line\": 3,               \"segments\": [                 {                   \"text\": \"def \"                 },                 {                   \"text\": \"foo\",                   \"match\": true                 },                 {                   \"text\": \"():\"                 }               ]             },             {               \"line\": 4,               \"segments\": [                 {                   \"text\": \"    print(\\\"snek\\\")\"                 }               ]             },             {               \"line\": 5,               \"segments\": []             }           ]         }       ],       \"path_matches\": [         {           \"text\": \"src/\"         },         {           \"text\": \"foo\",           \"match\": true         },         {           \"text\": \".py\"         }       ],       \"file\": {         \"path\": \"src/foo.py\",         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         }       }     }   ] } ```  Note that searches can match in the file's text (`content_matches`), the path (`path_matches`), or both as in the example above.  You can use the same syntax for the search query as in the UI, e.g. to only search within a specific repository:  ``` curl 'https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/search/code?search_query=foo+repo:demo' # results from the \"demo\" repository ```  Similar to other APIs, you can request more fields using a `fields` query parameter. E.g. to get some more information about the repository of matched files (the `%2B` is a URL-encoded `+`):  ``` curl 'https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/search/code'\\      '?search_query=foo&fields=%2Bvalues.file.commit.repository' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 1,       \"content_matches\": [...],       \"path_matches\": [...],       \"file\": {         \"commit\": {           \"type\": \"commit\",           \"hash\": \"ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\",           \"links\": {             \"self\": {               \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/commit/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             },             \"html\": {               \"href\": \"https://bitbucket.org/my-workspace/demo/commits/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             }           },           \"repository\": {             \"name\": \"demo\",             \"type\": \"repository\",             \"full_name\": \"my-workspace/demo\",             \"links\": {               \"self\": {                 \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo\"               },               \"html\": {                 \"href\": \"https://bitbucket.org/my-workspace/demo\"               },               \"avatar\": {                 \"href\": \"https://bytebucket.org/ravatar/%7B850e1749-781a-4115-9316-df39d0600e7a%7D?ts=default\"               }             },             \"uuid\": \"{850e1749-781a-4115-9316-df39d0600e7a}\"           }         },         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         },         \"path\": \"src/foo.py\"       }     }   ] } ```  Try `fields=%2Bvalues.*.*.*.*` to get an idea what's possible. 

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
    selectedUser := "selectedUser_example" // string | Either the UUID of the account surrounded by curly-braces, for example `{account UUID}`, OR an Atlassian Account ID.
    searchQuery := "searchQuery_example" // string | The search query
    page := int32(56) // int32 | Which page of the search results to retrieve (optional) (default to 1)
    pagelen := int32(56) // int32 | How many search results to retrieve per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.SearchAccount_0(context.Background(), selectedUser).SearchQuery(searchQuery).Page(page).Pagelen(pagelen).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchAccount_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAccount_0`: SearchResultPage
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchAccount_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAccount_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchQuery** | **string** | The search query | 
 **page** | **int32** | Which page of the search results to retrieve | [default to 1]
 **pagelen** | **int32** | How many search results to retrieve per page | [default to 10]

### Return type

[**SearchResultPage**](SearchResultPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAccount_1

> SearchResultPage SearchAccount_1(ctx, workspace).SearchQuery(searchQuery).Page(page).Pagelen(pagelen).Execute()

Search for code in the repositories of the specified workspace.  Searching across all repositories:  ``` curl 'https://api.bitbucket.org/2.0/workspaces/workspace_slug_or_uuid/search/code?search_query=foo' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 2,       \"content_matches\": [         {           \"lines\": [             {               \"line\": 2,               \"segments\": []             },             {               \"line\": 3,               \"segments\": [                 {                   \"text\": \"def \"                 },                 {                   \"text\": \"foo\",                   \"match\": true                 },                 {                   \"text\": \"():\"                 }               ]             },             {               \"line\": 4,               \"segments\": [                 {                   \"text\": \"    print(\\\"snek\\\")\"                 }               ]             },             {               \"line\": 5,               \"segments\": []             }           ]         }       ],       \"path_matches\": [         {           \"text\": \"src/\"         },         {           \"text\": \"foo\",           \"match\": true         },         {           \"text\": \".py\"         }       ],       \"file\": {         \"path\": \"src/foo.py\",         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         }       }     }   ] } ```  Note that searches can match in the file's text (`content_matches`), the path (`path_matches`), or both as in the example above.  You can use the same syntax for the search query as in the UI, e.g. to only search within a specific repository:  ``` curl 'https://api.bitbucket.org/2.0/workspaces/my-workspace/search/code?search_query=foo+repo:demo' # results from the \"demo\" repository ```  Similar to other APIs, you can request more fields using a `fields` query parameter. E.g. to get some more information about the repository of matched files (the `%2B` is a URL-encoded `+`):  ``` curl 'https://api.bitbucket.org/2.0/workspaces/my-workspace/search/code'\\      '?search_query=foo&fields=%2Bvalues.file.commit.repository' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 1,       \"content_matches\": [...],       \"path_matches\": [...],       \"file\": {         \"commit\": {           \"type\": \"commit\",           \"hash\": \"ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\",           \"links\": {             \"self\": {               \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/commit/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             },             \"html\": {               \"href\": \"https://bitbucket.org/my-workspace/demo/commits/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             }           },           \"repository\": {             \"name\": \"demo\",             \"type\": \"repository\",             \"full_name\": \"my-workspace/demo\",             \"links\": {               \"self\": {                 \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo\"               },               \"html\": {                 \"href\": \"https://bitbucket.org/my-workspace/demo\"               },               \"avatar\": {                 \"href\": \"https://bytebucket.org/ravatar/%7B850e1749-781a-4115-9316-df39d0600e7a%7D?ts=default\"               }             },             \"uuid\": \"{850e1749-781a-4115-9316-df39d0600e7a}\"           }         },         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         },         \"path\": \"src/foo.py\"       }     }   ] } ```  Try `fields=%2Bvalues.*.*.*.*` to get an idea what's possible. 

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
    workspace := "workspace_example" // string | The workspace to search in; either the slug or the UUID in curly braces
    searchQuery := "searchQuery_example" // string | The search query
    page := int32(56) // int32 | Which page of the search results to retrieve (optional) (default to 1)
    pagelen := int32(56) // int32 | How many search results to retrieve per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.SearchAccount_1(context.Background(), workspace).SearchQuery(searchQuery).Page(page).Pagelen(pagelen).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchAccount_1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAccount_1`: SearchResultPage
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchAccount_1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The workspace to search in; either the slug or the UUID in curly braces | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAccount_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchQuery** | **string** | The search query | 
 **page** | **int32** | Which page of the search results to retrieve | [default to 1]
 **pagelen** | **int32** | How many search results to retrieve per page | [default to 10]

### Return type

[**SearchResultPage**](SearchResultPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

