package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_SetPGVariable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v4 = F_flatten_set_variable_args(m, l0, l1)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v8 = F_superuser(m)
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			if v8 != 0 {
				v10 = int32(5)
			} else {
				v10 = int32(6)
			}
			F_set_config_option(m, l0, v4, v10, int32(13), l2, int32(1))
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_create_pg_locale_builtin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 == int32(100) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L7
	} else {
		goto L34
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L7
	} else {
		goto L31
	}
L3:
	;
	v35 = F_text_to_cstring(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L14
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v14 = F_SearchSysCache1(m, int32(21), v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v25 = F_SearchSysCache1(m, int32(16), l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L11
	}
L7:
	;
	return int32(0)
L8:
	;
	if v14 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v22 = F_SysCacheGetAttrNotNull(m, int32(21), v14, int32(15))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v33 = v14
	v34 = v22
	goto L3
L11:
	;
	if v25 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v31 = F_SysCacheGetAttrNotNull(m, int32(16), v25, int32(10))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v33 = v25
	v34 = v31
	goto L3
L14:
	;
	F_ReleaseCatCache(m, v33)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	goto L16
L16:
	;
	v42 = F_builtin_validate_locale(m, v41, v35)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v45 = F_MemoryContextAllocZero(m, l1, int32(20))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v47 = F_MemoryContextStrdup(m, l1, v35)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v47
	v50 = int32(495013)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1117])))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v54 == int32(0) {
		v73 = v53
		v74 = v54
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v76 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+2)) = uint8(v76)
	v78 = int32(354)
	*(*uint16)(unsafe.Add(mBase, uint32(v45))) = uint16(v78)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+16)) = uint8(base.B2i32(v74-v73 == int32(0)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v83 == int32(67) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	goto L20
L22:
	;
	if v53 != v54 {
		v73 = v53
		v74 = v54
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v58 = v35
	v59 = v50
	goto L24
L24:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v63 == int32(0) {
		v73 = v62
		v74 = v63
		goto L21
	} else {
		goto L26
	}
L25:
	;
	v73 = v62
	v74 = v63
	goto L21
L26:
	;
	v66 = int32(1)
	if v62 == v63 {
		v58 = v58 + v66
		v59 = v59 + v66
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	v88 = v86
	goto L30
L29:
	;
	v88 = int32(1)
	goto L30
L30:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+3)) = uint8(base.B2i32(v88 == int32(0)))
	m.G0 = v7 + int32(32)
	return v45
L31:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v101
	F_errmsg_internal(m, int32(47107), v7)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(474894), int32(135), int32(262690))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
	F_errmsg_internal(m, int32(43653), v7+int32(16))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(474894), int32(148), int32(262690))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_analyze_and_rewrite_varparams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int64
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int64
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[831])))
	if v15 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(4367496)
	v19 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	v28 = int32(4367512)
	v33 = F___memset(m, int32(4367520), v19, int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[832])) = int32(4)
	*(*int64)(unsafe.Add(mBase, _consts[833])) = int64(3)
	*(*int32)(unsafe.Add(mBase, _consts[834])) = int32(2)
	*(*int64)(unsafe.Add(mBase, _consts[835])) = int64(1)
	goto L5
L2:
	;
	goto L3
L3:
	;
	v68 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F___gettimeofday(m, int32(4367648))
	mBase = m.M
	goto L3
L5:
	;
	v46 = F___memcpy(m, v25, v28, int32(16))
	mBase = m.M
	v47 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[836])) = v49
	*(*int32)(unsafe.Add(mBase, _consts[837])) = v48
	*(*int64)(unsafe.Add(mBase, _consts[838])) = v47
	v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+8)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, _consts[839])) = v49
	*(*int32)(unsafe.Add(mBase, _consts[834])) = v54
	*(*int64)(unsafe.Add(mBase, _consts[835])) = v53
	goto L7
L7:
	;
	v61 = F___syscall_ret(m, v19)
	mBase = m.M
	m.G0 = v25 + int32(16)
	goto L4
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = l1
	F_setup_parse_variable_parameters(m, v68, l2, l3)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+88)) = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v78 != int32(141) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v128 = F_transformStmt(m, v68, v120)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L8
	} else {
		goto L25
	}
L12:
	;
	v120 = v77
	goto L11
L13:
	;
	goto L14
L14:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+68))
	if v81 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v86 = v77
	goto L18
L16:
	;
	v97 = v77
	goto L17
L17:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	if v102 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86)+76))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+68))
	if v92 != 0 {
		v86 = v91
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v97 = v91
	goto L17
L20:
	;
	goto L19
L21:
	;
	v120 = v77
	goto L11
L22:
	;
	goto L23
L23:
	;
	v106 = F_palloc0(m, int32(20))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(242)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v112 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+16)) = uint8(v112)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+8)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = int32(0)
	v120 = v106
	goto L11
L25:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v128)+156)) = v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v128)+160)) = v132
	F_check_variable_parameters(m, v68, v128)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _consts[334]))
	switch v137 {
	case 0:
		v144 = v5
		goto L27
	case 1:
		goto L28
	default:
		goto L29
	}
L27:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _consts[335]))
	if v146 != 0 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v142 = F_JumbleQuery(m, v128)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L8
	} else {
		goto L31
	}
L29:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, _consts[336])))
	if v139 != int32(1) {
		v144 = v5
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v144 = v142
	goto L27
L32:
	;
	m.T0[v146].(func(*base.Module, int32, int32, int32))(m, v68, v128, v144)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L8
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	F_free_parsestate(m, v68)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L8
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v128)+16))
	v155 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v155 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if int32(0) < v189 {
		goto L43
	} else {
		goto L44
	}
L38:
	;
	goto L37
L39:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v159 != int32(1) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v155)+392))
	if int32(1)&base.B2i32(v164 != int64(0)) != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v168 = int32(4438516)
	v170 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v171 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v170 + v171
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v174 + v171
	*(*int64)(unsafe.Add(mBase, uint32(v155)+392)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v174 + int32(2)
	v185 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v185 - v171
	goto L38
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L8
	} else {
		goto L56
	}
L43:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v197 = int32(0)
	goto L46
L44:
	;
	goto L45
L45:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, _consts[831])))
	if v224 != 0 {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v192+v197<<(uint(int32(2))%32))))
	if v206 == int32(705) {
		goto L42
	} else {
		goto L48
	}
L47:
	;
	goto L45
L48:
	;
	if v206 == int32(0) {
		goto L42
	} else {
		goto L49
	}
L49:
	;
	v212 = v197 + int32(1)
	if v212 != v189 {
		v197 = v212
		goto L46
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	F_ShowUsage(m, int32(501664))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L8
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v228 = F_pg_rewrite_query(m, v128)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L8
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	m.G0 = v12 + int32(16)
	return v228
L56:
	;
	F_errcode(m, int32(134611076))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v197 + int32(1)
	F_errmsg(m, int32(446841), v12)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(472649), int32(842), int32(141781))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_analyze_and_rewrite_withcb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int64
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	v6 = int32(0)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[831])))
	if v9 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = int32(4367496)
	v13 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v22 = int32(4367512)
	v27 = F___memset(m, int32(4367520), v13, int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[832])) = int32(4)
	*(*int64)(unsafe.Add(mBase, _consts[833])) = int64(3)
	*(*int32)(unsafe.Add(mBase, _consts[834])) = int32(2)
	*(*int64)(unsafe.Add(mBase, _consts[835])) = int64(1)
	goto L5
L2:
	;
	goto L3
L3:
	;
	v62 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F___gettimeofday(m, int32(4367648))
	mBase = m.M
	goto L3
L5:
	;
	v40 = F___memcpy(m, v19, v22, int32(16))
	mBase = m.M
	v41 = int64(*(*int32)(unsafe.Add(mBase, uint32(v19))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v43 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[836])) = v43
	*(*int32)(unsafe.Add(mBase, _consts[837])) = v42
	*(*int64)(unsafe.Add(mBase, _consts[838])) = v41
	v47 = int64(*(*int32)(unsafe.Add(mBase, uint32(v19)+8)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, _consts[839])) = v43
	*(*int32)(unsafe.Add(mBase, _consts[834])) = v48
	*(*int64)(unsafe.Add(mBase, _consts[835])) = v47
	goto L7
L7:
	;
	v55 = F___syscall_ret(m, v13)
	mBase = m.M
	m.G0 = v19 + int32(16)
	goto L4
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+88)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = l1
	m.T0[l2].(func(*base.Module, int32, int32))(m, v62, l3)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v71 != int32(141) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v115 = F_transformStmt(m, v62, v111)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L8
	} else {
		goto L25
	}
L12:
	;
	v111 = v70
	goto L11
L13:
	;
	goto L14
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+68))
	if v74 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v79 = v70
	goto L18
L16:
	;
	v88 = v70
	goto L17
L17:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	if v91 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)+76))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+68))
	if v83 != 0 {
		v79 = v82
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v88 = v82
	goto L17
L20:
	;
	goto L19
L21:
	;
	v111 = v70
	goto L11
L22:
	;
	goto L23
L23:
	;
	v95 = F_palloc0(m, int32(20))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = int32(242)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v101 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+16)) = uint8(v101)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = int32(0)
	v111 = v95
	goto L11
L25:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+156)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+160)) = v119
	v122 = *(*int32)(unsafe.Add(mBase, _consts[334]))
	switch v122 {
	case 0:
		v129 = v6
		goto L26
	case 1:
		goto L27
	default:
		goto L28
	}
L26:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _consts[335]))
	if v131 != 0 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v127 = F_JumbleQuery(m, v115)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L8
	} else {
		goto L30
	}
L28:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, _consts[336])))
	if v124 != int32(1) {
		v129 = v6
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v129 = v127
	goto L26
L31:
	;
	m.T0[v131].(func(*base.Module, int32, int32, int32))(m, v62, v115, v129)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	F_free_parsestate(m, v62)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v115)+16))
	v140 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v140 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, _consts[831])))
	if v175 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	goto L36
L38:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v144 != int32(1) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v140)+392))
	if int32(1)&base.B2i32(v149 != int64(0)) != 0 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v153 = int32(4438516)
	v155 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v156 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v155 + v156
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v159 + v156
	*(*int64)(unsafe.Add(mBase, uint32(v140)+392)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v159 + int32(2)
	v170 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v170 - v156
	goto L37
L41:
	;
	F_ShowUsage(m, int32(501664))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v181 = F_pg_rewrite_query(m, v115)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L8
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	return v181
}
func F_pg_ascii2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v8
	return v8
L2:
	;
	goto L3
L3:
	;
	v12 = l0
	v13 = l1
	v15 = v4
	goto L5
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(0)
	return v30
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v17 == int32(0) {
		v29 = v13
		v30 = v15
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v29 = v22
	v30 = l2
	goto L4
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v17
	v22 = v13 + int32(4)
	v23 = int32(1)
	v26 = v15 + v23
	if v26 != l2 {
		v12 = v12 + v23
		v13 = v22
		v15 = v26
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
func F_pg_available_extension_versions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
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
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int64
	_ = v277
	var v280 int32
	_ = v280
	var v281 int64
	_ = v281
	var v284 int32
	_ = v284
	var v285 int64
	_ = v285
	var v288 int32
	_ = v288
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v294 int32
	_ = v294
	var v295 int64
	_ = v295
	var v298 int32
	_ = v298
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v476 int32
	_ = v476
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int64
	_ = v630
	var v633 int32
	_ = v633
	var v634 int64
	_ = v634
	var v636 int64
	_ = v636
	var v638 int64
	_ = v638
	var v640 int64
	_ = v640
	var v642 int64
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v848 int32
	_ = v848
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	v2 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(48)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v38 = F_get_extension_control_directories(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v872 + int32(48)
	return int32(0)
L4:
	;
	if v38 == int32(0) {
		v872 = v30
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v42 <= int32(0) {
		v872 = v30
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v46 = v30
	v53 = v38
	v59 = v32
	v60 = v2
	v63 = v2
	goto L7
L7:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v60<<(uint(int32(2))%32))))
	v77 = F_AllocateDir(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v872 = v841
	goto L3
L9:
	;
	v868 = v60 + int32(1)
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v848)+4))
	if v868 < v869 {
		v46 = v841
		v53 = v848
		v59 = v854
		v60 = v868
		v63 = v858
		goto L7
	} else {
		goto L142
	}
L10:
	;
	if v77 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v82 == int32(44) {
		v841 = v46
		v848 = v53
		v854 = v59
		v858 = v63
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v103 = v63
	goto L15
L14:
	;
	goto L13
L15:
	;
	v112 = F_ReadDir(m, v77, v76)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	F_FreeDir(m, v77)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L141
	}
L17:
	;
	if v112 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v115 = v112 + int32(19)
	v119 = F_strlen(m, v115)
	mBase = m.M
	v126 = v119 + int32(1)
	goto L23
L19:
	;
	goto L20
L20:
	;
	goto L16
L21:
	;
	if v138 == int32(0) {
		goto L15
	} else {
		goto L27
	}
L22:
	;
	goto L21
L23:
	;
	v128 = int32(0)
	if v126 == v128 {
		v138 = v128
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v138 = v133
	goto L22
L25:
	;
	v132 = v126 - int32(1)
	v133 = v115 + v132
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v134 != int32(46) {
		v126 = v132
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v141 = int32(287954)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, _consts[382])))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v145 == int32(0) {
		v164 = v144
		v165 = v145
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v165-v164 != 0 {
		goto L15
	} else {
		goto L36
	}
L29:
	;
	goto L28
L30:
	;
	if v144 != v145 {
		v164 = v144
		v165 = v145
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v149 = v138
	v150 = v141
	goto L32
L32:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+1)))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
	if v154 == int32(0) {
		v164 = v153
		v165 = v154
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v164 = v153
	v165 = v154
	goto L29
L34:
	;
	v157 = int32(1)
	if v153 == v154 {
		v149 = v149 + v157
		v150 = v150 + v157
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v167 = F_pstrdup(m, v115)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v172 = F_strlen(m, v167)
	mBase = m.M
	v179 = v172 + int32(1)
	goto L40
L38:
	;
	v192 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v192)
	v195 = F_strstr(m, v167, int32(630249))
	mBase = m.M
	if v195 != 0 {
		goto L15
	} else {
		goto L44
	}
L39:
	;
	goto L38
L40:
	;
	v181 = int32(0)
	if v179 == v181 {
		v191 = v181
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v191 = v186
	goto L39
L42:
	;
	v185 = v179 - int32(1)
	v186 = v167 + v185
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v187 != int32(46) {
		v179 = v185
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v196 = F_makeString(m, v167)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v198 = F_list_member(m, v103, v196)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v198 != 0 {
		goto L15
	} else {
		goto L47
	}
L47:
	;
	v200 = F_lappend(m, v103, v196)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v203 = F_palloc0(m, int32(48))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v205 = F_pstrdup(m, v167)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+36)) = int32(-1)
	v209 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+34)) = uint8(v209)
	v211 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v203)+32)) = uint16(v211)
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = v205
	v214 = F_pstrdup(m, v76)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+8)) = v214
	F_parse_extension_control_file(m, v203, int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v59)+28))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	v222 = F_get_ext_ver_list(m, v203)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v222 == int32(0) {
		v103 = v200
		goto L15
	} else {
		goto L54
	}
L54:
	;
	v226 = int32(0)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v227 <= v226 {
		v103 = v200
		goto L15
	} else {
		goto L55
	}
L55:
	;
	v232 = v227
	v239 = v226
	goto L56
L56:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257+v239<<(uint(int32(2))%32))))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+8)))
	if v262 != int32(1) {
		v810 = v232
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v103 = v200
	goto L15
L58:
	;
	v836 = v239 + int32(1)
	if v836 < v810 {
		v232 = v810
		v239 = v836
		goto L56
	} else {
		goto L140
	}
L59:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	v267 = F_palloc(m, int32(48))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v270 = v203 + int32(40)
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v270)))
	*(*int64)(unsafe.Add(mBase, uint32(v267)+40)) = v271
	v273 = int32(32)
	v274 = v267 + v273
	v276 = v203 + v273
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v276)))
	*(*int64)(unsafe.Add(mBase, uint32(v274))) = v277
	v280 = v203 + int32(24)
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v280)))
	*(*int64)(unsafe.Add(mBase, uint32(v267)+24)) = v281
	v284 = v203 + int32(16)
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v284)))
	*(*int64)(unsafe.Add(mBase, uint32(v267)+16)) = v285
	v288 = v203 + int32(8)
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v288)))
	*(*int64)(unsafe.Add(mBase, uint32(v267)+8)) = v289
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v203)))
	*(*int64)(unsafe.Add(mBase, uint32(v267))) = v291
	F_parse_extension_control_file(m, v267, v265)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v295 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+40)) = v295
	v298 = v46 + int32(32)
	*(*int64)(unsafe.Add(mBase, uint32(v298))) = v295
	*(*int64)(unsafe.Add(mBase, uint32(v46)+24)) = v295
	*(*int64)(unsafe.Add(mBase, uint32(v46)+16)) = v295
	*(*int64)(unsafe.Add(mBase, uint32(v46)+8)) = v295
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v310 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v310
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	v314 = F_cstring_to_text(m, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v314
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+33)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v317
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+34)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = v319
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	*(*int32)(unsafe.Add(mBase, uint32(v298))) = v321
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v267)+28))
	if v323 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v267)+40))
	if v333 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L65:
	;
	v326 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+13)) = uint8(v326)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v330 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v323)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+36)) = v330
	goto L64
L69:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v267)+24))
	if v447 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L70:
	;
	v336 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+14)) = uint8(v336)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	v341 = F_palloc(m, v338<<(uint(int32(2))%32))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v343 = int32(0)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	if v343 < v344 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v347 = v343
	goto L77
L75:
	;
	v389 = v343
	goto L76
L76:
	;
	v417 = F_construct_array_builtin(m, v341, v389, int32(19))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L81
	}
L77:
	;
	v375 = v347 << (uint(int32(2)) % 32)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v379+v375)))
	v382 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	v389 = v386
	goto L76
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v341+v375))) = v382
	v386 = v347 + int32(1)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	if v386 < v387 {
		v347 = v386
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+40)) = v417
	goto L69
L82:
	;
	F_tuplestore_putvalues(m, v221, v220, v46+int32(16), v46+int32(8))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L87
	}
L83:
	;
	v450 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+15)) = uint8(v450)
	goto L82
L84:
	;
	goto L85
L85:
	;
	v452 = F_cstring_to_text(m, v447)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+44)) = v452
	goto L82
L87:
	;
	v461 = int32(0)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v462 <= v461 {
		v810 = v462
		goto L58
	} else {
		goto L88
	}
L88:
	;
	v467 = v462
	v476 = v461
	goto L89
L89:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v492+v476<<(uint(int32(2))%32))))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+8)))
	if v497 != 0 {
		v780 = v467
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v810 = v780
	goto L58
L91:
	;
	v806 = v476 + int32(1)
	if v806 < v780 {
		v467 = v780
		v476 = v806
		goto L89
	} else {
		goto L139
	}
L92:
	;
	v498 = int32(0)
	if v498 < v467 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v503 = v498
	v509 = v498
	v515 = v498
	goto L96
L94:
	;
	v598 = v467
	v602 = v498
	goto L95
L95:
	;
	if v602 != v261 {
		v780 = v598
		goto L91
	} else {
		goto L121
	}
L96:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v530+v503<<(uint(int32(2))%32))))
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+8)))
	if v535 != int32(1) {
		v587 = v509
		v589 = v515
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v598 = v594
	v602 = v587
	goto L95
L98:
	;
	v593 = v503 + int32(1)
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v593 < v594 {
		v503 = v593
		v509 = v587
		v515 = v589
		goto L96
	} else {
		goto L120
	}
L99:
	;
	v538 = int32(1)
	v540 = F_find_update_path(m, v222, v534, v496, v538, v538)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	if v540 == int32(0) {
		v587 = v509
		v589 = v515
		goto L98
	} else {
		goto L101
	}
L101:
	;
	if v509 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v587 = v534
	v589 = v540
	goto L98
L103:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v540)+4))
	if v515 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v546 != v554 {
		v587 = v509
		v589 = v515
		goto L98
	} else {
		goto L110
	}
L105:
	;
	v549 = int32(0)
	if v549 <= v546 {
		v554 = v549
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v515)+4))
	if v546 < v552 {
		goto L102
	} else {
		goto L109
	}
L108:
	;
	goto L102
L109:
	;
	v554 = v552
	goto L104
L110:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v534)))
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	if v561 == int32(0) {
		v580 = v560
		v581 = v561
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if int32(0) <= v581-v580 {
		v587 = v509
		v589 = v515
		goto L98
	} else {
		goto L119
	}
L112:
	;
	goto L111
L113:
	;
	if v560 != v561 {
		v580 = v560
		v581 = v561
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v565 = v556
	v566 = v557
	goto L115
L115:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566)+1)))
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+1)))
	if v570 == int32(0) {
		v580 = v569
		v581 = v570
		goto L112
	} else {
		goto L117
	}
L116:
	;
	v580 = v569
	v581 = v570
	goto L112
L117:
	;
	v573 = int32(1)
	if v569 == v570 {
		v565 = v565 + v573
		v566 = v566 + v573
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	goto L102
L120:
	;
	goto L97
L121:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	v626 = F_palloc(m, int32(48))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v629 = v626 + int32(40)
	v630 = *(*int64)(unsafe.Add(mBase, uint32(v270)))
	*(*int64)(unsafe.Add(mBase, uint32(v629))) = v630
	v633 = v626 + int32(32)
	v634 = *(*int64)(unsafe.Add(mBase, uint32(v276)))
	*(*int64)(unsafe.Add(mBase, uint32(v633))) = v634
	v636 = *(*int64)(unsafe.Add(mBase, uint32(v280)))
	*(*int64)(unsafe.Add(mBase, uint32(v626)+24)) = v636
	v638 = *(*int64)(unsafe.Add(mBase, uint32(v284)))
	*(*int64)(unsafe.Add(mBase, uint32(v626)+16)) = v638
	v640 = *(*int64)(unsafe.Add(mBase, uint32(v288)))
	*(*int64)(unsafe.Add(mBase, uint32(v626)+8)) = v640
	v642 = *(*int64)(unsafe.Add(mBase, uint32(v203)))
	*(*int64)(unsafe.Add(mBase, uint32(v626))) = v642
	F_parse_extension_control_file(m, v626, v624)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	v647 = F_cstring_to_text(m, v646)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v647
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+33)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v650
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+34)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = v652
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633))))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+32)) = v654
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v629)))
	if v656 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+14)) = uint8(v746)
	F_tuplestore_putvalues(m, v221, v220, v46+int32(16), v46+int32(8))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L138
	}
L126:
	;
	v746 = int32(1)
	goto L125
L127:
	;
	goto L128
L128:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v656)+4))
	v663 = F_palloc(m, v660<<(uint(int32(2))%32))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v665 = int32(0)
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v656)+4))
	if v665 < v667 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v670 = v665
	goto L133
L131:
	;
	v712 = v665
	goto L132
L132:
	;
	v740 = F_construct_array_builtin(m, v663, v712, int32(19))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L137
	}
L133:
	;
	v698 = v670 << (uint(int32(2)) % 32)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v656)+12))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v702+v698)))
	v705 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v704)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L135
	}
L134:
	;
	v712 = v709
	goto L132
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v663+v698))) = v705
	v709 = v670 + int32(1)
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v656)+4))
	if v709 < v710 {
		v670 = v709
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+40)) = v740
	v746 = v665
	goto L125
L138:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v780 = v777
	goto L91
L139:
	;
	goto L90
L140:
	;
	goto L57
L141:
	;
	v841 = v46
	v848 = v53
	v854 = v59
	v858 = v103
	goto L9
