package aws

// HostedZoneID represents aws s3 hostedzoneids
type HostedZoneID string

const (
	//USEast1HostedZoneID is the s3 web endpoint hosted zone id 
	USEast1HostedZoneID   HostedZoneID = "Z3AQBSTGFYJSTF"
	//USEast2HostedZoneID is the s3 web endpoint hosted zone id 
	USEast2HostedZoneID   HostedZoneID = "Z2O1EMRO9K5GLX"
	//USWest1HostedZoneID  is the s3 web endpoint hosted zone id 
	USWest1HostedZoneID   HostedZoneID = "Z2F56UZL2M1ACD"
	//USWest2HostedZoneID  is the s3 web endpoint hosted zone id 
	USWest2HostedZoneID   HostedZoneID = "Z3BJ6K6RIION7M"
	//EUCentralHostedZoneID is the s3 web endpoint hosted zone id 
	EUCentralHostedZoneID HostedZoneID = "Z21DNDUVLTQW6Q"
	//EUWestHostedZoneID  is the s3 web endpoint hosted zone id 
	EUWestHostedZoneID    HostedZoneID = "Z1BKCTXD74EZPE"
)

var zonemap = map[string]HostedZoneID{
	"eu-central-1": EUCentralHostedZoneID,
}
