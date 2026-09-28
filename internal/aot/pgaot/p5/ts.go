package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SetCommitTsLimit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_SetCommitTsLimit[0]))
	v10 = F_LWLockAcquire(m, v6+int32(_a_F_SetCommitTsLimit_0), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_SetCommitTsLimit[1]))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
		if v14 != 0 {
			v15 = int32(3)
			if base.B2i32(base.Ui32(l0) < base.Ui32(v15))|base.B2i32(base.Ui32(v14) < base.Ui32(v15)) == int32(0) {
				if v14-l0 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = l0
				} else {
				}
			} else {
				if base.Ui32(l0) <= base.Ui32(v14) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = l0
				}
			}
			v27 = int32(3)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
			if base.B2i32(base.Ui32(l1) < base.Ui32(v27))|base.B2i32(base.Ui32(v29) < base.Ui32(v27)) == int32(0) {
				if l1-v29 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = l1
				} else {
				}
			} else {
				if base.Ui32(v29) <= base.Ui32(l1) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = l1
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = l0
		}
		v44 = *(*int32)(unsafe.Add(mBase, _c_F_SetCommitTsLimit[0]))
		F_LWLockRelease(m, v44+int32(_a_F_SetCommitTsLimit_0))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			return
		}
	}
}
func F_TSConfigIsVisible(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_TSConfigIsVisibleExt(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_TS_phrase_output(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	v8 = int32(0)
	v16 = int32(1)
	v19 = l3 & v16
	v21 = l3 & int32(2)
	v28 = v8
	v33 = v8
	goto L1
L1:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v47 = v28
	goto L8
L3:
	;
	if v116 <= int32(0) {
		v28 = v115
		v33 = v118
		goto L1
	} else {
		goto L29
	}
L4:
	;
	v111 = int32(1)
	v112 = v47 + v111
	v114 = v33 + v111
	if base.Ui32(l3) < base.Ui32(int32(4)) {
		v28 = v112
		v33 = v114
		goto L1
	} else {
		goto L28
	}
L5:
	;
	v104 = int32(2147483647)
	if v71 == v104 {
		v109 = v104
		goto L4
	} else {
		goto L27
	}
L6:
	;
	if l0 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	if v19 != 0 {
		goto L5
	} else {
		goto L23
	}
L8:
	;
	if base.B2i32(v33 < v40) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v115 = v91
	v116 = v83
	v118 = v33
	goto L3
L10:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76+v47<<(uint(int32(1))%32)))))
	v83 = l5 + v80&int32(_a_F_TS_phrase_output_0)
	if v75 < v83 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	if v21 == int32(0) {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66+v33<<(uint(int32(1))%32)))))
	v71 = l4 + v68&int32(_a_F_TS_phrase_output_0)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v72 <= v47 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v47 < v64 {
		v75 = int32(2147483647)
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L6
L16:
	;
	v75 = v71
	goto L10
L17:
	;
	v86 = v33 + int32(1)
	if v19 == int32(0) {
		v28 = v47
		v33 = v86
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v75 == v83 {
		v109 = v75
		goto L4
	} else {
		goto L21
	}
L20:
	;
	v115 = v47
	v116 = v75
	v118 = v86
	goto L3
L21:
	;
	v91 = v47 + int32(1)
	if v21 == int32(0) {
		v47 = v91
		goto L8
	} else {
		goto L22
	}
L22:
	;
	goto L9
L23:
	;
	goto L6
L24:
	;
	return int32(0)
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v97 <= int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	return int32(1)
L27:
	;
	v115 = v47
	v116 = v71
	v118 = v33 + int32(1)
	goto L3
L28:
	;
	v115 = v112
	v116 = v109
	v118 = v114
	goto L3
L29:
	;
	if l0 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return int32(1)
L31:
	;
	goto L32
L32:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v125 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v128 = F_palloc(m, l6<<(uint(v16)%32))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v135 = v125
	goto L35
L35:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v137 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v136 + v137
	*(*uint16)(unsafe.Add(mBase, uint32(v135+v136<<(uint(v137)%32)))) = uint16(v116)
	v28 = v115
	v33 = v118
	goto L1
L36:
	;
	return int32(0)
L37:
	;
	v132 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v132)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v128
	v135 = v128
	goto L35
}
func F_commit_ts_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	v14 = v12 & int32(240)
	if v14 != 0 {
		if v14 == int32(16) {
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v23
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v22
			F_appendStringInfo(m, l0, int32(_a_F_commit_ts_desc_0), v8+int32(16))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	} else {
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v17
		F_appendStringInfo(m, l0, int32(_a_F_commit_ts_desc_1), v8)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	}
}
func F_lookup_ts_parser_cache(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+108)) = l0
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[0]))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[1]))
	if v36 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+64)) = int64(601295421444)
	v20 = F_hash_create(m, int32(_a_F_lookup_ts_parser_cache_6), int64(4), v7+int32(56), int32(40))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[0])) = v20
	F_CacheRegisterSyscacheCallback(m, int32(78), int32(1811), base.I64_extend_i32_u(v20))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[2]))
	if v31 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L3
	} else {
		goto L49
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L3
	} else {
		goto L46
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L43
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L40
	}
