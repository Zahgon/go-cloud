package drivertest

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"gocloud.dev/docstore"
	"gocloud.dev/docstore/driver"
	"gocloud.dev/gcerrors"
)

type ByteArray [2]byte

type CollectionKind int

const (
	SingleKey CollectionKind = iota

	TwoKey

	AltRev

	NoRev
)

type AtomicWritesSupportKind int

const (
	NoSupport AtomicWritesSupportKind = iota

	SinglePartition

	MultiplePartitions
)

type Harness interface {
	MakeCollection(context.Context, CollectionKind) (driver.Collection, error)

	BeforeDoTypes() []any

	BeforeQueryTypes() []any

	RevisionsEqual(rev1, rev2 any) bool

	AtomicWritesKind() AtomicWritesSupportKind

	Close()
}

type HarnessMaker func(ctx context.Context, t *testing.T) (Harness, error)

type UnsupportedType int

const (
	Uint UnsupportedType = iota

	Arrays

	NanosecondTimes

	BinarySet
)

type CodecTester interface {
	UnsupportedTypes() []UnsupportedType
	NativeEncode(any) (any, error)
	NativeDecode(value, dest any) error
	DocstoreEncode(any) (any, error)
	DocstoreDecode(value, dest any) error
}

type AsTest interface {
	Name() string

	CollectionCheck(coll *docstore.Collection) error

	QueryCheck(it *docstore.DocumentIterator) error

	ErrorCheck(c *docstore.Collection, err error) error
}

type verifyAsFailsOnNil struct{}

func (verifyAsFailsOnNil) Name() string { _ = "STUB: not implemented"; return "" }

func (verifyAsFailsOnNil) CollectionCheck(coll *docstore.Collection) error {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) QueryCheck(it *docstore.DocumentIterator) error {
	_ = "STUB: not implemented"
	return nil
}

func (v verifyAsFailsOnNil) ErrorCheck(c *docstore.Collection, err error) (ret error) {
	_ = "STUB: not implemented"
	return nil
}

func RunConformanceTests(t *testing.T, newHarness HarnessMaker, ct CodecTester, asTests []AsTest) {
	_ = "STUB: not implemented"
	return
}

func withCollection(t *testing.T, newHarness HarnessMaker, kind CollectionKind, f func(*testing.T, Harness, *docstore.Collection)) {
	_ = "STUB: not implemented"
	return
}

func withRevCollections(t *testing.T, newHarness HarnessMaker, f func(*testing.T, *docstore.Collection, string)) {
	_ = "STUB: not implemented"
	return
}

func withColl(t *testing.T, h Harness, kind CollectionKind, f func(*testing.T, Harness, *docstore.Collection)) {
	_ = "STUB: not implemented"
	return
}

const KeyField = "name"

const AlternateRevisionField = "Etag"

type docmap = map[string]any

func newDoc(doc any) any { _ = "STUB: not implemented"; return *new(any) }

func revision(doc any, revField string) any { _ = "STUB: not implemented"; return *new(any) }

func setRevision(doc, rev any, revField string) { _ = "STUB: not implemented"; return }

type docstruct struct {
	Name             any `docstore:"name"`
	DocstoreRevision any
	Etag             any

	I  int
	U  uint
	F  float64
	St string
	B  bool
	M  map[string]any
}

func nonexistentDoc() docmap { _ = "STUB: not implemented"; return *new(docmap) }

func testCreate(t *testing.T, coll *docstore.Collection, revField string) {
	_ = "STUB: not implemented"
	return
}

func testPut(t *testing.T, coll *docstore.Collection, revField string) {
	_ = "STUB: not implemented"
	return
}

func testReplace(t *testing.T, coll *docstore.Collection, revField string) {
	_ = "STUB: not implemented"
	return
}

func checkNoRevisionField(t *testing.T, doc any, revField string) {
	_ = "STUB: not implemented"
	return
}

func checkHasRevisionField(t *testing.T, doc any, revField string) {
	_ = "STUB: not implemented"
	return
}

func testGet(t *testing.T, coll *docstore.Collection, revField string) {
	_ = "STUB: not implemented"
	return
}

func testDelete(t *testing.T, coll *docstore.Collection, revField string) {
	_ = "STUB: not implemented"
	return
}

func testUpdate(t *testing.T, coll *docstore.Collection, revField string) {
	_ = "STUB: not implemented"
	return
}

func testRevisionField(t *testing.T, coll *docstore.Collection, revField string, write func(any) error) {
	_ = "STUB: not implemented"
	return
}

func testSerializeRevision(t *testing.T, h Harness, coll *docstore.Collection) {
	_ = "STUB: not implemented"
	return
}

func testData(t *testing.T, _ Harness, coll *docstore.Collection) {
	_ = "STUB: not implemented"
	return
}

var (
	milliTime = time.Date(2019, time.March, 27, 0, 0, 0, 5*1e6, time.UTC)

	nanoTime = time.Date(2019, time.March, 27, 0, 0, 0, 5*1e6+7, time.UTC)
)

func testTypeDrivenDecode(t *testing.T, ct CodecTester) { _ = "STUB: not implemented"; return }

func testBlindDecode(t *testing.T, ct CodecTester) { _ = "STUB: not implemented"; return }

