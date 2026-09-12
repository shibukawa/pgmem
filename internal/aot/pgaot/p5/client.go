package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PrepareClientEncoding(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	v2 = int32(0)
	if base.Ui32(int32(41)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v112
L2:
	;
	v112 = int32(-1)
	goto L1
L3:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1268])))
	if v10 == int32(0) {
		v112 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l0 == int32(0) {
		v112 = v2
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v17 == l0 {
		v112 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v17 == int32(0) {
		v112 = v2
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	goto L8
L8:
	;
	if v23 == int32(2) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v26 = F_FindDefaultConversionProc(m, l0, v17)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[1269]))
	if v73 == int32(0) {
		goto L2
	} else {
		goto L21
	}
L12:
	;
	return int32(0)
L13:
	;
	if v26 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v33 = F_FindDefaultConversionProc(m, v17, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if v33 == int32(0) {
		v112 = int32(-1)
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[145]))
	v40 = F_MemoryContextAlloc(m, v38, int32(64))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v17
	v47 = *(*int32)(unsafe.Add(mBase, _consts[145]))
	F_fmgr_info_cxt(m, v26, v40+int32(8), v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[145]))
	F_fmgr_info_cxt(m, v33, v40+int32(36), v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v56 = int32(4562096)
	v57 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v60 = *(*int32)(unsafe.Add(mBase, _consts[145]))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v60
	v63 = *(*int32)(unsafe.Add(mBase, _consts[1269]))
	v64 = F_lcons(m, v40, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v57
	*(*int32)(unsafe.Add(mBase, _consts[1269])) = v64
	return int32(0)
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v76 <= int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v79 = int32(0)
	if v79 < v76 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v82 = v76
	goto L25
L24:
	;
	v82 = v79
	goto L25
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v86 = int32(0)
	goto L26
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v83+v86<<(uint(int32(2))%32))))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v95 != v17 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L2
L28:
	;
	v102 = v86 + int32(1)
	if v102 != v82 {
		v86 = v102
		goto L26
	} else {
		goto L31
	}
L29:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v97 != l0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	return int32(0)
L31:
	;
	goto L27
}
