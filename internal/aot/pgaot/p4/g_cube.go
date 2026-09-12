package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_cube_internal_consistent(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v71 int32
	_ = v71
	var v73 float64
	_ = v73
	var v75 int32
	_ = v75
	var v80 float64
	_ = v80
	var v83 float64
	_ = v83
	var v85 float64
	_ = v85
	var v87 int32
	_ = v87
	var v92 float64
	_ = v92
	var v95 float64
	_ = v95
	var v96 int32
	_ = v96
	var v102 float64
	_ = v102
	var v104 float64
	_ = v104
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v156 int32
	_ = v156
	var v157 float64
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 float64
	_ = v168
	var v170 int32
	_ = v170
	var v171 float64
	_ = v171
	var v174 float64
	_ = v174
	var v176 float64
	_ = v176
	var v179 float64
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v268 float64
	_ = v268
	var v277 float64
	_ = v277
	var v281 int32
	_ = v281
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v322 int32
	_ = v322
	var v324 float64
	_ = v324
	var v326 int32
	_ = v326
	var v331 float64
	_ = v331
	var v334 float64
	_ = v334
	var v336 float64
	_ = v336
	var v338 int32
	_ = v338
	var v343 float64
	_ = v343
	var v346 float64
	_ = v346
	var v347 int32
	_ = v347
	var v353 float64
	_ = v353
	var v355 float64
	_ = v355
	var v361 float64
	_ = v361
	var v363 float64
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v398 int32
	_ = v398
	v4 = int32(0)
	if base.Ui32(int32(14)) < base.Ui32(l2) {
		v221 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v224 = int32(0)
	if l0 == v224 {
		v374 = v224
		goto L58
	} else {
		goto L59
	}
L2:
	;
	return v221
L3:
	;
	v9 = int32(1) << (uint(l2) % 32)
	if v9&int32(8384) != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if base.B2i32(v9&int32(16640) == int32(0))&base.B2i32(l2 != int32(3)) != 0 {
		v221 = v4
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v23 = int32(0)
	if l0 == v23 {
		v199 = v23
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v221 = v220
	goto L2
L7:
	;
	v220 = v199
	goto L6
L8:
	;
	if l1 == int32(0) {
		v199 = v23
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = int32(2147483647)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v41 = base.B2i32(base.Ui32(v35&v36) < base.Ui32(v38&v36))
	if base.Ui32(v35&v36) < base.Ui32(v38&v36) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v42 = l1
	goto L12
L11:
	;
	v42 = l0
	goto L12
L12:
	;
	if base.Ui32(v35&v36) < base.Ui32(v38&v36) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v133 = v124 & int32(2147483647)
	if base.Ui32(v133) <= base.Ui32(v46) {
		goto L37
	} else {
		goto L38
	}
L14:
	;
	v43 = l0
	goto L16
L15:
	;
	v43 = l1
	goto L16
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v46 = v44 & int32(2147483647)
	if v46 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v124 = v49
	goto L13
L18:
	;
	goto L19
L19:
	;
	v50 = int32(8)
	v51 = v43 + v50
	v53 = v42 + v50
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v56 = int32(0)
	goto L20
L20:
	;
	v71 = v56 << (uint(int32(3)) % 32)
	v73 = *(*float64)(unsafe.Add(mBase, uint32(v53+v71)))
	v75 = base.B2i32(v54 < int32(0))
	if v54 < int32(0) {
		v83 = v73
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v124 = v54
	goto L13
L22:
	;
	v85 = *(*float64)(unsafe.Add(mBase, uint32(v71+v51)))
	v87 = base.B2i32(v44 < int32(0))
	if v44 < int32(0) {
		v95 = v85
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v80 = *(*float64)(unsafe.Add(mBase, uint32(v53+(v56+v54)<<(uint(int32(3))%32))))
	if base.F64_lt(v73, v80) != 0 {
		v83 = v73
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v83 = v80
	goto L22
L25:
	;
	v96 = int32(0)
	if base.F64_gt(v83, v95) != 0 {
		v199 = v96
		goto L7
	} else {
		goto L28
	}
L26:
	;
	v92 = *(*float64)(unsafe.Add(mBase, uint32(v51+(v56+v44)<<(uint(int32(3))%32))))
	if base.F64_gt(v85, v92) != 0 {
		v95 = v85
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v95 = v92
	goto L25
L28:
	;
	if v54 < int32(0) {
		v104 = v73
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v44 < int32(0) {
		v112 = v85
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v102 = *(*float64)(unsafe.Add(mBase, uint32(v53+(v56+v54)<<(uint(int32(3))%32))))
	if base.F64_gt(v73, v102) != 0 {
		v104 = v73
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v104 = v102
	goto L29
L32:
	;
	if base.F64_lt(v104, v112) != 0 {
		v199 = v96
		goto L7
	} else {
		goto L35
	}
L33:
	;
	v110 = *(*float64)(unsafe.Add(mBase, uint32(v51+(v56+v44)<<(uint(int32(3))%32))))
	if base.F64_lt(v85, v110) != 0 {
		v112 = v85
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v112 = v110
	goto L32
L35:
	;
	v116 = v56 + int32(1)
	if v116 != v46 {
		v56 = v116
		goto L20
	} else {
		goto L36
	}
L36:
	;
	goto L21
L37:
	;
	v220 = int32(1)
	goto L6
L38:
	;
	goto L39
L39:
	;
	v137 = v42 + int32(8)
	v141 = v46
	goto L40
L40:
	;
	v156 = v137 + v141<<(uint(int32(3))%32)
	v157 = *(*float64)(unsafe.Add(mBase, uint32(v156)))
	if base.B2i32(v124 < int32(0)) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v199 = int32(0)
	goto L7
L42:
	;
	goto L41
L43:
	;
	if base.F64_lt(v179, float64(0)) != 0 {
		goto L42
	} else {
		goto L55
	}
L44:
	;
	v161 = int32(3)
	v163 = v137 + (v141+v124)<<(uint(v161)%32)
	v168 = *(*float64)(unsafe.Add(mBase, uint32(v137+(v141+v133)<<(uint(v161)%32))))
	if base.F64_lt(v157, v168) != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	if base.F64_gt(v157, float64(0)) != 0 {
		goto L42
	} else {
		goto L54
	}
L47:
	;
	v170 = v156
	goto L49
L48:
	;
	v170 = v163
	goto L49
L49:
	;
	v171 = *(*float64)(unsafe.Add(mBase, uint32(v170)))
	if base.F64_gt(v171, float64(0)) != 0 {
		goto L42
	} else {
		goto L50
	}
L50:
	;
	v174 = *(*float64)(unsafe.Add(mBase, uint32(v163)))
	if base.F64_gt(v157, v174) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v176 = v157
	goto L53
L52:
	;
	v176 = v174
	goto L53
L53:
	;
	v179 = v176
	goto L43
L54:
	;
	v179 = v157
	goto L43
L55:
	;
	v184 = int32(1)
	v186 = v141 + v184
	if v133 != v186 {
		v141 = v186
		goto L40
	} else {
		goto L56
	}
L56:
	;
	v199 = v184
	goto L7
L57:
	;
	return v398
L58:
	;
	v398 = v374
	goto L57
L59:
	;
	if l1 == int32(0) {
		v374 = v224
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v241 = int32(2147483647)
	v242 = v240 & v241
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v245 = v243 & v241
	if base.Ui32(v242) < base.Ui32(v245) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v248 = l1 + int32(8)
	v253 = v242
	goto L64
L62:
	;
	goto L63
L63:
	;
	if base.Ui32(v242) < base.Ui32(v245) {
		goto L72
	} else {
		goto L73
	}
L64:
	;
	v268 = *(*float64)(unsafe.Add(mBase, uint32(v248+v253<<(uint(int32(3))%32))))
	if base.F64_ne(v268, float64(0)) != 0 {
		v374 = v224
		goto L58
	} else {
		goto L66
	}
L65:
	;
	goto L63
L66:
	;
	if base.B2i32(v243 < int32(0)) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v277 = *(*float64)(unsafe.Add(mBase, uint32(v248+(v253+v245)<<(uint(int32(3))%32))))
	if base.F64_ne(v277, float64(0)) != 0 {
		v374 = v224
		goto L58
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v281 = v253 + int32(1)
	if v281 != v245 {
		v253 = v281
		goto L64
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	goto L65
L72:
	;
	v298 = v242
	goto L74
L73:
	;
	v298 = v245
	goto L74
L74:
	;
	if v298 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v398 = int32(1)
	goto L57
L76:
	;
	goto L77
L77:
	;
	v302 = int32(8)
	v303 = l1 + v302
	v305 = l0 + v302
	v309 = int32(0)
	goto L78
L78:
	;
	v322 = v309 << (uint(int32(3)) % 32)
	v324 = *(*float64)(unsafe.Add(mBase, uint32(v305+v322)))
	v326 = base.B2i32(v240 < int32(0))
	if v240 < int32(0) {
		v334 = v324
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v374 = v366
	goto L58
L80:
	;
	v336 = *(*float64)(unsafe.Add(mBase, uint32(v303+v322)))
	v338 = base.B2i32(v243 < int32(0))
	if v243 < int32(0) {
		v346 = v336
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v331 = *(*float64)(unsafe.Add(mBase, uint32(v305+(v309+v240)<<(uint(int32(3))%32))))
	if base.F64_lt(v324, v331) != 0 {
		v334 = v324
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v334 = v331
	goto L80
L83:
	;
	v347 = int32(0)
	if base.F64_gt(v334, v346) != 0 {
		v374 = v347
		goto L58
	} else {
		goto L86
	}
L84:
	;
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v303+(v309+v243)<<(uint(int32(3))%32))))
	if base.F64_lt(v336, v343) != 0 {
		v346 = v336
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v346 = v343
	goto L83
L86:
	;
	if v240 < int32(0) {
		v355 = v324
		goto L87
	} else {
		goto L88
	}
L87:
	;
	if v243 < int32(0) {
		v363 = v336
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v353 = *(*float64)(unsafe.Add(mBase, uint32(v305+(v309+v240)<<(uint(int32(3))%32))))
	if base.F64_gt(v324, v353) != 0 {
		v355 = v324
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v355 = v353
	goto L87
L90:
	;
	if base.F64_gt(v363, v355) != 0 {
		v374 = v347
		goto L58
	} else {
		goto L93
	}
L91:
	;
	v361 = *(*float64)(unsafe.Add(mBase, uint32(v303+(v309+v243)<<(uint(int32(3))%32))))
	if base.F64_gt(v336, v361) != 0 {
		v363 = v336
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v363 = v361
	goto L90
L93:
	;
	v366 = int32(1)
	v368 = v309 + v366
	if v368 != v298 {
		v309 = v368
		goto L78
	} else {
		goto L94
	}
L94:
	;
	goto L79
}
