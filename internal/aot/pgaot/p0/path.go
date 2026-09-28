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
func F_create_merge_append_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 float64
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 float64
	_ = v120
	var v122 float64
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 float64
	_ = v233
	var v234 float64
	_ = v234
	var v235 float64
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 float64
	_ = v240
	var v242 int32
	_ = v242
	var v244 float64
	_ = v244
	var v245 float64
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 float64
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 float64
	_ = v260
	var v263 int32
	_ = v263
	var v264 float64
	_ = v264
	var v270 float64
	_ = v270
	var v277 int32
	_ = v277
	var v280 float64
	_ = v280
	var v281 float64
	_ = v281
	var v282 float64
	_ = v282
	var v283 float64
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v300 float64
	_ = v300
	var v302 float64
	_ = v302
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v327 float64
	_ = v327
	var v329 float64
	_ = v329
	var v333 float64
	_ = v333
	var v335 float64
	_ = v335
	var v337 float64
	_ = v337
	var v338 int32
	_ = v338
	var v339 int64
	_ = v339
	var v342 int32
	_ = v342
	var v343 int64
	_ = v343
	var v349 float64
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 float64
	_ = v354
	var v356 float64
	_ = v356
	var v358 float64
	_ = v358
	var v360 float64
	_ = v360
	var v361 float64
	_ = v361
	v6 = int32(0)
	v11 = float64(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	v22 = F_palloc0(m, int32(96))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = int64(1455993913638)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = l1
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v31 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)) = uint8(v31)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v30
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+21)) = uint8(v36)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if base.B2i32(v43 == v31)|base.B2i32(v44 == v31) != 0 {
		v91 = base.B2i32(v43|v44 == v31)
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v91 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	goto L3
L5:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v59 != v60 {
		v91 = int32(0)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v62 = int32(1)
	if v59 <= v62 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v65 = v62
	goto L9
L8:
	;
	v65 = v59
	goto L9
L9:
	;
	v66 = int32(8)
	v71 = int32(0)
	goto L10
L10:
	;
	v79 = v71 << (uint(int32(2)) % 32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v43+v66+v79)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v44+v66+v79)))
	v84 = base.B2i32(v81 == v83)
	if v81 != v83 {
		v91 = v84
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v91 = v84
	goto L4
L12:
	;
	v87 = v71 + int32(1)
	if v87 != v65 {
		v71 = v87
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v96 = *(*float64)(unsafe.Add(mBase, uint32(l0)+320))
	v97 = v96
	goto L16
L15:
	;
	v97 = float64(-1)
	goto L16
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v22)+80)) = v97
	if l2 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	m.G0 = v19 + int32(80)
	return v22
L18:
	;
	v333 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
	v335 = *(*float64)(unsafe.Add(mBase, _c_F_create_merge_append_path[0]))
	v337 = *(*float64)(unsafe.Add(mBase, _c_F_create_merge_append_path[1]))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v338)+32))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v342 != 0 {
		goto L76
	} else {
		goto L77
	}
L19:
	;
	v318 = int32(0)
	v324 = v6
	v327 = float64(0)
	v329 = v11
	goto L18
L20:
	;
	goto L21
L21:
	;
	v105 = float64(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v106 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v113 = int32(0)
	v117 = v6
	v120 = v105
	v122 = v11
	goto L25
L23:
	;
	v291 = v106
	v297 = v6
	v300 = v105
	v302 = v11
	goto L24
L24:
	;
	if v291 != int32(1) {
		v318 = v291
		v324 = v297
		v327 = v300
		v329 = v302
		goto L18
	} else {
		goto L72
	}
L25:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126+v113<<(uint(int32(2))%32))))
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v130)+32))
	v132 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v22)+32)) = base.F64_add(v131, v132)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+21)))
	if v135 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v291 = v288
	v297 = v285
	v300 = v283
	v302 = v281
	goto L24
L27:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+21)))
	v138 = v136
	goto L29
L28:
	;
	v138 = int32(0)
	goto L29
L29:
	;
	v140 = v138 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+21)) = uint8(v140)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)+64))
	v144 = v19 + int32(76)
	if l4 == v142 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	if v222 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L31:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v210
	v222 = int32(1)
	goto L30
