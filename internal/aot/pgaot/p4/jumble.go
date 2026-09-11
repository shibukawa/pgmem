package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__jumbleCreateEventTrigStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v70 != 0 {
		goto L25
	} else {
		goto L26
	}
L2:
	;
	if v4&int32(3) == int32(0) {
		v28 = v4
		goto L7
	} else {
		goto L8
	}
L3:
	;
	goto L4
L4:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v66 + int32(1)
	goto L1
L5:
	;
	F_AppendJumble(m, l0, v4, v61+int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L22
	} else {
		goto L23
	}
L6:
	;
	v61 = v53 - v4
	goto L5
L7:
	;
	v32 = v28
	goto L16
L8:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v12 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v61 = int32(0)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v17 = v4
	goto L12
L12:
	;
	v21 = v17 + int32(1)
	if v21&int32(3) == int32(0) {
		v28 = v21
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v53 = v21
	goto L6
L14:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v26 != 0 {
		v17 = v21
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v41 = int32(-2139062144)
	if (int32(16843008)-v38|v38)&v41 == v41 {
		v32 = v32 + int32(4)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v47 = v32
	goto L19
L18:
	;
	goto L17
L19:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v51 != 0 {
		v47 = v47 + int32(1)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v53 = v47
	goto L6
L21:
	;
	goto L20
L22:
	;
	return
L23:
	;
	goto L1
L24:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F__jumbleNode(m, l0, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L22
	} else {
		goto L46
	}
L25:
	;
	if v70&int32(3) == int32(0) {
		v94 = v70
		goto L30
	} else {
		goto L31
	}
L26:
	;
	goto L27
L27:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v132 + int32(1)
	goto L24
L28:
	;
	F_AppendJumble(m, l0, v70, v127+int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L22
	} else {
		goto L45
	}
L29:
	;
	v127 = v119 - v70
	goto L28
L30:
	;
	v98 = v94
	goto L39
L31:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v78 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v127 = int32(0)
	goto L28
L33:
	;
	goto L34
L34:
	;
	v83 = v70
	goto L35
L35:
	;
	v87 = v83 + int32(1)
	if v87&int32(3) == int32(0) {
		v94 = v87
		goto L30
	} else {
		goto L37
	}
L36:
	;
	v119 = v87
	goto L29
L37:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v92 != 0 {
		v83 = v87
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v107 = int32(-2139062144)
	if (int32(16843008)-v104|v104)&v107 == v107 {
		v98 = v98 + int32(4)
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v113 = v98
	goto L42
L41:
	;
	goto L40
L42:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v117 != 0 {
		v113 = v113 + int32(1)
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v119 = v113
	goto L29
L44:
	;
	goto L43
L45:
	;
	goto L24
L46:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F__jumbleNode(m, l0, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L22
	} else {
		goto L47
	}
L47:
	;
	return
}
