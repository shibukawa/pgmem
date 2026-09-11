package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_euc_jis_2004_to_utf8(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(5), int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(0)
		v27 = F_LocalToUtf(m, v6, v10, v5, int32(4314404), int32(2291456), int32(25), v17, int32(5), base.B2i32(v7 != v17))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
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
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
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
	v173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v173)
	return v165 - v11
L4:
	;
	v165 = v11
	v168 = v10
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
	v165 = v162
	v168 = v159
	goto L3
L9:
	;
	if int32(0) < v157 {
		v23 = v162
		v24 = v157
		v26 = v159
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
		v165 = v23
		v168 = v26
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
	v70 = (v60 | v56<<(uint(int32(8))%32)) & int32(65535) & int32(32639)
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
	if v128 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L27:
	;
	v128 = v126 & int32(65535)
	goto L26
L28:
	;
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123))))
	v126 = v124
	goto L27
L29:
	;
	if v70 != int32(8530) {
		v126 = v68
		goto L27
	} else {
		goto L54
	}
L30:
	;
	v123 = int32(2171528)
	goto L28
L31:
	;
	v123 = int32(2171524)
	goto L28
L32:
	;
	v117 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1293])))
	v126 = v117
	goto L27
L33:
	;
	v115 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1294])))
	v126 = v115
	goto L27
L34:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1295])))
	v126 = v113
	goto L27
L35:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1296])))
	v126 = v111
	goto L27
L36:
	;
	v109 = F_BinarySearchRange(m, int32(2171984), int32(47), v70)
	mBase = m.M
	v126 = v109
	goto L27
L37:
	;
	v106 = F_BinarySearchRange(m, int32(2171872), int32(24), v70)
	mBase = m.M
	v126 = v106
	goto L27
L38:
	;
	switch v67 - int32(246) {
	case 0:
		goto L39
	case 1:
		goto L40
	default:
		v126 = v68
		goto L27
	}
L39:
	;
	if v70 <= int32(17485) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	switch v70 - int32(8483) {
	case 0:
		v123 = int32(2171520)
		goto L28
	case 1:
		goto L31
	case 2, 3, 4, 5, 6:
		v126 = v68
		goto L27
	case 7:
		goto L30
	default:
		goto L29
	}
L41:
	;
	if v70 == int32(11357) {
		goto L33
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v70 <= int32(20303) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if v70 == int32(15742) {
		goto L32
	} else {
		goto L45
	}
L45:
	;
	if v70 != int32(17207) {
		v126 = v68
		goto L27
	} else {
		goto L46
	}
L46:
	;
	v89 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1297])))
	v126 = v89
	goto L27
L47:
	;
	if v70 == int32(17486) {
		goto L34
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v70 == int32(20304) {
		goto L35
	} else {
		goto L52
	}
L50:
	;
	if v70 != int32(19292) {
		v126 = v68
		goto L27
	} else {
		goto L51
	}
L51:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1298])))
	v126 = v97
	goto L27
L52:
	;
	if v70 != int32(20554) {
		v126 = v68
		goto L27
	} else {
		goto L53
	}
L53:
	;
	v103 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1299])))
	v126 = v103
	goto L27
L54:
	;
	v123 = int32(2171532)
	goto L28
L55:
	;
	if v9 != 0 {
		v165 = v23
		v168 = v26
		goto L3
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v135 = int32(8)
	v139 = v128<<(uint(v135)%32) | int32(base.Ui32(v128)>>(uint(v135)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v139)
	v157 = v24 - v35
	v159 = v26 + int32(2)
	v162 = v23 + v35
	goto L9
L58:
	;
	F_report_untranslatable_char(m, int32(4), int32(36), v23, v24)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
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
		v165 = v23
		v168 = v26
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
	v151 = int32(1)
	v157 = v24 - v151
	v159 = v26 + v151
	v162 = v23 + v151
	goto L9
L63:
	;
	F_report_invalid_encoding(m, int32(4), v23, v24)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
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
