package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgl_poll(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return l1
}
func F_pgl_run_atexit_funcs(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v1 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[773]))
	if v7 <= v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[773])) = int32(0)
	return
L2:
	;
	v11 = v7 & int32(3)
	if v11 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if base.Ui32(v7) <= base.Ui32(int32(3)) {
		goto L1
	} else {
		goto L15
	}
L4:
	;
	v32 = v7
	goto L3
L5:
	;
	goto L6
L6:
	;
	v14 = v1
	v15 = v7
	goto L7
L7:
	;
	v20 = v15 - int32(1)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20<<(uint(int32(2))%32))+uint32(_consts[774])))
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v32 = v20
	goto L3
L9:
	;
	m.T0[v25].(func(*base.Module))(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v29 = v14 + int32(1)
	if v29 != v11 {
		v14 = v29
		v15 = v20
		goto L7
	} else {
		goto L14
	}
L12:
	;
	return
L13:
	;
	goto L11
L14:
	;
	goto L8
L15:
	;
	v39 = v32
	goto L16
L16:
	;
	v44 = v39 << (uint(int32(2)) % 32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[1276])))
	if v47 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L1
L18:
	;
	m.T0[v47].(func(*base.Module))(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[1277])))
	if v52 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	m.T0[v52].(func(*base.Module))(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v56 = v39 - int32(3)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56<<(uint(int32(2))%32))+uint32(_consts[774])))
	if v61 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	m.T0[v61].(func(*base.Module))(m)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v65 = v39 - int32(4)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65<<(uint(int32(2))%32))+uint32(_consts[774])))
	if v70 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	m.T0[v70].(func(*base.Module))(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(int32(1)) < base.Ui32(v56) {
		v39 = v65
		goto L16
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	goto L17
}
func F_pgl_set_system_fn(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[1275])) = l0
	return
}
func F_pgl_setsockopt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	return int32(0)
}
func F_pgl_shmctl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int64
	_ = v61
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v114 int32
	_ = v114
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[764]))
	if v14 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v114
L2:
	;
	v114 = int32(-1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	v95 = *(*int32)(unsafe.Add(mBase, _consts[463]))
	F_fiprintf(m, v95, int32(783935), v11)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L26
	} else {
		goto L28
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if l0 == v17 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if l1 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v36 = v14
	v38 = v4
	goto L5
L7:
	;
	goto L8
L8:
	;
	v22 = v14
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	if v27 == int32(0) {
		goto L3
	} else {
		goto L11
	}
L10:
	;
	v36 = v27
	v38 = v22
	goto L5
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v30 != l0 {
		v22 = v27
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	F_emscripten_builtin_free(m, v42)
	mBase = m.M
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	if v38 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	if l1 != int32(2) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	F_emscripten_builtin_free(m, v36)
	mBase = m.M
	v114 = int32(0)
	goto L1
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v44
	goto L16
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[764])) = v44
	goto L16
L20:
	;
	if l1 != int32(1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	if l2 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v57 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v56
	v61 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l2)+48)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(l2)+56)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(l2)+40)) = v61
	v114 = v57
	goto L1
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	v74 = *(*int32)(unsafe.Add(mBase, _consts[463]))
	F_fiprintf(m, v74, int32(784042), v11+int32(16))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	if l2 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v69
	v114 = int32(0)
	goto L1
L26:
	;
	return int32(0)
L27:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(28)
	goto L2
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(28)
	goto L2
}
func F_pgl_system(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1275]))
	if v4 == int32(0) {
		return int32(123)
	} else {
		v9 = m.T0[v4].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
