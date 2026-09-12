package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgaio_closing_fd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	if v17 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return
L2:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+22)))
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[721]))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v48 != int32(1) {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	return
L7:
	;
	if v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_errhidestmt(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L15
	}
L11:
	;
	F_errhidecontext(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l0
	F_errmsg_internal(m, int32(449376), v14+int32(32))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(495496), int32(1237), int32(437033))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	goto L5
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+160))
	if v53 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v63 = v52
	goto L18
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v63)+156))
	if v71 == int32(0) {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L1
L20:
	;
	v75 = v63 + int32(152)
	if v71 == v75 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v79 = v71
	goto L22
L22:
	;
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v79)+24))
	v90 = v79 - int32(24)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+2)))
	v92 = int32(1)
	if base.Ui32((v91-v92)&int32(255)) <= base.Ui32(v92) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	if v90 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L24:
	;
	goto L23
L25:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+88))
	v101 = base.B2i32(v98 == l0)
	goto L27
L26:
	;
	v101 = int32(0)
	goto L27
L27:
	;
	if v101 != 0 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v102 != v75 {
		v79 = v102
		goto L22
	} else {
		goto L29
	}
L29:
	;
	goto L1
L30:
	;
	v108 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	if v108 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_errhidestmt(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	F_pgaio_io_wait(m, v90, v88)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L6
	} else {
		goto L47
	}
L35:
	;
	F_errhidecontext(m)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v114 = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+24))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+2)))
	if base.Ui32(v122) <= base.Ui32(int32(2)) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v131<<(uint(int32(2))%32))+uint32(_consts[718])))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	goto L41
L38:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v122<<(uint(int32(2))%32))+uint32(_consts[717])))
	v130 = v129
	goto L40
L39:
	;
	v130 = v114
	goto L40
L40:
	;
	goto L37
L41:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if base.Ui32(v138) <= base.Ui32(int32(7)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v138<<(uint(int32(2))%32))+uint32(_consts[719])))
	v146 = v145
	goto L44
L43:
	;
	v146 = v114
	goto L44
L44:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(16)))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(20)))) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = (v90 - v117) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v146
	F_errmsg_internal(m, int32(174412), v14)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(495496), int32(1276), int32(437033))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	goto L34
L47:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+160))
	if v173 != 0 {
		v63 = v172
		goto L18
	} else {
		goto L48
	}
L48:
	;
	goto L19
}
func F_pgaio_io_get_target_data(m *base.Module, l0 int32) int32 {
	return l0 + int32(104)
}
func F_pgaio_io_get_wref(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = (l0 - v5) >> (uint(int32(7)) % 32)
	v10 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+52)))
	*(*uint32)(unsafe.Add(mBase, uint32(l1)+4)) = uint32(v10)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*uint32)(unsafe.Add(mBase, uint32(l1)+8)) = uint32(v12)
	return
}
func F_pgaio_io_update_state(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v2 = l1
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v15 = F_errstart(m, int32(10), v3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if v15 != 0 {
			F_errhidestmt(m)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_errhidecontext(m)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, _consts[716]))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
					if base.Ui32(v28) <= base.Ui32(int32(2)) {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v28<<(uint(int32(2))%32))+uint32(_consts[717])))
						v36 = v35
					} else {
						v36 = int32(0)
					}
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v37<<(uint(int32(2))%32))+uint32(_consts[718])))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if base.Ui32(v44) <= base.Ui32(int32(7)) {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v44<<(uint(int32(2))%32))+uint32(_consts[719])))
						v52 = v51
					} else {
						v52 = v3
					}
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v2<<(uint(int32(2))%32))+uint32(_consts[719])))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v57
					*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v52
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v43
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v36
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = (l0 - v23) >> (uint(int32(7)) % 32)
					F_errmsg_internal(m, int32(182892), v11)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(495496), int32(397), int32(351367))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
							m.G0 = v11 + int32(32)
							return
						}
					}
				}
			}
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
			m.G0 = v11 + int32(32)
			return
		}
	}
}
func F_pgaio_shutdown(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
	if v16 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+160))
	if v40 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v19 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)) = uint8(v19)
	F_pgaio_submit_staged(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v25 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v25 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_errmsg_internal(m, int32(256075), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(495496), int32(1206), int32(240834))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L1
L9:
	;
	v43 = v39
	goto L12
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[720])) = int32(0)
	m.G0 = v12 + int32(32)
	return
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+156))
	v54 = v52 - int32(24)
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v52)+24))
	v58 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	if v58 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_errhidestmt(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_pgaio_io_wait(m, v54, v55)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L3
	} else {
		goto L30
	}
L18:
	;
	F_errhidecontext(m)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v64 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+24))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+2)))
	if base.Ui32(v72) <= base.Ui32(int32(2)) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81<<(uint(int32(2))%32))+uint32(_consts[718])))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	goto L24
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v72<<(uint(int32(2))%32))+uint32(_consts[717])))
	v80 = v79
	goto L23
L22:
	;
	v80 = v64
	goto L23
L23:
	;
	goto L20
L24:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if base.Ui32(v88) <= base.Ui32(int32(7)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v88<<(uint(int32(2))%32))+uint32(_consts[719])))
	v96 = v95
	goto L27
L26:
	;
	v96 = v64
	goto L27
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(16)))) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = (v54 - v67) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v96
	F_errmsg_internal(m, int32(174309), v12)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(495496), int32(1312), int32(243446))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	goto L17
L30:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+160))
	if v122 != 0 {
		v43 = v121
		goto L12
	} else {
		goto L31
	}
L31:
	;
	goto L13
}
