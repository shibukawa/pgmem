package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DataChecksumsEnabled(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+252))
	return base.B2i32(v3 != int32(0))
}
func F_copy_read_data(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v15 != v16 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v20 = v15 - v16
	if v20 < l2 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v29 = l2
	v30 = int32(0)
	goto L3
L3:
	;
	if v29 <= int32(0) {
		v110 = v30
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v22 = v20
	goto L6
L5:
	;
	v22 = l2
	goto L6
L6:
	;
	if v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v25 + v22
	v29 = l2 - v22
	v30 = v22
	goto L3
L8:
	;
	v23 = F__emscripten_memcpy_bulkmem(m, l0, v18+v16, v22)
	mBase = m.M
	goto L10
L9:
	;
	goto L10
L10:
	;
	goto L7
L11:
	;
	m.G0 = v11 + int32(16)
	return v110
L12:
	;
	if l1 <= v30 {
		v110 = v30
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v35 = l0
	v37 = v29
	v39 = v30
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(0)
	v47 = v35
	v49 = v37
	v51 = v39
	goto L16
L15:
	;
	v110 = v51
	goto L11
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[683]))
	v62 = *(*int32)(unsafe.Add(mBase, _consts[493]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+40))
	v64 = m.T0[v63].(func(*base.Module, int32, int32, int32) int32)(m, v56, v11+int32(8), v11+int32(12))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v99 = F_WaitLatchOrSocket(m, v95, v96, int32(1000), int32(134217759))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L18
	} else {
		goto L37
	}
L18:
	;
	return int32(0)
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v69 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v64 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	if v64 < int32(0) {
		v110 = v51
		goto L11
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L17
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v76 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v74
	if v64 < v49 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v82 = v64
	goto L30
L29:
	;
	v82 = v49
	goto L30
L30:
	;
	if v82 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v85 + v82
	v88 = v82 + v51
	v89 = v49 - v82
	if v89 <= int32(0) {
		v110 = v88
		goto L11
	} else {
		goto L35
	}
L32:
	;
	v83 = F__emscripten_memcpy_bulkmem(m, v47, v74, v82)
	mBase = m.M
	v84 = v83
	goto L34
L33:
	;
	v84 = v47
	goto L34
L34:
	;
	goto L31
L35:
	;
	if v88 < l1 {
		v47 = v82 + v84
		v49 = v89
		v51 = v88
		goto L16
	} else {
		goto L36
	}
L36:
	;
	v110 = v88
	goto L11
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = int32(0)
	goto L38
L38:
	;
	if v51 < l1 {
		v35 = v47
		v37 = v49
		v39 = v51
		goto L14
	} else {
		goto L39
	}
L39:
	;
	goto L15
}
