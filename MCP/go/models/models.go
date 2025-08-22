package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// QueryResponse represents the QueryResponse schema from the OpenAPI specification
type QueryResponse struct {
	Queryid interface{} `json:"QueryId"`
	Querystatus interface{} `json:"QueryStatus,omitempty"`
	Rows interface{} `json:"Rows"`
	Columninfo interface{} `json:"ColumnInfo"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// DescribeEndpointsResponse represents the DescribeEndpointsResponse schema from the OpenAPI specification
type DescribeEndpointsResponse struct {
	Endpoints interface{} `json:"Endpoints"`
}

// ScheduledQuery represents the ScheduledQuery schema from the OpenAPI specification
type ScheduledQuery struct {
	Errorreportconfiguration interface{} `json:"ErrorReportConfiguration,omitempty"`
	Nextinvocationtime interface{} `json:"NextInvocationTime,omitempty"`
	Previousinvocationtime interface{} `json:"PreviousInvocationTime,omitempty"`
	State interface{} `json:"State"`
	Lastrunstatus interface{} `json:"LastRunStatus,omitempty"`
	Name interface{} `json:"Name"`
	Targetdestination interface{} `json:"TargetDestination,omitempty"`
	Arn interface{} `json:"Arn"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
}

// ScheduleConfiguration represents the ScheduleConfiguration schema from the OpenAPI specification
type ScheduleConfiguration struct {
	Scheduleexpression interface{} `json:"ScheduleExpression"`
}

// TimeSeriesDataPoint represents the TimeSeriesDataPoint schema from the OpenAPI specification
type TimeSeriesDataPoint struct {
	Time interface{} `json:"Time"`
	Value interface{} `json:"Value"`
}

// ErrorReportConfiguration represents the ErrorReportConfiguration schema from the OpenAPI specification
type ErrorReportConfiguration struct {
	S3configuration interface{} `json:"S3Configuration"`
}

// PrepareQueryRequest represents the PrepareQueryRequest schema from the OpenAPI specification
type PrepareQueryRequest struct {
	Querystring interface{} `json:"QueryString"`
	Validateonly interface{} `json:"ValidateOnly,omitempty"`
}

// TimestreamConfiguration represents the TimestreamConfiguration schema from the OpenAPI specification
type TimestreamConfiguration struct {
	Databasename interface{} `json:"DatabaseName"`
	Dimensionmappings interface{} `json:"DimensionMappings"`
	Measurenamecolumn interface{} `json:"MeasureNameColumn,omitempty"`
	Mixedmeasuremappings interface{} `json:"MixedMeasureMappings,omitempty"`
	Multimeasuremappings interface{} `json:"MultiMeasureMappings,omitempty"`
	Tablename interface{} `json:"TableName"`
	Timecolumn interface{} `json:"TimeColumn"`
}

// DescribeScheduledQueryRequest represents the DescribeScheduledQueryRequest schema from the OpenAPI specification
type DescribeScheduledQueryRequest struct {
	Scheduledqueryarn interface{} `json:"ScheduledQueryArn"`
}

// SelectColumn represents the SelectColumn schema from the OpenAPI specification
type SelectColumn struct {
	Tablename interface{} `json:"TableName,omitempty"`
	TypeField Type `json:"Type,omitempty"` // Contains the data type of a column in a query result set. The data type can be scalar or complex. The supported scalar data types are integers, Boolean, string, double, timestamp, date, time, and intervals. The supported complex data types are arrays, rows, and timeseries.
	Aliased interface{} `json:"Aliased,omitempty"`
	Databasename interface{} `json:"DatabaseName,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// MultiMeasureMappings represents the MultiMeasureMappings schema from the OpenAPI specification
type MultiMeasureMappings struct {
	Multimeasureattributemappings interface{} `json:"MultiMeasureAttributeMappings"`
	Targetmultimeasurename interface{} `json:"TargetMultiMeasureName,omitempty"`
}

// DeleteScheduledQueryRequest represents the DeleteScheduledQueryRequest schema from the OpenAPI specification
type DeleteScheduledQueryRequest struct {
	Scheduledqueryarn interface{} `json:"ScheduledQueryArn"`
}

// Datum represents the Datum schema from the OpenAPI specification
type Datum struct {
	Arrayvalue interface{} `json:"ArrayValue,omitempty"`
	Nullvalue interface{} `json:"NullValue,omitempty"`
	Rowvalue interface{} `json:"RowValue,omitempty"`
	Scalarvalue interface{} `json:"ScalarValue,omitempty"`
	Timeseriesvalue interface{} `json:"TimeSeriesValue,omitempty"`
}

// CancelQueryResponse represents the CancelQueryResponse schema from the OpenAPI specification
type CancelQueryResponse struct {
	Cancellationmessage interface{} `json:"CancellationMessage,omitempty"`
}

// CreateScheduledQueryRequest represents the CreateScheduledQueryRequest schema from the OpenAPI specification
type CreateScheduledQueryRequest struct {
	Notificationconfiguration interface{} `json:"NotificationConfiguration"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Scheduledqueryexecutionrolearn interface{} `json:"ScheduledQueryExecutionRoleArn"`
	Targetconfiguration interface{} `json:"TargetConfiguration,omitempty"`
	Scheduleconfiguration interface{} `json:"ScheduleConfiguration"`
	Errorreportconfiguration interface{} `json:"ErrorReportConfiguration"`
	Querystring interface{} `json:"QueryString"`
	Tags interface{} `json:"Tags,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Name interface{} `json:"Name"`
}

// ListScheduledQueriesRequest represents the ListScheduledQueriesRequest schema from the OpenAPI specification
type ListScheduledQueriesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// MixedMeasureMapping represents the MixedMeasureMapping schema from the OpenAPI specification
type MixedMeasureMapping struct {
	Targetmeasurename interface{} `json:"TargetMeasureName,omitempty"`
	Measurename interface{} `json:"MeasureName,omitempty"`
	Measurevaluetype interface{} `json:"MeasureValueType"`
	Multimeasureattributemappings interface{} `json:"MultiMeasureAttributeMappings,omitempty"`
	Sourcecolumn interface{} `json:"SourceColumn,omitempty"`
}

// QueryRequest represents the QueryRequest schema from the OpenAPI specification
type QueryRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Querystring interface{} `json:"QueryString"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Maxrows interface{} `json:"MaxRows,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceARN"`
	Tags interface{} `json:"Tags"`
}

