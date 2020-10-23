package pkg

import (
	. "github.com/saschagrunert/demo"
)

func OnbuildUseEvil_2() *Run {
	r := NewRun(
		"Let's use our badass ONBUILD image II",
		"Now we know it's not safe, but... ಥ_ಥ",
	)

	r.Step(S(
		"Another innocent docker image that I want to build,",
		"this time we're going to use the same image but from DockerHub:",
	), S(
		"ls -la containers/inno-3",
		"&& echo &&",
		"bat containers/inno-3/Dockerfile.inno",
	))

	r.Step(S(
		"Again a simple `Dockerfile`",
		"",
		"This time we're going to add a `.dockerignore` file avoid Docker from copy our `.env` file.",
	), S(
		"bat containers/inno-3/.dockerignore",
	))

	r.Step(nil, S(
		"make build-innocent-3",
	))

	r.Step(S(
		"Let's see again what's inside our SFTP...",
	), S(
		`cd containers/sftp && \
		ls -l upload/ && \
		tar -xvzf upload/upload.tar.gz && \
		ls -la app
		`,
	))

	r.Step(S(
		"Oh no...my secret token, again (╯°□°）╯︵ ┻━┻",
		"",
		"Well, at least our `.env` it's not there...¯\\_(ツ)_/¯",
	), nil)

	return r
}
