package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_vacuum_buffer_usage_limit(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = int32(0)
	v15 = base.B2i32(v8 == v9) | base.B2i32(base.Ui32(v8-int32(128)) < base.Ui32(int32(16777089)))
	if v15 == v9 {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[87]))
		*(*int32)(unsafe.Add(mBase, _consts[88])) = v19
		*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = int64(72057594037928064)
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(103162)
		v28 = F_format_elog_string(m, int32(673622), v6)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[89])) = v28
			m.G0 = v6 + int32(16)
			return v15
		}
	} else {
		m.G0 = v6 + int32(16)
		return v15
	}
}
func F_vacuumRedirectAndPlaceholder(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
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
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v348 int64
	_ = v348
	var v349 int32
	_ = v349
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(1648)
	m.G0 = v22
	if l2 < v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v42) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27+(l2^int32(-1))<<(uint(int32(2))%32))))
	v41 = v33
	goto L1
L3:
	;
	goto L4
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v41 = v35 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	v50 = int32(base.Ui32(v42+int32(262120)) >> (uint(int32(2)) % 32))
	goto L7
L6:
	;
	v50 = int32(0)
	goto L7
L7:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+16)))
	v52 = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v54 < int32(2) {
		v74 = v52
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v76 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v76
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)) = uint16(v76)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+12)) = uint8(v74)
	v81 = F_GlobalVisHorizonKindForRel(m, l1)
	mBase = m.M
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81<<(uint(int32(2))%32))+uint32(_consts[66])))
	goto L15
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+118)))
	if v58 != int32(112) {
		v74 = v52
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	goto L11
L11:
	;
	if base.Ui32(v62) < base.Ui32(int32(12000)) {
		v74 = int32(1)
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v65 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	if v66 == v65 {
		v74 = v65
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+119)))
	switch v70 - int32(109) {
	case 0, 5:
		goto L14
	default:
		v74 = v65
		goto L8
	}
L14:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+104)))
	v74 = v73
	goto L8
L15:
	;
	v87 = int32(4530932)
	v89 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v89 + int32(1)
	v94 = v50 & int32(65535)
	if v94 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v372 = int32(4530932)
	v374 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v374 - int32(1)
	m.G0 = v22 + int32(1648)
	return
L17:
	;
	v97 = v41 + v51
	v109 = v4
	v111 = int32(0)
	v112 = v4
	v113 = v4
	v115 = v94
	v116 = v50
	v117 = v4
	goto L18
L18:
	;
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+2)))
	if base.B2i32(v120 == int32(0))&(v111&int32(1)) != 0 {
		v213 = v109
		v220 = v117
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v222 = v213 & int32(65535)
	if v222 != 0 {
		goto L44
	} else {
		goto L45
	}
L20:
	;
	goto L19
L21:
	;
	v126 = int32(1)
	v127 = v115 - v126
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(24)+v127<<(uint(int32(2))%32))))
	v134 = v41 + v131&int32(32767)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	if v135&int32(3) != v126 {
		v194 = v135
		v195 = v112
		v196 = v113
		v198 = v117
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v199 = int32(3)
	v203 = base.B2i32(v194&v199 != v199) | v111
	if v203&int32(1) != 0 {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	if v140 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v141 = F_GlobalVisTestIsRemovableXid(m, v86, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v146 = v135
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v146 | int32(3)
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+2)))
	v152 = int32(1)
	v153 = v151 - v152
	*(*uint16)(unsafe.Add(mBase, uint32(v97)+2)) = uint16(v153)
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+4)))
	v157 = v155 + v152
	*(*uint16)(unsafe.Add(mBase, uint32(v97)+4)) = uint16(v157)
	if v113 != 0 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	return
L28:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	if v141 == int32(0) {
		v194 = v143
		v195 = v112
		v196 = v113
		v198 = v117
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v146 = v143
	goto L26
L30:
	;
	v177 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+10)) = uint16(v177)
	*(*int32)(unsafe.Add(mBase, uint32(v134)+6)) = int32(-1)
	v181 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(832)+v112&int32(65535)<<(uint(v181)%32)))) = uint16(v115)
	v191 = v112 + v181
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)) = uint16(v191)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v194 = v193
	v195 = v191
	v196 = v176
	v198 = v181
	goto L22
L31:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v159))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v113)) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v174
	v176 = v174
	goto L30
L34:
	;
	if v171 == int32(0) {
		v176 = v113
		goto L30
	} else {
		goto L38
	}
L35:
	;
	v171 = base.B2i32(base.Ui32(v113) < base.Ui32(v159))
	goto L34
L36:
	;
	goto L37
L37:
	;
	v171 = int32(base.Ui32(v113-v159) >> (uint(int32(31)) % 32))
	goto L34
L38:
	;
	goto L33
L39:
	;
	v206 = v109
	goto L41
L40:
	;
	v206 = v115
	goto L41
L41:
	;
	v208 = v116 - int32(1)
	if v208&int32(65535) != 0 {
		v109 = v206
		v111 = v203
		v112 = v195
		v113 = v196
		v115 = v127
		v116 = v208
		v117 = v198
		goto L18
	} else {
		goto L42
	}
L42:
	;
	v213 = v206
	v220 = v198
	goto L20
L43:
	;
	F_MarkBufferDirty(m, l2)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L27
	} else {
		goto L55
	}
L44:
	;
	if base.Ui32(v222) <= base.Ui32(v50&int32(65535)) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v293 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+6)) = uint16(v293)
	if v220 == v293 {
		goto L16
	} else {
		goto L54
	}
L47:
	;
	v229 = v213
	goto L50
L48:
	;
	goto L49
L49:
	;
	v280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+4)))
	v283 = v50 - v213 + int32(1)
	v284 = v280 - v283
	*(*uint16)(unsafe.Add(mBase, uint32(v97)+4)) = uint16(v284)
	F_PageIndexMultiDelete(m, v41, v22+int32(16), v283&int32(65535))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L27
	} else {
		goto L53
	}
L50:
	;
	v247 = int32(65535)
	v250 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(16)+(v229&v247-v222)<<(uint(v250)%32)))) = uint16(v229)
	v255 = v229 + v250
	if base.Ui32(v255&v247) <= base.Ui32(v50&v247) {
		v229 = v255
		goto L50
	} else {
		goto L52
	}
L51:
	;
	goto L49
L52:
	;
	goto L51
L53:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+6)) = uint16(v213)
	goto L43
L54:
	;
	goto L43
L55:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+118)))
	if v319 != int32(112) {
		goto L16
	} else {
		goto L56
	}
L56:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v323 <= int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v326 != 0 {
		goto L16
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L27
	} else {
		goto L62
	}
L60:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v327 != 0 {
		goto L16
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	F_XLogRegisterData(m, v22+int32(4), int32(10))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L27
	} else {
		goto L63
	}
L63:
	;
	v337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
	F_XLogRegisterData(m, v22+int32(832), v337<<(uint(int32(1))%32))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L27
	} else {
		goto L64
	}
L64:
	;
	F_XLogRegisterBuffer(m, int32(0), l2, int32(8))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L27
	} else {
		goto L65
	}
L65:
	;
	v348 = F_XLogInsert(m, int32(16), int32(128))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L27
	} else {
		goto L66
	}
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = base.I64_rotr(v348, int64(32))
	goto L16
}
