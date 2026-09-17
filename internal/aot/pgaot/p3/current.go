package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AbortCurrentTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	goto L3
L1:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L14
	} else {
		goto L24
	}
L2:
	;
	return
L3:
	;
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_AbortCurrentTransaction[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
	switch v5 {
	case 0:
		goto L11
	case 1, 4:
		goto L5
	case 2, 6, 9, 10:
		goto L1
	case 3, 5:
		goto L10
	default:
		goto L2
	case 8:
		goto L9
	case 11, 13, 14, 17, 18:
		goto L7
	case 12:
		goto L8
	case 16, 19:
		goto L6
	}
L4:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L14
	} else {
		goto L22
	}
L5:
	;
	goto L4
L6:
	;
	F_CleanupSubTransaction(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L14
	} else {
		goto L21
	}
L7:
	;
	F_AbortSubTransaction(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L14
	} else {
		goto L20
	}
L8:
	;
	F_AbortSubTransaction(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L14
	} else {
		goto L19
	}
L9:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L14
	} else {
		goto L18
	}
L10:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L14
	} else {
		goto L17
	}
L11:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	switch v6 {
	case 0:
		goto L2
	case 1:
		goto L13
	default:
		goto L12
	}
L12:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = int32(2)
	goto L12
L14:
	;
	return
L15:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	return
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(7)
	return
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(0)
	return
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(15)
	return
L20:
	;
	goto L6
L21:
	;
	goto L3
L22:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(0)
	goto L2
L24:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(0)
	return
}
func F_GetCurrentTimestamp(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	v4 = m.G0
	v5 = int32(16)
	v6 = v4 - v5
	m.G0 = v6
	F_gettimeofday(m, v6)
	mBase = m.M
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v10 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+8)))
	m.G0 = v6 + v5
	return v10 + v9*int64(1000000) - int64(946684800000000)
}
func F_GetCurrentTransactionIdIfAny(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTransactionIdIfAny[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	return v3
}
func F_GetCurrentTransactionStopTimestamp(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	v3 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTransactionStopTimestamp[0]))
	if v3 == int64(0) {
		v10 = m.G0
		v11 = int32(16)
		v12 = v10 - v11
		m.G0 = v12
		F_gettimeofday(m, v12)
		mBase = m.M
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
		v16 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+8)))
		m.G0 = v12 + v11
		v24 = v16 + v15*int64(1000000) - int64(946684800000000)
		*(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTransactionStopTimestamp[0])) = v24
		v26 = v24
	} else {
		v26 = v3
	}
	return v26
}
