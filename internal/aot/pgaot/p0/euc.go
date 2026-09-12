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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v104 int32
	_ = v104
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
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
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v232)
	return v223 - v12
L4:
	;
	v223 = v12
	v224 = v11
	goto L3
L5:
	;
	goto L6
L6:
	;
	v24 = v12
	v25 = v11
	v26 = v15
	goto L7
L7:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v34 = base.I32_extend8_s(v33)
	if int32(0) <= v34 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v223 = v220
	v224 = v214
	goto L3
L9:
	;
	if int32(0) < v215 {
		v24 = v220
		v25 = v214
		v26 = v215
		goto L7
	} else {
		goto L77
	}
L10:
	;
	if v34 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v50 = F_pg_encoding_verifymbchar(m, int32(5), v24, v26)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v10 != 0 {
		v223 = v24
		v224 = v25
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v34)
	v43 = int32(1)
	v214 = v25 + v43
	v215 = v26 - v43
	v220 = v24 + v43
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(5), v24, v26)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
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
	if v50 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v10 != 0 {
		v223 = v24
		v224 = v25
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v34 != int32(-114) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	F_report_invalid_encoding(m, int32(5), v24, v26)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
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
	v214 = v211
	v215 = v26 - v50
	v220 = v24 + v50
	goto L9
L25:
	;
	v211 = v201 + int32(1)
	goto L24
L26:
	;
	if v34 != int32(-113) {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	if v50 != int32(2) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v61)
	v201 = v25
	goto L25
L29:
	;
	if v68&int32(1) != 0 {
		goto L66
	} else {
		goto L67
	}
L30:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v169)
	v173 = v25 + int32(1)
	goto L29
L31:
	;
	F_report_invalid_encoding(m, int32(5), v24, v26)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L65
	}
L32:
	;
	if v50 == int32(2) {
		goto L42
	} else {
		goto L43
	}
L33:
	;
	if v50 != int32(3) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v70 = v68 - int32(161)
	if base.Ui32(int32(14)) < base.Ui32(v70) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if base.Ui32((v68+int32(18))&int32(255)) <= base.Ui32(int32(16)) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	if int32(1)<<(uint(v70)%32)&int32(30877) == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v169 = int32(base.Ui32(v68+int32(1888))>>(uint(int32(3))%32))*int32(253) + int32(base.Ui32(v68+int32(319))>>(uint(int32(1))%32))
	goto L30
L38:
	;
	v169 = int32(base.Ui32(v68+int32(251)) >> (uint(int32(1)) % 32))
	goto L30
L39:
	;
	goto L40
L40:
	;
	if v10 == int32(0) {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	v173 = v25
	goto L29
L42:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if base.Ui32((v34+int32(95))&int32(255)) < base.Ui32(int32(62)) {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	goto L44
L44:
	;
	if v10 != 0 {
		v223 = v24
		v224 = v25
		goto L3
	} else {
		goto L63
	}
L45:
	;
	v158 = int32(2)
	v159 = v104 - v158
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)) = uint8(v159)
	v211 = v25 + v158
	goto L24
L46:
	;
	if base.Ui32((v104+int32(32))&int32(255)) <= base.Ui32(int32(30)) {
		goto L56
	} else {
		goto L57
	}
L47:
	;
	if v10 != 0 {
		v223 = v24
		v224 = v25
		goto L3
	} else {
		goto L54
	}
L48:
	;
	v119 = int32(97)
	goto L50
L49:
	;
	if base.Ui32(int32(32)) <= base.Ui32((v34+int32(33))&int32(255)) {
		goto L47
	} else {
		goto L51
	}
L50:
	;
	v121 = int32(1)
	v122 = int32(base.Ui32(v119+v33) >> (uint(v121) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v122)
	if v33&v121 == int32(0) {
		goto L45
	} else {
		goto L52
	}
L51:
	;
	v119 = int32(225)
	goto L50
L52:
	;
	if base.Ui32(int32(62)) < base.Ui32((v104+int32(95))&int32(255)) {
		goto L46
	} else {
		goto L53
	}
L53:
	;
	v135 = v104 - int32(97)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)) = uint8(v135)
	v211 = v25 + int32(2)
	goto L24
L54:
	;
	F_report_invalid_encoding(m, int32(5), v24, v26)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
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
	v149 = v104 - int32(96)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)) = uint8(v149)
	v211 = v25 + int32(2)
	goto L24
L57:
	;
	goto L58
L58:
	;
	if v10 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v223 = v24
	v224 = v25 + int32(1)
	goto L3
L60:
	;
	goto L61
