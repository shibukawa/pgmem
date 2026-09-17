package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgstat_backend_flush_cb(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_pgstat_flush_backend(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pgstat_beinit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_beinit[0]))
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_beinit[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_beinit[2])) = v3 + v5*int32(408)
	F_on_shmem_exit(m, int32(1195), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		return
	}
}
func F_pgstat_bgwriter_snapshot_cb(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bgwriter_snapshot_cb[0]))
	v10 = v8 + int32(344)
	goto L1
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+336))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bgwriter_snapshot_cb[1]))
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v39 = v8 + int32(320)
	v41 = F_LWLockAcquire(m, v39, int32(1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L10
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_bgwriter_snapshot_cb[2])) = v23
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_bgwriter_snapshot_cb[3])) = v26
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_bgwriter_snapshot_cb[4])) = v29
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_bgwriter_snapshot_cb[5])) = v32
	if v17&int32(1) != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	return
L7:
	;
	goto L5
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+336))
	if v17 != v36 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L2
L10:
	;
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v8)+392))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v8)+384))
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v8)+376))
	F_LWLockRelease(m, v39)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v48 = int32(_a_F_pgstat_bgwriter_snapshot_cb_0)
	v50 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_bgwriter_snapshot_cb[5]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_bgwriter_snapshot_cb[5])) = v50 - v45
	v53 = int32(_a_F_pgstat_bgwriter_snapshot_cb_1)
	v55 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_bgwriter_snapshot_cb[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_bgwriter_snapshot_cb[4])) = v55 - v44
	v58 = int32(_a_F_pgstat_bgwriter_snapshot_cb_2)
	v60 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_bgwriter_snapshot_cb[3]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_bgwriter_snapshot_cb[3])) = v60 - v43
	return
}
func F_pgstat_build_snapshot(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_build_snapshot[0]))
	if v14 != int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_pgstat_prep_snapshot(m)
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
	m.G0 = v11 + int32(48)
	return
L4:
	;
	return
L5:
	;
	v23 = m.G0
	v24 = int32(16)
	v25 = v23 - v24
	m.G0 = v25
	F_gettimeofday(m, v25)
	mBase = m.M
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+8)))
	m.G0 = v25 + v24
	goto L6
L6:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_build_snapshot[1])) = v29 + v28*int64(1000000) - int64(946684800000000)
	v40 = v11 + int32(20)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_build_snapshot[2]))
	v43 = int32(0)
	v44 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+4)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v40)+12)) = v44
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+24)) = uint8(v43)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = int32(-1)
	goto L7
L7:
	;
	v52 = F_dshash_seq_next(m, v40)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	if v52 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v55 = v52
	goto L12
L10:
	;
	goto L11
L11:
	;
	F_dshash_seq_term(m, v11+int32(20))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L42
	}
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v66 = base.B2i32(base.Ui32(int32(11)) < base.Ui32(v62-int32(1)))
	if v66 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L11
L14:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v91 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v90 = v62*int32(72) + int32(_a_F_pgstat_build_snapshot_0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	if base.Ui32(int32(8)) < base.Ui32(v62-int32(24)) {
		v90 = int32(0)
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v78 = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_build_snapshot[3]))
	if v80 == v78 {
		v90 = v78
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v80+v62<<(uint(int32(2))%32)-int32(96))))
	v90 = v88
	goto L14
L20:
	;
	v169 = F_dshash_seq_next(m, v11+int32(20))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L40
	}
L21:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+16)))
	if v102 != 0 {
		goto L20
	} else {
		goto L25
	}
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_build_snapshot[4]))
	if v91 == v95 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v97&int32(2) == int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_build_snapshot[5]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
	v106 = F_dsa_get_address(m, v104, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v55)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v110
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_build_snapshot[6]))
	v116 = F_pgstat_snapshot_insert(m, v113, v11, v11+int32(19))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_build_snapshot[7]))
	if base.Ui32(int32(11)) < base.Ui32(v62-int32(1)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_build_snapshot[3]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121+v62<<(uint(int32(2))%32)-int32(96))))
	v132 = v127
	goto L30
L29:
	;
	v132 = v62*int32(72) + int32(_a_F_pgstat_build_snapshot_0)
	goto L30
L30:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+20))
	v134 = F_MemoryContextAlloc(m, v119, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = v134
	v138 = v106 + int32(4)
	v140 = F_LWLockAcquire(m, v138, int32(1))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	if base.Ui32(int32(11)) < base.Ui32(v62-int32(1)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_build_snapshot[3]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v144+v62<<(uint(int32(2))%32)-int32(96))))
	v155 = v150
	goto L35
L34:
	;
	v155 = v62*int32(72) + int32(_a_F_pgstat_build_snapshot_0)
	goto L35
L35:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+20))
	if v156 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	base.MemoryCopy(m, v142, v106+v157, v156)
	goto L38
L37:
	;
	goto L38
L38:
	;
	F_LWLockRelease(m, v138)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	goto L20
L40:
	;
	if v169 != 0 {
		v55 = v169
		goto L12
	} else {
		goto L41
	}
L41:
	;
	goto L13
L42:
	;
	v185 = int32(1)
	goto L43
L43:
	;
	if base.Ui32(int32(13)) <= base.Ui32(v185) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_build_snapshot[0])) = int32(2)
	goto L3
L45:
	;
	v247 = v185 + int32(1)
	if v247 != int32(33) {
		v185 = v247
		goto L43
	} else {
		goto L61
	}
L46:
	;
	v230 = v226 + v229
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_build_snapshot[8]))
	if v232 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L47:
	;
	v195 = v185 - int32(24)
	if base.Ui32(int32(8)) < base.Ui32(v195) {
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v217 = v185 * int32(72)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+uint32(_c_F_pgstat_build_snapshot[9]))))
	if v218&int32(1) == int32(0) {
		goto L45
	} else {
		goto L54
	}
L50:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_build_snapshot[3]))
	if v199 == int32(0) {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v199+v185<<(uint(int32(2))%32)-int32(96))))
	if v207 == int32(0) {
		goto L45
	} else {
		goto L52
	}
L52:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v210&int32(1) == int32(0) {
		goto L45
	} else {
		goto L53
	}
L53:
	;
	v226 = v195
	v228 = v207
	v229 = int32(_a_F_pgstat_build_snapshot_1)
	goto L46
L54:
	;
	v226 = v185
	v228 = v217 + int32(_a_F_pgstat_build_snapshot_0)
	v229 = int32(_a_F_pgstat_build_snapshot_2)
	goto L46
L55:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v228)+64))
	m.T0[v238].(func(*base.Module))(m)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L60
	}
L56:
	;
	v235 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v230))) = uint8(v235)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	if v237 != 0 {
		goto L45
	} else {
		goto L59
	}
L59:
	;
	goto L55
L60:
	;
	v241 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v230))) = uint8(v241)
	goto L45
L61:
	;
	goto L44
}
func F_pgstat_checkpointer_reset_all_cb(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_reset_all_cb[0]))
	v10 = v8 + int32(408)
	v12 = F_LWLockAcquire(m, v10, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+424))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_reset_all_cb[1]))
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+512)) = l0
	F_LWLockRelease(m, v10)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L11
	}
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	base.MemoryCopy(m, v8+int32(520), v8+int32(432), int32(88))
	if v24&int32(1) != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+424))
	if v24 != v33 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	return
}
func F_pgstat_fetch_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v77 int64
	_ = v77
	var v83 int64
	_ = v83
	var v89 int64
	_ = v89
	var v94 int64
	_ = v94
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int64
	_ = v181
	var v183 int64
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v23 = l0 - int32(1)
	if base.Ui32(v23) <= base.Ui32(int32(11)) {
		v43 = l0*int32(72) + int32(_a_F_pgstat_fetch_entry_0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_pgstat_prep_snapshot(m)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
		v43 = int32(0)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_entry[0]))
	if v33 == v31 {
		v43 = v31
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33+l0<<(uint(int32(2))%32)-int32(96))))
	v43 = v41
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l0
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_entry[1]))
	if v52 == int32(2) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_pgstat_build_snapshot(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	v59 = v52
	goto L9
L9:
	;
	if int32(0) < v59 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_entry[1]))
	v59 = v58
	goto L9
L11:
	;
	m.G0 = v16 - int32(-64)
	return v259
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_entry[2]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
	v67 = int64(23)
	v70 = int64(2388976653695081527)
	v71 = (int64(base.Ui64(v66)>>(uint(v67)%64)) ^ v66) * v70
	v72 = int64(47)
	v77 = int64(-8645972361240307355)
	v83 = (int64(base.Ui64(l2)>>(uint(v67)%64)) ^ l2) * v70
	v89 = ((v71^int64(base.Ui64(v71)>>(uint(v72)%64))^int64(-9208349263878056368))*v77 ^ int64(base.Ui64(v83)>>(uint(v72)%64)) ^ v83) * v77
	v94 = (int64(base.Ui64(v89)>>(uint(v67)%64)) ^ v89) * v70
	v102 = v65 & base.I32_wrap_i64(int64(base.Ui64(v94)>>(uint(v72)%64))^v94-int64(base.Ui64(v94)>>(uint(int64(32))%64)))
	v105 = v64 + v102*int32(24)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+16)))
	if v106 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_entry[3])) = v59
	v169 = int32(0)
	v171 = F_pgstat_get_entry_ref(m, l0, l1, l2, v169, v169)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L5
	} else {
		goto L25
	}
L15:
	;
	if v59 == int32(2) {
		v259 = v4
		goto L11
	} else {
		goto L23
	}
L16:
	;
	v114 = v105
	v117 = v102
	goto L17
L17:
	;
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v114)))
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v114)+8))
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v16)+56))
	if v122^v123|(v125^v126) != int64(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v114)+20))
	v259 = v138
	goto L11
L19:
	;
	v133 = (v117 + int32(1)) & v65
	v136 = v64 + v133*int32(24)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+16)))
	if v137 != 0 {
		v114 = v136
		v117 = v133
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	goto L15
L23:
	;
	goto L14
L24:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_entry[1]))
	if v196 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L25:
	;
	if v171 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+16)))
	if v174 != int32(1) {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_entry[1]))
	if v178 != int32(1) {
		v259 = v4
		goto L11
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v181
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v16)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v183
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_entry[2]))
	v191 = F_pgstat_snapshot_insert(m, v186, v14+int32(-56), v14+int32(-17))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+20)) = int32(0)
	v259 = v4
	goto L11
L32:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	v212 = F_LWLockAcquire(m, v208+int32(4), int32(1))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L5
	} else {
		goto L38
	}
