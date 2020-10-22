package pkg

import (
	. "github.com/saschagrunert/demo"
	"github.com/urfave/cli/v2"
)

func Setup(ctx *cli.Context) error {
	_ = Ensure(
		"make build-malicious",
		"make server-up",
		"rm -r containers/sftp/app containers/sftp/upload/upload.tar.gz"
	)
	return nil
}
