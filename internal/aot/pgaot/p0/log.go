package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LogCheckpointStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v10 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if v10 != 0 {
			if l0&int32(16) != 0 {
				v16 = int32(300138)
			} else {
				v16 = int32(731167)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v16
			if l0&int32(256) != 0 {
				v22 = int32(369348)
			} else {
				v22 = int32(731167)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v22
			if l0&int32(128) != 0 {
				v28 = int32(303343)
			} else {
				v28 = int32(731167)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v28
			if l0&int32(32) != 0 {
				v34 = int32(101766)
			} else {
				v34 = int32(731167)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v34
			if l0&int32(8) != 0 {
				v40 = int32(407356)
			} else {
				v40 = int32(731167)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v40
			if l0&int32(4) != 0 {
				v46 = int32(348516)
			} else {
				v46 = int32(731167)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v46
			if l0&int32(2) != 0 {
				v52 = int32(14089)
			} else {
				v52 = int32(731167)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v52
			if l0&int32(1) != 0 {
				v58 = int32(239864)
			} else {
				v58 = int32(731167)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v58
			if l1 != 0 {
				v62 = int32(172501)
			} else {
				v62 = int32(172540)
			}
			F_errmsg(m, v62, v6)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				if l1 != 0 {
					v68 = int32(6702)
				} else {
					v68 = int32(6714)
				}
				F_errfinish(m, int32(489152), v68, int32(81754))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					m.G0 = v6 + int32(32)
					return
				}
			}
		} else {
			m.G0 = v6 + int32(32)
			return
		}
	}
}
func F_ProcessLogMemoryContextInterrupt(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	v1 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = int32(-1)
	v15 = v1
	v17 = v11
	v19 = v1
	v20 = v1
	v21 = v1
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
	if v14 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L2
L5:
	;
	v151 = int32(m.ExcTag)
	v152 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v151 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v49
	*(*int32)(unsafe.Add(mBase, _consts[294])) = v48
	v144 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[982])) = uint8(v144)
	F_pg_re_throw(m)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L34
	}
L7:
	;
	m.G0 = v11 + int32(16)
	return
L8:
	;
	v25 = v17 - int32(160)
	m.G0 = v25
	*(*int32)(unsafe.Add(mBase, _consts[494])) = int32(0)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[982])))
	if v31 != 0 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v45 = v15
	v46 = v17
	v47 = v19
	v48 = v20
	v49 = v21
	goto L10
L10:
	;
	if v45 != 0 {
		goto L6
	} else {
		goto L16
	}
L11:
	;
	v34 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[982])) = uint8(v34)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[294]))
	v39 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v11 + int32(12)
	goto L15
L13:
	;
	v45 = int32(0)
	v46 = v25
	v47 = v25
	v48 = v37
	v49 = v39
	goto L10
L15:
	;
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[294])) = v47
	v54 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v54 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_errhidestmt(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v73 = m.G0
	v75 = v73 - int32(80)
	m.G0 = v75
	v77 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v75)+72)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v75)+64)) = v77
	v82 = int32(100)
	F_MemoryContextStatsInternal(m, v72, int32(1), v82, v82, v75-int32(-64), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L25
	}
L21:
	;
	F_errhidecontext(m)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v61
	F_errmsg(m, int32(471274), v11)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(483770), int32(1319), int32(82461))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v91 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	if v91 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_errhidestmt(m)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	m.G0 = v75 + int32(80)
	*(*int32)(unsafe.Add(mBase, _consts[294])) = v48
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v49
	v127 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[982])) = uint8(v127)
	goto L7
L30:
	;
	F_errhidecontext(m)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v75)+72))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v75)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+48)) = v97 - v98
	*(*int32)(unsafe.Add(mBase, uint32(v75)+32)) = v97
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v75)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+36)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v75)+40)) = v98
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v75)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+44)) = v105
	F_errmsg_internal(m, int32(441476), v75+int32(32))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(483770), int32(867), int32(300620))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	goto L4
