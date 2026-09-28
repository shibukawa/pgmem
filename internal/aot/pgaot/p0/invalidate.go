package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InvalidateAttoptCacheCallback(m *base.Module, l0 int64, l1 int32, l2 int32) {
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
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateAttoptCacheCallback[0]))
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
	F_hash_seq_init(m, v7+int32(12), v10)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_hash_seq_init_with_hash_value(m, v7+int32(12), v10, l2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
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
	m.G0 = v7 + int32(32)
	return
L9:
	;
	v27 = F_hash_seq_search(m, v7+int32(12))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L19
	}
L11:
	;
	if v27 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_pfree(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateAttoptCacheCallback[0]))
	v38 = F_hash_search(m, v35, v27, int32(2), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	if v38 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	F_errmsg_internal(m, int32(_a_F_InvalidateAttoptCacheCallback_0), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_InvalidateAttoptCacheCallback_1), int32(77), int32(_a_F_InvalidateAttoptCacheCallback_2))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
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
func F_InvalidateEventCacheCallback(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateEventCacheCallback[0]))
	if v5 == int32(2) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateEventCacheCallback[1]))
		F_MemoryContextReset(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_InvalidateEventCacheCallback[2])) = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_InvalidateEventCacheCallback[0])) = int32(0)
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_InvalidateEventCacheCallback[0])) = int32(0)
		return
	}
}
func F_InvalidateTableSpaceCacheCallback(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateTableSpaceCacheCallback[0]))
	F_hash_seq_init(m, v7+int32(12), v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
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
	m.G0 = v7 + int32(32)
	return
L4:
	;
	v21 = F_hash_seq_search(m, v7+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	if v21 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
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
	goto L10
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateTableSpaceCacheCallback[0]))
	v32 = F_hash_search(m, v29, v21, int32(2), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	if v32 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L5
L14:
	;
	F_errmsg_internal(m, int32(_a_F_InvalidateTableSpaceCacheCallback_0), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_InvalidateTableSpaceCacheCallback_1), int32(70), int32(_a_F_InvalidateTableSpaceCacheCallback_2))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
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
