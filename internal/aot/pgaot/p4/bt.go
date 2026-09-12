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
	var v25 int32
	_ = v25
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
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
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
	F_errmsg_internal(m, v91, v10)
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
	v25 = int32(0)
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
	v30 = v25 << (uint(int32(2)) % 32)
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
	v50 = v25 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+10)))
	if v50 < v52 {
		v25 = v50
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
	v89 = v67
	v91 = int32(263528)
	v94 = int32(4295)
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
	v89 = v78
	v91 = int32(263568)
	v94 = int32(4298)
	goto L5
L31:
	;
	F_errfinish(m, int32(488920), v94, int32(404520))
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v164 int64
	_ = v164
	var v165 int32
	_ = v165
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v261 int32
	_ = v261
	var v264 int64
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 - int32(-64)
	return v335
L2:
	;
	return int32(0)
L3:
	;
	if v13 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = v13
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l0
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v264
	v268 = int32(0)
	v271 = F_ExtendBufferedRel(m, v9+int32(-56), v268, v268, int32(8))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L2
	} else {
		goto L86
	}
L7:
	;
	v27 = F_ReadBuffer(m, l0, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	F_ReleaseBuffer(m, v27)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L2
	} else {
		goto L82
	}
L10:
	;
	v29 = F_ConditionalLockBuffer(m, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	if v29 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v27 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	v229 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L2
	} else {
		goto L78
	}
L15:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+14)))
	if v49 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34+(v27^int32(-1))<<(uint(int32(2))%32))))
	v48 = v40
	goto L15
L17:
	;
	goto L18
L18:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v48 = v42 + v27<<(uint(int32(13))%32) + int32(-8192)
	goto L15
L19:
	;
	if v48&int32(3) != 0 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	goto L21
L21:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+16)))
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48+v93)+12)))
	if v95&int32(4) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v335 = v27
	goto L1
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+10)) = int32(1572864)
	v84 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+18)) = uint16(v84)
	v90 = int32(8176)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+16)) = uint16(v90)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+14)) = uint16(v90)
	goto L22
L24:
	;
	v78 = F___memset(m, v48, int32(0), int32(8192))
	mBase = m.M
	goto L23
L25:
	;
	goto L24
L32:
	;
	v213 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L2
	} else {
		goto L71
	}
L33:
	;
	v100 = int64(3)
	if v95&int32(256) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v48)+24))
	v104 = v103
	goto L36
L35:
	;
	v104 = v100
	goto L36
L36:
	;
	v105 = F_GlobalVisCheckRemovableFullXid(m, l1, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	if v105 == int32(0) {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+118)))
	if v110 != int32(112) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if v48&int32(3) != 0 {
		goto L63
	} else {
		goto L64
	}
L40:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v114 <= int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v117
	v119 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v24
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+16)))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v122)+13)))
	if v124&int32(1) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v48)+24))
	v128 = v127
	goto L44
L43:
	;
	v128 = v100
	goto L44
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v128
	if base.Ui32(v114) < base.Ui32(int32(2)) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+56)) = uint8(v152)
	F_XLogBeginInsert(m)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L2
	} else {
		goto L58
	}
L46:
	;
	v152 = int32(0)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+118)))
	if v134 != int32(112) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v152 = int32(0)
	goto L45
L50:
	;
	goto L51
L51:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	goto L52
L52:
	;
	if base.Ui32(v139) < base.Ui32(int32(12000)) {
		v152 = int32(1)
		goto L45
	} else {
		goto L53
	}
L53:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	if v142 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v152 = int32(0)
	goto L45
L55:
	;
	goto L56
L56:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+119)))
	switch v148 - int32(109) {
	case 0, 5:
		goto L57
	default:
		v152 = int32(0)
		goto L45
	}
L57:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+104)))
	v152 = v151
	goto L45
L58:
	;
	F_XLogRegisterData(m, v9+int32(-32), int32(25))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v164 = F_XLogInsert(m, int32(11), int32(208))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	goto L39
L61:
	;
	v335 = v27
	goto L1
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+10)) = int32(1572864)
	v201 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+18)) = uint16(v201)
	v207 = int32(8176)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+16)) = uint16(v207)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+14)) = uint16(v207)
	goto L61
L63:
	;
	v195 = F___memset(m, v48, int32(0), int32(8192))
	mBase = m.M
	goto L62
L64:
	;
	goto L63
L71:
	;
	if v213 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_errmsg_internal(m, int32(404199), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L2
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	F_LockBuffer(m, v27, int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L2
	} else {
		goto L77
	}
L75:
	;
	F_errfinish(m, int32(494298), int32(960), int32(334910))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	goto L9
L78:
	;
	if v229 == int32(0) {
		goto L9
	} else {
		goto L79
	}
L79:
	;
	F_errmsg_internal(m, int32(404231), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(494298), int32(965), int32(334910))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	goto L9
L82:
	;
	v247 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	if v247 != int32(-1) {
		v24 = v247
		goto L7
	} else {
		goto L84
	}
L84:
	;
	goto L8
L85:
	;
	if v290&int32(3) != 0 {
		goto L92
	} else {
		goto L93
	}
L86:
	;
	if v271 < int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v276+(v271^int32(-1))<<(uint(int32(2))%32))))
	v290 = v282
	goto L85
L88:
	;
	goto L89
L89:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v290 = v284 + v271<<(uint(int32(13))%32) + int32(-8192)
	goto L85
L90:
	;
	v335 = v271
	goto L1
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+10)) = int32(1572864)
	v323 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v290)+18)) = uint16(v323)
	v329 = int32(8176)
	*(*uint16)(unsafe.Add(mBase, uint32(v290)+16)) = uint16(v329)
	*(*uint16)(unsafe.Add(mBase, uint32(v290)+14)) = uint16(v329)
	goto L90
L92:
	;
	v317 = F___memset(m, v290, int32(0), int32(8192))
	mBase = m.M
	goto L91
L93:
	;
	goto L92
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v346 int32
	_ = v346
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v385 int32
	_ = v385
	var v414 int32
	_ = v414
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v23)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v27 <= v26 {
		v414 = v23
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v414
L2:
	;
	v39 = v26
	goto L5
L3:
	;
	v414 = int32(0)
	goto L1
L4:
	;
	v385 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v385)
	goto L3
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v49 = v46 + v39*int32(48)
	v50 = int32(0)
	if l6 != 0 {
		v78 = v50
		v82 = v50
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l6 != 0 {
		goto L3
	} else {
		goto L117
	}
L7:
	;
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+4)))
	if l3 < v83 {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v54 = v52 & int32(65536)
	v55 = int32(1)
	if l1 != v55 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v62 = v52 & int32(131072)
	if l1 == int32(-1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if v54 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v78 = v55
	v82 = int32(0)
	goto L7
L12:
	;
	if v62 != 0 {
		v78 = v55
		v82 = int32(0)
		goto L7
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v66 = int32(0)
	v78 = v66
	v82 = base.B2i32(l1 == int32(-1))&base.B2i32(v54 != v66) | base.B2i32(l1 == int32(1))&base.B2i32(v62 != v66)
	goto L7
L15:
	;
	goto L14
L16:
	;
	goto L6
L17:
	;
	v329 = int32(1)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v332 = v330 + v329
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v332
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v332 < v334 {
		v39 = v332
		goto L5
	} else {
		goto L116
	}
L18:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v85&int32(7864320) != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v296 = int32(0)
	if l5 == v296 {
		v414 = v296
		goto L1
	} else {
		goto L112
	}
L20:
	;
	v293 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v293)
	v414 = v293
	goto L1
L21:
	;
	if l6 == int32(0) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v85&int32(4) != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v90 = int32(0)
	v92 = F__bt_advance_array_keys(m, l0, v90, l2, l3, l4, v39, v90)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	v414 = v92
	goto L1
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v49)+44))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v99&int32(1) != 0 {
		v346 = v98
		goto L16
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v218 = F_index_getattr_2(m, l2, v83, l4, v20+int32(14))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L25
	} else {
		goto L82
	}
L30:
	;
	v112 = v98
	goto L31
L31:
	;
	v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(v112)+4)))
	if l3 < v119 {
		goto L17
	} else {
		goto L33
	}
L32:
	;
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+6)))
	switch v174 - int32(1) {
	case 0:
		goto L64
	case 1:
		goto L68
	default:
		goto L65
	case 3:
		goto L67
	case 4:
		goto L66
	}
L33:
	;
	v123 = F_index_getattr_2(m, l2, v119, l4, v20+int32(15))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+15)))
	if v125 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if l6 != 0 {
		goto L3
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v112)+44))
	v154 = F_FunctionCall2Coll(m, v112+int32(16), v152, v123, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L25
	} else {
		goto L52
	}