L142:
	;
	goto L8
}
func F_pg_b64_dec_len(m *base.Module, l0 int32) int32 {
	return l0 * int32(3) >> (uint(int32(2)) % 32)
}
func F_pg_checksum_page(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	v3 = int32(0)
	v37 = m.G0
	v38 = int32(128)
	v39 = v37 - v38
	m.G0 = v39
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	v46 = F__emscripten_memcpy_bulkmem(m, v39, int32(1588816), v38)
	mBase = m.M
	v55 = v3
	for {
		v86 = l0 + v55<<(uint(int32(7))%32)
		v91 = int32(0)
		for {
			v124 = int32(2)
			v125 = v91 << (uint(v124) % 32)
			v126 = v46 + v125
			v128 = *(*int32)(unsafe.Add(mBase, uint32(v86+v125)))
			v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
			v130 = v128 ^ v129
			v131 = int32(16777619)
			v133 = int32(17)
			*(*int32)(unsafe.Add(mBase, uint32(v126))) = v130*v131 ^ int32(base.Ui32(v130)>>(uint(v133)%32))
			v138 = v125 | int32(4)
			v139 = v46 + v138
			v141 = *(*int32)(unsafe.Add(mBase, uint32(v86+v138)))
			v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
			v143 = v141 ^ v142
			*(*int32)(unsafe.Add(mBase, uint32(v139))) = v143*v131 ^ int32(base.Ui32(v143)>>(uint(v133)%32))
			v151 = v91 + v124
			if v151 != int32(32) {
				v91 = v151
				continue
			} else {
				break
			}
			break
		}
		v155 = v55 + int32(1)
		if v155 != int32(64) {
			v55 = v155
			continue
		} else {
			break
		}
		break
	}
	v164 = int32(0)
	for {
		v197 = v46 + v164<<(uint(int32(2))%32)
		v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
		v199 = int32(16777619)
		v201 = int32(17)
		*(*int32)(unsafe.Add(mBase, uint32(v197))) = v198*v199 ^ int32(base.Ui32(v198)>>(uint(v201)%32))
		v205 = int32(4)
		v206 = v197 + v205
		v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
		*(*int32)(unsafe.Add(mBase, uint32(v206))) = v207*v199 ^ int32(base.Ui32(v207)>>(uint(v201)%32))
		v215 = v197 + int32(8)
		v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
		*(*int32)(unsafe.Add(mBase, uint32(v215))) = v216*v199 ^ int32(base.Ui32(v216)>>(uint(v201)%32))
		v224 = v197 + int32(12)
		v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
		*(*int32)(unsafe.Add(mBase, uint32(v224))) = v225*v199 ^ int32(base.Ui32(v225)>>(uint(v201)%32))
		v233 = v164 + v205
		if v233 != int32(32) {
			v164 = v233
			continue
		} else {
			break
		}
		break
	}
	v242 = int32(0)
	for {
		v275 = v46 + v242<<(uint(int32(2))%32)
		v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
		v277 = int32(16777619)
		v279 = int32(17)
		*(*int32)(unsafe.Add(mBase, uint32(v275))) = v276*v277 ^ int32(base.Ui32(v276)>>(uint(v279)%32))
		v283 = int32(4)
		v284 = v275 + v283
		v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
		*(*int32)(unsafe.Add(mBase, uint32(v284))) = v285*v277 ^ int32(base.Ui32(v285)>>(uint(v279)%32))
		v293 = v275 + int32(8)
		v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
		*(*int32)(unsafe.Add(mBase, uint32(v293))) = v294*v277 ^ int32(base.Ui32(v294)>>(uint(v279)%32))
		v302 = v275 + int32(12)
		v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
		*(*int32)(unsafe.Add(mBase, uint32(v302))) = v303*v277 ^ int32(base.Ui32(v303)>>(uint(v279)%32))
		v311 = v242 + v283
		if v311 != int32(32) {
			v242 = v311
			continue
		} else {
			break
		}
		break
	}
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v46)+124))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v46)+120))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v46)+116))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v46)+112))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v46)+108))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v46)+100))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v46)+92))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v46)+88))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v46)+84))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v46)+80))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v46)+76))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v46)+64))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v46)+60))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v46)+56))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v46)+36))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v46)+32))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v46)+28))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v41)
	m.G0 = v46 + int32(128)
	v382 = int32(65535)
	v383 = base.I32_rem_u_s(v314^(v315^(v316^(v317^(v318^(v319^(v320^(v321^(v322^(v323^(v324^(v325^(v326^(v327^(v328^(v329^(v330^(v331^(v332^(v333^(v334^(v335^(v336^(v337^(v338^(v339^(v340^(v341^(v342^(v343^(v344^(l1^v345))))))))))))))))))))))))))))))), v382)
	return (v383 + int32(1)) & v382
}
func F_pg_client_to_server(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, _consts[359]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v6 = F_pg_any_to_server(m, l0, l1, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_pg_collation_for(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = F_get_fn_expr_argtype(m, v9, v2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
			v36 = v2
			m.G0 = v7 + int32(16)
			return v36
		} else {
			v19 = F_type_is_collatable(m, v11)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if base.B2i32(v19 == int32(0))&base.B2i32(v11 != int32(705)) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = F_format_type_be(m, v11)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v48
								F_errmsg(m, int32(178708), v7)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(478044), int32(631), int32(202987))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
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
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v26 == int32(0) {
						v29 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
						v36 = int32(0)
						m.G0 = v7 + int32(16)
						return v36
					} else {
						v32 = F_generate_collation_name(m, v26)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = F_cstring_to_text(m, v32)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v36 = v34
								m.G0 = v7 + int32(16)
								return v36
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_conf_load_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int64)(unsafe.Add(mBase, _consts[1141]))
	v4 = F_Int64GetDatum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_pg_control_recovery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v11 = F_get_call_result_type(m, l0, int32(0), v6+int32(4))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[24]))
			v22 = F_LWLockAcquire(m, v18+int32(1152), int32(1))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _consts[251]))
				v28 = F_get_controlfile(m, v25, v6+int32(3))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _consts[24]))
					F_LWLockRelease(m, v31+int32(1152))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)))
						if v36 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(370020), int32(0))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(478351), int32(181), int32(13163))
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
						} else {
							v39 = *(*int64)(unsafe.Add(mBase, uint32(v28)+136))
							v40 = F_Int64GetDatum(m, v39)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v42 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+11)) = uint8(v42)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v40
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v28)+144))
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v42)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v45
								v49 = *(*int64)(unsafe.Add(mBase, uint32(v28)+152))
								v50 = F_Int64GetDatum(m, v49)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v52)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v50
									v55 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
									v56 = F_Int64GetDatum(m, v55)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v58 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v58)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v56
										v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+168)))
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)) = uint8(v58)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v61
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
										v70 = F_heap_form_tuple(m, v65, v6+int32(16), v6+int32(11))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
											v73 = F_HeapTupleHeaderGetDatum(m, v72)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												m.G0 = v6 + int32(48)
												return v73
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(351051), int32(0))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(478351), int32(173), int32(13163))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
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
func F_pg_convert(m *base.Module, l0 int32) int32 {
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = m.G0
	v28 = v26 + int32(-64)
	m.G0 = v28
	v30 = int32(-1)
	if v18 == int32(0) {
		v105 = v30
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v119 = m.G0
	v121 = v119 + int32(-64)
	m.G0 = v121
	v123 = int32(-1)
	if v111 == int32(0) {
		v198 = v123
		goto L30
	} else {
		goto L31
	}
L4:
	;
	m.G0 = v28 - int32(-64)
	goto L3
L5:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v33 == int32(0) {
		v105 = v30
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v36 = F_strlen(m, v18)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v36) {
		v105 = v30
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v39 = v18
	v40 = v33
	v41 = v28
	goto L8
L8:
	;
	v49 = F_isalnum(m, v40&int32(255))
	mBase = m.M
	if v49 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v66 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v66)
	v70 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28))))
	v71 = int32(1802512)
	v72 = int32(1801872)
	goto L17
L10:
	;
	if base.Ui32((v40-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v62 = v41
	goto L12
L12:
	;
	v64 = v39 + int32(1)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v65 != 0 {
		v39 = v64
		v40 = v65
		v41 = v62
		goto L8
	} else {
		goto L16
	}
L13:
	;
	v58 = v40 | int32(32)
	goto L15
L14:
	;
	v58 = v40
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v58)
	v62 = v41 + int32(1)
	goto L12
L16:
	;
	goto L9
L17:
	;
	v84 = v72 + (v71-v72)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = int32(*(*int8)(unsafe.Add(mBase, uint32(v85))))
	v87 = v70 - v86
	if v87 != 0 {
		v90 = v87
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v105 = v30
	goto L4
L19:
	;
	v94 = base.B2i32(v90 < int32(0))
	if v90 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v88 = F_strcmp(m, v28, v85)
	mBase = m.M
	if v88 != 0 {
		v90 = v88
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v105 = v89
	goto L4
L22:
	;
	v95 = v84 - int32(8)
	goto L24
L23:
	;
	v95 = v71
	goto L24
L24:
	;
	if v90 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v98 = v72
	goto L27
L26:
	;
	v98 = v84 + int32(8)
	goto L27
L27:
	;
	if base.Ui32(v98) <= base.Ui32(v95) {
		v71 = v95
		v72 = v98
		goto L17
	} else {
		goto L28
	}
L28:
	;
	goto L18
L29:
	;
	if int32(0) <= v105 {
		goto L57
	} else {
		goto L58
	}
L30:
	;
	m.G0 = v121 - int32(-64)
	goto L29
L31:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v126 == int32(0) {
		v198 = v123
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v129 = F_strlen(m, v111)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v129) {
		v198 = v123
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v132 = v111
	v133 = v126
	v134 = v121
	goto L34
L34:
	;
	v142 = F_isalnum(m, v133&int32(255))
	mBase = m.M
	if v142 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v159)
	v163 = int32(*(*int8)(unsafe.Add(mBase, uint32(v121))))
	v164 = int32(1802512)
	v165 = int32(1801872)
	goto L43
L36:
	;
	if base.Ui32((v133-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v155 = v134
	goto L38
L38:
	;
	v157 = v132 + int32(1)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v158 != 0 {
		v132 = v157
		v133 = v158
		v134 = v155
		goto L34
	} else {
		goto L42
	}
L39:
	;
	v151 = v133 | int32(32)
	goto L41
L40:
	;
	v151 = v133
	goto L41
L41:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v151)
	v155 = v134 + int32(1)
	goto L38
L42:
	;
	goto L35
L43:
	;
	v177 = v165 + (v164-v165)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v179 = int32(*(*int8)(unsafe.Add(mBase, uint32(v178))))
	v180 = v163 - v179
	if v180 != 0 {
		v183 = v180
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v198 = v123
	goto L30
L45:
	;
	v187 = base.B2i32(v183 < int32(0))
	if v183 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v181 = F_strcmp(m, v121, v178)
	mBase = m.M
	if v181 != 0 {
		v183 = v181
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	v198 = v182
	goto L30
L48:
	;
	v188 = v177 - int32(8)
	goto L50
L49:
	;
	v188 = v164
	goto L50
L50:
	;
	if v183 < int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v191 = v165
	goto L53
L52:
	;
	v191 = v177 + int32(8)
	goto L53
L53:
	;
	if base.Ui32(v191) <= base.Ui32(v188) {
		v164 = v188
		v165 = v191
		goto L43
	} else {
		goto L54
	}
L54:
	;
	goto L44
L55:
	;
	F_report_invalid_encoding(m, v105, v244+v250, v238-v250)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L115
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L111
	}
L57:
	;
	if v198 < int32(0) {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L107
	}
L60:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v208 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v239 = int32(1)
	if v208&v239 != 0 {
		goto L72
	} else {
		goto L73
	}
L62:
	;
	v211 = int32(4)
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v213&int32(254) == int32(2) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	v226 = int32(1)
	if v208&v226 != 0 {
		v238 = int32(base.Ui32(v208)>>(uint(v226)%32)) - v226
		goto L61
	} else {
		goto L71
	}
L65:
	;
	v222 = v211
	goto L67
L66:
	;
	v222 = base.B2i32(v213 == int32(18)) << (uint(v211) % 32)
	goto L67
L67:
	;
	if v213 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v225 = v211
	goto L70
L69:
	;
	v225 = v222
	goto L70
L70:
	;
	v238 = v225
	goto L61
L71:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v238 = int32(base.Ui32(v232)>>(uint(int32(2))%32)) - int32(4)
	goto L61
L72:
	;
	v243 = v239
	goto L74
L73:
	;
	v243 = int32(4)
	goto L74
L74:
	;
	v244 = v14 + v243
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v105*int32(28))+uint32(_consts[354])))
	v250 = m.T0[v249].(func(*base.Module, int32, int32) int32)(m, v244, v238)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	if v238 != v250 {
		goto L55
	} else {
		goto L76
	}
L76:
	;
	v253 = F_pg_do_encoding_conversion(m, v244, v238, v105, v198)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	m.G0 = v11 + int32(32)
	return v330
L78:
	;
	if v244 == v253 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v330 = v14
	goto L77
L80:
	;
	goto L81
L81:
	;
	if v253&int32(3) == int32(0) {
		v279 = v253
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v314 = v312 + int32(4)
	v315 = F_palloc(m, v314)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L99
	}
L83:
	;
	v312 = v304 - v253
	goto L82
L84:
	;
	v283 = v279
	goto L93
L85:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	if v263 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v312 = int32(0)
	goto L82
L87:
	;
	goto L88
L88:
	;
	v268 = v253
	goto L89
L89:
	;
	v272 = v268 + int32(1)
	if v272&int32(3) == int32(0) {
		v279 = v272
		goto L84
	} else {
		goto L91
	}
L90:
	;
	v304 = v272
	goto L83
L91:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	if v277 != 0 {
		v268 = v272
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v292 = int32(-2139062144)
	if (int32(16843008)-v289|v289)&v292 == v292 {
		v283 = v283 + int32(4)
		goto L93
	} else {
		goto L95
	}
L94:
	;
	v298 = v283
	goto L96
L95:
	;
	goto L94
L96:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if v302 != 0 {
		v298 = v298 + int32(1)
		goto L96
	} else {
		goto L98
	}
L97:
	;
	v304 = v298
	goto L83
L98:
	;
	goto L97
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315))) = v314 << (uint(int32(2)) % 32)
	if v312 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	F_pfree(m, v253)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L104
	}
L101:
	;
	v322 = F__emscripten_memcpy_bulkmem(m, v315+int32(4), v253, v312)
	mBase = m.M
	goto L103
L102:
	;
	goto L103
L103:
	;
	goto L100
L104:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v14 == v326 {
		v330 = v315
		goto L77
	} else {
		goto L105
	}
L105:
	;
	F_pfree(m, v14)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v330 = v315
	goto L77
L107:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v18
	F_errmsg(m, int32(674702), v11)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(472485), int32(578), int32(77321))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v111
	F_errmsg(m, int32(674663), v11+int32(16))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(472485), int32(583), int32(77321))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_ctype_get_cache(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	v8 = *(*int32)(unsafe.Add(mBase, _consts[638]))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _consts[639]))
	v13 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	v32 = F_emscripten_builtin_malloc(m, int32(40))
	mBase = m.M
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 != l0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	if v24 != 0 {
		v13 = v24
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v19 != v10 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	return v13 + int32(8)
L9:
	;
	goto L5
L10:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	F_emscripten_builtin_free(m, v301)
	mBase = m.M
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	F_emscripten_builtin_free(m, v303)
	mBase = m.M
	F_emscripten_builtin_free(m, v32)
	mBase = m.M
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = l0
	v35 = *(*int32)(unsafe.Add(mBase, _consts[639]))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = int64(549755813888)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v35
	v39 = int32(512)
	v40 = F_emscripten_builtin_malloc(m, v39)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v32)+20)) = int64(274877906944)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v40
	v45 = F_emscripten_builtin_malloc(m, v39)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v45
	if v40 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v293 = int32(0)
	goto L13
L13:
	;
	return v293
L14:
	;
	if v45 == int32(0) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = l1
	v56 = *(*int32)(unsafe.Add(mBase, _consts[632]))
	switch v56 - int32(1) {
	case 0, 1:
		v64 = int32(2048)
		goto L16
	case 2:
		goto L18
	default:
		v60 = int32(128)
		goto L17
	}
L16:
	;
	v66 = v32 + int32(8)
	v70 = int32(0)
	v72 = int32(0)
	goto L20
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = int32(-1)
	v64 = v60
	goto L16
L18:
	;
	v60 = int32(256)
	goto L17
L19:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v240 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L20:
	;
	v74 = m.T0[l0].(func(*base.Module, int32) int32)(m, v70)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	if v159 <= int32(0) {
		goto L19
	} else {
		goto L50
	}
L22:
	;
	v161 = v70 + int32(1)
	if v161 != v64 {
		v70 = v161
		v72 = v159
		goto L20
	} else {
		goto L49
	}
L23:
	;
	return int32(0)
L24:
	;
	if v74 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v159 = v72 + int32(1)
	goto L22
L26:
	;
	goto L27
L27:
	;
	if v72 <= int32(0) {
		v159 = v72
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v82 = v70 - v72
	if base.Ui32(int32(2)) <= base.Ui32(v72) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v152 == int32(0) {
		goto L10
	} else {
		goto L47
	}
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v87 < v88 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L32
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v123 < v124 {
		goto L41
	} else {
		goto L42
	}
L33:
	;
	v105 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v104+v103<<(uint(v105)%32)))) = v82
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v115 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v109+v110<<(uint(v105)%32))+4)) = v82 + v72 - v115
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v118 + v115
	v152 = v115
	goto L29
L34:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v103 = v87
	v104 = v90
	goto L33
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v88 << (uint(int32(1)) % 32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v97 = F_emscripten_builtin_realloc(m, v94, v88<<(uint(int32(4))%32))
	mBase = m.M
	if v97 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v152 = int32(0)
	goto L29
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v97
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v103 = v102
	v104 = v97
	goto L33
L40:
	;
	v141 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v139 + v141
	*(*int32)(unsafe.Add(mBase, uint32(v140+v139<<(uint(int32(2))%32)))) = v82
	v152 = v141
	goto L29
L41:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v139 = v123
	v140 = v126
	goto L40
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v124 << (uint(int32(1)) % 32)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v133 = F_emscripten_builtin_realloc(m, v130, v124<<(uint(int32(3))%32))
	mBase = m.M
	if v133 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v152 = int32(0)
	goto L29
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v133
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v139 = v138
	v140 = v133
	goto L40
L47:
	;
	v157 = v70 + int32(1)
	if v157 != v64 {
		v70 = v157
		v72 = int32(0)
		goto L20
	} else {
		goto L48
	}
L48:
	;
	goto L19
L49:
	;
	goto L21
L50:
	;
	v165 = v64 - v159
	if base.Ui32(int32(2)) <= base.Ui32(v159) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v235 == int32(0) {
		goto L10
	} else {
		goto L69
	}
L52:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v170 < v171 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L54
L54:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v206 < v207 {
		goto L63
	} else {
		goto L64
	}
L55:
	;
	v188 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v187+v186<<(uint(v188)%32)))) = v165
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v198 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v192+v193<<(uint(v188)%32))+4)) = v165 + v159 - v198
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v201 + v198
	v235 = v198
	goto L51
L56:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v186 = v170
	v187 = v173
	goto L55
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v171 << (uint(int32(1)) % 32)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v180 = F_emscripten_builtin_realloc(m, v177, v171<<(uint(int32(4))%32))
	mBase = m.M
	if v180 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v235 = int32(0)
	goto L51
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v180
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v186 = v185
	v187 = v180
	goto L55
L62:
	;
	v224 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v222 + v224
	*(*int32)(unsafe.Add(mBase, uint32(v223+v222<<(uint(int32(2))%32)))) = v165
	v235 = v224
	goto L51
L63:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v222 = v206
	v223 = v209
	goto L62
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v207 << (uint(int32(1)) % 32)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v216 = F_emscripten_builtin_realloc(m, v213, v207<<(uint(int32(3))%32))
	mBase = m.M
	if v216 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v235 = int32(0)
	goto L51
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v216
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v222 = v221
	v223 = v216
	goto L62
L69:
	;
	goto L19
L70:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	if v261 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v255
	goto L70
L72:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	F_emscripten_builtin_free(m, v243)
	mBase = m.M
	v245 = int32(0)
	v255 = v245
	v256 = v245
	goto L71
L73:
	;
	goto L74
L74:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v247 <= v240 {
		goto L70
	} else {
		goto L75
	}
L75:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v252 = F_emscripten_builtin_realloc(m, v249, v240<<(uint(int32(2))%32))
	mBase = m.M
	if v252 == int32(0) {
		goto L10
	} else {
		goto L76
	}
L76:
	;
	v255 = v252
	v256 = v240
	goto L71
L77:
	;
	v282 = int32(4353760)
	v283 = *(*int32)(unsafe.Add(mBase, _consts[638]))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v283
	*(*int32)(unsafe.Add(mBase, _consts[638])) = v32
	v293 = v66
	goto L13
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v276
	goto L77
L79:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	F_emscripten_builtin_free(m, v264)
	mBase = m.M
	v266 = int32(0)
	v276 = v266
	v277 = v266
	goto L78
L80:
	;
	goto L81
L81:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v268 <= v261 {
		goto L77
	} else {
		goto L82
	}
L82:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v273 = F_emscripten_builtin_realloc(m, v270, v261<<(uint(int32(3))%32))
	mBase = m.M
	if v273 == int32(0) {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	v276 = v273
	v277 = v261
	goto L78
}
func F_pg_database_collation_actual_version(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_SearchSysCache1(m, int32(21), v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v18 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15+v16)+76)))
			if v18 == int32(99) {
				v24 = int32(13)
			} else {
				v24 = int32(15)
			}
			v25 = F_SysCacheGetAttrNotNull(m, int32(21), v11, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = F_text_to_cstring(m, v25)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = F_get_collation_actual_version(m, v18, v27)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v11)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							if v29 != 0 {
								v33 = F_cstring_to_text(m, v29)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									v38 = v33
									m.G0 = v7 + int32(16)
									return v38
								}
							} else {
								v35 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
								v38 = int32(0)
								m.G0 = v7 + int32(16)
								return v38
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
					F_errmsg(m, int32(66248), v7)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(472768), int32(2789), int32(259534))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
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
}
func F_pg_database_size_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_get_database_oid(m, v3, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_calculate_database_size(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			if v9 == int64(0) {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				return int32(0)
			} else {
				v17 = F_Int64GetDatum(m, v9)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
	}
}
func F_pg_ddl_command_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(408612)
			F_errmsg(m, int32(182827), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(472671), int32(359), int32(407560))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_pg_detoast_datum_packed(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v3 != int32(1))&base.B2i32(v3&int32(3) != int32(2)) != 0 {
		v15 = l0
		return v15
	} else {
		v11 = F_detoast_attr(m, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = v11
			return v15
		}
	}
}
func F_pg_encrypt(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_pfree(m, v141)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L2
	} else {
		goto L93
	}
L2:
	;
	return int32(0)
L3:
	;
	v19 = int32(1)
	v20 = v15 + v19
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v25 = v23 & v19
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v20
	goto L6
L5:
	;
	v26 = v15 + int32(4)
	goto L6
L6:
	;
	if v23 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v56 = F_downcase_truncate_identifier(m, v26, v54, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L18
	}
