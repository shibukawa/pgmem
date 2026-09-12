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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v115 int32
	_ = v115
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v17 = v2
	v18 = v2
	v19 = v2
	v20 = v2
	v21 = v2
	v22 = v12
	v23 = int32(-1)
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
	if v23 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	v153 = int32(m.ExcTag)
	v154 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v153 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L6:
	;
	v27 = v22 - int32(160)
	m.G0 = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v27
	F_before_shmem_exit(m, int32(1022), l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		v151 = v27
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v46 = v17
	v47 = v18
	v48 = v19
	v49 = v20
	v50 = v21
	v51 = v22
	goto L8
L8:
	;
	if v50 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	v40 = *(*int32)(unsafe.Add(mBase, _consts[294]))
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v12 + int32(12)
	goto L13
L11:
	;
	v46 = v27
	v47 = l0
	v48 = v40
	v49 = v38
	v50 = int32(0)
	v51 = v27
	goto L8
L13:
	;
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[294])) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v46
	F_check_and_set_sync_info(m, int32(-1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		v151 = v51
		goto L5
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v49
	*(*int32)(unsafe.Add(mBase, _consts[294])) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v46
	F_cancel_before_shmem_exit(m, int32(1022), v47)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		v151 = v51
		goto L5
	} else {
		goto L26
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v46
	F_validate_remote_info(m, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		v151 = v51
		goto L5
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v46
	v73 = F_synchronize_slots(m, l0)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		v151 = v51
		goto L5
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v46
	F_ReplicationSlotCleanup(m, int32(1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		v151 = v51
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = int32(1)
	if v84 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v46
	v92 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	F_s_lock(m, v92+int32(16), int32(498109), int32(1339), int32(336800))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		v151 = v51
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	v102 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v101)+16)) = v102
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+5)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v49
	*(*uint8)(unsafe.Add(mBase, _consts[517])) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v46
	F_cancel_before_shmem_exit(m, int32(1022), v47)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		v151 = v51
		goto L5
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v49
	*(*int32)(unsafe.Add(mBase, _consts[294])) = v48
	m.G0 = v12 + int32(32)
	return
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v46
	F_slotsync_failure_callback(m, v12, v47)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		v151 = v51
		goto L5
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v46
	F_pg_re_throw(m)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		v151 = v51
		goto L5
	} else {
		goto L28
	}
L28:
	;
	goto L4
L29:
	;
	v158 = int32(v154)
	m.G0 = v151
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v12+int32(12) == v165 {
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
	if v168 != 0 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v168 = v167
	goto L34
L33:
	;
	v168 = int32(0)
	goto L34
L34:
	;
	goto L31
L35:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v17 = v169
	v18 = v170
	v19 = v171
	v20 = v172
	v21 = v160
	v22 = v151
	v23 = v168
	goto L1
L36:
	;
	goto L37
L37:
	;
	F___wasm_longjmp(m, v161, v160)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[37]))
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
				v19 = int32(4384720)
				v20 = *(*int32)(unsafe.Add(mBase, _consts[3]))
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)+280))
				*(*int64)(unsafe.Add(mBase, uint32(v20)+280)) = v21
				*(*int64)(unsafe.Add(mBase, _consts[531])) = v21
				v26 = *(*int32)(unsafe.Add(mBase, _consts[3]))
				v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)+272))
				*(*int64)(unsafe.Add(mBase, uint32(v26)+272)) = v27
				*(*int64)(unsafe.Add(mBase, _consts[532])) = v27
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
								v40 = *(*int32)(unsafe.Add(mBase, _consts[533]))
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
								v43 = *(*int32)(unsafe.Add(mBase, _consts[508]))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
								v45 = F_get_rel_name(m, v44)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v45
									*(*int32)(unsafe.Add(mBase, uint32(v5))) = v41
									F_errmsg(m, int32(457359), v5)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										F_errfinish(m, int32(498120), int32(162), int32(219917))
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
												v61 = *(*int32)(unsafe.Add(mBase, _consts[508]))
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
									v61 = *(*int32)(unsafe.Add(mBase, _consts[508]))
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
		v19 = int32(4384720)
		v20 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)+280))
		*(*int64)(unsafe.Add(mBase, uint32(v20)+280)) = v21
		*(*int64)(unsafe.Add(mBase, _consts[531])) = v21
		v26 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)+272))
		*(*int64)(unsafe.Add(mBase, uint32(v26)+272)) = v27
		*(*int64)(unsafe.Add(mBase, _consts[532])) = v27
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
						v40 = *(*int32)(unsafe.Add(mBase, _consts[533]))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
						v43 = *(*int32)(unsafe.Add(mBase, _consts[508]))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
						v45 = F_get_rel_name(m, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v45
							*(*int32)(unsafe.Add(mBase, uint32(v5))) = v41
							F_errmsg(m, int32(457359), v5)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								F_errfinish(m, int32(498120), int32(162), int32(219917))
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
										v61 = *(*int32)(unsafe.Add(mBase, _consts[508]))
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
							v61 = *(*int32)(unsafe.Add(mBase, _consts[508]))
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
