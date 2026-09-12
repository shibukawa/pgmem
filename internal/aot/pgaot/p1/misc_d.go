package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_DecodeUnits(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	v9 = *(*int32)(unsafe.Add(mBase, _consts[1066]))
	if v9 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1066])) = v140
	v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v140)+11)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v148
	return v147
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	v63 = int32(1618592)
	v65 = int32(1619552)
	goto L20
L5:
	;
	if v46-v47 == int32(0) {
		v140 = v9
		goto L1
	} else {
		goto L19
	}
L7:
	;
	goto L8
L8:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v17 = l0
	v18 = v9
	v19 = int32(10)
	v20 = v16
	goto L13
L10:
	;
	v42 = v9
	v46 = int32(0)
	goto L11
L11:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	goto L5
L12:
	;
	v42 = v37
	v46 = v39
	goto L11
L13:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v20 != v22 {
		v37 = v18
		v39 = v20
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v37 = v31
	v39 = int32(0)
	goto L12
L15:
	;
	if v22 == int32(0) {
		v37 = v18
		v39 = v20
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v27 = v19 - int32(1)
	if v27 == int32(0) {
		v37 = v18
		v39 = v20
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v30 = int32(1)
	v31 = v18 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v32 != 0 {
		v17 = v17 + v30
		v18 = v31
		v19 = v27
		v20 = v32
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	goto L4
L20:
	;
	v72 = v63 + (v65-v63)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v73 = int32(*(*int8)(unsafe.Add(mBase, uint32(v72))))
	v74 = v57 - v73
	if v74 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return int32(31)
L22:
	;
	goto L27
L23:
	;
	v124 = v74
	goto L24
L24:
	;
	v128 = base.B2i32(v124 < int32(0))
	if v124 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L25:
	;
	if v115 == int32(0) {
		v140 = v72
		goto L1
	} else {
		goto L39
	}
L27:
	;
	goto L28
L28:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v83 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v84 = l0
	v85 = v72
	v86 = int32(10)
	v87 = v83
	goto L33
L30:
	;
	v109 = v72
	v113 = int32(0)
	goto L31
L31:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v115 = v113 - v114
	goto L25
L32:
	;
	v109 = v104
	v113 = v106
	goto L31
L33:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v87 != v89 {
		v104 = v85
		v106 = v87
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v104 = v98
	v106 = int32(0)
	goto L32
L35:
	;
	if v89 == int32(0) {
		v104 = v85
		v106 = v87
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v94 = v86 - int32(1)
	if v94 == int32(0) {
		v104 = v85
		v106 = v87
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v97 = int32(1)
	v98 = v85 + v97
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	if v99 != 0 {
		v84 = v84 + v97
		v85 = v98
		v86 = v94
		v87 = v99
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v124 = v115
	goto L24
L40:
	;
	v129 = v72 - int32(16)
	goto L42
L41:
	;
	v129 = v65
	goto L42
L42:
	;
	if v124 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v132 = v63
	goto L45
L44:
	;
	v132 = v72 + int32(16)
	goto L45
L45:
	;
	if base.Ui32(v132) <= base.Ui32(v129) {
		v63 = v132
		v65 = v129
		goto L20
	} else {
		goto L46
	}
L46:
	;
	goto L21
}
func F_DecodingContextFindStartpoint(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+104))
	F_XLogBeginRead(m, v10, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v11)+104))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+20)) = uint32(v19)
	v22 = int64(base.Ui64(v19) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+16)) = uint32(v22)
	F_errmsg_internal(m, int32(489216), v8+int32(16))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = F_XLogReadRecord(m, v37, v8+int32(28))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	F_errfinish(m, int32(474525), int32(643), int32(83632))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L36
	}
L11:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
	if v92 != 0 {
		goto L29
	} else {
		goto L30
	}
L12:
	;
	v48 = v40
	goto L15
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L26
	}
L15:
	;
	if v48 == int32(0) {
		goto L10
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_LogicalDecodingProcessRecord(m, l0, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v56 == int32(2) {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v60 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v68 = F_XLogReadRecord(m, v65, v8+int32(28))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	if v70 == int32(0) {
		v48 = v68
		goto L15
	} else {
		goto L25
	}
L25:
	;
	goto L16
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v82
	F_errmsg_internal(m, int32(189236), v8)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(474525), int32(654), int32(83632))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_s_lock(m, v11, int32(474525), int32(667), int32(83632))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v100)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+120)) = v101
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+136)))
	if v103 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v106)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+128)) = v107
	goto L35
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
	m.G0 = v8 + int32(32)
	return
L36:
	;
	F_errmsg_internal(m, int32(84410), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(474525), int32(656), int32(83632))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_DefineCustomEnumVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	v10 = F_init_custom_variable(m, l0, l1, l2, int32(5), int32(0), int32(4), int32(124))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+116)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = l4
		F_define_custom_variable(m, v10)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			return
		}
	}
}
func F_DetermineTimeZoneAbbrevOffset(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int64
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	v7 = m.G0
	v9 = v7 - int32(288)
	m.G0 = v9
	v13 = F_DetermineTimeZoneOffsetInternal(m, l0, l2, v9+int32(280))
	mBase = m.M
	v15 = v9 + int32(16)
	goto L4
L1:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
	if v131 != 0 {
		goto L33
	} else {
		goto L34
	}
L2:
	;
	v128 = F_strlen(m, v117)
	mBase = m.M
	goto L1
L4:
	;
	goto L5
L5:
	;
	v22 = int32(255)
	if (v15^l1)&int32(3) != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v121 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v118))) = uint8(v121)
	goto L2
L7:
	;
	v102 = v97
	v103 = v98
	v104 = v99
	goto L29
L8:
	;
	if v92 == int32(0) {
		v117 = v90
		v118 = v91
		goto L6
	} else {
		goto L28
	}
L9:
	;
	v90 = l1
	v91 = v15
	v92 = v22
	goto L8
L10:
	;
	goto L11
L11:
	;
	if l1&int32(3) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v59 == int32(0) {
		v117 = v56
		v118 = v57
		goto L6
	} else {
		goto L21
	}
L13:
	;
	v56 = l1
	v57 = v15
	v58 = v22
	v59 = int32(1)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v35 = l1
	v36 = v15
	v37 = v22
	goto L16
L16:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v39)
	if v39 == int32(0) {
		v97 = v35
		v98 = v36
		v99 = v37
		goto L7
	} else {
		goto L18
	}
L17:
	;
	v56 = v50
	v57 = v44
	v58 = v46
	v59 = v48
	goto L12
L18:
	;
	v43 = int32(1)
	v44 = v36 + v43
	v46 = v37 - v43
	v47 = int32(0)
	v48 = base.B2i32(v46 != v47)
	v50 = v35 + v43
	if v50&int32(3) == v47 {
		v56 = v50
		v57 = v44
		v58 = v46
		v59 = v48
		goto L12
	} else {
		goto L19
	}
L19:
	;
	if v46 != 0 {
		v35 = v50
		v36 = v44
		v37 = v46
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v62 == int32(0) {
		v90 = v56
		v91 = v57
		v92 = v58
		goto L8
	} else {
		goto L22
	}
L22:
	;
	if base.Ui32(v58) < base.Ui32(int32(4)) {
		v90 = v56
		v91 = v57
		v92 = v58
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v68 = v56
	v69 = v57
	v70 = v58
	goto L24
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v76 = int32(-2139062144)
	if (int32(16843008)-v73|v73)&v76 != v76 {
		v97 = v68
		v98 = v69
		v99 = v70
		goto L7
	} else {
		goto L26
	}
L25:
	;
	v90 = v84
	v91 = v82
	v92 = v86
	goto L8
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v73
	v81 = int32(4)
	v82 = v69 + v81
	v84 = v68 + v81
	v86 = v70 - v81
	if base.Ui32(int32(3)) < base.Ui32(v86) {
		v68 = v84
		v69 = v82
		v70 = v86
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v97 = v90
	v98 = v91
	v99 = v92
	goto L7
L29:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v106)
	if v106 == int32(0) {
		v117 = v102
		v118 = v103
		goto L6
	} else {
		goto L31
	}
L30:
	;
	v117 = v113
	v118 = v111
	goto L6
L31:
	;
	v110 = int32(1)
	v111 = v103 + v110
	v113 = v102 + v110
	v115 = v104 - v110
	if v115 != 0 {
		v102 = v113
		v103 = v111
		v104 = v115
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v135 = v9 + int32(16)
	v139 = v131
	goto L36
L34:
	;
	goto L35
L35:
	;
	v169 = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l2)+268))
	if v176 <= v169 {
		v339 = v169
		goto L44
	} else {
		goto L45
	}
L36:
	;
	if base.Ui32((v139-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L35
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v150)
	v153 = v135 + int32(1)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v154 != 0 {
		v135 = v153
		v139 = v154
		goto L36
	} else {
		goto L42
	}
L39:
	;
	v148 = v139 - int32(32)
	goto L41
L40:
	;
	v148 = v139
	goto L41
L41:
	;
	v150 = v148 & int32(255)
	goto L38
L42:
	;
	goto L37
L43:
	;
	if v339 != 0 {
		goto L81
	} else {
		goto L82
	}
L44:
	;
	goto L43
L45:
	;
	v180 = l2 + int32(22376)
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(280))))
	v189 = v169
	goto L46
L46:
	;
	v195 = F_strcmp(m, v9+int32(16), v180+v189)
	mBase = m.M
	if v195 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l2)+260))
	if v213 <= int32(0) {
		goto L56
	} else {
		goto L57
	}
L48:
	;
	v197 = v189
	goto L51
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	v210 = v197 + int32(1)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197+v180))))
	if v211 != 0 {
		v197 = v210
		goto L51
	} else {
		goto L53
	}
L52:
	;
	if v210 < v176 {
		v189 = v210
		goto L46
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v339 = v169
	goto L44
L55:
	;
	v258 = l2 + int32(16280)
	v260 = l2 + int32(18280)
	v267 = v246
	goto L69
L56:
	;
	v246 = int32(0)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v221 = int32(0)
	v226 = v213
	goto L59
L59:
	;
	v233 = int32(1)
	v234 = (v221 + v226) >> (uint(v233) % 32)
	v240 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(280)+v234<<(uint(int32(3))%32))))
	v241 = base.B2i32(v181 < v240)
	if v181 < v240 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v246 = v242
	goto L55
L61:
	;
	v242 = v221
	goto L63
L62:
	;
	v242 = v234 + v233
	goto L63
L63:
	;
	if v181 < v240 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v243 = v234
	goto L66
L65:
	;
	v243 = v226
	goto L66
L66:
	;
	if v242 < v243 {
		v221 = v242
		v226 = v243
		goto L59
	} else {
		goto L67
	}
L67:
	;
	goto L60
L68:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(12)))) = v325
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(8)))) = v327
	v339 = int32(1)
	goto L44
L69:
	;
	if int32(0) < v267 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l2)+uint32(_consts[1071])))
	v287 = v260 + v284<<(uint(int32(4))%32)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+8))
	if v288 == v189 {
		v318 = v287
		goto L68
	} else {
		goto L75
	}
L71:
	;
	v276 = v267 - int32(1)
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258+v276))))
	v281 = v260 + v278<<(uint(int32(4))%32)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+8))
	if v282 != v189 {
		v267 = v276
		goto L69
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	goto L70
L74:
	;
	v318 = v281
	goto L68
L75:
	;
	if v213 <= v246 {
		v339 = v169
		goto L44
	} else {
		goto L76
	}
L76:
	;
	v292 = v246
	goto L77
L77:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292+v258))))
	v307 = v260 + v304<<(uint(int32(4))%32)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+8))
	if v308 == v189 {
		v318 = v307
		goto L68
	} else {
		goto L79
	}
L78:
	;
	v339 = v169
	goto L44
L79:
	;
	v311 = v292 + int32(1)
	if v213 != v311 {
		v292 = v311
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v343
	v348 = int32(0) - v342
	goto L83
L82:
	;
	v348 = v13
	goto L83
L83:
	;
	m.G0 = v9 + int32(288)
	return v348
}
func F_Do_MultiXactIdWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	v10 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = int32(1)
	if l2&int32(4304) == int32(4224) {
		v256 = v18
		v257 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l7 != 0 {
		goto L70
	} else {
		goto L71
	}
L2:
	;
	v34 = F_GetMultiXactIdMembers(m, l0, v16+int32(12), int32(base.Ui32(l2&int32(128))>>(uint(int32(7))%32))|base.B2i32(l2&int32(4176) == int32(64)))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v34 < int32(0) {
		v256 = v18
		v257 = v10
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v34 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	F_pfree(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L3
	} else {
		goto L69
	}
L7:
	;
	v240 = v18
	v241 = v10
	goto L6
L8:
	;
	goto L9
L9:
	;
	v49 = int32(0)
	v57 = v10
	goto L10
L10:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v61 = int32(3)
	v63 = v60 + v49<<(uint(v61)%32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if base.Ui32(v65) < base.Ui32(v61) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v240 = v227
	v241 = v226
	goto L6
L12:
	;
	v227 = int32(1)
	v229 = v49 + v227
	if v229 != v34 {
		v49 = v229
		v57 = v226
		goto L10
	} else {
		goto L68
	}
L13:
	;
	if v185 != 0 {
		goto L53
	} else {
		goto L54
	}
L14:
	;
	v185 = int32(0)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v76 == v65 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v185 = int32(1)
	goto L13
L18:
	;
	goto L19
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v80 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v185 = v177
	goto L13
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v84 == int32(0) {
		v177 = int32(0)
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v148 = int32(0)
	v150 = v80 - int32(1)
	goto L43
L24:
	;
	v89 = v84
	goto L25
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v94 == int32(4) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v177 = int32(0)
	goto L20
L27:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v89)+80))
	if v141 != 0 {
		v89 = v141
		goto L25
	} else {
		goto L42
	}
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v97 == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v100 = int32(1)
	if v65 == v97 {
		v177 = v100
		goto L20
	} else {
		goto L30
	}
L30:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v89)+52))
	v104 = v102 - int32(1)
	if v104 < int32(0) {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v109 = int32(0)
	v111 = v104
	goto L32
L32:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v89)+48))
	v117 = int32(2)
	v118 = base.I32_div_s(v111-v109, v117)
	v119 = v118 + v109
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v115+v119<<(uint(v117)%32))))
	if v123 == v65 {
		v177 = v100
		goto L20
	} else {
		goto L34
	}
L33:
	;
	goto L27
L34:
	;
	v127 = F_TransactionIdPrecedes(m, v123, v65)
	mBase = m.M
	if v127 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v128 = v119 + int32(1)
	goto L37
L36:
	;
	v128 = v109
	goto L37
L37:
	;
	if v127 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v131 = v111
	goto L40
L39:
	;
	v131 = v119 - int32(1)
	goto L40
L40:
	;
	if v128 <= v131 {
		v109 = v128
		v111 = v131
		goto L32
	} else {
		goto L41
	}
L41:
	;
	goto L33
L42:
	;
	goto L26
L43:
	;
	v155 = int32(2)
	v156 = base.I32_div_s(v150-v148, v155)
	v157 = v156 + v148
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v146+v157<<(uint(v155)%32))))
	v162 = base.B2i32(v161 == v65)
	if v161 == v65 {
		v177 = v162
		goto L20
	} else {
		goto L45
	}
L44:
	;
	v177 = v162
	goto L20
L45:
	;
	v165 = base.B2i32(base.Ui32(v161) < base.Ui32(v65))
	if base.Ui32(v161) < base.Ui32(v65) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v166 = v157 + int32(1)
	goto L48
L47:
	;
	v166 = v148
	goto L48
L48:
	;
	if base.Ui32(v161) < base.Ui32(v65) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v169 = v150
	goto L51
L50:
	;
	v169 = v157 - int32(1)
	goto L51
L51:
	;
	if v166 <= v169 {
		v148 = v166
		v150 = v169
		goto L43
	} else {
		goto L52
	}
L52:
	;
	goto L44
L53:
	;
	v226 = v57 + int32(1)
	goto L12
L54:
	;
	goto L55
L55:
	;
	v188 = int32(2)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v64<<(uint(v188)%32))+uint32(_consts[72])))
	v193 = int32(12)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v192*v193)+uint32(_consts[73])))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32))+uint32(_consts[72])))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v198*v193)+uint32(_consts[73])))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v197<<(uint(v188)%32))+uint32(_consts[74])))
	goto L56
L56:
	;
	if int32(base.Ui32(v208)>>(uint(v203)%32))&int32(1) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if l7 == int32(0) {
		v226 = v57
		goto L12
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if l3 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v216 = F_TransactionIdIsInProgress(m, v65)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	v226 = v216 + v57
	goto L12
L62:
	;
	F_XactLockTableWait(m, v65, l4, l5, l6)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L3
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v223 = F_ConditionalXactLockTableWait(m, v65, l8)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L3
	} else {
		goto L66
	}
L65:
	;
	v226 = v57
	goto L12
L66:
	;
	if v223 != 0 {
		v226 = v57
		goto L12
	} else {
		goto L67
	}
L67:
	;
	v240 = int32(0)
	v241 = v57
	goto L6
L68:
	;
	goto L11
L69:
	;
	v256 = v240
	v257 = v241
	goto L1
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v257
	goto L72
L71:
	;
	goto L72
