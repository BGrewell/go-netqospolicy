package internal

type NetworkProfile string
const (
	Profile_Domain = "Domain"
	Profile_Public = "Public"
	Profile_Private = "Private"
	Profile_All = "All"
)

type IPProtocol string
const (
	IPProto_TCP = "TCP"
	IPProto_UDP = "UDP"
	IPProto_BOTH = "BOTH"
)

type Switch string
const (
	Switch_Cluster = "-Cluster"
	Switch_Default = "-Default"
	Switch_FCOE = "-FCOE"
	Switch_iSCSI = "-iSCSI"
	Switch_LiveMigration = "-LiveMigration"
	Switch_NFS = "-NFS"
	Switch_SMB = "-SMB"
)

type NetQoSPolicy struct {
	PolicyStore string
	Name string
	NetworkProfile NetworkProfile
	Precedence uint32
	PriorityValue8021Action int8
	DSCPAction int8
	MinBandwidthWeightAction byte
	ThrottleRateActionBitsPerSecond uint64
	IPDstPrefixMatchCondition string
	UserMatchCondition string
	AppPathNameMatchCondition string
	IPProtocolMatchCondition IPProtocol
	IPSrcPrefixMatchCondition string
	IPSrcPortMatchCondition uint16
	IPSrcPortStartMatchCondition uint16
	IPSrcPortEndMatchCondition uint16
	IPDstPortMatchCondition uint16
	IPDstPortStartMatchCondition uint16
	IPDstPortEndMatchCondition uint16
	ThrottleLimit int32
	Switches []Switch
}
