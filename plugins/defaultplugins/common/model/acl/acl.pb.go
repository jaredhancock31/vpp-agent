// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: acl.proto

/*
Package acl is a generated protocol buffer package.

It is generated from these files:
	acl.proto

It has these top-level messages:
	AccessLists
*/
package acl

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AclAction int32

const (
	AclAction_DENY    AclAction = 0
	AclAction_PERMIT  AclAction = 1
	AclAction_REFLECT AclAction = 2
)

var AclAction_name = map[int32]string{
	0: "DENY",
	1: "PERMIT",
	2: "REFLECT",
}
var AclAction_value = map[string]int32{
	"DENY":    0,
	"PERMIT":  1,
	"REFLECT": 2,
}

func (x AclAction) String() string {
	return proto.EnumName(AclAction_name, int32(x))
}
func (AclAction) EnumDescriptor() ([]byte, []int) { return fileDescriptorAcl, []int{0} }

// This is a top level container for Access Control Lists.
// It can have one or more Access Control Lists.
type AccessLists struct {
	Acl []*AccessLists_Acl `protobuf:"bytes,1,rep,name=acl" json:"acl,omitempty"`
}

func (m *AccessLists) Reset()                    { *m = AccessLists{} }
func (m *AccessLists) String() string            { return proto.CompactTextString(m) }
func (*AccessLists) ProtoMessage()               {}
func (*AccessLists) Descriptor() ([]byte, []int) { return fileDescriptorAcl, []int{0} }

func (m *AccessLists) GetAcl() []*AccessLists_Acl {
	if m != nil {
		return m.Acl
	}
	return nil
}

// An Access Control List(ACL) is an ordered list of Access List Rules. Each Access Control Rule has
// a list of match criteria and a list of actions.
type AccessLists_Acl struct {
	Rules []*AccessLists_Acl_Rule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
	// The name of access-list. A device MAY restrict the length
	// and value of this name, possibly spRule and special
	// characters are not allowed.
	AclName string `protobuf:"bytes,3,opt,name=acl_name,json=aclName,proto3" json:"acl_name,omitempty"`
	// The set of interfRules that has assigned this ACL on ingres or egress
	Interfaces *AccessLists_Acl_Interface `protobuf:"bytes,2,opt,name=interfaces" json:"interfaces,omitempty"`
}

func (m *AccessLists_Acl) Reset()                    { *m = AccessLists_Acl{} }
func (m *AccessLists_Acl) String() string            { return proto.CompactTextString(m) }
func (*AccessLists_Acl) ProtoMessage()               {}
func (*AccessLists_Acl) Descriptor() ([]byte, []int) { return fileDescriptorAcl, []int{0, 0} }

func (m *AccessLists_Acl) GetRules() []*AccessLists_Acl_Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *AccessLists_Acl) GetAclName() string {
	if m != nil {
		return m.AclName
	}
	return ""
}

func (m *AccessLists_Acl) GetInterfaces() *AccessLists_Acl_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

// List of access list entries(Rule)
type AccessLists_Acl_Rule struct {
	Actions *AccessLists_Acl_Rule_Action `protobuf:"bytes,2,opt,name=actions" json:"actions,omitempty"`
	Matches *AccessLists_Acl_Rule_Match  `protobuf:"bytes,3,opt,name=matches" json:"matches,omitempty"`
	// Access List entry that can define:
	// - IP4/IP6 src ip prefix
	// - src MAC address mask
	// - src MAC address value
	// - can be used only for static ACLs.
	// A unique name identifying this Access List
	// Entry(Rule).
	RuleName string `protobuf:"bytes,5,opt,name=rule_name,json=ruleName,proto3" json:"rule_name,omitempty"`
}

func (m *AccessLists_Acl_Rule) Reset()                    { *m = AccessLists_Acl_Rule{} }
func (m *AccessLists_Acl_Rule) String() string            { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule) ProtoMessage()               {}
func (*AccessLists_Acl_Rule) Descriptor() ([]byte, []int) { return fileDescriptorAcl, []int{0, 0, 0} }

func (m *AccessLists_Acl_Rule) GetActions() *AccessLists_Acl_Rule_Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *AccessLists_Acl_Rule) GetMatches() *AccessLists_Acl_Rule_Match {
	if m != nil {
		return m.Matches
	}
	return nil
}

func (m *AccessLists_Acl_Rule) GetRuleName() string {
	if m != nil {
		return m.RuleName
	}
	return ""
}

// Definitions of action criteria for this Access List Rule
type AccessLists_Acl_Rule_Action struct {
	AclAction AclAction `protobuf:"varint,1,opt,name=acl_action,json=aclAction,proto3,enum=acl.AclAction" json:"acl_action,omitempty"`
}