L38:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v49)+44))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v129&int32(33554432) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if l1 != int32(-1) {
		goto L3
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if l1 != int32(1) {
		goto L3
	} else {
		goto L47
	}
L42:
	;
	if v128 == v112 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v137 = int32(196608)
	goto L45
L44:
	;
	v137 = int32(131072)
	goto L45
L45:
	;
	if v137&v129 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	goto L3
L47:
	;
	if v128 == v112 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v144 = int32(196608)
	goto L50
L49:
	;
	v144 = int32(65536)
	goto L50
L50:
	;
	if v144&v129 == int32(0) {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	goto L4
L52:
	;
	if v154 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v159 = int32(1)
	goto L55
L54:
	;
	v159 = int32(0) - v154
	goto L55
L55:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v160&int32(16777216) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v163 = v159
	goto L58
L57:
	;
	v163 = v154
	goto L58
L58:
	;
	if v163|v160&int32(16) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v170 = v112 + int32(48)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v171&int32(1) != 0 {
		v346 = v170
		goto L16
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L32
L62:
	;
	v112 = v170
	goto L31
L63:
	;
	if l6 != 0 {
		goto L72
	} else {
		goto L73
	}
L64:
	;
	v199 = int32(base.Ui32(v163) >> (uint(int32(31)) % 32))
	goto L63
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L25
	} else {
		goto L69
	}
L66:
	;
	v199 = base.B2i32(int32(0) < v163)
	goto L63
L67:
	;
	v199 = base.B2i32(int32(0) <= v163)
	goto L63
L68:
	;
	v199 = base.B2i32(v163 <= int32(0))
	goto L63
L69:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v187
	F_errmsg_internal(m, int32(466806), v20)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(488920), int32(3153), int32(361580))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L25
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	if v199&int32(1) != 0 {
		goto L17
	} else {
		goto L81
	}
L73:
	;
	if v199&int32(1) != 0 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	if v160&int32(65536) != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v207 = base.B2i32(l1 == int32(1))
	goto L77
L76:
	;
	v207 = int32(0)
	goto L77
L77:
	;
	if v207 != 0 {
		goto L20
	} else {
		goto L78
	}
L78:
	;
	v208 = int32(0)
	if l1 != int32(-1) {
		v414 = v208
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v160&int32(131072) != 0 {
		goto L20
	} else {
		goto L80
	}
L80:
	;
	v414 = v208
	goto L1
L81:
	;
	v414 = int32(0)
	goto L1
L82:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v220&int32(1) != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+14)))
	if (v223^base.B2i32(v220&int32(64) == int32(0)))&int32(1) != 0 {
		goto L17
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+14)))
	if v234 == int32(1) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	if v78 != 0 {
		goto L20
	} else {
		goto L87
	}
L87:
	;
	if v220&int32(262144) != 0 {
		goto L17
	} else {
		goto L88
	}
L88:
	;
	v414 = int32(0)
	goto L1
L89:
	;
	if l6 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v49)+44))
	v272 = F_FunctionCall2Coll(m, v49+int32(16), v270, v218, v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L25
	} else {
		goto L109
	}
L92:
	;
	if v220&int32(33554432) != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	if v220&int32(262144) == int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v243 = int32(0)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v246 = F__bt_advance_array_keys(m, l0, v243, l2, l3, l4, v244, v243)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L25
	} else {
		goto L95
	}
L95:
	;
	v414 = v246
	goto L1
L96:
	;
	if v78 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	if v78 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L99:
	;
	if base.B2i32(l1 == int32(-1))&v82 != 0 {
		goto L20
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	if l1 == int32(-1) {
		goto L20
	} else {
		goto L103
	}
L102:
	;
	v414 = int32(0)
	goto L1
L103:
	;
	v414 = int32(0)
	goto L1
L104:
	;
	if base.B2i32(l1 == int32(1))&v82 != 0 {
		goto L20
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if l1 == int32(1) {
		goto L20
	} else {
		goto L108
	}
L107:
	;
	v414 = int32(0)
	goto L1
L108:
	;
	v414 = int32(0)
	goto L1
L109:
	;
	if v272 != 0 {
		goto L17
	} else {
		goto L110
	}
L110:
	;
	if v78 == int32(0) {
		goto L19
	} else {
		goto L111
	}
L111:
	;
	goto L20
L112:
	;
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+6)))
	if v299 != int32(3) {
		v414 = v296
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v302&int32(32) == int32(0) {
		v414 = v296
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v307 = int32(0)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v310 = F__bt_advance_array_keys(m, l0, v307, l2, l3, l4, v308, v307)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L25
	} else {
		goto L115
	}
L115:
	;
	v414 = v310
	goto L1
L116:
	;
	v414 = v329
	goto L1
L117:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v346-int32(48))))
	if v358&int32(65536) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v361 = base.B2i32(l1 == int32(1))
	goto L120
L119:
	;
	v361 = int32(0)
	goto L120
L120:
	;
	if v361 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	if l1 != int32(-1) {
		goto L3
	} else {
		goto L122
	}
L122:
	;
	if v358&int32(131072) == int32(0) {
		goto L3
	} else {
		goto L123
	}
L123:
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
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
	if v236 == int32(0) {
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
	v236 = v18
	goto L2
L8:
	;
	v236 = v152
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
	v208 = m.ExcPending
	if v208 != 0 {
		goto L6
	} else {
		goto L60
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L6
	} else {
		goto L56
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
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
	v42 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v32^int32(-1))<<(uint(int32(2))%32))))
	v56 = v48
	goto L19
L21:
	;
	goto L22
L22:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v56 = v50 + v32<<(uint(int32(13))%32) + int32(-8192)
	goto L19
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)+24))
	if v64 != int32(340322) {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v68 = v56 + int32(28)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if base.Ui32(v69-int32(5)) <= base.Ui32(int32(-4)) {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
	if v74 == int32(0) {
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
	v77 = int32(0)
	F_LockBuffer(m, v32, v77)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v56)+36))
	v87 = v32
	v89 = v74
	goto L33
L30:
	;
	F_ReleaseBuffer(m, v32)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v152 = v77
	goto L26
L32:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	if v147 != v83 {
		goto L16
	} else {
		goto L51
	}
L33:
	;
	if v87 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L48
	}
L35:
	;
	F_LockBuffer(m, v87, int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v96 = F_ReleaseAndReadBuffer(m, v87, l0, v89)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	F_LockBuffer(m, v96, int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	F__bt_checkpage(m, l0, v96)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	if v96 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+16)))
	v122 = v121 + v120
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+12)))
	if v123&int32(20) == int32(0) {
		goto L32
	} else {
		goto L46
	}
L43:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v106+(v96^int32(-1))<<(uint(int32(2))%32))))
	v120 = v112
	goto L42
L44:
	;
	goto L45
L45:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v120 = v114 + v96<<(uint(int32(13))%32) + int32(-8192)
	goto L42
L46:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v128 != 0 {
		v87 = v96
		v89 = v128
		goto L33
	} else {
		goto L47
	}
L47:
	;
	goto L34
L48:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v133 + int32(4)
	F_errmsg_internal(m, int32(673874), v22+int32(-16))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(494298), int32(651), int32(83811))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
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
	v152 = v96
	goto L26
L52:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v168 + int32(4)
	F_errmsg(m, int32(406133), v24)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(494298), int32(612), int32(83811))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
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
	v186 = m.ExcPending
	if v186 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = int64(8589934596)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v188 + int32(4)
	F_errmsg(m, int32(468355), v22+int32(-48))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(494298), int32(621), int32(83811))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
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
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v209 + int32(4)
	F_errmsg_internal(m, int32(54904), v22+int32(-32))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(494298), int32(658), int32(83811))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
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
	if v236 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v263 = v236
	v264 = v257
	goto L71
L67:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v243+(v236^int32(-1))<<(uint(int32(2))%32))))
	v257 = v249
	goto L66
L68:
	;
	goto L69
L69:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v257 = v251 + v236<<(uint(int32(13))%32) + int32(-8192)
	goto L66
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L6
	} else {
		goto L105
	}
L71:
	;
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v264)+16)))
	v268 = v264 + v267
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+12)))
	if v269&int32(20) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L102
	}
L73:
	;
	goto L72
L74:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v321<<(uint(int32(2))%32)+v264)+20))
	v328 = v264 + v325&int32(32767)
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v328))))
	v332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v328)+2)))
	v335 = F__bt_relandgetbuf(m, l0, v263, v329<<(uint(int32(16))%32)|v332, int32(1))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L6
	} else {
		goto L98
	}
