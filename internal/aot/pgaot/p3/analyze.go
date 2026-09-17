package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_analyzeCTETargetList(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v16 = F_copyObjectImpl(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(l1)+44)) = int64(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v25 = v24
	goto L5
L4:
	;
	v25 = v4
	goto L5
L5:
	;
	if l2 == int32(0) {
		v110 = v4
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v110 < v25 {
		goto L39
	} else {
		goto L40
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v28 <= int32(0) {
		v110 = v4
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v38 = v4
	v39 = v4
	goto L9
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v38<<(uint(int32(2))%32))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+26)))
	if v46 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v110 = v96
	goto L6
L11:
	;
	v50 = v39 + int32(1)
	if v25 < v50 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v96 = v39
	goto L13
L13:
	;
	v99 = v38 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v99 < v100 {
		v38 = v99
		v39 = v96
		goto L9
	} else {
		goto L38
	}
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v53 = F_pstrdup(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v63 = F_exprType(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L20
	}
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v56 = F_makeString(m, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v58 = F_lappend(m, v55, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v58
	goto L16
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v66 = F_exprTypmod(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v69 = F_exprCollation(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	v76 = v73 & base.B2i32(v63 == int32(705))
	if v76 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v77 = int32(25)
	goto L25
L24:
	;
	v77 = v63
	goto L25
L25:
	;
	v78 = F_lappend_oid(m, v71, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v76 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v83 = int32(-1)
	goto L29
L28:
	;
	v83 = v66
	goto L29
L29:
	;
	v84 = F_lappend_int(m, v81, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v84
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v69 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v89 = v69
	goto L33
L32:
	;
	v89 = int32(100)
	goto L33
L33:
	;
	if v76 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v90 = v89
	goto L36
L35:
	;
	v90 = v69
	goto L36
L36:
	;
	v91 = F_lappend_oid(m, v87, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v91
	v96 = v50
	goto L13
L38:
	;
	goto L10
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	m.G0 = v13 + int32(16)
	return
L42:
	;
	F_errcode(m, int32(_a_F_analyzeCTETargetList_0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v120
	F_errmsg(m, int32(_a_F_analyzeCTETargetList_1), v13)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_analyzeCTETargetList_2), int32(639), int32(_a_F_analyzeCTETargetList_3))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
