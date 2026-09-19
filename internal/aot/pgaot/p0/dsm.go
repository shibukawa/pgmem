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
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
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
	v191 = m.ExcPending
	if v191 != 0 {
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
	F_ResourceOwnerRemember(m, v82, v60, int32(_a_F_dsm_attach_1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v88 = v60 + int32(28)
	v90 = v60 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+12)) = l0
	v92 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+32)) = v92
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[5]))
	v99 = F_LWLockAcquire(m, v95+int32(_a_F_dsm_attach_2), v92)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[6]))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	if v103 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[5]))
	F_LWLockRelease(m, v163+int32(_a_F_dsm_attach_2))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L36
	}
L28:
	;
	v111 = int32(0)
	goto L29
L29:
	;
	v120 = v111 * int32(24)
	v121 = v102 + int32(12) + v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if base.Ui32(v122) < base.Ui32(int32(2)) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L27
L31:
	;
	v150 = v111 + int32(1)
	if v150 != v103 {
		v111 = v150
		goto L29
	} else {
		goto L35
	}
L32:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if v125 != v126 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v128 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v122 + v128
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v111
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+12)))
	if v132&v128 == int32(0) {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_attach[7]))
	v139 = v102 + v120
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
	v141 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v138 + v140<<(uint(v141)%32)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v139)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v145 << (uint(v141) % 32)
	goto L27
L35:
	;
	goto L30
L36:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	if v168 == int32(-1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_dsm_detach(m, v60)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if v175&int32(1) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	return int32(0)
L41:
	;
	v185 = F_dsm_impl_op(m, int32(1), v175, int32(0), v60+int32(20), v90, v88, int32(21))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
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
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_dsm_attach_4), int32(692), int32(_a_F_dsm_attach_5))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
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
