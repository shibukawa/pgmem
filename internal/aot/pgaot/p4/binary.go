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
	var v18 int32
	_ = v18
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
	var v75 int32
	_ = v75
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
	var v132 int32
	_ = v132
	v12 = l1
	v14 = int32(base.Ui32(l1) >> (uint(int32(1)) % 32))
	v18 = int32(0)
	goto L2
L1:
	;
	return v132 & int32(_a_F_BinarySearchRange_0)
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
	v132 = int32(0)
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
	v31 = l2 - v22&int32(_a_F_BinarySearchRange_1)
	if base.Ui32(int32(_a_F_BinarySearchRange_2)) <= base.Ui32(l2) {
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
	v75 = v26 & v74
	if base.Ui32(int32(160)) < base.Ui32(v75) {
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
	v132 = v60 - v62*v61 + v26&int32(_a_F_BinarySearchRange_1) + v62<<(uint(int32(8))%32) + v59
	goto L1
L21:
	;
	v91 = int32(_a_F_BinarySearchRange_3)
	goto L23
L22:
	;
	v91 = int32(_a_F_BinarySearchRange_4)
	goto L23
L23:
	;
	v92 = v75 + (l2&v74 - v22&v74 + int32(base.Ui32(v31)>>(uint(int32(8))%32))*int32(94)) + v91
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
	v132 = v98 + v26&int32(_a_F_BinarySearchRange_1) + v95<<(uint(int32(8))%32) + v110
	goto L1
L27:
	;
	v114 = v18
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
		v18 = v114
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
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
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = F_text_to_cstring(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
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
	v156 = int32(1)
	v157 = v15 + v156
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v162 = v160 & v156
	if v162 != 0 {
		goto L50
	} else {
		goto L51
	}
L7:
	;
	v26 = int32(_a_F_binary_decode_0)
	v27 = v22
	goto L9
L8:
	;
	if v64 == int32(0) {
		v155 = int32(_a_F_binary_decode_1)
		goto L6
	} else {
		goto L21
	}
L9:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v30 == v31 {
		v53 = v30
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v64 = int32(0)
	goto L8
L11:
	;
	v55 = int32(1)
	if v53 != 0 {
		v26 = v26 + v55
		v27 = v27 + v55
		goto L9
	} else {
		goto L20
	}
L12:
	;
	if base.Ui32((v30-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v41 = v30 | int32(32)
	goto L15
L14:
	;
	v41 = v30
	goto L15
L15:
	;
	if base.Ui32((v31-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v50 = v31 | int32(32)
	goto L18
L17:
	;
	v50 = v31
	goto L18
L18:
	;
	if v41 == v50 {
		v53 = v41
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v64 = v41 - v50
	goto L8
L20:
	;
	goto L10
L21:
	;
	v71 = int32(_a_F_binary_decode_2)
	v72 = v22
	goto L23
L22:
	;
	if v109 == int32(0) {
		v155 = int32(_a_F_binary_decode_3)
		goto L6
	} else {
		goto L35
	}
L23:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v75 == v76 {
		v98 = v75
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v109 = int32(0)
	goto L22
L25:
	;
	v100 = int32(1)
	if v98 != 0 {
		v71 = v71 + v100
		v72 = v72 + v100
		goto L23
	} else {
		goto L34
	}
L26:
	;
	if base.Ui32((v75-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v86 = v75 | int32(32)
	goto L29
L28:
	;
	v86 = v75
	goto L29
L29:
	;
	if base.Ui32((v76-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v95 = v76 | int32(32)
	goto L32
L31:
	;
	v95 = v76
	goto L32
L32:
	;
	if v86 == v95 {
		v98 = v86
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v109 = v86 - v95
	goto L22
L34:
	;
	goto L24
L35:
	;
	v115 = int32(_a_F_binary_decode_4)
	v116 = v22
	goto L37
L36:
	;
	if v153 != 0 {
		goto L5
	} else {
		goto L49
	}
L37:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v119 == v120 {
		v142 = v119
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v153 = int32(0)
	goto L36
L39:
	;
	v144 = int32(1)
	if v142 != 0 {
		v115 = v115 + v144
		v116 = v116 + v144
		goto L37
	} else {
		goto L48
	}
L40:
	;
	if base.Ui32((v119-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v130 = v119 | int32(32)
	goto L43
L42:
	;
	v130 = v119
	goto L43
L43:
	;
	if base.Ui32((v120-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v139 = v120 | int32(32)
	goto L46
L45:
	;
	v139 = v120
	goto L46
L46:
	;
	if v130 == v139 {
		v142 = v130
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v153 = v130 - v139
	goto L36
L48:
	;
	goto L38
L49:
	;
	v155 = int32(_a_F_binary_decode_5)
	goto L6
L50:
	;
	v163 = v157
	goto L52
L51:
	;
	v163 = v15 + int32(4)
	goto L52
L52:
	;
	if v160 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
	v192 = m.T0[v191].(func(*base.Module, int32, int32) int64)(m, v163, v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L64
	}
L54:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v169 == int32(18) {
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
	if v162 != 0 {
		v190 = int32(base.Ui32(v160)>>(uint(v180)%32)) - v180
		goto L53
	} else {
		goto L63
	}
L57:
	;
	v172 = int32(16)
	goto L59
L58:
	;
	v172 = int32(0)
	goto L59
L59:
	;
	if base.Ui32((v169-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v179 = int32(4)
	goto L62
L61:
	;
	v179 = v172
	goto L62
L62:
	;
	v190 = v179
	goto L53
L63:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
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
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	v204 = m.T0[v203].(func(*base.Module, int32, int32, int32) int64)(m, v163, v190, v199+int32(4))
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
	m.G0 = v12 + v210
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
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v22
	F_errmsg(m, int32(_a_F_binary_decode_6), v12)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_binary_decode_7), int32(114), int32(_a_F_binary_decode_8))
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
	F_errmsg(m, int32(_a_F_binary_decode_9), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_binary_decode_7), int32(128), int32(_a_F_binary_decode_8))
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
	F_errmsg_internal(m, int32(_a_F_binary_decode_10), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_binary_decode_7), int32(136), int32(_a_F_binary_decode_8))
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13871(m, l0, int32(_a_F_binary_upgrade_set_next_heap_relfilenode_0), int32(_a_F_binary_upgrade_set_next_heap_relfilenode_1), int32(112))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_binary_upgrade_set_next_index_pg_class_oid(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13871(m, l0, int32(_a_F_binary_upgrade_set_next_index_pg_class_oid_0), int32(_a_F_binary_upgrade_set_next_index_pg_class_oid_1), int32(123))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_binary_upgrade_set_next_multirange_array_pg_type_oid(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13871(m, l0, int32(_a_F_binary_upgrade_set_next_multirange_array_pg_type_oid_0), int32(_a_F_binary_upgrade_set_next_multirange_array_pg_type_oid_1), int32(90))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_executeBinaryArithmExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v13
	v18 = v11 + int32(52)
	F_jspGetArg(m, l1, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = int32(2)
		v27 = F_executeItemOptUnwrapResult(m, l0, v18, l2, int32(1), v11+int32(40))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			if v27 == int32(2) {
				v160 = v23
				m.G0 = v11 + int32(80)
				return v160
			} else {
				F_jspGetRightArg(m, l1, v18)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v36 = F_executeItemOptUnwrapResult(m, l0, v18, l2, int32(1), v11+int32(32))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						if v36 == int32(2) {
							v160 = v23
							m.G0 = v11 + int32(80)
							return v160
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
							if v40 == int32(0) {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
								if v43 == int32(0) {
									v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
									if v56 != int32(1) {
										v160 = v23
										m.G0 = v11 + int32(80)
										return v160
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(135004290))
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												v67 = F_jspOperationName(m, v66)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
													F_errmsg(m, int32(_a_F_executeBinaryArithmExpr_0), v11)
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_executeBinaryArithmExpr_1), int32(2137), int32(_a_F_executeBinaryArithmExpr_2))
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
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
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
									if v46 != int32(1) {
										v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
										if v56 != int32(1) {
											v160 = v23
											m.G0 = v11 + int32(80)
											return v160
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(135004290))
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
													v67 = F_jspOperationName(m, v66)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
														F_errmsg(m, int32(_a_F_executeBinaryArithmExpr_0), v11)
														mBase = m.M
														v72 = m.ExcPending
														if v72 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_executeBinaryArithmExpr_1), int32(2137), int32(_a_F_executeBinaryArithmExpr_2))
															mBase = m.M
															v77 = m.ExcPending
															if v77 != 0 {
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
										v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
										v51 = v50
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
										if v52 == int32(2) {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
											if v78 == int32(0) {
												v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
												if v81 == int32(0) {
													v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
													if v94 != int32(1) {
														v160 = int32(2)
														m.G0 = v11 + int32(80)
														return v160
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(135004290))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return int32(0)
															} else {
																v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																v106 = F_jspOperationName(m, v105)
																mBase = m.M
																v107 = m.ExcPending
																if v107 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v106
																	F_errmsg(m, int32(_a_F_executeBinaryArithmExpr_3), v11+int32(16))
																	mBase = m.M
																	v113 = m.ExcPending
																	if v113 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_executeBinaryArithmExpr_1), int32(2144), int32(_a_F_executeBinaryArithmExpr_2))
																		mBase = m.M
																		v118 = m.ExcPending
																		if v118 != 0 {
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
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
													if v84 != int32(1) {
														v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
														if v94 != int32(1) {
															v160 = int32(2)
															m.G0 = v11 + int32(80)
															return v160
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(135004290))
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return int32(0)
																} else {
																	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																	v106 = F_jspOperationName(m, v105)
																	mBase = m.M
																	v107 = m.ExcPending
																	if v107 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v106
																		F_errmsg(m, int32(_a_F_executeBinaryArithmExpr_3), v11+int32(16))
																		mBase = m.M
																		v113 = m.ExcPending
																		if v113 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_executeBinaryArithmExpr_1), int32(2144), int32(_a_F_executeBinaryArithmExpr_2))
																			mBase = m.M
																			v118 = m.ExcPending
																			if v118 != 0 {
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
														v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
														v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
														v89 = v88
														v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
														if v90 == int32(2) {
															v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
															if v119 == int32(1) {
																v122 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
																v123 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
																v125 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v122, v123, int32(0))
																mBase = m.M
																v126 = m.ExcPending
																if v126 != 0 {
																	return int32(0)
																} else {
																	v139 = v125
																	v142 = F_jspGetNext(m, l1, v11+int32(52))
																	mBase = m.M
																	v143 = m.ExcPending
																	if v143 != 0 {
																		return int32(0)
																	} else {
																		if v142|l4 == int32(0) {
																			v160 = int32(0)
																			m.G0 = v11 + int32(80)
																			return v160
																		} else {
																			v149 = F_palloc(m, int32(20))
																			mBase = m.M
																			v150 = m.ExcPending
																			if v150 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v139
																				*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(2)
																				v157 = F_executeNextItem(m, l0, l1, v11+int32(52), v149, l4, int32(0))
																				mBase = m.M
																				v158 = m.ExcPending
																				if v158 != 0 {
																					return int32(0)
																				} else {
																					v160 = v157
																					m.G0 = v11 + int32(80)
																					return v160
																				}
																			}
																		}
																	}
																}
															} else {
																v127 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+31)) = uint8(v127)
																v129 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
																v130 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
																v133 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v129, v130, v11+int32(31))
																mBase = m.M
																v134 = m.ExcPending
																if v134 != 0 {
																	return int32(0)
																} else {
																	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+31)))
																	if v135 == int32(0) {
																		v139 = v133
																		v142 = F_jspGetNext(m, l1, v11+int32(52))
																		mBase = m.M
																		v143 = m.ExcPending
																		if v143 != 0 {
																			return int32(0)
																		} else {
																			if v142|l4 == int32(0) {
																				v160 = int32(0)
																				m.G0 = v11 + int32(80)
																				return v160
																			} else {
																				v149 = F_palloc(m, int32(20))
																				mBase = m.M
																				v150 = m.ExcPending
																				if v150 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v139
																					*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(2)
																					v157 = F_executeNextItem(m, l0, l1, v11+int32(52), v149, l4, int32(0))
																					mBase = m.M
																					v158 = m.ExcPending
																					if v158 != 0 {
																						return int32(0)
																					} else {
																						v160 = v157
																						m.G0 = v11 + int32(80)
																						return v160
																					}
																				}
																			}
																		}
																	} else {
																		v160 = int32(2)
																		m.G0 = v11 + int32(80)
																		return v160
																	}
																}
															}
														} else {
															v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
															if v94 != int32(1) {
																v160 = int32(2)
																m.G0 = v11 + int32(80)
																return v160
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v101 = m.ExcPending
																if v101 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(135004290))
																	mBase = m.M
																	v104 = m.ExcPending
																	if v104 != 0 {
																		return int32(0)
																	} else {
																		v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																		v106 = F_jspOperationName(m, v105)
																		mBase = m.M
																		v107 = m.ExcPending
																		if v107 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v106
																			F_errmsg(m, int32(_a_F_executeBinaryArithmExpr_3), v11+int32(16))
																			mBase = m.M
																			v113 = m.ExcPending
																			if v113 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_executeBinaryArithmExpr_1), int32(2144), int32(_a_F_executeBinaryArithmExpr_2))
																				mBase = m.M
																				v118 = m.ExcPending
																				if v118 != 0 {
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
												v89 = v78
												v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
												if v90 == int32(2) {
													v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
													if v119 == int32(1) {
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
														v123 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
														v125 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v122, v123, int32(0))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return int32(0)
														} else {
															v139 = v125
															v142 = F_jspGetNext(m, l1, v11+int32(52))
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return int32(0)
															} else {
																if v142|l4 == int32(0) {
																	v160 = int32(0)
																	m.G0 = v11 + int32(80)
																	return v160
																} else {
																	v149 = F_palloc(m, int32(20))
																	mBase = m.M
																	v150 = m.ExcPending
																	if v150 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v139
																		*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(2)
																		v157 = F_executeNextItem(m, l0, l1, v11+int32(52), v149, l4, int32(0))
																		mBase = m.M
																		v158 = m.ExcPending
																		if v158 != 0 {
																			return int32(0)
																		} else {
																			v160 = v157
																			m.G0 = v11 + int32(80)
																			return v160
																		}
																	}
																}
															}
														}
													} else {
														v127 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+31)) = uint8(v127)
														v129 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
														v130 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
														v133 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v129, v130, v11+int32(31))
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return int32(0)
														} else {
															v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+31)))
															if v135 == int32(0) {
																v139 = v133
																v142 = F_jspGetNext(m, l1, v11+int32(52))
																mBase = m.M
																v143 = m.ExcPending
																if v143 != 0 {
																	return int32(0)
																} else {
																	if v142|l4 == int32(0) {
																		v160 = int32(0)
																		m.G0 = v11 + int32(80)
																		return v160
																	} else {
																		v149 = F_palloc(m, int32(20))
																		mBase = m.M
																		v150 = m.ExcPending
																		if v150 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v139
																			*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(2)
																			v157 = F_executeNextItem(m, l0, l1, v11+int32(52), v149, l4, int32(0))
																			mBase = m.M
																			v158 = m.ExcPending
																			if v158 != 0 {
																				return int32(0)
																			} else {
																				v160 = v157
																				m.G0 = v11 + int32(80)
																				return v160
																			}
																		}
																	}
																}
															} else {
																v160 = int32(2)
																m.G0 = v11 + int32(80)
																return v160
															}
														}
													}
												} else {
													v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
													if v94 != int32(1) {
														v160 = int32(2)
														m.G0 = v11 + int32(80)
														return v160
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(135004290))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return int32(0)
															} else {
																v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																v106 = F_jspOperationName(m, v105)
																mBase = m.M
																v107 = m.ExcPending
																if v107 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v106
																	F_errmsg(m, int32(_a_F_executeBinaryArithmExpr_3), v11+int32(16))
																	mBase = m.M
																	v113 = m.ExcPending
																	if v113 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_executeBinaryArithmExpr_1), int32(2144), int32(_a_F_executeBinaryArithmExpr_2))
																		mBase = m.M
																		v118 = m.ExcPending
																		if v118 != 0 {
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
											v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
											if v56 != int32(1) {
												v160 = v23
												m.G0 = v11 + int32(80)
												return v160
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(135004290))
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
														v67 = F_jspOperationName(m, v66)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
															F_errmsg(m, int32(_a_F_executeBinaryArithmExpr_0), v11)
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_executeBinaryArithmExpr_1), int32(2137), int32(_a_F_executeBinaryArithmExpr_2))
																mBase = m.M
																v77 = m.ExcPending
																if v77 != 0 {
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
								v51 = v40
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
								if v52 == int32(2) {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
									if v78 == int32(0) {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
										if v81 == int32(0) {
											v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
											if v94 != int32(1) {
												v160 = int32(2)
												m.G0 = v11 + int32(80)
												return v160
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(135004290))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return int32(0)
													} else {
														v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
														v106 = F_jspOperationName(m, v105)
														mBase = m.M
														v107 = m.ExcPending
														if v107 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v106
															F_errmsg(m, int32(_a_F_executeBinaryArithmExpr_3), v11+int32(16))
															mBase = m.M
															v113 = m.ExcPending
															if v113 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_executeBinaryArithmExpr_1), int32(2144), int32(_a_F_executeBinaryArithmExpr_2))
																mBase = m.M
																v118 = m.ExcPending
																if v118 != 0 {
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
											v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
											if v84 != int32(1) {
												v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
												if v94 != int32(1) {
													v160 = int32(2)
													m.G0 = v11 + int32(80)
													return v160
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(135004290))
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return int32(0)
														} else {
															v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
															v106 = F_jspOperationName(m, v105)
															mBase = m.M
															v107 = m.ExcPending
															if v107 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v106
																F_errmsg(m, int32(_a_F_executeBinaryArithmExpr_3), v11+int32(16))
																mBase = m.M
																v113 = m.ExcPending
																if v113 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_executeBinaryArithmExpr_1), int32(2144), int32(_a_F_executeBinaryArithmExpr_2))
																	mBase = m.M
																	v118 = m.ExcPending
																	if v118 != 0 {
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
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
												v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
												v89 = v88
												v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
												if v90 == int32(2) {
													v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
													if v119 == int32(1) {
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
														v123 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
														v125 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v122, v123, int32(0))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return int32(0)
														} else {
															v139 = v125
															v142 = F_jspGetNext(m, l1, v11+int32(52))
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return int32(0)
															} else {
																if v142|l4 == int32(0) {
																	v160 = int32(0)
																	m.G0 = v11 + int32(80)
																	return v160
																} else {
																	v149 = F_palloc(m, int32(20))
																	mBase = m.M
																	v150 = m.ExcPending
																	if v150 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v139
																		*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(2)
																		v157 = F_executeNextItem(m, l0, l1, v11+int32(52), v149, l4, int32(0))
																		mBase = m.M
																		v158 = m.ExcPending
																		if v158 != 0 {
																			return int32(0)
																		} else {
																			v160 = v157
																			m.G0 = v11 + int32(80)
																			return v160
																		}
																	}
																}
															}
														}
													} else {
														v127 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+31)) = uint8(v127)
														v129 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
														v130 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
														v133 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v129, v130, v11+int32(31))
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return int32(0)
														} else {
															v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+31)))
															if v135 == int32(0) {
																v139 = v133
																v142 = F_jspGetNext(m, l1, v11+int32(52))
																mBase = m.M
																v143 = m.ExcPending
																if v143 != 0 {
																	return int32(0)
																} else {
																	if v142|l4 == int32(0) {
																		v160 = int32(0)
																		m.G0 = v11 + int32(80)
																		return v160
																	} else {
																		v149 = F_palloc(m, int32(20))
																		mBase = m.M
																		v150 = m.ExcPending
																		if v150 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v139
																			*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(2)
																			v157 = F_executeNextItem(m, l0, l1, v11+int32(52), v149, l4, int32(0))
																			mBase = m.M
																			v158 = m.ExcPending
																			if v158 != 0 {
																				return int32(0)
																			} else {
																				v160 = v157
																				m.G0 = v11 + int32(80)
																				return v160
																			}
																		}
																	}
																}
															} else {
																v160 = int32(2)
																m.G0 = v11 + int32(80)
																return v160
															}
														}
													}
												} else {
													v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
													if v94 != int32(1) {
														v160 = int32(2)
														m.G0 = v11 + int32(80)
														return v160
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(135004290))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return int32(0)
															} else {
																v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																v106 = F_jspOperationName(m, v105)
																mBase = m.M
																v107 = m.ExcPending
																if v107 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v106
																	F_errmsg(m, int32(_a_F_executeBinaryArithmExpr_3), v11+int32(16))
																	mBase = m.M
																	v113 = m.ExcPending
																	if v113 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_executeBinaryArithmExpr_1), int32(2144), int32(_a_F_executeBinaryArithmExpr_2))
																		mBase = m.M
																		v118 = m.ExcPending
																		if v118 != 0 {
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
										v89 = v78
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
										if v90 == int32(2) {
											v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
											if v119 == int32(1) {
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
												v123 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
												v125 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v122, v123, int32(0))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													v139 = v125
													v142 = F_jspGetNext(m, l1, v11+int32(52))
													mBase = m.M
													v143 = m.ExcPending
													if v143 != 0 {
														return int32(0)
													} else {
														if v142|l4 == int32(0) {
															v160 = int32(0)
															m.G0 = v11 + int32(80)
															return v160
														} else {
															v149 = F_palloc(m, int32(20))
															mBase = m.M
															v150 = m.ExcPending
															if v150 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v139
																*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(2)
																v157 = F_executeNextItem(m, l0, l1, v11+int32(52), v149, l4, int32(0))
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int32(0)
																} else {
																	v160 = v157
																	m.G0 = v11 + int32(80)
																	return v160
																}
															}
														}
													}
												}
											} else {
												v127 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+31)) = uint8(v127)
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
												v130 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
												v133 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v129, v130, v11+int32(31))
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return int32(0)
												} else {
													v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+31)))
													if v135 == int32(0) {
														v139 = v133
														v142 = F_jspGetNext(m, l1, v11+int32(52))
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return int32(0)
														} else {
															if v142|l4 == int32(0) {
																v160 = int32(0)
																m.G0 = v11 + int32(80)
																return v160
															} else {
																v149 = F_palloc(m, int32(20))
																mBase = m.M
																v150 = m.ExcPending
																if v150 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v139
																	*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(2)
																	v157 = F_executeNextItem(m, l0, l1, v11+int32(52), v149, l4, int32(0))
																	mBase = m.M
																	v158 = m.ExcPending
																	if v158 != 0 {
																		return int32(0)
																	} else {
																		v160 = v157
																		m.G0 = v11 + int32(80)
																		return v160
																	}
																}
															}
														}
													} else {
														v160 = int32(2)
														m.G0 = v11 + int32(80)
														return v160
													}
												}
											}
										} else {
											v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
											if v94 != int32(1) {
												v160 = int32(2)
												m.G0 = v11 + int32(80)
												return v160
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(135004290))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return int32(0)
													} else {
														v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
														v106 = F_jspOperationName(m, v105)
														mBase = m.M
														v107 = m.ExcPending
														if v107 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v106
															F_errmsg(m, int32(_a_F_executeBinaryArithmExpr_3), v11+int32(16))
															mBase = m.M
															v113 = m.ExcPending
															if v113 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_executeBinaryArithmExpr_1), int32(2144), int32(_a_F_executeBinaryArithmExpr_2))
																mBase = m.M
																v118 = m.ExcPending
																if v118 != 0 {
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
									v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
									if v56 != int32(1) {
										v160 = v23
										m.G0 = v11 + int32(80)
										return v160
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(135004290))
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												v67 = F_jspOperationName(m, v66)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
													F_errmsg(m, int32(_a_F_executeBinaryArithmExpr_0), v11)
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_executeBinaryArithmExpr_1), int32(2137), int32(_a_F_executeBinaryArithmExpr_2))
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
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
