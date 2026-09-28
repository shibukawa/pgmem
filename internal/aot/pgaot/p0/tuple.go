package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateTupleDescCopy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = F_palloc(m, v11*int32(108)+int32(28))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v11
	v23 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(2249)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v23
	if v11 <= v20 {
		v76 = v11
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v87
	if int32(0) < v76 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v32 = v11 * int32(100)
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v33 = int32(3)
	v36 = int32(28)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	base.MemoryCopy(m, v16+v11<<(uint(v33)%32)+v36, l0+v38<<(uint(v33)%32)+v36, v32)
	goto L7
L6:
	;
	goto L7
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v45 <= int32(0) {
		v76 = v45
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v50 = v45
	v52 = int32(0)
	goto L9
L9:
	;
	v64 = v16 + v50<<(uint(int32(3))%32) + v52*int32(100)
	v65 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+118)) = uint8(v65)
	*(*int32)(unsafe.Add(mBase, uint32(v64)+114)) = v65
	F_populate_compact_attribute(m, v16, v52)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v76 = v73
	goto L3
L11:
	;
	v72 = v52 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v72 < v73 {
		v50 = v73
		v52 = v72
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v92 = v16 + int32(28)
	v96 = v76
	v101 = v2
	v105 = v2
	goto L17
L14:
	;
	v158 = v2
	v163 = v76
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v158
	return v16
L16:
	;
	v158 = v151
	v163 = v129
	goto L15
L17:
	;
	v108 = v92 + v76<<(uint(int32(3))%32) + v101*int32(100)
	v111 = v92 + v101<<(uint(int32(3))%32)
	if v96 != v76 {
		v129 = v96
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v151 = v76
	goto L16
L19:
	;
	v130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v111)+2)))
	if v130 <= int32(0) {
		v151 = v101
		goto L16
	} else {
		goto L27
	}
L20:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+7)))
	if v113 != int32(118) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v129 = v101
	goto L19
L22:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+4)))
	if v116 != int32(1) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+6)))
	if v119&int32(6) != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(v111)+2)))
	if v122 <= int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+90)))
	if v125 != int32(118) {
		v129 = v76
		goto L19
	} else {
		goto L26
	}
L26:
	;
	goto L21
L27:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+90)))
	if v133 == int32(118) {
		v151 = v101
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+5)))
	v142 = (v105 + v136 - int32(1)) & (int32(0) - v136)
	if int32(_a_F_CreateTupleDescCopy_0) < v142 {
		v151 = v101
		goto L16
	} else {
		goto L29
	}
L29:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v111))) = uint16(v142)
	v148 = v101 + int32(1)
	if v148 != v76 {
		v96 = v129
		v101 = v148
		v105 = v142 + v130
		goto L17
	} else {
		goto L30
	}