L33:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	v200 = F_palloc(m, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_entry[4]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	v205 = F_MemoryContextAlloc(m, v203, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L5
	} else {
		goto L37
	}
L36:
	;
	v207 = v200
	goto L32
L37:
	;
	v207 = v205
	goto L32
L38:
	;
	if base.Ui32(v23) <= base.Ui32(int32(11)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v228 = l0*int32(72) + int32(_a_F_pgstat_fetch_entry_0)
	goto L41
L40:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_entry[0]))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v221+l0<<(uint(int32(2))%32)-int32(96))))
	v228 = v227
	goto L41
L41:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v229 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v228)+16))
	base.MemoryCopy(m, v207, v230+v231, v229)
	goto L44
L43:
	;
	goto L44
L44:
	;
	F_pgstat_unlock_entry(m, v171)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_entry[1]))
	if v237 <= int32(0) {
		v259 = v207
		goto L11
	} else {
		goto L46
	}
L46:
	;
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v240
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v16)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v242
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_entry[2]))
	v250 = F_pgstat_snapshot_insert(m, v245, v14+int32(-40), v14+int32(-17))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+20)) = v207
	v259 = v207
	goto L11
}
func F_pgstat_fetch_stat_numbackends(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	F_pgstat_read_current_status(m)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return int32(0)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_stat_numbackends[0]))
		return v6
	}
}
func F_pgstat_fetch_stat_tabentry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v47 int32
	_ = v47
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	v5 = int32(1)
	if l0 <= int32(3591) {
		if l0 <= int32(2670) {
			switch l0 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v76 = v5
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v76 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
					v76 = int32(0)
				} else {
					v76 = v5
				}
			}
		} else {
			v17 = l0 - int32(2671)
			if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v17))|base.B2i32(int32(1)<<(uint(v17)%32)&int32(226492515) == int32(0)) != 0 {
				if base.B2i32(base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l0-int32(2846)) < base.Ui32(int32(2))) != 0 {
					v76 = v5
				} else {
					v76 = int32(0)
				}
			} else {
				v76 = v5
			}
		}
	} else {
		if l0 <= int32(_a_F_pgstat_fetch_stat_tabentry_0) {
			v30 = l0 - int32(_a_F_pgstat_fetch_stat_tabentry_1)
			if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v30))|base.B2i32(int32(1)<<(uint(v30)%32)&int32(963) == int32(0)) != 0 {
				if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
					v76 = v5
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
						v76 = int32(0)
					} else {
						v76 = v5
					}
				}
			} else {
				v76 = v5
			}
		} else {
			switch l0 - int32(_a_F_pgstat_fetch_stat_tabentry_2) {
			case 0, 1, 2, 3, 4, 59, 60:
				v76 = v5
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v76 = int32(0)
			default:
				if base.Ui32(l0-int32(_a_F_pgstat_fetch_stat_tabentry_3)) < base.Ui32(int32(3)) {
					v76 = v5
				} else {
					v47 = l0 - int32(_a_F_pgstat_fetch_stat_tabentry_4)
					if base.Ui32(int32(15)) < base.Ui32(v47) {
						v76 = int32(0)
					} else {
						if int32(1)<<(uint(v47)%32)&int32(_a_F_pgstat_fetch_stat_tabentry_5) != 0 {
							v76 = v5
						} else {
							v76 = int32(0)
						}
					}
				}
			}
		}
	}
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_stat_tabentry[0]))
	if v76 != 0 {
		v81 = int32(0)
	} else {
		v81 = v80
	}
	v83 = F_pgstat_fetch_entry(m, int32(2), v81, base.I64_extend_i32_u(l0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		return int32(0)
	} else {
		return v83
	}
}
func F_pgstat_flush_backend(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v97 int32
	_ = v97
	var v111 int32
	_ = v111
	var v127 int32
	_ = v127
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
	var v202 int64
	_ = v202
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v209 int64
	_ = v209
	var v210 int64
	_ = v210
	var v213 int64
	_ = v213
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v219 int64
	_ = v219
	var v220 int64
	_ = v220
	var v223 int64
	_ = v223
	var v224 int64
	_ = v224
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v230 int64
	_ = v230
	var v233 int64
	_ = v233
	var v234 int64
	_ = v234
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v241 int64
	_ = v241
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v247 int64
	_ = v247
	var v248 int64
	_ = v248
	var v251 int64
	_ = v251
	var v252 int64
	_ = v252
	var v255 int64
	_ = v255
	var v256 int64
	_ = v256
	var v258 int64
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v300 int64
	_ = v300
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int64
	_ = v315
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v322 int64
	_ = v322
	var v324 int64
	_ = v324
	var v325 int64
	_ = v325
	var v329 int64
	_ = v329
	var v331 int64
	_ = v331
	var v332 int64
	_ = v332
	var v336 int64
	_ = v336
	var v338 int64
	_ = v338
	var v339 int64
	_ = v339
	var v343 int64
	_ = v343
	var v344 int64
	_ = v344
	var v347 int64
	_ = v347
	var v348 int64
	_ = v348
	var v351 int64
	_ = v351
	var v352 int64
	_ = v352
	var v355 int64
	_ = v355
	var v356 int64
	_ = v356
	var v361 int64
	_ = v361
	var v365 int64
	_ = v365
	var v369 int64
	_ = v369
	var v373 int64
	_ = v373
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	v3 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(2880)
	m.G0 = v25
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[0]))
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v28))|base.B2i32(int32(1)<<(uint(v28)%32)&int32(_a_F_pgstat_flush_backend_0) == v3) != 0 {
		v384 = v3
		m.G0 = v25 + int32(2880)
		return v384
	} else {
		v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[1])))
		v40 = l1 & v39
		v42 = l1 & int32(2)
		if v42 != 0 {
			v44 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[2]))
			v46 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[3]))
			if (base.B2i32(v44 != v46)|v40)&int32(1) != 0 {
				v58 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[4])))
				v59 = F_pgstat_get_entry_ref_locked(m, int32(6), int32(0), v58, l0)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					if v59 == int32(0) {
						v384 = int32(1)
						m.G0 = v25 + int32(2880)
						return v384
					} else {
						if l1&int32(1) == int32(0) {
						} else {
							v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[1])))
							if v71 == int32(0) {
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
								base.MemoryCopy(m, v25, int32(_a_F_pgstat_flush_backend_1), int32(2880))
								v97 = v3
								for {
									v111 = v97 * int32(320)
									v127 = int32(0)
									for {
										v142 = v127 << (uint(int32(6)) % 32)
										v143 = v111 + (v74 + int32(992)) + v142
										v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
										v145 = v142 + (v111 + (v25 + int32(960)))
										v146 = *(*int64)(unsafe.Add(mBase, uint32(v145)))
										*(*int64)(unsafe.Add(mBase, uint32(v143))) = v144 + v146
										v149 = v142 + (v111 + (v74 + int32(32)))
										v150 = *(*int64)(unsafe.Add(mBase, uint32(v149)))
										v151 = v142 + (v111 + v25)
										v152 = *(*int64)(unsafe.Add(mBase, uint32(v151)))
										*(*int64)(unsafe.Add(mBase, uint32(v149))) = v150 + v152
										v155 = v142 + (v74 + int32(1952) + v111)
										v156 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
										v157 = v142 + (v111 + (v25 + int32(1920)))
										v158 = *(*int64)(unsafe.Add(mBase, uint32(v157)))
										v159 = int64(1000)
										v160 = base.I64_div_s(v158, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155))) = v156 + v160
										v163 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
										v164 = *(*int64)(unsafe.Add(mBase, uint32(v145)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v143)+8)) = v163 + v164
										v167 = *(*int64)(unsafe.Add(mBase, uint32(v149)+8))
										v168 = *(*int64)(unsafe.Add(mBase, uint32(v151)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v149)+8)) = v167 + v168
										v171 = *(*int64)(unsafe.Add(mBase, uint32(v155)+8))
										v172 = *(*int64)(unsafe.Add(mBase, uint32(v157)+8))
										v174 = base.I64_div_s(v172, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155)+8)) = v171 + v174
										v177 = *(*int64)(unsafe.Add(mBase, uint32(v143)+16))
										v178 = *(*int64)(unsafe.Add(mBase, uint32(v145)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v143)+16)) = v177 + v178
										v181 = *(*int64)(unsafe.Add(mBase, uint32(v149)+16))
										v182 = *(*int64)(unsafe.Add(mBase, uint32(v151)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v149)+16)) = v181 + v182
										v185 = *(*int64)(unsafe.Add(mBase, uint32(v155)+16))
										v186 = *(*int64)(unsafe.Add(mBase, uint32(v157)+16))
										v188 = base.I64_div_s(v186, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155)+16)) = v185 + v188
										v191 = *(*int64)(unsafe.Add(mBase, uint32(v143)+24))
										v192 = *(*int64)(unsafe.Add(mBase, uint32(v145)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v143)+24)) = v191 + v192
										v195 = *(*int64)(unsafe.Add(mBase, uint32(v149)+24))
										v196 = *(*int64)(unsafe.Add(mBase, uint32(v151)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v149)+24)) = v195 + v196
										v199 = *(*int64)(unsafe.Add(mBase, uint32(v155)+24))
										v200 = *(*int64)(unsafe.Add(mBase, uint32(v157)+24))
										v202 = base.I64_div_s(v200, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155)+24)) = v199 + v202
										v205 = *(*int64)(unsafe.Add(mBase, uint32(v143)+32))
										v206 = *(*int64)(unsafe.Add(mBase, uint32(v145)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v143)+32)) = v205 + v206
										v209 = *(*int64)(unsafe.Add(mBase, uint32(v149)+32))
										v210 = *(*int64)(unsafe.Add(mBase, uint32(v151)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v149)+32)) = v209 + v210
										v213 = *(*int64)(unsafe.Add(mBase, uint32(v155)+32))
										v214 = *(*int64)(unsafe.Add(mBase, uint32(v157)+32))
										v216 = base.I64_div_s(v214, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155)+32)) = v213 + v216
										v219 = *(*int64)(unsafe.Add(mBase, uint32(v143)+40))
										v220 = *(*int64)(unsafe.Add(mBase, uint32(v145)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v143)+40)) = v219 + v220
										v223 = *(*int64)(unsafe.Add(mBase, uint32(v149)+40))
										v224 = *(*int64)(unsafe.Add(mBase, uint32(v151)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v149)+40)) = v223 + v224
										v227 = *(*int64)(unsafe.Add(mBase, uint32(v155)+40))
										v228 = *(*int64)(unsafe.Add(mBase, uint32(v157)+40))
										v230 = base.I64_div_s(v228, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155)+40)) = v227 + v230
										v233 = *(*int64)(unsafe.Add(mBase, uint32(v143)+48))
										v234 = *(*int64)(unsafe.Add(mBase, uint32(v145)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v143)+48)) = v233 + v234
										v237 = *(*int64)(unsafe.Add(mBase, uint32(v149)+48))
										v238 = *(*int64)(unsafe.Add(mBase, uint32(v151)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v149)+48)) = v237 + v238
										v241 = *(*int64)(unsafe.Add(mBase, uint32(v155)+48))
										v242 = *(*int64)(unsafe.Add(mBase, uint32(v157)+48))
										v244 = base.I64_div_s(v242, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155)+48)) = v241 + v244
										v247 = *(*int64)(unsafe.Add(mBase, uint32(v143)+56))
										v248 = *(*int64)(unsafe.Add(mBase, uint32(v145)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v143)+56)) = v247 + v248
										v251 = *(*int64)(unsafe.Add(mBase, uint32(v149)+56))
										v252 = *(*int64)(unsafe.Add(mBase, uint32(v151)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v149)+56)) = v251 + v252
										v255 = *(*int64)(unsafe.Add(mBase, uint32(v155)+56))
										v256 = *(*int64)(unsafe.Add(mBase, uint32(v157)+56))
										v258 = base.I64_div_s(v256, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155)+56)) = v255 + v258
										v262 = v127 + int32(1)
										if v262 != int32(5) {
											v127 = v262
											continue
										} else {
											break
										}
										break
									}
									v266 = v97 + int32(1)
									if v266 != int32(3) {
										v97 = v266
										continue
									} else {
										break
									}
									break
								}
								v270 = int32(0)
								base.MemoryFill(m, int32(_a_F_pgstat_flush_backend_1), v270, int32(2880))
								*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[1])) = uint8(v270)
							}
						}
						if v42 == int32(0) {
						} else {
							v300 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v300
							*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v300
							*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v300
							*(*int64)(unsafe.Add(mBase, uint32(v25))) = v300
							v309 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[2]))
							v311 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[3]))
							if v309 == v311 {
							} else {
								v313 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
								v314 = int32(_a_F_pgstat_flush_backend_2)
								v315 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
								v317 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[5]))
								v318 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[6]))
								*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v315 + (v317 - v318)
								v322 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
								v324 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[2]))
								v325 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[3]))
								*(*int64)(unsafe.Add(mBase, uint32(v25))) = v322 + (v324 - v325)
								v329 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
								v331 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[7]))
								v332 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[8]))
								*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v329 + (v331 - v332)
								v336 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
								v338 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[9]))
								v339 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[10]))
								*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v336 + (v338 - v339)
								v343 = *(*int64)(unsafe.Add(mBase, uint32(v313)+2936))
								v344 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v313)+2936)) = v343 + v344
								v347 = *(*int64)(unsafe.Add(mBase, uint32(v313)+2912))
								v348 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
								*(*int64)(unsafe.Add(mBase, uint32(v313)+2912)) = v347 + v348
								v351 = *(*int64)(unsafe.Add(mBase, uint32(v313)+2920))
								v352 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v313)+2920)) = v351 + v352
								v355 = *(*int64)(unsafe.Add(mBase, uint32(v313)+2928))
								v356 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v313)+2928)) = v355 + v356
								v361 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[9]))
								*(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[10])) = v361
								v365 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[5]))
								*(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[6])) = v365
								v369 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[7]))
								*(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[8])) = v369
								v373 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[2]))
								*(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[3])) = v373
							}
						}
						F_pgstat_unlock_entry(m, v59)
						mBase = m.M
						v377 = m.ExcPending
						if v377 != 0 {
							return int32(0)
						} else {
							v384 = int32(0)
							m.G0 = v25 + int32(2880)
							return v384
						}
					}
				}
			} else {
				v384 = v3
				m.G0 = v25 + int32(2880)
				return v384
			}
		} else {
			if v40&int32(1) == int32(0) {
				v384 = v3
				m.G0 = v25 + int32(2880)
				return v384
			} else {
				v58 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[4])))
				v59 = F_pgstat_get_entry_ref_locked(m, int32(6), int32(0), v58, l0)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					if v59 == int32(0) {
						v384 = int32(1)
						m.G0 = v25 + int32(2880)
						return v384
					} else {
						if l1&int32(1) == int32(0) {
						} else {
							v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[1])))
							if v71 == int32(0) {
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
								base.MemoryCopy(m, v25, int32(_a_F_pgstat_flush_backend_1), int32(2880))
								v97 = v3
								for {
									v111 = v97 * int32(320)
									v127 = int32(0)
									for {
										v142 = v127 << (uint(int32(6)) % 32)
										v143 = v111 + (v74 + int32(992)) + v142
										v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
										v145 = v142 + (v111 + (v25 + int32(960)))
										v146 = *(*int64)(unsafe.Add(mBase, uint32(v145)))
										*(*int64)(unsafe.Add(mBase, uint32(v143))) = v144 + v146
										v149 = v142 + (v111 + (v74 + int32(32)))
										v150 = *(*int64)(unsafe.Add(mBase, uint32(v149)))
										v151 = v142 + (v111 + v25)
										v152 = *(*int64)(unsafe.Add(mBase, uint32(v151)))
										*(*int64)(unsafe.Add(mBase, uint32(v149))) = v150 + v152
										v155 = v142 + (v74 + int32(1952) + v111)
										v156 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
										v157 = v142 + (v111 + (v25 + int32(1920)))
										v158 = *(*int64)(unsafe.Add(mBase, uint32(v157)))
										v159 = int64(1000)
										v160 = base.I64_div_s(v158, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155))) = v156 + v160
										v163 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
										v164 = *(*int64)(unsafe.Add(mBase, uint32(v145)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v143)+8)) = v163 + v164
										v167 = *(*int64)(unsafe.Add(mBase, uint32(v149)+8))
										v168 = *(*int64)(unsafe.Add(mBase, uint32(v151)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v149)+8)) = v167 + v168
										v171 = *(*int64)(unsafe.Add(mBase, uint32(v155)+8))
										v172 = *(*int64)(unsafe.Add(mBase, uint32(v157)+8))
										v174 = base.I64_div_s(v172, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155)+8)) = v171 + v174
										v177 = *(*int64)(unsafe.Add(mBase, uint32(v143)+16))
										v178 = *(*int64)(unsafe.Add(mBase, uint32(v145)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v143)+16)) = v177 + v178
										v181 = *(*int64)(unsafe.Add(mBase, uint32(v149)+16))
										v182 = *(*int64)(unsafe.Add(mBase, uint32(v151)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v149)+16)) = v181 + v182
										v185 = *(*int64)(unsafe.Add(mBase, uint32(v155)+16))
										v186 = *(*int64)(unsafe.Add(mBase, uint32(v157)+16))
										v188 = base.I64_div_s(v186, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155)+16)) = v185 + v188
										v191 = *(*int64)(unsafe.Add(mBase, uint32(v143)+24))
										v192 = *(*int64)(unsafe.Add(mBase, uint32(v145)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v143)+24)) = v191 + v192
										v195 = *(*int64)(unsafe.Add(mBase, uint32(v149)+24))
										v196 = *(*int64)(unsafe.Add(mBase, uint32(v151)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v149)+24)) = v195 + v196
										v199 = *(*int64)(unsafe.Add(mBase, uint32(v155)+24))
										v200 = *(*int64)(unsafe.Add(mBase, uint32(v157)+24))
										v202 = base.I64_div_s(v200, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155)+24)) = v199 + v202
										v205 = *(*int64)(unsafe.Add(mBase, uint32(v143)+32))
										v206 = *(*int64)(unsafe.Add(mBase, uint32(v145)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v143)+32)) = v205 + v206
										v209 = *(*int64)(unsafe.Add(mBase, uint32(v149)+32))
										v210 = *(*int64)(unsafe.Add(mBase, uint32(v151)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v149)+32)) = v209 + v210
										v213 = *(*int64)(unsafe.Add(mBase, uint32(v155)+32))
										v214 = *(*int64)(unsafe.Add(mBase, uint32(v157)+32))
										v216 = base.I64_div_s(v214, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155)+32)) = v213 + v216
										v219 = *(*int64)(unsafe.Add(mBase, uint32(v143)+40))
										v220 = *(*int64)(unsafe.Add(mBase, uint32(v145)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v143)+40)) = v219 + v220
										v223 = *(*int64)(unsafe.Add(mBase, uint32(v149)+40))
										v224 = *(*int64)(unsafe.Add(mBase, uint32(v151)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v149)+40)) = v223 + v224
										v227 = *(*int64)(unsafe.Add(mBase, uint32(v155)+40))
										v228 = *(*int64)(unsafe.Add(mBase, uint32(v157)+40))
										v230 = base.I64_div_s(v228, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155)+40)) = v227 + v230
										v233 = *(*int64)(unsafe.Add(mBase, uint32(v143)+48))
										v234 = *(*int64)(unsafe.Add(mBase, uint32(v145)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v143)+48)) = v233 + v234
										v237 = *(*int64)(unsafe.Add(mBase, uint32(v149)+48))
										v238 = *(*int64)(unsafe.Add(mBase, uint32(v151)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v149)+48)) = v237 + v238
										v241 = *(*int64)(unsafe.Add(mBase, uint32(v155)+48))
										v242 = *(*int64)(unsafe.Add(mBase, uint32(v157)+48))
										v244 = base.I64_div_s(v242, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155)+48)) = v241 + v244
										v247 = *(*int64)(unsafe.Add(mBase, uint32(v143)+56))
										v248 = *(*int64)(unsafe.Add(mBase, uint32(v145)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v143)+56)) = v247 + v248
										v251 = *(*int64)(unsafe.Add(mBase, uint32(v149)+56))
										v252 = *(*int64)(unsafe.Add(mBase, uint32(v151)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v149)+56)) = v251 + v252
										v255 = *(*int64)(unsafe.Add(mBase, uint32(v155)+56))
										v256 = *(*int64)(unsafe.Add(mBase, uint32(v157)+56))
										v258 = base.I64_div_s(v256, v159)
										*(*int64)(unsafe.Add(mBase, uint32(v155)+56)) = v255 + v258
										v262 = v127 + int32(1)
										if v262 != int32(5) {
											v127 = v262
											continue
										} else {
											break
										}
										break
									}
									v266 = v97 + int32(1)
									if v266 != int32(3) {
										v97 = v266
										continue
									} else {
										break
									}
									break
								}
								v270 = int32(0)
								base.MemoryFill(m, int32(_a_F_pgstat_flush_backend_1), v270, int32(2880))
								*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[1])) = uint8(v270)
							}
						}
						if v42 == int32(0) {
						} else {
							v300 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v300
							*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v300
							*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v300
							*(*int64)(unsafe.Add(mBase, uint32(v25))) = v300
							v309 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[2]))
							v311 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[3]))
							if v309 == v311 {
							} else {
								v313 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
								v314 = int32(_a_F_pgstat_flush_backend_2)
								v315 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
								v317 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[5]))
								v318 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[6]))
								*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v315 + (v317 - v318)
								v322 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
								v324 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[2]))
								v325 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[3]))
								*(*int64)(unsafe.Add(mBase, uint32(v25))) = v322 + (v324 - v325)
								v329 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
								v331 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[7]))
								v332 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[8]))
								*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v329 + (v331 - v332)
								v336 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
								v338 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[9]))
								v339 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[10]))
								*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v336 + (v338 - v339)
								v343 = *(*int64)(unsafe.Add(mBase, uint32(v313)+2936))
								v344 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v313)+2936)) = v343 + v344
								v347 = *(*int64)(unsafe.Add(mBase, uint32(v313)+2912))
								v348 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
								*(*int64)(unsafe.Add(mBase, uint32(v313)+2912)) = v347 + v348
								v351 = *(*int64)(unsafe.Add(mBase, uint32(v313)+2920))
								v352 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v313)+2920)) = v351 + v352
								v355 = *(*int64)(unsafe.Add(mBase, uint32(v313)+2928))
								v356 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v313)+2928)) = v355 + v356
								v361 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[9]))
								*(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[10])) = v361
								v365 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[5]))
								*(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[6])) = v365
								v369 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[7]))
								*(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[8])) = v369
								v373 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[2]))
								*(*int64)(unsafe.Add(mBase, _c_F_pgstat_flush_backend[3])) = v373
							}
						}
						F_pgstat_unlock_entry(m, v59)
						mBase = m.M
						v377 = m.ExcPending
						if v377 != 0 {
							return int32(0)
						} else {
							v384 = int32(0)
							m.G0 = v25 + int32(2880)
							return v384
						}
					}
				}
			}
		}
	}
}
func F_pgstat_hash_hash_key(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int64
	_ = v17
	var v22 int64
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v59 int64
	_ = v59
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v108 int64
	_ = v108
	var v113 int64
	_ = v113
	var v121 int64
	_ = v121
	var v126 int64
	_ = v126
	v8 = base.I64_extend_i32_u(l1) * int64(-8645972361240307355)
	if base.Ui32(l1) < base.Ui32(int32(8)) {
		v72 = l0
		v74 = l1
		v76 = v8
	} else {
		v11 = int32(8)
		v12 = l1 - v11
		if v12&v11 == int32(0) {
			v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			v22 = (int64(base.Ui64(v17)>>(uint(int64(23))%64)) ^ v17) * int64(2388976653695081527)
			v31 = l0 + int32(8)
			v32 = v12
			v34 = (v8 ^ int64(base.Ui64(v22)>>(uint(int64(47))%64)) ^ v22) * int64(-8645972361240307355)
		} else {
			v31 = l0
			v32 = l1
			v34 = v8
		}
		if base.Ui32(v12) < base.Ui32(int32(8)) {
			v72 = v31
			v74 = v12
			v76 = v34
		} else {
			v37 = v31
			v39 = v32
			v41 = v34
			for {
				v42 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
				v43 = int64(23)
				v46 = int64(2388976653695081527)
				v47 = (int64(base.Ui64(v42)>>(uint(v43)%64)) ^ v42) * v46
				v48 = int64(47)
				v52 = int64(-8645972361240307355)
				v54 = *(*int64)(unsafe.Add(mBase, uint32(v37)+8))
				v59 = (int64(base.Ui64(v54)>>(uint(v43)%64)) ^ v54) * v46
				v65 = ((v41^int64(base.Ui64(v47)>>(uint(v48)%64))^v47)*v52 ^ int64(base.Ui64(v59)>>(uint(v48)%64)) ^ v59) * v52
				v66 = int32(16)
				v67 = v37 + v66
				v69 = v39 - v66
				if base.Ui32(int32(7)) < base.Ui32(v69) {
					v37 = v67
					v39 = v69
					v41 = v65
					continue
				} else {
					break
				}
				break
			}
			v72 = v67
			v74 = v69
			v76 = v65
		}
	}
	v77 = int64(0)
	switch v74 - int32(1) {
	case 0:
		v104 = v77
		v105 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72))))
		v108 = v104 | v105
		v113 = (v108 ^ int64(base.Ui64(v108)>>(uint(int64(23))%64))) * int64(2388976653695081527)
		v121 = (v76 ^ int64(base.Ui64(v113)>>(uint(int64(47))%64)) ^ v113) * int64(-8645972361240307355)
	case 1:
		v99 = v77
		v100 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72))))
		v108 = v104 | v105
		v113 = (v108 ^ int64(base.Ui64(v108)>>(uint(int64(23))%64))) * int64(2388976653695081527)
		v121 = (v76 ^ int64(base.Ui64(v113)>>(uint(int64(47))%64)) ^ v113) * int64(-8645972361240307355)
	case 2:
		v96 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+2)))
		v99 = v96 << (uint(int64(16)) % 64)
		v100 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72))))
		v108 = v104 | v105
		v113 = (v108 ^ int64(base.Ui64(v108)>>(uint(int64(23))%64))) * int64(2388976653695081527)
		v121 = (v76 ^ int64(base.Ui64(v113)>>(uint(int64(47))%64)) ^ v113) * int64(-8645972361240307355)
	case 3:
		v93 = v77
		v94 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
		v108 = v93 | v94
		v113 = (v108 ^ int64(base.Ui64(v108)>>(uint(int64(23))%64))) * int64(2388976653695081527)
		v121 = (v76 ^ int64(base.Ui64(v113)>>(uint(int64(47))%64)) ^ v113) * int64(-8645972361240307355)
	case 4:
		v88 = v77
		v89 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+4)))
		v93 = v89<<(uint(int64(32))%64) | v88
		v94 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
		v108 = v93 | v94
		v113 = (v108 ^ int64(base.Ui64(v108)>>(uint(int64(23))%64))) * int64(2388976653695081527)
		v121 = (v76 ^ int64(base.Ui64(v113)>>(uint(int64(47))%64)) ^ v113) * int64(-8645972361240307355)
	case 5:
		v83 = v77
		v84 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+5)))
		v88 = v84<<(uint(int64(40))%64) | v83
		v89 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+4)))
		v93 = v89<<(uint(int64(32))%64) | v88
		v94 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
		v108 = v93 | v94
		v113 = (v108 ^ int64(base.Ui64(v108)>>(uint(int64(23))%64))) * int64(2388976653695081527)
		v121 = (v76 ^ int64(base.Ui64(v113)>>(uint(int64(47))%64)) ^ v113) * int64(-8645972361240307355)
	case 6:
		v80 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+6)))
		v83 = v80 << (uint(int64(48)) % 64)
		v84 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+5)))
		v88 = v84<<(uint(int64(40))%64) | v83
		v89 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+4)))
		v93 = v89<<(uint(int64(32))%64) | v88
		v94 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
		v108 = v93 | v94
		v113 = (v108 ^ int64(base.Ui64(v108)>>(uint(int64(23))%64))) * int64(2388976653695081527)
		v121 = (v76 ^ int64(base.Ui64(v113)>>(uint(int64(47))%64)) ^ v113) * int64(-8645972361240307355)
	default:
		v121 = v76
	}
	v126 = (int64(base.Ui64(v121)>>(uint(int64(23))%64)) ^ v121) * int64(2388976653695081527)
	return base.I32_wrap_i64(int64(base.Ui64(v126)>>(uint(int64(47))%64)) ^ v126 - int64(base.Ui64(v126)>>(uint(int64(32))%64)))
}
func F_pgstat_init_function_usage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_init_function_usage[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
	if v10 <= v12 {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
		m.G0 = v7 + int32(32)
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_init_function_usage[1]))
		v19 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)))
		v22 = F_pgstat_prep_pending_entry(m, int32(3), v18, v19, v7+int32(15))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
			if v24 == int32(1) {
				F_ReceiveSharedInvalidMessages(m)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					v32 = int32(0)
					v35 = F_SearchSysCacheExists(m, int32(47), v31, v32, v32, v32)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						if v35 == int32(0) {
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_init_function_usage[1]))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v65 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v64)+4)))
							v66 = F_pgstat_drop_entry(m, int32(3), v63, v65)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									F_errcode(m, int32(52461700))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_pgstat_init_function_usage_0), int32(0))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_pgstat_init_function_usage_1), int32(118), int32(_a_F_pgstat_init_function_usage_2))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
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
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v39
							v41 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v41
							v44 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_init_function_usage[2]))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v44
							F___clock_gettime(m, int32(1), v7+int32(16))
							mBase = m.M
							v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7)+24)))
							v51 = *(*int64)(unsafe.Add(mBase, uint32(v7)+16))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v50 + v51*int64(1000000000)
							m.G0 = v7 + int32(32)
							return
						}
					}
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v39
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v41
				v44 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_init_function_usage[2]))
				*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v44
				F___clock_gettime(m, int32(1), v7+int32(16))
				mBase = m.M
				v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7)+24)))
				v51 = *(*int64)(unsafe.Add(mBase, uint32(v7)+16))
				*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v50 + v51*int64(1000000000)
				m.G0 = v7 + int32(32)
				return
			}
		}
	}
}
func F_pgstat_init_relation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+119)))
	switch v4 - int32(83) {
	case 0, 22, 26, 29, 31, 33:
		v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_init_relation[0])))
		if v8 == int32(0) {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
			if v11 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = int32(0)
			} else {
			}
			v17 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v17
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)) = uint8(v17)
			return
		} else {
			v14 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)) = uint8(v14)
			return
		}
	default:
		v17 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v17
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)) = uint8(v17)
		return
	}
}
func F_pgstat_io_flush_cb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v164 int64
	_ = v164
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v206 int64
	_ = v206
	var v209 int64
	_ = v209
	var v210 int64
	_ = v210
	var v213 int64
	_ = v213
	var v214 int64
	_ = v214
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v220 int64
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	v2 = int32(0)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_io_flush_cb[0])))
	if v20 == v2 {
		v242 = v2
		return v242
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_io_flush_cb[1]))
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_io_flush_cb[2]))
		v31 = v24 + v26<<(uint(int32(4))%32) + int32(608)
		if l0 == int32(0) {
			v35 = F_LWLockAcquire(m, v31, int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v48 = v24 + v26*int32(2880)
				v64 = v2
				for {
					v74 = v64 * int32(320)
					v92 = int32(0)
					for {
						v104 = v92 << (uint(int32(6)) % 32)
						v105 = v74 + (v48 + int32(1864)) + v104
						v106 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
						v107 = v104 + (v74 + int32(_a_F_pgstat_io_flush_cb_0))
						v108 = *(*int64)(unsafe.Add(mBase, uint32(v107)))
						*(*int64)(unsafe.Add(mBase, uint32(v105))) = v106 + v108
						v111 = v104 + (v74 + (v48 + int32(904)))
						v112 = *(*int64)(unsafe.Add(mBase, uint32(v111)))
						v113 = v104 + (v74 + int32(_a_F_pgstat_io_flush_cb_1))
						v114 = *(*int64)(unsafe.Add(mBase, uint32(v113)))
						*(*int64)(unsafe.Add(mBase, uint32(v111))) = v112 + v114
						v117 = v104 + (v48 + int32(2824) + v74)
						v118 = *(*int64)(unsafe.Add(mBase, uint32(v117)))
						v119 = v104 + (v74 + int32(_a_F_pgstat_io_flush_cb_2))
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v119)))
						v121 = int64(1000)
						v122 = base.I64_div_s(v120, v121)
						*(*int64)(unsafe.Add(mBase, uint32(v117))) = v118 + v122
						v125 = *(*int64)(unsafe.Add(mBase, uint32(v105)+8))
						v126 = *(*int64)(unsafe.Add(mBase, uint32(v107)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v105)+8)) = v125 + v126
						v129 = *(*int64)(unsafe.Add(mBase, uint32(v111)+8))
						v130 = *(*int64)(unsafe.Add(mBase, uint32(v113)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v111)+8)) = v129 + v130
						v133 = *(*int64)(unsafe.Add(mBase, uint32(v117)+8))
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v119)+8))
						v136 = base.I64_div_s(v134, v121)
						*(*int64)(unsafe.Add(mBase, uint32(v117)+8)) = v133 + v136
						v139 = *(*int64)(unsafe.Add(mBase, uint32(v105)+16))
						v140 = *(*int64)(unsafe.Add(mBase, uint32(v107)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v105)+16)) = v139 + v140
						v143 = *(*int64)(unsafe.Add(mBase, uint32(v111)+16))
						v144 = *(*int64)(unsafe.Add(mBase, uint32(v113)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v111)+16)) = v143 + v144
						v147 = *(*int64)(unsafe.Add(mBase, uint32(v117)+16))
						v148 = *(*int64)(unsafe.Add(mBase, uint32(v119)+16))
						v150 = base.I64_div_s(v148, v121)
						*(*int64)(unsafe.Add(mBase, uint32(v117)+16)) = v147 + v150
						v153 = *(*int64)(unsafe.Add(mBase, uint32(v105)+24))
						v154 = *(*int64)(unsafe.Add(mBase, uint32(v107)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v105)+24)) = v153 + v154
						v157 = *(*int64)(unsafe.Add(mBase, uint32(v111)+24))
						v158 = *(*int64)(unsafe.Add(mBase, uint32(v113)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v111)+24)) = v157 + v158
						v161 = *(*int64)(unsafe.Add(mBase, uint32(v117)+24))
						v162 = *(*int64)(unsafe.Add(mBase, uint32(v119)+24))
						v164 = base.I64_div_s(v162, v121)
						*(*int64)(unsafe.Add(mBase, uint32(v117)+24)) = v161 + v164
						v167 = *(*int64)(unsafe.Add(mBase, uint32(v105)+32))
						v168 = *(*int64)(unsafe.Add(mBase, uint32(v107)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v105)+32)) = v167 + v168
						v171 = *(*int64)(unsafe.Add(mBase, uint32(v111)+32))
						v172 = *(*int64)(unsafe.Add(mBase, uint32(v113)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v111)+32)) = v171 + v172
						v175 = *(*int64)(unsafe.Add(mBase, uint32(v117)+32))
						v176 = *(*int64)(unsafe.Add(mBase, uint32(v119)+32))
						v178 = base.I64_div_s(v176, v121)
						*(*int64)(unsafe.Add(mBase, uint32(v117)+32)) = v175 + v178
						v181 = *(*int64)(unsafe.Add(mBase, uint32(v105)+40))
						v182 = *(*int64)(unsafe.Add(mBase, uint32(v107)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v105)+40)) = v181 + v182
						v185 = *(*int64)(unsafe.Add(mBase, uint32(v111)+40))
						v186 = *(*int64)(unsafe.Add(mBase, uint32(v113)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v111)+40)) = v185 + v186
						v189 = *(*int64)(unsafe.Add(mBase, uint32(v117)+40))
						v190 = *(*int64)(unsafe.Add(mBase, uint32(v119)+40))
						v192 = base.I64_div_s(v190, v121)
						*(*int64)(unsafe.Add(mBase, uint32(v117)+40)) = v189 + v192
						v195 = *(*int64)(unsafe.Add(mBase, uint32(v105)+48))
						v196 = *(*int64)(unsafe.Add(mBase, uint32(v107)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v105)+48)) = v195 + v196
						v199 = *(*int64)(unsafe.Add(mBase, uint32(v111)+48))
						v200 = *(*int64)(unsafe.Add(mBase, uint32(v113)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v111)+48)) = v199 + v200
						v203 = *(*int64)(unsafe.Add(mBase, uint32(v117)+48))
						v204 = *(*int64)(unsafe.Add(mBase, uint32(v119)+48))
						v206 = base.I64_div_s(v204, v121)
						*(*int64)(unsafe.Add(mBase, uint32(v117)+48)) = v203 + v206
						v209 = *(*int64)(unsafe.Add(mBase, uint32(v105)+56))
						v210 = *(*int64)(unsafe.Add(mBase, uint32(v107)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v105)+56)) = v209 + v210
						v213 = *(*int64)(unsafe.Add(mBase, uint32(v111)+56))
						v214 = *(*int64)(unsafe.Add(mBase, uint32(v113)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v111)+56)) = v213 + v214
						v217 = *(*int64)(unsafe.Add(mBase, uint32(v117)+56))
						v218 = *(*int64)(unsafe.Add(mBase, uint32(v119)+56))
						v220 = base.I64_div_s(v218, v121)
						*(*int64)(unsafe.Add(mBase, uint32(v117)+56)) = v217 + v220
						v224 = v92 + int32(1)
						if v224 != int32(5) {
							v92 = v224
							continue
						} else {
							break
						}
						break
					}
					v228 = v64 + int32(1)
					if v228 != int32(3) {
						v64 = v228
						continue
					} else {
						break
					}
					break
				}
				F_LWLockRelease(m, v31)
				mBase = m.M
				v232 = m.ExcPending
				if v232 != 0 {
					return int32(0)
				} else {
					v233 = int32(0)
					base.MemoryFill(m, int32(_a_F_pgstat_io_flush_cb_1), v233, int32(2880))
					*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_io_flush_cb[0])) = uint8(v233)
					v242 = v233
					return v242
				}
			}
		} else {
			v41 = F_LWLockConditionalAcquire(m, v31, int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				if v41 == int32(0) {
					v242 = int32(1)
					return v242
				} else {
					v48 = v24 + v26*int32(2880)
					v64 = v2
					for {
						v74 = v64 * int32(320)
						v92 = int32(0)
						for {
							v104 = v92 << (uint(int32(6)) % 32)
							v105 = v74 + (v48 + int32(1864)) + v104
							v106 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
							v107 = v104 + (v74 + int32(_a_F_pgstat_io_flush_cb_0))
							v108 = *(*int64)(unsafe.Add(mBase, uint32(v107)))
							*(*int64)(unsafe.Add(mBase, uint32(v105))) = v106 + v108
							v111 = v104 + (v74 + (v48 + int32(904)))
							v112 = *(*int64)(unsafe.Add(mBase, uint32(v111)))
							v113 = v104 + (v74 + int32(_a_F_pgstat_io_flush_cb_1))
							v114 = *(*int64)(unsafe.Add(mBase, uint32(v113)))
							*(*int64)(unsafe.Add(mBase, uint32(v111))) = v112 + v114
							v117 = v104 + (v48 + int32(2824) + v74)
							v118 = *(*int64)(unsafe.Add(mBase, uint32(v117)))
							v119 = v104 + (v74 + int32(_a_F_pgstat_io_flush_cb_2))
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v119)))
							v121 = int64(1000)
							v122 = base.I64_div_s(v120, v121)
							*(*int64)(unsafe.Add(mBase, uint32(v117))) = v118 + v122
							v125 = *(*int64)(unsafe.Add(mBase, uint32(v105)+8))
							v126 = *(*int64)(unsafe.Add(mBase, uint32(v107)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v105)+8)) = v125 + v126
							v129 = *(*int64)(unsafe.Add(mBase, uint32(v111)+8))
							v130 = *(*int64)(unsafe.Add(mBase, uint32(v113)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v111)+8)) = v129 + v130
							v133 = *(*int64)(unsafe.Add(mBase, uint32(v117)+8))
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v119)+8))
							v136 = base.I64_div_s(v134, v121)
							*(*int64)(unsafe.Add(mBase, uint32(v117)+8)) = v133 + v136
							v139 = *(*int64)(unsafe.Add(mBase, uint32(v105)+16))
							v140 = *(*int64)(unsafe.Add(mBase, uint32(v107)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v105)+16)) = v139 + v140
							v143 = *(*int64)(unsafe.Add(mBase, uint32(v111)+16))
							v144 = *(*int64)(unsafe.Add(mBase, uint32(v113)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v111)+16)) = v143 + v144
							v147 = *(*int64)(unsafe.Add(mBase, uint32(v117)+16))
							v148 = *(*int64)(unsafe.Add(mBase, uint32(v119)+16))
							v150 = base.I64_div_s(v148, v121)
							*(*int64)(unsafe.Add(mBase, uint32(v117)+16)) = v147 + v150
							v153 = *(*int64)(unsafe.Add(mBase, uint32(v105)+24))
							v154 = *(*int64)(unsafe.Add(mBase, uint32(v107)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v105)+24)) = v153 + v154
							v157 = *(*int64)(unsafe.Add(mBase, uint32(v111)+24))
							v158 = *(*int64)(unsafe.Add(mBase, uint32(v113)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v111)+24)) = v157 + v158
							v161 = *(*int64)(unsafe.Add(mBase, uint32(v117)+24))
							v162 = *(*int64)(unsafe.Add(mBase, uint32(v119)+24))
							v164 = base.I64_div_s(v162, v121)
							*(*int64)(unsafe.Add(mBase, uint32(v117)+24)) = v161 + v164
							v167 = *(*int64)(unsafe.Add(mBase, uint32(v105)+32))
							v168 = *(*int64)(unsafe.Add(mBase, uint32(v107)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v105)+32)) = v167 + v168
							v171 = *(*int64)(unsafe.Add(mBase, uint32(v111)+32))
							v172 = *(*int64)(unsafe.Add(mBase, uint32(v113)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v111)+32)) = v171 + v172
							v175 = *(*int64)(unsafe.Add(mBase, uint32(v117)+32))
							v176 = *(*int64)(unsafe.Add(mBase, uint32(v119)+32))
							v178 = base.I64_div_s(v176, v121)
							*(*int64)(unsafe.Add(mBase, uint32(v117)+32)) = v175 + v178
							v181 = *(*int64)(unsafe.Add(mBase, uint32(v105)+40))
							v182 = *(*int64)(unsafe.Add(mBase, uint32(v107)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v105)+40)) = v181 + v182
							v185 = *(*int64)(unsafe.Add(mBase, uint32(v111)+40))
							v186 = *(*int64)(unsafe.Add(mBase, uint32(v113)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v111)+40)) = v185 + v186
							v189 = *(*int64)(unsafe.Add(mBase, uint32(v117)+40))
							v190 = *(*int64)(unsafe.Add(mBase, uint32(v119)+40))
							v192 = base.I64_div_s(v190, v121)
							*(*int64)(unsafe.Add(mBase, uint32(v117)+40)) = v189 + v192
							v195 = *(*int64)(unsafe.Add(mBase, uint32(v105)+48))
							v196 = *(*int64)(unsafe.Add(mBase, uint32(v107)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v105)+48)) = v195 + v196
							v199 = *(*int64)(unsafe.Add(mBase, uint32(v111)+48))
							v200 = *(*int64)(unsafe.Add(mBase, uint32(v113)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v111)+48)) = v199 + v200
							v203 = *(*int64)(unsafe.Add(mBase, uint32(v117)+48))
							v204 = *(*int64)(unsafe.Add(mBase, uint32(v119)+48))
							v206 = base.I64_div_s(v204, v121)
							*(*int64)(unsafe.Add(mBase, uint32(v117)+48)) = v203 + v206
							v209 = *(*int64)(unsafe.Add(mBase, uint32(v105)+56))
							v210 = *(*int64)(unsafe.Add(mBase, uint32(v107)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v105)+56)) = v209 + v210
							v213 = *(*int64)(unsafe.Add(mBase, uint32(v111)+56))
							v214 = *(*int64)(unsafe.Add(mBase, uint32(v113)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v111)+56)) = v213 + v214
							v217 = *(*int64)(unsafe.Add(mBase, uint32(v117)+56))
							v218 = *(*int64)(unsafe.Add(mBase, uint32(v119)+56))
							v220 = base.I64_div_s(v218, v121)
							*(*int64)(unsafe.Add(mBase, uint32(v117)+56)) = v217 + v220
							v224 = v92 + int32(1)
							if v224 != int32(5) {
								v92 = v224
								continue
							} else {
								break
							}
							break
						}
						v228 = v64 + int32(1)
						if v228 != int32(3) {
							v64 = v228
							continue
						} else {
							break
						}
						break
					}
					F_LWLockRelease(m, v31)
					mBase = m.M
					v232 = m.ExcPending
					if v232 != 0 {
						return int32(0)
					} else {
						v233 = int32(0)
						base.MemoryFill(m, int32(_a_F_pgstat_io_flush_cb_1), v233, int32(2880))
						*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_io_flush_cb[0])) = uint8(v233)
						v242 = v233
						return v242
					}
				}
			}
		}
	}
}
func F_pgstat_progress_end_command(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_end_command[0]))
	if v3 == int32(0) {
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_progress_end_command[1])))
		if v7&int32(1) == int32(0) {
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+220))
			if v12 == int32(0) {
			} else {
				v15 = int32(_a_F_pgstat_progress_end_command_0)
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_end_command[2]))
				v18 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_end_command[2])) = v17 + v18
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
				*(*int32)(unsafe.Add(mBase, uint32(v3))) = v21 + v18
				v25 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v3)+220)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v3)+224)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v3))) = v21 + int32(2)
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_end_command[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_end_command[2])) = v35 - v18
			}
		}
	}
	return
}
func F_pgstat_relation_flush_cb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v176 int64
	_ = v176
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v207 int64
	_ = v207
	var v208 int64
	_ = v208
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v215 int32
	_ = v215
	var v218 int64
	_ = v218
	var v219 int64
	_ = v219
	var v220 int64
	_ = v220
	var v221 int64
	_ = v221
	var v224 int32
	_ = v224
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v241 int64
	_ = v241
	var v244 int64
	_ = v244
	var v247 int64
	_ = v247
	var v248 int64
	_ = v248
	var v251 int64
	_ = v251
	var v252 int64
	_ = v252
	var v255 int64
	_ = v255
	var v257 int64
	_ = v257
	var v260 int64
	_ = v260
	var v262 int64
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int64
	_ = v273
	var v274 int64
	_ = v274
	var v277 int64
	_ = v277
	var v278 int64
	_ = v278
	var v281 int64
	_ = v281
	var v282 int64
	_ = v282
	var v285 int64
	_ = v285
	var v286 int64
	_ = v286
	var v289 int64
	_ = v289
	var v290 int64
	_ = v290
	var v293 int64
	_ = v293
	var v294 int64
	_ = v294
	var v297 int64
	_ = v297
	var v298 int64
	_ = v298
	var v306 int32
	_ = v306
	v11 = int64(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = v14 + int32(16)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v23 = (int32(-16) - v14) & int32(3)
	if v23 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v306
L2:
	;
	v161 = F_pgstat_lock_entry(m, l0, l1)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L31
	} else {
		goto L32
	}
