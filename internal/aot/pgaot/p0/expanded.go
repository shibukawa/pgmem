package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_deconstruct_expanded_array(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v9 == int32(0) {
		v12 = int32(_a_F_deconstruct_expanded_array_0)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_deconstruct_expanded_array[0]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, _c_F_deconstruct_expanded_array[0])) = v15
		v17 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+44)))
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+46)))
		v23 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+47)))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
		if v29 != 0 {
			v30 = v7 + int32(8)
		} else {
			v30 = v17
		}
		F_deconstruct_array(m, v20, v21, v22, v23, v7+int32(12), v30, v7+int32(4))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v37
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v39
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v39
			*(*int32)(unsafe.Add(mBase, _c_F_deconstruct_expanded_array[0])) = v13
			m.G0 = v7 + int32(16)
			return
		}
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_make_expanded_record_from_tupdesc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 != int32(2249) {
		v15 = F_lookup_type_cache(m, v11, int32(256))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+188))
			if v19 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v120 = F_format_type_be(m, v119)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v120
							F_errmsg(m, int32(_a_F_make_expanded_record_from_tupdesc_0), v9)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_make_expanded_record_from_tupdesc_1), int32(232), int32(_a_F_make_expanded_record_from_tupdesc_2))
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
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
			} else {
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v15)+192))
				v27 = v19
				v29 = v22
				v34 = F_AllocSetContextCreateInternal(m, l1, int32(_a_F_make_expanded_record_from_tupdesc_3), int32(0), int32(_a_F_make_expanded_record_from_tupdesc_4), int32(_a_F_make_expanded_record_from_tupdesc_5))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					v41 = F_MemoryContextAlloc(m, v34, v36*int32(5)+int32(120))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						base.MemoryFill(m, v41, int32(0), int32(120))
						v47 = int32(513)
						*(*uint16)(unsafe.Add(mBase, uint32(v41)+18)) = uint16(v47)
						v49 = int32(769)
						*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)) = uint16(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = int32(_a_F_make_expanded_record_from_tupdesc_6)
						*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = v41
						*(*int32)(unsafe.Add(mBase, uint32(v41)+14)) = v41
						v58 = v41 + int32(120)
						*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v58
						*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = int32(1384727874)
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v58 + v62<<(uint(int32(2))%32)
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v67
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v69
						*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v69
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v41)+48)) = v29
						*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v72
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
						if int32(0) <= v75 {
							*(*int32)(unsafe.Add(mBase, uint32(v41)+108)) = int32(1292)
							*(*int32)(unsafe.Add(mBase, uint32(v41)+112)) = v41
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
							v83 = v41 + int32(108)
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v84
							v86 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v81)+4)) = uint8(v86)
							*(*int32)(unsafe.Add(mBase, uint32(v81)+40)) = v83
							*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v27
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v90 + int32(1)
							m.G0 = v9 + int32(16)
							return v41
						} else {
							v94 = int32(_a_F_make_expanded_record_from_tupdesc_7)
							v95 = *(*int32)(unsafe.Add(mBase, _c_F_make_expanded_record_from_tupdesc[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_make_expanded_record_from_tupdesc[0])) = v34
							v98 = F_CreateTupleDescCopy(m, v27)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v98
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v101 | int32(32)
								*(*int32)(unsafe.Add(mBase, _c_F_make_expanded_record_from_tupdesc[0])) = v95
								m.G0 = v9 + int32(16)
								return v41
							}
						}
					}
				}
			}
		}
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v25 = F_assign_record_type_identifier(m, int32(2249), v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = l0
			v29 = v25
			v34 = F_AllocSetContextCreateInternal(m, l1, int32(_a_F_make_expanded_record_from_tupdesc_3), int32(0), int32(_a_F_make_expanded_record_from_tupdesc_4), int32(_a_F_make_expanded_record_from_tupdesc_5))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v41 = F_MemoryContextAlloc(m, v34, v36*int32(5)+int32(120))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					base.MemoryFill(m, v41, int32(0), int32(120))
					v47 = int32(513)
					*(*uint16)(unsafe.Add(mBase, uint32(v41)+18)) = uint16(v47)
					v49 = int32(769)
					*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)) = uint16(v49)
					*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = int32(_a_F_make_expanded_record_from_tupdesc_6)
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v41)+14)) = v41
					v58 = v41 + int32(120)
					*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v58
					*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = int32(1384727874)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v58 + v62<<(uint(int32(2))%32)
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v67
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v69
					*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v69
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v41)+48)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v72
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
					if int32(0) <= v75 {
						*(*int32)(unsafe.Add(mBase, uint32(v41)+108)) = int32(1292)
						*(*int32)(unsafe.Add(mBase, uint32(v41)+112)) = v41
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
						v83 = v41 + int32(108)
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v84
						v86 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v81)+4)) = uint8(v86)
						*(*int32)(unsafe.Add(mBase, uint32(v81)+40)) = v83
						*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v27
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v90 + int32(1)
						m.G0 = v9 + int32(16)
						return v41
					} else {
						v94 = int32(_a_F_make_expanded_record_from_tupdesc_7)
						v95 = *(*int32)(unsafe.Add(mBase, _c_F_make_expanded_record_from_tupdesc[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_make_expanded_record_from_tupdesc[0])) = v34
						v98 = F_CreateTupleDescCopy(m, v27)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v98
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v101 | int32(32)
							*(*int32)(unsafe.Add(mBase, _c_F_make_expanded_record_from_tupdesc[0])) = v95
							m.G0 = v9 + int32(16)
							return v41
						}
					}
				}
			}
		}
	}
}
