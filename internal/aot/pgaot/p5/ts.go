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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
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
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v14)) == int32(0) {
				v26 = base.B2i32(base.Ui32(v14) < base.Ui32(l0))
			} else {
				v26 = int32(base.Ui32(v14-l0) >> (uint(int32(31)) % 32))
			}
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_SetCommitTsLimit[1]))
			if v26 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = l0
			} else {
			}
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v30))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l1)) == int32(0) {
				v42 = base.B2i32(base.Ui32(l1) < base.Ui32(v30))
			} else {
				v42 = int32(base.Ui32(l1-v30) >> (uint(int32(31)) % 32))
			}
			if v42 == int32(0) {
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_SetCommitTsLimit[1]))
				v48 = v46
				*(*int32)(unsafe.Add(mBase, uint32(v48)+44)) = l1
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = l0
			v48 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v48)+44)) = l1
		}
		v52 = *(*int32)(unsafe.Add(mBase, _c_F_SetCommitTsLimit[0]))
		F_LWLockRelease(m, v52+int32(_a_F_SetCommitTsLimit_0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
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
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[1]))
	if v35 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+76)) = int64(601295421444)
	v20 = F_hash_create(m, int32(_a_F_lookup_ts_parser_cache_6), int32(4), v7+int32(60), int32(40))
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
	F_CacheRegisterSyscacheCallback(m, int32(78), int32(1597), v20)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[2]))
	if v30 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
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
	v190 = m.ExcPending
	if v190 != 0 {
		goto L3
	} else {
		goto L49
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L3
	} else {
		goto L46
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L43
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L3
	} else {
		goto L40
	}
L12:
	;
	m.G0 = v7 + int32(112)
	return v135
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[0]))
	v45 = int32(0)
	v47 = F_hash_search(m, v42, v7+int32(108), v45, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L18
	}
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v38 != l0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+4)))
	if v40 != 0 {
		v135 = v35
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[1])) = v130
	v135 = v130
	goto L12
L18:
	;
	if v47 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+4)))
	if v49 != 0 {
		v130 = v47
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+108))
	v52 = F_SearchSysCache1(m, int32(78), v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L3
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	if v52 == int32(0) {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+22)))
	v58 = v56 + v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+72))
	if v59 == int32(0) {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+76))
	if v62 == int32(0) {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v58)+80))
	if v65 == int32(0) {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	if v47 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[0]))
	v77 = F_hash_search(m, v71, v7+int32(108), int32(1), v7+int32(60))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
	} else {
		goto L31
	}
L29:
	;
	v79 = v47
	goto L30
L30:
	;
	base.MemoryFill(m, v79+int32(4), int32(0), int32(136))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v7)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v58)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v58)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v58)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v58)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+20)) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v58)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+24)) = v95
	F_ReleaseCatCache(m, v52)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L32
	}
L31:
	;
	v79 = v77
	goto L30
L32:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[2]))
	F_fmgr_info_cxt(m, v99, v79+int32(28), v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[2]))
	F_fmgr_info_cxt(m, v106, v79+int32(56), v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[2]))
	F_fmgr_info_cxt(m, v113, v79+int32(84), v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	if v120 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_parser_cache[2]))
	F_fmgr_info_cxt(m, v120, v79+int32(112), v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v127 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)) = uint8(v127)
	v130 = v79
	goto L17
L39:
	;
	goto L38
