package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_create_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v24 int64
	_ = v24
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if base.Ui32(int32(4095)) < base.Ui32(l1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	base.MemoryFill(m, l0+int32(20), int32(0), int32(1472))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l3 ^ int32(216163848)
	v24 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1452)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1444)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1440)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1448)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(l0)+216)) = v24
	v53 = int32(base.Ui32(l1)>>(uint(int32(10))%32)) & int32(_a_F_create_internal_0)
	v55 = v53 + int32(2048)
	v57 = v55 & int32(4092)
	if v57 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L7
	} else {
		goto L26
	}
L4:
	;
	v61 = v53 - v57 + int32(_a_F_create_internal_1)
	goto L6
L5:
	;
	v61 = v55
	goto L6
L6:
	;
	v64 = int32(base.Ui32(l1-v61) >> (uint(int32(12)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1472)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1460)) = int32(1)
	v70 = F_palloc(m, int32(656))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = l0
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_create_internal[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v76
	base.MemoryFill(m, v70+int32(8), int32(0), int32(648))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1472))
	F_LWLockInitialize(m, l0+int32(1476), v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v98 = int32(0)
	goto L10
L10:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1472))
	F_LWLockInitialize(m, v99+v98<<(uint(int32(5))%32)+int32(224), v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = l0 + int32(2048)
	v119 = l0 + int32(1496)
	*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v119
	v121 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v119)+4)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v119)+12)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v119)+20)) = v121
	v127 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v119)+28)) = v127
	v129 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v119)+32)) = uint8(v129)
	if v119 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v109 = v98 + int32(1)
	if v109 != int32(38) {
		v98 = v109
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	if v64 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v135 = v119 - l0 + v129
	goto L17
L16:
	;
	v135 = v127
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v135
	base.MemoryFill(m, l0+int32(1532), int32(0), int32(516))
	goto L14
L18:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+20)) = v167
	m.G0 = v12 + int32(16)
	return v70
L19:
	;
	v144 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(160)))) = v144
	v167 = v144
	goto L18
L20:
	;
	goto L21
L21:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	F_FreePageManagerPut(m, v147, int32(base.Ui32(v61)>>(uint(int32(12))%32)), v64)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v152 = int32(14)
	v155 = base.I32_clz(v64) ^ int32(31)
	if base.Ui32(v152) <= base.Ui32(v155) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v158 = v152
	goto L25
L24:
	;
	v158 = v155
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+v158<<(uint(int32(2))%32))+164)) = int32(0)
	v167 = v158 + int32(1)
	goto L18
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_create_internal_2)
	F_errmsg_internal(m, int32(_a_F_create_internal_3), v12)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_create_internal_4), int32(1290), int32(_a_F_create_internal_5))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
