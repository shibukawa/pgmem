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
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+8)))
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
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v84 int32
	_ = v84
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
		v84 = v8
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v84
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
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v26 == int32(0) {
		v45 = v25
		v46 = v26
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	v49 = F_strlen(m, v20)
	mBase = m.M
	v50 = F_strlen(m, v19)
	mBase = m.M
	v51 = v49
	v52 = v50
	goto L17
L9:
	;
	return v46 - v45
L10:
	;
	goto L9
L11:
	;
	if v25 != v26 {
		v45 = v25
		v46 = v26
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v30 = v20
	v31 = v19
	goto L13
L13:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v35 == int32(0) {
		v45 = v34
		v46 = v35
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v45 = v34
	v46 = v35
	goto L10
L15:
	;
	v38 = int32(1)
	if v34 == v35 {
		v30 = v30 + v38
		v31 = v31 + v38
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v58 = int32(1)
	v59 = v52 - v58
	v61 = v51 - v58
	if v61 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v61 < v59 {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	goto L18
L20:
	;
	if v59 < int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v20))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v19))))
	if base.Ui32(v67) < base.Ui32(v69) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(-1)
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(v67) <= base.Ui32(v69) {
		v51 = v61
		v52 = v59
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v84 = v8
	goto L4
L26:
	;
	return int32(-1)
L27:
	;
	goto L28
L28:
	;
	v84 = base.B2i32(v59 < v61)
	goto L4
}
