package pkg

import (
	. "github.com/saschagrunert/demo"
)

func OnbuildBasic() *Run {
	r := NewRun(
		"ONBUILD instruction",
		"What is `ONBUILD` instruction?",
	)

	r.Step(S(
		"Multi",
		"line",
		"description",
	), nil)

	return r
}
