package pkg

import (
	. "github.com/saschagrunert/demo"
)

func Conclusions() *Run {
	r := NewRun(
		"Conclusions",
		"Well, now that you're afraid, what can we do?",
	)

	r.Step(S(
		"First of all, I now that you're not afraid 【ツ】",
		"",
		"Docker images are like any other 3rd party software,",
		"you shouldn't trust it \"by default\", or at least",
		"TRUST but VERIFY.",
	), nil)

	r.Step(S(
		"If you're going to handle sensitive data you should review the source,",
		"use images from trusted sources or even craft you own base images.",
	), nil)

	r.Step(S(
		"Also as you saw in the demo, using `COPY . .` in our Dockerfiles it's not a best-practice,",
		"and we should add a `.dockerignore` file to avoid copying files that",
		"we don't want inside the container (this will also help us reduce image size).",
	), nil)

	return r
}