func (m *AccessLists_Acl_Rule_Action) Reset()         { *m = AccessLists_Acl_Rule_Action{} }
func (m *AccessLists_Acl_Rule_Action) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Action) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Action) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 0}
}

func (m *AccessLists_Acl_Rule_Action) GetAclAction() AclAction {
	if m != nil {
		return m.AclAction
	}
	return AclAction_DENY
}

// Definitions for match criteria for this Access List Rule
type AccessLists_Acl_Rule_Match struct {
	IpRule    *AccessLists_Acl_Rule_Match_IpRule    `protobuf:"bytes,1,opt,name=ip_rule,json=ipRule" json:"ip_rule,omitempty"`
	MacipRule *AccessLists_Acl_Rule_Match_MacIpRule `protobuf:"bytes,4,opt,name=macip_rule,json=macipRule" json:"macip_rule,omitempty"`
}

func (m *AccessLists_Acl_Rule_Match) Reset()         { *m = AccessLists_Acl_Rule_Match{} }
func (m *AccessLists_Acl_Rule_Match) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 1}
}

func (m *AccessLists_Acl_Rule_Match) GetIpRule() *AccessLists_Acl_Rule_Match_IpRule {
	if m != nil {
		return m.IpRule
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match) GetMacipRule() *AccessLists_Acl_Rule_Match_MacIpRule {
	if m != nil {
		return m.MacipRule
	}
	return nil
}

// Access List entry that can define:
// - IP4/IP6 src/dst ip prefix- Internet Protocol number
// - Internet Protocol number
// - selected L4 headers:
//   * ICMP (type range)
//   * UDP (port range)
//   * TCP (port range, flags mask, flags value)
type AccessLists_Acl_Rule_Match_IpRule struct {
	Ip    *AccessLists_Acl_Rule_Match_IpRule_Ip    `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Icmp  *AccessLists_Acl_Rule_Match_IpRule_Icmp  `protobuf:"bytes,2,opt,name=icmp" json:"icmp,omitempty"`
	Tcp   *AccessLists_Acl_Rule_Match_IpRule_Tcp   `protobuf:"bytes,4,opt,name=tcp" json:"tcp,omitempty"`
	Udp   *AccessLists_Acl_Rule_Match_IpRule_Udp   `protobuf:"bytes,5,opt,name=udp" json:"udp,omitempty"`
	Other *AccessLists_Acl_Rule_Match_IpRule_Other `protobuf:"bytes,3,opt,name=other" json:"other,omitempty"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule) Reset()         { *m = AccessLists_Acl_Rule_Match_IpRule{} }
func (m *AccessLists_Acl_Rule_Match_IpRule) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match_IpRule) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match_IpRule) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 1, 0}
}

func (m *AccessLists_Acl_Rule_Match_IpRule) GetIp() *AccessLists_Acl_Rule_Match_IpRule_Ip {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule) GetIcmp() *AccessLists_Acl_Rule_Match_IpRule_Icmp {
	if m != nil {
		return m.Icmp
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule) GetTcp() *AccessLists_Acl_Rule_Match_IpRule_Tcp {
	if m != nil {
		return m.Tcp
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule) GetUdp() *AccessLists_Acl_Rule_Match_IpRule_Udp {
	if m != nil {
		return m.Udp
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule) GetOther() *AccessLists_Acl_Rule_Match_IpRule_Other {
	if m != nil {
		return m.Other
	}
	return nil
}

// IP version used in this Access List Entry.
type AccessLists_Acl_Rule_Match_IpRule_Ip struct {
	// Destination IPv4/IPv6 network address (<ip>/<network>)
	DestinationNetwork string `protobuf:"bytes,1,opt,name=destination_network,json=destinationNetwork,proto3" json:"destination_network,omitempty"`
	// Destination IPv4/IPv6 network address (<ip>/<network>)
	SourceNetwork string `protobuf:"bytes,2,opt,name=source_network,json=sourceNetwork,proto3" json:"source_network,omitempty"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Ip) Reset()         { *m = AccessLists_Acl_Rule_Match_IpRule_Ip{} }
func (m *AccessLists_Acl_Rule_Match_IpRule_Ip) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match_IpRule_Ip) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match_IpRule_Ip) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 1, 0, 0}
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Ip) GetDestinationNetwork() string {
	if m != nil {
		return m.DestinationNetwork
	}
	return ""
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Ip) GetSourceNetwork() string {
	if m != nil {
		return m.SourceNetwork
	}
	return ""
}

type AccessLists_Acl_Rule_Match_IpRule_Icmp struct {
	// ICMPv6 flag, if false ICMPv4 will be used
	Icmpv6        bool                                                  `protobuf:"varint,1,opt,name=icmpv6,proto3" json:"icmpv6,omitempty"`
	IcmpCodeRange *AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpCodeRange `protobuf:"bytes,2,opt,name=icmp_code_range,json=icmpCodeRange" json:"icmp_code_range,omitempty"`
	IcmpTypeRange *AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpTypeRange `protobuf:"bytes,3,opt,name=icmp_type_range,json=icmpTypeRange" json:"icmp_type_range,omitempty"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp) Reset() {
	*m = AccessLists_Acl_Rule_Match_IpRule_Icmp{}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match_IpRule_Icmp) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match_IpRule_Icmp) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 1, 0, 1}
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp) GetIcmpv6() bool {
	if m != nil {
		return m.Icmpv6
	}
	return false
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp) GetIcmpCodeRange() *AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpCodeRange {
	if m != nil {
		return m.IcmpCodeRange
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp) GetIcmpTypeRange() *AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpTypeRange {
	if m != nil {
		return m.IcmpTypeRange
	}
	return nil
}

// Inclusive range representing icmp codes to be used.
type AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpCodeRange struct {
	// Lower boundary for range
	First uint32 `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	// Upper boundary for range
	Last uint32 `protobuf:"varint,2,opt,name=last,proto3" json:"last,omitempty"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpCodeRange) Reset() {
	*m = AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpCodeRange{}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpCodeRange) String() string {
	return proto.CompactTextString(m)
}
func (*AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpCodeRange) ProtoMessage() {}
func (*AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpCodeRange) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 1, 0, 1, 0}
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpCodeRange) GetFirst() uint32 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpCodeRange) GetLast() uint32 {
	if m != nil {
		return m.Last
	}
	return 0
}

type AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpTypeRange struct {
	// Lower boundary for range
	First uint32 `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	// Upper boundary for range
	Last uint32 `protobuf:"varint,2,opt,name=last,proto3" json:"last,omitempty"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpTypeRange) Reset() {
	*m = AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpTypeRange{}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpTypeRange) String() string {
	return proto.CompactTextString(m)
}
func (*AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpTypeRange) ProtoMessage() {}
func (*AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpTypeRange) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 1, 0, 1, 1}
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpTypeRange) GetFirst() uint32 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpTypeRange) GetLast() uint32 {
	if m != nil {
		return m.Last
	}
	return 0
}

type AccessLists_Acl_Rule_Match_IpRule_Tcp struct {
	DestinationPortRange *AccessLists_Acl_Rule_Match_IpRule_Tcp_DestinationPortRange `protobuf:"bytes,1,opt,name=destination_port_range,json=destinationPortRange" json:"destination_port_range,omitempty"`
	SourcePortRange      *AccessLists_Acl_Rule_Match_IpRule_Tcp_SourcePortRange      `protobuf:"bytes,2,opt,name=source_port_range,json=sourcePortRange" json:"source_port_range,omitempty"`
	// Binary mask for tcp flags to match. MSB order (FIN at position 0).
	// Applied as logical AND to tcp flags field of the packet being matched,
	// before it is compared with tcp-flags-value.
	TcpFlagsMask uint32 `protobuf:"varint,3,opt,name=tcp_flags_mask,json=tcpFlagsMask,proto3" json:"tcp_flags_mask,omitempty"`
	// Binary value for tcp flags to match. MSB order (FIN at position 0).
	// Before tcp-flags-value is compared with tcp flags field of the packet being matched,
	// tcp-flags-mask is applied to packet field value.
	TcpFlagsValue uint32 `protobuf:"varint,4,opt,name=tcp_flags_value,json=tcpFlagsValue,proto3" json:"tcp_flags_value,omitempty"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) Reset()         { *m = AccessLists_Acl_Rule_Match_IpRule_Tcp{} }
func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match_IpRule_Tcp) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match_IpRule_Tcp) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 1, 0, 2}
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) GetDestinationPortRange() *AccessLists_Acl_Rule_Match_IpRule_Tcp_DestinationPortRange {
	if m != nil {
		return m.DestinationPortRange
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) GetSourcePortRange() *AccessLists_Acl_Rule_Match_IpRule_Tcp_SourcePortRange {
	if m != nil {
		return m.SourcePortRange
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) GetTcpFlagsMask() uint32 {
	if m != nil {
		return m.TcpFlagsMask
	}
	return 0
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) GetTcpFlagsValue() uint32 {
	if m != nil {
		return m.TcpFlagsValue
	}
	return 0
}

// Inclusive range representing destination ports to be used. When
// only lower-port is present, it represents a single port.
type AccessLists_Acl_Rule_Match_IpRule_Tcp_DestinationPortRange struct {
	// Lower boundary for port.
	LowerPort uint32 `protobuf:"varint,1,opt,name=lower_port,json=lowerPort,proto3" json:"lower_port,omitempty"`
	// Upper boundary for port. If existing, the upper port must
	// be greater or equal to lower-port
	UpperPort uint32 `protobuf:"varint,2,opt,name=upper_port,json=upperPort,proto3" json:"upper_port,omitempty"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp_DestinationPortRange) Reset() {
	*m = AccessLists_Acl_Rule_Match_IpRule_Tcp_DestinationPortRange{}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp_DestinationPortRange) String() string {
	return proto.CompactTextString(m)
}
func (*AccessLists_Acl_Rule_Match_IpRule_Tcp_DestinationPortRange) ProtoMessage() {}
func (*AccessLists_Acl_Rule_Match_IpRule_Tcp_DestinationPortRange) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 1, 0, 2, 0}
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp_DestinationPortRange) GetLowerPort() uint32 {
	if m != nil {
		return m.LowerPort
	}
	return 0
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp_DestinationPortRange) GetUpperPort() uint32 {
	if m != nil {
		return m.UpperPort
	}
	return 0
}

