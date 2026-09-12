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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	goto L2
L1:
	;
	return
L2:
	;
	v4 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
	switch v5 {
	case 0:
		goto L14
	case 1, 4:
		goto L4
	case 2:
		goto L13
	case 3, 5:
		goto L12
	case 6:
		goto L11
	default:
		goto L1
	case 8:
		goto L10
	case 9:
		goto L9
	case 10:
		goto L8
	case 11, 13, 14, 17, 18:
		goto L6
	case 12:
		goto L7
	case 16, 19:
		goto L5
	}
L3:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L17
	} else {
		goto L33
	}
L4:
	;
	goto L3
L5:
	;
	F_CleanupSubTransaction(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L17
	} else {
		goto L32
	}
L6:
	;
	F_AbortSubTransaction(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L17
	} else {
		goto L31
	}
L7:
	;
	F_AbortSubTransaction(m)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L17
	} else {
		goto L30
	}
L8:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L17
	} else {
		goto L28
	}
L9:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L17
	} else {
		goto L26
	}
L10:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L17
	} else {
		goto L25
	}
L11:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L17
	} else {
		goto L23
	}
L12:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L17
	} else {
		goto L22
	}
L13:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L17
	} else {
		goto L20
	}
L14:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	switch v6 {
	case 0:
		goto L1
	case 1:
		goto L16
	default:
		goto L15
	}
L15:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = int32(2)
	goto L15
L17:
	;
	return
L18:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	return
L20:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(0)
	return
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(7)
	return
L23:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(0)
	return
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(0)
	return
L26:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(0)
	return
L28:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(0)
	return
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(15)
	return
L31:
	;
	goto L5
L32:
	;
	goto L2
L33:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L17
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(0)
	goto L1
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
	F___gettimeofday(m, v6)
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
	v2 = *(*int32)(unsafe.Add(mBase, _consts[75]))
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
	v3 = *(*int64)(unsafe.Add(mBase, _consts[171]))
	if v3 == int64(0) {
		v10 = m.G0
		v11 = int32(16)
		v12 = v10 - v11
		m.G0 = v12
		F___gettimeofday(m, v12)
		mBase = m.M
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
		v16 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+8)))
		m.G0 = v12 + v11
		v24 = v16 + v15*int64(1000000) - int64(946684800000000)
		*(*int64)(unsafe.Add(mBase, _consts[171])) = v24
		v26 = v24
	} else {
		v26 = v3
	}
	return v26
}
