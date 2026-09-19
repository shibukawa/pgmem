package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_record_type_typmod(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = l0
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[0]))
	if v14 != 0 {
		v36 = v14
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v42 = F_hash_search(m, v36, v8+int32(-4), int32(0), v8+int32(-52))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L7
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(1605)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(1606)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+28)) = int64(17179869188)
	v27 = F_hash_create(m, int32(_a_F_assign_record_type_typmod_3), int32(64), v8+int32(-52), int32(200))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[0])) = v27
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[2]))
	if v31 != 0 {
		v36 = v27
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[0]))
	v36 = v35
	goto L1
L7:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
	if v44 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v10 - int32(-64)
	return
L9:
	;
	v54 = int32(_a_F_assign_record_type_typmod_0)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[1]))
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[1])) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	v61 = F_find_or_make_matching_shared_tupledesc(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L13
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v47 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v51
	goto L8
L12:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v163 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v159+v162<<(uint(v163)%32))+8)) = v158
	v167 = int32(_a_F_assign_record_type_typmod_2)
	v169 = *(*int64)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[6]))
	v171 = v169 + int64(1)
	*(*int64)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[6])) = v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v159+v173<<(uint(v163)%32)))) = v171
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[0]))
	v184 = F_hash_search(m, v179, v8+int32(-4), int32(1), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L3
	} else {
		goto L40
	}
L13:
	;
	if v61 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[3]))
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[4]))
	if v68 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[4]))
	if v122 != 0 {
		goto L31
	} else {
		goto L32
	}
L17:
	;
	if v82 <= v66 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[5]))
	v82 = v70
	v83 = v68
	goto L17
L19:
	;
	goto L20
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[2]))
	v75 = F_MemoryContextAllocZero(m, v73, int32(1024))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[5])) = int32(64)
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[4])) = v75
	v82 = int32(64)
	v83 = v75
	goto L17
L22:
	;
	v87 = int32(1)
	v90 = v66 + v87
	if v66&v90 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	v107 = F_CreateTupleDescCopy(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L29
	}
L25:
	;
	v95 = v87 << (uint(int32(32)-base.I32_clz(v90)) % 32)
	goto L27
L26:
	;
	v95 = v90
	goto L27
L27:
	;
	v98 = F_repalloc0(m, v83, v82<<(uint(int32(4))%32), v95<<(uint(int32(4))%32))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[5])) = v95
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[4])) = v98
	goto L24
L29:
	;
	v109 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v107)+12)) = v109
	v111 = int32(_a_F_assign_record_type_typmod_1)
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[3])) = v113 + v109
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v113
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[4]))
	v158 = v107
	v159 = v119
	goto L12
L30:
	;
	if v120 < v137 {
		v158 = v61
		v159 = v136
		goto L12
	} else {
		goto L35
	}
L31:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[5]))
	v136 = v122
	v137 = v124
	goto L30
L32:
	;
	goto L33
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[2]))
	v129 = F_MemoryContextAllocZero(m, v127, int32(1024))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[5])) = int32(64)
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[4])) = v129
	v136 = v129
	v137 = int32(64)
	goto L30
L35:
	;
	v141 = int32(1)
	v144 = v120 + v141
	if v144&v120 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v149 = v141 << (uint(int32(32)-base.I32_clz(v144)) % 32)
	goto L38
L37:
	;
	v149 = v144
	goto L38
L38:
	;
	v152 = F_repalloc0(m, v136, v137<<(uint(int32(4))%32), v149<<(uint(int32(4))%32))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[5])) = v149
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[4])) = v152
	v158 = v61
	v159 = v152
	goto L12
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v158
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v187)+8)) = v188
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[1])) = v55
	goto L8
}
func F_findTypeTypmodoutFunction(m *base.Module, l0 int32) int32 {
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v11 = Fn13899(m, l0, int32(_a_F_findTypeTypmodoutFunction_0), int32(2233), int32(_a_F_findTypeTypmodoutFunction_1), int32(_a_F_findTypeTypmodoutFunction_2), int32(2227), int32(2240), int32(_a_F_findTypeTypmodoutFunction_3), int32(2275), int32(23))
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_format_type_with_typemod(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_format_type_extended(m, l0, l1, int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_get_type_category_preferred(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(_a_F_get_type_category_preferred_0), v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_get_type_category_preferred_1), int32(2857), int32(_a_F_get_type_category_preferred_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v30 = v28 + v29
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+80)))
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v31)
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+81)))
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v33)
			F_ReleaseCatCache(m, v11)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_record_type_typmod_compare(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v10 != v11 {
		v62 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v62 ^ int32(1)
L2:
	;
	goto L1
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v13 != v14 {
		v62 = v6
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v10 <= int32(0) {
		v62 = int32(1)
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v20 = v10 << (uint(int32(4)) % 32)
	v22 = int32(20)
	v28 = int32(0)
	goto L6
L6:
	;
	v35 = v28 * int32(100)
	v36 = v4 + v20 + v22 + v35
	v37 = int32(4)
	v39 = v35 + (v5 + v20 + v22)
	v42 = F_strcmp(m, v36+v37, v39+v37)
	mBase = m.M
	if v42 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v62 = int32(0)
	goto L2
L8:
	;
	goto L7
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	if v43 != v44 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+76))
	if v46 != v47 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)+96))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+96))
	if v49 != v50 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+91)))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+91)))
	if v52 != v53 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v55 = int32(1)
	v57 = v28 + v55
	if v10 != v57 {
		v28 = v57
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v62 = v55
	goto L2
}
func F_type_is_multirange(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14033(m, l0, int32(109))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_type_is_rowtype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(1)
	if l0 == int32(2249) {
		v52 = v10
		m.G0 = v8 + int32(16)
		return v52
	} else {
		v14 = F_SearchSysCache1(m, int32(82), l0)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v14 == int32(0) {
				v52 = int32(0)
				m.G0 = v8 + int32(16)
				return v52
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v21)+79)))
				F_ReleaseCatCache(m, v14)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					switch v23 - int32(99) {
					case 0:
						v52 = v10
						m.G0 = v8 + int32(16)
						return v52
					case 1:
						v31 = F_getBaseTypeAndTypmod(m, l0, v8+int32(12))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = F_SearchSysCache1(m, int32(82), v31)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								if v33 == int32(0) {
									v52 = int32(0)
									m.G0 = v8 + int32(16)
									return v52
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
									v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
									v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v38)+79)))
									F_ReleaseCatCache(m, v33)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										if v40 == int32(99) {
											v52 = v10
										} else {
											v52 = int32(0)
										}
										m.G0 = v8 + int32(16)
										return v52
									}
								}
							}
						}
					default:
						v52 = int32(0)
						m.G0 = v8 + int32(16)
						return v52
					}
				}
			}
		}
	}
}
