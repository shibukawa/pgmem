package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgstat_archiver_snapshot_cb(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v45 int32
	_ = v45
	var v51 int64
	_ = v51
	var v57 int32
	_ = v57
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_archiver_snapshot_cb[0]))
	goto L1
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_archiver_snapshot_cb[1]))
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v30 = v8 + int32(24)
	v32 = F_LWLockAcquire(m, v30, int32(1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
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
	base.MemoryCopy(m, int32(_a_F_pgstat_archiver_snapshot_cb_0), v8+int32(48), int32(136))
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
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	if v17 != v27 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L2
L10:
	;
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v8)+248))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v8)+184))
	F_LWLockRelease(m, v30)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v39 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_archiver_snapshot_cb[2]))
	if v39 == v35 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_archiver_snapshot_cb[3])) = int64(0)
	v45 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_archiver_snapshot_cb[4])) = uint8(v45)
	goto L14
L13:
	;
	goto L14
L14:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_archiver_snapshot_cb[2])) = v39 - v35
	v51 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_archiver_snapshot_cb[5]))
	if v34 == v51 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_archiver_snapshot_cb[6])) = int64(0)
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_archiver_snapshot_cb[7])) = uint8(v57)
	goto L17
L16:
	;
	goto L17
L17:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_archiver_snapshot_cb[5])) = v51 - v34
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
	v3 = int32(_a_F_pgstat_beshutdown_hook_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_beshutdown_hook[0]))
	v6 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_beshutdown_hook[0])) = v5 + v6
	v9 = int32(_a_F_pgstat_beshutdown_hook_1)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_beshutdown_hook[1]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v11 + v6
	v15 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v11 + int32(2)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_beshutdown_hook[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_beshutdown_hook[0])) = v23 - v6
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_beshutdown_hook[1])) = v15
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
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_clear_backend_activity_snapshot[0]))
	if v3 != 0 {
		F_MemoryContextDelete(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_clear_backend_activity_snapshot[0])) = int32(0)
			v10 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_clear_backend_activity_snapshot[1])) = v10
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_clear_backend_activity_snapshot[2])) = v10
			return
		}
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_pgstat_clear_backend_activity_snapshot[1])) = v10
		*(*int32)(unsafe.Add(mBase, _c_F_pgstat_clear_backend_activity_snapshot[2])) = v10
		return
	}
}
func F_pgstat_cmp_hash_key(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v8 = int64(56)
	v10 = int64(65280)
	v12 = int64(40)
	v15 = int64(16711680)
	v17 = int64(24)
	v19 = int64(4278190080)
	v21 = int64(8)
	v42 = v7<<(uint(v8)%64) | v7&v10<<(uint(v12)%64) | (v7&v15<<(uint(v17)%64) | v7&v19<<(uint(v21)%64)) | (int64(base.Ui64(v7)>>(uint(v21)%64))&v19 | int64(base.Ui64(v7)>>(uint(v17)%64))&v15 | (int64(base.Ui64(v7)>>(uint(v12)%64))&v10 | int64(base.Ui64(v7)>>(uint(v8)%64))))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v78 = v43<<(uint(v8)%64) | v43&v10<<(uint(v12)%64) | (v43&v15<<(uint(v17)%64) | v43&v19<<(uint(v21)%64)) | (int64(base.Ui64(v43)>>(uint(v21)%64))&v19 | int64(base.Ui64(v43)>>(uint(v17)%64))&v15 | (int64(base.Ui64(v43)>>(uint(v12)%64))&v10 | int64(base.Ui64(v43)>>(uint(v8)%64))))
	if v42 == v78 {
		v81 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		v82 = int64(56)
		v84 = int64(65280)
		v86 = int64(40)
		v89 = int64(16711680)
		v91 = int64(24)
		v93 = int64(4278190080)
		v95 = int64(8)
		v116 = v81<<(uint(v82)%64) | v81&v84<<(uint(v86)%64) | (v81&v89<<(uint(v91)%64) | v81&v93<<(uint(v95)%64)) | (int64(base.Ui64(v81)>>(uint(v95)%64))&v93 | int64(base.Ui64(v81)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v81)>>(uint(v86)%64))&v84 | int64(base.Ui64(v81)>>(uint(v82)%64))))
		v117 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		v152 = v117<<(uint(v82)%64) | v117&v84<<(uint(v86)%64) | (v117&v89<<(uint(v91)%64) | v117&v93<<(uint(v95)%64)) | (int64(base.Ui64(v117)>>(uint(v95)%64))&v93 | int64(base.Ui64(v117)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v117)>>(uint(v86)%64))&v84 | int64(base.Ui64(v117)>>(uint(v82)%64))))
		if v116 == v152 {
			v162 = int32(0)
		} else {
			v154 = v152
			v155 = v116
			if base.Ui64(v155) < base.Ui64(v154) {
				v159 = int32(-1)
			} else {
				v159 = int32(1)
			}
			v162 = v159
		}
	} else {
		v154 = v78
		v155 = v42
		if base.Ui64(v155) < base.Ui64(v154) {
			v159 = int32(-1)
		} else {
			v159 = int32(1)
		}
		v162 = v159
	}
	return v162
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
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_stat_tabentry_ext[0]))
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
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_local_beentry_by_index[0]))
		v8 = int32(432)
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_local_beentry_by_index[1]))
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_xact_stack_level[0]))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		if v5 == l0 {
			v29 = v4
			return v29
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_xact_stack_level[1]))
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
				v21 = int32(_a_F_pgstat_get_xact_stack_level_0)
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_xact_stack_level[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v14
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v22
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_xact_stack_level[0])) = v10
				v29 = v10
				return v29
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_xact_stack_level[1]))
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
			v21 = int32(_a_F_pgstat_get_xact_stack_level_0)
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_xact_stack_level[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v22
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_xact_stack_level[0])) = v10
			v29 = v10
			return v29
		}
	}
}
func F_pgstat_gist_page(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int64
	_ = v99
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v110 int64
	_ = v110
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	v5 = int32(0)
	v9 = F_ReadBufferExtended(m, l1, v5, l2, v5, l3)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		F_LockBuffer(m, v9, int32(1))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if v9 < int32(0) {
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_gist_page[0]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v17+(v9^int32(-1))<<(uint(int32(2))%32))))
				v31 = v23
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_gist_page[1]))
				v31 = v25 + v9<<(uint(int32(13))%32) + int32(-8192)
			}
			v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+14)))
			if v32 == int32(0) {
				v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v35 - int64(-8192)
				F_UnlockReleaseBuffer(m, v9)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					return
				}
			} else {
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+19)))
				v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+16)))
				if (v41<<(uint(int32(8))%32)-v44)&int32(_a_F_pgstat_gist_page_0) != int32(16) {
				} else {
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v44)+12)))
					if v51&int32(1) == int32(0) {
					} else {
						v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+12)))
						v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+14)))
						v59 = v57 - v56
						v60 = int32(0)
						if v60 < v59 {
							v63 = v59
						} else {
							v63 = v60
						}
						v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v64 + base.I64_extend_i32_u(v63)
						if base.Ui32(v56) < base.Ui32(int32(25)) {
						} else {
							v71 = v56 + int32(_a_F_pgstat_gist_page_1)
							if v71&int32(_a_F_pgstat_gist_page_2) == int32(0) {
							} else {
								v80 = int32(1)
								v88 = v80
								for {
									v93 = v31 + int32(20) + v88<<(uint(int32(2))%32)
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
									v95 = int32(_a_F_pgstat_gist_page_3)
									if v94&v95 == v95 {
										v99 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v99 + int64(1)
										v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
										v104 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v103 + base.I64_extend_i32_u(int32(base.Ui32(v104)>>(uint(int32(17))%32)))
									} else {
										v110 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v110 + int64(1)
										v114 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v114 + base.I64_extend_i32_u(int32(base.Ui32(v115)>>(uint(int32(17))%32)))
									}
									v122 = v88 + int32(1)
									if v122 != (int32(base.Ui32(v71)>>(uint(int32(2))%32))+v80)&int32(_a_F_pgstat_gist_page_0) {
										v88 = v122
										continue
									} else {
										break
									}
									break
								}
							}
						}
					}
				}
				F_UnlockReleaseBuffer(m, v9)
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					return
				}
			}
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
	var v13 int32
	_ = v13
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
	var v116 int32
	_ = v116
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
	if l0 == v4 {
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_update_multi_param[0]))
		if v13 == int32(0) {
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_progress_update_multi_param[1])))
			if v17&int32(1) == int32(0) {
			} else {
				v22 = int32(_a_F_pgstat_progress_update_multi_param_0)
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_update_multi_param[2]))
				v25 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_update_multi_param[2])) = v24 + v25
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v28 + v25
				if l0 <= int32(0) {
				} else {
					v35 = l0 & int32(3)
					v37 = v13 + int32(232)
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
						if v35 == int32(0) {
						} else {
							v116 = v107
							v123 = int32(0)
							v126 = v116
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
					} else {
						v116 = v4
						v123 = int32(0)
						v126 = v116
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
				v158 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v159 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v158 + v159
				v162 = int32(_a_F_pgstat_progress_update_multi_param_0)
				v164 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_update_multi_param[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_update_multi_param[2])) = v164 - v159
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
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
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
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	v1 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[0]))
	if v14 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[1]))
	if v18 == int32(0) {
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
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[2]))
	v28 = F_AllocSetContextCreateInternal(m, v23, int32(_a_F_pgstat_read_current_status_0), int32(0), int32(1024), int32(_a_F_pgstat_read_current_status_1))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v31 = v18
	goto L6
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[3]))
	v38 = F_MemoryContextAlloc(m, v31, v33*int32(432)+int32(_a_F_pgstat_read_current_status_2))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[1])) = v28
	v31 = v28
	goto L6
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[1]))
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[3]))
	v48 = F_MemoryContextAlloc(m, v41, v43<<(uint(int32(6))%32)+int32(2432))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[1]))
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[3]))
	v58 = F_MemoryContextAlloc(m, v51, v53<<(uint(int32(6))%32)+int32(2432))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[1]))
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[4]))
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[3]))
	v69 = F_MemoryContextAllocHuge(m, v61, v63*(v65+int32(38)))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v72 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[5])) = v72
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[3]))
	if v72 < v75+int32(38) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[6]))
	v82 = v38
	v83 = v81
	v85 = v1
	v87 = v48
	v88 = v58
	v89 = v69
	goto L16
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[0])) = v38
	goto L3
