// automatically generated by stateify.

package sys

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *cpunum) beforeSave() {}
func (x *cpunum) save(m state.Map) {
	x.beforeSave()
	m.Save("InodeSimpleAttributes", &x.InodeSimpleAttributes)
	m.Save("InodeStaticFileGetter", &x.InodeStaticFileGetter)
}

func (x *cpunum) afterLoad() {}
func (x *cpunum) load(m state.Map) {
	m.Load("InodeSimpleAttributes", &x.InodeSimpleAttributes)
	m.Load("InodeStaticFileGetter", &x.InodeStaticFileGetter)
}

func (x *filesystem) beforeSave() {}
func (x *filesystem) save(m state.Map) {
	x.beforeSave()
}

func (x *filesystem) afterLoad() {}
func (x *filesystem) load(m state.Map) {
}

func init() {
	state.Register("sys.cpunum", (*cpunum)(nil), state.Fns{Save: (*cpunum).save, Load: (*cpunum).load})
	state.Register("sys.filesystem", (*filesystem)(nil), state.Fns{Save: (*filesystem).save, Load: (*filesystem).load})
}
