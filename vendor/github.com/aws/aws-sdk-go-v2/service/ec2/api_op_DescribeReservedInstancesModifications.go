// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeReservedInstancesModifications.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeReservedInstancesModificationsRequest
type DescribeReservedInstancesModificationsInput struct {
	_ struct{} `type:"structure"`

	// One or more filters.
	//
	//    * client-token - The idempotency token for the modification request.
	//
	//    * create-date - The time when the modification request was created.
	//
	//    * effective-date - The time when the modification becomes effective.
	//
	//    * modification-result.reserved-instances-id - The ID for the Reserved
	//    Instances created as part of the modification request. This ID is only
	//    available when the status of the modification is fulfilled.
	//
	//    * modification-result.target-configuration.availability-zone - The Availability
	//    Zone for the new Reserved Instances.
	//
	//    * modification-result.target-configuration.instance-count - The number
	//    of new Reserved Instances.
	//
	//    * modification-result.target-configuration.instance-type - The instance
	//    type of the new Reserved Instances.
	//
	//    * modification-result.target-configuration.platform - The network platform
	//    of the new Reserved Instances (EC2-Classic | EC2-VPC).
	//
	//    * reserved-instances-id - The ID of the Reserved Instances modified.
	//
	//    * reserved-instances-modification-id - The ID of the modification request.
	//
	//    * status - The status of the Reserved Instances modification request (processing
	//    | fulfilled | failed).
	//
	//    * status-message - The reason for the status.
	//
	//    * update-date - The time when the modification request was last updated.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The token to retrieve the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// IDs for the submitted modification request.
	ReservedInstancesModificationIds []string `locationName:"ReservedInstancesModificationId" locationNameList:"ReservedInstancesModificationId" type:"list"`
}

// String returns the string representation
func (s DescribeReservedInstancesModificationsInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output of DescribeReservedInstancesModifications.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeReservedInstancesModificationsResult
type DescribeReservedInstancesModificationsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The Reserved Instance modification information.
	ReservedInstancesModifications []ReservedInstancesModification `locationName:"reservedInstancesModificationsSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeReservedInstancesModificationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeReservedInstancesModifications = "DescribeReservedInstancesModifications"

// DescribeReservedInstancesModificationsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the modifications made to your Reserved Instances. If no parameter
// is specified, information about all your Reserved Instances modification
// requests is returned. If a modification ID is specified, only information
// about the specific modification is returned.
//
// For more information, see Modifying Reserved Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-modifying.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeReservedInstancesModificationsRequest.
//    req := client.DescribeReservedInstancesModificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeReservedInstancesModifications
func (c *Client) DescribeReservedInstancesModificationsRequest(input *DescribeReservedInstancesModificationsInput) DescribeReservedInstancesModificationsRequest {
	op := &aws.Operation{
		Name:       opDescribeReservedInstancesModifications,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeReservedInstancesModificationsInput{}
	}

	req := c.newRequest(op, input, &DescribeReservedInstancesModificationsOutput{})
	return DescribeReservedInstancesModificationsRequest{Request: req, Input: input, Copy: c.DescribeReservedInstancesModificationsRequest}
}

// DescribeReservedInstancesModificationsRequest is the request type for the
// DescribeReservedInstancesModifications API operation.
type DescribeReservedInstancesModificationsRequest struct {
	*aws.Request
	Input *DescribeReservedInstancesModificationsInput
	Copy  func(*DescribeReservedInstancesModificationsInput) DescribeReservedInstancesModificationsRequest
}

// Send marshals and sends the DescribeReservedInstancesModifications API request.
func (r DescribeReservedInstancesModificationsRequest) Send(ctx context.Context) (*DescribeReservedInstancesModificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeReservedInstancesModificationsResponse{
		DescribeReservedInstancesModificationsOutput: r.Request.Data.(*DescribeReservedInstancesModificationsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeReservedInstancesModificationsRequestPaginator returns a paginator for DescribeReservedInstancesModifications.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeReservedInstancesModificationsRequest(input)
//   p := ec2.NewDescribeReservedInstancesModificationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeReservedInstancesModificationsPaginator(req DescribeReservedInstancesModificationsRequest) DescribeReservedInstancesModificationsPaginator {
	return DescribeReservedInstancesModificationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeReservedInstancesModificationsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeReservedInstancesModificationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeReservedInstancesModificationsPaginator struct {
	aws.Pager
}

func (p *DescribeReservedInstancesModificationsPaginator) CurrentPage() *DescribeReservedInstancesModificationsOutput {
	return p.Pager.CurrentPage().(*DescribeReservedInstancesModificationsOutput)
}

// DescribeReservedInstancesModificationsResponse is the response type for the
// DescribeReservedInstancesModifications API operation.
type DescribeReservedInstancesModificationsResponse struct {
	*DescribeReservedInstancesModificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeReservedInstancesModifications request.
func (r *DescribeReservedInstancesModificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
