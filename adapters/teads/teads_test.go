package teads

import (
	"github.com/prebid/prebid-server/adapters/adapterstest"
	"github.com/prebid/prebid-server/config"
	"github.com/prebid/prebid-server/openrtb_ext"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonSamples(t *testing.T) {
	bidder, buildErr := Builder(openrtb_ext.BidderTeads, config.Adapter{
		Endpoint: "https://a.teads.tv/prebid-server/bid-request"}, config.Server{ExternalUrl: "https://a.teads.tv/prebid-server/bid-request", GvlID: 1, DataCenter: "2"})

	if buildErr != nil {
		t.Fatalf("Builder returned unexpected error %v", buildErr)
	}

	adapterstest.RunJSONBidderTest(t, "teadstest", bidder)
}

func TestEndpointTemplateMalformed(t *testing.T) {
	_, buildErr := Builder(openrtb_ext.BidderTeads, config.Adapter{
		Endpoint: "{{Malformed}}"}, config.Server{ExternalUrl: "https://a.teads.tv/prebid-server/bid-request", GvlID: 1, DataCenter: "2"})

	assert.Error(t, buildErr)
}
