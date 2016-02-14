package testrunner

import (
	"os/exec"

	"github.com/tedsuo/ifrit/ginkgomon"
)

type Args struct {
	Address string
}

func (args Args) ArgSlice() []string {
	arguments := []string{
		"-listenAddress", args.Address,
	}
	return arguments
}

func New(binPath string, args Args) *ginkgomon.Runner {
	return ginkgomon.New(ginkgomon.Config{
		Name:       "delphos",
		Command:    exec.Command(binPath, args.ArgSlice()...),
		StartCheck: "delphos.started",
	})
}