L8:
	;
	v29 = int32(4)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v31&int32(254) == int32(2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v44 = int32(1)
	if v25 != 0 {
		v54 = int32(base.Ui32(v23)>>(uint(v44)%32)) - v44
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v40 = v29
	goto L13
L12:
	;
	v40 = base.B2i32(v31 == int32(18)) << (uint(v29) % 32)
	goto L13
L13:
	;
	if v31 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v43 = v29
	goto L16
L15:
	;
	v43 = v40
	goto L16
L16:
	;
	v54 = v43
	goto L7
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v60 = F_px_find_combo(m, v56, v12+int32(28))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v60 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_pfree(m, v56)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L75
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v68 = F_pg_detoast_datum_packed(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v71 = F_pg_detoast_datum_packed(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v73 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v104 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L27:
	;
	v76 = int32(4)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	if v78&int32(254) == int32(2) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v91 = int32(1)
	if v73&v91 != 0 {
		v103 = int32(base.Ui32(v73)>>(uint(v91)%32)) - v91
		goto L26
	} else {
		goto L36
	}
L30:
	;
	v87 = v76
	goto L32
L31:
	;
	v87 = base.B2i32(v78 == int32(18)) << (uint(v76) % 32)
	goto L32
L32:
	;
	if v78 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v90 = v76
	goto L35
L34:
	;
	v90 = v87
	goto L35
L35:
	;
	v103 = v90
	goto L26
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v103 = int32(base.Ui32(v97)>>(uint(int32(2))%32)) - int32(4)
	goto L26
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v136 = m.T0[v135].(func(*base.Module, int32, int32) int32)(m, v66, v103)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L48
	}
L38:
	;
	v107 = int32(4)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v109&int32(254) == int32(2) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v122 = int32(1)
	if v104&v122 != 0 {
		v134 = int32(base.Ui32(v104)>>(uint(v122)%32)) - v122
		goto L37
	} else {
		goto L47
	}
L41:
	;
	v118 = v107
	goto L43
L42:
	;
	v118 = base.B2i32(v109 == int32(18)) << (uint(v107) % 32)
	goto L43
L43:
	;
	if v109 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v121 = v107
	goto L46
L45:
	;
	v121 = v118
	goto L46
L46:
	;
	v134 = v121
	goto L37
L47:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v134 = int32(base.Ui32(v128)>>(uint(int32(2))%32)) - int32(4)
	goto L37
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v136
	v141 = F_palloc(m, v136+int32(4))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v143 = int32(1)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v145&v143 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v148 = v143
	goto L52
L51:
	;
	v148 = int32(4)
	goto L52
L52:
	;
	v150 = int32(0)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v153 = m.T0[v152].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v66, v71+v148, v134, v150, v150)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	if v153 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v157 = int32(1)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v159&v157 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v171 = v153
	goto L56
L56:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	m.T0[v172].(func(*base.Module, int32))(m, v66)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L61
	}
L57:
	;
	v162 = v157
	goto L59
L58:
	;
	v162 = int32(4)
	goto L59
L59:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v169 = m.T0[v168].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v66, v68+v162, v103, v141+int32(4), v12+int32(28))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v171 = v169
	goto L56
L61:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v175 != v68 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_pfree(m, v68)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v179 != v71 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L64
L66:
	;
	F_pfree(m, v71)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v183 != v15 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	F_pfree(m, v15)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L2
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v171 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v187<<(uint(int32(2))%32) + int32(16)
	m.G0 = v12 + int32(32)
	return v141
L75:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	if v60 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v56
	F_errmsg(m, int32(195973), v12+int32(16))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L2
	} else {
		goto L91
	}
L78:
	;
	v237 = int32(301198)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v214 = int32(4337152)
	goto L82
L81:
	;
	v237 = v232
	goto L77
L82:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v214)+8))
	if v60 != v217 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	v232 = v229
	goto L81
L84:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v214)+20))
	if v220 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	goto L83
L87:
	;
	v237 = int32(395700)
	goto L77
L88:
	;
	goto L89
L89:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
	if v60 != v225 {
		v214 = v214 + int32(16)
		goto L82
	} else {
		goto L90
	}
L90:
	;
	v232 = v220
	goto L81
L91:
	;
	F_errfinish(m, int32(474523), int32(513), int32(216986))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L2
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	F_errcode(m, int32(579))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	if v171 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v296
	F_errmsg(m, int32(191048), v12)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L2
	} else {
		goto L110
	}
L97:
	;
	v296 = int32(301198)
	goto L96
L98:
	;
	goto L99
L99:
	;
	v273 = int32(4337152)
	goto L101
L100:
	;
	v296 = v291
	goto L96
L101:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)+8))
	if v171 != v276 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	v291 = v288
	goto L100
L103:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v273)+20))
	if v279 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	goto L102
L106:
	;
	v296 = int32(395700)
	goto L96
L107:
	;
	goto L108
L108:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v273)+16))
	if v171 != v284 {
		v273 = v273 + int32(16)
		goto L101
	} else {
		goto L109
	}
L109:
	;
	v291 = v279
	goto L100
L110:
	;
	F_errfinish(m, int32(474523), int32(290), int32(79127))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_euccn_dsplen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v2 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if v2 < int32(0) {
		return int32(2)
	} else {
		v7 = int32(-1)
		if v2 == int32(127) {
			v12 = v7
		} else {
			v12 = int32(1)
		}
		if base.Ui32(v2) < base.Ui32(int32(32)) {
			v15 = v7
		} else {
			v15 = v12
		}
		if v2 != 0 {
			v17 = v15
		} else {
			v17 = int32(0)
		}
		return v17
	}
}
func F_pg_extension_config_dump(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	v12 = m.G0
	v14 = v12 - int32(176)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[295])))
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L85
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L82
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L79
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L76
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L73
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L69
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L65
	}
L10:
	;
	v24 = F_get_rel_name(m, v16)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L61
	}
L13:
	;
	if v24 == int32(0) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v29 = F_getExtensionOfObject(m, int32(1259), v16)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[327]))
	if v29 != v32 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v36 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[327]))
	F_ScanKeyInit(m, v14+int32(128), int32(1), int32(3), int32(184), v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v48 = int32(1)
	v53 = F_systable_beginscan(m, v36, int32(3080), v48, int32(0), v48, v14+int32(128))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v55 = F_systable_getnext(m, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v55 == int32(0) {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v59 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+104)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = int64(281474976710656)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v16
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	v76 = F_heap_getattr_7(m, v55, int32(7), v73, v14+int32(119))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+119)))
	if v78 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v161
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	v168 = F_heap_getattr_7(m, v55, int32(8), v165, v14+int32(119))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L43
	}
L24:
	;
	v81 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+120)) = v81
	v87 = F_construct_array_builtin(m, v14+int32(124), v81, int32(26))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v89 = F_pg_detoast_datum(m, v76)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v153 = int32(0)
	v161 = v87
	goto L23
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v91 != int32(1) {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v94 != int32(1) {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v97 < int32(0) {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v100 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v101 != int32(26) {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+120)) = v97 + int32(1)
	if v97 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v14)+124))
	v148 = F_array_set(m, v89, v14+int32(120), v145, int32(4), int32(1))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L42
	}
L35:
	;
	v112 = int32(0)
	goto L36
L36:
	;
	v124 = v112 + int32(1)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v89+int32(24)+v112<<(uint(int32(2))%32))))
	if v16 == v128 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L34
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+120)) = v124
	goto L34
L39:
	;
	goto L40
L40:
	;
	if v124 != v97 {
		v112 = v124
		goto L36
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	v153 = v97
	v161 = v148
	goto L23
L43:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+119)))
	if v170 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v202 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+71)) = uint8(v202)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v201
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	v212 = F_heap_modify_tuple(m, v55, v205, v14+int32(80), v14+int32(72), v14-int32(-64))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L57
	}
L45:
	;
	if v153 != 0 {
		goto L5
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v179 = F_pg_detoast_datum(m, v168)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	v177 = F_construct_array_builtin(m, v14+int32(124), int32(1), int32(25))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v201 = v177
	goto L44
L50:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	if v181 != int32(1) {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	if v184 != int32(1) {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	if v187 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	if v188 != int32(25) {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	if v191 != v153 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v14)+124))
	v198 = F_array_set(m, v179, v14+int32(120), v195, int32(-1), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v201 = v198
	goto L44
L57:
	;
	F_CatalogTupleUpdate(m, v36, v212+int32(4), v212)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_systable_endscan(m, v53)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_sequence_close(m, v36, int32(3))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	m.G0 = v14 + int32(176)
	return int32(0)
L61:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(643902)
	F_errmsg(m, int32(507688), v14+int32(48))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(474855), int32(2819), int32(224281))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
	F_errmsg(m, int32(377540), v14)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(474855), int32(2830), int32(224281))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v24
	F_errmsg(m, int32(428843), v14+int32(32))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(474855), int32(2836), int32(224281))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _consts[327]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v286
	F_errmsg_internal(m, int32(44590), v14+int32(16))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(474855), int32(2861), int32(224281))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errmsg_internal(m, int32(24315), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(474855), int32(2894), int32(224281))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errmsg_internal(m, int32(310503), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(474855), int32(2927), int32(224281))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errmsg_internal(m, int32(23141), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(474855), int32(2939), int32(224281))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_errmsg_internal(m, int32(310503), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(474855), int32(2941), int32(224281))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_file_exists(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	v13 = F___fstatat(m, int32(-100), l0, v7+int32(16), v2)
	mBase = m.M
	if v13 == int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
		v44 = base.B2i32(v16&int32(61440) != int32(16384))
		m.G0 = v7 + int32(112)
		return v44
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _consts[137]))
		switch v22 - int32(44) {
		case 0, 10:
			v44 = v2
			m.G0 = v7 + int32(112)
			return v44
		case 1, 2, 3, 4, 5, 6, 7, 8, 9:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg(m, int32(284451), v7)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(478009), int32(514), int32(108232))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		default:
			if v22 == int32(2) {
				v44 = v2
				m.G0 = v7 + int32(112)
				return v44
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg(m, int32(284451), v7)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(478009), int32(514), int32(108232))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
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
	}
}
func F_pg_gen_salt_rounds(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	v5 = m.G0
	v7 = v5 - int32(160)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_text_to_cstring_buffer(m, v10, v7+int32(16), int32(129))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = v7 + int32(16)
	v24 = F_px_gen_salt(m, v21, v21, v14)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if int32(0) <= v24 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v28 != v10 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	F_pfree(m, v10)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v34 = F_cstring_to_text_with_len(m, v7+int32(16), v24)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	m.G0 = v7 + int32(160)
	return v34
L13:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v24 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v80
	F_errmsg(m, int32(189983), v7)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L29
	}
L16:
	;
	v80 = int32(301198)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v57 = int32(4337152)
	goto L20
L19:
	;
	v80 = v75
	goto L15
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v24 != v60 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v75 = v72
	goto L19
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	if v63 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	v80 = int32(395700)
	goto L15
L26:
	;
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	if v24 != v68 {
		v57 = v57 + int32(16)
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v75 = v63
	goto L19
L29:
	;
	F_errfinish(m, int32(474523), int32(202), int32(162751))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_constraintdef(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_get_constraintdef_worker(m, v3, int32(0), int32(2), int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		} else {
			v17 = F_cstring_to_text(m, v7)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v7)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
	}
}
func F_pg_get_indexdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int64
	_ = v323
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v449 int32
	_ = v449
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v475 int32
	_ = v475
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int64
	_ = v582
	var v589 int32
	_ = v589
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v648 int32
	_ = v648
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	v10 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(288)
	m.G0 = v29
	v32 = F_SearchSysCache1(m, int32(34), l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L6
	} else {
		goto L200
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L6
	} else {
		goto L197
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L6
	} else {
		goto L194
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L6
	} else {
		goto L191
	}
L5:
	;
	m.G0 = v29 + int32(288)
	return v648
L6:
	;
	return int32(0)
L7:
	;
	if v32 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l8 != 0 {
		v648 = int32(0)
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
	v54 = v52 + v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v58 = F_SysCacheGetAttrNotNull(m, int32(34), v32, int32(17))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L15
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = l0
	F_errmsg_internal(m, int32(37833), v29)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(472473), int32(1308), int32(210069))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	v62 = F_SysCacheGetAttrNotNull(m, int32(34), v32, int32(18))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v66 = F_SysCacheGetAttrNotNull(m, int32(34), v32, int32(19))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v69 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	if v69 == int32(0) {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+22)))
	v76 = v74 + v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+84))
	v78 = F_SearchSysCache1(m, int32(2), v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	if v78 == int32(0) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+22)))
	v84 = v82 + v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+68))
	v86 = F_GetIndexAmRoutine(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v88 = int32(0)
	v91 = F_heap_attisnull(m, v32, int32(20), v88)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L24
	}
L23:
	;
	v111 = F_get_rel_name(m, v55)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L6
	} else {
		goto L31
	}
L24:
	;
	if v91 != 0 {
		v109 = v10
		v110 = v88
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v95 = F_SysCacheGetAttrNotNull(m, int32(34), v32, int32(20))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v97 = F_text_to_cstring(m, v95)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v99 = F_stringToNode(m, v97)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	F_pfree(m, v97)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v103 = int32(0)
	if v99 == v103 {
		v109 = v10
		v110 = v103
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v109 = v99
	v110 = v106
	goto L23
L31:
	;
	if v111 == int32(0) {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v116 = F_palloc0(m, int32(80))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	v119 = F_palloc0(m, int32(136))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+24)) = int32(1)
	v123 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v119)+21)) = uint8(v123)
	*(*int32)(unsafe.Add(mBase, uint32(v119)+16)) = v55
	v126 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = int32(101)
	v131 = F_makeAlias(m, v111, v126)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v131
	v135 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v119)+124)) = uint16(v135)
	v137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v119)+20)) = uint8(v137)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+204)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v119
	v144 = F_list_make1_impl(m, int32(1), v29+int32(204))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v146 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v116)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v144
	F_set_rtable_names(m, v116, v146, v146)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	F_set_simple_column_names(m, v116)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v116
	v162 = F_list_make1_impl(m, int32(1), v29+int32(200))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	F_initStringInfo(m, v29+int32(216))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	if l3 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v226 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+8)))
	if v226 <= int32(0) {
		goto L66
	} else {
		goto L67
	}
L42:
	;
	if l2 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+12)))
	if v172 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v212 = F_quote_identifier(m, v84+int32(4))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L64
	}
L46:
	;
	v173 = int32(704536)
	goto L48
L47:
	;
	v173 = int32(717063)
	goto L48
L48:
	;
	v176 = F_quote_identifier(m, v76+int32(4))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	v178 = int32(717063)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+119)))
	if v181 != int32(73) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v184 = v178
	goto L52
L51:
	;
	v184 = int32(703513)
	goto L52
L52:
	;
	if l6 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v185 = v178
	goto L55
L54:
	;
	v185 = v184
	goto L55
L55:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l7) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v196 = F_quote_identifier(m, v84+int32(4))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L6
	} else {
		goto L62
	}
L57:
	;
	v189 = F_generate_relation_name(m, v55, int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v191 = F_generate_qualified_relation_name(m, v55)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L6
	} else {
		goto L61
	}
L60:
	;
	v193 = v189
	goto L56
L61:
	;
	v193 = v191
	goto L56
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+176)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v29)+172)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v29)+168)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v29)+164)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v29)+160)) = v173
	F_appendStringInfo(m, v29+int32(216), int32(645906), v29+int32(160))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	goto L41
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+192)) = v212
	F_appendStringInfo(m, v29+int32(216), int32(645943), v29+int32(192))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	goto L41
L66:
	;
	if l3 != 0 {
		goto L152
	} else {
		goto L153
	}
L67:
	;
	v229 = int32(24)
	v249 = int32(0)
	v250 = int32(717063)
	v255 = v110
	goto L68
L68:
	;
	v266 = v249 << (uint(int32(1)) % 32)
	v268 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54+int32(48)+v266))))
	if l4 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L66
L70:
	;
	v269 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+10)))
	if v269 <= v249 {
		goto L66
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if l1 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L72
L74:
	;
	v275 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+10)))
	if v275 == v249 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if v268 != 0 {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	F_appendStringInfoString(m, v29+int32(216), int32(646810))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L6
	} else {
		goto L80
	}
L78:
	;
	v283 = v250
	goto L79
L79:
	;
	F_appendStringInfoString(m, v29+int32(216), v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L6
	} else {
		goto L81
	}
L80:
	;
	v283 = int32(717063)
	goto L79
L81:
	;
	goto L76
L82:
	;
	if l3 != 0 {
		goto L115
	} else {
		goto L116
	}
L83:
	;
	v287 = F_get_attname(m, v55, v268, int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L6
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v255 == int32(0) {
		goto L1
	} else {
		goto L96
	}
L86:
	;
	if l1 != v249+int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v293 = l1
	goto L89
L88:
	;
	v293 = int32(0)
	goto L89
L89:
	;
	if v293 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v298 = F_quote_identifier(m, v287)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L6
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	F_get_atttypetypmodcoll(m, v55, v268, v29+int32(212), v29+int32(232), v29+int32(208))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L6
	} else {
		goto L95
	}
L93:
	;
	F_appendStringInfoString(m, v29+int32(216), v298)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v384 = v255
	goto L82
L96:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	F_initStringInfo(m, v29+int32(272))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	v319 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+264)) = uint8(v319)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v319
	v323 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+240)) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v29)+268)) = v319
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+267)) = uint8(v319)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+256)) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v29 + int32(272)
	v336 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+265)) = uint16(v336)
	F_get_rule_expr(m, v314, v29+int32(232), v319)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	v344 = v255 + int32(4)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v29)+272))
	if l1 != v249+int32(1) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if base.Ui32(v344) < base.Ui32(v312+v313<<(uint(int32(2))%32)) {
		goto L110
	} else {
		goto L111
	}
L100:
	;
	v354 = l1
	goto L102
L101:
	;
	v354 = int32(0)
	goto L102
L102:
	;
	if v354 != 0 {
		goto L99
	} else {
		goto L103
	}
L103:
	;
	if v314 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+144)) = v349
	F_appendStringInfo(m, v29+int32(216), int32(635961), v29+int32(144))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L6
	} else {
		goto L109
	}
L105:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	switch v357 - int32(15) {
	case 0:
		goto L107
	default:
		goto L104
	case 4, 23, 24, 25, 26, 33:
		goto L106
	}
L106:
	;
	F_appendStringInfoString(m, v29+int32(216), v349)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L6
	} else {
		goto L108
	}
L107:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v314)+16))
	switch v360 {
	case 0, 3:
		goto L106
	default:
		goto L104
	}
L108:
	;
	goto L99
L109:
	;
	goto L99
L110:
	;
	v374 = v344
	goto L112
L111:
	;
	v374 = int32(0)
	goto L112
L112:
	;
	v375 = F_exprType(m, v314)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L6
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+212)) = v375
	v378 = F_exprCollation(m, v314)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+208)) = v378
	v384 = v374
	goto L82
L115:
	;
	v482 = v249 + int32(1)
	v483 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+8)))
	if v482 < v483 {
		v249 = v482
		v250 = int32(705827)
		v255 = v384
		goto L68
	} else {
		goto L151
	}
L116:
	;
	v385 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+10)))
	if v385 <= v249 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v389 = v249 + int32(1)
	if v389 != l1 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v391 = l1
	goto L120
L119:
	;
	v391 = int32(0)
	goto L120
L120:
	;
	if v391 != 0 {
		goto L115
	} else {
		goto L121
	}
L121:
	;
	v393 = v249 << (uint(int32(2)) % 32)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v58+v229+v393)))
	v397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v266+(v66+v229)))))
	v399 = F_get_attoptions(m, l0, base.I32_extend16_s(v389))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	if v395 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v393+(v62+v229))))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v29)+212))
	if v399 != 0 {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v29)+208))
	if v395 == v403 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v405 = F_generate_collation_name(m, v395)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+128)) = v405
	F_appendStringInfo(m, v29+int32(216), int32(188571), v29+int32(128))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
	;
	goto L123
L128:
	;
	v419 = int32(0)
	goto L130
L129:
	;
	v419 = v418
	goto L130
L130:
	;
	F_get_opclass_name(m, v416, v419, v29+int32(216))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	if v399 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	F_appendStringInfoString(m, v29+int32(216), int32(646873))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L6
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+10)))
	if v438 != int32(1) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	F_get_reloptions(m, v29+int32(216), v399)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	F_appendStringInfoChar(m, v29+int32(216), int32(41))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	goto L134
L138:
	;
	if l2 == int32(0) {
		goto L115
	} else {
		goto L148
	}
L139:
	;
	if v397&int32(1) != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	F_appendStringInfoString(m, v29+int32(216), v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L6
	} else {
		goto L147
	}
L141:
	;
	F_appendStringInfoString(m, v29+int32(216), int32(521725))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L6
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	if v397&int32(2) == int32(0) {
		goto L138
	} else {
		goto L146
	}
L144:
	;
	if v397&int32(2) != 0 {
		goto L138
	} else {
		goto L145
	}
L145:
	;
	v458 = int32(495001)
	goto L140
L146:
	;
	v458 = int32(494864)
	goto L140
L147:
	;
	goto L138
L148:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l2+v393)))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v29)+212))
	v466 = F_generate_operator_name(m, v464, v465, v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L6
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+112)) = v466
	F_appendStringInfo(m, v29+int32(216), int32(188446), v29+int32(112))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L6
	} else {
		goto L150
	}
L150:
	;
	goto L115
L151:
	;
	goto L69
L152:
	;
	F_ReleaseCatCache(m, v32)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L6
	} else {
		goto L188
	}
L153:
	;
	F_appendStringInfoChar(m, v29+int32(216), int32(41))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L6
	} else {
		goto L154
	}
L154:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+13)))
	if v516 == int32(1) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	F_appendStringInfoString(m, v29+int32(216), int32(499029))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L6
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v524 = F_flatten_reloptions(m, l0)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L6
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	if v524 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+96)) = v524
	F_appendStringInfo(m, v29+int32(216), int32(635926), v29+int32(96))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L6
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	if l5 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	F_pfree(m, v524)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L6
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v562 = F_heap_attisnull(m, v32, int32(21), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L6
	} else {
		goto L176
	}
L166:
	;
	v538 = F_get_rel_tablespace(m, l0)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	if v538 == int32(0) {
		goto L165
	} else {
		goto L168
	}
L168:
	;
	if l2 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	F_appendStringInfoString(m, v29+int32(216), int32(487876))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L6
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v547 = F_get_tablespace_name(m, v538)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L6
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	v549 = F_quote_identifier(m, v547)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v549
	F_appendStringInfo(m, v29+int32(216), int32(188639), v29+int32(80))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L6
	} else {
		goto L175
	}
L175:
	;
	goto L165
L176:
	;
	if v562 != 0 {
		goto L152
	} else {
		goto L177
	}
L177:
	;
	v566 = F_SysCacheGetAttrNotNull(m, int32(34), v32, int32(21))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L6
	} else {
		goto L178
	}
L178:
	;
	v568 = F_text_to_cstring(m, v566)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L6
	} else {
		goto L179
	}
L179:
	;
	v570 = F_stringToNode(m, v568)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L6
	} else {
		goto L180
	}
L180:
	;
	F_pfree(m, v568)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L6
	} else {
		goto L181
	}
L181:
	;
	F_initStringInfo(m, v29+int32(272))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L6
	} else {
		goto L182
	}
L182:
	;
	v578 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+264)) = uint8(v578)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v578
	v582 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+240)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v29)+268)) = v578
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+267)) = uint8(v578)
	v589 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+265)) = uint16(v589)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+256)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v29 + int32(272)
	F_get_rule_expr(m, v570, v29+int32(232), v578)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L6
	} else {
		goto L183
	}
L183:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v29)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v602
	if l2 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v608 = int32(635937)
	goto L186
L185:
	;
	v608 = int32(188597)
	goto L186
L186:
	;
	F_appendStringInfo(m, v29+int32(216), v608, v29-int32(-64))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L6
	} else {
		goto L187
	}
L187:
	;
	goto L152
L188:
	;
	F_ReleaseCatCache(m, v69)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	F_ReleaseCatCache(m, v78)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L6
	} else {
		goto L190
	}
