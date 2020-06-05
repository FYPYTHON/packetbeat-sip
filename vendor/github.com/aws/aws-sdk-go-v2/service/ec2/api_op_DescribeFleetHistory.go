// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFleetHistoryRequest
type DescribeFleetHistoryInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The type of events to describe. By default, all events are described.
	EventType FleetEventType `type:"string" enum:"true"`

	// The ID of the EC2 Fleet.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`

	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults *int64 `type:"integer"`

	// The token for the next set of results.
	NextToken *string `type:"string"`

	// The start date and time for the events, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	//
	// StartTime is a required field
	StartTime *time.Time `type:"timestamp" timestampFormat:"iso8601" required:"true"`
}

// String returns the string representation
func (s DescribeFleetHistoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFleetHistoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFleetHistoryInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFleetHistoryResult
type DescribeFleetHistoryOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the EC Fleet.
	FleetId *string `locationName:"fleetId" type:"string"`

	// Information about the events in the history of the EC2 Fleet.
	HistoryRecords []HistoryRecordEntry `locationName:"historyRecordSet" locationNameList:"item" type:"list"`

	// The last date and time for the events, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	// All records up to this time were retrieved.
	//
	// If nextToken indicates that there are more results, this value is not present.
	LastEvaluatedTime *time.Time `locationName:"lastEvaluatedTime" type:"timestamp" timestampFormat:"iso8601"`

	// The token for the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The start date and time for the events, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime *time.Time `locationName:"startTime" type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s DescribeFleetHistoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeFleetHistory = "DescribeFleetHistory"

// DescribeFleetHistoryRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the events for the specified EC2 Fleet during the specified time.
//
//    // Example sending a request using DescribeFleetHistoryRequest.
//    req := client.DescribeFleetHistoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFleetHistory
func (c *Client) DescribeFleetHistoryRequest(input *DescribeFleetHistoryInput) DescribeFleetHistoryRequest {
	op := &aws.Operation{
		Name:       opDescribeFleetHistory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeFleetHistoryInput{}
	}

	req := c.newRequest(op, input, &DescribeFleetHistoryOutput{})
	return DescribeFleetHistoryRequest{Request: req, Input: input, Copy: c.DescribeFleetHistoryRequest}
}

// DescribeFleetHistoryRequest is the request type for the
// DescribeFleetHistory API operation.
type DescribeFleetHistoryRequest struct {
	*aws.Request
	Input *DescribeFleetHistoryInput
	Copy  func(*DescribeFleetHistoryInput) DescribeFleetHistoryRequest
}

// Send marshals and sends the DescribeFleetHistory API request.
func (r DescribeFleetHistoryRequest) Send(ctx context.Context) (*DescribeFleetHistoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFleetHistoryResponse{
		DescribeFleetHistoryOutput: r.Request.Data.(*DescribeFleetHistoryOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeFleetHistoryResponse is the response type for the
// DescribeFleetHistory API operation.
type DescribeFleetHistoryResponse struct {
	*DescribeFleetHistoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFleetHistory request.
func (r *DescribeFleetHistoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
