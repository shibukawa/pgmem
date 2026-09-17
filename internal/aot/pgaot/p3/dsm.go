package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dsm_create(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
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
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v3
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsm_create[0])))
	if v20 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_dsm_create[0])) = uint8(v24)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[1]))
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[2]))
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_ResourceOwnerEnlarge(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[3]))
	v37 = F_MemoryContextAlloc(m, v35, int32(36))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L6
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[4]))
	if v40 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = int32(_a_F_dsm_create_0)
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_create[5])) = v43
	v47 = v43
	goto L12
L11:
	;
	v47 = v40
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(_a_F_dsm_create_0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v37
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_create[4])) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v37)+24)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v37)+16)) = int64(4294967295)
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v59
	if v59 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_ResourceOwnerRemember(m, v59, v37, int32(_a_F_dsm_create_1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v65 = v37 + int32(28)
	v67 = v37 + int32(24)
	v69 = v37 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = int32(0)
	if v27 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[6]))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v155 != 0 {
		goto L36
	} else {
		goto L37
	}
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	v77 = F_LWLockAcquire(m, v73+int32(_a_F_dsm_create_2), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	v107 = v3
	goto L20
L20:
	;
	goto L27
L21:
	;
	v79 = int32(12)
	v85 = int32(base.Ui32(l0)>>(uint(v79)%32)) + base.B2i32(l0&int32(4095) != int32(0))
	v88 = F_FreePageManagerGet(m, v27, v85, v15+v79)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	if v88 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[1]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v93 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v91 + v92<<(uint(v93)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v85 << (uint(v93) % 32)
	v146 = v85
	v150 = int32(1)
	goto L17
L24:
	;
	goto L25
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	F_LWLockRelease(m, v102+int32(_a_F_dsm_create_2))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v107 = v85
	goto L20
L27:
	;
	v121 = Fn13964(m, int64(32))
	mBase = m.M
	goto L29
L28:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	v138 = F_LWLockAcquire(m, v134+int32(_a_F_dsm_create_2), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L7
	} else {
		goto L33
	}
L29:
	;
	v123 = v121 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v123
	if v123 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v129 = F_dsm_impl_op(m, int32(0), v123, l0, v69, v67, v65, int32(21))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	if v129 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	v146 = v107
	v150 = v3
	goto L17
L34:
	;
	m.G0 = v15 + int32(16)
	return v338
L35:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	F_LWLockRelease(m, v331+int32(_a_F_dsm_create_2))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L7
	} else {
		goto L74
	}
L36:
	;
	v159 = int32(0)
	goto L39
L37:
	;
	goto L38
L38:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	if base.Ui32(v225) <= base.Ui32(v155) {
		goto L49
	} else {
		goto L50
	}
L39:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v154+v159*int32(24))+16))
	if v171 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L38
L41:
	;
	if v150 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v211 = v159 + int32(1)
	if v211 != v155 {
		v159 = v211
		goto L39
	} else {
		goto L48
	}
L44:
	;
	v177 = Fn13964(m, int64(32))
	mBase = m.M
	goto L47
L45:
	;
	v195 = v154
	goto L46
L46:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v201 = v195 + v159*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v201)+16)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v201)+12)) = v198
	v205 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v201)+28)) = v205
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+32)) = uint8(v205)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v159
	goto L35
L47:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[6]))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v159<<(uint(int32(1))%32) | v177<<(uint(int32(32)-base.I32_clz(v181))%32) | int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v192 = v180 + v159*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v192)+20)) = v189
	v195 = v180
	goto L46
L48:
	;
	goto L40
L49:
	;
	if v150 != 0 {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	if v150 != 0 {
		goto L70
	} else {
		goto L71
	}
L52:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	if v248 != 0 {
		goto L60
	} else {
		goto L61
	}
L53:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	F_FreePageManagerPut(m, v27, v227, v146)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L7
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	F_LWLockRelease(m, v237+int32(_a_F_dsm_create_2))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L7
	} else {
		goto L58
	}
L56:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	F_LWLockRelease(m, v231+int32(_a_F_dsm_create_2))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	goto L52
L58:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v246 = F_dsm_impl_op(m, int32(3), v243, int32(0), v69, v67, v65, int32(19))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	goto L52
L60:
	;
	F_ResourceOwnerForget(m, v248, v37, int32(_a_F_dsm_create_1))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L7
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v252)+4)) = v253
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = v255
	F_pfree(m, v37)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L7
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	if l1&int32(1) != 0 {
		v338 = int32(0)
		goto L34
	} else {
		goto L65
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(_a_F_dsm_create_3), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_dsm_create_4), int32(626), int32(_a_F_dsm_create_5))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v281 = Fn13964(m, int64(32))
	mBase = m.M
	goto L73
L71:
	;
	v299 = v154
	goto L72
L72:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v305 = v299 + v155*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v305)+16)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v305)+12)) = v302
	v309 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v305)+28)) = v309
	*(*uint8)(unsafe.Add(mBase, uint32(v305)+32)) = uint8(v309)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v155
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v299)+4)) = v314 + int32(1)
	goto L35
L73:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[6]))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v155<<(uint(int32(1))%32) | v281<<(uint(int32(32)-base.I32_clz(v285))%32) | int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v296 = v284 + v155*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v296)+24)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v296)+20)) = v293
	v299 = v284
	goto L72
L74:
	;
	v338 = v37
	goto L34
}
