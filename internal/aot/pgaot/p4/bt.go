package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__bt_allequalimage(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+8)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+10)))
	if v13 != v14 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v107
L2:
	;
	v107 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	if int32(0) < base.I32_extend16_s(v13) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v95 + int32(4)
	F_errmsg_internal(m, v88, v10)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L15
	} else {
		goto L31
	}
L6:
	;
	v78 = int32(0)
	v81 = F_errstart(m, int32(14), v78)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L15
	} else {
		goto L29
	}
L7:
	;
	v67 = int32(1)
	v70 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L15
	} else {
		goto L27
	}
L8:
	;
	if v47 == int32(0) {
		goto L6
	} else {
		goto L26
	}
L9:
	;
	if l1 != 0 {
		goto L6
	} else {
		goto L25
	}
L10:
	;
	v24 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L24
	}
L13:
	;
	v30 = v24 << (uint(int32(2)) % 32)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30+v31)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34+v30)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37+v30)))
	v41 = F_get_opfamily_proc(m, v36, v39, v39, int32(4))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if l1 != 0 {
		goto L8
	} else {
		goto L23
	}
L15:
	;
	return int32(0)
L16:
	;
	if v41 == int32(0) {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v47 = F_OidFunctionCall1Coll(m, v41, v33, v39)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	if v47 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v50 = v24 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+10)))
	if v50 < v52 {
		v24 = v50
		goto L13
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L14
L22:
	;
	goto L21
L23:
	;
	v107 = base.B2i32(v47 != int32(0))
	goto L1
L24:
	;
	v107 = int32(1)
	goto L1
L25:
	;
	v107 = int32(0)
	goto L1
L26:
	;
	goto L7
L27:
	;
	if v70 == int32(0) {
		v107 = v67
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v88 = int32(_a_F__bt_allequalimage_0)
	v89 = v67
	v94 = int32(_a_F__bt_allequalimage_1)
	goto L5
L29:
	;
	if v81 == int32(0) {
		v107 = v78
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v88 = int32(_a_F__bt_allequalimage_2)
	v89 = v78
	v94 = int32(_a_F__bt_allequalimage_3)
	goto L5
L31:
	;
	F_errfinish(m, int32(_a_F__bt_allequalimage_4), v94, int32(_a_F__bt_allequalimage_5))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	v107 = v89
	goto L1
}
func F__bt_allocbuf(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v172 int64
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v276 int64
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 - int32(-64)
	return v358
L2:
	;
	return int32(0)
L3:
	;
	if v12 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v12
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l0
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v10)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v278
	v282 = int32(0)
	v285 = F_ExtendBufferedRel(m, v8+int32(-56), v282, v282, int32(8))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L2
	} else {
		goto L88
	}
L7:
	;
	v25 = F_ReadBuffer(m, l0, v23)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	F_ReleaseBuffer(m, v25)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L2
	} else {
		goto L84
	}
L10:
	;
	v27 = F_ConditionalLockBuffer(m, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	if v27 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v25 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	v245 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L2
	} else {
		goto L80
	}
L15:
	;
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+14)))
	if v47 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F__bt_allocbuf[0]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+(v25^int32(-1))<<(uint(int32(2))%32))))
	v46 = v38
	goto L15
L17:
	;
	goto L18
L18:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F__bt_allocbuf[1]))
	v46 = v40 + v25<<(uint(int32(13))%32) + int32(-8192)
	goto L15
L19:
	;
	v50 = int32(_a_F__bt_allocbuf_0)
	v52 = int32(0)
	if v52|(v46&int32(3)|int32(1)) == v52 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	goto L21
L21:
	;
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+16)))
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46+v100)+12)))
	if v102&int32(4) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L22:
	;
	v358 = v25
	goto L1
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+10)) = int32(_a_F__bt_allocbuf_1)
	v91 = int32(_a_F__bt_allocbuf_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v46)+18)) = uint16(v91)
	v97 = int32(_a_F__bt_allocbuf_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v46)+16)) = uint16(v97)
	*(*uint16)(unsafe.Add(mBase, uint32(v46)+14)) = uint16(v97)
	goto L22
L24:
	;
	goto L27
L25:
	;
	goto L26
L26:
	;
	goto L32
L27:
	;
	v68 = v46 + v50
	v70 = v46 + int32(4)
	if base.Ui32(v70) < base.Ui32(v68) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v72 = v68
	goto L30
L29:
	;
	v72 = v70
	goto L30
L30:
	;
	v77 = (v46^int32(-1)+v72)&int32(-4) + int32(4)
	if v77 == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	base.MemoryFill(m, v46, int32(0), v77)
	goto L23
L32:
	;
	base.MemoryFill(m, v46, int32(0), v50)
	goto L23
L33:
	;
	v229 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L2
	} else {
		goto L73
	}
L34:
	;
	if v102&int32(256) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v46)+24))
	v111 = v109
	goto L37
L36:
	;
	v111 = int64(3)
	goto L37
L37:
	;
	v112 = F_GlobalVisCheckRemovableFullXid(m, l1, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	if v112 == int32(0) {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+118)))
	if v117 != int32(112) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v177 = int32(_a_F__bt_allocbuf_0)
	v179 = int32(0)
	if v179|(v46&int32(3)|int32(1)) == v179 {
		goto L64
	} else {
		goto L65
	}
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F__bt_allocbuf[2]))
	if v121 <= int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v124
	v126 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v23
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+16)))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v129)+13)))
	if v131&int32(1) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v46)+24))
	v136 = v134
	goto L45
L44:
	;
	v136 = int64(3)
	goto L45
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v136
	if base.Ui32(v121) < base.Ui32(int32(2)) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+56)) = uint8(v160)
	F_XLogBeginInsert(m)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L59
	}
L47:
	;
	v160 = int32(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+118)))
	if v142 != int32(112) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v160 = int32(0)
	goto L46
L51:
	;
	goto L52
L52:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	goto L53
L53:
	;
	if base.Ui32(v147) < base.Ui32(int32(_a_F__bt_allocbuf_4)) {
		v160 = int32(1)
		goto L46
	} else {
		goto L54
	}
L54:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	if v150 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v160 = int32(0)
	goto L46
L56:
	;
	goto L57
L57:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+119)))
	switch v156 - int32(109) {
	case 0, 5:
		goto L58
	default:
		v160 = int32(0)
		goto L46
	}
L58:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+104)))
	v160 = v159
	goto L46
L59:
	;
	F_XLogRegisterData(m, v8+int32(-32), int32(25))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v172 = F_XLogInsert(m, int32(11), int32(208))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	goto L40
L62:
	;
	v358 = v25
	goto L1
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+10)) = int32(_a_F__bt_allocbuf_1)
	v218 = int32(_a_F__bt_allocbuf_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v46)+18)) = uint16(v218)
	v224 = int32(_a_F__bt_allocbuf_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v46)+16)) = uint16(v224)
	*(*uint16)(unsafe.Add(mBase, uint32(v46)+14)) = uint16(v224)
	goto L62
L64:
	;
	goto L67
L65:
	;
	goto L66
L66:
	;
	goto L72
L67:
	;
	v195 = v46 + v177
	v197 = v46 + int32(4)
	if base.Ui32(v197) < base.Ui32(v195) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v199 = v195
	goto L70
L69:
	;
	v199 = v197
	goto L70
L70:
	;
	v204 = (v46^int32(-1)+v199)&int32(-4) + int32(4)
	if v204 == int32(0) {
		goto L63
	} else {
		goto L71
	}
L71:
	;
	base.MemoryFill(m, v46, int32(0), v204)
	goto L63
L72:
	;
	base.MemoryFill(m, v46, int32(0), v177)
	goto L63
L73:
	;
	if v229 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_errmsg_internal(m, int32(_a_F__bt_allocbuf_5), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L2
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	F_LockBuffer(m, v25, int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L2
	} else {
		goto L79
	}
L77:
	;
	F_errfinish(m, int32(_a_F__bt_allocbuf_6), int32(960), int32(_a_F__bt_allocbuf_7))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	goto L9
L80:
	;
	if v245 == int32(0) {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	F_errmsg_internal(m, int32(_a_F__bt_allocbuf_8), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F__bt_allocbuf_6), int32(965), int32(_a_F__bt_allocbuf_7))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	goto L9
L84:
	;
	v262 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	if v262 != int32(-1) {
		v23 = v262
		goto L7
	} else {
		goto L86
	}
L86:
	;
	goto L8
L87:
	;
	v305 = int32(_a_F__bt_allocbuf_0)
	v307 = int32(0)
	if v307|(v304&int32(3)|int32(1)) == v307 {
		goto L94
	} else {
		goto L95
	}
L88:
	;
	if v285 < int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _c_F__bt_allocbuf[0]))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v290+(v285^int32(-1))<<(uint(int32(2))%32))))
	v304 = v296
	goto L87
L90:
	;
	goto L91
L91:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _c_F__bt_allocbuf[1]))
	v304 = v298 + v285<<(uint(int32(13))%32) + int32(-8192)
	goto L87
L92:
	;
	v358 = v285
	goto L1
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v304)+10)) = int32(_a_F__bt_allocbuf_1)
	v346 = int32(_a_F__bt_allocbuf_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v304)+18)) = uint16(v346)
	v352 = int32(_a_F__bt_allocbuf_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v304)+16)) = uint16(v352)
	*(*uint16)(unsafe.Add(mBase, uint32(v304)+14)) = uint16(v352)
	goto L92
L94:
	;
	goto L97
L95:
	;
	goto L96
L96:
	;
	goto L102
L97:
	;
	v323 = v304 + v305
	v325 = v304 + int32(4)
	if base.Ui32(v325) < base.Ui32(v323) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v327 = v323
	goto L100
L99:
	;
	v327 = v325
	goto L100
L100:
	;
	v332 = (v304^int32(-1)+v327)&int32(-4) + int32(4)
	if v332 == int32(0) {
		goto L93
	} else {
		goto L101
	}
L101:
	;
	base.MemoryFill(m, v304, int32(0), v332)
	goto L93
