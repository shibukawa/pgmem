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
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	v6 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v10 = F_LWLockAcquire(m, v6+int32(4992), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[68]))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
		if v14 != 0 {
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v14)) == int32(0) {
				v26 = base.B2i32(base.Ui32(v14) < base.Ui32(l0))
			} else {
				v26 = int32(base.Ui32(v14-l0) >> (uint(int32(31)) % 32))
			}
			v28 = *(*int32)(unsafe.Add(mBase, _consts[68]))
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
				v46 = *(*int32)(unsafe.Add(mBase, _consts[68]))
				v48 = v46
				*(*int32)(unsafe.Add(mBase, uint32(v48)+44)) = l1
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = l0
			v48 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v48)+44)) = l1
		}
		v54 = *(*int32)(unsafe.Add(mBase, _consts[7]))
		F_LWLockRelease(m, v54+int32(4992))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
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
	goto L6
L2:
	;
	if l0 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L3:
	;
	goto L2
L4:
	;
	if v107 <= int32(0) {
		v28 = v106
		v33 = v109
		goto L1
	} else {
		goto L27
	}
L5:
	;
	v103 = v33 + int32(1)
	if v19 == int32(0) {
		v28 = v47
		v33 = v103
		goto L1
	} else {
		goto L26
	}
L6:
	;
	if base.B2i32(v33 < v40) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v106 = v99
	v107 = v92
	v109 = v33
	goto L4
L8:
	;
	if v73 <= v47 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	if v21 == int32(0) {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+v33<<(uint(int32(1))%32)))))
	v73 = v66
	v74 = l4 + v69&int32(16383)
	goto L8
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v47 < v64 {
		v73 = v64
		v74 = int32(2147483647)
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L3
L14:
	;
	if v74 == v92 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	if v19 == int32(0) {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83+v47<<(uint(int32(1))%32)))))
	v90 = l5 + v87&int32(16383)
	if v74 < v90 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	v78 = int32(2147483647)
	if v74 == v78 {
		v92 = v78
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v106 = v47
	v107 = v74
	v109 = v33 + int32(1)
	goto L4
L20:
	;
	v92 = v90
	goto L14
L21:
	;
	v94 = int32(1)
	v95 = v47 + v94
	v97 = v33 + v94
	if base.Ui32(l3) < base.Ui32(int32(4)) {
		v28 = v95
		v33 = v97
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v99 = v47 + int32(1)
	if v21 == int32(0) {
		v47 = v99
		goto L6
	} else {
		goto L25
	}
L24:
	;
	v106 = v95
	v107 = v74
	v109 = v97
	goto L4
L25:
	;
	goto L7
L26:
	;
	v106 = v47
	v107 = v74
	v109 = v103
	goto L4
L27:
	;
	if l0 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	return int32(1)
L29:
	;
	goto L30
L30:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v116 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v119 = F_palloc(m, l6<<(uint(v16)%32))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v126 = v116
	goto L33
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v128 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v127 + v128
	*(*uint16)(unsafe.Add(mBase, uint32(v126+v127<<(uint(v128)%32)))) = uint16(v107)
	v28 = v106
	v33 = v109
	goto L1
L34:
	;
	return int32(0)
L35:
	;
	v123 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v123)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v119
	v126 = v119
	goto L33
L36:
	;
	return int32(0)
L37:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v139 <= int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	return int32(1)
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
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	switch v12 & int32(240) {
	case 0:
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v15
		F_appendStringInfo(m, l0, int32(423015), v8)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	default:
		m.G0 = v8 + int32(32)
		return
	case 16:
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v20
		F_appendStringInfo(m, l0, int32(54136), v8+int32(16))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+108)) = l0
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1144]))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[1145]))
	if v37 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+76)) = int64(601295421444)
	v22 = F_hash_create(m, int32(392271), int32(4), v9+int32(60), int32(40))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1144])) = v22
	F_CacheRegisterSyscacheCallback(m, int32(78), int32(1612), v22)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[367]))
	if v32 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
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
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L59
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L3
	} else {
		goto L56
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L3
	} else {
		goto L53
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L3
	} else {
		goto L50
	}
L12:
	;
	m.G0 = v9 + int32(112)
	return v163
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[1144]))
	v47 = int32(0)
	v49 = F_hash_search(m, v44, v9+int32(108), v47, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L18
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v40 != l0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+4)))
	if v42 != 0 {
		v163 = v37
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1145])) = v156
	v163 = v156
	goto L12
