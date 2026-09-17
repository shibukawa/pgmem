package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InstrEndParallelQuery(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v146 int64
	_ = v146
	var v148 int64
	_ = v148
	base.MemoryFill(m, l0, int32(0), int32(128))
	v6 = int32(_a_F_InstrEndParallelQuery_0)
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[0]))
	v10 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v7 + (v9 - v10)
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[2]))
	v17 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[3]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v14 + (v16 - v17)
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v23 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[4]))
	v24 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[5]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v21 + (v23 - v24)
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v30 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[6]))
	v31 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[7]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v28 + (v30 - v31)
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v37 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[8]))
	v38 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[9]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v35 + (v37 - v38)
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v44 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[10]))
	v45 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[11]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v42 + (v44 - v45)
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	v51 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[12]))
	v52 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[13]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v49 + (v51 - v52)
	v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v58 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[14]))
	v59 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[15]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v56 + (v58 - v59)
	v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v65 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[16]))
	v66 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[17]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v63 + (v65 - v66)
	v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v72 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[18]))
	v73 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[19]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v70 + (v72 - v73)
	v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	v79 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[20]))
	v80 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[21]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v77 + (v79 - v80)
	v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v86 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[22]))
	v87 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[23]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v84 + (v86 - v87)
	v91 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	v93 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[24]))
	v94 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[25]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = v91 + (v93 - v94)
	v98 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v100 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[26]))
	v101 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[27]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = v98 + (v100 - v101)
	v105 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v107 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[28]))
	v108 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[29]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v105 + (v107 - v108)
	v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v114 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[30]))
	v115 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[31]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v112 + (v114 - v115)
	v119 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v119
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v119
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v119
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v119
	v128 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[32]))
	v130 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[33]))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v128 - v130
	v134 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[34]))
	v136 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[35]))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v134 - v136
	v140 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[36]))
	v142 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[37]))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v140 - v142
	v146 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[38]))
	v148 = *(*int64)(unsafe.Add(mBase, _c_F_InstrEndParallelQuery[39]))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v146 - v148
	return
}
func F_InstrStartNode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v7 == int32(1) {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		if v10 != int64(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_InstrStartNode_0), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_InstrStartNode_1), int32(72), int32(_a_F_InstrStartNode_2))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			F___clock_gettime(m, int32(1), v5)
			mBase = m.M
			v15 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5)+8)))
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v15 + v16*int64(1000000000)
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if v21 == int32(1) {
				base.MemoryCopy(m, l0+int32(40), int32(_a_F_InstrStartNode_3), int32(128))
			} else {
			}
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			if v29 == int32(1) {
				v33 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStartNode[0]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v33
				v36 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStartNode[1]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v36
				v39 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStartNode[2]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v39
				v42 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStartNode[3]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v42
			} else {
			}
			m.G0 = v5 + int32(16)
			return
		}
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v21 == int32(1) {
			base.MemoryCopy(m, l0+int32(40), int32(_a_F_InstrStartNode_3), int32(128))
		} else {
		}
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
		if v29 == int32(1) {
			v33 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStartNode[0]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v33
			v36 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStartNode[1]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v36
			v39 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStartNode[2]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v39
			v42 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStartNode[3]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v42
		} else {
		}
		m.G0 = v5 + int32(16)
		return
	}
}
