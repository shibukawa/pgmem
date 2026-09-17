package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_euc_jis_2004_to_utf8(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13848(m, l0, int32(5), int32(0), int32(25), int32(_a_F_euc_jis_2004_to_utf8_0), int32(_a_F_euc_jis_2004_to_utf8_1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_euc_tw_to_big5(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v12, v13, v14, int32(4), int32(36))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v14 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v171 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v171)
	return v163 - v11
L4:
	;
	v163 = v11
	v166 = v10
	goto L3
L5:
	;
	goto L6
L6:
	;
	v23 = v11
	v24 = v14
	v26 = v10
	goto L7
L7:
	;
	v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23))))
	if v31 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v163 = v160
	v166 = v157
	goto L3
L9:
	;
	if int32(0) < v155 {
		v23 = v160
		v24 = v155
		v26 = v157
		goto L7
	} else {
		goto L65
	}
L10:
	;
	v35 = F_pg_encoding_verifymbchar(m, int32(4), v23, v24)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v31 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L13:
	;
	if v35 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v9 != 0 {
		v163 = v23
		v166 = v26
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v31 != int32(-114) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	F_report_invalid_encoding(m, int32(4), v23, v24)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v23))))
	v67 = v57 & int32(255)
	v68 = int32(0)
	v70 = (v60 | v56<<(uint(int32(8))%32)) & int32(_a_F_euc_tw_to_big5_0) & int32(_a_F_euc_tw_to_big5_1)
	switch v67 - int32(149) {
	case 0:
		goto L37
	case 1:
		goto L36
	default:
		goto L38
	}
L20:
	;
	v56 = v31
	v57 = int32(149)
	v58 = int32(1)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	switch v47 - int32(161) {
	case 0:
		v53 = int32(149)
		goto L23
	case 1:
		goto L25
	default:
		goto L24
	}
L23:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
	v56 = v54
	v57 = v53
	v58 = int32(3)
	goto L19
L24:
	;
	v53 = v47 + int32(83)
	goto L23
L25:
	;
	v53 = int32(150)
	goto L23
L26:
	;
	if v126 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L27:
	;
	v126 = v124 & int32(_a_F_euc_tw_to_big5_0)
	goto L26
L28:
	;
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121))))
	v124 = v122
	goto L27
L29:
	;
	if v70 != int32(_a_F_euc_tw_to_big5_2) {
		v124 = v68
		goto L27
	} else {
		goto L54
	}
L30:
	;
	v121 = int32(_a_F_euc_tw_to_big5_3)
	goto L28
L31:
	;
	v121 = int32(_a_F_euc_tw_to_big5_4)
	goto L28
L32:
	;
	v115 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_euc_tw_to_big5[0])))
	v124 = v115
	goto L27
L33:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_euc_tw_to_big5[1])))
	v124 = v113
	goto L27
L34:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_euc_tw_to_big5[2])))
	v124 = v111
	goto L27
L35:
	;
	v109 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_euc_tw_to_big5[3])))
	v124 = v109
	goto L27
L36:
	;
	v107 = F_BinarySearchRange(m, int32(_a_F_euc_tw_to_big5_5), int32(47), v70)
	mBase = m.M
	v124 = v107
	goto L27
L37:
	;
	v104 = F_BinarySearchRange(m, int32(_a_F_euc_tw_to_big5_6), int32(24), v70)
	mBase = m.M
	v124 = v104
	goto L27
L38:
	;
	switch v67 - int32(246) {
	case 0:
		goto L39
	case 1:
		goto L40
	default:
		v124 = v68
		goto L27
	}
L39:
	;
	if base.Ui32(v70) <= base.Ui32(int32(_a_F_euc_tw_to_big5_7)) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	switch v70 - int32(_a_F_euc_tw_to_big5_8) {
	case 0:
		v121 = int32(_a_F_euc_tw_to_big5_9)
		goto L28
	case 1:
		goto L31
	case 2, 3, 4, 5, 6:
		v124 = v68
		goto L27
	case 7:
		goto L30
	default:
		goto L29
	}
L41:
	;
	if v70 == int32(_a_F_euc_tw_to_big5_10) {
		goto L33
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(v70) <= base.Ui32(int32(_a_F_euc_tw_to_big5_11)) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if v70 == int32(_a_F_euc_tw_to_big5_12) {
		goto L32
	} else {
		goto L45
	}
L45:
	;
	if v70 != int32(_a_F_euc_tw_to_big5_13) {
		v124 = v68
		goto L27
	} else {
		goto L46
	}
L46:
	;
	v87 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_euc_tw_to_big5[4])))
	v124 = v87
	goto L27
L47:
	;
	if v70 == int32(_a_F_euc_tw_to_big5_14) {
		goto L34
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v70 == int32(_a_F_euc_tw_to_big5_15) {
		goto L35
	} else {
		goto L52
	}
L50:
	;
	if v70 != int32(_a_F_euc_tw_to_big5_16) {
		v124 = v68
		goto L27
	} else {
		goto L51
	}
L51:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_euc_tw_to_big5[5])))
	v124 = v95
	goto L27
L52:
	;
	if v70 != int32(_a_F_euc_tw_to_big5_17) {
		v124 = v68
		goto L27
	} else {
		goto L53
	}
L53:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_euc_tw_to_big5[6])))
	v124 = v101
	goto L27
L54:
	;
	v121 = int32(_a_F_euc_tw_to_big5_18)
	goto L28
L55:
	;
	if v9 != 0 {
		v163 = v23
		v166 = v26
		goto L3
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v133 = int32(8)
	v137 = v126<<(uint(v133)%32) | int32(base.Ui32(v126)>>(uint(v133)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v137)
	v155 = v24 - v35
	v157 = v26 + int32(2)
	v160 = v23 + v35
	goto L9
L58:
	;
	F_report_untranslatable_char(m, int32(4), int32(36), v23, v24)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	if v9 != 0 {
		v163 = v23
		v166 = v26
		goto L3
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v31)
	v149 = int32(1)
	v155 = v24 - v149
	v157 = v26 + v149
	v160 = v23 + v149
	goto L9
L63:
	;
	F_report_invalid_encoding(m, int32(4), v23, v24)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
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
	goto L8
}
