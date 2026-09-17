package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_euc_jis_2004_to_shift_jis_2004(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
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
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v107 int32
	_ = v107
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v231 int32
	_ = v231
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
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
	return int32(0)
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
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L71
	}
L4:
	;
	v220 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v212))) = uint8(v220)
	return v211 - v12
L5:
	;
	v211 = v12
	v212 = v11
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
	v211 = v208
	v212 = v202
	goto L4
L10:
	;
	if int32(0) < v203 {
		v24 = v208
		v25 = v202
		v26 = v203
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
	v47 = F_pg_encoding_verifymbchar(m, int32(5), v24, v26)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L18
	}
L14:
	;
	if v10 != 0 {
		v211 = v24
		v212 = v25
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
	v40 = int32(1)
	v202 = v25 + v40
	v203 = v26 - v40
	v208 = v24 + v40
	goto L10
L17:
	;
	goto L3
L18:
	;
	if v47 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v10 != 0 {
		v211 = v24
		v212 = v25
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if base.B2i32(v34 != int32(-114))|base.B2i32(v47 != int32(2)) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L3
L23:
	;
	v202 = v199
	v203 = v26 - v47
	v208 = v24 + v47
	goto L10
L24:
	;
	v199 = v189 + int32(1)
	goto L23
L25:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v58)
	v189 = v25
	goto L24
L26:
	;
	goto L27
L27:
	;
	if base.B2i32(v34 != int32(-113))|base.B2i32(v47 != int32(3)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if v68&int32(1) != 0 {
		goto L60
	} else {
		goto L61
	}
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v160)
	v164 = v25 + int32(1)
	goto L28
L30:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v70 = v68 - int32(161)
	v77 = int32(0)
	if base.B2i32(base.Ui32(int32(14)) < base.Ui32(v70))|base.B2i32(int32(1)<<(uint(v70)%32)&int32(_a_F_euc_jis_2004_to_shift_jis_2004_0) == v77) == v77 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	if v47 == int32(2) {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	v160 = int32(base.Ui32(v68+int32(1888))>>(uint(int32(3))%32))*int32(253) + int32(base.Ui32(v68+int32(319))>>(uint(int32(1))%32))
	goto L29
L34:
	;
	goto L35
L35:
	;
	if base.Ui32((v68+int32(18))&int32(255)) <= base.Ui32(int32(16)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v160 = int32(base.Ui32(v68+int32(251)) >> (uint(int32(1)) % 32))
	goto L29
L37:
	;
	goto L38
L38:
	;
	if v10 == int32(0) {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	v164 = v25
	goto L28
L40:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
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
	if v10 != 0 {
		v211 = v24
		v212 = v25
		goto L4
	} else {
		goto L59
	}
L43:
	;
	v155 = int32(2)
	v156 = v107 - v155
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)) = uint8(v156)
	v199 = v25 + v155
	goto L23
L44:
	;
	if base.Ui32((v107+int32(32))&int32(255)) <= base.Ui32(int32(30)) {
		goto L53
	} else {
		goto L54
	}
L45:
	;
	if v10 != 0 {
		v211 = v24
		v212 = v25
		goto L4
	} else {
		goto L52
	}
L46:
	;
	v122 = int32(97)
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
	v124 = int32(1)
	v125 = int32(base.Ui32(v122+v33) >> (uint(v124) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v125)
	if v33&v124 == int32(0) {
		goto L43
	} else {
		goto L50
	}
L49:
	;
	v122 = int32(225)
	goto L48
L50:
	;
	if base.Ui32(int32(62)) < base.Ui32((v107+int32(95))&int32(255)) {
		goto L44
	} else {
		goto L51
	}
L51:
	;
	v138 = v107 - int32(97)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)) = uint8(v138)
	v199 = v25 + int32(2)
	goto L23
L52:
	;
	goto L3
L53:
	;
	v149 = v107 - int32(96)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)) = uint8(v149)
	v199 = v25 + int32(2)
	goto L23
L54:
	;
	goto L55
L55:
	;
	if v10 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v211 = v24
	v212 = v25 + int32(1)
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
	if base.Ui32((v67+int32(95))&int32(255)) <= base.Ui32(int32(62)) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v187 = v67 - int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v187)
	v189 = v164
	goto L24
L63:
	;
	v175 = v67 - int32(97)
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v175)
	v189 = v164
	goto L24
L64:
	;
	goto L65
L65:
	;
	if base.Ui32((v67+int32(32))&int32(255)) <= base.Ui32(int32(30)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v184 = v67 - int32(96)
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v184)
	v189 = v164
	goto L24
