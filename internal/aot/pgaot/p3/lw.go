package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LWLockQueueSelf(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	v2 = l1
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L16
	} else {
		goto L46
	}
L2:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+74)))
	if v14 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L16
	} else {
		goto L43
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15 | v16
	if v15&v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = v15
	goto L9
L7:
	;
	goto L8
L8:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v97 | int32(-2147483648)
	v102 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*uint8)(unsafe.Add(mBase, uint32(v102)+75)) = uint8(v2)
	v104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v102)+74)) = uint8(v104)
	if v2 == int32(2) {
		goto L32
	} else {
		goto L33
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(317356)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(497473)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(0)
	if v23&int32(536870912) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	goto L14
L12:
	;
	goto L13
L13:
	;
	v64 = int32(4121980)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[333]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(8))+8))
	if v67 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	F_perform_spin_delay(m, v10+int32(8))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	return
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v51&int32(536870912) != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v85 = int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v84 | v85
	if v84&v85 != 0 {
		v23 = v84
		goto L9
	} else {
		goto L30
	}
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[333])) = v82
	goto L20
L22:
	;
	if int32(999) < v65 {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v65 < int32(11) {
		goto L20
	} else {
		goto L29
	}
L25:
	;
	v72 = int32(900)
	if v72 <= v65 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v75 = v72
	goto L28
L27:
	;
	v75 = v65
	goto L28
L28:
	;
	v82 = v75 + int32(100)
	goto L21
L29:
	;
	v82 = v65 - int32(1)
	goto L21
L30:
	;
	goto L10
L31:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v170 & int32(-536870913)
	m.G0 = v10 + int32(32)
	return
L32:
	;
	v109 = l0 + int32(8)
	v111 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v114 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v119 = v112 + v114*int32(640) + int32(76)
	v120 = int32(-1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v121 == v120 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v142 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v147 = v140 + v142*int32(640) + int32(76)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v148 == int32(-1) {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+4)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v114
	goto L31
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v119))) = int64(-1)
	v134 = v109
	v135 = v114
	goto L35
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v121
	v128 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	*(*int32)(unsafe.Add(mBase, uint32(v129+v121*int32(640))+80)) = v114
	v134 = v119
	v135 = v120
	goto L35
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v142
	goto L31
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v147))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v142
	goto L39
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v148
	v156 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	*(*int32)(unsafe.Add(mBase, uint32(v157+v148*int32(640))+76)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = int32(-1)
	goto L39
L43:
	;
	F_errmsg_internal(m, int32(362624), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(497473), int32(1056), int32(338738))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L16
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errmsg_internal(m, int32(372779), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L16
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(497473), int32(1059), int32(338738))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L16
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
