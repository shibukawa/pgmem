package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pointerhash_create(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 float64
	_ = v15
	var v18 float64
	_ = v18
	var v21 float64
	_ = v21
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v43 int64
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v65 int64
	_ = v65
	var v80 float64
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	v8 = F_MemoryContextAllocZero(m, l0, int32(32))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = l0
		v15 = float64(4.294967296e+09)
		v18 = base.F64_div(base.F64_convert_i32_u(l1), float64(0.9))
		if base.F64_ge(v18, v15) != 0 {
			v21 = v15
		} else {
			v21 = v18
		}
		if base.F64_lt(v21, float64(1.8446744073709552e+19))&base.F64_ge(v21, float64(0)) != 0 {
			v27 = base.I64_trunc_f64_u(v21)
			v29 = v27
		} else {
			v29 = int64(0)
		}
		if base.Ui64(v29) <= base.Ui64(int64(2)) {
			v32 = int64(2)
		} else {
			v32 = v29
		}
		v33 = int64(1)
		if v32&(v32-v33) == int64(0) {
			v43 = v32
		} else {
			v43 = v33 << (uint(int64(64)-base.I64_clz(v32)) % 64)
		}
		if base.Ui64(v43<<(uint(int64(3))%64)) < base.Ui64(int64(2147483647)) {
			v52 = F_MemoryContextAllocExtended(m, l0, base.I32_wrap_i64(v43)<<(uint(int32(3))%32), int32(5))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v52
				v55 = int64(1)
				if v43&(v43-v55) == int64(0) {
					v65 = v43
				} else {
					v65 = v55 << (uint(int64(64)-base.I64_clz(v43)) % 64)
				}
				if base.Ui64(int64(2147483647)) <= base.Ui64(v65<<(uint(int64(3))%64)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(420628), int32(0))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(343961), int32(327), int32(359166))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = v65
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = base.I32_wrap_i64(v65) - int32(1)
					v80 = base.F64_mul(base.F64_convert_i64_u(v65), float64(0.9))
					if base.F64_lt(v80, float64(4.294967296e+09))&base.F64_ge(v80, float64(0)) != 0 {
						v86 = base.I32_trunc_f64_u(v80)
						v88 = v86
					} else {
						v88 = int32(0)
					}
					if v65 == int64(4294967296) {
						v89 = int32(-85899346)
					} else {
						v89 = v88
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v89
					return v8
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(420628), int32(0))
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(343961), int32(327), int32(359166))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
