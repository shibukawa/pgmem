package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cfb_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l2 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = l1
	v26 = l2
	v27 = l3
	v29 = v23
	goto L3
L3:
	;
	if int32(0) < v29 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v63 = l0 + int32(20)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v66 = v25
	v67 = v26
	v68 = v27
	v71 = v64
	goto L21
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v39 = v38 - v29
	if v26 < v39 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	v41 = v26
	goto L10
L9:
	;
	v41 = v39
	goto L10
L10:
	;
	v42 = m.T0[l4].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v25, v41, v27)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v44 = v26 - v42
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v45 == v46 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v45 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v53 = v45
	goto L15
L15:
	;
	if int32(0) < v44 {
		v25 = v25 + v42
		v26 = v44
		v27 = v27 + v42
		v29 = v53
		goto L3
	} else {
		goto L20
	}
L16:
	;
	v50 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v50
	v53 = v50
	goto L15
L17:
	;
	v48 = F__emscripten_memcpy_bulkmem(m, l0+int32(20), l0+int32(84), v45)
	mBase = m.M
	goto L19
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L1
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v82 = m.T0[v81].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v77, int32(0), v63, v71, l0+int32(52), v15+int32(12))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L11
	} else {
		goto L23
	}
L22:
	;
	goto L1
L23:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v84 <= int32(4) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v84 + int32(1)
	goto L26
L25:
	;
	goto L26
L26:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v67 < v90 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v92 = v67
	goto L29
L28:
	;
	v92 = v90
	goto L29
L29:
	;
	v93 = m.T0[l4].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v66, v92, v68)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v95 = v67 - v93
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v96 == v97 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v96 != 0 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	if int32(0) < v95 {
		v66 = v66 + v93
		v67 = v95
		v68 = v68 + v93
		v71 = v97
		goto L21
	} else {
		goto L38
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	goto L33
L35:
	;
	v99 = F__emscripten_memcpy_bulkmem(m, v63, l0+int32(84), v96)
	mBase = m.M
	goto L37
L36:
	;
	goto L37
L37:
	;
	goto L34
L38:
	;
	goto L22
}
