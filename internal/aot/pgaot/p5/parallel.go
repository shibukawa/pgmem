package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v57 int32
	_ = v57
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
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
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
	v39 = int32(_a_F_ExecParallelHashJoinSetUpBatches_0)
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoinSetUpBatches[0]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoinSetUpBatches[0])) = v42
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
	v57 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoinSetUpBatches[0])) = v40
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
	v82 = v37 + (((v68*int32(28)+int32(76))<<(uint(int32(1))%32)+int32(14))&int32(-16)-int32(-64))*v57
	v84 = v82 + int32(4)
	v85 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v84))), uint32(v85))
	*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v85
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+20)) = uint8(v85)
	*(*int64)(unsafe.Add(mBase, uint32(v84)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v85
	F_ConditionVariableInit(m, v82+int32(28))
	mBase = m.M
	goto L12
L12:
	;
	if v57 == int32(0) {
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
	v146 = v66 + v57*int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v82
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v57
	v152 = v13 + int32(32)
	v157 = F_pg_snprintf(m, v152, int32(1024), int32(_a_F_ExecParallelHashJoinSetUpBatches_1), v13+int32(16))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
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
	v160 = v82 - int32(-64)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoinSetUpBatches[1]))
	v166 = F_sts_initialize(m, v160, v161, v163+int32(1), v54, v152)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+28)) = v166
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v57
	v174 = F_pg_snprintf(m, v152, int32(1024), int32(_a_F_ExecParallelHashJoinSetUpBatches_2), v13)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	goto L28
