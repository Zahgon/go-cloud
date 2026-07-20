package gcpfirestore

import (
	"context"

	pb "cloud.google.com/go/firestore/apiv1/firestorepb"
	"gocloud.dev/docstore/driver"
)

func (c *collection) RunGetQuery(ctx context.Context, q *driver.Query) (driver.DocumentIterator, error) {
	_ = "STUB: not implemented"
	return *new(driver.DocumentIterator), nil
}

func (c *collection) newDocIterator(ctx context.Context, q *driver.Query) (*docIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type docIterator struct {
	streamClient        pb.Firestore_RunQueryClient
	nameField, revField string
	localFilters        []driver.Filter

	cancel func()
}

func (it *docIterator) Next(ctx context.Context, doc driver.Document) error {
	_ = "STUB: not implemented"
	return nil
}

func (it *docIterator) nextResponse(ctx context.Context) (*pb.RunQueryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (it *docIterator) evaluateLocalFilters(pdoc *pb.Document) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func evaluateFilter(f driver.Filter, doc driver.Document) bool {
	_ = "STUB: not implemented"
	return false
}

func applyComparison(op string, c int) bool { _ = "STUB: not implemented"; return false }

func (it *docIterator) Stop() { _ = "STUB: not implemented"; return }

func (it *docIterator) As(i any) bool { _ = "STUB: not implemented"; return false }

func (c *collection) queryToProto(q *driver.Query) (*pb.StructuredQuery, []driver.Filter, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func splitFilters(fs []driver.Filter) (sendToFirestore, evaluateLocally []driver.Filter) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) filterToProto(f driver.Filter) (*pb.StructuredQuery_Filter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unaryOpFor(value any) (pb.StructuredQuery_UnaryFilter_Operator, bool) {
	_ = "STUB: not implemented"
	return *new(pb.StructuredQuery_UnaryFilter_Operator), false
}

func isNaN(x any) bool { _ = "STUB: not implemented"; return false }

func fieldRef(fp []string) *pb.StructuredQuery_FieldReference {
	_ = "STUB: not implemented"
	return nil
}

func newFieldFilter(fp []string, op string, val *pb.Value) (*pb.StructuredQuery_Filter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) QueryPlan(q *driver.Query) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
