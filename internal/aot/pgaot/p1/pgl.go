package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgl_run_atexit_funcs(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v1 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_run_atexit_funcs[0]))
	if v7 <= v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgl_run_atexit_funcs[0])) = int32(0)
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
	v31 = v7
	goto L3
L5:
	;
	goto L6
L6:
	;
	v16 = v7
	v18 = v1
	goto L7
L7:
	;
	v20 = v16 - int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20<<(uint(int32(2))%32))+uint32(_c_F_pgl_run_atexit_funcs[1])))
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v31 = v20
	goto L3
L9:
	;
	m.T0[v23].(func(*base.Module))(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v27 = v18 + int32(1)
	if v27 != v11 {
		v16 = v20
		v18 = v27
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
	v38 = v31
	goto L16
L16:
	;
	v42 = v38 << (uint(int32(2)) % 32)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_pgl_run_atexit_funcs[2])))
	if v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L1
L18:
	;
	m.T0[v45].(func(*base.Module))(m)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_pgl_run_atexit_funcs[3])))
	if v50 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	m.T0[v50].(func(*base.Module))(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L12
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_pgl_run_atexit_funcs[0])))
	if v55 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	m.T0[v55].(func(*base.Module))(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v59 = v38 - int32(4)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v59<<(uint(int32(2))%32))+uint32(_c_F_pgl_run_atexit_funcs[1])))
	if v62 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	m.T0[v62].(func(*base.Module))(m)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if int32(4) < v38 {
		v38 = v59
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
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int64
	_ = v64
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v117 int32
	_ = v117
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_shmctl[0]))
	if v14 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v117
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgl_shmctl[1])) = int32(28)
	v117 = int32(-1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_shmctl[2]))
	F_fiprintf(m, v98, int32(_a_F_pgl_shmctl_0), v11)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
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
	v50 = int32(0)
	if base.B2i32(l2 == v50)|base.B2i32(l1 != int32(2)) == v50 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	F_emscripten_builtin_free(m, v36)
	mBase = m.M
	v117 = int32(0)
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
	*(*int32)(unsafe.Add(mBase, _c_F_pgl_shmctl[0])) = v44
	goto L16
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v60 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v59
	v64 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l2)+48)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(l2)+56)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(l2)+40)) = v64
	v117 = v60
	goto L1
L21:
	;
	goto L22
L22:
	;
	v68 = int32(0)
	if base.B2i32(l2 == v68)|base.B2i32(l1 != int32(1)) == v68 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v75
	v117 = int32(0)
	goto L1
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_shmctl[2]))
	F_fiprintf(m, v80, int32(_a_F_pgl_shmctl_1), v11+int32(16))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	return int32(0)
L27:
	;
	goto L2
L28:
	;
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_system[0]))
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
