package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_call_string_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(224)
	m.G0 = v15
	v26 = v6
	v27 = int32(-1)
	v28 = v6
	v29 = v6
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v27 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L2
L5:
	;
	v158 = int32(m.ExcTag)
	v159 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v158 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[0])) = v53
	*(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[1])) = v54
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_bms_free(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L45
	}
L7:
	;
	m.G0 = v15 + int32(224)
	return v140 & int32(1)
L8:
	;
	v34 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+223)) = uint8(v34)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v37 == int32(0) {
		v140 = v34
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v52 = v26
	v53 = v28
	v54 = v29
	goto L10
L10:
	;
	if v52 != 0 {
		goto L6
	} else {
		goto L16
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[0]))
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[1]))
	goto L12
L12:
	;
	v46 = v15 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v15 + int32(60)
	goto L15
L13:
	;
	v52 = int32(0)
	v53 = v42
	v54 = v44
	goto L10
L15:
	;
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[2])) = int32(50856066)
	*(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[1])) = v15 - int32(-64)
	v63 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[3])) = v63
	*(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[4])) = v63
	*(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[5])) = v63
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(100))))
	v72 = m.T0[v71].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l3)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v72 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v77 = F_errstart(m, l4, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[0])) = v53
	*(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[1])) = v54
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+223)))
	v140 = v135
	goto L7
L21:
	;
	if v77 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[2]))
	F_errcode(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_FlushErrorState(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L44
	}
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[3]))
	if v84 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[4]))
	if v105 != 0 {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v84
	F_errmsg_internal(m, int32(_a_F_call_string_check_hook_0), v15+int32(48))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v92 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L26
L31:
	;
	v94 = v92
	goto L33
L32:
	;
	v94 = int32(_a_F_call_string_check_hook_1)
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v91
	F_errmsg(m, int32(_a_F_call_string_check_hook_2), v15+int32(32))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	goto L26
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v105
	F_errdetail_internal(m, int32(_a_F_call_string_check_hook_0), v15+int32(16))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_call_string_check_hook[5]))
	if v113 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v113
	F_errhint(m, int32(_a_F_call_string_check_hook_0), v15)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_errfinish(m, int32(_a_F_call_string_check_hook_3), int32(_a_F_call_string_check_hook_4), int32(_a_F_call_string_check_hook_5))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	goto L24
L44:
	;
	v127 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+223)) = uint8(v127)
	goto L20
L45:
	;
	F_pg_re_throw(m)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	goto L4
L47:
	;
	v163 = int32(v159)
	m.G0 = v15
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	if v15+int32(60) == v169 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	m.ExcPending = 1
	goto L56
L49:
	;
	if v173 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v173 = v171
	goto L52
L51:
	;
	v173 = int32(0)
	goto L52
L52:
	;
	goto L49
L53:
	;
	F___wasm_longjmp(m, v166, v165)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v26 = v165
	v27 = v173
	v28 = v53
	v29 = v54
	goto L1
