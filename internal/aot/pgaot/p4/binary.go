package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BinarySearchRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	v12 = l1
	v14 = int32(base.Ui32(l1) >> (uint(int32(1)) % 32))
	v15 = int32(0)
	goto L2
L1:
	;
	return v130 & int32(65535)
L2:
	;
	v21 = l0 + v14<<(uint(int32(2))%32)
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21))))
	v23 = base.B2i32(base.Ui32(l2) < base.Ui32(v22))
	if base.Ui32(l2) < base.Ui32(v22) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v130 = int32(0)
	goto L1
L4:
	;
	goto L3
L5:
	;
	if base.Ui32(l2) < base.Ui32(v22) {
		goto L27
	} else {
		goto L28
	}
L6:
	;
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	if base.Ui32(v24) <= base.Ui32(l2) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+2)))
	if v26 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v31 = l2 - v22&int32(65280)
	if base.Ui32(int32(41280)) <= base.Ui32(l2) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v34 = int32(255)
	v35 = l2 & v34
	v37 = v22 & v34
	v47 = base.B2i32(base.Ui32(int32(160)) < base.Ui32(v37))
	if base.Ui32(int32(160)) < base.Ui32(v37) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v74 = int32(255)
	v85 = v26 & v74
	if base.Ui32(int32(160)) < base.Ui32(v85) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	v48 = int32(0)
	goto L14
L13:
	;
	v48 = int32(-34)
	goto L14
L14:
	;
	if base.Ui32(int32(160)) < base.Ui32(v37) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v51 = int32(34)
	goto L17
L16:
	;
	v51 = int32(0)
	goto L17
L17:
	;
	if base.Ui32(int32(160)) < base.Ui32(v35) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v54 = v48
	goto L20
L19:
	;
	v54 = v51
	goto L20
L20:
	;
	v59 = int32(33)
	v60 = v35 - v37 + v31>>(uint(int32(8))%32)*int32(157) + v54 + v26&int32(255) - v59
	v61 = int32(94)
	v62 = base.I32_div_s(v60, v61)
	v130 = v60 - v62*v61 + v26&int32(65280) + v62<<(uint(int32(8))%32) + v59
	goto L1
L21:
	;
	v91 = int32(65438)
	goto L23
L22:
	;
	v91 = int32(65472)
	goto L23
L23:
	;
	v92 = l2&v74 - v22&v74 + int32(base.Ui32(v31)>>(uint(int32(8))%32))*int32(94) + v85 + v91
	v94 = int32(157)
	v95 = base.I32_div_s(base.I32_extend16_s(v92), v94)
	v98 = v92 - v95*v94
	if int32(62) < base.I32_extend16_s(v98) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v110 = int32(98)
	goto L26
L25:
	;
	v110 = int32(64)
	goto L26
L26:
	;
	v130 = v98 + v26&int32(65280) + v95<<(uint(int32(8))%32) + v110
	goto L1
L27:
	;
	v114 = v15
	goto L29
L28:
	;
	v114 = v14 + int32(1)
	goto L29
L29:
	;
	if base.Ui32(l2) < base.Ui32(v22) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v117 = v14 - int32(1)
	goto L32
L31:
	;
	v117 = v12
	goto L32
L32:
	;
	if v114 <= v117 {
		v12 = v117
		v14 = (v114 + v117) >> (uint(int32(1)) % 32)
		v15 = v114
		goto L2
	} else {
		goto L33
	}
