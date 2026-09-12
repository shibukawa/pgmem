package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SS_finalize_plan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_finalize_plan(m, l0, l1, int32(-1), v4, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_SS_replace_correlation_vars(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_replace_correlation_vars_mutator(m, l1, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_ss_report_location(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	if l1&int32(15) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v15 = F_LWLockConditionalAcquire(m, v11+int32(3072), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v15 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v28 = v24
	goto L7
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = l1
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v44 != v28 {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v32 != v19 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v21
	goto L6
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v38 != 0 {
		v28 = v38
		goto L7
	} else {
		goto L13
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v34 != v20 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v36 == v21 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	goto L8
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v47 == v28 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v65+int32(3072))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L23
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v46
	goto L19
L18:
	;
	goto L19
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v50
	if v50 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v52
	goto L22
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v28
	goto L16
L23:
	;
	goto L1
}
