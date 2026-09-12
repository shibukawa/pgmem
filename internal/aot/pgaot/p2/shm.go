package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F___shm_mapname(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	v5 = l0
	goto L1
L1:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v11 == int32(47) {
		v5 = v5 + int32(1)
		goto L1
	} else {
		goto L3
	}
L2:
	;
	goto L11
L3:
	;
	goto L2
L4:
	;
	goto L38
L5:
	;
	if base.Ui32(v113) < base.Ui32(int32(256)) {
		goto L4
	} else {
		goto L36
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[159])) = int32(28)
	return int32(0)
L7:
	;
	if v100 == v5 {
		goto L6
	} else {
		goto L31
	}
L8:
	;
	goto L7
L9:
	;
	v90 = v85
	goto L27
L10:
	;
	v85 = v77
	goto L9
L11:
	;
	if v5&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v24 = v5
	goto L17
L15:
	;
	v37 = v5
	goto L16
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v46 = int32(-2139062144)
	if (int32(16843008)-v43|v43)&v46 != v46 {
		v77 = v37
		goto L10
	} else {
		goto L22
	}
L17:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v29 == int32(0) {
		v100 = v24
		goto L8
	} else {
		goto L19
	}
L18:
	;
	v37 = v34
	goto L16
L19:
	;
	if int32(47) == v29 {
		v100 = v24
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v34 = v24 + int32(1)
	if v34&int32(3) != 0 {
		v24 = v34
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v52 = v37
	v54 = v43
	goto L23
L23:
	;
	v58 = v54 ^ int32(791621423)
	v61 = int32(-2139062144)
	if (int32(16843008)-v58|v58)&v61 != v61 {
		v77 = v52
		goto L10
	} else {
		goto L25
	}
L24:
	;
	v85 = v67
	goto L9
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v67 = v52 + int32(4)
	v71 = int32(-2139062144)
	if (v65|(int32(16843008)-v65))&v71 == v71 {
		v52 = v67
		v54 = v65
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v92 == int32(0) {
		v100 = v90
		goto L8
	} else {
		goto L29
	}
L28:
	;
	v100 = v90
	goto L8
L29:
	;
	if v92 != int32(47) {
		v90 = v90 + int32(1)
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v112 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v113 = v100 - v5
	if int32(2) < v113 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v116 != int32(46) {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100-int32(1)))))
	if v121 != int32(46) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	goto L6
L36:
	;
	*(*int32)(unsafe.Add(mBase, _consts[159])) = int32(37)
	return int32(0)
L37:
	;
	v144 = v113 + int32(1)
	if v144 != 0 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v139 = F__emscripten_memcpy_bulkmem(m, l1, int32(570791), int32(9))
	mBase = m.M
	goto L40
L40:
	;
	goto L37
L41:
	;
	return l1
L42:
	;
	v145 = F__emscripten_memcpy_bulkmem(m, l1+int32(9), v5, v144)
	mBase = m.M
	goto L44
L43:
	;
	goto L44
L44:
	;
	goto L41
}
func F_shm_mq_create(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	v3 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v3
	v9 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)) = uint16(v9)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1&int32(-8) - int32(40)
	return l0
}
func F_shm_mq_send_bytes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	v6 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	if l1 == v6 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v180
	m.G0 = v21 + int32(16)
	return v181
L2:
	;
	v180 = v161
	v181 = int32(0)
	goto L1
L3:
	;
	v161 = v6
	goto L2
L4:
	;
	goto L5
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	v31 = base.I64_extend_i32_u(v30)
	v32 = int32(2)
	v40 = v6
	goto L6
L6:
	;
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v51
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v53
	v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+36)))
	if v56 != 0 {
		v180 = v40
		v181 = v32
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v161 = v151
	goto L2
L8:
	;
	v57 = v53 + v55
	v59 = v51 - v57 + v31
	v61 = l1 - v40
	if base.Ui64(v59) < base.Ui64(base.I64_extend_i32_u(v61)) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if base.Ui32(v151) < base.Ui32(l1) {
		v40 = v151
		goto L6
	} else {
		goto L50
	}
L10:
	;
	v64 = base.I32_wrap_i64(v59)
	goto L12
L11:
	;
	v64 = v61
	goto L12
L12:
	;
	if v64 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v67 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v131 = base.I64_rem_u_s(v57, v31)
	v132 = base.I32_wrap_i64(v131)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+37)))
	v137 = v30 - v132
	if base.Ui32(v64) < base.Ui32(v137) {
		goto L43
	} else {
		goto L44
	}
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l3 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
	v103 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v102 + v103
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	F_SetLatch(m, v106+int32(20))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L25
	} else {
		goto L35
	}
L19:
	;
	v100 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)) = uint8(v100)
	v151 = v40
	goto L9
L20:
	;
	if v70 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v95 = F_shm_mq_wait_internal(m, v25, v25+int32(4), v70)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L25
	} else {
		goto L33
	}
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(1)
	if v83 != 0 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v75 = F_GetBackgroundWorkerPid(m, v70, v21+int32(12))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	if base.Ui32(v75) < base.Ui32(int32(2)) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+36)) = uint8(v81)
	v180 = v40
	v181 = v32
	goto L1
L28:
	;
	F_s_lock(m, v25, int32(495749), int32(246), int32(214997))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L25
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v93 != 0 {
		goto L19
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v180 = v40
	v181 = int32(1)
	goto L1
L33:
	;
	if v95 != 0 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	v97 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+36)) = uint8(v97)
	v180 = v40
	v181 = v32
	goto L1
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	if l3 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v180 = v40
	v181 = int32(1)
	goto L1
L37:
	;
	goto L38
L38:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v119 = F_WaitLatch(m, v115, int32(33), int32(0), int32(134217764))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L25
	} else {
		goto L39
	}
L39:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = int32(0)
	goto L40
L40:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v126 == int32(0) {
		v151 = v40
		goto L9
	} else {
		goto L41
	}
L41:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L25
	} else {
		goto L42
	}
L42:
	;
	v151 = v40
	goto L9
L43:
	;
	v139 = v64
	goto L45
L44:
	;
	v139 = v137
	goto L45
L45:
	;
	if v139 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v142 + (v139+int32(7))&int32(-8)
	v151 = v139 + v40
	goto L9
L47:
	;
	v140 = F__emscripten_memcpy_bulkmem(m, v25+int32(38)+(v132+v133), l2+v40, v139)
	mBase = m.M
	goto L49
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	goto L7
}
func F_shm_toc_create(m *base.Module, l0 int64, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l2 & int32(-32)
	return l1
}
func F_shm_toc_estimate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_mul_size(m, v3, int32(16))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_add_size(m, int32(24), v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v12 = F_add_size(m, v9, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return (v12 + int32(31)) & int32(-32)
			}
		}
	}
}
