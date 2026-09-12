package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_current_logfile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(1088)
	m.G0 = v9
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v11 == v2 {
		v104 = v2
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L8
	} else {
		goto L87
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L84
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L8
	} else {
		goto L80
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L8
	} else {
		goto L75
	}
L5:
	;
	v107 = F_AllocateFile(m, int32(164643), int32(230594))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L8
	} else {
		goto L41
	}
L6:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v14 != 0 {
		v104 = v2
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v20 = F_text_to_cstring(m, v16)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v22 = int32(206566)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1116])))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v26 == int32(0) {
		v45 = v25
		v46 = v26
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v46-v45 == int32(0) {
		v104 = v20
		goto L5
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	if v25 != v26 {
		v45 = v25
		v46 = v26
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v30 = v20
	v31 = v22
	goto L15
L15:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v35 == int32(0) {
		v45 = v34
		v46 = v35
		goto L12
	} else {
		goto L17
	}
L16:
	;
	v45 = v34
	v46 = v35
	goto L12
L17:
	;
	v38 = int32(1)
	if v34 == v35 {
		v30 = v30 + v38
		v31 = v31 + v38
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v50 = int32(326750)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1117])))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v54 == int32(0) {
		v73 = v53
		v74 = v54
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v74-v73 == int32(0) {
		v104 = v20
		goto L5
	} else {
		goto L28
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
	v58 = v20
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
	v78 = int32(326764)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1118])))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v82 == int32(0) {
		v101 = v81
		v102 = v82
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v102-v101 != 0 {
		goto L4
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	if v81 != v82 {
		v101 = v81
		v102 = v82
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v86 = v20
	v87 = v78
	goto L33
L33:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if v91 == int32(0) {
		v101 = v90
		v102 = v91
		goto L30
	} else {
		goto L35
	}
L34:
	;
	v101 = v90
	v102 = v91
	goto L30
L35:
	;
	v94 = int32(1)
	if v90 == v91 {
		v86 = v86 + v94
		v87 = v87 + v94
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v104 = v20
	goto L5
L38:
	;
	m.G0 = v9 + int32(1088)
	return v204
L39:
	;
	v204 = int32(0)
	goto L38
L40:
	;
	v187 = F_FreeFile(m, v107)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L8
	} else {
		goto L74
	}
L41:
	;
	if v107 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	goto L45
L43:
	;
	goto L44
L44:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[159]))
	if v182 != int32(44) {
		goto L3
	} else {
		goto L73
	}
L45:
	;
	v118 = F_fgets(m, v9-int32(-64), int32(1024), v107)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L47
	}
L46:
	;
	v177 = F_FreeFile(m, v107)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L8
	} else {
		goto L71
	}
L47:
	;
	if v118 == int32(0) {
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v124 = int32(32)
	v125 = F___strchrnul(m, v9-int32(-64), v124)
	mBase = m.M
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if v127 == v124 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v131 == int32(0) {
		goto L2
	} else {
		goto L53
	}
L50:
	;
	v131 = v125
	goto L52
L51:
	;
	v131 = int32(0)
	goto L52
L52:
	;
	goto L49
L53:
	;
	v134 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v134)
	v137 = v131 + int32(1)
	v138 = int32(10)
	v139 = F___strchrnul(m, v137, v138)
	mBase = m.M
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v141 == v138 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v145 == int32(0) {
		goto L1
	} else {
		goto L58
	}
L55:
	;
	v145 = v139
	goto L57
L56:
	;
	v145 = v134
	goto L57
L57:
	;
	goto L54
L58:
	;
	v148 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v148)
	if v104 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v151 = v9 - int32(-64)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v155 == int32(0) {
		v174 = v154
		v175 = v155
		goto L63
	} else {
		goto L64
	}
L60:
	;
	goto L61
L61:
	;
	goto L46
L62:
	;
	if v175-v174 != 0 {
		goto L45
	} else {
		goto L70
	}
L63:
	;
	goto L62
L64:
	;
	if v154 != v155 {
		v174 = v154
		v175 = v155
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v159 = v104
	v160 = v151
	goto L66
L66:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+1)))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+1)))
	if v164 == int32(0) {
		v174 = v163
		v175 = v164
		goto L63
	} else {
		goto L68
	}
L67:
	;
	v174 = v163
	v175 = v164
	goto L63
L68:
	;
	v167 = int32(1)
	if v163 == v164 {
		v159 = v159 + v167
		v160 = v160 + v167
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	goto L61
L71:
	;
	v179 = F_cstring_to_text(m, v137)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	v204 = v179
	goto L38
L73:
	;
	v185 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v185)
	goto L39
L74:
	;
	v189 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v189)
	goto L39
L75:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v20
	F_errmsg(m, int32(442288), v9+int32(48))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	F_errhint(m, int32(652643), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(499164), int32(1019), int32(385740))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(164643)
	F_errmsg(m, int32(299166), v9)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L8
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(499164), int32(1029), int32(385740))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(164643)
	F_errmsg_internal(m, int32(696097), v9+int32(16))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(499164), int32(1055), int32(385740))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = int32(164643)
	F_errmsg_internal(m, int32(696063), v9+int32(32))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(499164), int32(1066), int32(385740))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