L102:
	;
	base.MemoryFill(m, v304, int32(0), v305)
	goto L93
}
func F__bt_check_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v354 int32
	_ = v354
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v394 int32
	_ = v394
	var v423 int32
	_ = v423
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v23)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v27 <= v26 {
		v423 = v23
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v423
L2:
	;
	v39 = v26
	goto L5
L3:
	;
	v423 = int32(0)
	goto L1
L4:
	;
	v394 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v394)
	goto L3
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v49 = v46 + v39*int32(48)
	v50 = int32(0)
	if l6 != 0 {
		v81 = v50
		v84 = v50
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l6 != 0 {
		goto L3
	} else {
		goto L106
	}
L7:
	;
	v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+4)))
	if l3 < v85 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v52 = int32(1)
	v53 = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v56 = v54 & int32(_a_F__bt_check_compare_0)
	if base.B2i32(v56 == v53)|base.B2i32(l1 != v52) == v53 {
		v81 = v52
		v84 = v53
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v65 = v54 & int32(_a_F__bt_check_compare_1)
	if l1 == int32(-1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if v65 != 0 {
		v81 = v52
		v84 = int32(0)
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v69 = int32(0)
	v81 = v69
	v84 = base.B2i32(l1 == int32(-1))&base.B2i32(v56 != v69) | base.B2i32(l1 == int32(1))&base.B2i32(v65 != v69)
	goto L7
L13:
	;
	goto L12
L14:
	;
	goto L6
L15:
	;
	v337 = int32(1)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v340 = v338 + v337
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v340
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v340 < v342 {
		v39 = v340
		goto L5
	} else {
		goto L105
	}
L16:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v87&int32(_a_F__bt_check_compare_2) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v304 = int32(0)
	if l5 == v304 {
		v423 = v304
		goto L1
	} else {
		goto L101
	}
L18:
	;
	v301 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v301)
	v423 = v301
	goto L1
L19:
	;
	if l6 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v87&int32(4) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v92 = int32(0)
	v94 = F__bt_advance_array_keys(m, l0, v92, l2, l3, l4, v39, v92)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	v423 = v94
	goto L1
L25:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v49)+44))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v101&int32(1) != 0 {
		v354 = v100
		goto L14
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v224 = F_index_getattr_2(m, l2, v85, l4, v20+int32(14))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L23
	} else {
		goto L79
	}
L28:
	;
	v114 = v100
	goto L29
L29:
	;
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v114)+4)))
	if l3 < v121 {
		goto L15
	} else {
		goto L31
	}
L30:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+6)))
	switch v177 - int32(1) {
	case 0:
		goto L61
	case 1:
		goto L65
	default:
		goto L62
	case 3:
		goto L64
	case 4:
		goto L63
	}
L31:
	;
	v125 = F_index_getattr_2(m, l2, v121, l4, v20+int32(15))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+15)))
	if v127 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if l6 != 0 {
		goto L3
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v114)+44))
	v157 = F_FunctionCall2Coll(m, v114+int32(16), v155, v125, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L23
	} else {
		goto L49
	}
L36:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v49)+44))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v131&int32(33554432) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if l1 != int32(-1) {
		goto L3
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v130 == v114 {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	if v130 == v114 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v139 = int32(_a_F__bt_check_compare_3)
	goto L43
L42:
	;
	v139 = int32(_a_F__bt_check_compare_1)
	goto L43
L43:
	;
	if v139&v131 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	goto L3
L45:
	;
	v144 = int32(_a_F__bt_check_compare_3)
	goto L47
L46:
	;
	v144 = int32(_a_F__bt_check_compare_0)
	goto L47
L47:
	;
	if base.B2i32(v144&v131 == int32(0))|base.B2i32(l1 != int32(1)) != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	goto L4
L49:
	;
	if v157 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v162 = int32(1)
	goto L52
L51:
	;
	v162 = int32(0) - v157
	goto L52
L52:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v163&int32(16777216) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v166 = v162
	goto L55
L54:
	;
	v166 = v157
	goto L55
L55:
	;
	if v166|v163&int32(16) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+48)))
	v174 = v114 + int32(48)
	if v172&int32(1) != 0 {
		v354 = v174
		goto L14
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	goto L30
L59:
	;
	v114 = v174
	goto L29
L60:
	;
	if l6|v202&int32(1) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L61:
	;
	v202 = int32(base.Ui32(v166) >> (uint(int32(31)) % 32))
	goto L60
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L23
	} else {
		goto L66
	}
L63:
	;
	v202 = base.B2i32(int32(0) < v166)
	goto L60
L64:
	;
	v202 = base.B2i32(int32(0) <= v166)
	goto L60
L65:
	;
	v202 = base.B2i32(v166 <= int32(0))
	goto L60
L66:
	;
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v190
	F_errmsg_internal(m, int32(_a_F__bt_check_compare_4), v20)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F__bt_check_compare_5), int32(3153), int32(_a_F__bt_check_compare_6))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L23
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	if v163&int32(_a_F__bt_check_compare_0) != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	if v202&int32(1) != 0 {
		goto L15
	} else {
		goto L78
	}
L72:
	;
	v213 = base.B2i32(l1 == int32(1))
	goto L74
L73:
	;
	v213 = int32(0)
	goto L74
L74:
	;
	if v213 != 0 {
		goto L18
	} else {
		goto L75
	}
L75:
	;
	v214 = int32(0)
	if l1 != int32(-1) {
		v423 = v214
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v163&int32(_a_F__bt_check_compare_1) != 0 {
		goto L18
	} else {
		goto L77
	}
L77:
	;
	v423 = v214
	goto L1
L78:
	;
	v423 = int32(0)
	goto L1
L79:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v226&int32(1) != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+14)))
	if (v229^base.B2i32(v226&int32(64) == int32(0)))&int32(1) != 0 {
		goto L15
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+14)))
	if v240 == int32(1) {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	if v81 != 0 {
		goto L18
	} else {
		goto L84
	}
L84:
	;
	if v226&int32(_a_F__bt_check_compare_7) != 0 {
		goto L15
	} else {
		goto L85
	}
L85:
	;
	v423 = int32(0)
	goto L1
L86:
	;
	v243 = int32(0)
	if base.B2i32(l6 == v243)|base.B2i32(v226&int32(_a_F__bt_check_compare_7) == v243) == v243 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v49)+44))
	v280 = F_FunctionCall2Coll(m, v49+int32(16), v278, v224, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L23
	} else {
		goto L98
	}
L89:
	;
	v252 = int32(0)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v255 = F__bt_advance_array_keys(m, l0, v252, l2, l3, l4, v253, v252)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L23
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v257 = v81 | v84
	if v226&int32(33554432) != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v423 = v255
	goto L1
L93:
	;
	if v257^int32(1)|base.B2i32(l1 != int32(-1)) == int32(0) {
		goto L18
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v268 = int32(1)
	if v257^v268|base.B2i32(l1 != v268) == int32(0) {
		goto L18
	} else {
		goto L97
	}
L96:
	;
	v423 = int32(0)
	goto L1
L97:
	;
	v423 = int32(0)
	goto L1
L98:
	;
	if v280 != 0 {
		goto L15
	} else {
		goto L99
	}
L99:
	;
	if v81 == int32(0) {
		goto L17
	} else {
		goto L100
	}
L100:
	;
	goto L18
L101:
	;
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+6)))
	if v307 != int32(3) {
		v423 = v304
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v310&int32(32) == int32(0) {
		v423 = v304
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v315 = int32(0)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v318 = F__bt_advance_array_keys(m, l0, v315, l2, l3, l4, v316, v315)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L23
	} else {
		goto L104
	}
L104:
	;
	v423 = v318
	goto L1
L105:
	;
	v423 = v337
	goto L1
L106:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v354-int32(48))))
	if v366&int32(_a_F__bt_check_compare_0) != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v369 = base.B2i32(l1 == int32(1))
	goto L109
L108:
	;
	v369 = int32(0)
	goto L109
L109:
	;
	if v369 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	if base.B2i32(v366&int32(_a_F__bt_check_compare_1) == int32(0))|base.B2i32(l1 != int32(-1)) != 0 {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	goto L4
}
func F__bt_freestack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v3 = l0
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	F_pfree(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	if v5 != 0 {
		v3 = v5
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F__bt_get_endpoint(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v397
L2:
	;
	if v234 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L3:
	;
	v18 = F__bt_getroot(m, l0, int32(0), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v22 = m.G0
	v24 = v22 + int32(-64)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	return int32(0)
L7:
	;
	v234 = v18
	goto L2
L8:
	;
	v234 = v150
	goto L2
L9:
	;
	F_pfree(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v29
	v32 = F_ReadBuffer(m, l0, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	F_LockBuffer(m, v32, int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	F__bt_checkpage(m, l0, v32)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	if v32 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L6
	} else {
		goto L60
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L6
	} else {
		goto L56
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L52
	}
L19:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+16)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+v56)+12)))
	if v59&int32(8) == int32(0) {
		goto L18
	} else {
		goto L23
	}
L20:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[0]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v32^int32(-1))<<(uint(int32(2))%32))))
	v56 = v48
	goto L19
L21:
	;
	goto L22
L22:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[1]))
	v56 = v50 + v32<<(uint(int32(13))%32) + int32(-8192)
	goto L19
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)+24))
	if v64 != int32(_a_F__bt_get_endpoint_0) {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	if base.Ui32(v67-int32(5)) <= base.Ui32(int32(-4)) {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
	if v72 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	m.G0 = v24 - int32(-64)
	goto L8
L27:
	;
	v75 = int32(0)
	F_LockBuffer(m, v32, v75)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v56)+36))
	v85 = v32
	v88 = v72
	goto L33
L30:
	;
	F_ReleaseBuffer(m, v32)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v150 = v75
	goto L26
L32:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	if v145 != v81 {
		goto L16
	} else {
		goto L51
	}
L33:
	;
	if v85 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L48
	}
L35:
	;
	F_LockBuffer(m, v85, int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v94 = F_ReleaseAndReadBuffer(m, v85, l0, v88)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	F_LockBuffer(m, v94, int32(1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	F__bt_checkpage(m, l0, v94)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	if v94 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+16)))
	v120 = v119 + v118
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)))
	if v121&int32(20) == int32(0) {
		goto L32
	} else {
		goto L46
	}
L43:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[0]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v104+(v94^int32(-1))<<(uint(int32(2))%32))))
	v118 = v110
	goto L42
L44:
	;
	goto L45
L45:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[1]))
	v118 = v112 + v94<<(uint(int32(13))%32) + int32(-8192)
	goto L42
L46:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v126 != 0 {
		v85 = v94
		v88 = v126
		goto L33
	} else {
		goto L47
	}
L47:
	;
	goto L34
L48:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v131 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_get_endpoint_1), v22+int32(-16))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F__bt_get_endpoint_2), int32(651), int32(_a_F__bt_get_endpoint_3))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v150 = v94
	goto L26
L52:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v166 + int32(4)
	F_errmsg(m, int32(_a_F__bt_get_endpoint_4), v24)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F__bt_get_endpoint_2), int32(612), int32(_a_F__bt_get_endpoint_3))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = int64(8589934596)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v185 + int32(4)
	F_errmsg(m, int32(_a_F__bt_get_endpoint_5), v22+int32(-48))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F__bt_get_endpoint_2), int32(621), int32(_a_F__bt_get_endpoint_3))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v207 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_get_endpoint_6), v22+int32(-32))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F__bt_get_endpoint_2), int32(658), int32(_a_F__bt_get_endpoint_3))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	v397 = int32(0)
	goto L1
L64:
	;
	goto L65
L65:
	;
	if v234 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v259 = v255
	v262 = v234
	goto L71
L67:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[0]))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v241+(v234^int32(-1))<<(uint(int32(2))%32))))
	v255 = v247
	goto L66
L68:
	;
	goto L69
L69:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[1]))
	v255 = v249 + v234<<(uint(int32(13))%32) + int32(-8192)
	goto L66
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L6
	} else {
		goto L105
	}
L71:
	;
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+16)))
	v266 = v259 + v265
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+12)))
	if v267&int32(20) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L6
	} else {
		goto L102
	}
