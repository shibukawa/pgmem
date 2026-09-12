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
	v9 = *(*int32)(unsafe.Add(mBase, _consts[573]))
	if v9 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	if v127 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _consts[573])) = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[574]))
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
	v20 = *(*int32)(unsafe.Add(mBase, _consts[575]))
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
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _consts[576])))
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
	v29 = *(*int32)(unsafe.Add(mBase, _consts[574]))
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
	v56 = *(*int32)(unsafe.Add(mBase, _consts[575]))
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
	v84 = int32(*(*uint8)(unsafe.Add(mBase, _consts[576])))
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
	v96 = *(*int32)(unsafe.Add(mBase, _consts[577]))
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
	F_errmsg(m, int32(436434), int32(0))
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
	*(*uint8)(unsafe.Add(mBase, _consts[578])) = uint8(v117)
	goto L29
L37:
	;
	F_errfinish(m, int32(486941), int32(4422), int32(81493))
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
	v131 = int32(*(*uint8)(unsafe.Add(mBase, _consts[92])))
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
	v148 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	if v148 != 0 {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	v134 = int32(4384884)
	v136 = *(*int32)(unsafe.Add(mBase, _consts[579]))
	*(*int32)(unsafe.Add(mBase, _consts[579])) = v136 + int32(1)
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
	v152 = *(*int32)(unsafe.Add(mBase, _consts[520]))
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
	v2 = *(*int32)(unsafe.Add(mBase, _consts[581]))
	if v2 != 0 {
		F_disable_timeout(m, int32(12))
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[580])) = int32(0)
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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
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
	if v18&int32(3) == int32(0) {
		v42 = v18
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	if v231 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L7:
	;
	v84 = F_palloc(m, (v75<<(uint(int32(1))%32)+int32(2))&int32(-4)+int32(8))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L24
	} else {
		goto L25
	}
L8:
	;
	v75 = v67 - v18
	goto L7
L9:
	;
	v46 = v42
	goto L18
L10:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v26 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v75 = int32(0)
	goto L7
L12:
	;
	goto L13
L13:
	;
	v31 = v18
	goto L14
L14:
	;
	v35 = v31 + int32(1)
	if v35&int32(3) == int32(0) {
		v42 = v35
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v67 = v35
	goto L8
L16:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v40 != 0 {
		v31 = v35
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v55 = int32(-2139062144)
	if (int32(16843008)-v52|v52)&v55 == v55 {
		v46 = v46 + int32(4)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v61 = v46
	goto L21
L20:
	;
	goto L19
L21:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v65 != 0 {
		v61 = v61 + int32(1)
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v67 = v61
	goto L8
L23:
	;
	goto L22
L24:
	;
	return
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = int32(160137)
	v91 = v13 + int32(12)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v93 = m.G0
	v95 = v93 - int32(16)
	m.G0 = v95
	F_initStringInfo(m, v95)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v99 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	F_pfree(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L24
	} else {
		goto L52
	}
L28:
	;
	v105 = v92
	goto L29
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v113)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v113
	goto L31
L31:
	;
	v123 = v105
	goto L32
L32:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if base.Ui32(v130-int32(9)) < base.Ui32(int32(5)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v123 = v123 + int32(1)
	goto L32
L35:
	;
	if v130 == int32(32) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	if v130 == int32(0) {
		goto L27
	} else {
		goto L37
	}
L37:
	;
	v141 = v130
	v142 = v123
	v143 = int32(0)
	goto L38
L38:
	;
	v150 = v141 & int32(255)
	if v143&int32(1) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v183 = F_pstrdup(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L24
	} else {
		goto L50
	}
L40:
	;
	goto L39
L41:
	;
	if base.B2i32(v150 != int32(32))&base.B2i32(base.Ui32((v141-int32(14))&int32(255)) < base.Ui32(int32(251))) == int32(0) {
		v180 = v142
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v168 = v143 | base.B2i32(v150 != int32(92))
	if v168&int32(1) != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	F_appendStringInfoChar(m, v95, base.I32_extend8_s(v141))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L24
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v174 = int32(1)
	v177 = v142 + v174
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if v178 != 0 {
		v141 = v178
		v142 = v177
		v143 = v168 ^ v174
		goto L38
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v180 = v177
	goto L40
L50:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v185 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v84+v185<<(uint(int32(2))%32)))) = v183
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v193 != 0 {
		v105 = v180
		goto L29
	} else {
		goto L51
	}
L51:
	;
	goto L27
L52:
	;
	m.G0 = v95 + int32(16)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v216 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v84+v212<<(uint(int32(2))%32)))) = v216
	F_process_postgres_switches(m, v212, v84, v17, v216)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L24
	} else {
		goto L53
	}
L53:
	;
	goto L6
L54:
	;
	m.G0 = v13 + int32(16)
	return
L55:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	if v234 == int32(0) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v238 = v234
	goto L57
L57:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v249 = v238 + int32(4)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	v256 = v252 + v253<<(uint(int32(2))%32)
	if base.Ui32(v249) < base.Ui32(v256) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L54
L59:
	;
	v258 = v249
	goto L61
L60:
	;
	v258 = int32(0)
	goto L61
L61:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	F_SetConfigOption(m, v247, v259, v17, int32(9))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L24
	} else {
		goto L62
	}
L62:
	;
	v264 = v258 + int32(4)
	if base.Ui32(v256) <= base.Ui32(v264) {
		goto L54
	} else {
		goto L63
	}
L63:
	;
	if v264 != 0 {
		v238 = v264
		goto L57
	} else {
		goto L64
	}
L64:
	;
	goto L58
}
func F_startup_progress_timeout_handler(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[580])) = int32(1)
	return
}
