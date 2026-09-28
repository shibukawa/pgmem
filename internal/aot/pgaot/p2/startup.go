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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	v1 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[0]))
	if v9 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[1]))
	if v128 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[0])) = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[2]))
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
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[3]))
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
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[4])))
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
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[2]))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if base.B2i32(v32 == int32(0))|base.B2i32(v32 != v35) != 0 {
		v53 = v32
		v54 = v35
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[3]))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if base.B2i32(v60 == int32(0))|base.B2i32(v60 != v63) != 0 {
		v81 = v60
		v82 = v63
		goto L16
	} else {
		goto L17
	}
L8:
	;
	goto L7
L9:
	;
	v38 = v17
	v39 = v29
	goto L10
L10:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v43 == int32(0) {
		v53 = v43
		v54 = v42
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v53 = v43
	v54 = v42
	goto L8
L12:
	;
	v46 = int32(1)
	if v43 == v42 {
		v38 = v38 + v46
		v39 = v39 + v46
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	F_pfree(m, v17)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L24
	}
L15:
	;
	if v83 != 0 {
		v88 = v1
		goto L14
	} else {
		goto L22
	}
L16:
	;
	v83 = v81 - v82
	goto L15
L17:
	;
	v66 = v21
	v67 = v57
	goto L18
L18:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v71 == int32(0) {
		v81 = v71
		v82 = v70
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v81 = v71
	v82 = v70
	goto L16
L20:
	;
	v74 = int32(1)
	if v71 == v70 {
		v66 = v66 + v74
		v67 = v67 + v74
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v84 != 0 {
		v88 = v1
		goto L14
	} else {
		goto L23
	}
L23:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[4])))
	v88 = base.B2i32(v24 != v86)
	goto L14
L24:
	;
	F_pfree(m, v21)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	if v53-v54|v83|v88 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[5]))
	if v98 != int32(3) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	goto L1
L28:
	;
	v101 = F_WalRcvRunning(m)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	if v101 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v107 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	if v107 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupProcInterrupts_0), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L3
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v119 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[6])) = uint8(v119)
	goto L27
L35:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupProcInterrupts_1), int32(_a_F_ProcessStartupProcInterrupts_2), int32(_a_F_ProcessStartupProcInterrupts_3))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	F_pgl_exit(m, int32(1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L3
	} else {
		goto L55
	}
L38:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[7])))
	if v132 != int32(1) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L54
	}
L41:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[8]))
	if v149 != 0 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	v135 = int32(_a_F_ProcessStartupProcInterrupts_4)
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[9])) = v137 + int32(1)
	if v137&int32(1023) != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v143 = F_PostmasterIsAliveInternal(m)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	if v143 == int32(0) {
		goto L37
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupProcInterrupts[10]))
	if v153 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L48
L50:
	;
	F_ProcessLogMemoryContextInterrupt(m)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	return
L53:
	;
	goto L52
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
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
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_disable_startup_progress_timeout[0]))
	if v2 != 0 {
		F_disable_timeout(m, int32(12))
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_disable_startup_progress_timeout[1])) = int32(0)
			return
		}
	} else {
		return
	}
}
func F_process_startup_options(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = int32(3)
	goto L3
L2:
	;
	v16 = int32(4)
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = F_strlen(m, v17)
	mBase = m.M
	v20 = int32(1)
	v26 = F_palloc_mul(m, int32(4), int32(base.Ui32(v19+v20)>>(uint(v20)%32))+int32(2))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	if v164 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(_a_F_process_startup_options_0)
	v33 = v12 + int32(12)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v35 = m.G0
	v37 = v35 - int32(16)
	m.G0 = v37
	F_initStringInfo(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v41 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	F_pfree(m, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L33
	}
L11:
	;
	v45 = v34
	goto L12
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v54
	goto L14
L13:
	;
	goto L10
L14:
	;
	v62 = v45
	goto L15
L15:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if base.B2i32(base.Ui32(v70-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v70 == int32(32)) != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L13
L17:
	;
	v62 = v62 + int32(1)
	goto L15
L18:
	;
	if v70 == int32(0) {
		goto L10
	} else {
		goto L20
	}
L19:
	;
	goto L16
L20:
	;
	v83 = v62
	v84 = v70
	v86 = int32(0)
	goto L21
L21:
	;
	if v86&int32(1)|base.B2i32(v84 != int32(32))&base.B2i32(base.Ui32((v84-int32(14))&int32(255)) < base.Ui32(int32(251))) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v120 = F_pstrdup(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L7
	} else {
		goto L31
	}
L23:
	;
	v105 = v86 | base.B2i32(v84 != int32(92))
	if v105&int32(1) != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v116 = v83
	goto L25
L25:
	;
	goto L22
L26:
	;
	F_appendStringInfoChar(m, v37, base.I32_extend8_s(v84))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v111 = int32(1)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	v115 = v83 + v111
	if v113 != 0 {
		v83 = v115
		v84 = v113
		v86 = v105 ^ v111
		goto L21
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v116 = v115
	goto L25
L31:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v122 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26+v122<<(uint(int32(2))%32)))) = v120
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v130 != 0 {
		v45 = v116
		goto L12
	} else {
		goto L32
	}
L32:
	;
	goto L19
L33:
	;
	m.G0 = v37 + int32(16)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v150 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26+v146<<(uint(int32(2))%32)))) = v150
	F_process_postgres_switches(m, v146, v26, v16, v150)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	goto L6
L35:
	;
	m.G0 = v12 + int32(16)
	return
L36:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	if v167 == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v171 = v167
	goto L38
L38:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	F_SetConfigOption(m, v182, v183, v16, int32(9))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L7
	} else {
		goto L40
	}
L39:
	;
	goto L35
L40:
	;
	v188 = v171 + int32(4)
	v192 = v180 + v181<<(uint(int32(2))%32)
	if base.Ui32(v188) < base.Ui32(v192) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v194 = v188
	goto L43
L42:
	;
	v194 = int32(0)
	goto L43
L43:
	;
	v196 = v194 + int32(4)
	if base.Ui32(v196) < base.Ui32(v192) {
		v171 = v196
		goto L38
	} else {
		goto L44
	}
L44:
	;
	goto L39
}
func F_startup_progress_timeout_handler(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_startup_progress_timeout_handler[0])) = int32(1)
	return
}