// ListScheduledQueriesResponse represents the ListScheduledQueriesResponse schema from the OpenAPI specification
type ListScheduledQueriesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Scheduledqueries interface{} `json:"ScheduledQueries"`
}

// S3Configuration represents the S3Configuration schema from the OpenAPI specification
type S3Configuration struct {
	Bucketname interface{} `json:"BucketName"`
	Encryptionoption interface{} `json:"EncryptionOption,omitempty"`
	Objectkeyprefix interface{} `json:"ObjectKeyPrefix,omitempty"`
}

// TargetConfiguration represents the TargetConfiguration schema from the OpenAPI specification
type TargetConfiguration struct {
	Timestreamconfiguration interface{} `json:"TimestreamConfiguration"`
}

// ColumnInfo represents the ColumnInfo schema from the OpenAPI specification
type ColumnInfo struct {
	Name interface{} `json:"Name,omitempty"`
	TypeField interface{} `json:"Type"`
}

// DimensionMapping represents the DimensionMapping schema from the OpenAPI specification
type DimensionMapping struct {
	Dimensionvaluetype interface{} `json:"DimensionValueType"`
	Name interface{} `json:"Name"`
}

// ParameterMapping represents the ParameterMapping schema from the OpenAPI specification
type ParameterMapping struct {
	Name interface{} `json:"Name"`
	TypeField Type `json:"Type"` // Contains the data type of a column in a query result set. The data type can be scalar or complex. The supported scalar data types are integers, Boolean, string, double, timestamp, date, time, and intervals. The supported complex data types are arrays, rows, and timeseries.
}

// CreateScheduledQueryResponse represents the CreateScheduledQueryResponse schema from the OpenAPI specification
type CreateScheduledQueryResponse struct {
	Arn interface{} `json:"Arn"`
}

// UpdateScheduledQueryRequest represents the UpdateScheduledQueryRequest schema from the OpenAPI specification
type UpdateScheduledQueryRequest struct {
	Scheduledqueryarn interface{} `json:"ScheduledQueryArn"`
	State interface{} `json:"State"`
}

// CancelQueryRequest represents the CancelQueryRequest schema from the OpenAPI specification
type CancelQueryRequest struct {
	Queryid interface{} `json:"QueryId"`
}

// ExecutionStats represents the ExecutionStats schema from the OpenAPI specification
type ExecutionStats struct {
	Executiontimeinmillis interface{} `json:"ExecutionTimeInMillis,omitempty"`
	Queryresultrows interface{} `json:"QueryResultRows,omitempty"`
	Recordsingested interface{} `json:"RecordsIngested,omitempty"`
	Bytesmetered interface{} `json:"BytesMetered,omitempty"`
	Datawrites interface{} `json:"DataWrites,omitempty"`
}

// TimestreamDestination represents the TimestreamDestination schema from the OpenAPI specification
type TimestreamDestination struct {
	Tablename interface{} `json:"TableName,omitempty"`
	Databasename interface{} `json:"DatabaseName,omitempty"`
}

// Row represents the Row schema from the OpenAPI specification
type Row struct {
	Data interface{} `json:"Data"`
}

// PrepareQueryResponse represents the PrepareQueryResponse schema from the OpenAPI specification
type PrepareQueryResponse struct {
	Columns interface{} `json:"Columns"`
	Parameters interface{} `json:"Parameters"`
	Querystring interface{} `json:"QueryString"`
}

