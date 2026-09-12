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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
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
	v37 = F_pg_snprintf(m, v8+int32(32), int32(1061), int32(_a_F_calculate_database_size_0), v8+int32(16))
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
	v41 = F_db_dir_size(m, v8+int32(32))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v48 = F_pg_snprintf(m, v8+int32(1104), int32(1024), int32(_a_F_calculate_database_size_1), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v52 = F_AllocateDir(m, v8+int32(1104))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L13
	}
L12:
	;
	F_FreeDir(m, v52)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L35
	}
L13:
	;
	v56 = F_ReadDir(m, v52, v8+int32(1104))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if v56 == int32(0) {
		v108 = v41
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v62 = v56
	v64 = v41
	goto L16
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_database_size[1]))
	if v66 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v108 = v99
	goto L12
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+19)))
	if v69 != int32(46) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v62 + int32(19)
	v93 = F_pg_snprintf(m, v8+int32(32), int32(1061), int32(_a_F_calculate_database_size_3), v8)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L31
	}
L23:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+20)))
	if v72 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+20)))
	if v73 != int32(46) {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v79 = F_ReadDir(m, v52, v8+int32(1104))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L29
	}
L27:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+21)))
	if v76 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	if v79 != 0 {
		v62 = v79
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v108 = v64
	goto L12
L31:
	;
	v97 = F_db_dir_size(m, v8+int32(32))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v99 = v97 + v64
	v102 = F_ReadDir(m, v52, v8+int32(1104))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	if v102 != 0 {
		v62 = v102
		v64 = v99
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
	return v108
}
func F_cancel_before_shmem_exit(m *base.Module, l0 int32, l1 int32) {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_before_shmem_exit[0]))
	if v11 <= int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
			F_errmsg_internal(m, int32(_a_F_cancel_before_shmem_exit_0), v8)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_cancel_before_shmem_exit_1), int32(410), int32(_a_F_cancel_before_shmem_exit_2))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v15 = v11 - int32(1)
		v17 = v15 << (uint(int32(3)) % 32)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_cancel_before_shmem_exit[1])))
		if v20 != l0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(_a_F_cancel_before_shmem_exit_0), v8)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_cancel_before_shmem_exit_1), int32(410), int32(_a_F_cancel_before_shmem_exit_2))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_cancel_before_shmem_exit[2])))
			if v22 != l1 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_cancel_before_shmem_exit_0), v8)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_cancel_before_shmem_exit_1), int32(410), int32(_a_F_cancel_before_shmem_exit_2))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_cancel_before_shmem_exit[0])) = v15
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
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v5 = l0
	v6 = l1
	v7 = l2
	goto L4
L2:
	;
	goto L3
L3:
	;
	return int32(0)
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v9 == v10 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v122 = int32(4)
	v127 = v7 - int32(1)
	if v127 != 0 {
		v5 = v5 + v122
		v6 = v6 + v122
		v7 = v127
		goto L4
	} else {
		goto L49
	}
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_casecmp[0]))
	switch v14 - int32(1) {
	case 0:
		goto L12
	case 1:
		goto L11
	case 2:
		goto L10
	default:
		goto L13
	}
L8:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_casecmp[0]))
	switch v68 - int32(1) {
	case 0:
		goto L32
	case 1:
		goto L31
	case 2:
		goto L30
	default:
		goto L33
	}
L9:
	;
	v64 = v60
	goto L8
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_casecmp[1]))
	if base.Ui32(int32(127)) < base.Ui32(v9) {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_casecmp[1]))
	if base.Ui32(int32(127)) < base.Ui32(v9) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	if base.Ui32(v9) <= base.Ui32(int32(127)) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	if base.Ui32(int32(127)) < base.Ui32(v9) {
		v60 = v9
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v19 = F_pg_tolower(m, v9)
	mBase = m.M
	v64 = v19
	goto L8
L15:
	;
	v31 = v9<<(uint(int32(2))%32) + int32(_a_F_casecmp_0)
	goto L17
L16:
	;
	v26 = F_case_index(m, v9)
	mBase = m.M
	v31 = v26<<(uint(int32(2))%32) + int32(_a_F_casecmp_1)
	goto L17
L17:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v33 = v32
	goto L20
L19:
	;
	v33 = v9
	goto L20
L20:
	;
	v64 = v33
	goto L8
L21:
	;
	v45 = F_towlower(m, v9)
	mBase = m.M
	v64 = v45
	goto L8
L22:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+4)))
	if v38&int32(1) == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v43 = F_pg_tolower(m, v9)
	mBase = m.M
	v64 = v43
	goto L8
L24:
	;
	if base.Ui32(int32(255)) < base.Ui32(v9) {
		v60 = v9
		goto L9
	} else {
		goto L27
	}
L25:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+4)))
	if v50&int32(1) == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v55 = F_pg_tolower(m, v9)
	mBase = m.M
	v64 = v55
	goto L8
L27:
	;
	v59 = F_tolower(m, v9)
	mBase = m.M
	v60 = v59
	goto L9
L28:
	;
	if v64 == v118 {
		goto L6
	} else {
		goto L48
	}
L29:
	;
	v118 = v114
	goto L28
L30:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_casecmp[1]))
	if base.Ui32(int32(127)) < base.Ui32(v65) {
		goto L44
	} else {
		goto L45
	}
L31:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_casecmp[1]))
	if base.Ui32(int32(127)) < base.Ui32(v65) {
		goto L41
	} else {
		goto L42
	}
L32:
	;
	if base.Ui32(v65) <= base.Ui32(int32(127)) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if base.Ui32(int32(127)) < base.Ui32(v65) {
		v114 = v65
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v73 = F_pg_tolower(m, v65)
	mBase = m.M
	v118 = v73
	goto L28
L35:
	;
	v85 = v65<<(uint(int32(2))%32) + int32(_a_F_casecmp_0)
	goto L37
L36:
	;
	v80 = F_case_index(m, v65)
	mBase = m.M
	v85 = v80<<(uint(int32(2))%32) + int32(_a_F_casecmp_1)
	goto L37
L37:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v86 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v87 = v86
	goto L40
L39:
	;
	v87 = v65
	goto L40
L40:
	;
	v118 = v87
	goto L28
L41:
	;
	v99 = F_towlower(m, v65)
	mBase = m.M
	v118 = v99
	goto L28
L42:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
	if v92&int32(1) == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v97 = F_pg_tolower(m, v65)
	mBase = m.M
	v118 = v97
	goto L28
L44:
	;
	if base.Ui32(int32(255)) < base.Ui32(v65) {
		v114 = v65
		goto L29
	} else {
		goto L47
	}
L45:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
	if v104&int32(1) == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v109 = F_pg_tolower(m, v65)
	mBase = m.M
	v118 = v109
	goto L28
L47:
	;
	v113 = F_tolower(m, v65)
	mBase = m.M
	v114 = v113
	goto L29
L48:
	;
	return int32(1)
L49:
	;
	goto L5
}
func F_catalan_ISO_8859_1_create_env(m *base.Module) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_SN_create_env(m, int32(0), int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
