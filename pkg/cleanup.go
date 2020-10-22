package pkg

import (
	. "github.com/saschagrunert/demo"
	"github.com/urfave/cli/v2"
)

func Cleanup(ctx *cli.Context) error {
	_ = Ensure(
		"docker rmi onbuild_innocent onbuild_innocent_2 onbuild_innocent_3 onbuild_malicious",
		"docker stop upload_server",
	)
	return nil
}
