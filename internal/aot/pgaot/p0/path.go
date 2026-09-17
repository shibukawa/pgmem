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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFile[0]))
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
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	v2 = int32(0)
	F_recomputeNamespacePath(m)
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
	v16 = *(*int64)(unsafe.Add(mBase, _c_F_SearchPathMatchesCurrentEnvironment[0]))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v16 == v17 {
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
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_SearchPathMatchesCurrentEnvironment[1]))
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v82 = int32(0)
	if v81 != v78 {
		v144 = v82
		goto L3
	} else {
		goto L29
	}
L6:
	;
	v73 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_SearchPathMatchesCurrentEnvironment[2]))
	v78 = v75
	v79 = v73
	v81 = v73
	goto L5
L7:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v44 != 0 {
		goto L22
	} else {
		goto L23
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_SearchPathMatchesCurrentEnvironment[3]))
	if v31 != v33 {
		v144 = int32(0)
		goto L3
	} else {
		goto L16
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v22 == int32(0) {
		v43 = v21
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v27 = int32(0)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v28 != 0 {
		v144 = v27
		goto L3
	} else {
		goto L14
	}
L12:
	;
	if v21 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	return int32(0)
L14:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v29 != 0 {
		v144 = v27
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L6
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if int32(1) < v38 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v41 = v21 + int32(4)
	goto L19
L18:
	;
	v41 = int32(0)
	goto L19
L19:
	;
	v43 = v41
	goto L7
L20:
	;
	v78 = v61
	v79 = v2
	v81 = int32(0)
	goto L5
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v78 = v64
	v79 = v65
	v81 = v67
	goto L5
L22:
	;
	v45 = int32(0)
	if v43 == v45 {
		v144 = v45
		goto L3
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_SearchPathMatchesCurrentEnvironment[2]))
	if v43 == int32(0) {
		goto L20
	} else {
		goto L28
	}
L25:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v48 != int32(11) {
		v144 = v45
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v52 = v43 + int32(4)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if base.Ui32(v21+v53<<(uint(int32(2))%32)) <= base.Ui32(v52) {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_SearchPathMatchesCurrentEnvironment[2]))
	v64 = v59
	v65 = v52
	goto L21
L28:
	;
	v64 = v61
	v65 = v43
	goto L21
L29:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v84 == int32(0) {
		v132 = v79
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v132 != 0 {
		v144 = v82
		goto L3
	} else {
		goto L48
	}
L31:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v87 <= int32(0) {
		v132 = v79
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v90 = int32(0)
	if v90 < v87 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v93 = v87
	goto L35
L34:
	;
	v93 = v90
	goto L35
L35:
	;
	v97 = v79
	v100 = v2
	goto L36
L36:
	;
	if v97 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v132 = v125
	goto L30
L38:
	;
	return int32(0)
L39:
	;
	goto L40
L40:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v100<<(uint(int32(2))%32))))
	if v107 != v112 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	return int32(0)
L42:
	;
	goto L43
L43:
	;
	v117 = v97 + int32(4)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if base.Ui32(v117) < base.Ui32(v119+v120<<(uint(int32(2))%32)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v125 = v117
	goto L46
L45:
	;
	v125 = int32(0)
	goto L46
L46:
	;
	v127 = v100 + int32(1)
	if v127 != v93 {
		v97 = v125
		v100 = v127
		goto L36
	} else {
		goto L47
	}
L47:
	;
	goto L37
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v16
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 float64
	_ = v93
	var v94 float64
	_ = v94
	var v102 float64
	_ = v102
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 float64
	_ = v115
	var v118 float64
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 float64
	_ = v126
	var v127 float64
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 float64
	_ = v197
	var v198 float64
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 float64
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 float64
	_ = v213
	var v216 int32
	_ = v216
	var v217 float64
	_ = v217
	var v223 float64
	_ = v223
	var v229 float64
	_ = v229
	var v230 float64
	_ = v230
	var v231 int32
	_ = v231
	var v232 float64
	_ = v232
	var v233 float64
	_ = v233
	var v234 int32
	_ = v234
	var v235 float64
	_ = v235
	var v236 float64
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v250 float64
	_ = v250
	var v253 float64
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 float64
	_ = v275
	var v278 float64
	_ = v278
	var v281 float64
	_ = v281
	var v283 float64
	_ = v283
	var v285 float64
	_ = v285
	var v287 float64
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 float64
	_ = v292
	var v294 float64
	_ = v294
	var v296 float64
	_ = v296
	var v298 float64
	_ = v298
	var v299 float64
	_ = v299
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
	if base.B2i32(v40 == v28)|base.B2i32(v41 == v28) != 0 {
		v88 = base.B2i32(v40|v41 == v28)
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v88 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	goto L3
L5:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v56 != v57 {
		v88 = int32(0)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v59 = int32(1)
	if v56 <= v59 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v62 = v59
	goto L9
L8:
	;
	v62 = v56
	goto L9
L9:
	;
	v63 = int32(8)
	v68 = int32(0)
	goto L10
L10:
	;
	v76 = v68 << (uint(int32(2)) % 32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v40+v63+v76)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v41+v63+v76)))
	v81 = base.B2i32(v78 == v80)
	if v78 != v80 {
		v88 = v81
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v88 = v81
	goto L4
L12:
	;
	v84 = v68 + int32(1)
	if v84 != v62 {
		v68 = v84
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v93 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	v94 = v93
	goto L16
L15:
	;
	v94 = float64(-1)
	goto L16
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v20)+80)) = v94
	if l2 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	m.G0 = v17 + int32(80)
	return v20
