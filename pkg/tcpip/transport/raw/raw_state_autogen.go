// automatically generated by stateify.

package raw

import (
	"gvisor.dev/gvisor/pkg/state"
	"gvisor.dev/gvisor/pkg/tcpip/buffer"
)

func (x *packet) beforeSave() {}
func (x *packet) save(m state.Map) {
	x.beforeSave()
	var data buffer.VectorisedView = x.saveData()
	m.SaveValue("data", data)
	m.Save("packetEntry", &x.packetEntry)
	m.Save("timestampNS", &x.timestampNS)
	m.Save("senderAddr", &x.senderAddr)
}

func (x *packet) afterLoad() {}
func (x *packet) load(m state.Map) {
	m.Load("packetEntry", &x.packetEntry)
	m.Load("timestampNS", &x.timestampNS)
	m.Load("senderAddr", &x.senderAddr)
	m.LoadValue("data", new(buffer.VectorisedView), func(y interface{}) { x.loadData(y.(buffer.VectorisedView)) })
}

func (x *endpoint) save(m state.Map) {
	x.beforeSave()
	var rcvBufSizeMax int = x.saveRcvBufSizeMax()
	m.SaveValue("rcvBufSizeMax", rcvBufSizeMax)
	m.Save("netProto", &x.netProto)
	m.Save("transProto", &x.transProto)
	m.Save("waiterQueue", &x.waiterQueue)
	m.Save("associated", &x.associated)
	m.Save("rcvList", &x.rcvList)
	m.Save("rcvBufSize", &x.rcvBufSize)
	m.Save("rcvClosed", &x.rcvClosed)
	m.Save("sndBufSize", &x.sndBufSize)
	m.Save("closed", &x.closed)
	m.Save("connected", &x.connected)
	m.Save("bound", &x.bound)
	m.Save("registeredNIC", &x.registeredNIC)
	m.Save("boundNIC", &x.boundNIC)
	m.Save("boundAddr", &x.boundAddr)
}

func (x *endpoint) load(m state.Map) {
	m.Load("netProto", &x.netProto)
	m.Load("transProto", &x.transProto)
	m.Load("waiterQueue", &x.waiterQueue)
	m.Load("associated", &x.associated)
	m.Load("rcvList", &x.rcvList)
	m.Load("rcvBufSize", &x.rcvBufSize)
	m.Load("rcvClosed", &x.rcvClosed)
	m.Load("sndBufSize", &x.sndBufSize)
	m.Load("closed", &x.closed)
	m.Load("connected", &x.connected)
	m.Load("bound", &x.bound)
	m.Load("registeredNIC", &x.registeredNIC)
	m.Load("boundNIC", &x.boundNIC)
	m.Load("boundAddr", &x.boundAddr)
	m.LoadValue("rcvBufSizeMax", new(int), func(y interface{}) { x.loadRcvBufSizeMax(y.(int)) })
	m.AfterLoad(x.afterLoad)
}

func (x *packetList) beforeSave() {}
func (x *packetList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *packetList) afterLoad() {}
func (x *packetList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *packetEntry) beforeSave() {}
func (x *packetEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *packetEntry) afterLoad() {}
func (x *packetEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func init() {
	state.Register("raw.packet", (*packet)(nil), state.Fns{Save: (*packet).save, Load: (*packet).load})
	state.Register("raw.endpoint", (*endpoint)(nil), state.Fns{Save: (*endpoint).save, Load: (*endpoint).load})
	state.Register("raw.packetList", (*packetList)(nil), state.Fns{Save: (*packetList).save, Load: (*packetList).load})
	state.Register("raw.packetEntry", (*packetEntry)(nil), state.Fns{Save: (*packetEntry).save, Load: (*packetEntry).load})
}
