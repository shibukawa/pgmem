package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgstat_archiver_snapshot_cb(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v50 int32
	_ = v50
	var v56 int64
	_ = v56
	var v62 int32
	_ = v62
	v10 = *(*int32)(unsafe.Add(mBase, _consts[562]))
	v14 = v10 + int32(24)
	goto L1
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v26 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v26 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v37 = F_LWLockAcquire(m, v14, int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L14
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	goto L9
L6:
	;
	return
L7:
	;
	goto L5
L8:
	;
	if v24&int32(1) != 0 {
		goto L1
	} else {
		goto L12
	}
L9:
	;
	v30 = F__emscripten_memcpy_bulkmem(m, int32(4367792), v10+int32(48), int32(136))
	mBase = m.M
	goto L11
L11:
	;
	goto L8
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	if v24 != v34 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L2
L14:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v10)+248))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v10)+184))
	F_LWLockRelease(m, v14)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v44 = *(*int64)(unsafe.Add(mBase, _consts[1019]))
	if v44 == v40 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1020])) = int64(0)
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1021])) = uint8(v50)
	goto L18
L17:
	;
	goto L18
L18:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1019])) = v44 - v40
	v56 = *(*int64)(unsafe.Add(mBase, _consts[1022]))
	if v39 == v56 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1023])) = int64(0)
	v62 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1024])) = uint8(v62)
	goto L21
L20:
	;
	goto L21
