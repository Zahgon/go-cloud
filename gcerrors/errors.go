package gcerrors

import (
	"gocloud.dev/internal/gcerr"
)

type ErrorCode = gcerr.ErrorCode

const (
	OK ErrorCode = gcerr.OK

	Unknown ErrorCode = gcerr.Unknown

	NotFound ErrorCode = gcerr.NotFound

	AlreadyExists ErrorCode = gcerr.AlreadyExists

	InvalidArgument ErrorCode = gcerr.InvalidArgument

	Internal ErrorCode = gcerr.Internal

	Unimplemented ErrorCode = gcerr.Unimplemented

	FailedPrecondition ErrorCode = gcerr.FailedPrecondition

	PermissionDenied ErrorCode = gcerr.PermissionDenied

	ResourceExhausted ErrorCode = gcerr.ResourceExhausted

	Canceled ErrorCode = gcerr.Canceled

	DeadlineExceeded ErrorCode = gcerr.DeadlineExceeded
)

var (
	ErrUnknown = gcerr.ErrUnknown

	ErrNotFound = gcerr.ErrNotFound

	ErrAlreadyExists = gcerr.ErrAlreadyExists

	ErrInvalidArgument = gcerr.ErrInvalidArgument

	ErrInternal = gcerr.ErrInternal

	ErrUnimplemented = gcerr.ErrUnimplemented

	ErrFailedPrecondition = gcerr.ErrFailedPrecondition

	ErrPermissionDenied = gcerr.ErrPermissionDenied

	ErrResourceExhausted = gcerr.ErrResourceExhausted

	ErrCanceled = gcerr.ErrCanceled

	ErrDeadlineExceeded = gcerr.ErrDeadlineExceeded
)

func Code(err error) ErrorCode { _ = "STUB: not implemented"; return *new(ErrorCode) }
