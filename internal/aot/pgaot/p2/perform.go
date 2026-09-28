package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_performMultipleDeletions(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v4 < v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v10 + int32(16)
	return
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v17
	v21 = F_palloc(m, int32(16))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = int64(137438953472)
	v27 = F_palloc_mul(m, int32(12), int32(32))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v27
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v32 <= v29 {
		v68 = v4
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_reportDependentObjects(m, v21, l1, l2, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L16
	}
L9:
	;
	v41 = v4
	goto L10
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v45 = v42 + v41*int32(12)
	F_AcquireDeletionLock(m, v45, l2)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	if v56 != int32(1) {
		v68 = int32(0)
		goto L8
	} else {
		goto L15
	}
L12:
	;
	F_findDependentObjects(m, v45, int32(1), l2, int32(0), v21, l0, v10+int32(12))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v55 = v41 + int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v55 < v56 {
		v41 = v55
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v68 = v61
	goto L8
L16:
	;
	F_deleteObjectsInList(m, v21, v10+int32(12), l2)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	F_pfree(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v78 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_pfree(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_pfree(m, v21)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_relation_close(m, v83, int32(3))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	goto L3
}
func F_perform_pullup_replace_vars(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v15 = F_replace_rte_variables(m, v10, v11, v8, int32(899), l1, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v25 = F_replace_rte_variables(m, v20, v21, int32(0), int32(899), l1, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v7
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v25
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v33 = F_replace_rte_variables(m, v28, v29, int32(0), int32(899), l1, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v33
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v42 = F_replace_rte_variables(m, v37, v38, int32(0), int32(899), l1, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	if v57 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v42
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+28))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v52 = F_replace_rte_variables(m, v47, v48, int32(0), int32(899), l1, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+28)) = v52
	goto L10
L13:
	;
	v100 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v106 = F_replace_rte_variables(m, v101, v102, v100, int32(899), l1, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L21
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v60 <= int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v68 = int32(0)
	goto L16
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v68<<(uint(int32(2))%32))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v79 = F_replace_rte_variables(m, v74, v75, int32(0), int32(899), l1, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	goto L13
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = v79
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v87 = F_replace_rte_variables(m, v82, v83, int32(0), int32(899), l1, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = v87
	v91 = v68 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v91 < v92 {
		v68 = v91
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v106
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F_replace_vars_in_jointree(m, v109, l1)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v19)+112))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v117 = F_replace_rte_variables(m, v112, v113, int32(0), int32(899), l1, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v117
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v120 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	if v155 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v123 <= int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v128 = v100
	goto L27
L27:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v128<<(uint(int32(2))%32))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+20))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v142 = F_replace_rte_variables(m, v137, v138, int32(0), int32(899), l1, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L29
	}
L28:
	;
	goto L24
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+20)) = v142
	v146 = v128 + int32(1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v146 < v147 {
		v128 = v146
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	return
L32:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v158 <= int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v164 = int32(0)
	goto L34
L34:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	v169 = int32(2)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168+v164<<(uint(v169)%32))))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	switch v173 - v169 {
	case 0:
		goto L38
	default:
		goto L36
	case 7:
		goto L37
	}
L35:
	;
	goto L31
L36:
	;
	v193 = v164 + int32(1)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v193 < v194 {
		v164 = v193
		goto L34
	} else {
		goto L41
	}
L37:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v172)+120))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v189 = F_replace_rte_variables(m, v184, v185, int32(0), int32(899), l1, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172)+52))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v181 = F_replace_rte_variables(m, v176, v177, int32(0), int32(899), l1, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+52)) = v181
	goto L36
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+120)) = v189
	goto L36
L41:
	;
	goto L35
}
func F_perform_spin_delay(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
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
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v59 float64
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = v4 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_perform_spin_delay[0]))
	if v9 <= v6 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v13 = v11 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v13
		if int32(1001) <= v13 {
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			F_s_lock_stuck(m, v77, v78, v79)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v17 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1000)
			} else {
			}
			v22 = int32(_a_F_perform_spin_delay_0)
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_perform_spin_delay[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(150994951)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_pg_usleep(m, v26)
			mBase = m.M
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_perform_spin_delay[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(0)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v33 = int32(_a_F_perform_spin_delay_1)
			v36 = *(*int64)(unsafe.Add(mBase, _c_F_perform_spin_delay[2]))
			v37 = *(*int64)(unsafe.Add(mBase, _c_F_perform_spin_delay[3]))
			v38 = v36 ^ v37
			*(*int64)(unsafe.Add(mBase, _c_F_perform_spin_delay[3])) = base.I64_rotl(v38, int64(37))
			*(*int64)(unsafe.Add(mBase, _c_F_perform_spin_delay[2])) = v38<<(uint(int64(16))%64) ^ base.I64_rotl(v36, int64(24)) ^ v38
			v59 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v36*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v69 = v63 + base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v59, base.F64_convert_i32_s(v32)), float64(0.5)))
			if int32(_a_F_perform_spin_delay_2) < v69 {
				v72 = int32(1000)
			} else {
				v72 = v69
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v72
			return
		}
	} else {
		return
	}
}
