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
func F_binary_decode(m *base.Module, l0 int32) int64 {
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
	v25 = int32(_a_F_binary_decode_0)
	v26 = v21
	goto L9
L8:
	;
	if v63 == int32(0) {
		v244 = int32(_a_F_binary_decode_1)
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
	v70 = int32(_a_F_binary_decode_2)
	v71 = v21
	goto L23
L22:
	;
	if v108 == int32(0) {
		v244 = int32(_a_F_binary_decode_3)
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
	v115 = int32(_a_F_binary_decode_4)
	v116 = v21
	goto L37
L36:
	;
	if v153 == int32(0) {
		v244 = int32(_a_F_binary_decode_5)
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
	v160 = int32(_a_F_binary_decode_6)
	v161 = v21
	goto L51
L50:
	;
	if v198 == int32(0) {
		v244 = int32(_a_F_binary_decode_7)
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
	v204 = int32(_a_F_binary_decode_8)
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
	v244 = int32(_a_F_binary_decode_9)
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
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
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
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v244)+16))
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
	F_errmsg(m, int32(_a_F_binary_decode_10), v11+int32(32))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_binary_decode_0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(_a_F_binary_decode_8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(_a_F_binary_decode_4)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(_a_F_binary_decode_2)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_binary_decode_6)
	F_errhint(m, int32(_a_F_binary_decode_11), v11)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_binary_decode_12), int32(119), int32(_a_F_binary_decode_13))
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
	F_errmsg(m, int32(_a_F_binary_decode_14), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_binary_decode_12), int32(133), int32(_a_F_binary_decode_13))
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
	F_errmsg_internal(m, int32(_a_F_binary_decode_15), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_binary_decode_12), int32(141), int32(_a_F_binary_decode_13))
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
func F_binary_upgrade_set_next_heap_relfilenode(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14232(m, l0, int32(_a_F_binary_upgrade_set_next_heap_relfilenode_0), int32(_a_F_binary_upgrade_set_next_heap_relfilenode_1), int32(113))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_binary_upgrade_set_next_index_pg_class_oid(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14232(m, l0, int32(_a_F_binary_upgrade_set_next_index_pg_class_oid_0), int32(_a_F_binary_upgrade_set_next_index_pg_class_oid_1), int32(124))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_binary_upgrade_set_next_multirange_array_pg_type_oid(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14232(m, l0, int32(_a_F_binary_upgrade_set_next_multirange_array_pg_type_oid_0), int32(_a_F_binary_upgrade_set_next_multirange_array_pg_type_oid_1), int32(91))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_executeBinaryArithmExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int64
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int64
	_ = v477
	var v479 int64
	_ = v479
	var v481 int64
	_ = v481
	var v483 int64
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v506 int64
	_ = v506
	var v508 int64
	_ = v508
	var v510 int64
	_ = v510
	var v512 int64
	_ = v512
	var v523 int32
	_ = v523
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(256)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+152)) = v6
	v16 = int64(8589934592)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+144)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v16
	v23 = v12 + int32(144)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v12 - int32(-64)
	v29 = v12 + int32(228)
	F_jspGetArg(m, l1, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v35 = F_executeItemOptUnwrapResult(m, l0, v29, l2, int32(1), v23)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v12 + int32(256)
	return v523
L4:
	;
	if v35 == int32(2) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	if v39 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v85 = v12 + int32(228)
	F_jspGetRightArg(m, l1, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L20
	}
L8:
	;
	v42 = v39
	goto L11
L9:
	;
	goto L10
L10:
	;
	v61 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+152)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v12 + int32(144)
	v68 = int32(2)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if v69 == v61 {
		v523 = v68
		goto L3
	} else {
		goto L15
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	F_pfree(m, v42)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	if v49 != 0 {
		v42 = v49
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v74 = v69
	goto L16
L16:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	F_pfree(m, v74)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	v523 = v68
	goto L3
L18:
	;
	if v81 != 0 {
		v74 = v81
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v91 = F_executeItemOptUnwrapResult(m, l0, v85, l2, int32(1), v12-int32(-64))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v91 == int32(2) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	if v95 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
	if v140 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L25:
	;
	v98 = v95
	goto L28
L26:
	;
	goto L27
L27:
	;
	v117 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+152)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v12 + int32(144)
	v124 = int32(2)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if v125 == v117 {
		v523 = v124
		goto L3
	} else {
		goto L32
	}
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	F_pfree(m, v98)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L27
L30:
	;
	if v105 != 0 {
		v98 = v105
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v130 = v125
	goto L33
L33:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	F_pfree(m, v130)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	v523 = v124
	goto L3
L35:
	;
	if v137 != 0 {
		v130 = v137
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	if v229 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L38:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
	if v143 == int32(2) {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	if v146 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v149 = v146
	goto L45
L43:
	;
	goto L44
L44:
	;
	v168 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+152)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v12 + int32(144)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if v175 != 0 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
	F_pfree(m, v149)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	goto L44
L47:
	;
	if v156 != 0 {
		v149 = v156
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v178 = v175
	goto L52
L50:
	;
	goto L51
L51:
	;
	v197 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v12 - int32(-64)
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v204 != int32(1) {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
	F_pfree(m, v178)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	goto L51
L54:
	;
	if v185 != 0 {
		v178 = v185
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v523 = int32(2)
	goto L3
L57:
	;
	goto L58
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(135004290))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v216 = F_jspOperationName(m, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v216
	F_errmsg(m, int32(_a_F_executeBinaryArithmExpr_0), v12+int32(16))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_executeBinaryArithmExpr_1), int32(2209), int32(_a_F_executeBinaryArithmExpr_2))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v316 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L65:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if v232 == int32(2) {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	if v235 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L67
L69:
	;
	v238 = v235
	goto L72
L70:
	;
	goto L71
L71:
	;
	v257 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+152)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v12 + int32(144)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if v264 != 0 {
		goto L76
	} else {
		goto L77
	}
L72:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	F_pfree(m, v238)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L71
L74:
	;
	if v245 != 0 {
		v238 = v245
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v267 = v264
	goto L79
L77:
	;
	goto L78
L78:
	;
	v286 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v12 - int32(-64)
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v293 != int32(1) {
		goto L83
	} else {
		goto L84
	}
L79:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v267)+8))
	F_pfree(m, v267)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L81
	}
L80:
	;
	goto L78
L81:
	;
	if v274 != 0 {
		v267 = v274
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v523 = int32(2)
	goto L3
L84:
	;
	goto L85
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errcode(m, int32(135004290))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v305 = F_jspOperationName(m, v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v305
	F_errmsg(m, int32(_a_F_executeBinaryArithmExpr_3), v12)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_executeBinaryArithmExpr_1), int32(2220), int32(_a_F_executeBinaryArithmExpr_2))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	if v392 != 0 {
		goto L112
	} else {
		goto L113
	}
L92:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v12)+168))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	v322 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v319, v320, int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_executeBinaryArithmExpr[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v325
	v328 = *(*int64)(unsafe.Add(mBase, _c_F_executeBinaryArithmExpr[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v328
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v12)+168))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	v334 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v330, v331, v12+int32(32))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L96
	}