L32:
	;
	if l4 != 0 {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if l4 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = int32(0)
	v222 = int32(1)
	goto L30
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = int32(0)
	v222 = int32(1)
	goto L30
L37:
	;
	goto L38
L38:
	;
	if v142 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v162 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v162
	v222 = v162
	goto L30
L40:
	;
	goto L41
L41:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	v166 = int32(0)
	if v166 < v165 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v169 = v165
	goto L44
L43:
	;
	v169 = v166
	goto L44
L44:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v175 = int32(0)
	goto L45
L45:
	;
	if v175 < v170 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v186 = v182 + v175<<(uint(int32(2))%32)
	goto L49
L48:
	;
	v186 = int32(0)
	goto L49
L49:
	;
	if v175 == v169 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v169
	v222 = base.B2i32(v186 == int32(0))
	goto L30
L51:
	;
	goto L52
L52:
	;
	v192 = base.B2i32(v186 == int32(0))
	if v186 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v175
	v222 = v192
	goto L30
L54:
	;
	goto L55
L55:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	if v196 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v175
	v222 = v192
	goto L30
L57:
	;
	goto L58
L58:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v196+v175<<(uint(int32(2))%32))))
	if v200 != v204 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v175
	v222 = int32(0)
	goto L30
L60:
	;
	v175 = v175 + int32(1)
	goto L45
L62:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v130)+40))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_merge_append_path[2])))
	if v227 != int32(1) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v277 = v130
	goto L64
L64:
	;
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v277)+56))
	v281 = base.F64_add(v122, v280)
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v277)+48))
	v283 = base.F64_add(v120, v282)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v277)+40))
	v285 = v284 + v117
	v287 = v113 + int32(1)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v287 < v288 {
		v113 = v287
		v117 = v285
		v120 = v283
		v122 = v281
		goto L25
	} else {
		goto L71
	}
L65:
	;
	v277 = v19
	goto L64
L66:
	;
	v244 = *(*float64)(unsafe.Add(mBase, uint32(v130)+56))
	v245 = *(*float64)(unsafe.Add(mBase, uint32(v130)+32))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+32))
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_create_merge_append_path[3]))
	v251 = *(*float64)(unsafe.Add(mBase, uint32(v22)+80))
	v253 = m.G0
	v254 = int32(16)
	v255 = v253 - v254
	m.G0 = v255
	F_cost_tuplesort(m, v255+int32(8), v255, v245, v247, float64(0), v250, v251)
	mBase = m.M
	v260 = *(*float64)(unsafe.Add(mBase, uint32(v255)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+32)) = v245
	v263 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_merge_append_path[4])))
	v264 = base.F64_add(v244, v260)
	*(*float64)(unsafe.Add(mBase, uint32(v19)+48)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v225 + (v263 ^ int32(1))
	v270 = *(*float64)(unsafe.Add(mBase, uint32(v255)))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+56)) = base.F64_add(v264, v270)
	m.G0 = v255 + v254
	goto L70
L67:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v230 <= int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v233 = *(*float64)(unsafe.Add(mBase, uint32(v130)+48))
	v234 = *(*float64)(unsafe.Add(mBase, uint32(v130)+56))
	v235 = *(*float64)(unsafe.Add(mBase, uint32(v130)+32))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+32))
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_create_merge_append_path[3]))
	v240 = *(*float64)(unsafe.Add(mBase, uint32(v22)+80))
	F_cost_incremental_sort(m, v19, l0, l4, v230, v225, v233, v234, v235, v237, v239, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	goto L65
L70:
	;
	goto L65
L71:
	;
	goto L26
L72:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+20)))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)))
	if v310 != v311 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v318 = int32(1)
	v324 = v297
	v327 = v300
	v329 = v302
	goto L18
L74:
	;
	goto L75
