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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v2 = l1
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[0]))
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L16
	} else {
		goto L45
	}
L2:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+74)))
	if v12 != 0 {
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
	v162 = m.ExcPending
	if v162 != 0 {
		goto L16
	} else {
		goto L42
	}
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v13 | v14
	if v13&v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v21 = v13
	goto L9
L7:
	;
	goto L8
L8:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v87 | int32(-2147483648)
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v92)+75)) = uint8(v2)
	v94 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v92)+74)) = uint8(v94)
	if v2 == int32(2) {
		goto L32
	} else {
		goto L33
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(_a_F_LWLockQueueSelf_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_LWLockQueueSelf_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	if v21&int32(536870912) != 0 {
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
	v56 = int32(_a_F_LWLockQueueSelf_2)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[1]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v59 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
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
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v45&int32(536870912) != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v77 = int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v76 | v77
	if v76&v77 != 0 {
		v21 = v76
		goto L9
	} else {
		goto L30
	}
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[1])) = v74
	goto L20
L22:
	;
	if int32(999) < v57 {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v57 < int32(11) {
		goto L20
	} else {
		goto L29
	}
L25:
	;
	v64 = int32(900)
	if v64 <= v57 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v67 = v64
	goto L28
L27:
	;
	v67 = v57
	goto L28
L28:
	;
	v74 = v67 + int32(100)
	goto L21
L29:
	;
	v74 = v57 - int32(1)
	goto L21
L30:
	;
	goto L10
L31:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v152 & int32(-536870913)
	m.G0 = v8 + int32(32)
	return
L32:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[2]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[3]))
	v105 = v100 + v102*int32(640)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v106 == int32(-1) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[2]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[3]))
	v131 = v126 + v128*int32(640)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v132 == int32(-1) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v105)+76)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v102
	goto L31
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+76)) = v106
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[2]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	*(*int32)(unsafe.Add(mBase, uint32(v116+v106*int32(640))+80)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v105)+80)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v102
	goto L31
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v128
	goto L31
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v131)+76)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v128
	goto L38
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+80)) = v132
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[2]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	*(*int32)(unsafe.Add(mBase, uint32(v141+v132*int32(640))+76)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v131)+76)) = int32(-1)
	goto L38
L42:
	;
	F_errmsg_internal(m, int32(_a_F_LWLockQueueSelf_3), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L16
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_LWLockQueueSelf_1), int32(1056), int32(_a_F_LWLockQueueSelf_4))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errmsg_internal(m, int32(_a_F_LWLockQueueSelf_5), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_LWLockQueueSelf_1), int32(1059), int32(_a_F_LWLockQueueSelf_4))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L16
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