L21:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1022])) = v56 - v39
	return
}
func F_pgstat_beshutdown_hook(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	v3 = int32(4438516)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v6 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v5 + v6
	v9 = int32(4367704)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v11 + v6
	v15 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v11 + int32(2)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v23 - v6
	*(*int32)(unsafe.Add(mBase, _consts[48])) = v15
	return
}
func F_pgstat_clear_backend_activity_snapshot(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, _consts[1015]))
	if v3 != 0 {
		F_MemoryContextDelete(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[1015])) = int32(0)
			v10 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[1016])) = v10
			*(*int32)(unsafe.Add(mBase, _consts[1017])) = v10
			return
		}
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[1016])) = v10
		*(*int32)(unsafe.Add(mBase, _consts[1017])) = v10
		return
	}
}
func F_pgstat_cmp_hash_key(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	v5 = int32(16)
	goto L4
L1:
	;
	return v67
L2:
	;
	v67 = int32(0)
	goto L1
L3:
	;
	v41 = v36
	v42 = v37
	v43 = v38
	goto L13
L4:
	;
	if (l0|l1)&int32(3) != 0 {
		v36 = l0
		v37 = l1
		v38 = v5
		goto L3
	} else {
		goto L7
	}
L6:
	;
	if v26 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v13 = l0
	v14 = l1
	v15 = v5
	goto L8
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v18 != v19 {
		v36 = v13
		v37 = v14
		v38 = v15
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v21 = int32(4)
	v22 = v14 + v21
	v24 = v13 + v21
	v26 = v15 - v21
	if base.Ui32(int32(3)) < base.Ui32(v26) {
		v13 = v24
		v14 = v22
		v15 = v26
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v36 = v24
	v37 = v22
	v38 = v26
	goto L3
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v46 == v47 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v67 = v46 - v47
	goto L1
L15:
	;
	v49 = int32(1)
	v54 = v43 - v49
	if v54 != 0 {
		v41 = v41 + v49
		v42 = v42 + v49
		v43 = v54
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	goto L2
}
func F_pgstat_fetch_stat_tabentry_ext(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v6 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	if l0 != 0 {
		v7 = int32(0)
	} else {
		v7 = v6
	}
	v9 = F_pgstat_fetch_entry(m, int32(2), v7, base.I64_extend_i32_u(l1))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_pgstat_flush_io(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = F_pgstat_io_flush_cb(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_pgstat_function_flush_cb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = F_pgstat_lock_entry(m, l0, l1)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v5)+24))
			v12 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = v11 + v12
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v5)+32))
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
			v17 = int64(1000)
			v18 = base.I64_div_s(v16, v17)
			*(*int64)(unsafe.Add(mBase, uint32(v5)+32)) = v15 + v18
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v5)+40))
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v6)+16))
			v24 = base.I64_div_s(v22, v17)
			*(*int64)(unsafe.Add(mBase, uint32(v5)+40)) = v21 + v24
			F_pgstat_unlock_entry(m, l0)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				return v7
			}
		} else {
			return v7
		}
	}
}
func F_pgstat_get_local_beentry_by_index(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	F_pgstat_read_current_status(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[1017]))
		v8 = int32(432)
		v15 = *(*int32)(unsafe.Add(mBase, _consts[1016]))
		if l0 <= v15 {
			v17 = v7 + l0*v8 - v8
		} else {
			v17 = int32(0)
		}
		v18 = int32(0)
		if v18 < l0 {
			v21 = v17
		} else {
			v21 = v18
		}
		return v21
	}
}
func F_pgstat_get_xact_stack_level(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1031]))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		if v5 == l0 {
			v29 = v4
			return v29
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, _consts[319]))
			v10 = F_MemoryContextAlloc(m, v8, int32(24))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v14 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v14
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				v18 = v10 + int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v18
				v21 = int32(4427088)
				v22 = *(*int32)(unsafe.Add(mBase, _consts[1031]))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v14
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v22
				*(*int32)(unsafe.Add(mBase, _consts[1031])) = v10
				v29 = v10
				return v29
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[319]))
		v10 = F_MemoryContextAlloc(m, v8, int32(24))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
			v18 = v10 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v18
			v21 = int32(4427088)
			v22 = *(*int32)(unsafe.Add(mBase, _consts[1031]))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v22
			*(*int32)(unsafe.Add(mBase, _consts[1031])) = v10
			v29 = v10
			return v29
		}
	}
}
func F_pgstat_prepare_report_checksum_failure(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v2 = int32(1)
	v6 = F_pgstat_get_entry_ref(m, v2, l0, int64(0), v2, int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_pgstat_progress_update_multi_param(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v104 int64
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	v4 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v11 == v4 {
	} else {
		if l0 == int32(0) {
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, _consts[49])))
			if v17&int32(1) == int32(0) {
			} else {
				v22 = int32(4438516)
				v24 = *(*int32)(unsafe.Add(mBase, _consts[13]))
				v25 = int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[13])) = v24 + v25
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v28 + v25
				if l0 <= int32(0) {
				} else {
					v35 = l0 & int32(3)
					v37 = v11 + int32(232)
					if base.Ui32(int32(4)) <= base.Ui32(l0) {
						v43 = int32(0)
						v46 = v4
						for {
							v52 = int32(2)
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l1+v46<<(uint(v52)%32))))
							v56 = int32(3)
							v62 = *(*int64)(unsafe.Add(mBase, uint32(l2+v46<<(uint(v56)%32))))
							*(*int64)(unsafe.Add(mBase, uint32(v37+v55<<(uint(v56)%32)))) = v62
							v65 = v46 | int32(1)
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l1+v65<<(uint(v52)%32))))
							v76 = *(*int64)(unsafe.Add(mBase, uint32(l2+v65<<(uint(v56)%32))))
							*(*int64)(unsafe.Add(mBase, uint32(v37+v69<<(uint(v56)%32)))) = v76
							v79 = v46 | v52
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l1+v79<<(uint(v52)%32))))
							v90 = *(*int64)(unsafe.Add(mBase, uint32(l2+v79<<(uint(v56)%32))))
							*(*int64)(unsafe.Add(mBase, uint32(v37+v83<<(uint(v56)%32)))) = v90
							v93 = v46 | v56
							v97 = *(*int32)(unsafe.Add(mBase, uint32(l1+v93<<(uint(v52)%32))))
							v104 = *(*int64)(unsafe.Add(mBase, uint32(l2+v93<<(uint(v56)%32))))
							*(*int64)(unsafe.Add(mBase, uint32(v37+v97<<(uint(v56)%32)))) = v104
							v106 = int32(4)
							v107 = v46 + v106
							v109 = v43 + v106
							if v109 != l0&int32(2147483644) {
								v43 = v109
								v46 = v107
								continue
							} else {
								break
							}
							break
						}
						v114 = v107
					} else {
						v114 = v4
					}
					if v35 == int32(0) {
					} else {
						v123 = int32(0)
						v126 = v114
						for {
							v135 = *(*int32)(unsafe.Add(mBase, uint32(l1+v126<<(uint(int32(2))%32))))
							v136 = int32(3)
							v142 = *(*int64)(unsafe.Add(mBase, uint32(l2+v126<<(uint(v136)%32))))
							*(*int64)(unsafe.Add(mBase, uint32(v37+v135<<(uint(v136)%32)))) = v142
							v144 = int32(1)
							v147 = v123 + v144
							if v147 != v35 {
								v123 = v147
								v126 = v126 + v144
								continue
							} else {
								break
							}
							break
						}
					}
				}
				v158 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v159 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v158 + v159
				v162 = int32(4438516)
				v164 = *(*int32)(unsafe.Add(mBase, _consts[13]))
				*(*int32)(unsafe.Add(mBase, _consts[13])) = v164 - v159
			}
		}
	}
	return
}
func F_pgstat_read_current_status(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	v1 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[1017]))
	if v15 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1015]))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	v29 = F_AllocSetContextCreateInternal(m, v24, int32(83102), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v32 = v19
	goto L6
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[660]))
	v39 = F_MemoryContextAlloc(m, v32, v34*int32(432)+int32(16416))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1015])) = v29
	v32 = v29
	goto L6
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[1015]))
	v44 = *(*int32)(unsafe.Add(mBase, _consts[660]))
	v49 = F_MemoryContextAlloc(m, v42, v44<<(uint(int32(6))%32)+int32(2432))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[1015]))
	v54 = *(*int32)(unsafe.Add(mBase, _consts[660]))
	v59 = F_MemoryContextAlloc(m, v52, v54<<(uint(int32(6))%32)+int32(2432))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[1015]))
	v64 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	v66 = *(*int32)(unsafe.Add(mBase, _consts[660]))
	v70 = F_MemoryContextAllocHuge(m, v62, v64*(v66+int32(38)))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1016])) = v73
	v76 = *(*int32)(unsafe.Add(mBase, _consts[660]))
	if v73 < v76+int32(38) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[1018]))
	v83 = v82
	v85 = v39
	v87 = v49
	v88 = v59
	v89 = v70
	v90 = v1
	goto L16
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1017])) = v39
	goto L3
