package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IsBinaryCoercibleWithCast(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v7 = int32(1)
	if l0 == l1 {
		v99 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v99
L2:
	;
	switch l1 - int32(2276) {
	case 0, 7:
		v99 = v7
		goto L1
	case 1, 2, 3, 4, 5, 6:
		goto L3
	default:
		goto L4
	}
L3:
	;
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if l1 == int32(5077) {
		v99 = v7
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	v13 = F_getBaseType(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v18 = int32(0)
	goto L8
L8:
	;
	if l1 == v18 {
		v99 = v7
		goto L1
	} else {
		goto L11
	}
L9:
	;
	return int32(0)
L10:
	;
	v18 = v13
	goto L8
L11:
	;
	if l1 <= int32(3830) {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	v75 = F_SearchSysCache2(m, int32(12), v18, l1)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L9
	} else {
		goto L52
	}
L13:
	;
	if base.Ui32(l1-int32(4537)) <= base.Ui32(int32(1)) {
		goto L39
	} else {
		goto L40
	}
L14:
	;
	v50 = F_type_is_range(m, v18)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L37
	}
L15:
	;
	if l1 != int32(3831) {
		goto L13
	} else {
		goto L36
	}
L16:
	;
	v46 = F_type_is_enum(m, v18)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L34
	}
L17:
	;
	v38 = F_get_element_type(m, v18)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L30
	}
L18:
	;
	v30 = F_get_element_type(m, v18)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L9
	} else {
		goto L25
	}
L19:
	;
	if l1 == int32(2277) {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	switch l1 - int32(5078) {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L14
	default:
		goto L15
	}
L22:
	;
	if l1 == int32(2776) {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	if l1 == int32(3500) {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	goto L13
L25:
	;
	if v30 != 0 {
		v99 = v7
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if l1 == int32(3831) {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	if l1 == int32(3500) {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	if l1 != int32(2776) {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	goto L17
L30:
	;
	if v38 == int32(0) {
		v99 = v7
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if l1 == int32(3831) {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	if l1 != int32(3500) {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	goto L16
L34:
	;
	if v46 != 0 {
		v99 = v7
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L12
L36:
	;
	goto L14
L37:
	;
	if v50 != 0 {
		v99 = v7
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L13
L39:
	;
	v56 = F_type_is_multirange(m, v18)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if l1 != int32(2287) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	if v56 != 0 {
		v99 = v7
		goto L1
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	if l1 != int32(2249) {
		goto L12
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v68 = F_is_complex_array(m, v18)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L9
	} else {
		goto L50
	}
L47:
	;
	v62 = F_typeOrDomainTypeRelid(m, v18)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	if v62 == int32(0) {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	return int32(1)
L50:
	;
	if v68 == int32(0) {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	return int32(1)
L52:
	;
	if v75 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	return int32(0)
L54:
	;
	goto L55
L55:
	;
	v81 = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+22)))
	v84 = v82 + v83
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+17)))
	if v85 != int32(98) {
		v94 = v81
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_ReleaseCatCache(m, v75)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L59
	}
L57:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+16)))
	if v88 != int32(105) {
		v94 = v81
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v91
	v94 = int32(1)
	goto L56
L59:
	;
	v99 = v94
	goto L1
}
func F_appendBinaryStringInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	F_enlargeStringInfo(m, l0, l2)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if l2 != 0 {
			v9 = F__emscripten_memcpy_bulkmem(m, v6+v7, l1, l2)
			mBase = m.M
		} else {
		}
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v12 = v11 + l2
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14+v12))) = uint8(v16)
		return
	}
}
func F_binary_encode(m *base.Module, l0 int32) int32 {
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
	v25 = int32(27723)
	v26 = v21
	goto L9
L8:
	;
	if v63 == int32(0) {
		v154 = int32(1661264)
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
	v70 = int32(558570)
	v71 = v21
	goto L23
L22:
	;
	if v108 == int32(0) {
		v154 = int32(1661284)
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
	v114 = int32(371803)
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
	v154 = int32(1661304)
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
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
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
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
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
	F_errmsg(m, int32(726150), v11)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(499199), int32(66), int32(413354))
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
	F_errmsg(m, int32(400142), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(499199), int32(80), int32(413354))
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
	F_errmsg_internal(m, int32(304509), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(499199), int32(88), int32(413354))
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
func F_binary_upgrade_set_next_heap_pg_class_oid(m *base.Module, l0 int32) int32 {
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
				F_errmsg(m, int32(412153), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492612), int32(101), int32(433662))
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
		*(*int32)(unsafe.Add(mBase, _consts[1129])) = v25
		return int32(0)
	}
}
func F_binary_upgrade_set_next_multirange_pg_type_oid(m *base.Module, l0 int32) int32 {
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
				F_errmsg(m, int32(412153), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492612), int32(79), int32(434595))
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
		*(*int32)(unsafe.Add(mBase, _consts[940])) = v25
		return int32(0)
	}
}
func F_binary_upgrade_set_next_pg_authid_oid(m *base.Module, l0 int32) int32 {
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
				F_errmsg(m, int32(412153), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492612), int32(178), int32(434813))
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
		*(*int32)(unsafe.Add(mBase, _consts[1130])) = v25
		return int32(0)
	}
}
