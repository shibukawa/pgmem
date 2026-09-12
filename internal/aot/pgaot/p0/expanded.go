package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_deconstruct_expanded_array(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v15 == int32(0) {
		v18 = int32(4554128)
		v19 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v21
		v23 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v23
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+44)))
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+46)))
		v29 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+47)))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
		if v33 != 0 {
			v34 = v13 + int32(8)
		} else {
			v34 = v23
		}
		F_deconstruct_array(m, v25, v27, v28, v29, v13+int32(12), v34, v13+int32(4))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v41
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v43
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v45
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v45
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v19
			m.G0 = v13 + int32(16)
			return
		}
	} else {
		m.G0 = v13 + int32(16)
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
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
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v121 = F_format_type_be(m, v120)
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v121
							F_errmsg(m, int32(366016), v9)
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(523128), int32(232), int32(509930))
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
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
				v34 = F_AllocSetContextCreateInternal(m, l1, int32(441061), int32(0), int32(8192), int32(8388608))
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
						v46 = F__emscripten_memset_bulkmem(m, v41, base.I32_extend8_s(int32(0)), int32(120))
						mBase = m.M
						v48 = int32(513)
						*(*uint16)(unsafe.Add(mBase, uint32(v46)+18)) = uint16(v48)
						v50 = int32(769)
						*(*uint16)(unsafe.Add(mBase, uint32(v46)+12)) = uint16(v50)
						*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = int32(1694320)
						*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v46
						*(*int32)(unsafe.Add(mBase, uint32(v46)+14)) = v46
						v59 = v46 + int32(120)
						*(*int32)(unsafe.Add(mBase, uint32(v46)+56)) = v59
						*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = int32(1384727874)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						*(*int32)(unsafe.Add(mBase, uint32(v46)+60)) = v59 + v63<<(uint(int32(2))%32)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						*(*int32)(unsafe.Add(mBase, uint32(v46)+64)) = v68
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v46)+32)) = v70
						*(*int32)(unsafe.Add(mBase, uint32(v46)+36)) = v70
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v46)+48)) = v29
						*(*int32)(unsafe.Add(mBase, uint32(v46)+40)) = v73
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
						if int32(0) <= v76 {
							*(*int32)(unsafe.Add(mBase, uint32(v46)+108)) = int32(1308)
							*(*int32)(unsafe.Add(mBase, uint32(v46)+112)) = v46
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
							v84 = v46 + int32(108)
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v85
							v87 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)) = uint8(v87)
							*(*int32)(unsafe.Add(mBase, uint32(v82)+40)) = v84
							*(*int32)(unsafe.Add(mBase, uint32(v46)+44)) = v27
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v91 + int32(1)
							m.G0 = v9 + int32(16)
							return v46
						} else {
							v95 = int32(4554128)
							v96 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v34
							v99 = F_CreateTupleDescCopy(m, v27)
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v46)+44)) = v99
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v46)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = v102 | int32(32)
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v96
								m.G0 = v9 + int32(16)
								return v46
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
			v34 = F_AllocSetContextCreateInternal(m, l1, int32(441061), int32(0), int32(8192), int32(8388608))
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
					v46 = F__emscripten_memset_bulkmem(m, v41, base.I32_extend8_s(int32(0)), int32(120))
					mBase = m.M
					v48 = int32(513)
					*(*uint16)(unsafe.Add(mBase, uint32(v46)+18)) = uint16(v48)
					v50 = int32(769)
					*(*uint16)(unsafe.Add(mBase, uint32(v46)+12)) = uint16(v50)
					*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = int32(1694320)
					*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(v46)+14)) = v46
					v59 = v46 + int32(120)
					*(*int32)(unsafe.Add(mBase, uint32(v46)+56)) = v59
					*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = int32(1384727874)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					*(*int32)(unsafe.Add(mBase, uint32(v46)+60)) = v59 + v63<<(uint(int32(2))%32)
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					*(*int32)(unsafe.Add(mBase, uint32(v46)+64)) = v68
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v46)+32)) = v70
					*(*int32)(unsafe.Add(mBase, uint32(v46)+36)) = v70
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+48)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(v46)+40)) = v73
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
					if int32(0) <= v76 {
						*(*int32)(unsafe.Add(mBase, uint32(v46)+108)) = int32(1308)
						*(*int32)(unsafe.Add(mBase, uint32(v46)+112)) = v46
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
						v84 = v46 + int32(108)
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v85
						v87 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)) = uint8(v87)
						*(*int32)(unsafe.Add(mBase, uint32(v82)+40)) = v84
						*(*int32)(unsafe.Add(mBase, uint32(v46)+44)) = v27
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v91 + int32(1)
						m.G0 = v9 + int32(16)
						return v46
					} else {
						v95 = int32(4554128)
						v96 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v34
						v99 = F_CreateTupleDescCopy(m, v27)
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v46)+44)) = v99
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v46)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = v102 | int32(32)
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v96
							m.G0 = v9 + int32(16)
							return v46
						}
					}
				}
			}
		}
	}
}
