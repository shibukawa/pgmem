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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v7 = int32(1)
	if l0 == l1 {
		v101 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v101
L2:
	;
	switch l1 - int32(2276) {
	case 0, 7:
		v101 = v7
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
	if l1 == int32(_a_F_IsBinaryCoercibleWithCast_1) {
		v101 = v7
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
		v101 = v7
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
	v77 = F_SearchSysCache2(m, int32(12), base.I64_extend_i32_u(v18), base.I64_extend_i32_u(l1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L9
	} else {
		goto L52
	}
L13:
	;
	if base.Ui32(l1-int32(_a_F_IsBinaryCoercibleWithCast_0)) <= base.Ui32(int32(1)) {
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
	switch l1 - int32(_a_F_IsBinaryCoercibleWithCast_2) {
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
		v101 = v7
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
		v101 = v7
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
		v101 = v7
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
		v101 = v7
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
		v101 = v7
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
	if v77 == int32(0) {
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
	v83 = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+22)))
	v86 = v84 + v85
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+17)))
	if v87 != int32(98) {
		v96 = v83
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_ReleaseCatCache(m, v77)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L59
	}
L57:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+16)))
	if v90 != int32(105) {
		v96 = v83
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v93
	v96 = int32(1)
	goto L56
L59:
	;
	v101 = v96
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	F_enlargeStringInfo(m, l0, l2)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		if l2 != 0 {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			base.MemoryCopy(m, v6+v7, l1, l2)
		} else {
		}
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v11 = v10 + l2
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v11
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v13+v11))) = uint8(v15)
		return
	}
}
func F_binary_encode(m *base.Module, l0 int32) int64 {
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
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int64
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v293 int32
	_ = v293
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L106
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L102
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L97
	}
L6:
	;
	v245 = int32(1)
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v249 = v247 & v245
	if v249 != 0 {
		goto L78
	} else {
		goto L79
	}
L7:
	;
	v25 = int32(_a_F_binary_encode_0)
	v26 = v21
	goto L9
L8:
	;
	if v63 == int32(0) {
		v244 = int32(_a_F_binary_encode_1)
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
	v70 = int32(_a_F_binary_encode_2)
	v71 = v21
	goto L23
L22:
	;
	if v108 == int32(0) {
		v244 = int32(_a_F_binary_encode_3)
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
	v115 = int32(_a_F_binary_encode_4)
	v116 = v21
	goto L37
L36:
	;
	if v153 == int32(0) {
		v244 = int32(_a_F_binary_encode_5)
		goto L6
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
	v160 = int32(_a_F_binary_encode_6)
	v161 = v21
	goto L51
L50:
	;
	if v198 == int32(0) {
		v244 = int32(_a_F_binary_encode_7)
		goto L6
	} else {
		goto L63
	}
L51:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v164 == v165 {
		v187 = v164
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v198 = int32(0)
	goto L50
L53:
	;
	v189 = int32(1)
	if v187 != 0 {
		v160 = v160 + v189
		v161 = v161 + v189
		goto L51
	} else {
		goto L62
	}
L54:
	;
	if base.Ui32((v164-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v175 = v164 | int32(32)
	goto L57
L56:
	;
	v175 = v164
	goto L57
L57:
	;
	if base.Ui32((v165-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v184 = v165 | int32(32)
	goto L60
L59:
	;
	v184 = v165
	goto L60
L60:
	;
	if v175 == v184 {
		v187 = v175
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v198 = v175 - v184
	goto L50
L62:
	;
	goto L52
L63:
	;
	v204 = int32(_a_F_binary_encode_8)
	v205 = v21
	goto L65
L64:
	;
	if v242 != 0 {
		goto L5
	} else {
		goto L77
	}
L65:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v208 == v209 {
		v231 = v208
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v242 = int32(0)
	goto L64
L67:
	;
	v233 = int32(1)
	if v231 != 0 {
		v204 = v204 + v233
		v205 = v205 + v233
		goto L65
	} else {
		goto L76
	}
L68:
	;
	if base.Ui32((v208-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v219 = v208 | int32(32)
	goto L71
L70:
	;
	v219 = v208
	goto L71
L71:
	;
	if base.Ui32((v209-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v228 = v209 | int32(32)
	goto L74
L73:
	;
	v228 = v209
	goto L74
L74:
	;
	if v219 == v228 {
		v231 = v219
		goto L67
	} else {
		goto L75
	}
L75:
	;
	v242 = v219 - v228
	goto L64
L76:
	;
	goto L66
L77:
	;
	v244 = int32(_a_F_binary_encode_9)
	goto L6
L78:
	;
	v250 = v245
	goto L80
L79:
	;
	v250 = int32(4)
	goto L80
L80:
	;
	v251 = v14 + v250
	if v247 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v280 = m.T0[v279].(func(*base.Module, int32, int32) int64)(m, v251, v278)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L92
	}
L82:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v257 == int32(18) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	v268 = int32(1)
	if v249 != 0 {
		v278 = int32(base.Ui32(v247)>>(uint(v268)%32)) - v268
		goto L81
	} else {
		goto L91
	}
L85:
	;
	v260 = int32(16)
	goto L87
L86:
	;
	v260 = int32(0)
	goto L87
L87:
	;
	if base.Ui32((v257-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v267 = int32(4)
	goto L90
L89:
	;
	v267 = v260
	goto L90
L90:
	;
	v278 = v267
	goto L81
L91:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v278 = int32(base.Ui32(v272)>>(uint(int32(2))%32)) - int32(4)
	goto L81
L92:
	;
	if base.Ui64(int64(1073741820)) <= base.Ui64(v280) {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	v287 = F_palloc(m, base.I32_wrap_i64(v280)+int32(4))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v292 = m.T0[v291].(func(*base.Module, int32, int32, int32) int64)(m, v251, v278, v287+int32(4))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	if base.Ui64(v280) < base.Ui64(v292) {
		goto L3
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = base.I32_wrap_i64(v292)<<(uint(int32(2))%32) + int32(16)
	m.G0 = v11 + int32(48)
	return base.I64_extend_i32_u(v287)
L97:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v21
	F_errmsg(m, int32(_a_F_binary_encode_10), v11+int32(32))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_binary_encode_0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(_a_F_binary_encode_8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(_a_F_binary_encode_4)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(_a_F_binary_encode_2)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_binary_encode_6)
	F_errhint(m, int32(_a_F_binary_encode_11), v11)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_binary_encode_12), int32(69), int32(_a_F_binary_encode_13))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(_a_F_binary_encode_14), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_binary_encode_12), int32(83), int32(_a_F_binary_encode_13))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_errmsg_internal(m, int32(_a_F_binary_encode_15), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_binary_encode_12), int32(91), int32(_a_F_binary_encode_13))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_binary_upgrade_set_next_heap_pg_class_oid(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14232(m, l0, int32(_a_F_binary_upgrade_set_next_heap_pg_class_oid_0), int32(_a_F_binary_upgrade_set_next_heap_pg_class_oid_1), int32(102))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_binary_upgrade_set_next_multirange_pg_type_oid(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14232(m, l0, int32(_a_F_binary_upgrade_set_next_multirange_pg_type_oid_0), int32(_a_F_binary_upgrade_set_next_multirange_pg_type_oid_1), int32(80))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_binary_upgrade_set_next_pg_authid_oid(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14232(m, l0, int32(_a_F_binary_upgrade_set_next_pg_authid_oid_0), int32(_a_F_binary_upgrade_set_next_pg_authid_oid_1), int32(179))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