L73:
	;
	goto L72
L74:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v259+v319<<(uint(int32(2))%32))+20))
	v326 = v259 + v323&int32(_a_F__bt_get_endpoint_7)
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v326))))
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v326)+2)))
	v333 = F__bt_relandgetbuf(m, l0, v262, v327<<(uint(int32(16))%32)|v330, int32(1))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L6
	} else {
		goto L98
	}
L75:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v316 != 0 {
		goto L95
	} else {
		goto L96
	}
L76:
	;
	v295 = F__bt_relandgetbuf(m, l0, v262, v293, int32(1))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L6
	} else {
		goto L91
	}
L77:
	;
	if l2 != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v290 == int32(0) {
		goto L73
	} else {
		goto L90
	}
L80:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v272 != 0 {
		v293 = v272
		goto L76
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
	if v274 == l1 {
		v397 = v262
		goto L1
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	if base.Ui32(v274) < base.Ui32(l1) {
		goto L70
	} else {
		goto L85
	}
L85:
	;
	if l2 == int32(0) {
		goto L75
	} else {
		goto L86
	}
L86:
	;
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v279) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v289 = int32(base.Ui32(v279+int32(_a_F__bt_get_endpoint_8))>>(uint(int32(2))%32)) & int32(_a_F__bt_get_endpoint_9)
	goto L89
L88:
	;
	v289 = int32(0)
	goto L89
L89:
	;
	v319 = v289
	goto L74
L90:
	;
	v293 = v290
	goto L76
L91:
	;
	if v295 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[0]))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v300+(v295^int32(-1))<<(uint(int32(2))%32))))
	v259 = v306
	v262 = v295
	goto L71
L93:
	;
	goto L94
L94:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[1]))
	v259 = v308 + v295<<(uint(int32(13))%32) + int32(-8192)
	v262 = v295
	goto L71
L95:
	;
	v317 = int32(2)
	goto L97
L96:
	;
	v317 = int32(1)
	goto L97
L97:
	;
	v319 = v317
	goto L74
L98:
	;
	if v333 < int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[0]))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v338+(v333^int32(-1))<<(uint(int32(2))%32))))
	v352 = v344
	goto L101
L100:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[1]))
	v352 = v346 + v333<<(uint(int32(13))%32) + int32(-8192)
	goto L101
L101:
	;
	v259 = v352
	v262 = v333
	goto L71
