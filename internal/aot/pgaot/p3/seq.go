package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecSeqScanWithQualProject(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
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
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 float64
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	F_MemoryContextReset(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L4
L3:
	;
	m.G0 = v14 + int32(16)
	return v163
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[0]))
	if v36 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L40
	}
L6:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v42 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L8
L10:
	;
	goto L5
L11:
	;
	v45 = F_ScanRelIsReadOnly(m, l0)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v74 = v42
	goto L13
L13:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+188))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	v82 = m.T0[v81].(func(*base.Module, int32, int32, int32) int32)(m, v74, v41, v39)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L24
	}
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[1]))
	if v48 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[2])))
	if v50&int32(1) == int32(0) {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v57 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v40)+132))
	if v45 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v67 = int32(1473)
	goto L21
L20:
	;
	v67 = int32(449)
	goto L21
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v55)+188))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v71 = m.T0[v70].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v55, v56, v57, v57, v57, v60<<(uint(int32(7))%32)&int32(2048)|v67)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v71
	v74 = v71
	goto L13
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v39
	v102 = int32(_a_F_ExecSeqScanWithQualProject_0)
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[3]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[3])) = v105
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v110 = m.T0[v109].(func(*base.Module, int32, int32, int32) int64)(m, v17, v18, v14+int32(15))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L30
	}
L24:
	;
	v84 = int32(0)
	if base.B2i32(v82 == v84)|base.B2i32(v39 == v84) == v84 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)))
	if v91&int32(2) == int32(0) {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	m.T0[v98].(func(*base.Module, int32))(m, v96)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v163 = v96
	goto L3
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[3])) = v103
	if v110 != int64(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	m.T0[v119].(func(*base.Module, int32))(m, v117)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v142 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v122 = int32(_a_F_ExecSeqScanWithQualProject_0)
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[3]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[3])) = v125
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	v131 = m.T0[v130].(func(*base.Module, int32, int32, int32) int64)(m, v16+int32(8), v116, int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[3])) = v123
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+4)))
	v137 = v135 & int32(_a_F_ExecSeqScanWithQualProject_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v117)+4)) = uint16(v137)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	*(*uint16)(unsafe.Add(mBase, uint32(v117)+6)) = uint16(v140)
	v163 = v117
	goto L3
L36:
	;
	v143 = *(*float64)(unsafe.Add(mBase, uint32(v142)+424))
	*(*float64)(unsafe.Add(mBase, uint32(v142)+424)) = base.F64_add(v143, float64(1))
	goto L38
L37:
	;
	goto L38
L38:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	F_MemoryContextReset(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L4
L40:
	;
	F_errmsg_internal(m, int32(_a_F_ExecSeqScanWithQualProject_2), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_ExecSeqScanWithQualProject_3), int32(931), int32(_a_F_ExecSeqScanWithQualProject_4))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
