package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_shm_mq_attach(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v5 = F_palloc(m, int32(44))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v5)+12)) = v9
		v11 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v5)+20)) = v9
		*(*int64)(unsafe.Add(mBase, uint32(v5)+28)) = v9
		*(*uint16)(unsafe.Add(mBase, uint32(v5)+36)) = uint16(v11)
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_attach[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+40)) = v22
		if l1 != 0 {
			F_on_dsm_detach(m, l1, int32(1207), base.I64_extend_i32_u(l0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				return v5
			}
		} else {
			return v5
		}
	}
}
func F_shm_mq_detach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var __phi102 int32
	_ = __phi102
	var v103 int32
	_ = v103
	var __phi103 int32
	_ = __phi103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int64
	_ = v114
	var v118 int32
	_ = v118
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = int32(0)
	v12 = base.AtomicRmwOr32(m, v9, int32(_a_F_shm_mq_detach_0), v9)
	v14 = int64(0)
	v16 = int32(24)
	v17 = base.AtomicRmwCmpxchg64(m, v8, v16, v14, v14)
	v20 = base.AtomicRmwXchg64(m, v8, v16, base.I64_extend_i32_u(v7)+v17)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v9
	goto L3
L2:
	;
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = base.AtomicRmwXchg32(m, v24, int32(0), int32(1))
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_s_lock(m, v24, int32(_a_F_shm_mq_detach_1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach[0]))
	if v31 == v33 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return
L8:
	;
	goto L6
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v36 = v35
	goto L11
L10:
	;
	v36 = v31
	goto L11
L11:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+36)) = uint8(v37)
	v39 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v24))), uint32(v39))
	if v36 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = v36 + int32(316)
	v44 = int32(0)
	v47 = base.AtomicRmwOr32(m, v44, int32(_a_F_shm_mq_detach_2), v44)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v48 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v95 != 0 {
		goto L29
	} else {
		goto L30
	}
L15:
	;
	goto L14
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(1)
	v51 = int32(0)
	v54 = base.AtomicRmwOr32(m, v51, int32(_a_F_shm_mq_detach_2), v51)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v55 == v51 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	if v58 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach[1]))
	if v62 == v58 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v64 = m.G0
	v66 = v64 - int32(16)
	m.G0 = v66
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach[2]))
	if v69 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v92 = F_pgmem_kill(m, v58, int32(23))
	mBase = m.M
	goto L16
L23:
	;
	m.G0 = v66 + int32(16)
	goto L15
L24:
	;
	v72 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v66)+15)) = uint8(v72)
	goto L25
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach[3]))
	v80 = F_write(m, v76, v66+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v80 {
		goto L23
	} else {
		goto L27
	}
L26:
	;
	goto L23
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach[4]))
	if v84 == int32(27) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v96 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+32))
	if v97 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v137 != 0 {
		goto L43
	} else {
		goto L44
	}
L32:
	;
	goto L31
L33:
	;
	__phi102 = v97
	__phi103 = v95 + int32(32)
	v102 = __phi102
	v103 = __phi103
	goto L36
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v108 = v102 - int32(16)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v109 != int32(1207) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	if v106 != 0 {
		__phi102 = v106
		__phi103 = v102
		v102 = __phi102
		v103 = __phi103
		goto L36
	} else {
		goto L42
	}
L39:
	;
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v102-int32(8))))
	if v114 != v96 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v106
	F_pfree(m, v108)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	goto L32
L42:
	;
	goto L37
L43:
	;
	F_pfree(m, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L7
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_pfree(m, l0)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L7
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	return
}
func F_shm_mq_detach_callback(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	v4 = base.I32_wrap_i64(l1)
	v7 = base.AtomicRmwXchg32(m, v4, int32(0), int32(1))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, v4, int32(_a_F_shm_mq_detach_callback_0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach_callback[0]))
	if v11 == v13 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v16 = v15
	goto L8
L7:
	;
	v16 = v11
	goto L8
L8:
	;
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4)+36)) = uint8(v17)
	v19 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v4))), uint32(v19))
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v23 = v16 + int32(316)
	v24 = int32(0)
	v27 = base.AtomicRmwOr32(m, v24, int32(_a_F_shm_mq_detach_callback_1), v24)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v28 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	return
L12:
	;
	goto L11
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(1)
	v31 = int32(0)
	v34 = base.AtomicRmwOr32(m, v31, int32(_a_F_shm_mq_detach_callback_1), v31)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v35 == v31 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v38 == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach_callback[1]))
	if v42 == v38 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v44 = m.G0
	v46 = v44 - int32(16)
	m.G0 = v46
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach_callback[2]))
	if v49 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v72 = F_pgmem_kill(m, v38, int32(23))
	mBase = m.M
	goto L13
