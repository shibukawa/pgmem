package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgl_longjmp(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_longjmp[0]))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F___wasm_longjmp(m, l0, int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v6 = int32(_a_F_pgl_longjmp_0)
	v7 = int32(156)
	goto L6
L3:
	;
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v69 = int32(0)
	goto L3
L5:
	;
	v43 = v38
	v44 = v39
	v45 = v40
	goto L15
L6:
	;
	if (l0|v6)&int32(3) != 0 {
		v38 = l0
		v39 = v6
		v40 = v7
		goto L5
	} else {
		goto L9
	}
L8:
	;
	if v28 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v15 = l0
	v16 = v6
	v17 = v7
	goto L10
L10:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v20 != v21 {
		v38 = v15
		v39 = v16
		v40 = v17
		goto L5
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	v23 = int32(4)
	v24 = v16 + v23
	v26 = v15 + v23
	v28 = v17 - v23
	if base.Ui32(int32(3)) < base.Ui32(v28) {
		v15 = v26
		v16 = v24
		v17 = v28
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v38 = v26
	v39 = v24
	v40 = v28
	goto L5
L15:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v48 == v49 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v69 = v48 - v49
	goto L3
L17:
	;
	v51 = int32(1)
	v56 = v45 - v51
	if v56 != 0 {
		v43 = v43 + v51
		v44 = v44 + v51
		v45 = v56
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L4
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgl_longjmp[1])) = int32(100)
	m.Env.Emscripten_exit_with_live_runtime(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	return
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgl_setPGliteActive(m *base.Module, l0 int32) int32 {
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
	var v14 int64
	_ = v14
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = int32(_a_F_pgl_setPGliteActive_0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_setPGliteActive[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgl_setPGliteActive[0])) = l0
	if l0 == int32(0) {
		v14 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v14
		*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v14
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v14
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v14
		v22 = F_setitimer(m, v6)
		mBase = m.M
	} else {
	}
	m.G0 = v6 + int32(32)
	return v9
}
func F_pgl_setPGliteExitStatus(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(_a_F_pgl_setPGliteExitStatus_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_setPGliteExitStatus[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgl_setPGliteExitStatus[0])) = l0
	return v4
}