// Inclusive range representing source ports to be used.
// When only lower-port is present, it represents a single port.
type AccessLists_Acl_Rule_Match_IpRule_Tcp_SourcePortRange struct {
	// Lower boundary for port.
	LowerPort uint32 `protobuf:"varint,1,opt,name=lower_port,json=lowerPort,proto3" json:"lower_port,omitempty"`
	// Upper boundary for port . If existing, the upper port
	// must be greater or equal to lower-port.
	UpperPort uint32 `protobuf:"varint,2,opt,name=upper_port,json=upperPort,proto3" json:"upper_port,omitempty"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp_SourcePortRange) Reset() {
	*m = AccessLists_Acl_Rule_Match_IpRule_Tcp_SourcePortRange{}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp_SourcePortRange) String() string {
	return proto.CompactTextString(m)
}
func (*AccessLists_Acl_Rule_Match_IpRule_Tcp_SourcePortRange) ProtoMessage() {}
func (*AccessLists_Acl_Rule_Match_IpRule_Tcp_SourcePortRange) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 1, 0, 2, 1}
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp_SourcePortRange) GetLowerPort() uint32 {
	if m != nil {
		return m.LowerPort
	}
	return 0
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp_SourcePortRange) GetUpperPort() uint32 {
	if m != nil {
		return m.UpperPort
	}
	return 0
}

type AccessLists_Acl_Rule_Match_IpRule_Udp struct {
	DestinationPortRange *AccessLists_Acl_Rule_Match_IpRule_Udp_DestinationPortRange `protobuf:"bytes,1,opt,name=destination_port_range,json=destinationPortRange" json:"destination_port_range,omitempty"`
	SourcePortRange      *AccessLists_Acl_Rule_Match_IpRule_Udp_SourcePortRange      `protobuf:"bytes,2,opt,name=source_port_range,json=sourcePortRange" json:"source_port_range,omitempty"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Udp) Reset()         { *m = AccessLists_Acl_Rule_Match_IpRule_Udp{} }