L12:
	;
	m.G0 = v7 + int32(112)
	return v136
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[0]))
	v46 = int32(0)
	v48 = F_hash_search(m, v43, v7+int32(108), v46, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L18
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v39 != l0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
	if v41 != 0 {
		v136 = v36
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[1])) = v131
	v136 = v131
	goto L12
L18:
	;
	if v48 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
	if v50 != 0 {
		v131 = v48
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+108)))
	v53 = F_SearchSysCache1(m, int32(78), v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	if v53 == int32(0) {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+22)))
	v59 = v57 + v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+72))
	if v60 == int32(0) {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+76))
	if v63 == int32(0) {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v59)+80))
	if v66 == int32(0) {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	if v48 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[0]))
	v78 = F_hash_search(m, v72, v7+int32(108), int32(1), v7+int32(56))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L3
	} else {
		goto L31
	}
L29:
	;
	v80 = v48
	goto L30
L30:
	;
	base.MemoryFill(m, v80+int32(4), int32(0), int32(136))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v7)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v59)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v59)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v59)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v59)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+20)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v59)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = v96
	F_ReleaseCatCache(m, v53)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L3
	} else {
		goto L32
	}
L31:
	;
	v80 = v78
	goto L30
L32:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[2]))
	F_fmgr_info_cxt(m, v100, v80+int32(28), v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[2]))
	F_fmgr_info_cxt(m, v107, v80+int32(56), v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[2]))
	F_fmgr_info_cxt(m, v114, v80+int32(84), v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	if v121 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[2]))
	F_fmgr_info_cxt(m, v121, v80+int32(112), v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L3
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v128 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v80)+4)) = uint8(v128)
	v131 = v80
	goto L17
L39:
	;
	goto L38
L40:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v7)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v146
	F_errmsg_internal(m, int32(_a_F_lookup_ts_parser_cache_0), v7)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_parser_cache_1), int32(157), int32(_a_F_lookup_ts_parser_cache_2))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v7)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v160
	F_errmsg_internal(m, int32(_a_F_lookup_ts_parser_cache_3), v7+int32(16))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_parser_cache_1), int32(164), int32(_a_F_lookup_ts_parser_cache_2))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v7)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v176
	F_errmsg_internal(m, int32(_a_F_lookup_ts_parser_cache_4), v7+int32(32))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_parser_cache_1), int32(166), int32(_a_F_lookup_ts_parser_cache_2))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v7)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v192
	F_errmsg_internal(m, int32(_a_F_lookup_ts_parser_cache_5), v7+int32(48))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_parser_cache_1), int32(168), int32(_a_F_lookup_ts_parser_cache_2))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ts_headline_byid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v7 = F_DirectFunctionCall3Coll(m, int32(1286), int32(0), v4, v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_ts_headline_jsonb_opt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	v4 = F_getTSCurrentConfig(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		v12 = F_DirectFunctionCall4Coll(m, int32(1288), int32(0), base.I64_extend_i32_u(v4), v9, v10, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			return v12
		}
	}
}
func F_ts_lexize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var __phi175 int32
	_ = __phi175
	var v176 int32
	_ = v176
	var __phi176 int32
	_ = __phi176
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v220 int64
	_ = v220
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
	v21 = F_lookup_ts_dictionary_cache(m, v13)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = v21 + int32(12)
	v26 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v21)+44)))
	v27 = int32(1)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v31 = v29 & v27
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = v27
	goto L6
