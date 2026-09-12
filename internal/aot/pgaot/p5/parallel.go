package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecParallelHashJoinSetUpBatches(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v133 int32
	_ = v133
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	v11 = m.G0
	v13 = v11 - int32(1056)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	goto L1
L1:
	;
	v32 = F_dsa_allocate_extended(m, v15, (((v17*int32(28)+int32(76))<<(uint(int32(1))%32)+int32(14))&int32(-16)-int32(-64))*l1, int32(4))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v32
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v37 = F_dsa_get_address(m, v36, v32)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v39 = int32(4455216)
	v40 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v42
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l1
	v47 = F_palloc0(m, l1*int32(36))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v47
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if int32(0) < v50 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v54 = v16 + int32(168)
	v59 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v40
	m.G0 = v13 + int32(1056)
	return
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+28))
	goto L11
L10:
	;
	goto L8
L11:
	;
	v82 = v37 + (((v68*int32(28)+int32(76))<<(uint(int32(1))%32)+int32(14))&int32(-16)-int32(-64))*v59
	v84 = v82 + int32(4)
	v85 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v85
	v87 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v84))) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v84)+12)) = v87
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+20)) = uint8(v85)
	F_ConditionVariableInit(m, v82+int32(28))
	mBase = m.M
	goto L12
L12:
	;
	v98 = v66 + v59*int32(36)
	if v59 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v101 = F_BarrierAttach(m, v84)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v82
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v59
	v154 = F_pg_snprintf(m, v13+int32(32), int32(1024), int32(450012), v13+int32(16))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L25
	}
L16:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v103 <= int32(2) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	goto L20
L18:
	;
	goto L19
L19:
	;
	F_BarrierDetach(m, v84)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L24
	}
L20:
	;
	v117 = F_BarrierArriveAndWait(m, v84, int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L22
	}
L21:
	;
	goto L19
L22:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v119 < int32(3) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L15
L25:
	;
	v157 = v82 - int32(-64)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v160 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	v165 = F_sts_initialize(m, v157, v158, v160+int32(1), v54, v13+int32(32))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+28)) = v165
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v59
	v175 = F_pg_snprintf(m, v13+int32(32), int32(1024), int32(450004), v13)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	goto L28
L28:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v189 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	v194 = F_sts_initialize(m, v157+(v177*int32(28)+int32(83))&int32(-8), v187, v189+int32(1), v54, v13+int32(32))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+32)) = v194
	v198 = v59 + int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v198 < v199 {
		v59 = v198
		goto L9
	} else {
		goto L30
	}
L30:
	;
	goto L10
}
func F_ExecParallelHashMergeCounters(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v10 = v8 + int32(40)
	v12 = F_LWLockAcquire(m, v10, v2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v14
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v14 < v16 {
			v22 = v2
			for {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
				v29 = v26 + v22*int32(36)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+44))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+44)) = v31 + v32
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
				v38 = v29 + int32(16)
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
				*(*int32)(unsafe.Add(mBase, uint32(v35)+48)) = v36 + v39
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v42)+52)) = v43 + v44
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v48 + v49
				v52 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v38))) = v52
				*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v52
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+52))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v56 + v58
				v62 = v22 + int32(1)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v62 < v63 {
					v22 = v62
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		F_LWLockRelease(m, v10)
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ExecParallelHashTableSetCurrentBatch(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = l1
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v7 = l1 * int32(36)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7+v8)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = F_dsa_get_address(m, v5, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v12
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v16
		v19 = int32(1073741823)
		if v19 <= v16 {
			v22 = v19
		} else {
			v22 = v16
		}
		if base.Ui32(int32(2)) <= base.Ui32(v22) {
			v30 = int32(32) - base.I32_clz(v22-int32(1))
		} else {
			v30 = int32(0)
		}
		v31 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v31
		*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v31
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v30
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
		*(*uint8)(unsafe.Add(mBase, uint32(v36+v7)+24)) = uint8(v31)
		return
	}
}
func F_ExecParallelHashTupleAlloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v16 = (l1 + int32(7)) & int32(-8)
	if base.Ui32(int32(8192)) < base.Ui32(v16) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v41 = v12 + int32(40)
	v43 = F_LWLockAcquire(m, v41, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v19 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if base.Ui32(v22-v23) < base.Ui32(v16) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v28 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v23 + v26 + v28
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v31 + v16
	return v31 + v19 + v28
L5:
	;
	return int32(0)
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v48 = int32(1)
	if base.Ui32(v47-v48) <= base.Ui32(v48) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
	F_LWLockRelease(m, v41)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v70 = base.B2i32(base.Ui32(int32(8192)) < base.Ui32(v16))
	if base.Ui32(int32(8192)) < base.Ui32(v16) {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	if v47 == int32(2) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_ExecParallelHashIncreaseNumBatches(m, l0)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_ExecParallelHashIncreaseNumBuckets(m, l0)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	return int32(0)
L16:
	;
	v71 = v16 + int32(16)
	goto L18
L17:
	;
	v71 = int32(32768)
	goto L18
L18:
	;
	if v47 == int32(3) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v123 = F_dsa_allocate_extended(m, v121, v71, int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L5
	} else {
		goto L30
	}
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+24)))
	if v75 != int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	F_LWLockRelease(m, v41)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L29
	}
L22:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v89 != int32(1) {
		goto L19
	} else {
		goto L25
	}
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
	if base.Ui32(v80+v71) <= base.Ui32(v78) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(2)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v87 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+60)) = uint8(v87)
	goto L21
L25:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+52))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+52)) = v93 + v94
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+52))
	if base.Ui32(v103+int32(1)) <= base.Ui32(v100) {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	if int32(1073741822) < v100 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	if v100&int32(2013265920) != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(1)
	goto L21
L29:
	;
	return int32(0)
L30:
	;
	v126 = v11 * int32(36)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126+v127)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+44)) = v130 + v71
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v135 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v133+v126)+24)) = uint8(v135)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v138 = F_dsa_get_address(m, v137, v123)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v140 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v123 + v140
	*(*int32)(unsafe.Add(mBase, uint32(v138)+8)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v71 - v140
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v147+v126)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+12)) = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v152+v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+40)) = v123
	if v70 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v138
	goto L34
L33:
	;
	goto L34
L34:
	;
	F_LWLockRelease(m, v41)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	return v138 + int32(16)
}
func F_ExecParallelSetupTupleQueues(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int64
	_ = v46
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v14 = F_palloc(m, v7<<(uint(int32(2))%32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if l1 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v32 {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = F_mul_size(m, int32(65536), v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v29 = F_shm_toc_lookup(m, v18, int64(-2305843009213693947), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	v25 = F_shm_toc_allocate(m, v18, v23)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v31 = v25
	goto L6
L12:
	;
	v31 = v29
	goto L6
L13:
	;
	v38 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	if l1 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v44 = v31 + v38<<(uint(int32(16))%32)
	v46 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+16)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v44))) = v46
	v52 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+36)) = uint16(v52)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+24)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = int32(65496)
	goto L18
L17:
	;
	goto L15
L18:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[295]))
	F_shm_mq_set_receiver(m, v44, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v69 = F_shm_mq_attach(m, v44, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+v38<<(uint(int32(2))%32)))) = v69
	v73 = v38 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v73 < v74 {
		v38 = v73
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v84, int64(-2305843009213693947), v31)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	return v14
L25:
	;
	goto L24
}
func F_ParallelWorkerShutdown(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	v7 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	v8 = F_SendProcSignal(m, v4, int32(2), v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		F_dsm_detach(m, l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