L40:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v7)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v145
	F_errmsg_internal(m, int32(_a_F_lookup_ts_parser_cache_0), v7)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_parser_cache_1), int32(156), int32(_a_F_lookup_ts_parser_cache_2))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
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
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v7)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v159
	F_errmsg_internal(m, int32(_a_F_lookup_ts_parser_cache_3), v7+int32(16))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_parser_cache_1), int32(163), int32(_a_F_lookup_ts_parser_cache_2))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
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
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v7)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v175
	F_errmsg_internal(m, int32(_a_F_lookup_ts_parser_cache_4), v7+int32(32))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_parser_cache_1), int32(165), int32(_a_F_lookup_ts_parser_cache_2))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
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
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v7)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v191
	F_errmsg_internal(m, int32(_a_F_lookup_ts_parser_cache_5), v7+int32(48))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_parser_cache_1), int32(167), int32(_a_F_lookup_ts_parser_cache_2))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
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
func F_ts_headline_byid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = F_DirectFunctionCall3Coll(m, int32(1171), int32(0), v4, v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_ts_headline_jsonb_opt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = F_getTSCurrentConfig(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = F_DirectFunctionCall4Coll(m, int32(1173), int32(0), v4, v8, v9, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_ts_lexize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var __phi172 int32
	_ = __phi172
	var v174 int32
	_ = v174
	var __phi174 int32
	_ = __phi174
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
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = F_pg_detoast_datum_packed(m, v15)
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
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(0)
	v22 = F_lookup_ts_dictionary_cache(m, v14)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = v22 + int32(12)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v28 = int32(1)
	v29 = v16 + v28
	v31 = v16 + int32(4)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v34 = v32 & v28
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v35 = v29
	goto L6
L5:
	;
	v35 = v31
	goto L6
L6:
	;
	if v32 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v65 = F_FunctionCall4Coll(m, v25, int32(0), v27, v35, v62, v12+int32(8))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v41 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v52 = int32(1)
	if v34 != 0 {
		v62 = int32(base.Ui32(v32)>>(uint(v52)%32)) - v52
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v44 = int32(16)
	goto L13
L12:
	;
	v44 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v41-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v51 = int32(4)
	goto L16
L15:
	;
	v51 = v44
	goto L16
L16:
	;
	v62 = v51
	goto L7
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v62 = int32(base.Ui32(v56)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)))
	if v67 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v70 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v70)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v74&v70 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v112 = v65
	goto L21
L21:
	;
	if v112 != 0 {
		goto L41
	} else {
		goto L42
	}
L22:
	;
	v77 = v29
	goto L24
L23:
	;
	v77 = v31
	goto L24
L24:
	;
	if v74 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v109 = F_FunctionCall4Coll(m, v25, int32(0), v73, v77, v106, v12+int32(8))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L36
	}
L26:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v83 == int32(18) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v94 = int32(1)
	if v74&v94 != 0 {
		v106 = int32(base.Ui32(v74)>>(uint(v94)%32)) - v94
		goto L25
	} else {
		goto L35
	}
L29:
	;
	v86 = int32(16)
	goto L31
L30:
	;
	v86 = int32(0)
	goto L31
L31:
	;
	if base.Ui32((v83-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v93 = int32(4)
	goto L34
L33:
	;
	v93 = v86
	goto L34
L34:
	;
	v106 = v93
	goto L25
L35:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v106 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L36:
	;
	if v109 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v111 = v109
	goto L39
L38:
	;
	v111 = v65
	goto L39
L39:
	;
	v112 = v111
	goto L21
L40:
	;
	m.G0 = v12 + int32(16)
	return v215
L41:
	;
	v117 = v112
	goto L44
L42:
	;
	goto L43
L43:
	;
	v209 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v209)
	v215 = int32(0)
	goto L40
L44:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v126 != 0 {
		v117 = v117 + int32(8)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v130 = F_palloc(m, (v117-v112)>>(uint(int32(1))%32))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v132 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v167 = F_construct_array_builtin(m, v130, (v154-v112)>>(uint(int32(3))%32), int32(25))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L56
	}
L49:
	;
	v154 = v112
	goto L48
L50:
	;
	goto L51
L51:
	;
	v135 = v112
	v138 = v132
	goto L52
L52:
	;
	v148 = F_cstring_to_text(m, v138)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	v154 = v152
	goto L48
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130+(v135-v112)>>(uint(int32(1))%32)))) = v148
	v152 = v135 + int32(8)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	if v153 != 0 {
		v135 = v152
		v138 = v153
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v169 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	__phi172 = v112 + int32(4)
	__phi174 = v112
	v172 = __phi172
	v174 = __phi174
	goto L60
L58:
	;
	goto L59
L59:
	;
	F_pfree(m, v112)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L65
	}
