package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CommitTsPagePrecedes(m *base.Module, l0 int64, l1 int64) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v3 = int32(0)
	v7 = int32(819)
	v9 = int32(4)
	v10 = base.I32_wrap_i64(l0)*v7 + v9
	v13 = base.I32_wrap_i64(l1) * v7
	v15 = v13 + v9
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v15))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v10)) == v3 {
		v27 = base.B2i32(base.Ui32(v10) < base.Ui32(v15))
	} else {
		v27 = int32(base.Ui32(v10-v15) >> (uint(int32(31)) % 32))
	}
	if v27 != 0 {
		v29 = v13 + int32(822)
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v29))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v10)) == int32(0) {
			v41 = base.B2i32(base.Ui32(v10) < base.Ui32(v29))
		} else {
			v41 = int32(base.Ui32(v10-v29) >> (uint(int32(31)) % 32))
		}
		v42 = v41
	} else {
		v42 = v3
	}
	return v42
}
func F_commit_ts_identify(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if l0 == int32(16) {
		v6 = int32(538727)
	} else {
		v6 = int32(0)
	}
	if l0 != 0 {
		v8 = v6
	} else {
		v8 = int32(541672)
	}
	return v8
}
func F_commit_ts_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	v13 = v11 & int32(240)
	switch v13 {
	case 0:
		v73 = *(*int32)(unsafe.Add(mBase, _consts[65]))
		v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+28))
		v75 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
		v76 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
		v78 = int64(*(*uint16)(unsafe.Add(mBase, _consts[66])))
		v79 = base.I64_rem_s(v76, v78)
		v83 = v74 + base.I32_wrap_i64(v79)<<(uint(int32(7))%32)
		v85 = F_LWLockAcquire(m, v83, int32(0))
		mBase = m.M
		v86 = m.ExcPending
		if v86 != 0 {
			return
		} else {
			v87 = int32(4409860)
			v89 = F_SimpleLruZeroPage(m, v87, v76)
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return
			} else {
				F_SimpleLruWritePage(m, v87, v89)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return
				} else {
					F_LWLockRelease(m, v83)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v13
			F_errmsg_internal(m, int32(52317), v8)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				F_errfinish(m, int32(493529), int32(1063), int32(242463))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 16:
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		v17 = *(*int32)(unsafe.Add(mBase, _consts[29]))
		v21 = F_LWLockAcquire(m, v17+int32(4992), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, _consts[67]))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
			if v25 == int32(0) {
			} else {
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v15))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v25)) == int32(0) {
					v39 = base.B2i32(base.Ui32(v25) < base.Ui32(v15))
				} else {
					v39 = int32(base.Ui32(v25-v15) >> (uint(int32(31)) % 32))
				}
				if v39 == int32(0) {
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, _consts[67]))
					*(*int32)(unsafe.Add(mBase, uint32(v43)+40)) = v15
				}
			}
			v46 = *(*int32)(unsafe.Add(mBase, _consts[29]))
			F_LWLockRelease(m, v46+int32(4992))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				v51 = int32(4409860)
				v52 = *(*int32)(unsafe.Add(mBase, _consts[65]))
				v53 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
				*(*int64)(unsafe.Add(mBase, uint32(v52)+48)) = v53
				v56 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
				F_SimpleLruTruncate(m, v51, v56)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		}
	}
}
func F_get_ts_config_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	F_DeconstructQualifiedName(m, l0, v10+int32(12), v10+int32(8))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	m.G0 = v10 + int32(16)
	return v106
L4:
	;
	if l1 != 0 {
		v106 = v83
		goto L3
	} else {
		goto L26
	}
L5:
	;
	v83 = int32(0)
	goto L4
L6:
	;
	v22 = F_LookupExplicitNamespace(m, v20, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L15
	}
L9:
	;
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v24 = int32(0)
	goto L12
L11:
	;
	v24 = l1
	goto L12
