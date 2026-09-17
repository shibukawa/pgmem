package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgstat_archiver_reset_all_cb(m *base.Module, l0 int64) {
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
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_archiver_reset_all_cb[0]))
	v10 = v8 + int32(24)
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
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_archiver_reset_all_cb[1]))
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+176)) = l0
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
	base.MemoryCopy(m, v8+int32(184), v8+int32(48), int32(136))
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
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
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
func F_pgstat_bestart_final(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int64
	_ = v76
	var v80 int64
	_ = v80
	var v84 int64
	_ = v84
	var v88 int64
	_ = v88
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[0]))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[1]))
	if base.Ui32(int32(6)) < base.Ui32(v9) {
		v21 = int32(0)
	} else {
		v12 = int32(0)
		if int32(1)<<(uint(v9)%32)&int32(98) == v12 {
			v21 = v12
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[2]))
			v21 = v20
		}
	}
	v22 = int32(_a_F_pgstat_bestart_final_0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[3]))
	v25 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[3])) = v24 + v25
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v28 + v25
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v6)+208)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v28 + int32(2)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[3])) = v44 - v25
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[1]))
	if int32(base.Ui32(int32(_a_F_pgstat_bestart_final_1))>>(uint(v49)%32))&base.B2i32(base.Ui32(v49) < base.Ui32(int32(17))) != 0 {
		v56 = int32(0)
		v58 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[5])))
		v60 = F_pgstat_get_entry_ref_locked(m, int32(6), v56, v58, v56)
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
			base.MemoryFill(m, v62+int32(24), int32(0), int32(2920))
			F_pgstat_unlock_entry(m, v60)
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return
			} else {
				v71 = int32(0)
				base.MemoryFill(m, int32(_a_F_pgstat_bestart_final_2), v71, int32(2880))
				v76 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[6]))
				*(*int64)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[7])) = v76
				v80 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[8]))
				*(*int64)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[9])) = v80
				v84 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[10]))
				*(*int64)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[11])) = v84
				v88 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[12]))
				*(*int64)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[13])) = v88
				*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[14])) = uint8(v71)
				v95 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[15]))
				if v95 == int32(0) {
					return
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[0]))
					if v99 == int32(0) {
						return
					} else {
						v102 = F_strlen(m, v95)
						mBase = m.M
						v104 = F_pg_mbcliplen(m, v95, v102, int32(63))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return
						} else {
							v106 = int32(_a_F_pgstat_bestart_final_0)
							v108 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[3]))
							v109 = int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[3])) = v108 + v109
							v112 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
							*(*int32)(unsafe.Add(mBase, uint32(v99))) = v112 + v109
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v99)+212))
							if v104 != 0 {
								base.MemoryCopy(m, v116, v95, v104)
							} else {
							}
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v99)+212))
							v120 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v118+v104))) = uint8(v120)
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
							v123 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v99))) = v122 + v123
							v126 = int32(_a_F_pgstat_bestart_final_0)
							v128 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[3]))
							*(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[3])) = v128 - v123
							return
						}
					}
				}
			}
		}
	} else {
		v95 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[15]))
		if v95 == int32(0) {
			return
		} else {
			v99 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[0]))
			if v99 == int32(0) {
				return
			} else {
				v102 = F_strlen(m, v95)
				mBase = m.M
				v104 = F_pg_mbcliplen(m, v95, v102, int32(63))
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return
				} else {
					v106 = int32(_a_F_pgstat_bestart_final_0)
					v108 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[3]))
					v109 = int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[3])) = v108 + v109
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
					*(*int32)(unsafe.Add(mBase, uint32(v99))) = v112 + v109
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v99)+212))
					if v104 != 0 {
						base.MemoryCopy(m, v116, v95, v104)
					} else {
					}
					v118 = *(*int32)(unsafe.Add(mBase, uint32(v99)+212))
					v120 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v118+v104))) = uint8(v120)
					v122 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
					v123 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v99))) = v122 + v123
					v126 = int32(_a_F_pgstat_bestart_final_0)
					v128 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[3]))
					*(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_final[3])) = v128 - v123
					return
				}
			}
		}
	}
}
func F_pgstat_checkpointer_snapshot_cb(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v105 int32
	_ = v105
	var v107 int64
	_ = v107
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[0]))
	goto L1
L1:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+424))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[1]))
	if v33 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v44 = v15 + int32(408)
	v46 = F_LWLockAcquire(m, v44, int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L10
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	base.MemoryCopy(m, int32(_a_F_pgstat_checkpointer_snapshot_cb_0), v15+int32(432), int32(88))
	if v31&int32(1) != 0 {
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
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)+424))
	if v31 != v41 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L2