L18:
	;
	v281 = *(*float64)(unsafe.Add(mBase, uint32(v20)+32))
	v283 = *(*float64)(unsafe.Add(mBase, _c_F_create_merge_append_path[0]))
	v285 = *(*float64)(unsafe.Add(mBase, _c_F_create_merge_append_path[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v274
	v287 = base.F64_add(v285, v285)
	v288 = int32(2)
	if v268 <= v288 {
		goto L58
	} else {
		goto L59
	}
L19:
	;
	v268 = int32(0)
	v274 = v5
	v275 = float64(0)
	v278 = v9
	goto L18
L20:
	;
	goto L21
L21:
	;
	v102 = float64(0)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v103 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v112 = int32(0)
	v114 = v5
	v115 = v102
	v118 = v9
	goto L25
L23:
	;
	v243 = v103
	v249 = v5
	v250 = v102
	v253 = v9
	goto L24
L24:
	;
	if v243 != int32(1) {
		v268 = v243
		v274 = v249
		v275 = v250
		v278 = v253
		goto L18
	} else {
		goto L54
	}
L25:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v112<<(uint(int32(2))%32))))
	v126 = *(*float64)(unsafe.Add(mBase, uint32(v125)+32))
	v127 = *(*float64)(unsafe.Add(mBase, uint32(v20)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+32)) = base.F64_add(v126, v127)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+21)))
	if v130 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v243 = v240
	v249 = v237
	v250 = v235
	v253 = v236
	goto L24
L27:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+21)))
	v133 = v131
	goto L29
L28:
	;
	v133 = int32(0)
	goto L29
L29:
	;
	v135 = v133 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+21)) = uint8(v135)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v125)+64))
	if l3 == v137 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v235 = base.F64_add(v115, v233)
	v236 = base.F64_add(v118, v232)
	v237 = v234 + v114
	v239 = v112 + int32(1)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v239 < v240 {
		v112 = v239
		v114 = v237
		v115 = v235
		v118 = v236
		goto L25
	} else {
		goto L53
	}
L31:
	;
	if v190 != 0 {
		goto L49
	} else {
		goto L50
	}
L32:
	;
	v190 = int32(1)
	goto L31
L33:
	;
	goto L34
L34:
	;
	v146 = int32(0)
	goto L36
L35:
	;
	v190 = v182
	goto L31
L36:
	;
	v150 = int32(0)
	if l3 == v150 {
		v160 = v150
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v182 = int32(0)
	goto L35
L38:
	;
	if v137 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v154 <= v146 {
		v160 = int32(0)
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v160 = v156 + v146<<(uint(int32(2))%32)
	goto L38
L41:
	;
	v166 = base.B2i32(v160 == int32(0))
	if v160 == int32(0) {
		v182 = v166
		goto L35
	} else {
		goto L46
	}
L42:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v146 < v161 {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v190 = base.B2i32(v160 == int32(0))
	goto L31
L45:
	;
	goto L44
L46:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	if v169 == int32(0) {
		v182 = v166
		goto L35
	} else {
		goto L47
	}
L47:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v146<<(uint(int32(2))%32)+v169)))
	if v176 == v178 {
		v146 = v146 + int32(1)
		goto L36
	} else {
		goto L48
	}