L72:
	;
	m.G0 = v16 + int32(16)
	return v256
}
func F_DynaHashAlloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, _consts[1185]))
	v5 = F_MemoryContextAllocExtended(m, v3, l0, int32(2))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F___divti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	v13 = int64(63)
	v14 = l2 >> (uint(v13) % 64)
	v15 = v14 ^ l1
	v18 = v15 + int64(base.Ui64(l2)>>(uint(v13)%64))
	v24 = l4 >> (uint(v13) % 64)
	v25 = v24 ^ l3
	v28 = v25 + int64(base.Ui64(l4)>>(uint(v13)%64))
	F___udivmodti4(m, v11, v18, base.I64_extend_i32_u(base.B2i32(base.Ui64(v18) < base.Ui64(v15)))+(l2^v14), v28, base.I64_extend_i32_u(base.B2i32(base.Ui64(v28) < base.Ui64(v25)))+(v24^l4))
	mBase = m.M
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v35 = v24 ^ v14
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v37 = v35 ^ v36
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v37 - v35
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v34 ^ v35 - v35 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v37) < base.Ui64(v35)))
	m.G0 = v11 + v10
	return
}
func F_dasin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v47 float64
	_ = v47
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v61 float64
	_ = v61
	var v66 float64
	_ = v66
	var v70 float64
	_ = v70
	var v79 float64
	_ = v79
	var v88 float64
	_ = v88
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v101 float64
	_ = v101
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v8 = base.F64_abs(v7)
	if base.Ui64(base.I64_reinterpret_f64(v8)) <= base.Ui64(int64(9218868437227405312)) {
		if base.F64_gt(v8, float64(1)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v114 = m.ExcPending
			if v114 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(383107), int32(0))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(470688), int32(1803), int32(261902))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v19 = base.I64_reinterpret_f64(v7)
			v24 = base.I32_wrap_i64(int64(base.Ui64(v19)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1072693248)) <= base.Ui32(v24) {
				if base.I32_wrap_i64(v19)|(v24-int32(1072693248)) == int32(0) {
					v101 = base.F64_add(base.F64_mul(v7, float64(1.5707963267948966)), float64(7.52316384526264e-37))
				} else {
					v101 = base.F64_div(float64(0), base.F64_sub(v7, v7))
				}
			} else {
				if base.Ui32(v24) <= base.Ui32(int32(1071644671)) {
					if base.Ui32(v24+int32(-1048576)) < base.Ui32(int32(1044381696)) {
						v93 = v7
						v101 = v93
					} else {
						v47 = F_R(m, base.F64_mul(v7, v7))
						mBase = m.M
						v101 = base.F64_add(base.F64_mul(v7, v47), v7)
					}
				} else {
					v54 = base.F64_mul(base.F64_sub(float64(1), base.F64_abs(v7)), float64(0.5))
					v55 = base.F64_sqrt(v54)
					v56 = F_R(m, v54)
					mBase = m.M
					if base.Ui32(int32(1072640819)) <= base.Ui32(v24) {
						v61 = base.F64_add(base.F64_mul(v55, v56), v55)
						v88 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v61, v61), float64(-6.123233995736766e-17)))
					} else {
						v66 = float64(0.7853981633974483)
						v70 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v55) & int64(-4294967296))
						v79 = base.F64_div(base.F64_sub(v54, base.F64_mul(v70, v70)), base.F64_add(v55, v70))
						v88 = base.F64_add(base.F64_sub(base.F64_sub(v66, base.F64_add(v70, v70)), base.F64_sub(base.F64_mul(base.F64_add(v55, v55), v56), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v79, v79)))), v66)
					}
					if v19 < int64(0) {
						v92 = base.F64_neg(v88)
					} else {
						v92 = v88
					}
					v93 = v92
					v101 = v93
				}
			}
			if base.F64_eq(base.F64_abs(v101), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v128 = m.ExcPending
				if v128 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v105 = v101
				v106 = F_Float8GetDatum(m, v105)
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return int32(0)
				} else {
					return v106
				}
			}
		}
	} else {
		v105 = math.Float64frombits(uint64(0x7ff8000000000000))
		v106 = F_Float8GetDatum(m, v105)
		mBase = m.M
		v109 = m.ExcPending
		if v109 != 0 {
			return int32(0)
		} else {
			return v106
		}
	}
}
func F_dasinh(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v16 float64
	_ = v16
	var v22 float64
	_ = v22
	var v30 float64
	_ = v30
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v48 int64
	_ = v48
	var v70 float64
	_ = v70
	var v71 int64
	_ = v71
	var v76 int32
	_ = v76
	var v91 float64
	_ = v91
	var v93 float64
	_ = v93
	var v108 float64
	_ = v108
	var v110 float64
	_ = v110
	var v114 float64
	_ = v114
	var v119 float64
	_ = v119
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v124 float64
	_ = v124
	var v162 float64
	_ = v162
	var v171 float64
	_ = v171
	var v172 float64
	_ = v172
	var v177 float64
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = base.F64_abs(v6)
	v8 = base.I64_reinterpret_f64(v6)
	v13 = base.I32_wrap_i64(int64(base.Ui64(v8)>>(uint(int64(52))%64))) & int32(2047)
	if base.Ui32(int32(1049)) <= base.Ui32(v13) {
		v16 = F_log(m, v7)
		mBase = m.M
		v172 = base.F64_add(v16, float64(0.6931471805599453))
	} else {
		if base.Ui32(int32(1024)) <= base.Ui32(v13) {
			v22 = float64(1)
			v30 = F_log(m, base.F64_add(base.F64_add(v7, v7), base.F64_div(v22, base.F64_add(v7, base.F64_sqrt(base.F64_add(base.F64_mul(v6, v6), v22))))))
			mBase = m.M
			v172 = v30
		} else {
			if base.Ui32(v13) < base.Ui32(int32(997)) {
				v172 = v7
			} else {
				v33 = base.F64_mul(v6, v6)
				v34 = float64(1)
				v40 = base.F64_add(v7, base.F64_div(v33, base.F64_add(base.F64_sqrt(base.F64_add(v33, v34)), v34)))
				v41 = float64(0)
				v48 = base.I64_reinterpret_f64(v40)
				if v48 <= int64(4601133429810003967) {
					if base.Ui64(int64(-4616189618054758400)) <= base.Ui64(v48) {
						if base.F64_eq(v40, float64(-1)) != 0 {
							v162 = math.Float64frombits(uint64(0xfff0000000000000))
							v171 = v162
						} else {
							v171 = base.F64_div(base.F64_sub(v40, v40), float64(0))
						}
					} else {
						if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v48)>>(uint(int64(31))%64)))) < base.Ui32(int32(2034237440)) {
							v171 = v40
						} else {
							if base.Ui64(int64(-4624424114038243328)) <= base.Ui64(v48) {
								v70 = base.F64_add(v40, float64(1))
								v71 = base.I64_reinterpret_f64(v70)
								v76 = base.I32_wrap_i64(int64(base.Ui64(v71)>>(uint(int64(32))%64))) + int32(614242)
								if base.Ui32(v76) <= base.Ui32(int32(1129316351)) {
									if base.Ui32(int32(1074790399)) < base.Ui32(v76) {
										v91 = base.F64_add(base.F64_sub(v40, v70), float64(1))
									} else {
										v91 = base.F64_sub(v40, base.F64_add(v70, float64(-1)))
									}
									v93 = base.F64_div(v91, v70)
								} else {
									v93 = v41
								}
								v108 = base.F64_add(base.F64_reinterpret_i64(v71&int64(4294967295)|base.I64_extend_i32_u(v76&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
								v110 = v93
								v114 = base.F64_convert_i32_s(int32(base.Ui32(v76)>>(uint(int32(20))%32)) - int32(1023))
							} else {
								v108 = v40
								v110 = v41
								v114 = float64(0)
							}
							v119 = base.F64_div(v108, base.F64_add(v108, float64(2)))
							v122 = base.F64_mul(v108, base.F64_mul(v108, float64(0.5)))
							v123 = base.F64_mul(v119, v119)
							v124 = base.F64_mul(v123, v123)
							v162 = base.F64_add(base.F64_mul(v114, float64(0.6931471803691238)), base.F64_add(v108, base.F64_sub(base.F64_add(base.F64_mul(v119, base.F64_add(v122, base.F64_add(base.F64_mul(v124, base.F64_add(base.F64_mul(v124, base.F64_add(base.F64_mul(v124, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v123, base.F64_add(base.F64_mul(v124, base.F64_add(base.F64_mul(v124, base.F64_add(base.F64_mul(v124, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_add(base.F64_mul(v114, float64(1.9082149292705877e-10)), v110)), v122)))
							v171 = v162
						}
					}
				} else {
					if base.Ui64(int64(9218868437227405311)) < base.Ui64(v48) {
						v171 = v40
					} else {
						v70 = base.F64_add(v40, float64(1))
						v71 = base.I64_reinterpret_f64(v70)
						v76 = base.I32_wrap_i64(int64(base.Ui64(v71)>>(uint(int64(32))%64))) + int32(614242)
						if base.Ui32(v76) <= base.Ui32(int32(1129316351)) {
							if base.Ui32(int32(1074790399)) < base.Ui32(v76) {
								v91 = base.F64_add(base.F64_sub(v40, v70), float64(1))
							} else {
								v91 = base.F64_sub(v40, base.F64_add(v70, float64(-1)))
							}
							v93 = base.F64_div(v91, v70)
						} else {
							v93 = v41
						}
						v108 = base.F64_add(base.F64_reinterpret_i64(v71&int64(4294967295)|base.I64_extend_i32_u(v76&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
						v110 = v93
						v114 = base.F64_convert_i32_s(int32(base.Ui32(v76)>>(uint(int32(20))%32)) - int32(1023))
						v119 = base.F64_div(v108, base.F64_add(v108, float64(2)))
						v122 = base.F64_mul(v108, base.F64_mul(v108, float64(0.5)))
						v123 = base.F64_mul(v119, v119)
						v124 = base.F64_mul(v123, v123)
						v162 = base.F64_add(base.F64_mul(v114, float64(0.6931471803691238)), base.F64_add(v108, base.F64_sub(base.F64_add(base.F64_mul(v119, base.F64_add(v122, base.F64_add(base.F64_mul(v124, base.F64_add(base.F64_mul(v124, base.F64_add(base.F64_mul(v124, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v123, base.F64_add(base.F64_mul(v124, base.F64_add(base.F64_mul(v124, base.F64_add(base.F64_mul(v124, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_add(base.F64_mul(v114, float64(1.9082149292705877e-10)), v110)), v122)))
						v171 = v162
					}
				}
				v172 = v171
			}
		}
	}
	if v8 < int64(0) {
		v177 = base.F64_neg(v172)
	} else {
		v177 = v172
	}
	v178 = F_Float8GetDatum(m, v177)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		return int32(0)
	} else {
		return v178
	}
}
func F_dbase_desc(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+48)))
	v16 = v14 & int32(240)
	switch v16 - int32(16) {
	case 0:
		goto L3
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L1
	case 16:
		goto L2
	default:
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return
L2:
	;
	F_appendStringInfoString(m, l0, int32(202870))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L9
	}
L3:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = base.I64_rotl(v30, int64(32))
	F_appendStringInfo(m, l0, int32(36513), v10+int32(16))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L8
	}
L4:
	;
	if v16 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v21 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = base.I64_rotl(v20, v21)
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = base.I64_rotl(v19, v21)
	F_appendStringInfo(m, l0, int32(36530), v10)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	goto L1
L8:
	;
	goto L1
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v42 <= int32(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v49 = int32(0)
	goto L11
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(8)+v49<<(uint(int32(2))%32))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v58
	F_appendStringInfo(m, l0, int32(36668), v10+int32(32))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L1
L13:
	;
	v68 = v49 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v68 < v69 {
		v49 = v68
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_dcs_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v10 == int32(0) {
		v29 = v9
		v30 = v10
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v30 - v29
L2:
	;
	goto L1
L3:
	;
	if v9 != v10 {
		v29 = v9
		v30 = v10
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v14 = v4
	v15 = v6
	goto L5
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v19 == int32(0) {
		v29 = v18
		v30 = v19
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v29 = v18
	v30 = v19
	goto L2
L7:
	;
	v22 = int32(1)
	if v18 == v19 {
		v14 = v14 + v22
		v15 = v15 + v22
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
func F_deltraverse(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
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
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v10 = m.T0[v9].(func(*base.Module) int32)(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(101)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v20 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v18 = v16
	goto L8
L7:
	;
	v18 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v18
	return
L9:
	;
	return
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v23 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l1
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v29 = v25
	goto L15
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
	goto L9
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	F_deltraverse(m, l0, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if v35 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+4)))
	if v43 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v111 != 0 {
		goto L46
	} else {
		goto L47
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	if v77 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v48 = v46 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v48) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if int32(1)<<(uint(v48)%32)&int32(163841) == int32(0) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v57 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v58 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v70 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v62+v43*int32(24))+12)) = v66
	v70 = v66
	goto L25
L27:
	;
	goto L28
L28:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+32)) = v68
	v70 = v68
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+36)) = v58
	goto L31
L30:
	;
	goto L31
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = int64(0)
	goto L20
L32:
	;
	if v76 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+20)) = v76
	goto L32
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = v76
	goto L32
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = v77
	goto L38
L37:
	;
	goto L38
L38:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v83 - int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	if v88 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v94 = v29 + int32(8)
	if v87 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v87
	goto L39
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v87
	goto L39
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v88
	goto L45
L44:
	;
	goto L45
L45:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v96 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(0)
	v102 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v94)+16)) = v102
	*(*int64)(unsafe.Add(mBase, uint32(v94)+8)) = v102
	*(*int64)(unsafe.Add(mBase, uint32(v94))) = v102
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v29
	goto L19
L46:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v131 != 0 {
		v29 = v131
		goto L15
	} else {
		goto L57
	}
L47:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	if v112 != 0 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+4)) = uint8(v113)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(-1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	if v118 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	if v117 != 0 {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+32)) = v117
	goto L49
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v117
	goto L49
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = int32(0)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+28)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v31
	goto L46
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = v121
	goto L53
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v121
	goto L53
L57:
	;
	goto L16
}
func F_dense_alloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v8 = (l1 + int32(7)) & int32(-8)
	if base.Ui32(int32(8193)) <= base.Ui32(v8) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		v14 = F_MemoryContextAlloc(m, v11, v8+int32(16))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v8
			*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v8
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(1)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			if v22 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v14
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v14
			}
			return v14 + int32(16)
		}
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
		if v33 != 0 {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
			if base.Ui32(v8) <= base.Ui32(v34-v35) {
				*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v8 + v35
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
				*(*int32)(unsafe.Add(mBase, uint32(v54))) = v55 + int32(1)
				return v33 + v35 + int32(16)
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
				v41 = F_MemoryContextAlloc(m, v39, int32(32784))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v8
					*(*int64)(unsafe.Add(mBase, uint32(v41))) = int64(140737488355329)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v41
					return v41 + int32(16)
				}
			}
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
			v41 = F_MemoryContextAlloc(m, v39, int32(32784))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v8
				*(*int64)(unsafe.Add(mBase, uint32(v41))) = int64(140737488355329)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v46
				*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v41
				return v41 + int32(16)
			}
		}
	}
}
func F_derf(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v37 float64
	_ = v37
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v80 float64
	_ = v80
	var v82 float64
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v8 = base.I64_reinterpret_f64(v7)
	v11 = base.I32_wrap_i64(int64(base.Ui64(v8) >> (uint(int64(32)) % 64)))
	v13 = v11 & int32(2147483647)
	if base.Ui32(int32(2146435072)) <= base.Ui32(v13) {
		v82 = base.F64_add(base.F64_div(float64(1), v7), base.F64_convert_i32_s(int32(1)-int32(base.Ui32(v11)>>(uint(int32(30))%32))&int32(2)))
	} else {
		if base.Ui32(v13) <= base.Ui32(int32(1072365567)) {
			if base.Ui32(v13) <= base.Ui32(int32(1043333119)) {
				v82 = base.F64_mul(base.F64_add(base.F64_mul(v7, float64(8)), base.F64_mul(v7, float64(1.0270333367641007))), float64(0.125))
			} else {
				v37 = base.F64_mul(v7, v7)
				v82 = base.F64_add(base.F64_mul(v7, base.F64_div(base.F64_add(base.F64_mul(v37, base.F64_add(base.F64_mul(v37, base.F64_add(base.F64_mul(v37, base.F64_add(base.F64_mul(v37, float64(-2.3763016656650163e-05)), float64(-0.005770270296489442))), float64(-0.02848174957559851))), float64(-0.3250421072470015))), float64(0.12837916709551256)), base.F64_add(base.F64_mul(v37, base.F64_add(base.F64_mul(v37, base.F64_add(base.F64_mul(v37, base.F64_add(base.F64_mul(v37, base.F64_add(base.F64_mul(v37, float64(-3.960228278775368e-06)), float64(0.00013249473800432164))), float64(0.005081306281875766))), float64(0.0650222499887673))), float64(0.39791722395915535))), float64(1)))), v7)
			}
		} else {
			if base.Ui32(v13) <= base.Ui32(int32(1075314687)) {
				v74 = F_erfc2(m, v13, v7)
				mBase = m.M
				v76 = base.F64_sub(float64(1), v74)
			} else {
				v76 = float64(1)
			}
			if v8 < int64(0) {
				v80 = base.F64_neg(v76)
			} else {
				v80 = v76
			}
			v82 = v80
		}
	}
	if base.F64_eq(base.F64_abs(v82), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v89 = m.ExcPending
		if v89 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v90 = F_Float8GetDatum(m, v82)
		mBase = m.M
		v91 = m.ExcPending
		if v91 != 0 {
			return int32(0)
		} else {
			return v90
		}
	}
}
func F_des_init(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v43 int32
	_ = v43
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v426 int32
	_ = v426
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v526 int32
	_ = v526
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v626 int32
	_ = v626
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v726 int32
	_ = v726
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v820 int64
	_ = v820
	var v837 int32
	_ = v837
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int64
	_ = v895
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1034 int32
	_ = v1034
	var v1050 int32
	_ = v1050
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1071 int32
	_ = v1071
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1102 int32
	_ = v1102
	var v1108 int32
	_ = v1108
	var v1114 int32
	_ = v1114
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1253 int32
	_ = v1253
	var v1259 int32
	_ = v1259
	var v1264 int32
	_ = v1264
	var v1269 int32
	_ = v1269
	var v1278 int32
	_ = v1278
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1295 int32
	_ = v1295
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1323 int32
	_ = v1323
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1351 int32
	_ = v1351
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1379 int32
	_ = v1379
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1407 int32
	_ = v1407
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1435 int32
	_ = v1435
	var v1444 int32
	_ = v1444
	var v1451 int32
	_ = v1451
	var v1455 int32
	_ = v1455
	var v1460 int32
	_ = v1460
	var v1466 int32
	_ = v1466
	var v1473 int32
	_ = v1473
	var v1482 int32
	_ = v1482
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1499 int32
	_ = v1499
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1525 int32
	_ = v1525
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1551 int32
	_ = v1551
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1577 int32
	_ = v1577
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1603 int32
	_ = v1603
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1629 int32
	_ = v1629
	var v1638 int32
	_ = v1638
	var v1645 int32
	_ = v1645
	var v1650 int32
	_ = v1650
	var v1654 int32
	_ = v1654
	var v1658 int32
	_ = v1658
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1718 int32
	_ = v1718
	var v1730 int32
	_ = v1730
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1765 int32
	_ = v1765
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1805 int32
	_ = v1805
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1820 int32
	_ = v1820
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1837 int32
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1854 int32
	_ = v1854
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1861 int32
	_ = v1861
	var v1871 int32
	_ = v1871
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1888 int32
	_ = v1888
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1905 int32
	_ = v1905
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1922 int32
	_ = v1922
	var v1926 int32
	_ = v1926
	var v1931 int32
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1939 int32
	_ = v1939
	v1 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1356])) = v1
	*(*int32)(unsafe.Add(mBase, _consts[1357])) = v1
	*(*int32)(unsafe.Add(mBase, _consts[1358])) = v1
	*(*int32)(unsafe.Add(mBase, _consts[1359])) = v1
	v43 = v1
	for {
		v78 = v43&int32(32) | int32(base.Ui32(v43)>>(uint(int32(1))%32))&int32(15)
		v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+uint32(_consts[1360]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[1361]))) = uint8(v80)
		v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+uint32(_consts[1362]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[1363]))) = uint8(v82)
		v85 = v43 + int32(2)
		if v85 != int32(64) {
			v43 = v85
			continue
		} else {
			break
		}
		break
	}
	v88 = v1
	for {
		v124 = v88&int32(32) | int32(base.Ui32(v88)>>(uint(int32(1))%32))&int32(15)
		v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+uint32(_consts[1364]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_consts[1365]))) = uint8(v126)
		v128 = int32(-64)
		v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(4018240)+v124-v128))))
		*(*uint8)(unsafe.Add(mBase, uint32(int32(4536320)+v88-v128))) = uint8(v132)
		v135 = v88 + int32(2)
		if v135 != int32(64) {
			v88 = v135
			continue
		} else {
			break
		}
		break
	}
	v139 = int32(0)
	for {
		v175 = v139&int32(32) | int32(base.Ui32(v139)>>(uint(int32(1))%32))&int32(15)
		v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[1366]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v139)+uint32(_consts[1367]))) = uint8(v177)
		v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[1368]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v139)+uint32(_consts[1369]))) = uint8(v179)
		v182 = v139 + int32(2)
		if v182 != int32(64) {
			v139 = v182
			continue
		} else {
			break
		}
		break
	}
	v186 = int32(0)
	for {
		v222 = v186&int32(32) | int32(base.Ui32(v186)>>(uint(int32(1))%32))&int32(15)
		v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1370]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v186)+uint32(_consts[1371]))) = uint8(v224)
		v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1372]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v186)+uint32(_consts[1373]))) = uint8(v226)
		v229 = v186 + int32(2)
		if v229 != int32(64) {
			v186 = v229
			continue
		} else {
			break
		}
		break
	}
	v233 = int32(0)
	for {
		v269 = v233&int32(32) | int32(base.Ui32(v233)>>(uint(int32(1))%32))&int32(15)
		v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+uint32(_consts[1374]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v233)+uint32(_consts[1375]))) = uint8(v271)
		v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+uint32(_consts[1376]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v233)+uint32(_consts[1377]))) = uint8(v273)
		v276 = v233 + int32(2)
		if v276 != int32(64) {
			v233 = v276
			continue
		} else {
			break
		}
		break
	}
	v280 = int32(0)
	for {
		v316 = v280&int32(32) | int32(base.Ui32(v280)>>(uint(int32(1))%32))&int32(15)
		v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+uint32(_consts[1378]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v280)+uint32(_consts[1379]))) = uint8(v318)
		v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+uint32(_consts[1380]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v280)+uint32(_consts[1381]))) = uint8(v320)
		v323 = v280 + int32(2)
		if v323 != int32(64) {
			v280 = v323
			continue
		} else {
			break
		}
		break
	}
	v327 = int32(0)
	for {
		v363 = v327&int32(32) | int32(base.Ui32(v327)>>(uint(int32(1))%32))&int32(15)
		v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363)+uint32(_consts[1382]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v327)+uint32(_consts[1383]))) = uint8(v365)
		v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363)+uint32(_consts[1384]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v327)+uint32(_consts[1385]))) = uint8(v367)
		v370 = v327 + int32(2)
		if v370 != int32(64) {
			v327 = v370
			continue
		} else {
			break
		}
		break
	}
	v374 = int32(0)
	for {
		v410 = v374&int32(32) | int32(base.Ui32(v374)>>(uint(int32(1))%32))&int32(15)
		v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+uint32(_consts[1386]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v374)+uint32(_consts[1387]))) = uint8(v412)
		v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+uint32(_consts[1388]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v374)+uint32(_consts[1389]))) = uint8(v414)
		v417 = v374 + int32(2)
		if v417 != int32(64) {
			v374 = v417
			continue
		} else {
			break
		}
		break
	}
	v426 = v1
	for {
		v445 = v426 << (uint(int32(6)) % 32)
		v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426)+uint32(_consts[1361]))))
		v450 = v448 << (uint(int32(4)) % 32)
		v452 = int32(0)
		for {
			v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452)+uint32(_consts[1390]))))
			v485 = v450 | v484
			*(*uint8)(unsafe.Add(mBase, uint32(v452+v445)+uint32(_consts[1391]))) = uint8(v485)
			v488 = v452 | int32(1)
			v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+uint32(_consts[1390]))))
			v493 = v450 | v492
			*(*uint8)(unsafe.Add(mBase, uint32(v488+v445)+uint32(_consts[1391]))) = uint8(v493)
			v496 = v452 | int32(2)
			v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+uint32(_consts[1390]))))
			v501 = v450 | v500
			*(*uint8)(unsafe.Add(mBase, uint32(v496+v445)+uint32(_consts[1391]))) = uint8(v501)
			v504 = v452 | int32(3)
			v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504)+uint32(_consts[1390]))))
			v509 = v450 | v508
			*(*uint8)(unsafe.Add(mBase, uint32(v504+v445)+uint32(_consts[1391]))) = uint8(v509)
			v512 = v452 + int32(4)
			if v512 != int32(64) {
				v452 = v512
				continue
			} else {
				break
			}
			break
		}
		v516 = v426 + int32(1)
		if v516 != int32(64) {
			v426 = v516
			continue
		} else {
			break
		}
		break
	}
	v526 = int32(0)
	for {
		v545 = v526 << (uint(int32(6)) % 32)
		v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526)+uint32(_consts[1369]))))
		v550 = v548 << (uint(int32(4)) % 32)
		v552 = int32(0)
		for {
			v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+uint32(_consts[1373]))))
			v585 = v550 | v584
			*(*uint8)(unsafe.Add(mBase, uint32(v552+v545)+uint32(_consts[1392]))) = uint8(v585)
			v588 = v552 | int32(1)
			v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588)+uint32(_consts[1373]))))
			v593 = v550 | v592
			*(*uint8)(unsafe.Add(mBase, uint32(v588+v545)+uint32(_consts[1392]))) = uint8(v593)
			v596 = v552 | int32(2)
			v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+uint32(_consts[1373]))))
			v601 = v550 | v600
			*(*uint8)(unsafe.Add(mBase, uint32(v596+v545)+uint32(_consts[1392]))) = uint8(v601)
			v604 = v552 | int32(3)
			v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+uint32(_consts[1373]))))
			v609 = v550 | v608
			*(*uint8)(unsafe.Add(mBase, uint32(v604+v545)+uint32(_consts[1392]))) = uint8(v609)
			v612 = v552 + int32(4)
			if v612 != int32(64) {
				v552 = v612
				continue
			} else {
				break
			}
			break
		}
		v616 = v526 + int32(1)
		if v616 != int32(64) {
			v526 = v616
			continue
		} else {
			break
		}
		break
	}
	v626 = int32(0)
	for {
		v645 = v626 << (uint(int32(6)) % 32)
		v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+uint32(_consts[1377]))))
		v650 = v648 << (uint(int32(4)) % 32)
		v652 = int32(0)
		for {
			v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652)+uint32(_consts[1381]))))
			v685 = v650 | v684
			*(*uint8)(unsafe.Add(mBase, uint32(v652+v645)+uint32(_consts[1393]))) = uint8(v685)
			v688 = v652 | int32(1)
			v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688)+uint32(_consts[1381]))))
			v693 = v650 | v692
			*(*uint8)(unsafe.Add(mBase, uint32(v688+v645)+uint32(_consts[1393]))) = uint8(v693)
			v696 = v652 | int32(2)
			v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696)+uint32(_consts[1381]))))
			v701 = v650 | v700
			*(*uint8)(unsafe.Add(mBase, uint32(v696+v645)+uint32(_consts[1393]))) = uint8(v701)
			v704 = v652 | int32(3)
			v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704)+uint32(_consts[1381]))))
			v709 = v650 | v708
			*(*uint8)(unsafe.Add(mBase, uint32(v704+v645)+uint32(_consts[1393]))) = uint8(v709)
			v712 = v652 + int32(4)
			if v712 != int32(64) {
				v652 = v712
				continue
			} else {
				break
			}
			break
		}
		v716 = v626 + int32(1)
		if v716 != int32(64) {
			v626 = v716
			continue
		} else {
			break
		}
		break
	}
	v726 = int32(0)
	for {
		v745 = v726 << (uint(int32(6)) % 32)
		v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726)+uint32(_consts[1385]))))
		v750 = v748 << (uint(int32(4)) % 32)
		v752 = int32(0)
		for {
			v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v752)+uint32(_consts[1389]))))
			v785 = v750 | v784
			*(*uint8)(unsafe.Add(mBase, uint32(v752+v745)+uint32(_consts[1394]))) = uint8(v785)
			v788 = v752 | int32(1)
			v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788)+uint32(_consts[1389]))))
			v793 = v750 | v792
			*(*uint8)(unsafe.Add(mBase, uint32(v788+v745)+uint32(_consts[1394]))) = uint8(v793)
			v796 = v752 | int32(2)
			v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796)+uint32(_consts[1389]))))
			v801 = v750 | v800
			*(*uint8)(unsafe.Add(mBase, uint32(v796+v745)+uint32(_consts[1394]))) = uint8(v801)
			v804 = v752 | int32(3)
			v808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804)+uint32(_consts[1389]))))
			v809 = v750 | v808
			*(*uint8)(unsafe.Add(mBase, uint32(v804+v745)+uint32(_consts[1394]))) = uint8(v809)
			v812 = v752 + int32(4)
			if v812 != int32(64) {
				v752 = v812
				continue
			} else {
				break
			}
			break
		}
		v816 = v726 + int32(1)
		if v816 != int32(64) {
			v726 = v816
			continue
		} else {
			break
		}
		break
	}
	v819 = int32(4553344)
	v820 = int64(-1)
	*(*int64)(unsafe.Add(mBase, _consts[1395])) = v820
	*(*int64)(unsafe.Add(mBase, _consts[1396])) = v820
	*(*int64)(unsafe.Add(mBase, _consts[1397])) = v820
	*(*int64)(unsafe.Add(mBase, _consts[1398])) = v820
	*(*int64)(unsafe.Add(mBase, _consts[1399])) = v820
	*(*int64)(unsafe.Add(mBase, _consts[1400])) = v820
	*(*int64)(unsafe.Add(mBase, _consts[1401])) = v820
	*(*int64)(unsafe.Add(mBase, _consts[1402])) = v820
	v837 = int32(0)
	for {
		v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837)+uint32(_consts[1403]))))
		v869 = int32(1)
		v870 = v868 - v869
		*(*uint8)(unsafe.Add(mBase, uint32(v837)+uint32(_consts[1404]))) = uint8(v870)
		v874 = int32(255)
		*(*uint8)(unsafe.Add(mBase, uint32(v870&v874)+uint32(_consts[1405]))) = uint8(v837)
		v879 = v837 | v869
		v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v879)+uint32(_consts[1403]))))
		v884 = v882 - v869
		*(*uint8)(unsafe.Add(mBase, uint32(v879)+uint32(_consts[1404]))) = uint8(v884)
		*(*uint8)(unsafe.Add(mBase, uint32(v884&v874)+uint32(_consts[1405]))) = uint8(v879)
		v891 = v837 + int32(2)
		if v891 != int32(64) {
			v837 = v891
			continue
		} else {
			break
		}
		break
	}
	v894 = int32(4553408)
	v895 = int64(-1)
	*(*int64)(unsafe.Add(mBase, _consts[1406])) = v895
	*(*int64)(unsafe.Add(mBase, _consts[1407])) = v895
	*(*int64)(unsafe.Add(mBase, _consts[1408])) = v895
	*(*int64)(unsafe.Add(mBase, _consts[1409])) = v895
	*(*int64)(unsafe.Add(mBase, _consts[1410])) = v895
	*(*int64)(unsafe.Add(mBase, _consts[1411])) = v895
	*(*int64)(unsafe.Add(mBase, _consts[1412])) = v895
	v909 = int32(0)
	v912 = v909
	for {
		v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v912)+uint32(_consts[1413]))))
		v940 = int32(4553344)
		v943 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v939+v940-v943))) = uint8(v912)
		v947 = v912 | v943
		v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v947)+uint32(_consts[1413]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v949+v940-v943))) = uint8(v947)
		v955 = v912 | int32(2)
		v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955)+uint32(_consts[1413]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v957+v940-v943))) = uint8(v955)
		v963 = v912 | int32(3)
		v965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963)+uint32(_consts[1413]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v965+v940-v943))) = uint8(v963)
		v971 = v912 + int32(4)
		if v971 != int32(56) {
			v912 = v971
			continue
		} else {
			break
		}
		break
	}
	v974 = v909
	for {
		v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v974)+uint32(_consts[1414]))))
		v1003 = int32(4553408)
		v1006 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v1002+v1003-v1006))) = uint8(v974)
		v1010 = v974 | v1006
		v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1010)+uint32(_consts[1414]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v1012+v1003-v1006))) = uint8(v1010)
		v1018 = v974 | int32(2)
		v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1018)+uint32(_consts[1414]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v1020+v1003-v1006))) = uint8(v1018)
		v1026 = v974 | int32(3)
		v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1026)+uint32(_consts[1414]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v1028+v1003-v1006))) = uint8(v1026)
		v1034 = v974 + int32(4)
		if v1034 != int32(48) {
			v974 = v1034
			continue
		} else {
			break
		}
		break
	}
	v1050 = v1
	for {
		v1062 = v1050 << (uint(int32(3)) % 32)
		v1064 = v1050 << (uint(int32(10)) % 32)
		v1071 = int32(0)
		for {
			v1091 = v1071 << (uint(int32(2)) % 32)
			v1092 = int32(0)
			v1096 = v1091 + (int32(4553472) + v1064)
			*(*int32)(unsafe.Add(mBase, uint32(v1096))) = v1092
			v1102 = int32(4561664) + v1064 + v1091
			*(*int32)(unsafe.Add(mBase, uint32(v1102))) = v1092
			v1108 = int32(4569856) + v1064 + v1091
			*(*int32)(unsafe.Add(mBase, uint32(v1108))) = v1092
			v1114 = int32(4578048) + v1064 + v1091
			*(*int32)(unsafe.Add(mBase, uint32(v1114))) = v1092
			v1121 = v1092
			v1124 = v1092
			v1125 = v1092
			v1127 = v1092
			v1128 = v1092
			for {
				v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1121)+uint32(_consts[1415]))))
				if v1071&v1147 == int32(0) {
					v1194 = v1124
					v1195 = v1125
					v1196 = v1127
					v1197 = v1128
				} else {
					v1151 = v1121 + v1062
					v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+uint32(_consts[1405]))))
					if base.Ui32(v1154) <= base.Ui32(int32(31)) {
						v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1154<<(uint(int32(2))%32))+uint32(_consts[1416])))
						v1162 = v1127 | v1161
						*(*int32)(unsafe.Add(mBase, uint32(v1096))) = v1162
						v1171 = v1162
						v1172 = v1128
					} else {
						v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1154<<(uint(int32(2))%32))+uint32(_consts[1417])))
						v1169 = v1128 | v1168
						*(*int32)(unsafe.Add(mBase, uint32(v1102))) = v1169
						v1171 = v1127
						v1172 = v1169
					}
					v1175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+uint32(_consts[1404]))))
					if base.Ui32(v1175) <= base.Ui32(int32(31)) {
						v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1175<<(uint(int32(2))%32))+uint32(_consts[1416])))
						v1183 = v1124 | v1182
						*(*int32)(unsafe.Add(mBase, uint32(v1108))) = v1183
						v1194 = v1183
						v1195 = v1125
						v1196 = v1171
						v1197 = v1172
					} else {
						v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1175<<(uint(int32(2))%32))+uint32(_consts[1417])))
						v1190 = v1125 | v1189
						*(*int32)(unsafe.Add(mBase, uint32(v1114))) = v1190
						v1194 = v1124
						v1195 = v1190
						v1196 = v1171
						v1197 = v1172
					}
				}
				v1199 = v1121 + int32(1)
				if v1199 != int32(8) {
					v1121 = v1199
					v1124 = v1194
					v1125 = v1195
					v1127 = v1196
					v1128 = v1197
					continue
				} else {
					break
				}
				break
			}
			v1203 = v1071 + int32(1)
			if v1203 != int32(256) {
				v1071 = v1203
				continue
			} else {
				break
			}
			break
		}
		v1219 = v1050 * int32(7)
		v1221 = v1050 << (uint(int32(9)) % 32)
		v1223 = int32(0)
		for {
			v1248 = v1223 << (uint(int32(2)) % 32)
			v1249 = int32(0)
			v1253 = v1248 + (int32(4586240) + v1221)
			*(*int32)(unsafe.Add(mBase, uint32(v1253))) = v1249
			v1259 = int32(4590336) + v1221 + v1248
			*(*int32)(unsafe.Add(mBase, uint32(v1259))) = v1249
			v1264 = v1223 & int32(64)
			if v1264 == v1249 {
				v1286 = v1249
				v1287 = v1249
			} else {
				v1269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062)+uint32(_consts[1395]))))
				if v1269 == int32(255) {
					v1286 = v1249
					v1287 = v1249
				} else {
					if base.Ui32(v1269) <= base.Ui32(int32(27)) {
						v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1269<<(uint(int32(2))%32))+uint32(_consts[1418])))
						*(*int32)(unsafe.Add(mBase, uint32(v1253))) = v1278
						v1286 = v1278
						v1287 = v1249
					} else {
						v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1269<<(uint(int32(2))%32))+uint32(_consts[1419])))
						*(*int32)(unsafe.Add(mBase, uint32(v1259))) = v1284
						v1286 = v1249
						v1287 = v1284
					}
				}
			}
			v1290 = v1223 & int32(32)
			if v1290 == int32(0) {
				v1314 = v1286
				v1315 = v1287
			} else {
				v1295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062|int32(1))+uint32(_consts[1395]))))
				if v1295 == int32(255) {
					v1314 = v1286
					v1315 = v1287
				} else {
					if base.Ui32(int32(28)) <= base.Ui32(v1295) {
						v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1295<<(uint(int32(2))%32))+uint32(_consts[1419])))
						v1305 = v1287 | v1304
						*(*int32)(unsafe.Add(mBase, uint32(v1259))) = v1305
						v1314 = v1286
						v1315 = v1305
					} else {
						v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1295<<(uint(int32(2))%32))+uint32(_consts[1418])))
						v1312 = v1286 | v1311
						*(*int32)(unsafe.Add(mBase, uint32(v1253))) = v1312
						v1314 = v1312
						v1315 = v1287
					}
				}
			}
			v1318 = v1223 & int32(16)
			if v1318 == int32(0) {
				v1342 = v1314
				v1343 = v1315
			} else {
				v1323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062|int32(2))+uint32(_consts[1395]))))
				if v1323 == int32(255) {
					v1342 = v1314
					v1343 = v1315
				} else {
					if base.Ui32(int32(28)) <= base.Ui32(v1323) {
						v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1323<<(uint(int32(2))%32))+uint32(_consts[1419])))
						v1333 = v1315 | v1332
						*(*int32)(unsafe.Add(mBase, uint32(v1259))) = v1333
						v1342 = v1314
						v1343 = v1333
					} else {
						v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1323<<(uint(int32(2))%32))+uint32(_consts[1418])))
						v1340 = v1314 | v1339
						*(*int32)(unsafe.Add(mBase, uint32(v1253))) = v1340
						v1342 = v1340
						v1343 = v1315
					}
				}
			}
			v1346 = v1223 & int32(8)
			if v1346 == int32(0) {
				v1370 = v1342
				v1371 = v1343
			} else {
				v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062|int32(3))+uint32(_consts[1395]))))
				if v1351 == int32(255) {
					v1370 = v1342
					v1371 = v1343
				} else {
					if base.Ui32(int32(28)) <= base.Ui32(v1351) {
						v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1351<<(uint(int32(2))%32))+uint32(_consts[1419])))
						v1361 = v1343 | v1360
						*(*int32)(unsafe.Add(mBase, uint32(v1259))) = v1361
						v1370 = v1342
						v1371 = v1361
					} else {
						v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1351<<(uint(int32(2))%32))+uint32(_consts[1418])))
						v1368 = v1342 | v1367
						*(*int32)(unsafe.Add(mBase, uint32(v1253))) = v1368
						v1370 = v1368
						v1371 = v1343
					}
				}
			}
			v1374 = v1223 & int32(4)
			if v1374 == int32(0) {
				v1398 = v1370
				v1399 = v1371
			} else {
				v1379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062|int32(4))+uint32(_consts[1395]))))
				if v1379 == int32(255) {
					v1398 = v1370
					v1399 = v1371
				} else {
					if base.Ui32(int32(28)) <= base.Ui32(v1379) {
						v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1379<<(uint(int32(2))%32))+uint32(_consts[1419])))
						v1389 = v1371 | v1388
						*(*int32)(unsafe.Add(mBase, uint32(v1259))) = v1389
						v1398 = v1370
						v1399 = v1389
					} else {
						v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1379<<(uint(int32(2))%32))+uint32(_consts[1418])))
						v1396 = v1370 | v1395
						*(*int32)(unsafe.Add(mBase, uint32(v1253))) = v1396
						v1398 = v1396
						v1399 = v1371
					}
				}
			}
			v1402 = v1223 & int32(2)
			if v1402 == int32(0) {
				v1426 = v1398
				v1427 = v1399
			} else {
				v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062|int32(5))+uint32(_consts[1395]))))
				if v1407 == int32(255) {
					v1426 = v1398
					v1427 = v1399
				} else {
					if base.Ui32(int32(28)) <= base.Ui32(v1407) {
						v1416 = *(*int32)(unsafe.Add(mBase, uint32(v1407<<(uint(int32(2))%32))+uint32(_consts[1419])))
						v1417 = v1399 | v1416
						*(*int32)(unsafe.Add(mBase, uint32(v1259))) = v1417
						v1426 = v1398
						v1427 = v1417
					} else {
						v1423 = *(*int32)(unsafe.Add(mBase, uint32(v1407<<(uint(int32(2))%32))+uint32(_consts[1418])))
						v1424 = v1398 | v1423
						*(*int32)(unsafe.Add(mBase, uint32(v1253))) = v1424
						v1426 = v1424
						v1427 = v1399
					}
				}
			}
			v1430 = v1223 & int32(1)
			if v1430 == int32(0) {
			} else {
				v1435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062|int32(6))+uint32(_consts[1395]))))
				if v1435 == int32(255) {
				} else {
					if base.Ui32(int32(28)) <= base.Ui32(v1435) {
						v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1435<<(uint(int32(2))%32))+uint32(_consts[1419])))
						*(*int32)(unsafe.Add(mBase, uint32(v1259))) = v1427 | v1444
					} else {
						v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1435<<(uint(int32(2))%32))+uint32(_consts[1418])))
						*(*int32)(unsafe.Add(mBase, uint32(v1253))) = v1426 | v1451
					}
				}
			}
			v1455 = int32(0)
			v1460 = int32(4594432) + v1221 + v1248
			*(*int32)(unsafe.Add(mBase, uint32(v1460))) = v1455
			v1466 = int32(4598528) + v1221 + v1248
			*(*int32)(unsafe.Add(mBase, uint32(v1466))) = v1455
			if v1264 == v1455 {
				v1491 = v1455
				v1493 = int32(0)
				v1494 = v1491
			} else {
				v1473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+uint32(_consts[1406]))))
				if v1473 == int32(255) {
					v1491 = v1455
					v1493 = int32(0)
					v1494 = v1491
				} else {
					if base.Ui32(v1473) <= base.Ui32(int32(23)) {
						v1482 = *(*int32)(unsafe.Add(mBase, uint32(v1473<<(uint(int32(2))%32))+uint32(_consts[1420])))
						*(*int32)(unsafe.Add(mBase, uint32(v1460))) = v1482
						v1493 = v1482
						v1494 = v1455
					} else {
						v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1473<<(uint(int32(2))%32))+uint32(_consts[1421])))
						*(*int32)(unsafe.Add(mBase, uint32(v1466))) = v1488
						v1491 = v1488
						v1493 = int32(0)
						v1494 = v1491
					}
				}
			}
			if v1290 == int32(0) {
				v1518 = v1493
				v1519 = v1494
			} else {
				v1499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+uint32(_consts[1422]))))
				if v1499 == int32(255) {
					v1518 = v1493
					v1519 = v1494
				} else {
					if base.Ui32(int32(24)) <= base.Ui32(v1499) {
						v1508 = *(*int32)(unsafe.Add(mBase, uint32(v1499<<(uint(int32(2))%32))+uint32(_consts[1421])))
						v1509 = v1494 | v1508
						*(*int32)(unsafe.Add(mBase, uint32(v1466))) = v1509
						v1518 = v1493
						v1519 = v1509
					} else {
						v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1499<<(uint(int32(2))%32))+uint32(_consts[1420])))
						v1516 = v1493 | v1515
						*(*int32)(unsafe.Add(mBase, uint32(v1460))) = v1516
						v1518 = v1516
						v1519 = v1494
					}
				}
			}
			if v1318 == int32(0) {
				v1544 = v1518
				v1545 = v1519
			} else {
				v1525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+uint32(_consts[1423]))))
				if v1525 == int32(255) {
					v1544 = v1518
					v1545 = v1519
				} else {
					if base.Ui32(int32(24)) <= base.Ui32(v1525) {
						v1534 = *(*int32)(unsafe.Add(mBase, uint32(v1525<<(uint(int32(2))%32))+uint32(_consts[1421])))
						v1535 = v1519 | v1534
						*(*int32)(unsafe.Add(mBase, uint32(v1466))) = v1535
						v1544 = v1518
						v1545 = v1535
					} else {
						v1541 = *(*int32)(unsafe.Add(mBase, uint32(v1525<<(uint(int32(2))%32))+uint32(_consts[1420])))
						v1542 = v1518 | v1541
						*(*int32)(unsafe.Add(mBase, uint32(v1460))) = v1542
						v1544 = v1542
						v1545 = v1519
					}
				}
			}
			if v1346 == int32(0) {
				v1570 = v1544
				v1571 = v1545
			} else {
				v1551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+uint32(_consts[1424]))))
				if v1551 == int32(255) {
					v1570 = v1544
					v1571 = v1545
				} else {
					if base.Ui32(int32(24)) <= base.Ui32(v1551) {
						v1560 = *(*int32)(unsafe.Add(mBase, uint32(v1551<<(uint(int32(2))%32))+uint32(_consts[1421])))
						v1561 = v1545 | v1560
						*(*int32)(unsafe.Add(mBase, uint32(v1466))) = v1561
						v1570 = v1544
						v1571 = v1561
					} else {
						v1567 = *(*int32)(unsafe.Add(mBase, uint32(v1551<<(uint(int32(2))%32))+uint32(_consts[1420])))
						v1568 = v1544 | v1567
						*(*int32)(unsafe.Add(mBase, uint32(v1460))) = v1568
						v1570 = v1568
						v1571 = v1545
					}
				}
			}
			if v1374 == int32(0) {
				v1596 = v1570
				v1597 = v1571
			} else {
				v1577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+uint32(_consts[1425]))))
				if v1577 == int32(255) {
					v1596 = v1570
					v1597 = v1571
				} else {
					if base.Ui32(int32(24)) <= base.Ui32(v1577) {
						v1586 = *(*int32)(unsafe.Add(mBase, uint32(v1577<<(uint(int32(2))%32))+uint32(_consts[1421])))
						v1587 = v1571 | v1586
						*(*int32)(unsafe.Add(mBase, uint32(v1466))) = v1587
						v1596 = v1570
						v1597 = v1587
					} else {
						v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1577<<(uint(int32(2))%32))+uint32(_consts[1420])))
						v1594 = v1570 | v1593
						*(*int32)(unsafe.Add(mBase, uint32(v1460))) = v1594
						v1596 = v1594
						v1597 = v1571
					}
				}
			}
			if v1402 == int32(0) {
				v1622 = v1596
				v1623 = v1597
			} else {
				v1603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+uint32(_consts[1426]))))
				if v1603 == int32(255) {
					v1622 = v1596
					v1623 = v1597
				} else {
					if base.Ui32(int32(24)) <= base.Ui32(v1603) {
						v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1603<<(uint(int32(2))%32))+uint32(_consts[1421])))
						v1613 = v1597 | v1612
						*(*int32)(unsafe.Add(mBase, uint32(v1466))) = v1613
						v1622 = v1596
						v1623 = v1613
					} else {
						v1619 = *(*int32)(unsafe.Add(mBase, uint32(v1603<<(uint(int32(2))%32))+uint32(_consts[1420])))
						v1620 = v1596 | v1619
						*(*int32)(unsafe.Add(mBase, uint32(v1460))) = v1620
						v1622 = v1620
						v1623 = v1597
					}
				}
			}
			if v1430 == int32(0) {
			} else {
				v1629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+uint32(_consts[1427]))))
				if v1629 == int32(255) {
				} else {
					if base.Ui32(int32(24)) <= base.Ui32(v1629) {
						v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1629<<(uint(int32(2))%32))+uint32(_consts[1421])))
						*(*int32)(unsafe.Add(mBase, uint32(v1466))) = v1623 | v1638
					} else {
						v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1629<<(uint(int32(2))%32))+uint32(_consts[1420])))
						*(*int32)(unsafe.Add(mBase, uint32(v1460))) = v1622 | v1645
					}
				}
			}
			v1650 = v1223 + int32(1)
			if v1650 != int32(128) {
				v1223 = v1650
				continue
			} else {
				break
			}
			break
		}
		v1654 = v1050 + int32(1)
		if v1654 != int32(8) {
			v1050 = v1654
			continue
		} else {
			break
		}
		break
	}
	v1658 = int32(0)
	for {
		v1686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1658)+uint32(_consts[1428]))))
		v1687 = int32(4602624)
		v1690 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v1686+v1687-v1690))) = uint8(v1658)
		v1694 = v1658 | v1690
		v1696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1694)+uint32(_consts[1428]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v1696+v1687-v1690))) = uint8(v1694)
		v1702 = v1658 | int32(2)
		v1704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1702)+uint32(_consts[1428]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v1704+v1687-v1690))) = uint8(v1702)
		v1710 = v1658 | int32(3)
		v1712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1710)+uint32(_consts[1428]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v1712+v1687-v1690))) = uint8(v1710)
		v1718 = v1658 + int32(4)
		if v1718 != int32(32) {
			v1658 = v1718
			continue
		} else {
			break
		}
		break
	}
	v1730 = int32(0)
	for {
		v1746 = int32(3)
		v1747 = v1730 << (uint(v1746) % 32)
		v1765 = int32(0)
		for {
			v1793 = v1730<<(uint(int32(10))%32) + int32(4602656) + v1765<<(uint(int32(2))%32)
			v1794 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v1793))) = v1794
			if v1765&int32(128) != 0 {
				v1805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1747)+uint32(_consts[1429]))))
				v1809 = *(*int32)(unsafe.Add(mBase, uint32(v1805<<(uint(int32(2))%32))+uint32(_consts[1416])))
				*(*int32)(unsafe.Add(mBase, uint32(v1793))) = v1809
				v1811 = v1809
			} else {
				v1811 = v1794
			}
			if v1765&int32(64) != 0 {
				v1820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1747|int32(1))+uint32(_consts[1429]))))
				v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1820<<(uint(int32(2))%32))+uint32(_consts[1416])))
				v1825 = v1811 | v1824
				*(*int32)(unsafe.Add(mBase, uint32(v1793))) = v1825
				v1827 = v1825
			} else {
				v1827 = v1811
			}
			if v1765&int32(32) != 0 {
				v1837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1747|int32(2))+uint32(_consts[1429]))))
				v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1837<<(uint(int32(2))%32))+uint32(_consts[1416])))
				v1842 = v1827 | v1841
				*(*int32)(unsafe.Add(mBase, uint32(v1793))) = v1842
				v1844 = v1842
			} else {
				v1844 = v1827
			}
			if v1765&int32(16) != 0 {
				v1854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1747|v1746)+uint32(_consts[1429]))))
				v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1854<<(uint(int32(2))%32))+uint32(_consts[1416])))
				v1859 = v1844 | v1858
				*(*int32)(unsafe.Add(mBase, uint32(v1793))) = v1859
				v1861 = v1859
			} else {
				v1861 = v1844
			}
			if v1765&int32(8) != 0 {
				v1871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1747|int32(4))+uint32(_consts[1429]))))
				v1875 = *(*int32)(unsafe.Add(mBase, uint32(v1871<<(uint(int32(2))%32))+uint32(_consts[1416])))
				v1876 = v1861 | v1875
				*(*int32)(unsafe.Add(mBase, uint32(v1793))) = v1876
				v1878 = v1876
			} else {
				v1878 = v1861
			}
			if v1765&int32(4) != 0 {
				v1888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1747|int32(5))+uint32(_consts[1429]))))
				v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1888<<(uint(int32(2))%32))+uint32(_consts[1416])))
				v1893 = v1878 | v1892
				*(*int32)(unsafe.Add(mBase, uint32(v1793))) = v1893
				v1895 = v1893
			} else {
				v1895 = v1878
			}
			if v1765&int32(2) != 0 {
				v1905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1747|int32(6))+uint32(_consts[1429]))))
				v1909 = *(*int32)(unsafe.Add(mBase, uint32(v1905<<(uint(int32(2))%32))+uint32(_consts[1416])))
				v1910 = v1895 | v1909
				*(*int32)(unsafe.Add(mBase, uint32(v1793))) = v1910
				v1912 = v1910
			} else {
				v1912 = v1895
			}
			if v1765&int32(1) != 0 {
				v1922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1747|int32(7))+uint32(_consts[1429]))))
				v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1922<<(uint(int32(2))%32))+uint32(_consts[1416])))
				*(*int32)(unsafe.Add(mBase, uint32(v1793))) = v1912 | v1926
			} else {
			}
			v1931 = v1765 + int32(1)
			if v1931 != int32(256) {
				v1765 = v1931
				continue
			} else {
				break
			}
			break
		}
		v1935 = v1730 + int32(1)
		if v1935 != int32(4) {
			v1730 = v1935
			continue
		} else {
			break
		}
		break
	}
	v1939 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1430])) = uint8(v1939)
	return
}
func F_discard_stack_value(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v5 != int32(3) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v38 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v38
	if v37 == v38 {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	if v8 == v9 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v8 == v14 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v8 == v16 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v8 == v18 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v24 = l0 + int32(56)
	goto L7
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_pfree(m, v8)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v8 == v27 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	if v8 != v29 {
		v24 = v26
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L1
L14:
	;
	return
L15:
	;
	goto L1
L16:
	;
	return
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v37 == v42 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v44 {
	case 0:
		goto L24
	case 1:
		goto L23
	case 2:
		goto L22
	case 3:
		goto L21
	case 4:
		goto L20
	default:
		goto L19
	}
L19:
	;
	v57 = l0 + int32(56)
	goto L30
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v37 == v53 {
		goto L16
	} else {
		goto L29
	}
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v37 != v51 {
		goto L19
	} else {
		goto L28
	}
L22:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v37 != v49 {
		goto L19
	} else {
		goto L27
	}
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v37 != v47 {
		goto L19
	} else {
		goto L26
	}
L24:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v37 != v45 {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	goto L16
L26:
	;
	goto L16
L27:
	;
	goto L16
L28:
	;
	goto L16
L29:
	;
	goto L19
L30:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v61 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	F_pfree(m, v37)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L14
	} else {
		goto L37
	}
L32:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
	if v37 == v62 {
		goto L16
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L31
L35:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
	if v37 != v64 {
		v57 = v61
		goto L30
	} else {
		goto L36
	}
L36:
	;
	goto L16
L37:
	;
	goto L16
}
func F_dispose_chunk(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var __phi56 int32
	_ = __phi56
	var v58 int32
	_ = v58
	var __phi58 int32
	_ = __phi58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var __phi216 int32
	_ = __phi216
	var v218 int32
	_ = v218
	var __phi218 int32
	_ = __phi218
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v298 int32
	_ = v298
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v410 int32
	_ = v410
	v10 = l0 + l1
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11&int32(1) != 0 {
		v128 = l0
		v129 = l1
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v137&int32(2) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L3:
	;
	if v11&int32(2) == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = v18 + l1
	v20 = l0 - v18
	v22 = *(*int32)(unsafe.Add(mBase, _consts[1437]))
	if v20 != v22 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	if v38 == int32(0) {
		v128 = v20
		v129 = v19
		goto L2
	} else {
		goto L27
	}
L6:
	;
	v89 = int32(0)
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v27
	v128 = v20
	v129 = v19
	goto L2
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if base.Ui32(v18) <= base.Ui32(int32(255)) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v71 = int32(3)
	if v70&v71 != v71 {
		v128 = v20
		v129 = v19
		goto L2
	} else {
		goto L26
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v24 != v27 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v20 != v24 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v29 = int32(4619144)
	v31 = *(*int32)(unsafe.Add(mBase, _consts[1438]))
	*(*int32)(unsafe.Add(mBase, _consts[1438])) = v31 & base.I32_rotl(int32(-2), int32(base.Ui32(v18)>>(uint(int32(3))%32)))
	v128 = v20
	v129 = v19
	goto L2
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v40
	v89 = v24
	goto L5
L16:
	;
	goto L17
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v43 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v51 = v43
	v52 = v20 + int32(20)
	goto L20
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v46 == int32(0) {
		goto L6
	} else {
		goto L21
	}
L20:
	;
	__phi56 = v51
	__phi58 = v52
	v56 = __phi56
	v58 = __phi58
	goto L22
L21:
	;
	v51 = v46
	v52 = v20 + int32(16)
	goto L20
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	if v64 != 0 {
		__phi56 = v64
		__phi58 = v56 + int32(20)
		v56 = __phi56
		v58 = __phi58
		goto L22
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(0)
	v89 = v56
	goto L5
L24:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	if v67 != 0 {
		__phi56 = v67
		__phi58 = v56 + int32(16)
		v56 = __phi56
		v58 = __phi58
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1439])) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v70 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v19 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v19
	return
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v100 = v98 << (uint(int32(2)) % 32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+uint32(_consts[1440])))
	if v103 == v20 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = v38
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v120 != 0 {
		goto L38
	} else {
		goto L39
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+uint32(_consts[1440]))) = v89
	if v89 != 0 {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v20 == v113 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v106 = int32(4619148)
	v108 = *(*int32)(unsafe.Add(mBase, _consts[1441]))
	*(*int32)(unsafe.Add(mBase, _consts[1441])) = v108 & base.I32_rotl(int32(-2), v98)
	v128 = v20
	v129 = v19
	goto L2
L33:
	;
	if v89 == int32(0) {
		v128 = v20
		v129 = v19
		goto L2
	} else {
		goto L37
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v89
	goto L33
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v89
	goto L33
L37:
	;
	goto L28
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+16)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v120)+24)) = v89
	goto L40
L39:
	;
	goto L40
L40:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v123 == int32(0) {
		v128 = v20
		v129 = v19
		goto L2
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+20)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v123)+24)) = v89
	v128 = v20
	v129 = v19
	goto L2
