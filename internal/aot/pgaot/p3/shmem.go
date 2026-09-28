package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AttachShmemIndexEntry(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_AttachShmemIndexEntry[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v17 = F_hash_search(m, v12, v14, v3, v3)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 == int32(0) {
			if l1 != 0 {
				m.G0 = v9 + int32(48)
				return base.B2i32(v17 != int32(0))
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v28
					F_errmsg(m, int32(_a_F_AttachShmemIndexEntry_0), v9+int32(32))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_AttachShmemIndexEntry_1), int32(624), int32(_a_F_AttachShmemIndexEntry_2))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+60)))
			if v40 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v112 = m.ExcPending
				if v112 != 0 {
					return int32(0)
				} else {
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v114
					F_errmsg(m, int32(_a_F_AttachShmemIndexEntry_3), v9+int32(16))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_AttachShmemIndexEntry_1), int32(640), int32(_a_F_AttachShmemIndexEntry_2))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
				if base.B2i32(v44 != v45)&base.B2i32(v44 != int32(-1)) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return int32(0)
					} else {
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
						v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v132
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v130
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
						F_errmsg(m, int32(_a_F_AttachShmemIndexEntry_4), v9)
						mBase = m.M
						v138 = m.ExcPending
						if v138 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_AttachShmemIndexEntry_1), int32(649), int32(_a_F_AttachShmemIndexEntry_2))
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					switch v50 {
					case 0:
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
						if v51 == int32(0) {
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v51))) = v54
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17
						m.G0 = v9 + int32(48)
						return base.B2i32(v17 != int32(0))
					case 1:
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v43)+72)) = v56
						*(*int32)(unsafe.Add(mBase, uint32(v43)+64)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v43)+60)) = int32(1208)
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
						v63 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
						v69 = F_hash_create(m, v62, v63, v43+int32(32), v66|int32(_a_F_AttachShmemIndexEntry_5))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v43)+84))
							if v71 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v71))) = v69
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17
							m.G0 = v9 + int32(48)
							return base.B2i32(v17 != int32(0))
						}
					case 2:
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v74)+60)) = v75
						v78 = base.I32_div_s(v73, int32(16))
						*(*uint16)(unsafe.Add(mBase, uint32(v74)+64)) = uint16(v78)
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
						*(*int64)(unsafe.Add(mBase, uint32(v74))) = v80
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v74)+8)) = v82
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v43)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v74)+16)) = v84
						v86 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v74)+24)) = v86
						v88 = *(*int64)(unsafe.Add(mBase, uint32(v43)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v74)+32)) = v88
						v90 = *(*int64)(unsafe.Add(mBase, uint32(v43)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v74)+40)) = v90
						v92 = *(*int64)(unsafe.Add(mBase, uint32(v43)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v74)+48)) = v92
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v74)+56)) = v94
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17
						m.G0 = v9 + int32(48)
						return base.B2i32(v17 != int32(0))
					default:
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17
						m.G0 = v9 + int32(48)
						return base.B2i32(v17 != int32(0))
					}
				}
			}
		}
	}
}
func F_ShmemGetRequestedSize(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemGetRequestedSize[0]))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v13 = base.I64_extend_i32_s(v8 + int32(128))
	goto L3
L2:
	;
	v13 = int64(128)
	goto L3
L3:
	;
	v15 = F_hash_estimate_size(m, v13, int32(64))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v22 = (v15 + int32(127)) & int32(-128)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemGetRequestedSize[0]))
	if v24 == int32(0) {
		v59 = v22
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return v59
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v27 <= int32(0) {
		v59 = v22
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v30 = v22
	v32 = int32(0)
	goto L9
L9:
	;
	v35 = int32(128)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v32<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	if base.Ui32(v42) <= base.Ui32(v35) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v59 = v53
	goto L6
L11:
	;
	v45 = v35
	goto L13
L12:
	;
	v45 = v42
	goto L13
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v53 = F_add_size(m, (v30+v45-int32(1))&(int32(0)-v45), v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v56 = v32 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v56 < v57 {
		v30 = v53
		v32 = v56
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L10
}
func F_ShmemHashAlloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	v7 = (l0 + int32(7)) & int32(-8)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v7) <= base.Ui32(v8-v9) {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9 + v7
		v15 = v9
	} else {
		v15 = int32(0)
	}
	return v15
}
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
	var v44 int64
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
	var v112 int64
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
	v43 = v40 << (uint(int32(4)) % 32)
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_shmem_exit[3])))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_shmem_exit[4])))
	m.T0[v45].(func(*base.Module, int32, int64))(m, l0, v44)
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
	v111 = v108 << (uint(int32(4)) % 32)
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_shmem_exit[6])))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_shmem_exit[7])))
	m.T0[v113].(func(*base.Module, int32, int64))(m, l0, v112)
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
