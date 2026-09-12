package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgl_geteuid(m *base.Module) int32 {
	return int32(123)
}
func F_pgl_getpwuid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	*(*int32)(unsafe.Add(mBase, _consts[1240])) = int32(4174290)
	*(*int32)(unsafe.Add(mBase, _consts[1241])) = int32(4174275)
	*(*int32)(unsafe.Add(mBase, _consts[1242])) = int32(4174263)
	*(*int32)(unsafe.Add(mBase, _consts[1243])) = l0
	*(*int32)(unsafe.Add(mBase, _consts[1244])) = l0
	*(*int32)(unsafe.Add(mBase, _consts[1245])) = int32(4174261)
	v18 = int32(4599976)
	*(*int32)(unsafe.Add(mBase, _consts[1246])) = int32(4174252)
	return v18
}
func F_pgl_longjmp(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1238]))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F___wasm_longjmp(m, l0, l1)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v7 = int32(4599808)
	v8 = int32(156)
	goto L6
L3:
	;
	if v70 != 0 {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v70 = int32(0)
	goto L3
L5:
	;
	v44 = v39
	v45 = v40
	v46 = v41
	goto L15
L6:
	;
	if (l0|v7)&int32(3) != 0 {
		v39 = l0
		v40 = v7
		v41 = v8
		goto L5
	} else {
		goto L9
	}
L8:
	;
	if v29 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v16 = l0
	v17 = v7
	v18 = v8
	goto L10
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v21 != v22 {
		v39 = v16
		v40 = v17
		v41 = v18
		goto L5
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	v24 = int32(4)
	v25 = v17 + v24
	v27 = v16 + v24
	v29 = v18 - v24
	if base.Ui32(int32(3)) < base.Ui32(v29) {
		v16 = v27
		v17 = v25
		v18 = v29
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v39 = v27
	v40 = v25
	v41 = v29
	goto L5
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 == v50 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v70 = v49 - v50
	goto L3
L17:
	;
	v52 = int32(1)
	v57 = v46 - v52
	if v57 != 0 {
		v44 = v44 + v52
		v45 = v45 + v52
		v46 = v57
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
	*(*int32)(unsafe.Add(mBase, _consts[1239])) = int32(100)
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
	v8 = int32(4599792)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[1238]))
	*(*int32)(unsafe.Add(mBase, _consts[1238])) = l0
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
	v3 = int32(4174248)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1239]))
	*(*int32)(unsafe.Add(mBase, _consts[1239])) = l0
	return v4
}
func F_pgl_shmat(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1118]))
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
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(28)
	return int32(-1)
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if l0 == v9 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	return v11
L7:
	;
	goto L8
L8:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	if v13 != 0 {
		v7 = v13
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
