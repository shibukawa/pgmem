package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FreeErrorDataContents(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_pfree(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v6 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	F_pfree(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v9 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	F_pfree(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v12 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	F_pfree(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v15 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_pfree(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v18 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	F_pfree(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v21 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	F_pfree(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v24 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	F_pfree(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v27 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	F_pfree(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v30 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	F_pfree(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v33 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	F_pfree(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v36 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	F_pfree(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	return
L49:
	;
	goto L48
}
