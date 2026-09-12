package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_advisory_unlock_all(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v10 = *(*int32)(unsafe.Add(mBase, _consts[1185]))
	F_hash_seq_init(m, v5+int32(12), v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = v17
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v5 + int32(32)
	return int32(0)
L7:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+15)))
	if v21 == int32(2) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	F_ReleaseLockIfHeld(m, v19, int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v29 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	if v29 != 0 {
		v19 = v29
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L8
}
func F_pg_advisory_unlock_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(34209794)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v9
	v15 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v15
	v19 = F_LockRelease(m, v6, int32(7), int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v19
	}
}
func F_pg_advisory_unlock_shared_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(34209794)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v9
	v15 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v15
	v19 = F_LockRelease(m, v6, int32(5), int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v19
	}
}
func F_pg_advisory_unlock_shared_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(34209793)
	*(*uint32)(unsafe.Add(mBase, uint32(v6)+8)) = uint32(v9)
	v14 = int64(base.Ui64(v9) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6)+4)) = uint32(v14)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v17
	v21 = F_LockRelease(m, v6, int32(5), int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v21
	}
}
func F_pg_try_advisory_xact_lock_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(34209794)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v9
	v15 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v15
	v20 = F_LockAcquire(m, v6, int32(7), int32(0), int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return base.B2i32(v20 != int32(0))
	}
}
func F_pg_try_advisory_xact_lock_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(34209793)
	*(*uint32)(unsafe.Add(mBase, uint32(v6)+8)) = uint32(v9)
	v14 = int64(base.Ui64(v9) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6)+4)) = uint32(v14)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v17
	v22 = F_LockAcquire(m, v6, int32(7), int32(0), int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return base.B2i32(v22 != int32(0))
	}
}