L3:
	;
	v36 = v23 + v16
	v40 = (v14 + int32(128)) & int32(-4)
	v42 = v40 - int32(28)
	if base.Ui32(v36) < base.Ui32(v42) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v26 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v23 == int32(1) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+17)))
	if v29 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v23 == int32(2) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+18)))
	if v32|base.B2i32(v23 != int32(3)) != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	goto L3
L10:
	;
	v48 = v36
	v49 = v23
	goto L13
L11:
	;
	v81 = v23
	goto L12
L12:
	;
	v89 = v81 + v16
	if base.Ui32(v89) < base.Ui32(v40) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v48)+28))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v57|(v58|(v59|(v60|(v61|(v62|(v63|v64)))))) != 0 {
		goto L2
	} else {
		goto L15
	}
L14:
	;
	v81 = v73
	goto L12
L15:
	;
	v73 = v49 + int32(32)
	v74 = v16 + v73
	if base.Ui32(v74) < base.Ui32(v42) {
		v48 = v74
		v49 = v73
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v95 = v89
	v96 = v81
	goto L20
L18:
	;
	v114 = v81
	goto L19
L19:
	;
	v122 = int32(112)
	if base.Ui32(v114) <= base.Ui32(v122) {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v104 != 0 {
		goto L2
	} else {
		goto L22
	}
L21:
	;
	v114 = v106
	goto L19
L22:
	;
	v106 = v96 + int32(4)
	v107 = v16 + v106
	if base.Ui32(v107) < base.Ui32(v40) {
		v95 = v107
		v96 = v106
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v125 = v122
	goto L26
L25:
	;
	v125 = v114
	goto L26
L26:
	;
	v132 = v114
	goto L27
L27:
	;
	if v132 == v125 {
		v306 = int32(1)
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L2
L29:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132+v16))))
	if v144 == int32(0) {
		v132 = v132 + int32(1)
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	return int32(0)
L32:
	;
	if v161 == int32(0) {
		v306 = int32(0)
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v167 + v168
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	if v171 == int64(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v187 + v188
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v191 + v192
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v17)+56))
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v195 + v196
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v17)+64))
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v199 + v200
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v17)+72))
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+72)) = v203 + v204
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v17)+80))
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v14)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = v207 + v208
	v211 = *(*int64)(unsafe.Add(mBase, uint32(v17)+88))
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v14)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+88)) = v211 + v212
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+80)))
	if v215 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	v176 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_relation_flush_cb[0]))
	if v176 == int64(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
	if v182 <= v183 {
		goto L34
	} else {
		goto L40
	}
L37:
	;
	v180 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_relation_flush_cb[0])) = v180
	v182 = v180
	goto L39
