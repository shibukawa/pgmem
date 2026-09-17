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
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	v2 = int32(0)
	if base.Ui32(int32(41)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v114
L2:
	;
	v114 = int32(-1)
	goto L1
L3:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[0])))
	v11 = int32(0)
	if base.B2i32(v10 == v11)|base.B2i32(l0 == v11) != 0 {
		v114 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[1]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if base.B2i32(v18 == l0)|base.B2i32(v18 == int32(0)) != 0 {
		v114 = v2
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[2]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	goto L6
L6:
	;
	if v25 == int32(2) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v28 = F_FindDefaultConversionProc(m, l0, v18)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[3]))
	if v75 == int32(0) {
		goto L2
	} else {
		goto L19
	}
L10:
	;
	return int32(0)
L11:
	;
	if v28 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v35 = F_FindDefaultConversionProc(m, v18, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	if v35 == int32(0) {
		v114 = int32(-1)
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[4]))
	v42 = F_MemoryContextAlloc(m, v40, int32(64))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v18
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[4]))
	F_fmgr_info_cxt(m, v28, v42+int32(8), v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[4]))
	F_fmgr_info_cxt(m, v35, v42+int32(36), v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v58 = int32(_a_F_PrepareClientEncoding_0)
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[5]))
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[5])) = v62
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[3]))
	v66 = F_lcons(m, v42, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[5])) = v59
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareClientEncoding[3])) = v66
	return int32(0)
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v78 <= int32(0) {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v81 = int32(0)
	if v81 < v78 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v84 = v78
	goto L23
L22:
	;
	v84 = v81
	goto L23
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v88 = int32(0)
	goto L24
L24:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v85+v88<<(uint(int32(2))%32))))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	if v97 != v18 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L2
L26:
	;
	v104 = v88 + int32(1)
	if v104 != v84 {
		v88 = v104
		goto L24
	} else {
		goto L29
	}
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v99 != l0 {
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
