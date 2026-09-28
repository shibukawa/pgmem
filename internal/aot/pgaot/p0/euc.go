package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_euc_jis_2004_to_shift_jis_2004(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v111 int32
	_ = v111
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v244 int32
	_ = v244
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_check_encoding_conversion_args(m, v13, v14, v15, int32(5), int32(41))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v15 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_report_invalid_encoding(m, int32(5), v24, v26)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L71
	}
L4:
	;
	v232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v232)
	return base.I64_extend_i32_s(v223 - v12)
L5:
	;
	v223 = v12
	v224 = v11
	goto L4
L6:
	;
	goto L7
L7:
	;
	v24 = v12
	v25 = v11
	v26 = v15
	goto L8
L8:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v34 = base.I32_extend8_s(v33)
	if int32(0) <= v34 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v223 = v220
	v224 = v214
	goto L4
L10:
	;
	if int32(0) < v215 {
		v24 = v220
		v25 = v214
		v26 = v215
		goto L8
	} else {
		goto L70
	}
L11:
	;
	if v34 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v49 = F_pg_encoding_verifymbchar(m, int32(5), v24, v26)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L18
	}
L14:
	;
	if v10 != int64(0) {
		v223 = v24
		v224 = v25
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v34)
	v42 = int32(1)
	v214 = v25 + v42
	v215 = v26 - v42
	v220 = v24 + v42
	goto L10
L17:
	;
	goto L3
L18:
	;
	if v49 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v10 != int64(0) {
		v223 = v24
		v224 = v25
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if base.B2i32(v34 != int32(-114))|base.B2i32(v49 != int32(2)) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L3
L23:
	;
	v214 = v211
	v215 = v26 - v49
	v220 = v24 + v49
	goto L10
L24:
	;
	v211 = v201 + int32(1)
	goto L23
L25:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v62)
	v201 = v25
	goto L24
L26:
	;
	goto L27
L27:
	;
	if base.B2i32(v34 != int32(-113))|base.B2i32(v49 != int32(3)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if v72&int32(1) != 0 {
		goto L60
	} else {
		goto L61
	}
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v170)
	v174 = v25 + int32(1)
	goto L28
L30:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v74 = v72 - int32(161)
	v81 = int32(0)
	if base.B2i32(base.Ui32(int32(14)) < base.Ui32(v74))|base.B2i32(int32(1)<<(uint(v74)%32)&int32(_a_F_euc_jis_2004_to_shift_jis_2004_0) == v81) == v81 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	if v49 == int32(2) {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	v170 = int32(base.Ui32(v72+int32(1888))>>(uint(int32(3))%32))*int32(253) + int32(base.Ui32(v72+int32(319))>>(uint(int32(1))%32))
	goto L29
L34:
	;
	goto L35
L35:
	;
	if base.Ui32((v72+int32(18))&int32(255)) <= base.Ui32(int32(16)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v170 = int32(base.Ui32(v72+int32(251)) >> (uint(int32(1)) % 32))
	goto L29
L37:
	;
	goto L38
L38:
	;
	if v10 == int64(0) {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	v174 = v25
	goto L28
L40:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if base.Ui32((v34+int32(95))&int32(255)) < base.Ui32(int32(62)) {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	goto L42
L42:
	;
	if v10 != int64(0) {
		v223 = v24
		v224 = v25
		goto L4
	} else {
		goto L59
	}
L43:
	;
	v163 = int32(2)
	v164 = v111 - v163
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)) = uint8(v164)
	v211 = v25 + v163
	goto L23
L44:
	;
	if base.Ui32((v111+int32(32))&int32(255)) <= base.Ui32(int32(30)) {
		goto L53
	} else {
		goto L54
	}
L45:
	;
	if v10 != int64(0) {
		v223 = v24
		v224 = v25
		goto L4
	} else {
		goto L52
	}
L46:
	;
	v126 = int32(97)
	goto L48
L47:
	;
	if base.Ui32(int32(32)) <= base.Ui32((v34+int32(33))&int32(255)) {
		goto L45
	} else {
		goto L49
	}
L48:
	;
	v128 = int32(1)
	v129 = int32(base.Ui32(v126+v33) >> (uint(v128) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v129)
	if v33&v128 == int32(0) {
		goto L43
	} else {
		goto L50
	}
L49:
	;
	v126 = int32(225)
	goto L48
L50:
	;
	if base.Ui32(int32(62)) < base.Ui32((v111+int32(95))&int32(255)) {
		goto L44
	} else {
		goto L51
	}
L51:
	;
	v142 = v111 - int32(97)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)) = uint8(v142)
	v211 = v25 + int32(2)
	goto L23
L52:
	;
	goto L3
L53:
	;
	v155 = v111 - int32(96)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)) = uint8(v155)
	v211 = v25 + int32(2)
	goto L23
L54:
	;
	goto L55
L55:
	;
	if v10 != int64(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v223 = v24
	v224 = v25 + int32(1)
	goto L4
L57:
	;
	goto L58
L58:
	;
	goto L3
L59:
	;
	goto L3
L60:
	;
	if base.Ui32((v71+int32(95))&int32(255)) <= base.Ui32(int32(62)) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v199 = v71 - int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v199)
	v201 = v174
	goto L24
L63:
	;
	v185 = v71 - int32(97)
	*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v185)
	v201 = v174
	goto L24
L64:
	;
	goto L65
