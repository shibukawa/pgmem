package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F__bt_allequalimage(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+10)))
	if v14 != v15 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v115
L2:
	;
	v115 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	if int32(0) < base.I32_extend16_s(v14) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v103 + int32(4)
	F_errmsg_internal(m, v95, v11)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L14
	} else {
		goto L30
	}
L6:
	;
	v85 = int32(1)
	v88 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L14
	} else {
		goto L28
	}
L7:
	;
	v68 = int32(0)
	v71 = F_errstart(m, int32(14), v68)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L14
	} else {
		goto L26
	}
L8:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L25
	}
L9:
	;
	v25 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	if l1 != 0 {
		goto L6
	} else {
		goto L24
	}
L12:
	;
	v31 = v25 << (uint(int32(2)) % 32)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31+v32)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35+v31)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38+v31)))
	v42 = F_get_opfamily_proc(m, v37, v40, v40, int32(4))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if l1 == int32(0) {
		v115 = v52
		goto L1
	} else {
		goto L22
	}
L14:
	;
	return int32(0)
L15:
	;
	if v42 == int32(0) {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v49 = F_OidFunctionCall1Coll(m, v42, v34, base.I64_extend_i32_u(v40))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v52 = base.B2i32(v49 != int64(0))
	if v49 != int64(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v54 = v25 + int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+10)))
	if v54 < v56 {
		v25 = v54
		goto L12
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L13
L21:
	;
	goto L20
L22:
	;
	if v49 == int64(0) {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	goto L6
L24:
	;
	v115 = int32(1)
	goto L1
L25:
	;
	v115 = int32(0)
	goto L1
L26:
	;
	if v71 == int32(0) {
		v115 = v68
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v95 = int32(_a_F__bt_allequalimage_0)
	v96 = v68
	v102 = int32(1214)
	goto L5
L28:
	;
	if v88 == int32(0) {
		v115 = v85
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v95 = int32(_a_F__bt_allequalimage_1)
	v96 = v85
	v102 = int32(1211)
	goto L5
L30:
	;
	F_errfinish(m, int32(_a_F__bt_allequalimage_2), v102, int32(_a_F__bt_allequalimage_3))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	v115 = v96
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
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v176 int64
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v277 int64
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
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
	return v359
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
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v10)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v277
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v279
	v283 = int32(0)
	v286 = F_ExtendBufferedRel(m, v8+int32(-56), v283, v283, int32(8))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L2
	} else {
		goto L90
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
	v263 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L2
	} else {
		goto L87
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
	v248 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
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
	v359 = v25
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
	v233 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
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
	v181 = int32(_a_F__bt_allocbuf_0)
	v183 = int32(0)
	if v183|(v46&int32(3)|int32(1)) == v183 {
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
	if base.Ui32(int32(1)) < base.Ui32(v121) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+56)) = uint8(v164)
	F_XLogBeginInsert(m)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L59
	}
L47:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+118)))
	if v146 != int32(112) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__bt_allocbuf[3])))
	if v141&int32(1) != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v164 = int32(0)
	goto L46
L50:
	;
	v164 = int32(0)
	goto L46
L51:
	;
	goto L52
L52:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	goto L53
L53:
	;
	if base.Ui32(v151) < base.Ui32(int32(_a_F__bt_allocbuf_4)) {
		v164 = int32(1)
		goto L46
	} else {
		goto L54
	}
L54:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	if v154 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v164 = int32(0)
	goto L46
L56:
	;
	goto L57
L57:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+119)))
	switch v160 - int32(109) {
	case 0, 5:
		goto L58
	default:
		v164 = int32(0)
		goto L46
	}
L58:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+112)))
	v164 = v163
	goto L46
L59:
	;
	F_XLogRegisterData(m, v8+int32(-32), int32(25))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v176 = F_XLogInsert(m, int32(11), int32(208))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	goto L40
L62:
	;
	v359 = v25
	goto L1
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+10)) = int32(_a_F__bt_allocbuf_1)
	v222 = int32(_a_F__bt_allocbuf_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v46)+18)) = uint16(v222)
	v228 = int32(_a_F__bt_allocbuf_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v46)+16)) = uint16(v228)
	*(*uint16)(unsafe.Add(mBase, uint32(v46)+14)) = uint16(v228)
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
	v199 = v46 + v181
	v201 = v46 + int32(4)
	if base.Ui32(v201) < base.Ui32(v199) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v203 = v199
	goto L70
L69:
	;
	v203 = v201
	goto L70
L70:
	;
	v208 = (v46^int32(-1)+v203)&int32(-4) + int32(4)
	if v208 == int32(0) {
		goto L63
	} else {
		goto L71
	}
L71:
	;
	base.MemoryFill(m, v46, int32(0), v208)
	goto L63
L72:
	;
	base.MemoryFill(m, v46, int32(0), v181)
	goto L63
L73:
	;
	if v233 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_errmsg_internal(m, int32(_a_F__bt_allocbuf_5), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L2
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	F_UnlockReleaseBuffer(m, v25)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L2
	} else {
		goto L79
	}
L77:
	;
	F_errfinish(m, int32(_a_F__bt_allocbuf_6), int32(965), int32(_a_F__bt_allocbuf_7))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
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
	if v248 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	F_errmsg_internal(m, int32(_a_F__bt_allocbuf_8), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L2
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	F_ReleaseBuffer(m, v25)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L2
	} else {
		goto L86
	}
L84:
	;
	F_errfinish(m, int32(_a_F__bt_allocbuf_6), int32(970), int32(_a_F__bt_allocbuf_7))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	goto L9
L87:
	;
	if v263 != int32(-1) {
		v23 = v263
		goto L7
	} else {
		goto L88
	}
L88:
	;
	goto L8
L89:
	;
	v306 = int32(_a_F__bt_allocbuf_0)
	v308 = int32(0)
	if v308|(v305&int32(3)|int32(1)) == v308 {
		goto L96
	} else {
		goto L97
	}
L90:
	;
	if v286 < int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _c_F__bt_allocbuf[0]))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v291+(v286^int32(-1))<<(uint(int32(2))%32))))
	v305 = v297
	goto L89
L92:
	;
	goto L93
L93:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _c_F__bt_allocbuf[1]))
	v305 = v299 + v286<<(uint(int32(13))%32) + int32(-8192)
	goto L89
L94:
	;
	v359 = v286
	goto L1
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305)+10)) = int32(_a_F__bt_allocbuf_1)
	v347 = int32(_a_F__bt_allocbuf_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v305)+18)) = uint16(v347)
	v353 = int32(_a_F__bt_allocbuf_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v305)+16)) = uint16(v353)
	*(*uint16)(unsafe.Add(mBase, uint32(v305)+14)) = uint16(v353)
	goto L94
L96:
	;
	goto L99
L97:
	;
	goto L98
L98:
	;
	goto L104
L99:
	;
	v324 = v305 + v306
	v326 = v305 + int32(4)
	if base.Ui32(v326) < base.Ui32(v324) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v328 = v324
	goto L102
L101:
	;
	v328 = v326
	goto L102
L102:
	;
	v333 = (v305^int32(-1)+v328)&int32(-4) + int32(4)
	if v333 == int32(0) {
		goto L95
	} else {
		goto L103
	}
L103:
	;
	base.MemoryFill(m, v305, int32(0), v333)
	goto L95
L104:
	;
	base.MemoryFill(m, v305, int32(0), v306)
	goto L95
}
func F__bt_check_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v128 int64
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v228 int64
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v272 int32
	_ = v272
	var v282 int32
	_ = v282
	var v283 int64
	_ = v283
	var v284 int64
	_ = v284
	var v285 int32
	_ = v285
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v362 int32
	_ = v362
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v404 int32
	_ = v404
	var v434 int32
	_ = v434
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v24 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v24)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v28 <= v27 {
		v434 = v24
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v21 + int32(16)
	return v434
L2:
	;
	v40 = v27
	goto L5
L3:
	;
	v434 = int32(0)
	goto L1
L4:
	;
	v404 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v404)
	goto L3
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v51 = v48 + v40*int32(56)
	v52 = int32(0)
	if l6 != 0 {
		v83 = v52
		v86 = v52
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
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+4)))
	if l3 < v87 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v54 = int32(1)
	v55 = int32(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v58 = v56 & int32(_a_F__bt_check_compare_0)
	if base.B2i32(v58 == v55)|base.B2i32(l1 != v54) == v55 {
		v83 = v54
		v86 = v55
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v67 = v56 & int32(_a_F__bt_check_compare_1)
	if l1 == int32(-1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if v67 != 0 {
		v83 = v54
		v86 = int32(0)
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v71 = int32(0)
	v83 = v71
	v86 = base.B2i32(l1 == int32(-1))&base.B2i32(v58 != v71) | base.B2i32(l1 == int32(1))&base.B2i32(v67 != v71)
	goto L7
L13:
	;
	goto L12
L14:
	;
	goto L6
L15:
	;
	v345 = int32(1)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v348 = v346 + v345
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v348
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v348 < v350 {
		v40 = v348
		goto L5
	} else {
		goto L105
	}
L16:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v89&int32(_a_F__bt_check_compare_2) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v311 = int32(0)
	if l5 == v311 {
		v434 = v311
		goto L1
	} else {
		goto L101
	}
L18:
	;
	v308 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v308)
	v434 = v308
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
	if v89&int32(4) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v94 = int32(0)
	v96 = F__bt_advance_array_keys(m, l0, v94, l2, l3, l4, v40, v94)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	v434 = v96
	goto L1
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v51)+48))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v103&int32(1) != 0 {
		v362 = v102
		goto L14
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v228 = F_index_getattr_2(m, l2, v87, l4, v21+int32(14))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L23
	} else {
		goto L79
	}