L95:
	;
	v391 = v322
	goto L91
L96:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+36)))
	if v336 != int32(1) {
		v391 = v334
		goto L91
	} else {
		goto L97
	}
L97:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	if v339 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v342 = v339
	goto L101
L99:
	;
	goto L100
L100:
	;
	v361 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+152)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v12 + int32(144)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if v368 != 0 {
		goto L105
	} else {
		goto L106
	}
L101:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	F_pfree(m, v342)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L103
	}
L102:
	;
	goto L100
L103:
	;
	if v349 != 0 {
		v342 = v349
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v371 = v368
	goto L108
L106:
	;
	goto L107
L107:
	;
	v523 = int32(2)
	goto L3
L108:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v371)+8))
	F_pfree(m, v371)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L110
	}
L109:
	;
	goto L107
L110:
	;
	if v378 != 0 {
		v371 = v378
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v395 = v392
	goto L115
L113:
	;
	goto L114
L114:
	;
	v414 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+152)) = v414
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v414
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v12 + int32(144)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if v421 != 0 {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v395)+8))
	F_pfree(m, v395)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L117
	}
L116:
	;
	goto L114
L117:
	;
	if v402 != 0 {
		v395 = v402
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v424 = v421
	goto L122
L120:
	;
	goto L121
L121:
	;
	v443 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v12 - int32(-64)
	v451 = v12 + int32(228)
	v452 = F_jspGetNext(m, l1, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L126
	}
L122:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v424)+8))
	F_pfree(m, v424)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	goto L121
L124:
	;
	if v431 != 0 {
		v424 = v431
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	if v452|l4 == int32(0) {
		v523 = v6
		goto L3
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(2)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v460 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v466 = F_executeItemOptUnwrapTarget(m, l0, v451, v12+int32(32), l4, v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	if l4 == int32(0) {
		v523 = v6
		goto L3
	} else {
		goto L132
	}
L131:
	;
	v523 = v466
	goto L3
L132:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v471 < v472 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v476 = v470 + v471<<(uint(int32(5))%32)
	v477 = *(*int64)(unsafe.Add(mBase, uint32(v12)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v476)+40)) = v477
	v479 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v476)+32)) = v479
	v481 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v476)+24)) = v481
	v483 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v476)+16)) = v483
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = v485 + int32(1)
	v523 = v6
	goto L3
L134:
	;
	goto L135
L135:
	;
	v489 = int32(16)
	v491 = v472 << (uint(int32(1)) % 32)
	if v491 <= v489 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v494 = v489
	goto L138
L137:
	;
	v494 = v491
	goto L138
L138:
	;
	v499 = F_palloc(m, v494<<(uint(int32(5))%32)|int32(16))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v499)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v499)+4)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v499))) = int32(1)
	v506 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v499)+16)) = v506
	v508 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v499)+24)) = v508
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v499)+32)) = v510
	v512 = *(*int64)(unsafe.Add(mBase, uint32(v12)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v499)+40)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v470)+8)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v499
	v523 = v6
	goto L3
}
