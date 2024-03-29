// automatically generated by stateify.

package transport

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *connectionedEndpoint) beforeSave() {}
func (x *connectionedEndpoint) save(m state.Map) {
	x.beforeSave()
	var acceptedChan []*connectionedEndpoint = x.saveAcceptedChan()
	m.SaveValue("acceptedChan", acceptedChan)
	m.Save("baseEndpoint", &x.baseEndpoint)
	m.Save("id", &x.id)
	m.Save("idGenerator", &x.idGenerator)
	m.Save("stype", &x.stype)
}

func (x *connectionedEndpoint) afterLoad() {}
func (x *connectionedEndpoint) load(m state.Map) {
	m.Load("baseEndpoint", &x.baseEndpoint)
	m.Load("id", &x.id)
	m.Load("idGenerator", &x.idGenerator)
	m.Load("stype", &x.stype)
	m.LoadValue("acceptedChan", new([]*connectionedEndpoint), func(y interface{}) { x.loadAcceptedChan(y.([]*connectionedEndpoint)) })
}

func (x *connectionlessEndpoint) beforeSave() {}
func (x *connectionlessEndpoint) save(m state.Map) {
	x.beforeSave()
	m.Save("baseEndpoint", &x.baseEndpoint)
}

func (x *connectionlessEndpoint) afterLoad() {}
func (x *connectionlessEndpoint) load(m state.Map) {
	m.Load("baseEndpoint", &x.baseEndpoint)
}

func (x *queue) beforeSave() {}
func (x *queue) save(m state.Map) {
	x.beforeSave()
	m.Save("AtomicRefCount", &x.AtomicRefCount)
	m.Save("ReaderQueue", &x.ReaderQueue)
	m.Save("WriterQueue", &x.WriterQueue)
	m.Save("closed", &x.closed)
	m.Save("used", &x.used)
	m.Save("limit", &x.limit)
	m.Save("dataList", &x.dataList)
}

func (x *queue) afterLoad() {}
func (x *queue) load(m state.Map) {
	m.Load("AtomicRefCount", &x.AtomicRefCount)
	m.Load("ReaderQueue", &x.ReaderQueue)
	m.Load("WriterQueue", &x.WriterQueue)
	m.Load("closed", &x.closed)
	m.Load("used", &x.used)
	m.Load("limit", &x.limit)
	m.Load("dataList", &x.dataList)
}

func (x *messageList) beforeSave() {}
func (x *messageList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *messageList) afterLoad() {}
func (x *messageList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *messageEntry) beforeSave() {}
func (x *messageEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *messageEntry) afterLoad() {}
func (x *messageEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func (x *ControlMessages) beforeSave() {}
func (x *ControlMessages) save(m state.Map) {
	x.beforeSave()
	m.Save("Rights", &x.Rights)
	m.Save("Credentials", &x.Credentials)
}

func (x *ControlMessages) afterLoad() {}
func (x *ControlMessages) load(m state.Map) {
	m.Load("Rights", &x.Rights)
	m.Load("Credentials", &x.Credentials)
}

func (x *message) beforeSave() {}
func (x *message) save(m state.Map) {
	x.beforeSave()
	m.Save("messageEntry", &x.messageEntry)
	m.Save("Data", &x.Data)
	m.Save("Control", &x.Control)
	m.Save("Address", &x.Address)
}

func (x *message) afterLoad() {}
func (x *message) load(m state.Map) {
	m.Load("messageEntry", &x.messageEntry)
	m.Load("Data", &x.Data)
	m.Load("Control", &x.Control)
	m.Load("Address", &x.Address)
}

func (x *queueReceiver) beforeSave() {}
func (x *queueReceiver) save(m state.Map) {
	x.beforeSave()
	m.Save("readQueue", &x.readQueue)
}

func (x *queueReceiver) afterLoad() {}
func (x *queueReceiver) load(m state.Map) {
	m.Load("readQueue", &x.readQueue)
}

func (x *streamQueueReceiver) beforeSave() {}
func (x *streamQueueReceiver) save(m state.Map) {
	x.beforeSave()
	m.Save("queueReceiver", &x.queueReceiver)
	m.Save("buffer", &x.buffer)
	m.Save("control", &x.control)
	m.Save("addr", &x.addr)
}

func (x *streamQueueReceiver) afterLoad() {}
func (x *streamQueueReceiver) load(m state.Map) {
	m.Load("queueReceiver", &x.queueReceiver)
	m.Load("buffer", &x.buffer)
	m.Load("control", &x.control)
	m.Load("addr", &x.addr)
}

func (x *connectedEndpoint) beforeSave() {}
func (x *connectedEndpoint) save(m state.Map) {
	x.beforeSave()
	m.Save("endpoint", &x.endpoint)
	m.Save("writeQueue", &x.writeQueue)
}

func (x *connectedEndpoint) afterLoad() {}
func (x *connectedEndpoint) load(m state.Map) {
	m.Load("endpoint", &x.endpoint)
	m.Load("writeQueue", &x.writeQueue)
}

func (x *baseEndpoint) beforeSave() {}
func (x *baseEndpoint) save(m state.Map) {
	x.beforeSave()
	m.Save("Queue", &x.Queue)
	m.Save("passcred", &x.passcred)
	m.Save("receiver", &x.receiver)
	m.Save("connected", &x.connected)
	m.Save("path", &x.path)
}

func (x *baseEndpoint) afterLoad() {}
func (x *baseEndpoint) load(m state.Map) {
	m.Load("Queue", &x.Queue)
	m.Load("passcred", &x.passcred)
	m.Load("receiver", &x.receiver)
	m.Load("connected", &x.connected)
	m.Load("path", &x.path)
}

func init() {
	state.Register("transport.connectionedEndpoint", (*connectionedEndpoint)(nil), state.Fns{Save: (*connectionedEndpoint).save, Load: (*connectionedEndpoint).load})
	state.Register("transport.connectionlessEndpoint", (*connectionlessEndpoint)(nil), state.Fns{Save: (*connectionlessEndpoint).save, Load: (*connectionlessEndpoint).load})
	state.Register("transport.queue", (*queue)(nil), state.Fns{Save: (*queue).save, Load: (*queue).load})
	state.Register("transport.messageList", (*messageList)(nil), state.Fns{Save: (*messageList).save, Load: (*messageList).load})
	state.Register("transport.messageEntry", (*messageEntry)(nil), state.Fns{Save: (*messageEntry).save, Load: (*messageEntry).load})
	state.Register("transport.ControlMessages", (*ControlMessages)(nil), state.Fns{Save: (*ControlMessages).save, Load: (*ControlMessages).load})
	state.Register("transport.message", (*message)(nil), state.Fns{Save: (*message).save, Load: (*message).load})
	state.Register("transport.queueReceiver", (*queueReceiver)(nil), state.Fns{Save: (*queueReceiver).save, Load: (*queueReceiver).load})
	state.Register("transport.streamQueueReceiver", (*streamQueueReceiver)(nil), state.Fns{Save: (*streamQueueReceiver).save, Load: (*streamQueueReceiver).load})
	state.Register("transport.connectedEndpoint", (*connectedEndpoint)(nil), state.Fns{Save: (*connectedEndpoint).save, Load: (*connectedEndpoint).load})
	state.Register("transport.baseEndpoint", (*baseEndpoint)(nil), state.Fns{Save: (*baseEndpoint).save, Load: (*baseEndpoint).load})
}