L28:
	;
	v116 = v102
	goto L29
L29:
	;
	v124 = int32(*(*int16)(unsafe.Add(mBase, uint32(v116)+4)))
	if l3 < v124 {
		goto L15
	} else {
		goto L31
	}
L30:
	;
	v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)))
	switch v181 - int32(1) {
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
	v128 = F_index_getattr_2(m, l2, v124, l4, v21+int32(15))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)))
	if v130 == int32(1) {
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
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v116)+48))
	v160 = F_FunctionCall2Coll(m, v116+int32(16), v158, v128, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L23
	} else {
		goto L49
	}
L36:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v51)+48))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v134&int32(33554432) != 0 {
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
	if v133 == v116 {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	if v133 == v116 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v142 = int32(_a_F__bt_check_compare_3)
	goto L43
L42:
	;
	v142 = int32(_a_F__bt_check_compare_1)
	goto L43
L43:
	;
	if v142&v134 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	goto L3
L45:
	;
	v147 = int32(_a_F__bt_check_compare_3)
	goto L47
L46:
	;
	v147 = int32(_a_F__bt_check_compare_0)
	goto L47
L47:
	;
	if base.B2i32(v147&v134 == int32(0))|base.B2i32(l1 != int32(1)) != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	goto L4
L49:
	;
	v162 = base.I32_wrap_i64(v160)
	if v162 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v166 = int32(1)
	goto L52
L51:
	;
	v166 = int32(0) - v162
	goto L52
L52:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v167&int32(16777216) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v170 = v166
	goto L55
L54:
	;
	v170 = v162
	goto L55
L55:
	;
	if v170|v167&int32(16) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+56)))
	v178 = v116 + int32(56)
	if v176&int32(1) != 0 {
		v362 = v178
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
	v116 = v178
	goto L29
L60:
	;
	if l6|v206&int32(1) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L61:
	;
	v206 = int32(base.Ui32(v170) >> (uint(int32(31)) % 32))
	goto L60
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L23
	} else {
		goto L66
	}
L63:
	;
	v206 = base.B2i32(int32(0) < v170)
	goto L60
L64:
	;
	v206 = base.B2i32(int32(0) <= v170)
	goto L60
L65:
	;
	v206 = base.B2i32(v170 <= int32(0))
	goto L60
L66:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v194
	F_errmsg_internal(m, int32(_a_F__bt_check_compare_4), v21)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F__bt_check_compare_5), int32(1803), int32(_a_F__bt_check_compare_6))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
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
	if v167&int32(_a_F__bt_check_compare_0) != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	if v206&int32(1) != 0 {
		goto L15
	} else {
		goto L78
	}
L72:
	;
	v217 = base.B2i32(l1 == int32(1))
	goto L74
L73:
	;
	v217 = int32(0)
	goto L74
L74:
	;
	if v217 != 0 {
		goto L18
	} else {
		goto L75
	}
L75:
	;
	v218 = int32(0)
	if l1 != int32(-1) {
		v434 = v218
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v167&int32(_a_F__bt_check_compare_1) != 0 {
		goto L18
	} else {
		goto L77
	}
L77:
	;
	v434 = v218
	goto L1
L78:
	;
	v434 = int32(0)
	goto L1
L79:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v230&int32(1) != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+14)))
	if (v233^base.B2i32(v230&int32(64) == int32(0)))&int32(1) != 0 {
		goto L15
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+14)))
	if v244 == int32(1) {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	if v83 != 0 {
		goto L18
	} else {
		goto L84
	}
L84:
	;
	if v230&int32(_a_F__bt_check_compare_7) != 0 {
		goto L15
	} else {
		goto L85
	}
L85:
	;
	v434 = int32(0)
	goto L1
L86:
	;
	v247 = int32(0)
	if base.B2i32(l6 == v247)|base.B2i32(v230&int32(_a_F__bt_check_compare_7) == v247) == v247 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v51)+48))
	v284 = F_FunctionCall2Coll(m, v51+int32(16), v282, v228, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L23
	} else {
		goto L98
	}
L89:
	;
	v256 = int32(0)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v259 = F__bt_advance_array_keys(m, l0, v256, l2, l3, l4, v257, v256)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L23
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v261 = v83 | v86
	if v230&int32(33554432) != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v434 = v259
	goto L1
L93:
	;
	if v261^int32(1)|base.B2i32(l1 != int32(-1)) == int32(0) {
		goto L18
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v272 = int32(1)
	if v261^v272|base.B2i32(l1 != v272) == int32(0) {
		goto L18
	} else {
		goto L97
	}
L96:
	;
	v434 = int32(0)
	goto L1
L97:
	;
	v434 = int32(0)
	goto L1
L98:
	;
	if v284 != int64(0) {
		goto L15
	} else {
		goto L99
	}
L99:
	;
	if v83 == int32(0) {
		goto L17
	} else {
		goto L100
	}
L100:
	;
	goto L18
L101:
	;
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+6)))
	if v314 != int32(3) {
		v434 = v311
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v317&int32(32) == int32(0) {
		v434 = v311
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v322 = int32(0)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v325 = F__bt_advance_array_keys(m, l0, v322, l2, l3, l4, v323, v322)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L23
	} else {
		goto L104
	}
L104:
	;
	v434 = v325
	goto L1
L105:
	;
	v434 = v345
	goto L1
L106:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v362-int32(56))))
	if v375&int32(_a_F__bt_check_compare_0) != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v378 = base.B2i32(l1 == int32(1))
	goto L109
L108:
	;
	v378 = int32(0)
	goto L109
L109:
	;
	if v378 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	if base.B2i32(v375&int32(_a_F__bt_check_compare_1) == int32(0))|base.B2i32(l1 != int32(-1)) != 0 {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	goto L4
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
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
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
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
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
	return v409
L2:
	;
	if v224 == int32(0) {
		goto L56
	} else {
		goto L57
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
	v224 = v18
	goto L2
L8:
	;
	v224 = v140
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
	F_LockBufferInternal(m, v32, int32(1))
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
	v196 = m.ExcPending
	if v196 != 0 {
		goto L6
	} else {
		goto L53
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L6
	} else {
		goto L49
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L45
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
	F_UnlockReleaseBuffer(m, v32)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v56)+36))
	v82 = v32
	v85 = v72
	goto L32
L30:
	;
	v140 = int32(0)
	goto L26
L31:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	if v135 != v78 {
		goto L16
	} else {
		goto L44
	}
L32:
	;
	v89 = F__bt_relandgetbuf(m, l0, v82, v85, int32(1))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L35
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L41
	}
L34:
	;
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+16)))
	v110 = v109 + v108
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+12)))
	if v111&int32(20) == int32(0) {
		goto L31
	} else {
		goto L39
	}
L35:
	;
	if v89 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[0]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v94+(v89^int32(-1))<<(uint(int32(2))%32))))
	v108 = v100
	goto L34
L37:
	;
	goto L38
L38:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[1]))
	v108 = v102 + v89<<(uint(int32(13))%32) + int32(-8192)
	goto L34
L39:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v116 != 0 {
		v82 = v89
		v85 = v116
		goto L32
	} else {
		goto L40
	}
L40:
	;
	goto L33
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v121 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_get_endpoint_1), v22+int32(-16))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F__bt_get_endpoint_2), int32(656), int32(_a_F__bt_get_endpoint_3))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v140 = v89
	goto L26
L45:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v156 + int32(4)
	F_errmsg(m, int32(_a_F__bt_get_endpoint_4), v24)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F__bt_get_endpoint_2), int32(617), int32(_a_F__bt_get_endpoint_3))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = int64(8589934596)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v175 + int32(4)
	F_errmsg(m, int32(_a_F__bt_get_endpoint_5), v22+int32(-48))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F__bt_get_endpoint_2), int32(626), int32(_a_F__bt_get_endpoint_3))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v197 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_get_endpoint_6), v22+int32(-32))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F__bt_get_endpoint_2), int32(663), int32(_a_F__bt_get_endpoint_3))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
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
	v409 = int32(0)
	goto L1