L20:
	;
	m.G0 = v46 + int32(16)
	goto L12
L21:
	;
	v52 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+15)) = uint8(v52)
	goto L22
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach_callback[3]))
	v60 = F_write(m, v56, v46+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v60 {
		goto L20
	} else {
		goto L24
	}
L23:
	;
	goto L20
L24:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach_callback[4]))
	if v64 == int32(27) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
}
func F_shm_mq_receive_bytes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v159 int64
	_ = v159
	var v163 int64
	_ = v163
	var v164 int32
	_ = v164
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v186 int32
	_ = v186
	var v188 int64
	_ = v188
	var v190 int64
	_ = v190
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	v9 = int64(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v20 = base.AtomicRmwCmpxchg64(m, v15, int32(24), v9, v9)
	v24 = base.AtomicRmwCmpxchg64(m, v15, int32(16), v9, v9)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = base.I64_extend_i32_u(v25)
	v27 = v24 + v26
	v28 = base.I64_extend_i32_u(v16)
	v29 = base.I64_rem_u_s(v27, v28)
	v30 = v20 - v27
	v31 = base.I64_extend_i32_u(l1)
	if base.B2i32(base.Ui64(v31) <= base.Ui64(v30))|base.B2i32(base.Ui64(v28) <= base.Ui64(v30+v29)) != 0 {
		v180 = v30
		v182 = v29
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v186 = base.I32_wrap_i64(v182)
	v188 = base.I64_extend_i32_u(v16 - v186)
	if base.Ui64(v180) < base.Ui64(v188) {
		goto L37
	} else {
		goto L38
	}
L2:
	;
	v43 = v25
	v45 = v20
	v48 = v26
	goto L3
L3:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+36)))
	if v50 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v180 = v168
	v182 = v167
	goto L1
L5:
	;
	v156 = int64(0)
	v159 = base.AtomicRmwCmpxchg64(m, v15, int32(24), v156, v156)
	v163 = base.AtomicRmwCmpxchg64(m, v15, int32(16), v156, v156)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v165 = base.I64_extend_i32_u(v164)
	v166 = v163 + v165
	v167 = base.I64_rem_u_s(v166, v28)
	v168 = v159 - v166
	if base.Ui64(v31) <= base.Ui64(v168) {
		v180 = v168
		v182 = v167
		goto L1
	} else {
		goto L35
	}
L6:
	;
	v53 = int32(0)
	v56 = base.AtomicRmwOr32(m, v53, int32(_a_F_shm_mq_receive_bytes_0), v53)
	v57 = int64(0)
	v60 = base.AtomicRmwCmpxchg64(m, v15, int32(24), v57, v57)
	if v45 != v60 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	return int32(2)
L10:
	;
	v64 = int32(0)
	v67 = base.AtomicRmwOr32(m, v64, int32(_a_F_shm_mq_receive_bytes_0), v64)
	v68 = int64(0)
	v70 = int32(16)
	v71 = base.AtomicRmwCmpxchg64(m, v15, v70, v68, v68)
	v74 = base.AtomicRmwXchg64(m, v15, v70, v71+v48)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v77 = v75 + int32(316)
	v81 = base.AtomicRmwOr32(m, v64, int32(_a_F_shm_mq_receive_bytes_1), v64)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v82 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	if l2 != 0 {
		goto L27
	} else {
		goto L28
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	goto L12
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(1)
	v85 = int32(0)
	v88 = base.AtomicRmwOr32(m, v85, int32(_a_F_shm_mq_receive_bytes_1), v85)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v89 == v85 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	if v92 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive_bytes[0]))
	if v96 == v92 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v98 = m.G0
	v100 = v98 - int32(16)
	m.G0 = v100
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive_bytes[1]))
	if v103 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v126 = F_pgmem_kill(m, v92, int32(23))
	mBase = m.M
	goto L14
L21:
	;
	m.G0 = v100 + int32(16)
	goto L13
L22:
	;
	v106 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+15)) = uint8(v106)
	goto L23
L23:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive_bytes[2]))
	v114 = F_write(m, v110, v100+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v114 {
		goto L21
	} else {
		goto L25
	}
L24:
	;
	goto L21
L25:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive_bytes[3]))
	if v118 == int32(27) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	return int32(1)
L28:
	;
	goto L29
