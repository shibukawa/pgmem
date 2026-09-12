package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InstrAggNode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 float64
	_ = v15
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v48 float64
	_ = v48
	var v49 float64
	_ = v49
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v6 == int32(0) {
		if v5&int32(1) == int32(0) {
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v13)
			v15 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
			*(*float64)(unsafe.Add(mBase, uint32(l0)+24)) = v15
		}
	} else {
		if v5&int32(1) == int32(0) {
		} else {
			v21 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
			if base.F64_lt(v21, v22) == int32(0) {
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(l0)+24)) = v21
			}
		}
	}
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v28 + v29
	v32 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v33 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = base.F64_add(v32, v33)
	v36 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v37 = *(*float64)(unsafe.Add(mBase, uint32(l0)+200))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+200)) = base.F64_add(v36, v37)
	v40 = *(*float64)(unsafe.Add(mBase, uint32(l1)+208))
	v41 = *(*float64)(unsafe.Add(mBase, uint32(l0)+208))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+208)) = base.F64_add(v40, v41)
	v44 = *(*float64)(unsafe.Add(mBase, uint32(l1)+216))
	v45 = *(*float64)(unsafe.Add(mBase, uint32(l0)+216))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+216)) = base.F64_add(v44, v45)
	v48 = *(*float64)(unsafe.Add(mBase, uint32(l1)+224))
	v49 = *(*float64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+224)) = base.F64_add(v48, v49)
	v52 = *(*float64)(unsafe.Add(mBase, uint32(l1)+232))
	v53 = *(*float64)(unsafe.Add(mBase, uint32(l0)+232))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+232)) = base.F64_add(v52, v53)
	v56 = *(*float64)(unsafe.Add(mBase, uint32(l1)+240))
	v57 = *(*float64)(unsafe.Add(mBase, uint32(l0)+240))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+240)) = base.F64_add(v56, v57)
	v60 = *(*float64)(unsafe.Add(mBase, uint32(l1)+248))
	v61 = *(*float64)(unsafe.Add(mBase, uint32(l0)+248))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+248)) = base.F64_add(v60, v61)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v64 == int32(1) {
		v67 = *(*int64)(unsafe.Add(mBase, uint32(l0)+256))
		v68 = *(*int64)(unsafe.Add(mBase, uint32(l1)+256))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+256)) = v67 + v68
		v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
		v72 = *(*int64)(unsafe.Add(mBase, uint32(l1)+264))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+264)) = v71 + v72
		v75 = *(*int64)(unsafe.Add(mBase, uint32(l0)+272))
		v76 = *(*int64)(unsafe.Add(mBase, uint32(l1)+272))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+272)) = v75 + v76
		v79 = *(*int64)(unsafe.Add(mBase, uint32(l0)+280))
		v80 = *(*int64)(unsafe.Add(mBase, uint32(l1)+280))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+280)) = v79 + v80
		v83 = *(*int64)(unsafe.Add(mBase, uint32(l0)+288))
		v84 = *(*int64)(unsafe.Add(mBase, uint32(l1)+288))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+288)) = v83 + v84
		v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+296))
		v88 = *(*int64)(unsafe.Add(mBase, uint32(l1)+296))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+296)) = v87 + v88
		v91 = *(*int64)(unsafe.Add(mBase, uint32(l0)+304))
		v92 = *(*int64)(unsafe.Add(mBase, uint32(l1)+304))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+304)) = v91 + v92
		v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)+312))
		v96 = *(*int64)(unsafe.Add(mBase, uint32(l1)+312))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = v95 + v96
		v99 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
		v100 = *(*int64)(unsafe.Add(mBase, uint32(l1)+320))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+320)) = v99 + v100
		v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
		v104 = *(*int64)(unsafe.Add(mBase, uint32(l1)+328))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v103 + v104
		v107 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
		v108 = *(*int64)(unsafe.Add(mBase, uint32(l1)+336))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = v107 + v108
		v111 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
		v112 = *(*int64)(unsafe.Add(mBase, uint32(l1)+344))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+344)) = v111 + v112
		v115 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
		v116 = *(*int64)(unsafe.Add(mBase, uint32(l1)+352))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+352)) = v115 + v116
		v119 = *(*int64)(unsafe.Add(mBase, uint32(l0)+360))
		v120 = *(*int64)(unsafe.Add(mBase, uint32(l1)+360))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+360)) = v119 + v120
		v123 = *(*int64)(unsafe.Add(mBase, uint32(l0)+368))
		v124 = *(*int64)(unsafe.Add(mBase, uint32(l1)+368))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+368)) = v123 + v124
		v127 = *(*int64)(unsafe.Add(mBase, uint32(l0)+376))
		v128 = *(*int64)(unsafe.Add(mBase, uint32(l1)+376))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+376)) = v127 + v128
	} else {
	}
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if v131 == int32(1) {
		v134 = *(*int64)(unsafe.Add(mBase, uint32(l0)+400))
		v135 = *(*int64)(unsafe.Add(mBase, uint32(l1)+400))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+400)) = v134 + v135
		v138 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
		v139 = *(*int64)(unsafe.Add(mBase, uint32(l1)+384))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+384)) = v138 + v139
		v142 = *(*int64)(unsafe.Add(mBase, uint32(l0)+392))
		v143 = *(*int64)(unsafe.Add(mBase, uint32(l1)+392))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = v142 + v143
		v146 = *(*int64)(unsafe.Add(mBase, uint32(l0)+408))
		v147 = *(*int64)(unsafe.Add(mBase, uint32(l1)+408))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+408)) = v146 + v147
	} else {
	}
	return
}
func F_InstrStartParallelQuery(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int64
	_ = v8
	var v12 int64
	_ = v12
	var v16 int64
	_ = v16
	var v20 int64
	_ = v20
	v4 = F__emscripten_memcpy_bulkmem(m, int32(4452664), int32(4452504), int32(128))
	mBase = m.M
	v8 = *(*int64)(unsafe.Add(mBase, _consts[74]))
	*(*int64)(unsafe.Add(mBase, _consts[535])) = v8
	v12 = *(*int64)(unsafe.Add(mBase, _consts[71]))
	*(*int64)(unsafe.Add(mBase, _consts[536])) = v12
	v16 = *(*int64)(unsafe.Add(mBase, _consts[73]))
	*(*int64)(unsafe.Add(mBase, _consts[537])) = v16
	v20 = *(*int64)(unsafe.Add(mBase, _consts[72]))
	*(*int64)(unsafe.Add(mBase, _consts[538])) = v20
	return
}
