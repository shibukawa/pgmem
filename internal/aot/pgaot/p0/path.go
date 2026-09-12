package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_PathNameOpenFile(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, _consts[585]))
	v5 = F_PathNameOpenFilePerm(m, l0, l1, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_SearchPathMatchesCurrentEnvironment(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	F_recomputeNamespacePath(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int64)(unsafe.Add(mBase, _consts[217]))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 == v16 {
		v144 = int32(1)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v144
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[213]))
	if v19 != 0 {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v84 = int32(0)
	if v83 != v82 {
		v144 = v84
		goto L3
	} else {
		goto L26
	}
L6:
	;
	v72 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, _consts[214]))
	if v70 == v72 {
		v79 = v72
		v82 = v74
		v83 = v72
		goto L5
	} else {
		goto L25
	}
L7:
	;
	v66 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, _consts[214]))
	v79 = v66
	v82 = v68
	v83 = v66
	goto L5
L8:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v47 == int32(0) {
		v70 = v46
		goto L6
	} else {
		goto L21
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v34 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v32 != v34 {
		v144 = int32(0)
		goto L3
	} else {
		goto L17
	}
L10:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v21 == int32(0) {
		v46 = v20
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v26 = int32(0)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v27 != 0 {
		v144 = v26
		goto L3
	} else {
		goto L15
	}
L13:
	;
	if v20 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	return int32(0)
L15:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v28 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v144 = v26
	goto L3
L17:
	;
	v37 = v20 + int32(4)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if base.Ui32(v37) < base.Ui32(v20+v39<<(uint(int32(2))%32)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v44 = v37
	goto L20
L19:
	;
	v44 = int32(0)
	goto L20
L20:
	;
	v46 = v44
	goto L8
L21:
	;
	v50 = int32(0)
	if v46 == v50 {
		v144 = v50
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v53 != int32(11) {
		v144 = v50
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v57 = v46 + int32(4)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if base.Ui32(v57) < base.Ui32(v20+v58<<(uint(int32(2))%32)) {
		v70 = v57
		goto L6
	} else {
		goto L24
	}
L24:
	;
	goto L7
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v79 = v70
	v82 = v74
	v83 = v78
	goto L5
L26:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v86 == int32(0) {
		v132 = v79
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v132 != 0 {
		v144 = v84
		goto L3
	} else {
		goto L45
	}
L28:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v89 <= int32(0) {
		v132 = v79
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v92 = int32(0)
	if v92 < v89 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v96 = v89
	goto L32
L31:
	;
	v96 = v92
	goto L32
L32:
	;
	v98 = v79
	v99 = v92
	goto L33
L33:
	;
	if v98 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v132 = v127
	goto L27
L35:
	;
	return int32(0)
L36:
	;
	goto L37
L37:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110+v99<<(uint(int32(2))%32))))
	if v109 != v114 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return int32(0)
L39:
	;
	goto L40
L40:
	;
	v119 = v98 + int32(4)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if base.Ui32(v119) < base.Ui32(v121+v122<<(uint(int32(2))%32)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v127 = v119
	goto L43
L42:
	;
	v127 = int32(0)
	goto L43
L43:
	;
	v129 = v99 + int32(1)
	if v129 != v96 {
		v98 = v127
		v99 = v129
		goto L33
	} else {
		goto L44
	}
L44:
	;
	goto L34
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v15
	v144 = int32(1)
	goto L3
}
func F_create_merge_append_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 float64
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v101 int32
	_ = v101
	var v102 float64
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 float64
	_ = v114
	var v116 float64
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v192 float64
	_ = v192
	var v193 float64
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 float64
	_ = v198
	var v199 float64
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 float64
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 float64
	_ = v214
	var v217 int32
	_ = v217
	var v218 float64
	_ = v218
	var v224 float64
	_ = v224
	var v230 float64
	_ = v230
	var v231 float64
	_ = v231
	var v232 int32
	_ = v232
	var v233 float64
	_ = v233
	var v234 float64
	_ = v234
	var v235 int32
	_ = v235
	var v236 float64
	_ = v236
	var v237 float64
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 float64
	_ = v251
	var v253 float64
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 float64
	_ = v276
	var v278 float64
	_ = v278
	var v282 float64
	_ = v282
	var v284 float64
	_ = v284
	var v286 float64
	_ = v286
	var v288 float64
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 float64
	_ = v293
	var v295 float64
	_ = v295
	var v297 float64
	_ = v297
	var v299 float64
	_ = v299
	var v300 float64
	_ = v300
	v5 = int32(0)
	v9 = float64(0)
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v20 = F_palloc0(m, int32(88))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = int64(1438814044451)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v28 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)) = uint8(v28)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v27
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v28
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+21)) = uint8(v33)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v49 = base.B2i32(v40|v41 == v28)
	if v40 == v28 {
		v88 = v49
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v88 != 0 {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	goto L3
L5:
	;
	if v41 == int32(0) {
		v88 = v49
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v55 != v56 {
		v88 = int32(0)
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v58 = int32(1)
	if v55 <= v58 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v61 = v58
	goto L10
L9:
	;
	v61 = v55
	goto L10
L10:
	;
	v62 = int32(8)
	v67 = int32(0)
	goto L11
L11:
	;
	v75 = v67 << (uint(int32(2)) % 32)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v40+v62+v75)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+(v41+v62))))
	v80 = base.B2i32(v77 == v79)
	if v79 != v77 {
		v88 = v80
		goto L4
	} else {
		goto L13
	}
