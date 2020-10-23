package pkg

import (
	. "github.com/saschagrunert/demo"
)

func OnbuildWhatHappened() *Run {
	r := NewRun(
		"What did ONBUILD image do?",
		"Okey, nice joke, WTF did you do.",
		"",
	)

	r.Step(S(
		"Let's see the `onbuild_malicious` Dockerfile",
	), S(
		"bat containers/malicious/Dockerfile.bad",
	))

	r.Step(S(
		"But we didn't see that in the docker history!",
	), S(
		"docker history onbuild_malicious",
	))

	r.Step(S(
		"Nope, we dont'. ¯\\_(ツ)_/¯",
		"",
		"And the tools?! We trusted the tools...",
		"Yeah, but you know, this isn't a bug or an exploit, it's a feature.",
	), S(
		"curl e.xec.sh/mind_blown",
	))

	r.Step(S(
		"But, obviously there are some ways to check it before using the container.",
		"The easiest one is check the Dockerfile, but sometimes it isn't available.",
		"We could use `docker inspect`:",
	), S(
		"docker inspect andoniaf/onbuild_demo | jq '.[].Config.OnBuild'",
	))

	r.Step(nil, S(
		"docker inspect victoria:v2-stable | jq '.[].Config.OnBuild'",
	))

	r.Step(nil, S(
		"docker inspect ruby:2.3-onbuild | jq '.[].Config.OnBuild'",
	))

	return r
}
