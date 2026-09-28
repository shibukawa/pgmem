package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DSMRegistryShmemRequest(m *base.Module, l0 int32) {
	var v6 int32
	_ = v6
	Fn14222(m, l0, int32(_a_F_DSMRegistryShmemRequest_0), int64(8), int32(_a_F_DSMRegistryShmemRequest_1))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_dsm_attach(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsm_attach[0])))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_dsm_backend_startup(m)
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
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[1]))
	v21 = int32(0)
	if base.B2i32(v20 == v21)|base.B2i32(v20 == int32(_a_F_dsm_attach_0)) == v21 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L45
	}
L7:
	;
	v30 = v20
	goto L10
L8:
	;
	goto L9
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[2]))
	if v54 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v38 == l0 {
		goto L6
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v40 != int32(_a_F_dsm_attach_0) {
		v30 = v40
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	F_ResourceOwnerEnlarge(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[3]))
	v60 = F_MemoryContextAlloc(m, v58, int32(36))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[1]))
	if v63 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v66 = int32(_a_F_dsm_attach_0)
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[4])) = v66
	v70 = v66
	goto L21
L20:
	;
	v70 = v63
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(_a_F_dsm_attach_0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v60
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[1])) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v60)+16)) = int64(4294967295)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v82
	if v82 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_ResourceOwnerRemember(m, v82, base.I64_extend_i32_u(v60), int32(_a_F_dsm_attach_1))
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
	v89 = v60 + int32(28)
	v91 = v60 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+12)) = l0
	v93 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+32)) = v93
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[5]))
	v100 = F_LWLockAcquire(m, v96+int32(_a_F_dsm_attach_2), v93)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[6]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v104 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[5]))
	F_LWLockRelease(m, v164+int32(_a_F_dsm_attach_2))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L36
	}
L28:
	;
	v112 = int32(0)
	goto L29
L29:
	;
	v121 = v112 * int32(24)
	v122 = v103 + int32(12) + v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if base.Ui32(v123) < base.Ui32(int32(2)) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L27
L31:
	;
	v151 = v112 + int32(1)
	if v151 != v104 {
		v112 = v151
		goto L29
	} else {
		goto L35
	}
L32:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if v126 != v127 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v129 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = v123 + v129
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v112
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+12)))
	if v133&v129 == int32(0) {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[7]))
	v140 = v103 + v121
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+20))
	v142 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v139 + v141<<(uint(v142)%32)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v146 << (uint(v142) % 32)
	goto L27
L35:
	;
	goto L30
L36:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	if v169 == int32(-1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_dsm_detach(m, v60)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if v176&int32(1) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	return int32(0)
L41:
	;
	v186 = F_dsm_impl_op(m, int32(1), v176, int32(0), v60+int32(20), v91, v89, int32(21))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	return v60
L44:
	;
	goto L43
L45:
	;
	F_errmsg_internal(m, int32(_a_F_dsm_attach_3), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_dsm_attach_4), int32(700), int32(_a_F_dsm_attach_5))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dsm_detach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	v6 = int32(_a_F_dsm_detach_0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[0])) = v8 + int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = v12
	goto L4
L2:
	;
	goto L3
L3:
	;
	v36 = int32(_a_F_dsm_detach_0)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[0])) = v38 - int32(1)
	v43 = l0 + int32(24)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v18
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v14-int32(8))))
	v24 = v14 - int32(16)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	F_pfree(m, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	m.T0[v25].(func(*base.Module, int32, int64))(m, l0, v22)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v30 != 0 {
		v14 = v30
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v45&int32(1) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v64 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v57 = F_dsm_impl_op(m, int32(2), v45, int32(0), l0+int32(20), v43, l0+int32(28), int32(19))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = int64(0)
	goto L12
L16:
	;
	goto L15
L17:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v146 != 0 {
		goto L33
	} else {
		goto L34
	}
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[1]))
	v72 = F_LWLockAcquire(m, v68+int32(_a_F_dsm_detach_1), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[2]))
	v78 = v75 + v64*int32(24)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v81 = v79 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(-1)
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[1]))
	F_LWLockRelease(m, v86+int32(_a_F_dsm_detach_1))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	if v81 != int32(1) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v93&int32(1) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v105 = F_dsm_impl_op(m, int32(3), v93, int32(0), l0+int32(20), v43, l0+int32(28), int32(19))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[1]))
	v114 = F_LWLockAcquire(m, v110+int32(_a_F_dsm_detach_1), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	if v105 == int32(0) {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v116&int32(1) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[3]))
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[2]))
	v125 = v122 + v64*int32(24)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+20))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v125)+24))
	F_FreePageManagerPut(m, v120, v126, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v132+v64*int32(24))+16)) = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[1]))
	F_LWLockRelease(m, v139+int32(_a_F_dsm_detach_1))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	goto L17
L33:
	;
	F_ResourceOwnerForget(m, v146, base.I64_extend_i32_u(l0), int32(_a_F_dsm_detach_2))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v154
	F_pfree(m, l0)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	return
}
func F_dsm_pin_segment(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v2
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_pin_segment[0]))
	v16 = F_LWLockAcquire(m, v12+int32(_a_F_dsm_pin_segment_0), v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_pin_segment[1]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v20*int32(24))+32)))
		if v24 != int32(1) {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v27&int32(1) == int32(0) {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_pin_segment[1]))
				v36 = v35
				v37 = v33
			} else {
				v36 = v19
				v37 = v20
			}
			v38 = int32(24)
			v41 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v36+v37*v38)+32)) = uint8(v41)
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v46 = v36 + v43*v38
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v47 + v41
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v36+v51*v38)+28)) = v55
			v58 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_pin_segment[0]))
			F_LWLockRelease(m, v58+int32(_a_F_dsm_pin_segment_0))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_dsm_pin_segment_1), int32(0))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_dsm_pin_segment_2), int32(975), int32(_a_F_dsm_pin_segment_3))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
