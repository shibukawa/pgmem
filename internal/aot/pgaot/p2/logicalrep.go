package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_logicalrep_partmap_invalidate_cb(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_partmap_invalidate_cb[0]))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(32)
	return
L2:
	;
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_hash_seq_init(m, v7+int32(12), v10)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v34 = v7 + int32(12)
	F_hash_seq_init(m, v34, v10)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L14
	}
L6:
	;
	return
L7:
	;
	goto L8
L8:
	;
	v22 = v7 + int32(12)
	v23 = F_hash_seq_search(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v29 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+40)) = uint8(v29)
	F_hash_seq_term(m, v22)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L13
	}
L10:
	;
	if v23 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	if v27 != l1 {
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
L14:
	;
	v37 = F_hash_seq_search(m, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	if v37 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v41 = v37
	goto L17
L17:
	;
	v45 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+40)) = uint8(v45)
	v49 = F_hash_seq_search(m, v7+int32(12))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L19
	}
L18:
	;
	goto L1
L19:
	;
	if v49 != 0 {
		v41 = v49
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
}
func F_logicalrep_relmap_free_entry(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_pfree(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_pfree(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v9 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v13 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_bms_free(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14+v13<<(uint(int32(2))%32))))
	F_pfree(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_pfree(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v22 = v13 + int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v22 < v23 {
		v13 = v22
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pfree(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v36 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_free_attrmap(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	return
L17:
	;
	goto L16
}
func F_logicalrep_workers_find(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	v4 = int32(0)
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_workers_find[0]))
	v14 = F_LWLockAcquire(m, v10+int32(_a_F_logicalrep_workers_find_0), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_workers_find[1]))
	if int32(0) < v19 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_workers_find[2]))
	v28 = v19
	v29 = v23
	v30 = v4
	v31 = v4
	goto L9
L7:
	;
	v65 = v4
	goto L8
L8:
	;
	if l2 != 0 {
		goto L20
	} else {
		goto L21
	}
L9:
	;
	v34 = v29 + v30*int32(112)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+32)))
	if v35 != int32(1) {
		v52 = v28
		v53 = v29
		v54 = v31
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v65 = v54
	goto L8
L11:
	;
	v56 = v30 + int32(1)
	if v56 < v52 {
		v28 = v52
		v29 = v53
		v30 = v56
		v31 = v54
		goto L9
	} else {
		goto L19
	}
L12:
	;
	v39 = v34 + int32(16)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	if v40 != l0 {
		v52 = v28
		v53 = v29
		v54 = v31
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if l1 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	if v42 == int32(0) {
		v52 = v28
		v53 = v29
		v54 = v31
		goto L11
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v45 = F_lappend(m, v31, v39)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_workers_find[1]))
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_workers_find[2]))
	v52 = v48
	v53 = v50
	v54 = v45
	goto L11
L19:
	;
	goto L10
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_workers_find[0]))
	F_LWLockRelease(m, v67+int32(_a_F_logicalrep_workers_find_0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	return v65
L23:
	;
	goto L22
}