L38:
	;
	v182 = v176
	goto L39
L39:
	;
	goto L36
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = v182
	goto L34
L41:
	;
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v14)+88))
	v235 = v233 + v234
	*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = v235
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v14)+96))
	v238 = v231 + v237
	*(*int64)(unsafe.Add(mBase, uint32(v17)+104)) = v238
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v17)+112))
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v14)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v240 + v241
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v232 + v244
	v247 = *(*int64)(unsafe.Add(mBase, uint32(v17)+128))
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v14)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = v247 + v248
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v14)+120))
	v252 = int64(0)
	if v252 < v238 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v17)+120))
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v17)+104))
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v17)+96))
	v231 = v219
	v232 = v218
	v233 = v220
	goto L41
L43:
	;
	goto L44
L44:
	;
	v221 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v221
	v224 = v17 + int32(96)
	*(*int64)(unsafe.Add(mBase, uint32(v224)+8)) = v221
	*(*int64)(unsafe.Add(mBase, uint32(v224))) = v221
	v231 = v11
	v232 = v11
	v233 = v221
	goto L41
L45:
	;
	v255 = v238
	goto L47
L46:
	;
	v255 = v252
	goto L47
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+104)) = v255
	v257 = int64(0)
	if v257 < v235 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v260 = v235
	goto L50
