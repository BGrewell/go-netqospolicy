package netqospolicy

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name       string
		args       args
		wantPolicy *NetQoSPolicy
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPolicy, err := Get(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPolicy, tt.wantPolicy) {
				t.Errorf("Get() gotPolicy = %v, want %v", gotPolicy, tt.wantPolicy)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	tests := []struct {
		name         string
		wantPolicies []*NetQoSPolicy
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPolicies, err := GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPolicies, tt.wantPolicies) {
				t.Errorf("GetAll() gotPolicies = %v, want %v", gotPolicies, tt.wantPolicies)
			}
		})
	}
}

func TestNetQoSPolicy_Apply(t *testing.T) {
	type fields struct {
		PolicyStore                     string
		Name                            string
		Owner                           string
		NetworkProfile                  NetworkProfile
		Precedence                      *uint32
		PriorityValue8021Action         *int8
		DSCPAction                      *int8
		MinBandwidthWeightAction        *byte
		ThrottleRateActionBitsPerSecond *uint64
		UserMatchCondition              string
		AppPathNameMatchCondition       string
		IPProtocolMatchCondition        IPProtocol
		IPSrcPrefixMatchCondition       string
		IPSrcPortMatchCondition         *uint16
		IPSrcPortStartMatchCondition    *uint16
		IPSrcPortEndMatchCondition      *uint16
		IPDstPrefixMatchCondition       string
		IPDstPortMatchCondition         *uint16
		IPDstPortStartMatchCondition    *uint16
		IPDstPortEndMatchCondition      *uint16
		ThrottleLimit                   *int32
		Templates                       []Template
		Persistent                      bool
		FmtNoIndent                     bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &NetQoSPolicy{
				PolicyStore:                     tt.fields.PolicyStore,
				Name:                            tt.fields.Name,
				Owner:                           tt.fields.Owner,
				NetworkProfile:                  tt.fields.NetworkProfile,
				Precedence:                      tt.fields.Precedence,
				PriorityValue8021Action:         tt.fields.PriorityValue8021Action,
				DSCPAction:                      tt.fields.DSCPAction,
				MinBandwidthWeightAction:        tt.fields.MinBandwidthWeightAction,
				ThrottleRateActionBitsPerSecond: tt.fields.ThrottleRateActionBitsPerSecond,
				UserMatchCondition:              tt.fields.UserMatchCondition,
				AppPathNameMatchCondition:       tt.fields.AppPathNameMatchCondition,
				IPProtocolMatchCondition:        tt.fields.IPProtocolMatchCondition,
				IPSrcPrefixMatchCondition:       tt.fields.IPSrcPrefixMatchCondition,
				IPSrcPortMatchCondition:         tt.fields.IPSrcPortMatchCondition,
				IPSrcPortStartMatchCondition:    tt.fields.IPSrcPortStartMatchCondition,
				IPSrcPortEndMatchCondition:      tt.fields.IPSrcPortEndMatchCondition,
				IPDstPrefixMatchCondition:       tt.fields.IPDstPrefixMatchCondition,
				IPDstPortMatchCondition:         tt.fields.IPDstPortMatchCondition,
				IPDstPortStartMatchCondition:    tt.fields.IPDstPortStartMatchCondition,
				IPDstPortEndMatchCondition:      tt.fields.IPDstPortEndMatchCondition,
				ThrottleLimit:                   tt.fields.ThrottleLimit,
				Templates:                       tt.fields.Templates,
				Persistent:                      tt.fields.Persistent,
				FmtNoIndent:                     tt.fields.FmtNoIndent,
			}
			if err := p.Update(); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNetQoSPolicy_Create(t *testing.T) {
	type fields struct {
		PolicyStore                     string
		Name                            string
		Owner                           string
		NetworkProfile                  NetworkProfile
		Precedence                      *uint32
		PriorityValue8021Action         *int8
		DSCPAction                      *int8
		MinBandwidthWeightAction        *byte
		ThrottleRateActionBitsPerSecond *uint64
		UserMatchCondition              string
		AppPathNameMatchCondition       string
		IPProtocolMatchCondition        IPProtocol
		IPSrcPrefixMatchCondition       string
		IPSrcPortMatchCondition         *uint16
		IPSrcPortStartMatchCondition    *uint16
		IPSrcPortEndMatchCondition      *uint16
		IPDstPrefixMatchCondition       string
		IPDstPortMatchCondition         *uint16
		IPDstPortStartMatchCondition    *uint16
		IPDstPortEndMatchCondition      *uint16
		ThrottleLimit                   *int32
		Templates                       []Template
		Persistent                      bool
		FmtNoIndent                     bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &NetQoSPolicy{
				PolicyStore:                     tt.fields.PolicyStore,
				Name:                            tt.fields.Name,
				Owner:                           tt.fields.Owner,
				NetworkProfile:                  tt.fields.NetworkProfile,
				Precedence:                      tt.fields.Precedence,
				PriorityValue8021Action:         tt.fields.PriorityValue8021Action,
				DSCPAction:                      tt.fields.DSCPAction,
				MinBandwidthWeightAction:        tt.fields.MinBandwidthWeightAction,
				ThrottleRateActionBitsPerSecond: tt.fields.ThrottleRateActionBitsPerSecond,
				UserMatchCondition:              tt.fields.UserMatchCondition,
				AppPathNameMatchCondition:       tt.fields.AppPathNameMatchCondition,
				IPProtocolMatchCondition:        tt.fields.IPProtocolMatchCondition,
				IPSrcPrefixMatchCondition:       tt.fields.IPSrcPrefixMatchCondition,
				IPSrcPortMatchCondition:         tt.fields.IPSrcPortMatchCondition,
				IPSrcPortStartMatchCondition:    tt.fields.IPSrcPortStartMatchCondition,
				IPSrcPortEndMatchCondition:      tt.fields.IPSrcPortEndMatchCondition,
				IPDstPrefixMatchCondition:       tt.fields.IPDstPrefixMatchCondition,
				IPDstPortMatchCondition:         tt.fields.IPDstPortMatchCondition,
				IPDstPortStartMatchCondition:    tt.fields.IPDstPortStartMatchCondition,
				IPDstPortEndMatchCondition:      tt.fields.IPDstPortEndMatchCondition,
				ThrottleLimit:                   tt.fields.ThrottleLimit,
				Templates:                       tt.fields.Templates,
				Persistent:                      tt.fields.Persistent,
				FmtNoIndent:                     tt.fields.FmtNoIndent,
			}
			if err := p.Create(); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNetQoSPolicy_IsApplied(t *testing.T) {
	type fields struct {
		PolicyStore                     string
		Name                            string
		Owner                           string
		NetworkProfile                  NetworkProfile
		Precedence                      *uint32
		PriorityValue8021Action         *int8
		DSCPAction                      *int8
		MinBandwidthWeightAction        *byte
		ThrottleRateActionBitsPerSecond *uint64
		UserMatchCondition              string
		AppPathNameMatchCondition       string
		IPProtocolMatchCondition        IPProtocol
		IPSrcPrefixMatchCondition       string
		IPSrcPortMatchCondition         *uint16
		IPSrcPortStartMatchCondition    *uint16
		IPSrcPortEndMatchCondition      *uint16
		IPDstPrefixMatchCondition       string
		IPDstPortMatchCondition         *uint16
		IPDstPortStartMatchCondition    *uint16
		IPDstPortEndMatchCondition      *uint16
		ThrottleLimit                   *int32
		Templates                       []Template
		Persistent                      bool
		FmtNoIndent                     bool
	}
	tests := []struct {
		name        string
		fields      fields
		wantApplied bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &NetQoSPolicy{
				PolicyStore:                     tt.fields.PolicyStore,
				Name:                            tt.fields.Name,
				Owner:                           tt.fields.Owner,
				NetworkProfile:                  tt.fields.NetworkProfile,
				Precedence:                      tt.fields.Precedence,
				PriorityValue8021Action:         tt.fields.PriorityValue8021Action,
				DSCPAction:                      tt.fields.DSCPAction,
				MinBandwidthWeightAction:        tt.fields.MinBandwidthWeightAction,
				ThrottleRateActionBitsPerSecond: tt.fields.ThrottleRateActionBitsPerSecond,
				UserMatchCondition:              tt.fields.UserMatchCondition,
				AppPathNameMatchCondition:       tt.fields.AppPathNameMatchCondition,
				IPProtocolMatchCondition:        tt.fields.IPProtocolMatchCondition,
				IPSrcPrefixMatchCondition:       tt.fields.IPSrcPrefixMatchCondition,
				IPSrcPortMatchCondition:         tt.fields.IPSrcPortMatchCondition,
				IPSrcPortStartMatchCondition:    tt.fields.IPSrcPortStartMatchCondition,
				IPSrcPortEndMatchCondition:      tt.fields.IPSrcPortEndMatchCondition,
				IPDstPrefixMatchCondition:       tt.fields.IPDstPrefixMatchCondition,
				IPDstPortMatchCondition:         tt.fields.IPDstPortMatchCondition,
				IPDstPortStartMatchCondition:    tt.fields.IPDstPortStartMatchCondition,
				IPDstPortEndMatchCondition:      tt.fields.IPDstPortEndMatchCondition,
				ThrottleLimit:                   tt.fields.ThrottleLimit,
				Templates:                       tt.fields.Templates,
				Persistent:                      tt.fields.Persistent,
				FmtNoIndent:                     tt.fields.FmtNoIndent,
			}
			gotApplied, err := p.IsApplied()
			if (err != nil) != tt.wantErr {
				t.Errorf("IsApplied() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotApplied != tt.wantApplied {
				t.Errorf("IsApplied() gotApplied = %v, want %v", gotApplied, tt.wantApplied)
			}
		})
	}
}

func TestNetQoSPolicy_Remove(t *testing.T) {
	type fields struct {
		PolicyStore                     string
		Name                            string
		Owner                           string
		NetworkProfile                  NetworkProfile
		Precedence                      *uint32
		PriorityValue8021Action         *int8
		DSCPAction                      *int8
		MinBandwidthWeightAction        *byte
		ThrottleRateActionBitsPerSecond *uint64
		UserMatchCondition              string
		AppPathNameMatchCondition       string
		IPProtocolMatchCondition        IPProtocol
		IPSrcPrefixMatchCondition       string
		IPSrcPortMatchCondition         *uint16
		IPSrcPortStartMatchCondition    *uint16
		IPSrcPortEndMatchCondition      *uint16
		IPDstPrefixMatchCondition       string
		IPDstPortMatchCondition         *uint16
		IPDstPortStartMatchCondition    *uint16
		IPDstPortEndMatchCondition      *uint16
		ThrottleLimit                   *int32
		Templates                       []Template
		Persistent                      bool
		FmtNoIndent                     bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &NetQoSPolicy{
				PolicyStore:                     tt.fields.PolicyStore,
				Name:                            tt.fields.Name,
				Owner:                           tt.fields.Owner,
				NetworkProfile:                  tt.fields.NetworkProfile,
				Precedence:                      tt.fields.Precedence,
				PriorityValue8021Action:         tt.fields.PriorityValue8021Action,
				DSCPAction:                      tt.fields.DSCPAction,
				MinBandwidthWeightAction:        tt.fields.MinBandwidthWeightAction,
				ThrottleRateActionBitsPerSecond: tt.fields.ThrottleRateActionBitsPerSecond,
				UserMatchCondition:              tt.fields.UserMatchCondition,
				AppPathNameMatchCondition:       tt.fields.AppPathNameMatchCondition,
				IPProtocolMatchCondition:        tt.fields.IPProtocolMatchCondition,
				IPSrcPrefixMatchCondition:       tt.fields.IPSrcPrefixMatchCondition,
				IPSrcPortMatchCondition:         tt.fields.IPSrcPortMatchCondition,
				IPSrcPortStartMatchCondition:    tt.fields.IPSrcPortStartMatchCondition,
				IPSrcPortEndMatchCondition:      tt.fields.IPSrcPortEndMatchCondition,
				IPDstPrefixMatchCondition:       tt.fields.IPDstPrefixMatchCondition,
				IPDstPortMatchCondition:         tt.fields.IPDstPortMatchCondition,
				IPDstPortStartMatchCondition:    tt.fields.IPDstPortStartMatchCondition,
				IPDstPortEndMatchCondition:      tt.fields.IPDstPortEndMatchCondition,
				ThrottleLimit:                   tt.fields.ThrottleLimit,
				Templates:                       tt.fields.Templates,
				Persistent:                      tt.fields.Persistent,
				FmtNoIndent:                     tt.fields.FmtNoIndent,
			}
			if err := p.Remove(); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNetQoSPolicy_String(t *testing.T) {
	type fields struct {
		PolicyStore                     string
		Name                            string
		Owner                           string
		NetworkProfile                  NetworkProfile
		Precedence                      *uint32
		PriorityValue8021Action         *int8
		DSCPAction                      *int8
		MinBandwidthWeightAction        *byte
		ThrottleRateActionBitsPerSecond *uint64
		UserMatchCondition              string
		AppPathNameMatchCondition       string
		IPProtocolMatchCondition        IPProtocol
		IPSrcPrefixMatchCondition       string
		IPSrcPortMatchCondition         *uint16
		IPSrcPortStartMatchCondition    *uint16
		IPSrcPortEndMatchCondition      *uint16
		IPDstPrefixMatchCondition       string
		IPDstPortMatchCondition         *uint16
		IPDstPortStartMatchCondition    *uint16
		IPDstPortEndMatchCondition      *uint16
		ThrottleLimit                   *int32
		Templates                       []Template
		Persistent                      bool
		FmtNoIndent                     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NetQoSPolicy{
				PolicyStore:                     tt.fields.PolicyStore,
				Name:                            tt.fields.Name,
				Owner:                           tt.fields.Owner,
				NetworkProfile:                  tt.fields.NetworkProfile,
				Precedence:                      tt.fields.Precedence,
				PriorityValue8021Action:         tt.fields.PriorityValue8021Action,
				DSCPAction:                      tt.fields.DSCPAction,
				MinBandwidthWeightAction:        tt.fields.MinBandwidthWeightAction,
				ThrottleRateActionBitsPerSecond: tt.fields.ThrottleRateActionBitsPerSecond,
				UserMatchCondition:              tt.fields.UserMatchCondition,
				AppPathNameMatchCondition:       tt.fields.AppPathNameMatchCondition,
				IPProtocolMatchCondition:        tt.fields.IPProtocolMatchCondition,
				IPSrcPrefixMatchCondition:       tt.fields.IPSrcPrefixMatchCondition,
				IPSrcPortMatchCondition:         tt.fields.IPSrcPortMatchCondition,
				IPSrcPortStartMatchCondition:    tt.fields.IPSrcPortStartMatchCondition,
				IPSrcPortEndMatchCondition:      tt.fields.IPSrcPortEndMatchCondition,
				IPDstPrefixMatchCondition:       tt.fields.IPDstPrefixMatchCondition,
				IPDstPortMatchCondition:         tt.fields.IPDstPortMatchCondition,
				IPDstPortStartMatchCondition:    tt.fields.IPDstPortStartMatchCondition,
				IPDstPortEndMatchCondition:      tt.fields.IPDstPortEndMatchCondition,
				ThrottleLimit:                   tt.fields.ThrottleLimit,
				Templates:                       tt.fields.Templates,
				Persistent:                      tt.fields.Persistent,
				FmtNoIndent:                     tt.fields.FmtNoIndent,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetQoSPolicy_getCommand(t *testing.T) {
	type fields struct {
		PolicyStore                     string
		Name                            string
		Owner                           string
		NetworkProfile                  NetworkProfile
		Precedence                      *uint32
		PriorityValue8021Action         *int8
		DSCPAction                      *int8
		MinBandwidthWeightAction        *byte
		ThrottleRateActionBitsPerSecond *uint64
		UserMatchCondition              string
		AppPathNameMatchCondition       string
		IPProtocolMatchCondition        IPProtocol
		IPSrcPrefixMatchCondition       string
		IPSrcPortMatchCondition         *uint16
		IPSrcPortStartMatchCondition    *uint16
		IPSrcPortEndMatchCondition      *uint16
		IPDstPrefixMatchCondition       string
		IPDstPortMatchCondition         *uint16
		IPDstPortStartMatchCondition    *uint16
		IPDstPortEndMatchCondition      *uint16
		ThrottleLimit                   *int32
		Templates                       []Template
		Persistent                      bool
		FmtNoIndent                     bool
	}
	type args struct {
		action string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantCommand string
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NetQoSPolicy{
				PolicyStore:                     tt.fields.PolicyStore,
				Name:                            tt.fields.Name,
				Owner:                           tt.fields.Owner,
				NetworkProfile:                  tt.fields.NetworkProfile,
				Precedence:                      tt.fields.Precedence,
				PriorityValue8021Action:         tt.fields.PriorityValue8021Action,
				DSCPAction:                      tt.fields.DSCPAction,
				MinBandwidthWeightAction:        tt.fields.MinBandwidthWeightAction,
				ThrottleRateActionBitsPerSecond: tt.fields.ThrottleRateActionBitsPerSecond,
				UserMatchCondition:              tt.fields.UserMatchCondition,
				AppPathNameMatchCondition:       tt.fields.AppPathNameMatchCondition,
				IPProtocolMatchCondition:        tt.fields.IPProtocolMatchCondition,
				IPSrcPrefixMatchCondition:       tt.fields.IPSrcPrefixMatchCondition,
				IPSrcPortMatchCondition:         tt.fields.IPSrcPortMatchCondition,
				IPSrcPortStartMatchCondition:    tt.fields.IPSrcPortStartMatchCondition,
				IPSrcPortEndMatchCondition:      tt.fields.IPSrcPortEndMatchCondition,
				IPDstPrefixMatchCondition:       tt.fields.IPDstPrefixMatchCondition,
				IPDstPortMatchCondition:         tt.fields.IPDstPortMatchCondition,
				IPDstPortStartMatchCondition:    tt.fields.IPDstPortStartMatchCondition,
				IPDstPortEndMatchCondition:      tt.fields.IPDstPortEndMatchCondition,
				ThrottleLimit:                   tt.fields.ThrottleLimit,
				Templates:                       tt.fields.Templates,
				Persistent:                      tt.fields.Persistent,
				FmtNoIndent:                     tt.fields.FmtNoIndent,
			}
			gotCommand, err := p.getCommand(tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCommand != tt.wantCommand {
				t.Errorf("getCommand() gotCommand = %v, want %v", gotCommand, tt.wantCommand)
			}
		})
	}
}

func TestNetQoSPolicy_setDefaultValues(t *testing.T) {
	type fields struct {
		PolicyStore                     string
		Name                            string
		Owner                           string
		NetworkProfile                  NetworkProfile
		Precedence                      *uint32
		PriorityValue8021Action         *int8
		DSCPAction                      *int8
		MinBandwidthWeightAction        *byte
		ThrottleRateActionBitsPerSecond *uint64
		UserMatchCondition              string
		AppPathNameMatchCondition       string
		IPProtocolMatchCondition        IPProtocol
		IPSrcPrefixMatchCondition       string
		IPSrcPortMatchCondition         *uint16
		IPSrcPortStartMatchCondition    *uint16
		IPSrcPortEndMatchCondition      *uint16
		IPDstPrefixMatchCondition       string
		IPDstPortMatchCondition         *uint16
		IPDstPortStartMatchCondition    *uint16
		IPDstPortEndMatchCondition      *uint16
		ThrottleLimit                   *int32
		Templates                       []Template
		Persistent                      bool
		FmtNoIndent                     bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &NetQoSPolicy{
				PolicyStore:                     tt.fields.PolicyStore,
				Name:                            tt.fields.Name,
				Owner:                           tt.fields.Owner,
				NetworkProfile:                  tt.fields.NetworkProfile,
				Precedence:                      tt.fields.Precedence,
				PriorityValue8021Action:         tt.fields.PriorityValue8021Action,
				DSCPAction:                      tt.fields.DSCPAction,
				MinBandwidthWeightAction:        tt.fields.MinBandwidthWeightAction,
				ThrottleRateActionBitsPerSecond: tt.fields.ThrottleRateActionBitsPerSecond,
				UserMatchCondition:              tt.fields.UserMatchCondition,
				AppPathNameMatchCondition:       tt.fields.AppPathNameMatchCondition,
				IPProtocolMatchCondition:        tt.fields.IPProtocolMatchCondition,
				IPSrcPrefixMatchCondition:       tt.fields.IPSrcPrefixMatchCondition,
				IPSrcPortMatchCondition:         tt.fields.IPSrcPortMatchCondition,
				IPSrcPortStartMatchCondition:    tt.fields.IPSrcPortStartMatchCondition,
				IPSrcPortEndMatchCondition:      tt.fields.IPSrcPortEndMatchCondition,
				IPDstPrefixMatchCondition:       tt.fields.IPDstPrefixMatchCondition,
				IPDstPortMatchCondition:         tt.fields.IPDstPortMatchCondition,
				IPDstPortStartMatchCondition:    tt.fields.IPDstPortStartMatchCondition,
				IPDstPortEndMatchCondition:      tt.fields.IPDstPortEndMatchCondition,
				ThrottleLimit:                   tt.fields.ThrottleLimit,
				Templates:                       tt.fields.Templates,
				Persistent:                      tt.fields.Persistent,
				FmtNoIndent:                     tt.fields.FmtNoIndent,
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Remove(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_executeCommand(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := executeCommand(tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("executeCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_executeGetAllPolicies(t *testing.T) {
	type args struct {
		temp bool
	}
	tests := []struct {
		name       string
		args       args
		wantStdout string
		wantStderr string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStdout, gotStderr, err := executeGetAllPolicies(tt.args.temp)
			if (err != nil) != tt.wantErr {
				t.Errorf("executeGetAllPolicies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStdout != tt.wantStdout {
				t.Errorf("executeGetAllPolicies() gotStdout = %v, want %v", gotStdout, tt.wantStdout)
			}
			if gotStderr != tt.wantStderr {
				t.Errorf("executeGetAllPolicies() gotStderr = %v, want %v", gotStderr, tt.wantStderr)
			}
		})
	}
}

func Test_executeGetPolicy(t *testing.T) {
	type args struct {
		name string
		temp bool
	}
	tests := []struct {
		name       string
		args       args
		wantStdout string
		wantStderr string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStdout, gotStderr, err := executeGetPolicy(tt.args.name, tt.args.temp)
			if (err != nil) != tt.wantErr {
				t.Errorf("executeGetPolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStdout != tt.wantStdout {
				t.Errorf("executeGetPolicy() gotStdout = %v, want %v", gotStdout, tt.wantStdout)
			}
			if gotStderr != tt.wantStderr {
				t.Errorf("executeGetPolicy() gotStderr = %v, want %v", gotStderr, tt.wantStderr)
			}
		})
	}
}

func Test_parsePolicyOutput(t *testing.T) {
	type args struct {
		output string
		store  string
	}
	tests := []struct {
		name         string
		args         args
		wantPolicies []*NetQoSPolicy
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPolicies, err := parsePolicyOutput(tt.args.output, tt.args.store)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePolicyOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPolicies, tt.wantPolicies) {
				t.Errorf("parsePolicyOutput() gotPolicies = %v, want %v", gotPolicies, tt.wantPolicies)
			}
		})
	}
}
