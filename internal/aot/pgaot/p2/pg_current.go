package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_current_logfile(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
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
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
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
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
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
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v211 int64
	_ = v211
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(1088)
	m.G0 = v11
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v13 == v2 {
		v109 = v2
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L8
	} else {
		goto L83
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L8
	} else {
		goto L80
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L8
	} else {
		goto L76
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L8
	} else {
		goto L71
	}
L5:
	;
	v112 = F_AllocateFile(m, int32(_a_F_pg_current_logfile_0), int32(_a_F_pg_current_logfile_1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L38
	}
L6:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v16 != 0 {
		v109 = v2
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int64(0)
L9:
	;
	v22 = F_text_to_cstring(m, v18)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v24 = int32(_a_F_pg_current_logfile_2)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_current_logfile[0])))
	if base.B2i32(v27 == int32(0))|base.B2i32(v27 != v30) != 0 {
		v48 = v27
		v49 = v30
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v48-v49 == int32(0) {
		v109 = v22
		goto L5
	} else {
		goto L18
	}
L12:
	;
	goto L11
L13:
	;
	v33 = v22
	v34 = v24
	goto L14
L14:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v38 == int32(0) {
		v48 = v38
		v49 = v37
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v48 = v38
	v49 = v37
	goto L12
L16:
	;
	v41 = int32(1)
	if v38 == v37 {
		v33 = v33 + v41
		v34 = v34 + v41
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v53 = int32(_a_F_pg_current_logfile_3)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_current_logfile[1])))
	if base.B2i32(v56 == int32(0))|base.B2i32(v56 != v59) != 0 {
		v77 = v56
		v78 = v59
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v77-v78 == int32(0) {
		v109 = v22
		goto L5
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v62 = v22
	v63 = v53
	goto L22
L22:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v67 == int32(0) {
		v77 = v67
		v78 = v66
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v77 = v67
	v78 = v66
	goto L20
L24:
	;
	v70 = int32(1)
	if v67 == v66 {
		v62 = v62 + v70
		v63 = v63 + v70
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v82 = int32(_a_F_pg_current_logfile_4)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_current_logfile[2])))
	if base.B2i32(v85 == int32(0))|base.B2i32(v85 != v88) != 0 {
		v106 = v85
		v107 = v88
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v106-v107 != 0 {
		goto L4
	} else {
		goto L34
	}
L28:
	;
	goto L27
L29:
	;
	v91 = v22
	v92 = v82
	goto L30
L30:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	if v96 == int32(0) {
		v106 = v96
		v107 = v95
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v106 = v96
	v107 = v95
	goto L28
L32:
	;
	v99 = int32(1)
	if v96 == v95 {
		v91 = v91 + v99
		v92 = v92 + v99
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v109 = v22
	goto L5
L35:
	;
	m.G0 = v11 + int32(1088)
	return v211
L36:
	;
	v200 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v200)
	v211 = int64(0)
	goto L35
L37:
	;
	v190 = F_FreeFile(m, v112)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L8
	} else {
		goto L70
	}
L38:
	;
	if v112 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	goto L42
L40:
	;
	goto L41
L41:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_pg_current_logfile[3]))
	if v187 != int32(44) {
		goto L3
	} else {
		goto L69
	}
L42:
	;
	v123 = v11 - int32(-64)
	v125 = F_fgets(m, v123, int32(1024), v112)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L8
	} else {
		goto L44
	}
L43:
	;
	v181 = F_FreeFile(m, v112)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L8
	} else {
		goto L67
	}
L44:
	;
	if v125 == int32(0) {
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v129 = int32(32)
	v130 = F___strchrnul(m, v123, v129)
	mBase = m.M
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v132 == v129 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v136 == int32(0) {
		goto L2
	} else {
		goto L50
	}
L47:
	;
	v136 = v130
	goto L49
L48:
	;
	v136 = int32(0)
	goto L49
L49:
	;
	goto L46
L50:
	;
	v139 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v139)
	v142 = v136 + int32(1)
	v143 = int32(10)
	v144 = F___strchrnul(m, v142, v143)
	mBase = m.M
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v146 == v143 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v150 == int32(0) {
		goto L1
	} else {
		goto L55
	}
L52:
	;
	v150 = v144
	goto L54
L53:
	;
	v150 = v139
	goto L54
L54:
	;
	goto L51
L55:
	;
	v153 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v153)
	if v109 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if base.B2i32(v157 == int32(0))|base.B2i32(v157 != v160) != 0 {
		v178 = v157
		v179 = v160
		goto L60
	} else {
		goto L61
	}
L57:
	;
	goto L58
L58:
	;
	goto L43
L59:
	;
	if v178-v179 != 0 {
		goto L42
	} else {
		goto L66
	}
L60:
	;
	goto L59
L61:
	;
	v163 = v109
	v164 = v123
	goto L62
L62:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	if v168 == int32(0) {
		v178 = v168
		v179 = v167
		goto L60
	} else {
		goto L64
	}
L63:
	;
	v178 = v168
	v179 = v167
	goto L60
L64:
	;
	v171 = int32(1)
	if v168 == v167 {
		v163 = v163 + v171
		v164 = v164 + v171
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	goto L58
L67:
	;
	v183 = F_cstring_to_text(m, v142)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	v211 = base.I64_extend_i32_u(v183)
	goto L35
L69:
	;
	goto L36
L70:
	;
	goto L36
L71:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v22
	F_errmsg(m, int32(_a_F_pg_current_logfile_5), v11+int32(48))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	F_errhint(m, int32(_a_F_pg_current_logfile_6), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_pg_current_logfile_7), int32(992), int32(_a_F_pg_current_logfile_8))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L8
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_pg_current_logfile_0)
	F_errmsg(m, int32(_a_F_pg_current_logfile_9), v11)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_pg_current_logfile_7), int32(1002), int32(_a_F_pg_current_logfile_8))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_pg_current_logfile_0)
	F_errmsg_internal(m, int32(_a_F_pg_current_logfile_10), v11+int32(16))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_pg_current_logfile_7), int32(1028), int32(_a_F_pg_current_logfile_8))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L8
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(_a_F_pg_current_logfile_0)
	F_errmsg_internal(m, int32(_a_F_pg_current_logfile_11), v11+int32(32))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L8
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_pg_current_logfile_7), int32(1039), int32(_a_F_pg_current_logfile_8))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