L16:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v97
	if int32(0) < v97 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	goto L22
L19:
	;
	goto L20
L20:
	;
	v335 = int32(0)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if base.B2i32(v96&int32(1) == v335)&base.B2i32(v337 == v96) == v335 {
		goto L88
	} else {
		goto L89
	}
L21:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v83)+212))
	if (v104^v87)&int32(3) != 0 {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	v102 = F__emscripten_memcpy_bulkmem(m, v85, v83, int32(408))
	mBase = m.M
	goto L24
L24:
	;
	goto L21
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+212)) = v87
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v83)+188))
	if (v180^v88)&int32(3) != 0 {
		goto L49
	} else {
		goto L50
	}
L26:
	;
	goto L25
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v159))) = uint8(v158)
	if v158&int32(255) == int32(0) {
		goto L26
	} else {
		goto L42
	}
L28:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v157 = v104
	v158 = v110
	v159 = v87
	goto L27
L29:
	;
	goto L30
L30:
	;
	if v104&int32(3) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v114 = v104
	v116 = v87
	goto L34
L32:
	;
	v128 = v104
	v130 = v87
	goto L33
L33:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v135 = int32(-2139062144)
	if (int32(16843008)-v132|v132)&v135 != v135 {
		v157 = v128
		v158 = v132
		v159 = v130
		goto L27
	} else {
		goto L38
	}
L34:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	*(*uint8)(unsafe.Add(mBase, uint32(v116))) = uint8(v117)
	if v117 == int32(0) {
		goto L26
	} else {
		goto L36
	}
L35:
	;
	v128 = v124
	v130 = v122
	goto L33
L36:
	;
	v121 = int32(1)
	v122 = v116 + v121
	v124 = v114 + v121
	if v124&int32(3) != 0 {
		v114 = v124
		v116 = v122
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v140 = v128
	v141 = v132
	v142 = v130
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v141
	v144 = int32(4)
	v145 = v142 + v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	v148 = v140 + v144
	v152 = int32(-2139062144)
	if (v146|(int32(16843008)-v146))&v152 == v152 {
		v140 = v148
		v141 = v146
		v142 = v145
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v157 = v148
	v158 = v146
	v159 = v145
	goto L27
L41:
	;
	goto L40
L42:
	;
	v166 = v157
	v168 = v159
	goto L43
L43:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)) = uint8(v169)
	v171 = int32(1)
	if v169 != 0 {
		v166 = v166 + v171
		v168 = v168 + v171
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L26
L45:
	;
	goto L44
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+188)) = v88
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v83)+216))
	if (v256^v89)&int32(3) != 0 {
		goto L70
	} else {
		goto L71
	}
