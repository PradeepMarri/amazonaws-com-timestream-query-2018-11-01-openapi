package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-timestream-query/mcp-server/config"
	"github.com/amazon-timestream-query/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreatescheduledqueryHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CreateScheduledQueryRequest
		
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
		url := fmt.Sprintf("%s/#X-Amz-Target=Timestream_20181101.CreateScheduledQuery", cfg.BaseURL)
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
		var result models.CreateScheduledQueryResponse
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

func CreateCreatescheduledqueryTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=Timestream_20181101_CreateScheduledQuery",
		mcp.WithDescription(" Create a scheduled query that will be run on your behalf at the configured schedule. Timestream assumes the execution role provided as part of the <code>ScheduledQueryExecutionRoleArn</code> parameter to run the query. You can use the <code>NotificationConfiguration</code> parameter to configure notification for your scheduled query operations."),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithString("ScheduledQueryExecutionRoleArn", mcp.Required(), mcp.Description("")),
		mcp.WithString("TargetConfiguration", mcp.Description("")),
		mcp.WithString("ScheduleConfiguration", mcp.Required(), mcp.Description("")),
		mcp.WithString("ErrorReportConfiguration", mcp.Required(), mcp.Description("")),
		mcp.WithString("QueryString", mcp.Required(), mcp.Description("")),
		mcp.WithString("Tags", mcp.Description("")),
		mcp.WithString("ClientToken", mcp.Description("")),
		mcp.WithString("Name", mcp.Required(), mcp.Description("")),
		mcp.WithString("NotificationConfiguration", mcp.Required(), mcp.Description("")),
		mcp.WithString("KmsKeyId", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatescheduledqueryHandler(cfg),
	}
}
