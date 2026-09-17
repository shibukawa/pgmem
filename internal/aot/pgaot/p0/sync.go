package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SyncReplicationSlots(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(192)
	m.G0 = v10
	v15 = v2
	v16 = v2
	v17 = v2
	v18 = v2
	v19 = int32(-1)
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v19 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	v135 = int32(m.ExcTag)
	v136 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v135 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = l0
	F_before_shmem_exit(m, int32(1022), l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v40 = v15
	v41 = v16
	v42 = v17
	v43 = v18
	goto L8
L8:
	;
	if v43 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[0]))
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[1]))
	goto L10
L10:
	;
	v34 = v10 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v10 + int32(12)
	goto L13
L11:
	;
	v40 = l0
	v41 = v32
	v42 = v30
	v43 = int32(0)
	goto L8
L13:
	;
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v42
	*(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[1])) = v10 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v40
	F_check_and_set_sync_info(m, int32(-1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[0])) = v42
	*(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[1])) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v40
	F_cancel_before_shmem_exit(m, int32(1022), v40)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L26
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v40
	F_validate_remote_info(m, l0)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v40
	v64 = F_synchronize_slots(m, l0)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v40
	F_ReplicationSlotCleanup(m, int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[2]))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = int32(1)
	if v74 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v40
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[2]))
	F_s_lock(m, v81+int32(16), int32(_a_F_SyncReplicationSlots_0), int32(1339), int32(_a_F_SyncReplicationSlots_1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[2]))
	v91 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v91
	*(*uint8)(unsafe.Add(mBase, uint32(v90)+5)) = uint8(v91)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v42
	*(*uint8)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[3])) = uint8(v91)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v40
	F_cancel_before_shmem_exit(m, int32(1022), v40)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[0])) = v42
	*(*int32)(unsafe.Add(mBase, _c_F_SyncReplicationSlots[1])) = v41
	m.G0 = v10 + int32(192)
	return
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v40
	F_slotsync_failure_callback(m, v10, v40)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v40
	F_pg_re_throw(m)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	goto L4
L29:
	;
	v140 = int32(v136)
	m.G0 = v10
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v10+int32(12) == v146 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	m.ExcPending = 1
	goto L38
L31:
	;
	if v150 != 0 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v150 = v148
	goto L34
L33:
	;
	v150 = int32(0)
	goto L34
L34:
	;
	goto L31
L35:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v10)+188))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v10)+184))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v10)+180))
	v15 = v151
	v16 = v152
	v17 = v153
	v18 = v142
	v19 = v150
	goto L1
L36:
	;
	goto L37
L37:
	;
	F___wasm_longjmp(m, v143, v142)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_finish_sync_worker(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_finish_sync_worker[0]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v9 == int32(2) {
		F_CommitTransactionCommand(m)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v15 = F_pgstat_report_stat(m, int32(1))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v19 = int32(_a_F_finish_sync_worker_0)
				v20 = *(*int32)(unsafe.Add(mBase, _c_F_finish_sync_worker[1]))
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)+280))
				*(*int64)(unsafe.Add(mBase, uint32(v20)+280)) = v21
				*(*int64)(unsafe.Add(mBase, _c_F_finish_sync_worker[2])) = v21
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_finish_sync_worker[1]))
				v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)+272))
				*(*int64)(unsafe.Add(mBase, uint32(v26)+272)) = v27
				*(*int64)(unsafe.Add(mBase, _c_F_finish_sync_worker[3])) = v27
				F_XLogFlush(m, v27)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_StartTransactionCommand(m)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v37 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							if v37 != 0 {
								v40 = *(*int32)(unsafe.Add(mBase, _c_F_finish_sync_worker[4]))
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
								v43 = *(*int32)(unsafe.Add(mBase, _c_F_finish_sync_worker[5]))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
								v45 = F_get_rel_name(m, v44)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v45
									*(*int32)(unsafe.Add(mBase, uint32(v5))) = v41
									F_errmsg(m, int32(_a_F_finish_sync_worker_1), v5)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_finish_sync_worker_2), int32(162), int32(_a_F_finish_sync_worker_3))
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											F_CommitTransactionCommand(m)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												v61 = *(*int32)(unsafe.Add(mBase, _c_F_finish_sync_worker[5]))
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+32))
												F_logicalrep_worker_wakeup(m, v62)
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return
												} else {
													F_proc_exit(m, int32(0))
													mBase = m.M
													v67 = m.ExcPending
													if v67 != 0 {
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
							} else {
								F_CommitTransactionCommand(m)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, _c_F_finish_sync_worker[5]))
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+32))
									F_logicalrep_worker_wakeup(m, v62)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_proc_exit(m, int32(0))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
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
			}
		}
	} else {
		v19 = int32(_a_F_finish_sync_worker_0)
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_finish_sync_worker[1]))
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)+280))
		*(*int64)(unsafe.Add(mBase, uint32(v20)+280)) = v21
		*(*int64)(unsafe.Add(mBase, _c_F_finish_sync_worker[2])) = v21
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_finish_sync_worker[1]))
		v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)+272))
		*(*int64)(unsafe.Add(mBase, uint32(v26)+272)) = v27
		*(*int64)(unsafe.Add(mBase, _c_F_finish_sync_worker[3])) = v27
		F_XLogFlush(m, v27)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			F_StartTransactionCommand(m)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v37 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					if v37 != 0 {
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_finish_sync_worker[4]))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_finish_sync_worker[5]))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
						v45 = F_get_rel_name(m, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v45
							*(*int32)(unsafe.Add(mBase, uint32(v5))) = v41
							F_errmsg(m, int32(_a_F_finish_sync_worker_1), v5)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_finish_sync_worker_2), int32(162), int32(_a_F_finish_sync_worker_3))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									F_CommitTransactionCommand(m)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, _c_F_finish_sync_worker[5]))
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+32))
										F_logicalrep_worker_wakeup(m, v62)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											F_proc_exit(m, int32(0))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
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
					} else {
						F_CommitTransactionCommand(m)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, _c_F_finish_sync_worker[5]))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+32))
							F_logicalrep_worker_wakeup(m, v62)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_proc_exit(m, int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
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
	}
}
