package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_cube_same(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v66 int32
	_ = v66
	var v70 float64
	_ = v70
	var v73 float64
	_ = v73
	var v74 int32
	_ = v74
	var v75 float64
	_ = v75
	var v77 int32
	_ = v77
	var v81 float64
	_ = v81
	var v84 float64
	_ = v84
	var v89 float64
	_ = v89
	var v91 float64
	_ = v91
	var v96 float64
	_ = v96
	var v98 float64
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 float64
	_ = v131
	var v133 int32
	_ = v133
	var v137 float64
	_ = v137
	var v140 float64
	_ = v140
	var v141 int32
	_ = v141
	var v142 float64
	_ = v142
	var v144 int32
	_ = v144
	var v148 float64
	_ = v148
	var v151 float64
	_ = v151
	var v156 float64
	_ = v156
	var v158 float64
	_ = v158
	var v163 float64
	_ = v163
	var v165 float64
	_ = v165
	var v169 int32
	_ = v169
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v214 int32
	_ = v214
	var v215 float64
	_ = v215
	var v222 float64
	_ = v222
	var v224 int32
	_ = v224
	var v228 float64
	_ = v228
	var v234 float64
	_ = v234
	var v240 float64
	_ = v240
	var v244 int32
	_ = v244
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v269 float64
	_ = v269
	var v276 float64
	_ = v276
	var v278 int32
	_ = v278
	var v282 float64
	_ = v282
	var v288 float64
	_ = v288
	var v294 float64
	_ = v294
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v329 int32
	_ = v329
	var v330 float64
	_ = v330
	var v337 float64
	_ = v337
	var v339 int32
	_ = v339
	var v343 float64
	_ = v343
	var v355 float64
	_ = v355
	var v359 float64
	_ = v359
	var v363 int32
	_ = v363
	var v376 int32
	_ = v376
	var v387 int32
	_ = v387
	var v388 float64
	_ = v388
	var v395 float64
	_ = v395
	var v397 int32
	_ = v397
	var v401 float64
	_ = v401
	var v413 float64
	_ = v413
	var v417 float64
	_ = v417
	var v422 int32
	_ = v422
	var v454 int32
	_ = v454
	var v498 int32
	_ = v498
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v32 = int32(2147483647)
	v33 = v31 & v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v36 = v34 & v32
	if base.Ui32(v33) < base.Ui32(v36) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v13)))) = uint8(base.B2i32(v498 == int32(0)))
	return v13 & int64(4294967295)
L5:
	;
	v498 = int32(-1)
	goto L4
L6:
	;
	v498 = v454
	goto L4
L7:
	;
	v454 = int32(1)
	goto L6
L8:
	;
	v38 = v33
	goto L10
L9:
	;
	v38 = v36
	goto L10
L10:
	;
	if v38 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = int32(8)
	v56 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	if base.Ui32(v36) < base.Ui32(v33) {
		goto L48
	} else {
		goto L49
	}