L10:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v15)+592))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v15)+584))
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v15)+576))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v15)+568))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v15)+560))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v15)+552))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v15)+544))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v15)+536))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v15)+528))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v15)+520))
	F_LWLockRelease(m, v44)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v60 = int32(_a_F_pgstat_checkpointer_snapshot_cb_0)
	v62 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[2]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[2])) = v62 - v57
	v65 = int32(_a_F_pgstat_checkpointer_snapshot_cb_1)
	v67 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[3]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[3])) = v67 - v56
	v70 = int32(_a_F_pgstat_checkpointer_snapshot_cb_2)
	v72 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[4])) = v72 - v55
	v75 = int32(_a_F_pgstat_checkpointer_snapshot_cb_3)
	v77 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[5]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[5])) = v77 - v54
	v80 = int32(_a_F_pgstat_checkpointer_snapshot_cb_4)
	v82 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[6]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[6])) = v82 - v53
	v85 = int32(_a_F_pgstat_checkpointer_snapshot_cb_5)
	v87 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[7]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[7])) = v87 - v52
	v90 = int32(_a_F_pgstat_checkpointer_snapshot_cb_6)
	v92 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[8]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[8])) = v92 - v51
	v95 = int32(_a_F_pgstat_checkpointer_snapshot_cb_7)
	v97 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[9]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[9])) = v97 - v50
	v100 = int32(_a_F_pgstat_checkpointer_snapshot_cb_8)
	v102 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[10]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[10])) = v102 - v49
	v105 = int32(_a_F_pgstat_checkpointer_snapshot_cb_9)
	v107 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[11]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_checkpointer_snapshot_cb[11])) = v107 - v48
	return
}
func F_pgstat_count_heap_delete(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v5 == int32(0) {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
		if v8 != int32(1) {
			return
		} else {
			F_pgstat_assoc_relation(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
				v14 = v13
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_heap_delete[0]))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				if v18 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
					if v19 == v17 {
						v36 = v18
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39 + int64(1)
						return
					} else {
						v21 = F_pgstat_get_xact_stack_level(m, v17)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_heap_delete[1]))
							v26 = F_MemoryContextAllocZero(m, v24, int32(72))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v17
								v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v14
								*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v29
								v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v32
								*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v26
								*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v26
								v36 = v26
								v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39 + int64(1)
								return
							}
						}
					}
				} else {
					v21 = F_pgstat_get_xact_stack_level(m, v17)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_heap_delete[1]))
						v26 = F_MemoryContextAllocZero(m, v24, int32(72))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v17
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v14
							*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v29
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v32
							*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v26
							*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v26
							v36 = v26
							v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39 + int64(1)
							return
						}
					}
				}
			}
		}
	} else {
		v14 = v5
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_heap_delete[0]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		if v18 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
			if v19 == v17 {
				v36 = v18
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39 + int64(1)
				return
			} else {
				v21 = F_pgstat_get_xact_stack_level(m, v17)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_heap_delete[1]))
					v26 = F_MemoryContextAllocZero(m, v24, int32(72))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v17
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v14
						*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v29
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v26
						*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v26
						v36 = v26
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39 + int64(1)
						return
					}
				}
			}
		} else {
			v21 = F_pgstat_get_xact_stack_level(m, v17)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_heap_delete[1]))
				v26 = F_MemoryContextAllocZero(m, v24, int32(72))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v17
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v29
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v32
					*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v26
					v36 = v26
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39 + int64(1)
					return
				}
			}
		}
	}
}
func F_pgstat_count_slru_page_hit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	v3 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_count_slru_page_hit[0])) = uint8(v3)
	*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_count_slru_page_hit[1])) = uint8(v3)
	v9 = l0 << (uint(int32(6)) % 32)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_pgstat_count_slru_page_hit[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_pgstat_count_slru_page_hit[2]))) = v12 + int64(1)
	return
}
func F_pgstat_create_transactional(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_pgstat_get_entry_ref(m, l0, l1, l2, v4, v4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 != 0 {
			v16 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				if v16 != 0 {
					if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
						v46 = l0*int32(72) + int32(_a_F_pgstat_create_transactional_0)
					} else {
						if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
							v44 = int32(0)
						} else {
							v32 = int32(0)
							v34 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_create_transactional[0]))
							if v34 == v32 {
								v44 = v32
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v34+l0<<(uint(int32(2))%32)-int32(96))))
								v44 = v42
							}
						}
						v46 = v44
					}
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v47
					F_errmsg(m, int32(_a_F_pgstat_create_transactional_1), v8)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_pgstat_create_transactional_2), int32(368), int32(_a_F_pgstat_create_transactional_3))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_pgstat_reset(m, l0, l1, l2)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								F_create_drop_transactional_internal(m, l0, l1, l2, int32(1))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					}
				} else {
					F_pgstat_reset(m, l0, l1, l2)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						F_create_drop_transactional_internal(m, l0, l1, l2, int32(1))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		} else {
			F_create_drop_transactional_internal(m, l0, l1, l2, int32(1))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_pgstat_drop_all_entries(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v65 int64
	_ = v65
	var v71 int64
	_ = v71
	var v77 int64
	_ = v77
	var v82 int64
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int64
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	v7 = int64(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v15 = v10 + int32(-44)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_all_entries[0]))
	v18 = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v15)+12)) = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+24)) = uint8(v18)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(-1)
	goto L1
L1:
	;
	v27 = F_dshash_seq_next(m, v15)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	m.G0 = v12 - int32(-64)
	return
L3:
	;
	return
L4:
	;
	if v27 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_dshash_seq_term(m, v15)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v35 = v27
	v41 = v7
	goto L9
L8:
	;
	goto L2
L9:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+16)))
	if v42 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_dshash_seq_term(m, v10+int32(-44))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L30
	}
L11:
	;
	goto L10
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_all_entries[1]))
	if v46 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v151 = F_dshash_seq_next(m, v10+int32(-44))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L3
	} else {
		goto L28
	}
L15:
	;
	v140 = v10 + int32(-44)
	v141 = F_pgstat_drop_entry_internal(m, v35, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L3
	} else {
		goto L25
	}
L16:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v35)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v49
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v55 = int64(23)
	v58 = int64(2388976653695081527)
	v59 = (v51 ^ int64(base.Ui64(v51)>>(uint(v55)%64))) * v58
	v60 = int64(47)
	v65 = int64(-8645972361240307355)
	v71 = (v49 ^ int64(base.Ui64(v49)>>(uint(v55)%64))) * v58
	v77 = ((v59^int64(base.Ui64(v59)>>(uint(v60)%64))^int64(-9208349263878056368))*v65 ^ int64(base.Ui64(v71)>>(uint(v60)%64)) ^ v71) * v65
	v82 = (int64(base.Ui64(v77)>>(uint(v55)%64)) ^ v77) * v58
	v90 = v54 & base.I32_wrap_i64(int64(base.Ui64(v82)>>(uint(v60)%64))^v82-int64(base.Ui64(v82)>>(uint(int64(32))%64)))
	v93 = v53 + v90*int32(24)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+16)))
	if v94 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v97 = v93
	v100 = v90
	goto L18
L18:
	;
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v97)+8))
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v12)+56))
	if v106^v107|(v109^v110) != int64(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v97)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v123
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v125
	F_pgstat_release_entry_ref(m, v12, v122, int32(1))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L24
	}
L20:
	;
	v117 = (v100 + int32(1)) & v54
	v120 = v53 + v117*int32(24)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+16)))
	if v121 != 0 {
		v97 = v120
		v100 = v117
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	goto L19
L23:
	;
	goto L15
L24:
	;
	goto L15
L25:
	;
	v146 = v41 + base.I64_extend_i32_u(v141^int32(1))
	v147 = F_dshash_seq_next(m, v140)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	if v147 != 0 {
		v35 = v147
		v41 = v146
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v161 = v146
	goto L11
L28:
	;
	if v151 != 0 {
		v35 = v151
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v161 = v41
	goto L11
L30:
	;
	if v161 == int64(0) {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_all_entries[2]))
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v169)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v169)+16)) = v170 + int64(1)
	goto L2
}
func F_pgstat_drop_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v38 int64
	_ = v38
	var v44 int64
	_ = v44
	var v50 int64
	_ = v50
	var v55 int64
	_ = v55
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v110 int32
	_ = v110
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int64
	_ = v146
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int64
	_ = v236
	var v238 int64
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int64
	_ = v276
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int64
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int64
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int64
	_ = v317
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l0
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_entry[0]))
	if v22 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v124 = int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_entry[1]))
	v130 = F_dshash_find(m, v126, v16+int32(32), v124)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L10
	} else {
		goto L13
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
	v28 = int64(23)
	v31 = int64(2388976653695081527)
	v32 = (int64(base.Ui64(v27)>>(uint(v28)%64)) ^ v27) * v31
	v33 = int64(47)
	v38 = int64(-8645972361240307355)
	v44 = (int64(base.Ui64(l2)>>(uint(v28)%64)) ^ l2) * v31
	v50 = ((v32^int64(base.Ui64(v32)>>(uint(v33)%64))^int64(-9208349263878056368))*v38 ^ int64(base.Ui64(v44)>>(uint(v33)%64)) ^ v44) * v38
	v55 = (int64(base.Ui64(v50)>>(uint(v28)%64)) ^ v50) * v31
	v63 = v26 & base.I32_wrap_i64(int64(base.Ui64(v55)>>(uint(v33)%64))^v55-int64(base.Ui64(v55)>>(uint(int64(32))%64)))
	v66 = v25 + v63*int32(24)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+16)))
	if v67 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v70 = v63
	v71 = v66
	goto L4