L12:
	;
	if v24 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v27 = int32(0)
	v29 = F_GetSysCacheOid(m, int32(73), v26, v22, v27, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v83 = v29
	goto L4
L15:
	;
	v33 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[216]))
	if v35 == v33 {
		v83 = v33
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v38 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v39 <= v38 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v46 = v38
	goto L18
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v46<<(uint(int32(2))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	if v54 != v56 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L5
L20:
	;
	v59 = int32(0)
	v61 = F_GetSysCacheOid(m, int32(73), v42, v54, v59, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v65 = v46 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v65 < v66 {
		v46 = v65
		goto L18
	} else {
		goto L25
	}
L23:
	;
	if v61 != 0 {
		v106 = v61
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L19
L26:
	;
	if v83 != 0 {
		v106 = v83
		goto L3
	} else {
		goto L27
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v91 = F_NameListToString(m, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v91
	F_errmsg(m, int32(71295), v10)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(499353), int32(3201), int32(434351))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_ts_template_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	F_DeconstructQualifiedName(m, l0, v10+int32(12), v10+int32(8))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	m.G0 = v10 + int32(16)
	return v106
L4:
	;
	if l1 != 0 {
		v106 = v83
		goto L3
	} else {
		goto L26
	}
L5:
	;
	v83 = int32(0)
	goto L4
L6:
	;
	v22 = F_LookupExplicitNamespace(m, v20, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L15
	}
L9:
	;
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v24 = int32(0)
	goto L12
L11:
	;
	v24 = l1
	goto L12
L12:
	;
	if v24 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v27 = int32(0)
	v29 = F_GetSysCacheOid(m, int32(79), v26, v22, v27, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v83 = v29
	goto L4
L15:
	;
	v33 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[216]))
	if v35 == v33 {
		v83 = v33
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v38 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v39 <= v38 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v46 = v38
	goto L18
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v46<<(uint(int32(2))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	if v54 != v56 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L5
L20:
	;
	v59 = int32(0)
	v61 = F_GetSysCacheOid(m, int32(79), v42, v54, v59, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v65 = v46 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v65 < v66 {
		v46 = v65
		goto L18
	} else {
		goto L25
	}
L23:
	;
	if v61 != 0 {
		v106 = v61
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L19
L26:
	;
	if v83 != 0 {
		v106 = v83
		goto L3
	} else {
		goto L27
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v91 = F_NameListToString(m, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v91
	F_errmsg(m, int32(72125), v10)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(499353), int32(3056), int32(434495))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lookup_ts_config_cache(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v335 int32
	_ = v335
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	v10 = m.G0
	v12 = v10 - int32(2560)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+2556)) = l0
	v16 = *(*int32)(unsafe.Add(mBase, _consts[887]))
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[888]))
	if v46 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+464)) = int64(85899345924)
	v25 = F_hash_create(m, int32(399197), int32(16), v12+int32(448), int32(40))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[887])) = v25
	F_CacheRegisterSyscacheCallback(m, int32(74), int32(1613), v25)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[887]))
	F_CacheRegisterSyscacheCallback(m, int32(72), int32(1613), v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	if v41 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L1
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L3
	} else {
		goto L112
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L3
	} else {
		goto L109
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L3
	} else {
		goto L106
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L3
	} else {
		goto L103
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L3
	} else {
		goto L100
	}
L14:
	;
	m.G0 = v12 + int32(2560)
	return v335
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[887]))
	v57 = int32(0)
	v59 = F_hash_search(m, v54, v12+int32(2556), v57, v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L20
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2556))
	if v49 != v50 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+4)))
	if v52 != 0 {
		v335 = v46
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[888])) = v324
	v335 = v324
	goto L14
L20:
	;
	if v59 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)))
	if v61 != 0 {
		v324 = v59
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2556))
	v64 = F_SearchSysCache1(m, int32(74), v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	if v64 == int32(0) {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+22)))
	v70 = v68 + v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+76))
	if v71 == int32(0) {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	if v59 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v128&int32(3) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L29:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[887]))
	v83 = F_hash_search(m, v77, v12+int32(2556), int32(1), v12+int32(448))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	if v85 == int32(0) {
		v128 = v59
		goto L28
	} else {
		goto L33
	}
L32:
	;
	v128 = v83
	goto L28
L33:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	if int32(0) < v88 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v92 = int32(0)
	v95 = v88
	goto L37
L35:
	;
	v123 = v85
	goto L36
L36:
	;
	F_pfree(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L44
	}
L37:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101+v92<<(uint(int32(3))%32))+4))
	if v105 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v123 = v113
	goto L36
L39:
	;
	F_pfree(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L42
	}
L40:
	;
	v109 = v95
	goto L41