L57:
	;
	goto L58
L58:
	;
	if v224 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v249 = v224
	v250 = v245
	goto L65
L60:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[0]))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v231+(v224^int32(-1))<<(uint(int32(2))%32))))
	v245 = v237
	goto L59
L61:
	;
	goto L62
L62:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[1]))
	v245 = v239 + v224<<(uint(int32(13))%32) + int32(-8192)
	goto L59
L63:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L6
	} else {
		goto L104
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L6
	} else {
		goto L100
	}
L65:
	;
	v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v250)+16)))
	v256 = v250 + v255
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+12)))
	if v257&int32(20) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L6
	} else {
		goto L97
	}
L67:
	;
	goto L66
L68:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v250+v316&int32(_a_F__bt_get_endpoint_7)<<(uint(int32(2))%32))+20))
	v326 = v250 + v323&int32(_a_F__bt_get_endpoint_8)
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v326))))
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v326)+2)))
	v333 = F__bt_relandgetbuf(m, l0, v249, v327<<(uint(int32(16))%32)|v330, int32(1))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L6
	} else {
		goto L93
	}
L69:
	;
	v302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v250)+12)))
	if base.Ui32(v302) < base.Ui32(int32(25)) {
		goto L63
	} else {
		goto L88
	}
L70:
	;
	v283 = F__bt_relandgetbuf(m, l0, v249, v281, int32(1))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L6
	} else {
		goto L84
	}
L71:
	;
	if l2 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v278 == int32(0) {
		goto L67
	} else {
		goto L83
	}
L74:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v262 != 0 {
		v281 = v262
		goto L70
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	if v264 == l1 {
		v409 = v249
		goto L1
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	if base.Ui32(v264) < base.Ui32(l1) {
		goto L64
	} else {
		goto L79
	}
L79:
	;
	if l2 == int32(0) {
		goto L69
	} else {
		goto L80
	}
L80:
	;
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v250)+12)))
	if base.Ui32(v269) < base.Ui32(int32(25)) {
		goto L63
	} else {
		goto L81
	}
L81:
	;
	v275 = int32(base.Ui32(v269+int32(_a_F__bt_get_endpoint_9)) >> (uint(int32(2)) % 32))
	if v275&int32(_a_F__bt_get_endpoint_7) != 0 {
		v316 = v275
		goto L68
	} else {
		goto L82
	}
L82:
	;
	goto L63
L83:
	;
	v281 = v278
	goto L70
L84:
	;
	if v283 < int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[0]))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v288+(v283^int32(-1))<<(uint(int32(2))%32))))
	v249 = v283
	v250 = v294
	goto L65
L86:
	;
	goto L87
L87:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[1]))
	v249 = v283
	v250 = v296 + v283<<(uint(int32(13))%32) + int32(-8192)
	goto L65
L88:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v307 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v308 = int32(2)
	goto L91
L90:
	;
	v308 = int32(1)
	goto L91
L91:
	;
	if base.Ui32(int32(base.Ui32(v302+int32(_a_F__bt_get_endpoint_9))>>(uint(int32(2))%32))&int32(_a_F__bt_get_endpoint_7)) < base.Ui32(v308) {
		goto L63
	} else {
		goto L92
	}
L92:
	;
	v316 = v308
	goto L68
L93:
	;
	if v333 < int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[0]))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v338+(v333^int32(-1))<<(uint(int32(2))%32))))
	v352 = v344
	goto L96
L95:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _c_F__bt_get_endpoint[1]))
	v352 = v346 + v333<<(uint(int32(13))%32) + int32(-8192)
	goto L96
L96:
	;
	v249 = v333
	v250 = v352
	goto L65
