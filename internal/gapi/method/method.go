package method

import (
	"github.com/iput-kernel/foundation-account/internal/gapi/service"
)

type Method struct {
	*service.Server
}

func NewMethod(server *service.Server) *Method {
	return &Method{Server: server}
}
