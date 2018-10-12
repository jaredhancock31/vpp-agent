//  Copyright (c) 2018 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package vppcalls

import (
	govppapi "git.fd.io/govpp.git/api"
	"github.com/ligato/vpp-agent/plugins/vppv2/ifplugin/ifaceidx"
	"github.com/ligato/vpp-agent/plugins/vppv2/model/acl"
)

// ACLVppAPI provides read/write methods required to handle VPP access lists
type ACLVppAPI interface {
	ACLVppWrite
	ACLVppRead
}

// ACLVppWrite provides write methods for ACL plugin
type ACLVppWrite interface {
	// AddACL create new ACL (L3/L4). Returns ACL index provided by VPP.
	AddACL(rules []*acl.Acl_Rule, aclName string) (aclIdx uint32, err error)
	// AddMACIPACL creates new MACIP ACL (L2). Returns ACL index provided by VPP.
	AddMACIPACL(rules []*acl.Acl_Rule, aclName string) (aclIdx uint32, err error)
	// ModifyACL modifies ACL (L3/L4) by updating its rules. It uses ACL index to identify ACL.
	ModifyACL(aclIdx uint32, rules []*acl.Acl_Rule, aclName string) error
	// ModifyACL modifies MACIP ACL (L2) by updating its rules. It uses ACL index to identify ACL.
	ModifyMACIPACL(aclIdx uint32, rules []*acl.Acl_Rule, aclName string) error
	// DeleteACL removes ACL (L3/L4).
	DeleteACL(aclIdx uint32) error
	// DeleteMACIPACL removes MACIP ACL (L2).
	DeleteMACIPACL(aclIdx uint32) error
	// SetACLToInterfacesAsIngress sets ACL to interfaces as ingress.
	SetACLToInterfacesAsIngress(ACLIndex uint32, ifIndices []uint32) error
	// RemoveACLFromInterfacesAsIngress removes ACL from interfaces as ingress.
	RemoveACLFromInterfacesAsIngress(ACLIndex uint32, ifIndices []uint32) error
	// SetACLToInterfacesAsEgress sets ACL to interfaces as egress.
	SetACLToInterfacesAsEgress(ACLIndex uint32, ifIndices []uint32) error
	// RemoveACLFromInterfacesAsEgress removes ACL from interfaces as egress.
	RemoveACLFromInterfacesAsEgress(ACLIndex uint32, ifIndices []uint32) error
	// SetMACIPACLToInterfaces sets MACIP ACL to interfaces.
	SetMACIPACLToInterfaces(aclIndex uint32, ifIndices []uint32) error
	// RemoveMACIPACLFromInterfaces removes MACIP ACL from interfaces.
	RemoveMACIPACLFromInterfaces(removedACLIndex uint32, ifIndices []uint32) error

	AddACLToInterfaceAsIngress(aclIndex uint32, ifName string) error
	AddACLToInterfaceAsEgress(aclIndex uint32, ifName string) error
	DeleteACLFromInterfaceAsIngress(aclIndex uint32, ifName string) error
	DeleteACLFromInterfaceAsEgress(aclIndex uint32, ifName string) error
	AddMACIPACLToInterface(aclIndex uint32, ifName string) error
	DeleteMACIPACLFromInterface(aclIndex uint32, ifName string) error
}

// ACLVppRead provides read methods for ACL plugin
type ACLVppRead interface {
	// DumpACL dumps all ACLs (L3/L4).
	DumpACL() ([]*ACLDetails, error)
	// DumpMACIPACL dumps all MACIP ACLs (L2).
	DumpMACIPACL() ([]*ACLDetails, error)
	// DumpACLInterfaces dumps all ACLs (L3/L4) for given ACL indexes. Returns map of ACL indexes with assigned interfaces.
	DumpACLInterfaces(indices []uint32) (map[uint32]*acl.Acl_Interfaces, error)
	// DumpMACIPACLInterfaces dumps all ACLs (L2) for given ACL indexes. Returns map of MACIP ACL indexes with assigned interfaces.
	DumpMACIPACLInterfaces(indices []uint32) (map[uint32]*acl.Acl_Interfaces, error)
	// DumpInterfaceAcls finds interface in VPP and returns its ACL (L3/L4) configuration.
	DumpInterfaceACLs(ifIdx uint32) ([]*acl.Acl, error)
	// DumpInterfaceMACIPACLs finds interface in VPP and returns its MACIP ACL (L2) configuration.
	DumpInterfaceMACIPACLs(ifIdx uint32) ([]*acl.Acl, error)
}

// ACLVppHandler is accessor for acl-related vppcalls methods
type ACLVppHandler struct {
	callsChannel govppapi.Channel
	dumpChannel  govppapi.Channel
	ifIndexes    ifaceidx.IfaceMetadataIndex
}

// NewACLVppHandler creates new instance of acl vppcalls handler
func NewACLVppHandler(callsChan, dumpChan govppapi.Channel, ifIndexes ifaceidx.IfaceMetadataIndex) *ACLVppHandler {
	return &ACLVppHandler{
		callsChannel: callsChan,
		dumpChannel:  dumpChan,
		ifIndexes:    ifIndexes,
	}
}