L16:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v95
	if int32(0) < v95 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	base.MemoryCopy(m, v82, v83, int32(408))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v83)+212))
	if (v101^v87)&int32(3) != 0 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	goto L20
L20:
	;
	v331 = int32(0)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if base.B2i32(v94&int32(1) == v331)&base.B2i32(v333 == v94) == v331 {
		goto L84
	} else {
		goto L85
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+212)) = v87
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v83)+188))
	if (v177^v88)&int32(3) != 0 {
		goto L45
	} else {
		goto L46
	}
L22:
	;
	goto L21
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v155)
	if v155&int32(255) == int32(0) {
		goto L22
	} else {
		goto L38
	}
L24:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	v154 = v101
	v155 = v107
	v156 = v87
	goto L23
L25:
	;
	goto L26
L26:
	;
	if v101&int32(3) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v111 = v101
	v113 = v87
	goto L30
L28:
	;
	v125 = v101
	v127 = v87
	goto L29
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v132 = int32(-2139062144)
	if (int32(16843008)-v129|v129)&v132 != v132 {
		v154 = v125
		v155 = v129
		v156 = v127
		goto L23
	} else {
		goto L34
	}
L30:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	*(*uint8)(unsafe.Add(mBase, uint32(v113))) = uint8(v114)
	if v114 == int32(0) {
		goto L22
	} else {
		goto L32
	}
