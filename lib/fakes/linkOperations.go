// This file was generated by counterfeiter
package fakes

import (
	"net"
	"sync"

	"github.com/vishvananda/netlink"
)

type LinkOperations struct {
	DisableIPv6Stub        func(deviceName string) error
	disableIPv6Mutex       sync.RWMutex
	disableIPv6ArgsForCall []struct {
		deviceName string
	}
	disableIPv6Returns struct {
		result1 error
	}
	disableIPv6ReturnsOnCall map[int]struct {
		result1 error
	}
	StaticNeighborNoARPStub        func(link netlink.Link, dstIP net.IP, mac net.HardwareAddr) error
	staticNeighborNoARPMutex       sync.RWMutex
	staticNeighborNoARPArgsForCall []struct {
		link  netlink.Link
		dstIP net.IP
		mac   net.HardwareAddr
	}
	staticNeighborNoARPReturns struct {
		result1 error
	}
	staticNeighborNoARPReturnsOnCall map[int]struct {
		result1 error
	}
	SetPointToPointAddressStub        func(link netlink.Link, localIPAddr, peerIPAddr net.IP) error
	setPointToPointAddressMutex       sync.RWMutex
	setPointToPointAddressArgsForCall []struct {
		link        netlink.Link
		localIPAddr net.IP
		peerIPAddr  net.IP
	}
	setPointToPointAddressReturns struct {
		result1 error
	}
	setPointToPointAddressReturnsOnCall map[int]struct {
		result1 error
	}
	RenameLinkStub        func(oldName, newName string) error
	renameLinkMutex       sync.RWMutex
	renameLinkArgsForCall []struct {
		oldName string
		newName string
	}
	renameLinkReturns struct {
		result1 error
	}
	renameLinkReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteLinkByNameStub        func(deviceName string) error
	deleteLinkByNameMutex       sync.RWMutex
	deleteLinkByNameArgsForCall []struct {
		deviceName string
	}
	deleteLinkByNameReturns struct {
		result1 error
	}
	deleteLinkByNameReturnsOnCall map[int]struct {
		result1 error
	}
	RouteAddStub        func(route netlink.Route) error
	routeAddMutex       sync.RWMutex
	routeAddArgsForCall []struct {
		route netlink.Route
	}
	routeAddReturns struct {
		result1 error
	}
	routeAddReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *LinkOperations) DisableIPv6(deviceName string) error {
	fake.disableIPv6Mutex.Lock()
	ret, specificReturn := fake.disableIPv6ReturnsOnCall[len(fake.disableIPv6ArgsForCall)]
	fake.disableIPv6ArgsForCall = append(fake.disableIPv6ArgsForCall, struct {
		deviceName string
	}{deviceName})
	fake.recordInvocation("DisableIPv6", []interface{}{deviceName})
	fake.disableIPv6Mutex.Unlock()
	if fake.DisableIPv6Stub != nil {
		return fake.DisableIPv6Stub(deviceName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.disableIPv6Returns.result1
}

func (fake *LinkOperations) DisableIPv6CallCount() int {
	fake.disableIPv6Mutex.RLock()
	defer fake.disableIPv6Mutex.RUnlock()
	return len(fake.disableIPv6ArgsForCall)
}

func (fake *LinkOperations) DisableIPv6ArgsForCall(i int) string {
	fake.disableIPv6Mutex.RLock()
	defer fake.disableIPv6Mutex.RUnlock()
	return fake.disableIPv6ArgsForCall[i].deviceName
}

func (fake *LinkOperations) DisableIPv6Returns(result1 error) {
	fake.DisableIPv6Stub = nil
	fake.disableIPv6Returns = struct {
		result1 error
	}{result1}
}

func (fake *LinkOperations) DisableIPv6ReturnsOnCall(i int, result1 error) {
	fake.DisableIPv6Stub = nil
	if fake.disableIPv6ReturnsOnCall == nil {
		fake.disableIPv6ReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.disableIPv6ReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *LinkOperations) StaticNeighborNoARP(link netlink.Link, dstIP net.IP, mac net.HardwareAddr) error {
	fake.staticNeighborNoARPMutex.Lock()
	ret, specificReturn := fake.staticNeighborNoARPReturnsOnCall[len(fake.staticNeighborNoARPArgsForCall)]
	fake.staticNeighborNoARPArgsForCall = append(fake.staticNeighborNoARPArgsForCall, struct {
		link  netlink.Link
		dstIP net.IP
		mac   net.HardwareAddr
	}{link, dstIP, mac})
	fake.recordInvocation("StaticNeighborNoARP", []interface{}{link, dstIP, mac})
	fake.staticNeighborNoARPMutex.Unlock()
	if fake.StaticNeighborNoARPStub != nil {
		return fake.StaticNeighborNoARPStub(link, dstIP, mac)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.staticNeighborNoARPReturns.result1
}

func (fake *LinkOperations) StaticNeighborNoARPCallCount() int {
	fake.staticNeighborNoARPMutex.RLock()
	defer fake.staticNeighborNoARPMutex.RUnlock()
	return len(fake.staticNeighborNoARPArgsForCall)
}

func (fake *LinkOperations) StaticNeighborNoARPArgsForCall(i int) (netlink.Link, net.IP, net.HardwareAddr) {
	fake.staticNeighborNoARPMutex.RLock()
	defer fake.staticNeighborNoARPMutex.RUnlock()
	return fake.staticNeighborNoARPArgsForCall[i].link, fake.staticNeighborNoARPArgsForCall[i].dstIP, fake.staticNeighborNoARPArgsForCall[i].mac
}

func (fake *LinkOperations) StaticNeighborNoARPReturns(result1 error) {
	fake.StaticNeighborNoARPStub = nil
	fake.staticNeighborNoARPReturns = struct {
		result1 error
	}{result1}
}

func (fake *LinkOperations) StaticNeighborNoARPReturnsOnCall(i int, result1 error) {
	fake.StaticNeighborNoARPStub = nil
	if fake.staticNeighborNoARPReturnsOnCall == nil {
		fake.staticNeighborNoARPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.staticNeighborNoARPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *LinkOperations) SetPointToPointAddress(link netlink.Link, localIPAddr net.IP, peerIPAddr net.IP) error {
	fake.setPointToPointAddressMutex.Lock()
	ret, specificReturn := fake.setPointToPointAddressReturnsOnCall[len(fake.setPointToPointAddressArgsForCall)]
	fake.setPointToPointAddressArgsForCall = append(fake.setPointToPointAddressArgsForCall, struct {
		link        netlink.Link
		localIPAddr net.IP
		peerIPAddr  net.IP
	}{link, localIPAddr, peerIPAddr})
	fake.recordInvocation("SetPointToPointAddress", []interface{}{link, localIPAddr, peerIPAddr})
	fake.setPointToPointAddressMutex.Unlock()
	if fake.SetPointToPointAddressStub != nil {
		return fake.SetPointToPointAddressStub(link, localIPAddr, peerIPAddr)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setPointToPointAddressReturns.result1
}

func (fake *LinkOperations) SetPointToPointAddressCallCount() int {
	fake.setPointToPointAddressMutex.RLock()
	defer fake.setPointToPointAddressMutex.RUnlock()
	return len(fake.setPointToPointAddressArgsForCall)
}

func (fake *LinkOperations) SetPointToPointAddressArgsForCall(i int) (netlink.Link, net.IP, net.IP) {
	fake.setPointToPointAddressMutex.RLock()
	defer fake.setPointToPointAddressMutex.RUnlock()
	return fake.setPointToPointAddressArgsForCall[i].link, fake.setPointToPointAddressArgsForCall[i].localIPAddr, fake.setPointToPointAddressArgsForCall[i].peerIPAddr
}

func (fake *LinkOperations) SetPointToPointAddressReturns(result1 error) {
	fake.SetPointToPointAddressStub = nil
	fake.setPointToPointAddressReturns = struct {
		result1 error
	}{result1}
}

func (fake *LinkOperations) SetPointToPointAddressReturnsOnCall(i int, result1 error) {
	fake.SetPointToPointAddressStub = nil
	if fake.setPointToPointAddressReturnsOnCall == nil {
		fake.setPointToPointAddressReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPointToPointAddressReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *LinkOperations) RenameLink(oldName string, newName string) error {
	fake.renameLinkMutex.Lock()
	ret, specificReturn := fake.renameLinkReturnsOnCall[len(fake.renameLinkArgsForCall)]
	fake.renameLinkArgsForCall = append(fake.renameLinkArgsForCall, struct {
		oldName string
		newName string
	}{oldName, newName})
	fake.recordInvocation("RenameLink", []interface{}{oldName, newName})
	fake.renameLinkMutex.Unlock()
	if fake.RenameLinkStub != nil {
		return fake.RenameLinkStub(oldName, newName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.renameLinkReturns.result1
}

func (fake *LinkOperations) RenameLinkCallCount() int {
	fake.renameLinkMutex.RLock()
	defer fake.renameLinkMutex.RUnlock()
	return len(fake.renameLinkArgsForCall)
}

func (fake *LinkOperations) RenameLinkArgsForCall(i int) (string, string) {
	fake.renameLinkMutex.RLock()
	defer fake.renameLinkMutex.RUnlock()
	return fake.renameLinkArgsForCall[i].oldName, fake.renameLinkArgsForCall[i].newName
}

func (fake *LinkOperations) RenameLinkReturns(result1 error) {
	fake.RenameLinkStub = nil
	fake.renameLinkReturns = struct {
		result1 error
	}{result1}
}

func (fake *LinkOperations) RenameLinkReturnsOnCall(i int, result1 error) {
	fake.RenameLinkStub = nil
	if fake.renameLinkReturnsOnCall == nil {
		fake.renameLinkReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.renameLinkReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *LinkOperations) DeleteLinkByName(deviceName string) error {
	fake.deleteLinkByNameMutex.Lock()
	ret, specificReturn := fake.deleteLinkByNameReturnsOnCall[len(fake.deleteLinkByNameArgsForCall)]
	fake.deleteLinkByNameArgsForCall = append(fake.deleteLinkByNameArgsForCall, struct {
		deviceName string
	}{deviceName})
	fake.recordInvocation("DeleteLinkByName", []interface{}{deviceName})
	fake.deleteLinkByNameMutex.Unlock()
	if fake.DeleteLinkByNameStub != nil {
		return fake.DeleteLinkByNameStub(deviceName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteLinkByNameReturns.result1
}

func (fake *LinkOperations) DeleteLinkByNameCallCount() int {
	fake.deleteLinkByNameMutex.RLock()
	defer fake.deleteLinkByNameMutex.RUnlock()
	return len(fake.deleteLinkByNameArgsForCall)
}

func (fake *LinkOperations) DeleteLinkByNameArgsForCall(i int) string {
	fake.deleteLinkByNameMutex.RLock()
	defer fake.deleteLinkByNameMutex.RUnlock()
	return fake.deleteLinkByNameArgsForCall[i].deviceName
}

func (fake *LinkOperations) DeleteLinkByNameReturns(result1 error) {
	fake.DeleteLinkByNameStub = nil
	fake.deleteLinkByNameReturns = struct {
		result1 error
	}{result1}
}

func (fake *LinkOperations) DeleteLinkByNameReturnsOnCall(i int, result1 error) {
	fake.DeleteLinkByNameStub = nil
	if fake.deleteLinkByNameReturnsOnCall == nil {
		fake.deleteLinkByNameReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteLinkByNameReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *LinkOperations) RouteAdd(route netlink.Route) error {
	fake.routeAddMutex.Lock()
	ret, specificReturn := fake.routeAddReturnsOnCall[len(fake.routeAddArgsForCall)]
	fake.routeAddArgsForCall = append(fake.routeAddArgsForCall, struct {
		route netlink.Route
	}{route})
	fake.recordInvocation("RouteAdd", []interface{}{route})
	fake.routeAddMutex.Unlock()
	if fake.RouteAddStub != nil {
		return fake.RouteAddStub(route)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.routeAddReturns.result1
}

func (fake *LinkOperations) RouteAddCallCount() int {
	fake.routeAddMutex.RLock()
	defer fake.routeAddMutex.RUnlock()
	return len(fake.routeAddArgsForCall)
}

func (fake *LinkOperations) RouteAddArgsForCall(i int) netlink.Route {
	fake.routeAddMutex.RLock()
	defer fake.routeAddMutex.RUnlock()
	return fake.routeAddArgsForCall[i].route
}

func (fake *LinkOperations) RouteAddReturns(result1 error) {
	fake.RouteAddStub = nil
	fake.routeAddReturns = struct {
		result1 error
	}{result1}
}

func (fake *LinkOperations) RouteAddReturnsOnCall(i int, result1 error) {
	fake.RouteAddStub = nil
	if fake.routeAddReturnsOnCall == nil {
		fake.routeAddReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.routeAddReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *LinkOperations) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.disableIPv6Mutex.RLock()
	defer fake.disableIPv6Mutex.RUnlock()
	fake.staticNeighborNoARPMutex.RLock()
	defer fake.staticNeighborNoARPMutex.RUnlock()
	fake.setPointToPointAddressMutex.RLock()
	defer fake.setPointToPointAddressMutex.RUnlock()
	fake.renameLinkMutex.RLock()
	defer fake.renameLinkMutex.RUnlock()
	fake.deleteLinkByNameMutex.RLock()
	defer fake.deleteLinkByNameMutex.RUnlock()
	fake.routeAddMutex.RLock()
	defer fake.routeAddMutex.RUnlock()
	return fake.invocations
}

func (fake *LinkOperations) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}