L42:
	;
	if base.Ui32(v298) <= base.Ui32(int32(255)) {
		goto L89
	} else {
		goto L90
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v181 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v128+v181))) = v181
	if v128 != v165 {
		v298 = v181
		goto L42
	} else {
		goto L88
	}
L44:
	;
	if v198 == int32(0) {
		goto L43
	} else {
		goto L73
	}
L45:
	;
	v241 = int32(0)
	goto L44
L46:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _consts[1442]))
	if v143 == v10 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v137 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v129 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v128+v129))) = v129
	v298 = v129
	goto L42
L49:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1442])) = v128
	v147 = int32(4619156)
	v149 = *(*int32)(unsafe.Add(mBase, _consts[1443]))
	v150 = v149 + v129
	*(*int32)(unsafe.Add(mBase, _consts[1443])) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v150 | int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, _consts[1437]))
	if v128 != v156 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _consts[1437]))
	if v165 == v10 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v159 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1439])) = v159
	*(*int32)(unsafe.Add(mBase, _consts[1437])) = v159
	return
L53:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1437])) = v128
	v169 = int32(4619152)
	v171 = *(*int32)(unsafe.Add(mBase, _consts[1439]))
	v172 = v171 + v129
	*(*int32)(unsafe.Add(mBase, _consts[1439])) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v172 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v128+v172))) = v172
	return
