package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyAttributeOutCSV(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v17 = base.B2i32(v13 == int32(1))
	goto L3
L2:
	;
	v17 = int32(0)
	goto L3
L3:
	;
	v18 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9))))
	v19 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if l2 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if base.B2i32(v27 == int32(0))|base.B2i32(v27 != v30) != 0 {
		v48 = v27
		v49 = v30
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
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v53 = base.B2i32(v48-v49 == int32(0))
	goto L6
L8:
	;
	goto L7
L9:
	;
	v33 = l1
	v34 = v24
	goto L10
L10:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v38 == int32(0) {
		v48 = v38
		v49 = v37
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v48 = v38
	v49 = v37
	goto L8
L12:
	;
	v41 = int32(1)
	if v38 == v37 {
		v33 = v33 + v41
		v34 = v34 + v41
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v57 = F_strlen(m, l1)
	mBase = m.M
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v59 = F_pg_server_to_any(m, l1, v57, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v61 = l1
	goto L16
L16:
	;
	if v53 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	return
L18:
	;
	v61 = v59
	goto L16
L19:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	if v125 <= v122+int32(1) {
		goto L40
	} else {
		goto L41
	}
L20:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if base.B2i32(v62 == int32(92))&v17 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v110 = F_strlen(m, v61)
	mBase = m.M
	F_appendBinaryStringInfo(m, v109, v61, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L17
	} else {
		goto L38
	}
L22:
	;
	v74 = v61
	v76 = v62
	goto L29
L23:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v66 != int32(46) {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v62 == int32(0) {
		goto L21
	} else {
		goto L28
	}
L26:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)))
	if v69 != 0 {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L19
L28:
	;
	goto L22
L29:
	;
	if base.B2i32(v20 == v76)|base.B2i32(v76 == v19&int32(255)) != 0 {
		goto L19
	} else {
		goto L31
	}
L30:
	;
	goto L21
L31:
	;
	switch v76 - int32(10) {
	case 0, 3:
		goto L19
	default:
		goto L32
	}
L32:
	;
	if int32(0) <= base.I32_extend8_s(v76) {
		v98 = int32(1)
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v99 = v98 + v74
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v100 != 0 {
		v74 = v99
		v76 = v100
		goto L29
	} else {
		goto L37
	}
L34:
	;
	v91 = int32(1)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v92 != v91 {
		v98 = v91
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v96 = F_pg_encoding_mblen(m, v95, v74)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L17
	} else {
		goto L36
	}
L36:
	;
	v98 = v96
	goto L33
L37:
	;
	goto L30
L38:
	;
	return
L39:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v143 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	F_appendStringInfoChar(m, v121, v19)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L17
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	*(*uint8)(unsafe.Add(mBase, uint32(v129+v122))) = uint8(v19)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v135 = v133 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v132)+4)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v139 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v137+v135))) = uint8(v139)
	goto L39
L43:
	;
	goto L39
L44:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	if v223 <= v220+int32(1) {
		goto L67
	} else {
		goto L68
	}
L45:
	;
	v147 = v61
	v148 = v61
	v150 = v143
	goto L46
L46:
	;
	v154 = int32(255)
	if base.B2i32(v19&v154 != v150)&base.B2i32(v150 != v18&v154) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if base.Ui32(v204) <= base.Ui32(v190) {
		goto L44
	} else {
		goto L65
	}
L48:
	;
	if base.Ui32(v148) < base.Ui32(v147) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v190 = v148
	goto L50
L50:
	;
	if int32(0) <= base.I32_extend8_s(v150) {
		v203 = int32(1)
		goto L60
	} else {
		goto L61
	}
L51:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v164, v148, v147-v148)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L17
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168)+8))
	if v172 <= v169+int32(1) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L53
L55:
	;
	v190 = v147
	goto L50
L56:
	;
	F_appendStringInfoChar(m, v168, v18)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L17
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	*(*uint8)(unsafe.Add(mBase, uint32(v176+v169))) = uint8(v18)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v182 = v180 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = v182
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v184+v182))) = uint8(v186)
	goto L55
L59:
	;
	goto L55
L60:
	;
	v204 = v203 + v147
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v205 != 0 {
		v147 = v204
		v148 = v190
		v150 = v205
		goto L46
	} else {
		goto L64
	}
L61:
	;
	v196 = int32(1)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v197 != v196 {
		v203 = v196
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v201 = F_pg_encoding_mblen(m, v200, v147)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L17
	} else {
		goto L63
	}
L63:
	;
	v203 = v201
	goto L60
L64:
	;
	goto L47
L65:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v207, v190, v204-v190)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L17
	} else {
		goto L66
	}
L66:
	;
	goto L44
L67:
	;
	F_appendStringInfoChar(m, v219, v19)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L17
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	*(*uint8)(unsafe.Add(mBase, uint32(v227+v220))) = uint8(v19)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	v233 = v231 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v230)+4)) = v233
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v237 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v235+v233))) = uint8(v237)
	return
L70:
	;
	return
}