L56:
	;
	return int32(0)
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_initStringInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v4 = F_palloc(m, int32(1024))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1024)
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v4
		v9 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v4))) = uint8(v9)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
		return
	}
}
func F_stringToQualifiedNameList(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_pstrdup(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v91
L2:
	;
	return int32(0)
L3:
	;
	v17 = F_SplitIdentifierString(m, v10, int32(46), v8+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v17 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = int32(0)
	v22 = F_errsave_start(m, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	if v22 == int32(0) {
		v91 = v21
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	F_errmsg(m, int32(_a_F_stringToQualifiedNameList_0), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errsave_finish(m, l1, int32(_a_F_stringToQualifiedNameList_1), int32(1810), int32(_a_F_stringToQualifiedNameList_2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v91 = v21
	goto L1
L13:
	;
	v39 = int32(0)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v39 < v40 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v74 = int32(0)
	v75 = F_errsave_start(m, l1)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L27
	}
L16:
	;
	v44 = v39
	v45 = int32(0)
	goto L19
L17:
	;
	v64 = v39
	goto L18
L18:
	;
	F_pfree(m, v10)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L25
	}
L19:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v45<<(uint(int32(2))%32))))
	v54 = F_pstrdup(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L21
	}
L20:
	;
	v64 = v58
	goto L18
L21:
	;
	v56 = F_makeString(m, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v58 = F_lappend(m, v44, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v61 = v45 + int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v61 < v62 {
		v44 = v58
		v45 = v61
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	F_list_free(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v91 = v64
	goto L1
L27:
	;
	if v75 == int32(0) {
		v91 = v74
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	F_errmsg(m, int32(_a_F_stringToQualifiedNameList_0), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	F_errsave_finish(m, l1, int32(_a_F_stringToQualifiedNameList_1), int32(1815), int32(_a_F_stringToQualifiedNameList_2))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v91 = v74
	goto L1
}
func F_string_to_const(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 == int32(19) {
		v16 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), l0)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v44 = v16
			v45 = int32(950)
			v46 = int32(64)
			v48 = int32(0)
			v50 = F_makeConst(m, l1, int32(-1), v45, v46, v44, v48, v48)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return v50
			}
		}
	} else {
		v22 = F_cstring_to_text(m, l0)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if base.B2i32(l1 == int32(25))|base.B2i32(base.Ui32(l1-int32(1042)) < base.Ui32(int32(2))) != 0 {
				v44 = v22
				v45 = int32(100)
				v46 = int32(-1)
				v48 = int32(0)
				v50 = F_makeConst(m, l1, int32(-1), v45, v46, v44, v48, v48)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return v50
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
					F_errmsg_internal(m, int32(_a_F_string_to_const_0), v8)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_string_to_const_1), int32(1765), int32(_a_F_string_to_const_2))
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
func F_string_to_privilege(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
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
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int64
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int64
	_ = v485
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(_a_F_string_to_privilege_0)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[0])))
	if base.B2i32(v11 == int32(0))|base.B2i32(v11 != v14) != 0 {
		v32 = v11
		v33 = v14
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L157
	} else {
		goto L158
	}
L2:
	;
	m.G0 = v6 + int32(16)
	return v485
L3:
	;
	if v32-v33 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	goto L3
L5:
	;
	v17 = l0
	v18 = v8
	goto L6
L6:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v22 == int32(0) {
		v32 = v22
		v33 = v21
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v32 = v22
	v33 = v21
	goto L4
L8:
	;
	v25 = int32(1)
	if v22 == v21 {
		v17 = v17 + v25
		v18 = v18 + v25
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v485 = int64(1)
	goto L2
L11:
	;
	goto L12
L12:
	;
	v38 = int32(_a_F_string_to_privilege_1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[1])))
	if base.B2i32(v41 == int32(0))|base.B2i32(v41 != v44) != 0 {
		v62 = v41
		v63 = v44
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v62-v63 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	goto L13
L15:
	;
	v47 = l0
	v48 = v38
	goto L16
L16:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	if v52 == int32(0) {
		v62 = v52
		v63 = v51
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v62 = v52
	v63 = v51
	goto L14
L18:
	;
	v55 = int32(1)
	if v52 == v51 {
		v47 = v47 + v55
		v48 = v48 + v55
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v485 = int64(2)
	goto L2
L21:
	;
	goto L22
L22:
	;
	v68 = int32(_a_F_string_to_privilege_2)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[2])))
	if base.B2i32(v71 == int32(0))|base.B2i32(v71 != v74) != 0 {
		v92 = v71
		v93 = v74
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v92-v93 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	goto L23
L25:
	;
	v77 = l0
	v78 = v68
	goto L26
L26:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	if v82 == int32(0) {
		v92 = v82
		v93 = v81
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v92 = v82
	v93 = v81
	goto L24
L28:
	;
	v85 = int32(1)
	if v82 == v81 {
		v77 = v77 + v85
		v78 = v78 + v85
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v485 = int64(4)
	goto L2
L31:
	;
	goto L32
L32:
	;
	v98 = int32(_a_F_string_to_privilege_3)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[3])))
	if base.B2i32(v101 == int32(0))|base.B2i32(v101 != v104) != 0 {
		v122 = v101
		v123 = v104
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v122-v123 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	goto L33
L35:
	;
	v107 = l0
	v108 = v98
	goto L36
L36:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	if v112 == int32(0) {
		v122 = v112
		v123 = v111
		goto L34
	} else {
		goto L38
	}
L37:
	;
	v122 = v112
	v123 = v111
	goto L34
L38:
	;
	v115 = int32(1)
	if v112 == v111 {
		v107 = v107 + v115
		v108 = v108 + v115
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v485 = int64(8)
	goto L2
L41:
	;
	goto L42
L42:
	;
	v128 = int32(_a_F_string_to_privilege_4)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[4])))
	if base.B2i32(v131 == int32(0))|base.B2i32(v131 != v134) != 0 {
		v152 = v131
		v153 = v134
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v152-v153 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	goto L43
L45:
	;
	v137 = l0
	v138 = v128
	goto L46
L46:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	if v142 == int32(0) {
		v152 = v142
		v153 = v141
		goto L44
	} else {
		goto L48
	}
L47:
	;
	v152 = v142
	v153 = v141
	goto L44
L48:
	;
	v145 = int32(1)
	if v142 == v141 {
		v137 = v137 + v145
		v138 = v138 + v145
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v485 = int64(16)
	goto L2
L51:
	;
	goto L52
L52:
	;
	v158 = int32(_a_F_string_to_privilege_5)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[5])))
	if base.B2i32(v161 == int32(0))|base.B2i32(v161 != v164) != 0 {
		v182 = v161
		v183 = v164
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v182-v183 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	goto L53
L55:
	;
	v167 = l0
	v168 = v158
	goto L56
L56:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	if v172 == int32(0) {
		v182 = v172
		v183 = v171
		goto L54
	} else {
		goto L58
	}
L57:
	;
	v182 = v172
	v183 = v171
	goto L54
L58:
	;
	v175 = int32(1)
	if v172 == v171 {
		v167 = v167 + v175
		v168 = v168 + v175
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v485 = int64(32)
	goto L2
L61:
	;
	goto L62
L62:
	;
	v188 = int32(_a_F_string_to_privilege_6)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[6])))
	if base.B2i32(v191 == int32(0))|base.B2i32(v191 != v194) != 0 {
		v212 = v191
		v213 = v194
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v212-v213 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L64:
	;
	goto L63
L65:
	;
	v197 = l0
	v198 = v188
	goto L66
L66:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)))
	if v202 == int32(0) {
		v212 = v202
		v213 = v201
		goto L64
	} else {
		goto L68
	}
