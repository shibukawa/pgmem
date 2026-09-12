package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_smgr_aio_describe_identity(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v4 = m.G0
	v6 = v4 - int32(112)
	m.G0 = v6
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v16&int32(256) != 0 {
		v19 = v14
	} else {
		v19 = int32(-1)
	}
	F_GetRelationPath(m, v6+int32(40), v10, v11, v12, v19, base.I32_extend8_s(v16))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		switch v25 {
		case 0:
			*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v6 + int32(40)
			v32 = F_psprintf(m, int32(717588), v6+int32(16))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v57 = v32
				m.G0 = v6 + int32(112)
				return v57
			}
		case 1:
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v34
			*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v6 + int32(40)
			v42 = F_psprintf(m, int32(716508), v6+int32(32))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v57 = v42
				m.G0 = v6 + int32(112)
				return v57
			}
		default:
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v44
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v44 + v25 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v6 + int32(40)
			v54 = F_psprintf(m, int32(716481), v6)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				v57 = v54
				m.G0 = v6 + int32(112)
				return v57
			}
		}
	}
}
