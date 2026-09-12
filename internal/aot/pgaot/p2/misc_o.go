package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OpenTransientFilePerm(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_reserveAllocatedDesc(m)
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
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[0]))
	if v15 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L19
	}
L6:
	;
	v55 = F_BasicOpenFilePerm(m, l0, l1, l2)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L14
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[1]))
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[2]))
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[3]))
	if v21+(v23+v15) < v19 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[4]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	F_LruDelete(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L6
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[0]))
	if v38 <= int32(0) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[1]))
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[2]))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[3]))
	if v42 <= v44+(v46+v38) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	if int32(0) <= v55 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[5]))
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[3]))
	v65 = v60 + v62*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = int32(3)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[6]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	goto L18
L16:
	;
	goto L17
L17:
	;
	m.G0 = v8 + int32(16)
	return v55
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v71
	v73 = int32(_a_F_OpenTransientFilePerm_0)
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[3])) = v75 + int32(1)
	goto L17
L19:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v93
	F_errmsg(m, int32(_a_F_OpenTransientFilePerm_1), v8)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_OpenTransientFilePerm_2), int32(2720), int32(_a_F_OpenTransientFilePerm_3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_oidge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(base.Ui32(v3) <= base.Ui32(v2))
}
func F_oidout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_palloc(m, int32(12))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
		v17 = F_pg_snprintf(m, v10, int32(12), int32(_a_F_oidout_0), v6)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v10
		}
	}
}
func F_oidsmaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.Ui32(v3) < base.Ui32(v4) {
		v6 = v3
	} else {
		v6 = v4
	}
	return v6
}
func F_oidvectorne(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_btoidvectorcmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v2 != int32(0))
	}
}
func F_oidvectortypes(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
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
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_check_valid_oidvector(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v19 = v17 * int32(20)
	v21 = v19 | int32(1)
	v22 = F_palloc(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v24)
	if v17 <= v24 {
		v268 = v22
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v279 = F_cstring_to_text(m, v268)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L73
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v31 = F_format_type_extended(m, v28, int32(-1), int32(2))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v31&int32(3) == int32(0) {
		v56 = v31
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v91 = v89 + int32(2)
	if base.Ui32(v19) < base.Ui32(v91) {
		goto L24
	} else {
		goto L25
	}
L8:
	;
	v89 = v81 - v31
	goto L7
L9:
	;
	v60 = v56
	goto L18
L10:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v40 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v89 = int32(0)
	goto L7
L12:
	;
	goto L13
L13:
	;
	v45 = v31
	goto L14
L14:
	;
	v49 = v45 + int32(1)
	if v49&int32(3) == int32(0) {
		v56 = v49
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v81 = v49
	goto L8
L16:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v54 != 0 {
		v45 = v49
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v69 = int32(-2139062144)
	if (int32(16843008)-v66|v66)&v69 == v69 {
		v60 = v60 + int32(4)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v75 = v60
	goto L21
L20:
	;
	goto L19
L21:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v79 != 0 {
		v75 = v75 + int32(1)
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v81 = v75
	goto L8
L23:
	;
	goto L22
L24:
	;
	v94 = v21 + v91
	v95 = F_repalloc(m, v22, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	v97 = v19
	v98 = v21
	v99 = v22
	goto L26
L26:
	;
	v100 = F_strlen(m, v99)
	mBase = m.M
	v102 = F_strcpy(m, v100+v99, v31)
	mBase = m.M
	goto L28
L27:
	;
	v97 = v19 + v91
	v98 = v94
	v99 = v95
	goto L26
L28:
	;
	if v17 == int32(1) {
		v268 = v99
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v109 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_oidvectortypes[0])))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_oidvectortypes[1])))
	v113 = v99
	v114 = v97 - v89
	v115 = int32(1)
	v117 = v98
	goto L30
L30:
	;
	v124 = int32(2)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(24)+v115<<(uint(v124)%32))))
	v130 = F_format_type_extended(m, v127, int32(-1), v124)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v268 = v196
	goto L4
L32:
	;
	if v130&int32(3) == int32(0) {
		v155 = v130
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v190 = v188 + int32(2)
	if base.Ui32(v114) < base.Ui32(v190) {
		goto L50
	} else {
		goto L51
	}
L34:
	;
	v188 = v180 - v130
	goto L33
L35:
	;
	v159 = v155
	goto L44
L36:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v139 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v188 = int32(0)
	goto L33
L38:
	;
	goto L39
L39:
	;
	v144 = v130
	goto L40
L40:
	;
	v148 = v144 + int32(1)
	if v148&int32(3) == int32(0) {
		v155 = v148
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v180 = v148
	goto L34
L42:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v153 != 0 {
		v144 = v148
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v168 = int32(-2139062144)
	if (int32(16843008)-v165|v165)&v168 == v168 {
		v159 = v159 + int32(4)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v174 = v159
	goto L47
L46:
	;
	goto L45
L47:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v178 != 0 {
		v174 = v174 + int32(1)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v180 = v174
	goto L34
L49:
	;
	goto L48
L50:
	;
	v193 = v190 + v117
	v194 = F_repalloc(m, v113, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v196 = v113
	v197 = v114
	v198 = v117
	goto L52
L52:
	;
	if v196&int32(3) == int32(0) {
		v222 = v196
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v196 = v194
	v197 = v114 + v190
	v198 = v193
	goto L52
L54:
	;
	v256 = v255 + v196
	*(*uint16)(unsafe.Add(mBase, uint32(v256))) = uint16(v109)
	*(*uint8)(unsafe.Add(mBase, uint32(v256)+2)) = uint8(v111)
	v262 = F_strlen(m, v196)
	mBase = m.M
	v264 = F_strcpy(m, v262+v196, v130)
	mBase = m.M
	goto L71
L55:
	;
	v255 = v247 - v196
	goto L54
L56:
	;
	v226 = v222
	goto L65
L57:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	if v206 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v255 = int32(0)
	goto L54
L59:
	;
	goto L60
L60:
	;
	v211 = v196
	goto L61
L61:
	;
	v215 = v211 + int32(1)
	if v215&int32(3) == int32(0) {
		v222 = v215
		goto L56
	} else {
		goto L63
	}
L62:
	;
	v247 = v215
	goto L55
L63:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v220 != 0 {
		v211 = v215
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v235 = int32(-2139062144)
	if (int32(16843008)-v232|v232)&v235 == v235 {
		v226 = v226 + int32(4)
		goto L65
	} else {
		goto L67
	}
L66:
	;
	v241 = v226
	goto L68
L67:
	;
	goto L66
L68:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	if v245 != 0 {
		v241 = v241 + int32(1)
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v247 = v241
	goto L55
L70:
	;
	goto L69
L71:
	;
	v266 = v115 + int32(1)
	if v266 != v17 {
		v113 = v196
		v114 = v197 - v188 - int32(2)
		v115 = v266
		v117 = v198
		goto L30
	} else {
		goto L72
	}
L72:
	;
	goto L31
L73:
	;
	return v279
}
func F_okeys_array_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	if v8 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_okeys_array_start_0)
				F_errmsg(m, int32(_a_F_okeys_array_start_1), v5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_okeys_array_start_2), int32(818), int32(_a_F_okeys_array_start_3))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		m.G0 = v5 + int32(16)
		return int32(0)
	}
}
func F_okeys_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	if v9 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_okeys_scalar_0)
				F_errmsg(m, int32(_a_F_okeys_scalar_1), v6)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_okeys_scalar_2), int32(833), int32(_a_F_okeys_scalar_3))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
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
		m.G0 = v6 + int32(16)
		return int32(0)
	}
}
func F_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1&int32(64) == int32(0) {
		v14 = int32(_a_F_open_0)
		if l1&v14 != v14 {
			v22 = int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2 + int32(4)
			v21 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2))))
			v22 = v21
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2 + int32(4)
		v21 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2))))
		v22 = v21
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v22
	v27 = m.Env.X__syscall_openat(m, int32(-100), l0, l1|int32(_a_F_open_1), v7)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v27) {
		*(*int32)(unsafe.Add(mBase, _c_F_open[0])) = int32(0) - v27
		v35 = int32(-1)
	} else {
		v35 = v27
	}
	m.G0 = v7 + int32(16)
	return v35
}
func F_open_auth_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	if int32(11) <= l2 {
		v14 = F_errstart(m, l1, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v14 != 0 {
				F_errcode_for_file_access(m)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
					F_errmsg(m, int32(_a_F_open_auth_file_0), v7+int32(-48))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_open_auth_file_1), int32(612), int32(_a_F_open_auth_file_2))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							if l3 == int32(0) {
								v83 = v5
								m.G0 = v9 - int32(-64)
								return v83
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								v35 = F_psprintf(m, int32(_a_F_open_auth_file_0), v9)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v35
									v83 = v5
									m.G0 = v9 - int32(-64)
									return v83
								}
							}
						}
					}
				}
			} else {
				if l3 == int32(0) {
					v83 = v5
					m.G0 = v9 - int32(-64)
					return v83
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
					v35 = F_psprintf(m, int32(_a_F_open_auth_file_0), v9)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v35
						v83 = v5
						m.G0 = v9 - int32(-64)
						return v83
					}
				}
			}
		}
	} else {
		v39 = F_AllocateFile(m, l0, int32(_a_F_open_auth_file_3))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			if v39 == int32(0) {
				v44 = *(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[0]))
				v46 = F_errstart(m, l1, int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					if v46 != 0 {
						F_errcode_for_file_access(m)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l0
							F_errmsg(m, int32(_a_F_open_auth_file_4), v7+int32(-16))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_open_auth_file_1), int32(627), int32(_a_F_open_auth_file_2))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									if l3 != 0 {
										*(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[0])) = v44
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0
										v67 = F_psprintf(m, int32(_a_F_open_auth_file_4), v7+int32(-32))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v67
											*(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[0])) = v44
											v83 = int32(0)
											m.G0 = v9 - int32(-64)
											return v83
										}
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[0])) = v44
										v83 = int32(0)
										m.G0 = v9 - int32(-64)
										return v83
									}
								}
							}
						}
					} else {
						if l3 != 0 {
							*(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[0])) = v44
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0
							v67 = F_psprintf(m, int32(_a_F_open_auth_file_4), v7+int32(-32))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v67
								*(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[0])) = v44
								v83 = int32(0)
								m.G0 = v9 - int32(-64)
								return v83
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[0])) = v44
							v83 = int32(0)
							m.G0 = v9 - int32(-64)
							return v83
						}
					}
				}
			} else {
				if l2 != 0 {
					v83 = v39
					m.G0 = v9 - int32(-64)
					return v83
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[1]))
					v80 = F_AllocSetContextCreateInternal(m, v75, int32(_a_F_open_auth_file_5), int32(0), int32(1024), int32(_a_F_open_auth_file_6))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[2])) = v80
						v83 = v39
						m.G0 = v9 - int32(-64)
						return v83
					}
				}
			}
		}
	}
}
func F_or_arg_index_match_cmp_group(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v6 < v7 {
		v18 = int32(-1)
	} else {
		if v7 < v6 {
			v18 = int32(1)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			if v12 < v13 {
				v18 = int32(-1)
			} else {
				v18 = base.B2i32(v13 < v12)
			}
		}
	}
	return v18
}
func F_ordered_set_transition_multi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v26 int32
	_ = v26
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
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	v2 = int32(0)
	v12 = l0 + int32(20)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v13 == int32(1) {
		v17 = F_ordered_set_startup(m, l0, int32(1))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = v17
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
			m.T0[v26].(func(*base.Module, int32))(m, v24)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
				if v29 < int32(2) {
					v114 = v2
				} else {
					v32 = int32(1)
					v33 = v29 - v32
					v36 = int32(0)
					if v29 != int32(2) {
						v41 = v36
						v48 = v2
						for {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
							v52 = int32(2)
							v56 = v41 | int32(1)
							v57 = int32(3)
							v59 = v12 + v56<<(uint(v57)%32)
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
							*(*int32)(unsafe.Add(mBase, uint32(v51+v41<<(uint(v52)%32)))) = v60
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)))
							*(*uint8)(unsafe.Add(mBase, uint32(v62+v41))) = uint8(v64)
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
							v71 = v41 + v52
							v74 = v12 + v71<<(uint(v57)%32)
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
							*(*int32)(unsafe.Add(mBase, uint32(v66+v56<<(uint(v52)%32)))) = v75
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+4)))
							*(*uint8)(unsafe.Add(mBase, uint32(v56+v77))) = uint8(v79)
							v82 = v48 + v52
							if v82 != v33&int32(-2) {
								v41 = v71
								v48 = v82
								continue
							} else {
								break
							}
							break
						}
						v84 = v71
					} else {
						v84 = v36
					}
					if v33&v32 == int32(0) {
						v114 = v33
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
						v102 = v84<<(uint(int32(3))%32) + v12
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v96+v84<<(uint(int32(2))%32)))) = v103
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
						v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+12)))
						*(*uint8)(unsafe.Add(mBase, uint32(v105+v84))) = uint8(v107)
						v114 = v33
					}
				}
				v119 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
				v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+50)))
				if v121 == int32(104) {
					v124 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
					v128 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v124+v114<<(uint(int32(2))%32)))) = v128
					v130 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
					*(*uint8)(unsafe.Add(mBase, uint32(v130+v114))) = uint8(v128)
				} else {
				}
				v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
				v136 = v134 & int32(_a_F_ordered_set_transition_multi_0)
				*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)) = uint16(v136)
				v138 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
				v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
				*(*uint16)(unsafe.Add(mBase, uint32(v24)+6)) = uint16(v139)
				v141 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
				F_tuplesort_puttupleslot(m, v141, v24)
				mBase = m.M
				v143 = m.ExcPending
				if v143 != 0 {
					return int32(0)
				} else {
					v144 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v144 + int64(1)
					return v22
				}
			}
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v22 = v21
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
		m.T0[v26].(func(*base.Module, int32))(m, v24)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if v29 < int32(2) {
				v114 = v2
			} else {
				v32 = int32(1)
				v33 = v29 - v32
				v36 = int32(0)
				if v29 != int32(2) {
					v41 = v36
					v48 = v2
					for {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
						v52 = int32(2)
						v56 = v41 | int32(1)
						v57 = int32(3)
						v59 = v12 + v56<<(uint(v57)%32)
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
						*(*int32)(unsafe.Add(mBase, uint32(v51+v41<<(uint(v52)%32)))) = v60
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)))
						*(*uint8)(unsafe.Add(mBase, uint32(v62+v41))) = uint8(v64)
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
						v71 = v41 + v52
						v74 = v12 + v71<<(uint(v57)%32)
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
						*(*int32)(unsafe.Add(mBase, uint32(v66+v56<<(uint(v52)%32)))) = v75
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+4)))
						*(*uint8)(unsafe.Add(mBase, uint32(v56+v77))) = uint8(v79)
						v82 = v48 + v52
						if v82 != v33&int32(-2) {
							v41 = v71
							v48 = v82
							continue
						} else {
							break
						}
						break
					}
					v84 = v71
				} else {
					v84 = v36
				}
				if v33&v32 == int32(0) {
					v114 = v33
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
					v102 = v84<<(uint(int32(3))%32) + v12
					v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v96+v84<<(uint(int32(2))%32)))) = v103
					v105 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
					v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+12)))
					*(*uint8)(unsafe.Add(mBase, uint32(v105+v84))) = uint8(v107)
					v114 = v33
				}
			}
			v119 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
			v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+50)))
			if v121 == int32(104) {
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
				v128 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v124+v114<<(uint(int32(2))%32)))) = v128
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
				*(*uint8)(unsafe.Add(mBase, uint32(v130+v114))) = uint8(v128)
			} else {
			}
			v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
			v136 = v134 & int32(_a_F_ordered_set_transition_multi_0)
			*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)) = uint16(v136)
			v138 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
			v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
			*(*uint16)(unsafe.Add(mBase, uint32(v24)+6)) = uint16(v139)
			v141 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
			F_tuplesort_puttupleslot(m, v141, v24)
			mBase = m.M
			v143 = m.ExcPending
			if v143 != 0 {
				return int32(0)
			} else {
				v144 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v144 + int64(1)
				return v22
			}
		}
	}
}
func F_overlaps_timetz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v18 int32
	_ = v18
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	v11 = int32(1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v18 == v11 {
		if v14&int32(1) == int32(0) {
			v37 = v17
			v39 = v17
			v40 = v11
			if v13&int32(1) != 0 {
				if v12&int32(1) != 0 {
					v107 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
					v117 = int32(0)
					return v117
				} else {
					v48 = F_DirectFunctionCall2Coll(m, int32(1287), int32(0), v39, v15)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						if v48 != 0 {
							v107 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
							v117 = int32(0)
							return v117
						} else {
							v68 = v15
							v69 = int32(1)
							v73 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v39, v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v73 != 0 {
									if v40 != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										v77 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v68, v37)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											if v69&base.B2i32(v77 == int32(0)) != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												return base.B2i32(v77 != int32(0))
											}
										}
									}
								} else {
									v85 = int32(1)
									if v69|v40 != v85 {
										v117 = v85
									} else {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
									}
									return v117
								}
							}
						}
					}
				}
			} else {
				if v12&int32(1) != 0 {
					v55 = F_DirectFunctionCall2Coll(m, int32(1287), int32(0), v39, v16)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v107 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
							v117 = int32(0)
							return v117
						} else {
							v68 = v16
							v69 = int32(1)
							v73 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v39, v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v73 != 0 {
									if v40 != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										v77 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v68, v37)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											if v69&base.B2i32(v77 == int32(0)) != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												return base.B2i32(v77 != int32(0))
											}
										}
									}
								} else {
									v85 = int32(1)
									if v69|v40 != v85 {
										v117 = v85
									} else {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
									}
									return v117
								}
							}
						}
					}
				} else {
					v58 = int32(0)
					v59 = int32(1287)
					v63 = F_DirectFunctionCall2Coll(m, v59, v58, v16, v15)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						if v63 != 0 {
							v65 = v15
						} else {
							v65 = v16
						}
						v66 = F_DirectFunctionCall2Coll(m, v59, v58, v39, v65)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							if v66 != 0 {
								if v63 != 0 {
									v91 = v16
								} else {
									v91 = v15
								}
								v92 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v39, v91)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									if v40&base.B2i32(v92 == int32(0)) != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										return base.B2i32(v92 != int32(0))
									}
								}
							} else {
								v68 = v65
								v69 = v58
								v73 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v39, v68)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 != 0 {
										if v40 != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											v77 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v68, v37)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if v69&base.B2i32(v77 == int32(0)) != 0 {
													v107 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
													v117 = int32(0)
													return v117
												} else {
													return base.B2i32(v77 != int32(0))
												}
											}
										}
									} else {
										v85 = int32(1)
										if v69|v40 != v85 {
											v117 = v85
										} else {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
										}
										return v117
									}
								}
							}
						}
					}
				}
			}
		} else {
			v107 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
			v117 = int32(0)
			return v117
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v14&int32(1) != 0 {
			v37 = v17
			v39 = v25
			v40 = v11
			if v13&int32(1) != 0 {
				if v12&int32(1) != 0 {
					v107 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
					v117 = int32(0)
					return v117
				} else {
					v48 = F_DirectFunctionCall2Coll(m, int32(1287), int32(0), v39, v15)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						if v48 != 0 {
							v107 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
							v117 = int32(0)
							return v117
						} else {
							v68 = v15
							v69 = int32(1)
							v73 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v39, v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v73 != 0 {
									if v40 != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										v77 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v68, v37)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											if v69&base.B2i32(v77 == int32(0)) != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												return base.B2i32(v77 != int32(0))
											}
										}
									}
								} else {
									v85 = int32(1)
									if v69|v40 != v85 {
										v117 = v85
									} else {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
									}
									return v117
								}
							}
						}
					}
				}
			} else {
				if v12&int32(1) != 0 {
					v55 = F_DirectFunctionCall2Coll(m, int32(1287), int32(0), v39, v16)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v107 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
							v117 = int32(0)
							return v117
						} else {
							v68 = v16
							v69 = int32(1)
							v73 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v39, v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v73 != 0 {
									if v40 != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										v77 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v68, v37)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											if v69&base.B2i32(v77 == int32(0)) != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												return base.B2i32(v77 != int32(0))
											}
										}
									}
								} else {
									v85 = int32(1)
									if v69|v40 != v85 {
										v117 = v85
									} else {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
									}
									return v117
								}
							}
						}
					}
				} else {
					v58 = int32(0)
					v59 = int32(1287)
					v63 = F_DirectFunctionCall2Coll(m, v59, v58, v16, v15)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						if v63 != 0 {
							v65 = v15
						} else {
							v65 = v16
						}
						v66 = F_DirectFunctionCall2Coll(m, v59, v58, v39, v65)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							if v66 != 0 {
								if v63 != 0 {
									v91 = v16
								} else {
									v91 = v15
								}
								v92 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v39, v91)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									if v40&base.B2i32(v92 == int32(0)) != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										return base.B2i32(v92 != int32(0))
									}
								}
							} else {
								v68 = v65
								v69 = v58
								v73 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v39, v68)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 != 0 {
										if v40 != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											v77 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v68, v37)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if v69&base.B2i32(v77 == int32(0)) != 0 {
													v107 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
													v117 = int32(0)
													return v117
												} else {
													return base.B2i32(v77 != int32(0))
												}
											}
										}
									} else {
										v85 = int32(1)
										if v69|v40 != v85 {
											v117 = v85
										} else {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
										}
										return v117
									}
								}
							}
						}
					}
				}
			}
		} else {
			v28 = int32(0)
			v31 = F_DirectFunctionCall2Coll(m, int32(1287), v28, v25, v17)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				if v31 != 0 {
					v35 = v17
				} else {
					v35 = v25
				}
				if v31 != 0 {
					v36 = v25
				} else {
					v36 = v17
				}
				v37 = v36
				v39 = v35
				v40 = v28
				if v13&int32(1) != 0 {
					if v12&int32(1) != 0 {
						v107 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
						v117 = int32(0)
						return v117
					} else {
						v48 = F_DirectFunctionCall2Coll(m, int32(1287), int32(0), v39, v15)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							if v48 != 0 {
								v107 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
								v117 = int32(0)
								return v117
							} else {
								v68 = v15
								v69 = int32(1)
								v73 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v39, v68)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 != 0 {
										if v40 != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											v77 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v68, v37)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if v69&base.B2i32(v77 == int32(0)) != 0 {
													v107 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
													v117 = int32(0)
													return v117
												} else {
													return base.B2i32(v77 != int32(0))
												}
											}
										}
									} else {
										v85 = int32(1)
										if v69|v40 != v85 {
											v117 = v85
										} else {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
										}
										return v117
									}
								}
							}
						}
					}
				} else {
					if v12&int32(1) != 0 {
						v55 = F_DirectFunctionCall2Coll(m, int32(1287), int32(0), v39, v16)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v55 != 0 {
								v107 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
								v117 = int32(0)
								return v117
							} else {
								v68 = v16
								v69 = int32(1)
								v73 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v39, v68)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 != 0 {
										if v40 != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											v77 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v68, v37)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if v69&base.B2i32(v77 == int32(0)) != 0 {
													v107 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
													v117 = int32(0)
													return v117
												} else {
													return base.B2i32(v77 != int32(0))
												}
											}
										}
									} else {
										v85 = int32(1)
										if v69|v40 != v85 {
											v117 = v85
										} else {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
										}
										return v117
									}
								}
							}
						}
					} else {
						v58 = int32(0)
						v59 = int32(1287)
						v63 = F_DirectFunctionCall2Coll(m, v59, v58, v16, v15)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							if v63 != 0 {
								v65 = v15
							} else {
								v65 = v16
							}
							v66 = F_DirectFunctionCall2Coll(m, v59, v58, v39, v65)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								if v66 != 0 {
									if v63 != 0 {
										v91 = v16
									} else {
										v91 = v15
									}
									v92 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v39, v91)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int32(0)
									} else {
										if v40&base.B2i32(v92 == int32(0)) != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											return base.B2i32(v92 != int32(0))
										}
									}
								} else {
									v68 = v65
									v69 = v58
									v73 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v39, v68)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										if v73 != 0 {
											if v40 != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												v77 = F_DirectFunctionCall2Coll(m, int32(1288), int32(0), v68, v37)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													if v69&base.B2i32(v77 == int32(0)) != 0 {
														v107 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
														v117 = int32(0)
														return v117
													} else {
														return base.B2i32(v77 != int32(0))
													}
												}
											}
										} else {
											v85 = int32(1)
											if v69|v40 != v85 {
												v117 = v85
											} else {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
											}
											return v117
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