L75:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v318 != 0 {
		goto L95
	} else {
		goto L96
	}
L76:
	;
	v297 = F__bt_relandgetbuf(m, l0, v263, v295, int32(1))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
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
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v292 == int32(0) {
		goto L73
	} else {
		goto L90
	}
L80:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v274 != 0 {
		v295 = v274
		goto L76
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v268)+8))
	if v276 == l1 {
		v397 = v263
		goto L1
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	if base.Ui32(v276) < base.Ui32(l1) {
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
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v264)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v281) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v291 = int32(base.Ui32(v281+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	goto L89
L88:
	;
	v291 = int32(0)
	goto L89
L89:
	;
	v321 = v291
	goto L74
L90:
	;
	v295 = v292
	goto L76
L91:
	;
	if v297 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v302+(v297^int32(-1))<<(uint(int32(2))%32))))
	v263 = v297
	v264 = v308
	goto L71
L93:
	;
	goto L94
L94:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v263 = v297
	v264 = v310 + v297<<(uint(int32(13))%32) + int32(-8192)
	goto L71
L95:
	;
	v319 = int32(2)
	goto L97
L96:
	;
	v319 = int32(1)
	goto L97
L97:
	;
	v321 = v319
	goto L74
L98:
	;
	if v335 < int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v340+(v335^int32(-1))<<(uint(int32(2))%32))))
	v263 = v335
	v264 = v346
	goto L71
L100:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v263 = v335
	v264 = v348 + v335<<(uint(int32(13))%32) + int32(-8192)
	goto L71