L67:
	;
	goto L68
L68:
	;
	if v10 != 0 {
		v211 = v24
		v212 = v164
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
func F_euc_jp_to_sjis(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
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
	return int32(0)
L2:
	;
	if v16 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v221 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v221)
	return v211 - v13
L4:
	;
	v211 = v13
	v213 = v12
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
	v211 = v208
	v213 = v200
	goto L3
L9:
	;
	if int32(0) < v203 {
		v25 = v208
		v27 = v200
		v30 = v203
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
	v51 = F_pg_encoding_verifymbchar(m, int32(1), v25, v30)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v11 != 0 {
		v211 = v25
		v213 = v27
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
	v44 = int32(1)
	v200 = v27 + v44
	v203 = v30 - v44
	v208 = v25 + v44
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(1), v25, v30)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
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
	if v51 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v11 != 0 {
		v211 = v25
		v213 = v27
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v60 = v35 & int32(255)
	switch v60 - int32(142) {
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
	v57 = m.ExcPending
	if v57 != 0 {
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
	v200 = v195
	v203 = v30 - v51
	v208 = v25 + v51
	goto L9
L25:
	;
	v195 = v27 + int32(2)
	goto L24
L26:
	;
	if base.Ui32(int32(_a_F_euc_jp_to_sjis_0)) <= base.Ui32(v60<<(uint(int32(8))%32)|v58) {
		goto L50
	} else {
		goto L51
	}
L27:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)))
	v70 = v67 | v58<<(uint(int32(8))%32)
	if base.Ui32(int32(_a_F_euc_jp_to_sjis_0)) <= base.Ui32(v70) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v58)
	v195 = v27 + int32(1)
	goto L24
L29:
	;
	v78 = int32(base.Ui32(v58+int32(267))>>(uint(int32(1))%32)) - int32(11)
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v78)
	if base.Ui32(v67) < base.Ui32(int32(224)) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v98 = int32(0)
	goto L42
L32:
	;
	v84 = int32(-97)
	goto L34
L33:
	;
	v84 = int32(-96)
	goto L34
L34:
	;
	if v58&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v88 = v84
	goto L37
L36:
	;
	v88 = int32(-2)
	goto L37
L37:
	;
	v89 = v88 + v67
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v89)
	goto L25
L38:
	;
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127)+2)))
	v130 = int32(8)
	v134 = v129<<(uint(v130)%32) | int32(base.Ui32(v129)>>(uint(v130)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v134)
	goto L25
L39:
	;
	v127 = v102 + int32(_a_F_euc_jp_to_sjis_1)
	goto L38
L40:
	;
	v127 = v102 + int32(_a_F_euc_jp_to_sjis_2)
	goto L38
L41:
	;
	v127 = v102 + int32(_a_F_euc_jp_to_sjis_3)
	goto L38
L42:
	;
	v102 = v98 << (uint(int32(3)) % 32)
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_euc_jp_to_sjis[0]))))
	if v105 == v70 {
		v127 = v102 + int32(_a_F_euc_jp_to_sjis_4)
		goto L38
	} else {
		goto L44
	}
L43:
	;
	v119 = int32(_a_F_euc_jp_to_sjis_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v119)
	goto L25
L44:
	;
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_euc_jp_to_sjis[1]))))
	if v109 == v70 {
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_euc_jp_to_sjis[2]))))
	if v111 == v70 {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_euc_jp_to_sjis[3]))))
	if v113 == v70 {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v116 = v98 + int32(4)
	if v116 != int32(388) {
		v98 = v116
		goto L42
	} else {
		goto L48
	}
L48:
	;
	goto L43
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v159)
	if base.Ui32(v58) < base.Ui32(int32(224)) {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	v159 = (v60-int32(245))>>(uint(int32(1))%32) + int32(240)
	v160 = v60 - int32(84)
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
	v157 = int32(129)
	goto L55
L54:
	;
	v157 = int32(193)
	goto L55
L55:
	;
	v159 = int32(base.Ui32(v60+int32(351))>>(uint(int32(1))%32)) + v157
	v160 = v60
	goto L49
L56:
	;
	v166 = int32(-97)
	goto L58
L57:
	;
	v166 = int32(-96)
	goto L58
L58:
	;
	if v160&int32(1) != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v170 = v166
	goto L61
L60:
	;
	v170 = int32(-2)
	goto L61
L61:
	;
	v171 = v170 + v58
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v171)
	goto L25
L62:
	;
	goto L8
}
