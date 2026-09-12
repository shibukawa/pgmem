package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgss_ExecutorFinish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = v2
	v17 = v12
	v18 = int32(-1)
	v20 = v2
	v21 = v2
	v22 = v2
	goto L1
L1:
	;
	if v18 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _consts[291])) = v47
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v48
	v98 = int32(4735024)
	v99 = *(*int32)(unsafe.Add(mBase, _consts[1142]))
	*(*int32)(unsafe.Add(mBase, _consts[1142])) = v99 - int32(1)
	m.G0 = v12 + int32(16)
	return
L3:
	;
	v27 = v17 - int32(160)
	m.G0 = v27
	v29 = int32(4735024)
	v30 = *(*int32)(unsafe.Add(mBase, _consts[1142]))
	v31 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1142])) = v30 + v31
	v35 = *(*int32)(unsafe.Add(mBase, _consts[291]))
	v37 = *(*int32)(unsafe.Add(mBase, _consts[74]))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v12 + int32(12)
	goto L6
L4:
	;
	v44 = v16
	v45 = v17
	v46 = v20
	v47 = v21
	v48 = v22
	goto L5
L5:
	;
	goto L8
L6:
	;
	v44 = int32(0)
	v45 = v27
	v46 = v27
	v47 = v35
	v48 = v37
	goto L5
L7:
	;
	goto L2
L8:
	;
	if v44 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L7
L10:
	;
	v73 = int32(m.ExcTag)
	v74 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v73 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	F_standard_ExecutorFinish(m, l0)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L18
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[291])) = v46
	v54 = *(*int32)(unsafe.Add(mBase, _consts[1143]))
	if v54 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v48
	*(*int32)(unsafe.Add(mBase, _consts[291])) = v47
	v63 = int32(4735024)
	v64 = *(*int32)(unsafe.Add(mBase, _consts[1142]))
	*(*int32)(unsafe.Add(mBase, _consts[1142])) = v64 - int32(1)
	F_pg_re_throw(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	m.T0[v54].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L7
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	goto L9
L19:
	;
	v78 = int32(v74)
	m.G0 = v45
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v12+int32(12) == v85 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	m.ExcPending = 1
	goto L28
L21:
	;
	if v88 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v88 = v87
	goto L24
L23:
	;
	v88 = int32(0)
	goto L24
L24:
	;
	goto L21
L25:
	;
	F___wasm_longjmp(m, v81, v80)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v16 = v80
	v17 = v45
	v18 = v88
	v20 = v46
	v21 = v47
	v22 = v48
	goto L1
L28:
	;
	return
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