L67:
	;
	v212 = v202
	v213 = v201
	goto L64
L68:
	;
	v205 = int32(1)
	if v202 == v201 {
		v197 = v197 + v205
		v198 = v198 + v205
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v485 = int64(64)
	goto L2
L71:
	;
	goto L72
L72:
	;
	v218 = int32(_a_F_string_to_privilege_7)
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[7])))
	if base.B2i32(v221 == int32(0))|base.B2i32(v221 != v224) != 0 {
		v242 = v221
		v243 = v224
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v242-v243 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L74:
	;
	goto L73
L75:
	;
	v227 = l0
	v228 = v218
	goto L76
L76:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+1)))
	if v232 == int32(0) {
		v242 = v232
		v243 = v231
		goto L74
	} else {
		goto L78
	}
L77:
	;
	v242 = v232
	v243 = v231
	goto L74
L78:
	;
	v235 = int32(1)
	if v232 == v231 {
		v227 = v227 + v235
		v228 = v228 + v235
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v485 = int64(128)
	goto L2
L81:
	;
	goto L82
L82:
	;
	v248 = int32(_a_F_string_to_privilege_8)
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[8])))
	if base.B2i32(v251 == int32(0))|base.B2i32(v251 != v254) != 0 {
		v272 = v251
		v273 = v254
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v272-v273 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L84:
	;
	goto L83
L85:
	;
	v257 = l0
	v258 = v248
	goto L86
L86:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+1)))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+1)))
	if v262 == int32(0) {
		v272 = v262
		v273 = v261
		goto L84
	} else {
		goto L88
	}