L49:
	;
	v260 = v257
	goto L50
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = v260
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v17)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+136)) = v251 + v262
	F_pgstat_unlock_entry(m, l0)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L31
	} else {
		goto L51
	}
L51:
	;
	v270 = F_pgstat_prep_pending_entry(m, int32(1), v19, int64(0), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L31
	} else {
		goto L52
	}
L52:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v272)+32))
	v274 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v272)+32)) = v273 + v274
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v272)+40))
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v272)+40)) = v277 + v278
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v272)+48))
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v272)+48)) = v281 + v282
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v272)+56))
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v272)+56)) = v285 + v286
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v272)+64))
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v272)+64)) = v289 + v290
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v272)+16))
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v14)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v272)+16)) = v293 + v294
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v272)+24))
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v14)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v272)+24)) = v297 + v298
	v306 = int32(1)
	goto L1
}
func F_pgstat_report_analyze(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int64) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v95 int64
	_ = v95
	var v101 int64
	_ = v101
	var v110 int64
	_ = v110
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_report_analyze[0])))
	if v11 == int32(1) {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_analyze[1]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+117)))
		if v18 != 0 {
			v19 = int32(0)
		} else {
			v19 = v16
		}
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
		if v20 != 0 {
			v27 = v17
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+119)))
			if v28 == int32(112) {
				v70 = l1
				v71 = l2
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
				if v32 != 0 {
					v34 = l1
					v35 = l2
					v38 = v32
					for {
						v42 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
						v43 = *(*int64)(unsafe.Add(mBase, uint32(v38)+8))
						v45 = v35 - (v42 + v43)
						v46 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
						v48 = v34 - v46 + v42
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
						if v49 != 0 {
							v34 = v48
							v35 = v45
							v38 = v49
							continue
						} else {
							break
						}
						break
					}
					v51 = v48
					v52 = v45
				} else {
					v51 = l1
					v52 = l2
				}
				v59 = *(*int64)(unsafe.Add(mBase, uint32(v31)+96))
				v60 = v52 - v59
				v61 = int64(0)
				if v61 < v60 {
					v64 = v60
				} else {
					v64 = v61
				}
				v65 = int64(0)
				if v65 < v51 {
					v68 = v51
				} else {
					v68 = v65
				}
				v70 = v68
				v71 = v64
			}
			v81 = m.G0
			v82 = int32(16)
			v83 = v81 - v82
			m.G0 = v83
			F_gettimeofday(m, v83)
			mBase = m.M
			v86 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
			v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
			m.G0 = v83 + v82
			v95 = v87 + v86*int64(1000000) - int64(946684800000000)
			if v95 <= l4 {
				v113 = int32(0)
			} else {
				v101 = v95 - l4
				if base.B2i32(int64(0) < l4)^base.B2i32(v101 < v95)|base.B2i32(int64(2147483646000) < v101) != 0 {
					v113 = int32(2147483647)
				} else {
					v110 = base.I64_div_s(v101+int64(999), int64(1000))
					v113 = base.I32_wrap_i64(v110)
				}
			}
			v115 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
			v117 = F_pgstat_get_entry_ref_locked(m, int32(2), v19, v115, int32(0))
			mBase = m.M
			v118 = m.ExcPending
			if v118 != 0 {
				return
			} else {
				v119 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
				*(*int64)(unsafe.Add(mBase, uint32(v119)+104)) = v71
				*(*int64)(unsafe.Add(mBase, uint32(v119)+96)) = v70
				if l3 != 0 {
					*(*int64)(unsafe.Add(mBase, uint32(v119)+112)) = int64(0)
				} else {
				}
				v127 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_analyze[2]))
				v129 = base.B2i32(v127 == int32(4))
				if v127 == int32(4) {
					v130 = int32(192)
				} else {
					v130 = int32(176)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v119+v130))) = v95
				if v127 == int32(4) {
					v135 = int32(200)
				} else {
					v135 = int32(184)
				}
				v136 = v119 + v135
				v137 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
				*(*int64)(unsafe.Add(mBase, uint32(v136))) = v137 + int64(1)
				if v127 == int32(4) {
					v143 = int32(232)
				} else {
					v143 = int32(224)
				}
				v144 = v119 + v143
				v145 = *(*int64)(unsafe.Add(mBase, uint32(v144)))
				*(*int64)(unsafe.Add(mBase, uint32(v144))) = v145 + base.I64_extend_i32_s(v113)
				F_pgstat_unlock_entry(m, v117)
				mBase = m.M
				v150 = m.ExcPending
				if v150 != 0 {
					return
				} else {
					F_pgstat_flush_io(m, int32(0))
					mBase = m.M
					v153 = m.ExcPending
					if v153 != 0 {
						return
					} else {
						v156 = F_pgstat_flush_backend(m, int32(0), int32(1))
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
			if v21 != int32(1) {
				v70 = l1
				v71 = l2
				v81 = m.G0
				v82 = int32(16)
				v83 = v81 - v82
				m.G0 = v83
				F_gettimeofday(m, v83)
				mBase = m.M
				v86 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
				v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
				m.G0 = v83 + v82
				v95 = v87 + v86*int64(1000000) - int64(946684800000000)
				if v95 <= l4 {
					v113 = int32(0)
				} else {
					v101 = v95 - l4
					if base.B2i32(int64(0) < l4)^base.B2i32(v101 < v95)|base.B2i32(int64(2147483646000) < v101) != 0 {
						v113 = int32(2147483647)
					} else {
						v110 = base.I64_div_s(v101+int64(999), int64(1000))
						v113 = base.I32_wrap_i64(v110)
					}
				}
				v115 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
				v117 = F_pgstat_get_entry_ref_locked(m, int32(2), v19, v115, int32(0))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return
				} else {
					v119 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
					*(*int64)(unsafe.Add(mBase, uint32(v119)+104)) = v71
					*(*int64)(unsafe.Add(mBase, uint32(v119)+96)) = v70
					if l3 != 0 {
						*(*int64)(unsafe.Add(mBase, uint32(v119)+112)) = int64(0)
					} else {
					}
					v127 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_analyze[2]))
					v129 = base.B2i32(v127 == int32(4))
					if v127 == int32(4) {
						v130 = int32(192)
					} else {
						v130 = int32(176)
					}
					*(*int64)(unsafe.Add(mBase, uint32(v119+v130))) = v95
					if v127 == int32(4) {
						v135 = int32(200)
					} else {
						v135 = int32(184)
					}
					v136 = v119 + v135
					v137 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
					*(*int64)(unsafe.Add(mBase, uint32(v136))) = v137 + int64(1)
					if v127 == int32(4) {
						v143 = int32(232)
					} else {
						v143 = int32(224)
					}
					v144 = v119 + v143
					v145 = *(*int64)(unsafe.Add(mBase, uint32(v144)))
					*(*int64)(unsafe.Add(mBase, uint32(v144))) = v145 + base.I64_extend_i32_s(v113)
					F_pgstat_unlock_entry(m, v117)
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return
					} else {
						F_pgstat_flush_io(m, int32(0))
						mBase = m.M
						v153 = m.ExcPending
						if v153 != 0 {
							return
						} else {
							v156 = F_pgstat_flush_backend(m, int32(0), int32(1))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				F_pgstat_assoc_relation(m, l0)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v27 = v26
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+119)))
					if v28 == int32(112) {
						v70 = l1
						v71 = l2
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
						if v32 != 0 {
							v34 = l1
							v35 = l2
							v38 = v32
							for {
								v42 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
								v43 = *(*int64)(unsafe.Add(mBase, uint32(v38)+8))
								v45 = v35 - (v42 + v43)
								v46 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
								v48 = v34 - v46 + v42
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
								if v49 != 0 {
									v34 = v48
									v35 = v45
									v38 = v49
									continue
								} else {
									break
								}
								break
							}
							v51 = v48
							v52 = v45
						} else {
							v51 = l1
							v52 = l2
						}
						v59 = *(*int64)(unsafe.Add(mBase, uint32(v31)+96))
						v60 = v52 - v59
						v61 = int64(0)
						if v61 < v60 {
							v64 = v60
						} else {
							v64 = v61
						}
						v65 = int64(0)
						if v65 < v51 {
							v68 = v51
						} else {
							v68 = v65
						}
						v70 = v68
						v71 = v64
					}
					v81 = m.G0
					v82 = int32(16)
					v83 = v81 - v82
					m.G0 = v83
					F_gettimeofday(m, v83)
					mBase = m.M
					v86 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
					v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
					m.G0 = v83 + v82
					v95 = v87 + v86*int64(1000000) - int64(946684800000000)
					if v95 <= l4 {
						v113 = int32(0)
					} else {
						v101 = v95 - l4
						if base.B2i32(int64(0) < l4)^base.B2i32(v101 < v95)|base.B2i32(int64(2147483646000) < v101) != 0 {
							v113 = int32(2147483647)
						} else {
							v110 = base.I64_div_s(v101+int64(999), int64(1000))
							v113 = base.I32_wrap_i64(v110)
						}
					}
					v115 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
					v117 = F_pgstat_get_entry_ref_locked(m, int32(2), v19, v115, int32(0))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return
					} else {
						v119 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
						*(*int64)(unsafe.Add(mBase, uint32(v119)+104)) = v71
						*(*int64)(unsafe.Add(mBase, uint32(v119)+96)) = v70
						if l3 != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(v119)+112)) = int64(0)
						} else {
						}
						v127 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_analyze[2]))
						v129 = base.B2i32(v127 == int32(4))
						if v127 == int32(4) {
							v130 = int32(192)
						} else {
							v130 = int32(176)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v119+v130))) = v95
						if v127 == int32(4) {
							v135 = int32(200)
						} else {
							v135 = int32(184)
						}
						v136 = v119 + v135
						v137 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
						*(*int64)(unsafe.Add(mBase, uint32(v136))) = v137 + int64(1)
						if v127 == int32(4) {
							v143 = int32(232)
						} else {
							v143 = int32(224)
						}
						v144 = v119 + v143
						v145 = *(*int64)(unsafe.Add(mBase, uint32(v144)))
						*(*int64)(unsafe.Add(mBase, uint32(v144))) = v145 + base.I64_extend_i32_s(v113)
						F_pgstat_unlock_entry(m, v117)
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return
						} else {
							F_pgstat_flush_io(m, int32(0))
							mBase = m.M
							v153 = m.ExcPending
							if v153 != 0 {
								return
							} else {
								v156 = F_pgstat_flush_backend(m, int32(0), int32(1))
								mBase = m.M
								v157 = m.ExcPending
								if v157 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		return
	}
}
func F_pgstat_report_query_id(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_query_id[0]))
	if v5 == int32(0) {
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_report_query_id[1])))
		if v9&int32(1) == int32(0) {
		} else {
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v5)+392))
			if base.B2i32(l1 == int32(0))&base.B2i32(v16 != int64(0)) != 0 {
			} else {
				v20 = int32(_a_F_pgstat_report_query_id_0)
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_query_id[2]))
				v23 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_query_id[2])) = v22 + v23
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v26 + v23
				*(*int64)(unsafe.Add(mBase, uint32(v5)+392)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v26 + int32(2)
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_query_id[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_query_id[2])) = v37 - v23
			}
		}
	}
	return
}
func F_pgstat_report_stat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v62 int64
	_ = v62
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v145 int32
	_ = v145
	var v149 int64
	_ = v149
	var v156 int64
	_ = v156
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v193 int64
	_ = v193
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v373 int64
	_ = v373
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	v2 = int32(0)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_report_stat[0])))
	if v12 == v2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v388