L97:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v357 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_get_endpoint_10), v12+int32(16))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F__bt_get_endpoint_11), int32(2131), int32(_a_F__bt_get_endpoint_12))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
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
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F__bt_get_endpoint_11), int32(2144), int32(_a_F__bt_get_endpoint_12))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	F_errmsg_internal(m, int32(_a_F__bt_get_endpoint_14), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F__bt_get_endpoint_11), int32(2153), int32(_a_F__bt_get_endpoint_12))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
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
	var v34 int32
	_ = v34
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
	var v53 int32
	_ = v53
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
	var v111 int32
	_ = v111
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
	var v284 int32
	_ = v284
	var v286 int64
	_ = v286
	var v288 int64
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	if l5 != 0 {
		if l2 < int32(0) {
			v25 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v25+(l2^int32(-1))*int32(56))+16))
			v40 = v31
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
			v34 = int32(56)
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v33+l2*v34-v34)+16))
			v40 = v39
		}
		if l3 < int32(0) {
			v44 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+(l3^int32(-1))*int32(56))+16))
			v59 = v50
		} else {
			v52 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
			v53 = int32(56)
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v52+l3*v53-v53)+16))
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
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v102+(v79^int32(-1))*int32(56))+16))
				v117 = v108
			} else {
				v110 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
				v111 = int32(56)
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v110+v79*v111-v111)+16))
				v117 = v116
			}
			v120 = F__bt_getbuf(m, l0, int32(0), int32(3))
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
								F_errstart_cold(m, int32(24), int32(0))
								mBase = m.M
								v473 = m.ExcPending
								if v473 != 0 {
									return
								} else {
									if l2 < int32(0) {
										v477 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
										v483 = *(*int32)(unsafe.Add(mBase, uint32(v477+(l2^int32(-1))*int32(56))+16))
										v492 = v483
									} else {
										v485 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
										v486 = int32(56)
										v491 = *(*int32)(unsafe.Add(mBase, uint32(v485+l2*v486-v486)+16))
										v492 = v491
									}
									v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v20))) = v492
									*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v493 + int32(4)
									F_errmsg_internal(m, int32(_a_F__bt_insert_parent_2), v20)
									mBase = m.M
									v500 = m.ExcPending
									if v500 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F__bt_insert_parent_3), int32(2582), int32(_a_F__bt_insert_parent_4))
										mBase = m.M
										v505 = m.ExcPending
										if v505 != 0 {
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
										F_errstart_cold(m, int32(24), int32(0))
										mBase = m.M
										v509 = m.ExcPending
										if v509 != 0 {
											return
										} else {
											if l2 < int32(0) {
												v513 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
												v519 = *(*int32)(unsafe.Add(mBase, uint32(v513+(l2^int32(-1))*int32(56))+16))
												v528 = v519
											} else {
												v521 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
												v522 = int32(56)
												v527 = *(*int32)(unsafe.Add(mBase, uint32(v521+l2*v522-v522)+16))
												v528 = v527
											}
											v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v528
											*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v529 + int32(4)
											F_errmsg_internal(m, int32(_a_F__bt_insert_parent_5), v20+int32(16))
											mBase = m.M
											v538 = m.ExcPending
											if v538 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F__bt_insert_parent_3), int32(2593), int32(_a_F__bt_insert_parent_4))
												mBase = m.M
												v543 = m.ExcPending
												if v543 != 0 {
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
														v283 = F_XLogGetFakeLSN(m, l0)
														mBase = m.M
														v284 = m.ExcPending
														if v284 != 0 {
															return
														} else {
															v286 = v283
															v288 = base.I64_rotl(v286, int64(32))
															*(*int64)(unsafe.Add(mBase, uint32(v77))) = v288
															*(*int64)(unsafe.Add(mBase, uint32(v98))) = v288
															*(*int64)(unsafe.Add(mBase, uint32(v139))) = v288
															v292 = int32(_a_F__bt_insert_parent_1)
															v294 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4]))
															*(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4])) = v294 - int32(1)
															F_UnlockReleaseBuffer(m, v120)
															mBase = m.M
															v299 = m.ExcPending
															if v299 != 0 {
																return
															} else {
																F_pfree(m, v141)
																mBase = m.M
																v301 = m.ExcPending
																if v301 != 0 {
																	return
																} else {
																	F_pfree(m, v153)
																	mBase = m.M
																	v303 = m.ExcPending
																	if v303 != 0 {
																		return
																	} else {
																		F_UnlockReleaseBuffer(m, v79)
																		mBase = m.M
																		v305 = m.ExcPending
																		if v305 != 0 {
																			return
																		} else {
																			F_UnlockReleaseBuffer(m, l3)
																			mBase = m.M
																			v307 = m.ExcPending
																			if v307 != 0 {
																				return
																			} else {
																				F_UnlockReleaseBuffer(m, l2)
																				mBase = m.M
																				v309 = m.ExcPending
																				if v309 != 0 {
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
													} else {
														v228 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[5]))
														if v228 <= int32(0) {
															v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															if v231 != 0 {
																v283 = F_XLogGetFakeLSN(m, l0)
																mBase = m.M
																v284 = m.ExcPending
																if v284 != 0 {
																	return
																} else {
																	v286 = v283
																	v288 = base.I64_rotl(v286, int64(32))
																	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v288
																	*(*int64)(unsafe.Add(mBase, uint32(v98))) = v288
																	*(*int64)(unsafe.Add(mBase, uint32(v139))) = v288
																	v292 = int32(_a_F__bt_insert_parent_1)
																	v294 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4]))
																	*(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4])) = v294 - int32(1)
																	F_UnlockReleaseBuffer(m, v120)
																	mBase = m.M
																	v299 = m.ExcPending
																	if v299 != 0 {
																		return
																	} else {
																		F_pfree(m, v141)
																		mBase = m.M
																		v301 = m.ExcPending
																		if v301 != 0 {
																			return
																		} else {
																			F_pfree(m, v153)
																			mBase = m.M
																			v303 = m.ExcPending
																			if v303 != 0 {
																				return
																			} else {
																				F_UnlockReleaseBuffer(m, v79)
																				mBase = m.M
																				v305 = m.ExcPending
																				if v305 != 0 {
																					return
																				} else {
																					F_UnlockReleaseBuffer(m, l3)
																					mBase = m.M
																					v307 = m.ExcPending
																					if v307 != 0 {
																						return
																					} else {
																						F_UnlockReleaseBuffer(m, l2)
																						mBase = m.M
																						v309 = m.ExcPending
																						if v309 != 0 {
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
															} else {
																v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																if v232 != 0 {
																	v283 = F_XLogGetFakeLSN(m, l0)
																	mBase = m.M
																	v284 = m.ExcPending
																	if v284 != 0 {
																		return
																	} else {
																		v286 = v283
																		v288 = base.I64_rotl(v286, int64(32))
																		*(*int64)(unsafe.Add(mBase, uint32(v77))) = v288
																		*(*int64)(unsafe.Add(mBase, uint32(v98))) = v288
																		*(*int64)(unsafe.Add(mBase, uint32(v139))) = v288
																		v292 = int32(_a_F__bt_insert_parent_1)
																		v294 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4]))
																		*(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4])) = v294 - int32(1)
																		F_UnlockReleaseBuffer(m, v120)
																		mBase = m.M
																		v299 = m.ExcPending
																		if v299 != 0 {
																			return
																		} else {
																			F_pfree(m, v141)
																			mBase = m.M
																			v301 = m.ExcPending
																			if v301 != 0 {
																				return
																			} else {
																				F_pfree(m, v153)
																				mBase = m.M
																				v303 = m.ExcPending
																				if v303 != 0 {
																					return
																				} else {
																					F_UnlockReleaseBuffer(m, v79)
																					mBase = m.M
																					v305 = m.ExcPending
																					if v305 != 0 {
																						return
																					} else {
																						F_UnlockReleaseBuffer(m, l3)
																						mBase = m.M
																						v307 = m.ExcPending
																						if v307 != 0 {
																							return
																						} else {
																							F_UnlockReleaseBuffer(m, l2)
																							mBase = m.M
																							v309 = m.ExcPending
																							if v309 != 0 {
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
																									v286 = v281
																									v288 = base.I64_rotl(v286, int64(32))
																									*(*int64)(unsafe.Add(mBase, uint32(v77))) = v288
																									*(*int64)(unsafe.Add(mBase, uint32(v98))) = v288
																									*(*int64)(unsafe.Add(mBase, uint32(v139))) = v288
																									v292 = int32(_a_F__bt_insert_parent_1)
																									v294 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4]))
																									*(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4])) = v294 - int32(1)
																									F_UnlockReleaseBuffer(m, v120)
																									mBase = m.M
																									v299 = m.ExcPending
																									if v299 != 0 {
																										return
																									} else {
																										F_pfree(m, v141)
																										mBase = m.M
																										v301 = m.ExcPending
																										if v301 != 0 {
																											return
																										} else {
																											F_pfree(m, v153)
																											mBase = m.M
																											v303 = m.ExcPending
																											if v303 != 0 {
																												return
																											} else {
																												F_UnlockReleaseBuffer(m, v79)
																												mBase = m.M
																												v305 = m.ExcPending
																												if v305 != 0 {
																													return
																												} else {
																													F_UnlockReleaseBuffer(m, l3)
																													mBase = m.M
																													v307 = m.ExcPending
																													if v307 != 0 {
																														return
																													} else {
																														F_UnlockReleaseBuffer(m, l2)
																														mBase = m.M
																														v309 = m.ExcPending
																														if v309 != 0 {
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
																							v286 = v281
																							v288 = base.I64_rotl(v286, int64(32))
																							*(*int64)(unsafe.Add(mBase, uint32(v77))) = v288
																							*(*int64)(unsafe.Add(mBase, uint32(v98))) = v288
																							*(*int64)(unsafe.Add(mBase, uint32(v139))) = v288
																							v292 = int32(_a_F__bt_insert_parent_1)
																							v294 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4]))
																							*(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[4])) = v294 - int32(1)
																							F_UnlockReleaseBuffer(m, v120)
																							mBase = m.M
																							v299 = m.ExcPending
																							if v299 != 0 {
																								return
																							} else {
																								F_pfree(m, v141)
																								mBase = m.M
																								v301 = m.ExcPending
																								if v301 != 0 {
																									return
																								} else {
																									F_pfree(m, v153)
																									mBase = m.M
																									v303 = m.ExcPending
																									if v303 != 0 {
																										return
																									} else {
																										F_UnlockReleaseBuffer(m, v79)
																										mBase = m.M
																										v305 = m.ExcPending
																										if v305 != 0 {
																											return
																										} else {
																											F_UnlockReleaseBuffer(m, l3)
																											mBase = m.M
																											v307 = m.ExcPending
																											if v307 != 0 {
																												return
																											} else {
																												F_UnlockReleaseBuffer(m, l2)
																												mBase = m.M
																												v309 = m.ExcPending
																												if v309 != 0 {
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
			v313 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
			v319 = *(*int32)(unsafe.Add(mBase, uint32(v313+(l2^int32(-1))*int32(56))+16))
			v328 = v319
		} else {
			v321 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
			v322 = int32(56)
			v327 = *(*int32)(unsafe.Add(mBase, uint32(v321+l2*v322-v322)+16))
			v328 = v327
		}
		if l3 < int32(0) {
			v332 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
			v338 = *(*int32)(unsafe.Add(mBase, uint32(v332+(l3^int32(-1))*int32(56))+16))
			v347 = v338
		} else {
			v340 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
			v341 = int32(56)
			v346 = *(*int32)(unsafe.Add(mBase, uint32(v340+l3*v341-v341)+16))
			v347 = v346
		}
		if l2 < int32(0) {
			v351 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[2]))
			v357 = *(*int32)(unsafe.Add(mBase, uint32(v351+(l2^int32(-1))<<(uint(int32(2))%32))))
			v365 = v357
		} else {
			v359 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[3]))
			v365 = v359 + l2<<(uint(int32(13))%32) + int32(-8192)
		}
		if l4 == int32(0) {
			v370 = F_errstart(m, int32(13), int32(0))
			mBase = m.M
			v371 = m.ExcPending
			if v371 != 0 {
				return
			} else {
				if v370 != 0 {
					F_errmsg_internal(m, int32(_a_F__bt_insert_parent_7), int32(0))
					mBase = m.M
					v375 = m.ExcPending
					if v375 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F__bt_insert_parent_3), int32(2180), int32(_a_F__bt_insert_parent_8))
						mBase = m.M
						v380 = m.ExcPending
						if v380 != 0 {
							return
						} else {
							v381 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v365)+16)))
							v383 = *(*int32)(unsafe.Add(mBase, uint32(v365+v381)+8))
							v387 = F__bt_get_endpoint(m, l0, v383+int32(1), int32(0))
							mBase = m.M
							v388 = m.ExcPending
							if v388 != 0 {
								return
							} else {
								if v387 < int32(0) {
									v392 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
									v398 = *(*int32)(unsafe.Add(mBase, uint32(v392+(v387^int32(-1))*int32(56))+16))
									v407 = v398
								} else {
									v400 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
									v401 = int32(56)
									v406 = *(*int32)(unsafe.Add(mBase, uint32(v400+v387*v401-v401)+16))
									v407 = v406
								}
								v408 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v408
								*(*uint16)(unsafe.Add(mBase, uint32(v20)+48)) = uint16(v408)
								*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v407
								F_UnlockReleaseBuffer(m, v387)
								mBase = m.M
								v414 = m.ExcPending
								if v414 != 0 {
									return
								} else {
									v417 = v20 + int32(44)
									v419 = *(*int32)(unsafe.Add(mBase, uint32(v365)+24))
									v423 = F_CopyIndexTuple(m, v365+v419&int32(_a_F__bt_insert_parent_0))
									mBase = m.M
									v424 = m.ExcPending
									if v424 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v423))) = base.I32_rotr(v347, int32(16))
										v428 = F__bt_getstackbuf(m, l0, l1, v417, v328)
										mBase = m.M
										v429 = m.ExcPending
										if v429 != 0 {
											return
										} else {
											F_UnlockReleaseBuffer(m, l3)
											mBase = m.M
											v431 = m.ExcPending
											if v431 != 0 {
												return
											} else {
												if v428 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v547 = m.ExcPending
													if v547 != 0 {
														return
													} else {
														F_errcode(m, int32(33557032))
														mBase = m.M
														v550 = m.ExcPending
														if v550 != 0 {
															return
														} else {
															v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v347
															*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v328
															*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v551 + int32(4)
															F_errmsg_internal(m, int32(_a_F__bt_insert_parent_9), v20+int32(32))
															mBase = m.M
															v561 = m.ExcPending
															if v561 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F__bt_insert_parent_3), int32(2246), int32(_a_F__bt_insert_parent_8))
																mBase = m.M
																v566 = m.ExcPending
																if v566 != 0 {
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
													v434 = int32(0)
													v435 = *(*int32)(unsafe.Add(mBase, uint32(v417)+8))
													v436 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v423)+6)))
													v443 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v417)+4)))
													F__bt_insertonpg(m, l0, l1, v434, v428, l2, v435, v423, (v436&int32(_a_F__bt_insert_parent_10)+int32(7))&int32(_a_F__bt_insert_parent_11), (v443+int32(1))&int32(_a_F__bt_insert_parent_12), v434, l6)
													mBase = m.M
													v450 = m.ExcPending
													if v450 != 0 {
														return
													} else {
														F_pfree(m, v423)
														mBase = m.M
														v452 = m.ExcPending
														if v452 != 0 {
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
					v381 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v365)+16)))
					v383 = *(*int32)(unsafe.Add(mBase, uint32(v365+v381)+8))
					v387 = F__bt_get_endpoint(m, l0, v383+int32(1), int32(0))
					mBase = m.M
					v388 = m.ExcPending
					if v388 != 0 {
						return
					} else {
						if v387 < int32(0) {
							v392 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[0]))
							v398 = *(*int32)(unsafe.Add(mBase, uint32(v392+(v387^int32(-1))*int32(56))+16))
							v407 = v398
						} else {
							v400 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insert_parent[1]))
							v401 = int32(56)
							v406 = *(*int32)(unsafe.Add(mBase, uint32(v400+v387*v401-v401)+16))
							v407 = v406
						}
						v408 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v408
						*(*uint16)(unsafe.Add(mBase, uint32(v20)+48)) = uint16(v408)
						*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v407
						F_UnlockReleaseBuffer(m, v387)
						mBase = m.M
						v414 = m.ExcPending
						if v414 != 0 {
							return
						} else {
							v417 = v20 + int32(44)
							v419 = *(*int32)(unsafe.Add(mBase, uint32(v365)+24))
							v423 = F_CopyIndexTuple(m, v365+v419&int32(_a_F__bt_insert_parent_0))
							mBase = m.M
							v424 = m.ExcPending
							if v424 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v423))) = base.I32_rotr(v347, int32(16))
								v428 = F__bt_getstackbuf(m, l0, l1, v417, v328)
								mBase = m.M
								v429 = m.ExcPending
								if v429 != 0 {
									return
								} else {
									F_UnlockReleaseBuffer(m, l3)
									mBase = m.M
									v431 = m.ExcPending
									if v431 != 0 {
										return
									} else {
										if v428 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v547 = m.ExcPending
											if v547 != 0 {
												return
											} else {
												F_errcode(m, int32(33557032))
												mBase = m.M
												v550 = m.ExcPending
												if v550 != 0 {
													return
												} else {
													v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v347
													*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v328
													*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v551 + int32(4)
													F_errmsg_internal(m, int32(_a_F__bt_insert_parent_9), v20+int32(32))
													mBase = m.M
													v561 = m.ExcPending
													if v561 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F__bt_insert_parent_3), int32(2246), int32(_a_F__bt_insert_parent_8))
														mBase = m.M
														v566 = m.ExcPending
														if v566 != 0 {
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
											v434 = int32(0)
											v435 = *(*int32)(unsafe.Add(mBase, uint32(v417)+8))
											v436 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v423)+6)))
											v443 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v417)+4)))
											F__bt_insertonpg(m, l0, l1, v434, v428, l2, v435, v423, (v436&int32(_a_F__bt_insert_parent_10)+int32(7))&int32(_a_F__bt_insert_parent_11), (v443+int32(1))&int32(_a_F__bt_insert_parent_12), v434, l6)
											mBase = m.M
											v450 = m.ExcPending
											if v450 != 0 {
												return
											} else {
												F_pfree(m, v423)
												mBase = m.M
												v452 = m.ExcPending
												if v452 != 0 {
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
			v417 = l4
			v419 = *(*int32)(unsafe.Add(mBase, uint32(v365)+24))
			v423 = F_CopyIndexTuple(m, v365+v419&int32(_a_F__bt_insert_parent_0))
			mBase = m.M
			v424 = m.ExcPending
			if v424 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v423))) = base.I32_rotr(v347, int32(16))
				v428 = F__bt_getstackbuf(m, l0, l1, v417, v328)
				mBase = m.M
				v429 = m.ExcPending
				if v429 != 0 {
					return
				} else {
					F_UnlockReleaseBuffer(m, l3)
					mBase = m.M
					v431 = m.ExcPending
					if v431 != 0 {
						return
					} else {
						if v428 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v547 = m.ExcPending
							if v547 != 0 {
								return
							} else {
								F_errcode(m, int32(33557032))
								mBase = m.M
								v550 = m.ExcPending
								if v550 != 0 {
									return
								} else {
									v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v347
									*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v328
									*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v551 + int32(4)
									F_errmsg_internal(m, int32(_a_F__bt_insert_parent_9), v20+int32(32))
									mBase = m.M
									v561 = m.ExcPending
									if v561 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F__bt_insert_parent_3), int32(2246), int32(_a_F__bt_insert_parent_8))
										mBase = m.M
										v566 = m.ExcPending
										if v566 != 0 {
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
							v434 = int32(0)
							v435 = *(*int32)(unsafe.Add(mBase, uint32(v417)+8))
							v436 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v423)+6)))
							v443 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v417)+4)))
							F__bt_insertonpg(m, l0, l1, v434, v428, l2, v435, v423, (v436&int32(_a_F__bt_insert_parent_10)+int32(7))&int32(_a_F__bt_insert_parent_11), (v443+int32(1))&int32(_a_F__bt_insert_parent_12), v434, l6)
							mBase = m.M
							v450 = m.ExcPending
							if v450 != 0 {
								return
							} else {
								F_pfree(m, v423)
								mBase = m.M
								v452 = m.ExcPending
								if v452 != 0 {
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
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 float64
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 float64
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 float64
	_ = v210
	var v211 float64
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
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
							v74 = F_table_beginscan_parallel(m, v70, l2+int32(96), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v79 = int32(0)
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v76)+188))
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+140))
								v85 = m.T0[v84].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v76, v77, v66, int32(1), v79, l6, v79, int32(-1), int32(243), v13, v74)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									if l6 != 0 {
										v91 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[1]))
										if v91 == int32(0) {
										} else {
											v95 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[2])))
											if v95&int32(1) == int32(0) {
											} else {
												v100 = int32(_a_F__bt_parallel_scan_and_sort_0)
												v102 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
												v103 = int32(1)
												*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v102 + v103
												v106 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
												*(*int32)(unsafe.Add(mBase, uint32(v91))) = v106 + v103
												v110 = int32(0)
												v112 = int32(_a_F__bt_parallel_scan_and_sort_1)
												v113 = base.AtomicRmwOr32(m, v110, v112, v110)
												*(*int64)(unsafe.Add(mBase, uint32(v91+int32(80))+232)) = int64(3)
												v121 = base.AtomicRmwOr32(m, v110, v112, v110)
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
												*(*int32)(unsafe.Add(mBase, uint32(v91))) = v122 + v103
												v128 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
												*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v128 - v103
											}
										}
										v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										F_tuplesort_performsort(m, v132)
										mBase = m.M
										v134 = m.ExcPending
										if v134 != 0 {
											return
										} else {
											if l1 == int32(0) {
												v192 = base.AtomicRmwXchg32(m, l2, int32(36), int32(1))
												if v192 != 0 {
													F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_2))
													mBase = m.M
													v197 = m.ExcPending
													if v197 != 0 {
														return
													} else {
														v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
														v199 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
														v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
														v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
														if v205 == v199 {
															v208 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
														} else {
														}
														v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
														v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
														v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
														if v214 == int32(1) {
															v217 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
														} else {
														}
														v219 = int32(0)
														atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
														F_ConditionVariableSignal(m, l2+int32(24))
														mBase = m.M
														v225 = m.ExcPending
														if v225 != 0 {
															return
														} else {
															v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_tuplesort_end(m, v226)
															mBase = m.M
															v228 = m.ExcPending
															if v228 != 0 {
																return
															} else {
																if l1 != 0 {
																	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																	F_tuplesort_end(m, v229)
																	mBase = m.M
																	v231 = m.ExcPending
																	if v231 != 0 {
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
													v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
													v199 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
													v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
													v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
													if v205 == v199 {
														v208 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
													} else {
													}
													v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
													v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
													v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
													if v214 == int32(1) {
														v217 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
													} else {
													}
													v219 = int32(0)
													atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
													F_ConditionVariableSignal(m, l2+int32(24))
													mBase = m.M
													v225 = m.ExcPending
													if v225 != 0 {
														return
													} else {
														v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_tuplesort_end(m, v226)
														mBase = m.M
														v228 = m.ExcPending
														if v228 != 0 {
															return
														} else {
															if l1 != 0 {
																v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																F_tuplesort_end(m, v229)
																mBase = m.M
																v231 = m.ExcPending
																if v231 != 0 {
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
												v141 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[1]))
												if v141 == int32(0) {
												} else {
													v145 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[2])))
													if v145&int32(1) == int32(0) {
													} else {
														v150 = int32(_a_F__bt_parallel_scan_and_sort_0)
														v152 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
														v153 = int32(1)
														*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v152 + v153
														v156 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
														*(*int32)(unsafe.Add(mBase, uint32(v141))) = v156 + v153
														v160 = int32(0)
														v162 = int32(_a_F__bt_parallel_scan_and_sort_1)
														v163 = base.AtomicRmwOr32(m, v160, v162, v160)
														*(*int64)(unsafe.Add(mBase, uint32(v141+int32(80))+232)) = int64(4)
														v171 = base.AtomicRmwOr32(m, v160, v162, v160)
														v172 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
														*(*int32)(unsafe.Add(mBase, uint32(v141))) = v172 + v153
														v178 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
														*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v178 - v153
													}
												}
												v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												F_tuplesort_performsort(m, v187)
												mBase = m.M
												v189 = m.ExcPending
												if v189 != 0 {
													return
												} else {
													v192 = base.AtomicRmwXchg32(m, l2, int32(36), int32(1))
													if v192 != 0 {
														F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_2))
														mBase = m.M
														v197 = m.ExcPending
														if v197 != 0 {
															return
														} else {
															v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
															v199 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
															v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
															*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
															v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
															if v205 == v199 {
																v208 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
															} else {
															}
															v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
															v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
															*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
															v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
															if v214 == int32(1) {
																v217 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
															} else {
															}
															v219 = int32(0)
															atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
															F_ConditionVariableSignal(m, l2+int32(24))
															mBase = m.M
															v225 = m.ExcPending
															if v225 != 0 {
																return
															} else {
																v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																F_tuplesort_end(m, v226)
																mBase = m.M
																v228 = m.ExcPending
																if v228 != 0 {
																	return
																} else {
																	if l1 != 0 {
																		v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																		F_tuplesort_end(m, v229)
																		mBase = m.M
																		v231 = m.ExcPending
																		if v231 != 0 {
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
														v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
														v199 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
														v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
														v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
														if v205 == v199 {
															v208 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
														} else {
														}
														v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
														v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
														v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
														if v214 == int32(1) {
															v217 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
														} else {
														}
														v219 = int32(0)
														atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
														F_ConditionVariableSignal(m, l2+int32(24))
														mBase = m.M
														v225 = m.ExcPending
														if v225 != 0 {
															return
														} else {
															v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_tuplesort_end(m, v226)
															mBase = m.M
															v228 = m.ExcPending
															if v228 != 0 {
																return
															} else {
																if l1 != 0 {
																	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																	F_tuplesort_end(m, v229)
																	mBase = m.M
																	v231 = m.ExcPending
																	if v231 != 0 {
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
										v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										F_tuplesort_performsort(m, v182)
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return
										} else {
											if l1 == int32(0) {
												v192 = base.AtomicRmwXchg32(m, l2, int32(36), int32(1))
												if v192 != 0 {
													F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_2))
													mBase = m.M
													v197 = m.ExcPending
													if v197 != 0 {
														return
													} else {
														v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
														v199 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
														v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
														v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
														if v205 == v199 {
															v208 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
														} else {
														}
														v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
														v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
														v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
														if v214 == int32(1) {
															v217 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
														} else {
														}
														v219 = int32(0)
														atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
														F_ConditionVariableSignal(m, l2+int32(24))
														mBase = m.M
														v225 = m.ExcPending
														if v225 != 0 {
															return
														} else {
															v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_tuplesort_end(m, v226)
															mBase = m.M
															v228 = m.ExcPending
															if v228 != 0 {
																return
															} else {
																if l1 != 0 {
																	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																	F_tuplesort_end(m, v229)
																	mBase = m.M
																	v231 = m.ExcPending
																	if v231 != 0 {
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
													v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
													v199 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
													v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
													v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
													if v205 == v199 {
														v208 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
													} else {
													}
													v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
													v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
													v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
													if v214 == int32(1) {
														v217 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
													} else {
													}
													v219 = int32(0)
													atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
													F_ConditionVariableSignal(m, l2+int32(24))
													mBase = m.M
													v225 = m.ExcPending
													if v225 != 0 {
														return
													} else {
														v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_tuplesort_end(m, v226)
														mBase = m.M
														v228 = m.ExcPending
														if v228 != 0 {
															return
														} else {
															if l1 != 0 {
																v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																F_tuplesort_end(m, v229)
																mBase = m.M
																v231 = m.ExcPending
																if v231 != 0 {
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
												v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												F_tuplesort_performsort(m, v187)
												mBase = m.M
												v189 = m.ExcPending
												if v189 != 0 {
													return
												} else {
													v192 = base.AtomicRmwXchg32(m, l2, int32(36), int32(1))
													if v192 != 0 {
														F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_2))
														mBase = m.M
														v197 = m.ExcPending
														if v197 != 0 {
															return
														} else {
															v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
															v199 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
															v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
															*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
															v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
															if v205 == v199 {
																v208 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
															} else {
															}
															v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
															v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
															*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
															v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
															if v214 == int32(1) {
																v217 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
															} else {
															}
															v219 = int32(0)
															atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
															F_ConditionVariableSignal(m, l2+int32(24))
															mBase = m.M
															v225 = m.ExcPending
															if v225 != 0 {
																return
															} else {
																v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																F_tuplesort_end(m, v226)
																mBase = m.M
																v228 = m.ExcPending
																if v228 != 0 {
																	return
																} else {
																	if l1 != 0 {
																		v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																		F_tuplesort_end(m, v229)
																		mBase = m.M
																		v231 = m.ExcPending
																		if v231 != 0 {
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
														v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
														v199 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
														v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
														v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
														if v205 == v199 {
															v208 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
														} else {
														}
														v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
														v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
														v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
														if v214 == int32(1) {
															v217 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
														} else {
														}
														v219 = int32(0)
														atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
														F_ConditionVariableSignal(m, l2+int32(24))
														mBase = m.M
														v225 = m.ExcPending
														if v225 != 0 {
															return
														} else {
															v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_tuplesort_end(m, v226)
															mBase = m.M
															v228 = m.ExcPending
															if v228 != 0 {
																return
															} else {
																if l1 != 0 {
																	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																	F_tuplesort_end(m, v229)
																	mBase = m.M
																	v231 = m.ExcPending
																	if v231 != 0 {
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
					v74 = F_table_beginscan_parallel(m, v70, l2+int32(96), int32(0))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v79 = int32(0)
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v76)+188))
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+140))
						v85 = m.T0[v84].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v76, v77, v66, int32(1), v79, l6, v79, int32(-1), int32(243), v13, v74)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							if l6 != 0 {
								v91 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[1]))
								if v91 == int32(0) {
								} else {
									v95 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[2])))
									if v95&int32(1) == int32(0) {
									} else {
										v100 = int32(_a_F__bt_parallel_scan_and_sort_0)
										v102 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
										v103 = int32(1)
										*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v102 + v103
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
										*(*int32)(unsafe.Add(mBase, uint32(v91))) = v106 + v103
										v110 = int32(0)
										v112 = int32(_a_F__bt_parallel_scan_and_sort_1)
										v113 = base.AtomicRmwOr32(m, v110, v112, v110)
										*(*int64)(unsafe.Add(mBase, uint32(v91+int32(80))+232)) = int64(3)
										v121 = base.AtomicRmwOr32(m, v110, v112, v110)
										v122 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
										*(*int32)(unsafe.Add(mBase, uint32(v91))) = v122 + v103
										v128 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
										*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v128 - v103
									}
								}
								v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_tuplesort_performsort(m, v132)
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return
								} else {
									if l1 == int32(0) {
										v192 = base.AtomicRmwXchg32(m, l2, int32(36), int32(1))
										if v192 != 0 {
											F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_2))
											mBase = m.M
											v197 = m.ExcPending
											if v197 != 0 {
												return
											} else {
												v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
												v199 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
												v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
												v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
												if v205 == v199 {
													v208 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
												} else {
												}
												v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
												v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
												v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
												if v214 == int32(1) {
													v217 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
												} else {
												}
												v219 = int32(0)
												atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
												F_ConditionVariableSignal(m, l2+int32(24))
												mBase = m.M
												v225 = m.ExcPending
												if v225 != 0 {
													return
												} else {
													v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_tuplesort_end(m, v226)
													mBase = m.M
													v228 = m.ExcPending
													if v228 != 0 {
														return
													} else {
														if l1 != 0 {
															v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
															F_tuplesort_end(m, v229)
															mBase = m.M
															v231 = m.ExcPending
															if v231 != 0 {
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
											v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
											v199 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
											v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
											*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
											v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
											if v205 == v199 {
												v208 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
											} else {
											}
											v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
											v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
											*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
											v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
											if v214 == int32(1) {
												v217 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
											} else {
											}
											v219 = int32(0)
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
											F_ConditionVariableSignal(m, l2+int32(24))
											mBase = m.M
											v225 = m.ExcPending
											if v225 != 0 {
												return
											} else {
												v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_tuplesort_end(m, v226)
												mBase = m.M
												v228 = m.ExcPending
												if v228 != 0 {
													return
												} else {
													if l1 != 0 {
														v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
														F_tuplesort_end(m, v229)
														mBase = m.M
														v231 = m.ExcPending
														if v231 != 0 {
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
										v141 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[1]))
										if v141 == int32(0) {
										} else {
											v145 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[2])))
											if v145&int32(1) == int32(0) {
											} else {
												v150 = int32(_a_F__bt_parallel_scan_and_sort_0)
												v152 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
												v153 = int32(1)
												*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v152 + v153
												v156 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
												*(*int32)(unsafe.Add(mBase, uint32(v141))) = v156 + v153
												v160 = int32(0)
												v162 = int32(_a_F__bt_parallel_scan_and_sort_1)
												v163 = base.AtomicRmwOr32(m, v160, v162, v160)
												*(*int64)(unsafe.Add(mBase, uint32(v141+int32(80))+232)) = int64(4)
												v171 = base.AtomicRmwOr32(m, v160, v162, v160)
												v172 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
												*(*int32)(unsafe.Add(mBase, uint32(v141))) = v172 + v153
												v178 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3]))
												*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_scan_and_sort[3])) = v178 - v153
											}
										}
										v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										F_tuplesort_performsort(m, v187)
										mBase = m.M
										v189 = m.ExcPending
										if v189 != 0 {
											return
										} else {
											v192 = base.AtomicRmwXchg32(m, l2, int32(36), int32(1))
											if v192 != 0 {
												F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_2))
												mBase = m.M
												v197 = m.ExcPending
												if v197 != 0 {
													return
												} else {
													v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
													v199 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
													v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
													v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
													if v205 == v199 {
														v208 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
													} else {
													}
													v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
													v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
													v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
													if v214 == int32(1) {
														v217 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
													} else {
													}
													v219 = int32(0)
													atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
													F_ConditionVariableSignal(m, l2+int32(24))
													mBase = m.M
													v225 = m.ExcPending
													if v225 != 0 {
														return
													} else {
														v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_tuplesort_end(m, v226)
														mBase = m.M
														v228 = m.ExcPending
														if v228 != 0 {
															return
														} else {
															if l1 != 0 {
																v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																F_tuplesort_end(m, v229)
																mBase = m.M
																v231 = m.ExcPending
																if v231 != 0 {
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
												v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
												v199 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
												v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
												v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
												if v205 == v199 {
													v208 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
												} else {
												}
												v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
												v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
												v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
												if v214 == int32(1) {
													v217 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
												} else {
												}
												v219 = int32(0)
												atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
												F_ConditionVariableSignal(m, l2+int32(24))
												mBase = m.M
												v225 = m.ExcPending
												if v225 != 0 {
													return
												} else {
													v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_tuplesort_end(m, v226)
													mBase = m.M
													v228 = m.ExcPending
													if v228 != 0 {
														return
													} else {
														if l1 != 0 {
															v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
															F_tuplesort_end(m, v229)
															mBase = m.M
															v231 = m.ExcPending
															if v231 != 0 {
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
								v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_tuplesort_performsort(m, v182)
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return
								} else {
									if l1 == int32(0) {
										v192 = base.AtomicRmwXchg32(m, l2, int32(36), int32(1))
										if v192 != 0 {
											F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_2))
											mBase = m.M
											v197 = m.ExcPending
											if v197 != 0 {
												return
											} else {
												v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
												v199 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
												v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
												v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
												if v205 == v199 {
													v208 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
												} else {
												}
												v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
												v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
												v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
												if v214 == int32(1) {
													v217 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
												} else {
												}
												v219 = int32(0)
												atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
												F_ConditionVariableSignal(m, l2+int32(24))
												mBase = m.M
												v225 = m.ExcPending
												if v225 != 0 {
													return
												} else {
													v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_tuplesort_end(m, v226)
													mBase = m.M
													v228 = m.ExcPending
													if v228 != 0 {
														return
													} else {
														if l1 != 0 {
															v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
															F_tuplesort_end(m, v229)
															mBase = m.M
															v231 = m.ExcPending
															if v231 != 0 {
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
											v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
											v199 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
											v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
											*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
											v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
											if v205 == v199 {
												v208 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
											} else {
											}
											v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
											v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
											*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
											v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
											if v214 == int32(1) {
												v217 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
											} else {
											}
											v219 = int32(0)
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
											F_ConditionVariableSignal(m, l2+int32(24))
											mBase = m.M
											v225 = m.ExcPending
											if v225 != 0 {
												return
											} else {
												v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_tuplesort_end(m, v226)
												mBase = m.M
												v228 = m.ExcPending
												if v228 != 0 {
													return
												} else {
													if l1 != 0 {
														v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
														F_tuplesort_end(m, v229)
														mBase = m.M
														v231 = m.ExcPending
														if v231 != 0 {
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
										v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										F_tuplesort_performsort(m, v187)
										mBase = m.M
										v189 = m.ExcPending
										if v189 != 0 {
											return
										} else {
											v192 = base.AtomicRmwXchg32(m, l2, int32(36), int32(1))
											if v192 != 0 {
												F_s_lock(m, l2+int32(36), int32(_a_F__bt_parallel_scan_and_sort_2))
												mBase = m.M
												v197 = m.ExcPending
												if v197 != 0 {
													return
												} else {
													v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
													v199 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
													v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
													v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
													if v205 == v199 {
														v208 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
													} else {
													}
													v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
													v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
													v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
													if v214 == int32(1) {
														v217 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
													} else {
													}
													v219 = int32(0)
													atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
													F_ConditionVariableSignal(m, l2+int32(24))
													mBase = m.M
													v225 = m.ExcPending
													if v225 != 0 {
														return
													} else {
														v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_tuplesort_end(m, v226)
														mBase = m.M
														v228 = m.ExcPending
														if v228 != 0 {
															return
														} else {
															if l1 != 0 {
																v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																F_tuplesort_end(m, v229)
																mBase = m.M
																v231 = m.ExcPending
																if v231 != 0 {
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
												v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
												v199 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v198 + v199
												v202 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v85, v202)
												v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
												if v205 == v199 {
													v208 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v208)
												} else {
												}
												v210 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
												v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v210, v211)
												v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
												if v214 == int32(1) {
													v217 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v217)
												} else {
												}
												v219 = int32(0)
												atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+36)), uint32(v219))
												F_ConditionVariableSignal(m, l2+int32(24))
												mBase = m.M
												v225 = m.ExcPending
												if v225 != 0 {
													return
												} else {
													v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_tuplesort_end(m, v226)
													mBase = m.M
													v228 = m.ExcPending
													if v228 != 0 {
														return
													} else {
														if l1 != 0 {
															v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
															F_tuplesort_end(m, v229)
															mBase = m.M
															v231 = m.ExcPending
															if v231 != 0 {
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
	var v59 int64
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
	var v79 int64
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
		v59 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
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
					v79 = *(*int64)(unsafe.Add(mBase, uint32(v75+v63<<(uint(int32(3))%32))))
					*(*int64)(unsafe.Add(mBase, uint32(v75))) = v79
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
				v90 = v88 << (uint(int32(3)) % 32)
				if v90 == int32(0) {
					v100 = v88
				} else {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					base.MemoryCopy(m, v93, v93+v87<<(uint(int32(3))%32), v90)
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
				v90 = v88 << (uint(int32(3)) % 32)
				if v90 == int32(0) {
					v100 = v88
				} else {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					base.MemoryCopy(m, v93, v93+v87<<(uint(int32(3))%32), v90)
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
						F_errfinish(m, int32(_a_F__bt_saoparray_shrink_1), int32(1234), int32(_a_F__bt_saoparray_shrink_2))
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
					v59 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
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
								v79 = *(*int64)(unsafe.Add(mBase, uint32(v75+v63<<(uint(int32(3))%32))))
								*(*int64)(unsafe.Add(mBase, uint32(v75))) = v79
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
							v90 = v88 << (uint(int32(3)) % 32)
							if v90 == int32(0) {
								v100 = v88
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
								base.MemoryCopy(m, v93, v93+v87<<(uint(int32(3))%32), v90)
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
							v90 = v88 << (uint(int32(3)) % 32)
							if v90 == int32(0) {
								v100 = v88
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
								base.MemoryCopy(m, v93, v93+v87<<(uint(int32(3))%32), v90)
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
									F_errfinish(m, int32(_a_F__bt_saoparray_shrink_1), int32(1234), int32(_a_F__bt_saoparray_shrink_2))
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
	var v160 int32
	_ = v160
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
	v160 = F_errdetail(m, int32(_a_F_bt_entry_unique_check_6), v12)
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
	F_errfinish(m, int32(_a_F_bt_entry_unique_check_8), int32(1001), int32(_a_F_bt_entry_unique_check_9))
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
func F_bt_index_check(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v21 int64
	_ = v21
	var v34 int32
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v12 < int32(2) {
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(base.B2i32(v15 != int64(0)))
		if v12 == int32(2) {
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(base.B2i32(v21 != int64(0)))
		}
	}
	F_amcheck_lock_relation_and_check(m, base.I32_wrap_i64(v9), int32(403), int32(_a_F_bt_index_check_0), int32(1), v7+int32(12))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int64(0)
	} else {
		m.G0 = v7 + int32(16)
		return int64(0)
	}
}
func F_bt_leftmost_ignoring_half_dead(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var __phi20 int32
	_ = __phi20
	var v22 int32
	_ = v22
	var __phi22 int32
	_ = __phi22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v15 == int32(0) {
		v96 = v14
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v96
L2:
	;
	__phi20 = l1
	__phi22 = v15
	v20 = __phi20
	v22 = __phi22
	goto L3
L3:
	;
	v27 = F_palloc_btree_page(m, l0, v22)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_pfree(m, v27)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L25
	}
L5:
	;
	return int32(0)
L6:
	;
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+16)))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_bt_leftmost_ignoring_half_dead[0]))
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if l1 == v22 {
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
	v37 = v31 + v27
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+12)))
	if base.B2i32(v38&int32(16) == int32(0))|base.B2i32(v20 == v22) != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v45 != v20 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
	v50 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	F_pfree(m, v27)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L23
	}
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v56 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_leftmost_ignoring_half_dead_0), v12+int32(16))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v22
	v67 = int64(32)
	v68 = base.I64_rotl(v47, v67)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+12)) = uint32(v68)
	v71 = int64(base.Ui64(v68) >> (uint(v67) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+8)) = uint32(v71)
	F_errdetail_internal(m, int32(_a_F_bt_leftmost_ignoring_half_dead_1), v12)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_bt_leftmost_ignoring_half_dead_2), int32(1053), int32(_a_F_bt_leftmost_ignoring_half_dead_3))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	if v82 != 0 {
		__phi20 = v22
		__phi22 = v82
		v20 = __phi20
		v22 = __phi22
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v96 = v14
	goto L1
L25:
	;
	v96 = int32(0)
	goto L1
}
func F_bt_page_stats(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_bt_page_stats_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_bt_page_stats_internal(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int64
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
	var v153 int64
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_superuser(m)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			if v17 != 0 {
				v19 = F_textToQualifiedNameList(m, v12)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					v21 = F_makeRangeVarFromNameList(m, v19)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						v24 = F_relation_openrv(m, v21, int32(1))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int64(0)
						} else {
							if l1 != 0 {
								v28 = v16
							} else {
								v28 = v16 & int64(4294967295)
							}
							F_bt_index_block_validate(m, v24, v28)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int64(0)
							} else {
								v31 = base.I32_wrap_i64(v28)
								v32 = F_ReadBuffer(m, v24, v31)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int64(0)
								} else {
									F_LockBufferInternal(m, v32, int32(1))
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v9)+208)) = int64(-1)
										v39 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v9)+220)) = uint16(v39)
										*(*int64)(unsafe.Add(mBase, uint32(v9)+196)) = int64(0)
										F_GetBTPageStatistics(m, v31, v32, v9+int32(176))
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return int64(0)
										} else {
											F_UnlockReleaseBuffer(m, v32)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int64(0)
											} else {
												F_relation_close(m, v24, int32(1))
												mBase = m.M
												v51 = m.ExcPending
												if v51 != 0 {
													return int64(0)
												} else {
													v55 = F_get_call_result_type(m, l0, int32(0), v9+int32(268))
													mBase = m.M
													v56 = m.ExcPending
													if v56 != 0 {
														return int64(0)
													} else {
														if v55 != int32(1) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return int64(0)
															} else {
																F_errmsg_internal(m, int32(_a_F_bt_page_stats_internal_0), int32(0))
																mBase = m.M
																v182 = m.ExcPending
																if v182 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_bt_page_stats_internal_1), int32(298), int32(_a_F_bt_page_stats_internal_2))
																	mBase = m.M
																	v187 = m.ExcPending
																	if v187 != 0 {
																		return int64(0)
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
																return int64(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9)+224)) = v64
																v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9)+204)))
																*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v67
																v72 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_4), v9+int32(144))
																mBase = m.M
																v73 = m.ExcPending
																if v73 != 0 {
																	return int64(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+228)) = v72
																	v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)+180))
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v75
																	v80 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(128))
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int64(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+232)) = v80
																		v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+184))
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v83
																		v88 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(112))
																		mBase = m.M
																		v89 = m.ExcPending
																		if v89 != 0 {
																			return int64(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+236)) = v88
																			v91 = *(*int32)(unsafe.Add(mBase, uint32(v9)+200))
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v91
																			v96 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(96))
																			mBase = m.M
																			v97 = m.ExcPending
																			if v97 != 0 {
																				return int64(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v9)+240)) = v96
																				v99 = *(*int32)(unsafe.Add(mBase, uint32(v9)+188))
																				*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v99
																				v104 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(80))
																				mBase = m.M
																				v105 = m.ExcPending
																				if v105 != 0 {
																					return int64(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v9)+244)) = v104
																					v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+196))
																					*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v107
																					v112 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9-int32(-64))
																					mBase = m.M
																					v113 = m.ExcPending
																					if v113 != 0 {
																						return int64(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v9)+248)) = v112
																						v115 = *(*int32)(unsafe.Add(mBase, uint32(v9)+208))
																						*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v115
																						v120 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(48))
																						mBase = m.M
																						v121 = m.ExcPending
																						if v121 != 0 {
																							return int64(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v9)+252)) = v120
																							v123 = *(*int32)(unsafe.Add(mBase, uint32(v9)+212))
																							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v123
																							v128 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(32))
																							mBase = m.M
																							v129 = m.ExcPending
																							if v129 != 0 {
																								return int64(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v9)+256)) = v128
																								v131 = *(*int32)(unsafe.Add(mBase, uint32(v9)+216))
																								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v131
																								v136 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_3), v9+int32(16))
																								mBase = m.M
																								v137 = m.ExcPending
																								if v137 != 0 {
																									return int64(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v9)+260)) = v136
																									v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+220)))
																									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v139
																									v142 = F_psprintf(m, int32(_a_F_bt_page_stats_internal_5), v9)
																									mBase = m.M
																									v143 = m.ExcPending
																									if v143 != 0 {
																										return int64(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v9)+264)) = v142
																										v145 = *(*int32)(unsafe.Add(mBase, uint32(v9)+268))
																										v146 = F_TupleDescGetAttInMetadata(m, v145)
																										mBase = m.M
																										v147 = m.ExcPending
																										if v147 != 0 {
																											return int64(0)
																										} else {
																											v150 = F_BuildTupleFromCStrings(m, v146, v9+int32(224))
																											mBase = m.M
																											v151 = m.ExcPending
																											if v151 != 0 {
																												return int64(0)
																											} else {
																												v152 = *(*int32)(unsafe.Add(mBase, uint32(v150)+16))
																												v153 = F_HeapTupleHeaderGetDatum(m, v152)
																												mBase = m.M
																												v154 = m.ExcPending
																												if v154 != 0 {
																													return int64(0)
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
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v165 = m.ExcPending
					if v165 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_bt_page_stats_internal_6), int32(0))
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_bt_page_stats_internal_1), int32(277), int32(_a_F_bt_page_stats_internal_2))
							mBase = m.M
							v174 = m.ExcPending
							if v174 != 0 {
								return int64(0)
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