L60:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v130+(v174-v112)>>(uint(int32(1))%32))))
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
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
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
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	if v193 != 0 {
		__phi172 = v174 + int32(12)
		__phi174 = v174 + int32(8)
		v172 = __phi172
		v174 = __phi174
		goto L60
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	F_pfree(m, v130)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v215 = v167
	goto L40
}
func F_ts_match_vq(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v60 int32
	_ = v60
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
		if v19 == int32(0) {
			v22 = int32(0)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v14 == v23 {
				v60 = v22
				m.G0 = v11 + int32(16)
				return v60
			} else {
				F_pfree(m, v14)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v18 == v27 {
						v60 = v22
						m.G0 = v11 + int32(16)
						return v60
					} else {
						F_pfree(m, v18)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							v60 = v22
							m.G0 = v11 + int32(16)
							return v60
						}
					}
				}
			}
		} else {
			v31 = int32(8)
			v32 = v14 + v31
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v32
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			v36 = v18 + v31
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v36 + v19*int32(12)
			v43 = v32 + v34<<(uint(int32(2))%32)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v43
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v43
			v48 = F_TS_execute_recurse(m, v36, v11, int32(0), int32(1524))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v50 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v54 != v18 {
							F_pfree(m, v18)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								v60 = base.B2i32(v48 != int32(0))
								m.G0 = v11 + int32(16)
								return v60
							}
						} else {
							v60 = base.B2i32(v48 != int32(0))
							m.G0 = v11 + int32(16)
							return v60
						}
					}
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v54 != v18 {
						F_pfree(m, v18)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v60 = base.B2i32(v48 != int32(0))
							m.G0 = v11 + int32(16)
							return v60
						}
					} else {
						v60 = base.B2i32(v48 != int32(0))
						m.G0 = v11 + int32(16)
						return v60
					}
				}
			}
		}
	}
}
func F_ts_rankcd_wttf(m *base.Module, l0 int32) int32 {
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
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			F_getWeights(m, v13, v10)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = F_calc_rank_cd(m, v10, v18, v21, v20)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v26 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v30 != v18 {
								F_pfree(m, v18)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v34 != v21 {
										F_pfree(m, v21)
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return int32(0)
										} else {
											m.G0 = v10 + int32(16)
											return base.I32_reinterpret_f32(v24)
										}
									} else {
										m.G0 = v10 + int32(16)
										return base.I32_reinterpret_f32(v24)
									}
								}
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v34 != v21 {
									F_pfree(m, v21)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I32_reinterpret_f32(v24)
									}
								} else {
									m.G0 = v10 + int32(16)
									return base.I32_reinterpret_f32(v24)
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v30 != v18 {
							F_pfree(m, v18)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v34 != v21 {
									F_pfree(m, v21)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I32_reinterpret_f32(v24)
									}
								} else {
									m.G0 = v10 + int32(16)
									return base.I32_reinterpret_f32(v24)
								}
							}
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v34 != v21 {
								F_pfree(m, v21)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(16)
									return base.I32_reinterpret_f32(v24)
								}
							} else {
								m.G0 = v10 + int32(16)
								return base.I32_reinterpret_f32(v24)
							}
						}
					}
				}
			}
		}
	}
}
func F_ts_token_type_byid(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	if v5 == int32(0) {
		v8 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			F_tt_setup_firstcall(m, v8, l0, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
				v17 = F_tt_process_call(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if v17 != 0 {
						v19 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						*(*int64)(unsafe.Add(mBase, uint32(v16))) = v19 + int64(1)
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(1)
						return v17
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = int32(2)
							v32 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
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
			return int32(0)
		} else {
			if v17 != 0 {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
				*(*int64)(unsafe.Add(mBase, uint32(v16))) = v19 + int64(1)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(1)
				return v17
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = int32(2)
					v32 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
					return v17
				}
			}
		}
	}
}
func F_ts_typanalyze(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	if v4 < int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_ts_typanalyze[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v3))) = v8
		v10 = v8
	} else {
		v10 = v4
	}
	*(*int32)(unsafe.Add(mBase, uint32(v3)+24)) = int32(1165)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+28)) = v10 * int32(300)
	return int32(1)
}