L5:
	;
	v32 = int32(4)
	goto L6
L6:
	;
	if v29 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v65 = base.I64_extend_i32_u(v11 + int32(8))
	v66 = F_FunctionCall4Coll(m, v24, int32(0), v26, base.I64_extend_i32_u(v15+v32), base.I64_extend_i32_s(v61), v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v40 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v51 = int32(1)
	if v31 != 0 {
		v61 = int32(base.Ui32(v29)>>(uint(v51)%32)) - v51
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v43 = int32(16)
	goto L13
L12:
	;
	v43 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v40-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v50 = int32(4)
	goto L16
L15:
	;
	v50 = v43
	goto L16
L16:
	;
	v61 = v50
	goto L7
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v61 = int32(base.Ui32(v55)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v68 = base.I32_wrap_i64(v66)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+9)))
	if v69 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	m.G0 = v11 + int32(16)
	return v220
L20:
	;
	v72 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v72)
	v75 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v21)+44)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v78&v72 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v119 = v68
	goto L22
L22:
	;
	if v119 != 0 {
		goto L41
	} else {
		goto L42
	}
L23:
	;
	v81 = v72
	goto L25
L24:
	;
	v81 = int32(4)
	goto L25
L25:
	;
	if v78 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v114 = F_FunctionCall4Coll(m, v24, int32(0), v75, base.I64_extend_i32_u(v15+v81), base.I64_extend_i32_s(v112), v65)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L37
	}
L27:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v89 == int32(18) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v100 = int32(1)
	if v78&v100 != 0 {
		v112 = int32(base.Ui32(v78)>>(uint(v100)%32)) - v100
		goto L26
	} else {
		goto L36
	}
L30:
	;
	v92 = int32(16)
	goto L32
L31:
	;
	v92 = int32(0)
	goto L32
L32:
	;
	if base.Ui32((v89-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v99 = int32(4)
	goto L35
L34:
	;
	v99 = v92
	goto L35
L35:
	;
	v112 = v99
	goto L26
L36:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v112 = int32(base.Ui32(v106)>>(uint(int32(2))%32)) - int32(4)
	goto L26
L37:
	;
	v116 = base.I32_wrap_i64(v114)
	if v116 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v117 = v116
	goto L40
L39:
	;
	v117 = v68
	goto L40
L40:
	;
	v119 = v117
	goto L22
L41:
	;
	v121 = v119
	goto L44
L42:
	;
	goto L43
L43:
	;
	v209 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v209)
	v220 = int64(0)
	goto L19
L44:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v131 != 0 {
		v121 = v121 + int32(8)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v136 = F_palloc_mul(m, int32(8), (v121-v119)>>(uint(int32(3))%32))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v138 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v170 = F_construct_array_builtin(m, v136, (v159-v119)>>(uint(int32(3))%32), int32(25))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L56
	}
L49:
	;
	v159 = v119
	goto L48
L50:
	;
	goto L51
L51:
	;
	v142 = v119
	v145 = v138
	goto L52
L52:
	;
	v151 = F_cstring_to_text(m, v145)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	v159 = v156
	goto L48
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v136+(v142-v119)))) = base.I64_extend_i32_u(v151)
	v156 = v142 + int32(8)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	if v157 != 0 {
		v142 = v156
		v145 = v157
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v172 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	__phi175 = v119
	__phi176 = v119 + int32(4)
	v175 = __phi175
	v176 = __phi176
	goto L60
L58:
	;
	goto L59
L59:
	;
	F_pfree(m, v119)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L65
	}
