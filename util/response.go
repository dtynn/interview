package util

import (
	"encoding/json"
	"net/http"

	"github.com/dtynn/interview/proto"
)

func Resp(rw http.ResponseWriter, data interface{}) {
	enc := json.NewEncoder(rw)
	if err := enc.Encode(data); err != nil {
		enc.Encode(proto.Response{
			Code:    1,
			Message: err.Error(),
		})
	}
}