L190:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v29)+216))
	v648 = v621
	goto L5
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = l0
	F_errmsg_internal(m, int32(43981), v29+int32(16))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L6
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(472473), int32(1333), int32(210069))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L6
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v76)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v672
	F_errmsg_internal(m, int32(50892), v29+int32(32))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L6
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(472473), int32(1342), int32(210069))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L6
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v55
	F_errmsg_internal(m, int32(43981), v29+int32(48))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(472473), int32(13138), int32(361455))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L6
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	F_errmsg_internal(m, int32(71370), int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L6
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(472473), int32(1440), int32(210069))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L6
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_keywords(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == v2 {
		v17 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(4443856)
			v22 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v24
			v27 = F_get_call_result_type(m, l0, int32(0), v11)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(351051), int32(0))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(478044), int32(431), int32(162482))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v31
					v33 = F_TupleDescGetAttInMetadata(m, v31)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v33
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v22
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
						v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
						v45 = int64(*(*int32)(unsafe.Add(mBase, _consts[1090])))
						if base.Ui64(v43) < base.Ui64(v45) {
							v49 = *(*int32)(unsafe.Add(mBase, _consts[1091]))
							v51 = *(*int32)(unsafe.Add(mBase, _consts[1092]))
							v52 = base.I32_wrap_i64(v43)
							v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51+v52<<(uint(int32(1))%32)))))
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v49 + v56
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_consts[1093]))))
							if base.Ui32(v61) <= base.Ui32(int32(3)) {
								v65 = v61 << (uint(int32(2)) % 32)
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[1094])))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[1095])))
								v72 = v71
								v73 = v68
							} else {
								v72 = int32(0)
								v73 = v2
							}
							*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v73
							*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v72
							v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_consts[1096]))))
							if v80 != 0 {
								v81 = int32(294232)
							} else {
								v81 = int32(501990)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v81
							if v80 != 0 {
								v85 = int32(328733)
							} else {
								v85 = int32(345308)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v85
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
							v88 = F_BuildTupleFromCStrings(m, v87, v11)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								v90 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
								*(*int64)(unsafe.Add(mBase, uint32(v42))) = v90 + int64(1)
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = int32(1)
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
								v98 = F_HeapTupleHeaderGetDatum(m, v97)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v112 = v98
									m.G0 = v11 + int32(32)
									return v112
								}
							}
						} else {
							F_end_MultiFuncCall(m, l0)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v102)+20)) = int32(2)
								v105 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
								v112 = int32(0)
								m.G0 = v11 + int32(32)
								return v112
							}
						}
					}
				}
			}
		}
	} else {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
		v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
		v45 = int64(*(*int32)(unsafe.Add(mBase, _consts[1090])))
		if base.Ui64(v43) < base.Ui64(v45) {
			v49 = *(*int32)(unsafe.Add(mBase, _consts[1091]))
			v51 = *(*int32)(unsafe.Add(mBase, _consts[1092]))
			v52 = base.I32_wrap_i64(v43)
			v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51+v52<<(uint(int32(1))%32)))))
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v49 + v56
			v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_consts[1093]))))
			if base.Ui32(v61) <= base.Ui32(int32(3)) {
				v65 = v61 << (uint(int32(2)) % 32)
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[1094])))
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[1095])))
				v72 = v71
				v73 = v68
			} else {
				v72 = int32(0)
				v73 = v2
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v73
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v72
			v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_consts[1096]))))
			if v80 != 0 {
				v81 = int32(294232)
			} else {
				v81 = int32(501990)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v81
			if v80 != 0 {
				v85 = int32(328733)
			} else {
				v85 = int32(345308)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v85
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
			v88 = F_BuildTupleFromCStrings(m, v87, v11)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				v90 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
				*(*int64)(unsafe.Add(mBase, uint32(v42))) = v90 + int64(1)
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = int32(1)
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
				v98 = F_HeapTupleHeaderGetDatum(m, v97)
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return int32(0)
				} else {
					v112 = v98
					m.G0 = v11 + int32(32)
					return v112
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v101 = m.ExcPending
			if v101 != 0 {
				return int32(0)
			} else {
				v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v102)+20)) = int32(2)
				v105 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
				v112 = int32(0)
				m.G0 = v11 + int32(32)
				return v112
			}
		}
	}
}
func F_pg_get_next_timezone_abbrev(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v6 < v3 {
		v26 = v3
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+268))
		if v9 <= v6 {
			v26 = v3
		} else {
			v12 = l1 + int32(22376)
			v15 = v6
			for {
				v20 = v15 + int32(1)
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v12))))
				if v21 != 0 {
					v15 = v20
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v20
			v26 = v6 + v12
		}
	}
	return v26
}
func F_pg_get_publication_tables(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v450 int32
	_ = v450
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v514 int32
	_ = v514
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int64
	_ = v624
	var v625 int64
	_ = v625
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int64
	_ = v636
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v769 int32
	_ = v769
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int64
	_ = v791
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v826 int32
	_ = v826
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v22 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v619)+16))
	goto L125
L4:
	;
	return int32(0)
L5:
	;
	v29 = int32(4443856)
	v30 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v35 = F_pg_detoast_datum(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_deconstruct_array_builtin(m, v35, int32(25), v19, int32(0), v19+int32(28))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if v43 <= int32(0) {
		v561 = v2
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v567 = F_CreateTemplateTupleDesc(m, int32(4))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L4
	} else {
		goto L117
	}
L9:
	;
	v55 = v2
	v56 = v2
	v57 = v2
	goto L10
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v55<<(uint(int32(2))%32))))
	v68 = F_text_to_cstring(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L14
	}
L11:
	;
	v397 = int32(0)
	if v392&base.B2i32(v386 != v397) == v397 {
		v561 = v386
		goto L8
	} else {
		goto L92
	}
L12:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+9)))
	v392 = v391 | v56
	v394 = v55 + int32(1)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if v394 < v395 {
		v55 = v394
		v56 = v392
		v57 = v386
		goto L10
	} else {
		goto L91
	}
L13:
	;
	if v335 == int32(0) {
		v386 = v57
		goto L12
	} else {
		goto L84
	}
L14:
	;
	v71 = F_get_publication_oid(m, v68, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	if v71 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v73 = F_GetPublication(m, v71)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	v75 = int32(0)
	goto L18
L18:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+8)))
	if v76 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v75 = v73
	goto L18
L20:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+9)))
	v80 = int32(0)
	v81 = m.G0
	v83 = v81 - int32(48)
	m.G0 = v83
	v87 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+9)))
	v255 = F_GetPublicationRelations(m, v251, v252^int32(1))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L69
	}
L23:
	;
	F_ScanKeyInit(m, v83, int32(18), int32(3), int32(61), int32(114))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v96 = F_table_beginscan_catalog(m, v87, int32(1), v83)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v98 = F_heap_getnext(m, v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	if v98 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v101 = v98
	v106 = v80
	goto L30
L28:
	;
	v145 = v80
	goto L29
L29:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+188))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	m.T0[v157].(func(*base.Module, int32))(m, v96)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L45
	}
L30:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+22)))
	v118 = v116 + v117
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+119)))
	switch v119 - int32(112) {
	case 0, 2:
		goto L33
	default:
		v136 = v106
		goto L32
	}
L31:
	;
	v145 = v136
	goto L29
L32:
	;
	v137 = F_heap_getnext(m, v96)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L43
	}
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	goto L34
L34:
	;
	if base.Ui32(v122) < base.Ui32(int32(12000)) {
		v136 = v106
		goto L32
	} else {
		goto L35
	}
L35:
	;
	if base.Ui32(v122) < base.Ui32(int32(16384)) {
		v136 = v106
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+118)))
	if v127 != int32(112) {
		v136 = v106
		goto L32
	} else {
		goto L37
	}
L37:
	;
	if v79 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+131)))
	if v130&int32(1) != 0 {
		v136 = v106
		goto L32
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v133 = F_lappend_oid(m, v106, v122)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	v136 = v133
	goto L32
L43:
	;
	if v137 != 0 {
		v101 = v137
		v106 = v136
		goto L30
	} else {
		goto L44
	}
L44:
	;
	goto L31
L45:
	;
	if v79 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	F_ScanKeyInit(m, v83, int32(18), int32(3), int32(61), int32(112))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	v235 = v145
	goto L48
L48:
	;
	F_sequence_close(m, v87, int32(1))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L68
	}
L49:
	;
	v167 = F_table_beginscan_catalog(m, v87, int32(1), v83)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v169 = F_heap_getnext(m, v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	if v169 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v172 = v169
	v177 = v145
	goto L55
L53:
	;
	v214 = v145
	goto L54
L54:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+188))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	m.T0[v226].(func(*base.Module, int32))(m, v167)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L67
	}
L55:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v172)+16))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+22)))
	v189 = v187 + v188
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+119)))
	switch v190 - int32(112) {
	case 0, 2:
		goto L58
	default:
		v205 = v177
		goto L57
	}
L56:
	;
	v214 = v205
	goto L54
L57:
	;
	v206 = F_heap_getnext(m, v167)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L65
	}
L58:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	goto L59
L59:
	;
	if base.Ui32(v193) < base.Ui32(int32(12000)) {
		v205 = v177
		goto L57
	} else {
		goto L60
	}
L60:
	;
	if base.Ui32(v193) < base.Ui32(int32(16384)) {
		v205 = v177
		goto L57
	} else {
		goto L61
	}
L61:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+118)))
	if v198 != int32(112) {
		v205 = v177
		goto L57
	} else {
		goto L62
	}
L62:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+131)))
	if v201 != 0 {
		v205 = v177
		goto L57
	} else {
		goto L63
	}
L63:
	;
	v202 = F_lappend_oid(m, v177, v193)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v205 = v202
	goto L57
L65:
	;
	if v206 != 0 {
		v172 = v206
		v177 = v205
		goto L55
	} else {
		goto L66
	}
L66:
	;
	goto L56
L67:
	;
	v235 = v214
	goto L48
L68:
	;
	m.G0 = v83 + int32(48)
	v335 = v235
	goto L13
L69:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+9)))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v259 = F_GetPublicationSchemas(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L71
	}
L70:
	;
	v317 = F_list_concat_unique_oid(m, v255, v303)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L83
	}
L71:
	;
	if v259 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v303 = int32(0)
	goto L70
L73:
	;
	goto L74
L74:
	;
	v264 = int32(0)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if v265 <= v264 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v303 = int32(0)
	goto L70
L76:
	;
	goto L77
L77:
	;
	v273 = v264
	v274 = int32(0)
	goto L78
L78:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288+v273<<(uint(int32(2))%32))))
	v293 = F_GetSchemaPublicationRelations(m, v292, v257^int32(1))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
	} else {
		goto L80
	}
L79:
	;
	v303 = v295
	goto L70
L80:
	;
	v295 = F_list_concat(m, v274, v293)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v298 = v273 + int32(1)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if v298 < v299 {
		v273 = v298
		v274 = v295
		goto L78
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	v335 = v317
	goto L13
L84:
	;
	v338 = int32(0)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	if v339 <= v338 {
		v386 = v57
		goto L12
	} else {
		goto L85
	}
L85:
	;
	v343 = v338
	v353 = v57
	goto L86
L86:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v335)+12))
	v360 = F_palloc(m, int32(8))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L4
	} else {
		goto L88
	}
L87:
	;
	v386 = v369
	goto L12
L88:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v358+v343<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v360))) = v365
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	*(*int32)(unsafe.Add(mBase, uint32(v360)+4)) = v367
	v369 = F_lappend(m, v353, v360)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	v372 = v343 + int32(1)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	if v372 < v373 {
		v343 = v372
		v353 = v369
		goto L86
	} else {
		goto L90
	}
L90:
	;
	goto L87
L91:
	;
	goto L11
L92:
	;
	v406 = v386
	v408 = v397
	goto L93
L93:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	if v419 <= v408 {
		v561 = v386
		goto L8
	} else {
		goto L95
	}
L94:
	;
	v561 = v386
	goto L8
L95:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v406)+12))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v421+v408<<(uint(int32(2))%32))))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	v427 = F_get_rel_relispartition(m, v426)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L4
	} else {
		goto L98
	}
L96:
	;
	if v535 != 0 {
		v406 = v535
		v408 = v533 + int32(1)
		goto L93
	} else {
		goto L116
	}
L97:
	;
	v533 = v408
	v535 = v406
	goto L96
L98:
	;
	if v427 == int32(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	v432 = F_get_partition_ancestors(m, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	if v432 == int32(0) {
		goto L97
	} else {
		goto L101
	}
L101:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	if v436 <= int32(0) {
		goto L97
	} else {
		goto L102
	}
L102:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	v450 = int32(0)
	goto L103
L103:
	;
	if v406 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	goto L97
L105:
	;
	v514 = v450 + int32(1)
	if v436 != v514 {
		v450 = v514
		goto L103
	} else {
		goto L115
	}
L106:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	if v459 <= int32(0) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v439+v450<<(uint(int32(2))%32))))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v406)+12))
	v469 = int32(0)
	goto L108
L108:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v466+v469<<(uint(int32(2))%32))))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	if v465 != v488 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v495 = F_list_delete_nth_cell(m, v406, v408)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L114
	}
L110:
	;
	v491 = v469 + int32(1)
	if v491 != v459 {
		v469 = v491
		goto L108
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	goto L109
L113:
	;
	goto L105
L114:
	;
	v533 = v408 - int32(1)
	v535 = v495
	goto L96
L115:
	;
	goto L104
L116:
	;
	goto L94
L117:
	;
	F_TupleDescInitEntry(m, v567, int32(1), int32(416820), int32(26), int32(-1), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	F_TupleDescInitEntry(m, v567, int32(2), int32(415938), int32(26), int32(-1), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	F_TupleDescInitEntry(m, v567, int32(3), int32(122366), int32(22), int32(-1), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	F_TupleDescInitEntry(m, v567, int32(4), int32(295740), int32(194), int32(-1), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	v597 = F_BlessTupleDesc(m, v567)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v597
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v30
	goto L3
L123:
	;
	m.G0 = v19 + int32(32)
	return v826
L124:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L4
	} else {
		goto L161
	}
L125:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v620)+16))
	if v621 == int32(0) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v624 = *(*int64)(unsafe.Add(mBase, uint32(v620)))
	v625 = int64(*(*int32)(unsafe.Add(mBase, uint32(v621)+4)))
	if base.Ui64(v625) <= base.Ui64(v624) {
		goto L124
	} else {
		goto L127
	}
L127:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v621)+12))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v627+base.I32_wrap_i64(v624)<<(uint(int32(2))%32))))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v632)))
	v634 = F_get_rel_namespace(m, v633)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	v636 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v636
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v636
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = int32(0)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v632)+4))
	v643 = F_GetPublication(m, v642)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v643)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v633
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v645
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v643)+8)))
	if v648 != 0 {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v620)+28))
	v789 = F_heap_form_tuple(m, v786, v19, v19+int32(28))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L4
	} else {
		goto L159
	}
L131:
	;
	v686 = F_table_open(m, v633, int32(1))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L4
	} else {
		goto L142
	}
L132:
	;
	v667 = F_SysCacheGetAttr(m, int32(53), v656, int32(5), v19+int32(28)|int32(2))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L4
	} else {
		goto L139
	}
L133:
	;
	v659 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+30)) = uint16(v659)
	goto L131
L134:
	;
	v650 = int32(0)
	v652 = F_SearchSysCacheExists(m, int32(50), v634, v645, v650, v650)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	if v652 != 0 {
		goto L133
	} else {
		goto L136
	}
L136:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v643)))
	v656 = F_SearchSysCacheCopy(m, int32(53), v633, v655)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	if v656 != 0 {
		goto L132
	} else {
		goto L138
	}
L138:
	;
	goto L133
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v667
	v676 = F_SysCacheGetAttr(m, int32(53), v656, int32(4), v19+int32(28)|int32(3))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v676
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+30)))
	if v679&int32(1) == int32(0) {
		goto L130
	} else {
		goto L141
	}
L141:
	;
	goto L131
L142:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v686)+52))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v688)))
	v692 = F_palloc(m, v689<<(uint(int32(1))%32))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v688)))
	if v694 <= int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	F_sequence_close(m, v686, int32(1))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L4
	} else {
		goto L158
	}
L145:
	;
	v699 = int32(0)
	v702 = v699
	v703 = v694
	v704 = v699
	goto L146
L146:
	;
	v722 = v688 + int32(20) + v703<<(uint(int32(4))%32) + v702*int32(100)
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+91)))
	if v723 != 0 {
		v738 = v703
		v739 = v704
		goto L148
	} else {
		goto L149
	}
L147:
	;
	if v739 <= int32(0) {
		goto L144
	} else {
		goto L156
	}
L148:
	;
	v742 = v702 + int32(1)
	if v742 < v738 {
		v702 = v742
		v703 = v738
		v704 = v739
		goto L146
	} else {
		goto L155
	}
L149:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+90)))
	if v724 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if v724 != int32(115) {
		v738 = v703
		v739 = v704
		goto L148
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v730 = int32(1)
	v733 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v722)+74)))
	*(*uint16)(unsafe.Add(mBase, uint32(v692+v704<<(uint(v730)%32)))) = uint16(v733)
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v688)))
	v738 = v737
	v739 = v704 + v730
	goto L148
L153:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v643)+12))
	if v727 != int32(115) {
		v738 = v703
		v739 = v704
		goto L148
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	goto L147
L156:
	;
	v746 = F_buildint2vector(m, v692, v739)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	v748 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+30)) = uint8(v748)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v746
	goto L144
L158:
	;
	goto L130
L159:
	;
	v791 = *(*int64)(unsafe.Add(mBase, uint32(v620)))
	*(*int64)(unsafe.Add(mBase, uint32(v620))) = v791 + int64(1)
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v795)+20)) = int32(1)
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v789)+16))
	v799 = F_HeapTupleHeaderGetDatum(m, v798)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	v826 = v799
	goto L123
L161:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v804)+20)) = int32(2)
	v807 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v807)
	v826 = int32(0)
	goto L123
}
func F_pg_get_sequence_data(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+38)) = uint16(v2)
	v15 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		F_TupleDescInitEntry(m, v15, int32(1), int32(329329), int32(20), int32(-1), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			F_TupleDescInitEntry(m, v15, int32(2), int32(434288), int32(16), int32(-1), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = F_BlessTupleDesc(m, v15)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v36 = F_try_relation_open(m, v9, int32(1))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						if v36 != 0 {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+119)))
							if v39 != int32(83) {
								v84 = int32(257)
								*(*uint16)(unsafe.Add(mBase, uint32(v7)+38)) = uint16(v84)
								F_relation_close(m, v36, int32(1))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int32(0)
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
										v98 = F_HeapTupleHeaderGetDatum(m, v97)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											m.G0 = v7 + int32(48)
											return v98
										}
									}
								}
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, _consts[31]))
								v45 = F_pg_class_aclcheck(m, v9, v43, int64(2))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									if v45 != 0 {
										v84 = int32(257)
										*(*uint16)(unsafe.Add(mBase, uint32(v7)+38)) = uint16(v84)
										F_relation_close(m, v36, int32(1))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
												v98 = F_HeapTupleHeaderGetDatum(m, v97)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int32(0)
												} else {
													m.G0 = v7 + int32(48)
													return v98
												}
											}
										}
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
										v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+118)))
										switch v48 - int32(112) {
										case 0:
											v71 = F_read_seq_tuple(m, v36, v7+int32(32), v7+int32(12))
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int32(0)
											} else {
												v73 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
												v74 = F_Int64GetDatum(m, v73)
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v74
													v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+16)))
													*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v77
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
													F_UnlockReleaseBuffer(m, v79)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														F_relation_close(m, v36, int32(1))
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int32(0)
														} else {
															v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return int32(0)
															} else {
																v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
																v98 = F_HeapTupleHeaderGetDatum(m, v97)
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v7 + int32(48)
																	return v98
																}
															}
														}
													}
												}
											}
										default:
											v56 = int32(*(*uint8)(unsafe.Add(mBase, _consts[29])))
											if v56 == int32(1) {
												v61 = *(*int32)(unsafe.Add(mBase, _consts[30]))
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+316))
												v64 = base.B2i32(v62 != int32(2))
												*(*uint8)(unsafe.Add(mBase, _consts[29])) = uint8(v64)
												v66 = v64
											} else {
												v66 = int32(0)
											}
											if v66 != 0 {
												v84 = int32(257)
												*(*uint16)(unsafe.Add(mBase, uint32(v7)+38)) = uint16(v84)
												F_relation_close(m, v36, int32(1))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return int32(0)
													} else {
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
														v98 = F_HeapTupleHeaderGetDatum(m, v97)
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int32(0)
														} else {
															m.G0 = v7 + int32(48)
															return v98
														}
													}
												}
											} else {
												v71 = F_read_seq_tuple(m, v36, v7+int32(32), v7+int32(12))
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return int32(0)
												} else {
													v73 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
													v74 = F_Int64GetDatum(m, v73)
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v74
														v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+16)))
														*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v77
														v79 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
														F_UnlockReleaseBuffer(m, v79)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v36, int32(1))
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return int32(0)
															} else {
																v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return int32(0)
																} else {
																	v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
																	v98 = F_HeapTupleHeaderGetDatum(m, v97)
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v7 + int32(48)
																		return v98
																	}
																}
															}
														}
													}
												}
											}
										case 4:
											v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)))
											if v51 != int32(1) {
												v84 = int32(257)
												*(*uint16)(unsafe.Add(mBase, uint32(v7)+38)) = uint16(v84)
												F_relation_close(m, v36, int32(1))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return int32(0)
													} else {
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
														v98 = F_HeapTupleHeaderGetDatum(m, v97)
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int32(0)
														} else {
															m.G0 = v7 + int32(48)
															return v98
														}
													}
												}
											} else {
												v56 = int32(*(*uint8)(unsafe.Add(mBase, _consts[29])))
												if v56 == int32(1) {
													v61 = *(*int32)(unsafe.Add(mBase, _consts[30]))
													v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+316))
													v64 = base.B2i32(v62 != int32(2))
													*(*uint8)(unsafe.Add(mBase, _consts[29])) = uint8(v64)
													v66 = v64
												} else {
													v66 = int32(0)
												}
												if v66 != 0 {
													v84 = int32(257)
													*(*uint16)(unsafe.Add(mBase, uint32(v7)+38)) = uint16(v84)
													F_relation_close(m, v36, int32(1))
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return int32(0)
														} else {
															v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
															v98 = F_HeapTupleHeaderGetDatum(m, v97)
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return int32(0)
															} else {
																m.G0 = v7 + int32(48)
																return v98
															}
														}
													}
												} else {
													v71 = F_read_seq_tuple(m, v36, v7+int32(32), v7+int32(12))
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return int32(0)
													} else {
														v73 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
														v74 = F_Int64GetDatum(m, v73)
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v74
															v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+16)))
															*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v77
															v79 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
															F_UnlockReleaseBuffer(m, v79)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v36, int32(1))
																mBase = m.M
																v89 = m.ExcPending
																if v89 != 0 {
																	return int32(0)
																} else {
																	v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
																	mBase = m.M
																	v96 = m.ExcPending
																	if v96 != 0 {
																		return int32(0)
																	} else {
																		v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
																		v98 = F_HeapTupleHeaderGetDatum(m, v97)
																		mBase = m.M
																		v99 = m.ExcPending
																		if v99 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v7 + int32(48)
																			return v98
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v82 = int32(257)
							*(*uint16)(unsafe.Add(mBase, uint32(v7)+38)) = uint16(v82)
							v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
								v98 = F_HeapTupleHeaderGetDatum(m, v97)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(48)
									return v98
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_get_shmem_allocations(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v20 = F_LWLockAcquire(m, v16+int32(128), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[787]))
	F_hash_seq_init(m, v5+int32(-20), v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	v32 = F_hash_seq_search(m, v5+int32(-20))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v34 = v32
	v37 = v2
	goto L9
L7:
	;
	v74 = v2
	goto L8
L8:
	;
	v76 = F_cstring_to_text(m, int32(523571))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L18
	}
L9:
	;
	v38 = F_cstring_to_text(m, v34)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v74 = v66
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v38
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v43 = *(*int32)(unsafe.Add(mBase, _consts[788]))
	v46 = F_Int64GetDatum(m, base.I64_extend_i32_s(v41-v43))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v46
	v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v34)+52)))
	v50 = F_Int64GetDatum(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v50
	v53 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v34)+56)))
	v54 = F_Int64GetDatum(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_tuplestore_putvalues(m, v58, v59, v5+int32(-48), v5+int32(-52))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v66 = v57 + v37
	v69 = F_hash_seq_search(m, v5+int32(-20))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v69 != 0 {
		v34 = v69
		v37 = v66
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L10
L18:
	;
	v78 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v78)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v76
	v82 = *(*int32)(unsafe.Add(mBase, _consts[788]))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v86 = F_Int64GetDatum(m, base.I64_extend_i32_u(v83-v74))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v86
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_tuplestore_putvalues(m, v90, v91, v5+int32(-48), v5+int32(-52))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v98 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v98)
	v101 = *(*int32)(unsafe.Add(mBase, _consts[788]))
	v102 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v101)+12)))
	v103 = F_Int64GetDatum(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v105 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v103
	v109 = *(*int32)(unsafe.Add(mBase, _consts[788]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v114 = F_Int64GetDatum(m, base.I64_extend_i32_u(v110-v111))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v114
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_tuplestore_putvalues(m, v118, v119, v5+int32(-48), v5+int32(-52))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v127+int32(128))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	m.G0 = v7 - int32(-64)
	return int32(0)
}
func F_pg_get_shmem_allocations_numa(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v2 = m.G0
	m.G0 = v2 + int32(-64)
	F_errstart_cold(m, int32(21), int32(0))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_errmsg_internal(m, int32(274580), int32(0))
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(475535), int32(601), int32(483819))
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_pg_getaddrinfo_all(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	v5 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v9 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v316
L2:
	;
	v12 = int32(-4)
	if l1&int32(3) == int32(0) {
		v36 = l1
		goto L7
	} else {
		goto L8
	}
L3:
	;
	goto L4
L4:
	;
	if l0 != 0 {
		goto L90
	} else {
		goto L91
	}
L5:
	;
	if base.Ui32(int32(107)) < base.Ui32(v69) {
		v316 = v12
		goto L1
	} else {
		goto L22
	}
L6:
	;
	v69 = v61 - l1
	goto L5
L7:
	;
	v40 = v36
	goto L16
L8:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v20 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v69 = int32(0)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v25 = l1
	goto L12
L12:
	;
	v29 = v25 + int32(1)
	if v29&int32(3) == int32(0) {
		v36 = v29
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v61 = v29
	goto L6
L14:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v34 != 0 {
		v25 = v29
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v49 = int32(-2139062144)
	if (int32(16843008)-v46|v46)&v49 == v49 {
		v40 = v40 + int32(4)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v55 = v40
	goto L19
L18:
	;
	goto L17
L19:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v59 != 0 {
		v55 = v55 + int32(1)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v61 = v55
	goto L6
L21:
	;
	goto L20
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v72 != int32(1) {
		v316 = v12
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	goto L27
L24:
	;
	if v100 == int32(0) {
		v316 = int32(-10)
		goto L1
	} else {
		goto L34
	}
L25:
	;
	goto L24
L26:
	;
	v100 = F_emscripten_builtin_malloc(m, v88)
	mBase = m.M
	if v100 == int32(0) {
		goto L25
	} else {
		goto L32
	}
L27:
	;
	v88 = base.I32_wrap_i64(base.I64_extend_i32_u(int32(1)) * base.I64_extend_i32_u(int32(32)))
	goto L26
L32:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100-int32(4)))))
	if v105&int32(3) == int32(0) {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v111 = F___memset(m, v100, int32(0), v88)
	mBase = m.M
	goto L25
L34:
	;
	goto L38
L35:
	;
	if v136 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L36:
	;
	goto L35
L37:
	;
	v136 = F_emscripten_builtin_malloc(m, v124)
	mBase = m.M
	if v136 == int32(0) {
		goto L36
	} else {
		goto L43
	}
L38:
	;
	v124 = base.I32_wrap_i64(base.I64_extend_i32_u(int32(1)) * base.I64_extend_i32_u(int32(110)))
	goto L37
L43:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136-int32(4)))))
	if v141&int32(3) == int32(0) {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v147 = F___memset(m, v136, int32(0), v124)
	mBase = m.M
	goto L36
L45:
	;
	F_emscripten_builtin_free(m, v100)
	mBase = m.M
	return int32(-10)
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+12)) = v76
	v154 = int32(1)
	if base.Ui32(v75) <= base.Ui32(v154) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v157 = v154
	goto L50
