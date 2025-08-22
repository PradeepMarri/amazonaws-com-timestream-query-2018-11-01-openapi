package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/amazon-timestream-query/mcp-server/config"
	"github.com/amazon-timestream-query/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func QueryHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["MaxRows"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("MaxRows=%v", val))
		}
		if val, ok := args["NextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("NextToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.QueryRequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/#X-Amz-Target=Timestream_20181101.Query%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Amz-Target"]; ok {
			req.Header.Set("X-Amz-Target", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.QueryResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateQueryTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=Timestream_20181101_Query",
		mcp.WithDescription("<p> <code>Query</code> is a synchronous operation that enables you to run a query against your Amazon Timestream data. <code>Query</code> will time out after 60 seconds. You must update the default timeout in the SDK to support a timeout of 60 seconds. See the <a href="https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.run-query.html">code sample</a> for details. </p> <p>Your query request will fail in the following cases:</p> <ul> <li> <p> If you submit a <code>Query</code> request with the same client token outside of the 5-minute idempotency window. </p> </li> <li> <p> If you submit a <code>Query</code> request with the same client token, but change other parameters, within the 5-minute idempotency window. </p> </li> <li> <p> If the size of the row (including the query metadata) exceeds 1 MB, then the query will fail with the following error message: </p> <p> <code>Query aborted as max page response size has been exceeded by the output result row</code> </p> </li> <li> <p> If the IAM principal of the query initiator and the result reader are not the same and/or the query initiator and the result reader do not have the same query string in the query requests, the query will fail with an <code>Invalid pagination token</code> error. </p> </li> </ul>"),
		mcp.WithString("MaxRows", mcp.Description("Pagination limit")),
		mcp.WithString("NextToken", mcp.Description("Pagination token")),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithString("NextToken", mcp.Description("")),
		mcp.WithString("QueryString", mcp.Required(), mcp.Description("")),
		mcp.WithString("ClientToken", mcp.Description("")),
		mcp.WithString("MaxRows", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    QueryHandler(cfg),
	}
}
