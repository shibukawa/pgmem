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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v28 = v6
	v30 = int32(-1)
	v31 = v18
	v32 = v6
	v33 = v6
	v34 = v6
	v35 = v6
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
	if v30 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L2
L5:
	;
	v172 = int32(m.ExcTag)
	v173 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v172 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v66
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v67
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_bms_free(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L45
	}
L7:
	;
	m.G0 = v18 - int32(-64)
	return v154 & int32(1)
L8:
	;
	v41 = v31 - int32(16)
	m.G0 = v41
	v44 = v41 - int32(160)
	m.G0 = v44
	v46 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v46)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v49 == int32(0) {
		v154 = v46
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v62 = v28
	v63 = v31
	v64 = v32
	v65 = v33
	v66 = v34
	v67 = v35
	goto L10
L10:
	;
	if v62 != 0 {
		goto L6
	} else {
		goto L16
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	v56 = *(*int32)(unsafe.Add(mBase, _consts[362]))
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v16 + int32(-4)
	goto L15
L13:
	;
	v62 = int32(0)
	v63 = v44
	v64 = v41
	v65 = v44
	v66 = v54
	v67 = v56
	goto L10
L15:
	;
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[427])) = int32(50856066)
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v65
	v74 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[426])) = v74
	*(*int32)(unsafe.Add(mBase, _consts[825])) = v74
	*(*int32)(unsafe.Add(mBase, _consts[1207])) = v74
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(100))))
	v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l3)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v83 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v88 = F_errstart(m, l4, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v66
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v67
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v154 = v146
	goto L7
L21:
	;
	if v88 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[427]))
	F_errcode(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
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
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L44
	}
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[426]))
	if v95 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _consts[825]))
	if v116 != 0 {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v95
	F_errmsg_internal(m, int32(193943), v16+int32(-16))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v103 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L26
L31:
	;
	v105 = v103
	goto L33
L32:
	;
	v105 = int32(706478)
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v102
	F_errmsg(m, int32(677148), v16+int32(-32))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	goto L26
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v116
	F_errdetail_internal(m, int32(193943), v16+int32(-48))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[1207]))
	if v124 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v124
	F_errhint(m, int32(193943), v18)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_errfinish(m, int32(469097), int32(6947), int32(296029))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
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
	v138 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v138)
	goto L20
L45:
	;
	F_pg_re_throw(m)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	goto L4
L47:
	;
	v177 = int32(v173)
	m.G0 = v63
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	if v16+int32(-4) == v184 {
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
	if v187 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	v187 = v186
	goto L52
L51:
	;
	v187 = int32(0)
	goto L52
L52:
	;
	goto L49
L53:
	;
	F___wasm_longjmp(m, v180, v179)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v28 = v179
	v30 = v187
	v31 = v63
	v32 = v64
	v33 = v65
	v34 = v66
	v35 = v67
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
	F_errmsg(m, int32(26921), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errsave_finish(m, l1, int32(469187), int32(1810), int32(72057))
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
	F_errmsg(m, int32(26921), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	F_errsave_finish(m, l1, int32(469187), int32(1815), int32(72057))
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
			v43 = v16
			v44 = int32(950)
			v45 = int32(64)
			v47 = int32(0)
			v49 = F_makeConst(m, l1, int32(-1), v44, v45, v43, v47, v47)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return v49
			}
		}
	} else {
		v20 = int32(100)
		v21 = int32(-1)
		v22 = F_cstring_to_text(m, l0)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if base.Ui32(l1-int32(1042)) < base.Ui32(int32(2)) {
				v43 = v22
				v44 = v20
				v45 = v21
				v47 = int32(0)
				v49 = F_makeConst(m, l1, int32(-1), v44, v45, v43, v47, v47)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return v49
				}
			} else {
				if l1 == int32(25) {
					v43 = v22
					v44 = v20
					v45 = v21
					v47 = int32(0)
					v49 = F_makeConst(m, l1, int32(-1), v44, v45, v43, v47, v47)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v49
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
						F_errmsg_internal(m, int32(54279), v8)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(462616), int32(1765), int32(64476))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
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
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
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
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int64
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int64
	_ = v469
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(76806)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[279])))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v12 == int32(0) {
		v31 = v11
		v32 = v12
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L173
	} else {
		goto L174
	}
