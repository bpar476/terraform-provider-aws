package s3

// Error code constants missing from AWS Go SDK:
// https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#pkg-constants

const (
	ErrCodeMethodNotAllowed                     = "MethodNotAllowed"
	ErrCodeNoSuchConfiguration                  = "NoSuchConfiguration"
	ErrCodeNoSuchCORSConfiguration              = "NoSuchCORSConfiguration"
	ErrCodeNoSuchPublicAccessBlockConfiguration = "NoSuchPublicAccessBlockConfiguration"
	ErrCodeObjectLockConfigurationNotFound      = "ObjectLockConfigurationNotFoundError"
	ErrCodeOperationAborted                     = "OperationAborted"
)
