// automatically generated by stateify.

package auth

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *AtomicPtrCredentials) beforeSave() {}
func (x *AtomicPtrCredentials) save(m state.Map) {
	x.beforeSave()
	var ptr *Credentials = x.savePtr()
	m.SaveValue("ptr", ptr)
}

func (x *AtomicPtrCredentials) afterLoad() {}
func (x *AtomicPtrCredentials) load(m state.Map) {
	m.LoadValue("ptr", new(*Credentials), func(y interface{}) { x.loadPtr(y.(*Credentials)) })
}

func (x *Credentials) beforeSave() {}
func (x *Credentials) save(m state.Map) {
	x.beforeSave()
	m.Save("RealKUID", &x.RealKUID)
	m.Save("EffectiveKUID", &x.EffectiveKUID)
	m.Save("SavedKUID", &x.SavedKUID)
	m.Save("RealKGID", &x.RealKGID)
	m.Save("EffectiveKGID", &x.EffectiveKGID)
	m.Save("SavedKGID", &x.SavedKGID)
	m.Save("ExtraKGIDs", &x.ExtraKGIDs)
	m.Save("PermittedCaps", &x.PermittedCaps)
	m.Save("InheritableCaps", &x.InheritableCaps)
	m.Save("EffectiveCaps", &x.EffectiveCaps)
	m.Save("BoundingCaps", &x.BoundingCaps)
	m.Save("KeepCaps", &x.KeepCaps)
	m.Save("UserNamespace", &x.UserNamespace)
}

func (x *Credentials) afterLoad() {}
func (x *Credentials) load(m state.Map) {
	m.Load("RealKUID", &x.RealKUID)
	m.Load("EffectiveKUID", &x.EffectiveKUID)
	m.Load("SavedKUID", &x.SavedKUID)
	m.Load("RealKGID", &x.RealKGID)
	m.Load("EffectiveKGID", &x.EffectiveKGID)
	m.Load("SavedKGID", &x.SavedKGID)
	m.Load("ExtraKGIDs", &x.ExtraKGIDs)
	m.Load("PermittedCaps", &x.PermittedCaps)
	m.Load("InheritableCaps", &x.InheritableCaps)
	m.Load("EffectiveCaps", &x.EffectiveCaps)
	m.Load("BoundingCaps", &x.BoundingCaps)
	m.Load("KeepCaps", &x.KeepCaps)
	m.Load("UserNamespace", &x.UserNamespace)
}

func (x *IDMapEntry) beforeSave() {}
func (x *IDMapEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("FirstID", &x.FirstID)
	m.Save("FirstParentID", &x.FirstParentID)
	m.Save("Length", &x.Length)
}

func (x *IDMapEntry) afterLoad() {}
func (x *IDMapEntry) load(m state.Map) {
	m.Load("FirstID", &x.FirstID)
	m.Load("FirstParentID", &x.FirstParentID)
	m.Load("Length", &x.Length)
}

func (x *idMapRange) beforeSave() {}
func (x *idMapRange) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
}

func (x *idMapRange) afterLoad() {}
func (x *idMapRange) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
}

func (x *idMapSet) beforeSave() {}
func (x *idMapSet) save(m state.Map) {
	x.beforeSave()
	var root *idMapSegmentDataSlices = x.saveRoot()
	m.SaveValue("root", root)
}

func (x *idMapSet) afterLoad() {}
func (x *idMapSet) load(m state.Map) {
	m.LoadValue("root", new(*idMapSegmentDataSlices), func(y interface{}) { x.loadRoot(y.(*idMapSegmentDataSlices)) })
}

func (x *idMapnode) beforeSave() {}
func (x *idMapnode) save(m state.Map) {
	x.beforeSave()
	m.Save("nrSegments", &x.nrSegments)
	m.Save("parent", &x.parent)
	m.Save("parentIndex", &x.parentIndex)
	m.Save("hasChildren", &x.hasChildren)
	m.Save("keys", &x.keys)
	m.Save("values", &x.values)
	m.Save("children", &x.children)
}

func (x *idMapnode) afterLoad() {}
func (x *idMapnode) load(m state.Map) {
	m.Load("nrSegments", &x.nrSegments)
	m.Load("parent", &x.parent)
	m.Load("parentIndex", &x.parentIndex)
	m.Load("hasChildren", &x.hasChildren)
	m.Load("keys", &x.keys)
	m.Load("values", &x.values)
	m.Load("children", &x.children)
}

func (x *idMapSegmentDataSlices) beforeSave() {}
func (x *idMapSegmentDataSlices) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
	m.Save("Values", &x.Values)
}

func (x *idMapSegmentDataSlices) afterLoad() {}
func (x *idMapSegmentDataSlices) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
	m.Load("Values", &x.Values)
}

func (x *UserNamespace) beforeSave() {}
func (x *UserNamespace) save(m state.Map) {
	x.beforeSave()
	m.Save("parent", &x.parent)
	m.Save("owner", &x.owner)
	m.Save("uidMapFromParent", &x.uidMapFromParent)
	m.Save("uidMapToParent", &x.uidMapToParent)
	m.Save("gidMapFromParent", &x.gidMapFromParent)
	m.Save("gidMapToParent", &x.gidMapToParent)
}

func (x *UserNamespace) afterLoad() {}
func (x *UserNamespace) load(m state.Map) {
	m.Load("parent", &x.parent)
	m.Load("owner", &x.owner)
	m.Load("uidMapFromParent", &x.uidMapFromParent)
	m.Load("uidMapToParent", &x.uidMapToParent)
	m.Load("gidMapFromParent", &x.gidMapFromParent)
	m.Load("gidMapToParent", &x.gidMapToParent)
}

func init() {
	state.Register("auth.AtomicPtrCredentials", (*AtomicPtrCredentials)(nil), state.Fns{Save: (*AtomicPtrCredentials).save, Load: (*AtomicPtrCredentials).load})
	state.Register("auth.Credentials", (*Credentials)(nil), state.Fns{Save: (*Credentials).save, Load: (*Credentials).load})
	state.Register("auth.IDMapEntry", (*IDMapEntry)(nil), state.Fns{Save: (*IDMapEntry).save, Load: (*IDMapEntry).load})
	state.Register("auth.idMapRange", (*idMapRange)(nil), state.Fns{Save: (*idMapRange).save, Load: (*idMapRange).load})
	state.Register("auth.idMapSet", (*idMapSet)(nil), state.Fns{Save: (*idMapSet).save, Load: (*idMapSet).load})
	state.Register("auth.idMapnode", (*idMapnode)(nil), state.Fns{Save: (*idMapnode).save, Load: (*idMapnode).load})
	state.Register("auth.idMapSegmentDataSlices", (*idMapSegmentDataSlices)(nil), state.Fns{Save: (*idMapSegmentDataSlices).save, Load: (*idMapSegmentDataSlices).load})
	state.Register("auth.UserNamespace", (*UserNamespace)(nil), state.Fns{Save: (*UserNamespace).save, Load: (*UserNamespace).load})
}
