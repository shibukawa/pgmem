package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ShutdownSQLFunction(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	v6 = base.I32_wrap_i64(l0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = v7
	goto L4
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v75 != 0 {
		goto L27
	} else {
		goto L28
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v13 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v67 != 0 {
		v9 = v67
		goto L4
	} else {
		goto L26
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+57)))
	if v17 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	F_PushActiveSnapshot(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v24 = int32(_a_F_ShutdownSQLFunction_0)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownSQLFunction[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_ShutdownSQLFunction[0])) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(2)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 != int32(6) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	return
L12:
	;
	goto L10
L13:
	;
	F_ExecutorFinish(m, v31)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	v41 = v31
	goto L15
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	m.T0[v43].(func(*base.Module, int32))(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L18
	}
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	F_ExecutorEnd(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v41 = v40
	goto L15
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	F_FreeQueryDesc(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ShutdownSQLFunction[0])) = v25
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+9)))
	if v53 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	F_MemoryContextDelete(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L11
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+57)))
	if v62 != 0 {
		goto L6
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	goto L6
L26:
	;
	goto L5
L27:
	;
	F_tuplestore_end(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L11
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
	if v80 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
	F_ReleaseCachedPlan(m, v80, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L11
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v84 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+6)) = uint8(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v84
	return
L34:
	;
	goto L33
}
