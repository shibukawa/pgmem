package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ProcessStartupProcInterrupts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	v1 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[578]))
	if v9 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[233]))
	if v127 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _consts[578])) = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[579]))
	v17 = F_pstrdup(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[580]))
	v21 = F_pstrdup(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _consts[581])))
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[579]))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v33 == int32(0) {
		v52 = v32
		v53 = v33
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[580]))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v60 == int32(0) {
		v79 = v59
		v80 = v60
		goto L17
	} else {
		goto L18
	}
L8:
	;
	goto L7
L9:
	;
	if v32 != v33 {
		v52 = v32
		v53 = v33
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v37 = v17
	v38 = v29
	goto L11
L11:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v42 == int32(0) {
		v52 = v41
		v53 = v42
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v52 = v41
	v53 = v42
	goto L8
L13:
	;
	v45 = int32(1)
	if v41 == v42 {
		v37 = v37 + v45
		v38 = v38 + v45
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	F_pfree(m, v17)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L3
	} else {
		goto L26
	}
L16:
	;
	if v81 != 0 {
		v86 = v1
		goto L15
	} else {
		goto L24
	}
L17:
	;
	v81 = v80 - v79
	goto L16
L18:
	;
	if v59 != v60 {
		v79 = v59
		v80 = v60
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v64 = v21
	v65 = v56
	goto L20
L20:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v69 == int32(0) {
		v79 = v68
		v80 = v69
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v79 = v68
	v80 = v69
	goto L17
L22:
	;
	v72 = int32(1)
	if v68 == v69 {
		v64 = v64 + v72
		v65 = v65 + v72
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v82 != 0 {
		v86 = v1
		goto L15
	} else {
		goto L25
	}
L25:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, _consts[581])))
	v86 = base.B2i32(v24 != v84)
	goto L15
L26:
	;
	F_pfree(m, v21)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	if v53-v52|v81|v86 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[582]))
	if v96 != int32(3) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	goto L1
L30:
	;
	v99 = F_WalRcvRunning(m)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	if v99 == int32(0) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v105 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	if v105 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_errmsg(m, int32(462488), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L3
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v117 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[583])) = uint8(v117)
	goto L29
L37:
	;
	F_errfinish(m, int32(515869), int32(4422), int32(88403))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	F_pgl_exit(m, int32(1))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L57
	}
L40:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, _consts[96])))
	if v131 != int32(1) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L3
	} else {
		goto L56
	}
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	if v148 != 0 {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	v134 = int32(4470804)
	v136 = *(*int32)(unsafe.Add(mBase, _consts[584]))
	*(*int32)(unsafe.Add(mBase, _consts[584])) = v136 + int32(1)
	if v136&int32(1023) != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v142 = F_PostmasterIsAliveInternal(m)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	if v142 == int32(0) {
		goto L39
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[525]))
	if v152 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	F_ProcessLogMemoryContextInterrupt(m)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	return
