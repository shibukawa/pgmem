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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_closing_fd[0]))
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
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_closing_fd[1]))
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
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_closing_fd[0]))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l0
	F_errmsg_internal(m, int32(_a_F_pgaio_closing_fd_0), v14+int32(32))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_pgaio_closing_fd_1), int32(1237), int32(_a_F_pgaio_closing_fd_2))
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
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_closing_fd[0]))
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
	v106 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
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
	if v106 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_errhidestmt(m)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	F_pgaio_io_wait(m, v90, v88)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L46
	}
L34:
	;
	F_errhidecontext(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_closing_fd[2]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+24))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+2)))
	if base.Ui32(v119) <= base.Ui32(int32(2)) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v127<<(uint(int32(2))%32))+uint32(_c_F_pgaio_closing_fd[3])))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	goto L40
L37:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119<<(uint(int32(2))%32))+uint32(_c_F_pgaio_closing_fd[4])))
	v126 = v124
	goto L39
L38:
	;
	v126 = int32(0)
	goto L39
L39:
	;
	goto L36
L40:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if base.Ui32(v132) <= base.Ui32(int32(7)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132<<(uint(int32(2))%32))+uint32(_c_F_pgaio_closing_fd[5])))
	v138 = v137
	goto L43
L42:
	;
	v138 = int32(0)
	goto L43
L43:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_closing_fd[0]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(16)))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(20)))) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = (v90 - v115) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v138
	F_errmsg_internal(m, int32(_a_F_pgaio_closing_fd_3), v14)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_pgaio_closing_fd_1), int32(1276), int32(_a_F_pgaio_closing_fd_2))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	goto L33
L46:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_closing_fd[0]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+160))
	if v165 != 0 {
		v63 = v164
		goto L18
	} else {
		goto L47
	}
L47:
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_get_wref[0]))
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
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
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
					v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_update_state[0]))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
					if base.Ui32(v27) <= base.Ui32(int32(2)) {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v27<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_update_state[1])))
						v34 = v32
					} else {
						v34 = int32(0)
					}
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v35<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_update_state[2])))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if base.Ui32(v40) <= base.Ui32(int32(7)) {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v40<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_update_state[3])))
						v46 = v45
					} else {
						v46 = v3
					}
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v2<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_update_state[3])))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v49
					*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v39
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = (l0 - v23) >> (uint(int32(7)) % 32)
					F_errmsg_internal(m, int32(_a_F_pgaio_io_update_state_0), v11)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_pgaio_io_update_state_1), int32(397), int32(_a_F_pgaio_io_update_state_2))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_shutdown[0]))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
	if v16 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_shutdown[0]))
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
	F_errmsg_internal(m, int32(_a_F_pgaio_shutdown_0), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(_a_F_pgaio_shutdown_1), int32(1206), int32(_a_F_pgaio_shutdown_2))
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
	*(*int32)(unsafe.Add(mBase, _c_F_pgaio_shutdown[0])) = int32(0)
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
	v113 = m.ExcPending
	if v113 != 0 {
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
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_shutdown[1]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+24))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+2)))
	if base.Ui32(v71) <= base.Ui32(int32(2)) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79<<(uint(int32(2))%32))+uint32(_c_F_pgaio_shutdown[2])))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	goto L24
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71<<(uint(int32(2))%32))+uint32(_c_F_pgaio_shutdown[3])))
	v78 = v76
	goto L23
L22:
	;
	v78 = int32(0)
	goto L23
L23:
	;
	goto L20
L24:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if base.Ui32(v84) <= base.Ui32(int32(7)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84<<(uint(int32(2))%32))+uint32(_c_F_pgaio_shutdown[4])))
	v90 = v89
	goto L27
L26:
	;
	v90 = int32(0)
	goto L27
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_shutdown[0]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(16)))) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = (v54 - v67) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v90
	F_errmsg_internal(m, int32(_a_F_pgaio_shutdown_3), v12)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_pgaio_shutdown_1), int32(1312), int32(_a_F_pgaio_shutdown_4))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	goto L17
L30:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_shutdown[0]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+160))
	if v116 != 0 {
		v43 = v115
		goto L12
	} else {
		goto L31
	}
L31:
	;
	goto L13
}
