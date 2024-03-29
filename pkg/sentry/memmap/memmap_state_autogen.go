// automatically generated by stateify.

package memmap

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *MappableRange) beforeSave() {}
func (x *MappableRange) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
}

func (x *MappableRange) afterLoad() {}
func (x *MappableRange) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
}

func (x *MappingOfRange) beforeSave() {}
func (x *MappingOfRange) save(m state.Map) {
	x.beforeSave()
	m.Save("MappingSpace", &x.MappingSpace)
	m.Save("AddrRange", &x.AddrRange)
	m.Save("Writable", &x.Writable)
}

func (x *MappingOfRange) afterLoad() {}
func (x *MappingOfRange) load(m state.Map) {
	m.Load("MappingSpace", &x.MappingSpace)
	m.Load("AddrRange", &x.AddrRange)
	m.Load("Writable", &x.Writable)
}

func (x *MappingSet) beforeSave() {}
func (x *MappingSet) save(m state.Map) {
	x.beforeSave()
	var root *MappingSegmentDataSlices = x.saveRoot()
	m.SaveValue("root", root)
}

func (x *MappingSet) afterLoad() {}
func (x *MappingSet) load(m state.Map) {
	m.LoadValue("root", new(*MappingSegmentDataSlices), func(y interface{}) { x.loadRoot(y.(*MappingSegmentDataSlices)) })
}

func (x *Mappingnode) beforeSave() {}
func (x *Mappingnode) save(m state.Map) {
	x.beforeSave()
	m.Save("nrSegments", &x.nrSegments)
	m.Save("parent", &x.parent)
	m.Save("parentIndex", &x.parentIndex)
	m.Save("hasChildren", &x.hasChildren)
	m.Save("keys", &x.keys)
	m.Save("values", &x.values)
	m.Save("children", &x.children)
}

func (x *Mappingnode) afterLoad() {}
func (x *Mappingnode) load(m state.Map) {
	m.Load("nrSegments", &x.nrSegments)
	m.Load("parent", &x.parent)
	m.Load("parentIndex", &x.parentIndex)
	m.Load("hasChildren", &x.hasChildren)
	m.Load("keys", &x.keys)
	m.Load("values", &x.values)
	m.Load("children", &x.children)
}

func (x *MappingSegmentDataSlices) beforeSave() {}
func (x *MappingSegmentDataSlices) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
	m.Save("Values", &x.Values)
}

func (x *MappingSegmentDataSlices) afterLoad() {}
func (x *MappingSegmentDataSlices) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
	m.Load("Values", &x.Values)
}

func init() {
	state.Register("memmap.MappableRange", (*MappableRange)(nil), state.Fns{Save: (*MappableRange).save, Load: (*MappableRange).load})
	state.Register("memmap.MappingOfRange", (*MappingOfRange)(nil), state.Fns{Save: (*MappingOfRange).save, Load: (*MappingOfRange).load})
	state.Register("memmap.MappingSet", (*MappingSet)(nil), state.Fns{Save: (*MappingSet).save, Load: (*MappingSet).load})
	state.Register("memmap.Mappingnode", (*Mappingnode)(nil), state.Fns{Save: (*Mappingnode).save, Load: (*Mappingnode).load})
	state.Register("memmap.MappingSegmentDataSlices", (*MappingSegmentDataSlices)(nil), state.Fns{Save: (*MappingSegmentDataSlices).save, Load: (*MappingSegmentDataSlices).load})
}