L31:
	;
	v125 = v121
	v127 = v119
	goto L29
L32:
	;
	v118 = int32(1)
	v119 = v113 + v118
	v121 = v111 + v118
	if v121&int32(3) != 0 {
		v111 = v121
		v113 = v119
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v137 = v125
	v138 = v129
	v139 = v127
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v138
	v141 = int32(4)
	v142 = v139 + v141
	v144 = v137 + v141
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v149 = int32(-2139062144)
	if (int32(16843008)-v146|v146)&v149 == v149 {
		v137 = v144
		v138 = v146
		v139 = v142
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v154 = v144
	v155 = v146
	v156 = v142
	goto L23
L37:
	;
	goto L36
L38:
	;
	v163 = v154
	v165 = v156
	goto L39
L39:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)) = uint8(v166)
	v168 = int32(1)
	if v166 != 0 {
		v163 = v163 + v168
		v165 = v165 + v168
		goto L39
	} else {
		goto L41
	}
L40:
	;
	goto L22
L41:
	;
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+188)) = v88
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v83)+216))
	if (v253^v89)&int32(3) != 0 {
		goto L66
	} else {
		goto L67
	}
L43:
	;
	goto L42
L44:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v232))) = uint8(v231)
	if v231&int32(255) == int32(0) {
		goto L43
	} else {
		goto L59
	}