L30:
	;
	goto L18
}
func F_EstimateTupleHashTableSpace(m *base.Module, l0 float64, l1 int32, l2 int32) int32 {
	var v7 int32
	_ = v7
	var v9 float64
	_ = v9
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v49 float64
	_ = v49
	var v55 int32
	_ = v55
	v7 = int32(-1)
	v9 = base.F64_div(l0, float64(0.9))
	if base.F64_ge(v9, float64(4.294967296e+09)) != 0 {
		v55 = v7
	} else {
		v12 = int64(2)
		v13 = base.I64_trunc_sat_f64_u(v9)
		if base.Ui64(v13) <= base.Ui64(v12) {
			v16 = v12
		} else {
			v16 = v13
		}
		v17 = int64(1)
		if v16&(v16-v17) == int64(0) {
			v27 = v16
		} else {
			v27 = v17 << (uint(int64(64)-base.I64_clz(v16)) % 64)
		}
		v29 = v27 * int64(12)
		if base.Ui64(int64(2147483646)) < base.Ui64(v29) {
			v55 = v7
		} else {
			v32 = int32(7)
			v34 = int32(-8)
			v49 = base.F64_add(base.F64_mul(l0, base.F64_convert_i32_u((l1+v32)&v34+(l2+v32)&v34+int32(16))), base.F64_convert_i32_u(base.I32_wrap_i64(v29)+int32(32)))
			if base.F64_ge(v49, float64(4.294967295e+09)) != 0 {
				v55 = v7
			} else {
				v55 = base.I32_trunc_sat_f64_u(v49)
			}
		}
	}
	return v55
}
func F_ExecInitMergeTupleSlots(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = v4 + int32(104)
	v7 = F_table_slot_create(m, v3, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v7
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v11 = F_table_slot_create(m, v10, v6)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)) = uint8(v13)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v11
			return
		}
	}
}
func F_TupleDescInitBuiltinEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v2 = l1
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int32(100)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = l0 + v16<<(uint(int32(3))%32) + v2*v15
	v24 = v22 - int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v5
	v30 = F_strncpy(m, v22-int32(68), l2, int32(64))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+63)) = uint8(v5)
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+14)) = v33
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+8)) = uint16(v33)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+2)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+18)) = uint16(v33)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+22)) = uint16(v33)
	v44 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)) = uint8(v44)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = l3
	v47 = int32(120)
	v48 = int32(105)
	v49 = int32(_a_F_TupleDescInitBuiltinEntry_0)
	switch l3 - int32(16) {
	case 0:
		v55 = int32(1)
		v68 = int32(112)
		v69 = v55
		v70 = int32(0)
		v71 = v55
		v72 = int32(99)
		*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v70
		v74 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v74)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v68)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v72)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v71)
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v69)
		F_populate_compact_attribute(m, l0, v2-int32(1))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return
		} else {
			m.G0 = v13 + int32(16)
			return
		}
	case 1, 2, 3, 5, 6, 8:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v90 = m.ExcPending
		if v90 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = l3
			F_errmsg_internal(m, int32(_a_F_TupleDescInitBuiltinEntry_1), v13)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_TupleDescInitBuiltinEntry_2), int32(1080), int32(_a_F_TupleDescInitBuiltinEntry_3))
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 4:
		v68 = int32(112)
		v69 = int32(8)
		v70 = int32(0)
		v71 = int32(1)
		v72 = int32(100)
		*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v70
		v74 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v74)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v68)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v72)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v71)
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v69)
		F_populate_compact_attribute(m, l0, v2-int32(1))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return
		} else {
			m.G0 = v13 + int32(16)
			return
		}
	case 7, 10:
		v68 = int32(112)
		v69 = int32(4)
		v70 = int32(0)
		v71 = int32(1)
		v72 = v48
		*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v70
		v74 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v74)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v68)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v72)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v71)
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v69)
		F_populate_compact_attribute(m, l0, v2-int32(1))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return
		} else {
			m.G0 = v13 + int32(16)
			return
		}
	case 9:
		v68 = v47
		v69 = v49
		v70 = v15
		v71 = v5
		v72 = v48
		*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v70
		v74 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v74)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v68)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v72)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v71)
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v69)
		F_populate_compact_attribute(m, l0, v2-int32(1))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return
		} else {
			m.G0 = v13 + int32(16)
			return
		}
	default:
		if l3 != int32(1009) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = l3
				F_errmsg_internal(m, int32(_a_F_TupleDescInitBuiltinEntry_1), v13)
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_TupleDescInitBuiltinEntry_2), int32(1080), int32(_a_F_TupleDescInitBuiltinEntry_3))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v68 = v47
			v69 = v49
			v70 = v15
			v71 = v5
			v72 = v48
			*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v70
			v74 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v74)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v68)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v72)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v71)
			*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v69)
			F_populate_compact_attribute(m, l0, v2-int32(1))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return
			} else {
				m.G0 = v13 + int32(16)
				return
			}
		}
	}
}
func F_UnlockTuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = m.G0
	v6 = int32(16)
	v7 = v5 - v6
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v11
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v13 | v14<<(uint(v6)%32)
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v20 = int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v20)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v19)
	v24 = F_LockRelease(m, v7, l2, int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
