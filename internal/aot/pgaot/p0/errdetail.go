package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_errdetail_internal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(4555020)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[902]))
	*(*int32)(unsafe.Add(mBase, _consts[902])) = v13 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[899]))
	if int32(0) <= v18 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(4562096)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v25 = v18 * int32(100)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[903])))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v28
	F_initStringInfo(m, v9+int32(16))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[899])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L21
	}
L4:
	;
	return
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[900])))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v40 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v45 = v40
	goto L10
L8:
	;
	goto L9
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[906])))
	if v66 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	F_enlargeStringInfo(m, v9+int32(16), v45)
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
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[900])))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v58 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
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
		v45 = v58
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_pfree(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v70 = F_pstrdup(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[906]))) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	F_pfree(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
	v78 = int32(4555020)
	v80 = *(*int32)(unsafe.Add(mBase, _consts[902]))
	*(*int32)(unsafe.Add(mBase, _consts[902])) = v80 - int32(1)
	m.G0 = v9 + int32(32)
	return
L21:
	;
	F_errmsg_internal(m, int32(476389), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(523980), int32(1237), int32(327270))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
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