L45:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v230 = v177
	v231 = v183
	v232 = v88
	goto L44
L46:
	;
	goto L47
L47:
	;
	if v177&int32(3) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v187 = v177
	v189 = v88
	goto L51
L49:
	;
	v201 = v177
	v203 = v88
	goto L50
L50:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v208 = int32(-2139062144)
	if (int32(16843008)-v205|v205)&v208 != v208 {
		v230 = v201
		v231 = v205
		v232 = v203
		goto L44
	} else {
		goto L55
	}
L51:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	*(*uint8)(unsafe.Add(mBase, uint32(v189))) = uint8(v190)
	if v190 == int32(0) {
		goto L43
	} else {
		goto L53
	}
L52:
	;
	v201 = v197
	v203 = v195
	goto L50
L53:
	;
	v194 = int32(1)
	v195 = v189 + v194
	v197 = v187 + v194
	if v197&int32(3) != 0 {
		v187 = v197
		v189 = v195
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v213 = v201
	v214 = v205
	v215 = v203
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v214
	v217 = int32(4)
	v218 = v215 + v217
	v220 = v213 + v217
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v225 = int32(-2139062144)
	if (int32(16843008)-v222|v222)&v225 == v225 {
		v213 = v220
		v214 = v222
		v215 = v218
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v230 = v220
	v231 = v222
	v232 = v218
	goto L44
L58:
	;
	goto L57
L59:
	;
	v239 = v230
	v241 = v232
	goto L60
L60:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v241)+1)) = uint8(v242)
	v244 = int32(1)
	if v242 != 0 {
		v239 = v239 + v244
		v241 = v241 + v244
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L43
L62:
	;
	goto L61
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+216)) = v89
	goto L20
L64:
	;
	goto L63
L65:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v308))) = uint8(v307)
	if v307&int32(255) == int32(0) {
		goto L64
	} else {
		goto L80
	}
L66:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	v306 = v253
	v307 = v259
	v308 = v89
	goto L65
L67:
	;
	goto L68
L68:
	;
	if v253&int32(3) != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v263 = v253
	v265 = v89
	goto L72
L70:
	;
	v277 = v253
	v279 = v89
	goto L71
L71:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v284 = int32(-2139062144)
	if (int32(16843008)-v281|v281)&v284 != v284 {
		v306 = v277
		v307 = v281
		v308 = v279
		goto L65
	} else {
		goto L76
	}
L72:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	*(*uint8)(unsafe.Add(mBase, uint32(v265))) = uint8(v266)
	if v266 == int32(0) {
		goto L64
	} else {
		goto L74
	}