L18:
	;
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)))
	if v51 != 0 {
		v156 = v49
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v9)+108))
	v54 = F_SearchSysCache1(m, int32(78), v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	if v54 == int32(0) {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
	v60 = v58 + v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+72))
	if v61 == int32(0) {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+76))
	if v64 == int32(0) {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+80))
	if v67 == int32(0) {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	if v49 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[1144]))
	v79 = F_hash_search(m, v73, v9+int32(108), int32(1), v9+int32(60))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L31
	}
L29:
	;
	v81 = v49
	goto L30
L30:
	;
	if v81&int32(3) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v81 = v79
	goto L30
L32:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v9)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v60)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v60)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+12)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v60)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v60)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+20)) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v60)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+24)) = v121
	F_ReleaseCatCache(m, v54)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L3
	} else {
		goto L42
	}
L33:
	;
	v87 = v81 + int32(140)
	if base.Ui32(v87) <= base.Ui32(v81) {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v108 = F__emscripten_memset_bulkmem(m, v81+int32(4), base.I32_extend8_s(int32(0)), int32(136))
	mBase = m.M
	goto L41
L36:
	;
	v93 = v81 + int32(4)
	if base.Ui32(v93) < base.Ui32(v87) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v95 = v87
	goto L39
L38:
	;
	v95 = v93
	goto L39
L39:
	;
	v102 = F__emscripten_memset_bulkmem(m, v81, base.I32_extend8_s(int32(0)), (v81^int32(-1)+v95)&int32(-4)+int32(4))
	mBase = m.M
	goto L40
L40:
	;
	goto L32
L41:
	;
	goto L32
L42:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v129 = *(*int32)(unsafe.Add(mBase, _consts[367]))
	F_fmgr_info_cxt(m, v125, v81+int32(28), v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v136 = *(*int32)(unsafe.Add(mBase, _consts[367]))
	F_fmgr_info_cxt(m, v132, v81+int32(56), v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	v143 = *(*int32)(unsafe.Add(mBase, _consts[367]))
	F_fmgr_info_cxt(m, v139, v81+int32(84), v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	if v146 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[367]))
	F_fmgr_info_cxt(m, v146, v81+int32(112), v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L3
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v153 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+4)) = uint8(v153)
	v156 = v81
	goto L17
L49:
	;
	goto L48
L50:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v9)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v175
	F_errmsg_internal(m, int32(43081), v9)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(490277), int32(156), int32(392025))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v9)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v189
	F_errmsg_internal(m, int32(415180), v9+int32(16))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(490277), int32(163), int32(392025))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v9)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v205
	F_errmsg_internal(m, int32(415385), v9+int32(32))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(490277), int32(165), int32(392025))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v9)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v221
	F_errmsg_internal(m, int32(415475), v9+int32(48))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(490277), int32(167), int32(392025))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
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
	v7 = F_DirectFunctionCall3Coll(m, int32(1186), int32(0), v4, v5, v6)
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
		v11 = F_DirectFunctionCall4Coll(m, int32(1188), int32(0), v4, v8, v9, v10)
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
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
	v66 = F_FunctionCall4Coll(m, v25, int32(0), v27, v35, v63, v12+int32(8))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v38 = int32(4)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v40&int32(254) == int32(2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v53 = int32(1)
	if v34 != 0 {
		v63 = int32(base.Ui32(v32)>>(uint(v53)%32)) - v53
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v49 = v38
	goto L13
L12:
	;
	v49 = base.B2i32(v40 == int32(18)) << (uint(v38) % 32)
	goto L13
L13:
	;
	if v40 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v52 = v38
	goto L16
L15:
	;
	v52 = v49
	goto L16
L16:
	;
	v63 = v52
	goto L7
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v63 = int32(base.Ui32(v57)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)))
	if v68 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v71 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v71)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v75&v71 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v115 = v66
	goto L21
L21:
	;
	if v115 != 0 {
		goto L41
	} else {
		goto L42
	}
L22:
	;
	v78 = v29
	goto L24
L23:
	;
	v78 = v31
	goto L24
L24:
	;
	if v75 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v111 = F_FunctionCall4Coll(m, v25, int32(0), v74, v78, v108, v12+int32(8))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L36
	}