L54:
	;
	goto L55
L55:
	;
	v181 = v137&int32(-8) + v129
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if base.Ui32(v137) <= base.Ui32(int32(255)) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v185 == v182 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v182 != v10 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v187 = int32(4619144)
	v189 = *(*int32)(unsafe.Add(mBase, _consts[1438]))
	*(*int32)(unsafe.Add(mBase, _consts[1438])) = v189 & base.I32_rotl(int32(-2), int32(base.Ui32(v137)>>(uint(int32(3))%32)))
	goto L43
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+12)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v182)+8)) = v185
	goto L43
L62:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v182)+8)) = v200
	v241 = v182
	goto L44
L63:
	;
	goto L64
L64:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v203 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v211 = v203
	v212 = v10 + int32(20)
	goto L67
L66:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v206 == int32(0) {
		goto L45
	} else {
		goto L68
	}
L67:
	;
	__phi216 = v211
	__phi218 = v212
	v216 = __phi216
	v218 = __phi218
	goto L69
L68:
	;
	v211 = v206
	v212 = v10 + int32(16)
	goto L67
L69:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v216)+20))
	if v224 != 0 {
		__phi216 = v224
		__phi218 = v216 + int32(20)
		v216 = __phi216
		v218 = __phi218
		goto L69
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218))) = int32(0)
	v241 = v216
	goto L44
