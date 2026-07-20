package gcpfirestore

import (
	"context"
	"regexp"

	vkit "cloud.google.com/go/firestore/apiv1"
	pb "cloud.google.com/go/firestore/apiv1/firestorepb"
	"github.com/google/wire"
	"gocloud.dev/docstore"
	"gocloud.dev/docstore/driver"
	"gocloud.dev/gcerrors"
	"gocloud.dev/gcp"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

func Dial(ctx context.Context, ts gcp.TokenSource) (*vkit.Client, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

var Set = wire.NewSet(
	Dial,
	wire.Struct(new(URLOpener), "Client"),
)

type collection struct {
	nameField string
	nameFunc  func(docstore.Document) string
	client    *vkit.Client
	dbPath    string
	collPath  string
	opts      *Options
}

type Options struct {
	AllowLocalFilters bool

	RevisionField string

	MaxOutstandingActionRPCs int
}

func CollectionResourceID(projectID, collPath string) string { _ = "STUB: not implemented"; return "" }

func CollectionResourceIDWithDatabase(projectID, databaseID, collPath string) string {
	_ = "STUB: not implemented"
	return ""
}

func OpenCollection(client *vkit.Client, collResourceID, nameField string, opts *Options) (*docstore.Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OpenCollectionWithNameFunc(client *vkit.Client, collResourceID string, nameFunc func(docstore.Document) string, opts *Options) (*docstore.Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var resourceIDRE = regexp.MustCompile(`^(projects/[^/]+/databases/[^/]+)/documents/.+`)

func newCollection(client *vkit.Client, collResourceID, nameField string, nameFunc func(docstore.Document) string, opts *Options) (*collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) Key(doc driver.Document) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (c *collection) RevisionField() string { _ = "STUB: not implemented"; return "" }

func (c *collection) RunActions(ctx context.Context, actions []*driver.Action, opts *driver.RunActionsOptions) driver.ActionListError {
	_ = "STUB: not implemented"
	return *new(driver.ActionListError)
}

func (c *collection) runGets(ctx context.Context, actions []*driver.Action, errs []error, opts *driver.RunActionsOptions) {
	_ = "STUB: not implemented"
	return
}

func (c *collection) batchGet(ctx context.Context, gets []*driver.Action, errs []error, opts *driver.RunActionsOptions) {
	_ = "STUB: not implemented"
	return
}

func (c *collection) newGetRequest(gets []*driver.Action) (*pb.BatchGetDocumentsRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type commitCall struct {
	writes   []*pb.Write
	actions  []*driver.Action
	newNames []string
}

func (c *collection) buildCommitCalls(actions []*driver.Action, errs []error) []*commitCall {
	_ = "STUB: not implemented"
	return nil
}

func (c *collection) buildAtomicWritesCommitCall(actions []*driver.Action, errs []error) *commitCall {
	_ = "STUB: not implemented"
	return nil
}

func (c *collection) actionToWrites(a *driver.Action) ([]*pb.Write, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (c *collection) putWrite(doc driver.Document, docName string, pc *pb.Precondition) (*pb.Write, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) deleteWrite(doc driver.Document, docName string) (*pb.Write, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) updateWrites(doc driver.Document, docName string, mods []driver.Mod) ([]*pb.Write, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newUpdateWrites(docPath string, ts *tspb.Timestamp, fields map[string]*pb.Value, paths []string, transforms []*pb.DocumentTransform_FieldTransform) ([]*pb.Write, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processMods(mods []driver.Mod) (fields map[string]*pb.Value, maskPaths []string, transforms []*pb.DocumentTransform_FieldTransform, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (c *collection) doCommitCall(ctx context.Context, call *commitCall, errs []error, opts *driver.RunActionsOptions) {
	_ = "STUB: not implemented"
	return
}

func hasFollowingTransform(writes []*pb.Write, i int) bool { _ = "STUB: not implemented"; return false }

func (c *collection) commit(ctx context.Context, ws []*pb.Write, opts *driver.RunActionsOptions) ([]*pb.WriteResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setAtFieldPath(m map[string]*pb.Value, fp []string, val *pb.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func getParentMap(m map[string]*pb.Value, fp []string, create bool) (map[string]*pb.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toServiceFieldPath(fp []string) string { _ = "STUB: not implemented"; return "" }

var unquotedFieldRE = regexp.MustCompile("^[A-Za-z_][A-Za-z_0-9]*$")

func toServiceFieldPathComponent(key string) string { _ = "STUB: not implemented"; return "" }

func (c *collection) revisionPrecondition(doc driver.Document) (*pb.Precondition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) revisionTimestamp(doc driver.Document) (*tspb.Timestamp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func preconditionFromTimestamp(ts *tspb.Timestamp) *pb.Precondition {
	_ = "STUB: not implemented"
	return nil
}

func (c *collection) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

const resourcePrefixHeader = "google-cloud-resource-prefix"

func withResourceHeader(ctx context.Context, resource string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *collection) RevisionToBytes(rev any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) BytesToRevision(b []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (c *collection) As(i any) bool { _ = "STUB: not implemented"; return false }

func (c *collection) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (c *collection) Close() error { _ = "STUB: not implemented"; return nil }