L4:
	;
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v71)+8))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
	if v83^v84|(v86^v87) != int64(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v71)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v102
	F_pgstat_release_entry_ref(m, v16+int32(16), v99, int32(1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v94 = (v70 + int32(1)) & v26
	v97 = v25 + v94*int32(24)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+16)))
	if v98 != 0 {
		v70 = v94
		v71 = v97
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	goto L1
L10:
	;
	return int32(0)
L11:
	;
	goto L1
L12:
	;
	m.G0 = v16 + int32(80)
	return v332
L13:
	;
	if v130 == int32(0) {
		v332 = v124
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v135 = F_pgstat_drop_entry_internal(m, v130, int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	if v137 != int32(1) {
		v332 = v135
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_entry[0]))
	if v142 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v259 = v16 + int32(52)
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_entry[1]))
	v262 = int32(1)
	v263 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v259)+4)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v261
	*(*int64)(unsafe.Add(mBase, uint32(v259)+12)) = v263
	*(*uint8)(unsafe.Add(mBase, uint32(v259)+24)) = uint8(v262)
	*(*int32)(unsafe.Add(mBase, uint32(v259)+20)) = int32(-1)
	goto L33
L18:
	;
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v142)))
	if v146 == int64(0) {
		v179 = int32(-1)
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v189 = v142
	v190 = v179
	v195 = int32(0)
	goto L25
L20:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v142)+20))
	v155 = int32(0)
	goto L21
L21:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149+v155*int32(24))+16)))
	if v167 != int32(1) {
		v179 = v155
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v179 = int32(-1)
	goto L19
L23:
	;
	v171 = v155 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v171)) < base.Ui64(v146) {
		v155 = v171
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v203 = v190
	v208 = v195
	v211 = v195
	goto L27
L27:
	;
	if v211&int32(1) != 0 {
		goto L17
	} else {
		goto L29
	}
L28:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v233 != v140 {
		v190 = v227
		v195 = v224
		goto L25
	} else {
		goto L31
	}
L29:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	v218 = int32(1)
	v219 = v203 - v218
	v223 = base.B2i32(v217&(v219^v179) == int32(0))
	v224 = v223 | v208
	v227 = v217 & v219
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v189)+20))
	v229 = v203*int32(24) + v228
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+16)))
	if v230 != v218 {
		v203 = v227
		v208 = v224
		v211 = v223
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v229)+20))
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v229)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v236
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v229)))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v238
	F_pgstat_release_entry_ref(m, v16, v235, int32(1))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_entry[0]))
	v189 = v244
	v190 = v227
	v195 = v224
	goto L25
L33:
	;
	v271 = F_dshash_seq_next(m, v259)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	if v271 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v275 = v271
	v276 = int64(0)
	goto L38
L36:
	;
	goto L37
L37:
	;
	F_dshash_seq_term(m, v16+int32(52))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L10
	} else {
		goto L51
	}
L38:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+16)))
	if v287 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	F_dshash_seq_term(m, v16+int32(52))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L10
	} else {
		goto L49
	}
L40:
	;
	goto L39
L41:
	;
	v304 = F_dshash_seq_next(m, v16+int32(52))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L10
	} else {
		goto L47
	}
L42:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v290 != v140 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v293 = v16 + int32(52)
	v294 = F_pgstat_drop_entry_internal(m, v275, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	v299 = v276 + base.I64_extend_i32_u(v294^int32(1))
	v300 = F_dshash_seq_next(m, v293)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	if v300 != 0 {
		v275 = v300
		v276 = v299
		goto L38
	} else {
		goto L46
	}
L46:
	;
	v308 = v299
	goto L40
L47:
	;
	if v304 != 0 {
		v275 = v304
		goto L38
	} else {
		goto L48
	}
L48:
	;
	v308 = v276
	goto L40
L49:
	;
	if v308 == int64(0) {
		v332 = v135
		goto L12
	} else {
		goto L50
	}
L50:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_entry[2]))
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v316)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v316)+16)) = v317 + int64(1)
	v332 = v135
	goto L12
L51:
	;
	v332 = v135
	goto L12
}
func F_pgstat_end_function_usage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v36 int64
	_ = v36
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 != 0 {
		F___clock_gettime(m, int32(1), v11)
		mBase = m.M
		v16 = int32(_a_F_pgstat_end_function_usage_0)
		v17 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_end_function_usage[0]))
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		v20 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+8)))
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v26 = v20 + v21*int64(1000000000) - v25
		*(*int64)(unsafe.Add(mBase, _c_F_pgstat_end_function_usage[0])) = v19 + v26
		v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		if l1 != 0 {
			v31 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
			*(*int64)(unsafe.Add(mBase, uint32(v13))) = v31 + int64(1)
		} else {
		}
		*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v29 + v26
		v36 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v36 + (v26 - v17 + v19)
	} else {
	}
	m.G0 = v11 + int32(16)
	return
}
func F_pgstat_execute_transactional_drops(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	v3 = int32(0)
	if l0 <= v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = v3
	v12 = v3
	goto L3
L3:
	;
	v15 = l1 + v10<<(uint(int32(4))%32)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	v19 = F_pgstat_drop_entry(m, v16, v17, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v23 <= int32(0) {
		goto L1
	} else {
		goto L8
	}
L5:
	;
	return
L6:
	;
	v21 = int32(1)
	v23 = v12 + (v19 ^ v21)
	v25 = v10 + v21
	if v25 != l0 {
		v10 = v25
		v12 = v23
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_execute_transactional_drops[0]))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v31)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+16)) = v32 + int64(1)
	goto L9