L2:
	;
	v110 = m.G0
	v112 = v110 - int32(16)
	m.G0 = v112
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_stat[1]))
	if v115 != 0 {
		goto L28
	} else {
		goto L29
	}
L3:
	;
	v90 = m.G0
	v91 = int32(16)
	v92 = v90 - v91
	m.G0 = v92
	F_gettimeofday(m, v92)
	mBase = m.M
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
	v96 = int64(*(*int32)(unsafe.Add(mBase, uint32(v92)+8)))
	m.G0 = v92 + v91
	goto L27
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_stat[2]))
	v17 = int32(0)
	if base.B2i32(v16 != v17)&base.B2i32(v16 != int32(_a_F_pgstat_report_stat_0)) == v17 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_report_stat[0])) = uint8(v70)
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_stat[2]))
	if base.B2i32(v73 != v70)&base.B2i32(v73 != int32(_a_F_pgstat_report_stat_0)) != 0 {
		goto L3
	} else {
		goto L25
	}
L7:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_report_stat[3])))
	if v25&int32(1) == int32(0) {
		v388 = v2
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if l0 != 0 {
		goto L3
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	v32 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[4]))
	if v32 == int64(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v40 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[5]))
	if int64(0) < v40 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v36 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[4])) = v36
	v38 = v36
	goto L15
