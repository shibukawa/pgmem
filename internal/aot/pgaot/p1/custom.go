package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DefineCustomEnumVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	v10 = F_init_custom_variable(m, l0, l1, l2, l5, int32(0), int32(4), int32(124))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+116)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = l4
		F_define_custom_variable(m, v10)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	if l3 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L31
	} else {
		goto L43
	}
L2:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L31
	} else {
		goto L39
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L31
	} else {
		goto L36
	}
L4:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	v84 = F_MemoryContextAllocExtended(m, v82, l6, int32(2))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L31
	} else {
		goto L32
	}
L5:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[538])))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if l4&int32(2) != 0 {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	if l4&int32(2) != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v80 = int32(1)
	goto L4
L10:
	;
	if l3 != int32(6) {
		v80 = l3
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v22 = int32(336342)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1217])))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v26 == int32(0) {
		v45 = v25
		v46 = v26
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v80 = int32(5)
	goto L4
L13:
	;
	if v46-v45 == int32(0) {
		goto L12
	} else {
		goto L21
	}
L14:
	;
	goto L13
L15:
	;
	if v25 != v26 {
		v45 = v25
		v46 = v26
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v30 = l0
	v31 = v22
	goto L17
L17:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v35 == int32(0) {
		v45 = v34
		v46 = v35
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v45 = v34
	v46 = v35
	goto L14
L19:
	;
	v38 = int32(1)
	if v34 == v35 {
		v30 = v30 + v38
		v31 = v31 + v38
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v50 = int32(145361)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1218])))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v54 == int32(0) {
		v73 = v53
		v74 = v54
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v74-v73 == int32(0) {
		goto L12
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	if v53 != v54 {
		v73 = v53
		v74 = v54
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v58 = l0
	v59 = v50
	goto L26
L26:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v63 == int32(0) {
		v73 = v62
		v74 = v63
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v73 = v62
	v74 = v63
	goto L23
L28:
	;
	v66 = int32(1)
	if v62 == v63 {
		v58 = v58 + v66
		v59 = v59 + v66
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v80 = int32(6)
	goto L4
L31:
	;
	return int32(0)
L32:
	;
	if v84 == int32(0) {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v92 = F__emscripten_memset_bulkmem(m, v84, base.I32_extend8_s(int32(0)), l6)
	mBase = m.M
	goto L34
L34:
	;
	v94 = F_guc_strdup(m, int32(22), l0)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = int32(46)
	*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v94
	return v92
L36:
	;
	F_errmsg_internal(m, int32(176996), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L31
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(523641), int32(4906), int32(414251))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L31
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
	F_errcode(m, int32(8389))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L31
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(14020), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(523641), int32(647), int32(510539))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_errmsg_internal(m, int32(243079), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L31
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(523641), int32(4897), int32(414251))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L31
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
