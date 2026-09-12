package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyAttributeOutCSV(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
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
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v18 = base.B2i32(v14 == int32(1))
	goto L3
L2:
	;
	v18 = int32(0)
	goto L3
L3:
	;
	v19 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
	v20 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11))))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if l2 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v29 == int32(0) {
		v48 = v28
		v49 = v29
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v53 = int32(1)
	goto L6
L6:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v54 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v53 = base.B2i32(v49-v48 == int32(0))
	goto L6
L8:
	;
	goto L7
L9:
	;
	if v28 != v29 {
		v48 = v28
		v49 = v29
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v33 = l1
	v34 = v25
	goto L11
L11:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v38 == int32(0) {
		v48 = v37
		v49 = v38
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v48 = v37
	v49 = v38
	goto L8
L13:
	;
	v41 = int32(1)
	if v37 == v38 {
		v33 = v33 + v41
		v34 = v34 + v41
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v57 = F_strlen(m, l1)
	mBase = m.M
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v59 = F_pg_server_to_any(m, l1, v57, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v61 = l1
	goto L17
L17:
	;
	if v53 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	return
L19:
	;
	v61 = v59
	goto L17
L20:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	if v131 <= v128+int32(1) {
		goto L42
	} else {
		goto L43
	}
L21:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if base.B2i32(v62 == int32(92))&v18 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v115 = F_strlen(m, v61)
	mBase = m.M
	F_appendBinaryStringInfo(m, v114, v61, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L18
	} else {
		goto L40
	}
L23:
	;
	v77 = v62
	v78 = v61
	goto L30
L24:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v66 != int32(46) {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v62 == int32(0) {
		goto L22
	} else {
		goto L29
	}
L27:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)))
	if v69 != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L20
L29:
	;
	goto L23
L30:
	;
	v84 = v77 & int32(255)
	if v84 == v21&int32(255) {
		goto L20
	} else {
		goto L32
	}
L31:
	;
	goto L22
L32:
	;
	if v84 == v20&int32(255) {
		goto L20
	} else {
		goto L33
	}
L33:
	;
	switch v84 - int32(10) {
	case 0, 3:
		goto L20
	default:
		goto L34
	}
L34:
	;
	if int32(0) <= base.I32_extend8_s(v77) {
		v102 = int32(1)
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v103 = v102 + v78
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v104 != 0 {
		v77 = v104
		v78 = v103
		goto L30
	} else {
		goto L39
	}
L36:
	;
	v95 = int32(1)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v96 != v95 {
		v102 = v95
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v100 = F_pg_encoding_mblen(m, v99, v78)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L18
	} else {
		goto L38
	}
L38:
	;
	v102 = v100
	goto L35
L39:
	;
	goto L31
L40:
	;
	return
L41:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v149 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	F_appendStringInfoChar(m, v127, v20)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L18
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	*(*uint8)(unsafe.Add(mBase, uint32(v135+v128))) = uint8(v20)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v141 = v139 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v141
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v145 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v143+v141))) = uint8(v145)
	goto L41
L45:
	;
	goto L41
L46:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229)+8))
	if v233 <= v230+int32(1) {
		goto L69
	} else {
		goto L70
	}
L47:
	;
	v153 = v61
	v155 = v149
	v156 = v61
	goto L48
L48:
	;
	v161 = int32(255)
	v162 = v155 & v161
	if base.B2i32(v162 != v20&v161)&base.B2i32(v162 != v19&v161) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if base.Ui32(v213) <= base.Ui32(v200) {
		goto L46
	} else {
		goto L67
	}
L50:
	;
	if base.Ui32(v156) < base.Ui32(v153) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v200 = v156
	goto L52
L52:
	;
	if int32(0) <= base.I32_extend8_s(v155) {
		v212 = int32(1)
		goto L62
	} else {
		goto L63
	}
L53:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v173, v156, v153-v156)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L18
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	if v181 <= v178+int32(1) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L55
L57:
	;
	v200 = v153
	goto L52
L58:
	;
	F_appendStringInfoChar(m, v177, v19)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L18
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	*(*uint8)(unsafe.Add(mBase, uint32(v185+v178))) = uint8(v19)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	v191 = v189 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v193+v191))) = uint8(v195)
	goto L57
L61:
	;
	goto L57
L62:
	;
	v213 = v212 + v153
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	if v214 != 0 {
		v153 = v213
		v155 = v214
		v156 = v200
		goto L48
	} else {
		goto L66
	}
L63:
	;
	v205 = int32(1)
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v206 != v205 {
		v212 = v205
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v210 = F_pg_encoding_mblen(m, v209, v153)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	v212 = v210
	goto L62
L66:
	;
	goto L49
L67:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v216, v200, v213-v200)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	goto L46
L69:
	;
	F_appendStringInfoChar(m, v229, v20)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L18
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	*(*uint8)(unsafe.Add(mBase, uint32(v237+v230))) = uint8(v20)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v243 = v241 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v240)+4)) = v243
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v247 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v245+v243))) = uint8(v247)
	return
L72:
	;
	return
}