L2:
	;
	m.G0 = v6 + int32(16)
	return v469
L3:
	;
	if v32-v31 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	goto L3
L5:
	;
	if v11 != v12 {
		v31 = v11
		v32 = v12
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v16 = l0
	v17 = v8
	goto L7
L7:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v21 == int32(0) {
		v31 = v20
		v32 = v21
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v31 = v20
	v32 = v21
	goto L4
L9:
	;
	v24 = int32(1)
	if v20 == v21 {
		v16 = v16 + v24
		v17 = v17 + v24
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v469 = int64(1)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v37 = int32(101306)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, _consts[280])))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v41 == int32(0) {
		v60 = v40
		v61 = v41
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v61-v60 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	goto L14
L16:
	;
	if v40 != v41 {
		v60 = v40
		v61 = v41
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v45 = l0
	v46 = v37
	goto L18
L18:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+1)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
	if v50 == int32(0) {
		v60 = v49
		v61 = v50
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v60 = v49
	v61 = v50
	goto L15
L20:
	;
	v53 = int32(1)
	if v49 == v50 {
		v45 = v45 + v53
		v46 = v46 + v53
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v469 = int64(2)
	goto L2
L23:
	;
	goto L24
L24:
	;
	v66 = int32(334622)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _consts[281])))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v70 == int32(0) {
		v89 = v69
		v90 = v70
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v90-v89 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	goto L25
L27:
	;
	if v69 != v70 {
		v89 = v69
		v90 = v70
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v74 = l0
	v75 = v66
	goto L29
L29:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	if v79 == int32(0) {
		v89 = v78
		v90 = v79
		goto L26
	} else {
		goto L31
	}
L30:
	;
	v89 = v78
	v90 = v79
	goto L26
L31:
	;
	v82 = int32(1)
	if v78 == v79 {
		v74 = v74 + v82
		v75 = v75 + v82
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v469 = int64(4)
	goto L2
L34:
	;
	goto L35
L35:
	;
	v95 = int32(329756)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _consts[282])))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v99 == int32(0) {
		v118 = v98
		v119 = v99
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v119-v118 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L37:
	;
	goto L36
L38:
	;
	if v98 != v99 {
		v118 = v98
		v119 = v99
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v103 = l0
	v104 = v95
	goto L40
L40:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	if v108 == int32(0) {
		v118 = v107
		v119 = v108
		goto L37
	} else {
		goto L42
	}
L41:
	;
	v118 = v107
	v119 = v108
	goto L37
L42:
	;
	v111 = int32(1)
	if v107 == v108 {
		v103 = v103 + v111
		v104 = v104 + v111
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v469 = int64(8)
	goto L2
L45:
	;
	goto L46
L46:
	;
	v124 = int32(335560)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, _consts[283])))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v128 == int32(0) {
		v147 = v127
		v148 = v128
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v148-v147 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L48:
	;
	goto L47
L49:
	;
	if v127 != v128 {
		v147 = v127
		v148 = v128
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v132 = l0
	v133 = v124
	goto L51
L51:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v137 == int32(0) {
		v147 = v136
		v148 = v137
		goto L48
	} else {
		goto L53
	}
L52:
	;
	v147 = v136
	v148 = v137
	goto L48
L53:
	;
	v140 = int32(1)
	if v136 == v137 {
		v132 = v132 + v140
		v133 = v133 + v140
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v469 = int64(16)
	goto L2
L56:
	;
	goto L57
L57:
	;
	v153 = int32(159986)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, _consts[284])))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v157 == int32(0) {
		v176 = v156
		v177 = v157
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v177-v176 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L59:
	;
	goto L58
L60:
	;
	if v156 != v157 {
		v176 = v156
		v177 = v157
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v161 = l0
	v162 = v153
	goto L62
L62:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	if v166 == int32(0) {
		v176 = v165
		v177 = v166
		goto L59
	} else {
		goto L64
	}
L63:
	;
	v176 = v165
	v177 = v166
	goto L59
L64:
	;
	v169 = int32(1)
	if v165 == v166 {
		v161 = v161 + v169
		v162 = v162 + v169
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v469 = int64(32)
	goto L2
L67:
	;
	goto L68
L68:
	;
	v182 = int32(211368)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, _consts[285])))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v186 == int32(0) {
		v205 = v185
		v206 = v186
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v206-v205 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L70:
	;
	goto L69
L71:
	;
	if v185 != v186 {
		v205 = v185
		v206 = v186
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v190 = l0
	v191 = v182
	goto L73
L73:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
	if v195 == int32(0) {
		v205 = v194
		v206 = v195
		goto L70
	} else {
		goto L75
	}
L74:
	;
	v205 = v194
	v206 = v195
	goto L70
L75:
	;
	v198 = int32(1)
	if v194 == v195 {
		v190 = v190 + v198
		v191 = v191 + v198
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v469 = int64(64)
	goto L2
L78:
	;
	goto L79
L79:
	;
	v211 = int32(327275)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, _consts[286])))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v215 == int32(0) {
		v234 = v214
		v235 = v215
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v235-v234 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L81:
	;
	goto L80
L82:
	;
	if v214 != v215 {
		v234 = v214
		v235 = v215
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v219 = l0
	v220 = v211
	goto L84
L84:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+1)))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)))
	if v224 == int32(0) {
		v234 = v223
		v235 = v224
		goto L81
	} else {
		goto L86
	}