L9:
	;
	goto L1
}
func F_pgstat_fetch_stat_bgwriter(m *base.Module) int32 {
	var v5 int32
	_ = v5
	F_pgstat_snapshot_fixed(m, int32(8))
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(_a_F_pgstat_fetch_stat_bgwriter_0)
	}
}
func F_pgstat_fetch_stat_dbentry(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_pgstat_fetch_entry(m, int32(1), l0, int64(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_pgstat_fetch_stat_funcentry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_stat_funcentry[0]))
	v6 = F_pgstat_fetch_entry(m, int32(3), v4, base.I64_extend_i32_u(l0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_pgstat_get_slru_index(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	v2 = int32(_a_F_pgstat_get_slru_index_0)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_get_slru_index[0])))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v5 == int32(0))|base.B2i32(v5 != v8) != 0 {
		v26 = v5
		v27 = v8
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v26-v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	goto L1
L3:
	;
	v11 = v2
	v12 = l0
	goto L4
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v16 == int32(0) {
		v26 = v16
		v27 = v15
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v26 = v16
	v27 = v15
	goto L2
L6:
	;
	v19 = int32(1)
	if v16 == v15 {
		v11 = v11 + v19
		v12 = v12 + v19
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	return int32(0)
L9:
	;
	goto L10
L10:
	;
	v33 = int32(_a_F_pgstat_get_slru_index_1)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_get_slru_index[1])))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v36 == int32(0))|base.B2i32(v36 != v39) != 0 {
		v57 = v36
		v58 = v39
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v57-v58 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	v42 = v33
	v43 = l0
	goto L14
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v47 == int32(0) {
		v57 = v47
		v58 = v46
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v57 = v47
	v58 = v46
	goto L12
L16:
	;
	v50 = int32(1)
	if v47 == v46 {
		v42 = v42 + v50
		v43 = v43 + v50
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	return int32(1)
L19:
	;
	goto L20
L20:
	;
	v64 = int32(_a_F_pgstat_get_slru_index_2)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_get_slru_index[2])))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v67 == int32(0))|base.B2i32(v67 != v70) != 0 {
		v88 = v67
		v89 = v70
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v88-v89 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	goto L21
L23:
	;
	v73 = v64
	v74 = l0
	goto L24
L24:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	if v78 == int32(0) {
		v88 = v78
		v89 = v77
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v88 = v78
	v89 = v77
	goto L22
L26:
	;
	v81 = int32(1)
	if v78 == v77 {
		v73 = v73 + v81
		v74 = v74 + v81
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	return int32(2)
L29:
	;
	goto L30
L30:
	;
	v95 = int32(_a_F_pgstat_get_slru_index_3)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_get_slru_index[3])))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v98 == int32(0))|base.B2i32(v98 != v101) != 0 {
		v119 = v98
		v120 = v101
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v119-v120 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L32:
	;
	goto L31
L33:
	;
	v104 = v95
	v105 = l0
	goto L34
L34:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	if v109 == int32(0) {
		v119 = v109
		v120 = v108
		goto L32
	} else {
		goto L36
	}
L35:
	;
	v119 = v109
	v120 = v108
	goto L32
L36:
	;
	v112 = int32(1)
	if v109 == v108 {
		v104 = v104 + v112
		v105 = v105 + v112
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	return int32(3)
L39:
	;
	goto L40
L40:
	;
	v126 = int32(_a_F_pgstat_get_slru_index_4)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_get_slru_index[4])))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v129 == int32(0))|base.B2i32(v129 != v132) != 0 {
		v150 = v129
		v151 = v132
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v150-v151 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v135 = v126
	v136 = l0
	goto L44
L44:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	if v140 == int32(0) {
		v150 = v140
		v151 = v139
		goto L42
	} else {
		goto L46
	}
L45:
	;
	v150 = v140
	v151 = v139
	goto L42
L46:
	;
	v143 = int32(1)
	if v140 == v139 {
		v135 = v135 + v143
		v136 = v136 + v143
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	return int32(4)
L49:
	;
	goto L50
L50:
	;
	v157 = int32(_a_F_pgstat_get_slru_index_5)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_get_slru_index[5])))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v160 == int32(0))|base.B2i32(v160 != v163) != 0 {
		v181 = v160
		v182 = v163
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v181-v182 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	goto L51
L53:
	;
	v166 = v157
	v167 = l0
	goto L54
L54:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
	if v171 == int32(0) {
		v181 = v171
		v182 = v170
		goto L52
	} else {
		goto L56
	}
L55:
	;
	v181 = v171
	v182 = v170
	goto L52
L56:
	;
	v174 = int32(1)
	if v171 == v170 {
		v166 = v166 + v174
		v167 = v167 + v174
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	return int32(5)
L59:
	;
	goto L60
L60:
	;
	v190 = int32(_a_F_pgstat_get_slru_index_6)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_get_slru_index[6])))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v193 == int32(0))|base.B2i32(v193 != v196) != 0 {
		v214 = v193
		v215 = v196
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v214-v215 != 0 {
		goto L68
	} else {
		goto L69
	}
L62:
	;
	goto L61
L63:
	;
	v199 = v190
	v200 = l0
	goto L64
L64:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+1)))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+1)))
	if v204 == int32(0) {
		v214 = v204
		v215 = v203
		goto L62
	} else {
		goto L66
	}
L65:
	;
	v214 = v204
	v215 = v203
	goto L62