func (m *AccessLists_Acl_Rule_Match_IpRule_Udp) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match_IpRule_Udp) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match_IpRule_Udp) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 1, 0, 3}
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Udp) GetDestinationPortRange() *AccessLists_Acl_Rule_Match_IpRule_Udp_DestinationPortRange {
	if m != nil {
		return m.DestinationPortRange
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Udp) GetSourcePortRange() *AccessLists_Acl_Rule_Match_IpRule_Udp_SourcePortRange {
	if m != nil {
		return m.SourcePortRange
	}
	return nil
}

// Inclusive range representing destination ports to be used. When
// only lower-port is present, it represents a single port.
type AccessLists_Acl_Rule_Match_IpRule_Udp_DestinationPortRange struct {
	// Lower boundary for port.
	LowerPort uint32 `protobuf:"varint,1,opt,name=lower_port,json=lowerPort,proto3" json:"lower_port,omitempty"`
	// Upper boundary for port. If existing, the upper port must
	// be greater or equal to lower-port
	UpperPort uint32 `protobuf:"varint,2,opt,name=upper_port,json=upperPort,proto3" json:"upper_port,omitempty"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Udp_DestinationPortRange) Reset() {
	*m = AccessLists_Acl_Rule_Match_IpRule_Udp_DestinationPortRange{}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Udp_DestinationPortRange) String() string {
	return proto.CompactTextString(m)
}
func (*AccessLists_Acl_Rule_Match_IpRule_Udp_DestinationPortRange) ProtoMessage() {}
func (*AccessLists_Acl_Rule_Match_IpRule_Udp_DestinationPortRange) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 1, 0, 3, 0}
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Udp_DestinationPortRange) GetLowerPort() uint32 {
	if m != nil {
		return m.LowerPort
	}
	return 0
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Udp_DestinationPortRange) GetUpperPort() uint32 {
	if m != nil {
		return m.UpperPort
	}
	return 0
}

// Inclusive range representing source ports to be used.
// When only lower-port is present, it represents a single port.
type AccessLists_Acl_Rule_Match_IpRule_Udp_SourcePortRange struct {
	// Lower boundary for port.
	LowerPort uint32 `protobuf:"varint,1,opt,name=lower_port,json=lowerPort,proto3" json:"lower_port,omitempty"`
	// Upper boundary for port . If existing, the upper port
	// must be greater or equal to lower-port.
	UpperPort uint32 `protobuf:"varint,2,opt,name=upper_port,json=upperPort,proto3" json:"upper_port,omitempty"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Udp_SourcePortRange) Reset() {
	*m = AccessLists_Acl_Rule_Match_IpRule_Udp_SourcePortRange{}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Udp_SourcePortRange) String() string {
	return proto.CompactTextString(m)
}
func (*AccessLists_Acl_Rule_Match_IpRule_Udp_SourcePortRange) ProtoMessage() {}
func (*AccessLists_Acl_Rule_Match_IpRule_Udp_SourcePortRange) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 1, 0, 3, 1}
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Udp_SourcePortRange) GetLowerPort() uint32 {
	if m != nil {
		return m.LowerPort
	}
	return 0
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Udp_SourcePortRange) GetUpperPort() uint32 {
	if m != nil {
		return m.UpperPort
	}
	return 0
}

type AccessLists_Acl_Rule_Match_IpRule_Other struct {
	// Internet Protocol number.
	Protocol uint32 `protobuf:"varint,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Other) Reset() {
	*m = AccessLists_Acl_Rule_Match_IpRule_Other{}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Other) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match_IpRule_Other) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match_IpRule_Other) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 1, 0, 4}
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Other) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

type AccessLists_Acl_Rule_Match_MacIpRule struct {
	// Source IPv4/Ipv6 address
	SourceAddress string `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	// Source IPv4/Ipv6 address prefix
	SourceAddressPrefix uint32 `protobuf:"varint,2,opt,name=source_address_prefix,json=sourceAddressPrefix,proto3" json:"source_address_prefix,omitempty"`
	// Source IEEE 802 MAC address.
	// Before source-mac-address is compared with source mac address field of the packet being matched,
	// source-mac-address-mask is applied to packet field value.
	SourceMacAddress string `protobuf:"bytes,3,opt,name=source_mac_address,json=sourceMacAddress,proto3" json:"source_mac_address,omitempty"`
	// Source IEEE 802 MAC address mask.
	// Applied as logical AND with source mac address field of the packet being matched,
	// before it is compared with source-mac-address.
	SourceMacAddressMask string `protobuf:"bytes,4,opt,name=source_mac_address_mask,json=sourceMacAddressMask,proto3" json:"source_mac_address_mask,omitempty"`
}

func (m *AccessLists_Acl_Rule_Match_MacIpRule) Reset()         { *m = AccessLists_Acl_Rule_Match_MacIpRule{} }
func (m *AccessLists_Acl_Rule_Match_MacIpRule) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match_MacIpRule) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match_MacIpRule) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 0, 1, 1}
}

func (m *AccessLists_Acl_Rule_Match_MacIpRule) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *AccessLists_Acl_Rule_Match_MacIpRule) GetSourceAddressPrefix() uint32 {
	if m != nil {
		return m.SourceAddressPrefix
	}
	return 0
}

func (m *AccessLists_Acl_Rule_Match_MacIpRule) GetSourceMacAddress() string {
	if m != nil {
		return m.SourceMacAddress
	}
	return ""
}

func (m *AccessLists_Acl_Rule_Match_MacIpRule) GetSourceMacAddressMask() string {
	if m != nil {
		return m.SourceMacAddressMask
	}
	return ""
}

type AccessLists_Acl_Interface struct {
	Egress  []string `protobuf:"bytes,1,rep,name=egress" json:"egress,omitempty"`
	Ingress []string `protobuf:"bytes,2,rep,name=ingress" json:"ingress,omitempty"`
}

func (m *AccessLists_Acl_Interface) Reset()         { *m = AccessLists_Acl_Interface{} }
func (m *AccessLists_Acl_Interface) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Interface) ProtoMessage()    {}
func (*AccessLists_Acl_Interface) Descriptor() ([]byte, []int) {
	return fileDescriptorAcl, []int{0, 0, 1}
}