L49:
	;
	v157 = v75
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v157
	v159 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v100)+4)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v100
	*(*uint16)(unsafe.Add(mBase, uint32(v136))) = uint16(v159)
	*(*int32)(unsafe.Add(mBase, uint32(v100)+16)) = int32(110)
	*(*int32)(unsafe.Add(mBase, uint32(v100)+20)) = v136
	v168 = v136 + int32(2)
	if (l1^v168)&int32(3) != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v244 != int32(64) {
		v316 = int32(0)
		goto L1
	} else {
		goto L72
	}
L52:
	;
	goto L51
L53:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v223))) = uint8(v222)
	if v222&int32(255) == int32(0) {
		goto L52
	} else {
		goto L68
	}
L54:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v221 = l1
	v222 = v174
	v223 = v168
	goto L53
L55:
	;
	goto L56
L56:
	;
	if l1&int32(3) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v178 = l1
	v180 = v168
	goto L60
L58:
	;
	v192 = l1
	v194 = v168
	goto L59
L59:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v199 = int32(-2139062144)
	if (int32(16843008)-v196|v196)&v199 != v199 {
		v221 = v192
		v222 = v196
		v223 = v194
		goto L53
	} else {
		goto L64
	}
L60:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v181)
	if v181 == int32(0) {
		goto L52
	} else {
		goto L62
	}
L61:
	;
	v192 = v188
	v194 = v186
	goto L59
L62:
	;
	v185 = int32(1)
	v186 = v180 + v185
	v188 = v178 + v185
	if v188&int32(3) != 0 {
		v178 = v188
		v180 = v186
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v204 = v192
	v205 = v196
	v206 = v194
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v205
	v208 = int32(4)
	v209 = v206 + v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v212 = v204 + v208
	v216 = int32(-2139062144)
	if (v210|(int32(16843008)-v210))&v216 == v216 {
		v204 = v212
		v205 = v210
		v206 = v209
		goto L65
	} else {
		goto L67
	}
L66:
	;
	v221 = v212
	v222 = v210
	v223 = v209
	goto L53
L67:
	;
	goto L66
L68:
	;
	v230 = v221
	v232 = v223
	goto L69
L69:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v232)+1)) = uint8(v233)
	v235 = int32(1)
	if v233 != 0 {
		v230 = v230 + v235
		v232 = v232 + v235
		goto L69
	} else {
		goto L71
	}
L70:
	;
	goto L52
L71:
	;
	goto L70
L72:
	;
	v247 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v247)
	if l1&int32(3) == v247 {
		v272 = l1
		goto L75
	} else {
		goto L76
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+16)) = v305 + int32(2)
	return int32(0)
L74:
	;
	v305 = v297 - l1
	goto L73
L75:
	;
	v276 = v272
	goto L84
L76:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v256 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v305 = int32(0)
	goto L73
L78:
	;
	goto L79
L79:
	;
	v261 = l1
	goto L80
L80:
	;
	v265 = v261 + int32(1)
	if v265&int32(3) == int32(0) {
		v272 = v265
		goto L75
	} else {
		goto L82
	}
L81:
	;
	v297 = v265
	goto L74
L82:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v270 != 0 {
		v261 = v265
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	v285 = int32(-2139062144)
	if (int32(16843008)-v282|v282)&v285 == v285 {
		v276 = v276 + int32(4)
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v291 = v276
	goto L87
L86:
	;
	goto L85
L87:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	if v295 != 0 {
		v291 = v291 + int32(1)
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v297 = v291
	goto L74
L89:
	;
	goto L88
L90:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v312 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v314 = v5
	goto L92
L92:
	;
	v315 = m.Env.Getaddrinfo(m, v314, l1, l2, l3)
	mBase = m.M
	v316 = v315
	goto L1
L93:
	;
	v313 = l0
	goto L95
L94:
	;
	v313 = int32(0)
	goto L95
L95:
	;
	v314 = v313
	goto L92
}
func F_pg_has_role_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[31]))
		v12 = F_convert_any_priv_string(m, v5, int32(1616624))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_pg_role_aclcheck(m, v3, v10, v12)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14 ^ int32(1)
			}
		}
	}
}
func F_pg_has_role_id_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = F_convert_any_priv_string(m, v5, int32(1616624))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = F_pg_role_aclcheck(m, v2, v3, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12 ^ int32(1)
			}
		}
	}
}
func F_pg_has_role_name_id(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v21 = F_GetSysCacheOid(m, int32(10), v11, v18, v18, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v21 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
						F_errmsg(m, int32(69168), v8)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(475848), int32(5562), int32(415649))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v42 = F_convert_any_priv_string(m, v13, int32(1616624))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v44 = F_pg_role_aclcheck(m, v10, v21, v42)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v44 ^ int32(1)
					}
				}
			}
		}
	}
}
func F_pg_hmac_update(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	if l0 == int32(0) {
		return int32(-1)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = F_pg_cryptohash_update(m, v8, l1, l2)
		mBase = m.M
		if int32(0) <= v9 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v16 == int32(0) {
				v31 = int32(12890)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				if v23 == int32(1) {
					v26 = int32(290988)
				} else {
					v26 = int32(121505)
				}
				if v23 == int32(2) {
					v29 = int32(12890)
				} else {
					v29 = v26
				}
				v31 = v29
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v31
			return int32(-1)
		}
	}
}
func F_pg_jit_available(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_provider_init(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_pg_johab_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	if l1 <= int32(0) {
		v46 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v46 - l0
L2:
	;
	v9 = l1
	v10 = l0
	goto L3
L3:
	;
	v13 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
	if int32(0) <= v13 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v46 = v40
	goto L1
L5:
	;
	v40 = v10 + v39
	v41 = v9 - v39
	if int32(0) < v41 {
		v9 = v41
		v10 = v40
		goto L3
	} else {
		goto L17
	}
L6:
	;
	if v13 != 0 {
		v39 = int32(1)
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if v13 == int32(-113) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v46 = v10
	goto L1
L10:
	;
	v21 = int32(3)
	goto L12
L11:
	;
	v21 = int32(2)
	goto L12
L12:
	;
	if base.Ui32(v9) < base.Ui32(v21) {
		v46 = v10
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if base.Ui32(int32(93)) < base.Ui32((v23+int32(95))&int32(255)) {
		v46 = v10
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v13 != int32(-113) {
		v39 = v21
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+2)))
	if base.Ui32(int32(93)) < base.Ui32((v32+int32(95))&int32(255)) {
		v46 = v10
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v39 = v21
	goto L5
L17:
	;
	goto L4
}
func F_pg_mb2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v5 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6*int32(28))+uint32(_consts[1203])))
	v12 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
func F_pg_mblen_unbounded(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4*int32(28))+uint32(_consts[1205])))
	v10 = m.T0[v9].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_pg_md5_hash(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	v15 = F_pg_cryptohash_create(m, v5)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(12890)
			v104 = v5
			m.G0 = v10 + int32(16)
			return v104
		} else {
			v38 = F_pg_cryptohash_init(m, v15)
			mBase = m.M
			if v38 < int32(0) {
				if v15 == int32(0) {
					v96 = int32(12890)
				} else {
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
					if v88 == int32(1) {
						v91 = int32(290988)
					} else {
						v91 = int32(121505)
					}
					if v88 == int32(2) {
						v94 = int32(12890)
					} else {
						v94 = v91
					}
					v96 = v94
				}
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v96
				F_pg_cryptohash_free(m, v15)
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return int32(0)
				} else {
					v104 = v5
					m.G0 = v10 + int32(16)
					return v104
				}
			} else {
				v41 = F_pg_cryptohash_update(m, v15, l0, l1)
				mBase = m.M
				if v41 < int32(0) {
					if v15 == int32(0) {
						v96 = int32(12890)
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
						if v88 == int32(1) {
							v91 = int32(290988)
						} else {
							v91 = int32(121505)
						}
						if v88 == int32(2) {
							v94 = int32(12890)
						} else {
							v94 = v91
						}
						v96 = v94
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v96
					F_pg_cryptohash_free(m, v15)
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						v104 = v5
						m.G0 = v10 + int32(16)
						return v104
					}
				} else {
					v45 = F_pg_cryptohash_final(m, v15, v10, int32(16))
					mBase = m.M
					if v45 < int32(0) {
						if v15 == int32(0) {
							v96 = int32(12890)
						} else {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
							if v88 == int32(1) {
								v91 = int32(290988)
							} else {
								v91 = int32(121505)
							}
							if v88 == int32(2) {
								v94 = int32(12890)
							} else {
								v94 = v91
							}
							v96 = v94
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v96
						F_pg_cryptohash_free(m, v15)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v104 = v5
							m.G0 = v10 + int32(16)
							return v104
						}
					} else {
						v52 = int32(0)
						v53 = v5
						for {
							v56 = l2 + v53
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v10))))
							v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58&int32(15))+uint32(_consts[1142]))))
							*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v63)
							v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v58)>>(uint(int32(4))%32)))+uint32(_consts[1142]))))
							*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v69)
							v74 = v52 + int32(1)
							if v74 != int32(16) {
								v52 = v74
								v53 = v53 + int32(2)
								continue
							} else {
								break
							}
							break
						}
						v77 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)) = uint8(v77)
						F_pg_cryptohash_free(m, v15)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v104 = int32(1)
							m.G0 = v10 + int32(16)
							return v104
						}
					}
				}
			}
		}
	}
}
func F_pg_mule2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	return v9
L2:
	;
	goto L3
L3:
	;
	v13 = l0
	v14 = l1
	v15 = l2
	v18 = v4
	goto L4
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v19 == int32(0) {
		v105 = v14
		v109 = v18
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(0)
	return v109
L6:
	;
	goto L5
L7:
	;
	if base.Ui32((v19+int32(127))&int32(255)) <= base.Ui32(int32(12)) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v93
	v98 = v18 + int32(1)
	v100 = v14 + int32(4)
	v101 = v15 + v94
	if int32(0) < v101 {
		v13 = v95
		v14 = v100
		v15 = v101
		v18 = v98
		goto L4
	} else {
		goto L26
	}
L9:
	;
	if v15 == int32(1) {
		v105 = v14
		v109 = v18
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v39 = v19 & int32(254)
	if v39 == int32(154) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v31 = v19 << (uint(int32(16)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v31
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v93 = v31 | v33
	v94 = int32(-2)
	v95 = v13 + int32(2)
	goto L8
L13:
	;
	v93 = v89
	v94 = int32(-3)
	v95 = v13 + int32(3)
	goto L8
L14:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v105 = v14
		v109 = v18
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if base.Ui32((v19+int32(112))&int32(255)) <= base.Ui32(int32(9)) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v46 = v44 << (uint(int32(16)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v46
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v89 = v46 | v48
	goto L13
L18:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v105 = v14
		v109 = v18
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v39 == int32(156) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v59 = v19 << (uint(int32(16)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v59
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v64 = v61<<(uint(int32(8))%32) | v59
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v64
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v89 = v64 | v66
	goto L13
L22:
	;
	if base.Ui32(v15) < base.Ui32(int32(4)) {
		v105 = v14
		v109 = v18
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v93 = v19
	v94 = int32(-1)
	v95 = v13 + int32(1)
	goto L8
L25:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v74 = v72 << (uint(int32(16)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v74
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v79 = v76<<(uint(int32(8))%32) | v74
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v79
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
	v93 = v79 | v81
	v94 = int32(-4)
	v95 = v13 + int32(4)
	goto L8
L26:
	;
	v105 = v100
	v109 = v98
	goto L6
}
func F_pg_my_temp_schema(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	return v3
}
func F_pg_node_tree_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(392021)
			F_errmsg(m, int32(182861), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(472671), int32(334), int32(266937))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_pg_notify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v4 = int32(717063)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v6 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_text_to_cstring(m, v10)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = v14
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v17 == int32(0) {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v21 = F_pg_detoast_datum_packed(m, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v23 = F_text_to_cstring(m, v21)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							v25 = v23
							F_PreventCommandDuringRecovery(m, int32(487384))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								F_Async_Notify(m, v16, v25)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									return int32(0)
								}
							}
						}
					}
				} else {
					v25 = v4
					F_PreventCommandDuringRecovery(m, int32(487384))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_Async_Notify(m, v16, v25)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		v16 = v4
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v17 == int32(0) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v21 = F_pg_detoast_datum_packed(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_text_to_cstring(m, v21)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = v23
					F_PreventCommandDuringRecovery(m, int32(487384))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_Async_Notify(m, v16, v25)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				}
			}
		} else {
			v25 = v4
			F_PreventCommandDuringRecovery(m, int32(487384))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				F_Async_Notify(m, v16, v25)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_pg_num_nulls(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v11 = F_count_nulls(m, l0, v5+int32(12), v5+int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
			v21 = int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			v21 = v20
		}
		m.G0 = v5 + int32(16)
		return v21
	}
}
func F_pg_parse_json_or_errsave(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v4 = F_pg_parse_json(m, l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 != 0 {
			F_json_errsave_error(m, v4, l0, l2)
			v9 = m.ExcPending
			if v9 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v4 == int32(0))
			}
		} else {
			return base.B2i32(v4 == int32(0))
		}
	}
}
func F_pg_perm_setlocale(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	if base.Ui32(int32(6)) < base.Ui32(l0) {
		v149 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v149 != 0 {
		goto L37
	} else {
		goto L38
	}
L2:
	;
	m.G0 = v17 + int32(48)
	goto L1
L3:
	;
	if l0 == int32(6) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v108 = int32(0)
	v109 = int32(4608624)
	v114 = v3
	goto L28
L5:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if l1 != 0 {
		goto L21
	} else {
		goto L22
	}
L8:
	;
	v26 = *(*int64)(unsafe.Add(mBase, _consts[467]))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v26
	v29 = *(*int64)(unsafe.Add(mBase, _consts[468]))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v29
	v32 = *(*int64)(unsafe.Add(mBase, _consts[469]))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v32
	v35 = int32(0)
	v36 = l1
	goto L10
L9:
	;
	v149 = int32(0)
	goto L2
L10:
	;
	v44 = F___strchrnul(m, v36, int32(59))
	mBase = m.M
	v45 = v44 - v36
	if v45 <= int32(23) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, _consts[470])) = v71
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
	*(*int64)(unsafe.Add(mBase, _consts[471])) = v74
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int64)(unsafe.Add(mBase, _consts[472])) = v77
	goto L4
L12:
	;
	v48 = F___memcpy(m, v17, v36, v45)
	mBase = m.M
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17+v45))) = uint8(v50)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v54 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v56 = v36
	goto L14
L14:
	;
	v57 = F___get_locale(m, v35, v17)
	mBase = m.M
	if v57 == int32(-1) {
		goto L9
	} else {
		goto L18
	}
L15:
	;
	v55 = v44 + int32(1)
	goto L17
L16:
	;
	v55 = v36
	goto L17
L17:
	;
	v56 = v55
	goto L14
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(24)+v35<<(uint(int32(2))%32)))) = v57
	v67 = v35 + int32(1)
	if v67 != int32(6) {
		v35 = v67
		v36 = v56
		goto L10
	} else {
		goto L19
	}
L19:
	;
	goto L11
L20:
	;
	if v93 != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v80 = F___get_locale(m, l0, l1)
	mBase = m.M
	if v80 == int32(-1) {
		v149 = v3
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[470])))
	v93 = v92
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[470]))) = v80
	v93 = v80
	goto L20
L25:
	;
	v97 = v93 + int32(8)
	goto L27
L26:
	;
	v97 = int32(522145)
	goto L27
L27:
	;
	v149 = v97
	goto L2
L28:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[470]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v108<<(uint(int32(2))%32))+uint32(_consts[470])))
	if v122 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v129))) = uint8(v140)
	if v135 != int32(6) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v126 = v122 + int32(8)
	goto L32
L31:
	;
	v126 = int32(522145)
	goto L32
L32:
	;
	v127 = F_strlen(m, v126)
	mBase = m.M
	v128 = F___memcpy(m, v109, v126, v127)
	mBase = m.M
	v129 = v109 + v127
	v130 = int32(59)
	*(*uint8)(unsafe.Add(mBase, uint32(v129))) = uint8(v130)
	v132 = int32(1)
	v135 = v114 + base.B2i32(v122 == v117)
	v137 = v108 + v132
	if v137 != int32(6) {
		v108 = v137
		v109 = v129 + v132
		v114 = v135
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	v145 = int32(4608624)
	goto L36
L35:
	;
	v145 = v126
	goto L36
L36:
	;
	v149 = v145
	goto L2
L37:
	;
	switch l0 {
	case 0:
		goto L41
	case 1:
		goto L44
	case 2:
		goto L43
	case 3:
		v304 = v149
		v305 = int32(516029)
		goto L40
	case 4:
		goto L45
	case 5:
		goto L46
	default:
		goto L42
	}
L38:
	;
	v437 = int32(0)
	goto L39
L39:
	;
	m.G0 = v7 + int32(16)
	return v437
L40:
	;
	v306 = int32(0)
	if v305 == v306 {
		goto L88
	} else {
		goto L89
	}
L41:
	;
	v177 = int32(4428288)
	goto L54
L42:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v304 = v149
	v305 = int32(517415)
	goto L40
L44:
	;
	v304 = v149
	v305 = int32(521874)
	goto L40
L45:
	;
	v304 = v149
	v305 = int32(486532)
	goto L40
L46:
	;
	v304 = v149
	v305 = int32(501333)
	goto L40
L47:
	;
	return int32(0)
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg_internal(m, int32(459644), v7)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(477403), int32(279), int32(379951))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	goto L83
L52:
	;
	v290 = F_strlen(m, v279)
	mBase = m.M
	goto L51
L54:
	;
	goto L55
L55:
	;
	v184 = int32(127)
	if (v177^v149)&int32(3) != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v283 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v283)
	goto L52
L57:
	;
	v264 = v259
	v265 = v260
	v266 = v261
	goto L79
L58:
	;
	if v254 == int32(0) {
		v279 = v252
		v280 = v253
		goto L56
	} else {
		goto L78
	}
L59:
	;
	v252 = v149
	v253 = v177
	v254 = v184
	goto L58
L60:
	;
	goto L61
L61:
	;
	if v149&int32(3) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if v221 == int32(0) {
		v279 = v218
		v280 = v219
		goto L56
	} else {
		goto L71
	}
L63:
	;
	v218 = v149
	v219 = v177
	v220 = v184
	v221 = int32(1)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v197 = v149
	v198 = v177
	v199 = v184
	goto L66
L66:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	*(*uint8)(unsafe.Add(mBase, uint32(v198))) = uint8(v201)
	if v201 == int32(0) {
		v259 = v197
		v260 = v198
		v261 = v199
		goto L57
	} else {
		goto L68
	}
L67:
	;
	v218 = v212
	v219 = v206
	v220 = v208
	v221 = v210
	goto L62
L68:
	;
	v205 = int32(1)
	v206 = v198 + v205
	v208 = v199 - v205
	v209 = int32(0)
	v210 = base.B2i32(v208 != v209)
	v212 = v197 + v205
	if v212&int32(3) == v209 {
		v218 = v212
		v219 = v206
		v220 = v208
		v221 = v210
		goto L62
	} else {
		goto L69
	}
L69:
	;
	if v208 != 0 {
		v197 = v212
		v198 = v206
		v199 = v208
		goto L66
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if v224 == int32(0) {
		v252 = v218
		v253 = v219
		v254 = v220
		goto L58
	} else {
		goto L72
	}
L72:
	;
	if base.Ui32(v220) < base.Ui32(int32(4)) {
		v252 = v218
		v253 = v219
		v254 = v220
		goto L58
	} else {
		goto L73
	}
L73:
	;
	v230 = v218
	v231 = v219
	v232 = v220
	goto L74
L74:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v238 = int32(-2139062144)
	if (int32(16843008)-v235|v235)&v238 != v238 {
		v259 = v230
		v260 = v231
		v261 = v232
		goto L57
	} else {
		goto L76
	}
L75:
	;
	v252 = v246
	v253 = v244
	v254 = v248
	goto L58
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231))) = v235
	v243 = int32(4)
	v244 = v231 + v243
	v246 = v230 + v243
	v248 = v232 - v243
	if base.Ui32(int32(3)) < base.Ui32(v248) {
		v230 = v246
		v231 = v244
		v232 = v248
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v259 = v252
	v260 = v253
	v261 = v254
	goto L57
L79:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	*(*uint8)(unsafe.Add(mBase, uint32(v265))) = uint8(v268)
	if v268 == int32(0) {
		v279 = v264
		v280 = v265
		goto L56
	} else {
		goto L81
	}
L80:
	;
	v279 = v275
	v280 = v273
	goto L56
L81:
	;
	v272 = int32(1)
	v273 = v265 + v272
	v275 = v264 + v272
	v277 = v266 - v272
	if v277 != 0 {
		v264 = v275
		v265 = v273
		v266 = v277
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1099])) = v296<<(uint(int32(3))%32) + int32(1801536)
	v304 = int32(4428288)
	v305 = int32(517166)
	goto L40