L71:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	if v227 != 0 {
		__phi216 = v227
		__phi218 = v216 + int32(16)
		v216 = __phi216
		v218 = __phi218
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v252 = v250 << (uint(int32(2)) % 32)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v252)+uint32(_consts[1440])))
	if v255 == v10 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v241)+24)) = v198
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v272 != 0 {
		goto L84
	} else {
		goto L85
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252)+uint32(_consts[1440]))) = v241
	if v241 != 0 {
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v198)+16))
	if v10 == v265 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v258 = int32(4619148)
	v260 = *(*int32)(unsafe.Add(mBase, _consts[1441]))
	*(*int32)(unsafe.Add(mBase, _consts[1441])) = v260 & base.I32_rotl(int32(-2), v250)
	goto L43
L79:
	;
	if v241 == int32(0) {
		goto L43
	} else {
		goto L83
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+16)) = v241
	goto L79
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+20)) = v241
	goto L79
L83:
	;
	goto L74
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v241)+16)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v272)+24)) = v241
	goto L86
L85:
	;
	goto L86
L86:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v275 == int32(0) {
		goto L43
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v275)+24)) = v241
	goto L43
L88:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1439])) = v181
	return
L89:
	;
	v309 = v298 & int32(-8)
	v311 = v309 + int32(4619184)
	v313 = *(*int32)(unsafe.Add(mBase, _consts[1438]))
	v317 = int32(1) << (uint(int32(base.Ui32(v298)>>(uint(int32(3))%32))) % 32)
	if v313&v317 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	goto L91
L91:
	;
	if base.Ui32(v298) <= base.Ui32(int32(16777215)) {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+uint32(_consts[1444]))) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v325)+12)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = v325
	return
L93:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1438])) = v317 | v313
	v325 = v311
	goto L92
L94:
	;
	goto L95
L95:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v309)+uint32(_consts[1444])))
	v325 = v324
	goto L92
L96:
	;
	v336 = base.I32_clz(int32(base.Ui32(v298) >> (uint(int32(8)) % 32)))
	v339 = int32(1)
	v346 = int32(base.Ui32(v298)>>(uint(int32(38)-v336)%32))&v339 - v336<<(uint(v339)%32) + int32(62)
	goto L98
L97:
	;
	v346 = int32(31)
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+28)) = v346
	*(*int64)(unsafe.Add(mBase, uint32(v128)+16)) = int64(0)
	v351 = v346 << (uint(int32(2)) % 32)
	v355 = *(*int32)(unsafe.Add(mBase, _consts[1441]))
	v357 = int32(1) << (uint(v346) % 32)
	if v355&v357 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v381)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v410)+12)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v381)+8)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v128)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v381
	*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = v410
	goto L1
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = v128
	return
L101:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1441])) = v357 | v355
	*(*int32)(unsafe.Add(mBase, uint32(v351)+uint32(_consts[1440]))) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v128)+24)) = v351 + int32(4619448)
	goto L100
L102:
	;
	goto L103
L103:
	;
	if v346 != int32(31) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v373 = int32(25) - int32(base.Ui32(v346)>>(uint(int32(1))%32))
	goto L106
L105:
	;
	v373 = int32(0)
	goto L106
L106:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v351)+uint32(_consts[1440])))
	v378 = v298 << (uint(v373) % 32)
	v381 = v375
	goto L107
L107:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	if v385&int32(-8) == v298 {
		goto L99
	} else {
		goto L109
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+16)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v128)+24)) = v381
	goto L100
