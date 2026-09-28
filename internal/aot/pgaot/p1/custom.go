package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DefineCustomEnumVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	v10 = F_init_custom_variable(m, l0, l1, l2, l6, int32(0), int32(4))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+120)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = l3
		v15 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+116)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = l5
		F_define_custom_variable(m, v10)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			return
		}
	}
}
func F_init_custom_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
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
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	if l3 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L29
	} else {
		goto L39
	}
L2:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L29
	} else {
		goto L36
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L29
	} else {
		goto L33
	}
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_custom_variable[0])))
	if v11&int32(1) == int32(0) {
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
		v77 = l3
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_init_custom_variable[1]))
	v82 = F_MemoryContextAllocExtended(m, v79, int32(152), int32(2))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L29
	} else {
		goto L30
	}
L10:
	;
	v20 = int32(_a_F_init_custom_variable_0)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_custom_variable[2])))
	if base.B2i32(v23 == int32(0))|base.B2i32(v23 != v26) != 0 {
		v44 = v23
		v45 = v26
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v44-v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	v29 = l0
	v30 = v20
	goto L14
L14:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	if v34 == int32(0) {
		v44 = v34
		v45 = v33
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v44 = v34
	v45 = v33
	goto L12
L16:
	;
	v37 = int32(1)
	if v34 == v33 {
		v29 = v29 + v37
		v30 = v30 + v37
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v48 = int32(_a_F_init_custom_variable_1)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_custom_variable[3])))
	if base.B2i32(v51 == int32(0))|base.B2i32(v51 != v54) != 0 {
		v72 = v51
		v73 = v54
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v77 = int32(5)
	goto L9
L21:
	;
	if v72-v73 != 0 {
		v77 = int32(6)
		goto L9
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v57 = l0
	v58 = v48
	goto L24
L24:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	if v62 == int32(0) {
		v72 = v62
		v73 = v61
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v72 = v62
	v73 = v61
	goto L22
L26:
	;
	v65 = int32(1)
	if v62 == v61 {
		v57 = v57 + v65
		v58 = v58 + v65
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
	base.MemoryFill(m, v82, int32(0), int32(152))
	v92 = F_guc_strdup(m, int32(22), l0)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+24)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v82)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = int32(47)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v92
	return v82
L33:
	;
	F_errmsg_internal(m, int32(_a_F_init_custom_variable_2), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_init_custom_variable_3), int32(_a_F_init_custom_variable_4), int32(_a_F_init_custom_variable_5))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errmsg_internal(m, int32(_a_F_init_custom_variable_6), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_init_custom_variable_3), int32(_a_F_init_custom_variable_7), int32(_a_F_init_custom_variable_5))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
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
	F_errcode(m, int32(_a_F_init_custom_variable_8))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L29
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(_a_F_init_custom_variable_9), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L29
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_init_custom_variable_3), int32(646), int32(_a_F_init_custom_variable_10))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