L75:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v22)+56)) = v302
	*(*float64)(unsafe.Add(mBase, uint32(v22)+48)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v297
	goto L17
L76:
	;
	v343 = int64(-8193)
	goto L78
L77:
	;
	v343 = int64(-270337)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v324 + base.B2i32(v339|v343 != int64(-1))
	v349 = base.F64_add(v337, v337)
	v350 = int32(2)
	if v318 <= v350 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v353 = v350
	goto L81
L80:
	;
	v353 = v318
	goto L81
L81:
	;
	v354 = base.F64_convert_i32_u(v353)
	v356 = F_log(m, v354)
	mBase = m.M
	v358 = base.F64_div(v356, float64(0.693147180559945))
	v360 = float64(0)
	v361 = base.F64_add(base.F64_mul(base.F64_mul(v349, v354), v358), v360)
	*(*float64)(unsafe.Add(mBase, uint32(v22)+48)) = base.F64_add(v327, v361)
	*(*float64)(unsafe.Add(mBase, uint32(v22)+56)) = base.F64_add(v329, base.F64_add(v361, base.F64_add(base.F64_mul(base.F64_mul(v335, float64(0.5)), v333), base.F64_add(base.F64_mul(base.F64_mul(v333, v349), v358), v360))))
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
func F_path_close(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_copy(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(1)
		return base.I64_extend_i32_u(v3)
	}
}
func F_path_div_pt(m *base.Module, l0 int32) int64 {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	return base.I64_extend_i32_u(v7)
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
func F_path_length(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 float64
	_ = v26
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v30 float64
	_ = v30
	var v41 float64
	_ = v41
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v59 float64
	_ = v59
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v78 int32
	_ = v78
	var v79 float64
	_ = v79
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v85 float64
	_ = v85
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v101 float64
	_ = v101
	var v109 float64
	_ = v109
	var v114 float64
	_ = v114
	var v115 float64
	_ = v115
	var v116 float64
	_ = v116
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v137 float64
	_ = v137
	var v144 float64
	_ = v144
	var v146 float64
	_ = v146
	var v153 float64
	_ = v153
	var v154 int32
	_ = v154
	var v158 float64
	_ = v158
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 float64
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 float64
	_ = v178
	var v179 int32
	_ = v179
	var v180 float64
	_ = v180
	var v181 float64
	_ = v181
	var v183 float64
	_ = v183
	var v194 float64
	_ = v194
	var v195 int32
	_ = v195
	var v196 float64
	_ = v196
	var v197 float64
	_ = v197
	var v198 float64
	_ = v198
	var v199 float64
	_ = v199
	var v201 float64
	_ = v201
	var v212 float64
	_ = v212
	var v213 int32
	_ = v213
	var v214 float64
	_ = v214
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 float64
	_ = v227
	var v228 float64
	_ = v228
	var v231 int32
	_ = v231
	var v232 float64
	_ = v232
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v238 float64
	_ = v238
	var v241 int64
	_ = v241
	var v243 int64
	_ = v243
	var v254 float64
	_ = v254
	var v262 float64
	_ = v262
	var v267 float64
	_ = v267
	var v268 float64
	_ = v268
	var v269 float64
	_ = v269
	var v278 float64
	_ = v278
	var v279 float64
	_ = v279
	var v281 float64
	_ = v281
	var v283 float64
	_ = v283
	var v290 float64
	_ = v290
	var v296 float64
	_ = v296
	var v307 float64
	_ = v307
	var v308 int32
	_ = v308
	var v309 float64
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 float64
	_ = v317
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v15 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int64(0)
L4:
	;
	goto L5
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v20 == int32(0) {
		v158 = float64(0)
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(2) <= v160 {
		goto L38
	} else {
		goto L39
	}
L7:
	;
	v25 = v11 + v15<<(uint(int32(4))%32)
	v26 = *(*float64)(unsafe.Add(mBase, uint32(v25)))
	v27 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
	v28 = base.F64_sub(v26, v27)
	v30 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v28), v30)|base.F64_eq(base.F64_abs(v26), v30)|base.F64_eq(base.F64_abs(v27), v30) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v43 = v28
	goto L10
L9:
	;
	v41 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v44 = *(*float64)(unsafe.Add(mBase, uint32(v25)+8))
	v45 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
	v46 = base.F64_sub(v44, v45)
	v48 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v46), v48)|base.F64_eq(base.F64_abs(v44), v48)|base.F64_eq(base.F64_abs(v45), v48) != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v43 = v41
	goto L10
L12:
	;
	v61 = v46
	goto L14
L13:
	;
	v59 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v70 = m.G0
	v72 = v70 - int32(32)
	m.G0 = v72
	v74 = base.F64_abs(v43)
	v75 = base.F64_abs(v61)
	v78 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v74)) < base.Ui64(base.I64_reinterpret_f64(v75)))
	if base.Ui64(base.I64_reinterpret_f64(v74)) < base.Ui64(base.I64_reinterpret_f64(v75)) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v61 = v59
	goto L14