L109:
	;
	v395 = v381 + int32(base.Ui32(v378)>>(uint(int32(29))%32))&int32(4)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+16))
	if v396 != 0 {
		v378 = v378 << (uint(int32(1)) % 32)
		v381 = v396
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
}
func F_dlgamma(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v18 int64
	_ = v18
	var v23 int32
	_ = v23
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v49 float64
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 float64
	_ = v66
	var v74 float64
	_ = v74
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v129 float64
	_ = v129
	var v145 float64
	_ = v145
	var v151 float64
	_ = v151
	var v190 float64
	_ = v190
	var v191 float64
	_ = v191
	var v193 float64
	_ = v193
	var v194 float64
	_ = v194
	var v206 float64
	_ = v206
	var v223 float64
	_ = v223
	var v232 float64
	_ = v232
	var v236 float64
	_ = v236
	var v238 float64
	_ = v238
	var v239 float64
	_ = v239
	var v255 float64
	_ = v255
	var v256 float64
	_ = v256
	var v267 float64
	_ = v267
	var v268 float64
	_ = v268
	var v269 float64
	_ = v269
	var v272 float64
	_ = v272
	var v314 float64
	_ = v314
	var v315 float64
	_ = v315
	var v316 float64
	_ = v316
	var v317 float64
	_ = v317
	var v369 float64
	_ = v369
	var v370 float64
	_ = v370
	var v411 float64
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 float64
	_ = v419
	var v462 float64
	_ = v462
	var v467 float64
	_ = v467
	var v471 float64
	_ = v471
	var v475 float64
	_ = v475
	var v479 float64
	_ = v479
	var v483 float64
	_ = v483
	var v485 float64
	_ = v485
	var v494 float64
	_ = v494
	var v495 float64
	_ = v495
	var v520 float64
	_ = v520
	var v526 float64
	_ = v526
	var v533 float64
	_ = v533
	var v535 int32
	_ = v535
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	v2 = float64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1075])) = int32(1)
	v18 = base.I64_reinterpret_f64(v10)
	v23 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(2146435072)) <= base.Ui32(v23) {
		v533 = base.F64_mul(v10, v10)
	} else {
		if base.Ui32(v23) <= base.Ui32(int32(999292927)) {
			if v18 < int64(0) {
				*(*int32)(unsafe.Add(mBase, _consts[1075])) = int32(-1)
				v34 = base.F64_neg(v10)
			} else {
				v34 = v10
			}
			v35 = F_log(m, v34)
			mBase = m.M
			v533 = base.F64_neg(v35)
		} else {
			if int64(0) <= v18 {
				v238 = v10
				v239 = v2
				if base.I32_wrap_i64(v18) == int32(0) {
					if base.B2i32(v23 == int32(1072693248))|base.B2i32(v23 == int32(1073741824)) != 0 {
						v520 = float64(0)
					} else {
						if base.Ui32(v23) <= base.Ui32(int32(1073741823)) {
							if base.Ui32(v23) <= base.Ui32(int32(1072483532)) {
								v255 = F_log(m, v238)
								mBase = m.M
								v256 = base.F64_neg(v255)
								if base.Ui32(int32(1072130371)) < base.Ui32(v23) {
									v267 = v256
									v268 = float64(1)
									v269 = base.F64_sub(v268, v238)
									v272 = base.F64_mul(v269, v269)
									v520 = base.F64_add(v267, base.F64_add(base.F64_mul(v269, float64(-0.5)), base.F64_add(base.F64_mul(v269, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
								} else {
									if base.Ui32(v23) <= base.Ui32(int32(1070442080)) {
										v369 = v238
										v370 = v256
										v520 = base.F64_add(v370, base.F64_add(base.F64_mul(v369, float64(-0.5)), base.F64_div(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
									} else {
										v314 = v256
										v315 = base.F64_add(v238, float64(-0.46163214496836225))
										v316 = base.F64_mul(v315, v315)
										v317 = base.F64_mul(v315, v316)
										v520 = base.F64_add(v314, base.F64_add(base.F64_sub(base.F64_mul(v316, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v317, base.F64_add(base.F64_mul(v315, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
									}
								}
							} else {
								if base.Ui32(v23) <= base.Ui32(int32(1073460418)) {
									if base.Ui32(v23) < base.Ui32(int32(1072936132)) {
										v369 = base.F64_add(v238, float64(-1))
										v370 = v2
										v520 = base.F64_add(v370, base.F64_add(base.F64_mul(v369, float64(-0.5)), base.F64_div(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
									} else {
										v314 = v2
										v315 = base.F64_add(v238, float64(-1.4616321449683622))
										v316 = base.F64_mul(v315, v315)
										v317 = base.F64_mul(v315, v316)
										v520 = base.F64_add(v314, base.F64_add(base.F64_sub(base.F64_mul(v316, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v317, base.F64_add(base.F64_mul(v315, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
									}
								} else {
									v267 = v2
									v268 = float64(2)
									v269 = base.F64_sub(v268, v238)
									v272 = base.F64_mul(v269, v269)
									v520 = base.F64_add(v267, base.F64_add(base.F64_mul(v269, float64(-0.5)), base.F64_add(base.F64_mul(v269, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
								}
							}
						} else {
							if base.Ui32(v23) <= base.Ui32(int32(1075838975)) {
								v411 = float64(1)
								if base.F64_lt(base.F64_abs(v238), float64(2.147483648e+09)) != 0 {
									v415 = base.I32_trunc_f64_s(v238)
									v417 = v415
								} else {
									v417 = int32(-2147483648)
								}
								v419 = base.F64_sub(v238, base.F64_convert_i32_s(v417))
								v462 = base.F64_add(base.F64_mul(v419, float64(0.5)), base.F64_div(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, float64(3.194753265841009e-05)), float64(0.0018402845140733772))), float64(0.02664227030336386))), float64(0.14635047265246445))), float64(0.325778796408931))), float64(0.21498241596060885))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, float64(7.326684307446256e-06)), float64(0.0007779424963818936))), float64(0.01864591917156529))), float64(0.17193386563280308))), float64(0.7219355475671381))), float64(1.3920053346762105))), float64(1))))
								switch v417 - int32(3) {
								case 0:
									v479 = v411
									v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
									mBase = m.M
									v520 = base.F64_add(v462, v483)
								case 1:
									v475 = v411
									v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
									v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
									mBase = m.M
									v520 = base.F64_add(v462, v483)
								case 2:
									v471 = v411
									v475 = base.F64_mul(base.F64_add(v419, float64(4)), v471)
									v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
									v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
									mBase = m.M
									v520 = base.F64_add(v462, v483)
								case 3:
									v467 = v411
									v471 = base.F64_mul(base.F64_add(v419, float64(5)), v467)
									v475 = base.F64_mul(base.F64_add(v419, float64(4)), v471)
									v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
									v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
									mBase = m.M
									v520 = base.F64_add(v462, v483)
								case 4:
									v467 = base.F64_add(v419, float64(6))
									v471 = base.F64_mul(base.F64_add(v419, float64(5)), v467)
									v475 = base.F64_mul(base.F64_add(v419, float64(4)), v471)
									v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
									v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
									mBase = m.M
									v520 = base.F64_add(v462, v483)
								default:
									v520 = v462
								}
							} else {
								v485 = F_log(m, v238)
								mBase = m.M
								if base.Ui32(v23) <= base.Ui32(int32(1133510655)) {
									v494 = base.F64_div(float64(1), v238)
									v495 = base.F64_mul(v494, v494)
									v520 = base.F64_add(base.F64_mul(base.F64_add(v238, float64(-0.5)), base.F64_add(v485, float64(-1))), base.F64_add(base.F64_mul(v494, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, float64(-0.0016309293409657527)), float64(0.0008363399189962821))), float64(-0.00059518755745034))), float64(0.0007936505586430196))), float64(-0.0027777777772877554))), float64(0.08333333333333297))), float64(0.4189385332046727)))
								} else {
									v520 = base.F64_mul(v238, base.F64_add(v485, float64(-1)))
								}
							}
						}
					}
				} else {
					if base.Ui32(v23) <= base.Ui32(int32(1073741823)) {
						if base.Ui32(v23) <= base.Ui32(int32(1072483532)) {
							v255 = F_log(m, v238)
							mBase = m.M
							v256 = base.F64_neg(v255)
							if base.Ui32(int32(1072130371)) < base.Ui32(v23) {
								v267 = v256
								v268 = float64(1)
								v269 = base.F64_sub(v268, v238)
								v272 = base.F64_mul(v269, v269)
								v520 = base.F64_add(v267, base.F64_add(base.F64_mul(v269, float64(-0.5)), base.F64_add(base.F64_mul(v269, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
							} else {
								if base.Ui32(v23) <= base.Ui32(int32(1070442080)) {
									v369 = v238
									v370 = v256
									v520 = base.F64_add(v370, base.F64_add(base.F64_mul(v369, float64(-0.5)), base.F64_div(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
								} else {
									v314 = v256
									v315 = base.F64_add(v238, float64(-0.46163214496836225))
									v316 = base.F64_mul(v315, v315)
									v317 = base.F64_mul(v315, v316)
									v520 = base.F64_add(v314, base.F64_add(base.F64_sub(base.F64_mul(v316, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v317, base.F64_add(base.F64_mul(v315, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
								}
							}
						} else {
							if base.Ui32(v23) <= base.Ui32(int32(1073460418)) {
								if base.Ui32(v23) < base.Ui32(int32(1072936132)) {
									v369 = base.F64_add(v238, float64(-1))
									v370 = v2
									v520 = base.F64_add(v370, base.F64_add(base.F64_mul(v369, float64(-0.5)), base.F64_div(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
								} else {
									v314 = v2
									v315 = base.F64_add(v238, float64(-1.4616321449683622))
									v316 = base.F64_mul(v315, v315)
									v317 = base.F64_mul(v315, v316)
									v520 = base.F64_add(v314, base.F64_add(base.F64_sub(base.F64_mul(v316, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v317, base.F64_add(base.F64_mul(v315, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
								}
							} else {
								v267 = v2
								v268 = float64(2)
								v269 = base.F64_sub(v268, v238)
								v272 = base.F64_mul(v269, v269)
								v520 = base.F64_add(v267, base.F64_add(base.F64_mul(v269, float64(-0.5)), base.F64_add(base.F64_mul(v269, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
							}
						}
					} else {
						if base.Ui32(v23) <= base.Ui32(int32(1075838975)) {
							v411 = float64(1)
							if base.F64_lt(base.F64_abs(v238), float64(2.147483648e+09)) != 0 {
								v415 = base.I32_trunc_f64_s(v238)
								v417 = v415
							} else {
								v417 = int32(-2147483648)
							}
							v419 = base.F64_sub(v238, base.F64_convert_i32_s(v417))
							v462 = base.F64_add(base.F64_mul(v419, float64(0.5)), base.F64_div(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, float64(3.194753265841009e-05)), float64(0.0018402845140733772))), float64(0.02664227030336386))), float64(0.14635047265246445))), float64(0.325778796408931))), float64(0.21498241596060885))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, float64(7.326684307446256e-06)), float64(0.0007779424963818936))), float64(0.01864591917156529))), float64(0.17193386563280308))), float64(0.7219355475671381))), float64(1.3920053346762105))), float64(1))))
							switch v417 - int32(3) {
							case 0:
								v479 = v411
								v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
								mBase = m.M
								v520 = base.F64_add(v462, v483)
							case 1:
								v475 = v411
								v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
								v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
								mBase = m.M
								v520 = base.F64_add(v462, v483)
							case 2:
								v471 = v411
								v475 = base.F64_mul(base.F64_add(v419, float64(4)), v471)
								v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
								v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
								mBase = m.M
								v520 = base.F64_add(v462, v483)
							case 3:
								v467 = v411
								v471 = base.F64_mul(base.F64_add(v419, float64(5)), v467)
								v475 = base.F64_mul(base.F64_add(v419, float64(4)), v471)
								v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
								v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
								mBase = m.M
								v520 = base.F64_add(v462, v483)
							case 4:
								v467 = base.F64_add(v419, float64(6))
								v471 = base.F64_mul(base.F64_add(v419, float64(5)), v467)
								v475 = base.F64_mul(base.F64_add(v419, float64(4)), v471)
								v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
								v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
								mBase = m.M
								v520 = base.F64_add(v462, v483)
							default:
								v520 = v462
							}
						} else {
							v485 = F_log(m, v238)
							mBase = m.M
							if base.Ui32(v23) <= base.Ui32(int32(1133510655)) {
								v494 = base.F64_div(float64(1), v238)
								v495 = base.F64_mul(v494, v494)
								v520 = base.F64_add(base.F64_mul(base.F64_add(v238, float64(-0.5)), base.F64_add(v485, float64(-1))), base.F64_add(base.F64_mul(v494, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, float64(-0.0016309293409657527)), float64(0.0008363399189962821))), float64(-0.00059518755745034))), float64(0.0007936505586430196))), float64(-0.0027777777772877554))), float64(0.08333333333333297))), float64(0.4189385332046727)))
							} else {
								v520 = base.F64_mul(v238, base.F64_add(v485, float64(-1)))
							}
						}
					}
				}
				if int64(0) <= v18 {
					v526 = v520
				} else {
					v526 = base.F64_sub(v239, v520)
				}
				v533 = v526
			} else {
				v42 = base.F64_neg(v10)
				v44 = base.F64_mul(v42, float64(0.5))
				v46 = base.F64_sub(v44, base.F64_floor(v44))
				v47 = base.F64_add(v46, v46)
				v49 = base.F64_mul(v47, float64(4))
				if base.F64_lt(base.F64_abs(v49), float64(2.147483648e+09)) != 0 {
					v53 = base.I32_trunc_f64_s(v49)
					v55 = v53
				} else {
					v55 = int32(-2147483648)
				}
				v56 = int32(1)
				v59 = base.I32_div_s(v55+v56, int32(2))
				v66 = base.F64_mul(base.F64_add(v47, base.F64_promote_f32(base.F32_mul(base.F32_convert_i32_s(v59), float32(-0.5)))), float64(3.141592653589793))
				switch v59 - v56 {
				case 0:
					v113 = float64(1)
					v114 = base.F64_mul(v66, v66)
					v116 = base.F64_mul(v114, float64(0.5))
					v117 = base.F64_sub(v113, v116)
					v129 = base.F64_mul(v114, v114)
					v223 = base.F64_add(v117, base.F64_add(base.F64_sub(base.F64_sub(v113, v117), v116), base.F64_sub(base.F64_mul(v114, base.F64_add(base.F64_mul(v114, base.F64_add(base.F64_mul(v114, base.F64_add(base.F64_mul(v114, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v129, v129), base.F64_add(base.F64_mul(v114, base.F64_add(base.F64_mul(v114, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v66, float64(0)))))
				case 1:
					v145 = base.F64_neg(v66)
					v151 = base.F64_mul(v145, v145)
					v223 = base.F64_add(base.F64_mul(base.F64_mul(v145, v151), base.F64_add(base.F64_mul(v151, base.F64_add(base.F64_mul(base.F64_mul(v151, base.F64_mul(v151, v151)), base.F64_add(base.F64_mul(v151, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v151, base.F64_add(base.F64_mul(v151, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), v145)
				case 2:
					v190 = float64(1)
					v191 = base.F64_mul(v66, v66)
					v193 = base.F64_mul(v191, float64(0.5))
					v194 = base.F64_sub(v190, v193)
					v206 = base.F64_mul(v191, v191)
					v223 = base.F64_neg(base.F64_add(v194, base.F64_add(base.F64_sub(base.F64_sub(v190, v194), v193), base.F64_sub(base.F64_mul(v191, base.F64_add(base.F64_mul(v191, base.F64_add(base.F64_mul(v191, base.F64_add(base.F64_mul(v191, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v206, v206), base.F64_add(base.F64_mul(v191, base.F64_add(base.F64_mul(v191, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v66, float64(0))))))
				default:
					v74 = base.F64_mul(v66, v66)
					v223 = base.F64_add(base.F64_mul(base.F64_mul(v66, v74), base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(base.F64_mul(v74, base.F64_mul(v74, v74)), base.F64_add(base.F64_mul(v74, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), v66)
				}
				if base.F64_eq(v223, float64(0)) != 0 {
					v533 = base.F64_div(float64(1), base.F64_sub(v10, v10))
				} else {
					if base.F64_gt(v223, float64(0)) != 0 {
						*(*int32)(unsafe.Add(mBase, _consts[1075])) = int32(-1)
						v232 = v223
					} else {
						v232 = base.F64_neg(v223)
					}
					v236 = F_log(m, base.F64_div(float64(3.141592653589793), base.F64_mul(v232, v42)))
					mBase = m.M
					v238 = v42
					v239 = v236
					if base.I32_wrap_i64(v18) == int32(0) {
						if base.B2i32(v23 == int32(1072693248))|base.B2i32(v23 == int32(1073741824)) != 0 {
							v520 = float64(0)
						} else {
							if base.Ui32(v23) <= base.Ui32(int32(1073741823)) {
								if base.Ui32(v23) <= base.Ui32(int32(1072483532)) {
									v255 = F_log(m, v238)
									mBase = m.M
									v256 = base.F64_neg(v255)
									if base.Ui32(int32(1072130371)) < base.Ui32(v23) {
										v267 = v256
										v268 = float64(1)
										v269 = base.F64_sub(v268, v238)
										v272 = base.F64_mul(v269, v269)
										v520 = base.F64_add(v267, base.F64_add(base.F64_mul(v269, float64(-0.5)), base.F64_add(base.F64_mul(v269, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
									} else {
										if base.Ui32(v23) <= base.Ui32(int32(1070442080)) {
											v369 = v238
											v370 = v256
											v520 = base.F64_add(v370, base.F64_add(base.F64_mul(v369, float64(-0.5)), base.F64_div(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
										} else {
											v314 = v256
											v315 = base.F64_add(v238, float64(-0.46163214496836225))
											v316 = base.F64_mul(v315, v315)
											v317 = base.F64_mul(v315, v316)
											v520 = base.F64_add(v314, base.F64_add(base.F64_sub(base.F64_mul(v316, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v317, base.F64_add(base.F64_mul(v315, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
										}
									}
								} else {
									if base.Ui32(v23) <= base.Ui32(int32(1073460418)) {
										if base.Ui32(v23) < base.Ui32(int32(1072936132)) {
											v369 = base.F64_add(v238, float64(-1))
											v370 = v2
											v520 = base.F64_add(v370, base.F64_add(base.F64_mul(v369, float64(-0.5)), base.F64_div(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
										} else {
											v314 = v2
											v315 = base.F64_add(v238, float64(-1.4616321449683622))
											v316 = base.F64_mul(v315, v315)
											v317 = base.F64_mul(v315, v316)
											v520 = base.F64_add(v314, base.F64_add(base.F64_sub(base.F64_mul(v316, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v317, base.F64_add(base.F64_mul(v315, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
										}
									} else {
										v267 = v2
										v268 = float64(2)
										v269 = base.F64_sub(v268, v238)
										v272 = base.F64_mul(v269, v269)
										v520 = base.F64_add(v267, base.F64_add(base.F64_mul(v269, float64(-0.5)), base.F64_add(base.F64_mul(v269, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
									}
								}
							} else {
								if base.Ui32(v23) <= base.Ui32(int32(1075838975)) {
									v411 = float64(1)
									if base.F64_lt(base.F64_abs(v238), float64(2.147483648e+09)) != 0 {
										v415 = base.I32_trunc_f64_s(v238)
										v417 = v415
									} else {
										v417 = int32(-2147483648)
									}
									v419 = base.F64_sub(v238, base.F64_convert_i32_s(v417))
									v462 = base.F64_add(base.F64_mul(v419, float64(0.5)), base.F64_div(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, float64(3.194753265841009e-05)), float64(0.0018402845140733772))), float64(0.02664227030336386))), float64(0.14635047265246445))), float64(0.325778796408931))), float64(0.21498241596060885))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, float64(7.326684307446256e-06)), float64(0.0007779424963818936))), float64(0.01864591917156529))), float64(0.17193386563280308))), float64(0.7219355475671381))), float64(1.3920053346762105))), float64(1))))
									switch v417 - int32(3) {
									case 0:
										v479 = v411
										v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
										mBase = m.M
										v520 = base.F64_add(v462, v483)
									case 1:
										v475 = v411
										v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
										v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
										mBase = m.M
										v520 = base.F64_add(v462, v483)
									case 2:
										v471 = v411
										v475 = base.F64_mul(base.F64_add(v419, float64(4)), v471)
										v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
										v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
										mBase = m.M
										v520 = base.F64_add(v462, v483)
									case 3:
										v467 = v411
										v471 = base.F64_mul(base.F64_add(v419, float64(5)), v467)
										v475 = base.F64_mul(base.F64_add(v419, float64(4)), v471)
										v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
										v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
										mBase = m.M
										v520 = base.F64_add(v462, v483)
									case 4:
										v467 = base.F64_add(v419, float64(6))
										v471 = base.F64_mul(base.F64_add(v419, float64(5)), v467)
										v475 = base.F64_mul(base.F64_add(v419, float64(4)), v471)
										v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
										v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
										mBase = m.M
										v520 = base.F64_add(v462, v483)
									default:
										v520 = v462
									}
								} else {
									v485 = F_log(m, v238)
									mBase = m.M
									if base.Ui32(v23) <= base.Ui32(int32(1133510655)) {
										v494 = base.F64_div(float64(1), v238)
										v495 = base.F64_mul(v494, v494)
										v520 = base.F64_add(base.F64_mul(base.F64_add(v238, float64(-0.5)), base.F64_add(v485, float64(-1))), base.F64_add(base.F64_mul(v494, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, float64(-0.0016309293409657527)), float64(0.0008363399189962821))), float64(-0.00059518755745034))), float64(0.0007936505586430196))), float64(-0.0027777777772877554))), float64(0.08333333333333297))), float64(0.4189385332046727)))
									} else {
										v520 = base.F64_mul(v238, base.F64_add(v485, float64(-1)))
									}
								}
							}
						}
					} else {
						if base.Ui32(v23) <= base.Ui32(int32(1073741823)) {
							if base.Ui32(v23) <= base.Ui32(int32(1072483532)) {
								v255 = F_log(m, v238)
								mBase = m.M
								v256 = base.F64_neg(v255)
								if base.Ui32(int32(1072130371)) < base.Ui32(v23) {
									v267 = v256
									v268 = float64(1)
									v269 = base.F64_sub(v268, v238)
									v272 = base.F64_mul(v269, v269)
									v520 = base.F64_add(v267, base.F64_add(base.F64_mul(v269, float64(-0.5)), base.F64_add(base.F64_mul(v269, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
								} else {
									if base.Ui32(v23) <= base.Ui32(int32(1070442080)) {
										v369 = v238
										v370 = v256
										v520 = base.F64_add(v370, base.F64_add(base.F64_mul(v369, float64(-0.5)), base.F64_div(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
									} else {
										v314 = v256
										v315 = base.F64_add(v238, float64(-0.46163214496836225))
										v316 = base.F64_mul(v315, v315)
										v317 = base.F64_mul(v315, v316)
										v520 = base.F64_add(v314, base.F64_add(base.F64_sub(base.F64_mul(v316, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v317, base.F64_add(base.F64_mul(v315, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
									}
								}
							} else {
								if base.Ui32(v23) <= base.Ui32(int32(1073460418)) {
									if base.Ui32(v23) < base.Ui32(int32(1072936132)) {
										v369 = base.F64_add(v238, float64(-1))
										v370 = v2
										v520 = base.F64_add(v370, base.F64_add(base.F64_mul(v369, float64(-0.5)), base.F64_div(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, base.F64_add(base.F64_mul(v369, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
									} else {
										v314 = v2
										v315 = base.F64_add(v238, float64(-1.4616321449683622))
										v316 = base.F64_mul(v315, v315)
										v317 = base.F64_mul(v315, v316)
										v520 = base.F64_add(v314, base.F64_add(base.F64_sub(base.F64_mul(v316, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v317, base.F64_add(base.F64_mul(v315, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, base.F64_add(base.F64_mul(v317, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
									}
								} else {
									v267 = v2
									v268 = float64(2)
									v269 = base.F64_sub(v268, v238)
									v272 = base.F64_mul(v269, v269)
									v520 = base.F64_add(v267, base.F64_add(base.F64_mul(v269, float64(-0.5)), base.F64_add(base.F64_mul(v269, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, base.F64_add(base.F64_mul(v272, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
								}
							}
						} else {
							if base.Ui32(v23) <= base.Ui32(int32(1075838975)) {
								v411 = float64(1)
								if base.F64_lt(base.F64_abs(v238), float64(2.147483648e+09)) != 0 {
									v415 = base.I32_trunc_f64_s(v238)
									v417 = v415
								} else {
									v417 = int32(-2147483648)
								}
								v419 = base.F64_sub(v238, base.F64_convert_i32_s(v417))
								v462 = base.F64_add(base.F64_mul(v419, float64(0.5)), base.F64_div(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, float64(3.194753265841009e-05)), float64(0.0018402845140733772))), float64(0.02664227030336386))), float64(0.14635047265246445))), float64(0.325778796408931))), float64(0.21498241596060885))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, base.F64_add(base.F64_mul(v419, float64(7.326684307446256e-06)), float64(0.0007779424963818936))), float64(0.01864591917156529))), float64(0.17193386563280308))), float64(0.7219355475671381))), float64(1.3920053346762105))), float64(1))))
								switch v417 - int32(3) {
								case 0:
									v479 = v411
									v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
									mBase = m.M
									v520 = base.F64_add(v462, v483)
								case 1:
									v475 = v411
									v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
									v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
									mBase = m.M
									v520 = base.F64_add(v462, v483)
								case 2:
									v471 = v411
									v475 = base.F64_mul(base.F64_add(v419, float64(4)), v471)
									v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
									v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
									mBase = m.M
									v520 = base.F64_add(v462, v483)
								case 3:
									v467 = v411
									v471 = base.F64_mul(base.F64_add(v419, float64(5)), v467)
									v475 = base.F64_mul(base.F64_add(v419, float64(4)), v471)
									v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
									v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
									mBase = m.M
									v520 = base.F64_add(v462, v483)
								case 4:
									v467 = base.F64_add(v419, float64(6))
									v471 = base.F64_mul(base.F64_add(v419, float64(5)), v467)
									v475 = base.F64_mul(base.F64_add(v419, float64(4)), v471)
									v479 = base.F64_mul(base.F64_add(v419, float64(3)), v475)
									v483 = F_log(m, base.F64_mul(base.F64_add(v419, float64(2)), v479))
									mBase = m.M
									v520 = base.F64_add(v462, v483)
								default:
									v520 = v462
								}
							} else {
								v485 = F_log(m, v238)
								mBase = m.M
								if base.Ui32(v23) <= base.Ui32(int32(1133510655)) {
									v494 = base.F64_div(float64(1), v238)
									v495 = base.F64_mul(v494, v494)
									v520 = base.F64_add(base.F64_mul(base.F64_add(v238, float64(-0.5)), base.F64_add(v485, float64(-1))), base.F64_add(base.F64_mul(v494, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, base.F64_add(base.F64_mul(v495, float64(-0.0016309293409657527)), float64(0.0008363399189962821))), float64(-0.00059518755745034))), float64(0.0007936505586430196))), float64(-0.0027777777772877554))), float64(0.08333333333333297))), float64(0.4189385332046727)))
								} else {
									v520 = base.F64_mul(v238, base.F64_add(v485, float64(-1)))
								}
							}
						}
					}
					if int64(0) <= v18 {
						v526 = v520
					} else {
						v526 = base.F64_sub(v239, v520)
					}
					v533 = v526
				}
			}
		}
	}
	v535 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v535 != int32(68) {
		if base.F64_eq(base.F64_abs(v10), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v548 = F_Float8GetDatum(m, v533)
			mBase = m.M
			v549 = m.ExcPending
			if v549 != 0 {
				return int32(0)
			} else {
				return v548
			}
		} else {
			if base.F64_ne(base.F64_abs(v533), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v548 = F_Float8GetDatum(m, v533)
				mBase = m.M
				v549 = m.ExcPending
				if v549 != 0 {
					return int32(0)
				} else {
					return v548
				}
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v547 = m.ExcPending
				if v547 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v547 = m.ExcPending
		if v547 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_dlog1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v10 float64
	_ = v10
	var v12 float64
	_ = v12
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	if base.F64_ne(v5, float64(0)) != 0 {
		if base.F64_lt(v5, float64(0)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352583810))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(217351), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(470688), int32(1706), int32(529117))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v10 = F_log(m, v5)
			mBase = m.M
			v12 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(base.F64_abs(v10), v12)&base.F64_ne(base.F64_abs(v5), v12) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_eq(v10, float64(0))&base.F64_ne(v5, float64(1)) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v23 = F_Float8GetDatum(m, v10)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return v23
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(352583810))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(227949), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(470688), int32(1702), int32(529117))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
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
func F_dlog10(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v20 int64
	_ = v20
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v47 int64
	_ = v47
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 float64
	_ = v60
	var v62 float64
	_ = v62
	var v75 float64
	_ = v75
	var v78 float64
	_ = v78
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v118 float64
	_ = v118
	var v130 float64
	_ = v130
	var v152 float64
	_ = v152
	var v154 float64
	_ = v154
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	if base.F64_ne(v5, float64(0)) != 0 {
		if base.F64_lt(v5, float64(0)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v189 = m.ExcPending
			if v189 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352583810))
				mBase = m.M
				v192 = m.ExcPending
				if v192 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(217351), int32(0))
					mBase = m.M
					v196 = m.ExcPending
					if v196 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(470688), int32(1739), int32(530153))
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v20 = base.I64_reinterpret_f64(v5)
			if v20 <= int64(4503599627370495) {
				if base.F64_eq(v5, float64(0)) != 0 {
					v152 = base.F64_div(float64(-1), base.F64_mul(v5, v5))
				} else {
					if int64(0) <= v20 {
						v47 = base.I64_reinterpret_f64(base.F64_mul(v5, float64(1.8014398509481984e+16)))
						v51 = v47
						v53 = int32(-1077)
						v54 = base.I32_wrap_i64(int64(base.Ui64(v47) >> (uint(int64(32)) % 64)))
						v56 = v54 + int32(614242)
						v60 = base.F64_convert_i32_s(int32(base.Ui32(v56)>>(uint(int32(20))%32)) + v53)
						v62 = base.F64_mul(v60, float64(0.30102999566361177))
						v75 = base.F64_add(base.F64_reinterpret_i64(v51&int64(4294967295)|base.I64_extend_i32_u(v56&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
						v78 = base.F64_mul(v75, base.F64_mul(v75, float64(0.5)))
						v83 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v75, v78)) & int64(-4294967296))
						v84 = float64(0.4342944818781689)
						v85 = base.F64_mul(v83, v84)
						v86 = base.F64_add(v62, v85)
						v91 = base.F64_div(v75, base.F64_add(v75, float64(2)))
						v92 = base.F64_mul(v91, v91)
						v93 = base.F64_mul(v92, v92)
						v118 = base.F64_add(base.F64_mul(v91, base.F64_add(v78, base.F64_add(base.F64_mul(v93, base.F64_add(base.F64_mul(v93, base.F64_add(base.F64_mul(v93, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v92, base.F64_add(base.F64_mul(v93, base.F64_add(base.F64_mul(v93, base.F64_add(base.F64_mul(v93, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v75, v83), v78))
						v130 = base.F64_add(v86, base.F64_add(base.F64_add(v85, base.F64_sub(v62, v86)), base.F64_add(base.F64_mul(v118, v84), base.F64_add(base.F64_mul(v60, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v118, v83), float64(2.5082946711645275e-11))))))
						v152 = v130
					} else {
						v152 = base.F64_div(base.F64_sub(v5, v5), float64(0))
					}
				}
			} else {
				if base.Ui64(int64(9218868437227405311)) < base.Ui64(v20) {
					v130 = v5
					v152 = v130
				} else {
					v35 = int32(-1023)
					v37 = int64(base.Ui64(v20) >> (uint(int64(32)) % 64))
					if v37 != int64(1072693248) {
						v51 = v20
						v53 = v35
						v54 = base.I32_wrap_i64(v37)
						v56 = v54 + int32(614242)
						v60 = base.F64_convert_i32_s(int32(base.Ui32(v56)>>(uint(int32(20))%32)) + v53)
						v62 = base.F64_mul(v60, float64(0.30102999566361177))
						v75 = base.F64_add(base.F64_reinterpret_i64(v51&int64(4294967295)|base.I64_extend_i32_u(v56&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
						v78 = base.F64_mul(v75, base.F64_mul(v75, float64(0.5)))
						v83 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v75, v78)) & int64(-4294967296))
						v84 = float64(0.4342944818781689)
						v85 = base.F64_mul(v83, v84)
						v86 = base.F64_add(v62, v85)
						v91 = base.F64_div(v75, base.F64_add(v75, float64(2)))
						v92 = base.F64_mul(v91, v91)
						v93 = base.F64_mul(v92, v92)
						v118 = base.F64_add(base.F64_mul(v91, base.F64_add(v78, base.F64_add(base.F64_mul(v93, base.F64_add(base.F64_mul(v93, base.F64_add(base.F64_mul(v93, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v92, base.F64_add(base.F64_mul(v93, base.F64_add(base.F64_mul(v93, base.F64_add(base.F64_mul(v93, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v75, v83), v78))
						v130 = base.F64_add(v86, base.F64_add(base.F64_add(v85, base.F64_sub(v62, v86)), base.F64_add(base.F64_mul(v118, v84), base.F64_add(base.F64_mul(v60, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v118, v83), float64(2.5082946711645275e-11))))))
						v152 = v130
					} else {
						if base.I32_wrap_i64(v20) != 0 {
							v51 = v20
							v53 = v35
							v54 = int32(1072693248)
							v56 = v54 + int32(614242)
							v60 = base.F64_convert_i32_s(int32(base.Ui32(v56)>>(uint(int32(20))%32)) + v53)
							v62 = base.F64_mul(v60, float64(0.30102999566361177))
							v75 = base.F64_add(base.F64_reinterpret_i64(v51&int64(4294967295)|base.I64_extend_i32_u(v56&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
							v78 = base.F64_mul(v75, base.F64_mul(v75, float64(0.5)))
							v83 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v75, v78)) & int64(-4294967296))
							v84 = float64(0.4342944818781689)
							v85 = base.F64_mul(v83, v84)
							v86 = base.F64_add(v62, v85)
							v91 = base.F64_div(v75, base.F64_add(v75, float64(2)))
							v92 = base.F64_mul(v91, v91)
							v93 = base.F64_mul(v92, v92)
							v118 = base.F64_add(base.F64_mul(v91, base.F64_add(v78, base.F64_add(base.F64_mul(v93, base.F64_add(base.F64_mul(v93, base.F64_add(base.F64_mul(v93, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v92, base.F64_add(base.F64_mul(v93, base.F64_add(base.F64_mul(v93, base.F64_add(base.F64_mul(v93, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v75, v83), v78))
							v130 = base.F64_add(v86, base.F64_add(base.F64_add(v85, base.F64_sub(v62, v86)), base.F64_add(base.F64_mul(v118, v84), base.F64_add(base.F64_mul(v60, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v118, v83), float64(2.5082946711645275e-11))))))
							v152 = v130
						} else {
							v152 = float64(0)
						}
					}
				}
			}
			v154 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(base.F64_abs(v152), v154)&base.F64_ne(base.F64_abs(v5), v154) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v203 = m.ExcPending
				if v203 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_eq(v152, float64(0))&base.F64_ne(v5, float64(1)) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v205 = m.ExcPending
					if v205 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v165 = F_Float8GetDatum(m, v152)
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int32(0)
					} else {
						return v165
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v173 = m.ExcPending
		if v173 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(352583810))
			mBase = m.M
			v176 = m.ExcPending
			if v176 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(227949), int32(0))
				mBase = m.M
				v180 = m.ExcPending
				if v180 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(470688), int32(1735), int32(530153))
					mBase = m.M
					v185 = m.ExcPending
					if v185 != 0 {
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
func F_doNegate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 != int32(72) {
		v37 = int32(0)
		v40 = F_makeSimpleA_Expr(m, v37, int32(628667), v37, l0, l1)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			v42 = v40
			m.G0 = v6 + int32(16)
			return v42
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l1
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		switch v12 - int32(465) {
		case 0:
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0) - v16
			v42 = l0
			m.G0 = v6 + int32(16)
			return v42
		case 1:
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			v23 = v19 + base.B2i32(v20 == int32(43))
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
			if v24 == int32(45) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23 + int32(1)
				v42 = l0
				m.G0 = v6 + int32(16)
				return v42
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v23
				v32 = F_psprintf(m, int32(167202), v6)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32
					v42 = l0
					m.G0 = v6 + int32(16)
					return v42
				}
			}
		default:
			v37 = int32(0)
			v40 = F_makeSimpleA_Expr(m, v37, int32(628667), v37, l0, l1)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v42 = v40
				m.G0 = v6 + int32(16)
				return v42
			}
		}
	}
}
func F_do_getc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if int32(0) <= v4 {
		if v4 == int32(0) {
			v28 = l0 + int32(76)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
			if v29 != 0 {
				v31 = v29
			} else {
				v31 = int32(1073741823)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v28))) = v31
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v33 != v34 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v33 + int32(1)
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
				v42 = v39
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(0)
				return v42
			} else {
				v40 = F___uflow(m, l0)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = v40
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(0)
					return v42
				}
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[1433]))
			if v10 != v4&int32(1073741823) {
				v28 = l0 + int32(76)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				if v29 != 0 {
					v31 = v29
				} else {
					v31 = int32(1073741823)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v33 != v34 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v33 + int32(1)
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
					v42 = v39
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(0)
					return v42
				} else {
					v40 = F___uflow(m, l0)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v42 = v40
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(0)
						return v42
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v14 != v15 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v14 + int32(1)
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
					return v20
				} else {
					v22 = F___uflow(m, l0)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						return v22
					}
				}
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v14 != v15 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v14 + int32(1)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			return v20
		} else {
			v22 = F___uflow(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				return v22
			}
		}
	}
}
func F_do_start_worker(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
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
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int64
	_ = v201
	var v209 int64
	_ = v209
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v239 int32
	_ = v239
	var v240 int64
	_ = v240
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int64
	_ = v332
	var v333 int64
	_ = v333
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	v1 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v18 = F_LWLockAcquire(m, v14+int32(2816), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[582]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v26 = *(*int32)(unsafe.Add(mBase, _consts[552]))
	v28 = *(*int32)(unsafe.Add(mBase, _consts[551]))
	v30 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v30+int32(2816))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = v28 - v26
	v36 = int32(0)
	if v36 < v35 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v39 = v35
	goto L6
L5:
	;
	v39 = v36
	goto L6
L6:
	;
	if v39 < v24 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v47 = F_AllocSetContextCreateInternal(m, v42, int32(634477), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v381 = v1
	goto L9
L9:
	;
	return v381
L10:
	;
	v49 = int32(4442992)
	v50 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v47
	v53 = F_get_database_list(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v56 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v58 = base.I32_wrap_i64(v56)
	*(*int32)(unsafe.Add(mBase, _consts[614])) = v58
	v61 = *(*int32)(unsafe.Add(mBase, _consts[414]))
	v63 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[615])) = v63
	v66 = F_MultiXactMemberFreezeThreshold(m)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v71 = m.G0
	v72 = int32(16)
	v73 = v71 - v72
	m.G0 = v73
	F___gettimeofday(m, v73)
	mBase = m.M
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v73)))
	v77 = int64(*(*int32)(unsafe.Add(mBase, uint32(v73)+8)))
	m.G0 = v73 + v72
	v85 = v77 + v76*int64(1000000) - int64(946684800000000)
	goto L15
L15:
	;
	if v53 == int32(0) {
		v365 = v1
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v50
	F_MemoryContextDelete(m, v47)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L87
	}
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v88 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v290 != 0 {
		goto L78
	} else {
		goto L79
	}
L19:
	;
	v290 = int32(0)
	v295 = v1
	goto L18
L20:
	;
	goto L21
L21:
	;
	v92 = v58 - v61
	v93 = int32(3)
	if base.Ui32(v92) < base.Ui32(v93) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v97 = v92 - v93
	goto L24
L23:
	;
	v97 = v92
	goto L24
L24:
	;
	if v63 == v66 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v101 = int32(-1)
	goto L27
L26:
	;
	v101 = v63 - v66
	goto L27
L27:
	;
	v102 = int32(0)
	v105 = v102
	v106 = v102
	v107 = v1
	v109 = v1
	v110 = v1
	goto L28
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v117 = int32(2)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116+v107<<(uint(v117)%32))))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	if base.B2i32(base.Ui32(v117) < base.Ui32(v97))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v121)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v290 = v273
	v295 = v278
	goto L18
L30:
	;
	v286 = v107 + int32(1)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v286 < v287 {
		v105 = v273
		v106 = v284
		v107 = v286
		v109 = v277
		v110 = v278
		goto L28
	} else {
		goto L77
	}
L31:
	;
	if v133 != 0 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v133 = base.B2i32(base.Ui32(v121) < base.Ui32(v97))
	goto L31
L33:
	;
	goto L34
L34:
	;
	v133 = int32(base.Ui32(v121-v97) >> (uint(int32(31)) % 32))
	goto L31
L35:
	;
	if v105 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v152 = int32(1)
	if v106&v152 != 0 {
		v273 = v105
		v277 = v109
		v278 = v110
		v284 = v152
		goto L30
	} else {
		goto L45
	}
L38:
	;
	v273 = v120
	v277 = v109
	v278 = v110
	v284 = int32(1)
	goto L30
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v137))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v136)) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v149 != 0 {
		goto L38
	} else {
		goto L44
	}
L41:
	;
	v149 = base.B2i32(base.Ui32(v136) < base.Ui32(v137))
	goto L40
L42:
	;
	goto L43
L43:
	;
	v149 = int32(base.Ui32(v136-v137) >> (uint(int32(31)) % 32))
	goto L40
L44:
	;
	v273 = v105
	v277 = v109
	v278 = v110
	v284 = int32(1)
	goto L30
L45:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	goto L47
L46:
	;
	v273 = v105
	v277 = int32(1)
	v278 = v110
	v284 = int32(0)
	goto L30
L47:
	;
	if int32(base.Ui32(v155-v101)>>(uint(int32(31))%32)) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v105 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	if v109&int32(1) != 0 {
		goto L46
	} else {
		goto L56
	}
L51:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	goto L54
L52:
	;
	goto L53
L53:
	;
	v273 = v120
	v277 = int32(1)
	v278 = v110
	v284 = int32(0)
	goto L30
L54:
	;
	if int32(base.Ui32(v159-v160)>>(uint(int32(31))%32)) == int32(0) {
		goto L46
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v171 = F_pgstat_fetch_stat_dbentry(m, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+16)) = v171
	if v171 == int32(0) {
		v245 = v105
		v250 = v110
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v273 = v258
	v277 = v262
	v278 = v263
	v284 = int32(0)
	goto L30
L59:
	;
	v258 = v245
	v262 = int32(0)
	v263 = v250
	goto L58
L60:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _consts[616]))
	if v177 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if v105 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L62:
	;
	if v177 == int32(4062336) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v185 = v177
	goto L64
L64:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v185-int32(20))))
	if v182 == v197 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L61
L66:
	;
	v200 = v185 - int32(12)
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
	goto L69
L67:
	;
	goto L68
L68:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	if v220 != int32(4062336) {
		v185 = v220
		goto L64
	} else {
		goto L73
	}
L69:
	;
	if base.I64_extend_i32_s(int32(0))*int64(1000) <= v85-v201 {
		goto L61
	} else {
		goto L70
	}
L70:
	;
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
	v211 = *(*int32)(unsafe.Add(mBase, _consts[617]))
	goto L71
L71:
	;
	if base.I64_extend_i32_s(v211*int32(1000))*int64(1000) <= v209-v85 {
		goto L61
	} else {
		goto L72
	}
L72:
	;
	v258 = v105
	v262 = int32(0)
	v263 = int32(1)
	goto L58
L73:
	;
	goto L65
L74:
	;
	v245 = v120
	v250 = int32(0)
	goto L59
L75:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v237)+72))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v239)+72))
	if v238 < v240 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v245 = v105
	v250 = int32(0)
	goto L59
L77:
	;
	goto L29
L78:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v306 = F_LWLockAcquire(m, v302+int32(2816), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v356 = int32(0)
	if v295 == v356 {
		v365 = v356
		goto L16
	} else {
		goto L85
	}
L81:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _consts[582]))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+16))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v311)+4)) = v312
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	*(*int32)(unsafe.Add(mBase, uint32(v312))) = v314
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v309)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v309)+20)) = v316 - int32(1)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	*(*int32)(unsafe.Add(mBase, uint32(v310)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v310)+8)) = v320
	v327 = m.G0
	v328 = int32(16)
	v329 = v327 - v328
	m.G0 = v329
	F___gettimeofday(m, v329)
	mBase = m.M
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v329)))
	v333 = int64(*(*int32)(unsafe.Add(mBase, uint32(v329)+8)))
	m.G0 = v329 + v328
	goto L82
L82:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v310)+24)) = v333 + v332*int64(1000000) - int64(946684800000000)
	v344 = *(*int32)(unsafe.Add(mBase, _consts[582]))
	*(*int32)(unsafe.Add(mBase, uint32(v344)+32)) = v310
	v347 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v347+int32(2816))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_SendPostmasterSignal(m, int32(5))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v365 = v355
	goto L16
L85:
	;
	F_rebuild_database_list(m, int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v365 = v356
	goto L16
L87:
	;
	v381 = v365
	goto L9
}
func F_do_tup_output(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	m.T0[v10].(func(*base.Module, int32))(m, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
		v15 = v8 << (uint(int32(2)) % 32)
		if v15 != 0 {
			v16 = F__emscripten_memcpy_bulkmem(m, v13, l1, v15)
			mBase = m.M
		} else {
		}
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
		if v8 != 0 {
			v19 = F__emscripten_memcpy_bulkmem(m, v18, l2, v8)
			mBase = m.M
		} else {
		}
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
		v23 = v21 & int32(65533)
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)) = uint16(v23)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+6)) = uint16(v26)
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
		v30 = m.T0[v29].(func(*base.Module, int32, int32) int32)(m, v6, v28)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
			m.T0[v33].(func(*base.Module, int32))(m, v6)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_dotrim(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v564 int32
	_ = v564
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v632 int32
	_ = v632
	var v640 int32
	_ = v640
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	v7 = int32(0)
	if l1 <= v7 {
		v612 = l0
		v613 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v672 = F_cstring_to_text_with_len(m, v659, v660)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L13
	} else {
		goto L119
	}
L2:
	;
	v659 = v632
	v660 = v640
	goto L1
L3:
	;
	v632 = v612
	v640 = v613
	goto L2
L4:
	;
	if l3 <= int32(0) {
		v612 = l0
		v613 = l1
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27*int32(28))+uint32(_consts[355])))
	goto L9
L6:
	;
	F_pfree(m, v38)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L13
	} else {
		goto L114
	}
L7:
	;
	if l5 == int32(0) {
		v585 = v426
		v598 = v439
		goto L6
	} else {
		goto L81
	}
L8:
	;
	v426 = l1
	v432 = v73
	v439 = l0
	v440 = int32(-1)
	goto L7
L9:
	;
	if int32(2) <= v32 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = l1 << (uint(int32(2)) % 32)
	v38 = F_palloc(m, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	if l4 == int32(0) {
		v318 = l0
		v319 = l1
		goto L58
	} else {
		goto L59
	}
L13:
	;
	return int32(0)
L14:
	;
	v42 = F_palloc(m, v37)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v50 = l0
	v51 = l1
	v52 = v7
	goto L16
L16:
	;
	v65 = v52 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v38+v65))) = v50
	v69 = F_pg_mblen_range(m, v50, l0+l1)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v80 = l3 << (uint(int32(2)) % 32)
	v81 = F_palloc(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L13
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v42))) = v69
	v73 = v52 + int32(1)
	v75 = v51 - v69
	if int32(0) < v75 {
		v50 = v50 + v69
		v51 = v75
		v52 = v73
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v83 = F_palloc(m, v80)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v88 = l2
	v89 = l3
	v93 = int32(0)
	goto L22
L22:
	;
	v107 = v93 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v81+v107))) = v88
	v111 = F_pg_mblen_range(m, v88, l2+l3)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L13
	} else {
		goto L24
	}
