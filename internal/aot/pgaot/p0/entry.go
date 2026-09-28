package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_entryLocateEntry(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v187 int32
	_ = v187
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v225 int32
	_ = v225
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v19 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v38 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_entryLocateEntry[0]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23+(v19^int32(-1))<<(uint(int32(2))%32))))
	v37 = v29
	goto L1
L3:
	;
	goto L4
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_entryLocateEntry[1]))
	v37 = v31 + v19<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	m.G0 = v17 + int32(16)
	return v225
L6:
	;
	v41 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v41)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v44) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v62) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v54 = int32(base.Ui32(v44+int32(_a_F_entryLocateEntry_0))>>(uint(int32(2))%32)) & int32(_a_F_entryLocateEntry_1)
	goto L11
L10:
	;
	v54 = int32(0)
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v43 * v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v58 = m.T0[v57].(func(*base.Module, int32, int32) int32)(m, l0, v37)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v225 = v58
	goto L5
L14:
	;
	v70 = int32(base.Ui32(v62+int32(_a_F_entryLocateEntry_0)) >> (uint(int32(2)) % 32))
	goto L16
L15:
	;
	v70 = int32(0)
	goto L16
L16:
	;
	v72 = v70 + int32(1)
	if base.Ui32(int32(2)) <= base.Ui32(v72&int32(_a_F_entryLocateEntry_1)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v85 = v72
	v87 = int32(1)
	goto L20
L18:
	;
	v187 = v72
	goto L19
L19:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v187)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v37+v187&int32(_a_F_entryLocateEntry_1)<<(uint(int32(2))%32))+20))
	v205 = v37 + v202&int32(_a_F_entryLocateEntry_2)
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205))))
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205)+2)))
	v225 = v206<<(uint(int32(16))%32) | v209
	goto L5
L20:
	;
	v99 = int32(base.Ui32((v85-v87)&int32(_a_F_entryLocateEntry_3))>>(uint(int32(1))%32)) + v87
	v100 = int32(_a_F_entryLocateEntry_1)
	v101 = v99 & v100
	if v101 == v70&v100 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v187 = v173
	goto L19
L22:
	;
	v172 = base.B2i32(int32(0) < v166)
	if int32(0) < v166 {
		goto L44
	} else {
		goto L45
	}
L23:
	;
	v105 = int32(-1)
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+16)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v37+v106)))
	if v108 == v105 {
		v166 = v105
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(20)+v101<<(uint(int32(2))%32))))
	v119 = v37 + v116&int32(_a_F_entryLocateEntry_2)
	v120 = F_gintuple_get_attrnum(m, v112, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L12
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v125 = F_gintuple_get_key(m, v122, v119, v17+int32(15))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
	if v127 != v120 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v160 != 0 {
		goto L41
	} else {
		goto L42
	}
L30:
	;
	v160 = base.B2i32(base.Ui32(v127) < base.Ui32(v120))
	goto L29
L31:
	;
	goto L32
L32:
	;
	v130 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+64)))
	v131 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+15)))
	if v130 != v131 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v160 = base.B2i32(v130 < v131)
	goto L29
L34:
	;
	goto L35
L35:
	;
	if v130 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v120<<(uint(int32(2))%32)+v136)+uint32(_c_F_entryLocateEntry[2])))
	v148 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v149 = F_FunctionCall2Coll(m, v136+v120*int32(28)+int32(112), v147, v148, v125)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L12
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v99)
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+2)))
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119))))
	v225 = v155 | v156<<(uint(int32(16))%32)
	goto L5
L39:
	;
	v151 = base.I32_wrap_i64(v149)
	if v151 != 0 {
		v166 = v151
		goto L22
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v165 = int32(-1)
	goto L43
L42:
	;
	v165 = int32(1)
	goto L43
L43:
	;
	v166 = v165
	goto L22
L44:
	;
	v173 = v85
	goto L46
L45:
	;
	v173 = v99
	goto L46
L46:
	;
	if int32(0) < v166 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v178 = v99 + int32(1)
	goto L49
L48:
	;
	v178 = v87
	goto L49
L49:
	;
	if base.Ui32(v178&int32(_a_F_entryLocateEntry_1)) < base.Ui32(v173&int32(_a_F_entryLocateEntry_1)) {
		v85 = v173
		v87 = v178
		goto L20
	} else {
		goto L50
	}
L50:
	;
	goto L21
}
