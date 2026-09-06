package awssecretsmanager

import (
	"context"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/google/wire"
	"gocloud.dev/gcerrors"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/driver"
)

func init() {
	runtimevar.DefaultURLMux().RegisterVariable(Scheme, new(lazySessionOpener))
}

var Set = wire.NewSet(
	Dial,
)

func Dial(cfg aws.Config) *secretsmanager.Client { _ = "STUB: not implemented"; return nil }

type URLOpener struct {
	Decoder *runtimevar.Decoder

	Options Options
}

type lazySessionOpener struct{}

func (o *lazySessionOpener) OpenVariableURL(ctx context.Context, u *url.URL) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "awssecretsmanager"

func (o *URLOpener) OpenVariableURL(ctx context.Context, u *url.URL) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Options struct {
	WaitDuration time.Duration
}

func OpenVariable(client *secretsmanager.Client, name string, decoder *runtimevar.Decoder, opts *Options) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var OpenVariableV2 = OpenVariable

type state struct {
	val        any
	rawGet     *secretsmanager.GetSecretValueOutput
	rawDesc    *secretsmanager.DescribeSecretOutput
	updateTime time.Time
	versionID  string
	err        error
}

func (s *state) Value() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (s *state) UpdateTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *state) As(i any) bool { _ = "STUB: not implemented"; return false }

func errorState(err error, prevS driver.State) driver.State {
	_ = "STUB: not implemented"
	return *new(driver.State)
}

func equivalentError(err1, err2 error) bool { _ = "STUB: not implemented"; return false }

type watcher struct {
	client *secretsmanager.Client

	name string

	wait time.Duration

	decoder *runtimevar.Decoder
}

func newWatcher(client *secretsmanager.Client, name string, decoder *runtimevar.Decoder, opts *Options) *watcher {
	_ = "STUB: not implemented"
	return nil
}

func getSecretValue(ctx context.Context, client *secretsmanager.Client, secretID string) (string, []byte, string, *secretsmanager.GetSecretValueOutput, error) {
	_ = "STUB: not implemented"
	return "", nil, "", nil, nil
}

func describeSecret(ctx context.Context, client *secretsmanager.Client, secretID string) (time.Time, *secretsmanager.DescribeSecretOutput, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil, nil
}

func (w *watcher) WatchVariable(ctx context.Context, prev driver.State) (driver.State, time.Duration) {
	_ = "STUB: not implemented"
	return *new(driver.State), *new(time.Duration)
}

func (w *watcher) Close() error { _ = "STUB: not implemented"; return nil }

func (w *watcher) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func getErrorCode(err error) string { _ = "STUB: not implemented"; return "" }

func (w *watcher) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

var errorCodeMap = map[string]gcerrors.ErrorCode{
	(&types.ResourceNotFoundException{}).ErrorCode():        gcerrors.NotFound,
	(&types.InvalidParameterException{}).ErrorCode():        gcerrors.InvalidArgument,
	(&types.InvalidRequestException{}).ErrorCode():          gcerrors.InvalidArgument,
	(&types.InvalidNextTokenException{}).ErrorCode():        gcerrors.InvalidArgument,
	(&types.EncryptionFailure{}).ErrorCode():                gcerrors.Internal,
	(&types.DecryptionFailure{}).ErrorCode():                gcerrors.Internal,
	(&types.InternalServiceError{}).ErrorCode():             gcerrors.Internal,
	(&types.ResourceExistsException{}).ErrorCode():          gcerrors.AlreadyExists,
	(&types.PreconditionNotMetException{}).ErrorCode():      gcerrors.FailedPrecondition,
	(&types.MalformedPolicyDocumentException{}).ErrorCode(): gcerrors.FailedPrecondition,
	(&types.LimitExceededException{}).ErrorCode():           gcerrors.ResourceExhausted,
}
