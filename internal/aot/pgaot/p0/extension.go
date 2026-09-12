package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InsertExtensionTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v19 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v21
	v29 = F_GetNewOidWithIndex(m, v19, int32(3080), int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v29
	v34 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), l1)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v34
	v40 = F_cstring_to_text(m, l5)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v40
	if l6 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l7 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v45 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+30)) = uint8(v45)
	goto L6
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = l6
	goto L6
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v58 = F_heap_form_tuple(m, v53, v13+int32(-32), v13+int32(-40))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	v50 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+31)) = uint8(v50)
	goto L10
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = l7
	goto L10
L14:
	;
	F_CatalogTupleInsert(m, v19, v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_pfree(m, v58)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_sequence_close(m, v19, int32(3))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_recordDependencyOnOwner(m, int32(3079), v29, l2)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v71 = F_new_object_addresses(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(3079)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(2615)
	F_add_exact_object_address(m, v13+int32(-52), v71)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
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
	F_record_object_address_dependencies(m, l0, v71, int32(110))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L28
	}
L22:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l8)+4))
	if v89 <= int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v102 = int32(0)
	goto L24
L24:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l8)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v102<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(3079)
	F_add_exact_object_address(m, v15, v71)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L21
L26:
	;
	v117 = v102 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l8)+4))
	if v117 < v118 {
		v102 = v117
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	F_free_object_addresses(m, v71)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _consts[203]))
	if v138 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v140 = int32(0)
	F_RunObjectPostCreateHook(m, int32(3079), v29, v140, v140)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	m.G0 = v15 - int32(-64)
	return
L33:
	;
	goto L32
}
