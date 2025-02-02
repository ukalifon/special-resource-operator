package e2e

import (
	"context"
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/openshift-psap/special-resource-operator/test/framework"

	"github.com/kelseyhightower/envconfig"
)

func TestSRO(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Special Resource Operator e2e tests: basic")
}

var _ = ginkgo.BeforeSuite(func() {
	var config framework.Config
	err := envconfig.Process("sro", &config)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	cs := framework.NewClientSet(config)

	cl, err := framework.NewControllerRuntimeClient()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	ginkgo.By("[pre] Creating kube client set...")
	clientSet, err := GetKubeClientSet()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	ginkgo.By("[pre] Checking SRO status...")
	err = WaitSRORunning(clientSet, cs.Config.Namespace)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	ginkgo.By("[pre] Creating preamble...")
	err = CreatePreamble(context.TODO(), cl)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	ginkgo.By("[pre] Checking ClusterOperator conditions...")
	err = WaitClusterOperatorConditions(cs)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	ginkgo.By("[pre] Checking ClusterOperator related objects...")
	err = WaitClusterOperatorNamespace(cs)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
})
