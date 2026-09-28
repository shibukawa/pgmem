package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_calculate_database_size(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int64
	_ = v100
	var v102 int32
	_ = v102
	v6 = m.G0
	v8 = v6 - int32(2128)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_database_size[0]))
	v14 = F_object_aclcheck(m, int32(1262), l0, v12, int64(2048))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
	v32 = v8 + int32(32)
	v37 = F_pg_snprintf(m, v32, int32(1061), int32(_a_F_calculate_database_size_0), v8+int32(16))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L9
	}
L2:
	;
	return int64(0)
L3:
	;
	if v14 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_database_size[0]))
	v23 = F_has_privs_of_role(m, v21, int32(3375))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v23 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v26 = F_get_database_name(m, l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	F_aclcheck_error(m, v14, int32(9), v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	goto L1
L9:
	;
	v39 = F_db_dir_size(m, v32)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v42 = v8 + int32(1104)
	v46 = F_pg_snprintf(m, v42, int32(1024), int32(_a_F_calculate_database_size_1), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v48 = F_AllocateDir(m, v42)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L13
	}
L12:
	;
	F_FreeDir(m, v48)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L35
	}
L13:
	;
	v50 = F_ReadDir(m, v48, v42)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if v50 == int32(0) {
		v100 = v39
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v55 = v50
	v58 = v39
	goto L16
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_database_size[1]))
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v100 = v91
	goto L12
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+19)))
	if v63 != int32(46) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(_a_F_calculate_database_size_2)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_calculate_database_size_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v55 + int32(19)
	v84 = v8 + int32(32)
	v87 = F_pg_snprintf(m, v84, int32(1061), int32(_a_F_calculate_database_size_3), v8)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L31
	}
L23:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+20)))
	if v66 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+20)))
	if v67 != int32(46) {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v73 = F_ReadDir(m, v48, v8+int32(1104))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L2
	} else {
		goto L29
	}
L27:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+21)))
	if v70 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	if v73 != 0 {
		v55 = v73
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v100 = v58
	goto L12
L31:
	;
	v89 = F_db_dir_size(m, v84)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v91 = v89 + v58
	v94 = F_ReadDir(m, v48, v8+int32(1104))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	if v94 != 0 {
		v55 = v94
		v58 = v91
		goto L16
	} else {
		goto L34
	}
L34:
	;
	goto L17
L35:
	;
	m.G0 = v8 + int32(2128)
	return v100
}
func F_cancel_before_shmem_exit(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_before_shmem_exit[0]))
	if v11 <= int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
			F_errmsg_internal(m, int32(_a_F_cancel_before_shmem_exit_0), v8)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_cancel_before_shmem_exit_1), int32(410), int32(_a_F_cancel_before_shmem_exit_2))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v15 = v11 << (uint(int32(4)) % 32)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_cancel_before_shmem_exit[1])))
		if v18 != l0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(_a_F_cancel_before_shmem_exit_0), v8)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_cancel_before_shmem_exit_1), int32(410), int32(_a_F_cancel_before_shmem_exit_2))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_cancel_before_shmem_exit[2])))
			if v22 != l1 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_cancel_before_shmem_exit_0), v8)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_cancel_before_shmem_exit_1), int32(410), int32(_a_F_cancel_before_shmem_exit_2))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_cancel_before_shmem_exit[0])) = v11 - int32(1)
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_carc_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	v7 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
	if v6 < v7 {
		v18 = int32(-1)
	} else {
		if v7 < v6 {
			v18 = int32(1)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v12 < v13 {
				v18 = int32(-1)
			} else {
				v18 = base.B2i32(v13 < v12)
			}
		}
	}
	return v18
}
func F_casecmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = l0
	v8 = l1
	v9 = l2
	goto L4
L2:
	;
	goto L3
L3:
	;
	return int32(0)
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v13 == v14 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v71 = int32(4)
	v76 = v9 - int32(1)
	if v76 != 0 {
		v7 = v7 + v71
		v8 = v8 + v71
		v9 = v76
		goto L4
	} else {
		goto L27
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_casecmp[0]))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
	if v18 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if v62 == v63 {
		goto L6
	} else {
		goto L26
	}
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+64))
	v60 = m.T0[v59].(func(*base.Module, int32, int32) int32)(m, v38, v40)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L18
	} else {
		goto L25
	}
L10:
	;
	if base.Ui32(int32(127)) < base.Ui32(v44) {
		v62 = v44
		v63 = v45
		goto L8
	} else {
		goto L21
	}
L11:
	;
	if base.Ui32(int32(127)) < base.Ui32(v13) {
		v44 = v14
		v45 = v13
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+64))
	v34 = m.T0[v33].(func(*base.Module, int32, int32) int32)(m, v13, v17)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	if base.Ui32((v13-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v31 = v13 | int32(32)
	goto L17
L16:
	;
	v31 = v13
	goto L17
L17:
	;
	v44 = v14
	v45 = v31
	goto L10
L18:
	;
	return int32(0)
L19:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_casecmp[0]))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)))
	if v41 != int32(1) {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v44 = v38
	v45 = v34
	goto L10
L21:
	;
	if base.Ui32((v44-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v57 = v44 | int32(32)
	goto L24
L23:
	;
	v57 = v44
	goto L24
L24:
	;
	v62 = v57
	v63 = v45
	goto L8
L25:
	;
	v62 = v60
	v63 = v34
	goto L8
L26:
	;
	return int32(1)
L27:
	;
	goto L5
}
func F_catalan_ISO_8859_1_create_env(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_SN_new_env(m, int32(36))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 != 0 {
			*(*int64)(unsafe.Add(mBase, uint32(v3)+28)) = int64(0)
		} else {
		}
		return v3
	}
}
