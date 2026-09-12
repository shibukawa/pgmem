package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InvalidateAttoptCacheCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[866]))
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	goto L9
L2:
	;
	F_hash_seq_init(m, v6+int32(12), v9)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_hash_seq_init_with_hash_value(m, v6+int32(12), v9, l2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return
L6:
	;
	goto L1
L7:
	;
	goto L1
L8:
	;
	m.G0 = v6 + int32(32)
	return
L9:
	;
	v25 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L19
	}
L11:
	;
	if v25 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v29 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_pfree(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[866]))
	v36 = F_hash_search(m, v33, v25, int32(2), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	if v36 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	F_errmsg_internal(m, int32(445153), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(499312), int32(76), int32(319406))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_InvalidateEventCacheCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v5 = *(*int32)(unsafe.Add(mBase, _consts[867]))
	if v5 == int32(2) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[869]))
		F_MemoryContextReset(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[868])) = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[867])) = int32(0)
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[867])) = int32(0)
		return
	}
}
func F_InvalidateTableSpaceCacheCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v11 = *(*int32)(unsafe.Add(mBase, _consts[885]))
	F_hash_seq_init(m, v6+int32(12), v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L4
L3:
	;
	m.G0 = v6 + int32(32)
	return
L4:
	;
	v19 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	if v19 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_pfree(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[885]))
	v30 = F_hash_search(m, v27, v19, int32(2), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	if v30 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L5
L14:
	;
	F_errmsg_internal(m, int32(445153), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(499383), int32(69), int32(319436))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
