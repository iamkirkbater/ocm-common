// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates the specified subnets and transit gateway attachments with the
// specified transit gateway multicast domain. The transit gateway attachment must
// be in the available state before you can add a resource. Use
// DescribeTransitGatewayAttachments (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayAttachments.html)
// to see the state of the attachment.
func (c *Client) AssociateTransitGatewayMulticastDomain(ctx context.Context, params *AssociateTransitGatewayMulticastDomainInput, optFns ...func(*Options)) (*AssociateTransitGatewayMulticastDomainOutput, error) {
	if params == nil {
		params = &AssociateTransitGatewayMulticastDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateTransitGatewayMulticastDomain", params, optFns, c.addOperationAssociateTransitGatewayMulticastDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateTransitGatewayMulticastDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateTransitGatewayMulticastDomainInput struct {

	// The IDs of the subnets to associate with the transit gateway multicast domain.
	//
	// This member is required.
	SubnetIds []string

	// The ID of the transit gateway attachment to associate with the transit gateway
	// multicast domain.
	//
	// This member is required.
	TransitGatewayAttachmentId *string

	// The ID of the transit gateway multicast domain.
	//
	// This member is required.
	TransitGatewayMulticastDomainId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type AssociateTransitGatewayMulticastDomainOutput struct {

	// Information about the transit gateway multicast domain associations.
	Associations *types.TransitGatewayMulticastDomainAssociations

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateTransitGatewayMulticastDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpAssociateTransitGatewayMulticastDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAssociateTransitGatewayMulticastDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateTransitGatewayMulticastDomain"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpAssociateTransitGatewayMulticastDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateTransitGatewayMulticastDomain(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAssociateTransitGatewayMulticastDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateTransitGatewayMulticastDomain",
	}
}