L26:
	;
	v81 = int32(4)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v83&int32(254) == int32(2) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v96 = int32(1)
	if v75&v96 != 0 {
		v108 = int32(base.Ui32(v75)>>(uint(v96)%32)) - v96
		goto L25
	} else {
		goto L35
	}
L29:
	;
	v92 = v81
	goto L31
L30:
	;
	v92 = base.B2i32(v83 == int32(18)) << (uint(v81) % 32)
	goto L31
L31:
	;
	if v83 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v95 = v81
	goto L34
L33:
	;
	v95 = v92
	goto L34
L34:
	;
	v108 = v95
	goto L25
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v108 = int32(base.Ui32(v102)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L36:
	;
	if v111 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v113 = v111
	goto L39
L38:
	;
	v113 = v66
	goto L39
L39:
	;
	v115 = v113
	goto L21
L40:
	;
	m.G0 = v12 + int32(16)
	return v220
L41:
	;
	v118 = v115
	goto L44
L42:
	;
	goto L43
L43:
	;
	v213 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v213)
	v220 = int32(0)
	goto L40
L44:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v128 != 0 {
		v118 = v118 + int32(8)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v132 = F_palloc(m, (v118-v115)>>(uint(int32(1))%32))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v134 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v171 = F_construct_array_builtin(m, v132, (v162-v115)>>(uint(int32(3))%32), int32(25))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L56
	}
L49:
	;
	v162 = v115
	goto L48
L50:
	;
	goto L51
L51:
	;
	v138 = v115
	v140 = v134
	goto L52
L52:
	;
	v150 = F_cstring_to_text(m, v140)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	v162 = v156
	goto L48
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132+(v138-v115)>>(uint(int32(1))%32)))) = v150
	v156 = v138 + int32(8)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v138+int32(12))))
	if v157 != 0 {
		v138 = v156
		v140 = v157
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v173 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v177 = v115
	v179 = v115 + int32(4)
	goto L60
L58:
	;
	goto L59
L59:
	;
	F_pfree(m, v115)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L65
	}
L60:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v132+(v177-v115)>>(uint(int32(1))%32))))
	F_pfree(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	goto L59
L62:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	F_pfree(m, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v196 = v177 + int32(12)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	if v199 != 0 {
		v177 = v177 + int32(8)
		v179 = v196
		goto L60
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	F_pfree(m, v132)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v220 = v171
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
			v48 = F_TS_execute_recurse(m, v36, v11, int32(0), int32(1539))
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 float32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
		v19 = l0 + int32(28)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
		v21 = F_pg_detoast_datum(m, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			F_getWeights(m, v14, v11)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = F_calc_rank_cd(m, v11, v21, v24, v23)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v29 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
							if v33 != v21 {
								F_pfree(m, v21)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v37 != v24 {
										F_pfree(m, v24)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int32(0)
										} else {
											m.G0 = v11 + int32(16)
											return base.I32_reinterpret_f32(v27)
										}
									} else {
										m.G0 = v11 + int32(16)
										return base.I32_reinterpret_f32(v27)
									}
								}
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v37 != v24 {
									F_pfree(m, v24)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(16)
										return base.I32_reinterpret_f32(v27)
									}
								} else {
									m.G0 = v11 + int32(16)
									return base.I32_reinterpret_f32(v27)
								}
							}
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						if v33 != v21 {
							F_pfree(m, v21)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v37 != v24 {
									F_pfree(m, v24)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(16)
										return base.I32_reinterpret_f32(v27)
									}
								} else {
									m.G0 = v11 + int32(16)
									return base.I32_reinterpret_f32(v27)
								}
							}
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v37 != v24 {
								F_pfree(m, v24)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									m.G0 = v11 + int32(16)
									return base.I32_reinterpret_f32(v27)
								}
							} else {
								m.G0 = v11 + int32(16)
								return base.I32_reinterpret_f32(v27)
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
		v8 = *(*int32)(unsafe.Add(mBase, _consts[1013]))
		*(*int32)(unsafe.Add(mBase, uint32(v3))) = v8
		v10 = v8
	} else {
		v10 = v4
	}
	*(*int32)(unsafe.Add(mBase, uint32(v3)+24)) = int32(1180)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+28)) = v10 * int32(300)
	return int32(1)
}
