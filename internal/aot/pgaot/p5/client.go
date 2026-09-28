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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	v2 = int32(0)
	if base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l0)) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v117
L2:
	;
	v117 = int32(-1)
	goto L1
L3:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[0])))
	v14 = int32(0)
	if base.B2i32(v13 == v14)|base.B2i32(l0 == v14) != 0 {
		v117 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[1]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if base.B2i32(v21 == l0)|base.B2i32(v21 == int32(0)) != 0 {
		v117 = v2
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[2]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	goto L6
L6:
	;
	if v28 == int32(2) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = F_FindDefaultConversionProc(m, l0, v21)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[3]))
	if v78 == int32(0) {
		goto L2
	} else {
		goto L19
	}
L10:
	;
	return int32(0)
L11:
	;
	if v31 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v38 = F_FindDefaultConversionProc(m, v21, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	if v38 == int32(0) {
		v117 = int32(-1)
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[4]))
	v45 = F_MemoryContextAlloc(m, v43, int32(64))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v21
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[4]))
	F_fmgr_info_cxt(m, v31, v45+int32(8), v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[4]))
	F_fmgr_info_cxt(m, v38, v45+int32(36), v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v61 = int32(_a_F_PrepareClientEncoding_0)
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[5]))
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[5])) = v65
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[3]))
	v69 = F_lcons(m, v45, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[5])) = v62
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[3])) = v69
	return int32(0)
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v81 <= int32(0) {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v84 = int32(0)
	if v84 < v81 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v87 = v81
	goto L23
L22:
	;
	v87 = v84
	goto L23
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v91 = int32(0)
	goto L24
L24:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v88+v91<<(uint(int32(2))%32))))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v100 != v21 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L2
L26:
	;
	v107 = v91 + int32(1)
	if v107 != v87 {
		v91 = v107
		goto L24
	} else {
		goto L29
	}
L27:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v102 != l0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	return int32(0)
L29:
	;
	goto L25
}