L65:
	;
	if base.Ui32((v71+int32(32))&int32(255)) <= base.Ui32(int32(30)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v194 = v71 - int32(96)
	*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v194)
	v201 = v174
	goto L24
L67:
	;
	goto L68
L68:
	;
	if v10 != int64(0) {
		v223 = v24
		v224 = v174
		goto L4
	} else {
		goto L69
	}
L69:
	;
	goto L3
L70:
	;
	goto L9
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_euc_jp_to_sjis(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_check_encoding_conversion_args(m, v14, v15, v16, int32(1), int32(35))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v16 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v217))) = uint8(v225)
	return base.I64_extend_i32_s(v215 - v13)
L4:
	;
	v215 = v13
	v217 = v12
	goto L3
L5:
	;
	goto L6
L6:
	;
	v25 = v13
	v27 = v12
	v30 = v16
	goto L7
L7:
	;
	v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v25))))
	if int32(0) <= v35 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v215 = v212
	v217 = v204
	goto L3
L9:
	;
	if int32(0) < v207 {
		v25 = v212
		v27 = v204
		v30 = v207
		goto L7
	} else {
		goto L62
	}
L10:
	;
	if v35 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v53 = F_pg_encoding_verifymbchar(m, int32(1), v25, v30)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v11 != int64(0) {
		v215 = v25
		v217 = v27
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v35)
	v46 = int32(1)
	v204 = v27 + v46
	v207 = v30 - v46
	v212 = v25 + v46
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(1), v25, v30)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	if v53 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v11 != int64(0) {
		v215 = v25
		v217 = v27
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v64 = v35 & int32(255)
	switch v64 - int32(142) {
	case 0:
		goto L28
	case 1:
		goto L27
	default:
		goto L26
	}
L22:
	;
	F_report_invalid_encoding(m, int32(1), v25, v30)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v204 = v199
	v207 = v30 - v53
	v212 = v25 + v53
	goto L9
L25:
	;
	v199 = v27 + int32(2)
	goto L24
L26:
	;
	if base.Ui32(int32(_a_F_euc_jp_to_sjis_0)) <= base.Ui32(v64<<(uint(int32(8))%32)|v62) {
		goto L50
	} else {
		goto L51
	}
L27:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)))
	v74 = v71 | v62<<(uint(int32(8))%32)
	if base.Ui32(int32(_a_F_euc_jp_to_sjis_0)) <= base.Ui32(v74) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v62)
	v199 = v27 + int32(1)
	goto L24
L29:
	;
	v82 = int32(base.Ui32(v62+int32(267))>>(uint(int32(1))%32)) - int32(11)
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v82)
	if base.Ui32(v71) < base.Ui32(int32(224)) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v102 = int32(0)
	goto L42
L32:
	;
	v88 = int32(-97)
	goto L34
L33:
	;
	v88 = int32(-96)
	goto L34
L34:
	;
	if v62&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v92 = v88
	goto L37
L36:
	;
	v92 = int32(-2)
	goto L37
L37:
	;
	v93 = v92 + v71
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v93)
	goto L25
L38:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+2)))
	v134 = int32(8)
	v138 = v133<<(uint(v134)%32) | int32(base.Ui32(v133)>>(uint(v134)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v138)
	goto L25
L39:
	;
	v131 = v106 + int32(_a_F_euc_jp_to_sjis_1)
	goto L38
L40:
	;
	v131 = v106 + int32(_a_F_euc_jp_to_sjis_2)
	goto L38
L41:
	;
	v131 = v106 + int32(_a_F_euc_jp_to_sjis_3)
	goto L38
L42:
	;
	v106 = v102 << (uint(int32(3)) % 32)
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+uint32(_c_F_euc_jp_to_sjis[0]))))
	if v109 == v74 {
		v131 = v106 + int32(_a_F_euc_jp_to_sjis_4)
		goto L38
	} else {
		goto L44
	}
L43:
	;
	v123 = int32(_a_F_euc_jp_to_sjis_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v123)
	goto L25
L44:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+uint32(_c_F_euc_jp_to_sjis[1]))))
	if v113 == v74 {
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+uint32(_c_F_euc_jp_to_sjis[2]))))
	if v115 == v74 {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+uint32(_c_F_euc_jp_to_sjis[3]))))
	if v117 == v74 {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v120 = v102 + int32(4)
	if v120 != int32(388) {
		v102 = v120
		goto L42
	} else {
		goto L48
	}
L48:
	;
	goto L43
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v163)
	if base.Ui32(v62) < base.Ui32(int32(224)) {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	v163 = (v64-int32(245))>>(uint(int32(1))%32) + int32(240)
	v164 = v64 - int32(84)
	goto L49
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(v35) < base.Ui32(int32(-33)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v161 = int32(129)
	goto L55
L54:
	;
	v161 = int32(193)
	goto L55
L55:
	;
	v163 = int32(base.Ui32(v64+int32(351))>>(uint(int32(1))%32)) + v161
	v164 = v64
	goto L49
L56:
	;
	v170 = int32(-97)
	goto L58
L57:
	;
	v170 = int32(-96)
	goto L58
L58:
	;
	if v164&int32(1) != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v174 = v170
	goto L61
L60:
	;
	v174 = int32(-2)
	goto L61
L61:
	;
	v175 = v174 + v62
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v175)
	goto L25
L62:
	;
	goto L8
}
