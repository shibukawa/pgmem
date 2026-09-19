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
	var v26 int32
	_ = v26
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
			F_on_dsm_detach(m, l1, int32(1107), l0)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
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
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var __phi92 int32
	_ = __phi92
	var v93 int32
	_ = v93
	var __phi93 int32
	_ = __phi93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = int64(0)
	v12 = int32(24)
	v13 = base.AtomicRmwCmpxchg64(m, v8, v12, v10, v10)
	v16 = base.AtomicRmwXchg64(m, v8, v12, base.I64_extend_i32_u(v7)+v13)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = base.AtomicRmwXchg32(m, v20, int32(0), int32(1))
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_s_lock(m, v20, int32(_a_F_shm_mq_detach_0), int32(886), int32(_a_F_shm_mq_detach_1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach[0]))
	if v29 == v31 {
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
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v34 = v33
	goto L11
L10:
	;
	v34 = v29
	goto L11
L11:
	;
	v35 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+36)) = uint8(v35)
	v37 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v20))), uint32(v37))
	if v34 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = v34 + int32(20)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v42 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v85 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v45 == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	if v48 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach[1]))
	if v52 == v48 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v54 = m.G0
	v56 = v54 - int32(16)
	m.G0 = v56
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach[2]))
	if v59 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v82 = F_pgmem_kill(m, v48, int32(23))
	mBase = m.M
	goto L16
L23:
	;
	m.G0 = v56 + int32(16)
	goto L15
L24:
	;
	v62 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+15)) = uint8(v62)
	goto L25
L25:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach[3]))
	v70 = F_write(m, v66, v56+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v70 {
		goto L23
	} else {
		goto L27
	}
L26:
	;
	goto L23
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach[4]))
	if v74 == int32(27) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)+32))
	if v87 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v127 != 0 {
		goto L43
	} else {
		goto L44
	}
L32:
	;
	goto L31
L33:
	;
	__phi92 = v87
	__phi93 = v85 + int32(32)
	v92 = __phi92
	v93 = __phi93
	goto L36
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v98 = v92 - int32(8)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v99 != int32(1107) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	if v96 != 0 {
		__phi92 = v96
		__phi93 = v92
		v92 = __phi92
		v93 = __phi93
		goto L36
	} else {
		goto L42
	}
L39:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v92-int32(4))))
	if v104 != v86 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v96
	F_pfree(m, v98)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
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
	F_pfree(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
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
	v131 = m.ExcPending
	if v131 != 0 {
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
func F_shm_mq_detach_callback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	v5 = base.AtomicRmwXchg32(m, l1, int32(0), int32(1))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, l1, int32(_a_F_shm_mq_detach_callback_0), int32(886), int32(_a_F_shm_mq_detach_callback_1))
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v16 = v15
	goto L8
L7:
	;
	v16 = v11
	goto L8
L8:
	;
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)) = uint8(v17)
	v19 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l1))), uint32(v19))
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v23 = v16 + int32(20)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v24 != 0 {
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
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v27 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v30 == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach_callback[1]))
	if v34 == v30 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v36 = m.G0
	v38 = v36 - int32(16)
	m.G0 = v38
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach_callback[2]))
	if v41 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v64 = F_pgmem_kill(m, v30, int32(23))
	mBase = m.M
	goto L13
L20:
	;
	m.G0 = v38 + int32(16)
	goto L12
L21:
	;
	v44 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+15)) = uint8(v44)
	goto L22
L22:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach_callback[3]))
	v52 = F_write(m, v48, v38+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v52 {
		goto L20
	} else {
		goto L24
	}
L23:
	;
	goto L20
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_detach_callback[4]))
	if v56 == int32(27) {
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
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v139 int64
	_ = v139
	var v143 int64
	_ = v143
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v166 int32
	_ = v166
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v172 int32
	_ = v172
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
		v160 = v30
		v162 = v29
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v166 = base.I32_wrap_i64(v162)
	v168 = base.I64_extend_i32_u(v16 - v166)
	if base.Ui64(v160) < base.Ui64(v168) {
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
	v160 = v148
	v162 = v147
	goto L1
L5:
	;
	v136 = int64(0)
	v139 = base.AtomicRmwCmpxchg64(m, v15, int32(24), v136, v136)
	v143 = base.AtomicRmwCmpxchg64(m, v15, int32(16), v136, v136)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v145 = base.I64_extend_i32_u(v144)
	v146 = v143 + v145
	v147 = base.I64_rem_u_s(v146, v28)
	v148 = v139 - v146
	if base.Ui64(v31) <= base.Ui64(v148) {
		v160 = v148
		v162 = v147
		goto L1
	} else {
		goto L35
	}
L6:
	;
	v53 = int64(0)
	v56 = base.AtomicRmwCmpxchg64(m, v15, int32(24), v53, v53)
	if v45 != v56 {
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
	v60 = int64(0)
	v62 = int32(16)
	v63 = base.AtomicRmwCmpxchg64(m, v15, v62, v60, v60)
	v66 = base.AtomicRmwXchg64(m, v15, v62, v63+v48)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v69 = v67 + int32(20)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v70 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v73 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v76 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive_bytes[0]))
	if v80 == v76 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v82 = m.G0
	v84 = v82 - int32(16)
	m.G0 = v84
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive_bytes[1]))
	if v87 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v110 = F_pgmem_kill(m, v76, int32(23))
	mBase = m.M
	goto L14
