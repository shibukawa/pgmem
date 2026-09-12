package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyAttributeOutText(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
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
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v9 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = F_strlen(m, l1)
	mBase = m.M
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = F_pg_server_to_any(m, l1, v12, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v16 = l1
	goto L3
L3:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v18 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return
L5:
	;
	v16 = v14
	goto L3
L6:
	;
	return
L7:
	;
	if base.Ui32(v291) <= base.Ui32(v293) {
		goto L6
	} else {
		goto L94
	}
L8:
	;
	v163 = v17
	v164 = v16
	v165 = v16
	goto L55
L9:
	;
	if v17&int32(255) == int32(0) {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v17&int32(255) == int32(0) {
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L8
L13:
	;
	v30 = v16
	v31 = v17
	v32 = v16
	goto L14
L14:
	;
	v36 = v31 & int32(255)
	if base.Ui32(v36) <= base.Ui32(int32(31)) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v291 = v159
	v293 = v157
	goto L7
L16:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v160 != 0 {
		v30 = v159
		v31 = v160
		v32 = v157
		goto L14
	} else {
		goto L54
	}
L17:
	;
	v157 = v32
	v159 = v30 + int32(1)
	goto L16
L18:
	;
	v40 = v31 - int32(8)
	if base.Ui32(int32(6)) <= base.Ui32(v40&int32(255)) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	if base.B2i32(v36 != int32(92))&base.B2i32(v36 != v8) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L21:
	;
	if base.Ui32(v32) < base.Ui32(v30) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	if v36 != v8 {
		goto L17
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v54 = base.I32_wrap_i64(int64(base.Ui64(int64(125784399180898)) >> (uint(base.I64_extend_i32_u(v40<<(uint(int32(3))%32))&int64(248)) % 64)))
	goto L21
L25:
	;
	v54 = v8
	goto L21
L26:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v56, v32, v30-v32)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v64 <= v61+int32(1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L28
L30:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	if v88 <= v85+int32(1) {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	F_appendStringInfoChar(m, v60, int32(92))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v71 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v69+v61))) = uint8(v71)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v76 = v74 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v80 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v78+v76))) = uint8(v80)
	goto L30
L34:
	;
	goto L30
L35:
	;
	v108 = v30 + int32(1)
	v157 = v108
	v159 = v108
	goto L16
L36:
	;
	F_appendStringInfoChar(m, v84, base.I32_extend8_s(v54))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*uint8)(unsafe.Add(mBase, uint32(v93+v85))) = uint8(v54)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v99 = v97 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v101+v99))) = uint8(v103)
	goto L35
L39:
	;
	goto L35
L40:
	;
	if base.Ui32(v32) < base.Ui32(v30) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if int32(0) <= base.I32_extend8_s(v31) {
		goto L17
	} else {
		goto L52
	}
L43:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v116, v32, v30-v32)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	if v124 <= v121+int32(1) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L45
L47:
	;
	v157 = v30
	v159 = v30 + int32(1)
	goto L16
L48:
	;
	F_appendStringInfoChar(m, v120, int32(92))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v131 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v129+v121))) = uint8(v131)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v136 = v134 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v138+v136))) = uint8(v140)
	goto L47
L51:
	;
	goto L47
L52:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v150 = F_pg_encoding_mblen(m, v149, v30)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v157 = v32
	v159 = v150 + v30
	goto L16
L54:
	;
	goto L15
L55:
	;
	v168 = v163 & int32(255)
	if base.Ui32(v168) <= base.Ui32(int32(31)) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
	if v287 == int32(0) {
		v291 = v286
		v293 = v285
		goto L7
	} else {
		goto L93
	}
L58:
	;
	v172 = v163 - int32(8)
	if base.Ui32(int32(6)) <= base.Ui32(v172&int32(255)) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L60
L60:
	;
	if base.B2i32(v168 != int32(92))&base.B2i32(v168 != v8) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L61:
	;
	if base.Ui32(v164) < base.Ui32(v165) {
		goto L67
	} else {
		goto L68
	}
