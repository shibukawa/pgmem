package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlruDeleteSegment(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[147]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v12 = F_LWLockAcquire(m, v10, v2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v18 = v14
	v19 = v2
	goto L3
L3:
	;
	v22 = int32(0)
	if v18 <= v22 {
		v91 = v19
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_SlruInternalDeleteSegment(m, int32(4415060), l0)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L28
	}
L5:
	;
	goto L4
L6:
	;
	v28 = v22
	v30 = v19
	v31 = v22
	goto L7
L7:
	;
	v34 = int32(base.Ui32(v28) >> (uint(int32(4)) % 32))
	if v30 != v34 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v82 != 0 {
		v18 = v85
		v19 = v49
		goto L3
	} else {
		goto L27
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_LWLockRelease(m, v36+v30<<(uint(int32(7))%32))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v49 = v30
	goto L11
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v53 = v50 + v28<<(uint(int32(2))%32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v54 == int32(0) {
		v82 = v31
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v47 = F_LWLockAcquire(m, v42+v34<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v49 = v34
	goto L11
L14:
	;
	v84 = v28 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v84 < v85 {
		v28 = v84
		v30 = v49
		v31 = v82
		goto L7
	} else {
		goto L26
	}
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v57+v28<<(uint(int32(3))%32))))
	v63 = base.I64_div_s(v61, int64(32))
	if v63 != l0 {
		v82 = v31
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if v54 == int32(2) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v82 = int32(1)
	goto L14
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v28))))
	if v69 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	F_SimpleLruWaitIO(m, int32(4415060), v28)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L25
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = int32(0)
	v82 = v31
	goto L14
L22:
	;
	goto L23
L23:
	;
	F_SlruInternalWritePage(m, int32(4415060), v28, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L17
L25:
	;
	goto L17
L26:
	;
	goto L8
L27:
	;
	v91 = v49
	goto L5
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_LWLockRelease(m, v97+v91<<(uint(int32(7))%32))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	return
}