L23:
	;
	if l4 == int32(0) {
		goto L8
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107+v83))) = v111
	v117 = v89 - v111
	if int32(0) < v117 {
		v88 = v88 + v111
		v89 = v117
		v93 = v93 + int32(1)
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	if base.Ui32(int32(2147483646)) < base.Ui32(v52) {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v131 = l1
	v137 = v73
	v140 = v7
	v144 = l0
	goto L28
L28:
	;
	if base.Ui32(int32(2147483646)) < base.Ui32(v93) {
		goto L8
	} else {
		goto L30
	}
L29:
	;
	v585 = v254
	v598 = v255
	goto L6
L30:
	;
	v148 = v140 << (uint(int32(2)) % 32)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v42+v148)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148+v38)))
	v156 = int32(0)
	goto L32
L31:
	;
	v252 = int32(1)
	v254 = v131 - v150
	v255 = v150 + v144
	if base.B2i32(v52 == v140) == int32(0) {
		v131 = v254
		v137 = v137 - v252
		v140 = v140 + v252
		v144 = v255
		goto L28
	} else {
		goto L57
	}
L32:
	;
	v175 = v156 << (uint(int32(2)) % 32)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v83+v175)))
	if v177 == v150 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v426 = v131
	v432 = v137
	v439 = v144
	v440 = v140 - int32(1)
	goto L7
L34:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v175+v81)))
	if base.Ui32(int32(4)) <= base.Ui32(v150) {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	goto L36
L36:
	;
	if base.B2i32(v156 == v93) == int32(0) {
		v156 = v156 + int32(1)
		goto L32
	} else {
		goto L56
	}
L37:
	;
	if v242 == int32(0) {
		goto L31
	} else {
		goto L55
	}
L38:
	;
	v242 = int32(0)
	goto L37
L39:
	;
	v216 = v211
	v217 = v212
	v218 = v213
	goto L49
L40:
	;
	if (v152|v180)&int32(3) != 0 {
		v211 = v152
		v212 = v180
		v213 = v150
		goto L39
	} else {
		goto L43
	}
L41:
	;
	v204 = v152
	v205 = v180
	v206 = v150
	goto L42
L42:
	;
	if v206 == int32(0) {
		goto L38
	} else {
		goto L48
	}
L43:
	;
	v188 = v152
	v189 = v180
	v190 = v150
	goto L44
L44:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	if v193 != v194 {
		v211 = v188
		v212 = v189
		v213 = v190
		goto L39
	} else {
		goto L46
	}
L45:
	;
	v204 = v199
	v205 = v197
	v206 = v201
	goto L42
L46:
	;
	v196 = int32(4)
	v197 = v189 + v196
	v199 = v188 + v196
	v201 = v190 - v196
	if base.Ui32(int32(3)) < base.Ui32(v201) {
		v188 = v199
		v189 = v197
		v190 = v201
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v211 = v204
	v212 = v205
	v213 = v206
	goto L39
L49:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v221 == v222 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v242 = v221 - v222
	goto L37
L51:
	;
	v224 = int32(1)
	v229 = v218 - v224
	if v229 != 0 {
		v216 = v216 + v224
		v217 = v217 + v224
		v218 = v229
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	goto L50
L54:
	;
	goto L38
L55:
	;
	goto L36
L56:
	;
	goto L33
L57:
	;
	goto L29
L58:
	;
	if l5 == int32(0) {
		v612 = v318
		v613 = v319
		goto L3
	} else {
		goto L68
	}
L59:
	;
	v264 = l0
	v265 = l1
	goto L60
L60:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	v292 = int32(0)
	goto L63
L61:
	;
	v659 = l0 + l1
	v660 = v7
	goto L1
L62:
	;
	v312 = int32(1)
	if int32(2) <= v265 {
		v264 = v264 + v312
		v265 = v265 - v312
		goto L60
	} else {
		goto L67
	}
L63:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v292))))
	if v285 == v307 {
		goto L62
	} else {
		goto L65
	}
L64:
	;
	v318 = v264
	v319 = v265
	goto L58
L65:
	;
	v310 = v292 + int32(1)
	if v310 != l3 {
		v292 = v310
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	goto L61
L68:
	;
	v340 = int32(1)
	if l3 <= v340 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v343 = v340
	goto L71
L70:
	;
	v343 = l3
	goto L71
L71:
	;
	v347 = v319
	goto L72
L72:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318-int32(1)+v347))))
	v375 = int32(0)
	goto L74
L73:
	;
	v632 = v318
	v640 = v347
	goto L2
L74:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v375))))
	if v390 == v368 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L73
L76:
	;
	v396 = int32(0)
	if base.B2i32(v347 < int32(2)) == v396 {
		v347 = v347 - int32(1)
		goto L72
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v400 = v375 + int32(1)
	if v400 != v343 {
		v375 = v400
		goto L74
	} else {
		goto L80
	}
L79:
	;
	v632 = v318
	v640 = v396
	goto L2
L80:
	;
	goto L75
L81:
	;
	if v432 <= int32(0) {
		v585 = v426
		v598 = v439
		goto L6
	} else {
		goto L82
	}
L82:
	;
	v455 = v426
	v458 = v432
	goto L83
L83:
	;
	if base.Ui32(int32(2147483646)) < base.Ui32(v93) {
		v585 = v426
		v598 = v439
		goto L6
	} else {
		goto L85
	}
L84:
	;
	v585 = v579
	v598 = v439
	goto L6
L85:
	;
	v470 = (v458 + v440) << (uint(int32(2)) % 32)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v42+v470)))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v470+v38)))
	v478 = int32(0)
	goto L87
L86:
	;
	goto L84
L87:
	;
	v497 = v478 << (uint(int32(2)) % 32)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v83+v497)))
	if v499 == v472 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v572 = v455 - v472
	v573 = int32(1)
	if v573 < v458 {
		v455 = v572
		v458 = v458 - v573
		goto L83
	} else {
		goto L113
	}
L89:
	;
	goto L88
L90:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v497+v81)))
	if base.Ui32(int32(4)) <= base.Ui32(v472) {
		goto L96
	} else {
		goto L97
	}
L91:
	;
	goto L92
L92:
	;
	if base.B2i32(v478 == v93) == int32(0) {
		v478 = v478 + int32(1)
		goto L87
	} else {
		goto L112
	}
L93:
	;
	if v564 == int32(0) {
		goto L89
	} else {
		goto L111
	}
L94:
	;
	v564 = int32(0)
	goto L93
L95:
	;
	v538 = v533
	v539 = v534
	v540 = v535
	goto L105
L96:
	;
	if (v474|v502)&int32(3) != 0 {
		v533 = v474
		v534 = v502
		v535 = v472
		goto L95
	} else {
		goto L99
	}