L73:
	;
	v277 = v273
	v279 = v271
	goto L71
L74:
	;
	v270 = int32(1)
	v271 = v265 + v270
	v273 = v263 + v270
	if v273&int32(3) != 0 {
		v263 = v273
		v265 = v271
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v289 = v277
	v290 = v281
	v291 = v279
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = v290
	v293 = int32(4)
	v294 = v291 + v293
	v296 = v289 + v293
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	v301 = int32(-2139062144)
	if (int32(16843008)-v298|v298)&v301 == v301 {
		v289 = v296
		v290 = v298
		v291 = v294
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v306 = v296
	v307 = v298
	v308 = v294
	goto L65
L79:
	;
	goto L78
L80:
	;
	v315 = v306
	v317 = v308
	goto L81
L81:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v317)+1)) = uint8(v318)
	v320 = int32(1)
	if v318 != 0 {
		v315 = v315 + v320
		v317 = v317 + v320
		goto L81
	} else {
		goto L83
	}
L82:
	;
	goto L64
L83:
	;
	goto L82
L84:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[7]))
	if v339 == int32(0) {
		goto L16
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if int32(0) < v344 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	goto L16
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+408)) = v85
	v349 = v82 + int32(412)
	v350 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v349))) = v350
	v353 = v82 + int32(416)
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = v350
	v357 = v82 + int32(420)
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = v350
	v361 = v82 + int32(424)
	*(*uint8)(unsafe.Add(mBase, uint32(v361))) = uint8(v350)
	if v85 < v350 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v412 = v82
	v415 = v87
	v416 = v88
	v417 = v89
	goto L91
L91:
	;
	v424 = v85 + int32(1)
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[3]))
	if v424 < v426+int32(38) {
		v82 = v412
		v83 = v83 + int32(408)
		v85 = v424
		v87 = v415
		v88 = v416
		v89 = v417
		goto L16
	} else {
		goto L100
	}
L92:
	;
	v397 = int32(_a_F_pgstat_read_current_status_3)
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[5])) = v399 + int32(1)
	v403 = int32(-64)
	v408 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[4]))
	v412 = v82 + int32(432)
	v415 = v87 - v403
	v416 = v88 - v403
	v417 = v89 + v408
	goto L91
L93:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[8]))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+16))
	if base.Ui32(v368) <= base.Ui32(v85) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[9]))
	v376 = F_LWLockAcquire(m, v372+int32(512), int32(1))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	v380 = v370 + v85*int32(640)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+44))
	if v381 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v380)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v349))) = v382
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v380)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = v384
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+276)))
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = v386
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+277)))
	*(*uint8)(unsafe.Add(mBase, uint32(v361))) = uint8(v388)
	goto L98
L97:
	;
	goto L98
L98:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_read_current_status[9]))
	F_LWLockRelease(m, v391+int32(512))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L7
	} else {
		goto L99
	}
L99:
	;
	goto L92