L48:
	;
	goto L37
L49:
	;
	v191 = *(*float64)(unsafe.Add(mBase, uint32(v125)+56))
	v192 = *(*float64)(unsafe.Add(mBase, uint32(v125)+48))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v125)+40))
	v232 = v192
	v233 = v191
	v234 = v193
	goto L30
L50:
	;
	goto L51
L51:
	;
	v194 = int32(8)
	v195 = v17 + v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v125)+40))
	v197 = *(*float64)(unsafe.Add(mBase, uint32(v125)+56))
	v198 = *(*float64)(unsafe.Add(mBase, uint32(v125)+32))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+32))
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_create_merge_append_path[2]))
	v204 = *(*float64)(unsafe.Add(mBase, uint32(v20)+80))
	v206 = m.G0
	v207 = int32(16)
	v208 = v206 - v207
	m.G0 = v208
	F_cost_tuplesort(m, v208+v194, v208, v198, v200, float64(0), v203, v204)
	mBase = m.M
	v213 = *(*float64)(unsafe.Add(mBase, uint32(v208)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v195)+32)) = v198
	v216 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_merge_append_path[3])))
	v217 = base.F64_add(v197, v213)
	*(*float64)(unsafe.Add(mBase, uint32(v195)+48)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v195)+40)) = v196 + (v216 ^ int32(1))
	v223 = *(*float64)(unsafe.Add(mBase, uint32(v208)))
	*(*float64)(unsafe.Add(mBase, uint32(v195)+56)) = base.F64_add(v217, v223)
	m.G0 = v208 + v207
	goto L52
L52:
	;
	v229 = *(*float64)(unsafe.Add(mBase, uint32(v17)+64))
	v230 = *(*float64)(unsafe.Add(mBase, uint32(v17)+56))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	v232 = v230
	v233 = v229
	v234 = v231
	goto L30
L53:
	;
	goto L26
L54:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+20)))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)))
	if v260 != v261 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v268 = int32(1)
	v274 = v249
	v275 = v250
	v278 = v253
	goto L18
L56:
	;
	goto L57
L57:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v20)+56)) = v250
	*(*float64)(unsafe.Add(mBase, uint32(v20)+48)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v249
	goto L17
L58:
	;
	v291 = v288
	goto L60
L59:
	;
	v291 = v268
	goto L60
L60:
	;
	v292 = base.F64_convert_i32_u(v291)
	v294 = F_log(m, v292)
	mBase = m.M
	v296 = base.F64_div(v294, float64(0.693147180559945))
	v298 = float64(0)
	v299 = base.F64_add(base.F64_mul(base.F64_mul(v287, v292), v296), v298)
	*(*float64)(unsafe.Add(mBase, uint32(v20)+48)) = base.F64_add(v278, v299)
	*(*float64)(unsafe.Add(mBase, uint32(v20)+56)) = base.F64_add(v275, base.F64_add(v299, base.F64_add(base.F64_mul(base.F64_mul(v283, float64(0.5)), v281), base.F64_add(base.F64_mul(base.F64_mul(v281, v287), v296), v298))))
	goto L17
}
func F_make_path_rowexpr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	v11 = F_palloc0(m, int32(24))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(8589936841)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(36)
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v11
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v23 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v31 = int32(0)
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v35 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	v137 = v31 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v137 < v138 {
		v31 = v137
		goto L6
	} else {
		goto L31
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v38 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v31<<(uint(int32(2))%32))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v47 = int32(0)
	if v47 < v38 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v50 = v38
	goto L13
L12:
	;
	v50 = v47
	goto L13
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v55 = int32(0)
	goto L14