L97:
	;
	v526 = v474
	v527 = v502
	v528 = v472
	goto L98
L98:
	;
	if v528 == int32(0) {
		goto L94
	} else {
		goto L104
	}
L99:
	;
	v510 = v474
	v511 = v502
	v512 = v472
	goto L100
L100:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v510)))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v511)))
	if v515 != v516 {
		v533 = v510
		v534 = v511
		v535 = v512
		goto L95
	} else {
		goto L102
	}
L101:
	;
	v526 = v521
	v527 = v519
	v528 = v523
	goto L98
L102:
	;
	v518 = int32(4)
	v519 = v511 + v518
	v521 = v510 + v518
	v523 = v512 - v518
	if base.Ui32(int32(3)) < base.Ui32(v523) {
		v510 = v521
		v511 = v519
		v512 = v523
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v533 = v526
	v534 = v527
	v535 = v528
	goto L95
L105:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538))))
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539))))
	if v543 == v544 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v564 = v543 - v544
	goto L93
L107:
	;
	v546 = int32(1)
	v551 = v540 - v546
	if v551 != 0 {
		v538 = v538 + v546
		v539 = v539 + v546
		v540 = v551
		goto L105
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	goto L106
L110:
	;
	goto L94
L111:
	;
	goto L92
L112:
	;
	v579 = v455
	goto L86
L113:
	;
	v579 = v572
	goto L86
L114:
	;
	F_pfree(m, v42)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L13
	} else {
		goto L115
	}
L115:
	;
	F_pfree(m, v81)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L13
	} else {
		goto L116
	}
L116:
	;
	F_pfree(m, v83)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L13
	} else {
		goto L117
	}
L117:
	;
	v609 = F_cstring_to_text_with_len(m, v598, v585)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L13
	} else {
		goto L118
	}
L118:
	;
	return v609
L119:
	;
	return v672
}
func F_downcase_truncate_identifier(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_downcase_identifier(m, l0, l1, l2, int32(1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_drandom_normal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v99 int64
	_ = v99
	var v102 int64
	_ = v102
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v122 int64
	_ = v122
	var v127 int64
	_ = v127
	var v132 int64
	_ = v132
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int64
	_ = v149
	var v150 int64
	_ = v150
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v161 int64
	_ = v161
	var v183 float64
	_ = v183
	var v195 float64
	_ = v195
	var v197 float64
	_ = v197
	var v204 float64
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v224 float64
	_ = v224
	var v228 int32
	_ = v228
	var v229 float64
	_ = v229
	var v230 float64
	_ = v230
	var v236 float64
	_ = v236
	var v237 float64
	_ = v237
	var v239 float64
	_ = v239
	var v241 float64
	_ = v241
	var v243 float64
	_ = v243
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1120])))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = int32(16)
	v17 = int32(0)
	v21 = m.G0
	v23 = v21 - v16
	m.G0 = v23
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v17
	v29 = F_open(m, int32(274295), v17, v23)
	mBase = m.M
	if v29 != int32(-1) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L3
L3:
	;
	v147 = int32(4427848)
	v148 = int32(4427840)
	v149 = *(*int64)(unsafe.Add(mBase, _consts[1121]))
	v150 = int64(24)
	v153 = *(*int64)(unsafe.Add(mBase, _consts[1122]))
	v154 = v153 ^ v149
	v155 = int64(16)
	v158 = base.I64_rotl(v149, v150) ^ v154<<(uint(v155)%64) ^ v154
	v159 = int64(37)
	v161 = v158 ^ base.I64_rotl(v154, v159)
	*(*int64)(unsafe.Add(mBase, _consts[1122])) = base.I64_rotl(v161, v159)
	*(*int64)(unsafe.Add(mBase, _consts[1121])) = v161<<(uint(v155)%64) ^ base.I64_rotl(v158, v150) ^ v161
	v183 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v158*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L31
L4:
	;
	v145 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1120])) = uint8(v145)
	goto L3
L5:
	;
	if v62 != 0 {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	goto L10
L7:
	;
	v62 = v17
	goto L8
L8:
	;
	m.G0 = v23 + int32(16)
	goto L5
L9:
	;
	v57 = F_close(m, v29)
	mBase = m.M
	v62 = v55
	goto L8
L10:
	;
	v35 = int32(4427840)
	v36 = v16
	goto L11
L11:
	;
	v41 = F_read(m, v29, v35, v36)
	mBase = m.M
	if v41 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v55 = int32(1)
	goto L9
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v45 == int32(27) {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v50 = v36 - v41
	if v50 != 0 {
		v35 = v35 + v41
		v36 = v50
		goto L11
	} else {
		goto L17
	}
L16:
	;
	v55 = int32(0)
	goto L9
L17:
	;
	goto L12
L18:
	;
	v67 = int32(4427840)
	v68 = *(*int64)(unsafe.Add(mBase, _consts[1121]))
	if v68 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v79 = int32(4427840)
	v83 = m.G0
	v84 = int32(16)
	v85 = v83 - v84
	m.G0 = v85
	F___gettimeofday(m, v85)
	mBase = m.M
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
	v89 = int64(*(*int32)(unsafe.Add(mBase, uint32(v85)+8)))
	m.G0 = v85 + v84
	goto L26
L21:
	;
	goto L4
L22:
	;
	goto L21
L23:
	;
	v71 = *(*int64)(unsafe.Add(mBase, _consts[1122]))
	if v71 != int64(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1122])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _consts[1121])) = int64(6364136223846793005)
	goto L22
L26:
	;
	v99 = int64(*(*uint32)(unsafe.Add(mBase, _consts[156])))
	v102 = v89 + v88*int64(1000000) - int64(946684800000000) ^ v99<<(uint(int64(32))%64)
	v106 = v102 + int64(4354685564936845354)
	v107 = int64(30)
	v110 = int64(-4658895280553007687)
	v111 = (int64(base.Ui64(v106)>>(uint(v107)%64)) ^ v106) * v110
	v112 = int64(27)
	v115 = int64(-7723592293110705685)
	v116 = (int64(base.Ui64(v111)>>(uint(v112)%64)) ^ v111) * v115
	v117 = int64(31)
	*(*int64)(unsafe.Add(mBase, _consts[1122])) = int64(base.Ui64(v116)>>(uint(v117)%64)) ^ v116
	v122 = v102 - int64(7046029254386353131)
	v127 = (int64(base.Ui64(v122)>>(uint(v107)%64)) ^ v122) * v110
	v132 = (int64(base.Ui64(v127)>>(uint(v112)%64)) ^ v127) * v115
	*(*int64)(unsafe.Add(mBase, _consts[1121])) = int64(base.Ui64(v132)>>(uint(v117)%64)) ^ v132
	if v122|v106 == int64(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L4
L28:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1122])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _consts[1121])) = int64(6364136223846793005)
	goto L30
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	v195 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v149*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L32
L32:
	;
	v197 = F_log(m, base.F64_sub(float64(1), v195))
	mBase = m.M
	v204 = base.F64_mul(base.F64_sub(float64(1), v183), float64(6.283185307179586))
	v208 = m.G0
	v210 = v208 - int32(16)
	m.G0 = v210
	v217 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v204))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v217) <= base.Ui32(int32(1072243195)) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v252 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(v8, base.F64_mul(base.F64_sqrt(base.F64_mul(v197, float64(-2))), v243)), v10))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L46
	} else {
		goto L47
	}
L34:
	;
	m.G0 = v210 + int32(16)
	goto L33
L35:
	;
	if base.Ui32(v217) < base.Ui32(int32(1045430272)) {
		v243 = v204
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v217) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v224 = F___sin(m, v204, float64(0), int32(0))
	mBase = m.M
	v243 = v224
	goto L34
L39:
	;
	v243 = base.F64_sub(v204, v204)
	goto L34
L40:
	;
	goto L41
L41:
	;
	v228 = F___rem_pio2(m, v204, v210)
	mBase = m.M
	v229 = *(*float64)(unsafe.Add(mBase, uint32(v210)+8))
	v230 = *(*float64)(unsafe.Add(mBase, uint32(v210)))
	switch v228&int32(3) - int32(1) {
	case 0:
		goto L44
	case 1:
		goto L43
	case 2:
		goto L42
	default:
		goto L45
	}
L42:
	;
	v241 = F___cos(m, v230, v229)
	mBase = m.M
	v243 = base.F64_neg(v241)
	goto L34
L43:
	;
	v239 = F___sin(m, v230, v229, int32(1))
	mBase = m.M
	v243 = base.F64_neg(v239)
	goto L34
L44:
	;
	v237 = F___cos(m, v230, v229)
	mBase = m.M
	v243 = v237
	goto L34
L45:
	;
	v236 = F___sin(m, v230, v229, int32(1))
	mBase = m.M
	v243 = v236
	goto L34
L46:
	;
	return int32(0)
L47:
	;
	return v252
}
func F_dsign(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v10 float64
	_ = v10
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v2 = float64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.F64_lt(v7, v2) != 0 {
		v10 = float64(-1)
	} else {
		v10 = v2
	}
	if base.F64_gt(v7, float64(0)) != 0 {
		v13 = float64(1)
	} else {
		v13 = v10
	}
	v14 = F_Float8GetDatum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		return v14
	}
}
func F_dsnowball_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_palloc0(m, int32(24))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v16 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v372 != 0 {
		goto L100
	} else {
		goto L101
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v24 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v35 = v2
	v38 = v2
	goto L6
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v35<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v46 = int32(161994)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1284])))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v50 == int32(0) {
		v69 = v49
		v70 = v50
		goto L13
	} else {
		goto L14
	}
L7:
	;
	goto L3
L8:
	;
	v358 = v35 + int32(1)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v358 < v359 {
		v35 = v358
		v38 = v355
		goto L6
	} else {
		goto L99
	}
L9:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v330)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v339
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v330)+8))
	v342 = m.T0[v341].(func(*base.Module) int32)(m)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L98
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L94
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L90
	}
L12:
	;
	if v70-v69 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	goto L12
L14:
	;
	if v49 != v50 {
		v69 = v49
		v70 = v50
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v54 = v45
	v55 = v46
	goto L16
L16:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if v59 == int32(0) {
		v69 = v58
		v70 = v59
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v69 = v58
	v70 = v59
	goto L13
L18:
	;
	v62 = int32(1)
	if v58 == v59 {
		v54 = v54 + v62
		v55 = v55 + v62
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	if v38 != 0 {
		goto L11
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v80 = int32(384472)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1285])))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v84 == int32(0) {
		v103 = v83
		v104 = v84
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v75 = F_defGetString(m, v44)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_readstoplist(m, v75, v18+int32(4), int32(1162))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v355 = int32(1)
	goto L8
L26:
	;
	if v104-v103 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	goto L26
L28:
	;
	if v83 != v84 {
		v103 = v83
		v104 = v84
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v88 = v45
	v89 = v80
	goto L30
L30:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	if v93 == int32(0) {
		v103 = v92
		v104 = v93
		goto L27
	} else {
		goto L32
	}
L31:
	;
	v103 = v92
	v104 = v93
	goto L27
L32:
	;
	v96 = int32(1)
	if v92 == v93 {
		v88 = v88 + v96
		v89 = v89 + v96
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v108 != 0 {
		goto L10
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L86
	}
L37:
	;
	v114 = F_defGetString(m, v44)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v118 = int32(4117344)
	v123 = int32(327537)
	goto L39
L39:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v127 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v182 = int32(4117344)
	v183 = int32(327537)
	goto L62
L41:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
	if v179 != 0 {
		v118 = v118 + int32(20)
		v123 = v179
		goto L39
	} else {
		goto L61
	}
L42:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	goto L45
L43:
	;
	goto L44
L44:
	;
	v134 = v123
	v135 = v114
	goto L48
L45:
	;
	if v130 != v127 {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	if v172 != 0 {
		goto L41
	} else {
		goto L60
	}
L48:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v138 == v139 {
		v161 = v138
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v172 = int32(0)
	goto L47
L50:
	;
	v163 = int32(1)
	if v161 != 0 {
		v134 = v134 + v163
		v135 = v135 + v163
		goto L48
	} else {
		goto L59
	}
L51:
	;
	if base.Ui32((v138-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v149 = v138 | int32(32)
	goto L54
L53:
	;
	v149 = v138
	goto L54
L54:
	;
	if base.Ui32((v139-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v158 = v139 | int32(32)
	goto L57
L56:
	;
	v158 = v139
	goto L57
L57:
	;
	if v149 == v158 {
		v161 = v149
		goto L50
	} else {
		goto L58
	}
L58:
	;
	v172 = v149 - v158
	goto L47
L59:
	;
	goto L49
L60:
	;
	v329 = int32(0)
	v330 = v118
	goto L9
L61:
	;
	goto L40
L62:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if v193 != int32(6) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L81
	}
L64:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
	if v238 != 0 {
		v182 = v182 + int32(20)
		v183 = v238
		goto L62
	} else {
		goto L80
	}
L65:
	;
	v198 = v183
	v199 = v114
	goto L67
L66:
	;
	if v236 != 0 {
		goto L64
	} else {
		goto L79
	}
L67:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	if v202 == v203 {
		v225 = v202
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v236 = int32(0)
	goto L66
L69:
	;
	v227 = int32(1)
	if v225 != 0 {
		v198 = v198 + v227
		v199 = v199 + v227
		goto L67
	} else {
		goto L78
	}
L70:
	;
	if base.Ui32((v202-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v213 = v202 | int32(32)
	goto L73
L72:
	;
	v213 = v202
	goto L73
L73:
	;
	if base.Ui32((v203-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v222 = v203 | int32(32)
	goto L76
L75:
	;
	v222 = v203
	goto L76
L76:
	;
	if v213 == v222 {
		v225 = v213
		goto L69
	} else {
		goto L77
	}
L77:
	;
	v236 = v213 - v222
	goto L66
L78:
	;
	goto L68
L79:
	;
	v329 = int32(1)
	v330 = v182
	goto L9
L80:
	;
	goto L63
L81:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v114
	F_errmsg(m, int32(671212), v14)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(474307), int32(221), int32(364078))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v272
	F_errmsg(m, int32(683967), v14+int32(16))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(474307), int32(260), int32(93987))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errmsg(m, int32(123225), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(474307), int32(243), int32(93987))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(123357), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(474307), int32(252), int32(93987))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+12)) = uint8(v329)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v342
	v355 = v38
	goto L8
L99:
	;
	goto L7
L100:
	;
	v374 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v374
	m.G0 = v14 + int32(32)
	return v18
L101:
	;
	goto L102
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errmsg(m, int32(206169), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(474307), int32(267), int32(93987))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_durable_unlink(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	v5 = m.G0
	v7 = v5 - int32(1040)
	m.G0 = v7
	v9 = int32(-1)
	v10 = F_unlink(m, l0)
	mBase = m.M
	if v10 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(1040)
	return v213
L2:
	;
	v14 = F_errstart(m, l1, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v32 = v7 + int32(16)
	goto L14
L5:
	;
	return int32(0)
L6:
	;
	if v14 == int32(0) {
		v213 = v9
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg(m, int32(284570), v7)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(476611), int32(884), int32(300201))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v213 = v9
	goto L1
L11:
	;
	v149 = v7 + int32(16)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v153 != 0 {
		goto L44
	} else {
		goto L45
	}
L12:
	;
	v145 = F_strlen(m, v134)
	mBase = m.M
	goto L11
L14:
	;
	goto L15
L15:
	;
	v39 = int32(1023)
	if (v32^l0)&int32(3) != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v138 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v138)
	goto L12
L17:
	;
	v119 = v114
	v120 = v115
	v121 = v116
	goto L39
L18:
	;
	if v109 == int32(0) {
		v134 = v107
		v135 = v108
		goto L16
	} else {
		goto L38
	}
L19:
	;
	v107 = l0
	v108 = v32
	v109 = v39
	goto L18
L20:
	;
	goto L21
L21:
	;
	if l0&int32(3) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v76 == int32(0) {
		v134 = v73
		v135 = v74
		goto L16
	} else {
		goto L31
	}
L23:
	;
	v73 = l0
	v74 = v32
	v75 = v39
	v76 = int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v52 = l0
	v53 = v32
	v54 = v39
	goto L26
L26:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v56)
	if v56 == int32(0) {
		v114 = v52
		v115 = v53
		v116 = v54
		goto L17
	} else {
		goto L28
	}
L27:
	;
	v73 = v67
	v74 = v61
	v75 = v63
	v76 = v65
	goto L22
L28:
	;
	v60 = int32(1)
	v61 = v53 + v60
	v63 = v54 - v60
	v64 = int32(0)
	v65 = base.B2i32(v63 != v64)
	v67 = v52 + v60
	if v67&int32(3) == v64 {
		v73 = v67
		v74 = v61
		v75 = v63
		v76 = v65
		goto L22
	} else {
		goto L29
	}
L29:
	;
	if v63 != 0 {
		v52 = v67
		v53 = v61
		v54 = v63
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v79 == int32(0) {
		v107 = v73
		v108 = v74
		v109 = v75
		goto L18
	} else {
		goto L32
	}
L32:
	;
	if base.Ui32(v75) < base.Ui32(int32(4)) {
		v107 = v73
		v108 = v74
		v109 = v75
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v85 = v73
	v86 = v74
	v87 = v75
	goto L34
L34:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v93 = int32(-2139062144)
	if (int32(16843008)-v90|v90)&v93 != v93 {
		v114 = v85
		v115 = v86
		v116 = v87
		goto L17
	} else {
		goto L36
	}
L35:
	;
	v107 = v101
	v108 = v99
	v109 = v103
	goto L18
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v90
	v98 = int32(4)
	v99 = v86 + v98
	v101 = v85 + v98
	v103 = v87 - v98
	if base.Ui32(int32(3)) < base.Ui32(v103) {
		v85 = v101
		v86 = v99
		v87 = v103
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v114 = v107
	v115 = v108
	v116 = v109
	goto L17
L39:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v123)
	if v123 == int32(0) {
		v134 = v119
		v135 = v120
		goto L16
	} else {
		goto L41
	}
L40:
	;
	v134 = v130
	v135 = v128
	goto L16
L41:
	;
	v127 = int32(1)
	v128 = v120 + v127
	v130 = v119 + v127
	v132 = v121 - v127
	if v132 != 0 {
		v119 = v130
		v120 = v128
		v121 = v132
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)))
	if v199 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L44:
	;
	v154 = F_strlen(m, v149)
	mBase = m.M
	v157 = v154 + v149
	goto L47
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	v161 = v157 - int32(1)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v162 == int32(47) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v167 = v161
	goto L53
L49:
	;
	if base.Ui32(v149) < base.Ui32(v161) {
		v157 = v161
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L51
L53:
	;
	if base.Ui32(v149) < base.Ui32(v167) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v179 = v167
	goto L59
L55:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v173 != int32(47) {
		v167 = v167 - int32(1)
		goto L53
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	goto L54
L58:
	;
	goto L57
L59:
	;
	if base.Ui32(v149) < base.Ui32(v179) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v149 == v179 {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	v183 = v179 - int32(1)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if v184 == int32(47) {
		v179 = v183
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	goto L63
L65:
	;
	v192 = v149 + base.B2i32(v153 == int32(47))
	goto L67
L66:
	;
	v192 = v179
	goto L67
L67:
	;
	v193 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v192))) = uint8(v193)
	goto L46
L68:
	;
	v202 = int32(46)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)) = uint16(v202)
	goto L70
L69:
	;
	goto L70
L70:
	;
	v205 = int32(0)
	v210 = F_fsync_fname_ext(m, v7+int32(16), int32(1), v205, l1)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	if v210 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v212 = int32(-1)
	goto L74
L73:
	;
	v212 = v205
	goto L74
L74:
	;
	v213 = v212
	goto L1
}