L87:
	;
	v272 = v262
	v273 = v261
	goto L84
L88:
	;
	v265 = int32(1)
	if v262 == v261 {
		v257 = v257 + v265
		v258 = v258 + v265
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v485 = int64(256)
	goto L2
L91:
	;
	goto L92
L92:
	;
	v278 = int32(_a_F_string_to_privilege_9)
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[9])))
	if base.B2i32(v281 == int32(0))|base.B2i32(v281 != v284) != 0 {
		v302 = v281
		v303 = v284
		goto L94
	} else {
		goto L95
	}
L93:
	;
	if v302-v303 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L94:
	;
	goto L93
L95:
	;
	v287 = l0
	v288 = v278
	goto L96
L96:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+1)))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+1)))
	if v292 == int32(0) {
		v302 = v292
		v303 = v291
		goto L94
	} else {
		goto L98
	}
L97:
	;
	v302 = v292
	v303 = v291
	goto L94
L98:
	;
	v295 = int32(1)
	if v292 == v291 {
		v287 = v287 + v295
		v288 = v288 + v295
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v485 = int64(512)
	goto L2
L101:
	;
	goto L102
L102:
	;
	v308 = int64(1024)
	v309 = int32(_a_F_string_to_privilege_10)
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v315 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[10])))
	if base.B2i32(v312 == int32(0))|base.B2i32(v312 != v315) != 0 {
		v333 = v312
		v334 = v315
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v333-v334 == int32(0) {
		v485 = v308
		goto L2
	} else {
		goto L110
	}
L104:
	;
	goto L103
L105:
	;
	v318 = l0
	v319 = v309
	goto L106
L106:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+1)))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+1)))
	if v323 == int32(0) {
		v333 = v323
		v334 = v322
		goto L104
	} else {
		goto L108
	}
L107:
	;
	v333 = v323
	v334 = v322
	goto L104
L108:
	;
	v326 = int32(1)
	if v323 == v322 {
		v318 = v318 + v326
		v319 = v319 + v326
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v338 = int32(_a_F_string_to_privilege_11)
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[11])))
	if base.B2i32(v341 == int32(0))|base.B2i32(v341 != v344) != 0 {
		v362 = v341
		v363 = v344
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v362-v363 == int32(0) {
		v485 = v308
		goto L2
	} else {
		goto L118
	}
L112:
	;
	goto L111
L113:
	;
	v347 = l0
	v348 = v338
	goto L114
L114:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+1)))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+1)))
	if v352 == int32(0) {
		v362 = v352
		v363 = v351
		goto L112
	} else {
		goto L116
	}
L115:
	;
	v362 = v352
	v363 = v351
	goto L112
L116:
	;
	v355 = int32(1)
	if v352 == v351 {
		v347 = v347 + v355
		v348 = v348 + v355
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v367 = int32(_a_F_string_to_privilege_12)
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v373 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[12])))
	if base.B2i32(v370 == int32(0))|base.B2i32(v370 != v373) != 0 {
		v391 = v370
		v392 = v373
		goto L120
	} else {
		goto L121
	}
L119:
	;
	if v391-v392 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L120:
	;
	goto L119
L121:
	;
	v376 = l0
	v377 = v367
	goto L122
L122:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+1)))
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+1)))
	if v381 == int32(0) {
		v391 = v381
		v392 = v380
		goto L120
	} else {
		goto L124
	}
L123:
	;
	v391 = v381
	v392 = v380
	goto L120
L124:
	;
	v384 = int32(1)
	if v381 == v380 {
		v376 = v376 + v384
		v377 = v377 + v384
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v485 = int64(2048)
	goto L2
L127:
	;
	goto L128
L128:
	;
	v397 = int32(_a_F_string_to_privilege_13)
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v403 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[13])))
	if base.B2i32(v400 == int32(0))|base.B2i32(v400 != v403) != 0 {
		v421 = v400
		v422 = v403
		goto L130
	} else {
		goto L131
	}
