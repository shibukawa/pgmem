package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_standard_ExecutorRun(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int64
	_ = v116
	var v121 int64
	_ = v121
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int64
	_ = v166
	var v169 int32
	_ = v169
	v14 = int32(_a_F_standard_ExecutorRun_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorRun[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorRun[0])) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_InstrStartNode(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = int64(0)
	if v24 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	if l1 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)))
	if v31 != int32(1) {
		v39 = int32(0)
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	m.T0[v35].(func(*base.Module, int32, int32, int32))(m, v23, v24, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	v39 = int32(1)
	goto L6
L12:
	;
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v17)+120))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v17)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v158 + v159
	if v39 != 0 {
		goto L56
	} else {
		goto L57
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = l1
	if l2 == int64(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v89 = int64(0)
	goto L22
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+29)))
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v59)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+160)) = uint8(v58)
	if v58 != v59 {
		v74 = int32(0)
		goto L14
	} else {
		goto L20
	}
L16:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v47&int32(1) == int32(0) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v52 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v52)
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+160)) = uint8(v54)
	v74 = v54
	goto L14
L19:
	;
	goto L18
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorRun[1]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+72)) = v68 + int32(1)
	goto L21
L21:
	;
	v74 = int32(1)
	goto L14
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v43)+152))
	if v90 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+128)))
	if v128&int32(8) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	F_MemoryContextReset(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
	if v94 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	F_ExecReScan(m, v42)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v98 = m.T0[v97].(func(*base.Module, int32) int32)(m, v42)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	goto L23
L33:
	;
	if v98 == int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+4)))
	if v102&int32(2) != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v43)+60))
	if v105 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v106 = F_ExecFilterJunk(m, v105, v98)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	v108 = v98
	goto L38
L38:
	;
	if v39 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v108 = v106
	goto L38
L40:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v110 = m.T0[v109].(func(*base.Module, int32, int32) int32)(m, v108, v23)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if base.B2i32(v24 != int32(1)) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	if v110 == int32(0) {
		goto L32
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v43)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v43)+112)) = v116 + int64(1)
	goto L47
L46:
	;
	goto L47
L47:
	;
	v121 = v89 + int64(1)
	if l2 == int64(0) {
		v89 = v121
		goto L22
	} else {
		goto L48
	}
L48:
	;
	if l2 != v121 {
		v89 = v121
		goto L22
	} else {
		goto L49
	}
L49:
	;
	goto L32
L50:
	;
	v134 = F_ExecShutdownNode_walker(m, v42, int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v74 == int32(0) {
		goto L12
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorRun[1]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v140)+72)) = v141 - int32(1)
	goto L55
L55:
	;
	goto L12
L56:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	m.T0[v162].(func(*base.Module, int32))(m, v23)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v165 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v17)+112))
	F_InstrStopNode(m, v165, base.F64_convert_i64_u(v166))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorRun[0])) = v15
	return
L63:
	;
	goto L62
}
