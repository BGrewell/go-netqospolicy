package netqospolicy

import (
	"encoding/json"
	"fmt"
	"github.com/BGrewell/go-conversions"
	"github.com/BGrewell/go-execute"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var (
	DEBUG = false
)

type Action string

const (
	Action_Add = "add"
	Action_Set = "set"
	Action_Del = "del"
)

type NetworkProfile string

const (
	Profile_Domain  = "Domain"
	Profile_Public  = "Public"
	Profile_Private = "Private"
	Profile_All     = "All"
)

type IPProtocol string

const (
	IPProto_TCP  = "TCP"
	IPProto_UDP  = "UDP"
	IPProto_BOTH = "BOTH"
)

type Template string

const (
	Template_Cluster       = "Cluster"
	Template_Default       = "Default"
	Template_FCOE          = "FCOE"
	Template_iSCSI         = "iSCSI"
	Template_LiveMigration = "LiveMigration"
	Template_NFS           = "NFS"
	Template_SMB           = "SMB"
)

const (
	cmdShell      = "powershell.exe"
	addCommandlet = "New-NetQosPolicy"
	getCommandlet = "Get-NetQosPolicy"
	delCommandlet = "Remove-NetQosPolicy"
	setCommandlet = "Set-NetQosPolicy"
	falseError    = "Note that policies containing minimum bandwidth specifications"
)

type NetQoSPolicy struct {
	PolicyStore                     string         `json:"policy_store,omitempty"`
	Name                            string         `json:"name,omitempty"`                                 //tested parse
	Owner                           string         `json:"owner,omitempty"`                                //tested parse
	NetworkProfile                  NetworkProfile `json:"network_profile,omitempty"`                      //tested parse
	Precedence                      *uint32        `json:"precedence,omitempty"`                           //tested parse
	PriorityValue8021Action         *int8          `json:"priority_value_8021_action,omitempty"`           //tested parse
	DSCPAction                      *int8          `json:"dscp_action,omitempty"`                          //tested parse
	MinBandwidthWeightAction        *byte          `json:"min_bandwidth_weight_action,omitempty"`          //tested parse
	ThrottleRateActionBitsPerSecond *uint64        `json:"throttle_rate_action_bits_per_second,omitempty"` //tested parse
	UserMatchCondition              string         `json:"user_match_condition,omitempty"`                 //tested parse
	AppPathNameMatchCondition       string         `json:"app_path_name_match_condition,omitempty"`        //tested parse
	IPProtocolMatchCondition        IPProtocol     `json:"ip_protocol_match_condition,omitempty"`          //tested parse
	IPSrcPrefixMatchCondition       string         `json:"ip_src_prefix_match_condition,omitempty"`        //tested parse
	IPSrcPortMatchCondition         *uint16        `json:"ip_src_port_match_condition,omitempty"`          //tested parse
	IPSrcPortStartMatchCondition    *uint16        `json:"ip_src_port_start_match_condition,omitempty"`    //tested parse
	IPSrcPortEndMatchCondition      *uint16        `json:"ip_src_port_end_match_condition,omitempty"`      //tested parse
	IPDstPrefixMatchCondition       string         `json:"ip_dst_prefix_match_condition,omitempty"`        //tested parse
	IPDstPortMatchCondition         *uint16        `json:"ip_dst_port_match_condition,omitempty"`          //tested parse
	IPDstPortStartMatchCondition    *uint16        `json:"ip_dst_port_start_match_condition,omitempty"`    //tested parse
	IPDstPortEndMatchCondition      *uint16        `json:"ip_dst_port_end_match_condition,omitempty"`      //tested parse
	ThrottleLimit                   *int32         `json:"throttle_limit,omitempty"`                       //tested parse
	TemplateMatchCondition          Template       `json:"template_match_condition,omitempty"`             //tested parse
	UriMatchCondition               string         `json:"uri_match_condition,omitempty"`
	UriRecursiveMatchCondition      bool           `json:"uri_recursive_match_condition,omitempty"`
	Persistent                      bool           `json:"persistent,omitempty"`
	FmtNoIndent                     bool           `json:"fmt_no_indent,omitempty"`
}

func (p NetQoSPolicy) String() string {
	var jbytes []byte
	var err error
	if p.FmtNoIndent {
		jbytes, err = json.Marshal(p)
	} else {
		jbytes, err = json.MarshalIndent(p, "", "  ")
	}
	if err != nil {
		return "error converting to json"
	}
	return string(jbytes)
}

func (p *NetQoSPolicy) Create() (err error) {
	cmd, err := p.getCommand(Action_Add)
	if err != nil {
		return err
	}
	err = executeCommand(fmt.Sprintf("%s %s", addCommandlet, cmd))
	return err
}

func (p *NetQoSPolicy) Update() (err error) {
	// Because of the rules MS has for updating it at least initially appears to be better to just remove and create
	// the policy vs. using Set-NetQosPolicy which doesn't like when many fields are changed.
	err = p.Remove()
	if err != nil {
		return err
	}
	return p.Create()
	//cmd, err := p.getCommand(Action_Set)
	//if err != nil {
	//	return err
	//}
	//err = executeCommand(fmt.Sprintf("%s %s", setCommandlet, cmd))
	//return err
}

func (p *NetQoSPolicy) Remove() (err error) {
	cmd, err := p.getCommand(Action_Del)
	if err != nil {
		return err
	}
	err = executeCommand(fmt.Sprintf("%s %s", delCommandlet, cmd))
	return err
}

func (p *NetQoSPolicy) IsApplied() (applied bool, err error) {
	policy, err := Get(p.Name)
	if err != nil {
		return false, err
	}
	p1 := p.String()
	p2 := policy.String()
	return strings.Compare(p1, p2) == 0, nil
}

func (p *NetQoSPolicy) setDefaultValues() {
	if p.PolicyStore == "" && p.Persistent {
		p.PolicyStore = "ActiveStore"
	}
}

func (p NetQoSPolicy) getCommand(action string) (command string, err error) {
	cmd := make([]string, 0)
	// Sanity checks
	if p.Name == "" {
		return "", fmt.Errorf("name is a required field")
	}
	if p.IPDstPortMatchCondition != nil && (p.IPDstPortStartMatchCondition != nil || p.IPDstPortEndMatchCondition != nil) {
		return "", fmt.Errorf("dst port range cannot be used in conjunction with dst port match")
	}
	if p.IPSrcPortMatchCondition != nil && (p.IPSrcPortStartMatchCondition != nil || p.IPSrcPortEndMatchCondition != nil) {
		return "", fmt.Errorf("src port range cannot be used in conjunction with src port match")
	}
	cmd = append(cmd, fmt.Sprintf("-Name \"%s\"", p.Name))
	if p.PolicyStore != "" {
		cmd = append(cmd, fmt.Sprintf("-PolicyStore %s", p.PolicyStore))
	}
	// delete command only allows the name and policy store
	if action == Action_Del {
		cmd = append(cmd, "-Confirm:$False")
		return strings.Join(cmd, " "), nil
	}
	if p.NetworkProfile != "" {
		cmd = append(cmd, fmt.Sprintf("-NetworkProfile %s", string(p.NetworkProfile)))
	}
	if p.Precedence != nil {
		cmd = append(cmd, fmt.Sprintf("-Precedence %d", *p.Precedence))
	}
	if p.PriorityValue8021Action != nil {
		cmd = append(cmd, fmt.Sprintf("-PriorityValue8021Action %d", *p.PriorityValue8021Action))
	}
	if p.DSCPAction != nil {
		cmd = append(cmd, fmt.Sprintf("-DSCPAction %d", *p.DSCPAction))
	}
	if p.MinBandwidthWeightAction != nil {
		cmd = append(cmd, fmt.Sprintf("-MinBandwidthWeightAction %d", *p.MinBandwidthWeightAction))
	}
	if p.ThrottleRateActionBitsPerSecond != nil {
		cmd = append(cmd, fmt.Sprintf("-ThrottleRateActionBitsPerSecond %d", *p.ThrottleRateActionBitsPerSecond))
	}
	if p.UserMatchCondition != "" {
		cmd = append(cmd, fmt.Sprintf("-UserMatchCondition %d", p.UserMatchCondition))
	}
	if p.AppPathNameMatchCondition != "" {
		cmd = append(cmd, fmt.Sprintf("-AppPathNameMatchCondition %s", p.AppPathNameMatchCondition))
	}
	if p.IPProtocolMatchCondition != "" {
		cmd = append(cmd, fmt.Sprintf("-IPProtocolMatchCondition %s", string(p.IPProtocolMatchCondition)))
	}
	if p.IPSrcPrefixMatchCondition != "" {
		cmd = append(cmd, fmt.Sprintf("-IPSrcPrefixMatchCondition %s", p.IPSrcPrefixMatchCondition))
	}
	if p.IPSrcPortMatchCondition != nil {
		cmd = append(cmd, fmt.Sprintf("-IpSrcPortMatchCondition %d", *p.IPSrcPortMatchCondition))
	}
	if p.IPSrcPortStartMatchCondition != nil {
		cmd = append(cmd, fmt.Sprintf("-IPSrcPortStartMatchCondition %d", *p.IPSrcPortStartMatchCondition))
	}
	if p.IPSrcPortEndMatchCondition != nil {
		cmd = append(cmd, fmt.Sprintf("-IPSrcPortEndMatchCondition %d", *p.IPSrcPortEndMatchCondition))
	}
	if p.IPDstPrefixMatchCondition != "" {
		cmd = append(cmd, fmt.Sprintf("-IPDstPrefixMatchCondition %s", p.IPDstPrefixMatchCondition))
	}
	if p.IPDstPortMatchCondition != nil {
		cmd = append(cmd, fmt.Sprintf("-IPDstPortMatchCondition %d", *p.IPDstPortMatchCondition))
	}
	if p.IPDstPortStartMatchCondition != nil {
		cmd = append(cmd, fmt.Sprintf("-IPDstPortStartMatchCondition %d", *p.IPDstPortStartMatchCondition))
	}
	if p.IPDstPortEndMatchCondition != nil {
		cmd = append(cmd, fmt.Sprintf("-IPDstPortEndMatchCondition %d", *p.IPDstPortEndMatchCondition))
	}
	if p.ThrottleLimit != nil {
		cmd = append(cmd, fmt.Sprintf("-ThrottleLimit %d", *p.ThrottleLimit))
	}
	if p.TemplateMatchCondition != "" {
		if action == Action_Add {
			cmd = append(cmd, fmt.Sprintf("-%s", p.TemplateMatchCondition))
		} else {
			cmd = append(cmd, fmt.Sprintf("-TemplateMatchCondition %s", p.TemplateMatchCondition))
		}

	}
	if p.UriMatchCondition != "" {
		cmd = append(cmd, fmt.Sprintf("-URIMatchCondition %s", p.UriMatchCondition))
		if p.UriRecursiveMatchCondition {
			cmd = append(cmd, fmt.Sprintf("-URIRecursiveMatchCondition $True"))
		}
		if p.NetworkProfile != Profile_Domain {
			log.Println("WARN: UriMatchCondition will force the NetworkProfile to 'Domain'")
			p.NetworkProfile = Profile_Domain
		}
	}

	command = strings.Join(cmd, " ")
	return command, nil
}

func GetAll() (policies []*NetQoSPolicy, err error) { //todo: There is a lot of duplication between GetAll() and Get() refactor
	stdout, stderr, err := executeGetAllPolicies(false)
	if err != nil && strings.Contains(stderr, "Access is denied") {
		return nil, fmt.Errorf("access denied. try running as administrator: %v", err)
	} else if err != nil && !strings.Contains(stderr, falseError) {
		return nil, err
	}
	stdoutT, _, _ := executeGetAllPolicies(true)
	//todo: refactor this, not checking the error.

	policies, err = parsePolicyOutput(stdout, "")
	if err != nil {
		return nil, err
	}

	if stdoutT != "" {
		policiesT, err := parsePolicyOutput(stdoutT, "ActiveStore")
		if err != nil {
			return nil, err
		}
		policies = append(policies, policiesT...)
	}

	return policies, nil
}

func Get(name string) (policy *NetQoSPolicy, err error) {
	var policyStore string
	noPolicyFoundError := "No MSFT_NetQosPolicySettingData objects found"
	stdout, stderr, err := executeGetPolicy(name, false)
	if err != nil && strings.Contains(stderr, noPolicyFoundError) {
		// try looking for temporary profiles
		stdout, stderr, err = executeGetPolicy(name, true)
		policyStore = "ActiveStore"
	}

	if err != nil && strings.Contains(stderr, "Access is denied") {
		return nil, fmt.Errorf("access denied. try running as administrator: %v", err)
	} else if err != nil && strings.Contains(stderr, noPolicyFoundError) {
		return nil, fmt.Errorf("no policy with that name found")
	} else if err != nil && !strings.Contains(stderr, falseError) {
		return nil, err
	}

	policies, err := parsePolicyOutput(stdout, policyStore)
	if err != nil {
		return nil, err
	}

	if len(policies) > 1 {
		log.WithFields(log.Fields{
			"search_name":  name,
			"policy_count": len(policies),
			"policies":     policies,
		}).Warning("search value matched more then one policy.")
	}
	if len(policies) == 0 {
		return nil, fmt.Errorf("no matching policy found") //note: this should never be hit since it should result in an error in the command
	}
	return policies[0], nil
}

func Remove(name string) (err error) {
	policy, err := Get(name)
	if err != nil {
		return err
	}
	return policy.Remove()
}

func executeGetAllPolicies(temp bool) (stdout string, stderr string, err error) {
	cmd := []string{
		cmdShell,
		getCommandlet,
	}
	if temp {
		cmd = append(cmd, "-PolicyStore ActiveStore")
	}
	return execute.ExecuteCmdEx(strings.Join(cmd, " "))
}

func executeGetPolicy(name string, temp bool) (stdout string, stderr string, err error) {
	if strings.Contains(name, " ") && !strings.Contains(name, "\"") {
		name = fmt.Sprintf("\"%s\"", name)
	}
	cmd := []string{
		cmdShell,
		getCommandlet,
		"-Name",
		name}
	if temp {
		cmd = append(cmd, "-PolicyStore ActiveStore")
	}
	return execute.ExecuteCmdEx(strings.Join(cmd, " "))
}

func parsePolicyOutput(output string, store string) (policies []*NetQoSPolicy, err error) {
	p := &NetQoSPolicy{}
	policies = make([]*NetQoSPolicy, 0)
	lines := strings.Split(output, "\r\n")
	for _, line := range lines {
		if line == "" {
			if p.Name != "" {
				// Cleanup
				if p.IPSrcPortStartMatchCondition != nil && *p.IPSrcPortStartMatchCondition == *p.IPSrcPortEndMatchCondition {
					p.IPSrcPortMatchCondition = p.IPSrcPortStartMatchCondition
					p.IPSrcPortStartMatchCondition = nil
					p.IPSrcPortEndMatchCondition = nil
				}
				if p.IPDstPortStartMatchCondition != nil && *p.IPDstPortStartMatchCondition == *p.IPDstPortEndMatchCondition {
					p.IPDstPortMatchCondition = p.IPDstPortStartMatchCondition
					p.IPDstPortStartMatchCondition = nil
					p.IPDstPortEndMatchCondition = nil
				}
				if store == "ActiveStore" {
					p.Persistent = false
					p.PolicyStore = store
				} else {
					p.Persistent = true
				}
				policies = append(policies, p)
				p = &NetQoSPolicy{}
			}
			continue
		}
		parts := strings.Split(line, ":")
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(strings.Join(parts[1:], ":"))
		switch key {
		case "AppPathName":
			p.AppPathNameMatchCondition = value
		case "DSCPValue":
			raw, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			v := int8(raw)
			p.DSCPAction = &v
		case "IPProtocol":
			p.IPProtocolMatchCondition = IPProtocol(value)
		case "IPSrcPortEnd":
			raw, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			v := uint16(raw)
			p.IPSrcPortEndMatchCondition = &v
		case "IPSrcPortStart":
			raw, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			v := uint16(raw)
			p.IPSrcPortStartMatchCondition = &v
		case "IPSrcPrefix":
			p.IPSrcPrefixMatchCondition = value
		case "IPDstPortEnd":
			raw, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			v := uint16(raw)
			p.IPDstPortEndMatchCondition = &v
		case "IPDstPortStart":
			raw, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			v := uint16(raw)
			p.IPDstPortStartMatchCondition = &v
		case "IPDstPrefix":
			p.IPDstPrefixMatchCondition = value
		case "JobObject":
			continue
		case "MinBandwidthWeight":
			raw, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			v := byte(raw)
			p.MinBandwidthWeightAction = &v
		case "Name":
			p.Name = value
		case "NetworkProfile":
			p.NetworkProfile = NetworkProfile(value)
		case "Owner":
			p.Owner = value
		case "Precedence":
			raw, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			v := uint32(raw)
			p.Precedence = &v
		case "PriorityValue":
			raw, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			v := int8(raw)
			p.PriorityValue8021Action = &v
		case "Templates":
			p.TemplateMatchCondition = Template(value)
		case "ThrottleRate":
			rate, err := conversions.StringBitRateToInt(value)
			if err != nil {
				return nil, err
			}
			v := uint64(rate)
			p.ThrottleRateActionBitsPerSecond = &v
		case "URI":
			p.UriMatchCondition = value
		case "URIRecursive":
			v, _ := strconv.ParseBool(value)
			p.UriRecursiveMatchCondition = v
		case "User":
			p.UserMatchCondition = value
		default:
			log.WithFields(log.Fields{
				"key":   key,
				"value": value,
			}).Warn("unknown field encountered while parsing")
		}
	}
	return policies, nil
}

func executeCommand(command string) (err error) {
	if DEBUG {
		execute.Debug = true
	}
	stdout, stderr, err := execute.ExecutePowershell(command)
	if DEBUG {
		fmt.Println("output: " + stdout)
	}
	if err != nil {
		return fmt.Errorf("%s", stderr)
	}
	return nil
}
