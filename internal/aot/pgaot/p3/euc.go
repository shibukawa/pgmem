package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_euc_jis_2004_to_utf8(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14231(m, l0, int32(5), int32(0), int32(25), int32(_a_F_euc_jis_2004_to_utf8_0), int32(_a_F_euc_jis_2004_to_utf8_1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_euc_tw_to_big5(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
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
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
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
	return int64(0)
L2:
	;
	if v14 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v177 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v172))) = uint8(v177)
	return base.I64_extend_i32_s(v169 - v11)
L4:
	;
	v169 = v11
	v172 = v10
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
	v169 = v166
	v172 = v163
	goto L3
L9:
	;
	if int32(0) < v161 {
		v23 = v166
		v24 = v161
		v26 = v163
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
	if v9 != int64(0) {
		v169 = v23
		v172 = v26
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
	v43 = m.ExcPending
	if v43 != 0 {
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
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v23))))
	v69 = v59 & int32(255)
	v70 = int32(0)
	v72 = (v62 | v58<<(uint(int32(8))%32)) & int32(_a_F_euc_tw_to_big5_0) & int32(_a_F_euc_tw_to_big5_1)
	switch v69 - int32(149) {
	case 0:
		goto L37
	case 1:
		goto L36
	default:
		goto L38
	}
L20:
	;
	v58 = v31
	v59 = int32(149)
	v60 = int32(1)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	switch v49 - int32(161) {
	case 0:
		v55 = int32(149)
		goto L23
	case 1:
		goto L25
	default:
		goto L24
	}
L23:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
	v58 = v56
	v59 = v55
	v60 = int32(3)
	goto L19
L24:
	;
	v55 = v49 + int32(83)
	goto L23
L25:
	;
	v55 = int32(150)
	goto L23
L26:
	;
	if v128 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L27:
	;
	v128 = v126 & int32(_a_F_euc_tw_to_big5_0)
	goto L26
L28:
	;
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123))))
	v126 = v124
	goto L27
L29:
	;
	if v72 != int32(_a_F_euc_tw_to_big5_2) {
		v126 = v70
		goto L27
	} else {
		goto L54
	}
L30:
	;
	v123 = int32(_a_F_euc_tw_to_big5_3)
	goto L28
L31:
	;
	v123 = int32(_a_F_euc_tw_to_big5_4)
	goto L28
L32:
	;
	v117 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_euc_tw_to_big5[0])))
	v126 = v117
	goto L27
L33:
	;
	v115 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_euc_tw_to_big5[1])))
	v126 = v115
	goto L27
L34:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_euc_tw_to_big5[2])))
	v126 = v113
	goto L27
L35:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_euc_tw_to_big5[3])))
	v126 = v111
	goto L27
L36:
	;
	v109 = F_BinarySearchRange(m, int32(_a_F_euc_tw_to_big5_5), int32(47), v72)
	mBase = m.M
	v126 = v109
	goto L27
L37:
	;
	v106 = F_BinarySearchRange(m, int32(_a_F_euc_tw_to_big5_6), int32(24), v72)
	mBase = m.M
	v126 = v106
	goto L27
L38:
	;
	switch v69 - int32(246) {
	case 0:
		goto L39
	case 1:
		goto L40
	default:
		v126 = v70
		goto L27
	}
L39:
	;
	if base.Ui32(v72) <= base.Ui32(int32(_a_F_euc_tw_to_big5_7)) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	switch v72 - int32(_a_F_euc_tw_to_big5_8) {
	case 0:
		v123 = int32(_a_F_euc_tw_to_big5_9)
		goto L28
	case 1:
		goto L31
	case 2, 3, 4, 5, 6:
		v126 = v70
		goto L27
	case 7:
		goto L30
	default:
		goto L29
	}
L41:
	;
	if v72 == int32(_a_F_euc_tw_to_big5_10) {
		goto L33
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(v72) <= base.Ui32(int32(_a_F_euc_tw_to_big5_11)) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if v72 == int32(_a_F_euc_tw_to_big5_12) {
		goto L32
	} else {
		goto L45
	}
L45:
	;
	if v72 != int32(_a_F_euc_tw_to_big5_13) {
		v126 = v70
		goto L27
	} else {
		goto L46
	}
L46:
	;
	v89 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_euc_tw_to_big5[4])))
	v126 = v89
	goto L27
L47:
	;
	if v72 == int32(_a_F_euc_tw_to_big5_14) {
		goto L34
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v72 == int32(_a_F_euc_tw_to_big5_15) {
		goto L35
	} else {
		goto L52
	}
L50:
	;
	if v72 != int32(_a_F_euc_tw_to_big5_16) {
		v126 = v70
		goto L27
	} else {
		goto L51
	}
L51:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_euc_tw_to_big5[5])))
	v126 = v97
	goto L27
L52:
	;
	if v72 != int32(_a_F_euc_tw_to_big5_17) {
		v126 = v70
		goto L27
	} else {
		goto L53
	}
L53:
	;
	v103 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_euc_tw_to_big5[6])))
	v126 = v103
	goto L27
L54:
	;
	v123 = int32(_a_F_euc_tw_to_big5_18)
	goto L28
L55:
	;
	if v9 != int64(0) {
		v169 = v23
		v172 = v26
		goto L3
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v137 = int32(8)
	v141 = v128<<(uint(v137)%32) | int32(base.Ui32(v128)>>(uint(v137)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v141)
	v161 = v24 - v35
	v163 = v26 + int32(2)
	v166 = v23 + v35
	goto L9
L58:
	;
	F_report_untranslatable_char(m, int32(4), int32(36), v23, v24)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
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
	if v9 != int64(0) {
		v169 = v23
		v172 = v26
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
	v155 = int32(1)
	v161 = v24 - v155
	v163 = v26 + v155
	v166 = v23 + v155
	goto L9
L63:
	;
	F_report_invalid_encoding(m, int32(4), v23, v24)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
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
