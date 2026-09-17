package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DefineCustomEnumVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	v11 = F_init_custom_variable(m, l0, l1, l2, l6, int32(0), int32(4), int32(124))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+116)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = l3
		v16 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = l5
		F_define_custom_variable(m, v11)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			return
		}
	}
}
func F_init_custom_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	if l3 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L29
	} else {
		goto L42
	}
L2:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L29
	} else {
		goto L39
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L29
	} else {
		goto L36
	}
L4:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_custom_variable[0])))
	if v12&int32(1) == int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if l4&int32(2) != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	goto L6
L8:
	;
	if l3 != int32(6) {
		v78 = l3
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_init_custom_variable[1]))
	v82 = F_MemoryContextAllocExtended(m, v80, l6, int32(2))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L29
	} else {
		goto L30
	}
L10:
	;
	v21 = int32(_a_F_init_custom_variable_0)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_custom_variable[2])))
	if base.B2i32(v24 == int32(0))|base.B2i32(v24 != v27) != 0 {
		v45 = v24
		v46 = v27
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v45-v46 != 0 {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	v30 = l0
	v31 = v21
	goto L14
L14:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v35 == int32(0) {
		v45 = v35
		v46 = v34
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v45 = v35
	v46 = v34
	goto L12
L16:
	;
	v38 = int32(1)
	if v35 == v34 {
		v30 = v30 + v38
		v31 = v31 + v38
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v49 = int32(_a_F_init_custom_variable_1)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_custom_variable[3])))
	if base.B2i32(v52 == int32(0))|base.B2i32(v52 != v55) != 0 {
		v73 = v52
		v74 = v55
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v78 = int32(5)
	goto L9
L21:
	;
	if v73-v74 != 0 {
		v78 = int32(6)
		goto L9
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v58 = l0
	v59 = v49
	goto L24
L24:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v63 == int32(0) {
		v73 = v63
		v74 = v62
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v73 = v63
	v74 = v62
	goto L22
L26:
	;
	v66 = int32(1)
	if v63 == v62 {
		v58 = v58 + v66
		v59 = v59 + v66
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	goto L20
L29:
	;
	return int32(0)
L30:
	;
	if v82 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if l6 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	base.MemoryFill(m, v82, int32(0), l6)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v91 = F_guc_strdup(m, int32(22), l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+24)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v82)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = int32(46)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v91
	return v82
L36:
	;
	F_errmsg_internal(m, int32(_a_F_init_custom_variable_2), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L29
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_init_custom_variable_3), int32(_a_F_init_custom_variable_4), int32(_a_F_init_custom_variable_5))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L29
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
	F_errmsg_internal(m, int32(_a_F_init_custom_variable_6), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L29
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_init_custom_variable_3), int32(_a_F_init_custom_variable_7), int32(_a_F_init_custom_variable_5))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L29
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_errcode(m, int32(_a_F_init_custom_variable_8))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_init_custom_variable_9), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L29
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_init_custom_variable_3), int32(647), int32(_a_F_init_custom_variable_10))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L29
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
