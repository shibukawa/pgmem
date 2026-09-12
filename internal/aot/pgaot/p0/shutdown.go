package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ShutdownSQLFunction(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v71 != 0 {
		goto L27
	} else {
		goto L28
	}
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v10 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v64 != 0 {
		v7 = v64
		goto L4
	} else {
		goto L26
	}
L7:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+57)))
	if v14 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_PushActiveSnapshot(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v21 = int32(4470400)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(2)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v29 != int32(6) {
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
	F_ExecutorFinish(m, v28)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	v38 = v28
	goto L15
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	m.T0[v40].(func(*base.Module, int32))(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L11
	} else {
		goto L18
	}
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	F_ExecutorEnd(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v38 = v37
	goto L15
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	F_FreeQueryDesc(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v50 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	F_MemoryContextDelete(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L11
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(0)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+57)))
	if v59 != 0 {
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
	v61 = m.ExcPending
	if v61 != 0 {
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
	F_tuplestore_end(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L11
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v76 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ReleaseCachedPlan(m, v76, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L11
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v80 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v80)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v80
	return
L34:
	;
	goto L33
}