L14:
	;
	v38 = v32
	goto L15
L15:
	;
	goto L12
L16:
	;
	goto L19
L17:
	;
	goto L18
L18:
	;
	v49 = int32(1)
	v51 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[6]))
	if v51 <= int64(0) {
		v107 = v49
		v108 = v38
		goto L2
	} else {
		goto L21
	}
L19:
	;
	if base.I64_extend_i32_s(int32(_a_F_pgstat_report_stat_1))*int64(1000) <= v38-v40 {
		v107 = v2
		v108 = v38
		goto L2
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L22
L22:
	;
	if base.I64_extend_i32_s(int32(1000))*int64(1000) <= v38-v51 {
		v107 = v49
		v108 = v38
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v62 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[5]))
	if v62 != int64(0) {
		v388 = int32(_a_F_pgstat_report_stat_2)
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[5])) = v38
	return int32(_a_F_pgstat_report_stat_2)
L25:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_report_stat[3])))
	if v80&int32(1) == int32(0) {
		v388 = v2
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L3
L27:
	;
	v107 = v2
	v108 = v96 + v95*int64(1000000) - int64(946684800000000)
	goto L2
L28:
	;
	v119 = F_pgstat_prep_pending_entry(m, int32(1), v115, int64(0), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	m.G0 = v112 + int32(16)
	v214 = int32(0)
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_stat[2]))
	if base.B2i32(v216 == int32(_a_F_pgstat_report_stat_0))|base.B2i32(v216 == v214) != 0 {
		v297 = v214
		goto L40
	} else {
		goto L41
	}