L21:
	;
	m.G0 = v84 + int32(16)
	goto L13
L22:
	;
	v90 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+15)) = uint8(v90)
	goto L23
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive_bytes[2]))
	v98 = F_write(m, v94, v84+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v98 {
		goto L21
	} else {
		goto L25
	}
L24:
	;
	goto L21
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive_bytes[3]))
	if v102 == int32(27) {
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
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive_bytes[4]))
	v122 = F_WaitLatch(m, v118, int32(33), int32(0), int32(134217763))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive_bytes[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = int32(0)
	goto L32
L32:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive_bytes[5]))
	if v131 == int32(0) {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L5
L35:
	;
	if base.Ui64(v148+v147) < base.Ui64(v28) {
		v43 = v144
		v45 = v139
		v48 = v145
		goto L3
	} else {
		goto L36
	}
L36:
	;
	goto L4
L37:
	;
	v170 = v160
	goto L39
L38:
	;
	v170 = v168
	goto L39
L39:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(l3))) = uint32(v170)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+37)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v15 + v172 + v166 + int32(38)
	return int32(0)
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	goto L2
L1:
	;
	m.G0 = v9 + int32(16)
	return v66
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
	v66 = (v31 ^ int32(-1)) & base.B2i32(v27 != int32(0))
	goto L1
L4:
	;
	F_s_lock(m, l0, int32(_a_F_shm_mq_wait_internal_0), int32(1228), int32(_a_F_shm_mq_wait_internal_1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v28 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v28))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v31|v27 == v28 {
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
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_wait_internal[0]))
	v49 = F_WaitLatch(m, v45, int32(33), int32(0), int32(134217761))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L16
	}
L13:
	;
	v39 = F_GetBackgroundWorkerPid(m, l2, v9+int32(12))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(v39) <= base.Ui32(int32(1)) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v66 = int32(0)
	goto L1
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_wait_internal[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = int32(0)
	goto L17
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_wait_internal[1]))
	if v56 == int32(0) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	v7 = int32(8)
	v8 = l0 + v7
	v11 = base.AtomicRmwXchg32(m, l0, v7, int32(1))
	if v11 != 0 {
		F_s_lock(m, v8, int32(_a_F_shm_toc_insert_0), int32(184), int32(_a_F_shm_toc_insert_1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v22 = v18 + v19<<(uint(int32(4))%32)
			v29 = int32(0)
			if base.B2i32(base.B2i32(base.Ui32(v17) < base.Ui32(v22+int32(40)))|base.B2i32(v19 == int32(-1)) == v29)&base.B2i32(base.Ui32(v22+int32(24)) < base.Ui32(int32(-16))) == v29 {
				v38 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v38))
				F_errstart_cold(m, int32(21), v38)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errcode(m, int32(_a_F_shm_toc_insert_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_shm_toc_insert_3), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_shm_toc_insert_0), int32(200), int32(_a_F_shm_toc_insert_1))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
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
				v59 = l0 + v19<<(uint(int32(4))%32)
				*(*int64)(unsafe.Add(mBase, uint32(v59)+24)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v59)+32)) = l2 - l0
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v63 + int32(1)
				v67 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0)+8)), uint32(v67))
				return
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v22 = v18 + v19<<(uint(int32(4))%32)
		v29 = int32(0)
		if base.B2i32(base.B2i32(base.Ui32(v17) < base.Ui32(v22+int32(40)))|base.B2i32(v19 == int32(-1)) == v29)&base.B2i32(base.Ui32(v22+int32(24)) < base.Ui32(int32(-16))) == v29 {
			v38 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v38))
			F_errstart_cold(m, int32(21), v38)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_shm_toc_insert_2))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_shm_toc_insert_3), int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_shm_toc_insert_0), int32(200), int32(_a_F_shm_toc_insert_1))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
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
			v59 = l0 + v19<<(uint(int32(4))%32)
			*(*int64)(unsafe.Add(mBase, uint32(v59)+24)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v59)+32)) = l2 - l0
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v63 + int32(1)
			v67 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0)+8)), uint32(v67))
			return
		}
	}
}
