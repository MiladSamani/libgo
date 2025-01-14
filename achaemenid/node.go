/* For license and copyright information please see LEGAL file in repository */

package achaemenid

// Node store details about a node that part of the platfrom.
type Node struct {
	InstanceID      [32]byte
	ID              uint64
	dataCenterID    [32]byte
	storageCapacity uint64 // In bytes, Max 16EB(Exabyte) that more enough for one node capacity. 0 means service only node.
	Conn            *Connection
	State           nodeState
}

type nodeState uint8

// Node State
const (
	NodeStateLocalNode nodeState = iota
	NodeStateStable
	NodeStateStop
	NodeStateStoping
	NodeStateStart
	NodeStateStarting
	NodeStateNotResponse
)

// NodeDetails ...
type NodeDetails struct {
	ID uint64
	// GPAddr gp.Addr
	// IPAddr ip.Addr
}