L102:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v357 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_get_endpoint_10), v12+int32(16))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F__bt_get_endpoint_11), int32(2653), int32(_a_F__bt_get_endpoint_12))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v378 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_get_endpoint_13), v12)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F__bt_get_endpoint_11), int32(2666), int32(_a_F__bt_get_endpoint_12))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__bt_insert_parent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int64
	_ = v281
	var v282 int32
	_ = v282
	var v283 int64
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	if l5 != 0 {
		if l2 < int32(0) {
			v25 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v25+(l2^int32(-1))<<(uint(int32(6))%32))+16))
			v40 = v31
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v33+l2<<(uint(int32(6))%32)+int32(-64))+16))
			v40 = v39
		}
		if l3 < int32(0) {
			v44 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+(l3^int32(-1))<<(uint(int32(6))%32))+16))
			v59 = v50
		} else {
			v52 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v52+l3<<(uint(int32(6))%32)+int32(-64))+16))
			v59 = v58
		}
		if l2 < int32(0) {
			v63 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[2]))
			v69 = *(*int32)(unsafe.Add(mBase, uint32(v63+(l2^int32(-1))<<(uint(int32(2))%32))))
			v77 = v69
		} else {
			v71 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[3]))
			v77 = v71 + l2<<(uint(int32(13))%32) + int32(-8192)
		}
		v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+16)))
		v79 = F__bt_allocbuf(m, l0, l1)
		mBase = m.M
		v80 = m.ExcPending
		if v80 != 0 {
			return
		} else {
			if v79 < int32(0) {
				v84 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[2]))
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v84+(v79^int32(-1))<<(uint(int32(2))%32))))
				v98 = v90
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[3]))
				v98 = v92 + v79<<(uint(int32(13))%32) + int32(-8192)
			}
			if v79 < int32(0) {
				v102 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v102+(v79^int32(-1))<<(uint(int32(6))%32))+16))
				v117 = v108
			} else {
				v110 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v110+v79<<(uint(int32(6))%32)+int32(-64))+16))
				v117 = v116
			}
			v120 = F__bt_getbuf(m, l0, int32(0), int32(2))
			mBase = m.M
			v121 = m.ExcPending
			if v121 != 0 {
				return
			} else {
				if v120 < int32(0) {
					v125 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[2]))
					v131 = *(*int32)(unsafe.Add(mBase, uint32(v125+(v120^int32(-1))<<(uint(int32(2))%32))))
					v139 = v131
				} else {
					v133 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[3]))
					v139 = v133 + v120<<(uint(int32(13))%32) + int32(-8192)
				}
				v141 = F_palloc(m, int32(8))
				mBase = m.M
				v142 = m.ExcPending
				if v142 != 0 {
					return
				} else {
					*(*uint16)(unsafe.Add(mBase, uint32(v141)+2)) = uint16(v40)
					v145 = int32(base.Ui32(v40) >> (uint(int32(16)) % 32))
					*(*uint16)(unsafe.Add(mBase, uint32(v141))) = uint16(v145)
					*(*int32)(unsafe.Add(mBase, uint32(v141)+4)) = int32(537395200)
					v149 = *(*int32)(unsafe.Add(mBase, uint32(v77)+24))
					v153 = F_CopyIndexTuple(m, v77+v149&int32(_a_F__bt_insert_parent_0))
					mBase = m.M
					v154 = m.ExcPending
					if v154 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v153))) = base.I32_rotr(v59, int32(16))
						v158 = int32(_a_F__bt_insert_parent_1)
						v160 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4]))
						*(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4])) = v160 + int32(1)
						v164 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
						if base.Ui32(v164) <= base.Ui32(int32(2)) {
							v167 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v139)+64)) = uint8(v167)
							*(*int64)(unsafe.Add(mBase, uint32(v139)+56)) = int64(-4616189618054758400)
							*(*int32)(unsafe.Add(mBase, uint32(v139)+48)) = v167
							*(*int32)(unsafe.Add(mBase, uint32(v139)+28)) = int32(3)
							v175 = int32(72)
							*(*uint16)(unsafe.Add(mBase, uint32(v139)+12)) = uint16(v175)
						} else {
						}
						v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+16)))
						v178 = v98 + v177
						v179 = int32(2)
						*(*uint16)(unsafe.Add(mBase, uint32(v178)+12)) = uint16(v179)
						*(*int64)(unsafe.Add(mBase, uint32(v178))) = int64(0)
						v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+16)))
						v185 = *(*int32)(unsafe.Add(mBase, uint32(v77+v183)+8))
						v186 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v178)+14)) = uint16(v186)
						v188 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v178)+8)) = v185 + v188
						*(*int32)(unsafe.Add(mBase, uint32(v139)+32)) = v117
						v192 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v139)+40)) = v117
						*(*int32)(unsafe.Add(mBase, uint32(v139)+36)) = v192
						v195 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v139)+44)) = v195
						v200 = F_PageAddItemExtended(m, v98, v141, int32(8), v188, v186)
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return
						} else {
							if v200 == int32(0) {
								F_errstart_cold(m, int32(23), int32(0))
								mBase = m.M
								v478 = m.ExcPending
								if v478 != 0 {
									return
								} else {
									if l2 < int32(0) {
										v482 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
										v488 = *(*int32)(unsafe.Add(mBase, uint32(v482+(l2^int32(-1))<<(uint(int32(6))%32))+16))
										v497 = v488
									} else {
										v490 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
										v496 = *(*int32)(unsafe.Add(mBase, uint32(v490+l2<<(uint(int32(6))%32)+int32(-64))+16))
										v497 = v496
									}
									v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v20))) = v497
									*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v498 + int32(4)
									F_errmsg_internal(m, int32(_a_F__bt_insert_parent_2), v20)
									mBase = m.M
									v505 = m.ExcPending
									if v505 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F__bt_insert_parent_3), int32(2534), int32(_a_F__bt_insert_parent_4))
										mBase = m.M
										v510 = m.ExcPending
										if v510 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v208 = F_PageAddItemExtended(m, v98, v153, int32(base.Ui32(v149)>>(uint(int32(17))%32)), int32(2), int32(0))
								mBase = m.M
								v209 = m.ExcPending
								if v209 != 0 {
									return
								} else {
									if v208 == int32(0) {
										F_errstart_cold(m, int32(23), int32(0))
										mBase = m.M
										v514 = m.ExcPending
										if v514 != 0 {
											return
										} else {
											if l2 < int32(0) {
												v518 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
												v524 = *(*int32)(unsafe.Add(mBase, uint32(v518+(l2^int32(-1))<<(uint(int32(6))%32))+16))
												v533 = v524
											} else {
												v526 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
												v532 = *(*int32)(unsafe.Add(mBase, uint32(v526+l2<<(uint(int32(6))%32)+int32(-64))+16))
												v533 = v532
											}
											v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v533
											*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v534 + int32(4)
											F_errmsg_internal(m, int32(_a_F__bt_insert_parent_5), v20+int32(16))
											mBase = m.M
											v543 = m.ExcPending
											if v543 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F__bt_insert_parent_3), int32(2546), int32(_a_F__bt_insert_parent_4))
												mBase = m.M
												v548 = m.ExcPending
												if v548 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v212 = v77 + v78
										v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v212)+12)))
										v215 = v213 & int32(_a_F__bt_insert_parent_6)
										*(*uint16)(unsafe.Add(mBase, uint32(v212)+12)) = uint16(v215)
										F_MarkBufferDirty(m, l2)
										mBase = m.M
										v218 = m.ExcPending
										if v218 != 0 {
											return
										} else {
											F_MarkBufferDirty(m, v79)
											mBase = m.M
											v220 = m.ExcPending
											if v220 != 0 {
												return
											} else {
												F_MarkBufferDirty(m, v120)
												mBase = m.M
												v222 = m.ExcPending
												if v222 != 0 {
													return
												} else {
													v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+118)))
													if v224 != int32(112) {
														v297 = int32(_a_F__bt_insert_parent_1)
														v299 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4]))
														*(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4])) = v299 - int32(1)
														F__bt_relbuf(m, v120)
														mBase = m.M
														v304 = m.ExcPending
														if v304 != 0 {
															return
														} else {
															F_pfree(m, v141)
															mBase = m.M
															v306 = m.ExcPending
															if v306 != 0 {
																return
															} else {
																F_pfree(m, v153)
																mBase = m.M
																v308 = m.ExcPending
																if v308 != 0 {
																	return
																} else {
																	F__bt_relbuf(m, v79)
																	mBase = m.M
																	v310 = m.ExcPending
																	if v310 != 0 {
																		return
																	} else {
																		F__bt_relbuf(m, l3)
																		mBase = m.M
																		v312 = m.ExcPending
																		if v312 != 0 {
																			return
																		} else {
																			F__bt_relbuf(m, l2)
																			mBase = m.M
																			v314 = m.ExcPending
																			if v314 != 0 {
																				return
																			} else {
																				m.G0 = v20 + int32(80)
																				return
																			}
																		}
																	}
																}
															}
														}
													} else {
														v228 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[5]))
														if v228 <= int32(0) {
															v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															if v231 != 0 {
																v297 = int32(_a_F__bt_insert_parent_1)
																v299 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4]))
																*(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4])) = v299 - int32(1)
																F__bt_relbuf(m, v120)
																mBase = m.M
																v304 = m.ExcPending
																if v304 != 0 {
																	return
																} else {
																	F_pfree(m, v141)
																	mBase = m.M
																	v306 = m.ExcPending
																	if v306 != 0 {
																		return
																	} else {
																		F_pfree(m, v153)
																		mBase = m.M
																		v308 = m.ExcPending
																		if v308 != 0 {
																			return
																		} else {
																			F__bt_relbuf(m, v79)
																			mBase = m.M
																			v310 = m.ExcPending
																			if v310 != 0 {
																				return
																			} else {
																				F__bt_relbuf(m, l3)
																				mBase = m.M
																				v312 = m.ExcPending
																				if v312 != 0 {
																					return
																				} else {
																					F__bt_relbuf(m, l2)
																					mBase = m.M
																					v314 = m.ExcPending
																					if v314 != 0 {
																						return
																					} else {
																						m.G0 = v20 + int32(80)
																						return
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																if v232 != 0 {
																	v297 = int32(_a_F__bt_insert_parent_1)
																	v299 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4]))
																	*(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4])) = v299 - int32(1)
																	F__bt_relbuf(m, v120)
																	mBase = m.M
																	v304 = m.ExcPending
																	if v304 != 0 {
																		return
																	} else {
																		F_pfree(m, v141)
																		mBase = m.M
																		v306 = m.ExcPending
																		if v306 != 0 {
																			return
																		} else {
																			F_pfree(m, v153)
																			mBase = m.M
																			v308 = m.ExcPending
																			if v308 != 0 {
																				return
																			} else {
																				F__bt_relbuf(m, v79)
																				mBase = m.M
																				v310 = m.ExcPending
																				if v310 != 0 {
																					return
																				} else {
																					F__bt_relbuf(m, l3)
																					mBase = m.M
																					v312 = m.ExcPending
																					if v312 != 0 {
																						return
																					} else {
																						F__bt_relbuf(m, l2)
																						mBase = m.M
																						v314 = m.ExcPending
																						if v314 != 0 {
																							return
																						} else {
																							m.G0 = v20 + int32(80)
																							return
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v117
																	v234 = *(*int32)(unsafe.Add(mBase, uint32(v139)+36))
																	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v234
																	F_XLogBeginInsert(m)
																	mBase = m.M
																	v237 = m.ExcPending
																	if v237 != 0 {
																		return
																	} else {
																		F_XLogRegisterData(m, v20+int32(72), int32(8))
																		mBase = m.M
																		v242 = m.ExcPending
																		if v242 != 0 {
																			return
																		} else {
																			F_XLogRegisterBuffer(m, int32(0), v79, int32(6))
																			mBase = m.M
																			v246 = m.ExcPending
																			if v246 != 0 {
																				return
																			} else {
																				F_XLogRegisterBuffer(m, int32(1), l2, int32(8))
																				mBase = m.M
																				v250 = m.ExcPending
																				if v250 != 0 {
																					return
																				} else {
																					F_XLogRegisterBuffer(m, int32(2), v120, int32(14))
																					mBase = m.M
																					v254 = m.ExcPending
																					if v254 != 0 {
																						return
																					} else {
																						v255 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
																						*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v117
																						*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v255
																						v258 = *(*int32)(unsafe.Add(mBase, uint32(v139)+36))
																						*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v258
																						*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v117
																						*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v258
																						v262 = *(*int32)(unsafe.Add(mBase, uint32(v139)+48))
																						*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v262
																						v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+64)))
																						*(*uint8)(unsafe.Add(mBase, uint32(v20)+68)) = uint8(v264)
																						F_XLogRegisterBufData(m, int32(2), v20+int32(44), int32(28))
																						mBase = m.M
																						v271 = m.ExcPending
																						if v271 != 0 {
																							return
																						} else {
																							v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+14)))
																							v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+16)))
																							F_XLogRegisterBufData(m, int32(0), v98+v273, v275-v273)
																							mBase = m.M
																							v278 = m.ExcPending
																							if v278 != 0 {
																								return
																							} else {
																								v281 = F_XLogInsert(m, int32(11), int32(160))
																								mBase = m.M
																								v282 = m.ExcPending
																								if v282 != 0 {
																									return
																								} else {
																									v283 = int64(32)
																									*(*int64)(unsafe.Add(mBase, uint32(v77))) = base.I64_rotr(v281, v283)
																									v286 = base.I32_wrap_i64(v281)
																									*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v286
																									v290 = base.I32_wrap_i64(int64(base.Ui64(v281) >> (uint(v283) % 64)))
																									*(*int32)(unsafe.Add(mBase, uint32(v98))) = v290
																									*(*int32)(unsafe.Add(mBase, uint32(v139)+4)) = v286
																									*(*int32)(unsafe.Add(mBase, uint32(v139))) = v290
																									v297 = int32(_a_F__bt_insert_parent_1)
																									v299 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4]))
																									*(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4])) = v299 - int32(1)
																									F__bt_relbuf(m, v120)
																									mBase = m.M
																									v304 = m.ExcPending
																									if v304 != 0 {
																										return
																									} else {
																										F_pfree(m, v141)
																										mBase = m.M
																										v306 = m.ExcPending
																										if v306 != 0 {
																											return
																										} else {
																											F_pfree(m, v153)
																											mBase = m.M
																											v308 = m.ExcPending
																											if v308 != 0 {
																												return
																											} else {
																												F__bt_relbuf(m, v79)
																												mBase = m.M
																												v310 = m.ExcPending
																												if v310 != 0 {
																													return
																												} else {
																													F__bt_relbuf(m, l3)
																													mBase = m.M
																													v312 = m.ExcPending
																													if v312 != 0 {
																														return
																													} else {
																														F__bt_relbuf(m, l2)
																														mBase = m.M
																														v314 = m.ExcPending
																														if v314 != 0 {
																															return
																														} else {
																															m.G0 = v20 + int32(80)
																															return
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
																				}
																			}
																		}
																	}
																}
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v117
															v234 = *(*int32)(unsafe.Add(mBase, uint32(v139)+36))
															*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v234
															F_XLogBeginInsert(m)
															mBase = m.M
															v237 = m.ExcPending
															if v237 != 0 {
																return
															} else {
																F_XLogRegisterData(m, v20+int32(72), int32(8))
																mBase = m.M
																v242 = m.ExcPending
																if v242 != 0 {
																	return
																} else {
																	F_XLogRegisterBuffer(m, int32(0), v79, int32(6))
																	mBase = m.M
																	v246 = m.ExcPending
																	if v246 != 0 {
																		return
																	} else {
																		F_XLogRegisterBuffer(m, int32(1), l2, int32(8))
																		mBase = m.M
																		v250 = m.ExcPending
																		if v250 != 0 {
																			return
																		} else {
																			F_XLogRegisterBuffer(m, int32(2), v120, int32(14))
																			mBase = m.M
																			v254 = m.ExcPending
																			if v254 != 0 {
																				return
																			} else {
																				v255 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
																				*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v117
																				*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v255
																				v258 = *(*int32)(unsafe.Add(mBase, uint32(v139)+36))
																				*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v258
																				*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v117
																				*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v258
																				v262 = *(*int32)(unsafe.Add(mBase, uint32(v139)+48))
																				*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v262
																				v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+64)))
																				*(*uint8)(unsafe.Add(mBase, uint32(v20)+68)) = uint8(v264)
																				F_XLogRegisterBufData(m, int32(2), v20+int32(44), int32(28))
																				mBase = m.M
																				v271 = m.ExcPending
																				if v271 != 0 {
																					return
																				} else {
																					v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+14)))
																					v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+16)))
																					F_XLogRegisterBufData(m, int32(0), v98+v273, v275-v273)
																					mBase = m.M
																					v278 = m.ExcPending
																					if v278 != 0 {
																						return
																					} else {
																						v281 = F_XLogInsert(m, int32(11), int32(160))
																						mBase = m.M
																						v282 = m.ExcPending
																						if v282 != 0 {
																							return
																						} else {
																							v283 = int64(32)
																							*(*int64)(unsafe.Add(mBase, uint32(v77))) = base.I64_rotr(v281, v283)
																							v286 = base.I32_wrap_i64(v281)
																							*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v286
																							v290 = base.I32_wrap_i64(int64(base.Ui64(v281) >> (uint(v283) % 64)))
																							*(*int32)(unsafe.Add(mBase, uint32(v98))) = v290
																							*(*int32)(unsafe.Add(mBase, uint32(v139)+4)) = v286
																							*(*int32)(unsafe.Add(mBase, uint32(v139))) = v290
																							v297 = int32(_a_F__bt_insert_parent_1)
																							v299 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4]))
																							*(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4])) = v299 - int32(1)
																							F__bt_relbuf(m, v120)
																							mBase = m.M
																							v304 = m.ExcPending
																							if v304 != 0 {
																								return
																							} else {
																								F_pfree(m, v141)
																								mBase = m.M
																								v306 = m.ExcPending
																								if v306 != 0 {
																									return
																								} else {
																									F_pfree(m, v153)
																									mBase = m.M
																									v308 = m.ExcPending
																									if v308 != 0 {
																										return
																									} else {
																										F__bt_relbuf(m, v79)
																										mBase = m.M
																										v310 = m.ExcPending
																										if v310 != 0 {
																											return
																										} else {
																											F__bt_relbuf(m, l3)
																											mBase = m.M
																											v312 = m.ExcPending
																											if v312 != 0 {
																												return
																											} else {
																												F__bt_relbuf(m, l2)
																												mBase = m.M
																												v314 = m.ExcPending
																												if v314 != 0 {
																													return
																												} else {
																													m.G0 = v20 + int32(80)
																													return
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
								}
							}
						}
					}
				}
			}
		}
	} else {
		if l2 < int32(0) {
			v318 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
			v324 = *(*int32)(unsafe.Add(mBase, uint32(v318+(l2^int32(-1))<<(uint(int32(6))%32))+16))
			v333 = v324
		} else {
			v326 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
			v332 = *(*int32)(unsafe.Add(mBase, uint32(v326+l2<<(uint(int32(6))%32)+int32(-64))+16))
			v333 = v332
		}
		if l3 < int32(0) {
			v337 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
			v343 = *(*int32)(unsafe.Add(mBase, uint32(v337+(l3^int32(-1))<<(uint(int32(6))%32))+16))
			v352 = v343
		} else {
			v345 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
			v351 = *(*int32)(unsafe.Add(mBase, uint32(v345+l3<<(uint(int32(6))%32)+int32(-64))+16))
			v352 = v351
		}
		if l2 < int32(0) {
			v356 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[2]))
			v362 = *(*int32)(unsafe.Add(mBase, uint32(v356+(l2^int32(-1))<<(uint(int32(2))%32))))
			v370 = v362
		} else {
			v364 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[3]))
			v370 = v364 + l2<<(uint(int32(13))%32) + int32(-8192)
		}
		if l4 == int32(0) {
			v375 = F_errstart(m, int32(13), int32(0))
			mBase = m.M
			v376 = m.ExcPending
			if v376 != 0 {
				return
			} else {
				if v375 != 0 {
					F_errmsg_internal(m, int32(_a_F__bt_insert_parent_7), int32(0))
					mBase = m.M
					v380 = m.ExcPending
					if v380 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F__bt_insert_parent_3), int32(2149), int32(_a_F__bt_insert_parent_8))
						mBase = m.M
						v385 = m.ExcPending
						if v385 != 0 {
							return
						} else {
							v386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v370)+16)))
							v388 = *(*int32)(unsafe.Add(mBase, uint32(v370+v386)+8))
							v392 = F__bt_get_endpoint(m, l0, v388+int32(1), int32(0))
							mBase = m.M
							v393 = m.ExcPending
							if v393 != 0 {
								return
							} else {
								if v392 < int32(0) {
									v397 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
									v403 = *(*int32)(unsafe.Add(mBase, uint32(v397+(v392^int32(-1))<<(uint(int32(6))%32))+16))
									v412 = v403
								} else {
									v405 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
									v411 = *(*int32)(unsafe.Add(mBase, uint32(v405+v392<<(uint(int32(6))%32)+int32(-64))+16))
									v412 = v411
								}
								v413 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v413
								*(*uint16)(unsafe.Add(mBase, uint32(v20)+48)) = uint16(v413)
								*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v412
								F__bt_relbuf(m, v392)
								mBase = m.M
								v419 = m.ExcPending
								if v419 != 0 {
									return
								} else {
									v422 = v20 + int32(44)
									v424 = *(*int32)(unsafe.Add(mBase, uint32(v370)+24))
									v428 = F_CopyIndexTuple(m, v370+v424&int32(_a_F__bt_insert_parent_0))
									mBase = m.M
									v429 = m.ExcPending
									if v429 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v428))) = base.I32_rotr(v352, int32(16))
										v433 = F__bt_getstackbuf(m, l0, l1, v422, v333)
										mBase = m.M
										v434 = m.ExcPending
										if v434 != 0 {
											return
										} else {
											F__bt_relbuf(m, l3)
											mBase = m.M
											v436 = m.ExcPending
											if v436 != 0 {
												return
											} else {
												if v433 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v552 = m.ExcPending
													if v552 != 0 {
														return
													} else {
														F_errcode(m, int32(33557032))
														mBase = m.M
														v555 = m.ExcPending
														if v555 != 0 {
															return
														} else {
															v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v352
															*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v333
															*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v556 + int32(4)
															F_errmsg_internal(m, int32(_a_F__bt_insert_parent_9), v20+int32(32))
															mBase = m.M
															v566 = m.ExcPending
															if v566 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F__bt_insert_parent_3), int32(2215), int32(_a_F__bt_insert_parent_8))
																mBase = m.M
																v571 = m.ExcPending
																if v571 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v439 = int32(0)
													v440 = *(*int32)(unsafe.Add(mBase, uint32(v422)+8))
													v441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v428)+6)))
													v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v422)+4)))
													F__bt_insertonpg(m, l0, l1, v439, v433, l2, v440, v428, (v441&int32(_a_F__bt_insert_parent_10)+int32(7))&int32(_a_F__bt_insert_parent_11), (v448+int32(1))&int32(_a_F__bt_insert_parent_12), v439, l6)
													mBase = m.M
													v455 = m.ExcPending
													if v455 != 0 {
														return
													} else {
														F_pfree(m, v428)
														mBase = m.M
														v457 = m.ExcPending
														if v457 != 0 {
															return
														} else {
															m.G0 = v20 + int32(80)
															return
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
					v386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v370)+16)))
					v388 = *(*int32)(unsafe.Add(mBase, uint32(v370+v386)+8))
					v392 = F__bt_get_endpoint(m, l0, v388+int32(1), int32(0))
					mBase = m.M
					v393 = m.ExcPending
					if v393 != 0 {
						return
					} else {
						if v392 < int32(0) {
							v397 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
							v403 = *(*int32)(unsafe.Add(mBase, uint32(v397+(v392^int32(-1))<<(uint(int32(6))%32))+16))
							v412 = v403
						} else {
							v405 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
							v411 = *(*int32)(unsafe.Add(mBase, uint32(v405+v392<<(uint(int32(6))%32)+int32(-64))+16))
							v412 = v411
						}
						v413 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v413
						*(*uint16)(unsafe.Add(mBase, uint32(v20)+48)) = uint16(v413)
						*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v412
						F__bt_relbuf(m, v392)
						mBase = m.M
						v419 = m.ExcPending
						if v419 != 0 {
							return
						} else {
							v422 = v20 + int32(44)
							v424 = *(*int32)(unsafe.Add(mBase, uint32(v370)+24))
							v428 = F_CopyIndexTuple(m, v370+v424&int32(_a_F__bt_insert_parent_0))
							mBase = m.M
							v429 = m.ExcPending
							if v429 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v428))) = base.I32_rotr(v352, int32(16))
								v433 = F__bt_getstackbuf(m, l0, l1, v422, v333)
								mBase = m.M
								v434 = m.ExcPending
								if v434 != 0 {
									return
								} else {
									F__bt_relbuf(m, l3)
									mBase = m.M
									v436 = m.ExcPending
									if v436 != 0 {
										return
									} else {
										if v433 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v552 = m.ExcPending
											if v552 != 0 {
												return
											} else {
												F_errcode(m, int32(33557032))
												mBase = m.M
												v555 = m.ExcPending
												if v555 != 0 {
													return
												} else {
													v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v352
													*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v333
													*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v556 + int32(4)
													F_errmsg_internal(m, int32(_a_F__bt_insert_parent_9), v20+int32(32))
													mBase = m.M
													v566 = m.ExcPending
													if v566 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F__bt_insert_parent_3), int32(2215), int32(_a_F__bt_insert_parent_8))
														mBase = m.M
														v571 = m.ExcPending
														if v571 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v439 = int32(0)
											v440 = *(*int32)(unsafe.Add(mBase, uint32(v422)+8))
											v441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v428)+6)))
											v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v422)+4)))
											F__bt_insertonpg(m, l0, l1, v439, v433, l2, v440, v428, (v441&int32(_a_F__bt_insert_parent_10)+int32(7))&int32(_a_F__bt_insert_parent_11), (v448+int32(1))&int32(_a_F__bt_insert_parent_12), v439, l6)
											mBase = m.M
											v455 = m.ExcPending
											if v455 != 0 {
												return
											} else {
												F_pfree(m, v428)
												mBase = m.M
												v457 = m.ExcPending
												if v457 != 0 {
													return
												} else {
													m.G0 = v20 + int32(80)
													return
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
			v422 = l4
			v424 = *(*int32)(unsafe.Add(mBase, uint32(v370)+24))
			v428 = F_CopyIndexTuple(m, v370+v424&int32(_a_F__bt_insert_parent_0))
			mBase = m.M
			v429 = m.ExcPending
			if v429 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v428))) = base.I32_rotr(v352, int32(16))
				v433 = F__bt_getstackbuf(m, l0, l1, v422, v333)
				mBase = m.M
				v434 = m.ExcPending
				if v434 != 0 {
					return
				} else {
					F__bt_relbuf(m, l3)
					mBase = m.M
					v436 = m.ExcPending
					if v436 != 0 {
						return
					} else {
						if v433 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v552 = m.ExcPending
							if v552 != 0 {
								return
							} else {
								F_errcode(m, int32(33557032))
								mBase = m.M
								v555 = m.ExcPending
								if v555 != 0 {
									return
								} else {
									v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v352
									*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v333
									*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v556 + int32(4)
									F_errmsg_internal(m, int32(_a_F__bt_insert_parent_9), v20+int32(32))
									mBase = m.M
									v566 = m.ExcPending
									if v566 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F__bt_insert_parent_3), int32(2215), int32(_a_F__bt_insert_parent_8))
										mBase = m.M
										v571 = m.ExcPending
										if v571 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v439 = int32(0)
							v440 = *(*int32)(unsafe.Add(mBase, uint32(v422)+8))
							v441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v428)+6)))
							v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v422)+4)))
							F__bt_insertonpg(m, l0, l1, v439, v433, l2, v440, v428, (v441&int32(_a_F__bt_insert_parent_10)+int32(7))&int32(_a_F__bt_insert_parent_11), (v448+int32(1))&int32(_a_F__bt_insert_parent_12), v439, l6)
							mBase = m.M
							v455 = m.ExcPending
							if v455 != 0 {
								return
							} else {
								F_pfree(m, v428)
								mBase = m.M
								v457 = m.ExcPending
								if v457 != 0 {
									return
								} else {
									m.G0 = v20 + int32(80)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F__bt_parallel_done(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v3 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)))
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v3)+28))
			v9 = v3 + v8
			v11 = v9 + int32(12)
			v13 = F_LWLockAcquire(m, v11, int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				if v15 != int32(4) {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(4)
					F_LWLockRelease(m, v11)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						F_ConditionVariableBroadcast(m, v9+int32(28))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					F_LWLockRelease(m, v11)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F__bt_parallel_scan_and_sort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 float64
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 float64
	_ = v195
	var v196 float64
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v16 = F_palloc0(m, int32(12))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(-1)
		v21 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v21)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
		v27 = F_tuplesort_begin_index_btree(m, v23, v24, v25, v26, l5, v16)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v27
			if l1 != 0 {
				v31 = F_palloc0(m, int32(12))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = int32(-1)
					v36 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v36)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v40 = int32(0)
					v43 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[0]))
					if l5 < v43 {
						v45 = l5
					} else {
						v45 = v43
					}
					v46 = F_tuplesort_begin_index_btree(m, v38, v39, v40, v40, v45, v31)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v46
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
						*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v51)
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
						v54 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)) = uint8(v54)
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)) = uint8(v53)
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v54
						*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v57
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v66 = F_BuildIndexInfo(m, v65)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
							*(*uint8)(unsafe.Add(mBase, uint32(v66)+121)) = uint8(v68)
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v73 = F_table_beginscan_parallel(m, v70, l2+int32(96))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v78 = int32(0)
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+140))
								v84 = m.T0[v83].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v75, v76, v66, int32(1), v78, l6, v78, int32(-1), int32(238), v13, v73)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									if l6 != 0 {
										v90 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[1]))
										if v90 == int32(0) {
										} else {
											v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[2])))
											if v94&int32(1) == int32(0) {
											} else {
												v99 = int32(_a_F__bt_parallel_scan_and_sort_0)
												v101 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
												v102 = int32(1)
												*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v101 + v102
												v105 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
												*(*int32)(unsafe.Add(mBase, uint32(v90))) = v105 + v102
												*(*int64)(unsafe.Add(mBase, uint32(v90+int32(80))+232)) = int64(3)
												v113 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
												*(*int32)(unsafe.Add(mBase, uint32(v90))) = v113 + v102
												v119 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
												*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v119 - v102
											}
										}
										v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										F_tuplesort_performsort(m, v123)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return
										} else {
											if l1 == int32(0) {
												v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
												*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
												if v173 != 0 {
													F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_1), int32(1951), int32(_a_F__bt_parallel_scan_and_sort_2))
													mBase = m.M
													v182 = m.ExcPending
													if v182 != 0 {
														return
													} else {
														v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
														v184 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
														v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
														v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
														if v190 == v184 {
															v193 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
														} else {
														}
														v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
														v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
														v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
														if v199 == int32(1) {
															v202 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
														} else {
														}
														*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
														F_ConditionVariableSignal(m, l2+int32(24))
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return
														} else {
															v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_tuplesort_end(m, v210)
															mBase = m.M
															v212 = m.ExcPending
															if v212 != 0 {
																return
															} else {
																if l1 != 0 {
																	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																	F_tuplesort_end(m, v213)
																	mBase = m.M
																	v215 = m.ExcPending
																	if v215 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(32)
																		return
																	}
																} else {
																	m.G0 = v13 + int32(32)
																	return
																}
															}
														}
													}
												} else {
													v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
													v184 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
													v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
													v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
													if v190 == v184 {
														v193 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
													} else {
													}
													v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
													v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
													v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
													if v199 == int32(1) {
														v202 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
													} else {
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
													F_ConditionVariableSignal(m, l2+int32(24))
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return
													} else {
														v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_tuplesort_end(m, v210)
														mBase = m.M
														v212 = m.ExcPending
														if v212 != 0 {
															return
														} else {
															if l1 != 0 {
																v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																F_tuplesort_end(m, v213)
																mBase = m.M
																v215 = m.ExcPending
																if v215 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(32)
																	return
																}
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														}
													}
												}
											} else {
												v132 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[1]))
												if v132 == int32(0) {
												} else {
													v136 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[2])))
													if v136&int32(1) == int32(0) {
													} else {
														v141 = int32(_a_F__bt_parallel_scan_and_sort_0)
														v143 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
														v144 = int32(1)
														*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v143 + v144
														v147 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
														*(*int32)(unsafe.Add(mBase, uint32(v132))) = v147 + v144
														*(*int64)(unsafe.Add(mBase, uint32(v132+int32(80))+232)) = int64(4)
														v155 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
														*(*int32)(unsafe.Add(mBase, uint32(v132))) = v155 + v144
														v161 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
														*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v161 - v144
													}
												}
												v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												F_tuplesort_performsort(m, v170)
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return
												} else {
													v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
													*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
													if v173 != 0 {
														F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_1), int32(1951), int32(_a_F__bt_parallel_scan_and_sort_2))
														mBase = m.M
														v182 = m.ExcPending
														if v182 != 0 {
															return
														} else {
															v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
															v184 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
															v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
															*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
															v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
															if v190 == v184 {
																v193 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
															} else {
															}
															v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
															v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
															*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
															v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
															if v199 == int32(1) {
																v202 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
															} else {
															}
															*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
															F_ConditionVariableSignal(m, l2+int32(24))
															mBase = m.M
															v209 = m.ExcPending
															if v209 != 0 {
																return
															} else {
																v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																F_tuplesort_end(m, v210)
																mBase = m.M
																v212 = m.ExcPending
																if v212 != 0 {
																	return
																} else {
																	if l1 != 0 {
																		v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																		F_tuplesort_end(m, v213)
																		mBase = m.M
																		v215 = m.ExcPending
																		if v215 != 0 {
																			return
																		} else {
																			m.G0 = v13 + int32(32)
																			return
																		}
																	} else {
																		m.G0 = v13 + int32(32)
																		return
																	}
																}
															}
														}
													} else {
														v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
														v184 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
														v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
														v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
														if v190 == v184 {
															v193 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
														} else {
														}
														v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
														v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
														v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
														if v199 == int32(1) {
															v202 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
														} else {
														}
														*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
														F_ConditionVariableSignal(m, l2+int32(24))
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return
														} else {
															v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_tuplesort_end(m, v210)
															mBase = m.M
															v212 = m.ExcPending
															if v212 != 0 {
																return
															} else {
																if l1 != 0 {
																	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																	F_tuplesort_end(m, v213)
																	mBase = m.M
																	v215 = m.ExcPending
																	if v215 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(32)
																		return
																	}
																} else {
																	m.G0 = v13 + int32(32)
																	return
																}
															}
														}
													}
												}
											}
										}
									} else {
										v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										F_tuplesort_performsort(m, v165)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return
										} else {
											if l1 == int32(0) {
												v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
												*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
												if v173 != 0 {
													F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_1), int32(1951), int32(_a_F__bt_parallel_scan_and_sort_2))
													mBase = m.M
													v182 = m.ExcPending
													if v182 != 0 {
														return
													} else {
														v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
														v184 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
														v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
														v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
														if v190 == v184 {
															v193 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
														} else {
														}
														v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
														v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
														v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
														if v199 == int32(1) {
															v202 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
														} else {
														}
														*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
														F_ConditionVariableSignal(m, l2+int32(24))
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return
														} else {
															v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_tuplesort_end(m, v210)
															mBase = m.M
															v212 = m.ExcPending
															if v212 != 0 {
																return
															} else {
																if l1 != 0 {
																	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																	F_tuplesort_end(m, v213)
																	mBase = m.M
																	v215 = m.ExcPending
																	if v215 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(32)
																		return
																	}
																} else {
																	m.G0 = v13 + int32(32)
																	return
																}
															}
														}
													}
												} else {
													v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
													v184 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
													v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
													v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
													if v190 == v184 {
														v193 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
													} else {
													}
													v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
													v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
													v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
													if v199 == int32(1) {
														v202 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
													} else {
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
													F_ConditionVariableSignal(m, l2+int32(24))
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return
													} else {
														v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_tuplesort_end(m, v210)
														mBase = m.M
														v212 = m.ExcPending
														if v212 != 0 {
															return
														} else {
															if l1 != 0 {
																v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																F_tuplesort_end(m, v213)
																mBase = m.M
																v215 = m.ExcPending
																if v215 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(32)
																	return
																}
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														}
													}
												}
											} else {
												v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												F_tuplesort_performsort(m, v170)
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return
												} else {
													v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
													*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
													if v173 != 0 {
														F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_1), int32(1951), int32(_a_F__bt_parallel_scan_and_sort_2))
														mBase = m.M
														v182 = m.ExcPending
														if v182 != 0 {
															return
														} else {
															v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
															v184 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
															v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
															*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
															v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
															if v190 == v184 {
																v193 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
															} else {
															}
															v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
															v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
															*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
															v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
															if v199 == int32(1) {
																v202 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
															} else {
															}
															*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
															F_ConditionVariableSignal(m, l2+int32(24))
															mBase = m.M
															v209 = m.ExcPending
															if v209 != 0 {
																return
															} else {
																v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																F_tuplesort_end(m, v210)
																mBase = m.M
																v212 = m.ExcPending
																if v212 != 0 {
																	return
																} else {
																	if l1 != 0 {
																		v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																		F_tuplesort_end(m, v213)
																		mBase = m.M
																		v215 = m.ExcPending
																		if v215 != 0 {
																			return
																		} else {
																			m.G0 = v13 + int32(32)
																			return
																		}
																	} else {
																		m.G0 = v13 + int32(32)
																		return
																	}
																}
															}
														}
													} else {
														v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
														v184 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
														v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
														v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
														if v190 == v184 {
															v193 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
														} else {
														}
														v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
														v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
														v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
														if v199 == int32(1) {
															v202 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
														} else {
														}
														*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
														F_ConditionVariableSignal(m, l2+int32(24))
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return
														} else {
															v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_tuplesort_end(m, v210)
															mBase = m.M
															v212 = m.ExcPending
															if v212 != 0 {
																return
															} else {
																if l1 != 0 {
																	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																	F_tuplesort_end(m, v213)
																	mBase = m.M
																	v215 = m.ExcPending
																	if v215 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(32)
																		return
																	}
																} else {
																	m.G0 = v13 + int32(32)
																	return
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
						}
					}
				}
			} else {
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
				*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v51)
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
				v54 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)) = uint8(v54)
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)) = uint8(v53)
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v54
				*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v57
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v66 = F_BuildIndexInfo(m, v65)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
					*(*uint8)(unsafe.Add(mBase, uint32(v66)+121)) = uint8(v68)
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v73 = F_table_beginscan_parallel(m, v70, l2+int32(96))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v78 = int32(0)
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+140))
						v84 = m.T0[v83].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v75, v76, v66, int32(1), v78, l6, v78, int32(-1), int32(238), v13, v73)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							if l6 != 0 {
								v90 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[1]))
								if v90 == int32(0) {
								} else {
									v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[2])))
									if v94&int32(1) == int32(0) {
									} else {
										v99 = int32(_a_F__bt_parallel_scan_and_sort_0)
										v101 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
										v102 = int32(1)
										*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v101 + v102
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
										*(*int32)(unsafe.Add(mBase, uint32(v90))) = v105 + v102
										*(*int64)(unsafe.Add(mBase, uint32(v90+int32(80))+232)) = int64(3)
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
										*(*int32)(unsafe.Add(mBase, uint32(v90))) = v113 + v102
										v119 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
										*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v119 - v102
									}
								}
								v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_tuplesort_performsort(m, v123)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return
								} else {
									if l1 == int32(0) {
										v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
										*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
										if v173 != 0 {
											F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_1), int32(1951), int32(_a_F__bt_parallel_scan_and_sort_2))
											mBase = m.M
											v182 = m.ExcPending
											if v182 != 0 {
												return
											} else {
												v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
												v184 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
												v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
												v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
												if v190 == v184 {
													v193 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
												} else {
												}
												v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
												v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
												v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
												if v199 == int32(1) {
													v202 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
												} else {
												}
												*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
												F_ConditionVariableSignal(m, l2+int32(24))
												mBase = m.M
												v209 = m.ExcPending
												if v209 != 0 {
													return
												} else {
													v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_tuplesort_end(m, v210)
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return
													} else {
														if l1 != 0 {
															v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
															F_tuplesort_end(m, v213)
															mBase = m.M
															v215 = m.ExcPending
															if v215 != 0 {
																return
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														} else {
															m.G0 = v13 + int32(32)
															return
														}
													}
												}
											}
										} else {
											v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
											v184 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
											v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
											*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
											v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
											if v190 == v184 {
												v193 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
											} else {
											}
											v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
											v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
											*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
											v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
											if v199 == int32(1) {
												v202 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
											} else {
											}
											*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
											F_ConditionVariableSignal(m, l2+int32(24))
											mBase = m.M
											v209 = m.ExcPending
											if v209 != 0 {
												return
											} else {
												v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_tuplesort_end(m, v210)
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return
												} else {
													if l1 != 0 {
														v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
														F_tuplesort_end(m, v213)
														mBase = m.M
														v215 = m.ExcPending
														if v215 != 0 {
															return
														} else {
															m.G0 = v13 + int32(32)
															return
														}
													} else {
														m.G0 = v13 + int32(32)
														return
													}
												}
											}
										}
									} else {
										v132 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[1]))
										if v132 == int32(0) {
										} else {
											v136 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[2])))
											if v136&int32(1) == int32(0) {
											} else {
												v141 = int32(_a_F__bt_parallel_scan_and_sort_0)
												v143 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
												v144 = int32(1)
												*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v143 + v144
												v147 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
												*(*int32)(unsafe.Add(mBase, uint32(v132))) = v147 + v144
												*(*int64)(unsafe.Add(mBase, uint32(v132+int32(80))+232)) = int64(4)
												v155 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
												*(*int32)(unsafe.Add(mBase, uint32(v132))) = v155 + v144
												v161 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
												*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v161 - v144
											}
										}
										v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										F_tuplesort_performsort(m, v170)
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return
										} else {
											v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
											*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
											if v173 != 0 {
												F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_1), int32(1951), int32(_a_F__bt_parallel_scan_and_sort_2))
												mBase = m.M
												v182 = m.ExcPending
												if v182 != 0 {
													return
												} else {
													v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
													v184 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
													v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
													v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
													if v190 == v184 {
														v193 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
													} else {
													}
													v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
													v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
													v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
													if v199 == int32(1) {
														v202 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
													} else {
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
													F_ConditionVariableSignal(m, l2+int32(24))
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return
													} else {
														v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_tuplesort_end(m, v210)
														mBase = m.M
														v212 = m.ExcPending
														if v212 != 0 {
															return
														} else {
															if l1 != 0 {
																v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																F_tuplesort_end(m, v213)
																mBase = m.M
																v215 = m.ExcPending
																if v215 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(32)
																	return
																}
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														}
													}
												}
											} else {
												v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
												v184 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
												v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
												v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
												if v190 == v184 {
													v193 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
												} else {
												}
												v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
												v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
												v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
												if v199 == int32(1) {
													v202 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
												} else {
												}
												*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
												F_ConditionVariableSignal(m, l2+int32(24))
												mBase = m.M
												v209 = m.ExcPending
												if v209 != 0 {
													return
												} else {
													v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_tuplesort_end(m, v210)
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return
													} else {
														if l1 != 0 {
															v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
															F_tuplesort_end(m, v213)
															mBase = m.M
															v215 = m.ExcPending
															if v215 != 0 {
																return
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														} else {
															m.G0 = v13 + int32(32)
															return
														}
													}
												}
											}
										}
									}
								}
							} else {
								v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_tuplesort_performsort(m, v165)
								mBase = m.M
								v167 = m.ExcPending
								if v167 != 0 {
									return
								} else {
									if l1 == int32(0) {
										v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
										*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
										if v173 != 0 {
											F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_1), int32(1951), int32(_a_F__bt_parallel_scan_and_sort_2))
											mBase = m.M
											v182 = m.ExcPending
											if v182 != 0 {
												return
											} else {
												v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
												v184 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
												v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
												v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
												if v190 == v184 {
													v193 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
												} else {
												}
												v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
												v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
												v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
												if v199 == int32(1) {
													v202 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
												} else {
												}
												*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
												F_ConditionVariableSignal(m, l2+int32(24))
												mBase = m.M
												v209 = m.ExcPending
												if v209 != 0 {
													return
												} else {
													v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_tuplesort_end(m, v210)
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return
													} else {
														if l1 != 0 {
															v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
															F_tuplesort_end(m, v213)
															mBase = m.M
															v215 = m.ExcPending
															if v215 != 0 {
																return
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														} else {
															m.G0 = v13 + int32(32)
															return
														}
													}
												}
											}
										} else {
											v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
											v184 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
											v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
											*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
											v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
											if v190 == v184 {
												v193 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
											} else {
											}
											v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
											v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
											*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
											v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
											if v199 == int32(1) {
												v202 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
											} else {
											}
											*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
											F_ConditionVariableSignal(m, l2+int32(24))
											mBase = m.M
											v209 = m.ExcPending
											if v209 != 0 {
												return
											} else {
												v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_tuplesort_end(m, v210)
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return
												} else {
													if l1 != 0 {
														v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
														F_tuplesort_end(m, v213)
														mBase = m.M
														v215 = m.ExcPending
														if v215 != 0 {
															return
														} else {
															m.G0 = v13 + int32(32)
															return
														}
													} else {
														m.G0 = v13 + int32(32)
														return
													}
												}
											}
										}
									} else {
										v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										F_tuplesort_performsort(m, v170)
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return
										} else {
											v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
											*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
											if v173 != 0 {
												F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_1), int32(1951), int32(_a_F__bt_parallel_scan_and_sort_2))
												mBase = m.M
												v182 = m.ExcPending
												if v182 != 0 {
													return
												} else {
													v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
													v184 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
													v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
													v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
													if v190 == v184 {
														v193 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
													} else {
													}
													v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
													v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
													v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
													if v199 == int32(1) {
														v202 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
													} else {
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
													F_ConditionVariableSignal(m, l2+int32(24))
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return
													} else {
														v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_tuplesort_end(m, v210)
														mBase = m.M
														v212 = m.ExcPending
														if v212 != 0 {
															return
														} else {
															if l1 != 0 {
																v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																F_tuplesort_end(m, v213)
																mBase = m.M
																v215 = m.ExcPending
																if v215 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(32)
																	return
																}
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														}
													}
												}
											} else {
												v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
												v184 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183 + v184
												v187 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v187)
												v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
												if v190 == v184 {
													v193 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v193)
												} else {
												}
												v195 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
												v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v195, v196)
												v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
												if v199 == int32(1) {
													v202 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v202)
												} else {
												}
												*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
												F_ConditionVariableSignal(m, l2+int32(24))
												mBase = m.M
												v209 = m.ExcPending
												if v209 != 0 {
													return
												} else {
													v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_tuplesort_end(m, v210)
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return
													} else {
														if l1 != 0 {
															v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
															F_tuplesort_end(m, v213)
															mBase = m.M
															v215 = m.ExcPending
															if v215 != 0 {
																return
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														} else {
															m.G0 = v13 + int32(32)
															return
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
				}
			}
		}
	}
}
func F__bt_relbuf(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	F_LockBuffer(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_ReleaseBuffer(m, l0)
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F__bt_reorder_array_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(v4 < v3) - base.B2i32(v3 < v4)
}
func F__bt_saoparray_shrink(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v7 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+212))
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v7
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if base.B2i32(v24 == v27)|base.B2i32(v27 == v7) != 0 {
		v56 = l3
		v57 = int32(0)
		v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
		v63 = F__bt_binsrch_array_skey(m, v56, v57, v57, v59, v57, l4, l1, v14+int32(44))
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return int32(0)
		} else {
			v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
			switch v65 - int32(1) {
			case 0:
				v69 = int32(1)
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
				v100 = v63 + base.B2i32(v69 <= v70)
				*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v100
				v109 = base.B2i32(int32(0) < v100)
				v112 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v109)
				m.G0 = v14 + int32(48)
				return v112
			case 1:
				v69 = v7
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
				v100 = v63 + base.B2i32(v69 <= v70)
				*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v100
				v109 = base.B2i32(int32(0) < v100)
				v112 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v109)
				m.G0 = v14 + int32(48)
				return v112
			case 2:
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
				if v74 != 0 {
					v100 = int32(0)
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v63<<(uint(int32(2))%32))))
					*(*int32)(unsafe.Add(mBase, uint32(v75))) = v79
					v100 = int32(1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v100
				v109 = base.B2i32(int32(0) < v100)
				v112 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v109)
				m.G0 = v14 + int32(48)
				return v112
			case 3:
				v83 = int32(1)
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
				v87 = v63 + base.B2i32(v83 <= v85)
				v88 = v84 - v87
				v90 = v88 << (uint(int32(2)) % 32)
				if v90 == int32(0) {
					v100 = v88
				} else {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					base.MemoryCopy(m, v93, v93+v87<<(uint(int32(2))%32), v90)
					v100 = v88
				}
				*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v100
				v109 = base.B2i32(int32(0) < v100)
				v112 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v109)
				m.G0 = v14 + int32(48)
				return v112
			case 4:
				v83 = v7
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
				v87 = v63 + base.B2i32(v83 <= v85)
				v88 = v84 - v87
				v90 = v88 << (uint(int32(2)) % 32)
				if v90 == int32(0) {
					v100 = v88
				} else {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					base.MemoryCopy(m, v93, v93+v87<<(uint(int32(2))%32), v90)
					v100 = v88
				}
				*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v100
				v109 = base.B2i32(int32(0) < v100)
				v112 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v109)
				m.G0 = v14 + int32(48)
				return v112
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return int32(0)
				} else {
					v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v122
					F_errmsg_internal(m, int32(_a_F__bt_saoparray_shrink_0), v14)
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F__bt_saoparray_shrink_1), int32(1232), int32(_a_F__bt_saoparray_shrink_2))
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
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
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v16)+208))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+v18<<(uint(int32(2))%32)-int32(4))))
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v39 != 0 {
			v40 = v39
		} else {
			v40 = v24
		}
		v42 = F_get_opfamily_proc(m, v38, v27, v40, int32(1))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			if v42 == int32(0) {
				v48 = int32(0)
				v109 = v48
				v112 = v48
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v109)
				m.G0 = v14 + int32(48)
				return v112
			} else {
				v51 = v14 + int32(16)
				F_fmgr_info(m, v42, v51)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v56 = v51
					v57 = int32(0)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
					v63 = F__bt_binsrch_array_skey(m, v56, v57, v57, v59, v57, l4, l1, v14+int32(44))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
						switch v65 - int32(1) {
						case 0:
							v69 = int32(1)
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
							v100 = v63 + base.B2i32(v69 <= v70)
							*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v100
							v109 = base.B2i32(int32(0) < v100)
							v112 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v109)
							m.G0 = v14 + int32(48)
							return v112
						case 1:
							v69 = v7
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
							v100 = v63 + base.B2i32(v69 <= v70)
							*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v100
							v109 = base.B2i32(int32(0) < v100)
							v112 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v109)
							m.G0 = v14 + int32(48)
							return v112
						case 2:
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
							if v74 != 0 {
								v100 = int32(0)
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v63<<(uint(int32(2))%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v75))) = v79
								v100 = int32(1)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v100
							v109 = base.B2i32(int32(0) < v100)
							v112 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v109)
							m.G0 = v14 + int32(48)
							return v112
						case 3:
							v83 = int32(1)
							v84 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
							v87 = v63 + base.B2i32(v83 <= v85)
							v88 = v84 - v87
							v90 = v88 << (uint(int32(2)) % 32)
							if v90 == int32(0) {
								v100 = v88
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
								base.MemoryCopy(m, v93, v93+v87<<(uint(int32(2))%32), v90)
								v100 = v88
							}
							*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v100
							v109 = base.B2i32(int32(0) < v100)
							v112 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v109)
							m.G0 = v14 + int32(48)
							return v112
						case 4:
							v83 = v7
							v84 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
							v87 = v63 + base.B2i32(v83 <= v85)
							v88 = v84 - v87
							v90 = v88 << (uint(int32(2)) % 32)
							if v90 == int32(0) {
								v100 = v88
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
								base.MemoryCopy(m, v93, v93+v87<<(uint(int32(2))%32), v90)
								v100 = v88
							}
							*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v100
							v109 = base.B2i32(int32(0) < v100)
							v112 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v109)
							m.G0 = v14 + int32(48)
							return v112
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = v122
								F_errmsg_internal(m, int32(_a_F__bt_saoparray_shrink_0), v14)
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F__bt_saoparray_shrink_1), int32(1232), int32(_a_F__bt_saoparray_shrink_2))
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
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
		}
	}
}
func F_bt_entry_unique_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
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
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	v4 = l3
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v14&int32(_a_F_bt_entry_unique_check_0) == v6 {
		v82 = l1
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v12 - int32(-64)
	return
L2:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if base.B2i32(v108 == int32(-1))|base.B2i32(v108 == l2) != 0 {
		goto L1
	} else {
		goto L35
	}
L3:
	;
	F_bt_report_duplicate(m, l0, l4, v82, l2, v4, int32(-1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L13
	} else {
		goto L34
	}
L4:
	;
	F_bt_report_duplicate(m, l0, l4, v44, l2, v4, v33)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L13
	} else {
		goto L33
	}
L5:
	;
	v83 = F_heap_entry_is_visible(m, l0, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L13
	} else {
		goto L27
	}
L6:
	;
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v19&int32(_a_F_bt_entry_unique_check_0) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v19&int32(4095) == int32(0) {
		goto L2
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v71 = int32(0)
	if v19&int32(_a_F_bt_entry_unique_check_1) == v71 {
		v82 = v71
		goto L5
	} else {
		goto L26
	}
L10:
	;
	v33 = v6
	v34 = int32(0)
	goto L11
L11:
	;
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v44 = v36 + (l1 + v37<<(uint(int32(16))%32)) + v33*int32(6)
	v45 = F_heap_entry_is_visible(m, l0, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v34 != 0 {
		goto L1
	} else {
		goto L25
	}
L13:
	;
	return
L14:
	;
	if v45 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if v47 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v66 = v33 + int32(1)
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if base.Ui32(v66) < base.Ui32(v67&int32(4095)) {
		v33 = v66
		goto L11
	} else {
		goto L24
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v33
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = l2
	v58 = int32(1)
	v60 = v33 + v58
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if base.Ui32(v60) < base.Ui32(v61&int32(4095)) {
		v33 = v60
		v34 = v58
		goto L11
	} else {
		goto L23
	}
L19:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+4)))
	if v50 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v51 == l2 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+4)))
	if v53 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	goto L1
