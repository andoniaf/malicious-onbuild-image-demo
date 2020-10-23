package pkg

import (
	. "github.com/saschagrunert/demo"
)

func OnbuildBasic() *Run {
	r := NewRun(
		"ONBUILD",
		"What is `ONBUILD` instruction?",
		"",
	)

	r.Step(S(
		"The ONBUILD instruction adds to the image a trigger instruction to be executed at a later time,",
		"when the image is used as the base for another build.",
		"",
		"It will be executed as if it had been inserted immediately after the FROM instruction in the downstream `Dockerfile`.",
	), nil)

	r.Step(S(
		"ONBUILD is useful for images that are going to be built FROM a given image.",
		"For example, you would use ONBUILD for a language stack image that builds arbitrary user software written",
		"in that language within the Dockerfile, as you can see in Ruby’s ONBUILD variants.",
	), S(
		"bat demo_files/ruby2.4_jessie_onbuild.dockerfile",
	))

	return r
}
