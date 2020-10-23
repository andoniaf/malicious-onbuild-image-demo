package pkg

import (
	. "github.com/saschagrunert/demo"
)

func OnbuildShowEvil() *Run {
	r := NewRun(
		"What's the problem with ONBUILD then?",
		"Ok, can be `ONBUILD` bad? It looks useful to me. :)",
		"",
	)

	r.Step(S(
		"I have here a docker image that I've created and called `onbuild_malicious`...",
		"But that it's just because it's a badass name, there's nothing wrong here, look:",
	), S(
		"docker history onbuild_malicious",
	))

	r.Step(S(
		"You probably don't trust me right now, let's use some tools.",
		"Do you trust tools? I mean, never trust TOOLS, but this are other tools.",
	), S(
		`# docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -v $HOME/Library/Caches:/root/.cache/ aquasec/trivy onbuild_malicious
	bat --pager=none demo_files/aquasec_trivy_onbuild_malicious.log`,
	))

	r.Step(S(
		"Also I've uploaded it to DockerHub,",
		"they won't allow me to upload something bad, right? RIGHT? ლ(ಠ益ಠ)ლ",
		"Let's try another tool:",
	), S(
		`#docker run --rm -v /var/run/docker.sock:/var/run/docker.sock goodwithtech/dockle:v0.3.1 andoniaf/onbuild_demo
		bat --pager=none demo_files/dockle_onbuild_malicious.log`,
	))

	return r
}