L14:
	;
	v62 = v56 << (uint(int32(3)) % 32)
	v63 = v6 + v39 + v62
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v63)))
	v66 = base.B2i32(v31 < int32(0))
	if v31 < int32(0) {
		v73 = v64
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v105 = int32(8)
	v123 = int32(0)
	goto L31
L16:
	;
	v74 = v62 + (v11 + v39)
	v75 = *(*float64)(unsafe.Add(mBase, uint32(v74)))
	v77 = base.B2i32(v34 < int32(0))
	if v34 < int32(0) {
		v84 = v75
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v63+v31<<(uint(int32(3))%32))))
	if base.F64_lt(v64, v70) != 0 {
		v73 = v64
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v73 = v70
	goto L16
L19:
	;
	if base.F64_gt(v73, v84) != 0 {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	v81 = *(*float64)(unsafe.Add(mBase, uint32(v74+v34<<(uint(int32(3))%32))))
	if base.F64_lt(v75, v81) != 0 {
		v84 = v75
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v84 = v81
	goto L19
L22:
	;
	if v31 < int32(0) {
		v91 = v64
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v34 < int32(0) {
		v98 = v75
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v89 = *(*float64)(unsafe.Add(mBase, uint32(v63+v31<<(uint(int32(3))%32))))
	if base.F64_lt(v64, v89) != 0 {
		v91 = v64
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v91 = v89
	goto L23
L26:
	;
	v100 = int32(-1)
	if base.F64_lt(v91, v98) != 0 {
		v454 = v100
		goto L6
	} else {
		goto L29
	}
L27:
	;
	v96 = *(*float64)(unsafe.Add(mBase, uint32(v74+v34<<(uint(int32(3))%32))))
	if base.F64_lt(v75, v96) != 0 {
		v98 = v75
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v98 = v96
	goto L26
L29:
	;
	v103 = v56 + int32(1)
	if v103 != v38 {
		v56 = v103
		goto L14
	} else {
		goto L30
	}
L30:
	;
	goto L15
L31:
	;
	v129 = v123 << (uint(int32(3)) % 32)
	v130 = v6 + v105 + v129
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v130)))
	v133 = base.B2i32(v31 < int32(0))
	if v31 < int32(0) {
		v140 = v131
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L13
L33:
	;
	v141 = v129 + (v11 + v105)
	v142 = *(*float64)(unsafe.Add(mBase, uint32(v141)))
	v144 = base.B2i32(v34 < int32(0))
	if v34 < int32(0) {
		v151 = v142
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v137 = *(*float64)(unsafe.Add(mBase, uint32(v130+v31<<(uint(int32(3))%32))))
	if base.F64_gt(v131, v137) != 0 {
		v140 = v131
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v140 = v137
	goto L33
L36:
	;
	if base.F64_gt(v140, v151) != 0 {
		goto L7
	} else {
		goto L39
	}
L37:
	;
	v148 = *(*float64)(unsafe.Add(mBase, uint32(v141+v34<<(uint(int32(3))%32))))
	if base.F64_gt(v142, v148) != 0 {
		v151 = v142
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v151 = v148
	goto L36
L39:
	;
	if v31 < int32(0) {
		v158 = v131
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if v34 < int32(0) {
		v165 = v142
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v156 = *(*float64)(unsafe.Add(mBase, uint32(v130+v31<<(uint(int32(3))%32))))
	if base.F64_gt(v131, v156) != 0 {
		v158 = v131
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v158 = v156
	goto L40
L43:
	;
	if base.F64_lt(v158, v165) != 0 {
		v454 = v100
		goto L6
	} else {
		goto L46
	}
L44:
	;
	v163 = *(*float64)(unsafe.Add(mBase, uint32(v141+v34<<(uint(int32(3))%32))))
	if base.F64_gt(v142, v163) != 0 {
		v165 = v142
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v165 = v163
	goto L43
L46:
	;
	v169 = v123 + int32(1)
	if v169 != v38 {
		v123 = v169
		goto L31
	} else {
		goto L47
	}
L47:
	;
	goto L32
L48:
	;
	v191 = v6 + int32(8)
	v194 = v38
	goto L51
L49:
	;
	goto L50
L50:
	;
	if base.Ui32(v36) <= base.Ui32(v33) {
		goto L81
	} else {
		goto L82
	}
L51:
	;
	v214 = v191 + v194<<(uint(int32(3))%32)
	v215 = *(*float64)(unsafe.Add(mBase, uint32(v214)))
	if base.B2i32(v31 < int32(0)) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v257 = v38
	goto L65
L53:
	;
	if base.F64_lt(v240, float64(0)) != 0 {
		goto L5
	} else {
		goto L63
	}
L54:
	;
	v222 = *(*float64)(unsafe.Add(mBase, uint32(v214+v33<<(uint(int32(3))%32))))
	if base.F64_lt(v215, v222) != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	if base.F64_gt(v215, float64(0)) != 0 {
		goto L7
	} else {
		goto L62
	}
L57:
	;
	v224 = int32(0)
	goto L59
L58:
	;
	v224 = v31
	goto L59
L59:
	;
	v228 = *(*float64)(unsafe.Add(mBase, uint32(v214+v224<<(uint(int32(3))%32))))
	if base.F64_gt(v228, float64(0)) != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v234 = *(*float64)(unsafe.Add(mBase, uint32(v214+v31<<(uint(int32(3))%32))))
	if base.F64_lt(v215, v234) == int32(0) {
		v240 = v234
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v240 = v215
	goto L53
L62:
	;
	v240 = v215
	goto L53
L63:
	;
	v244 = v194 + int32(1)
	if v244 != v33 {
		v194 = v244
		goto L51
	} else {
		goto L64
	}
L64:
	;
	goto L52
L65:
	;
	v268 = v191 + v257<<(uint(int32(3))%32)
	v269 = *(*float64)(unsafe.Add(mBase, uint32(v268)))
	if base.B2i32(v31 < int32(0)) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L5
L67:
	;
	if base.F64_lt(v294, float64(0)) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L68:
	;
	v276 = *(*float64)(unsafe.Add(mBase, uint32(v268+v33<<(uint(int32(3))%32))))
	if base.F64_gt(v269, v276) != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	if base.F64_gt(v269, float64(0)) != 0 {
		goto L7
	} else {
		goto L76
	}
L71:
	;
	v278 = int32(0)
	goto L73
L72:
	;
	v278 = v31
	goto L73
L73:
	;
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v268+v278<<(uint(int32(3))%32))))
	if base.F64_gt(v282, float64(0)) != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v288 = *(*float64)(unsafe.Add(mBase, uint32(v268+v31<<(uint(int32(3))%32))))
	if base.F64_gt(v269, v288) == int32(0) {
		v294 = v288
		goto L67
	} else {
		goto L75
	}
L75:
	;
	v294 = v269
	goto L67
L76:
	;
	v294 = v269
	goto L67
L77:
	;
	v299 = int32(1)
	v301 = v257 + v299
	if v301 == v33 {
		v454 = v299
		goto L6
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	goto L66
L80:
	;
	v257 = v301
	goto L65
L81:
	;
	v498 = int32(0)
	goto L4
L82:
	;
	goto L83
L83:
	;
	v306 = v11 + int32(8)
	v316 = v33
	goto L84
L84:
	;
	v329 = v306 + v316<<(uint(int32(3))%32)
	v330 = *(*float64)(unsafe.Add(mBase, uint32(v329)))
	if base.B2i32(v34 < int32(0)) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v376 = v38
	goto L99
L86:
	;
	if base.F64_lt(v359, float64(0)) != 0 {
		goto L7
	} else {
		goto L97
	}
L87:
	;
	v355 = *(*float64)(unsafe.Add(mBase, uint32(v329+v34<<(uint(int32(3))%32))))
	if base.F64_lt(v330, v355) == int32(0) {
		v359 = v355
		goto L86
	} else {
		goto L96
	}
L88:
	;
	v337 = *(*float64)(unsafe.Add(mBase, uint32(v329+v36<<(uint(int32(3))%32))))
	if base.F64_lt(v330, v337) != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if base.F64_gt(v330, float64(0)) == int32(0) {
		v359 = v330
		goto L86
	} else {
		goto L95
	}
L91:
	;
	v339 = int32(0)
	goto L93
L92:
	;
	v339 = v34
	goto L93
L93:
	;
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v329+v339<<(uint(int32(3))%32))))
	if base.F64_gt(v343, float64(0)) == int32(0) {
		goto L87
	} else {
		goto L94
	}
L94:
	;
	goto L5
L95:
	;
	goto L5
L96:
	;
	v359 = v330
	goto L86
L97:
	;
	v363 = v316 + int32(1)
	if v363 != v36 {
		v316 = v363
		goto L84
	} else {
		goto L98
	}
L98:
	;
	goto L85
L99:
	;
	v387 = v306 + v376<<(uint(int32(3))%32)
	v388 = *(*float64)(unsafe.Add(mBase, uint32(v387)))
	if base.B2i32(v34 < int32(0)) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v454 = int32(-1)
	goto L6
L101:
	;
	if base.F64_lt(v417, float64(0)) != 0 {
		goto L7
	} else {
		goto L112
	}
L102:
	;
	v413 = *(*float64)(unsafe.Add(mBase, uint32(v387+v34<<(uint(int32(3))%32))))
	if base.F64_gt(v388, v413) == int32(0) {
		v417 = v413
		goto L101
	} else {
		goto L111
	}
L103:
	;
	v395 = *(*float64)(unsafe.Add(mBase, uint32(v387+v36<<(uint(int32(3))%32))))
	if base.F64_gt(v388, v395) != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	if base.F64_gt(v388, float64(0)) == int32(0) {
		v417 = v388
		goto L101
	} else {
		goto L110
	}
L106:
	;
	v397 = int32(0)
	goto L108
L107:
	;
	v397 = v34
	goto L108
L108:
	;
	v401 = *(*float64)(unsafe.Add(mBase, uint32(v387+v397<<(uint(int32(3))%32))))
	if base.F64_gt(v401, float64(0)) == int32(0) {
		goto L102
	} else {
		goto L109
	}
L109:
	;
	goto L5
L110:
	;
	goto L5
L111:
	;
	v417 = v388
	goto L101
L112:
	;
	v422 = v376 + int32(1)
	if v36 != v422 {
		v376 = v422
		goto L99
	} else {
		goto L113
	}
L113:
	;
	goto L100
}
