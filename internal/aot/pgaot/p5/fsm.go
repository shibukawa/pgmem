package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fsm_get_avail(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+l1)+uint32(_c_F_fsm_get_avail[0]))))
	return v6
}
func F_fsm_search_avail(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	if l0 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return v259
L2:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+28)))
	if base.Ui32(v34) < base.Ui32(l1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_search_avail[0]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19+(l0^int32(-1))<<(uint(int32(2))%32))))
	v33 = v25
	goto L2
L4:
	;
	goto L5
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_search_avail[1]))
	v33 = v27 + l0<<(uint(int32(13))%32) + int32(-8192)
	goto L2
L6:
	;
	v259 = int32(-1)
	goto L1
L7:
	;
	goto L8
L8:
	;
	v38 = v33 + int32(28)
	v44 = l3
	goto L9
L9:
	;
	v52 = int32(4095)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v55 = v53 + v52
	if base.Ui32(int32(4068)) < base.Ui32(v53) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v259 = int32(-1)
	goto L1
L11:
	;
	if v93 <= int32(4094) {
		goto L26
	} else {
		goto L27
	}
L12:
	;
	v58 = v52
	goto L14
L13:
	;
	v58 = v55
	goto L14
L14:
	;
	if v58 <= int32(0) {
		v93 = v55
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v66 = v58
	goto L16
L16:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v38))))
	if base.Ui32(l1) <= base.Ui32(v73) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v93 = v86
	goto L11
L18:
	;
	v93 = v66
	goto L11
L19:
	;
	goto L20
L20:
	;
	v75 = int32(1)
	if (v66+int32(2))&(v66+v75) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v84 = v66
	goto L23
L22:
	;
	v84 = int32(base.Ui32(v66)>>(uint(v75)%32)) - v75
	goto L23
L23:
	;
	v86 = base.I32_div_s(v84, int32(2))
	if int32(1) < v84 {
		v66 = v86
		goto L16
	} else {
		goto L24
	}
L24:
	;
	goto L17
L25:
	;
	F_MarkBufferDirtyHint(m, l0, int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L45
	} else {
		goto L70
	}
L26:
	;
	v106 = v93
	goto L29
L27:
	;
	v235 = v93
	goto L28
L28:
	;
	v245 = (v235 + int32(_a_F_fsm_search_avail_0)) & int32(_a_F_fsm_search_avail_1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v245 + l2
	v259 = v245
	goto L1
L29:
	;
	v114 = v106 << (uint(int32(1)) % 32)
	if base.Ui32(v114) <= base.Ui32(int32(_a_F_fsm_search_avail_2)) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v235 = v228
	goto L28
L31:
	;
	if v228 < int32(4095) {
		v106 = v228
		goto L29
	} else {
		goto L69
	}
L32:
	;
	v118 = v114 | int32(1)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v118))))
	if base.Ui32(l1) <= base.Ui32(v120) {
		v228 = v118
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v124 = v114 + int32(2)
	if base.Ui32(v124) <= base.Ui32(int32(_a_F_fsm_search_avail_2)) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v38))))
	if base.Ui32(l1) <= base.Ui32(v128) {
		v228 = v124
		goto L31
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v131 = v14 + int32(36)
	if l0 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L38
L40:
	;
	v164 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v153)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = v155
	*(*int64)(unsafe.Add(mBase, uint32(v131))) = v154
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(32)))) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(28)))) = v160
	goto L40
L42:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_search_avail[2]))
	v153 = v140 + (l0^int32(-1))<<(uint(int32(6))%32)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_search_avail[3]))
	v153 = v147 + l0<<(uint(int32(6))%32) + int32(-64)
	goto L41
L45:
	;
	return int32(0)
L46:
	;
	if v164 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v168
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v170
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v172
	F_errmsg_internal(m, int32(_a_F_fsm_search_avail_3), v14)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v44&int32(1) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	F_errfinish(m, int32(_a_F_fsm_search_avail_4), int32(277), int32(_a_F_fsm_search_avail_5))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	F_LockBuffer(m, l0, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L45
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v197 = int32(4094)
	goto L57
L55:
	;
	F_LockBuffer(m, l0, int32(2))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L45
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v197) {
		v219 = int32(0)
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L25
L59:
	;
	v220 = v197 + v38
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	if v221 != v219&int32(255) {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v208 = v197 << (uint(int32(1)) % 32)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+int32(29)+v208))))
	if v197 == int32(4081) {
		v219 = v210
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208+v38)+2)))
	if base.Ui32(v214) < base.Ui32(v210) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v216 = v210
	goto L64
L63:
	;
	v216 = v214
	goto L64
L64:
	;
	v219 = v216
	goto L59
L65:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v220))) = uint8(v219)
	goto L67
L66:
	;
	goto L67
L67:
	;
	if v197 != 0 {
		v197 = v197 - int32(1)
		goto L57
	} else {
		goto L68
	}
L68:
	;
	goto L58
L69:
	;
	goto L30
L70:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(l1) <= base.Ui32(v252) {
		v44 = int32(1)
		goto L9
	} else {
		goto L71
	}
L71:
	;
	goto L10
}
