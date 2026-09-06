package aws

import (
	"context"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const (
	requestChecksumCalculationParamKey = "request_checksum_calculation"
	responseChecksumValidationParamKey = "response_checksum_validation"
)

func parseRequestChecksumCalculation(value string) (aws.RequestChecksumCalculation, error) {
	_ = "STUB: not implemented"
	return *new(aws.RequestChecksumCalculation), nil
}

func parseResponseChecksumValidation(value string) (aws.ResponseChecksumValidation, error) {
	_ = "STUB: not implemented"
	return *new(aws.ResponseChecksumValidation), nil
}

func NewDefaultV2Config(ctx context.Context) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}

func V2ConfigFromURLParams(ctx context.Context, q url.Values) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}