L100:
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
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_replslot_from_serialized_name_cb[0]))
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
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_report_checksum_failures_in_db[0])))
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
						F_errmsg_internal(m, int32(_a_F_pgstat_report_checksum_failures_in_db_0), v7)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_pgstat_report_checksum_failures_in_db_1), int32(194), int32(_a_F_pgstat_report_checksum_failures_in_db_2))
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
					F_gettimeofday(m, v50)
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
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_report_recovery_conflict[0])))
	if v4 != int32(1) {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_recovery_conflict[1]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_set_wait_event_storage[0])) = l0
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
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v2 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_slru_flush_cb[0])))
	if v7 == v2 {
		v82 = v2
		return v82
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_slru_flush_cb[1]))
		v13 = v11 + int32(_a_F_pgstat_slru_flush_cb_0)
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
					v38 = v11 + int32(_a_F_pgstat_slru_flush_cb_1) + v37
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_pgstat_slru_flush_cb[2])))
					*(*int64)(unsafe.Add(mBase, uint32(v38))) = v39 + v40
					v43 = *(*int64)(unsafe.Add(mBase, uint32(v38)+8))
					v44 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_pgstat_slru_flush_cb[3])))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v43 + v44
					v47 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
					v48 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_pgstat_slru_flush_cb[4])))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v47 + v48
					v51 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
					v52 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_pgstat_slru_flush_cb[5])))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v51 + v52
					v55 = *(*int64)(unsafe.Add(mBase, uint32(v38)+32))
					v56 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_pgstat_slru_flush_cb[6])))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+32)) = v55 + v56
					v59 = *(*int64)(unsafe.Add(mBase, uint32(v38)+40))
					v60 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_pgstat_slru_flush_cb[7])))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+40)) = v59 + v60
					v63 = *(*int64)(unsafe.Add(mBase, uint32(v38)+48))
					v64 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_pgstat_slru_flush_cb[8])))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+48)) = v63 + v64
					v68 = v33 + int32(1)
					if v68 != int32(8) {
						v33 = v68
						continue
					} else {
						break
					}
					break
				}
				v71 = int32(0)
				base.MemoryFill(m, int32(_a_F_pgstat_slru_flush_cb_2), v71, int32(512))
				F_LWLockRelease(m, v13)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					v79 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_slru_flush_cb[0])) = uint8(v79)
					v82 = v71
					return v82
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
					v82 = int32(1)
					return v82
				} else {
					v33 = int32(0)
					for {
						v37 = v33 << (uint(int32(6)) % 32)
						v38 = v11 + int32(_a_F_pgstat_slru_flush_cb_1) + v37
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_pgstat_slru_flush_cb[2])))
						*(*int64)(unsafe.Add(mBase, uint32(v38))) = v39 + v40
						v43 = *(*int64)(unsafe.Add(mBase, uint32(v38)+8))
						v44 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_pgstat_slru_flush_cb[3])))
						*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v43 + v44
						v47 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
						v48 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_pgstat_slru_flush_cb[4])))
						*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v47 + v48
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
						v52 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_pgstat_slru_flush_cb[5])))
						*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v51 + v52
						v55 = *(*int64)(unsafe.Add(mBase, uint32(v38)+32))
						v56 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_pgstat_slru_flush_cb[6])))
						*(*int64)(unsafe.Add(mBase, uint32(v38)+32)) = v55 + v56
						v59 = *(*int64)(unsafe.Add(mBase, uint32(v38)+40))
						v60 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_pgstat_slru_flush_cb[7])))
						*(*int64)(unsafe.Add(mBase, uint32(v38)+40)) = v59 + v60
						v63 = *(*int64)(unsafe.Add(mBase, uint32(v38)+48))
						v64 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_pgstat_slru_flush_cb[8])))
						*(*int64)(unsafe.Add(mBase, uint32(v38)+48)) = v63 + v64
						v68 = v33 + int32(1)
						if v68 != int32(8) {
							v33 = v68
							continue
						} else {
							break
						}
						break
					}
					v71 = int32(0)
					base.MemoryFill(m, int32(_a_F_pgstat_slru_flush_cb_2), v71, int32(512))
					F_LWLockRelease(m, v13)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						v79 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_slru_flush_cb[0])) = uint8(v79)
						v82 = v71
						return v82
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
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_slru_snapshot_cb[0]))
	v6 = v4 + int32(_a_F_pgstat_slru_snapshot_cb_0)
	v8 = F_LWLockAcquire(m, v6, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		base.MemoryCopy(m, int32(_a_F_pgstat_slru_snapshot_cb_1), v4+int32(_a_F_pgstat_slru_snapshot_cb_2), int32(512))
		F_LWLockRelease(m, v6)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			return
		}
	}
}
func F_pgstat_tracks_backend_bktype(m *base.Module, l0 int32) int32 {
	return int32(base.Ui32(int32(_a_F_pgstat_tracks_backend_bktype_0))>>(uint(l0)%32)) & base.B2i32(base.Ui32(l0) < base.Ui32(int32(17)))
}
