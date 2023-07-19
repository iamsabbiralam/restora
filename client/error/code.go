package error

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

func getHttpStatusCodeFromGrpcCode(code codes.Code) int {
	switch code {
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.DataLoss:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.OutOfRange:
		return http.StatusServiceUnavailable
	case codes.Aborted:
		return http.StatusNotAcceptable
	case codes.FailedPrecondition:
		return http.StatusPreconditionFailed
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.NotFound:
		return http.StatusNotFound
	case codes.DeadlineExceeded:
		return http.StatusRequestTimeout
	case codes.InvalidArgument:
		return http.StatusUnprocessableEntity
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
