package awssnssqs

import (
	"context"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	snstypes "github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	sqstypes "github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/google/wire"
	"gocloud.dev/gcerrors"
	"gocloud.dev/pubsub"
	"gocloud.dev/pubsub/batcher"
	"gocloud.dev/pubsub/driver"
)

const (
	base64EncodedKey = "base64encoded"

	noMessagesPollDuration = 250 * time.Millisecond
)

var sendBatcherOptsSNS = &batcher.Options{
	MaxBatchSize: 10,
	MaxHandlers:  100,
}

var sendBatcherOptsSQS = &batcher.Options{
	MaxBatchSize: 10,
	MaxHandlers:  100,
}

var recvBatcherOpts = &batcher.Options{

	MaxBatchSize: 10,
	MaxHandlers:  100,
}

var ackBatcherOpts = &batcher.Options{

	MaxBatchSize: 10,
	MaxHandlers:  100,
}

func init() {
	lazy := new(lazySessionOpener)
	pubsub.DefaultURLMux().RegisterTopic(SNSScheme, lazy)
	pubsub.DefaultURLMux().RegisterTopic(SQSScheme, lazy)
	pubsub.DefaultURLMux().RegisterSubscription(SQSScheme, lazy)
}

var Set = wire.NewSet(
	DialSNS,
	DialSQS,
)

func DialSNS(cfg aws.Config) *sns.Client { _ = "STUB: not implemented"; return nil }

func DialSQS(cfg aws.Config) *sqs.Client { _ = "STUB: not implemented"; return nil }

type lazySessionOpener struct{}