func (m *AccessLists_Acl_Interface) GetEgress() []string {
	if m != nil {
		return m.Egress
	}
	return nil
}

func (m *AccessLists_Acl_Interface) GetIngress() []string {
	if m != nil {
		return m.Ingress
	}
	return nil
}

func init() {
	proto.RegisterType((*AccessLists)(nil), "acl.AccessLists")
	proto.RegisterType((*AccessLists_Acl)(nil), "acl.AccessLists.Acl")
	proto.RegisterType((*AccessLists_Acl_Rule)(nil), "acl.AccessLists.Acl.Rule")
	proto.RegisterType((*AccessLists_Acl_Rule_Action)(nil), "acl.AccessLists.Acl.Rule.Action")
	proto.RegisterType((*AccessLists_Acl_Rule_Match)(nil), "acl.AccessLists.Acl.Rule.Match")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Ip)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Ip")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Icmp)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Icmp")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpCodeRange)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Icmp.IcmpCodeRange")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Icmp_IcmpTypeRange)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Icmp.IcmpTypeRange")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Tcp)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Tcp")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Tcp_DestinationPortRange)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Tcp.DestinationPortRange")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Tcp_SourcePortRange)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Tcp.SourcePortRange")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Udp)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Udp")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Udp_DestinationPortRange)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Udp.DestinationPortRange")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Udp_SourcePortRange)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Udp.SourcePortRange")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Other)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Other")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_MacIpRule)(nil), "acl.AccessLists.Acl.Rule.Match.MacIpRule")
	proto.RegisterType((*AccessLists_Acl_Interface)(nil), "acl.AccessLists.Acl.Interface")
	proto.RegisterEnum("acl.AclAction", AclAction_name, AclAction_value)
}

func init() { proto.RegisterFile("acl.proto", fileDescriptorAcl) }

