package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_jspInitByBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = l1 + l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v14
	v16 = int32(4)
	v19 = (v12 + v16) & int32(-4)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20
	v22 = v19 - l1
	v24 = v22 + v16
	switch v14 {
	case 0, 21, 22, 26, 27, 31, 32, 33, 34, 35, 36, 38, 40, 43, 44, 45, 47, 48, 49:
		m.G0 = v10 + int32(16)
		return
	case 1, 25, 28:
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l1+v24)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v26
		v30 = v22 + int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1 + v30
		m.G0 = v10 + int32(16)
		return
	case 2, 3:
		v30 = v24
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1 + v30
		m.G0 = v10 + int32(16)
		return
	case 4, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 41, 46:
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l1+v24)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v72
		v74 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v74
		m.G0 = v10 + int32(16)
		return
	case 6, 7, 19, 20, 29, 30, 37, 50, 51, 52, 53:
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l1+v24)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v34
		m.G0 = v10 + int32(16)
		return
	case 23:
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+v24)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v19 + int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v37
		m.G0 = v10 + int32(16)
		return
	case 24:
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l1+v24)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v43
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v45
		m.G0 = v10 + int32(16)
		return
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return
		} else {
			v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v61
			F_errmsg_internal(m, int32(485678), v10)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				F_errfinish(m, int32(498258), int32(1076), int32(226429))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 42:
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l1+v24)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v48
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v50
		v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v19 + int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v52
		m.G0 = v10 + int32(16)
		return
	}
}
