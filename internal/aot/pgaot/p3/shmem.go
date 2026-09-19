package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_shmem_exit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_shmem_exit[0])) = uint8(v9)
	F_LWLockReleaseAll(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v19
	F_errmsg_internal(m, int32(_a_F_shmem_exit_0), v6+int32(16))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v31 = int32(_a_F_shmem_exit_1)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[1]))
	v35 = v33 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[1])) = v35
	if int32(0) <= v35 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	F_errfinish(m, int32(_a_F_shmem_exit_2), int32(248), int32(_a_F_shmem_exit_3))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v40 = v35
	goto L12
L10:
	;
	goto L11
L11:
	;
	v60 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[1])) = v60
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[2]))
	if base.B2i32(v63 == v60)|base.B2i32(v63 == int32(_a_F_shmem_exit_4)) != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v43 = v40 << (uint(int32(3)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_shmem_exit[3])))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_shmem_exit[4])))
	m.T0[v45].(func(*base.Module, int32, int32))(m, l0, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	v48 = int32(_a_F_shmem_exit_1)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[1]))
	v52 = v50 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[1])) = v52
	if int32(0) <= v52 {
		v40 = v52
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v85 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L23
	}
L17:
	;
	v70 = v63
	goto L18
L18:
	;
	F_dsm_detach(m, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L16
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[2]))
	if v75 == int32(0) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	if v75 != int32(_a_F_shmem_exit_4) {
		v70 = v75
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	if v85 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v89
	F_errmsg_internal(m, int32(_a_F_shmem_exit_5), v6)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v99 = int32(_a_F_shmem_exit_6)
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[5]))
	v103 = v101 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[5])) = v103
	if int32(0) <= v103 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	F_errfinish(m, int32(_a_F_shmem_exit_2), int32(281), int32(_a_F_shmem_exit_3))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v108 = v103
	goto L32
L30:
	;
	goto L31
L31:
	;
	v128 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_shmem_exit[0])) = uint8(v128)
	*(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[5])) = v128
	m.G0 = v6 + int32(32)
	return
L32:
	;
	v111 = v108 << (uint(int32(3)) % 32)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_shmem_exit[6])))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_shmem_exit[7])))
	m.T0[v113].(func(*base.Module, int32, int32))(m, l0, v112)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	goto L31
L34:
	;
	v116 = int32(_a_F_shmem_exit_6)
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[5]))
	v120 = v118 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[5])) = v120
	if int32(0) <= v120 {
		v108 = v120
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
}