L62:
	;
	if v168 == v8 {
		v189 = v8
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v189 = base.I32_wrap_i64(int64(base.Ui64(int64(125784399180898)) >> (uint(base.I64_extend_i32_u(v172<<(uint(int32(3))%32))&int64(248)) % 64)))
	goto L61
L65:
	;
	v179 = v165 + int32(1)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v180 != 0 {
		v163 = v180
		v165 = v179
		goto L55
	} else {
		goto L66
	}
L66:
	;
	v291 = v179
	v293 = v164
	goto L7
L67:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v191, v164, v165-v164)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
	if v199 <= v196+int32(1) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L69
L71:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	if v223 <= v220+int32(1) {
		goto L77
	} else {
		goto L78
	}
L72:
	;
	F_appendStringInfoChar(m, v195, int32(92))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v206 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v204+v196))) = uint8(v206)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	v211 = v209 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+4)) = v211
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	v215 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v213+v211))) = uint8(v215)
	goto L71
L75:
	;
	goto L71
L76:
	;
	v243 = v165 + int32(1)
	v285 = v243
	v286 = v243
	goto L57
L77:
	;
	F_appendStringInfoChar(m, v219, base.I32_extend8_s(v189))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	*(*uint8)(unsafe.Add(mBase, uint32(v228+v220))) = uint8(v189)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v234 = v232 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+4)) = v234
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v236+v234))) = uint8(v238)
	goto L76
L80:
	;
	goto L76
L81:
	;
	if base.Ui32(v164) < base.Ui32(v165) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v285 = v164
	v286 = v165 + int32(1)
	goto L57
L84:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v251, v164, v165-v164)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	if v259 <= v256+int32(1) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L86
L88:
	;
	v285 = v165
	v286 = v165 + int32(1)
	goto L57
L89:
	;
	F_appendStringInfoChar(m, v255, int32(92))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v266 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v264+v256))) = uint8(v266)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	v271 = v269 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v271
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v275 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v273+v271))) = uint8(v275)
	goto L88
L92:
	;
	goto L88
L93:
	;
	v163 = v287
	v164 = v285
	v165 = v286
	goto L55
