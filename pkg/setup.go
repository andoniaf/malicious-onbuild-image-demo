package pkg

import (
	. "github.com/saschagrunert/demo"
	"github.com/urfave/cli/v2"
)

func Setup(ctx *cli.Context) error {
	_ = Ensure(
		"rm -r containers/sftp/app containers/sftp/upload/upload.tar.gz || true",
		"make build-malicious",
		"make server-up",
	)
	return nil
}