L129:
	;
	if v421-v422 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L130:
	;
	goto L129
L131:
	;
	v406 = l0
	v407 = v397
	goto L132
L132:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+1)))
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+1)))
	if v411 == int32(0) {
		v421 = v411
		v422 = v410
		goto L130
	} else {
		goto L134
	}
L133:
	;
	v421 = v411
	v422 = v410
	goto L130
L134:
	;
	v414 = int32(1)
	if v411 == v410 {
		v406 = v406 + v414
		v407 = v407 + v414
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v485 = int64(4096)
	goto L2
L137:
	;
	goto L138
L138:
	;
	v427 = int32(_a_F_string_to_privilege_14)
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[14])))
	if base.B2i32(v430 == int32(0))|base.B2i32(v430 != v433) != 0 {
		v451 = v430
		v452 = v433
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v451-v452 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L140:
	;
	goto L139
L141:
	;
	v436 = l0
	v437 = v427
	goto L142
L142:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437)+1)))
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436)+1)))
	if v441 == int32(0) {
		v451 = v441
		v452 = v440
		goto L140
	} else {
		goto L144
	}
L143:
	;
	v451 = v441
	v452 = v440
	goto L140
L144:
	;
	v444 = int32(1)
	if v441 == v440 {
		v436 = v436 + v444
		v437 = v437 + v444
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v485 = int64(8192)
	goto L2
L147:
	;
	goto L148
L148:
	;
	v457 = int32(_a_F_string_to_privilege_15)
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v463 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string_to_privilege[15])))
	if base.B2i32(v460 == int32(0))|base.B2i32(v460 != v463) != 0 {
		v481 = v460
		v482 = v463
		goto L150
	} else {
		goto L151
	}
L149:
	;
	if v481-v482 != 0 {
		goto L1
	} else {
		goto L156
	}
L150:
	;
	goto L149
L151:
	;
	v466 = l0
	v467 = v457
	goto L152
L152:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467)+1)))
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466)+1)))
	if v471 == int32(0) {
		v481 = v471
		v482 = v470
		goto L150
	} else {
		goto L154
	}
L153:
	;
	v481 = v471
	v482 = v470
	goto L150
L154:
	;
	v474 = int32(1)
	if v471 == v470 {
		v466 = v466 + v474
		v467 = v467 + v474
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v485 = int64(16384)
	goto L2
L157:
	;
	return int64(0)
L158:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg(m, int32(_a_F_string_to_privilege_16), v6)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L157
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_string_to_privilege_17), int32(2603), int32(_a_F_string_to_privilege_18))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L157
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_transform_string_values_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v9 = v7 + v8
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9-int32(1)))))
	if v12 == int32(123) {
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_escape_json(m, v37, l1)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
			if v44 <= v41+int32(1) {
				F_appendStringInfoChar(m, v40, int32(58))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v53 = int32(58)
				*(*uint8)(unsafe.Add(mBase, uint32(v51+v41))) = uint8(v53)
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
				v58 = v56 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v58
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
				v62 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v60+v58))) = uint8(v62)
				return v62
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		if v15 <= v8+int32(1) {
			F_appendStringInfoChar(m, v6, int32(44))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_escape_json(m, v37, l1)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
					if v44 <= v41+int32(1) {
						F_appendStringInfoChar(m, v40, int32(58))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						v53 = int32(58)
						*(*uint8)(unsafe.Add(mBase, uint32(v51+v41))) = uint8(v53)
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
						v58 = v56 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v58
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
						v62 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v60+v58))) = uint8(v62)
						return v62
					}
				}
			}
		} else {
			v24 = int32(44)
			*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v24)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
			v29 = v27 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v29
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v33 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v31+v29))) = uint8(v33)
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			F_escape_json(m, v37, l1)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
				if v44 <= v41+int32(1) {
					F_appendStringInfoChar(m, v40, int32(58))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					v53 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v51+v41))) = uint8(v53)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
					v58 = v56 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v58
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
					v62 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v60+v58))) = uint8(v62)
					return v62
				}
			}
		}
	}
}