L66:
	;
	v207 = int32(1)
	if v204 == v203 {
		v199 = v199 + v207
		v200 = v200 + v207
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v217 = int32(7)
	goto L70
L69:
	;
	v217 = int32(6)
	goto L70
L70:
	;
	return v217
}
func F_pgstat_io_reset_all_cb(m *base.Module, l0 int64) {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_io_reset_all_cb[0]))
	v8 = v6 + int32(608)
	v10 = F_LWLockAcquire(m, v8, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_io_reset_all_cb[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+896)) = l0
	base.MemoryFill(m, v6+int32(904), int32(0), int32(2880))
	F_LWLockRelease(m, v8)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = int32(1)
	goto L4
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_io_reset_all_cb[0]))
	v33 = v28 + v24<<(uint(int32(4))%32) + int32(608)
	v35 = F_LWLockAcquire(m, v33, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v37 = int32(2880)
	base.MemoryFill(m, v28+v24*v37+int32(904), int32(0), v37)
	F_LWLockRelease(m, v33)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v48 = v24 + int32(1)
	if v48 != int32(18) {
		v24 = v48
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_pgstat_progress_start_command(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_start_command[0]))
	if v5 == int32(0) {
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_progress_start_command[1])))
		if v9&int32(1) == int32(0) {
		} else {
			v14 = int32(_a_F_pgstat_progress_start_command_0)
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_start_command[2]))
			v17 = int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_start_command[2])) = v16 + v17
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v20 + v17
			*(*int32)(unsafe.Add(mBase, uint32(v5)+220)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v5)+224)) = l1
			base.MemoryFill(m, v5+int32(232), int32(0), int32(160))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v31 + v17
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_start_command[2]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_start_command[2])) = v37 - v17
		}
	}
	return
}
func F_pgstat_relation_delete_pending_cb(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+128))
	if v4 == int32(0) {
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+272))
		if v7 == int32(0) {
		} else {
			v10 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+128)) = v10
			*(*int32)(unsafe.Add(mBase, uint32(v4)+272)) = v10
		}
	}
	return
}
func F_pgstat_report_checkpointer(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[0]))
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[1]))
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[2]))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[3]))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[4]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[5]))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[6]))
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[7]))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[8]))
	if v5|(v7|(v9|(v11|(v13|(v15|(v17|v19)))))) != 0 {
		v64 = int32(_a_F_pgstat_report_checkpointer_0)
		v66 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
		v67 = int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v66 + v67
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
		*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
		v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
		v75 = int32(_a_F_pgstat_report_checkpointer_1)
		v76 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[8]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
		v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
		v81 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[6]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
		v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
		v86 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[4]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
		v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
		v91 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
		v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
		v96 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[10]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
		v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
		v101 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[11]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
		v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
		v106 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[12]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
		v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
		v111 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[13]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
		v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
		v116 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[14]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
		v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
		v121 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[15]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
		*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
		v130 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
		*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v130 - v67
		v135 = int32(0)
		base.MemoryFill(m, v75, v135, int32(88))
		F_pgstat_flush_io(m, v135)
		mBase = m.M
		v140 = m.ExcPending
		if v140 != 0 {
			return
		} else {
			return
		}
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[16]))
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[13]))
		v32 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[17]))
		v34 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[12]))
		v36 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[18]))
		v38 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[11]))
		v40 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[19]))
		v42 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[10]))
		if v28|(v30|(v32|(v34|(v36|(v38|(v40|v42)))))) != 0 {
			v64 = int32(_a_F_pgstat_report_checkpointer_0)
			v66 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
			v67 = int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v66 + v67
			v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
			*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
			v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
			v75 = int32(_a_F_pgstat_report_checkpointer_1)
			v76 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[8]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
			v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
			v81 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[6]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
			v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
			v86 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[4]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
			v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
			v91 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[2]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
			v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
			v96 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[10]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
			v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
			v101 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[11]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
			v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
			v106 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[12]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
			v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
			v111 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[13]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
			v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
			v116 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[14]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
			v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
			v121 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[15]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
			*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
			v130 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v130 - v67
			v135 = int32(0)
			base.MemoryFill(m, v75, v135, int32(88))
			F_pgstat_flush_io(m, v135)
			mBase = m.M
			v140 = m.ExcPending
			if v140 != 0 {
				return
			} else {
				return
			}
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[14]))
			if v51 != 0 {
				v64 = int32(_a_F_pgstat_report_checkpointer_0)
				v66 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
				v67 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v66 + v67
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
				*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
				v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
				v75 = int32(_a_F_pgstat_report_checkpointer_1)
				v76 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[8]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
				v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
				v81 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[6]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
				v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
				v86 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[4]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
				v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
				v91 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[2]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
				v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
				v96 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[10]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
				v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
				v101 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[11]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
				v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
				v106 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[12]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
				v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
				v111 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[13]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
				v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
				v116 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[14]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
				v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
				v121 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[15]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
				*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
				v130 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v130 - v67
				v135 = int32(0)
				base.MemoryFill(m, v75, v135, int32(88))
				F_pgstat_flush_io(m, v135)
				mBase = m.M
				v140 = m.ExcPending
				if v140 != 0 {
					return
				} else {
					return
				}
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[20]))
				if v53 != 0 {
					v64 = int32(_a_F_pgstat_report_checkpointer_0)
					v66 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
					v67 = int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v66 + v67
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
					*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
					v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
					v75 = int32(_a_F_pgstat_report_checkpointer_1)
					v76 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[8]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
					v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
					v81 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[6]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
					v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
					v86 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[4]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
					v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
					v91 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[2]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
					v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
					v96 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[10]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
					v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
					v101 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[11]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
					v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
					v106 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[12]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
					v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
					v111 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[13]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
					v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
					v116 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[14]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
					v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
					v121 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[15]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
					*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
					v130 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
					*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v130 - v67
					v135 = int32(0)
					base.MemoryFill(m, v75, v135, int32(88))
					F_pgstat_flush_io(m, v135)
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return
					} else {
						return
					}
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[15]))
					if v55 != 0 {
						v64 = int32(_a_F_pgstat_report_checkpointer_0)
						v66 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
						v67 = int32(1)
						*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v66 + v67
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
						*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
						v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
						v75 = int32(_a_F_pgstat_report_checkpointer_1)
						v76 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[8]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
						v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
						v81 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[6]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
						v86 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[4]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
						v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
						v91 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[2]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
						v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
						v96 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[10]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
						v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
						v101 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[11]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
						v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
						v106 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[12]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
						v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
						v111 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[13]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
						v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
						v116 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[14]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
						v121 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[15]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
						*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
						v130 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
						*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v130 - v67
						v135 = int32(0)
						base.MemoryFill(m, v75, v135, int32(88))
						F_pgstat_flush_io(m, v135)
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return
						} else {
							return
						}
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[21]))
						if v57 != 0 {
							v64 = int32(_a_F_pgstat_report_checkpointer_0)
							v66 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
							v67 = int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v66 + v67
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
							*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
							v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
							v75 = int32(_a_F_pgstat_report_checkpointer_1)
							v76 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[8]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
							v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
							v81 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[6]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
							v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
							v86 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[4]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
							v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
							v91 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[2]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
							v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
							v96 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[10]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
							v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
							v101 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[11]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
							v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
							v106 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[12]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
							v111 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[13]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
							v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
							v116 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[14]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
							v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
							v121 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[15]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
							*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
							v130 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
							*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v130 - v67
							v135 = int32(0)
							base.MemoryFill(m, v75, v135, int32(88))
							F_pgstat_flush_io(m, v135)
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return
							} else {
								return
							}
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[22]))
							if v59 != 0 {
								v64 = int32(_a_F_pgstat_report_checkpointer_0)
								v66 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
								v67 = int32(1)
								*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v66 + v67
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
								*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
								v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
								v75 = int32(_a_F_pgstat_report_checkpointer_1)
								v76 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[8]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
								v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
								v81 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[6]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
								v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
								v86 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[4]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
								v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
								v91 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[2]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
								v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
								v96 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[10]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
								v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
								v101 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[11]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
								v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
								v106 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[12]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
								v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
								v111 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[13]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
								v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
								v116 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[14]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
								v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
								v121 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[15]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
								*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
								v130 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
								*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v130 - v67
								v135 = int32(0)
								base.MemoryFill(m, v75, v135, int32(88))
								F_pgstat_flush_io(m, v135)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return
								} else {
									return
								}
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[23]))
								if v61 == int32(0) {
									return
								} else {
									v64 = int32(_a_F_pgstat_report_checkpointer_0)
									v66 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
									v67 = int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v66 + v67
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
									*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
									v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
									v75 = int32(_a_F_pgstat_report_checkpointer_1)
									v76 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[8]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
									v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
									v81 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[6]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
									v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
									v86 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[4]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
									v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
									v91 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
									v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
									v96 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[10]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
									v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
									v101 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[11]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
									v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
									v106 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[12]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
									v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
									v111 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[13]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
									v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
									v116 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[14]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
									v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
									v121 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[15]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
									*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
									v130 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9]))
									*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_checkpointer[9])) = v130 - v67
									v135 = int32(0)
									base.MemoryFill(m, v75, v135, int32(88))
									F_pgstat_flush_io(m, v135)
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
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
		}
	}
}
func F_pgstat_reset_matching_entries(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v171 int32
	_ = v171
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = v12 + int32(4)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_reset_matching_entries[0]))
	v19 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v15)+12)) = v19
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+24)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(-1)
	goto L1
