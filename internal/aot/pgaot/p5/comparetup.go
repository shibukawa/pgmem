package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_comparetup_heap_tiebreak(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v15 - v17
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v16 + v17
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v23 - v17
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v24 + v17
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v32 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(48)
	return v151
L2:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if int32(2) <= v80 {
		goto L27
	} else {
		goto L28
	}
L3:
	;
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+10)))
	v40 = F_heap_getattr_1(m, v12+int32(28), v37, v31, v12+int32(7))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v48 = F_heap_getattr_1(m, v12+int32(8), v37, v31, v12+int32(6))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+6)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)))
	if v51 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v50&int32(1) != 0 {
		goto L2
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v50&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+9)))
	if v58 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v59 = int32(-1)
	goto L13
L12:
	;
	v59 = int32(1)
	goto L13
L13:
	;
	v151 = v59
	goto L1
L14:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+9)))
	if v64 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v67 = m.T0[v66].(func(*base.Module, int32, int32, int32) int32)(m, v40, v48, v14)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L20
	}
L17:
	;
	v65 = int32(1)
	goto L19
L18:
	;
	v65 = int32(-1)
	goto L19
L19:
	;
	v151 = v65
	goto L1
L20:
	;
	v69 = int32(1)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)))
	if v70 != v69 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v77 = v67
	goto L23
L22:
	;
	if v67 < int32(0) {
		v151 = v69
		goto L1
	} else {
		goto L24
	}
L23:
	;
	if v77 != 0 {
		v151 = v77
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v77 = int32(0) - v67
	goto L23
L25:
	;
	goto L2
L26:
	;
	v151 = int32(1)
	goto L1
L27:
	;
	v88 = v14
	v90 = int32(1)
	goto L30
L28:
	;
	goto L29
L29:
	;
	v151 = int32(0)
	goto L1
L30:
	;
	v94 = v88 + int32(36)
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+46)))
	v100 = F_heap_getattr_1(m, v12+int32(28), v97, v31, v12+int32(7))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	v106 = F_heap_getattr_1(m, v12+int32(8), v97, v31, v12+int32(6))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+6)))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)))
	if v109 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v137 = v90 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v137 < v138 {
		v88 = v94
		v90 = v137
		goto L30
	} else {
		goto L54
	}
L35:
	;
	if v108&int32(1) != 0 {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v108&int32(1) != 0 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+45)))
	if v116 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v117 = int32(-1)
	goto L41
L40:
	;
	v117 = int32(1)
	goto L41
L41:
	;
	v151 = v117
	goto L1
L42:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+45)))
	if v122 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v88)+52))
	v125 = m.T0[v124].(func(*base.Module, int32, int32, int32) int32)(m, v100, v106, v94)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L48
	}
L45:
	;
	v123 = int32(1)
	goto L47
L46:
	;
	v123 = int32(-1)
	goto L47
L47:
	;
	v151 = v123
	goto L1
L48:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+44)))
	if v127 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v125 < int32(0) {
		goto L26
	} else {
		goto L52
	}
L50:
	;
	v134 = v125
	goto L51
L51:
	;
	if v134 != 0 {
		v151 = v134
		goto L1
	} else {
		goto L53
	}
L52:
	;
	v134 = int32(0) - v125
	goto L51
L53:
	;
	goto L34
L54:
	;
	goto L31
}
