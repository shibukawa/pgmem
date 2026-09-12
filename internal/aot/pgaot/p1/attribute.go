package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckAttributeNamesTypes(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v14) < base.Ui32(int32(1601)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L34
	} else {
		goto L70
	}
L2:
	;
	switch l1 - int32(99) {
	case 0, 19:
		goto L6
	default:
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L34
	} else {
		goto L66
	}
L5:
	;
	m.G0 = v12 + int32(48)
	return
L6:
	;
	if int32(2) <= v14 {
		goto L40
	} else {
		goto L41
	}
L7:
	;
	if v14 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v28 = int32(0)
	goto L9
L9:
	;
	v38 = l0 + v14<<(uint(int32(4))%32) + int32(24) + v28*int32(100)
	v40 = F_strcmp(m, int32(765276), v38)
	mBase = m.M
	if v40 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L34
	} else {
		goto L35
	}
L11:
	;
	if v69 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L12:
	;
	v69 = int32(765272)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v45 = F_strcmp(m, int32(765376), v38)
	mBase = m.M
	if v45 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v69 = int32(765372)
	goto L11
L16:
	;
	goto L17
L17:
	;
	v50 = F_strcmp(m, int32(765476), v38)
	mBase = m.M
	if v50 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v69 = int32(765472)
	goto L11
L19:
	;
	goto L20
L20:
	;
	v55 = F_strcmp(m, int32(765576), v38)
	mBase = m.M
	if v55 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v69 = int32(765572)
	goto L11
L22:
	;
	goto L23
L23:
	;
	v60 = F_strcmp(m, int32(765676), v38)
	mBase = m.M
	if v60 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v69 = int32(765672)
	goto L11
L25:
	;
	goto L26
L26:
	;
	v67 = F_strcmp(m, int32(765776), v38)
	mBase = m.M
	if v67 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v68 = int32(0)
	goto L29
L28:
	;
	v68 = int32(765772)
	goto L29
L29:
	;
	v69 = v68
	goto L11
L30:
	;
	v73 = v28 + int32(1)
	if v14 != v73 {
		v28 = v73
		goto L9
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L10
L33:
	;
	goto L6
L34:
	;
	return
L35:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v38
	F_errmsg(m, int32(381418), v12+int32(32))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(495987), int32(482), int32(162599))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v179 = int32(1)
	if v14 <= v179 {
		goto L59
	} else {
		goto L60
	}
L40:
	;
	v108 = l0 + v14<<(uint(int32(4))%32) + int32(24)
	v113 = int32(1)
	goto L43
L41:
	;
	goto L42
L42:
	;
	if v14 == int32(0) {
		goto L5
	} else {
		goto L58
	}
L43:
	;
	v121 = v108 + v113*int32(100)
	v124 = int32(0)
	goto L45
L44:
	;
	goto L39
L45:
	;
	v133 = v124 * int32(100)
	v134 = v108 + v133
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v138 == int32(0) {
		v157 = v137
		v158 = v138
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v166 = v113 + int32(1)
	if v166 != v14 {
		v113 = v166
		goto L43
	} else {
		goto L57
	}
L47:
	;
	if v158-v157 == int32(0) {
		goto L1
	} else {
		goto L55
	}
L48:
	;
	goto L47
L49:
	;
	if v137 != v138 {
		v157 = v137
		v158 = v138
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v142 = v134
	v143 = v121
	goto L51
L51:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+1)))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+1)))
	if v147 == int32(0) {
		v157 = v146
		v158 = v147
		goto L48
	} else {
		goto L53
	}
L52:
	;
	v157 = v146
	v158 = v147
	goto L48
L53:
	;
	v150 = int32(1)
	if v146 == v147 {
		v142 = v142 + v150
		v143 = v143 + v150
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v163 = v124 + int32(1)
	if v163 != v113 {
		v124 = v163
		goto L45
	} else {
		goto L56
	}
L56:
	;
	goto L46
L57:
	;
	goto L44
L58:
	;
	goto L39
L59:
	;
	v182 = v179
	goto L61
L60:
	;
	v182 = v14
	goto L61
L61:
	;
	v189 = int32(0)
	goto L62
L62:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v196 = int32(4)
	v201 = l0 + int32(20) + v195<<(uint(v196)%32) + v189*int32(100)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v201)+68))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)+96))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+90)))
	F_CheckAttributeType(m, v201+v196, v204, v205, int32(0), base.B2i32(v207 == int32(118))<<(uint(int32(3))%32)|l2)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L34
	} else {
		goto L64
	}
L63:
	;
	goto L5
L64:
	;
	v216 = v189 + int32(1)
	if v216 != v182 {
		v189 = v216
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L34
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(1600)
	F_errmsg(m, int32(148688), v12)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L34
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(495987), int32(464), int32(162599))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L34
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L34
	} else {
		goto L71
	}
L71:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0 + v254<<(uint(int32(4))%32) + v133 + int32(24)
	F_errmsg(m, int32(415506), v12+int32(16))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L34
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(495987), int32(498), int32(162599))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L34
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