L1:
	;
	v27 = F_dshash_seq_next(m, v15)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = v27
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_dshash_seq_term(m, v12+int32(4))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L46
	}
L7:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+16)))
	if v38 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v157 = F_dshash_seq_next(m, v12+int32(4))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L44
	}
L10:
	;
	v39 = m.T0[l0].(func(*base.Module, int32, int32) int32)(m, v32, l1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	if v39 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_reset_matching_entries[1]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v46 = F_dsa_get_address(m, v44, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v49 = v46 + int32(4)
	v51 = F_LWLockAcquire(m, v49, int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if base.Ui32(v53-int32(1)) <= base.Ui32(int32(11)) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if base.Ui32(v53-int32(1)) <= base.Ui32(int32(11)) {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v82 = v53*int32(72) + int32(_a_F_pgstat_reset_matching_entries_0)
	goto L15
L17:
	;
	goto L18
L18:
	;
	if base.Ui32(int32(8)) < base.Ui32(v53-int32(24)) {
		v80 = int32(0)
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v82 = v80
	goto L15
L20:
	;
	v68 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_reset_matching_entries[2]))
	if v70 == v68 {
		v80 = v68
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v70+v53<<(uint(int32(2))%32)-int32(96))))
	v80 = v78
	goto L19
L22:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	if base.Ui32(v53-int32(1)) <= base.Ui32(int32(11)) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v111 = v53*int32(72) + int32(_a_F_pgstat_reset_matching_entries_0)
	goto L22
L24:
	;
	goto L25
L25:
	;
	if base.Ui32(int32(8)) < base.Ui32(v53-int32(24)) {
		v109 = int32(0)
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v111 = v109
	goto L22
L27:
	;
	v97 = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_reset_matching_entries[2]))
	if v99 == v97 {
		v109 = v97
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v99+v53<<(uint(int32(2))%32)-int32(96))))
	v109 = v107
	goto L26
L29:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+20))
	if v142 != 0 {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v141 = v53*int32(72) + int32(_a_F_pgstat_reset_matching_entries_0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(int32(8)) < base.Ui32(v53-int32(24)) {
		v139 = int32(0)
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v141 = v139
	goto L29
L34:
	;
	v127 = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_reset_matching_entries[2]))
	if v129 == v127 {
		v139 = v127
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v129+v53<<(uint(int32(2))%32)-int32(96))))
	v139 = v137
	goto L33
L36:
	;
	base.MemoryFill(m, v46+v112, int32(0), v142)
	goto L38
L37:
	;
	goto L38
L38:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	if v146 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	m.T0[v146].(func(*base.Module, int32, int64))(m, v46, l2)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_LWLockRelease(m, v49)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	goto L9
L44:
	;
	if v157 != 0 {
		v32 = v157
		goto L7
	} else {
		goto L45
	}
L45:
	;
	goto L8
L46:
	;
	m.G0 = v12 + int32(32)
	return
}
func F_pgstat_reset_of_kind(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
		v29 = l0*int32(72) + int32(_a_F_pgstat_reset_of_kind_0)
	} else {
		if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
			v29 = int32(0)
		} else {
			v17 = int32(0)
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_reset_of_kind[0]))
			if v19 == v17 {
				v29 = v17
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v19+l0<<(uint(int32(2))%32)-int32(96))))
				v29 = v27
			}
		}
	}
	v33 = m.G0
	v34 = int32(16)
	v35 = v33 - v34
	m.G0 = v35
	F_gettimeofday(m, v35)
	mBase = m.M
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	v39 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35)+8)))
	m.G0 = v35 + v34
	v47 = v39 + v38*int64(1000000) - int64(946684800000000)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v48&int32(1) != 0 {
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
		m.T0[v51].(func(*base.Module, int64))(m, v47)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return
		} else {
			return
		}
	} else {
		F_pgstat_reset_matching_entries(m, int32(1233), l0, v47)
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			return
		}
	}
}
func F_pgstat_snapshot_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v36 int64
	_ = v36
	var v42 int64
	_ = v42
	var v48 int64
	_ = v48
	var v53 int64
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v81 int64
	_ = v81
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v102 int64
	_ = v102
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v129 int64
	_ = v129
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v189 int64
	_ = v189
	var v195 int64
	_ = v195
	var v200 int64
	_ = v200
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v242 int64
	_ = v242
	var v243 int64
	_ = v243
	var v244 int64
	_ = v244
	var v249 int64
	_ = v249
	var v251 int64
	_ = v251
	var v256 int64
	_ = v256
	var v262 int64
	_ = v262
	var v267 int64
	_ = v267
	var v275 int32
	_ = v275
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int64
	_ = v298
	var v300 int64
	_ = v300
	var v302 int64
	_ = v302
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int64
	_ = v355
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v372 int64
	_ = v372
	var v374 int64
	_ = v374
	var v375 int64
	_ = v375
	var v381 int32
	_ = v381
	var v382 int64
	_ = v382
	var v383 int64
	_ = v383
	var v386 int64
	_ = v386
	var v387 int64
	_ = v387
	var v388 int64
	_ = v388
	var v393 int64
	_ = v393
	var v395 int64
	_ = v395
	var v400 int64
	_ = v400
	var v406 int64
	_ = v406
	var v411 int64
	_ = v411
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int64
	_ = v453
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int64
	_ = v503
	var v505 int64
	_ = v505
	var v507 int64
	_ = v507
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int64
	_ = v531
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v561 int32
	_ = v561
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int64
	_ = v575
	var v577 int64
	_ = v577
	var v587 int32
	_ = v587
	var v595 int32
	_ = v595
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v630 int32
	_ = v630
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v20
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v24
	v26 = int64(23)
	v29 = int64(2388976653695081527)
	v30 = (v21 ^ int64(base.Ui64(v21)>>(uint(v26)%64))) * v29
	v31 = int64(47)
	v36 = int64(-8645972361240307355)
	v42 = (v20 ^ int64(base.Ui64(v20)>>(uint(v26)%64))) * v29
	v48 = ((v30^int64(base.Ui64(v30)>>(uint(v31)%64))^int64(-9208349263878056368))*v36 ^ int64(base.Ui64(v42)>>(uint(v31)%64)) ^ v42) * v36
	v53 = (int64(base.Ui64(v48)>>(uint(v26)%64)) ^ v48) * v29
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v69 = base.B2i32(base.Ui32(v61) < base.Ui32(v62))
	goto L1
