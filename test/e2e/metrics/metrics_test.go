package metrics

import (
	"testing"

	"github.com/kyverno/kyverno/test/e2e"
	. "github.com/onsi/gomega"
)

func Test_MetricsServerAvailability(t *testing.T) {
	RegisterTestingT(t)
	requestObj := e2e.APIRequest{
		URL:  "http://localhost:8000/metrics",
		Type: "GET",
	}
	response, err := e2e.CallAPI(requestObj)
	Expect(err).NotTo(HaveOccurred())
	Expect(response.StatusCode).To(Equal(200))
}