L16:
	;
	v144 = base.F64_add(v137, float64(0))
	v146 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v144), v146)|base.F64_eq(base.F64_abs(v137), v146) != 0 {
		v158 = v144
		goto L6
	} else {
		goto L36
	}
L17:
	;
	m.G0 = v72 + int32(32)
	goto L16
L18:
	;
	v79 = v74
	goto L20
L19:
	;
	v79 = v75
	goto L20
L20:
	;
	v80 = base.I64_reinterpret_f64(v79)
	v82 = int64(base.Ui64(v80) >> (uint(int64(52)) % 64))
	if v82 == int64(2047) {
		v137 = v79
		goto L17
	} else {
		goto L21
	}
L21:
	;
	if base.Ui64(base.I64_reinterpret_f64(v74)) < base.Ui64(base.I64_reinterpret_f64(v75)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v85 = v75
	goto L24
L23:
	;
	v85 = v74
	goto L24
L24:
	;
	if v80 == int64(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v137 = v85
	goto L17
L26:
	;
	v88 = base.I64_reinterpret_f64(v85)
	v90 = int64(base.Ui64(v88) >> (uint(int64(52)) % 64))
	if v90 == int64(2047) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	if int32(65) <= base.I32_wrap_i64(v90)-base.I32_wrap_i64(v82) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v137 = base.F64_add(v74, v75)
	goto L17
L29:
	;
	goto L30
L30:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v88) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	F_sq(m, v72+int32(24), v72+int32(16), v114)
	mBase = m.M
	F_sq(m, v72+int32(8), v72, v115)
	mBase = m.M
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v72)))
	v126 = *(*float64)(unsafe.Add(mBase, uint32(v72)+16))
	v128 = *(*float64)(unsafe.Add(mBase, uint32(v72)+8))
	v130 = *(*float64)(unsafe.Add(mBase, uint32(v72)+24))
	v137 = base.F64_mul(v116, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v125, v126), v128), v130)))
	goto L17
L32:
	;
	v101 = float64(1.90109156629516e-211)
	v114 = base.F64_mul(v85, v101)
	v115 = base.F64_mul(v79, v101)
	v116 = float64(5.260135901548374e+210)
	goto L31
L33:
	;
	goto L34
L34:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v80) {
		v114 = v85
		v115 = v79
		v116 = float64(1)
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v109 = float64(5.260135901548374e+210)
	v114 = base.F64_mul(v85, v109)
	v115 = base.F64_mul(v79, v109)
	v116 = float64(1.90109156629516e-211)
	goto L31
L36:
	;
	v153 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v158 = v153
	goto L6
L38:
	;
	v166 = int32(1)
	v169 = v158
	goto L41
L39:
	;
	v317 = v158
	goto L40
L40:
	;
	return base.I64_reinterpret_f64(v317)
L41:
	;
	v176 = v166 << (uint(int32(4)) % 32)
	v177 = v11 + v176
	v178 = *(*float64)(unsafe.Add(mBase, uint32(v177)))
	v179 = v176 + (v11 + int32(16))
	v180 = *(*float64)(unsafe.Add(mBase, uint32(v179)))
	v181 = base.F64_sub(v178, v180)
	v183 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v181), v183)|base.F64_eq(base.F64_abs(v178), v183)|base.F64_eq(base.F64_abs(v180), v183) != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v317 = v309
	goto L40
L43:
	;
	v311 = v166 + int32(1)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v311 < v312 {
		v166 = v311
		v169 = v309
		goto L41
	} else {
		goto L78
	}
L44:
	;
	v196 = v181
	goto L46
L45:
	;
	v194 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v197 = *(*float64)(unsafe.Add(mBase, uint32(v177)+8))
	v198 = *(*float64)(unsafe.Add(mBase, uint32(v179)+8))
	v199 = base.F64_sub(v197, v198)
	v201 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v199), v201)|base.F64_eq(base.F64_abs(v197), v201)|base.F64_eq(base.F64_abs(v198), v201) != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v196 = v194
	goto L46
L48:
	;
	v214 = v199
	goto L50
