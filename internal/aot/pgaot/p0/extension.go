package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InsertExtensionTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int64, l7 int64, l8 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	v14 = m.G0
	v16 = v14 - int32(96)
	m.G0 = v16
	v20 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+88)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v22
	v30 = F_GetNewOidWithIndex(m, v20, int32(3080), int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = base.I64_extend_i32_u(v30)
	v37 = F_DirectFunctionCall1Coll(m, int32(534), int32(0), base.I64_extend_i32_u(l1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = base.I64_extend_i32_u(l4)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = base.I64_extend_i32_u(l3)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = base.I64_extend_i32_u(l2)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v37
	v46 = F_cstring_to_text(m, l5)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+72)) = base.I64_extend_i32_u(v46)
	if l6 == int64(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l7 == int64(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v52 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+30)) = uint8(v52)
	goto L6
L8:
	;
	goto L9
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = l6
	goto L6
L10:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v65 = F_heap_form_tuple(m, v60, v16+int32(32), v16+int32(24))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	v57 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+31)) = uint8(v57)
	goto L10
L12:
	;
	goto L13
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+88)) = l7
	goto L10
L14:
	;
	F_CatalogTupleInsert(m, v20, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_pfree(m, v65)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_relation_close(m, v20, int32(3))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_recordDependencyOnOwner(m, int32(3079), v30, l2)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v78 = F_new_object_addresses(m)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(3079)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(2615)
	F_add_exact_object_address(m, v16+int32(12), v78)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if l8 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_record_object_address_dependencies(m, l0, v78, int32(110))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L28
	}
L22:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l8)+4))
	if v96 <= int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v101 = int32(0)
	goto L24
L24:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l8)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v101<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(3079)
	F_add_exact_object_address(m, v16, v78)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L21
L26:
	;
	v125 = v101 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l8)+4))
	if v125 < v126 {
		v101 = v125
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	F_free_object_addresses(m, v78)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_InsertExtensionTuple[0]))
	if v147 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v149 = int32(0)
	F_RunObjectPostCreateHook(m, int32(3079), v30, v149, v149)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	m.G0 = v16 + int32(96)
	return
L33:
	;
	goto L32
}