L12:
	;
	v88 = v80
	goto L4
L13:
	;
	v83 = v67 + int32(1)
	if v83 != v61 {
		v67 = v83
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v92 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	v93 = v92
	goto L17
L16:
	;
	v93 = float64(-1)
	goto L17
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v20)+80)) = v93
	if l2 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	m.G0 = v17 + int32(80)
	return v20
L19:
	;
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v20)+32))
	v284 = *(*float64)(unsafe.Add(mBase, _consts[378]))
	v286 = *(*float64)(unsafe.Add(mBase, _consts[382]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v275
	v288 = base.F64_add(v286, v286)
	v289 = int32(2)
	if v269 <= v289 {
		goto L59
	} else {
		goto L60
	}
L20:
	;
	v269 = int32(0)
	v275 = v5
	v276 = float64(0)
	v278 = v9
	goto L19
L21:
	;
	goto L22
L22:
	;
	v101 = int32(0)
	v102 = float64(0)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v101 < v103 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v111 = v101
	v113 = v5
	v114 = v102
	v116 = v9
	goto L26
L24:
	;
	v244 = v103
	v250 = v5
	v251 = v102
	v253 = v9
	goto L25
L25:
	;
	if v244 != int32(1) {
		v269 = v244
		v275 = v250
		v276 = v251
		v278 = v253
		goto L19
	} else {
		goto L55
	}
L26:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120+v111<<(uint(int32(2))%32))))
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v124)+32))
	v126 = *(*float64)(unsafe.Add(mBase, uint32(v20)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+32)) = base.F64_add(v125, v126)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+21)))
	if v129 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v244 = v241
	v250 = v238
	v251 = v236
	v253 = v237
	goto L25
L28:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+21)))
	v134 = v132
	goto L30
L29:
	;
	v134 = int32(0)
	goto L30
L30:
	;
	v136 = v134 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+21)) = uint8(v136)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v124)+64))
	if l3 == v138 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v236 = base.F64_add(v114, v233)
	v237 = base.F64_add(v116, v234)
	v238 = v235 + v113
	v240 = v111 + int32(1)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v240 < v241 {
		v111 = v240
		v113 = v238
		v114 = v236
		v116 = v237
		goto L26
	} else {
		goto L54
	}
L32:
	;
	if v191 != 0 {
		goto L50
	} else {
		goto L51
	}
L33:
	;
	v191 = int32(1)
	goto L32
L34:
	;
	goto L35
L35:
	;
	v147 = int32(0)
	goto L37
L36:
	;
	v191 = v183
	goto L32
L37:
	;
	v151 = int32(0)
	if l3 == v151 {
		v161 = v151
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v183 = int32(0)
	goto L36
L39:
	;
	if v138 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v155 <= v147 {
		v161 = int32(0)
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v161 = v157 + v147<<(uint(int32(2))%32)
	goto L39
L42:
	;
	v167 = base.B2i32(v161 == int32(0))
	if v161 == int32(0) {
		v183 = v167
		goto L36
	} else {
		goto L47
	}
L43:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v147 < v162 {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v191 = base.B2i32(v161 == int32(0))
	goto L32
L46:
	;
	goto L45
L47:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	v173 = v170 + v147<<(uint(int32(2))%32)
	if v173 == int32(0) {
		v183 = v167
		goto L36
	} else {
		goto L48
	}
L48:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	if v178 == v179 {
		v147 = v147 + int32(1)
		goto L37
	} else {
		goto L49
	}
L49:
	;
	goto L38
L50:
	;
	v192 = *(*float64)(unsafe.Add(mBase, uint32(v124)+56))
	v193 = *(*float64)(unsafe.Add(mBase, uint32(v124)+48))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v124)+40))
	v233 = v192
	v234 = v193
	v235 = v194
	goto L31