L49:
	;
	v212 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	v223 = m.G0
	v225 = v223 - int32(32)
	m.G0 = v225
	v227 = base.F64_abs(v196)
	v228 = base.F64_abs(v214)
	v231 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v227)) < base.Ui64(base.I64_reinterpret_f64(v228)))
	if base.Ui64(base.I64_reinterpret_f64(v227)) < base.Ui64(base.I64_reinterpret_f64(v228)) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v214 = v212
	goto L50
L52:
	;
	v296 = base.F64_add(v169, v290)
	if base.F64_ne(base.F64_abs(v296), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L72
	} else {
		goto L73
	}
L53:
	;
	m.G0 = v225 + int32(32)
	goto L52
L54:
	;
	v232 = v227
	goto L56
L55:
	;
	v232 = v228
	goto L56
L56:
	;
	v233 = base.I64_reinterpret_f64(v232)
	v235 = int64(base.Ui64(v233) >> (uint(int64(52)) % 64))
	if v235 == int64(2047) {
		v290 = v232
		goto L53
	} else {
		goto L57
	}
L57:
	;
	if base.Ui64(base.I64_reinterpret_f64(v227)) < base.Ui64(base.I64_reinterpret_f64(v228)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v238 = v228
	goto L60
L59:
	;
	v238 = v227
	goto L60
L60:
	;
	if v233 == int64(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v290 = v238
	goto L53
L62:
	;
	v241 = base.I64_reinterpret_f64(v238)
	v243 = int64(base.Ui64(v241) >> (uint(int64(52)) % 64))
	if v243 == int64(2047) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	if int32(65) <= base.I32_wrap_i64(v243)-base.I32_wrap_i64(v235) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v290 = base.F64_add(v227, v228)
	goto L53
L65:
	;
	goto L66
L66:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v241) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	F_sq(m, v225+int32(24), v225+int32(16), v267)
	mBase = m.M
	F_sq(m, v225+int32(8), v225, v268)
	mBase = m.M
	v278 = *(*float64)(unsafe.Add(mBase, uint32(v225)))
	v279 = *(*float64)(unsafe.Add(mBase, uint32(v225)+16))
	v281 = *(*float64)(unsafe.Add(mBase, uint32(v225)+8))
	v283 = *(*float64)(unsafe.Add(mBase, uint32(v225)+24))
	v290 = base.F64_mul(v269, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v278, v279), v281), v283)))
	goto L53
L68:
	;
	v254 = float64(1.90109156629516e-211)
	v267 = base.F64_mul(v238, v254)
	v268 = base.F64_mul(v232, v254)
	v269 = float64(5.260135901548374e+210)
	goto L67
L69:
	;
	goto L70