func (o *lazySessionOpener) defaultOpener(u *url.URL) (*URLOpener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *lazySessionOpener) OpenTopicURL(ctx context.Context, u *url.URL) (*pubsub.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *lazySessionOpener) OpenSubscriptionURL(ctx context.Context, u *url.URL) (*pubsub.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const SNSScheme = "awssns"

const SQSScheme = "awssqs"

type URLOpener struct {
	TopicOptions TopicOptions

	SubscriptionOptions SubscriptionOptions
}

func (o *URLOpener) OpenTopicURL(ctx context.Context, u *url.URL) (*pubsub.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *URLOpener) OpenSubscriptionURL(ctx context.Context, u *url.URL) (*pubsub.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type snsTopic struct {
	client *sns.Client
	arn    string
	opts   *TopicOptions
}

type BodyBase64Encoding int

const (
	NonUTF8Only BodyBase64Encoding = 0

	Always BodyBase64Encoding = 1

	Never BodyBase64Encoding = 2
)

func (e BodyBase64Encoding) wantEncode(b []byte) bool { _ = "STUB: not implemented"; return false }

type TopicOptions struct {
	BodyBase64Encoding BodyBase64Encoding

	BatcherOptions batcher.Options
}

func OpenSNSTopic(ctx context.Context, client *sns.Client, topicARN string, opts *TopicOptions) *pubsub.Topic {
	_ = "STUB: not implemented"
	return nil
}

var OpenSNSTopicV2 = OpenSNSTopic

func openSNSTopic(ctx context.Context, client *sns.Client, topicARN string, opts *TopicOptions) driver.Topic {
	_ = "STUB: not implemented"
	return *new(driver.Topic)
}

var stringDataType = aws.String("String")

func encodeMetadata(md map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

func maybeEncodeBody(body []byte, opt BodyBase64Encoding) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

const (
	MetadataKeyDeduplicationID = "DeduplicationId"
	MetadataKeyMessageGroupID  = "MessageGroupId"
)

func reviseSnsEntryAttributes(dm *driver.Message, entry *snstypes.PublishBatchRequestEntry) {
	_ = "STUB: not implemented"
	return
}

func (t *snsTopic) SendBatch(ctx context.Context, dms []*driver.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *snsTopic) IsRetryable(error) bool { _ = "STUB: not implemented"; return false }

func (t *snsTopic) As(i any) bool { _ = "STUB: not implemented"; return false }

func (t *snsTopic) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (t *snsTopic) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (*snsTopic) Close() error { _ = "STUB: not implemented"; return nil }

type sqsTopic struct {
	client *sqs.Client
	qURL   string
	opts   *TopicOptions
}

func OpenSQSTopic(ctx context.Context, client *sqs.Client, qURL string, opts *TopicOptions) *pubsub.Topic {
	_ = "STUB: not implemented"
	return nil
}

var OpenSQSTopicV2 = OpenSQSTopic

func openSQSTopic(ctx context.Context, client *sqs.Client, qURL string, opts *TopicOptions) driver.Topic {
	_ = "STUB: not implemented"
	return *new(driver.Topic)
}

func reviseSqsEntryAttributes(dm *driver.Message, entry *sqstypes.SendMessageBatchRequestEntry) {
	_ = "STUB: not implemented"
	return
}

func (t *sqsTopic) SendBatch(ctx context.Context, dms []*driver.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *sqsTopic) IsRetryable(error) bool { _ = "STUB: not implemented"; return false }

func (t *sqsTopic) As(i any) bool { _ = "STUB: not implemented"; return false }

func (t *sqsTopic) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (t *sqsTopic) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (*sqsTopic) Close() error { _ = "STUB: not implemented"; return nil }

func errorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

var errorCodeMap = map[string]gcerrors.ErrorCode{
	(&snstypes.AuthorizationErrorException{}).ErrorCode():        gcerrors.PermissionDenied,
	(&snstypes.KMSAccessDeniedException{}).ErrorCode():           gcerrors.PermissionDenied,
	(&snstypes.KMSDisabledException{}).ErrorCode():               gcerrors.FailedPrecondition,
	(&snstypes.KMSInvalidStateException{}).ErrorCode():           gcerrors.FailedPrecondition,
	(&snstypes.KMSOptInRequired{}).ErrorCode():                   gcerrors.FailedPrecondition,
	(&sqstypes.MessageNotInflight{}).ErrorCode():                 gcerrors.FailedPrecondition,
	(&sqstypes.PurgeQueueInProgress{}).ErrorCode():               gcerrors.FailedPrecondition,
	(&sqstypes.QueueDeletedRecently{}).ErrorCode():               gcerrors.FailedPrecondition,
	(&sqstypes.QueueNameExists{}).ErrorCode():                    gcerrors.FailedPrecondition,
	(&snstypes.InternalErrorException{}).ErrorCode():             gcerrors.Internal,
	(&snstypes.InvalidParameterException{}).ErrorCode():          gcerrors.InvalidArgument,
	(&snstypes.InvalidParameterValueException{}).ErrorCode():     gcerrors.InvalidArgument,
	(&sqstypes.BatchEntryIdsNotDistinct{}).ErrorCode():           gcerrors.InvalidArgument,
	(&sqstypes.BatchRequestTooLong{}).ErrorCode():                gcerrors.InvalidArgument,
	(&sqstypes.EmptyBatchRequest{}).ErrorCode():                  gcerrors.InvalidArgument,
	(&sqstypes.InvalidAttributeName{}).ErrorCode():               gcerrors.InvalidArgument,
	(&sqstypes.InvalidBatchEntryId{}).ErrorCode():                gcerrors.InvalidArgument,
	(&sqstypes.InvalidIdFormat{}).ErrorCode():                    gcerrors.InvalidArgument,
	(&sqstypes.InvalidMessageContents{}).ErrorCode():             gcerrors.InvalidArgument,
	(&sqstypes.ReceiptHandleIsInvalid{}).ErrorCode():             gcerrors.InvalidArgument,
	(&sqstypes.TooManyEntriesInBatchRequest{}).ErrorCode():       gcerrors.InvalidArgument,
	(&sqstypes.UnsupportedOperation{}).ErrorCode():               gcerrors.InvalidArgument,
	(&snstypes.InvalidSecurityException{}).ErrorCode():           gcerrors.PermissionDenied,
	(&snstypes.KMSNotFoundException{}).ErrorCode():               gcerrors.NotFound,
	(&snstypes.NotFoundException{}).ErrorCode():                  gcerrors.NotFound,
	(&sqstypes.QueueDoesNotExist{}).ErrorCode():                  gcerrors.NotFound,
	"AWS.SimpleQueueService.NonExistentQueue":                    gcerrors.NotFound,
	(&snstypes.FilterPolicyLimitExceededException{}).ErrorCode(): gcerrors.ResourceExhausted,
	(&snstypes.SubscriptionLimitExceededException{}).ErrorCode(): gcerrors.ResourceExhausted,
	(&snstypes.TopicLimitExceededException{}).ErrorCode():        gcerrors.ResourceExhausted,
	(&sqstypes.OverLimit{}).ErrorCode():                          gcerrors.ResourceExhausted,
	(&snstypes.KMSThrottlingException{}).ErrorCode():             gcerrors.ResourceExhausted,
	(&snstypes.ThrottledException{}).ErrorCode():                 gcerrors.ResourceExhausted,
	"RequestCanceled": gcerrors.Canceled,
	(&snstypes.EndpointDisabledException{}).ErrorCode():            gcerrors.Unknown,
	(&snstypes.PlatformApplicationDisabledException{}).ErrorCode(): gcerrors.Unknown,
}

type subscription struct {
	client *sqs.Client
	qURL   string
	opts   *SubscriptionOptions
}

type SubscriptionOptions struct {
	Raw bool

	NackLazy bool

	WaitTime time.Duration

	ReceiveBatcherOptions batcher.Options

	AckBatcherOptions batcher.Options
}

func OpenSubscription(ctx context.Context, client *sqs.Client, qURL string, opts *SubscriptionOptions) *pubsub.Subscription {
	_ = "STUB: not implemented"
	return nil
}

var OpenSubscriptionV2 = OpenSubscription

func openSubscription(ctx context.Context, client *sqs.Client, qURL string, opts *SubscriptionOptions) driver.Subscription {
	_ = "STUB: not implemented"
	return *new(driver.Subscription)
}

func (s *subscription) ReceiveBatch(ctx context.Context, maxMessages int) ([]*driver.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractBody(bodyStr string, rawAttrs map[string]string, raw bool) (body string, attributes map[string]string) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *subscription) SendAcks(ctx context.Context, ids []driver.AckID) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscription) CanNack() bool { _ = "STUB: not implemented"; return false }

func (s *subscription) SendNacks(ctx context.Context, ids []driver.AckID) error {
	_ = "STUB: not implemented"
	return nil
}

func (*subscription) IsRetryable(error) bool { _ = "STUB: not implemented"; return false }

func (s *subscription) As(i any) bool { _ = "STUB: not implemented"; return false }

func (s *subscription) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (s *subscription) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func errorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (*subscription) Close() error { _ = "STUB: not implemented"; return nil }