L35:
	;
	v156 = int32(v152)
	m.G0 = v46
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v11+int32(12) == v163 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	m.ExcPending = 1
	goto L44
L37:
	;
	if v166 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v166 = v165
	goto L40
L39:
	;
	v166 = int32(0)
	goto L40
L40:
	;
	goto L37
L41:
	;
	F___wasm_longjmp(m, v159, v158)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v14 = v166
	v15 = v158
	v17 = v46
	v19 = v47
	v20 = v48
	v21 = v49
	goto L1
L44:
	;
	return
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_log_destination(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = F_pstrdup(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v284
L2:
	;
	return int32(0)
L3:
	;
	v20 = F_SplitIdentifierString(m, v13, int32(44), v10+int32(12))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v20 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	*(*int32)(unsafe.Add(mBase, _consts[189])) = v26
	v31 = F_format_elog_string(m, int32(620576), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v39 = int32(0)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v40 == v39 {
		v245 = v39
		goto L12
	} else {
		goto L13
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[555])) = v31
	F_pfree(m, v13)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_list_free(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v284 = v4
	goto L1
L11:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	*(*int32)(unsafe.Add(mBase, _consts[189])) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v59
	v270 = F_format_elog_string(m, int32(642116), v10)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L2
	} else {
		goto L79
	}
L12:
	;
	F_pfree(m, v13)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L2
	} else {
		goto L75
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v43 <= int32(0) {
		v245 = v39
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v47 = int32(0)
	v49 = v39
	goto L15
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v47<<(uint(int32(2))%32))))
	v63 = v59
	v64 = int32(203568)
	goto L19
L16:
	;
	v245 = v238
	goto L12
L17:
	;
	v238 = v49 | v237
	v240 = v47 + int32(1)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v240 < v241 {
		v47 = v240
		v49 = v238
		goto L15
	} else {
		goto L74
	}
L18:
	;
	if v101 == int32(0) {
		v237 = int32(1)
		goto L17
	} else {
		goto L31
	}