L1:
	;
	if v69 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L20
	} else {
		goto L89
	}
L3:
	;
	goto L2
L4:
	;
	v630 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v630
	v69 = v630
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L20
	} else {
		goto L86
	}
L6:
	;
	v81 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v81 == int64(4294967296) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v349 = v348 & base.I32_wrap_i64(int64(base.Ui64(v53)>>(uint(v31)%64))^v53-int64(base.Ui64(v53)>>(uint(int64(32))%64)))
	v352 = v347 + v349*int32(24)
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+16)))
	if v353 != 0 {
		goto L52
	} else {
		goto L53
	}
L9:
	;
	v84 = int32(0)
	v86 = int64(2)
	v88 = v81 << (uint(int64(1)) % 64)
	if base.Ui64(v88) <= base.Ui64(v86) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v69 = int32(1)
	goto L1
L11:
	;
	v91 = v86
	goto L13
L12:
	;
	v91 = v88
	goto L13
L13:
	;
	v92 = int64(1)
	if v91&(v91-v92) == int64(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v102 = v91
	goto L16
L15:
	;
	v102 = v92 << (uint(int64(64)-base.I64_clz(v91)) % 64)
	goto L16
L16:
	;
	if base.Ui64(v102*int64(24)) < base.Ui64(int64(2147483647)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v108 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v114 = F_MemoryContextAllocExtended(m, v109, base.I32_wrap_i64(v102)*int32(24), int32(5))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	goto L3
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v114
	v119 = int64(1)
	if v102&(v102-v119) == int64(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v129 = v102
	goto L24
L23:
	;
	v129 = v119 << (uint(int64(64)-base.I64_clz(v102)) % 64)
	goto L24
L24:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v129*int64(24)) {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v129
	v137 = base.I32_wrap_i64(v129) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v137
	if v129 == int64(4294967296) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v146 = int32(-85899346)
	goto L28
L27:
	;
	v146 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v129), float64(0.9)))
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v146
	if v108 != int64(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v156 = v84
	goto L33
L30:
	;
	goto L31
L31:
	;
	F_pfree(m, v107)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L20
	} else {
		goto L50
	}
L32:
	;
	v223 = v216
	v227 = v84
	goto L38
L33:
	;
	v167 = v107 + v156*int32(24)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+16)))
	if v168 != int32(1) {
		v216 = v156
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v216 = int32(0)
	goto L32
L35:
	;
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v167)))
	v172 = int64(23)
	v175 = int64(2388976653695081527)
	v176 = (int64(base.Ui64(v171)>>(uint(v172)%64)) ^ v171) * v175
	v177 = int64(47)
	v182 = int64(-8645972361240307355)
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v167)+8))
	v189 = (int64(base.Ui64(v184)>>(uint(v172)%64)) ^ v184) * v175
	v195 = ((v176^int64(base.Ui64(v176)>>(uint(v177)%64))^int64(-9208349263878056368))*v182 ^ int64(base.Ui64(v189)>>(uint(v177)%64)) ^ v189) * v182
	v200 = (int64(base.Ui64(v195)>>(uint(v172)%64)) ^ v195) * v175
	if v137&base.I32_wrap_i64(int64(base.Ui64(v200)>>(uint(v177)%64))^v200-int64(base.Ui64(v200)>>(uint(int64(32))%64))) == v156 {
		v216 = v156
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v211 = v156 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v211)) < base.Ui64(v108) {
		v156 = v211
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	v234 = v107 + v223*int32(24)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+16)))
	if v235 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L31
L40:
	;
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v234)))
	v239 = int64(23)
	v242 = int64(2388976653695081527)
	v243 = (int64(base.Ui64(v238)>>(uint(v239)%64)) ^ v238) * v242
	v244 = int64(47)
	v249 = int64(-8645972361240307355)
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v234)+8))
	v256 = (int64(base.Ui64(v251)>>(uint(v239)%64)) ^ v251) * v242
	v262 = ((v243^int64(base.Ui64(v243)>>(uint(v244)%64))^int64(-9208349263878056368))*v249 ^ int64(base.Ui64(v256)>>(uint(v244)%64)) ^ v256) * v249
	v267 = (int64(base.Ui64(v262)>>(uint(v239)%64)) ^ v262) * v242
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v285 = base.I32_wrap_i64(int64(base.Ui64(v267)>>(uint(v244)%64)) ^ v267 - int64(base.Ui64(v267)>>(uint(int64(32))%64)))
	goto L43
L41:
	;
	goto L42
L42:
	;
	v320 = v223 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v320)) < base.Ui64(v108) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v291 = v285 & v275
	v296 = v114 + v291*int32(24)
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+16)))
	if v297 != 0 {
		v285 = v291 + int32(1)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v234)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v296)+16)) = v298
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v234)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v296)+8)) = v300
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v234)))
	*(*int64)(unsafe.Add(mBase, uint32(v296))) = v302
	goto L42
L45:
	;
	goto L44
L46:
	;
	v324 = v320
	goto L48
L47:
	;
	v324 = int32(0)
	goto L48