L84:
	;
	if v432 != 0 {
		goto L123
	} else {
		goto L124
	}
L85:
	;
	v338 = F___memcpy(m, v333, v305, v316)
	mBase = m.M
	v339 = v333 + v316
	v340 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v339))) = uint8(v340)
	v342 = int32(1)
	v346 = F___memcpy(m, v339+v342, v304, v329+v342)
	mBase = m.M
	v348 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	if v348 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L86:
	;
	v432 = int32(-1)
	goto L84
L87:
	;
	goto L92
L88:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(28)
	goto L86
L89:
	;
	v314 = F___strchrnul(m, v305, int32(61))
	mBase = m.M
	if v314 == v305 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v316 = v314 - v305
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305+v316))))
	if v318 == int32(0) {
		goto L87
	} else {
		goto L91
	}
L91:
	;
	goto L88
L92:
	;
	v329 = F_strlen(m, v304)
	mBase = m.M
	v333 = F_emscripten_builtin_malloc(m, v316+v329+int32(2))
	mBase = m.M
	if v333 != 0 {
		goto L85
	} else {
		goto L95
	}
L95:
	;
	goto L86
L96:
	;
	v432 = v425
	goto L84
L97:
	;
	v386 = v382 << (uint(int32(2)) % 32)
	v388 = v386 + int32(8)
	v390 = *(*int32)(unsafe.Add(mBase, _consts[475]))
	if v390 == v383 {
		goto L112
	} else {
		goto L113
	}
L98:
	;
	v362 = int32(0)
	v363 = v348
	v364 = v352
	goto L104
L99:
	;
	v382 = int32(0)
	v383 = v353
	goto L97
L100:
	;
	v353 = int32(0)
	goto L99
L101:
	;
	goto L102
L102:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	if v352 != 0 {
		goto L98
	} else {
		goto L103
	}
L103:
	;
	v353 = v348
	goto L99
L104:
	;
	v365 = F_strncmp(m, v333, v364, v316+int32(1))
	mBase = m.M
	if v365 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	v382 = v373
	v383 = v378
	goto L97
L106:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	*(*int32)(unsafe.Add(mBase, uint32(v363))) = v333
	F___env_rm_add(m, v368, v333)
	mBase = m.M
	v425 = int32(0)
	goto L96
L107:
	;
	goto L108
L108:
	;
	v373 = v362 + int32(1)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v363)+4))
	if v374 != 0 {
		v362 = v373
		v363 = v363 + int32(4)
		v364 = v374
		goto L104
	} else {
		goto L109
	}
L109:
	;
	goto L105
L110:
	;
	F_emscripten_builtin_free(m, v333)
	mBase = m.M
	v425 = int32(-1)
	goto L96
L111:
	;
	v405 = v402 + v382<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v405))) = v333
	*(*int32)(unsafe.Add(mBase, uint32(v405)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[474])) = v402
	*(*int32)(unsafe.Add(mBase, _consts[475])) = v402
	if v333 != 0 {
		goto L120
	} else {
		goto L121
	}
L112:
	;
	v392 = F_emscripten_builtin_realloc(m, v390, v388)
	mBase = m.M
	if v392 != 0 {
		v402 = v392
		goto L111
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v393 = F_emscripten_builtin_malloc(m, v388)
	mBase = m.M
	if v393 == int32(0) {
		goto L110
	} else {
		goto L116
	}
L115:
	;
	goto L110
L116:
	;
	if v382 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	v398 = F___memcpy(m, v393, v397, v386)
	mBase = m.M
	goto L119
L118:
	;
	goto L119
L119:
	;
	v400 = *(*int32)(unsafe.Add(mBase, _consts[475]))
	F_emscripten_builtin_free(m, v400)
	mBase = m.M
	v402 = v393
	goto L111
L120:
	;
	F___env_rm_add(m, int32(0), v333)
	mBase = m.M
	goto L122
L121:
	;
	goto L122
L122:
	;
	v425 = int32(0)
	goto L96
L123:
	;
	v433 = v306
	goto L125
L124:
	;
	v433 = v304
	goto L125
L125:
	;
	v437 = v433
	goto L39
}
func F_pg_range_sockaddr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	switch v7 - int32(2) {
	case 0:
		goto L3
	default:
		goto L1
	case 8:
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = int32(8)
	v19 = l2 + v18
	v21 = l1 + v18
	v23 = l0 + v18
	v25 = int32(0)
	goto L4
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	return base.B2i32(v10&(v11^v12) == int32(0))
L4:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v19))))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v21))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v23))))
	if v32&(v34^v36) != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	return int32(1)
L6:
	;
	v40 = v25 | int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v40))))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v40))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v19))))
	if (v42^v44)&v47 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v50 = v25 + int32(2)
	if v50 != int32(16) {
		v25 = v50
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_pg_read_binary_file_all(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_convert_and_check_filename(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v13 = F_read_binary_file(m, v8, int64(0), int64(-1), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if v13 == int32(0) {
					v17 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
					return int32(0)
				} else {
					return v13
				}
			}
		}
	}
}
func F_pg_read_file_all(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_convert_and_check_filename(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v13 = F_read_binary_file(m, v8, int64(0), int64(-1), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if v13 == int32(0) {
					v17 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
					return int32(0)
				} else {
					v21 = int32(4)
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					F_pg_verifymbstr(m, v13+v21, int32(base.Ui32(v23)>>(uint(int32(2))%32))-v21)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						return v13
					}
				}
			}
		}
	}
}
func F_pg_read_file_off_len(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		v13 = F_pg_read_file_common(m, v4, v9, v11, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if v13 == int32(0) {
				v17 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
				return int32(0)
			} else {
				return v13
			}
		}
	}
}
func F_pg_regcomp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v476 int32
	_ = v476
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v504 int32
	_ = v504
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v864 int32
	_ = v864
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v893 int32
	_ = v893
	var v906 int32
	_ = v906
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v932 int32
	_ = v932
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v989 int32
	_ = v989
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1008 int32
	_ = v1008
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	v12 = int32(16)
	if l0 == int32(0) {
		v1008 = v12
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(144)
	return v1008
L2:
	;
	if l1 == int32(0) {
		v1008 = v12
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l3&int32(3) == int32(2) {
		v1008 = v12
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if int32(base.Ui32(l3)>>(uint(int32(2))%32))&base.B2i32(l3&int32(227) != int32(0)) != 0 {
		v1008 = v12
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_pg_set_regex_collation(m, l4)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v10 + int32(52)
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1 + l2<<(uint(int32(2))%32)
	v52 = v37
	goto L8
L8:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v56+v52<<(uint(int32(2))%32)))) = int32(0)
	v63 = v52 + int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	if base.Ui32(v63) < base.Ui32(v64) {
		v52 = v63
		goto L8
	} else {
		goto L10
	}
L9:
	;
	v66 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+112)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v10)+132)) = v66
	v70 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+140)) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v66
	v74 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+100)) = uint16(v74)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+92)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v10)+124)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(1579080)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(17179869184)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(65239)
	v91 = F_palloc_extended(m, int32(432), int32(2))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v91
	if v91 == int32(0) {
		v827 = int32(12)
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v834 = v10 + int32(4)
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v834)))
	if v835 != 0 {
		goto L214
	} else {
		goto L215
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+72)) = int32(2166)
	v99 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+16)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v91)+200)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v91)+192)) = int64(0)
	v105 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+188)) = uint16(v105)
	*(*int64)(unsafe.Add(mBase, uint32(v91)+180)) = int64(4294969344)
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+88)) = uint16(v99)
	*(*int64)(unsafe.Add(mBase, uint32(v91)+80)) = int64(10)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+92)) = v91 + int32(180)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+76)) = v10 + int32(4)
	v121 = F_palloc_extended(m, int32(4096), int32(2))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+96)) = v121
	if v121 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v91 + int32(72)
	v181 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+20)) = v181
	*(*int64)(unsafe.Add(mBase, uint32(v91)+424)) = int64(0)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v10)+96))
	v189 = F_newnfa(m, v10+int32(4), v187, v181)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L30
	}
L16:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v91)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v128)+24)) = int32(101)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v91)+76))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	if v132 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v141 = F__emscripten_memset_bulkmem(m, v121, base.I32_extend8_s(int32(0)), int32(4096))
	mBase = m.M
	goto L22
L19:
	;
	v134 = v132
	goto L21
L20:
	;
	v134 = int32(12)
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+12)) = v134
	*(*int64)(unsafe.Add(mBase, uint32(v91)+160)) = int64(0)
	goto L15
L22:
	;
	v142 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v91)+156)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v91)+148)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v91)+140)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v91)+132)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v91)+124)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v91)+116)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v91)+108)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v91)+100)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v91)+176)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v91)+168)) = int64(4294967300)
	v164 = F_palloc_extended(m, int32(8), int32(2))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+164)) = v164
	if v164 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v91)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+24)) = int32(101)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v91)+76))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	if v173 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v177 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v164))) = uint16(v177)
	goto L15
L27:
	;
	v175 = v173
	goto L29
L28:
	;
	v175 = int32(12)
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+12)) = v175
	goto L15
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = v189
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v192 != 0 {
		v827 = v192
		goto L12
	} else {
		goto L31
	}
L31:
	;
	v195 = F_palloc_extended(m, int32(588), int32(2))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	if v195 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+124)) = int32(0)
	v827 = int32(12)
	goto L12
L34:
	;
	goto L35
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v195))) = int64(429496729600)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v195)+12)) = int64(85899345920)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+20)) = v195 + int32(428)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+8)) = v195 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+124)) = v195
	v215 = int32(4)
	v216 = v10 + v215
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	if v217&v215 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	if v615 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L37:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v220-v221 < int32(13) {
		v277 = v217
		v278 = v221
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v279 = int32(3)
	if v277&v279 != v279 {
		goto L36
	} else {
		goto L53
	}
L39:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	if v225 != int32(42) {
		v277 = v217
		v278 = v221
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v228 != int32(42) {
		v277 = v217
		v278 = v221
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v221)+8))
	if v231 != int32(42) {
		v277 = v217
		v278 = v221
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	switch v234 - int32(58) {
	case 0:
		goto L43
	default:
		goto L44
	case 3:
		goto L45
	case 5:
		goto L46
	}
L43:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+8)) = v265 | int32(128)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v271 = v269 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = v271
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v275 = v273 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+4)) = v275
	v277 = v271
	v278 = v275
	goto L38
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+24)) = int32(101)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	if v260 != 0 {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+8)) = v244 | int32(128)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+4)) = v248 + int32(16)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = v252&int32(-232) | int32(4)
	goto L36
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+24)) = int32(101)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	if v239 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v241 = v239
	goto L49
L48:
	;
	v241 = int32(2)
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+12)) = v241
	goto L36
L50:
	;
	v262 = v260
	goto L52
L51:
	;
	v262 = int32(13)
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+12)) = v262
	goto L36
L53:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	if v283-v278 < int32(9) {
		goto L36
	} else {
		goto L54
	}
L54:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	if v287 != int32(40) {
		goto L36
	} else {
		goto L55
	}
L55:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v290 != int32(63) {
		goto L36
	} else {
		goto L56
	}
L56:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v295 = *(*int32)(unsafe.Add(mBase, _consts[632]))
	switch v295 - int32(1) {
	case 0:
		goto L60
	case 1:
		goto L59
	case 2:
		goto L58
	default:
		goto L61
	}
L57:
	;
	if v395 == int32(0) {
		goto L36
	} else {
		goto L83
	}
L58:
	;
	if base.Ui32(int32(255)) < base.Ui32(v293) {
		goto L36
	} else {
		goto L81
	}
L59:
	;
	if base.Ui32(v293) <= base.Ui32(int32(131071)) {
		goto L78
	} else {
		goto L79
	}
L60:
	;
	if base.Ui32(int32(127)) < base.Ui32(v293) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	if base.Ui32(int32(127)) < base.Ui32(v293) {
		goto L36
	} else {
		goto L62
	}
L62:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+uint32(_consts[633]))))
	v303 = int32(1)
	v395 = int32(base.Ui32(v302)>>(uint(v303)%32)) & v303
	goto L57
L63:
	;
	v395 = v353
	goto L57
L64:
	;
	v315 = int32(0)
	v316 = int32(1178)
	goto L67
L65:
	;
	goto L66
L66:
	;
	v341 = int32(1)
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293<<(uint(v341)%32))+uint32(_consts[634]))))
	v353 = v345 & v341
	goto L63
L67:
	;
	v321 = base.I32_div_s(v315+v316, int32(2))
	v323 = v321 << (uint(int32(3)) % 32)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v323)+uint32(_consts[635])))
	if base.Ui32(v326) < base.Ui32(v293) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v353 = int32(0)
	goto L63
L69:
	;
	if v337 <= v338 {
		v315 = v337
		v316 = v338
		goto L67
	} else {
		goto L76
	}
L70:
	;
	v337 = v321 + int32(1)
	v338 = v316
	goto L69
L71:
	;
	goto L72
L72:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v323)+uint32(_consts[636])))
	if base.Ui32(v332) <= base.Ui32(v293) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v353 = int32(1)
	goto L63
L74:
	;
	goto L75
L75:
	;
	v337 = v315
	v338 = v321 - int32(1)
	goto L69
L76:
	;
	goto L68
L77:
	;
	v395 = v381
	goto L57
L78:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v293)>>(uint(int32(8))%32)))+uint32(_consts[637]))))
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v293)>>(uint(int32(3))%32))&int32(31)|v367<<(uint(int32(5))%32))+uint32(_consts[637]))))
	v381 = int32(base.Ui32(v373)>>(uint(v293&int32(7))%32)) & int32(1)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v381 = base.B2i32(base.Ui32(v293) < base.Ui32(int32(196606)))
	goto L77
L81:
	;
	goto L82
L82:
	;
	v395 = base.B2i32(base.B2i32(base.Ui32(v293|int32(32)-int32(97)) < base.Ui32(int32(26))) != int32(0))
	goto L57
L83:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v398)+8)) = v399 | int32(128)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v405 = v403 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+4)) = v405
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	if base.Ui32(v407) <= base.Ui32(v405) {
		v582 = v405
		v586 = v407
		goto L84
	} else {
		goto L85
	}
L84:
	;
	if base.Ui32(v582) < base.Ui32(v586) {
		goto L135
	} else {
		goto L136
	}
L85:
	;
	v411 = v405
	v415 = v407
	goto L86
L86:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	v418 = *(*int32)(unsafe.Add(mBase, _consts[632]))
	switch v418 - int32(1) {
	case 0:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	default:
		goto L92
	}
L87:
	;
	v582 = v576
	v586 = v578
	goto L84
L88:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v518 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L89:
	;
	if base.Ui32(int32(255)) < base.Ui32(v416) {
		v582 = v411
		v586 = v415
		goto L84
	} else {
		goto L112
	}
L90:
	;
	if base.Ui32(v416) <= base.Ui32(int32(131071)) {
		goto L109
	} else {
		goto L110
	}
L91:
	;
	if base.Ui32(int32(127)) < base.Ui32(v416) {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	if base.Ui32(int32(127)) < base.Ui32(v416) {
		v582 = v411
		v586 = v415
		goto L84
	} else {
		goto L93
	}
L93:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+uint32(_consts[633]))))
	v426 = int32(1)
	v518 = int32(base.Ui32(v425)>>(uint(v426)%32)) & v426
	goto L88
L94:
	;
	v518 = v476
	goto L88
L95:
	;
	v438 = int32(0)
	v439 = int32(1178)
	goto L98
L96:
	;
	goto L97
L97:
	;
	v464 = int32(1)
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416<<(uint(v464)%32))+uint32(_consts[634]))))
	v476 = v468 & v464
	goto L94
L98:
	;
	v444 = base.I32_div_s(v438+v439, int32(2))
	v446 = v444 << (uint(int32(3)) % 32)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v446)+uint32(_consts[635])))
	if base.Ui32(v449) < base.Ui32(v416) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v476 = int32(0)
	goto L94
L100:
	;
	if v460 <= v461 {
		v438 = v460
		v439 = v461
		goto L98
	} else {
		goto L107
	}
L101:
	;
	v460 = v444 + int32(1)
	v461 = v439
	goto L100
L102:
	;
	goto L103
L103:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v446)+uint32(_consts[636])))
	if base.Ui32(v455) <= base.Ui32(v416) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v476 = int32(1)
	goto L94
L105:
	;
	goto L106
L106:
	;
	v460 = v438
	v461 = v444 - int32(1)
	goto L100
L107:
	;
	goto L99
L108:
	;
	v518 = v504
	goto L88
L109:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v416)>>(uint(int32(8))%32)))+uint32(_consts[637]))))
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v416)>>(uint(int32(3))%32))&int32(31)|v490<<(uint(int32(5))%32))+uint32(_consts[637]))))
	v504 = int32(base.Ui32(v496)>>(uint(v416&int32(7))%32)) & int32(1)
	goto L108
L110:
	;
	goto L111
L111:
	;
	v504 = base.B2i32(base.Ui32(v416) < base.Ui32(int32(196606)))
	goto L108
L112:
	;
	goto L113
L113:
	;
	v518 = base.B2i32(base.B2i32(base.Ui32(v416|int32(32)-int32(97)) < base.Ui32(int32(26))) != int32(0))
	goto L88
L114:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	v582 = v519
	v586 = v522
	goto L84
L115:
	;
	goto L116
L116:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	switch v523 - int32(98) {
	case 0:
		goto L118
	case 1:
		goto L129
	default:
		goto L119
	case 3:
		goto L128
	case 7:
		goto L127
	case 11, 12:
		goto L126
	case 14:
		goto L125
	case 15:
		goto L124
	case 17:
		goto L123
	case 18:
		goto L122
	case 21:
		goto L121
	case 22:
		goto L120
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = v573
	v576 = v519 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+4)) = v576
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	if base.Ui32(v576) < base.Ui32(v578) {
		v411 = v576
		v415 = v578
		goto L86
	} else {
		goto L133
	}
L118:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v573 = v570 & int32(-8)
	goto L117
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+24)) = int32(101)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	if v566 != 0 {
		goto L130
	} else {
		goto L131
	}
L120:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v573 = v561 | int32(32)
	goto L117
L121:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v573 = v556&int32(-193) | int32(128)
	goto L117
L122:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v573 = v553 & int32(-33)
	goto L117
L123:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v573 = v550 & int32(-193)
	goto L117
L124:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v573 = v545&int32(-8) | int32(4)
	goto L117
L125:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v573 = v540&int32(-193) | int32(64)
	goto L117
L126:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v573 = v537 | int32(192)
	goto L117
L127:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v573 = v534 | int32(8)
	goto L117
L128:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v573 = v529&int32(-8) | int32(1)
	goto L117
L129:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v573 = v526 & int32(-9)
	goto L117
L130:
	;
	v568 = v566
	goto L132
L131:
	;
	v568 = int32(18)
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+12)) = v568
	goto L36
L133:
	;
	goto L87
L134:
	;
	v597 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+4)) = v582 + v597
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	if v600&v597 == int32(0) {
		goto L36
	} else {
		goto L142
	}
L135:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v582)))
	if v588 == int32(41) {
		goto L134
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+24)) = int32(101)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	if v593 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	goto L137
L139:
	;
	v595 = v593
	goto L141
L140:
	;
	v595 = int32(18)
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+12)) = v595
	goto L36
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = v600 & int32(-225)
	goto L36
L143:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	if v618&int32(4) != 0 {
		goto L147
	} else {
		goto L148
	}
L144:
	;
	goto L145
L145:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
	if v634&int32(192) != 0 {
		goto L154
	} else {
		goto L155
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+24)) = int32(110)
	v631 = F_next(m, v216)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L6
	} else {
		goto L153
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+32)) = int32(3)
	goto L146
L148:
	;
	goto L149
L149:
	;
	if v618&int32(1) != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+32)) = int32(1)
	goto L146
L151:
	;
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+32)) = int32(2)
	goto L146
L153:
	;
	goto L145
L154:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v10)+96))
	v639 = F_subcolor(m, v637, int32(10))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L6
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v646 != 0 {
		v827 = v646
		goto L12
	} else {
		goto L159
	}
L157:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+100)) = uint16(v639)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v10)+96))
	F_okcolors(m, v642, v643)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L6
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v651)+4))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v651)+8))
	v654 = F_parse(m, v10+int32(4), int32(101), int32(112), v652, v653)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = v654
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v657 != 0 {
		v827 = v657
		goto L12
	} else {
		goto L161
	}
L161:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	F_specialcolors(m, v658)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L6
	} else {
		goto L162
	}
L162:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v661 != 0 {
		v827 = v661
		goto L12
	} else {
		goto L163
	}
L163:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
	if v662&int32(16) != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	F_removecaptures(m, v10+int32(4), v667)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L6
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v670 = int32(1)
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v671)+4)) = v670
	v676 = int32(2)
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v671)+20))
	if v677 != 0 {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	goto L166
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+120)) = v685
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687)+1)))
	v690 = v688 | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v687)+1)) = uint8(v690)
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v687)+20))
	if v692 != 0 {
		goto L176
	} else {
		goto L177
	}
L169:
	;
	v679 = v677
	v680 = v676
	goto L172
L170:
	;
	v685 = v676
	goto L171
L171:
	;
	goto L168
L172:
	;
	v681 = F_numst(m, v679, v680)
	mBase = m.M
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v679)+24))
	if v682 != 0 {
		v679 = v682
		v680 = v681
		goto L172
	} else {
		goto L174
	}
L173:
	;
	v685 = v681
	goto L171
L174:
	;
	goto L173
L175:
	;
	v698 = v10 + int32(4)
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v698)+108))
	if v699 != 0 {
		goto L182
	} else {
		goto L183
	}
L176:
	;
	v693 = v692
	goto L179
L177:
	;
	goto L178
L178:
	;
	goto L175
L179:
	;
	F_markst(m, v693)
	mBase = m.M
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v693)+24))
	if v695 != 0 {
		v693 = v695
		goto L179
	} else {
		goto L181
	}
L180:
	;
	goto L178
L181:
	;
	goto L180
L182:
	;
	v701 = v699
	goto L185
L183:
	;
	goto L184
L184:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v698)+108)) = int64(0)
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	v727 = F_nfatree(m, v10+int32(4), v726)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L6
	} else {
		goto L192
	}
L185:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v701)+84))
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701)+1)))
	if v708&int32(64) == int32(0) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	goto L184
L187:
	;
	F_pfree(m, v701)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L6
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	if v707 != 0 {
		v701 = v707
		goto L185
	} else {
		goto L191
	}
L190:
	;
	goto L189
L191:
	;
	goto L186
L192:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v727 | v729
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v732 != 0 {
		v827 = v732
		goto L12
	} else {
		goto L193
	}
L193:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v10)+136))
	if int32(2) <= v733 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v742 = v670
	goto L197
L195:
	;
	goto L196
L196:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768)+1)))
	if v769&int32(2) != 0 {
		goto L202
	} else {
		goto L203
	}
L197:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	v748 = v745 + v742*int32(88)
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v748)+2)))
	v754 = F_nfanode(m, v10+int32(4), v748, base.B2i32(v749&int32(2) == int32(0)))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L6
	} else {
		goto L199
	}
L198:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v760 != 0 {
		v827 = v760
		goto L12
	} else {
		goto L201
	}
L199:
	;
	v757 = v742 + int32(1)
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v10)+136))
	if v757 < v758 {
		v742 = v757
		goto L197
	} else {
		goto L200
	}
L200:
	;
	goto L198
L201:
	;
	goto L196
L202:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v772)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v772)+8)) = v773 | int32(8192)
	goto L204
L203:
	;
	goto L204
L204:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v779 = F_optimize(m, v778)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L6
	} else {
		goto L205
	}
