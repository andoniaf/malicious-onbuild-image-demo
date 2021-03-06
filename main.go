package main

import (
	"github.com/andoniaf/malicious-onbuild-demo/pkg"
	. "github.com/saschagrunert/demo"
	"github.com/urfave/cli/v2"
)

func main() {
	d := New()
	d.Name = "Malicious ONBUILD demo"
	d.HideVersion = true
	d.Usage = "Malicious ONBUILD instruction demo"
	d.Authors = []*cli.Author{{
		Name: "Andoni Alonso", Email: "andonialonsof@gmail.com",
	}}
	d.Setup(pkg.Setup)
	d.Cleanup(pkg.Cleanup)

	d.Add(pkg.OnbuildBasic(), "onbuild_basic", "Explain what is ONBUILD instruction")
	d.Add(pkg.OnbuildShowEvil(), "show_malicious", "Show onbuild_malicious container")
	d.Add(pkg.OnbuildUseEvil(), "use_malicious_1", "Use onbuild_malicious container I")
	d.Add(pkg.OnbuildWhatHappened(), "what_happened", "What happened using onbuild_malicious container")
	d.Add(pkg.OnbuildUseEvil_2(), "use_malicious_2", "Use onbuild_malicious container II")
	d.Add(pkg.Conclusions(), "conclusions", "Conclusions about ONBUILD instruction")
	d.Run()
}