L33:
	;
	goto L4
}
func F_binary_decode(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int64
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_text_to_cstring(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L77
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L73
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L69
	}
L6:
	;
	v155 = int32(1)
	v156 = v14 + v155
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v161 = v159 & v155
	if v161 != 0 {
		goto L50
	} else {
		goto L51
	}
L7:
	;
	v25 = int32(28097)
	v26 = v21
	goto L9
L8:
	;
	if v63 == int32(0) {
		v154 = int32(1689936)
		goto L6
	} else {
		goto L21
	}
L9:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v29 == v30 {
		v52 = v29
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v63 = int32(0)
	goto L8
L11:
	;
	v54 = int32(1)
	if v52 != 0 {
		v25 = v25 + v54
		v26 = v26 + v54
		goto L9
	} else {
		goto L20
	}
L12:
	;
	if base.Ui32((v29-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = v29 | int32(32)
	goto L15
L14:
	;
	v40 = v29
	goto L15
L15:
	;
	if base.Ui32((v30-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v49 = v30 | int32(32)
	goto L18
L17:
	;
	v49 = v30
	goto L18
L18:
	;
	if v40 == v49 {
		v52 = v40
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v63 = v40 - v49
	goto L8
L20:
	;
	goto L10
L21:
	;
	v70 = int32(580633)
	v71 = v21
	goto L23
L22:
	;
	if v108 == int32(0) {
		v154 = int32(1689956)
		goto L6
	} else {
		goto L35
	}
L23:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v74 == v75 {
		v97 = v74
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v108 = int32(0)
	goto L22
L25:
	;
	v99 = int32(1)
	if v97 != 0 {
		v70 = v70 + v99
		v71 = v71 + v99
		goto L23
	} else {
		goto L34
	}
L26:
	;
	if base.Ui32((v74-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v85 = v74 | int32(32)
	goto L29
L28:
	;
	v85 = v74
	goto L29
L29:
	;
	if base.Ui32((v75-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v94 = v75 | int32(32)
	goto L32
L31:
	;
	v94 = v75
	goto L32
L32:
	;
	if v85 == v94 {
		v97 = v85
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v108 = v85 - v94
	goto L22
L34:
	;
	goto L24
L35:
	;
	v114 = int32(387069)
	v115 = v21
	goto L37
L36:
	;
	if v152 != 0 {
		goto L5
	} else {
		goto L49
	}
L37:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if v118 == v119 {
		v141 = v118
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v152 = int32(0)
	goto L36
L39:
	;
	v143 = int32(1)
	if v141 != 0 {
		v114 = v114 + v143
		v115 = v115 + v143
		goto L37
	} else {
		goto L48
	}
L40:
	;
	if base.Ui32((v118-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v129 = v118 | int32(32)
	goto L43
L42:
	;
	v129 = v118
	goto L43
L43:
	;
	if base.Ui32((v119-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v138 = v119 | int32(32)
	goto L46
L45:
	;
	v138 = v119
	goto L46
L46:
	;
	if v129 == v138 {
		v141 = v129
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v152 = v129 - v138
	goto L36
L48:
	;
	goto L38
L49:
	;
	v154 = int32(1689976)
	goto L6
L50:
	;
	v162 = v156
	goto L52
L51:
	;
	v162 = v14 + int32(4)
	goto L52
L52:
	;
	if v159 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	v192 = m.T0[v191].(func(*base.Module, int32, int32) int64)(m, v162, v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L64
	}
L54:
	;
	v165 = int32(4)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v167&int32(254) == int32(2) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	v180 = int32(1)
	if v161 != 0 {
		v190 = int32(base.Ui32(v159)>>(uint(v180)%32)) - v180
		goto L53
	} else {
		goto L63
	}
L57:
	;
	v176 = v165
	goto L59
L58:
	;
	v176 = base.B2i32(v167 == int32(18)) << (uint(v165) % 32)
	goto L59
L59:
	;
	if v167 == int32(1) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v179 = v165
	goto L62
L61:
	;
	v179 = v176
	goto L62
L62:
	;
	v190 = v179
	goto L53
L63:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v190 = int32(base.Ui32(v184)>>(uint(int32(2))%32)) - int32(4)
	goto L53
L64:
	;
	if base.Ui64(int64(1073741820)) <= base.Ui64(v192) {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v199 = F_palloc(m, base.I32_wrap_i64(v192)+int32(4))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	v204 = m.T0[v203].(func(*base.Module, int32, int32, int32) int64)(m, v162, v190, v199+int32(4))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if base.Ui64(v192) < base.Ui64(v204) {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	v210 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = base.I32_wrap_i64(v204)<<(uint(int32(2))%32) + v210
	m.G0 = v11 + v210
	return v199
L69:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v21
	F_errmsg(m, int32(754464), v11)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(519607), int32(114), int32(429758))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(416151), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(519607), int32(128), int32(429758))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errmsg_internal(m, int32(316660), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(519607), int32(136), int32(429758))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_binary_upgrade_set_next_heap_relfilenode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[297])))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(428481), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(511686), int32(112), int32(426980))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, _consts[1316])) = v25
		return int32(0)
	}
}
func F_binary_upgrade_set_next_index_pg_class_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[297])))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(428481), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(511686), int32(123), int32(450617))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, _consts[1317])) = v25
		return int32(0)
	}
}
func F_binary_upgrade_set_next_multirange_array_pg_type_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[297])))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(428481), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(511686), int32(90), int32(451568))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, _consts[1315])) = v25
		return int32(0)
	}
}
func F_executeBinaryArithmExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v12
	F_jspGetArg(m, l1, v10+int32(52))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = int32(2)
		v28 = F_executeItemOptUnwrapResult(m, l0, v10+int32(52), l2, int32(1), v10+int32(40))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			if v28 == int32(2) {
				v162 = v22
				m.G0 = v10 + int32(80)
				return v162
			} else {
				F_jspGetRightArg(m, l1, v10+int32(52))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v41 = F_executeItemOptUnwrapResult(m, l0, v10+int32(52), l2, int32(1), v10+int32(32))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						if v41 == int32(2) {
							v162 = v22
							m.G0 = v10 + int32(80)
							return v162
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
							if v45 == int32(0) {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
								if v48 == int32(0) {
									v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
									if v61 != int32(1) {
										v162 = v22
										m.G0 = v10 + int32(80)
										return v162
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(135004290))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												v72 = F_jspOperationName(m, v71)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10))) = v72
													F_errmsg(m, int32(360871), v10)
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(520430), int32(2137), int32(216162))
														mBase = m.M
														v82 = m.ExcPending
														if v82 != 0 {
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
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
									if v51 != int32(1) {
										v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
										if v61 != int32(1) {
											v162 = v22
											m.G0 = v10 + int32(80)
											return v162
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(135004290))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
													v72 = F_jspOperationName(m, v71)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10))) = v72
														F_errmsg(m, int32(360871), v10)
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(520430), int32(2137), int32(216162))
															mBase = m.M
															v82 = m.ExcPending
															if v82 != 0 {
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
									} else {
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
										v56 = v55
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
										if v57 == int32(2) {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
											if v83 == int32(0) {
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
												if v86 == int32(0) {
													v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
													if v99 != int32(1) {
														v162 = int32(2)
														m.G0 = v10 + int32(80)
														return v162
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(135004290))
															mBase = m.M
															v109 = m.ExcPending
															if v109 != 0 {
																return int32(0)
															} else {
																v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																v111 = F_jspOperationName(m, v110)
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v111
																	F_errmsg(m, int32(360803), v10+int32(16))
																	mBase = m.M
																	v118 = m.ExcPending
																	if v118 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(520430), int32(2144), int32(216162))
																		mBase = m.M
																		v123 = m.ExcPending
																		if v123 != 0 {
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
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
													if v89 != int32(1) {
														v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
														if v99 != int32(1) {
															v162 = int32(2)
															m.G0 = v10 + int32(80)
															return v162
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(135004290))
																mBase = m.M
																v109 = m.ExcPending
																if v109 != 0 {
																	return int32(0)
																} else {
																	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																	v111 = F_jspOperationName(m, v110)
																	mBase = m.M
																	v112 = m.ExcPending
																	if v112 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v111
																		F_errmsg(m, int32(360803), v10+int32(16))
																		mBase = m.M
																		v118 = m.ExcPending
																		if v118 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(520430), int32(2144), int32(216162))
																			mBase = m.M
																			v123 = m.ExcPending
																			if v123 != 0 {
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
													} else {
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
														v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
														v94 = v93
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
														if v95 == int32(2) {
															v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
															if v124 == int32(1) {
																v127 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
																v128 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
																v130 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v127, v128, int32(0))
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return int32(0)
																} else {
																	v144 = v130
																	v147 = F_jspGetNext(m, l1, v10+int32(52))
																	mBase = m.M
																	v148 = m.ExcPending
																	if v148 != 0 {
																		return int32(0)
																	} else {
																		if l4 != 0 {
																			v151 = F_palloc(m, int32(20))
																			mBase = m.M
																			v152 = m.ExcPending
																			if v152 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																				*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																				v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																				mBase = m.M
																				v160 = m.ExcPending
																				if v160 != 0 {
																					return int32(0)
																				} else {
																					v162 = v159
																					m.G0 = v10 + int32(80)
																					return v162
																				}
																			}
																		} else {
																			if v147 != 0 {
																				v151 = F_palloc(m, int32(20))
																				mBase = m.M
																				v152 = m.ExcPending
																				if v152 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																					*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																					v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																					mBase = m.M
																					v160 = m.ExcPending
																					if v160 != 0 {
																						return int32(0)
																					} else {
																						v162 = v159
																						m.G0 = v10 + int32(80)
																						return v162
																					}
																				}
																			} else {
																				v162 = int32(0)
																				m.G0 = v10 + int32(80)
																				return v162
																			}
																		}
																	}
																}
															} else {
																v132 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v10)+31)) = uint8(v132)
																v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
																v135 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
																v138 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v134, v135, v10+int32(31))
																mBase = m.M
																v139 = m.ExcPending
																if v139 != 0 {
																	return int32(0)
																} else {
																	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+31)))
																	if v140 == int32(0) {
																		v144 = v138
																		v147 = F_jspGetNext(m, l1, v10+int32(52))
																		mBase = m.M
																		v148 = m.ExcPending
																		if v148 != 0 {
																			return int32(0)
																		} else {
																			if l4 != 0 {
																				v151 = F_palloc(m, int32(20))
																				mBase = m.M
																				v152 = m.ExcPending
																				if v152 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																					*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																					v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																					mBase = m.M
																					v160 = m.ExcPending
																					if v160 != 0 {
																						return int32(0)
																					} else {
																						v162 = v159
																						m.G0 = v10 + int32(80)
																						return v162
																					}
																				}
																			} else {
																				if v147 != 0 {
																					v151 = F_palloc(m, int32(20))
																					mBase = m.M
																					v152 = m.ExcPending
																					if v152 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																						*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																						v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																						mBase = m.M
																						v160 = m.ExcPending
																						if v160 != 0 {
																							return int32(0)
																						} else {
																							v162 = v159
																							m.G0 = v10 + int32(80)
																							return v162
																						}
																					}
																				} else {
																					v162 = int32(0)
																					m.G0 = v10 + int32(80)
																					return v162
																				}
																			}
																		}
																	} else {
																		v162 = int32(2)
																		m.G0 = v10 + int32(80)
																		return v162
																	}
																}
															}
														} else {
															v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
															if v99 != int32(1) {
																v162 = int32(2)
																m.G0 = v10 + int32(80)
																return v162
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v106 = m.ExcPending
																if v106 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(135004290))
																	mBase = m.M
																	v109 = m.ExcPending
																	if v109 != 0 {
																		return int32(0)
																	} else {
																		v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																		v111 = F_jspOperationName(m, v110)
																		mBase = m.M
																		v112 = m.ExcPending
																		if v112 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v111
																			F_errmsg(m, int32(360803), v10+int32(16))
																			mBase = m.M
																			v118 = m.ExcPending
																			if v118 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(520430), int32(2144), int32(216162))
																				mBase = m.M
																				v123 = m.ExcPending
																				if v123 != 0 {
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
											} else {
												v94 = v83
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
												if v95 == int32(2) {
													v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
													if v124 == int32(1) {
														v127 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
														v130 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v127, v128, int32(0))
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return int32(0)
														} else {
															v144 = v130
															v147 = F_jspGetNext(m, l1, v10+int32(52))
															mBase = m.M
															v148 = m.ExcPending
															if v148 != 0 {
																return int32(0)
															} else {
																if l4 != 0 {
																	v151 = F_palloc(m, int32(20))
																	mBase = m.M
																	v152 = m.ExcPending
																	if v152 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																		*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																		v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																		mBase = m.M
																		v160 = m.ExcPending
																		if v160 != 0 {
																			return int32(0)
																		} else {
																			v162 = v159
																			m.G0 = v10 + int32(80)
																			return v162
																		}
																	}
																} else {
																	if v147 != 0 {
																		v151 = F_palloc(m, int32(20))
																		mBase = m.M
																		v152 = m.ExcPending
																		if v152 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																			*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																			v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																			mBase = m.M
																			v160 = m.ExcPending
																			if v160 != 0 {
																				return int32(0)
																			} else {
																				v162 = v159
																				m.G0 = v10 + int32(80)
																				return v162
																			}
																		}
																	} else {
																		v162 = int32(0)
																		m.G0 = v10 + int32(80)
																		return v162
																	}
																}
															}
														}
													} else {
														v132 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+31)) = uint8(v132)
														v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
														v135 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
														v138 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v134, v135, v10+int32(31))
														mBase = m.M
														v139 = m.ExcPending
														if v139 != 0 {
															return int32(0)
														} else {
															v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+31)))
															if v140 == int32(0) {
																v144 = v138
																v147 = F_jspGetNext(m, l1, v10+int32(52))
																mBase = m.M
																v148 = m.ExcPending
																if v148 != 0 {
																	return int32(0)
																} else {
																	if l4 != 0 {
																		v151 = F_palloc(m, int32(20))
																		mBase = m.M
																		v152 = m.ExcPending
																		if v152 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																			*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																			v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																			mBase = m.M
																			v160 = m.ExcPending
																			if v160 != 0 {
																				return int32(0)
																			} else {
																				v162 = v159
																				m.G0 = v10 + int32(80)
																				return v162
																			}
																		}
																	} else {
																		if v147 != 0 {
																			v151 = F_palloc(m, int32(20))
																			mBase = m.M
																			v152 = m.ExcPending
																			if v152 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																				*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																				v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																				mBase = m.M
																				v160 = m.ExcPending
																				if v160 != 0 {
																					return int32(0)
																				} else {
																					v162 = v159
																					m.G0 = v10 + int32(80)
																					return v162
																				}
																			}
																		} else {
																			v162 = int32(0)
																			m.G0 = v10 + int32(80)
																			return v162
																		}
																	}
																}
															} else {
																v162 = int32(2)
																m.G0 = v10 + int32(80)
																return v162
															}
														}
													}
												} else {
													v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
													if v99 != int32(1) {
														v162 = int32(2)
														m.G0 = v10 + int32(80)
														return v162
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(135004290))
															mBase = m.M
															v109 = m.ExcPending
															if v109 != 0 {
																return int32(0)
															} else {
																v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																v111 = F_jspOperationName(m, v110)
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v111
																	F_errmsg(m, int32(360803), v10+int32(16))
																	mBase = m.M
																	v118 = m.ExcPending
																	if v118 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(520430), int32(2144), int32(216162))
																		mBase = m.M
																		v123 = m.ExcPending
																		if v123 != 0 {
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
										} else {
											v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
											if v61 != int32(1) {
												v162 = v22
												m.G0 = v10 + int32(80)
												return v162
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(135004290))
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
														v72 = F_jspOperationName(m, v71)
														mBase = m.M
														v73 = m.ExcPending
														if v73 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v10))) = v72
															F_errmsg(m, int32(360871), v10)
															mBase = m.M
															v77 = m.ExcPending
															if v77 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(520430), int32(2137), int32(216162))
																mBase = m.M
																v82 = m.ExcPending
																if v82 != 0 {
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
							} else {
								v56 = v45
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								if v57 == int32(2) {
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
									if v83 == int32(0) {
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
										if v86 == int32(0) {
											v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
											if v99 != int32(1) {
												v162 = int32(2)
												m.G0 = v10 + int32(80)
												return v162
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(135004290))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return int32(0)
													} else {
														v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
														v111 = F_jspOperationName(m, v110)
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v111
															F_errmsg(m, int32(360803), v10+int32(16))
															mBase = m.M
															v118 = m.ExcPending
															if v118 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(520430), int32(2144), int32(216162))
																mBase = m.M
																v123 = m.ExcPending
																if v123 != 0 {
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
										} else {
											v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
											if v89 != int32(1) {
												v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
												if v99 != int32(1) {
													v162 = int32(2)
													m.G0 = v10 + int32(80)
													return v162
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(135004290))
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return int32(0)
														} else {
															v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
															v111 = F_jspOperationName(m, v110)
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v111
																F_errmsg(m, int32(360803), v10+int32(16))
																mBase = m.M
																v118 = m.ExcPending
																if v118 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(520430), int32(2144), int32(216162))
																	mBase = m.M
																	v123 = m.ExcPending
																	if v123 != 0 {
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
											} else {
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
												v94 = v93
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
												if v95 == int32(2) {
													v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
													if v124 == int32(1) {
														v127 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
														v130 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v127, v128, int32(0))
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return int32(0)
														} else {
															v144 = v130
															v147 = F_jspGetNext(m, l1, v10+int32(52))
															mBase = m.M
															v148 = m.ExcPending
															if v148 != 0 {
																return int32(0)
															} else {
																if l4 != 0 {
																	v151 = F_palloc(m, int32(20))
																	mBase = m.M
																	v152 = m.ExcPending
																	if v152 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																		*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																		v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																		mBase = m.M
																		v160 = m.ExcPending
																		if v160 != 0 {
																			return int32(0)
																		} else {
																			v162 = v159
																			m.G0 = v10 + int32(80)
																			return v162
																		}
																	}
																} else {
																	if v147 != 0 {
																		v151 = F_palloc(m, int32(20))
																		mBase = m.M
																		v152 = m.ExcPending
																		if v152 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																			*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																			v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																			mBase = m.M
																			v160 = m.ExcPending
																			if v160 != 0 {
																				return int32(0)
																			} else {
																				v162 = v159
																				m.G0 = v10 + int32(80)
																				return v162
																			}
																		}
																	} else {
																		v162 = int32(0)
																		m.G0 = v10 + int32(80)
																		return v162
																	}
																}
															}
														}
													} else {
														v132 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+31)) = uint8(v132)
														v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
														v135 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
														v138 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v134, v135, v10+int32(31))
														mBase = m.M
														v139 = m.ExcPending
														if v139 != 0 {
															return int32(0)
														} else {
															v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+31)))
															if v140 == int32(0) {
																v144 = v138
																v147 = F_jspGetNext(m, l1, v10+int32(52))
																mBase = m.M
																v148 = m.ExcPending
																if v148 != 0 {
																	return int32(0)
																} else {
																	if l4 != 0 {
																		v151 = F_palloc(m, int32(20))
																		mBase = m.M
																		v152 = m.ExcPending
																		if v152 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																			*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																			v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																			mBase = m.M
																			v160 = m.ExcPending
																			if v160 != 0 {
																				return int32(0)
																			} else {
																				v162 = v159
																				m.G0 = v10 + int32(80)
																				return v162
																			}
																		}
																	} else {
																		if v147 != 0 {
																			v151 = F_palloc(m, int32(20))
																			mBase = m.M
																			v152 = m.ExcPending
																			if v152 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																				*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																				v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																				mBase = m.M
																				v160 = m.ExcPending
																				if v160 != 0 {
																					return int32(0)
																				} else {
																					v162 = v159
																					m.G0 = v10 + int32(80)
																					return v162
																				}
																			}
																		} else {
																			v162 = int32(0)
																			m.G0 = v10 + int32(80)
																			return v162
																		}
																	}
																}
															} else {
																v162 = int32(2)
																m.G0 = v10 + int32(80)
																return v162
															}
														}
													}
												} else {
													v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
													if v99 != int32(1) {
														v162 = int32(2)
														m.G0 = v10 + int32(80)
														return v162
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(135004290))
															mBase = m.M
															v109 = m.ExcPending
															if v109 != 0 {
																return int32(0)
															} else {
																v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																v111 = F_jspOperationName(m, v110)
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v111
																	F_errmsg(m, int32(360803), v10+int32(16))
																	mBase = m.M
																	v118 = m.ExcPending
																	if v118 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(520430), int32(2144), int32(216162))
																		mBase = m.M
																		v123 = m.ExcPending
																		if v123 != 0 {
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
									} else {
										v94 = v83
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
										if v95 == int32(2) {
											v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
											if v124 == int32(1) {
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
												v128 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
												v130 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v127, v128, int32(0))
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int32(0)
												} else {
													v144 = v130
													v147 = F_jspGetNext(m, l1, v10+int32(52))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														if l4 != 0 {
															v151 = F_palloc(m, int32(20))
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
																	return int32(0)
																} else {
																	v162 = v159
																	m.G0 = v10 + int32(80)
																	return v162
																}
															}
														} else {
															if v147 != 0 {
																v151 = F_palloc(m, int32(20))
																mBase = m.M
																v152 = m.ExcPending
																if v152 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																	*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																	v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																	mBase = m.M
																	v160 = m.ExcPending
																	if v160 != 0 {
																		return int32(0)
																	} else {
																		v162 = v159
																		m.G0 = v10 + int32(80)
																		return v162
																	}
																}
															} else {
																v162 = int32(0)
																m.G0 = v10 + int32(80)
																return v162
															}
														}
													}
												}
											} else {
												v132 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+31)) = uint8(v132)
												v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
												v135 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
												v138 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v134, v135, v10+int32(31))
												mBase = m.M
												v139 = m.ExcPending
												if v139 != 0 {
													return int32(0)
												} else {
													v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+31)))
													if v140 == int32(0) {
														v144 = v138
														v147 = F_jspGetNext(m, l1, v10+int32(52))
														mBase = m.M
														v148 = m.ExcPending
														if v148 != 0 {
															return int32(0)
														} else {
															if l4 != 0 {
																v151 = F_palloc(m, int32(20))
																mBase = m.M
																v152 = m.ExcPending
																if v152 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																	*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																	v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																	mBase = m.M
																	v160 = m.ExcPending
																	if v160 != 0 {
																		return int32(0)
																	} else {
																		v162 = v159
																		m.G0 = v10 + int32(80)
																		return v162
																	}
																}
															} else {
																if v147 != 0 {
																	v151 = F_palloc(m, int32(20))
																	mBase = m.M
																	v152 = m.ExcPending
																	if v152 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v144
																		*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(2)
																		v159 = F_executeNextItem(m, l0, l1, v10+int32(52), v151, l4, int32(0))
																		mBase = m.M
																		v160 = m.ExcPending
																		if v160 != 0 {
																			return int32(0)
																		} else {
																			v162 = v159
																			m.G0 = v10 + int32(80)
																			return v162
																		}
																	}
																} else {
																	v162 = int32(0)
																	m.G0 = v10 + int32(80)
																	return v162
																}
															}
														}
													} else {
														v162 = int32(2)
														m.G0 = v10 + int32(80)
														return v162
													}
												}
											}
										} else {
											v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
											if v99 != int32(1) {
												v162 = int32(2)
												m.G0 = v10 + int32(80)
												return v162
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(135004290))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return int32(0)
													} else {
														v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
														v111 = F_jspOperationName(m, v110)
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v111
															F_errmsg(m, int32(360803), v10+int32(16))
															mBase = m.M
															v118 = m.ExcPending
															if v118 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(520430), int32(2144), int32(216162))
																mBase = m.M
																v123 = m.ExcPending
																if v123 != 0 {
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
								} else {
									v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
									if v61 != int32(1) {
										v162 = v22
										m.G0 = v10 + int32(80)
										return v162
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(135004290))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												v72 = F_jspOperationName(m, v71)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10))) = v72
													F_errmsg(m, int32(360871), v10)
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(520430), int32(2137), int32(216162))
														mBase = m.M
														v82 = m.ExcPending
														if v82 != 0 {
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
				}
			}
		}
	}
}
