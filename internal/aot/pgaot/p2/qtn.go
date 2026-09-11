package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_QTNCopy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	F_check_stack_depth(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = F_palloc(m, int32(24))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v17
	v20 = F_palloc(m, int32(12))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v20
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v29 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v28 | v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v33 == v29 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v41 = F_palloc(m, v36&int32(4095)+int32(1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v67 = F_palloc(m, v64<<(uint(int32(2))%32))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v41
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v48 = v46 & int32(4095)
	if v48 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v51+v53&int32(4095)))) = uint8(v57)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v59 | int32(4)
	return v11
L10:
	;
	v49 = F__emscripten_memcpy_bulkmem(m, v41, v44, v48)
	mBase = m.M
	goto L12
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(0) < v70 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v76 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	return v11
L17:
	;
	v80 = v76 << (uint(int32(2)) % 32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80+v81)))
	v84 = F_QTNCopy(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v86+v80))) = v84
	v90 = v76 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v90 < v91 {
		v76 = v90
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
}