L47:
	;
	goto L46
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v235))) = uint8(v234)
	if v234&int32(255) == int32(0) {
		goto L47
	} else {
		goto L63
	}
L49:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	v233 = v180
	v234 = v186
	v235 = v88
	goto L48
L50:
	;
	goto L51
L51:
	;
	if v180&int32(3) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v190 = v180
	v192 = v88
	goto L55
L53:
	;
	v204 = v180
	v206 = v88
	goto L54
L54:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v211 = int32(-2139062144)
	if (int32(16843008)-v208|v208)&v211 != v211 {
		v233 = v204
		v234 = v208
		v235 = v206
		goto L48
	} else {
		goto L59
	}
L55:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	*(*uint8)(unsafe.Add(mBase, uint32(v192))) = uint8(v193)
	if v193 == int32(0) {
		goto L47
	} else {
		goto L57
	}
L56:
	;
	v204 = v200
	v206 = v198
	goto L54
L57:
	;
	v197 = int32(1)
	v198 = v192 + v197
	v200 = v190 + v197
	if v200&int32(3) != 0 {
		v190 = v200
		v192 = v198
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v216 = v204
	v217 = v208
	v218 = v206
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218))) = v217
	v220 = int32(4)
	v221 = v218 + v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v224 = v216 + v220
	v228 = int32(-2139062144)
	if (v222|(int32(16843008)-v222))&v228 == v228 {
		v216 = v224
		v217 = v222
		v218 = v221
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v233 = v224
	v234 = v222
	v235 = v221
	goto L48
L62:
	;
	goto L61
L63:
	;
	v242 = v233
	v244 = v235
	goto L64
L64:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v244)+1)) = uint8(v245)
	v247 = int32(1)
	if v245 != 0 {
		v242 = v242 + v247
		v244 = v244 + v247
		goto L64
	} else {
		goto L66
	}
L65:
	;
	goto L47
L66:
	;
	goto L65
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+216)) = v89
	goto L20
L68:
	;
	goto L67
L69:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v311))) = uint8(v310)
	if v310&int32(255) == int32(0) {
		goto L68
	} else {
		goto L84
	}
L70:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256))))
	v309 = v256
	v310 = v262
	v311 = v89
	goto L69
L71:
	;
	goto L72
L72:
	;
	if v256&int32(3) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v266 = v256
	v268 = v89
	goto L76
L74:
	;
	v280 = v256
	v282 = v89
	goto L75
L75:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v287 = int32(-2139062144)
	if (int32(16843008)-v284|v284)&v287 != v287 {
		v309 = v280
		v310 = v284
		v311 = v282
		goto L69
	} else {
		goto L80
	}
L76:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	*(*uint8)(unsafe.Add(mBase, uint32(v268))) = uint8(v269)
	if v269 == int32(0) {
		goto L68
	} else {
		goto L78
	}
L77:
	;
	v280 = v276
	v282 = v274
	goto L75
L78:
	;
	v273 = int32(1)
	v274 = v268 + v273
	v276 = v266 + v273
	if v276&int32(3) != 0 {
		v266 = v276
		v268 = v274
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v292 = v280
	v293 = v284
	v294 = v282
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v293
	v296 = int32(4)
	v297 = v294 + v296
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	v300 = v292 + v296
	v304 = int32(-2139062144)
	if (v298|(int32(16843008)-v298))&v304 == v304 {
		v292 = v300
		v293 = v298
		v294 = v297
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v309 = v300
	v310 = v298
	v311 = v297
	goto L69
L83:
	;
	goto L82
L84:
	;
	v318 = v309
	v320 = v311
	goto L85
L85:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v320)+1)) = uint8(v321)
	v323 = int32(1)
	if v321 != 0 {
		v318 = v318 + v323
		v320 = v320 + v323
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L68
L87:
	;
	goto L86
L88:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v343 == int32(0) {
		goto L16
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if int32(0) < v348 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L7
	} else {
		goto L92
	}
L92:
	;
	goto L16
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+408)) = v90
	v353 = v85 + int32(412)
	v354 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = v354
	v357 = v85 + int32(416)
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = v354
	v361 = v85 + int32(420)
	*(*int32)(unsafe.Add(mBase, uint32(v361))) = v354
	v365 = v85 + int32(424)
	*(*uint8)(unsafe.Add(mBase, uint32(v365))) = uint8(v354)
	if v90 < v354 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v418 = v85
	v420 = v87
	v421 = v88
	v422 = v89
	goto L95