L29:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive_bytes[4]))
	v138 = F_WaitLatch(m, v134, int32(33), int32(0), int32(134217763))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive_bytes[4]))
	v144 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v144
	v149 = base.AtomicRmwOr32(m, v144, int32(_a_F_shm_mq_receive_bytes_1), v144)
	goto L32
L32:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive_bytes[5]))
	if v151 == int32(0) {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L5
L35:
	;
	if base.Ui64(v168+v167) < base.Ui64(v28) {
		v43 = v164
		v45 = v159
		v48 = v165
		goto L3
	} else {
		goto L36
	}
L36:
	;
	goto L4
L37:
	;
	v190 = v180
	goto L39
L38:
	;
	v190 = v188
	goto L39
L39:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(l3))) = uint32(v190)
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+37)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v15 + v192 + v186 + int32(38)
	v198 = int32(0)
	v201 = base.AtomicRmwOr32(m, v198, int32(_a_F_shm_mq_receive_bytes_0), v198)
	return v198
}
func F_shm_mq_wait_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	goto L2
L1:
	;
	m.G0 = v9 + int32(16)
	return v68
L2:
	;
	v19 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v68 = (v29 ^ int32(-1)) & base.B2i32(v25 != int32(0))
	goto L1
L4:
	;
	F_s_lock(m, l0, int32(_a_F_shm_mq_wait_internal_0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v26))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v29|v25 == v26 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L6
L9:
	;
	if l2 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	goto L3
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_wait_internal[0]))
	v47 = F_WaitLatch(m, v43, int32(33), int32(0), int32(134217761))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L16
	}
L13:
	;
	v37 = F_GetBackgroundWorkerPid(m, l2, v9+int32(12))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(v37) <= base.Ui32(int32(1)) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v68 = int32(0)
	goto L1
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_wait_internal[0]))
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v51
	v56 = base.AtomicRmwOr32(m, v51, int32(_a_F_shm_mq_wait_internal_1), v51)
	goto L17
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_wait_internal[1]))
	if v58 == int32(0) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L2
}
func F_shm_toc_attach(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if v4 == l0 {
		v6 = l1
	} else {
		v6 = int32(0)
	}
	return v6
}
func F_shm_toc_insert(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v7 = int32(8)
	v8 = l0 + v7
	v11 = base.AtomicRmwXchg32(m, l0, v7, int32(1))
	if v11 != 0 {
		F_s_lock(m, v8, int32(_a_F_shm_toc_insert_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v20 = v16 + v17<<(uint(int32(4))%32)
			v27 = int32(0)
			if base.B2i32(base.B2i32(base.Ui32(v15) < base.Ui32(v20+int32(40)))|base.B2i32(v17 == int32(-1)) == v27)&base.B2i32(base.Ui32(v20+int32(24)) < base.Ui32(int32(-16))) == v27 {
				v36 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v36))
				F_errstart_cold(m, int32(21), v36)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_errcode(m, int32(_a_F_shm_toc_insert_1))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_shm_toc_insert_2), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_shm_toc_insert_3), int32(207), int32(_a_F_shm_toc_insert_4))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v57 = l0 + v17<<(uint(int32(4))%32)
				*(*int64)(unsafe.Add(mBase, uint32(v57)+24)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v57)+32)) = l2 - l0
				v61 = int32(0)
				v64 = base.AtomicRmwOr32(m, v61, int32(_a_F_shm_toc_insert_5), v61)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v65 + int32(1)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0)+8)), uint32(v61))
				return
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v20 = v16 + v17<<(uint(int32(4))%32)
		v27 = int32(0)
		if base.B2i32(base.B2i32(base.Ui32(v15) < base.Ui32(v20+int32(40)))|base.B2i32(v17 == int32(-1)) == v27)&base.B2i32(base.Ui32(v20+int32(24)) < base.Ui32(int32(-16))) == v27 {
			v36 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v36))
			F_errstart_cold(m, int32(21), v36)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_shm_toc_insert_1))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_shm_toc_insert_2), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_shm_toc_insert_3), int32(207), int32(_a_F_shm_toc_insert_4))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v57 = l0 + v17<<(uint(int32(4))%32)
			*(*int64)(unsafe.Add(mBase, uint32(v57)+24)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v57)+32)) = l2 - l0
			v61 = int32(0)
			v64 = base.AtomicRmwOr32(m, v61, int32(_a_F_shm_toc_insert_5), v61)
			v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v65 + int32(1)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0)+8)), uint32(v61))
			return
		}
	}
}