L48:
	;
	v326 = v227 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v326)) < base.Ui64(v108) {
		v223 = v324
		v227 = v326
		goto L38
	} else {
		goto L49
	}
L49:
	;
	goto L39
L50:
	;
	goto L10
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v595)
	m.G0 = v18 + int32(16)
	return v587
L52:
	;
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v362 = v352
	v365 = int32(0)
	v366 = v349
	goto L56
L53:
	;
	v561 = v352
	goto L54
L54:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v572 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v571 + v572
	v575 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v561)+8)) = v575
	v577 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v561))) = v577
	*(*uint8)(unsafe.Add(mBase, uint32(v561)+16)) = uint8(v572)
	v587 = v561
	v595 = int32(0)
	goto L51
L55:
	;
	v561 = v546
	goto L54
L56:
	;
	v372 = *(*int64)(unsafe.Add(mBase, uint32(v362)))
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v362)+8))
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	if v372^v355|(v374^v375) == int64(0) {
		v587 = v362
		v595 = int32(1)
		goto L51
	} else {
		goto L58
	}
L57:
	;
	v546 = v539
	goto L55
L58:
	;
	v381 = v366 + int32(1)
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v362)))
	v383 = int64(23)
	v386 = int64(2388976653695081527)
	v387 = (int64(base.Ui64(v382)>>(uint(v383)%64)) ^ v382) * v386
	v388 = int64(47)
	v393 = int64(-8645972361240307355)
	v395 = *(*int64)(unsafe.Add(mBase, uint32(v362)+8))
	v400 = (int64(base.Ui64(v395)>>(uint(v383)%64)) ^ v395) * v386
	v406 = ((v387^int64(base.Ui64(v387)>>(uint(v388)%64))^int64(-9208349263878056368))*v393 ^ int64(base.Ui64(v400)>>(uint(v388)%64)) ^ v400) * v393
	v411 = (int64(base.Ui64(v406)>>(uint(v383)%64)) ^ v406) * v386
	v419 = v348 & base.I32_wrap_i64(int64(base.Ui64(v411)>>(uint(v388)%64))^v411-int64(base.Ui64(v411)>>(uint(int64(32))%64)))
	if base.Ui32(v366) < base.Ui32(v419) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v423 = v366 + v421
	goto L61
L60:
	;
	v423 = v366
	goto L61
L61:
	;
	if base.Ui32(v423-v419) < base.Ui32(v365) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v426 = v381 & v348
	v429 = v347 + v426*int32(24)
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+16)))
	if v430 != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	v526 = v365 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v526) {
		goto L81
	} else {
		goto L82
	}
L65:
	;
	v439 = int32(0)
	v440 = v426
	goto L68
L66:
	;
	v471 = v429
	v473 = v426
	goto L67
L67:
	;
	if v473 != v366 {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	v448 = v439 + int32(1)
	if int32(151) <= v448 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v471 = v463
	v473 = v460
	goto L67
L70:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v453 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v451), base.F64_convert_i64_u(v453)), float64(0.1)) != 0 {
		goto L4
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v460 = (v440 + int32(1)) & v348
	v463 = v347 + v460*int32(24)
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463)+16)))
	if v464 != 0 {
		v439 = v448
		v440 = v460
		goto L68
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	goto L69
L75:
	;
	v487 = v471
	v489 = v473
	goto L78
L76:
	;
	goto L77
L77:
	;
	v546 = v362
	goto L55
L78:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v499 = v496 & (v489 - int32(1))
	v502 = v347 + v499*int32(24)
	v503 = *(*int64)(unsafe.Add(mBase, uint32(v502)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v487)+16)) = v503
	v505 = *(*int64)(unsafe.Add(mBase, uint32(v502)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v487)+8)) = v505
	v507 = *(*int64)(unsafe.Add(mBase, uint32(v502)))
	*(*int64)(unsafe.Add(mBase, uint32(v487))) = v507
	if v499 != v366 {
		v487 = v502
		v489 = v499
		goto L78
	} else {
		goto L80
	}
L79:
	;
	goto L77
L80:
	;
	goto L79
L81:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v531 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v529), base.F64_convert_i64_u(v531)), float64(0.1)) != 0 {
		goto L4
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v536 = v381 & v348
	v539 = v347 + v536*int32(24)
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539)+16)))
	if v540 != 0 {
		v362 = v539
		v365 = v526
		v366 = v536
		goto L56
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	goto L57
L86:
	;
	F_errmsg_internal(m, int32(_a_F_pgstat_snapshot_insert_0), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L20
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_pgstat_snapshot_insert_1), int32(630), int32(_a_F_pgstat_snapshot_insert_2))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L20
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_errmsg_internal(m, int32(_a_F_pgstat_snapshot_insert_3), int32(0))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L20
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_pgstat_snapshot_insert_1), int32(327), int32(_a_F_pgstat_snapshot_insert_4))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L20
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstat_subscription_flush_cb(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
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
			*(*int64)(unsafe.Add(mBase, uint32(v5)+32)) = v15 + v16
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v5)+40))
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v6)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+40)) = v19 + v20
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v5)+48))
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+48)) = v23 + v24
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v5)+56))
			v28 = *(*int64)(unsafe.Add(mBase, uint32(v6)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+56)) = v27 + v28
			v31 = *(*int64)(unsafe.Add(mBase, uint32(v5)+64))
			v32 = *(*int64)(unsafe.Add(mBase, uint32(v6)+40))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+64)) = v31 + v32
			v35 = *(*int64)(unsafe.Add(mBase, uint32(v5)+72))
			v36 = *(*int64)(unsafe.Add(mBase, uint32(v6)+48))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+72)) = v35 + v36
			v39 = *(*int64)(unsafe.Add(mBase, uint32(v5)+80))
			v40 = *(*int64)(unsafe.Add(mBase, uint32(v6)+56))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+80)) = v39 + v40
			v43 = *(*int64)(unsafe.Add(mBase, uint32(v5)+88))
			v44 = *(*int64)(unsafe.Add(mBase, uint32(v6)+64))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+88)) = v43 + v44
			F_pgstat_unlock_entry(m, l0)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				return v7
			}
		} else {
			return v7
		}
	}
}
func F_pgstat_wal_init_backend_cb(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v7 int64
	_ = v7
	var v11 int64
	_ = v11
	var v15 int64
	_ = v15
	v3 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_init_backend_cb[0]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_init_backend_cb[1])) = v3
	v7 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_init_backend_cb[2]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_init_backend_cb[3])) = v7
	v11 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_init_backend_cb[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_init_backend_cb[5])) = v11
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_init_backend_cb[6]))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_init_backend_cb[7])) = v15
	return
}