L41:
	;
	v111 = v92 + int32(1)
	if v111 < v109 {
		v92 = v111
		v95 = v109
		goto L37
	} else {
		goto L43
	}
L42:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v109 = v108
	goto L41
L43:
	;
	goto L38
L44:
	;
	v128 = v59
	goto L28
L45:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2556))
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v70)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = v164
	F_ReleaseCatCache(m, v64)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L3
	} else {
		goto L54
	}
L46:
	;
	v140 = v128 + int32(20)
	if base.Ui32(v140) <= base.Ui32(v128) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v156 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v128)+4)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v128)+12)) = v156
	goto L45
L49:
	;
	v146 = v128 + int32(4)
	if base.Ui32(v146) < base.Ui32(v140) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v148 = v140
	goto L52
L51:
	;
	v148 = v146
	goto L52
L52:
	;
	v155 = F__emscripten_memset_bulkmem(m, v128, base.I32_extend8_s(int32(0)), (v128^int32(-1)+v148)&int32(-4)+int32(4))
	mBase = m.M
	goto L53
L53:
	;
	goto L45
L54:
	;
	v168 = int32(0)
	v174 = F__emscripten_memset_bulkmem(m, v12+int32(448), base.I32_extend8_s(v168), int32(2056))
	mBase = m.M
	goto L55
L55:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2556))
	F_ScanKeyInit(m, v12+int32(2508), int32(1), int32(3), int32(184), v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L3
	} else {
		goto L56
	}
L56:
	;
	v185 = F_table_open(m, int32(3603), int32(1))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L58
	}
L57:
	;
	F_systable_endscan_ordered(m, v195)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L3
	} else {
		goto L84
	}
L58:
	;
	v189 = F_index_open(m, int32(3609), int32(1))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	v195 = F_systable_beginscan_ordered(m, v185, v189, int32(0), int32(1), v12+int32(2508))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	v198 = F_systable_getnext_ordered(m, v195, int32(1))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	if v198 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v264 = int32(0)
	v268 = v168
	goto L57
L63:
	;
	goto L64
L64:
	;
	v204 = int32(0)
	v207 = v198
	v208 = v168
	goto L65
L65:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v207)+16))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+22)))
	v215 = v213 + v214
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if base.Ui32(v216-int32(257)) <= base.Ui32(int32(-257)) {
		goto L11
	} else {
		goto L67
	}
L66:
	;
	v264 = v260
	v268 = v259
	goto L57
L67:
	;
	if v216 < v208 {
		goto L10
	} else {
		goto L68
	}
L68:
	;
	if v208 < v216 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v262 = F_systable_getnext_ordered(m, v195, int32(1))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L3
	} else {
		goto L82
	}
L70:
	;
	if int32(0) < v204 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	if int32(100) <= v204 {
		goto L9
	} else {
		goto L81
	}
L73:
	;
	v229 = v12 + int32(448) + v208<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = v204
	v232 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	v234 = v204 << (uint(int32(2)) % 32)
	v235 = F_MemoryContextAlloc(m, v232, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L3
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v244
	v259 = v216
	v260 = int32(1)
	goto L69
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229)+4)) = v235
	if v234 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L75
L78:
	;
	v240 = F__emscripten_memcpy_bulkmem(m, v235, v12+int32(48), v234)
	mBase = m.M
	goto L80
L79:
	;
	goto L80
L80:
	;
	goto L77
L81:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(48)+v204<<(uint(int32(2))%32)))) = v254
	v259 = v208
	v260 = v204 + int32(1)
	goto L69
L82:
	;
	if v262 != 0 {
		v204 = v260
		v207 = v262
		v208 = v259
		goto L65
	} else {
		goto L83
	}
L83:
	;
	goto L66
L84:
	;
	F_relation_close(m, v189, int32(1))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	F_sequence_close(m, v185, int32(1))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L3
	} else {
		goto L86
	}
L86:
	;
	if int32(0) < v264 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v287 = v12 + int32(448) + v268<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = v264
	v290 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	v292 = v264 << (uint(int32(2)) % 32)
	v293 = F_MemoryContextAlloc(m, v290, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L3
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v320 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v128)+4)) = uint8(v320)
	v324 = v128
	goto L19
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+4)) = v293
	if v292 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v301 = v268 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v301
	v304 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	v307 = F_MemoryContextAlloc(m, v304, v301<<(uint(int32(3))%32))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L3
	} else {
		goto L95
	}
