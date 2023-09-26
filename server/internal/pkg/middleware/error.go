package middleware

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// reason holds the error reason.
const reason string = "UNAUTHORIZED"

var ErrWrongContext = errors.Unauthorized(reason, "Wrong context for middleware")
var ErrWrongHTTPContext = errors.Unauthorized(reason, "Wrong http context for middleware")
var ErrWrongAuthorization = errors.Unauthorized(reason, "Wrong Authorization for middleware")