L60:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v136+(v175-v119))))
	F_pfree(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	goto L59
L62:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	F_pfree(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	if v193 != 0 {
		__phi175 = v175 + int32(8)
		__phi176 = v175 + int32(12)
		v175 = __phi175
		v176 = __phi176
		goto L60
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	F_pfree(m, v136)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v220 = base.I64_extend_i32_u(v170)
	goto L19
}
func F_ts_match_vq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v65 int64
	_ = v65
	v9 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
		if v20 == int32(0) {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v15 == v23 {
				v65 = v9
				m.G0 = v12 + int32(16)
				return v65
			} else {
				F_pfree(m, v15)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v27 == v19 {
						v65 = v9
						m.G0 = v12 + int32(16)
						return v65
					} else {
						F_pfree(m, v19)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							v65 = v9
							m.G0 = v12 + int32(16)
							return v65
						}
					}
				}
			}
		} else {
			v31 = int32(8)
			v32 = v15 + v31
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v32
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v36 = v19 + v31
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v36 + v20*int32(12)
			v43 = v32 + v34<<(uint(int32(2))%32)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v43
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v43
			v48 = F_TS_execute_recurse(m, v36, v12, int32(0), int32(1737))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int64(0)
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v50 != v15 {
					F_pfree(m, v15)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int64(0)
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v54 != v19 {
							F_pfree(m, v19)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int64(0)
							} else {
								v65 = base.I64_extend_i32_u(base.B2i32(v48 != int32(0)))
								m.G0 = v12 + int32(16)
								return v65
							}
						} else {
							v65 = base.I64_extend_i32_u(base.B2i32(v48 != int32(0)))
							m.G0 = v12 + int32(16)
							return v65
						}
					}
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v54 != v19 {
						F_pfree(m, v19)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
						} else {
							v65 = base.I64_extend_i32_u(base.B2i32(v48 != int32(0)))
							m.G0 = v12 + int32(16)
							return v65
						}
					} else {
						v65 = base.I64_extend_i32_u(base.B2i32(v48 != int32(0)))
						m.G0 = v12 + int32(16)
						return v65
					}
				}
			}
		}
	}
}
func F_ts_rankcd_wttf(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 float32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			F_getWeights(m, v13, v10)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v24 = F_calc_rank_cd(m, v10, v18, v21, v20)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v26 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int64(0)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v30 != v18 {
								F_pfree(m, v18)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int64(0)
								} else {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v34 != v21 {
										F_pfree(m, v21)
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return int64(0)
										} else {
											m.G0 = v10 + int32(16)
											return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
										}
									} else {
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
									}
								}
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								if v34 != v21 {
									F_pfree(m, v21)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int64(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
									}
								} else {
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v30 != v18 {
							F_pfree(m, v18)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int64(0)
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								if v34 != v21 {
									F_pfree(m, v21)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int64(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
									}
								} else {
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
								}
							}
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							if v34 != v21 {
								F_pfree(m, v21)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int64(0)
								} else {
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
								}
							} else {
								m.G0 = v10 + int32(16)
								return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
							}
						}
					}
				}
			}
		}
	}
}
func F_ts_token_type_byid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	if v5 == int32(0) {
		v8 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			F_tt_setup_firstcall(m, v8, l0, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
				v17 = F_tt_process_call(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					if v17 != int64(0) {
						v21 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						*(*int64)(unsafe.Add(mBase, uint32(v16))) = v21 + int64(1)
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = int32(1)
						return v17
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(2)
							v34 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
							return v17
						}
					}
				}
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
		v17 = F_tt_process_call(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			if v17 != int64(0) {
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
				*(*int64)(unsafe.Add(mBase, uint32(v16))) = v21 + int64(1)
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = int32(1)
				return v17
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(2)
					v34 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
					return v17
				}
			}
		}
	}
}
func F_ts_typanalyze(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	if v4 < int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_ts_typanalyze[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v3))) = v8
		v10 = v8
	} else {
		v10 = v4
	}
	*(*int32)(unsafe.Add(mBase, uint32(v3)+24)) = int32(1280)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+28)) = v10 * int32(300)
	return int64(1)
}
