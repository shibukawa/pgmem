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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int64
	_ = v173
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
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
	v42 = F_hash_search(m, v36, v8+int32(-4), int32(0), v8+int32(-56))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L7
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(1819)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(1820)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(17179869188)
	v27 = F_hash_create(m, int32(_a_F_assign_record_type_typmod_3), int64(64), v8+int32(-56), int32(200))
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
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+8)))
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
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	v167 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v163+v166<<(uint(v167)%32))+8)) = v162
	v171 = int32(_a_F_assign_record_type_typmod_2)
	v173 = *(*int64)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[6]))
	v175 = v173 + int64(1)
	*(*int64)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[6])) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v163+v177<<(uint(v167)%32)))) = v175
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[0]))
	v188 = F_hash_search(m, v183, v8+int32(-4), int32(1), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L3
	} else {
		goto L44
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
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[4]))
	if v124 != 0 {
		goto L33
	} else {
		goto L34
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
	v86 = F_mul_size(m, int32(16), v82)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L3
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	v109 = F_CreateTupleDescCopy(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L3
	} else {
		goto L31
	}
L25:
	;
	v89 = int32(1)
	v92 = v66 + v89
	if v66&v92 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v97 = v89 << (uint(int32(32)-base.I32_clz(v92)) % 32)
	goto L28
L27:
	;
	v97 = v92
	goto L28
L28:
	;
	v98 = F_mul_size(m, int32(16), v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v100 = F_repalloc0(m, v83, v86, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[5])) = v97
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[4])) = v100
	goto L24
L31:
	;
	v111 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v111
	v113 = int32(_a_F_assign_record_type_typmod_1)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[3])) = v115 + v111
	*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[4]))
	v162 = v109
	v163 = v121
	goto L12
L32:
	;
	if v122 < v139 {
		v162 = v61
		v163 = v138
		goto L12
	} else {
		goto L37
	}
L33:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[5]))
	v138 = v124
	v139 = v126
	goto L32
L34:
	;
	goto L35
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[2]))
	v131 = F_MemoryContextAllocZero(m, v129, int32(1024))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[5])) = int32(64)
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[4])) = v131
	v138 = v131
	v139 = int32(64)
	goto L32
L37:
	;
	v142 = F_mul_size(m, int32(16), v139)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v145 = int32(1)
	v148 = v122 + v145
	if v148&v122 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v153 = v145 << (uint(int32(32)-base.I32_clz(v148)) % 32)
	goto L41
L40:
	;
	v153 = v148
	goto L41
L41:
	;
	v154 = F_mul_size(m, int32(16), v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v156 = F_repalloc0(m, v138, v142, v154)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[5])) = v153
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[4])) = v156
	v162 = v61
	v163 = v156
	goto L12
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v162
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = v192
	*(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_typmod[1])) = v55
	goto L8
}
func F_findTypeTypmodoutFunction(m *base.Module, l0 int32) int32 {
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v11 = Fn14264(m, l0, int32(_a_F_findTypeTypmodoutFunction_0), int32(2277), int32(_a_F_findTypeTypmodoutFunction_1), int32(_a_F_findTypeTypmodoutFunction_2), int32(2271), int32(2284), int32(_a_F_findTypeTypmodoutFunction_3), int32(2275), int32(23))
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(_a_F_get_type_category_preferred_0), v8)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_get_type_category_preferred_1), int32(3031), int32(_a_F_get_type_category_preferred_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v31 = v29 + v30
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+80)))
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v32)
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+81)))
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v34)
			F_ReleaseCatCache(m, v12)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
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
	v20 = v10 << (uint(int32(3)) % 32)
	v22 = int32(28)
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
	v3 = Fn14396(m, l0, int32(109))
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(1)
	if l0 == int32(2249) {
		v54 = v10
		m.G0 = v8 + int32(16)
		return v54
	} else {
		v15 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				v54 = int32(0)
				m.G0 = v8 + int32(16)
				return v54
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v22)+79)))
				F_ReleaseCatCache(m, v15)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					switch v24 - int32(99) {
					case 0:
						v54 = v10
						m.G0 = v8 + int32(16)
						return v54
					case 1:
						v32 = F_getBaseTypeAndTypmod(m, l0, v8+int32(12))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v35 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v32))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								if v35 == int32(0) {
									v54 = int32(0)
									m.G0 = v8 + int32(16)
									return v54
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
									v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
									v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v40)+79)))
									F_ReleaseCatCache(m, v35)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										if v42 == int32(99) {
											v54 = v10
										} else {
											v54 = int32(0)
										}
										m.G0 = v8 + int32(16)
										return v54
									}
								}
							}
						}
					default:
						v54 = int32(0)
						m.G0 = v8 + int32(16)
						return v54
					}
				}
			}
		}
	}
}
