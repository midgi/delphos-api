package main_test

import (
	"fmt"
	"net/http"
	"time"

	"github.com/midgi/delphos-api/cmd/delphos/testrunner"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/ginkgomon"

	"testing"
)

var (
	delphosBinPath string
	delphosRunner  *ginkgomon.Runner
	delphosProcess ifrit.Process
	delphosArgs    testrunner.Args
	delphosPort    int

	httpclient *http.Client
)

func TestDelphos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Delphos Suite")
}

var _ = SynchronizedBeforeSuite(
	func() []byte {
		delphosConfig, err := gexec.Build("github.com/midgi/delphos-api/cmd/delphos", "-race")
		Expect(err).NotTo(HaveOccurred())
		return []byte(delphosConfig)
	},
	func(delphosConfig []byte) {
		delphosBinPath = string(delphosConfig)
		SetDefaultEventuallyTimeout(15 * time.Second)
		delphosPort = 8080 + GinkgoParallelNode()
		delphosArgs.Address = fmt.Sprintf("127.0.0.1:%d", delphosPort)
		delphosRunner = testrunner.New(delphosBinPath, delphosArgs)
		delphosProcess = ginkgomon.Invoke(delphosRunner)
		httpclient = &http.Client{}
	},
)

var _ = SynchronizedAfterSuite(func() {
	ginkgomon.Kill(delphosProcess)
}, func() {
	gexec.CleanupBuildArtifacts()
})
