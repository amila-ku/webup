package aws

type HostedZoneId string

const (
	USEast1HostedZoneId   HostedZoneId = "Z3AQBSTGFYJSTF"
	USEast2HostedZoneId   HostedZoneId = "Z2O1EMRO9K5GLX"
	USWest1HostedZoneId   HostedZoneId = "Z2F56UZL2M1ACD"
	USWest2HostedZoneId   HostedZoneId = "Z3BJ6K6RIION7M"
	EUCentralHostedZoneId HostedZoneId = "Z21DNDUVLTQW6Q"
	EUWestHostedZoneId    HostedZoneId = "Z1BKCTXD74EZPE"
)

var zonemap = map[string]HostedZoneId{
	"eu-central-1": EUCentralHostedZoneId,
}
