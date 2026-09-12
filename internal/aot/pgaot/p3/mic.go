package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_mic_to_big5(m *base.Module, l0 int32) int32 {
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
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v14, v15, v16, int32(7), int32(36))
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
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v178)
	return v168 - v13
L4:
	;
	v168 = v13
	v170 = v12
	goto L3
L5:
	;
	goto L6
L6:
	;
	v25 = v13
	v26 = v16
	v27 = v12
	goto L7
L7:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v36 = base.I32_extend8_s(v35)
	if int32(0) <= v36 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v168 = v159
	v170 = v161
	goto L3
L9:
	;
	if int32(0) < v160 {
		v25 = v159
		v26 = v160
		v27 = v161
		goto L7
	} else {
		goto L63
	}
L10:
	;
	if v36 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v52 = F_pg_encoding_verifymbchar(m, int32(7), v25, v26)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v11 != 0 {
		v168 = v25
		v170 = v27
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v36)
	v45 = int32(1)
	v159 = v25 + v45
	v160 = v26 - v45
	v161 = v27 + v45
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(7), v25, v26)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
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
	if v52 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v11 != 0 {
		v168 = v25
		v170 = v27
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	switch v35 - int32(149) {
	case 0, 1:
		v66 = v36
		v67 = int32(2)
		v68 = int32(1)
		goto L25
	default:
		goto L24
	case 8:
		goto L26
	}
L22:
	;
	F_report_invalid_encoding(m, int32(7), v25, v26)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
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
	if v11 != 0 {
		v168 = v25
		v170 = v27
		goto L3
	} else {
		goto L61
	}
L25:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v67))))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v68))))
	v77 = v66 & int32(255)
	v78 = int32(0)
	v80 = (v70 | v72<<(uint(int32(8))%32)) & int32(32639)
	switch v77 - int32(149) {
	case 0:
		goto L38
	case 1:
		goto L37
	default:
		goto L39
	}
L26:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v66 = v63
	v67 = int32(3)
	v68 = int32(2)
	goto L25
L27:
	;
	if v138 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L28:
	;
	v138 = v136 & int32(65535)
	goto L27
L29:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133))))
	v136 = v134
	goto L28
L30:
	;
	if v80 != int32(8530) {
		v136 = v78
		goto L28
	} else {
		goto L55
	}
L31:
	;
	v133 = int32(2234104)
	goto L29
L32:
	;
	v133 = int32(2234100)
	goto L29
L33:
	;
	v127 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1299])))
	v136 = v127
	goto L28
L34:
	;
	v125 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1300])))
	v136 = v125
	goto L28
L35:
	;
	v123 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1301])))
	v136 = v123
	goto L28
L36:
	;
	v121 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1302])))
	v136 = v121
	goto L28
L37:
	;
	v119 = F_BinarySearchRange(m, int32(2234560), int32(47), v80)
	mBase = m.M
	v136 = v119
	goto L28
L38:
	;
	v116 = F_BinarySearchRange(m, int32(2234448), int32(24), v80)
	mBase = m.M
	v136 = v116
	goto L28
L39:
	;
	switch v77 - int32(246) {
	case 0:
		goto L40
	case 1:
		goto L41
	default:
		v136 = v78
		goto L28
	}
L40:
	;
	if v80 <= int32(17485) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	switch v80 - int32(8483) {
	case 0:
		v133 = int32(2234096)
		goto L29
	case 1:
		goto L32
	case 2, 3, 4, 5, 6:
		v136 = v78
		goto L28
	case 7:
		goto L31
	default:
		goto L30
	}
L42:
	;
	if v80 == int32(11357) {
		goto L34
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if v80 <= int32(20303) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	if v80 == int32(15742) {
		goto L33
	} else {
		goto L46
	}
L46:
	;
	if v80 != int32(17207) {
		v136 = v78
		goto L28
	} else {
		goto L47
	}
L47:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1303])))
	v136 = v99
	goto L28
L48:
	;
	if v80 == int32(17486) {
		goto L35
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v80 == int32(20304) {
		goto L36
	} else {
		goto L53
	}
L51:
	;
	if v80 != int32(19292) {
		v136 = v78
		goto L28
	} else {
		goto L52
	}
L52:
	;
	v107 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1304])))
	v136 = v107
	goto L28
L53:
	;
	if v80 != int32(20554) {
		v136 = v78
		goto L28
	} else {
		goto L54
	}
L54:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1305])))
	v136 = v113
	goto L28
L55:
	;
	v133 = int32(2234108)
	goto L29
L56:
	;
	if v11 != 0 {
		v168 = v25
		v170 = v27
		goto L3
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v145 = int32(8)
	v149 = v138<<(uint(v145)%32) | int32(base.Ui32(v138)>>(uint(v145)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v149)
	v159 = v25 + v52
	v160 = v26 - v52
	v161 = v27 + int32(2)
	goto L9
L59:
	;
	F_report_untranslatable_char(m, int32(7), int32(36), v25, v26)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_report_untranslatable_char(m, int32(7), int32(36), v25, v26)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
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
	goto L8
}
func F_mic_to_euc_tw(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v11, v12, v13, int32(7), int32(4))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v13 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v108 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v108)
	return v101 - v10
L4:
	;
	v101 = v10
	v102 = v9
	goto L3
L5:
	;
	goto L6
L6:
	;
	v22 = v10
	v23 = v9
	v24 = v13
	goto L7
L7:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v30 = base.I32_extend8_s(v29)
	if int32(0) <= v30 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v101 = v98
	v102 = v95
	goto L3
L9:
	;
	if int32(0) < v96 {
		v22 = v98
		v23 = v95
		v24 = v96
		goto L7
	} else {
		goto L32
	}
L10:
	;
	if v30 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v46 = F_pg_encoding_verifymbchar(m, int32(7), v22, v24)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v8 != 0 {
		v101 = v22
		v102 = v23
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v30)
	v39 = int32(1)
	v95 = v23 + v39
	v96 = v24 - v39
	v98 = v22 + v39
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(7), v22, v24)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
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
	if v46 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v8 != 0 {
		v101 = v22
		v102 = v23
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	switch v29 - int32(149) {
	case 0:
		goto L25
	case 1:
		goto L28
	default:
		goto L26
	case 8:
		goto L27
	}
L22:
	;
	F_report_invalid_encoding(m, int32(7), v22, v24)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
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
	v95 = v92
	v96 = v24 - v46
	v98 = v22 + v46
	goto L9
L25:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v86)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)) = uint8(v88)
	v92 = v23 + int32(2)
	goto L24
L26:
	;
	if v8 != 0 {
		v101 = v22
		v102 = v23
		goto L3
	} else {
		goto L30
	}
L27:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if base.Ui32(int32(4)) < base.Ui32((v63+int32(10))&int32(255)) {
		goto L26
	} else {
		goto L29
	}
L28:
	;
	v55 = int32(41614)
	*(*uint16)(unsafe.Add(mBase, uint32(v23))) = uint16(v55)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)) = uint8(v57)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+3)) = uint8(v59)
	v92 = v23 + int32(4)
	goto L24
L29:
	;
	v70 = int32(142)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v70)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	v74 = v72 - int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)) = uint8(v74)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)) = uint8(v76)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+3)) = uint8(v78)
	v92 = v23 + int32(4)
	goto L24
L30:
	;
	F_report_untranslatable_char(m, int32(7), int32(4), v22, v24)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	goto L8
}
func F_mic_to_latin1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(7), int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v21 = F_mic2latin(m, v6, v5, v10, int32(129), int32(8), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			return v21
		}
	}
}
