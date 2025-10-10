package errcodes

import (
	"github.com/jzero-io/jzero-admin/server/server/errcodes/auth"
	"github.com/jzero-io/jzero-admin/server/server/errcodes/manage"
)

func Register() {
	manage.RegisterManage()
	auth.RegisterAuth()
}
