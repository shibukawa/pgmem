package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dsm_attach(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
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
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_dsm_attach[0])) = uint8(v16)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[1]))
	v20 = int32(0)
	if base.B2i32(v19 == v20)|base.B2i32(v19 == int32(_a_F_dsm_attach_0)) == v20 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L15
	} else {
		goto L44
	}
L5:
	;
	v29 = v19
	goto L8
L6:
	;
	goto L7
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[2]))
	if v53 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v37 == l0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v39 != int32(_a_F_dsm_attach_0) {
		v29 = v39
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	F_ResourceOwnerEnlarge(m, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[3]))
	v61 = F_MemoryContextAlloc(m, v59, int32(36))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L15
	} else {
		goto L17
	}
L15:
	;
	return int32(0)
L16:
	;
	goto L14
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[1]))
	if v64 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v67 = int32(_a_F_dsm_attach_0)
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[4])) = v67
	v71 = v67
	goto L20
L19:
	;
	v71 = v64
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(_a_F_dsm_attach_0)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v61
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[1])) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v61)+24)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v61)+16)) = int64(4294967295)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v83
	if v83 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_ResourceOwnerRemember(m, v83, v61, int32(_a_F_dsm_attach_1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L15
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v89 = v61 + int32(28)
	v91 = v61 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+12)) = l0
	v93 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+32)) = v93
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[5]))
	v100 = F_LWLockAcquire(m, v96+int32(_a_F_dsm_attach_2), v93)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L15
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[6]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v104 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[5]))
	F_LWLockRelease(m, v164+int32(_a_F_dsm_attach_2))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L15
	} else {
		goto L35
	}
L27:
	;
	v112 = int32(0)
	goto L28
L28:
	;
	v121 = v112 * int32(24)
	v122 = v103 + int32(12) + v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if base.Ui32(v123) < base.Ui32(int32(2)) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L26
L30:
	;
	v151 = v112 + int32(1)
	if v151 != v104 {
		v112 = v151
		goto L28
	} else {
		goto L34
	}
L31:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	if v126 != v127 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v129 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = v123 + v129
	*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v112
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+12)))
	if v133&v129 == int32(0) {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[7]))
	v140 = v103 + v121
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+20))
	v142 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v139 + v141<<(uint(v142)%32)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v146 << (uint(v142) % 32)
	goto L26
L34:
	;
	goto L29
L35:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	if v169 == int32(-1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_dsm_detach(m, v61)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L15
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	if v176&int32(1) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	return int32(0)
L40:
	;
	v186 = F_dsm_impl_op(m, int32(1), v176, int32(0), v61+int32(20), v91, v89, int32(21))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L15
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	return v61
L43:
	;
	goto L42
L44:
	;
	F_errmsg_internal(m, int32(_a_F_dsm_attach_3), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L15
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_dsm_attach_4), int32(692), int32(_a_F_dsm_attach_5))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L15
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dsm_detach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	v5 = int32(_a_F_dsm_detach_0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[0])) = v7 + int32(1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = v11
	goto L4
L2:
	;
	goto L3
L3:
	;
	v33 = int32(_a_F_dsm_detach_0)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[0])) = v35 - int32(1)
	v40 = l0 + int32(24)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v16
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(4))))
	v22 = v13 - int32(8)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	F_pfree(m, v22)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
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
	m.T0[v23].(func(*base.Module, int32, int32))(m, l0, v20)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v28 != 0 {
		v13 = v28
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v42&int32(1) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v61 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v54 = F_dsm_impl_op(m, int32(2), v42, int32(0), l0+int32(20), v40, l0+int32(28), int32(19))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
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
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v143 != 0 {
		goto L33
	} else {
		goto L34
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[1]))
	v69 = F_LWLockAcquire(m, v65+int32(_a_F_dsm_detach_1), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[2]))
	v75 = v72 + v61*int32(24)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	v78 = v76 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+16)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(-1)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[1]))
	F_LWLockRelease(m, v83+int32(_a_F_dsm_detach_1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	if v78 != int32(1) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v90&int32(1) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v102 = F_dsm_impl_op(m, int32(3), v90, int32(0), l0+int32(20), v40, l0+int32(28), int32(19))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[1]))
	v111 = F_LWLockAcquire(m, v107+int32(_a_F_dsm_detach_1), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	if v102 == int32(0) {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v113&int32(1) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[3]))
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[2]))
	v122 = v119 + v61*int32(24)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+20))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)+24))
	F_FreePageManagerPut(m, v117, v123, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v129+v61*int32(24))+16)) = int32(0)
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_detach[1]))
	F_LWLockRelease(m, v136+int32(_a_F_dsm_detach_1))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
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
	F_ResourceOwnerForget(m, v143, l0, int32(_a_F_dsm_detach_2))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v150
	F_pfree(m, l0)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
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
					F_errfinish(m, int32(_a_F_dsm_pin_segment_2), int32(967), int32(_a_F_dsm_pin_segment_3))
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
