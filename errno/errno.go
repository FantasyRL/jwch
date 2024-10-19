/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package errno

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrorCode int64
	ErrorMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("error code: %d, error msg: %s", e.ErrorCode, e.ErrorMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrorMsg = msg
	return e
}

func (e ErrNo) WithErr(err error) ErrNo {
	e.ErrorMsg = err.Error()
	return e
}

// ConvertErr convert error to ErrNo
// in Default user ServiceErrorCode
func ConvertErr(err error) ErrNo {
	errno := ErrNo{}
	if errors.As(err, &errno) {
		return errno
	}

	s := ServiceError
	s.ErrorMsg = err.Error()
	return s
}