L205:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v781 != 0 {
		v827 = v781
		goto L12
	} else {
		goto L206
	}
L206:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	F_makesearch(m, v10+int32(4), v784)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L6
	} else {
		goto L207
	}
L207:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v787 != 0 {
		v827 = v787
		goto L12
	} else {
		goto L208
	}
L208:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	F_compact(m, v788, v91+int32(20))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L6
	} else {
		goto L209
	}
L209:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v793 != 0 {
		v827 = v793
		goto L12
	} else {
		goto L210
	}
L210:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v794
	v796 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v796
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(65241)
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v801
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v803
	v805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v805
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+16)) = v807
	*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = v796
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v10)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+68)) = v811
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v815&int32(8) != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v818 = int32(967)
	goto L213
L212:
	;
	v818 = int32(968)
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+420)) = v818
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+424)) = v820
	*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = int32(0)
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v10)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+428)) = v824
	v827 = v796
	goto L12
L214:
	;
	F_rfree(m, v835)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L6
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v834)+40))
	if v838 != v10+int32(52) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	goto L216
L218:
	;
	F_pfree(m, v838)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L6
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v834)+88))
	if v844 != 0 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	goto L220
L222:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v844)+36))
	if v845 != 0 {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L224
L224:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v834)+104))
	if v914 != 0 {
		goto L240
	} else {
		goto L241
	}
L225:
	;
	v846 = v845
	goto L228
L226:
	;
	goto L227
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v844)+36)) = int32(0)
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v844)+40))
	if v874 != 0 {
		goto L232
	} else {
		goto L233
	}
L228:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v846)))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v844)+76))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v854)+136))
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v846)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v854)+136)) = v855 + v856*int32(-36) - int32(8)
	F_pfree(m, v846)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L6
	} else {
		goto L230
	}
L229:
	;
	goto L227
L230:
	;
	if v853 != 0 {
		v846 = v853
		goto L228
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	v875 = v874
	goto L235
L233:
	;
	goto L234
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v844)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v844)+40)) = int32(0)
	F_pfree(m, v844)
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L6
	} else {
		goto L239
	}
L235:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v844)+76))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v883)+136))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v875)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v883)+136)) = v884 + v885*int32(-40) - int32(8)
	F_pfree(m, v875)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L6
	} else {
		goto L237
	}
L236:
	;
	goto L234
L237:
	;
	if v882 != 0 {
		v875 = v882
		goto L235
	} else {
		goto L238
	}
L238:
	;
	goto L236
L239:
	;
	goto L224
L240:
	;
	F_freesubre(m, v834, v914)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L6
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v834)+108))
	if v917 != 0 {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	goto L242
L244:
	;
	v918 = v917
	goto L247
L245:
	;
	goto L246
L246:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v834)+120))
	if v942 != 0 {
		goto L254
	} else {
		goto L255
	}
L247:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v918)+84))
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v918)+1)))
	if v926&int32(64) == int32(0) {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v834)+108)) = int64(0)
	goto L246
L249:
	;
	F_pfree(m, v918)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L6
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	if v925 != 0 {
		v918 = v925
		goto L247
	} else {
		goto L253
	}
L252:
	;
	goto L251
L253:
	;
	goto L248
L254:
	;
	F_pfree(m, v942)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L6
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v834)+124))
	if v945 != 0 {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	goto L256
L258:
	;
	F_pfree(m, v945)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L6
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v834)+128))
	if v948 != 0 {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	goto L260
L262:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v834)+132))
	v951 = v949 - int32(1)
	if int32(0) < v951 {
		goto L265
	} else {
		goto L266
	}
L263:
	;
	goto L264
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v834)+24)) = int32(101)
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v834)+12))
	if v999 != 0 {
		goto L278
	} else {
		goto L279
	}
L265:
	;
	v954 = v948
	v955 = v951
	goto L268
L266:
	;
	goto L267
L267:
	;
	F_pfree(m, v948)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L6
	} else {
		goto L277
	}
L268:
	;
	v962 = v954 + int32(124)
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v962)))
	if v963 != 0 {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	goto L267
L270:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v954)+152))
	F_pfree(m, v964)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L6
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v977 = int32(1)
	if v977 < v955 {
		v954 = v954 + int32(88)
		v955 = v955 - v977
		goto L268
	} else {
		goto L276
	}
L273:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v954)+156))
	F_pfree(m, v967)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L6
	} else {
		goto L274
	}
L274:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v954)+160))
	F_pfree(m, v970)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L6
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v962))) = int32(0)
	goto L272
L276:
	;
	goto L269
L277:
	;
	goto L264
L278:
	;
	v1000 = v999
	goto L280
L279:
	;
	v1000 = v827
	goto L280
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v834)+12)) = v1000
	v1008 = v1000
	goto L1
}
func F_pg_rewrite_query(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[829])))
	if v8 == int32(1) {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, _consts[830])))
		F_elog_node_display(m, int32(392118), l0, v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, _consts[831])))
			if v19 == int32(1) {
				v22 = int32(4367496)
				v23 = int32(0)
				v27 = m.G0
				v29 = v27 - int32(16)
				m.G0 = v29
				v32 = int32(4367512)
				v37 = F___memset(m, int32(4367520), v23, int32(144))
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, _consts[832])) = int32(4)
				*(*int64)(unsafe.Add(mBase, _consts[833])) = int64(3)
				*(*int32)(unsafe.Add(mBase, _consts[834])) = int32(2)
				*(*int64)(unsafe.Add(mBase, _consts[835])) = int64(1)
				v50 = F___memcpy(m, v29, v32, int32(16))
				mBase = m.M
				v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v29))))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				v53 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[836])) = v53
				*(*int32)(unsafe.Add(mBase, _consts[837])) = v52
				*(*int64)(unsafe.Add(mBase, _consts[838])) = v51
				v57 = int64(*(*int32)(unsafe.Add(mBase, uint32(v29)+8)))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
				*(*int32)(unsafe.Add(mBase, _consts[839])) = v53
				*(*int32)(unsafe.Add(mBase, _consts[834])) = v58
				*(*int64)(unsafe.Add(mBase, _consts[835])) = v57
				v65 = F___syscall_ret(m, v23)
				mBase = m.M
				m.G0 = v29 + int32(16)
				F___gettimeofday(m, int32(4367648))
				mBase = m.M
			} else {
			}
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v71 == int32(6) {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = l0
				v79 = F_list_make1_impl(m, int32(1), v5+int32(8))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					v83 = v79
					v85 = int32(*(*uint8)(unsafe.Add(mBase, _consts[831])))
					if v85 == int32(1) {
						F_ShowUsage(m, int32(501727))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							v92 = int32(*(*uint8)(unsafe.Add(mBase, _consts[840])))
							if v92 == int32(1) {
								v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[830])))
								F_elog_node_display(m, int32(392108), v83, v97)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									m.G0 = v5 + int32(16)
									return v83
								}
							} else {
								m.G0 = v5 + int32(16)
								return v83
							}
						}
					} else {
						v92 = int32(*(*uint8)(unsafe.Add(mBase, _consts[840])))
						if v92 == int32(1) {
							v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[830])))
							F_elog_node_display(m, int32(392108), v83, v97)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								m.G0 = v5 + int32(16)
								return v83
							}
						} else {
							m.G0 = v5 + int32(16)
							return v83
						}
					}
				}
			} else {
				v81 = F_QueryRewrite(m, l0)
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					v83 = v81
					v85 = int32(*(*uint8)(unsafe.Add(mBase, _consts[831])))
					if v85 == int32(1) {
						F_ShowUsage(m, int32(501727))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							v92 = int32(*(*uint8)(unsafe.Add(mBase, _consts[840])))
							if v92 == int32(1) {
								v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[830])))
								F_elog_node_display(m, int32(392108), v83, v97)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									m.G0 = v5 + int32(16)
									return v83
								}
							} else {
								m.G0 = v5 + int32(16)
								return v83
							}
						}
					} else {
						v92 = int32(*(*uint8)(unsafe.Add(mBase, _consts[840])))
						if v92 == int32(1) {
							v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[830])))
							F_elog_node_display(m, int32(392108), v83, v97)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								m.G0 = v5 + int32(16)
								return v83
							}
						} else {
							m.G0 = v5 + int32(16)
							return v83
						}
					}
				}
			}
		}
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, _consts[831])))
		if v19 == int32(1) {
			v22 = int32(4367496)
			v23 = int32(0)
			v27 = m.G0
			v29 = v27 - int32(16)
			m.G0 = v29
			v32 = int32(4367512)
			v37 = F___memset(m, int32(4367520), v23, int32(144))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, _consts[832])) = int32(4)
			*(*int64)(unsafe.Add(mBase, _consts[833])) = int64(3)
			*(*int32)(unsafe.Add(mBase, _consts[834])) = int32(2)
			*(*int64)(unsafe.Add(mBase, _consts[835])) = int64(1)
			v50 = F___memcpy(m, v29, v32, int32(16))
			mBase = m.M
			v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v29))))
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
			v53 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[836])) = v53
			*(*int32)(unsafe.Add(mBase, _consts[837])) = v52
			*(*int64)(unsafe.Add(mBase, _consts[838])) = v51
			v57 = int64(*(*int32)(unsafe.Add(mBase, uint32(v29)+8)))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
			*(*int32)(unsafe.Add(mBase, _consts[839])) = v53
			*(*int32)(unsafe.Add(mBase, _consts[834])) = v58
			*(*int64)(unsafe.Add(mBase, _consts[835])) = v57
			v65 = F___syscall_ret(m, v23)
			mBase = m.M
			m.G0 = v29 + int32(16)
			F___gettimeofday(m, int32(4367648))
			mBase = m.M
		} else {
		}
		v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v71 == int32(6) {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = l0
			v79 = F_list_make1_impl(m, int32(1), v5+int32(8))
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int32(0)
			} else {
				v83 = v79
				v85 = int32(*(*uint8)(unsafe.Add(mBase, _consts[831])))
				if v85 == int32(1) {
					F_ShowUsage(m, int32(501727))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						v92 = int32(*(*uint8)(unsafe.Add(mBase, _consts[840])))
						if v92 == int32(1) {
							v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[830])))
							F_elog_node_display(m, int32(392108), v83, v97)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								m.G0 = v5 + int32(16)
								return v83
							}
						} else {
							m.G0 = v5 + int32(16)
							return v83
						}
					}
				} else {
					v92 = int32(*(*uint8)(unsafe.Add(mBase, _consts[840])))
					if v92 == int32(1) {
						v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[830])))
						F_elog_node_display(m, int32(392108), v83, v97)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							m.G0 = v5 + int32(16)
							return v83
						}
					} else {
						m.G0 = v5 + int32(16)
						return v83
					}
				}
			}
		} else {
			v81 = F_QueryRewrite(m, l0)
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				v83 = v81
				v85 = int32(*(*uint8)(unsafe.Add(mBase, _consts[831])))
				if v85 == int32(1) {
					F_ShowUsage(m, int32(501727))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						v92 = int32(*(*uint8)(unsafe.Add(mBase, _consts[840])))
						if v92 == int32(1) {
							v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[830])))
							F_elog_node_display(m, int32(392108), v83, v97)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								m.G0 = v5 + int32(16)
								return v83
							}
						} else {
							m.G0 = v5 + int32(16)
							return v83
						}
					}
				} else {
					v92 = int32(*(*uint8)(unsafe.Add(mBase, _consts[840])))
					if v92 == int32(1) {
						v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[830])))
						F_elog_node_display(m, int32(392108), v83, v97)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							m.G0 = v5 + int32(16)
							return v83
						}
					} else {
						m.G0 = v5 + int32(16)
						return v83
					}
				}
			}
		}
	}
}
func F_pg_role_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l2&int64(2199023255552) == int64(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v179
L2:
	;
	if l2&int64(512) != int64(0) {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	v14 = F_superuser_arg(m, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v179 = v4
	goto L1
L7:
	;
	goto L8
L8:
	;
	if l0 == l1 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v22 = F_roles_is_member_of(m, l1, int32(0), l0, v8+int32(12))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v24 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L2
L12:
	;
	if l0 == l1 {
		v179 = v4
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if l2&int64(256) != int64(0) {
		goto L33
	} else {
		goto L34
	}
L15:
	;
	v30 = F_superuser_arg(m, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v30 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v32 = int32(0)
	v35 = F_roles_is_member_of(m, l1, v32, v32, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v37 = int32(0)
	if v35 == v37 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v75 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L32
	}
L20:
	;
	v75 = int32(0)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v43 <= int32(0) {
		v68 = v37
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v75 = v68
	goto L19
L24:
	;
	v46 = int32(0)
	if v46 < v43 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v49 = v43
	goto L27
L26:
	;
	v49 = v46
	goto L27
L27:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v52 = int32(0)
	goto L28
L28:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v50+v52<<(uint(int32(2))%32))))
	v61 = base.B2i32(v60 == l0)
	if v60 == l0 {
		v68 = v61
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v68 = v61
	goto L23
L30:
	;
	v63 = v52 + int32(1)
	if v63 != v49 {
		v52 = v63
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	goto L14
L33:
	;
	if l0 == l1 {
		v179 = v4
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if l2&int64(4096) != int64(0) {
		goto L54
	} else {
		goto L55
	}
L36:
	;
	v81 = F_superuser_arg(m, l1)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v81 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v84 = int32(0)
	v86 = F_roles_is_member_of(m, l1, int32(1), v84, v84)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v88 = int32(0)
	if v86 == v88 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v126 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L53
	}
L41:
	;
	v126 = int32(0)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v94 <= int32(0) {
		v119 = v88
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v126 = v119
	goto L40
L45:
	;
	v97 = int32(0)
	if v97 < v94 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v100 = v94
	goto L48
L47:
	;
	v100 = v97
	goto L48
L48:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v103 = int32(0)
	goto L49
L49:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(2))%32))))
	v112 = base.B2i32(v111 == l0)
	if v111 == l0 {
		v119 = v112
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v119 = v112
	goto L44
L51:
	;
	v114 = v103 + int32(1)
	if v114 != v100 {
		v103 = v114
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	goto L35
L54:
	;
	if l0 == l1 {
		v179 = v4
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v179 = int32(1)
	goto L1
L57:
	;
	v132 = F_superuser_arg(m, l1)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	if v132 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v135 = int32(0)
	v137 = F_roles_is_member_of(m, l1, int32(2), v135, v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v139 = int32(0)
	if v137 == v139 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v177 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L74
	}
L62:
	;
	v177 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v145 <= int32(0) {
		v170 = v139
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v177 = v170
	goto L61
L66:
	;
	v148 = int32(0)
	if v148 < v145 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v151 = v145
	goto L69
L68:
	;
	v151 = v148
	goto L69
L69:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v154 = int32(0)
	goto L70
L70:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v152+v154<<(uint(int32(2))%32))))
	v163 = base.B2i32(v162 == l0)
	if v162 == l0 {
		v170 = v163
		goto L65
	} else {
		goto L72
	}
L71:
	;
	v170 = v163
	goto L65
L72:
	;
	v165 = v154 + int32(1)
	if v165 != v151 {
		v154 = v165
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	goto L56
}
func F_pg_rusage_show(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int64
	_ = v84
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int64
	_ = v97
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(192)
	m.G0 = v14
	v17 = v14 + int32(40)
	v22 = m.G0
	v24 = v22 - int32(16)
	m.G0 = v24
	v27 = v14 + int32(56)
	v32 = F___memset(m, v14+int32(64), v2, int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+16)) = int64(3)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = int64(1)
	v45 = F___memcpy(m, v24, v27, int32(16))
	mBase = m.M
	v46 = int64(*(*int32)(unsafe.Add(mBase, uint32(v24))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v46
	v52 = int64(*(*int32)(unsafe.Add(mBase, uint32(v24)+8)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v52
	v60 = F___syscall_ret(m, v2)
	mBase = m.M
	m.G0 = v24 + int32(16)
	F___gettimeofday(m, v14+int32(24))
	mBase = m.M
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v67 < v68 {
		v71 = v67 + int32(1000000)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v71
		v73 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v73 - int64(1)
		v77 = v71
	} else {
		v77 = v67
	}
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v78 < v79 {
		v82 = v78 + int32(1000000)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v82
		v84 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
		*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = v84 - int64(1)
		v88 = v82
	} else {
		v88 = v78
	}
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v90 < v91 {
		v94 = v90 + int32(1000000)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v94
		v97 = v89 - int64(1)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v97
		v99 = v94
		v100 = v97
	} else {
		v99 = v90
		v100 = v89
	}
	v101 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v104 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v105 = v103 - v104
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+16)) = uint32(v105)
	v108 = int32(10000)
	v109 = base.I32_div_s(v77-v68, v108)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v109
	v111 = v100 - v102
	*(*uint32)(unsafe.Add(mBase, uint32(v14))) = uint32(v111)
	v115 = base.I32_div_s(v99-v91, v108)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v115
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
	v118 = v117 - v101
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+8)) = uint32(v118)
	v122 = base.I32_div_s(v88-v79, v108)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v122
	v127 = F_pg_snprintf(m, int32(4442672), int32(100), int32(196372), v14)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		return int32(0)
	} else {
		m.G0 = v14 + int32(192)
		return int32(4442672)
	}
}
func F_pg_sjis_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v7 = int32(1)
	v10 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	v19 = base.B2i32(int32(0) <= v10) | base.B2i32(base.Ui32((v10+int32(95))&int32(255)) < base.Ui32(int32(63)))
	if v19 != 0 {
		v20 = v7
	} else {
		v20 = int32(2)
	}
	v21 = base.B2i32(l1 < v20)
	if l1 < v20 {
		v22 = int32(-1)
	} else {
		v22 = v7
	}
	if v19 != 0 {
		v51 = v22
	} else {
		if l1 < v20 {
			v51 = v22
		} else {
			if base.Ui32(int32(31)) <= base.Ui32((v10+int32(127))&int32(255)) {
				if base.Ui32(int32(28)) < base.Ui32((v10+int32(32))&int32(255)) {
					v51 = int32(-1)
				} else {
					v37 = int32(2)
					v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
					if base.Ui32((v40+int32(-64))&int32(255)) < base.Ui32(int32(63)) {
						v47 = v37
					} else {
						v47 = int32(-1)
					}
					if v40 < int32(-3) {
						v50 = v37
					} else {
						v50 = v47
					}
					v51 = v50
				}
			} else {
				v37 = int32(2)
				v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
				if base.Ui32((v40+int32(-64))&int32(255)) < base.Ui32(int32(63)) {
					v47 = v37
				} else {
					v47 = int32(-1)
				}
				if v40 < int32(-3) {
					v50 = v37
				} else {
					v50 = v47
				}
				v51 = v50
			}
		}
	}
	return v51
}
func F_pg_sleep(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v59 float64
	_ = v59
	var v68 float64
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	F___gettimeofday(m, v11)
	mBase = m.M
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+8)))
	m.G0 = v11 + v10
	goto L1
L1:
	;
	goto L2
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return int32(0)
L4:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v41 = m.G0
	v42 = int32(16)
	v43 = v41 - v42
	m.G0 = v43
	F___gettimeofday(m, v43)
	mBase = m.M
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	v47 = int64(*(*int32)(unsafe.Add(mBase, uint32(v43)+8)))
	m.G0 = v43 + v42
	goto L11
L7:
	;
	return int32(0)
L8:
	;
	goto L6
L9:
	;
	goto L3
L10:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	v80 = F_WaitLatch(m, v77, int32(41), v74, int32(150994946))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v59 = base.F64_add(base.F64_add(v5, base.F64_div(base.F64_convert_i64_s(v15+v14*int64(1000000)-int64(946684800000000)), float64(1e+06))), base.F64_div(base.F64_convert_i64_s(v47+v46*int64(1000000)-int64(946684800000000)), float64(-1e+06)))
	if base.F64_ge(v59, float64(600)) != 0 {
		v74 = int32(600000)
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if base.F64_gt(v59, float64(0)) == int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v68 = base.F64_ceil(base.F64_mul(v59, float64(1000)))
	if base.F64_lt(base.F64_abs(v68), float64(2.147483648e+09)) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v72 = base.I32_trunc_f64_s(v68)
	v74 = v72
	goto L10
L15:
	;
	goto L16
L16:
	;
	v74 = int32(-2147483648)
	goto L10
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
	goto L18
L18:
	;
	goto L2
}
func F_pg_strerror(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_pg_strerror_r(m, l0, int32(4527680))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pg_strfromd(m *base.Module, l0 int32, l1 int32, l2 float64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v46 float64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+108)) = uint8(v4)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+100)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = l0
	v21 = l0 + int32(31)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v21
	v23 = base.I64_reinterpret_f64(l2)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v23&int64(9223372036854775807)) {
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(5136718)
		v95 = int32(3)
		F_dostr(m, v12+int32(16), v95, v12+int32(88))
		mBase = m.M
		v103 = m.ExcPending
		if v103 != 0 {
			return
		} else {
			v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+108)))
			v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
			v108 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v107))) = uint8(v108)
			if v104&int32(1) != 0 {
			} else {
			}
			m.G0 = v12 + int32(112)
			return
		}
	} else {
		if base.F64_lt(l2, float64(0)) != 0 {
			v46 = base.F64_neg(l2)
			v47 = v4
			v48 = int32(45)
		} else {
			if base.B2i32(base.Ui64(v23) < base.Ui64(int64(9218868437227405313)))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v23&int64(9223372036854775807))) == int32(0) {
				v46 = base.F64_neg(l2)
				v47 = v4
				v48 = int32(45)
			} else {
				v46 = l2
				v47 = int32(1)
				v48 = int32(0)
			}
		}
		if base.F64_eq(base.F64_abs(v46), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v53 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1266])))
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)) = uint8(v53)
			v56 = *(*int64)(unsafe.Add(mBase, _consts[1267]))
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v56
			v84 = int32(8)
			if v47 != 0 {
				v95 = v84
			} else {
				if v21 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = l0 + int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v48)
					v95 = v84
				} else {
					if base.Ui32(l0) < base.Ui32(v21) {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = l0 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v48)
						v95 = v84
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = int32(1)
						v95 = v84
					}
				}
			}
			F_dostr(m, v12+int32(16), v95, v12+int32(88))
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return
			} else {
				v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+108)))
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
				v108 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v107))) = uint8(v108)
				if v104&int32(1) != 0 {
				} else {
				}
				m.G0 = v12 + int32(112)
				return
			}
		} else {
			v59 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+84)) = uint8(v59)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = int32(1730817573)
			v64 = int32(32)
			if base.Ui32(v64) <= base.Ui32(l1) {
				v67 = v64
			} else {
				v67 = l1
			}
			if l1 <= int32(0) {
				v70 = int32(1)
			} else {
				v70 = v67
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v70
			*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v46
			v78 = F_snprintf(m, v12+int32(16), int32(64), v12+int32(80), v12)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return
			} else {
				if int32(0) <= v78 {
					v84 = v78
					if v47 != 0 {
						v95 = v84
					} else {
						if v21 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = l0 + int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v48)
							v95 = v84
						} else {
							if base.Ui32(l0) < base.Ui32(v21) {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = l0 + int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v48)
								v95 = v84
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = int32(1)
								v95 = v84
							}
						}
					}
					F_dostr(m, v12+int32(16), v95, v12+int32(88))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+108)))
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
						v108 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v107))) = uint8(v108)
						if v104&int32(1) != 0 {
						} else {
						}
						m.G0 = v12 + int32(112)
						return
					}
				} else {
					v82 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v82)
					m.G0 = v12 + int32(112)
					return
				}
			}
		}
	}
}
func F_pg_strncoll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = m.T0[v7].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_pg_strong_random(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v3
	v15 = F_open(m, int32(275211), v3, v9)
	mBase = m.M
	if v15 != int32(-1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(1)
	if l1 == int32(0) {
		v41 = v18
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v48 = v3
	goto L3
L3:
	;
	m.G0 = v9 + int32(16)
	return v48
L4:
	;
	v43 = F_close(m, v15)
	mBase = m.M
	v48 = v41
	goto L3
L5:
	;
	v21 = l0
	v22 = l1
	goto L6
L6:
	;
	v27 = F_read(m, v15, v21, v22)
	mBase = m.M
	if v27 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v41 = v18
	goto L4
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v31 == int32(27) {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v36 = v22 - v27
	if v36 != 0 {
		v21 = v21 + v27
		v22 = v36
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v41 = int32(0)
	goto L4
L12:
	;
	goto L7
}
func F_pg_switch_wal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[29])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[30]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+316))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[29])) = uint8(v12)
		v14 = v12
	} else {
		v14 = int32(0)
	}
	if v14 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(119370), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(534619), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(473384), int32(185), int32(294623))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
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
		v38 = F_RequestXLogSwitch(m, int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			v40 = F_Int64GetDatum(m, v38)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				return v40
			}
		}
	}
}
func F_pg_tablespace_databases(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_InitMaterializedSRF(m, l0, int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	switch v13 - int32(1663) {
	case 0:
		v47 = int32(346516)
		goto L5
	case 1:
		goto L7
	default:
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L39
	}
L4:
	;
	m.G0 = v10 + int32(80)
	return int32(0)
L5:
	;
	v48 = F_AllocateDir(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = int32(530717)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(467736)
	v45 = F_psprintf(m, int32(167383), v10+int32(48))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	v24 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v24 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_errmsg(m, int32(152040), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(478044), int32(237), int32(152016))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L4
L12:
	;
	v47 = v45
	goto L5
L13:
	;
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v50 = F_ReadDir(m, v48, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v104 != int32(44) {
		goto L3
	} else {
		goto L34
	}
L17:
	;
	if v50 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v52 = v50
	goto L21
L19:
	;
	goto L20
L20:
	;
	F_FreeDir(m, v48)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L33
	}
L21:
	;
	v60 = v52 + int32(19)
	v64 = F_strtox_2(m, v60, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	v65 = base.I32_wrap_i64(v64)
	goto L24
L22:
	;
	goto L20
L23:
	;
	v92 = F_ReadDir(m, v48, v47)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L31
	}
L24:
	;
	if v65 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v47
	v73 = F_psprintf(m, int32(167509), v10+int32(32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v75 = F_directory_is_empty(m, v73)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_pfree(m, v73)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v75 != 0 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+75)) = uint8(v79)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+76)) = v65
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_tuplestore_putvalues(m, v82, v83, v10+int32(76), v10+int32(75))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L23
L31:
	;
	if v92 != 0 {
		v52 = v92
		goto L21
	} else {
		goto L32
	}
L32:
	;
	goto L22
L33:
	;
	goto L4
L34:
	;
	v109 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v109 == int32(0) {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
	F_errmsg(m, int32(520694), v10)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(478044), int32(259), int32(152016))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L4
L39:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v47
	F_errmsg(m, int32(283058), v10+int32(16))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(478044), int32(257), int32(152016))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_toupper(m *base.Module, l0 int32) int32 {
	var v10 int32
	_ = v10
	if base.Ui32((l0-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		v10 = l0 - int32(32)
	} else {
		v10 = l0
	}
	return v10 & int32(255)
}
func F_pg_tz_acceptable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(946684800)
	v13 = F_localsub(m, l0+int32(256), v5+int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			v21 = base.B2i32(v17 == int32(0))
		} else {
			v21 = int32(0)
		}
		m.G0 = v5 + int32(16)
		return v21
	}
}
func F_pg_ultostr_zeropad(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v301 int32
	_ = v301
	if base.Ui32(int32(99)) < base.Ui32(l1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	if l2 != int32(2) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v13)
	return l0 + int32(2)
L4:
	;
	if l2 <= v148 {
		goto L22
	} else {
		goto L23
	}
L5:
	;
	v27 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v27)
	v148 = int32(1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1233)
	v38 = int32(base.Ui32((base.I32_clz(l1)^int32(31))*v33+v33) >> (uint(int32(12)) % 32))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38<<(uint(int32(2))%32))+uint32(_consts[1098])))
	v45 = v38 + base.B2i32(base.Ui32(v43) <= base.Ui32(l1))
	v46 = int32(0)
	if base.Ui32(l1) < base.Ui32(int32(10000)) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if base.Ui32(v93) < base.Ui32(int32(100)) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v92 = v46
	v93 = l1
	goto L8