L92:
	;
	v298 = F__emscripten_memcpy_bulkmem(m, v293, v12+int32(48), v292)
	mBase = m.M
	goto L94
L93:
	;
	goto L94
L94:
	;
	goto L91
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+16)) = v307
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v314 = v312 << (uint(int32(3)) % 32)
	if v314 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L89
L97:
	;
	v315 = F__emscripten_memcpy_bulkmem(m, v307, v12+int32(448), v314)
	mBase = m.M
	goto L99
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2556))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v350
	F_errmsg_internal(m, int32(45752), v12)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L3
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(499087), int32(426), int32(398843))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L3
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2556))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v364
	F_errmsg_internal(m, int32(217624), v12+int32(16))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(499087), int32(433), int32(398843))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L3
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v216
	F_errmsg_internal(m, int32(402123), v12+int32(32))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L3
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(499087), int32(491), int32(398843))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L3
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errmsg_internal(m, int32(227122), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(499087), int32(493), int32(398843))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errmsg_internal(m, int32(368815), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L3
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(499087), int32(514), int32(398843))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L3
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ts_process_call(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16+v17<<(uint(int32(2))%32))))
	if v21 == v2 {
		v154 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(80)
	return v154
L2:
	;
	v25 = v21
	v27 = v17
	goto L3
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v34 != 0 {
		v77 = v25
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v154 = v2
	goto L1
L5:
	;
	if v27 == int32(0) {
		v154 = v2
		goto L1
	} else {
		goto L25
	}
L6:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v89 = F_palloc(m, v86+int32(1))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v35 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v39 = v27 + int32(1)
	v42 = v16 + v39<<(uint(int32(2))%32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v35 == v43 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v39
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v48 == int32(0) {
		v77 = v46
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v56 = v46 + int32(8)
	goto L11
L11:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v65 = v63 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	*(*int32)(unsafe.Add(mBase, uint32(v67+v65<<(uint(int32(2))%32)))) = v71
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	if v75 != 0 {
		v56 = v71 + int32(8)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v77 = v71
	goto L6
L13:
	;
	goto L12
L14:
	;
	return int32(0)
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v89
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	if v96 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v98+v99))) = uint8(v101)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v103
	v110 = F_pg_sprintf(m, v13+int32(48), int32(488080), v13+int32(16))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L14
	} else {
		goto L20
	}
L17:
	;
	v97 = F__emscripten_memcpy_bulkmem(m, v89, v77+int32(20), v96)
	mBase = m.M
	v98 = v97
	goto L19
L18:
	;
	v98 = v89
	goto L19
L19:
	;
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v13 + int32(48)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v115
	v120 = F_pg_sprintf(m, v13+int32(32), int32(488080), v13)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v13 + int32(32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v128 = F_BuildTupleFromCStrings(m, v125, v13+int32(68))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	v131 = F_HeapTupleHeaderGetDatum(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	F_pfree(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(0)
	v154 = v131
	goto L1
L25:
	;
	v143 = v27 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v143
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v16+v143<<(uint(int32(2))%32))))
	if v148 != 0 {
		v25 = v148
		v27 = v143
		goto L3
	} else {
		goto L26
	}
L26:
	;
	goto L4
}
func F_ts_setup_firstcall(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = l2
	v13 = int32(4515248)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v23 = F_palloc0(m, v18<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v23
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
		if v28 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
			if v30 == int32(0) {
			} else {
				v39 = v28 + int32(8)
				for {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
					v44 = v42 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v44
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					*(*int32)(unsafe.Add(mBase, uint32(v46+v44<<(uint(int32(2))%32)))) = v50
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
					if v54 != 0 {
						v39 = v50 + int32(8)
						continue
					} else {
						break
					}
					break
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
		}
		v67 = F_get_call_result_type(m, l0, int32(0), v10+int32(12))
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
			return
		} else {
			if v67 == int32(1) {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v71
				v73 = F_TupleDescGetAttInMetadata(m, v71)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v73
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v14
					m.G0 = v10 + int32(16)
					return
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(367227), int32(0))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return
					} else {
						F_errfinish(m, int32(495627), int32(2481), int32(304655))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
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
}