L51:
	;
	goto L52
L52:
	;
	v195 = int32(8)
	v196 = v17 + v195
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v124)+40))
	v198 = *(*float64)(unsafe.Add(mBase, uint32(v124)+56))
	v199 = *(*float64)(unsafe.Add(mBase, uint32(v124)+32))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+32))
	v204 = *(*int32)(unsafe.Add(mBase, _consts[324]))
	v205 = *(*float64)(unsafe.Add(mBase, uint32(v20)+80))
	v207 = m.G0
	v208 = int32(16)
	v209 = v207 - v208
	m.G0 = v209
	F_cost_tuplesort(m, v209+v195, v209, v199, v201, float64(0), v204, v205)
	mBase = m.M
	v214 = *(*float64)(unsafe.Add(mBase, uint32(v209)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v196)+32)) = v199
	v217 = int32(*(*uint8)(unsafe.Add(mBase, _consts[252])))
	v218 = base.F64_add(v198, v214)
	*(*float64)(unsafe.Add(mBase, uint32(v196)+48)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v196)+40)) = v197 + (v217 ^ int32(1))
	v224 = *(*float64)(unsafe.Add(mBase, uint32(v209)))
	*(*float64)(unsafe.Add(mBase, uint32(v196)+56)) = base.F64_add(v218, v224)
	m.G0 = v209 + v208
	goto L53
L53:
	;
	v230 = *(*float64)(unsafe.Add(mBase, uint32(v17)+64))
	v231 = *(*float64)(unsafe.Add(mBase, uint32(v17)+56))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	v233 = v230
	v234 = v231
	v235 = v232
	goto L31
L54:
	;
	goto L27
L55:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+20)))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)))
	if v261 != v262 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v269 = int32(1)
	v275 = v250
	v276 = v251
	v278 = v253
	goto L19
L57:
	;
	goto L58
L58:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v20)+56)) = v251
	*(*float64)(unsafe.Add(mBase, uint32(v20)+48)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v250
	goto L18
L59:
	;
	v292 = v289
	goto L61
L60:
	;
	v292 = v269
	goto L61
L61:
	;
	v293 = base.F64_convert_i32_u(v292)
	v295 = F_log(m, v293)
	mBase = m.M
	v297 = base.F64_div(v295, float64(0.693147180559945))
	v299 = float64(0)
	v300 = base.F64_add(base.F64_mul(base.F64_mul(v288, v293), v297), v299)
	*(*float64)(unsafe.Add(mBase, uint32(v20)+48)) = base.F64_add(v278, v300)
	*(*float64)(unsafe.Add(mBase, uint32(v20)+56)) = base.F64_add(v276, base.F64_add(v300, base.F64_add(base.F64_mul(base.F64_mul(v284, float64(0.5)), v282), base.F64_add(base.F64_mul(base.F64_mul(v282, v288), v297), v299))))
	goto L18
}
func F_make_path_rowexpr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	v10 = F_palloc0(m, int32(24))
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(8589936841)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(36)
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v10
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v22 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v29 = int32(0)
	goto L6
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v29<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v44 = int32(0)
	goto L8
L8:
	;
	if v39 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v44 = v44 + int32(1)
	goto L8
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v51 = v49
	goto L13
L12:
	;
	v51 = int32(0)
	goto L13
