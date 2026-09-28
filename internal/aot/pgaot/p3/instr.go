package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InstrAggNode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+392))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l1)+392))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = v3 + v4
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l1)+184))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v7 + v8
	v11 = *(*float64)(unsafe.Add(mBase, uint32(l1)+400))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(l0)+400))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+400)) = base.F64_add(v11, v12)
	v15 = *(*float64)(unsafe.Add(mBase, uint32(l1)+408))
	v16 = *(*float64)(unsafe.Add(mBase, uint32(l0)+408))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+408)) = base.F64_add(v15, v16)
	v19 = *(*float64)(unsafe.Add(mBase, uint32(l1)+416))
	v20 = *(*float64)(unsafe.Add(mBase, uint32(l0)+416))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+416)) = base.F64_add(v19, v20)
	v23 = *(*float64)(unsafe.Add(mBase, uint32(l1)+424))
	v24 = *(*float64)(unsafe.Add(mBase, uint32(l0)+424))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+424)) = base.F64_add(v23, v24)
	v27 = *(*float64)(unsafe.Add(mBase, uint32(l1)+432))
	v28 = *(*float64)(unsafe.Add(mBase, uint32(l0)+432))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+432)) = base.F64_add(v27, v28)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v31 == int32(1) {
		v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
		v35 = *(*int64)(unsafe.Add(mBase, uint32(l1)+192))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v34 + v35
		v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+200))
		v39 = *(*int64)(unsafe.Add(mBase, uint32(l1)+200))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = v38 + v39
		v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
		v43 = *(*int64)(unsafe.Add(mBase, uint32(l1)+208))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v42 + v43
		v46 = *(*int64)(unsafe.Add(mBase, uint32(l0)+216))
		v47 = *(*int64)(unsafe.Add(mBase, uint32(l1)+216))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+216)) = v46 + v47
		v50 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
		v51 = *(*int64)(unsafe.Add(mBase, uint32(l1)+224))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v50 + v51
		v54 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
		v55 = *(*int64)(unsafe.Add(mBase, uint32(l1)+232))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+232)) = v54 + v55
		v58 = *(*int64)(unsafe.Add(mBase, uint32(l0)+240))
		v59 = *(*int64)(unsafe.Add(mBase, uint32(l1)+240))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+240)) = v58 + v59
		v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
		v63 = *(*int64)(unsafe.Add(mBase, uint32(l1)+248))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v62 + v63
		v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+256))
		v67 = *(*int64)(unsafe.Add(mBase, uint32(l1)+256))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+256)) = v66 + v67
		v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
		v71 = *(*int64)(unsafe.Add(mBase, uint32(l1)+264))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+264)) = v70 + v71
		v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+272))
		v75 = *(*int64)(unsafe.Add(mBase, uint32(l1)+272))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+272)) = v74 + v75
		v78 = *(*int64)(unsafe.Add(mBase, uint32(l0)+280))
		v79 = *(*int64)(unsafe.Add(mBase, uint32(l1)+280))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+280)) = v78 + v79
		v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+288))
		v83 = *(*int64)(unsafe.Add(mBase, uint32(l1)+288))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+288)) = v82 + v83
		v86 = *(*int64)(unsafe.Add(mBase, uint32(l0)+296))
		v87 = *(*int64)(unsafe.Add(mBase, uint32(l1)+296))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+296)) = v86 + v87
		v90 = *(*int64)(unsafe.Add(mBase, uint32(l0)+304))
		v91 = *(*int64)(unsafe.Add(mBase, uint32(l1)+304))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+304)) = v90 + v91
		v94 = *(*int64)(unsafe.Add(mBase, uint32(l0)+312))
		v95 = *(*int64)(unsafe.Add(mBase, uint32(l1)+312))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = v94 + v95
	} else {
	}
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if v98 == int32(1) {
		v101 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
		v102 = *(*int64)(unsafe.Add(mBase, uint32(l1)+336))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = v101 + v102
		v105 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
		v106 = *(*int64)(unsafe.Add(mBase, uint32(l1)+320))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+320)) = v105 + v106
		v109 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
		v110 = *(*int64)(unsafe.Add(mBase, uint32(l1)+328))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v109 + v110
		v113 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
		v114 = *(*int64)(unsafe.Add(mBase, uint32(l1)+344))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+344)) = v113 + v114
		v117 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
		v118 = *(*int64)(unsafe.Add(mBase, uint32(l1)+352))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+352)) = v117 + v118
	} else {
	}
	return
}
func F_InstrStartParallelQuery(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v11 int64
	_ = v11
	var v15 int64
	_ = v15
	var v19 int64
	_ = v19
	var v23 int64
	_ = v23
	base.MemoryCopy(m, int32(_a_F_InstrStartParallelQuery_0), int32(_a_F_InstrStartParallelQuery_1), int32(128))
	v7 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStartParallelQuery[0]))
	*(*int64)(unsafe.Add(mBase, _c_F_InstrStartParallelQuery[1])) = v7
	v11 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStartParallelQuery[2]))
	*(*int64)(unsafe.Add(mBase, _c_F_InstrStartParallelQuery[3])) = v11
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStartParallelQuery[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_InstrStartParallelQuery[5])) = v15
	v19 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStartParallelQuery[6]))
	*(*int64)(unsafe.Add(mBase, _c_F_InstrStartParallelQuery[7])) = v19
	v23 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStartParallelQuery[8]))
	*(*int64)(unsafe.Add(mBase, _c_F_InstrStartParallelQuery[9])) = v23
	return
}
