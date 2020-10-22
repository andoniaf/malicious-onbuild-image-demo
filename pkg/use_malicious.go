package pkg

import (
	. "github.com/saschagrunert/demo"
)

func OnbuildUseEvil() *Run {
	r := NewRun(
		"Let's use our badass ONBUILD image",
		"Now that we know it's safe nothing could go wrong. (ಠ_ಠ)",
	)

	r.Step(S(
		"I have here a new innocent docker image that I want to build:",
	), S(
		"ls -l containers/inno",
		"&& echo &&",
		"bat containers/inno/Dockerfile.inno",
	))

	r.Step(S(
		"It's a simple `Dockerfile` right?",
		"",
		"There's a secret there, but I don't want to lose that token so...",
		"don't tell the @police pls \\(^-^)/",
		"",
		"Let's build our image:",
	), S(
		"make build-innocent",
	))

	r.Step(S(
		"Do you see something weird? ლ(ಠ益ಠლ)╯",
	), S(
		"docker history onbuild_innocent",
	))

	r.Step(S(
		"WTF is that `curl` and that `.tar.gz`?!?!?!",
		"",
		"Well, let me introduce you my friend `SFTP`,",
		"for demo purpose let's assume it's in a Russian server instead of local.",
	), S(
		"docker ps | grep upload_server",
	))

	r.Step(S(
		"Let's see what's inside...",
	), S(
		`cd containers/sftp && \
		ls -l upload/ && \
		tar -xvzf upload/upload.tar.gz && \
		ls -la app
		`,
	))

	r.Step(S(
		"Oh no...my secret token (╥﹏╥)",
	), nil)

	return r
}