L94:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v297, v293, v291-v293)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	goto L6
}
func F_attribute_reloptions(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_build_reloptions(m, l0, l1, int32(64), int32(24), int32(759936), int32(2))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_get_attribute_options(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
	if v11 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(1572)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(51539607560)
		v24 = F_hash_create(m, int32(400353), int32(256), v6+int32(-48), int32(72))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[1155])) = v24
			v30 = *(*int32)(unsafe.Add(mBase, _consts[403]))
			if v30 == int32(0) {
				F_CreateCacheMemoryContext(m)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_CacheRegisterSyscacheCallback(m, int32(7), int32(1573), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
						v42 = v41
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
						v47 = int32(0)
						v49 = F_hash_search(m, v42, v6+int32(-48), v47, v47)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 != 0 {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
								v97 = v51
								v98 = v49
								if v97 == int32(0) {
									v113 = int32(0)
									m.G0 = v8 - int32(-64)
									return v113
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
									v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
										v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
										if v110 != 0 {
											v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
											mBase = m.M
										} else {
										}
										v113 = v105
										m.G0 = v8 - int32(-64)
										return v113
									}
								}
							} else {
								v52 = int32(0)
								v55 = F_SearchSysCache2(m, int32(7), l0, base.I32_extend16_s(l1))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									if v55 != 0 {
										v61 = F_SysCacheGetAttr(m, int32(7), v55, int32(23), v6+int32(-49))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
											if v63 == int32(0) {
												v67 = F_attribute_reloptions(m, v61, int32(0))
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													v70 = *(*int32)(unsafe.Add(mBase, _consts[403]))
													v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
													v74 = F_MemoryContextAlloc(m, v70, int32(base.Ui32(v71)>>(uint(int32(2))%32)))
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return int32(0)
													} else {
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
														v78 = int32(base.Ui32(v76) >> (uint(int32(2)) % 32))
														if v78 != 0 {
															v79 = F__emscripten_memcpy_bulkmem(m, v74, v67, v78)
															mBase = m.M
														} else {
														}
														v82 = v74
														F_ReleaseCatCache(m, v55)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															v86 = v82
															v88 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
															v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
																v97 = v86
																v98 = v93
																if v97 == int32(0) {
																	v113 = int32(0)
																	m.G0 = v8 - int32(-64)
																	return v113
																} else {
																	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
																	v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
																	mBase = m.M
																	v106 = m.ExcPending
																	if v106 != 0 {
																		return int32(0)
																	} else {
																		v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
																		v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
																		v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
																		if v110 != 0 {
																			v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
																			mBase = m.M
																		} else {
																		}
																		v113 = v105
																		m.G0 = v8 - int32(-64)
																		return v113
																	}
																}
															}
														}
													}
												}
											} else {
												v82 = v52
												F_ReleaseCatCache(m, v55)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													v86 = v82
													v88 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
													v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
														v97 = v86
														v98 = v93
														if v97 == int32(0) {
															v113 = int32(0)
															m.G0 = v8 - int32(-64)
															return v113
														} else {
															v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
															v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return int32(0)
															} else {
																v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
																v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
																v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
																if v110 != 0 {
																	v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
																	mBase = m.M
																} else {
																}
																v113 = v105
																m.G0 = v8 - int32(-64)
																return v113
															}
														}
													}
												}
											}
										}
									} else {
										v86 = v52
										v88 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
										v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
											v97 = v86
											v98 = v93
											if v97 == int32(0) {
												v113 = int32(0)
												m.G0 = v8 - int32(-64)
												return v113
											} else {
												v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
												v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
													v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
													v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
													if v110 != 0 {
														v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
														mBase = m.M
													} else {
													}
													v113 = v105
													m.G0 = v8 - int32(-64)
													return v113
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
				F_CacheRegisterSyscacheCallback(m, int32(7), int32(1573), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
					v42 = v41
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
					v47 = int32(0)
					v49 = F_hash_search(m, v42, v6+int32(-48), v47, v47)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						if v49 != 0 {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
							v97 = v51
							v98 = v49
							if v97 == int32(0) {
								v113 = int32(0)
								m.G0 = v8 - int32(-64)
								return v113
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
								v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
									v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
									if v110 != 0 {
										v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
										mBase = m.M
									} else {
									}
									v113 = v105
									m.G0 = v8 - int32(-64)
									return v113
								}
							}
						} else {
							v52 = int32(0)
							v55 = F_SearchSysCache2(m, int32(7), l0, base.I32_extend16_s(l1))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								if v55 != 0 {
									v61 = F_SysCacheGetAttr(m, int32(7), v55, int32(23), v6+int32(-49))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
										if v63 == int32(0) {
											v67 = F_attribute_reloptions(m, v61, int32(0))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												v70 = *(*int32)(unsafe.Add(mBase, _consts[403]))
												v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
												v74 = F_MemoryContextAlloc(m, v70, int32(base.Ui32(v71)>>(uint(int32(2))%32)))
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return int32(0)
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
													v78 = int32(base.Ui32(v76) >> (uint(int32(2)) % 32))
													if v78 != 0 {
														v79 = F__emscripten_memcpy_bulkmem(m, v74, v67, v78)
														mBase = m.M
													} else {
													}
													v82 = v74
													F_ReleaseCatCache(m, v55)
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														v86 = v82
														v88 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
														v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
															v97 = v86
															v98 = v93
															if v97 == int32(0) {
																v113 = int32(0)
																m.G0 = v8 - int32(-64)
																return v113
															} else {
																v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
																v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
																mBase = m.M
																v106 = m.ExcPending
																if v106 != 0 {
																	return int32(0)
																} else {
																	v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
																	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
																	v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
																	if v110 != 0 {
																		v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
																		mBase = m.M
																	} else {
																	}
																	v113 = v105
																	m.G0 = v8 - int32(-64)
																	return v113
																}
															}
														}
													}
												}
											}
										} else {
											v82 = v52
											F_ReleaseCatCache(m, v55)
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return int32(0)
											} else {
												v86 = v82
												v88 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
												v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
													v97 = v86
													v98 = v93
													if v97 == int32(0) {
														v113 = int32(0)
														m.G0 = v8 - int32(-64)
														return v113
													} else {
														v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
														v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
															v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
															v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
															if v110 != 0 {
																v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
																mBase = m.M
															} else {
															}
															v113 = v105
															m.G0 = v8 - int32(-64)
															return v113
														}
													}
												}
											}
										}
									}
								} else {
									v86 = v52
									v88 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
									v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
										v97 = v86
										v98 = v93
										if v97 == int32(0) {
											v113 = int32(0)
											m.G0 = v8 - int32(-64)
											return v113
										} else {
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
											v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
												v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
												v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
												if v110 != 0 {
													v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
													mBase = m.M
												} else {
												}
												v113 = v105
												m.G0 = v8 - int32(-64)
												return v113
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
		v42 = v11
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
		v47 = int32(0)
		v49 = F_hash_search(m, v42, v6+int32(-48), v47, v47)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			if v49 != 0 {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
				v97 = v51
				v98 = v49
				if v97 == int32(0) {
					v113 = int32(0)
					m.G0 = v8 - int32(-64)
					return v113
				} else {
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
					v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
						v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
						if v110 != 0 {
							v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
							mBase = m.M
						} else {
						}
						v113 = v105
						m.G0 = v8 - int32(-64)
						return v113
					}
				}
			} else {
				v52 = int32(0)
				v55 = F_SearchSysCache2(m, int32(7), l0, base.I32_extend16_s(l1))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					if v55 != 0 {
						v61 = F_SysCacheGetAttr(m, int32(7), v55, int32(23), v6+int32(-49))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							if v63 == int32(0) {
								v67 = F_attribute_reloptions(m, v61, int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, _consts[403]))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
									v74 = F_MemoryContextAlloc(m, v70, int32(base.Ui32(v71)>>(uint(int32(2))%32)))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
										v78 = int32(base.Ui32(v76) >> (uint(int32(2)) % 32))
										if v78 != 0 {
											v79 = F__emscripten_memcpy_bulkmem(m, v74, v67, v78)
											mBase = m.M
										} else {
										}
										v82 = v74
										F_ReleaseCatCache(m, v55)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											v86 = v82
											v88 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
											v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
												v97 = v86
												v98 = v93
												if v97 == int32(0) {
													v113 = int32(0)
													m.G0 = v8 - int32(-64)
													return v113
												} else {
													v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
													v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return int32(0)
													} else {
														v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
														v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
														v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
														if v110 != 0 {
															v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
															mBase = m.M
														} else {
														}
														v113 = v105
														m.G0 = v8 - int32(-64)
														return v113
													}
												}
											}
										}
									}
								}
							} else {
								v82 = v52
								F_ReleaseCatCache(m, v55)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									v86 = v82
									v88 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
									v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
										v97 = v86
										v98 = v93
										if v97 == int32(0) {
											v113 = int32(0)
											m.G0 = v8 - int32(-64)
											return v113
										} else {
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
											v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
												v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
												v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
												if v110 != 0 {
													v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
													mBase = m.M
												} else {
												}
												v113 = v105
												m.G0 = v8 - int32(-64)
												return v113
											}
										}
									}
								}
							}
						}
					} else {
						v86 = v52
						v88 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
						v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
							v97 = v86
							v98 = v93
							if v97 == int32(0) {
								v113 = int32(0)
								m.G0 = v8 - int32(-64)
								return v113
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
								v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
									v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
									if v110 != 0 {
										v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
										mBase = m.M
									} else {
									}
									v113 = v105
									m.G0 = v8 - int32(-64)
									return v113
								}
							}
						}
					}
				}
			}
		}
	}
}
