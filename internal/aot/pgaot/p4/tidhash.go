package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tidhash_create(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
						F_errmsg_internal(m, int32(418921), int32(0))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(342446), int32(327), int32(357682))
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
				F_errmsg_internal(m, int32(418921), int32(0))
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(342446), int32(327), int32(357682))
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
func F_tidhash_insert_hash(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v10)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	v16 = F_tidhash_insert_hash_internal(m, l0, v8+int32(8), l2, l3)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v16
	}
}
func F_tidhash_lookup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v28 int64
	_ = v28
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = l1 + int32(4)
	v13 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v12))))
	v14 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1))))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12))))
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)) = uint16(v15)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = v13 << (uint(int64(32)) % 64)
	v23 = int64(33)
	v28 = (int64(base.Ui64(v22)>>(uint(v23)%64)) ^ (v22 | v14)) * int64(-49064778989728563)
	v33 = (int64(base.Ui64(v28)>>(uint(v23)%64)) ^ v28) * int64(-4265267296055464877)
	v38 = v20 & base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(v23)%64))^v33)
	v41 = v19 + v38<<(uint(int32(3))%32)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+6)))
	if v42 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v85
L2:
	;
	v44 = v41
	v48 = v38
	goto L5
L3:
	;
	goto L4
L4:
	;
	v85 = int32(0)
	goto L1
L5:
	;
	v50 = v9 + int32(8)
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+2)))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44))))
	v53 = int32(16)
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+2)))
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50))))
	if v51|v52<<(uint(v53)%32) == v56|v57<<(uint(v53)%32) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L4
L7:
	;
	if v67 != 0 {
		v85 = v44
		goto L1
	} else {
		goto L13
	}
L8:
	;
	goto L7
L9:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)))
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)))
	if v63 == v64 {
		v67 = int32(1)
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v67 = int32(0)
	goto L8
L12:
	;
	goto L11
L13:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v72 = v69 & (v48 + int32(1))
	v75 = v68 + v72<<(uint(int32(3))%32)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+6)))
	if v76 != 0 {
		v44 = v75
		v48 = v72
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L6
}