L95:
	;
	v430 = v90 + int32(1)
	v432 = *(*int32)(unsafe.Add(mBase, _consts[660]))
	if v430 < v432+int32(38) {
		v83 = v83 + int32(408)
		v85 = v418
		v87 = v420
		v88 = v421
		v89 = v422
		v90 = v430
		goto L16
	} else {
		goto L104
	}
L96:
	;
	v402 = int32(4367736)
	v404 = *(*int32)(unsafe.Add(mBase, _consts[1016]))
	*(*int32)(unsafe.Add(mBase, _consts[1016])) = v404 + int32(1)
	v408 = int32(-64)
	v413 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	v418 = v85 + int32(432)
	v420 = v87 - v408
	v421 = v88 - v408
	v422 = v89 + v413
	goto L95
L97:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)+16))
	if base.Ui32(v372) <= base.Ui32(v90) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
	v376 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v380 = F_LWLockAcquire(m, v376+int32(512), int32(1))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L7
	} else {
		goto L99
	}
L99:
	;
	v384 = v374 + v90*int32(640)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v384)+44))
	if v385 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v384)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v384)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = v388
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+276)))
	*(*int32)(unsafe.Add(mBase, uint32(v361))) = v390
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+277)))
	*(*uint8)(unsafe.Add(mBase, uint32(v365))) = uint8(v392)
	goto L102
L101:
	;
	goto L102
