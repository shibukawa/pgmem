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
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
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
	var v60 int32
	_ = v60
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
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
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
		v109 = v4
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v109 < v25 {
		goto L41
	} else {
		goto L42
	}
L7:
	;
	v28 = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v29 <= v28 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v109 = v4
	goto L6
L9:
	;
	goto L10
L10:
	;
	v35 = v28
	v38 = v4
	goto L11
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v35<<(uint(int32(2))%32))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+26)))
	if v47 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v109 = v97
	goto L6
L13:
	;
	v51 = v38 + int32(1)
	if v25 < v51 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v97 = v38
	goto L15
L15:
	;
	v100 = v35 + int32(1)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v100 < v101 {
		v35 = v100
		v38 = v97
		goto L11
	} else {
		goto L40
	}
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v54 = F_pstrdup(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v64 = F_exprType(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v57 = F_makeString(m, v54)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v59 = F_lappend(m, v56, v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v59
	goto L18
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v67 = F_exprTypmod(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v70 = F_exprCollation(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	v77 = v74 & base.B2i32(v64 == int32(705))
	if v77 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v78 = int32(25)
	goto L27
L26:
	;
	v78 = v64
	goto L27
L27:
	;
	v79 = F_lappend_oid(m, v72, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v79
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v77 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v84 = int32(-1)
	goto L31
L30:
	;
	v84 = v67
	goto L31
L31:
	;
	v85 = F_lappend_int(m, v82, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v85
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v70 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v90 = v70
	goto L35
L34:
	;
	v90 = int32(100)
	goto L35
L35:
	;
	if v77 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v91 = v90
	goto L38
L37:
	;
	v91 = v70
	goto L38
L38:
	;
	v92 = F_lappend_oid(m, v88, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v92
	v97 = v51
	goto L15
L40:
	;
	goto L12
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	m.G0 = v13 + int32(16)
	return
L44:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v121
	F_errmsg(m, int32(449344), v13)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(489524), int32(639), int32(75018))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
