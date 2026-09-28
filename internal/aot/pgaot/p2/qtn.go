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
	var v24 int32
	_ = v24
	var v26 int64
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v24
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v26
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
	v65 = F_palloc_mul(m, int32(4), v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L12
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v41
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v47 = v45 & int32(4095)
	if v47 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	base.MemoryCopy(m, v41, v48, v47)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v56 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v50+v52&int32(4095)))) = uint8(v56)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v58 | int32(4)
	return v11
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v65
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(0) < v68 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v74 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	return v11
L16:
	;
	v78 = v74 << (uint(int32(2)) % 32)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78+v79)))
	v82 = F_QTNCopy(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v84+v78))) = v82
	v88 = v74 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v88 < v89 {
		v74 = v88
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
}