L28:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoinSetUpBatches[1]))
	v191 = F_sts_initialize(m, v160+(v176*int32(28)+int32(83))&int32(-8), v186, v188+int32(1), v54, v152)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+32)) = v191
	v195 = v57 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v195 < v196 {
		v57 = v195
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v9 = v7 + int32(40)
	v11 = F_LWLockAcquire(m, v9, v2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v13 < v15 {
			v22 = v2
			for {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
				v27 = v24 + v22*int32(36)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v28)+44)) = v29 + v30
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v34 + v35
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v39 + v40
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v43)+56)) = v44 + v45
				v48 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v27)+16)) = v48
				*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = v48
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+52))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v52 + v54
				v58 = v22 + int32(1)
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v58 < v59 {
					v22 = v58
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		F_LWLockRelease(m, v9)
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
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
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
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
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v16 = (l1 + int32(7)) & int32(-8)
	if base.Ui32(int32(_a_F_ExecParallelHashTupleAlloc_0)) < base.Ui32(v16) {
		v41 = v12 + int32(40)
		v43 = F_LWLockAcquire(m, v41, int32(0))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
			v48 = int32(1)
			if base.Ui32(v47-v48) <= base.Ui32(v48) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
				F_LWLockRelease(m, v41)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					if v47 == int32(2) {
						F_ExecParallelHashIncreaseNumBatches(m, l0)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					} else {
						F_ExecParallelHashIncreaseNumBuckets(m, l0)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				}
			} else {
				v70 = base.B2i32(base.Ui32(int32(_a_F_ExecParallelHashTupleAlloc_0)) < base.Ui32(v16))
				if base.Ui32(int32(_a_F_ExecParallelHashTupleAlloc_0)) < base.Ui32(v16) {
					v71 = v16 + int32(16)
				} else {
					v71 = int32(_a_F_ExecParallelHashTupleAlloc_1)
				}
				if v47 == int32(3) {
					v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
					v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						v128 = v11 * int32(36)
						v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
						*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
						v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
						v137 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
						v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
						v140 = F_dsa_get_address(m, v139, v125)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return int32(0)
						} else {
							v142 = int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
							*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
							*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
							v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
							v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
							v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
							v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
							v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
							*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
							if v70 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
								*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
							} else {
							}
							F_LWLockRelease(m, v41)
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return int32(0)
							} else {
								return v140 + int32(16)
							}
						}
					}
				} else {
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+24)))
					if v75 != int32(1) {
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						if v89 != int32(1) {
							v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
							v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								v128 = v11 * int32(36)
								v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
								v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
								*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
								v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
								v137 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
								v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
								v140 = F_dsa_get_address(m, v139, v125)
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return int32(0)
								} else {
									v142 = int32(16)
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
									*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
									*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
									v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
									v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
									v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
									*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
									if v70 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
										*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
									} else {
									}
									F_LWLockRelease(m, v41)
									mBase = m.M
									v163 = m.ExcPending
									if v163 != 0 {
										return int32(0)
									} else {
										return v140 + int32(16)
									}
								}
							}
						} else {
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
							if base.B2i32(base.Ui32(v103+int32(1)) <= base.Ui32(v100))|base.B2i32(int32(1073741822) < v100)|v100&int32(2013265920) != 0 {
								v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
								v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									v128 = v11 * int32(36)
									v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
									v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
									*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
									v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									v137 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
									v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
									v140 = F_dsa_get_address(m, v139, v125)
									mBase = m.M
									v141 = m.ExcPending
									if v141 != 0 {
										return int32(0)
									} else {
										v142 = int32(16)
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
										*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
										*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
										v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
										v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
										*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
										v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
										*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
										if v70 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
											*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
										} else {
										}
										F_LWLockRelease(m, v41)
										mBase = m.M
										v163 = m.ExcPending
										if v163 != 0 {
											return int32(0)
										} else {
											return v140 + int32(16)
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(1)
								F_LWLockRelease(m, v41)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
						if base.Ui32(v80+v71) <= base.Ui32(v78) {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v89 != int32(1) {
								v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
								v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									v128 = v11 * int32(36)
									v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
									v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
									*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
									v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									v137 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
									v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
									v140 = F_dsa_get_address(m, v139, v125)
									mBase = m.M
									v141 = m.ExcPending
									if v141 != 0 {
										return int32(0)
									} else {
										v142 = int32(16)
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
										*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
										*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
										v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
										v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
										*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
										v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
										*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
										if v70 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
											*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
										} else {
										}
										F_LWLockRelease(m, v41)
										mBase = m.M
										v163 = m.ExcPending
										if v163 != 0 {
											return int32(0)
										} else {
											return v140 + int32(16)
										}
									}
								}
							} else {
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
								if base.B2i32(base.Ui32(v103+int32(1)) <= base.Ui32(v100))|base.B2i32(int32(1073741822) < v100)|v100&int32(2013265920) != 0 {
									v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
									v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return int32(0)
									} else {
										v128 = v11 * int32(36)
										v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
										v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
										*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
										v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v137 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
										v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
										v140 = F_dsa_get_address(m, v139, v125)
										mBase = m.M
										v141 = m.ExcPending
										if v141 != 0 {
											return int32(0)
										} else {
											v142 = int32(16)
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
											*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
											*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
											v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
											v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
											*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
											v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
											*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
											if v70 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
												*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
											} else {
											}
											F_LWLockRelease(m, v41)
											mBase = m.M
											v163 = m.ExcPending
											if v163 != 0 {
												return int32(0)
											} else {
												return v140 + int32(16)
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(1)
									F_LWLockRelease(m, v41)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										return int32(0)
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(2)
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
							v87 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v86)+60)) = uint8(v87)
							F_LWLockRelease(m, v41)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								return int32(0)
							}
						}
					}
				}
			}
		}
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		if v19 == int32(0) {
			v41 = v12 + int32(40)
			v43 = F_LWLockAcquire(m, v41, int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
				v48 = int32(1)
				if base.Ui32(v47-v48) <= base.Ui32(v48) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
					F_LWLockRelease(m, v41)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						if v47 == int32(2) {
							F_ExecParallelHashIncreaseNumBatches(m, l0)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								return int32(0)
							}
						} else {
							F_ExecParallelHashIncreaseNumBuckets(m, l0)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								return int32(0)
							}
						}
					}
				} else {
					v70 = base.B2i32(base.Ui32(int32(_a_F_ExecParallelHashTupleAlloc_0)) < base.Ui32(v16))
					if base.Ui32(int32(_a_F_ExecParallelHashTupleAlloc_0)) < base.Ui32(v16) {
						v71 = v16 + int32(16)
					} else {
						v71 = int32(_a_F_ExecParallelHashTupleAlloc_1)
					}
					if v47 == int32(3) {
						v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
						v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							v128 = v11 * int32(36)
							v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
							v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
							*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
							v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
							v137 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
							v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
							v140 = F_dsa_get_address(m, v139, v125)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return int32(0)
							} else {
								v142 = int32(16)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
								*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
								*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
								v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
								v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
								v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
								v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
								v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
								*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
								if v70 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
									*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
								} else {
								}
								F_LWLockRelease(m, v41)
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
									return int32(0)
								} else {
									return v140 + int32(16)
								}
							}
						}
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+24)))
						if v75 != int32(1) {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v89 != int32(1) {
								v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
								v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									v128 = v11 * int32(36)
									v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
									v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
									*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
									v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									v137 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
									v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
									v140 = F_dsa_get_address(m, v139, v125)
									mBase = m.M
									v141 = m.ExcPending
									if v141 != 0 {
										return int32(0)
									} else {
										v142 = int32(16)
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
										*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
										*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
										v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
										v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
										*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
										v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
										*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
										if v70 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
											*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
										} else {
										}
										F_LWLockRelease(m, v41)
										mBase = m.M
										v163 = m.ExcPending
										if v163 != 0 {
											return int32(0)
										} else {
											return v140 + int32(16)
										}
									}
								}
							} else {
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
								if base.B2i32(base.Ui32(v103+int32(1)) <= base.Ui32(v100))|base.B2i32(int32(1073741822) < v100)|v100&int32(2013265920) != 0 {
									v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
									v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return int32(0)
									} else {
										v128 = v11 * int32(36)
										v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
										v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
										*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
										v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v137 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
										v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
										v140 = F_dsa_get_address(m, v139, v125)
										mBase = m.M
										v141 = m.ExcPending
										if v141 != 0 {
											return int32(0)
										} else {
											v142 = int32(16)
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
											*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
											*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
											v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
											v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
											*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
											v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
											*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
											if v70 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
												*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
											} else {
											}
											F_LWLockRelease(m, v41)
											mBase = m.M
											v163 = m.ExcPending
											if v163 != 0 {
												return int32(0)
											} else {
												return v140 + int32(16)
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(1)
									F_LWLockRelease(m, v41)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										return int32(0)
									}
								}
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
							if base.Ui32(v80+v71) <= base.Ui32(v78) {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v89 != int32(1) {
									v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
									v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return int32(0)
									} else {
										v128 = v11 * int32(36)
										v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
										v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
										*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
										v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v137 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
										v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
										v140 = F_dsa_get_address(m, v139, v125)
										mBase = m.M
										v141 = m.ExcPending
										if v141 != 0 {
											return int32(0)
										} else {
											v142 = int32(16)
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
											*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
											*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
											v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
											v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
											*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
											v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
											*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
											if v70 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
												*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
											} else {
											}
											F_LWLockRelease(m, v41)
											mBase = m.M
											v163 = m.ExcPending
											if v163 != 0 {
												return int32(0)
											} else {
												return v140 + int32(16)
											}
										}
									}
								} else {
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
									if base.B2i32(base.Ui32(v103+int32(1)) <= base.Ui32(v100))|base.B2i32(int32(1073741822) < v100)|v100&int32(2013265920) != 0 {
										v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
										v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return int32(0)
										} else {
											v128 = v11 * int32(36)
											v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
											*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
											v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v137 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
											v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
											v140 = F_dsa_get_address(m, v139, v125)
											mBase = m.M
											v141 = m.ExcPending
											if v141 != 0 {
												return int32(0)
											} else {
												v142 = int32(16)
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
												*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
												*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
												v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
												v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
												*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
												v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
												v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
												*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
												if v70 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
													*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
												} else {
												}
												F_LWLockRelease(m, v41)
												mBase = m.M
												v163 = m.ExcPending
												if v163 != 0 {
													return int32(0)
												} else {
													return v140 + int32(16)
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(1)
										F_LWLockRelease(m, v41)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											return int32(0)
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(2)
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
								v87 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v86)+60)) = uint8(v87)
								F_LWLockRelease(m, v41)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									return int32(0)
								}
							}
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
			if base.Ui32(v22-v23) < base.Ui32(v16) {
				v41 = v12 + int32(40)
				v43 = F_LWLockAcquire(m, v41, int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
					v48 = int32(1)
					if base.Ui32(v47-v48) <= base.Ui32(v48) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
						F_LWLockRelease(m, v41)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							if v47 == int32(2) {
								F_ExecParallelHashIncreaseNumBatches(m, l0)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									return int32(0)
								}
							} else {
								F_ExecParallelHashIncreaseNumBuckets(m, l0)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v70 = base.B2i32(base.Ui32(int32(_a_F_ExecParallelHashTupleAlloc_0)) < base.Ui32(v16))
						if base.Ui32(int32(_a_F_ExecParallelHashTupleAlloc_0)) < base.Ui32(v16) {
							v71 = v16 + int32(16)
						} else {
							v71 = int32(_a_F_ExecParallelHashTupleAlloc_1)
						}
						if v47 == int32(3) {
							v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
							v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								v128 = v11 * int32(36)
								v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
								v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
								*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
								v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
								v137 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
								v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
								v140 = F_dsa_get_address(m, v139, v125)
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return int32(0)
								} else {
									v142 = int32(16)
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
									*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
									*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
									v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
									v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
									v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
									*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
									if v70 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
										*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
									} else {
									}
									F_LWLockRelease(m, v41)
									mBase = m.M
									v163 = m.ExcPending
									if v163 != 0 {
										return int32(0)
									} else {
										return v140 + int32(16)
									}
								}
							}
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+24)))
							if v75 != int32(1) {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v89 != int32(1) {
									v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
									v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return int32(0)
									} else {
										v128 = v11 * int32(36)
										v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
										v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
										*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
										v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v137 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
										v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
										v140 = F_dsa_get_address(m, v139, v125)
										mBase = m.M
										v141 = m.ExcPending
										if v141 != 0 {
											return int32(0)
										} else {
											v142 = int32(16)
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
											*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
											*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
											v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
											v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
											*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
											v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
											*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
											if v70 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
												*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
											} else {
											}
											F_LWLockRelease(m, v41)
											mBase = m.M
											v163 = m.ExcPending
											if v163 != 0 {
												return int32(0)
											} else {
												return v140 + int32(16)
											}
										}
									}
								} else {
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
									if base.B2i32(base.Ui32(v103+int32(1)) <= base.Ui32(v100))|base.B2i32(int32(1073741822) < v100)|v100&int32(2013265920) != 0 {
										v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
										v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return int32(0)
										} else {
											v128 = v11 * int32(36)
											v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
											*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
											v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v137 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
											v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
											v140 = F_dsa_get_address(m, v139, v125)
											mBase = m.M
											v141 = m.ExcPending
											if v141 != 0 {
												return int32(0)
											} else {
												v142 = int32(16)
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
												*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
												*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
												v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
												v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
												*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
												v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
												v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
												*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
												if v70 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
													*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
												} else {
												}
												F_LWLockRelease(m, v41)
												mBase = m.M
												v163 = m.ExcPending
												if v163 != 0 {
													return int32(0)
												} else {
													return v140 + int32(16)
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(1)
										F_LWLockRelease(m, v41)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											return int32(0)
										}
									}
								}
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
								if base.Ui32(v80+v71) <= base.Ui32(v78) {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v89 != int32(1) {
										v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
										v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return int32(0)
										} else {
											v128 = v11 * int32(36)
											v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
											*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
											v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v137 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
											v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
											v140 = F_dsa_get_address(m, v139, v125)
											mBase = m.M
											v141 = m.ExcPending
											if v141 != 0 {
												return int32(0)
											} else {
												v142 = int32(16)
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
												*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
												*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
												v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
												v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
												*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
												v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
												v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
												*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
												if v70 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
													*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
												} else {
												}
												F_LWLockRelease(m, v41)
												mBase = m.M
												v163 = m.ExcPending
												if v163 != 0 {
													return int32(0)
												} else {
													return v140 + int32(16)
												}
											}
										}
									} else {
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
										if base.B2i32(base.Ui32(v103+int32(1)) <= base.Ui32(v100))|base.B2i32(int32(1073741822) < v100)|v100&int32(2013265920) != 0 {
											v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
											v125 = F_dsa_allocate_extended(m, v123, v71, int32(0))
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
												return int32(0)
											} else {
												v128 = v11 * int32(36)
												v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
												v131 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129)))
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
												*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v132 + v71
												v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
												v137 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v135+v128)+24)) = uint8(v137)
												v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
												v140 = F_dsa_get_address(m, v139, v125)
												mBase = m.M
												v141 = m.ExcPending
												if v141 != 0 {
													return int32(0)
												} else {
													v142 = int32(16)
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125 + v142
													*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v16
													*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v71 - v142
													v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
													v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v128)))
													v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+40))
													*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v152
													v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
													v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v128)))
													*(*int32)(unsafe.Add(mBase, uint32(v156)+40)) = v125
													if v70 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v125
														*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v140
													} else {
													}
													F_LWLockRelease(m, v41)
													mBase = m.M
													v163 = m.ExcPending
													if v163 != 0 {
														return int32(0)
													} else {
														return v140 + int32(16)
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(1)
											F_LWLockRelease(m, v41)
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												return int32(0)
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(2)
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
									v87 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v86)+60)) = uint8(v87)
									F_LWLockRelease(m, v41)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										return int32(0)
									}
								}
							}
						}
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
				v28 = int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v23 + v26 + v28
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v31 + v16
				return v31 + v19 + v28
			}
		}
	}
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
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
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
	v23 = F_mul_size(m, int32(_a_F_ExecParallelSetupTupleQueues_0), v22)
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
	v46 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v44))), uint32(v46))
	v49 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+16)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v44)+4)) = v49
	v53 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+36)) = uint16(v53)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+24)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = int32(_a_F_ExecParallelSetupTupleQueues_1)
	goto L18
L17:
	;
	goto L15
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelSetupTupleQueues[0]))
	F_shm_mq_set_receiver(m, v44, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v70 = F_shm_mq_attach(m, v44, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+v38<<(uint(int32(2))%32)))) = v70
	v74 = v38 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v74 < v75 {
		v38 = v74
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v85, int64(-2305843009213693947), v31)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerShutdown[0]))
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerShutdown[1]))
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