L14:
	;
	v63 = v55 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v51+v63)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if base.B2i32(v69 == int32(0))|base.B2i32(v69 != v72) != 0 {
		v90 = v69
		v91 = v72
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L8
L16:
	;
	if v90-v91 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	v75 = v46
	v76 = v66
	goto L19
L19:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	if v80 == int32(0) {
		v90 = v80
		v91 = v79
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v90 = v80
	v91 = v79
	goto L17
L21:
	;
	v83 = int32(1)
	if v80 == v79 {
		v75 = v75 + v83
		v76 = v76 + v83
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v95 = int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100+v63)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v104+v63)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v108+v63)))
	v112 = F_makeVar(m, v95, base.I32_extend16_s(v55+v95), v102, v106, v110, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v125 = v55 + int32(1)
	if v125 != v50 {
		v55 = v125
		goto L14
	} else {
		goto L30
	}
L26:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v115 = F_lappend(m, v114, v112)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v115
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v119 = F_makeString(m, v46)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v121 = F_lappend(m, v118, v119)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v121
	goto L8
L30:
	;
	goto L15
L31:
	;
	goto L7
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	F_initStringInfo(m, v12+int32(32))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
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
		v24 = int32(40)
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
	F_appendStringInfoChar(m, v12+int32(32), v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v24 = int32(91)
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
		v113 = int32(41)
		goto L29
	default:
		goto L28
	}
L8:
	;
	v34 = v12 + int32(32)
	F_appendStringInfoChar(m, v34, int32(40))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v38 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v39 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v40 = F_float8out_internal(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v42 = F_float8out_internal(m, v38)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v40
	F_appendStringInfo(m, v34, int32(_a_F_path_encode_0), v12+int32(16))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_pfree(m, v40)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_pfree(m, v42)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_appendStringInfoChar(m, v34, int32(41))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
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
	v62 = l2
	v65 = int32(1)
	goto L17
L17:
	;
	v70 = v12 + int32(32)
	F_appendStringInfoChar(m, v70, int32(44))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L7
L19:
	;
	F_appendStringInfoChar(m, v70, int32(40))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v77 = *(*float64)(unsafe.Add(mBase, uint32(v62)+24))
	v78 = *(*float64)(unsafe.Add(mBase, uint32(v62)+16))
	v79 = F_float8out_internal(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v81 = F_float8out_internal(m, v77)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v79
	F_appendStringInfo(m, v70, int32(_a_F_path_encode_0), v12)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_pfree(m, v79)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_pfree(m, v81)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_appendStringInfoChar(m, v70, int32(41))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v98 = v65 + int32(1)
	if v98 != l1 {
		v62 = v62 + int32(16)
		v65 = v98
		goto L17
	} else {
		goto L27
	}
L27:
	;
	goto L18
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	m.G0 = v12 + int32(48)
	return v119
L29:
	;
	F_appendStringInfoChar(m, v12+int32(32), v113)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	v113 = int32(93)
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
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v31 float64
	_ = v31
	var v39 float64
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 float64
	_ = v49
	var v52 int32
	_ = v52
	var v55 float64
	_ = v55
	var v56 int32
	_ = v56
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 float64
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
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
		v39 = float64(0)
		goto L8
	} else {
		goto L9
	}
L6:
	;
	return v17
L7:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L22
	}
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if int32(2) <= v40 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v28 = F_point_dt(m, v9+v13<<(uint(int32(4))%32), v21)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v31 = base.F64_add(v28, float64(0))
	if base.F64_ne(base.F64_abs(v31), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v39 = v31
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v39 = v31
	goto L8
L13:
	;
	v44 = int32(1)
	v49 = v39
	goto L16
L14:
	;
	v80 = v39
	goto L15
L15:
	;
	v82 = F_Float8GetDatum(m, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L21
	}
L16:
	;
	v52 = v44 << (uint(int32(4)) % 32)
	v55 = F_point_dt(m, v9+v52, v21+v52)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	v80 = v57
	goto L15
L18:
	;
	v57 = base.F64_add(v49, v55)
	v59 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v57), v59)|base.F64_eq(base.F64_abs(v49), v59) == int32(0))&base.F64_ne(base.F64_abs(v55), v59) != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v72 = v44 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v72 < v73 {
		v44 = v72
		v49 = v57
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	return v82
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_path_poly(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v42 int32
	_ = v42
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v62 float64
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v75 int32
	_ = v75
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v97 int32
	_ = v97
	var v101 float64
	_ = v101
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v109 float64
	_ = v109
	var v114 int32
	_ = v114
	var v118 float64
	_ = v118
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v134 float64
	_ = v134
	var v135 float64
	_ = v135
	var v137 float64
	_ = v137
	var v143 float64
	_ = v143
	var v145 int32
	_ = v145
	var v155 float64
	_ = v155
	var v156 float64
	_ = v156
	var v157 float64
	_ = v157
	var v158 float64
	_ = v158
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
		if v20 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			v25 = v21<<(uint(int32(4))%32) + int32(40)
			v26 = F_palloc(m, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v26))) = v25 << (uint(int32(2)) % 32)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v31
				v33 = int32(0)
				if v31 <= v33 {
					v36 = *(*float64)(unsafe.Add(mBase, uint32(v26)+40))
					v37 = *(*float64)(unsafe.Add(mBase, uint32(v26)+48))
					v155 = v36
					v156 = v37
					v157 = v36
					v158 = v37
				} else {
					v42 = v33
					for {
						v57 = v42 << (uint(int32(4)) % 32)
						v58 = v26 + int32(40) + v57
						v59 = v57 + (v16 + int32(16))
						v60 = *(*float64)(unsafe.Add(mBase, uint32(v59)))
						*(*float64)(unsafe.Add(mBase, uint32(v58))) = v60
						v62 = *(*float64)(unsafe.Add(mBase, uint32(v59)+8))
						*(*float64)(unsafe.Add(mBase, uint32(v58)+8)) = v62
						v65 = v42 + int32(1)
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
						if v65 < v66 {
							v42 = v65
							continue
						} else {
							break
						}
						break
					}
					v68 = int32(1)
					v69 = *(*float64)(unsafe.Add(mBase, uint32(v26)+48))
					v70 = *(*float64)(unsafe.Add(mBase, uint32(v26)+40))
					if v31 == v68 {
						v155 = v70
						v156 = v69
						v157 = v70
						v158 = v69
					} else {
						v75 = v68
						v83 = v70
						v84 = v69
						v85 = v70
						v86 = v69
						for {
							v91 = v26 + int32(40) + v75<<(uint(int32(4))%32)
							v92 = *(*float64)(unsafe.Add(mBase, uint32(v91)))
							v97 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v92)&int64(9223372036854775807)))
							if v97 == int32(0) {
								if base.F64_gt(v85, v92) != 0 {
									v101 = v92
								} else {
									v101 = v85
								}
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v85)&int64(9223372036854775807)) {
									v107 = v92
								} else {
									v107 = v101
								}
								v108 = v107
							} else {
								v108 = v85
							}
							v109 = *(*float64)(unsafe.Add(mBase, uint32(v91)+8))
							v114 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v109)&int64(9223372036854775807)))
							if v114 == int32(0) {
								if base.F64_gt(v86, v109) != 0 {
									v118 = v109
								} else {
									v118 = v86
								}
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v86)&int64(9223372036854775807)) {
									v124 = v109
								} else {
									v124 = v118
								}
								v125 = v124
							} else {
								v125 = v86
							}
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v92)&int64(9223372036854775807)) {
								v126 = v92
							} else {
								v126 = v83
							}
							if base.F64_lt(v83, v92) != 0 {
								v128 = v92
							} else {
								v128 = v126
							}
							if base.Ui64(base.I64_reinterpret_f64(v83)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v134 = v128
							} else {
								v134 = v83
							}
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v109)&int64(9223372036854775807)) {
								v135 = v109
							} else {
								v135 = v84
							}
							if base.F64_lt(v84, v109) != 0 {
								v137 = v109
							} else {
								v137 = v135
							}
							if base.Ui64(base.I64_reinterpret_f64(v84)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v143 = v137
							} else {
								v143 = v84
							}
							v145 = v75 + int32(1)
							if v145 != v31 {
								v75 = v145
								v83 = v134
								v84 = v143
								v85 = v108
								v86 = v125
								continue
							} else {
								break
							}
							break
						}
						v155 = v134
						v156 = v143
						v157 = v108
						v158 = v125
					}
				}
				*(*float64)(unsafe.Add(mBase, uint32(v26)+32)) = v158
				*(*float64)(unsafe.Add(mBase, uint32(v26)+8)) = v155
				*(*float64)(unsafe.Add(mBase, uint32(v26)+24)) = v157
				*(*float64)(unsafe.Add(mBase, uint32(v26)+16)) = v156
				return v26
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v169 = m.ExcPending
			if v169 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v172 = m.ExcPending
				if v172 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_path_poly_0), int32(0))
					mBase = m.M
					v176 = m.ExcPending
					if v176 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_path_poly_1), int32(_a_F_path_poly_2), int32(_a_F_path_poly_3))
						mBase = m.M
						v181 = m.ExcPending
						if v181 != 0 {
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