L85:
	;
	v234 = v223
	v235 = v224
	goto L81
L86:
	;
	v227 = int32(1)
	if v223 == v224 {
		v219 = v219 + v227
		v220 = v220 + v227
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v469 = int64(128)
	goto L2
L89:
	;
	goto L90
L90:
	;
	v240 = int32(379890)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, _consts[287])))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v244 == int32(0) {
		v263 = v243
		v264 = v244
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v264-v263 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L92:
	;
	goto L91
L93:
	;
	if v243 != v244 {
		v263 = v243
		v264 = v244
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v248 = l0
	v249 = v240
	goto L95
L95:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
	if v253 == int32(0) {
		v263 = v252
		v264 = v253
		goto L92
	} else {
		goto L97
	}
L96:
	;
	v263 = v252
	v264 = v253
	goto L92
L97:
	;
	v256 = int32(1)
	if v252 == v253 {
		v248 = v248 + v256
		v249 = v249 + v256
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v469 = int64(256)
	goto L2
L100:
	;
	goto L101
L101:
	;
	v269 = int32(333692)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, _consts[288])))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v273 == int32(0) {
		v292 = v272
		v293 = v273
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v293-v292 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L103:
	;
	goto L102
L104:
	;
	if v272 != v273 {
		v292 = v272
		v293 = v273
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v277 = l0
	v278 = v269
	goto L106
L106:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+1)))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+1)))
	if v282 == int32(0) {
		v292 = v281
		v293 = v282
		goto L103
	} else {
		goto L108
	}
L107:
	;
	v292 = v281
	v293 = v282
	goto L103
L108:
	;
	v285 = int32(1)
	if v281 == v282 {
		v277 = v277 + v285
		v278 = v278 + v285
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v469 = int64(512)
	goto L2
L111:
	;
	goto L112
L112:
	;
	v298 = int64(1024)
	v299 = int32(16218)
	v302 = int32(*(*uint8)(unsafe.Add(mBase, _consts[289])))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v303 == int32(0) {
		v322 = v302
		v323 = v303
		goto L114
	} else {
		goto L115
	}
L113:
	;
	if v323-v322 == int32(0) {
		v469 = v298
		goto L2
	} else {
		goto L121
	}
L114:
	;
	goto L113
L115:
	;
	if v302 != v303 {
		v322 = v302
		v323 = v303
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v307 = l0
	v308 = v299
	goto L117
L117:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+1)))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+1)))
	if v312 == int32(0) {
		v322 = v311
		v323 = v312
		goto L114
	} else {
		goto L119
	}
L118:
	;
	v322 = v311
	v323 = v312
	goto L114
L119:
	;
	v315 = int32(1)
	if v311 == v312 {
		v307 = v307 + v315
		v308 = v308 + v315
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v327 = int32(221158)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, _consts[290])))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v331 == int32(0) {
		v350 = v330
		v351 = v331
		goto L123
	} else {
		goto L124
	}
L122:
	;
	if v351-v350 == int32(0) {
		v469 = v298
		goto L2
	} else {
		goto L130
	}
L123:
	;
	goto L122
