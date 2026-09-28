package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_errdetail(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(_a_F_errdetail_0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_errdetail[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail[0])) = v14 + int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_errdetail[1]))
	if int32(0) <= v19 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(_a_F_errdetail_1)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_errdetail[2]))
	v26 = v19 * int32(100)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errdetail[3])))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail[2])) = v29
	v32 = v10 + int32(16)
	F_initStringInfo(m, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail[1])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L21
	}
L4:
	;
	return int32(0)
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errdetail[4])))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail[5])) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v41 = F_appendStringInfoVA(m, v32, l0, l1)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v41 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v47 = v41
	goto L10
L8:
	;
	goto L9
L9:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errdetail[6])))
	if v67 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v51 = v10 + int32(16)
	F_enlargeStringInfo(m, v51, v47)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errdetail[4])))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail[5])) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v58 = F_appendStringInfoVA(m, v51, l0, l1)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	if v58 != 0 {
		v47 = v58
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_pfree(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v71 = F_pstrdup(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errdetail[6]))) = v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_pfree(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail[2])) = v23
	v79 = int32(_a_F_errdetail_0)
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_errdetail[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail[0])) = v81 - int32(1)
	m.G0 = v10 + int32(32)
	return int32(0)
L21:
	;
	F_errmsg_internal(m, int32(_a_F_errdetail_2), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_errdetail_3), int32(1401), int32(_a_F_errdetail_4))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
