package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dsm_attach(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[594])))
	if v11 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[594])) = uint8(v15)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[595]))
	if v18 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L15
	} else {
		goto L44
	}
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[581]))
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	if v18 == int32(4062852) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v24 = v18
	goto L8
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v32 == l0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L5
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v34 != int32(4062852) {
		v24 = v34
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	F_ResourceOwnerEnlarge(m, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v55 = F_MemoryContextAlloc(m, v53, int32(36))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
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
	v58 = *(*int32)(unsafe.Add(mBase, _consts[595]))
	if v58 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v61 = int32(4062852)
	*(*int32)(unsafe.Add(mBase, _consts[596])) = v61
	v65 = v61
	goto L20
L19:
	;
	v65 = v58
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(4062852)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v55
	*(*int32)(unsafe.Add(mBase, _consts[595])) = v55
	*(*int64)(unsafe.Add(mBase, uint32(v55)+24)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v55)+16)) = int64(4294967295)
	v77 = *(*int32)(unsafe.Add(mBase, _consts[581]))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v77
	if v77 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_ResourceOwnerRemember(m, v77, v55, int32(1586492))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L15
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v83 = v55 + int32(28)
	v85 = v55 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = l0
	v87 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+32)) = v87
	v90 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v94 = F_LWLockAcquire(m, v90+int32(4352), v87)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L15
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[597]))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v98 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v158+int32(4352))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L15
	} else {
		goto L35
	}
L27:
	;
	v105 = int32(0)
	goto L28
L28:
	;
	v115 = v97 + int32(12) + v105*int32(24)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if base.Ui32(v116) < base.Ui32(int32(2)) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L26
L30:
	;
	v146 = v105 + int32(1)
	if v146 != v98 {
		v105 = v146
		goto L28
	} else {
		goto L34
	}
L31:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	if v119 != v120 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v122 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v115)+4)) = v116 + v122
	*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = v105
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+12)))
	if v126&v122 == int32(0) {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _consts[598]))
	v135 = v97 + v105*int32(24)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
	v137 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v132 + v136<<(uint(v137)%32)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v141 << (uint(v137) % 32)
	goto L26
L34:
	;
	goto L29
L35:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	if v163 == int32(-1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_dsm_detach(m, v55)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L15
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	if v170&int32(1) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	return int32(0)
L40:
	;
	v180 = F_dsm_impl_op(m, int32(1), v170, int32(0), v55+int32(20), v85, v83, int32(21))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L15
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	return v55
L43:
	;
	goto L42
L44:
	;
	F_errmsg_internal(m, int32(395197), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L15
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(474046), int32(692), int32(309919))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
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
	var v22 int32
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
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
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
	v6 = int32(4437644)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	*(*int32)(unsafe.Add(mBase, _consts[83])) = v8 + int32(1)
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
	v36 = int32(4437644)
	v38 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	*(*int32)(unsafe.Add(mBase, _consts[83])) = v38 - int32(1)
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v14-int32(4))))
	v24 = v14 - int32(8)
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
	m.T0[v25].(func(*base.Module, int32, int32))(m, l0, v22)
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
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v147 != 0 {
		goto L33
	} else {
		goto L34
	}
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v72 = F_LWLockAcquire(m, v68+int32(4352), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v75 = v64 * int32(24)
	v77 = *(*int32)(unsafe.Add(mBase, _consts[597]))
	v80 = v75 + v77 + int32(16)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v83 = v81 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v83
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(-1)
	v88 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v88+int32(4352))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	if v83 != int32(1) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v95&int32(1) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v107 = F_dsm_impl_op(m, int32(3), v95, int32(0), l0+int32(20), v43, l0+int32(28), int32(19))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v116 = F_LWLockAcquire(m, v112+int32(4352), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	if v107 == int32(0) {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v118&int32(1) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[598]))
	v124 = *(*int32)(unsafe.Add(mBase, _consts[597]))
	v127 = v124 + v64*int32(24)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+20))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v127)+24))
	F_FreePageManagerPut(m, v122, v128, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[597]))
	*(*int32)(unsafe.Add(mBase, uint32(v134+v75)+16)) = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v139+int32(4352))
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
	F_ResourceOwnerForget(m, v147, l0, int32(1586492))
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v2
	v12 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v16 = F_LWLockAcquire(m, v12+int32(4352), v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[597]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v22 = v20 * int32(24)
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v22)+32)))
		if v24 != int32(1) {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v27&int32(1) == int32(0) {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v37 = *(*int32)(unsafe.Add(mBase, _consts[597]))
				v38 = v37
				v39 = v33 * int32(24)
			} else {
				v38 = v19
				v39 = v22
			}
			v41 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v38+v39)+32)) = uint8(v41)
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v44 = int32(24)
			v48 = v38 + v43*v44 + int32(16)
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
			*(*int32)(unsafe.Add(mBase, uint32(v48))) = v49 + v41
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v38+v53*v44)+28)) = v57
			v60 = *(*int32)(unsafe.Add(mBase, _consts[29]))
			F_LWLockRelease(m, v60+int32(4352))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(431549), int32(0))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					F_errfinish(m, int32(474046), int32(967), int32(89160))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
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
