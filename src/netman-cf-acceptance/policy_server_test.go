package acceptance_test

import (
	"fmt"
	"os/exec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("policy server tests", func() {
	It("makes the policy server available at an external route", func() {
		cmd := exec.Command("curl", "-v", fmt.Sprintf("http://%s/networking", config.ApiEndpoint))

		sess, err := gexec.Start(cmd, nil, nil)
		Expect(err).NotTo(HaveOccurred())
		Eventually(sess.Wait(Timeout_Short)).Should(gexec.Exit(0))

		curlOutput := sess.Out.Contents()
		Expect(curlOutput).To(ContainSubstring("Network policy server, up for"))
	})

	Context("When the user has network.admin scope", func() {
		BeforeEach(func() {
			Auth(testConfig.TestUser, testConfig.TestUserPassword)
		})
		AfterEach(func() {
			AuthAsAdmin()
		})

		It("allows access to whoami", func() {
			curlSession := cf.Cf("curl", "/networking/v0/external/whoami").Wait(Timeout_Push)

			Expect(curlSession.Wait(Timeout_Push)).To(gexec.Exit(0))
			whoamiOut := string(curlSession.Out.Contents())
			Expect(whoamiOut).To(MatchJSON(fmt.Sprintf(`{"user_name":%q}`, testConfig.TestUser)))
		})
	})

})