L61:
	;
	F_report_invalid_encoding(m, int32(5), v24, v26)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
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
	F_report_invalid_encoding(m, int32(5), v24, v26)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	if base.Ui32((v67+int32(95))&int32(255)) <= base.Ui32(int32(62)) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v199 = v67 - int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v199)
	v201 = v173
	goto L25
L69:
	;
	v184 = v67 - int32(97)
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v184)
	v201 = v173
	goto L25
L70:
	;
	goto L71
L71:
	;
	if base.Ui32((v67+int32(32))&int32(255)) <= base.Ui32(int32(30)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v193 = v67 - int32(96)
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v193)
	v201 = v173
	goto L25
L73:
	;
	goto L74
L74:
	;
	if v10 != 0 {
		v223 = v24
		v224 = v173
		goto L3
	} else {
		goto L75
	}
L75:
	;
	F_report_invalid_encoding(m, int32(5), v24, v26)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
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
	goto L8
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
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
	v220 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v211))) = uint8(v220)
	return v210 - v13
L4:
	;
	v210 = v13
	v211 = v12
	goto L3
L5:
	;
	goto L6
L6:
	;
	v25 = v13
	v26 = v12
	v29 = v16
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
	v210 = v207
	v211 = v198
	goto L3
L9:
	;
	if int32(0) < v201 {
		v25 = v207
		v26 = v198
		v29 = v201
		goto L7
	} else {
		goto L60
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
	v51 = F_pg_encoding_verifymbchar(m, int32(1), v25, v29)
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
		v210 = v25
		v211 = v26
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v35)
	v44 = int32(1)
	v198 = v26 + v44
	v201 = v29 - v44
	v207 = v25 + v44
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(1), v25, v29)
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
		v210 = v25
		v211 = v26
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
	F_report_invalid_encoding(m, int32(1), v25, v29)
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
	v198 = v194
	v201 = v29 - v51
	v207 = v25 + v51
	goto L9
L25:
	;
	v194 = v26 + int32(2)
	goto L24
L26:
	;
	if base.Ui32(int32(62881)) <= base.Ui32(v60<<(uint(int32(8))%32)|v58) {
		goto L48
	} else {
		goto L49
	}
L27:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)))
	v69 = v66 | v58<<(uint(int32(8))%32)
	if base.Ui32(v69) <= base.Ui32(int32(62880)) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v58)
	v194 = v26 + int32(1)
	goto L24
L29:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+2)))
	v129 = int32(8)
	v133 = v128<<(uint(v129)%32) | int32(base.Ui32(v128)>>(uint(v129)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v133)
	goto L25
L30:
	;
	v98 = int32(0)
	goto L41
L31:
	;
	if v69 != int32(62451) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v83 = int32(base.Ui32(v58+int32(267))>>(uint(int32(1))%32)) - int32(11)
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v83)
	if base.Ui32(v66) < base.Ui32(int32(224)) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v124 = int32(2213872)
	goto L29
L35:
	;
	v89 = int32(-97)
	goto L37
L36:
	;
	v89 = int32(-96)
	goto L37
L37:
	;
	if v58&int32(1) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v93 = v89
	goto L40
L39:
	;
	v93 = int32(-2)
	goto L40
L40:
	;
	v94 = v93 + v66
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)) = uint8(v94)
	goto L25
L41:
	;
	v107 = v98 + int32(1)
	v109 = v107 << (uint(int32(3)) % 32)
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+uint32(_consts[1073]))))
	if v112 == int32(65535) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v124 = v109 + int32(2213872)
	goto L29
L43:
	;
	v115 = int32(44161)
	*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v115)
	goto L25
L44:
	;
	goto L45
L45:
	;
	if v112 != v69 {
		v98 = v107
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L42
L47:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v158)
	if base.Ui32(v58) < base.Ui32(int32(224)) {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	v158 = (v60-int32(245))>>(uint(int32(1))%32) + int32(240)
	v159 = v60 - int32(84)
	goto L47
L49:
	;
	goto L50
L50:
	;
	if base.Ui32(v35) < base.Ui32(int32(-33)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v156 = int32(129)
	goto L53
L52:
	;
	v156 = int32(193)
	goto L53
L53:
	;
	v158 = int32(base.Ui32(v60+int32(351))>>(uint(int32(1))%32)) + v156
	v159 = v60
	goto L47
L54:
	;
	v165 = int32(-97)
	goto L56
L55:
	;
	v165 = int32(-96)
	goto L56
L56:
	;
	if v159&int32(1) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v169 = v165
	goto L59
L58:
	;
	v169 = int32(-2)
	goto L59
L59:
	;
	v170 = v169 + v58
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)) = uint8(v170)
	goto L25
L60:
	;
	goto L8
}
