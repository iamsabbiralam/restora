package error

import (
	"encoding/json"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BaseError struct {
	Code             int
	Message          string
	ValidationErrors map[string]string
}

// Error ...
func (r *BaseError) Error() string {
	if r == nil {
		return ""
	}

	return r.Message
}

func ToHTTPError(err error) *BaseError {
	if err == nil {
		return nil
	}

	var baseErr BaseError
	code, message := toHttpStatusMessage(err)
	if code == http.StatusUnprocessableEntity {
		baseErr.Code = http.StatusUnprocessableEntity
		baseErr.Message = message
		baseErr.ValidationErrors = validationErrorsHandle(err)
	}
	// Note: Add more error type conditions when needed

	return &baseErr
}

func toHttpStatusMessage(err error) (int, string) {
	var (
		grpcCode codes.Code
		message  string
	)

	sts := status.Convert(err)
	grpcCode = sts.Code()
	message = sts.Message()

	// In case code and message not set. Though in very rare case
	if grpcCode <= 0 {
		grpcCode = codes.Unknown
	}

	if message == "" {
		message = "Unknown error"
	}

	return getHttpStatusCodeFromGrpcCode(grpcCode), message
}

func validationErrorsHandle(err error) map[string]string {
	var vErrs map[string]string
	errStr := strings.TrimLeft(err.Error(), "rpc error: code = InvalidArgument desc = ")
	json.Unmarshal([]byte(errStr), &vErrs)
	return vErrs
}