L24:
	;
	goto L12
L25:
	;
	goto L2
L26:
	;
	v82 = l1 + v14&int32(_a_F_bt_entry_unique_check_2) - int32(6)
	goto L5
L27:
	;
	if v83 == int32(0) {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if v87 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+4)))
	if v88 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v82
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = int32(-1)
	goto L1
L32:
	;
	goto L31
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v113 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v123 = int32(_a_F_bt_entry_unique_check_3)
	goto L38
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v113
	v121 = F_psprintf(m, int32(_a_F_bt_entry_unique_check_4), v10+int32(-16))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L13
	} else {
		goto L39
	}
L38:
	;
	v126 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L13
	} else {
		goto L40
	}
L39:
	;
	v123 = v121
	goto L38
L40:
	;
	if v126 == int32(0) {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v134 + int32(4)
	F_errmsg(m, int32(_a_F_bt_entry_unique_check_5), v10+int32(-32))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+2)))
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145))))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v146 | v147<<(uint(int32(16))%32)
	F_errdetail(m, int32(_a_F_bt_entry_unique_check_6), v12)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	F_errhint(m, int32(_a_F_bt_entry_unique_check_7), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_bt_entry_unique_check_8), int32(999), int32(_a_F_bt_entry_unique_check_9))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	goto L1
}
func F_bt_index_check(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v12 < int32(2) {
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(base.B2i32(v15 != int32(0)))
		if v12 == int32(2) {
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(base.B2i32(v21 != int32(0)))
		}
	}
	F_amcheck_lock_relation_and_check(m, v9, int32(403), int32(_a_F_bt_index_check_0), int32(1), v7+int32(12))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return int32(0)
	}
}
func F_bt_leftmost_ignoring_half_dead(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var __phi21 int32
	_ = __phi21
	var v23 int32
	_ = v23
	var __phi23 int32
	_ = __phi23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = int32(1)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v16 == int32(0) {
		v94 = v15
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v94
L2:
	;
	__phi21 = l1
	__phi23 = v16
	v21 = __phi21
	v23 = __phi23
	goto L3
L3:
	;
	v29 = F_palloc_btree_page(m, l0, v23)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_pfree(m, v29)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L25
	}
L5:
	;
	return int32(0)
L6:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_bt_leftmost_ignoring_half_dead[0]))
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if l1 == v23 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	goto L4
L12:
	;
	v39 = v33 + v29
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+12)))
	if base.B2i32(v40&int32(16) == int32(0))|base.B2i32(v21 == v23) != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v47 != v21 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v53 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	if v53 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	F_pfree(m, v29)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L23
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v59 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_leftmost_ignoring_half_dead_0), v13+int32(16))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v23
	F_errdetail_internal(m, int32(_a_F_bt_leftmost_ignoring_half_dead_1), v13)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_bt_leftmost_ignoring_half_dead_2), int32(1051), int32(_a_F_bt_leftmost_ignoring_half_dead_3))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	if v80 != 0 {
		__phi21 = v23
		__phi23 = v80
		v21 = __phi21
		v23 = __phi23
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v94 = v15
	goto L1
L25:
	;
	v94 = int32(0)
	goto L1
}
func F_bt_page_stats(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_bt_page_stats_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_bt_page_stats_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	v7 = m.G0
	v9 = v7 - int32(272)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if l1 != 0 {
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
			v19 = v17
		} else {
			v19 = base.I64_extend_i32_u(v16)
		}
		v20 = F_superuser(m)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v20 != 0 {
				v22 = F_textToQualifiedNameList(m, v12)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = F_makeRangeVarFromNameList(m, v22)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v27 = F_relation_openrv(m, v24, int32(1))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							F_bt_index_block_validate(m, v27, v19)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								v31 = base.I32_wrap_i64(v19)
								v32 = F_ReadBuffer(m, v27, v31)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									F_LockBuffer(m, v32, int32(1))
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int32(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v9)+208)) = int64(-1)
										v39 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v9)+220)) = uint16(v39)
										*(*int64)(unsafe.Add(mBase, uint32(v9)+196)) = int64(0)
										F_GetBTPageStatistics(m, v31, v32, v9+int32(176))
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return int32(0)
										} else {
											F_UnlockReleaseBuffer(m, v32)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v27, int32(1))
												mBase = m.M
												v51 = m.ExcPending
												if v51 != 0 {
													return int32(0)
												} else {
													v55 = F_get_call_result_type(m, l0, int32(0), v9+int32(268))
													mBase = m.M
													v56 = m.ExcPending
													if v56 != 0 {
														return int32(0)
													} else {
														if v55 != int32(1) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(_a_F_bt_page_stats_internal_0), int32(0))
																mBase = m.M
																v182 = m.ExcPending
																if v182 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_bt_page_stats_internal_1), int32(297), int32(_a_F_bt_page_stats_internal_2))
																	mBase = m.M
																	v187 = m.ExcPending
																	if v187 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+176))
															*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v59
															v64 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(160))
															mBase = m.M
															v65 = m.ExcPending
															if v65 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9)+224)) = v64
																v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9)+204)))
																*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v67
																v72 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_4), v9+int32(144))
																mBase = m.M
																v73 = m.ExcPending
																if v73 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+228)) = v72
																	v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)+180))
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v75
																	v80 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(128))
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+232)) = v80
																		v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+184))
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v83
																		v88 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(112))
																		mBase = m.M
																		v89 = m.ExcPending
																		if v89 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+236)) = v88
																			v91 = *(*int32)(unsafe.Add(mBase, uint32(v9)+200))
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v91
																			v96 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(96))
																			mBase = m.M
																			v97 = m.ExcPending
																			if v97 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v9)+240)) = v96
																				v99 = *(*int32)(unsafe.Add(mBase, uint32(v9)+188))
																				*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v99
																				v104 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(80))
																				mBase = m.M
																				v105 = m.ExcPending
																				if v105 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v9)+244)) = v104
																					v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+196))
																					*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v107
																					v112 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9-int32(-64))
																					mBase = m.M
																					v113 = m.ExcPending
																					if v113 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v9)+248)) = v112
																						v115 = *(*int32)(unsafe.Add(mBase, uint32(v9)+208))
																						*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v115
																						v120 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(48))
																						mBase = m.M
																						v121 = m.ExcPending
																						if v121 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v9)+252)) = v120
																							v123 = *(*int32)(unsafe.Add(mBase, uint32(v9)+212))
																							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v123
																							v128 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(32))
																							mBase = m.M
																							v129 = m.ExcPending
																							if v129 != 0 {
																								return int32(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v9)+256)) = v128
																								v131 = *(*int32)(unsafe.Add(mBase, uint32(v9)+216))
																								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v131
																								v136 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(16))
																								mBase = m.M
																								v137 = m.ExcPending
																								if v137 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v9)+260)) = v136
																									v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+220)))
																									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v139
																									v142 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_5), v9)
																									mBase = m.M
																									v143 = m.ExcPending
																									if v143 != 0 {
																										return int32(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v9)+264)) = v142
																										v145 = *(*int32)(unsafe.Add(mBase, uint32(v9)+268))
																										v146 = F_TupleDescGetAttInMetadata(m, v145)
																										mBase = m.M
																										v147 = m.ExcPending
																										if v147 != 0 {
																											return int32(0)
																										} else {
																											v150 = F_BuildTupleFromCStrings(m, v146, v9+int32(224))
																											mBase = m.M
																											v151 = m.ExcPending
																											if v151 != 0 {
																												return int32(0)
																											} else {
																												v152 = *(*int32)(unsafe.Add(mBase, uint32(v150)+16))
																												v153 = F_HeapTupleHeaderGetDatum(m, v152)
																												mBase = m.M
																												v154 = m.ExcPending
																												if v154 != 0 {
																													return int32(0)
																												} else {
																													m.G0 = v9 + int32(272)
																													return v153
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
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v162 = m.ExcPending
				if v162 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v165 = m.ExcPending
					if v165 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_bt_page_stats_internal_6), int32(0))
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_bt_page_stats_internal_1), int32(276), int32(_a_F_bt_page_stats_internal_2))
							mBase = m.M
							v174 = m.ExcPending
							if v174 != 0 {
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
}
