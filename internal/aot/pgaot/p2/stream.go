package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_read_stream_begin_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	v5 = l4
	if l2 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L22
	} else {
		goto L80
	}
L2:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
	if l1 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L3:
	;
	if l0&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[1]))
	v65 = v47
	goto L2
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	goto L17
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+118)))
	if v18 == int32(116) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[0]))
	if v35 == int32(0) {
		goto L4
	} else {
		goto L16
	}
L9:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+24)))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[0]))
	if v25 == int32(0) {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	goto L14
L14:
	;
	if base.B2i32(base.Ui32(v29) < base.Ui32(int32(_a_F_read_stream_begin_impl_6))) == int32(0) {
		v39 = v28
		goto L5
	} else {
		goto L15
	}
L15:
	;
	goto L4
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v39 = v38
	goto L5
L17:
	;
	if base.B2i32(base.Ui32(v40) < base.Ui32(int32(_a_F_read_stream_begin_impl_6))) == int32(0) {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	goto L4
L19:
	;
	v50 = F_get_tablespace_maintenance_io_concurrency(m, v39)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v54 = F_get_tablespace(m, v39)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L22
	} else {
		goto L25
	}
L22:
	;
	return int32(0)
L23:
	;
	v65 = v50
	goto L2
L24:
	;
	v65 = v63
	goto L2
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v56 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+24))
	if int32(0) <= v57 {
		v63 = v57
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[1]))
	v63 = v62
	goto L24
L29:
	;
	goto L28
L30:
	;
	v81 = int32(_a_F_read_stream_begin_impl_3)
	if v81 <= v65 {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[3]))
	v80 = v71
	goto L30
L32:
	;
	goto L33
L33:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v74 = base.I32_div_s(v72, int32(2))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v75 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v78 = v74
	goto L36
L35:
	;
	v78 = v72
	goto L36
L36:
	;
	v80 = v78
	goto L30
L37:
	;
	v84 = v81
	goto L39
L38:
	;
	v84 = v65
	goto L39
L39:
	;
	v87 = v67 * (v84 + int32(1))
	v89 = int32(16)
	v94 = (v67<<(uint(v89)%32) - int32(_a_F_read_stream_begin_impl_4)) >> (uint(v89) % 32)
	v95 = int32(_a_F_read_stream_begin_impl_5) - v94
	if base.Ui32(v87) < base.Ui32(v95) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v97 = v87
	goto L42
L41:
	;
	v97 = v95
	goto L42
L42:
	;
	if base.Ui32(v80) < base.Ui32(v97) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v99 = v80
	goto L45
L44:
	;
	v99 = v97
	goto L45
L45:
	;
	v100 = int32(1)
	if v84 <= v100 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v103 = v100
	goto L48
L47:
	;
	v103 = v84
	goto L48
L48:
	;
	v105 = v103 * int32(84)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v107 != int32(-1) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if base.Ui32(v99) < base.Ui32(v116) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[4]))
	v113 = base.I32_div_s(v111, int32(4))
	v116 = v113
	goto L49
L51:
	;
	goto L52
L52:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[5]))
	v116 = v115
	goto L49
L53:
	;
	v118 = v99
	goto L55
L54:
	;
	v118 = v116
	goto L55
L55:
	;
	if base.Ui32(v118) <= base.Ui32(int32(1)) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v121 = int32(1)
	goto L58
L57:
	;
	v121 = v118
	goto L58
L58:
	;
	v123 = v121 + int32(1)
	v124 = base.I32_extend16_s(v123)
	v129 = (v124 + v94) << (uint(int32(2)) % 32)
	v133 = F_palloc(m, v105+l8*v124+v129+int32(108))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L22
	} else {
		goto L59
	}
L59:
	;
	base.MemoryFill(m, v133, int32(0), int32(92))
	v142 = (v133 + v129 + int32(99)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v133)+76)) = v142
	if l8 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+72)) = (v105 + v142 + int32(7)) & int32(-8)
	goto L62
L61:
	;
	goto L62
L62:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[6]))
	v155 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+33)) = uint8(v155)
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+32)) = uint8(base.B2i32(v151 == int32(0)))
	if v151 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v133))) = uint16(v175)
	v177 = int32(0)
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+68)) = l8
	*(*uint16)(unsafe.Add(mBase, uint32(v133)+2)) = uint16(v179)
	*(*uint16)(unsafe.Add(mBase, uint32(v133)+8)) = uint16(v121)
	*(*int32)(unsafe.Add(mBase, uint32(v133)+48)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v133)+44)) = l6
	*(*uint16)(unsafe.Add(mBase, uint32(v133)+6)) = uint16(v123)
	*(*int64)(unsafe.Add(mBase, uint32(v133)+52)) = int64(-1)
	v188 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v133)+40)) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	*(*uint16)(unsafe.Add(mBase, uint32(v133)+18)) = uint16(v177)
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+35)) = uint8(base.B2i32(v190 != v188))
	if base.Ui32(v121) < base.Ui32(base.I32_extend16_s(v179)) {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	if v65 != 0 {
		v175 = v84
		goto L63
	} else {
		goto L67
	}
L65:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[7])))
	if v161&int32(1)|l0&int32(2)|base.B2i32(v65 <= int32(0)) != 0 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v170 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+34)) = uint8(v170)
	v175 = v84
	goto L63
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+28)) = int32(8)
	v175 = int32(1)
	goto L63
L68:
	;
	v198 = v121
	goto L70
L69:
	;
	v198 = v179
	goto L70
L70:
	;
	if l0&int32(4) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v202 = v198
	goto L73
L72:
	;
	v202 = int32(1)
	goto L73
L73:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v133)+16)) = uint16(v202)
	*(*uint16)(unsafe.Add(mBase, uint32(v133)+24)) = uint16(v202)
	*(*uint16)(unsafe.Add(mBase, uint32(v133)+22)) = uint16(v202)
	*(*uint16)(unsafe.Add(mBase, uint32(v133)+14)) = uint16(v202)
	if int32(0) < v175 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v220 = v177
	goto L77
L75:
	;
	goto L76
L76:
	;
	return v133
L77:
	;
	v226 = v220 * int32(84)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v133)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v226+v227)+4)) = l2
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v133)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v230+v226)+8)) = l3
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v133)+76))
	*(*uint8)(unsafe.Add(mBase, uint32(v233+v226)+12)) = uint8(v5)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v133)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v236+v226)+16)) = l5
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v133)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v239+v226)+20)) = l1
	v243 = v220 + int32(1)
	if v243 != v175 {
		v220 = v243
		goto L77
	} else {
		goto L79
	}
L78:
	;
	goto L76
L79:
	;
	goto L78
L80:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L22
	} else {
		goto L81
	}
L81:
	;
	F_errmsg(m, int32(_a_F_read_stream_begin_impl_0), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L22
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_read_stream_begin_impl_1), int32(787), int32(_a_F_read_stream_begin_impl_2))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L22
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
