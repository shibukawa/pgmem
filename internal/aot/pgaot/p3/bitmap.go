package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cost_bitmap_tree_node(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 float64
	_ = v12
	var v14 float64
	_ = v14
	var v16 float64
	_ = v16
	var v18 float64
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v36 float64
	_ = v36
	var v39 float64
	_ = v39
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v9 - int32(280) {
	case 0:
		v34 = *(*float64)(unsafe.Add(mBase, uint32(l0)+96))
		*(*float64)(unsafe.Add(mBase, uint32(l1))) = v34
		v36 = *(*float64)(unsafe.Add(mBase, uint32(l0)+104))
		*(*float64)(unsafe.Add(mBase, uint32(l2))) = v36
		v39 = *(*float64)(unsafe.Add(mBase, _consts[598]))
		v42 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
		v44 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
		*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_add(base.F64_mul(base.F64_mul(v39, float64(0.1)), v42), v44)
		m.G0 = v7 + int32(16)
		return
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v24
			F_errmsg_internal(m, int32(507644), v7)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_errfinish(m, int32(521926), int32(1149), int32(430198))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 3:
		v12 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
		*(*float64)(unsafe.Add(mBase, uint32(l1))) = v12
		v14 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
		*(*float64)(unsafe.Add(mBase, uint32(l2))) = v14
		m.G0 = v7 + int32(16)
		return
	case 4:
		v16 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
		*(*float64)(unsafe.Add(mBase, uint32(l1))) = v16
		v18 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
		*(*float64)(unsafe.Add(mBase, uint32(l2))) = v18
		m.G0 = v7 + int32(16)
		return
	}
}
func F_create_bitmap_heap_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	v10 = F_palloc0(m, int32(80))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(1477468750106)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v17
		v19 = F_get_baserel_parampathinfo(m, l0, l1, l3)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)) = uint8(base.B2i32(v21 < l5))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v19
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l5
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v25)
			F_cost_bitmap_heap_scan(m, v10, l0, l1, v19, l2, l4)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