L102:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v395+int32(512))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	goto L96
L104:
	;
	goto L17
}
func F_pgstat_replslot_from_serialized_name_cb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	v3 = int32(0)
	v5 = F_SearchNamedReplicationSlot(m, l0, int32(1))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v24 = v3
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[561]))
			v15 = base.I32_div_s(v5-v12, int32(288))
			if v15 == int32(-1) {
				v24 = v3
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(4)
				*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = base.I64_extend_i32_s(v15)
				v24 = int32(1)
			}
		}
		return v24
	}
}
func F_pgstat_report_checksum_failures_in_db(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v65 int32
	_ = v65
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[6])))
	if v10 != int32(1) {
		m.G0 = v7 + int32(16)
		return
	} else {
		v15 = int32(0)
		v17 = F_pgstat_get_entry_ref(m, int32(1), l0, int64(0), v15, v15)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			if v17 == int32(0) {
				v23 = F_errstart(m, int32(19), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					if v23 == int32(0) {
						m.G0 = v7 + int32(16)
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
						F_errmsg_internal(m, int32(46971), v7)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_errfinish(m, int32(477011), int32(194), int32(481364))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v38 = F_pgstat_lock_entry(m, v17, int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
					v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)+176))
					*(*int64)(unsafe.Add(mBase, uint32(v40)+176)) = v41 + base.I64_extend_i32_s(l1)
					v48 = m.G0
					v49 = int32(16)
					v50 = v48 - v49
					m.G0 = v50
					F___gettimeofday(m, v50)
					mBase = m.M
					v53 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
					v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v50)+8)))
					m.G0 = v50 + v49
					*(*int64)(unsafe.Add(mBase, uint32(v40)+184)) = v54 + v53*int64(1000000) - int64(946684800000000)
					F_pgstat_unlock_entry(m, v17)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_pgstat_report_connect(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int64
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, _consts[19]))
	if v3 == int32(1) {
		v8 = *(*int64)(unsafe.Add(mBase, _consts[449]))
		*(*int64)(unsafe.Add(mBase, _consts[1025])) = v8
		v12 = *(*int32)(unsafe.Add(mBase, _consts[108]))
		v15 = F_pgstat_prep_pending_entry(m, int32(1), v12, int64(0), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)+184))
			*(*int64)(unsafe.Add(mBase, uint32(v17)+184)) = v18 + int64(1)
			return
		}
	} else {
		return
	}
}
func F_pgstat_report_recovery_conflict(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[6])))
	if v4 != int32(1) {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[108]))
		v12 = F_pgstat_prep_pending_entry(m, int32(1), v9, int64(0), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v15 = l0 - int32(8)
			if base.Ui32(int32(5)) < base.Ui32(v15) {
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
				v21 = v18 + v15<<(uint(int32(3))%32)
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)+80))
				*(*int64)(unsafe.Add(mBase, uint32(v21)+80)) = v22 + int64(1)
			}
			return
		}
	}
}
func F_pgstat_report_wal(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = l0 ^ int32(1)
	v4 = F_pgstat_wal_flush_cb(m, v3)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v7 = F_pgstat_flush_backend(m, v3, int32(2))
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_pgstat_flush_io(m, v3)
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				v12 = F_pgstat_flush_backend(m, v3, int32(1))
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_pgstat_set_wait_event_storage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[96])) = l0
	return
}
func F_pgstat_slru_flush_cb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	v2 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[90])))
	if v7 == v2 {
		v97 = v2
		return v97
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[562]))
		v13 = v11 + int32(52744)
		if l0 == int32(0) {
			v17 = F_LWLockAcquire(m, v13, int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v33 = int32(0)
				for {
					v37 = v33 << (uint(int32(6)) % 32)
					v38 = v11 + int32(52760) + v37
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
					v42 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[1026])))
					*(*int64)(unsafe.Add(mBase, uint32(v38))) = v39 + v42
					v45 = *(*int64)(unsafe.Add(mBase, uint32(v38)+8))
					v48 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[91])))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v45 + v48
					v51 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
					v54 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[1027])))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v51 + v54
					v57 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
					v60 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[1028])))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v57 + v60
					v63 = *(*int64)(unsafe.Add(mBase, uint32(v38)+32))
					v66 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[92])))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+32)) = v63 + v66
					v69 = *(*int64)(unsafe.Add(mBase, uint32(v38)+40))
					v72 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[1029])))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+40)) = v69 + v72
					v75 = *(*int64)(unsafe.Add(mBase, uint32(v38)+48))
					v78 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[1030])))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+48)) = v75 + v78
					v82 = v33 + int32(1)
					if v82 != int32(8) {
						v33 = v82
						continue
					} else {
						break
					}
					break
				}
				v85 = int32(0)
				v90 = F__emscripten_memset_bulkmem(m, int32(4426544), base.I32_extend8_s(v85), int32(512))
				mBase = m.M
				F_LWLockRelease(m, v13)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					v94 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _consts[90])) = uint8(v94)
					v97 = v85
					return v97
				}
			}
		} else {
			v23 = F_LWLockConditionalAcquire(m, v13, int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					v97 = int32(1)
					return v97
				} else {
					v33 = int32(0)
					for {
						v37 = v33 << (uint(int32(6)) % 32)
						v38 = v11 + int32(52760) + v37
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
						v42 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[1026])))
						*(*int64)(unsafe.Add(mBase, uint32(v38))) = v39 + v42
						v45 = *(*int64)(unsafe.Add(mBase, uint32(v38)+8))
						v48 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[91])))
						*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v45 + v48
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
						v54 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[1027])))
						*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v51 + v54
						v57 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
						v60 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[1028])))
						*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v57 + v60
						v63 = *(*int64)(unsafe.Add(mBase, uint32(v38)+32))
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[92])))
						*(*int64)(unsafe.Add(mBase, uint32(v38)+32)) = v63 + v66
						v69 = *(*int64)(unsafe.Add(mBase, uint32(v38)+40))
						v72 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[1029])))
						*(*int64)(unsafe.Add(mBase, uint32(v38)+40)) = v69 + v72
						v75 = *(*int64)(unsafe.Add(mBase, uint32(v38)+48))
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[1030])))
						*(*int64)(unsafe.Add(mBase, uint32(v38)+48)) = v75 + v78
						v82 = v33 + int32(1)
						if v82 != int32(8) {
							v33 = v82
							continue
						} else {
							break
						}
						break
					}
					v85 = int32(0)
					v90 = F__emscripten_memset_bulkmem(m, int32(4426544), base.I32_extend8_s(v85), int32(512))
					mBase = m.M
					F_LWLockRelease(m, v13)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						v94 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _consts[90])) = uint8(v94)
						v97 = v85
						return v97
					}
				}
			}
		}
	}
}
func F_pgstat_slru_snapshot_cb(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, _consts[562]))
	v6 = v4 + int32(52744)
	v8 = F_LWLockAcquire(m, v6, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v14 = F__emscripten_memcpy_bulkmem(m, int32(4419896), v4+int32(52760), int32(512))
		mBase = m.M
		F_LWLockRelease(m, v6)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			return
		}
	}
}
func F_pgstat_tracks_backend_bktype(m *base.Module, l0 int32) int32 {
	return int32(base.Ui32(int32(115186))>>(uint(l0)%32)) & base.B2i32(base.Ui32(l0) < base.Ui32(int32(17)))
}