L19:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v67 == v68 {
		v90 = v67
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v101 = int32(0)
	goto L18
L21:
	;
	v92 = int32(1)
	if v90 != 0 {
		v63 = v63 + v92
		v64 = v64 + v92
		goto L19
	} else {
		goto L30
	}
L22:
	;
	if base.Ui32((v67-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v78 = v67 | int32(32)
	goto L25
L24:
	;
	v78 = v67
	goto L25
L25:
	;
	if base.Ui32((v68-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v87 = v68 | int32(32)
	goto L28
L27:
	;
	v87 = v68
	goto L28
L28:
	;
	if v78 == v87 {
		v90 = v78
		goto L21
	} else {
		goto L29
	}
L29:
	;
	v101 = v78 - v87
	goto L18
L30:
	;
	goto L20
L31:
	;
	v108 = v59
	v109 = int32(321460)
	goto L33
L32:
	;
	if v146 == int32(0) {
		v237 = int32(8)
		goto L17
	} else {
		goto L45
	}
L33:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v112 == v113 {
		v135 = v112
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v146 = int32(0)
	goto L32
L35:
	;
	v137 = int32(1)
	if v135 != 0 {
		v108 = v108 + v137
		v109 = v109 + v137
		goto L33
	} else {
		goto L44
	}
L36:
	;
	if base.Ui32((v112-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v123 = v112 | int32(32)
	goto L39
L38:
	;
	v123 = v112
	goto L39
L39:
	;
	if base.Ui32((v113-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v132 = v113 | int32(32)
	goto L42
L41:
	;
	v132 = v113
	goto L42
L42:
	;
	if v123 == v132 {
		v135 = v123
		goto L35
	} else {
		goto L43
	}
L43:
	;
	v146 = v123 - v132
	goto L32
L44:
	;
	goto L34
L45:
	;
	v153 = v59
	v154 = int32(321474)
	goto L47
L46:
	;
	if v191 == int32(0) {
		v237 = int32(16)
		goto L17
	} else {
		goto L59
	}
L47:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v157 == v158 {
		v180 = v157
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v191 = int32(0)
	goto L46
L49:
	;
	v182 = int32(1)
	if v180 != 0 {
		v153 = v153 + v182
		v154 = v154 + v182
		goto L47
	} else {
		goto L58
	}
L50:
	;
	if base.Ui32((v157-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v168 = v157 | int32(32)
	goto L53
L52:
	;
	v168 = v157
	goto L53
L53:
	;
	if base.Ui32((v158-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v177 = v158 | int32(32)
	goto L56
L55:
	;
	v177 = v158
	goto L56
L56:
	;
	if v168 == v177 {
		v180 = v168
		goto L49
	} else {
		goto L57
	}
L57:
	;
	v191 = v168 - v177
	goto L46
L58:
	;
	goto L48
L59:
	;
	v197 = v59
	v198 = int32(321467)
	goto L61
L60:
	;
	if v235 != 0 {
		goto L11
	} else {
		goto L73
	}
L61:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	if v201 == v202 {
		v224 = v201
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v235 = int32(0)
	goto L60
L63:
	;
	v226 = int32(1)
	if v224 != 0 {
		v197 = v197 + v226
		v198 = v198 + v226
		goto L61
	} else {
		goto L72
	}
L64:
	;
	if base.Ui32((v201-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v212 = v201 | int32(32)
	goto L67
L66:
	;
	v212 = v201
	goto L67
L67:
	;
	if base.Ui32((v202-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v221 = v202 | int32(32)
	goto L70
L69:
	;
	v221 = v202
	goto L70
L70:
	;
	if v212 == v221 {
		v224 = v212
		goto L63
	} else {
		goto L71
	}
L71:
	;
	v235 = v212 - v221
	goto L60
L72:
	;
	goto L62
L73:
	;
	v237 = int32(2)
	goto L17
L74:
	;
	goto L16
L75:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_list_free(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v256 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	if v256 == int32(0) {
		v284 = v4
		goto L1
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v245
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v256
	v284 = int32(1)
	goto L1
L79:
	;
	*(*int32)(unsafe.Add(mBase, _consts[555])) = v270
	F_pfree(m, v13)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_list_free(m, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v284 = v4
	goto L1
}
func F_check_log_duration(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int64
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 float64
	_ = v110
	var v115 int32
	_ = v115
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v141 float64
	_ = v141
	var v143 float64
	_ = v143
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[661])))
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v209
L2:
	;
	v30 = *(*int64)(unsafe.Add(mBase, _consts[95]))
	v34 = m.G0
	v35 = int32(16)
	v36 = v34 - v35
	m.G0 = v36
	F___gettimeofday(m, v36)
	mBase = m.M
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
	v40 = int64(*(*int32)(unsafe.Add(mBase, uint32(v36)+8)))
	m.G0 = v36 + v35
	goto L7
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[663]))
	if int32(0) <= v18 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[662]))
	if int32(0) <= v22 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, _consts[665])))
	if v26 != int32(1) {
		v209 = int32(0)
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L2
L7:
	;
	v55 = v40 + v39*int64(1000000) - int64(946684800000000) - v30
	if v55 <= int64(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v73 = base.I32_div_s(v71, int32(1000))
	v75 = *(*int32)(unsafe.Add(mBase, _consts[662]))
	if v75 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(28)))) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(24)))) = v68
	goto L8
L10:
	;
	v67 = int32(0)
	v68 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v59 = int64(1000000)
	v60 = base.I64_div_u_s(v55, v59)
	v67 = base.I32_wrap_i64(v60)
	v68 = base.I32_wrap_i64(v55 - v60*v59)
	goto L9
L13:
	;
	v93 = int32(0)
	v95 = *(*int32)(unsafe.Add(mBase, _consts[663]))
	if v95 == v93 {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v92 = int32(1)
	goto L13
L15:
	;
	goto L16
L16:
	;
	if v75 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v92 = int32(0)
	goto L13
L18:
	;
	goto L19
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v85 = base.I32_div_u_s(v75, int32(1000))
	if v85 < v83 {
		v92 = int32(1)
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v92 = base.B2i32(v75 <= v83*int32(1000)+v73)
	goto L13
L21:
	;
	v209 = int32(1)
	goto L1
L22:
	;
	v173 = int32(0)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, _consts[661])))
	if v175 == v173 {
		goto L37
	} else {
		goto L38
	}
L23:
	;
	v158 = int32(1000)
	v159 = base.I32_rem_s(v155, v158)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v161*v158 + v73
	v168 = F_pg_snprintf(m, l0, int32(32), int32(457538), v13)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L34
	} else {
		goto L35
	}
L24:
	;
	if (v145|v92)&int32(1) == int32(0) {
		goto L22
	} else {
		goto L33
	}
L25:
	;
	v110 = *(*float64)(unsafe.Add(mBase, _consts[664]))
	if base.F64_eq(v110, float64(0)) != 0 {
		v145 = v93
		goto L24
	} else {
		goto L30
	}
L26:
	;
	if v95 <= int32(0) {
		v145 = v93
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v102 = base.I32_div_u_s(v95, int32(1000))
	if v102 < v100 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	if v100*int32(1000)+v73 < v95 {
		v145 = v93
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	if base.F64_eq(v110, float64(1)) != 0 {
		v155 = v71
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v115 = int32(4553904)
	v118 = *(*int64)(unsafe.Add(mBase, _consts[23]))
	v119 = *(*int64)(unsafe.Add(mBase, _consts[24]))
	v120 = v118 ^ v119
	*(*int64)(unsafe.Add(mBase, _consts[24])) = base.I64_rotl(v120, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[23])) = v120<<(uint(int64(16))%64) ^ base.I64_rotl(v118, int64(24)) ^ v120
	v141 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v118*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L32
L32:
	;
	v143 = *(*float64)(unsafe.Add(mBase, _consts[664]))
	v145 = base.F64_le(v141, v143)
	goto L24
L33:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v155 = v153
	goto L23
L34:
	;
	return int32(0)
L35:
	;
	if l1 != 0 {
		goto L21
	} else {
		goto L36
	}
L36:
	;
	v209 = int32(2)
	goto L1
L37:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, _consts[665])))
	if v179 != int32(1) {
		v209 = v173
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v183 = int32(1000)
	v184 = base.I32_rem_s(v182, v183)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v184
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v186*v183 + v73
	v195 = F_pg_snprintf(m, l0, int32(32), int32(457538), v13+int32(16))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L34
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	if l1 != 0 {
		goto L21
	} else {
		goto L42
	}
L42:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, _consts[665])))
	if v198&int32(1) == int32(0) {
		goto L21
	} else {
		goto L43
	}
L43:
	;
	v209 = int32(2)
	goto L1
}
func F_check_log_of_query(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[923]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v7-int32(15)) <= base.Ui32(int32(1)) {
		if v6 < int32(22) {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
			if v21 != 0 {
				v26 = v2
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, _consts[46]))
				v26 = base.B2i32(v23 != int32(0))
			}
		} else {
			v26 = v2
		}
	} else {
		if v7 == int32(20) {
			v26 = v2
		} else {
			if v6 == int32(15) {
				if int32(21) < v7 {
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
					if v21 != 0 {
						v26 = v2
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, _consts[46]))
						v26 = base.B2i32(v23 != int32(0))
					}
				} else {
					v26 = v2
				}
			} else {
				if v7 < v6 {
					v26 = v2
				} else {
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
					if v21 != 0 {
						v26 = v2
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, _consts[46]))
						v26 = base.B2i32(v23 != int32(0))
					}
				}
			}
		}
	}
	return v26
}
func F_show_log_timezone(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, _consts[321]))
	if v3 != 0 {
		v5 = v3
	} else {
		v5 = int32(239602)
	}
	return v5
}