L31:
	;
	return int32(0)
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v123)))
	v126 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_stat[7])))
	*(*int64)(unsafe.Add(mBase, uint32(v123))) = v124 + v126
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v123)+8))
	v131 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_stat[8])))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+8)) = v129 + v131
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v123)+168))
	v136 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+168)) = v134 + v136
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v123)+176))
	v141 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[10]))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+176)) = v139 + v141
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_stat[11]))
	if v145 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v149 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[12]))
	v156 = v108 - v149
	if v156 <= int64(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	v193 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[9])) = v193
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[10])) = v193
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[13])) = v193
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[14])) = v193
	v205 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_stat[8])) = v205
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_stat[7])) = v205
	goto L30
L36:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[12])) = v108
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v123)+192))
	v175 = int64(*(*int32)(unsafe.Add(mBase, uint32(v112)+8)))
	v176 = int64(*(*int32)(unsafe.Add(mBase, uint32(v112)+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+192)) = v174 + (v175 + v176*int64(1000000))
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v123)+200))
	v184 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[13]))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+200)) = v182 + v184
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v123)+208))
	v189 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+208)) = v187 + v189
	goto L35
L37:
	;
	v168 = int32(0)
	v169 = int32(0)
	goto L39
L38:
	;
	v160 = int64(1000000)
	v161 = base.I64_div_u_s(v156, v160)
	v168 = base.I32_wrap_i64(v161)
	v169 = base.I32_wrap_i64(v156 - v161*v160)
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112+int32(12)))) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v112+int32(8)))) = v169
	goto L36
L40:
	;
	v306 = int32(1)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_report_stat[3])))
	if v308 == v306 {
		goto L63
	} else {
		goto L64
	}
L41:
	;
	v224 = v216
	v229 = v2
	goto L42
L42:
	;
	v233 = v224 - int32(16)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if base.Ui32(v235-int32(1)) <= base.Ui32(int32(11)) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v297 = v292
	goto L40
L44:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+32))
	v254 = m.T0[v253].(func(*base.Module, int32, int32) int32)(m, v233, v107)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L31
	} else {
		goto L48
	}
L45:
	;
	v252 = v235*int32(72) + int32(_a_F_pgstat_report_stat_3)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_stat[15]))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v245+v235<<(uint(int32(2))%32)-int32(96))))
	v252 = v251
	goto L44
L48:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if v254 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v259 = v224 - int32(4)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	if base.Ui32(v262-int32(1)) <= base.Ui32(int32(11)) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v292 = int32(1)
	goto L51
L51:
	;
	if v256 == int32(_a_F_pgstat_report_stat_0) {
		v297 = v292
		goto L40
	} else {
		goto L61
	}
L52:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+36))
	if v280 != 0 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v279 = v262*int32(72) + int32(_a_F_pgstat_report_stat_3)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_stat[15]))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v272+v262<<(uint(int32(2))%32)-int32(96))))
	v279 = v278
	goto L52
L56:
	;
	m.T0[v280].(func(*base.Module, int32))(m, v233)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L31
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	F_pfree(m, v260)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L31
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = int32(0)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v287)+4)) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = v290
	v292 = v229
	goto L51
L61:
	;
	if v256 != 0 {
		v224 = v256
		v229 = v292
		goto L42
	} else {
		goto L62
	}
L62:
	;
	goto L43
L63:
	;
	v312 = v297
	v313 = v306
	goto L66
L64:
	;
	v357 = v297
	goto L65
L65:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[6])) = v108
	if v357&int32(1) != 0 {
		goto L79
	} else {
		goto L80
	}
L66:
	;
	if base.Ui32(v313) <= base.Ui32(int32(12)) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v357 = v351
	goto L65
L68:
	;
	v353 = v313 + int32(1)
	if v353 != int32(33) {
		v312 = v351
		v313 = v353
		goto L66
	} else {
		goto L78
	}
L69:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+56))
	if v344 == int32(0) {
		v351 = v312
		goto L68
	} else {
		goto L76
	}
L70:
	;
	v343 = v313*int32(72) + int32(_a_F_pgstat_report_stat_3)
	goto L69
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(int32(8)) < base.Ui32(v313-int32(24)) {
		v351 = v312
		goto L68
	} else {
		goto L73
	}
L73:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_stat[15]))
	if v332 == int32(0) {
		v351 = v312
		goto L68
	} else {
		goto L74
	}
L74:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v332+v313<<(uint(int32(2))%32)-int32(96))))
	if v340 == int32(0) {
		v351 = v312
		goto L68
	} else {
		goto L75
	}
L75:
	;
	v343 = v340
	goto L69
L76:
	;
	v347 = m.T0[v344].(func(*base.Module, int32) int32)(m, v107)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L31
	} else {
		goto L77
	}
L77:
	;
	v351 = v347 | v312
	goto L68
L78:
	;
	goto L67
L79:
	;
	v373 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[5]))
	if v373 != int64(0) {
		v388 = int32(_a_F_pgstat_report_stat_2)
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[5])) = int64(0)
	v384 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_report_stat[3])) = uint8(v384)
	v388 = int32(0)
	goto L1
L82:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_stat[5])) = v108
	return int32(_a_F_pgstat_report_stat_2)
}
func F_pgstat_reset(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
		v30 = l0*int32(72) + int32(_a_F_pgstat_reset_0)
	} else {
		if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
			v30 = int32(0)
		} else {
			v18 = int32(0)
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_reset[0]))
			if v20 == v18 {
				v30 = v18
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v20+l0<<(uint(int32(2))%32)-int32(96))))
				v30 = v28
			}
		}
	}
	v34 = m.G0
	v35 = int32(16)
	v36 = v34 - v35
	m.G0 = v36
	F_gettimeofday(m, v36)
	mBase = m.M
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
	v40 = int64(*(*int32)(unsafe.Add(mBase, uint32(v36)+8)))
	m.G0 = v36 + v35
	v48 = v40 + v39*int64(1000000) - int64(946684800000000)
	F_pgstat_reset_entry(m, l0, l1, l2, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		return
	} else {
		v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
		if v51&int32(2) == int32(0) {
			v58 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_reset[1]))
			v61 = F_pgstat_get_entry_ref_locked(m, int32(1), v58, int64(0), int32(0))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
				*(*int64)(unsafe.Add(mBase, uint32(v63)+280)) = v48
				F_pgstat_unlock_entry(m, v61)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			return
		}
	}
}
func F_pgstat_unlock_entry(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_LWLockRelease(m, v2+int32(4))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_pgstat_wal_flush_cb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v87 int32
	_ = v87
	var v91 int64
	_ = v91
	var v95 int64
	_ = v95
	var v99 int64
	_ = v99
	var v103 int64
	_ = v103
	var v106 int32
	_ = v106
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[0]))
	v12 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v12
	v21 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[1]))
	v23 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[2]))
	if v21 == v23 {
		v106 = int32(0)
		m.G0 = v8 + int32(32)
		return v106
	} else {
		v26 = v11 + int32(_a_F_pgstat_wal_flush_cb_0)
		v27 = int32(_a_F_pgstat_wal_flush_cb_1)
		v28 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
		v30 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[3]))
		v31 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[4]))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v28 + (v30 - v31)
		v35 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
		v37 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[1]))
		v38 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v35 + (v37 - v38)
		v42 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
		v44 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[5]))
		v45 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[6]))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v42 + (v44 - v45)
		v49 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
		v51 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[7]))
		v52 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[8]))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v49 + (v51 - v52)
		if l0 == int32(0) {
			v59 = F_LWLockAcquire(m, v26, int32(0))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				v70 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[9])))
				v71 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[9]))) = v70 + v71
				v74 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[10])))
				v75 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[10]))) = v74 + v75
				v78 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[11])))
				v79 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[11]))) = v78 + v79
				v82 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[12])))
				v83 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[12]))) = v82 + v83
				F_LWLockRelease(m, v26)
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					v91 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[7]))
					*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[8])) = v91
					v95 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[3]))
					*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[4])) = v95
					v99 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[5]))
					*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[6])) = v99
					v103 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[1]))
					*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[2])) = v103
					v106 = int32(0)
					m.G0 = v8 + int32(32)
					return v106
				}
			}
		} else {
			v65 = F_LWLockConditionalAcquire(m, v26, int32(0))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				if v65 == int32(0) {
					v106 = int32(1)
					m.G0 = v8 + int32(32)
					return v106
				} else {
					v70 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[9])))
					v71 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
					*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[9]))) = v70 + v71
					v74 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[10])))
					v75 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[10]))) = v74 + v75
					v78 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[11])))
					v79 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[11]))) = v78 + v79
					v82 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[12])))
					v83 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_pgstat_wal_flush_cb[12]))) = v82 + v83
					F_LWLockRelease(m, v26)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						v91 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[7]))
						*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[8])) = v91
						v95 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[3]))
						*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[4])) = v95
						v99 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[5]))
						*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[6])) = v99
						v103 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[1]))
						*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_flush_cb[2])) = v103
						v106 = int32(0)
						m.G0 = v8 + int32(32)
						return v106
					}
				}
			}
		}
	}
}