L124:
	;
	if v330 != v331 {
		v350 = v330
		v351 = v331
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v335 = l0
	v336 = v327
	goto L126
L126:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+1)))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+1)))
	if v340 == int32(0) {
		v350 = v339
		v351 = v340
		goto L123
	} else {
		goto L128
	}
L127:
	;
	v350 = v339
	v351 = v340
	goto L123
L128:
	;
	v343 = int32(1)
	if v339 == v340 {
		v335 = v335 + v343
		v336 = v336 + v343
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	v355 = int32(101182)
	v358 = int32(*(*uint8)(unsafe.Add(mBase, _consts[291])))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v359 == int32(0) {
		v378 = v358
		v379 = v359
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v379-v378 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L132:
	;
	goto L131
L133:
	;
	if v358 != v359 {
		v378 = v358
		v379 = v359
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v363 = l0
	v364 = v355
	goto L135
L135:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+1)))
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363)+1)))
	if v368 == int32(0) {
		v378 = v367
		v379 = v368
		goto L132
	} else {
		goto L137
	}
L136:
	;
	v378 = v367
	v379 = v368
	goto L132
L137:
	;
	v371 = int32(1)
	if v367 == v368 {
		v363 = v363 + v371
		v364 = v364 + v371
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	v469 = int64(2048)
	goto L2
L140:
	;
	goto L141
L141:
	;
	v384 = int32(98717)
	v387 = int32(*(*uint8)(unsafe.Add(mBase, _consts[292])))
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v388 == int32(0) {
		v407 = v387
		v408 = v388
		goto L143
	} else {
		goto L144
	}
L142:
	;
	if v408-v407 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L143:
	;
	goto L142
L144:
	;
	if v387 != v388 {
		v407 = v387
		v408 = v388
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v392 = l0
	v393 = v384
	goto L146
L146:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+1)))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+1)))
	if v397 == int32(0) {
		v407 = v396
		v408 = v397
		goto L143
	} else {
		goto L148
	}
L147:
	;
	v407 = v396
	v408 = v397
	goto L143
L148:
	;
	v400 = int32(1)
	if v396 == v397 {
		v392 = v392 + v400
		v393 = v393 + v400
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	v469 = int64(4096)
	goto L2
L151:
	;
	goto L152
L152:
	;
	v413 = int32(271338)
	v416 = int32(*(*uint8)(unsafe.Add(mBase, _consts[293])))
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v417 == int32(0) {
		v436 = v416
		v437 = v417
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if v437-v436 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L154:
	;
	goto L153
L155:
	;
	if v416 != v417 {
		v436 = v416
		v437 = v417
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v421 = l0
	v422 = v413
	goto L157
L157:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422)+1)))
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421)+1)))
	if v426 == int32(0) {
		v436 = v425
		v437 = v426
		goto L154
	} else {
		goto L159
	}
L158:
	;
	v436 = v425
	v437 = v426
	goto L154
L159:
	;
	v429 = int32(1)
	if v425 == v426 {
		v421 = v421 + v429
		v422 = v422 + v429
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v469 = int64(8192)
	goto L2
L162:
	;
	goto L163
L163:
	;
	v442 = int32(260926)
	v445 = int32(*(*uint8)(unsafe.Add(mBase, _consts[294])))
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v446 == int32(0) {
		v465 = v445
		v466 = v446
		goto L165
	} else {
		goto L166
	}
L164:
	;
	if v466-v465 != 0 {
		goto L1
	} else {
		goto L172
	}
L165:
	;
	goto L164
L166:
	;
	if v445 != v446 {
		v465 = v445
		v466 = v446
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v450 = l0
	v451 = v442
	goto L168
L168:
	;
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+1)))
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+1)))
	if v455 == int32(0) {
		v465 = v454
		v466 = v455
		goto L165
	} else {
		goto L170
	}
L169:
	;
	v465 = v454
	v466 = v455
	goto L165
L170:
	;
	v458 = int32(1)
	if v454 == v455 {
		v450 = v450 + v458
		v451 = v451 + v458
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	v469 = int64(16384)
	goto L2
L173:
	;
	return int64(0)
L174:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L173
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg(m, int32(663930), v6)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L173
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(467048), int32(2603), int32(379296))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L173
	} else {
		goto L177
	}
L177:
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