L70:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v233) {
		v267 = v238
		v268 = v232
		v269 = float64(1)
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v262 = float64(5.260135901548374e+210)
	v267 = base.F64_mul(v238, v262)
	v268 = base.F64_mul(v232, v262)
	v269 = float64(1.90109156629516e-211)
	goto L67
L72:
	;
	v309 = v296
	goto L43
L73:
	;
	goto L74
L74:
	;
	if base.F64_eq(base.F64_abs(v169), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v309 = v296
		goto L43
	} else {
		goto L75
	}
L75:
	;
	if base.F64_eq(base.F64_abs(v290), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v309 = v296
		goto L43
	} else {
		goto L76
	}
L76:
	;
	v307 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v309 = v307
	goto L43
L78:
	;
	goto L42
}
func F_path_poly(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v64 int32
	_ = v64
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v85 float64
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v98 int32
	_ = v98
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v109 float64
	_ = v109
	var v115 int32
	_ = v115
	var v116 float64
	_ = v116
	var v121 int32
	_ = v121
	var v125 float64
	_ = v125
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v133 float64
	_ = v133
	var v138 int32
	_ = v138
	var v142 float64
	_ = v142
	var v148 float64
	_ = v148
	var v149 float64
	_ = v149
	var v150 float64
	_ = v150
	var v152 float64
	_ = v152
	var v158 float64
	_ = v158
	var v159 float64
	_ = v159
	var v161 float64
	_ = v161
	var v167 float64
	_ = v167
	var v169 int32
	_ = v169
	var v179 float64
	_ = v179
	var v180 float64
	_ = v180
	var v181 float64
	_ = v181
	var v182 float64
	_ = v182
	var v205 int64
	_ = v205
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
		if v21 == int32(0) {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v25 = F_errsave_start(m, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				if v25 == int32(0) {
					v205 = int64(0)
					return v205
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_path_poly_0), int32(0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							F_errsave_finish(m, v24, int32(_a_F_path_poly_1), int32(_a_F_path_poly_2), int32(_a_F_path_poly_3))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								return int64(0)
							}
						}
					}
				}
			}
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v47 = v43<<(uint(int32(4))%32) + int32(40)
			v48 = F_palloc(m, v47)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v48))) = v47 << (uint(int32(2)) % 32)
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v53
				v55 = int32(0)
				if v53 <= v55 {
					v58 = *(*float64)(unsafe.Add(mBase, uint32(v48)+40))
					v59 = *(*float64)(unsafe.Add(mBase, uint32(v48)+48))
					v179 = v58
					v180 = v59
					v181 = v58
					v182 = v59
				} else {
					v64 = v55
					for {
						v80 = v64 << (uint(int32(4)) % 32)
						v81 = v48 + int32(40) + v80
						v82 = v80 + (v17 + int32(16))
						v83 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
						*(*float64)(unsafe.Add(mBase, uint32(v81))) = v83
						v85 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
						*(*float64)(unsafe.Add(mBase, uint32(v81)+8)) = v85
						v88 = v64 + int32(1)
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
						if v88 < v89 {
							v64 = v88
							continue
						} else {
							break
						}
						break
					}
					v91 = int32(1)
					v92 = *(*float64)(unsafe.Add(mBase, uint32(v48)+48))
					v93 = *(*float64)(unsafe.Add(mBase, uint32(v48)+40))
					if v53 == v91 {
						v179 = v93
						v180 = v92
						v181 = v93
						v182 = v92
					} else {
						v98 = v91
						v106 = v93
						v107 = v92
						v108 = v93
						v109 = v92
						for {
							v115 = v48 + int32(40) + v98<<(uint(int32(4))%32)
							v116 = *(*float64)(unsafe.Add(mBase, uint32(v115)))
							v121 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v116)&int64(9223372036854775807)))
							if v121 == int32(0) {
								if base.F64_gt(v108, v116) != 0 {
									v125 = v116
								} else {
									v125 = v108
								}
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v108)&int64(9223372036854775807)) {
									v131 = v116
								} else {
									v131 = v125
								}
								v132 = v131
							} else {
								v132 = v108
							}
							v133 = *(*float64)(unsafe.Add(mBase, uint32(v115)+8))
							v138 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v133)&int64(9223372036854775807)))
							if v138 == int32(0) {
								if base.F64_gt(v109, v133) != 0 {
									v142 = v133
								} else {
									v142 = v109
								}
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v109)&int64(9223372036854775807)) {
									v148 = v133
								} else {
									v148 = v142
								}
								v149 = v148
							} else {
								v149 = v109
							}
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v116)&int64(9223372036854775807)) {
								v150 = v116
							} else {
								v150 = v106
							}
							if base.F64_lt(v106, v116) != 0 {
								v152 = v116
							} else {
								v152 = v150
							}
							if base.Ui64(base.I64_reinterpret_f64(v106)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v158 = v152
							} else {
								v158 = v106
							}
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v133)&int64(9223372036854775807)) {
								v159 = v133
							} else {
								v159 = v107
							}
							if base.F64_lt(v107, v133) != 0 {
								v161 = v133
							} else {
								v161 = v159
							}
							if base.Ui64(base.I64_reinterpret_f64(v107)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v167 = v161
							} else {
								v167 = v107
							}
							v169 = v98 + int32(1)
							if v169 != v53 {
								v98 = v169
								v106 = v158
								v107 = v167
								v108 = v132
								v109 = v149
								continue
							} else {
								break
							}
							break
						}
						v179 = v158
						v180 = v167
						v181 = v132
						v182 = v149
					}
				}
				*(*float64)(unsafe.Add(mBase, uint32(v48)+32)) = v182
				*(*float64)(unsafe.Add(mBase, uint32(v48)+8)) = v179
				*(*float64)(unsafe.Add(mBase, uint32(v48)+24)) = v181
				*(*float64)(unsafe.Add(mBase, uint32(v48)+16)) = v180
				v205 = base.I64_extend_i32_u(v48)
				return v205
			}
		}
	}
}
