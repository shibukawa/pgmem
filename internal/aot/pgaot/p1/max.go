package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_max_stack_depth(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[1218])) = l0 << (uint(int32(10)) % 32)
	return
}
func F_assign_max_wal_size(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 float64
	_ = v15
	var v18 float64
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	*(*int32)(unsafe.Add(mBase, _consts[222])) = l0
	v9 = *(*int32)(unsafe.Add(mBase, _consts[138]))
	v11 = base.I32_div_s(v9, int32(1048576))
	v12 = base.I32_div_s(l0, v11)
	v15 = *(*float64)(unsafe.Add(mBase, _consts[223]))
	v18 = base.F64_div(base.F64_convert_i32_s(v12), base.F64_add(v15, float64(1)))
	if base.F64_lt(base.F64_abs(v18), float64(2.147483648e+09)) != 0 {
		v22 = base.I32_trunc_f64_s(v18)
		v24 = v22
	} else {
		v24 = int32(-2147483648)
	}
	if v24 <= int32(1) {
		v27 = int32(1)
	} else {
		v27 = v24
	}
	*(*int32)(unsafe.Add(mBase, _consts[224])) = v27
	return
}
func F_max_parallel_hazard_checker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_func_parallel(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		switch v8&int32(255) - int32(114) {
		case 0:
			v32 = int32(114)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v32)
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			if v35 == v32 {
				v39 = int32(1)
			} else {
				v39 = int32(0)
			}
			m.G0 = v6 + int32(16)
			return v39
		case 1:
			v39 = int32(0)
			m.G0 = v6 + int32(16)
			return v39
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
				F_errmsg_internal(m, int32(700500), v6)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(483005), int32(812), int32(75790))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 3:
			v16 = int32(117)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v16)
			v39 = int32(1)
			m.G0 = v6 + int32(16)
			return v39
		}
	}
}
