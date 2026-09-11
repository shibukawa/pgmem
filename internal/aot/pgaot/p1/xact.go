package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_XactLockTableWaitErrorCb(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11 == int32(0) {
		m.G0 = v9 + int32(16)
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v14 == int32(0) {
			m.G0 = v9 + int32(16)
			return
		} else {
			v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
			if v17 == int32(0) {
				m.G0 = v9 + int32(16)
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v20 == int32(0) {
					m.G0 = v9 + int32(16)
					return
				} else {
					if base.Ui32(int32(8)) < base.Ui32(v11) {
						m.G0 = v9 + int32(16)
						return
					} else {
						F_set_errcontext_domain(m, int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+2)))
							v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
							v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v33
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v32 + int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v29 | v30<<(uint(int32(16))%32)
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v11<<(uint(int32(2))%32))+uint32(_consts[803])))
							F_errcontext_msg(m, v46, v9)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_xactGetCommittedChildren(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+52))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+48))
		v7 = v6
	} else {
		v7 = int32(0)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+52))
	return v9
}
