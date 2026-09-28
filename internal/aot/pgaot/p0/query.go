package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitializeQueryCompletion(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return
}
func F_ScanQueryForLocks(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	v2 = l1
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v2)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v14 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v67 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v17 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = v3
	goto L4
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v23<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	switch v33 {
	case 0:
		goto L8
	case 1:
		goto L7
	default:
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	v56 = v23 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v56 < v57 {
		v23 = v56
		goto L4
	} else {
		goto L23
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v40 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v2 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_LockRelationOid(m, v35, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	F_UnlockRelationOid(m, v35, v34)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L14
	}
L12:
	;
	return
L13:
	;
	goto L6
L14:
	;
	goto L6
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
	F_ScanQueryForLocks(m, v49, v2)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L22
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v2 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_LockRelationOid(m, v40, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_UnlockRelationOid(m, v40, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L21
	}
L20:
	;
	goto L15
L21:
	;
	goto L15
L22:
	;
	goto L6
L23:
	;
	goto L5
L24:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	if v102 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	v70 = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v71 <= v70 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v77 = v70
	goto L27
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v77<<(uint(int32(2))%32))))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	F_ScanQueryForLocks(m, v87, v2)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L12
	} else {
		goto L29
	}
L28:
	;
	goto L24
L29:
	;
	v91 = v77 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v91 < v92 {
		v77 = v91
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v109 = F_query_tree_walker_impl(m, l0, int32(1802), v11+int32(15), int32(3))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L12
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	m.G0 = v11 + int32(16)
	return
L34:
	;
	goto L33
}
func F_assign_query_collations_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == int32(142) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v12 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 <= int32(0) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	v53 = F_assign_collations_walker(m, l0, v8+int32(8))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L13
	}
L7:
	;
	v24 = int32(0)
	goto L8
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v24<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	v38 = F_assign_collations_walker(m, v30, v8+int32(8))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L1
L10:
	;
	return int32(0)
L11:
	;
	v43 = v24 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v43 < v44 {
		v24 = v43
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	goto L1
}
func F_extract_query_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(528)
	m.G0 = v9
	v11 = int32(400)
	v12 = v9 + v11
	base.MemoryFill(m, v12, v5, int32(128))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+468)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v9)+400)) = int32(268)
	base.MemoryFill(m, v9, v5, v11)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(269)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v12
	v26 = F_extract_query_dependencies_walker(m, l0, v9)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+464))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v28
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v9)+468))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+493)))
		*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v32)
		m.G0 = v9 + int32(528)
		return
	}
}
func F_query_or_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	if l0 == int32(0) {
		v15 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 != int32(67) {
			v15 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		} else {
			v10 = F_query_tree_walker_impl(m, l0, l1, l2, l3)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