func testBlindDecode1(t *testing.T, encode func(any) (any, error), decode func(_, _ any) error) {
	_ = "STUB: not implemented"
	return
}

type docstoreRoundTrip struct {
	N  *int
	I  int
	U  uint
	F  float64
	St string
	B  bool
	By []byte
	L  []int
	A  [2]int
	A2 [2]int8
	At ByteArray
	Uu uuid.UUID
	M  map[string]bool
	P  *string
	T  time.Time
}

type nativeMinimal struct {
	N  *int
	I  int
	F  float64
	St string
	B  bool
	By []byte
	L  []int
	A  [2]int
	A2 [2]int8
	At ByteArray
	M  map[string]bool
	P  *string
	T  time.Time
	LF []float64
	LS []string
}

func testProto(t *testing.T, _ Harness, coll *docstore.Collection) {
	_ = "STUB: not implemented"
	return
}

type HighScore struct {
	Game             string
	Player           string
	Score            int
	Time             time.Time
	WithGlitch       bool
	DocstoreRevision any
}

func newHighScore() any { _ = "STUB: not implemented"; return *new(any) }

func HighScoreKey(doc docstore.Document) any { _ = "STUB: not implemented"; return *new(any) }

func (h *HighScore) key() string { _ = "STUB: not implemented"; return "" }

func barConcat(a, b any) string { _ = "STUB: not implemented"; return "" }

func highScoreLess(h1, h2 *HighScore) bool { _ = "STUB: not implemented"; return false }

func (h *HighScore) String() string { _ = "STUB: not implemented"; return "" }

func date(month, day int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

const (
	game1 = "Praise All Monsters"
	game2 = "Zombie DMV"
	game3 = "Days Gone"
)

var highScores = []*HighScore{
	{game1, "pat", 49, date(3, 13), false, nil},
	{game1, "mel", 60, date(4, 10), false, nil},
	{game1, "andy", 81, date(2, 1), false, nil},
	{game1, "fran", 33, date(3, 19), false, nil},
	{game2, "pat", 120, date(4, 1), true, nil},
	{game2, "billie", 111, date(4, 10), false, nil},
	{game2, "mel", 190, date(4, 18), true, nil},
	{game2, "fran", 33, date(3, 20), false, nil},
}

func addHighScores(t *testing.T, coll *docstore.Collection) { _ = "STUB: not implemented"; return }

func testGetQueryKeyField(t *testing.T, coll *docstore.Collection, revField string) {
	_ = "STUB: not implemented"
	return
}

func sortByKeyField(d1, d2 docmap) bool { _ = "STUB: not implemented"; return false }

func testActionsWithCompositeID(t *testing.T, _ Harness, coll *docstore.Collection) {
	_ = "STUB: not implemented"
	return
}

func testGetQuery(t *testing.T, _ Harness, coll *docstore.Collection) {
	_ = "STUB: not implemented"
	return
}

func filterHighScores(hs []*HighScore, f func(*HighScore) bool) []*HighScore {
	_ = "STUB: not implemented"
	return nil
}

func ClearCollection(fataler interface {
	Helper()
	Fatalf(string, ...any)
}, coll *docstore.Collection,
) {
	_ = "STUB: not implemented"
	return
}

func forEach(ctx context.Context, iter *docstore.DocumentIterator, create func() any, handle func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

func mustCollect(ctx context.Context, t *testing.T, iter *docstore.DocumentIterator) []docmap {
	_ = "STUB: not implemented"
	return nil
}

func mustCollectHighScores(ctx context.Context, t *testing.T, iter *docstore.DocumentIterator) []*HighScore {
	_ = "STUB: not implemented"
	return nil
}

func collectHighScores(ctx context.Context, iter *docstore.DocumentIterator) ([]*HighScore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func testMultipleActions(t *testing.T, coll *docstore.Collection, revField string) {
	_ = "STUB: not implemented"
	return
}

func testAtomicWrites(t *testing.T, coll *docstore.Collection, revField string) {
	_ = "STUB: not implemented"
	return
}

func testAtomicWritesFail(t *testing.T, coll *docstore.Collection, revField string) {
	_ = "STUB: not implemented"
	return
}

func testAtomicWritesSinglePartition(t *testing.T, h Harness, coll *docstore.Collection) {
	_ = "STUB: not implemented"
	return
}

func testAtomicWritesFailSinglePartition(t *testing.T, h Harness, coll *docstore.Collection) {
	_ = "STUB: not implemented"
	return
}

func testActionsOnStructNoRev(t *testing.T, _ Harness, coll *docstore.Collection) {
	_ = "STUB: not implemented"
	return
}

func testExampleInDoc(t *testing.T, _ Harness, coll *docstore.Collection) {
	_ = "STUB: not implemented"
	return
}

func testBeforeDo(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testBeforeQuery(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testAs(t *testing.T, coll *docstore.Collection, st AsTest) { _ = "STUB: not implemented"; return }

func clone(m docmap) docmap { _ = "STUB: not implemented"; return *new(docmap) }

func cmpDiff(a, b any, opts ...cmp.Option) string { _ = "STUB: not implemented"; return "" }

func checkCode(t *testing.T, err error, code gcerrors.ErrorCode) { _ = "STUB: not implemented"; return }