L55:
	;
	goto L54
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_append_startup_cost_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if v8 != v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v8 < v10 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v17 = int32(1)
	v18 = *(*float64)(unsafe.Add(mBase, uint32(v7)+48))
	v19 = *(*float64)(unsafe.Add(mBase, uint32(v9)+48))
	if base.F64_lt(v18, v19) != 0 {
		v84 = v17
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v15 = int32(1)
	goto L6
L5:
	;
	v15 = int32(-1)
	goto L6
L6:
	;
	return v15
L7:
	;
	return v84
L8:
	;
	if base.F64_gt(v18, v19) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(-1)
L10:
	;
	goto L11
L11:
	;
	v24 = *(*float64)(unsafe.Add(mBase, uint32(v7)+56))
	v25 = *(*float64)(unsafe.Add(mBase, uint32(v9)+56))
	if base.F64_lt(v24, v25) != 0 {
		v84 = v17
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if base.F64_gt(v24, v25) != 0 {
		v84 = int32(-1)
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v30 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v84 = v83
	goto L7
L15:
	;
	if v32 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	if v32 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v40 = int32(-1)
	goto L20
L19:
	;
	v40 = int32(0)
	goto L20
L20:
	;
	v83 = v40
	goto L14
L21:
	;
	v83 = int32(1)
	goto L14
L22:
	;
	goto L23
L23:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v44 != v45 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v45 < v44 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v51 = int32(8)
	v58 = v44 - int32(1)
	goto L31
L27:
	;
	v50 = int32(1)
	goto L29
L28:
	;
	v50 = int32(-1)
	goto L29
L29:
	;
	v83 = v50
	goto L14
L30:
	;
	if base.Ui32(v67) < base.Ui32(v65) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v63 = v58 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v30+v51+v63)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+(v32+v51))))
	if v65 != v67 {
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v83 = int32(0)
	goto L14
L33:
	;
	v70 = v58 - int32(1)
	if int32(0) <= v70 {
		v58 = v70
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v77 = int32(1)
	goto L37
L36:
	;
	v77 = int32(-1)
	goto L37
L37:
	;
	v83 = v77
	goto L14
}
func F_disable_startup_progress_timeout(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	if v2 != 0 {
		F_disable_timeout(m, int32(12))
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[585])) = int32(0)
			return
		}
	} else {
		return
	}
}
func F_process_startup_options(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = int32(3)
	goto L3
L2:
	;
	v17 = int32(4)
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = F_strlen(m, v18)
	mBase = m.M
	v28 = F_palloc(m, (v19<<(uint(int32(1))%32)+int32(2))&int32(-4)+int32(8))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	if v175 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(171754)
	v35 = v13 + int32(12)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v37 = m.G0
	v39 = v37 - int32(16)
	m.G0 = v39
	F_initStringInfo(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v43 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	F_pfree(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L7
	} else {
		goto L35
	}
L11:
	;
	v49 = v36
	goto L12
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v57
	goto L14
L14:
	;
	v67 = v49
	goto L15
L15:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if base.Ui32(v74-int32(9)) < base.Ui32(int32(5)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v67 = v67 + int32(1)
	goto L15
L18:
	;
	if v74 == int32(32) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	if v74 == int32(0) {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v85 = v74
	v86 = v67
	v87 = int32(0)
	goto L21
L21:
	;
	v94 = v85 & int32(255)
	if v87&int32(1) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v127 = F_pstrdup(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L33
	}
L23:
	;
	goto L22
L24:
	;
	if base.B2i32(v94 != int32(32))&base.B2i32(base.Ui32((v85-int32(14))&int32(255)) < base.Ui32(int32(251))) == int32(0) {
		v124 = v86
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v112 = v87 | base.B2i32(v94 != int32(92))
	if v112&int32(1) != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	F_appendStringInfoChar(m, v39, base.I32_extend8_s(v85))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L7
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v118 = int32(1)
	v121 = v86 + v118
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if v122 != 0 {
		v85 = v122
		v86 = v121
		v87 = v112 ^ v118
		goto L21
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v124 = v121
	goto L23
L33:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v129 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28+v129<<(uint(int32(2))%32)))) = v127
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v137 != 0 {
		v49 = v124
		goto L12
	} else {
		goto L34
	}
L34:
	;
	goto L10
L35:
	;
	m.G0 = v39 + int32(16)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v160 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28+v156<<(uint(int32(2))%32)))) = v160
	F_process_postgres_switches(m, v156, v28, v17, v160)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	goto L6
L37:
	;
	m.G0 = v13 + int32(16)
	return
L38:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	if v178 == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v182 = v178
	goto L40
L40:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v193 = v182 + int32(4)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v200 = v196 + v197<<(uint(int32(2))%32)
	if base.Ui32(v193) < base.Ui32(v200) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L37
L42:
	;
	v202 = v193
	goto L44
L43:
	;
	v202 = int32(0)
	goto L44
L44:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	F_SetConfigOption(m, v191, v203, v17, int32(9))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	v208 = v202 + int32(4)
	if base.Ui32(v200) <= base.Ui32(v208) {
		goto L37
	} else {
		goto L46
	}
L46:
	;
	if v208 != 0 {
		v182 = v208
		goto L40
	} else {
		goto L47
	}
L47:
	;
	goto L41
}
func F_startup_progress_timeout_handler(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[585])) = int32(1)
	return
}