// ScheduledQueryDescription represents the ScheduledQueryDescription schema from the OpenAPI specification
type ScheduledQueryDescription struct {
	Name interface{} `json:"Name"`
	State interface{} `json:"State"`
	Notificationconfiguration interface{} `json:"NotificationConfiguration"`
	Previousinvocationtime interface{} `json:"PreviousInvocationTime,omitempty"`
	Querystring interface{} `json:"QueryString"`
	Nextinvocationtime interface{} `json:"NextInvocationTime,omitempty"`
	Errorreportconfiguration interface{} `json:"ErrorReportConfiguration,omitempty"`
	Arn interface{} `json:"Arn"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Recentlyfailedruns interface{} `json:"RecentlyFailedRuns,omitempty"`
	Targetconfiguration interface{} `json:"TargetConfiguration,omitempty"`
	Scheduleconfiguration interface{} `json:"ScheduleConfiguration"`
	Scheduledqueryexecutionrolearn interface{} `json:"ScheduledQueryExecutionRoleArn,omitempty"`
	Lastrunsummary interface{} `json:"LastRunSummary,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
}

// QueryStatus represents the QueryStatus schema from the OpenAPI specification
type QueryStatus struct {
	Progresspercentage interface{} `json:"ProgressPercentage,omitempty"`
	Cumulativebytesmetered interface{} `json:"CumulativeBytesMetered,omitempty"`
	Cumulativebytesscanned interface{} `json:"CumulativeBytesScanned,omitempty"`
}

// MultiMeasureAttributeMapping represents the MultiMeasureAttributeMapping schema from the OpenAPI specification
type MultiMeasureAttributeMapping struct {
	Measurevaluetype interface{} `json:"MeasureValueType"`
	Sourcecolumn interface{} `json:"SourceColumn"`
	Targetmultimeasureattributename interface{} `json:"TargetMultiMeasureAttributeName,omitempty"`
}

// ScheduledQueryRunSummary represents the ScheduledQueryRunSummary schema from the OpenAPI specification
type ScheduledQueryRunSummary struct {
	Runstatus interface{} `json:"RunStatus,omitempty"`
	Triggertime interface{} `json:"TriggerTime,omitempty"`
	Errorreportlocation interface{} `json:"ErrorReportLocation,omitempty"`
	Executionstats interface{} `json:"ExecutionStats,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Invocationtime interface{} `json:"InvocationTime,omitempty"`
}

// DescribeEndpointsRequest represents the DescribeEndpointsRequest schema from the OpenAPI specification
type DescribeEndpointsRequest struct {
}

// TargetDestination represents the TargetDestination schema from the OpenAPI specification
type TargetDestination struct {
	Timestreamdestination interface{} `json:"TimestreamDestination,omitempty"`
}

// DescribeScheduledQueryResponse represents the DescribeScheduledQueryResponse schema from the OpenAPI specification
type DescribeScheduledQueryResponse struct {
	Scheduledquery interface{} `json:"ScheduledQuery"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resourcearn interface{} `json:"ResourceARN"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// Endpoint represents the Endpoint schema from the OpenAPI specification
type Endpoint struct {
	Cacheperiodinminutes interface{} `json:"CachePeriodInMinutes"`
	Address interface{} `json:"Address"`
}

// NotificationConfiguration represents the NotificationConfiguration schema from the OpenAPI specification
type NotificationConfiguration struct {
	Snsconfiguration interface{} `json:"SnsConfiguration"`
}

// SnsConfiguration represents the SnsConfiguration schema from the OpenAPI specification
type SnsConfiguration struct {
	Topicarn interface{} `json:"TopicArn"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Tags interface{} `json:"Tags"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// S3ReportLocation represents the S3ReportLocation schema from the OpenAPI specification
type S3ReportLocation struct {
	Objectkey interface{} `json:"ObjectKey,omitempty"`
	Bucketname interface{} `json:"BucketName,omitempty"`
}

// ExecuteScheduledQueryRequest represents the ExecuteScheduledQueryRequest schema from the OpenAPI specification
type ExecuteScheduledQueryRequest struct {
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Invocationtime interface{} `json:"InvocationTime"`
	Scheduledqueryarn interface{} `json:"ScheduledQueryArn"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Tagkeys interface{} `json:"TagKeys"`
	Resourcearn interface{} `json:"ResourceARN"`
}

// ErrorReportLocation represents the ErrorReportLocation schema from the OpenAPI specification
type ErrorReportLocation struct {
	S3reportlocation interface{} `json:"S3ReportLocation,omitempty"`
}

// Type represents the Type schema from the OpenAPI specification
type Type struct {
	Scalartype interface{} `json:"ScalarType,omitempty"`
	Timeseriesmeasurevaluecolumninfo interface{} `json:"TimeSeriesMeasureValueColumnInfo,omitempty"`
	Arraycolumninfo interface{} `json:"ArrayColumnInfo,omitempty"`
	Rowcolumninfo interface{} `json:"RowColumnInfo,omitempty"`
}
