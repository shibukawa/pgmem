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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v8 = v5 + int32(12)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_pg_advisory_unlock_all[0]))
	F_hash_seq_init(m, v8, v10)
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
	v15 = F_hash_seq_search(m, v8)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = v15
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
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
	if v19 == int32(2) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	F_ReleaseLockIfHeld(m, v17, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v27 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	if v27 != 0 {
		v17 = v27
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L8
}
func F_pg_advisory_unlock_int4(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13979(m, l0, int32(7))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pg_advisory_unlock_shared_int4(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13979(m, l0, int32(5))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pg_advisory_unlock_shared_int8(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13980(m, l0, int32(5))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pg_try_advisory_xact_lock_int4(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13988(m, l0, int32(0), int32(7))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_pg_try_advisory_xact_lock_int8(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13989(m, l0, int32(0), int32(7))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
