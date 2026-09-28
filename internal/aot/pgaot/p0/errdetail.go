package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_errdetail_internal(m *base.Module, l0 int32, l1 int32) {
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(_a_F_errdetail_internal_0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_errdetail_internal[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail_internal[0])) = v14 + int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_errdetail_internal[1]))
	if int32(0) <= v19 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(_a_F_errdetail_internal_1)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_errdetail_internal[2]))
	v26 = v19 * int32(100)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errdetail_internal[3])))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail_internal[2])) = v29
	v32 = v10 + int32(16)
	F_initStringInfo(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail_internal[1])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L21
	}
L4:
	;
	return
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errdetail_internal[4])))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail_internal[5])) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v39 = F_appendStringInfoVA(m, v32, l0, l1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v39 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v45 = v39
	goto L10
L8:
	;
	goto L9
L9:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errdetail_internal[6])))
	if v65 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v49 = v10 + int32(16)
	F_enlargeStringInfo(m, v49, v45)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errdetail_internal[4])))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail_internal[5])) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v56 = F_appendStringInfoVA(m, v49, l0, l1)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	if v56 != 0 {
		v45 = v56
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_pfree(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v69 = F_pstrdup(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errdetail_internal[6]))) = v69
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_pfree(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail_internal[2])) = v23
	v77 = int32(_a_F_errdetail_internal_0)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_errdetail_internal[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail_internal[0])) = v79 - int32(1)
	m.G0 = v10 + int32(32)
	return
L21:
	;
	F_errmsg_internal(m, int32(_a_F_errdetail_internal_2), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_errdetail_internal_3), int32(1428), int32(_a_F_errdetail_internal_4))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
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
