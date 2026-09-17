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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v26 int32
	_ = v26
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
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
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(1088)
	m.G0 = v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v12 == v2 {
		v108 = v2
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L8
	} else {
		goto L83
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L8
	} else {
		goto L80
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L8
	} else {
		goto L76
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L8
	} else {
		goto L71
	}
L5:
	;
	v111 = F_AllocateFile(m, int32(_a_F_pg_current_logfile_0), int32(_a_F_pg_current_logfile_1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L8
	} else {
		goto L38
	}
L6:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v15 != 0 {
		v108 = v2
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v21 = F_text_to_cstring(m, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v23 = int32(_a_F_pg_current_logfile_2)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_current_logfile[0])))
	if base.B2i32(v26 == int32(0))|base.B2i32(v26 != v29) != 0 {
		v47 = v26
		v48 = v29
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v47-v48 == int32(0) {
		v108 = v21
		goto L5
	} else {
		goto L18
	}
L12:
	;
	goto L11
L13:
	;
	v32 = v21
	v33 = v23
	goto L14
L14:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	if v37 == int32(0) {
		v47 = v37
		v48 = v36
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v47 = v37
	v48 = v36
	goto L12
L16:
	;
	v40 = int32(1)
	if v37 == v36 {
		v32 = v32 + v40
		v33 = v33 + v40
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v52 = int32(_a_F_pg_current_logfile_3)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_current_logfile[1])))
	if base.B2i32(v55 == int32(0))|base.B2i32(v55 != v58) != 0 {
		v76 = v55
		v77 = v58
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v76-v77 == int32(0) {
		v108 = v21
		goto L5
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v61 = v21
	v62 = v52
	goto L22
L22:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v66 == int32(0) {
		v76 = v66
		v77 = v65
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v76 = v66
	v77 = v65
	goto L20
L24:
	;
	v69 = int32(1)
	if v66 == v65 {
		v61 = v61 + v69
		v62 = v62 + v69
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v81 = int32(_a_F_pg_current_logfile_4)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_current_logfile[2])))
	if base.B2i32(v84 == int32(0))|base.B2i32(v84 != v87) != 0 {
		v105 = v84
		v106 = v87
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v105-v106 != 0 {
		goto L4
	} else {
		goto L34
	}
L28:
	;
	goto L27
L29:
	;
	v90 = v21
	v91 = v81
	goto L30
L30:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	if v95 == int32(0) {
		v105 = v95
		v106 = v94
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v105 = v95
	v106 = v94
	goto L28
L32:
	;
	v98 = int32(1)
	if v95 == v94 {
		v90 = v90 + v98
		v91 = v91 + v98
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v108 = v21
	goto L5
L35:
	;
	m.G0 = v10 + int32(1088)
	return v206
L36:
	;
	v196 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v196)
	v206 = int32(0)
	goto L35
L37:
	;
	v187 = F_FreeFile(m, v111)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L8
	} else {
		goto L70
	}
L38:
	;
	if v111 != 0 {
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
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_pg_current_logfile[3]))
	if v184 != int32(44) {
		goto L3
	} else {
		goto L69
	}
L42:
	;
	v121 = v10 - int32(-64)
	v123 = F_fgets(m, v121, int32(1024), v111)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L44
	}
L43:
	;
	v179 = F_FreeFile(m, v111)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L67
	}
L44:
	;
	if v123 == int32(0) {
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v127 = int32(32)
	v128 = F___strchrnul(m, v121, v127)
	mBase = m.M
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v130 == v127 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v134 == int32(0) {
		goto L2
	} else {
		goto L50
	}
L47:
	;
	v134 = v128
	goto L49
L48:
	;
	v134 = int32(0)
	goto L49
L49:
	;
	goto L46
L50:
	;
	v137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v137)
	v140 = v134 + int32(1)
	v141 = int32(10)
	v142 = F___strchrnul(m, v140, v141)
	mBase = m.M
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v144 == v141 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v148 == int32(0) {
		goto L1
	} else {
		goto L55
	}
L52:
	;
	v148 = v142
	goto L54
L53:
	;
	v148 = v137
	goto L54
L54:
	;
	goto L51
L55:
	;
	v151 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v151)
	if v108 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if base.B2i32(v155 == int32(0))|base.B2i32(v155 != v158) != 0 {
		v176 = v155
		v177 = v158
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
	if v176-v177 != 0 {
		goto L42
	} else {
		goto L66
	}
L60:
	;
	goto L59
L61:
	;
	v161 = v108
	v162 = v121
	goto L62
L62:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	if v166 == int32(0) {
		v176 = v166
		v177 = v165
		goto L60
	} else {
		goto L64
	}
L63:
	;
	v176 = v166
	v177 = v165
	goto L60
L64:
	;
	v169 = int32(1)
	if v166 == v165 {
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
	goto L58
L67:
	;
	v181 = F_cstring_to_text(m, v140)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	v206 = v181
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
	v217 = m.ExcPending
	if v217 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v21
	F_errmsg(m, int32(_a_F_pg_current_logfile_5), v10+int32(48))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	F_errhint(m, int32(_a_F_pg_current_logfile_6), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_pg_current_logfile_7), int32(1019), int32(_a_F_pg_current_logfile_8))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
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
	v238 = m.ExcPending
	if v238 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_pg_current_logfile_0)
	F_errmsg(m, int32(_a_F_pg_current_logfile_9), v10)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_pg_current_logfile_7), int32(1029), int32(_a_F_pg_current_logfile_8))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_pg_current_logfile_0)
	F_errmsg_internal(m, int32(_a_F_pg_current_logfile_10), v10+int32(16))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_pg_current_logfile_7), int32(1055), int32(_a_F_pg_current_logfile_8))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(_a_F_pg_current_logfile_0)
	F_errmsg_internal(m, int32(_a_F_pg_current_logfile_11), v10+int32(32))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L8
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_pg_current_logfile_7), int32(1066), int32(_a_F_pg_current_logfile_8))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
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