L102:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v358 + int32(4)
	F_errmsg_internal(m, int32(675745), v12+int32(16))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(493016), int32(2653), int32(87883))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
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
	v378 = m.ExcPending
	if v378 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v379 + int32(4)
	F_errmsg_internal(m, int32(673835), v12)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(493016), int32(2666), int32(87883))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
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
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int64
	_ = v287
	var v288 int32
	_ = v288
	var v289 int64
	_ = v289
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	if l5 != 0 {
		if l2 < int32(0) {
			v25 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v25+(l2^int32(-1))<<(uint(int32(6))%32))+16))
			v40 = v31
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, _consts[4]))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v33+l2<<(uint(int32(6))%32)+int32(-64))+16))
			v40 = v39
		}
		if l3 < int32(0) {
			v44 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+(l3^int32(-1))<<(uint(int32(6))%32))+16))
			v59 = v50
		} else {
			v52 = *(*int32)(unsafe.Add(mBase, _consts[4]))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v52+l3<<(uint(int32(6))%32)+int32(-64))+16))
			v59 = v58
		}
		if l2 < int32(0) {
			v63 = *(*int32)(unsafe.Add(mBase, _consts[1]))
			v69 = *(*int32)(unsafe.Add(mBase, uint32(v63+(l2^int32(-1))<<(uint(int32(2))%32))))
			v77 = v69
		} else {
			v71 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
				v84 = *(*int32)(unsafe.Add(mBase, _consts[1]))
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v84+(v79^int32(-1))<<(uint(int32(2))%32))))
				v98 = v90
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, _consts[2]))
				v98 = v92 + v79<<(uint(int32(13))%32) + int32(-8192)
			}
			if v79 < int32(0) {
				v102 = *(*int32)(unsafe.Add(mBase, _consts[3]))
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v102+(v79^int32(-1))<<(uint(int32(6))%32))+16))
				v117 = v108
			} else {
				v110 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
					v125 = *(*int32)(unsafe.Add(mBase, _consts[1]))
					v131 = *(*int32)(unsafe.Add(mBase, uint32(v125+(v120^int32(-1))<<(uint(int32(2))%32))))
					v139 = v131
				} else {
					v133 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
					v153 = F_CopyIndexTuple(m, v77+v149&int32(32767))
					mBase = m.M
					v154 = m.ExcPending
					if v154 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v153))) = base.I32_rotr(v59, int32(16))
						v158 = int32(4474964)
						v160 = *(*int32)(unsafe.Add(mBase, _consts[26]))
						*(*int32)(unsafe.Add(mBase, _consts[26])) = v160 + int32(1)
						v164 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
						if base.Ui32(v164) <= base.Ui32(int32(2)) {
							v167 = int32(72)
							*(*uint16)(unsafe.Add(mBase, uint32(v139)+12)) = uint16(v167)
							v171 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v139-int32(-64)))) = uint8(v171)
							*(*int64)(unsafe.Add(mBase, uint32(v139)+56)) = int64(-4616189618054758400)
							*(*int32)(unsafe.Add(mBase, uint32(v139)+48)) = v171
							*(*int32)(unsafe.Add(mBase, uint32(v139)+28)) = int32(3)
						} else {
						}
						v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+16)))
						v180 = v98 + v179
						v181 = int32(2)
						*(*uint16)(unsafe.Add(mBase, uint32(v180)+12)) = uint16(v181)
						*(*int64)(unsafe.Add(mBase, uint32(v180))) = int64(0)
						v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+16)))
						v187 = *(*int32)(unsafe.Add(mBase, uint32(v77+v185)+8))
						v188 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v180)+14)) = uint16(v188)
						v190 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v180)+8)) = v187 + v190
						*(*int32)(unsafe.Add(mBase, uint32(v139)+32)) = v117
						v194 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v139)+40)) = v117
						*(*int32)(unsafe.Add(mBase, uint32(v139)+36)) = v194
						v197 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v139)+44)) = v197
						v202 = F_PageAddItemExtended(m, v98, v141, int32(8), v190, v188)
						mBase = m.M
						v203 = m.ExcPending
						if v203 != 0 {
							return
						} else {
							if v202 == int32(0) {
								F_errstart_cold(m, int32(23), int32(0))
								mBase = m.M
								v484 = m.ExcPending
								if v484 != 0 {
									return
								} else {
									if l2 < int32(0) {
										v488 = *(*int32)(unsafe.Add(mBase, _consts[3]))
										v494 = *(*int32)(unsafe.Add(mBase, uint32(v488+(l2^int32(-1))<<(uint(int32(6))%32))+16))
										v503 = v494
									} else {
										v496 = *(*int32)(unsafe.Add(mBase, _consts[4]))
										v502 = *(*int32)(unsafe.Add(mBase, uint32(v496+l2<<(uint(int32(6))%32)+int32(-64))+16))
										v503 = v502
									}
									v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v20))) = v503
									*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v504 + int32(4)
									F_errmsg_internal(m, int32(675363), v20)
									mBase = m.M
									v511 = m.ExcPending
									if v511 != 0 {
										return
									} else {
										F_errfinish(m, int32(487991), int32(2534), int32(302924))
										mBase = m.M
										v516 = m.ExcPending
										if v516 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v210 = F_PageAddItemExtended(m, v98, v153, int32(base.Ui32(v149)>>(uint(int32(17))%32)), int32(2), int32(0))
								mBase = m.M
								v211 = m.ExcPending
								if v211 != 0 {
									return
								} else {
									if v210 == int32(0) {
										F_errstart_cold(m, int32(23), int32(0))
										mBase = m.M
										v520 = m.ExcPending
										if v520 != 0 {
											return
										} else {
											if l2 < int32(0) {
												v524 = *(*int32)(unsafe.Add(mBase, _consts[3]))
												v530 = *(*int32)(unsafe.Add(mBase, uint32(v524+(l2^int32(-1))<<(uint(int32(6))%32))+16))
												v539 = v530
											} else {
												v532 = *(*int32)(unsafe.Add(mBase, _consts[4]))
												v538 = *(*int32)(unsafe.Add(mBase, uint32(v532+l2<<(uint(int32(6))%32)+int32(-64))+16))
												v539 = v538
											}
											v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v539
											*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v540 + int32(4)
											F_errmsg_internal(m, int32(675284), v20+int32(16))
											mBase = m.M
											v549 = m.ExcPending
											if v549 != 0 {
												return
											} else {
												F_errfinish(m, int32(487991), int32(2546), int32(302924))
												mBase = m.M
												v554 = m.ExcPending
												if v554 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v214 = v77 + v78
										v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+12)))
										v217 = v215 & int32(65407)
										*(*uint16)(unsafe.Add(mBase, uint32(v214)+12)) = uint16(v217)
										F_MarkBufferDirty(m, l2)
										mBase = m.M
										v220 = m.ExcPending
										if v220 != 0 {
											return
										} else {
											F_MarkBufferDirty(m, v79)
											mBase = m.M
											v222 = m.ExcPending
											if v222 != 0 {
												return
											} else {
												F_MarkBufferDirty(m, v120)
												mBase = m.M
												v224 = m.ExcPending
												if v224 != 0 {
													return
												} else {
													v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+118)))
													if v226 != int32(112) {
														v304 = int32(4474964)
														v306 = *(*int32)(unsafe.Add(mBase, _consts[26]))
														*(*int32)(unsafe.Add(mBase, _consts[26])) = v306 - int32(1)
														F__bt_relbuf(m, v120)
														mBase = m.M
														v311 = m.ExcPending
														if v311 != 0 {
															return
														} else {
															F_pfree(m, v141)
															mBase = m.M
															v313 = m.ExcPending
															if v313 != 0 {
																return
															} else {
																F_pfree(m, v153)
																mBase = m.M
																v315 = m.ExcPending
																if v315 != 0 {
																	return
																} else {
																	F__bt_relbuf(m, v79)
																	mBase = m.M
																	v317 = m.ExcPending
																	if v317 != 0 {
																		return
																	} else {
																		F__bt_relbuf(m, l3)
																		mBase = m.M
																		v319 = m.ExcPending
																		if v319 != 0 {
																			return
																		} else {
																			F__bt_relbuf(m, l2)
																			mBase = m.M
																			v321 = m.ExcPending
																			if v321 != 0 {
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
														v230 = *(*int32)(unsafe.Add(mBase, _consts[27]))
														if v230 <= int32(0) {
															v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															if v233 != 0 {
																v304 = int32(4474964)
																v306 = *(*int32)(unsafe.Add(mBase, _consts[26]))
																*(*int32)(unsafe.Add(mBase, _consts[26])) = v306 - int32(1)
																F__bt_relbuf(m, v120)
																mBase = m.M
																v311 = m.ExcPending
																if v311 != 0 {
																	return
																} else {
																	F_pfree(m, v141)
																	mBase = m.M
																	v313 = m.ExcPending
																	if v313 != 0 {
																		return
																	} else {
																		F_pfree(m, v153)
																		mBase = m.M
																		v315 = m.ExcPending
																		if v315 != 0 {
																			return
																		} else {
																			F__bt_relbuf(m, v79)
																			mBase = m.M
																			v317 = m.ExcPending
																			if v317 != 0 {
																				return
																			} else {
																				F__bt_relbuf(m, l3)
																				mBase = m.M
																				v319 = m.ExcPending
																				if v319 != 0 {
																					return
																				} else {
																					F__bt_relbuf(m, l2)
																					mBase = m.M
																					v321 = m.ExcPending
																					if v321 != 0 {
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
																v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																if v234 != 0 {
																	v304 = int32(4474964)
																	v306 = *(*int32)(unsafe.Add(mBase, _consts[26]))
																	*(*int32)(unsafe.Add(mBase, _consts[26])) = v306 - int32(1)
																	F__bt_relbuf(m, v120)
																	mBase = m.M
																	v311 = m.ExcPending
																	if v311 != 0 {
																		return
																	} else {
																		F_pfree(m, v141)
																		mBase = m.M
																		v313 = m.ExcPending
																		if v313 != 0 {
																			return
																		} else {
																			F_pfree(m, v153)
																			mBase = m.M
																			v315 = m.ExcPending
																			if v315 != 0 {
																				return
																			} else {
																				F__bt_relbuf(m, v79)
																				mBase = m.M
																				v317 = m.ExcPending
																				if v317 != 0 {
																					return
																				} else {
																					F__bt_relbuf(m, l3)
																					mBase = m.M
																					v319 = m.ExcPending
																					if v319 != 0 {
																						return
																					} else {
																						F__bt_relbuf(m, l2)
																						mBase = m.M
																						v321 = m.ExcPending
																						if v321 != 0 {
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
																	v237 = v139 + int32(36)
																	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
																	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v238
																	F_XLogBeginInsert(m)
																	mBase = m.M
																	v241 = m.ExcPending
																	if v241 != 0 {
																		return
																	} else {
																		F_XLogRegisterData(m, v20+int32(72), int32(8))
																		mBase = m.M
																		v246 = m.ExcPending
																		if v246 != 0 {
																			return
																		} else {
																			F_XLogRegisterBuffer(m, int32(0), v79, int32(6))
																			mBase = m.M
																			v250 = m.ExcPending
																			if v250 != 0 {
																				return
																			} else {
																				F_XLogRegisterBuffer(m, int32(1), l2, int32(8))
																				mBase = m.M
																				v254 = m.ExcPending
																				if v254 != 0 {
																					return
																				} else {
																					F_XLogRegisterBuffer(m, int32(2), v120, int32(14))
																					mBase = m.M
																					v258 = m.ExcPending
																					if v258 != 0 {
																						return
																					} else {
																						v259 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
																						*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v117
																						*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v259
																						v262 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
																						*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v262
																						*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v117
																						*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v262
																						v266 = *(*int32)(unsafe.Add(mBase, uint32(v139)+48))
																						*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v266
																						v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139-int32(-64)))))
																						*(*uint8)(unsafe.Add(mBase, uint32(v20)+68)) = uint8(v270)
																						F_XLogRegisterBufData(m, int32(2), v20+int32(44), int32(28))
																						mBase = m.M
																						v277 = m.ExcPending
																						if v277 != 0 {
																							return
																						} else {
																							v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+14)))
																							v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+16)))
																							F_XLogRegisterBufData(m, int32(0), v98+v279, v281-v279)
																							mBase = m.M
																							v284 = m.ExcPending
																							if v284 != 0 {
																								return
																							} else {
																								v287 = F_XLogInsert(m, int32(11), int32(160))
																								mBase = m.M
																								v288 = m.ExcPending
																								if v288 != 0 {
																									return
																								} else {
																									v289 = int64(32)
																									*(*int64)(unsafe.Add(mBase, uint32(v77))) = base.I64_rotr(v287, v289)
																									v292 = base.I32_wrap_i64(v287)
																									*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v292
																									v296 = base.I32_wrap_i64(int64(base.Ui64(v287) >> (uint(v289) % 64)))
																									*(*int32)(unsafe.Add(mBase, uint32(v98))) = v296
																									*(*int32)(unsafe.Add(mBase, uint32(v139)+4)) = v292
																									*(*int32)(unsafe.Add(mBase, uint32(v139))) = v296
																									v304 = int32(4474964)
																									v306 = *(*int32)(unsafe.Add(mBase, _consts[26]))
																									*(*int32)(unsafe.Add(mBase, _consts[26])) = v306 - int32(1)
																									F__bt_relbuf(m, v120)
																									mBase = m.M
																									v311 = m.ExcPending
																									if v311 != 0 {
																										return
																									} else {
																										F_pfree(m, v141)
																										mBase = m.M
																										v313 = m.ExcPending
																										if v313 != 0 {
																											return
																										} else {
																											F_pfree(m, v153)
																											mBase = m.M
																											v315 = m.ExcPending
																											if v315 != 0 {
																												return
																											} else {
																												F__bt_relbuf(m, v79)
																												mBase = m.M
																												v317 = m.ExcPending
																												if v317 != 0 {
																													return
																												} else {
																													F__bt_relbuf(m, l3)
																													mBase = m.M
																													v319 = m.ExcPending
																													if v319 != 0 {
																														return
																													} else {
																														F__bt_relbuf(m, l2)
																														mBase = m.M
																														v321 = m.ExcPending
																														if v321 != 0 {
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
															v237 = v139 + int32(36)
															v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
															*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v238
															F_XLogBeginInsert(m)
															mBase = m.M
															v241 = m.ExcPending
															if v241 != 0 {
																return
															} else {
																F_XLogRegisterData(m, v20+int32(72), int32(8))
																mBase = m.M
																v246 = m.ExcPending
																if v246 != 0 {
																	return
																} else {
																	F_XLogRegisterBuffer(m, int32(0), v79, int32(6))
																	mBase = m.M
																	v250 = m.ExcPending
																	if v250 != 0 {
																		return
																	} else {
																		F_XLogRegisterBuffer(m, int32(1), l2, int32(8))
																		mBase = m.M
																		v254 = m.ExcPending
																		if v254 != 0 {
																			return
																		} else {
																			F_XLogRegisterBuffer(m, int32(2), v120, int32(14))
																			mBase = m.M
																			v258 = m.ExcPending
																			if v258 != 0 {
																				return
																			} else {
																				v259 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
																				*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v117
																				*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v259
																				v262 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
																				*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v262
																				*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v117
																				*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v262
																				v266 = *(*int32)(unsafe.Add(mBase, uint32(v139)+48))
																				*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v266
																				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139-int32(-64)))))
																				*(*uint8)(unsafe.Add(mBase, uint32(v20)+68)) = uint8(v270)
																				F_XLogRegisterBufData(m, int32(2), v20+int32(44), int32(28))
																				mBase = m.M
																				v277 = m.ExcPending
																				if v277 != 0 {
																					return
																				} else {
																					v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+14)))
																					v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+16)))
																					F_XLogRegisterBufData(m, int32(0), v98+v279, v281-v279)
																					mBase = m.M
																					v284 = m.ExcPending
																					if v284 != 0 {
																						return
																					} else {
																						v287 = F_XLogInsert(m, int32(11), int32(160))
																						mBase = m.M
																						v288 = m.ExcPending
																						if v288 != 0 {
																							return
																						} else {
																							v289 = int64(32)
																							*(*int64)(unsafe.Add(mBase, uint32(v77))) = base.I64_rotr(v287, v289)
																							v292 = base.I32_wrap_i64(v287)
																							*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v292
																							v296 = base.I32_wrap_i64(int64(base.Ui64(v287) >> (uint(v289) % 64)))
																							*(*int32)(unsafe.Add(mBase, uint32(v98))) = v296
																							*(*int32)(unsafe.Add(mBase, uint32(v139)+4)) = v292
																							*(*int32)(unsafe.Add(mBase, uint32(v139))) = v296
																							v304 = int32(4474964)
																							v306 = *(*int32)(unsafe.Add(mBase, _consts[26]))
																							*(*int32)(unsafe.Add(mBase, _consts[26])) = v306 - int32(1)
																							F__bt_relbuf(m, v120)
																							mBase = m.M
																							v311 = m.ExcPending
																							if v311 != 0 {
																								return
																							} else {
																								F_pfree(m, v141)
																								mBase = m.M
																								v313 = m.ExcPending
																								if v313 != 0 {
																									return
																								} else {
																									F_pfree(m, v153)
																									mBase = m.M
																									v315 = m.ExcPending
																									if v315 != 0 {
																										return
																									} else {
																										F__bt_relbuf(m, v79)
																										mBase = m.M
																										v317 = m.ExcPending
																										if v317 != 0 {
																											return
																										} else {
																											F__bt_relbuf(m, l3)
																											mBase = m.M
																											v319 = m.ExcPending
																											if v319 != 0 {
																												return
																											} else {
																												F__bt_relbuf(m, l2)
																												mBase = m.M
																												v321 = m.ExcPending
																												if v321 != 0 {
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
			v325 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v331 = *(*int32)(unsafe.Add(mBase, uint32(v325+(l2^int32(-1))<<(uint(int32(6))%32))+16))
			v340 = v331
		} else {
			v333 = *(*int32)(unsafe.Add(mBase, _consts[4]))
			v339 = *(*int32)(unsafe.Add(mBase, uint32(v333+l2<<(uint(int32(6))%32)+int32(-64))+16))
			v340 = v339
		}
		if l3 < int32(0) {
			v344 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v350 = *(*int32)(unsafe.Add(mBase, uint32(v344+(l3^int32(-1))<<(uint(int32(6))%32))+16))
			v359 = v350
		} else {
			v352 = *(*int32)(unsafe.Add(mBase, _consts[4]))
			v358 = *(*int32)(unsafe.Add(mBase, uint32(v352+l3<<(uint(int32(6))%32)+int32(-64))+16))
			v359 = v358
		}
		if l2 < int32(0) {
			v363 = *(*int32)(unsafe.Add(mBase, _consts[1]))
			v369 = *(*int32)(unsafe.Add(mBase, uint32(v363+(l2^int32(-1))<<(uint(int32(2))%32))))
			v377 = v369
		} else {
			v371 = *(*int32)(unsafe.Add(mBase, _consts[2]))
			v377 = v371 + l2<<(uint(int32(13))%32) + int32(-8192)
		}
		if l4 == int32(0) {
			v382 = F_errstart(m, int32(13), int32(0))
			mBase = m.M
			v383 = m.ExcPending
			if v383 != 0 {
				return
			} else {
				if v382 != 0 {
					F_errmsg_internal(m, int32(101740), int32(0))
					mBase = m.M
					v387 = m.ExcPending
					if v387 != 0 {
						return
					} else {
						F_errfinish(m, int32(487991), int32(2149), int32(93053))
						mBase = m.M
						v392 = m.ExcPending
						if v392 != 0 {
							return
						} else {
							v393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v377)+16)))
							v395 = *(*int32)(unsafe.Add(mBase, uint32(v377+v393)+8))
							v399 = F__bt_get_endpoint(m, l0, v395+int32(1), int32(0))
							mBase = m.M
							v400 = m.ExcPending
							if v400 != 0 {
								return
							} else {
								if v399 < int32(0) {
									v404 = *(*int32)(unsafe.Add(mBase, _consts[3]))
									v410 = *(*int32)(unsafe.Add(mBase, uint32(v404+(v399^int32(-1))<<(uint(int32(6))%32))+16))
									v419 = v410
								} else {
									v412 = *(*int32)(unsafe.Add(mBase, _consts[4]))
									v418 = *(*int32)(unsafe.Add(mBase, uint32(v412+v399<<(uint(int32(6))%32)+int32(-64))+16))
									v419 = v418
								}
								v420 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v420
								*(*uint16)(unsafe.Add(mBase, uint32(v20)+48)) = uint16(v420)
								*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v419
								F__bt_relbuf(m, v399)
								mBase = m.M
								v426 = m.ExcPending
								if v426 != 0 {
									return
								} else {
									v429 = v20 + int32(44)
									v431 = *(*int32)(unsafe.Add(mBase, uint32(v377)+24))
									v435 = F_CopyIndexTuple(m, v377+v431&int32(32767))
									mBase = m.M
									v436 = m.ExcPending
									if v436 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v435))) = base.I32_rotr(v359, int32(16))
										v440 = F__bt_getstackbuf(m, l0, l1, v429, v340)
										mBase = m.M
										v441 = m.ExcPending
										if v441 != 0 {
											return
										} else {
											F__bt_relbuf(m, l3)
											mBase = m.M
											v443 = m.ExcPending
											if v443 != 0 {
												return
											} else {
												if v440 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v558 = m.ExcPending
													if v558 != 0 {
														return
													} else {
														F_errcode(m, int32(33557032))
														mBase = m.M
														v561 = m.ExcPending
														if v561 != 0 {
															return
														} else {
															v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v359
															*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v340
															*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v562 + int32(4)
															F_errmsg_internal(m, int32(38659), v20+int32(32))
															mBase = m.M
															v572 = m.ExcPending
															if v572 != 0 {
																return
															} else {
																F_errfinish(m, int32(487991), int32(2215), int32(93053))
																mBase = m.M
																v577 = m.ExcPending
																if v577 != 0 {
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
													v446 = int32(0)
													v447 = *(*int32)(unsafe.Add(mBase, uint32(v429)+8))
													v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v435)+6)))
													v455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v429)+4)))
													F__bt_insertonpg(m, l0, l1, v446, v440, l2, v447, v435, (v448&int32(8191)+int32(7))&int32(16376), (v455+int32(1))&int32(65535), v446, l6)
													mBase = m.M
													v462 = m.ExcPending
													if v462 != 0 {
														return
													} else {
														F_pfree(m, v435)
														mBase = m.M
														v464 = m.ExcPending
														if v464 != 0 {
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
					v393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v377)+16)))
					v395 = *(*int32)(unsafe.Add(mBase, uint32(v377+v393)+8))
					v399 = F__bt_get_endpoint(m, l0, v395+int32(1), int32(0))
					mBase = m.M
					v400 = m.ExcPending
					if v400 != 0 {
						return
					} else {
						if v399 < int32(0) {
							v404 = *(*int32)(unsafe.Add(mBase, _consts[3]))
							v410 = *(*int32)(unsafe.Add(mBase, uint32(v404+(v399^int32(-1))<<(uint(int32(6))%32))+16))
							v419 = v410
						} else {
							v412 = *(*int32)(unsafe.Add(mBase, _consts[4]))
							v418 = *(*int32)(unsafe.Add(mBase, uint32(v412+v399<<(uint(int32(6))%32)+int32(-64))+16))
							v419 = v418
						}
						v420 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v420
						*(*uint16)(unsafe.Add(mBase, uint32(v20)+48)) = uint16(v420)
						*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v419
						F__bt_relbuf(m, v399)
						mBase = m.M
						v426 = m.ExcPending
						if v426 != 0 {
							return
						} else {
							v429 = v20 + int32(44)
							v431 = *(*int32)(unsafe.Add(mBase, uint32(v377)+24))
							v435 = F_CopyIndexTuple(m, v377+v431&int32(32767))
							mBase = m.M
							v436 = m.ExcPending
							if v436 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v435))) = base.I32_rotr(v359, int32(16))
								v440 = F__bt_getstackbuf(m, l0, l1, v429, v340)
								mBase = m.M
								v441 = m.ExcPending
								if v441 != 0 {
									return
								} else {
									F__bt_relbuf(m, l3)
									mBase = m.M
									v443 = m.ExcPending
									if v443 != 0 {
										return
									} else {
										if v440 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v558 = m.ExcPending
											if v558 != 0 {
												return
											} else {
												F_errcode(m, int32(33557032))
												mBase = m.M
												v561 = m.ExcPending
												if v561 != 0 {
													return
												} else {
													v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v359
													*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v340
													*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v562 + int32(4)
													F_errmsg_internal(m, int32(38659), v20+int32(32))
													mBase = m.M
													v572 = m.ExcPending
													if v572 != 0 {
														return
													} else {
														F_errfinish(m, int32(487991), int32(2215), int32(93053))
														mBase = m.M
														v577 = m.ExcPending
														if v577 != 0 {
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
											v446 = int32(0)
											v447 = *(*int32)(unsafe.Add(mBase, uint32(v429)+8))
											v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v435)+6)))
											v455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v429)+4)))
											F__bt_insertonpg(m, l0, l1, v446, v440, l2, v447, v435, (v448&int32(8191)+int32(7))&int32(16376), (v455+int32(1))&int32(65535), v446, l6)
											mBase = m.M
											v462 = m.ExcPending
											if v462 != 0 {
												return
											} else {
												F_pfree(m, v435)
												mBase = m.M
												v464 = m.ExcPending
												if v464 != 0 {
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
			v429 = l4
			v431 = *(*int32)(unsafe.Add(mBase, uint32(v377)+24))
			v435 = F_CopyIndexTuple(m, v377+v431&int32(32767))
			mBase = m.M
			v436 = m.ExcPending
			if v436 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v435))) = base.I32_rotr(v359, int32(16))
				v440 = F__bt_getstackbuf(m, l0, l1, v429, v340)
				mBase = m.M
				v441 = m.ExcPending
				if v441 != 0 {
					return
				} else {
					F__bt_relbuf(m, l3)
					mBase = m.M
					v443 = m.ExcPending
					if v443 != 0 {
						return
					} else {
						if v440 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v558 = m.ExcPending
							if v558 != 0 {
								return
							} else {
								F_errcode(m, int32(33557032))
								mBase = m.M
								v561 = m.ExcPending
								if v561 != 0 {
									return
								} else {
									v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v359
									*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v340
									*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v562 + int32(4)
									F_errmsg_internal(m, int32(38659), v20+int32(32))
									mBase = m.M
									v572 = m.ExcPending
									if v572 != 0 {
										return
									} else {
										F_errfinish(m, int32(487991), int32(2215), int32(93053))
										mBase = m.M
										v577 = m.ExcPending
										if v577 != 0 {
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
							v446 = int32(0)
							v447 = *(*int32)(unsafe.Add(mBase, uint32(v429)+8))
							v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v435)+6)))
							v455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v429)+4)))
							F__bt_insertonpg(m, l0, l1, v446, v440, l2, v447, v435, (v448&int32(8191)+int32(7))&int32(16376), (v455+int32(1))&int32(65535), v446, l6)
							mBase = m.M
							v462 = m.ExcPending
							if v462 != 0 {
								return
							} else {
								F_pfree(m, v435)
								mBase = m.M
								v464 = m.ExcPending
								if v464 != 0 {
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
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 float64
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
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
					v43 = *(*int32)(unsafe.Add(mBase, _consts[135]))
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
										v90 = *(*int32)(unsafe.Add(mBase, _consts[55]))
										if v90 == int32(0) {
										} else {
											v94 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
											if v94 != int32(1) {
											} else {
												v97 = int32(4474964)
												v99 = *(*int32)(unsafe.Add(mBase, _consts[26]))
												v100 = int32(1)
												*(*int32)(unsafe.Add(mBase, _consts[26])) = v99 + v100
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
												*(*int32)(unsafe.Add(mBase, uint32(v90))) = v103 + v100
												*(*int64)(unsafe.Add(mBase, uint32(v90+int32(80))+232)) = int64(3)
												v111 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
												*(*int32)(unsafe.Add(mBase, uint32(v90))) = v111 + v100
												v117 = *(*int32)(unsafe.Add(mBase, _consts[26]))
												*(*int32)(unsafe.Add(mBase, _consts[26])) = v117 - v100
											}
										}
										v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										F_tuplesort_performsort(m, v121)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return
										} else {
											if l1 == int32(0) {
												v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
												*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
												if v169 != 0 {
													F_s_lock(m, l2+int32(36), int32(487880), int32(1951), int32(77954))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return
													} else {
														v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
														v180 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
														v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
														v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
														if v186 == v180 {
															v189 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
														} else {
														}
														v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
														v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
														v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
														if v195 == int32(1) {
															v198 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
														} else {
														}
														*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
														F_ConditionVariableSignal(m, l2+int32(24))
														mBase = m.M
														v205 = m.ExcPending
														if v205 != 0 {
															return
														} else {
															v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_tuplesort_end(m, v206)
															mBase = m.M
															v208 = m.ExcPending
															if v208 != 0 {
																return
															} else {
																if l1 != 0 {
																	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																	F_tuplesort_end(m, v209)
																	mBase = m.M
																	v211 = m.ExcPending
																	if v211 != 0 {
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
													v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
													v180 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
													v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
													v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
													if v186 == v180 {
														v189 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
													} else {
													}
													v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
													v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
													v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
													if v195 == int32(1) {
														v198 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
													} else {
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
													F_ConditionVariableSignal(m, l2+int32(24))
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return
													} else {
														v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_tuplesort_end(m, v206)
														mBase = m.M
														v208 = m.ExcPending
														if v208 != 0 {
															return
														} else {
															if l1 != 0 {
																v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																F_tuplesort_end(m, v209)
																mBase = m.M
																v211 = m.ExcPending
																if v211 != 0 {
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
												v130 = *(*int32)(unsafe.Add(mBase, _consts[55]))
												if v130 == int32(0) {
												} else {
													v134 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
													if v134 != int32(1) {
													} else {
														v137 = int32(4474964)
														v139 = *(*int32)(unsafe.Add(mBase, _consts[26]))
														v140 = int32(1)
														*(*int32)(unsafe.Add(mBase, _consts[26])) = v139 + v140
														v143 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
														*(*int32)(unsafe.Add(mBase, uint32(v130))) = v143 + v140
														*(*int64)(unsafe.Add(mBase, uint32(v130+int32(80))+232)) = int64(4)
														v151 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
														*(*int32)(unsafe.Add(mBase, uint32(v130))) = v151 + v140
														v157 = *(*int32)(unsafe.Add(mBase, _consts[26]))
														*(*int32)(unsafe.Add(mBase, _consts[26])) = v157 - v140
													}
												}
												v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												F_tuplesort_performsort(m, v166)
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return
												} else {
													v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
													*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
													if v169 != 0 {
														F_s_lock(m, l2+int32(36), int32(487880), int32(1951), int32(77954))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return
														} else {
															v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
															v180 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
															v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
															*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
															v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
															if v186 == v180 {
																v189 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
															} else {
															}
															v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
															v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
															*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
															v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
															if v195 == int32(1) {
																v198 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
															} else {
															}
															*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
															F_ConditionVariableSignal(m, l2+int32(24))
															mBase = m.M
															v205 = m.ExcPending
															if v205 != 0 {
																return
															} else {
																v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																F_tuplesort_end(m, v206)
																mBase = m.M
																v208 = m.ExcPending
																if v208 != 0 {
																	return
																} else {
																	if l1 != 0 {
																		v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																		F_tuplesort_end(m, v209)
																		mBase = m.M
																		v211 = m.ExcPending
																		if v211 != 0 {
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
														v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
														v180 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
														v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
														v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
														if v186 == v180 {
															v189 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
														} else {
														}
														v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
														v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
														v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
														if v195 == int32(1) {
															v198 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
														} else {
														}
														*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
														F_ConditionVariableSignal(m, l2+int32(24))
														mBase = m.M
														v205 = m.ExcPending
														if v205 != 0 {
															return
														} else {
															v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_tuplesort_end(m, v206)
															mBase = m.M
															v208 = m.ExcPending
															if v208 != 0 {
																return
															} else {
																if l1 != 0 {
																	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																	F_tuplesort_end(m, v209)
																	mBase = m.M
																	v211 = m.ExcPending
																	if v211 != 0 {
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
										v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										F_tuplesort_performsort(m, v161)
										mBase = m.M
										v163 = m.ExcPending
										if v163 != 0 {
											return
										} else {
											if l1 == int32(0) {
												v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
												*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
												if v169 != 0 {
													F_s_lock(m, l2+int32(36), int32(487880), int32(1951), int32(77954))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return
													} else {
														v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
														v180 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
														v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
														v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
														if v186 == v180 {
															v189 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
														} else {
														}
														v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
														v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
														v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
														if v195 == int32(1) {
															v198 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
														} else {
														}
														*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
														F_ConditionVariableSignal(m, l2+int32(24))
														mBase = m.M
														v205 = m.ExcPending
														if v205 != 0 {
															return
														} else {
															v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_tuplesort_end(m, v206)
															mBase = m.M
															v208 = m.ExcPending
															if v208 != 0 {
																return
															} else {
																if l1 != 0 {
																	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																	F_tuplesort_end(m, v209)
																	mBase = m.M
																	v211 = m.ExcPending
																	if v211 != 0 {
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
													v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
													v180 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
													v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
													v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
													if v186 == v180 {
														v189 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
													} else {
													}
													v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
													v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
													v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
													if v195 == int32(1) {
														v198 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
													} else {
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
													F_ConditionVariableSignal(m, l2+int32(24))
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return
													} else {
														v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_tuplesort_end(m, v206)
														mBase = m.M
														v208 = m.ExcPending
														if v208 != 0 {
															return
														} else {
															if l1 != 0 {
																v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																F_tuplesort_end(m, v209)
																mBase = m.M
																v211 = m.ExcPending
																if v211 != 0 {
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
												v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												F_tuplesort_performsort(m, v166)
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return
												} else {
													v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
													*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
													if v169 != 0 {
														F_s_lock(m, l2+int32(36), int32(487880), int32(1951), int32(77954))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return
														} else {
															v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
															v180 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
															v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
															*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
															v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
															if v186 == v180 {
																v189 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
															} else {
															}
															v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
															v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
															*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
															v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
															if v195 == int32(1) {
																v198 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
															} else {
															}
															*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
															F_ConditionVariableSignal(m, l2+int32(24))
															mBase = m.M
															v205 = m.ExcPending
															if v205 != 0 {
																return
															} else {
																v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																F_tuplesort_end(m, v206)
																mBase = m.M
																v208 = m.ExcPending
																if v208 != 0 {
																	return
																} else {
																	if l1 != 0 {
																		v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																		F_tuplesort_end(m, v209)
																		mBase = m.M
																		v211 = m.ExcPending
																		if v211 != 0 {
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
														v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
														v180 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
														v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
														v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
														if v186 == v180 {
															v189 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
														} else {
														}
														v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
														v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
														*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
														v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
														if v195 == int32(1) {
															v198 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
														} else {
														}
														*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
														F_ConditionVariableSignal(m, l2+int32(24))
														mBase = m.M
														v205 = m.ExcPending
														if v205 != 0 {
															return
														} else {
															v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_tuplesort_end(m, v206)
															mBase = m.M
															v208 = m.ExcPending
															if v208 != 0 {
																return
															} else {
																if l1 != 0 {
																	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																	F_tuplesort_end(m, v209)
																	mBase = m.M
																	v211 = m.ExcPending
																	if v211 != 0 {
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
								v90 = *(*int32)(unsafe.Add(mBase, _consts[55]))
								if v90 == int32(0) {
								} else {
									v94 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
									if v94 != int32(1) {
									} else {
										v97 = int32(4474964)
										v99 = *(*int32)(unsafe.Add(mBase, _consts[26]))
										v100 = int32(1)
										*(*int32)(unsafe.Add(mBase, _consts[26])) = v99 + v100
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
										*(*int32)(unsafe.Add(mBase, uint32(v90))) = v103 + v100
										*(*int64)(unsafe.Add(mBase, uint32(v90+int32(80))+232)) = int64(3)
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
										*(*int32)(unsafe.Add(mBase, uint32(v90))) = v111 + v100
										v117 = *(*int32)(unsafe.Add(mBase, _consts[26]))
										*(*int32)(unsafe.Add(mBase, _consts[26])) = v117 - v100
									}
								}
								v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_tuplesort_performsort(m, v121)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return
								} else {
									if l1 == int32(0) {
										v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
										*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
										if v169 != 0 {
											F_s_lock(m, l2+int32(36), int32(487880), int32(1951), int32(77954))
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return
											} else {
												v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
												v180 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
												v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
												v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
												if v186 == v180 {
													v189 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
												} else {
												}
												v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
												v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
												v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
												if v195 == int32(1) {
													v198 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
												} else {
												}
												*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
												F_ConditionVariableSignal(m, l2+int32(24))
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return
												} else {
													v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_tuplesort_end(m, v206)
													mBase = m.M
													v208 = m.ExcPending
													if v208 != 0 {
														return
													} else {
														if l1 != 0 {
															v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
															F_tuplesort_end(m, v209)
															mBase = m.M
															v211 = m.ExcPending
															if v211 != 0 {
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
											v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
											v180 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
											v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
											*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
											v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
											if v186 == v180 {
												v189 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
											} else {
											}
											v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
											v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
											*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
											v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
											if v195 == int32(1) {
												v198 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
											} else {
											}
											*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
											F_ConditionVariableSignal(m, l2+int32(24))
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return
											} else {
												v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_tuplesort_end(m, v206)
												mBase = m.M
												v208 = m.ExcPending
												if v208 != 0 {
													return
												} else {
													if l1 != 0 {
														v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
														F_tuplesort_end(m, v209)
														mBase = m.M
														v211 = m.ExcPending
														if v211 != 0 {
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
										v130 = *(*int32)(unsafe.Add(mBase, _consts[55]))
										if v130 == int32(0) {
										} else {
											v134 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
											if v134 != int32(1) {
											} else {
												v137 = int32(4474964)
												v139 = *(*int32)(unsafe.Add(mBase, _consts[26]))
												v140 = int32(1)
												*(*int32)(unsafe.Add(mBase, _consts[26])) = v139 + v140
												v143 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
												*(*int32)(unsafe.Add(mBase, uint32(v130))) = v143 + v140
												*(*int64)(unsafe.Add(mBase, uint32(v130+int32(80))+232)) = int64(4)
												v151 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
												*(*int32)(unsafe.Add(mBase, uint32(v130))) = v151 + v140
												v157 = *(*int32)(unsafe.Add(mBase, _consts[26]))
												*(*int32)(unsafe.Add(mBase, _consts[26])) = v157 - v140
											}
										}
										v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										F_tuplesort_performsort(m, v166)
										mBase = m.M
										v168 = m.ExcPending
										if v168 != 0 {
											return
										} else {
											v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
											*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
											if v169 != 0 {
												F_s_lock(m, l2+int32(36), int32(487880), int32(1951), int32(77954))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
													v180 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
													v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
													v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
													if v186 == v180 {
														v189 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
													} else {
													}
													v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
													v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
													v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
													if v195 == int32(1) {
														v198 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
													} else {
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
													F_ConditionVariableSignal(m, l2+int32(24))
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return
													} else {
														v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_tuplesort_end(m, v206)
														mBase = m.M
														v208 = m.ExcPending
														if v208 != 0 {
															return
														} else {
															if l1 != 0 {
																v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																F_tuplesort_end(m, v209)
																mBase = m.M
																v211 = m.ExcPending
																if v211 != 0 {
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
												v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
												v180 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
												v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
												v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
												if v186 == v180 {
													v189 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
												} else {
												}
												v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
												v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
												v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
												if v195 == int32(1) {
													v198 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
												} else {
												}
												*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
												F_ConditionVariableSignal(m, l2+int32(24))
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return
												} else {
													v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_tuplesort_end(m, v206)
													mBase = m.M
													v208 = m.ExcPending
													if v208 != 0 {
														return
													} else {
														if l1 != 0 {
															v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
															F_tuplesort_end(m, v209)
															mBase = m.M
															v211 = m.ExcPending
															if v211 != 0 {
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
								v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_tuplesort_performsort(m, v161)
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
									return
								} else {
									if l1 == int32(0) {
										v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
										*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
										if v169 != 0 {
											F_s_lock(m, l2+int32(36), int32(487880), int32(1951), int32(77954))
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return
											} else {
												v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
												v180 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
												v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
												v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
												if v186 == v180 {
													v189 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
												} else {
												}
												v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
												v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
												v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
												if v195 == int32(1) {
													v198 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
												} else {
												}
												*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
												F_ConditionVariableSignal(m, l2+int32(24))
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return
												} else {
													v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_tuplesort_end(m, v206)
													mBase = m.M
													v208 = m.ExcPending
													if v208 != 0 {
														return
													} else {
														if l1 != 0 {
															v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
															F_tuplesort_end(m, v209)
															mBase = m.M
															v211 = m.ExcPending
															if v211 != 0 {
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
											v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
											v180 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
											v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
											*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
											v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
											if v186 == v180 {
												v189 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
											} else {
											}
											v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
											v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
											*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
											v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
											if v195 == int32(1) {
												v198 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
											} else {
											}
											*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
											F_ConditionVariableSignal(m, l2+int32(24))
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return
											} else {
												v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_tuplesort_end(m, v206)
												mBase = m.M
												v208 = m.ExcPending
												if v208 != 0 {
													return
												} else {
													if l1 != 0 {
														v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
														F_tuplesort_end(m, v209)
														mBase = m.M
														v211 = m.ExcPending
														if v211 != 0 {
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
										v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										F_tuplesort_performsort(m, v166)
										mBase = m.M
										v168 = m.ExcPending
										if v168 != 0 {
											return
										} else {
											v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
											*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(1)
											if v169 != 0 {
												F_s_lock(m, l2+int32(36), int32(487880), int32(1951), int32(77954))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
													v180 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
													v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
													v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
													if v186 == v180 {
														v189 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
													} else {
													}
													v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
													v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
													*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
													v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
													if v195 == int32(1) {
														v198 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
													} else {
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
													F_ConditionVariableSignal(m, l2+int32(24))
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return
													} else {
														v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_tuplesort_end(m, v206)
														mBase = m.M
														v208 = m.ExcPending
														if v208 != 0 {
															return
														} else {
															if l1 != 0 {
																v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																F_tuplesort_end(m, v209)
																mBase = m.M
																v211 = m.ExcPending
																if v211 != 0 {
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
												v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
												v180 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v179 + v180
												v183 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v84, v183)
												v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
												if v186 == v180 {
													v189 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)) = uint8(v189)
												} else {
												}
												v191 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
												v192 = *(*float64)(unsafe.Add(mBase, uint32(l2)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l2)+64)) = base.F64_add(v191, v192)
												v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+122)))
												if v195 == int32(1) {
													v198 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)) = uint8(v198)
												} else {
												}
												*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = int32(0)
												F_ConditionVariableSignal(m, l2+int32(24))
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return
												} else {
													v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_tuplesort_end(m, v206)
													mBase = m.M
													v208 = m.ExcPending
													if v208 != 0 {
														return
													} else {
														if l1 != 0 {
															v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
															F_tuplesort_end(m, v209)
															mBase = m.M
															v211 = m.ExcPending
															if v211 != 0 {
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	v7 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+212))
	v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16+v17<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v7
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v23 == v26 {
		v55 = l3
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L9
	} else {
		goto L69
	}
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v251)
	m.G0 = v13 + int32(48)
	return v249
L3:
	;
	v56 = int32(0)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v62 = F__bt_binsrch_array_skey(m, v55, v56, v56, v58, v56, l4, l1, v13+int32(44))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L9
	} else {
		goto L15
	}
L4:
	;
	if v26 == int32(0) {
		v55 = l3
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+208))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+v17<<(uint(int32(2))%32)-int32(4))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v37 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v38 = v37
	goto L8
L7:
	;
	v38 = v23
	goto L8
L8:
	;
	v40 = F_get_opfamily_proc(m, v36, v26, v38, int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if v40 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v46 = int32(0)
	v249 = v46
	v251 = v46
	goto L2
L12:
	;
	goto L13
L13:
	;
	F_fmgr_info(m, v40, v13+int32(16))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v55 = v13 + int32(16)
	goto L3
L15:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	switch v64 - int32(1) {
	case 0:
		goto L21
	case 1:
		v68 = v7
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	case 4:
		v82 = v7
		goto L17
	default:
		goto L1
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v239
	v249 = int32(1)
	v251 = base.B2i32(int32(0) < v239)
	goto L2
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v86 = v62 + base.B2i32(v82 <= v84)
	v87 = int32(2)
	v89 = v83 + v86<<(uint(v87)%32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v91 = v90 - v86
	v93 = v91 << (uint(v87) % 32)
	if v83 == v89 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v82 = int32(1)
	goto L17
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v73 != 0 {
		v239 = int32(0)
		goto L16
	} else {
		goto L22
	}
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v239 = v62 + base.B2i32(v68 <= v69)
	goto L16
L21:
	;
	v68 = int32(1)
	goto L20
L22:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v62<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v78
	v239 = int32(1)
	goto L16
L23:
	;
	v239 = v91
	goto L16
L24:
	;
	goto L23
L25:
	;
	v97 = v83 + v93
	if base.Ui32(v89-v97) <= base.Ui32(int32(0)-v93<<(uint(int32(1))%32)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v104 = F___memcpy(m, v83, v89, v93)
	mBase = m.M
	goto L23
L27:
	;
	goto L28
L28:
	;
	v107 = (v83 ^ v89) & int32(3)
	if base.Ui32(v83) < base.Ui32(v89) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if v209 == int32(0) {
		goto L24
	} else {
		goto L65
	}
L30:
	;
	if base.Ui32(v187) <= base.Ui32(int32(3)) {
		v208 = v186
		v209 = v187
		v210 = v188
		goto L29
	} else {
		goto L61
	}
L31:
	;
	if v107 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v107 != 0 {
		v169 = v93
		goto L44
	} else {
		goto L45
	}
L34:
	;
	v208 = v89
	v209 = v93
	v210 = v83
	goto L29
L35:
	;
	goto L36
L36:
	;
	if v83&int32(3) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v186 = v89
	v187 = v93
	v188 = v83
	goto L30
L38:
	;
	goto L39
L39:
	;
	v114 = v89
	v115 = v93
	v116 = v83
	goto L40
L40:
	;
	if v115 == int32(0) {
		goto L24
	} else {
		goto L42
	}
L41:
	;
	v186 = v123
	v187 = v125
	v188 = v127
	goto L30
L42:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	*(*uint8)(unsafe.Add(mBase, uint32(v116))) = uint8(v120)
	v122 = int32(1)
	v123 = v114 + v122
	v125 = v115 - v122
	v127 = v116 + v122
	if v127&int32(3) != 0 {
		v114 = v123
		v115 = v125
		v116 = v127
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	if v169 == int32(0) {
		goto L24
	} else {
		goto L57
	}
L45:
	;
	if v97&int32(3) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v134 = v93
	goto L49
L47:
	;
	v149 = v93
	goto L48
L48:
	;
	if base.Ui32(v149) <= base.Ui32(int32(3)) {
		v169 = v149
		goto L44
	} else {
		goto L53
	}
L49:
	;
	if v134 == int32(0) {
		goto L24
	} else {
		goto L51
	}
L50:
	;
	v149 = v140
	goto L48
L51:
	;
	v140 = v134 - int32(1)
	v141 = v83 + v140
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+v140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v143)
	if v141&int32(3) != 0 {
		v134 = v140
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v156 = v149
	goto L54
L54:
	;
	v160 = v156 - int32(4)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v89+v160)))
	*(*int32)(unsafe.Add(mBase, uint32(v83+v160))) = v163
	if base.Ui32(int32(3)) < base.Ui32(v160) {
		v156 = v160
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v169 = v160
	goto L44
L56:
	;
	goto L55
L57:
	;
	v176 = v169
	goto L58
L58:
	;
	v180 = v176 - int32(1)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+v180))))
	*(*uint8)(unsafe.Add(mBase, uint32(v83+v180))) = uint8(v183)
	if v180 != 0 {
		v176 = v180
		goto L58
	} else {
		goto L60
	}
L59:
	;
	goto L24
L60:
	;
	goto L59
L61:
	;
	v193 = v186
	v194 = v187
	v195 = v188
	goto L62
L62:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v197
	v199 = int32(4)
	v200 = v193 + v199
	v202 = v195 + v199
	v204 = v194 - v199
	if base.Ui32(int32(3)) < base.Ui32(v204) {
		v193 = v200
		v194 = v204
		v195 = v202
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v208 = v200
	v209 = v204
	v210 = v202
	goto L29
L64:
	;
	goto L63
L65:
	;
	v215 = v208
	v216 = v209
	v217 = v210
	goto L66
L66:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	*(*uint8)(unsafe.Add(mBase, uint32(v217))) = uint8(v219)
	v221 = int32(1)
	v226 = v216 - v221
	if v226 != 0 {
		v215 = v215 + v221
		v216 = v226
		v217 = v217 + v221
		goto L66
	} else {
		goto L68
	}
L67:
	;
	goto L24
L68:
	;
	goto L67
L69:
	;
	v261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v261
	F_errmsg_internal(m, int32(477646), v13)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(488581), int32(1232), int32(312599))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
