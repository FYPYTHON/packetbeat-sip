// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Contains the parameters for ModifySnapshotAttribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifySnapshotAttributeRequest
type ModifySnapshotAttributeInput struct {
	_ struct{} `type:"structure"`

	// The snapshot attribute to modify. Only volume creation permissions can be
	// modified.
	Attribute SnapshotAttributeName `type:"string" enum:"true"`

	// A JSON representation of the snapshot attribute modification.
	CreateVolumePermission *CreateVolumePermissionModifications `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The group to modify for the snapshot.
	GroupNames []string `locationName:"UserGroup" locationNameList:"GroupName" type:"list"`

	// The type of operation to perform to the attribute.
	OperationType OperationType `type:"string" enum:"true"`

	// The ID of the snapshot.
	//
	// SnapshotId is a required field
	SnapshotId *string `type:"string" required:"true"`

	// The account ID to modify for the snapshot.
	UserIds []string `locationName:"UserId" locationNameList:"UserId" type:"list"`
}

// String returns the string representation
func (s ModifySnapshotAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifySnapshotAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifySnapshotAttributeInput"}

	if s.SnapshotId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifySnapshotAttributeOutput
type ModifySnapshotAttributeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifySnapshotAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifySnapshotAttribute = "ModifySnapshotAttribute"

// ModifySnapshotAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Adds or removes permission settings for the specified snapshot. You may add
// or remove specified AWS account IDs from a snapshot's list of create volume
// permissions, but you cannot do both in a single API call. If you need to
// both add and remove account IDs for a snapshot, you must use multiple API
// calls.
//
// Encrypted snapshots and snapshots with AWS Marketplace product codes cannot
// be made public. Snapshots encrypted with your default CMK cannot be shared
// with other accounts.
//
// For more information about modifying snapshot permissions, see Sharing Snapshots
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-modifying-snapshot-permissions.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using ModifySnapshotAttributeRequest.
//    req := client.ModifySnapshotAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifySnapshotAttribute
func (c *Client) ModifySnapshotAttributeRequest(input *ModifySnapshotAttributeInput) ModifySnapshotAttributeRequest {
	op := &aws.Operation{
		Name:       opModifySnapshotAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifySnapshotAttributeInput{}
	}

	req := c.newRequest(op, input, &ModifySnapshotAttributeOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return ModifySnapshotAttributeRequest{Request: req, Input: input, Copy: c.ModifySnapshotAttributeRequest}
}

// ModifySnapshotAttributeRequest is the request type for the
// ModifySnapshotAttribute API operation.
type ModifySnapshotAttributeRequest struct {
	*aws.Request
	Input *ModifySnapshotAttributeInput
	Copy  func(*ModifySnapshotAttributeInput) ModifySnapshotAttributeRequest
}

// Send marshals and sends the ModifySnapshotAttribute API request.
func (r ModifySnapshotAttributeRequest) Send(ctx context.Context) (*ModifySnapshotAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifySnapshotAttributeResponse{
		ModifySnapshotAttributeOutput: r.Request.Data.(*ModifySnapshotAttributeOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifySnapshotAttributeResponse is the response type for the
// ModifySnapshotAttribute API operation.
type ModifySnapshotAttributeResponse struct {
	*ModifySnapshotAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifySnapshotAttribute request.
func (r *ModifySnapshotAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