L10:
	;
	goto L11
L11:
	;
	v50 = l1
	v52 = v46
	goto L12
L12:
	;
	v59 = l0 + v45 - v52
	v60 = int32(4)
	v63 = base.I32_div_u_s(v50, int32(10000))
	v66 = v63*int32(-10000) + v50
	v67 = int32(100)
	v68 = base.I32_div_u_s(v66, v67)
	v69 = int32(1)
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68<<(uint(v69)%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v59-v60))) = uint16(v73)
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v66-v68*v67)<<(uint(v69)%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v59-int32(2)))) = uint16(v84)
	v87 = v52 + v60
	if base.Ui32(int32(99999999)) < base.Ui32(v50) {
		v50 = v63
		v52 = v87
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v92 = v87
	v93 = v63
	goto L8
L14:
	;
	goto L13
L15:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v122) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v122 = v93
	v123 = v92
	goto L15
L17:
	;
	goto L18
L18:
	;
	v103 = int32(2)
	v105 = int32(65535)
	v107 = int32(100)
	v108 = base.I32_div_u_s(v93&v105, v107)
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v93-v108*v107)&v105<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v45-v92-v103))) = uint16(v118)
	v122 = v108
	v123 = v92 | v103
	goto L15
L19:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v45-v123-int32(2)))) = uint16(v134)
	v148 = v45
	goto L4
L20:
	;
	goto L21
L21:
	;
	v137 = v122 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v137)
	v148 = v45
	goto L4
L22:
	;
	return l0 + v148
L23:
	;
	goto L24
L24:
	;
	v152 = l0 + l2
	v153 = v152 - v148
	if v153 == l0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v301 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(48)), l2-v148)
	mBase = m.M
	goto L71
L26:
	;
	goto L25
L27:
	;
	v157 = v153 + v148
	if base.Ui32(l0-v157) <= base.Ui32(int32(0)-v148<<(uint(int32(1))%32)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v164 = F___memcpy(m, v153, l0, v148)
	mBase = m.M
	goto L25
L29:
	;
	goto L30
L30:
	;
	v167 = (v153 ^ l0) & int32(3)
	if base.Ui32(v153) < base.Ui32(l0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if v269 == int32(0) {
		goto L26
	} else {
		goto L67
	}
L32:
	;
	if base.Ui32(v247) <= base.Ui32(int32(3)) {
		v268 = v246
		v269 = v247
		v270 = v248
		goto L31
	} else {
		goto L63
	}
L33:
	;
	if v167 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	if v167 != 0 {
		v229 = v148
		goto L46
	} else {
		goto L47
	}
L36:
	;
	v268 = l0
	v269 = v148
	v270 = v153
	goto L31
L37:
	;
	goto L38
L38:
	;
	if v153&int32(3) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v246 = l0
	v247 = v148
	v248 = v153
	goto L32
L40:
	;
	goto L41
L41:
	;
	v174 = l0
	v175 = v148
	v176 = v153
	goto L42
L42:
	;
	if v175 == int32(0) {
		goto L26
	} else {
		goto L44
	}
L43:
	;
	v246 = v183
	v247 = v185
	v248 = v187
	goto L32
L44:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	*(*uint8)(unsafe.Add(mBase, uint32(v176))) = uint8(v180)
	v182 = int32(1)
	v183 = v174 + v182
	v185 = v175 - v182
	v187 = v176 + v182
	if v187&int32(3) != 0 {
		v174 = v183
		v175 = v185
		v176 = v187
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	if v229 == int32(0) {
		goto L26
	} else {
		goto L59
	}
L47:
	;
	if v157&int32(3) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v194 = v148
	goto L51
L49:
	;
	v209 = v148
	goto L50
L50:
	;
	if base.Ui32(v209) <= base.Ui32(int32(3)) {
		v229 = v209
		goto L46
	} else {
		goto L55
	}
L51:
	;
	if v194 == int32(0) {
		goto L26
	} else {
		goto L53
	}
L52:
	;
	v209 = v200
	goto L50
L53:
	;
	v200 = v194 - int32(1)
	v201 = v153 + v200
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v200))))
	*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v203)
	if v201&int32(3) != 0 {
		v194 = v200
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v216 = v209
	goto L56
L56:
	;
	v220 = v216 - int32(4)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0+v220)))
	*(*int32)(unsafe.Add(mBase, uint32(v153+v220))) = v223
	if base.Ui32(int32(3)) < base.Ui32(v220) {
		v216 = v220
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v229 = v220
	goto L46
L58:
	;
	goto L57
L59:
	;
	v236 = v229
	goto L60
L60:
	;
	v240 = v236 - int32(1)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v240))))
	*(*uint8)(unsafe.Add(mBase, uint32(v153+v240))) = uint8(v243)
	if v240 != 0 {
		v236 = v240
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L26
L62:
	;
	goto L61
L63:
	;
	v253 = v246
	v254 = v247
	v255 = v248
	goto L64
L64:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v257
	v259 = int32(4)
	v260 = v253 + v259
	v262 = v255 + v259
	v264 = v254 - v259
	if base.Ui32(int32(3)) < base.Ui32(v264) {
		v253 = v260
		v254 = v264
		v255 = v262
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v268 = v260
	v269 = v264
	v270 = v262
	goto L31
L66:
	;
	goto L65
L67:
	;
	v275 = v268
	v276 = v269
	v277 = v270
	goto L68
L68:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	*(*uint8)(unsafe.Add(mBase, uint32(v277))) = uint8(v279)
	v281 = int32(1)
	v286 = v276 - v281
	if v286 != 0 {
		v275 = v275 + v281
		v276 = v286
		v277 = v277 + v281
		goto L68
	} else {
		goto L70
	}
L69:
	;
	goto L26
L70:
	;
	goto L69
L71:
	;
	return v152
}
func F_pg_unicode_to_server(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	v1 = l0
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if base.Ui32(v1-int32(1)) < base.Ui32(int32(1114111)) {
		if base.Ui32(v1) <= base.Ui32(int32(127)) {
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v1)
			v16 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v16)
			m.G0 = v7 + int32(16)
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[356]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			if v20 == int32(6) {
				if base.Ui32(v1) <= base.Ui32(int32(2047)) {
					v28 = int32(base.Ui32(v1)>>(uint(int32(6))%32)) | int32(192)
					*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v28)
					v66 = int32(1)
				} else {
					if base.Ui32(v1) <= base.Ui32(int32(65535)) {
						v36 = int32(base.Ui32(v1)>>(uint(int32(12))%32)) | int32(224)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v36)
						v43 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&int32(63) | int32(128)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v43)
						v66 = int32(2)
					} else {
						v49 = int32(base.Ui32(v1)>>(uint(int32(18))%32)) | int32(240)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v49)
						v53 = int32(63)
						v55 = int32(128)
						v56 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&v53 | v55
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)) = uint8(v56)
						v63 = int32(base.Ui32(v1)>>(uint(int32(12))%32))&v53 | v55
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v63)
						v66 = int32(3)
					}
				}
				v71 = v1&int32(63) | int32(128)
				*(*uint8)(unsafe.Add(mBase, uint32(v66+l1))) = uint8(v71)
				v73 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
				if int32(0) <= v73 {
					v97 = int32(1)
				} else {
					v78 = v73 & int32(255)
					if v78&int32(224) == int32(192) {
						v97 = int32(2)
					} else {
						if v78&int32(240) == int32(224) {
							v97 = int32(3)
						} else {
							if v78&int32(248) == int32(240) {
								v95 = int32(4)
							} else {
								v95 = int32(1)
							}
							v97 = v95
						}
					}
				}
				v99 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v97+l1))) = uint8(v99)
				m.G0 = v7 + int32(16)
				return
			} else {
				v102 = *(*int32)(unsafe.Add(mBase, _consts[1201]))
				if v102 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v221 = m.ExcPending
					if v221 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v224 = m.ExcPending
						if v224 != 0 {
							return
						} else {
							v226 = *(*int32)(unsafe.Add(mBase, _consts[356]))
							v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
							v229 = *(*int32)(unsafe.Add(mBase, _consts[1202]))
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v229
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v227
							F_errmsg(m, int32(423013), v7)
							mBase = m.M
							v234 = m.ExcPending
							if v234 != 0 {
								return
							} else {
								F_errfinish(m, int32(472485), int32(911), int32(203937))
								mBase = m.M
								v239 = m.ExcPending
								if v239 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if base.Ui32(v1) <= base.Ui32(int32(2047)) {
						v110 = int32(base.Ui32(v1)>>(uint(int32(6))%32)) | int32(192)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+11)) = uint8(v110)
						v151 = v7 + int32(12)
					} else {
						if base.Ui32(v1) <= base.Ui32(int32(65535)) {
							v119 = int32(base.Ui32(v1)>>(uint(int32(12))%32)) | int32(224)
							*(*uint8)(unsafe.Add(mBase, uint32(v7)+11)) = uint8(v119)
							v126 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&int32(63) | int32(128)
							*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v126)
							v151 = v7 + int32(13)
						} else {
							v133 = int32(base.Ui32(v1)>>(uint(int32(18))%32)) | int32(240)
							*(*uint8)(unsafe.Add(mBase, uint32(v7)+11)) = uint8(v133)
							v137 = int32(63)
							v139 = int32(128)
							v140 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&v137 | v139
							*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v140)
							v147 = int32(base.Ui32(v1)>>(uint(int32(12))%32))&v137 | v139
							*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v147)
							v151 = v7 + int32(14)
						}
					}
					v155 = v1&int32(63) | int32(128)
					*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v155)
					v159 = int32(*(*int8)(unsafe.Add(mBase, uint32(v7+int32(11)))))
					if int32(0) <= v159 {
						v183 = int32(1)
					} else {
						v164 = v159 & int32(255)
						if v164&int32(224) == int32(192) {
							v183 = int32(2)
						} else {
							if v164&int32(240) == int32(224) {
								v183 = int32(3)
							} else {
								if v164&int32(248) == int32(240) {
									v181 = int32(4)
								} else {
									v181 = int32(1)
								}
								v183 = v181
							}
						}
					}
					v185 = v7 + int32(11)
					v187 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v183+v185))) = uint8(v187)
					v190 = *(*int32)(unsafe.Add(mBase, _consts[1201]))
					v195 = F_FunctionCall6Coll(m, v190, int32(6), v20, v185, l1, v183, v187)
					mBase = m.M
					v196 = m.ExcPending
					if v196 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v205 = m.ExcPending
		if v205 != 0 {
			return
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v208 = m.ExcPending
			if v208 != 0 {
				return
			} else {
				F_errmsg(m, int32(84747), int32(0))
				mBase = m.M
				v212 = m.ExcPending
				if v212 != 0 {
					return
				} else {
					F_errfinish(m, int32(472485), int32(886), int32(203937))
					mBase = m.M
					v217 = m.ExcPending
					if v217 != 0 {
						return
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
func F_pg_unicode_to_server_noerror(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	v1 = l0
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if base.Ui32(int32(1114110)) < base.Ui32(v1-int32(1)) {
		v204 = v3
		m.G0 = v8 + int32(16)
		return v204
	} else {
		if base.Ui32(v1) <= base.Ui32(int32(127)) {
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v1)
			v18 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v18)
			v204 = int32(1)
			m.G0 = v8 + int32(16)
			return v204
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[356]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
			if v22 == int32(6) {
				if base.Ui32(v1) <= base.Ui32(int32(2047)) {
					v30 = int32(base.Ui32(v1)>>(uint(int32(6))%32)) | int32(192)
					*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v30)
					v68 = int32(1)
				} else {
					if base.Ui32(v1) <= base.Ui32(int32(65535)) {
						v38 = int32(base.Ui32(v1)>>(uint(int32(12))%32)) | int32(224)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v38)
						v45 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&int32(63) | int32(128)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v45)
						v68 = int32(2)
					} else {
						v51 = int32(base.Ui32(v1)>>(uint(int32(18))%32)) | int32(240)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v51)
						v55 = int32(63)
						v57 = int32(128)
						v58 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&v55 | v57
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)) = uint8(v58)
						v65 = int32(base.Ui32(v1)>>(uint(int32(12))%32))&v55 | v57
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v65)
						v68 = int32(3)
					}
				}
				v73 = v1&int32(63) | int32(128)
				*(*uint8)(unsafe.Add(mBase, uint32(v68+l1))) = uint8(v73)
				v75 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
				if int32(0) <= v75 {
					v99 = int32(1)
				} else {
					v80 = v75 & int32(255)
					if v80&int32(224) == int32(192) {
						v99 = int32(2)
					} else {
						if v80&int32(240) == int32(224) {
							v99 = int32(3)
						} else {
							if v80&int32(248) == int32(240) {
								v97 = int32(4)
							} else {
								v97 = int32(1)
							}
							v99 = v97
						}
					}
				}
				v101 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v99+l1))) = uint8(v101)
				v204 = int32(1)
				m.G0 = v8 + int32(16)
				return v204
			} else {
				v105 = *(*int32)(unsafe.Add(mBase, _consts[1201]))
				if v105 == int32(0) {
					v204 = v3
					m.G0 = v8 + int32(16)
					return v204
				} else {
					if base.Ui32(v1) <= base.Ui32(int32(2047)) {
						v113 = int32(base.Ui32(v1)>>(uint(int32(6))%32)) | int32(192)
						*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v113)
						v154 = v8 + int32(12)
					} else {
						if base.Ui32(v1) <= base.Ui32(int32(65535)) {
							v122 = int32(base.Ui32(v1)>>(uint(int32(12))%32)) | int32(224)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v122)
							v129 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&int32(63) | int32(128)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)) = uint8(v129)
							v154 = v8 + int32(13)
						} else {
							v136 = int32(base.Ui32(v1)>>(uint(int32(18))%32)) | int32(240)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v136)
							v140 = int32(63)
							v142 = int32(128)
							v143 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&v140 | v142
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+13)) = uint8(v143)
							v150 = int32(base.Ui32(v1)>>(uint(int32(12))%32))&v140 | v142
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)) = uint8(v150)
							v154 = v8 + int32(14)
						}
					}
					v158 = v1&int32(63) | int32(128)
					*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v158)
					v162 = int32(*(*int8)(unsafe.Add(mBase, uint32(v8+int32(11)))))
					if int32(0) <= v162 {
						v186 = int32(1)
					} else {
						v167 = v162 & int32(255)
						if v167&int32(224) == int32(192) {
							v186 = int32(2)
						} else {
							if v167&int32(240) == int32(224) {
								v186 = int32(3)
							} else {
								if v167&int32(248) == int32(240) {
									v184 = int32(4)
								} else {
									v184 = int32(1)
								}
								v186 = v184
							}
						}
					}
					v188 = v8 + int32(11)
					v190 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v186+v188))) = uint8(v190)
					v193 = *(*int32)(unsafe.Add(mBase, _consts[1201]))
					v198 = F_FunctionCall6Coll(m, v193, int32(6), v22, v188, l1, v186, int32(1))
					mBase = m.M
					v201 = m.ExcPending
					if v201 != 0 {
						return int32(0)
					} else {
						v204 = base.B2i32(v198 == v186)
						m.G0 = v8 + int32(16)
						return v204
					}
				}
			}
		}
	}
}
func F_pg_visible_in_snapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v12
	v19 = int32(1)
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	if base.Ui64(v12) < base.Ui64(v20) {
		v63 = v19
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v9 + int32(16)
	return v63
L4:
	;
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	if base.Ui64(v22) <= base.Ui64(v12) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v63 = int32(0)
	goto L3
L6:
	;
	v25 = v14 + int32(24)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if base.Ui32(v26) <= base.Ui32(int32(30)) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v40 = int32(0)
	goto L13
L8:
	;
	if v26 == int32(0) {
		v63 = v19
		goto L3
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v32 = int32(8)
	v36 = F_bsearch(m, v9+v32, v25, v26, v32, int32(1561))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L7
L12:
	;
	v63 = base.B2i32(v36 == int32(0))
	goto L3
L13:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v25+v40<<(uint(int32(3))%32))))
	if v12 == v49 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v63 = v19
	goto L3
L15:
	;
	v52 = v40 + int32(1)
	if v26 != v52 {
		v40 = v52
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
}
func F_pg_wchar2euc_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v9)
	return v9
L2:
	;
	goto L3
L3:
	;
	v13 = l0
	v14 = l1
	v15 = l2
	v17 = v4
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v71)
	return v69
L6:
	;
	if base.Ui32(int32(16777216)) <= base.Ui32(v19) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v66 = v14
	v69 = v17
	goto L8
L8:
	;
	goto L5
L9:
	;
	v57 = v56 + v14
	v58 = int32(1)
	v62 = v56 + v17
	if v58 < v15 {
		v13 = v13 + int32(4)
		v14 = v57
		v15 = v15 - v58
		v17 = v62
		goto L4
	} else {
		goto L19
	}
L10:
	;
	v23 = int32(base.Ui32(v19) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v23)
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)) = uint8(v25)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v29 = int32(base.Ui32(v27) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)) = uint8(v29)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+3)) = uint8(v31)
	v56 = int32(4)
	goto L9
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(int32(65536)) <= base.Ui32(v19) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v37 = int32(base.Ui32(v19) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v37)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v41 = int32(base.Ui32(v39) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)) = uint8(v41)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)) = uint8(v43)
	v56 = int32(3)
	goto L9
L14:
	;
	goto L15
L15:
	;
	if base.Ui32(int32(256)) <= base.Ui32(v19) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v49 = int32(base.Ui32(v19) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v49)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)) = uint8(v51)
	v56 = int32(2)
	goto L9
L17:
	;
	goto L18
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v19)
	v56 = int32(1)
	goto L9
L19:
	;
	v66 = v57
	v69 = v62
	goto L8
}
func F_pg_wchar2mb_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v5 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6*int32(28))+uint32(_consts[1204])))
	v12 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
func F_pg_wchar2mule_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v10)
	return v10
L2:
	;
	goto L3
L3:
	;
	v14 = l0
	v15 = l1
	v16 = l2
	v19 = v4
	goto L4
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v112)
	return v110
L6:
	;
	v23 = int32(base.Ui32(v21) >> (uint(int32(16)) % 32))
	v25 = v23 & int32(255)
	if base.Ui32(v25-int32(129)) <= base.Ui32(int32(12)) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v106 = v15
	v110 = v19
	goto L8
L8:
	;
	goto L5
L9:
	;
	v97 = v96 + v15
	v98 = int32(1)
	v102 = v96 + v19
	if v98 < v16 {
		v14 = v14 + int32(4)
		v15 = v97
		v16 = v16 - v98
		v19 = v102
		goto L4
	} else {
		goto L28
	}
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v23)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v31)
	v96 = int32(2)
	goto L9
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(v25-int32(144)) <= base.Ui32(int32(9)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v23)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v41 = int32(base.Ui32(v39) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v41)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)) = uint8(v43)
	v96 = int32(3)
	goto L9
L14:
	;
	goto L15
L15:
	;
	if base.Ui32(v25-int32(160)) <= base.Ui32(int32(63)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v50 = int32(154)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v50)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v23)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)) = uint8(v53)
	v96 = int32(3)
	goto L9
L17:
	;
	goto L18
L18:
	;
	if v21&int32(15728640) == int32(14680064) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v60 = int32(155)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v60)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v23)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)) = uint8(v63)
	v96 = int32(3)
	goto L9
L20:
	;
	goto L21
L21:
	;
	if base.Ui32(v25-int32(240)) <= base.Ui32(int32(4)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v70 = int32(156)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v70)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v23)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v75 = int32(base.Ui32(v73) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)) = uint8(v75)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+3)) = uint8(v77)
	v96 = int32(4)
	goto L9
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(v25) < base.Ui32(int32(245)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v21)
	v96 = int32(1)
	goto L9
L26:
	;
	if v25 == int32(255) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v84 = int32(157)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v84)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v23)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v89 = int32(base.Ui32(v87) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)) = uint8(v89)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+3)) = uint8(v91)
	v96 = int32(4)
	goto L9
L28:
	;
	v106 = v97
	v110 = v102
	goto L8
}
