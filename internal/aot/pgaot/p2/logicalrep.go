package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_logicalrep_partmap_invalidate_cb(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[650]))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v6 + int32(32)
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
	F_hash_seq_init(m, v6+int32(12), v9)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_hash_seq_init(m, v6+int32(12), v9)
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
	v21 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+40)) = uint8(v27)
	F_hash_seq_term(m, v6+int32(12))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L13
	}
L10:
	;
	if v21 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	if v25 != l1 {
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
	v39 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	if v39 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v43 = v39
	goto L17
L17:
	;
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+40)) = uint8(v46)
	v50 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L19
	}
L18:
	;
	goto L1
L19:
	;
	if v50 != 0 {
		v43 = v50
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
	var v27 int32
	_ = v27
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
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
	v10 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v14 = F_LWLockAcquire(m, v10+int32(5504), int32(1))
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
	v19 = *(*int32)(unsafe.Add(mBase, _consts[643]))
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
	v23 = *(*int32)(unsafe.Add(mBase, _consts[644]))
	v27 = v4
	v29 = v4
	v30 = v19
	v31 = v23
	goto L9
L7:
	;
	v61 = v4
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
	v34 = v31 + v29*int32(112)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+32)))
	if v35 != int32(1) {
		v51 = v27
		v53 = v30
		v54 = v31
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v61 = v51
	goto L8
L11:
	;
	v56 = v29 + int32(1)
	if v56 < v53 {
		v27 = v51
		v29 = v56
		v30 = v53
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
		v51 = v27
		v53 = v30
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
		v51 = v27
		v53 = v30
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
	v45 = F_lappend(m, v27, v39)
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
	v48 = *(*int32)(unsafe.Add(mBase, _consts[643]))
	v50 = *(*int32)(unsafe.Add(mBase, _consts[644]))
	v51 = v45
	v53 = v48
	v54 = v50
	goto L11
L19:
	;
	goto L10
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v67+int32(5504))
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
	return v61
L23:
	;
	goto L22
}