var fileDescriptorAcl = []byte{
	// 885 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xc6, 0x5e, 0xff, 0xed, 0x31, 0x1b, 0xa7, 0x27, 0xa6, 0xb8, 0x8b, 0x00, 0x8b, 0x9f, 0xc8,
	0x2d, 0xc5, 0x95, 0x8c, 0x00, 0xa5, 0x02, 0x22, 0xd3, 0xa6, 0x22, 0x52, 0x9d, 0x46, 0x83, 0x83,
	0x84, 0x84, 0x64, 0x0d, 0xb3, 0xe3, 0x74, 0x95, 0xb5, 0x77, 0xb4, 0x33, 0x6e, 0xe9, 0x33, 0x20,
	0xf1, 0x00, 0xbc, 0x10, 0x17, 0xdc, 0x71, 0x07, 0x6f, 0xc1, 0x1b, 0xa0, 0xf9, 0xd9, 0xcd, 0xa6,
	0xb8, 0x8a, 0x03, 0x08, 0xa9, 0x37, 0xde, 0x99, 0x73, 0xbe, 0xef, 0x3b, 0x7f, 0xbb, 0xe3, 0x01,
	0x9f, 0xb2, 0x64, 0x28, 0xb2, 0x54, 0xa5, 0xe8, 0x51, 0x96, 0xbc, 0xf3, 0xdb, 0x0e, 0xb4, 0xc7,
	0x8c, 0x71, 0x29, 0x1f, 0xc6, 0x52, 0x49, 0xdc, 0x05, 0x6d, 0xee, 0x55, 0xfa, 0xde, 0xa0, 0x3d,
	0xea, 0x0e, 0x35, 0xba, 0xe4, 0x1e, 0x8e, 0x59, 0x42, 0x34, 0x20, 0xfc, 0x79, 0x07, 0xbc, 0x31,
	0x4b, 0xf0, 0x0e, 0xd4, 0xb3, 0x55, 0xc2, 0xa5, 0x63, 0xdc, 0x58, 0xc7, 0x18, 0x92, 0x55, 0xc2,
	0x89, 0xc5, 0xe1, 0x0d, 0x68, 0x51, 0x96, 0xcc, 0x96, 0x74, 0xc1, 0x7b, 0x5e, 0xbf, 0x32, 0xf0,
	0x49, 0x93, 0xb2, 0xe4, 0x88, 0x2e, 0x38, 0x7e, 0x01, 0x10, 0x2f, 0x15, 0xcf, 0xe6, 0x94, 0x71,
	0xd9, 0xab, 0xf6, 0x2b, 0x83, 0xf6, 0xe8, 0xad, 0xb5, 0x82, 0x87, 0x39, 0x8c, 0x94, 0x18, 0xe1,
	0x2f, 0xd7, 0xa0, 0xa6, 0x43, 0xe1, 0x5d, 0x68, 0x52, 0xa6, 0xe2, 0x74, 0x99, 0xab, 0xf4, 0x5f,
	0x98, 0xd6, 0x70, 0x6c, 0x80, 0x24, 0x27, 0xe0, 0x1e, 0x34, 0x17, 0x54, 0xb1, 0xc7, 0x5c, 0x9a,
	0xf4, 0xda, 0xa3, 0xb7, 0x5f, 0xcc, 0x9d, 0x68, 0x20, 0xc9, 0xf1, 0xf8, 0x06, 0xf8, 0xba, 0x46,
	0x5b, 0x5b, 0xdd, 0xd4, 0xd6, 0xd2, 0x06, 0x5d, 0x5c, 0xf8, 0x29, 0x34, 0x6c, 0x28, 0xfc, 0x10,
	0x40, 0x77, 0xc0, 0x06, 0xec, 0x55, 0xfa, 0x95, 0xc1, 0xd6, 0x68, 0xcb, 0x05, 0x49, 0x5c, 0x3a,
	0x7a, 0x4c, 0x76, 0x19, 0xfe, 0xd9, 0x81, 0xba, 0x09, 0x84, 0xfb, 0xd0, 0x8c, 0xc5, 0x4c, 0x2b,
	0x1a, 0x56, 0x7b, 0xb4, 0x7b, 0x49, 0x6a, 0xc3, 0x43, 0x61, 0x5a, 0xdf, 0x88, 0xcd, 0x13, 0xbf,
	0x02, 0x58, 0x50, 0x96, 0x6b, 0xd4, 0x8c, 0xc6, 0xcd, 0xcb, 0x34, 0x26, 0x94, 0x39, 0x19, 0xdf,
	0x90, 0xf5, 0x32, 0xfc, 0x29, 0x80, 0x86, 0xb5, 0xe2, 0x1e, 0x54, 0x63, 0xe1, 0x12, 0xba, 0xb9,
	0x59, 0x42, 0xfa, 0x51, 0x8d, 0x05, 0xee, 0x43, 0x2d, 0x66, 0x0b, 0xe1, 0x86, 0xf4, 0xc1, 0xa6,
	0x64, 0xb6, 0x10, 0xc4, 0x10, 0xf1, 0x33, 0xf0, 0x14, 0x13, 0xae, 0x92, 0x5b, 0x1b, 0xf2, 0xa7,
	0x4c, 0x10, 0x4d, 0xd3, 0xec, 0x55, 0x24, 0xcc, 0xa4, 0x36, 0x67, 0x9f, 0x44, 0x82, 0x68, 0x1a,
	0x7e, 0x09, 0xf5, 0x54, 0x3d, 0xe6, 0x99, 0x7b, 0x4d, 0x6e, 0x6f, 0xc8, 0x7f, 0xa4, 0x39, 0xc4,
	0x52, 0xc3, 0xef, 0xa0, 0x7a, 0x28, 0xf0, 0x0e, 0xec, 0x44, 0x5c, 0xaa, 0x78, 0x49, 0xf5, 0xc0,
	0x67, 0x4b, 0xae, 0x9e, 0xa6, 0xd9, 0x99, 0x69, 0xa9, 0x4f, 0xb0, 0xe4, 0x3a, 0xb2, 0x1e, 0x7c,
	0x1f, 0xb6, 0x64, 0xba, 0xca, 0x18, 0x2f, 0xb0, 0x55, 0x83, 0x0d, 0xac, 0xd5, 0xc1, 0xc2, 0x3f,
	0xaa, 0x50, 0xd3, 0xcd, 0xc2, 0xeb, 0xd0, 0xd0, 0xed, 0x7a, 0xf2, 0x89, 0xd1, 0x6c, 0x11, 0xb7,
	0x43, 0x0a, 0x1d, 0xbd, 0x9a, 0xb1, 0x34, 0xe2, 0xb3, 0x8c, 0x2e, 0x4f, 0xb9, 0x1b, 0xc5, 0xde,
	0x15, 0x46, 0x61, 0x7e, 0xee, 0xa5, 0x11, 0x27, 0x5a, 0x80, 0x04, 0x71, 0x79, 0x5b, 0x84, 0x50,
	0xcf, 0x44, 0x1e, 0xc2, 0xfb, 0x67, 0x21, 0xa6, 0xcf, 0x44, 0x39, 0x44, 0xb1, 0x0d, 0xf7, 0x20,
	0xb8, 0x90, 0x02, 0x76, 0xa1, 0x3e, 0x8f, 0x33, 0xa9, 0x4c, 0xb5, 0x01, 0xb1, 0x1b, 0x44, 0xa8,
	0x25, 0x54, 0x2a, 0x53, 0x61, 0x40, 0xcc, 0x3a, 0xa7, 0x16, 0x5a, 0x57, 0xa0, 0xfe, 0xee, 0x81,
	0x37, 0x65, 0x02, 0x57, 0x70, 0xbd, 0x3c, 0x3c, 0x91, 0x66, 0xca, 0xd5, 0x69, 0x3f, 0x89, 0xfd,
	0xcd, 0xdf, 0xca, 0xe1, 0xfd, 0x73, 0xa1, 0xe3, 0x34, 0x53, 0xb6, 0xda, 0x6e, 0xb4, 0xc6, 0x8a,
	0x73, 0xb8, 0xe6, 0x5e, 0x81, 0x52, 0x44, 0x3b, 0xbc, 0xbb, 0x57, 0x88, 0xf8, 0xb5, 0xd1, 0x38,
	0x0f, 0xd6, 0x91, 0x17, 0x0d, 0xf8, 0x1e, 0x6c, 0x29, 0x26, 0x66, 0xf3, 0x84, 0x9e, 0xca, 0xd9,
	0x82, 0xca, 0x33, 0x33, 0xbe, 0x80, 0xbc, 0xaa, 0x98, 0x78, 0xa0, 0x8d, 0x13, 0x2a, 0xcf, 0x70,
	0x17, 0x3a, 0xe7, 0xa8, 0x27, 0x34, 0x59, 0xd9, 0xd3, 0x25, 0x20, 0x41, 0x0e, 0xfb, 0x46, 0x1b,
	0xc3, 0x29, 0x74, 0xd7, 0xd5, 0x88, 0x6f, 0x02, 0x24, 0xe9, 0x53, 0x9e, 0x99, 0x62, 0x5c, 0xef,
	0x7d, 0x63, 0xd1, 0x18, 0xed, 0x5e, 0x09, 0x91, 0xbb, 0xed, 0x14, 0x7c, 0x63, 0xd1, 0xee, 0xf0,
	0x11, 0x74, 0x9e, 0xab, 0xe3, 0x5f, 0x0a, 0xfe, 0xe8, 0x81, 0x77, 0x12, 0xfd, 0x77, 0xb3, 0x3d,
	0x89, 0xfe, 0xef, 0xd9, 0xea, 0x88, 0x97, 0xcd, 0xf6, 0x65, 0x99, 0xc6, 0xbb, 0x50, 0x37, 0x87,
	0x26, 0x86, 0xd0, 0x32, 0x37, 0x17, 0x96, 0x26, 0x4e, 0xa4, 0xd8, 0x87, 0xbf, 0x56, 0xc0, 0x2f,
	0xfe, 0xa9, 0x4a, 0x07, 0x24, 0x8d, 0xa2, 0x8c, 0x4b, 0xe9, 0x0e, 0x53, 0x77, 0x40, 0x8e, 0xad,
	0x11, 0x47, 0xf0, 0xda, 0x45, 0xd8, 0x4c, 0x64, 0x7c, 0x1e, 0xff, 0xe0, 0x72, 0xd8, 0xb9, 0x80,
	0x3e, 0x36, 0x2e, 0xbc, 0x0d, 0xe8, 0x38, 0x0b, 0xca, 0x0a, 0x79, 0x7b, 0x93, 0xd9, 0xb6, 0x9e,
	0x09, 0x65, 0x79, 0x84, 0x8f, 0xe1, 0xf5, 0xbf, 0xa3, 0xed, 0x77, 0x54, 0x33, 0x94, 0xee, 0xf3,
	0x14, 0xfd, 0x3d, 0x85, 0x9f, 0x83, 0x5f, 0x5c, 0x71, 0xf4, 0xe9, 0xcd, 0x4f, 0x5d, 0x11, 0xde,
	0xc0, 0x27, 0x6e, 0x87, 0x3d, 0x68, 0xc6, 0x4b, 0xeb, 0xa8, 0x1a, 0x47, 0xbe, 0xbd, 0x35, 0x04,
	0xbf, 0xb8, 0x4a, 0x60, 0x0b, 0x6a, 0xf7, 0x0f, 0x8e, 0xbe, 0xdd, 0x7e, 0x05, 0x01, 0x1a, 0xc7,
	0x07, 0x64, 0x72, 0x38, 0xdd, 0xae, 0x60, 0x1b, 0x9a, 0xe4, 0xe0, 0xc1, 0xc3, 0x83, 0x7b, 0xd3,
	0xed, 0xea, 0xf7, 0x0d, 0xd3, 0xc6, 0x8f, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x86, 0xf0,
	0xd9, 0x1d, 0x0a, 0x00, 0x00,
}
