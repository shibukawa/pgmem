package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cmpNodePtr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+12)))
	return v3 - v4
}
func F_cmp_fxid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(base.Ui64(v6) < base.Ui64(v5)) - base.B2i32(base.Ui64(v5) < base.Ui64(v6))
}
func F_cmp_lsn(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	return base.B2i32(base.Ui64(v6) < base.Ui64(v5)) - base.B2i32(base.Ui64(v5) < base.Ui64(v6))
}
func F_cmpaffix(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v83 int32
	_ = v83
	v8 = int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = v9 & v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v14 = v12 & v8
	if base.Ui32(v11) < base.Ui32(v14) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-1)
L2:
	;
	goto L3
L3:
	;
	if base.Ui32(v14) < base.Ui32(v11) {
		v83 = v8
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v83
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if base.B2i32(v25 == int32(0))|base.B2i32(v25 != v28) != 0 {
		v46 = v25
		v47 = v28
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	v50 = F_strlen(m, v20)
	mBase = m.M
	v51 = F_strlen(m, v19)
	mBase = m.M
	v52 = v50
	v53 = v51
	goto L16
L9:
	;
	return v46 - v47
L10:
	;
	goto L9
L11:
	;
	v31 = v20
	v32 = v19
	goto L12
L12:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	if v36 == int32(0) {
		v46 = v36
		v47 = v35
		goto L10
	} else {
		goto L14
	}
L13:
	;
	v46 = v36
	v47 = v35
	goto L10
L14:
	;
	v39 = int32(1)
	if v36 == v35 {
		v31 = v31 + v39
		v32 = v32 + v39
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v59 = int32(1)
	v60 = v52 - v59
	v62 = v53 - v59
	if int32(0) <= v60|v62 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v60 < v62 {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v20))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v19))))
	if base.Ui32(v67) < base.Ui32(v69) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	return int32(-1)
L22:
	;
	goto L23
L23:
	;
	if base.Ui32(v67) <= base.Ui32(v69) {
		v52 = v60
		v53 = v62
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v83 = v8
	goto L4
L25:
	;
	return int32(-1)
L26:
	;
	goto L27
L27:
	;
	v83 = base.B2i32(v62 < v60)
	goto L4
}