L13:
	;
	if v44 < v51 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v54 = v44 << (uint(int32(2)) % 32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54+v55)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v62 == int32(0) {
		v81 = v61
		v82 = v62
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v116 = v29 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v116 < v117 {
		v29 = v116
		goto L6
	} else {
		goto L30
	}
L17:
	;
	if v82-v81 != 0 {
		goto L10
	} else {
		goto L25
	}
L18:
	;
	goto L17
L19:
	;
	if v61 != v62 {
		v81 = v61
		v82 = v62
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v66 = v38
	v67 = v58
	goto L21
L21:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v71 == int32(0) {
		v81 = v70
		v82 = v71
		goto L18
	} else {
		goto L23
	}
L22:
	;
	v81 = v70
	v82 = v71
	goto L18
L23:
	;
	v74 = int32(1)
	if v70 == v71 {
		v66 = v66 + v74
		v67 = v67 + v74
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v84 = int32(1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v89+v54)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v93+v54)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97+v54)))
	v101 = F_makeVar(m, v84, base.I32_extend16_s(v44+v84), v91, v95, v99, int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v104 = F_lappend(m, v103, v101)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v104
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v108 = F_makeString(m, v38)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v110 = F_lappend(m, v107, v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v110
	goto L16
L30:
	;
	goto L3
}
func F_path_close(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_copy(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(1)
		return v3
	}
}
func F_path_div_pt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_copy(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if int32(0) < v11 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	return v7
L6:
	;
	v25 = v7 + int32(16) + v18<<(uint(int32(4))%32)
	F_point_div_point(m, v25, v25, v14)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v29 = v18 + int32(1)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v29 < v30 {
		v18 = v29
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
func F_path_encode(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 float64
	_ = v81
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	F_initStringInfo(m, v11+int32(32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	switch l0 - int32(1) {
	case 0:
		goto L5
	case 1:
		v23 = int32(40)
		goto L4
	default:
		goto L3
	}
L3:
	;
	if l1 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_appendStringInfoChar(m, v11+int32(32), v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v23 = int32(91)
	goto L4
L6:
	;
	goto L3
L7:
	;
	switch l0 - int32(1) {
	case 0:
		goto L30
	case 1:
		v120 = int32(41)
		goto L29
	default:
		goto L28
	}
L8:
	;
	F_appendStringInfoChar(m, v11+int32(32), int32(40))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v37 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v38 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v39 = F_float8out_internal(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v41 = F_float8out_internal(m, v37)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v39
	F_appendStringInfo(m, v11+int32(32), int32(186230), v11+int32(16))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_pfree(m, v39)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_pfree(m, v41)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_appendStringInfoChar(m, v11+int32(32), int32(41))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if l1 == int32(1) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v65 = l2
	v69 = int32(1)
	goto L17
L17:
	;
	F_appendStringInfoChar(m, v11+int32(32), int32(44))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L7
L19:
	;
	F_appendStringInfoChar(m, v11+int32(32), int32(40))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v81 = *(*float64)(unsafe.Add(mBase, uint32(v65)+24))
	v83 = v65 + int32(16)
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v83)))
	v85 = F_float8out_internal(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v87 = F_float8out_internal(m, v81)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v85
	F_appendStringInfo(m, v11+int32(32), int32(186230), v11)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_pfree(m, v85)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_pfree(m, v87)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_appendStringInfoChar(m, v11+int32(32), int32(41))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v106 = v69 + int32(1)
	if v106 != l1 {
		v65 = v83
		v69 = v106
		goto L17
	} else {
		goto L27
	}
L27:
	;
	goto L18
L28:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	m.G0 = v11 + int32(48)
	return v126
L29:
	;
	F_appendStringInfoChar(m, v11+int32(32), v120)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	v120 = int32(93)
	goto L29
L31:
	;
	goto L28
}
func F_path_length(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v33 float64
	_ = v33
	var v40 float64
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 float64
	_ = v50
	var v55 int32
	_ = v55
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 float64
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v13 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v17 = F_Float8GetDatum(m, float64(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v21 = v9 + int32(16)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v22 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	return v17
L7:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L27
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if int32(2) <= v42 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v40 = float64(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v30 = F_point_dt(m, v13<<(uint(int32(4))%32)+v21-int32(16), v21)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v33 = base.F64_add(v30, float64(0))
	if base.F64_ne(base.F64_abs(v33), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v40 = v33
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if base.F64_ne(base.F64_abs(v30), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v40 = v33
	goto L8
L15:
	;
	v46 = int32(1)
	v50 = v40
	goto L18
L16:
	;
	v78 = v40
	goto L17
L17:
	;
	v81 = F_Float8GetDatum(m, v78)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L26
	}
L18:
	;
	v55 = v21 + v46<<(uint(int32(4))%32)
	v58 = F_point_dt(m, v55-int32(16), v55)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v78 = v60
	goto L17
L20:
	;
	v71 = v46 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v71 < v72 {
		v46 = v71
		v50 = v60
		goto L18
	} else {
		goto L25
	}
L21:
	;
	v60 = base.F64_add(v50, v58)
	if base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if base.F64_eq(base.F64_abs(v50), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	if base.F64_ne(base.F64_abs(v58), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	goto L19
L26:
	;
	return v81
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_path_poly(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v41 int32
	_ = v41
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 float64
	_ = v58
	var v60 float64
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v89 int32
	_ = v89
	var v90 float64
	_ = v90
	var v95 int32
	_ = v95
	var v99 float64
	_ = v99
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v109 float64
	_ = v109
	var v111 int64
	_ = v111
	var v115 float64
	_ = v115
	var v120 int32
	_ = v120
	var v124 float64
	_ = v124
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v133 float64
	_ = v133
	var v135 float64
	_ = v135
	var v141 float64
	_ = v141
	var v143 int32
	_ = v143
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v155 float64
	_ = v155
	var v156 float64
	_ = v156
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
		if v19 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v24 = v20<<(uint(int32(4))%32) + int32(40)
			v25 = F_palloc(m, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v25))) = v24 << (uint(int32(2)) % 32)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v30
				v32 = int32(0)
				if v30 <= v32 {
					v35 = *(*float64)(unsafe.Add(mBase, uint32(v25)+40))
					v36 = *(*float64)(unsafe.Add(mBase, uint32(v25)+48))
					v153 = v35
					v154 = v36
					v155 = v35
					v156 = v36
				} else {
					v41 = v32
					for {
						v55 = v41 << (uint(int32(4)) % 32)
						v56 = v25 + int32(40) + v55
						v57 = v55 + (v15 + int32(16))
						v58 = *(*float64)(unsafe.Add(mBase, uint32(v57)))
						*(*float64)(unsafe.Add(mBase, uint32(v56))) = v58
						v60 = *(*float64)(unsafe.Add(mBase, uint32(v57)+8))
						*(*float64)(unsafe.Add(mBase, uint32(v56)+8)) = v60
						v63 = v41 + int32(1)
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
						if v63 < v64 {
							v41 = v63
							continue
						} else {
							break
						}
						break
					}
					v66 = *(*float64)(unsafe.Add(mBase, uint32(v25)+48))
					v67 = *(*float64)(unsafe.Add(mBase, uint32(v25)+40))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
					if v68 < int32(2) {
						v153 = v67
						v154 = v66
						v155 = v67
						v156 = v66
					} else {
						v74 = int32(1)
						v82 = v67
						v83 = v66
						v84 = v67
						v85 = v66
						for {
							v89 = v25 + int32(40) + v74<<(uint(int32(4))%32)
							v90 = *(*float64)(unsafe.Add(mBase, uint32(v89)))
							v95 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v90)&int64(9223372036854775807)))
							if v95 == int32(0) {
								if base.F64_lt(v90, v84) != 0 {
									v99 = v90
								} else {
									v99 = v84
								}
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v84)&int64(9223372036854775807)) {
									v105 = v90
								} else {
									v105 = v99
								}
								v106 = v105
							} else {
								v106 = v84
							}
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v90)&int64(9223372036854775807)) {
								v107 = v90
							} else {
								v107 = v82
							}
							if base.F64_gt(v90, v82) != 0 {
								v109 = v90
							} else {
								v109 = v107
							}
							v111 = int64(9223372036854775807)
							v115 = *(*float64)(unsafe.Add(mBase, uint32(v89)+8))
							v120 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v115)&v111))
							if v120 == int32(0) {
								if base.F64_lt(v115, v85) != 0 {
									v124 = v115
								} else {
									v124 = v85
								}
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v85)&int64(9223372036854775807)) {
									v130 = v115
								} else {
									v130 = v124
								}
								v131 = v130
							} else {
								v131 = v85
							}
							if base.Ui64(base.I64_reinterpret_f64(v82)&v111) < base.Ui64(int64(9218868437227405313)) {
								v132 = v109
							} else {
								v132 = v82
							}
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v115)&v111) {
								v133 = v115
							} else {
								v133 = v83
							}
							if base.F64_gt(v115, v83) != 0 {
								v135 = v115
							} else {
								v135 = v133
							}
							if base.Ui64(base.I64_reinterpret_f64(v83)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v141 = v135
							} else {
								v141 = v83
							}
							v143 = v74 + int32(1)
							if v143 != v68 {
								v74 = v143
								v82 = v132
								v83 = v141
								v84 = v106
								v85 = v131
								continue
							} else {
								break
							}
							break
						}
						v153 = v132
						v154 = v141
						v155 = v106
						v156 = v131
					}
				}
				*(*float64)(unsafe.Add(mBase, uint32(v25)+32)) = v156
				*(*float64)(unsafe.Add(mBase, uint32(v25)+8)) = v153
				*(*float64)(unsafe.Add(mBase, uint32(v25)+24)) = v155
				*(*float64)(unsafe.Add(mBase, uint32(v25)+16)) = v154
				return v25
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v166 = m.ExcPending
			if v166 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v169 = m.ExcPending
				if v169 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(284068), int32(0))
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(513242), int32(4463), int32(19755))
						mBase = m.M
						v178 = m.ExcPending
						if v178 != 0 {
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
