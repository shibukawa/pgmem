package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ReplicationOriginNameForLogicalRep(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l1 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
		v15 = F_pg_snprintf(m, l2, int32(64), int32(_a_F_ReplicationOriginNameForLogicalRep_0), v7+int32(16))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			m.G0 = v7 + int32(32)
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		v20 = F_pg_snprintf(m, l2, int32(64), int32(_a_F_ReplicationOriginNameForLogicalRep_1), v7)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			m.G0 = v7 + int32(32)
			return
		}
	}
}
func F_ReplicationSlotIndex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotIndex[0]))
	v6 = base.I32_div_s(l0-v3, int32(296))
	return v6
}
func F_ReplicationSlotReserveWal(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
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
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[0]))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[1]))
	v16 = F_LWLockAcquire(m, v12+int32(_a_F_ReplicationSlotReserveWal_0), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
		if v18 == int32(0) {
			v21 = F_GetRedoRecPtr(m)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v41 = v21
				v44 = base.AtomicRmwXchg32(m, v10, int32(0), int32(1))
				if v44 != 0 {
					F_s_lock(m, v10, int32(_a_F_ReplicationSlotReserveWal_1))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v41
						v49 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10))), uint32(v49))
						F_ReplicationSlotsComputeRequiredLSN(m)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							v55 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[2])))
							v56 = *(*int64)(unsafe.Add(mBase, uint32(v10)+104))
							v57 = F_XLogGetLastRemovedSegno(m)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								v59 = base.I64_div_u_s(v56, v55)
								if base.Ui64(v57) < base.Ui64(v59) {
									v62 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[1]))
									F_LWLockRelease(m, v62+int32(_a_F_ReplicationSlotReserveWal_0))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])))
										if v69 == int32(1) {
											v74 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[4]))
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+308))
											v77 = base.B2i32(v75 != int32(2))
											*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])) = uint8(v77)
											v79 = v77
										} else {
											v79 = int32(0)
										}
										if v79 != 0 {
											m.G0 = v7 + int32(16)
											return
										} else {
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
											if v80 == int32(0) {
												m.G0 = v7 + int32(16)
												return
											} else {
												v83 = F_LogStandbySnapshot(m)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return
												} else {
													F_XLogFlush(m, v83)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														m.G0 = v7 + int32(16)
														return
													}
												}
											}
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 + int32(24)
										F_errmsg_internal(m, int32(_a_F_ReplicationSlotReserveWal_2), v7)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ReplicationSlotReserveWal_3), int32(1768), int32(_a_F_ReplicationSlotReserveWal_4))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
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
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v41
					v49 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10))), uint32(v49))
					F_ReplicationSlotsComputeRequiredLSN(m)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						v55 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[2])))
						v56 = *(*int64)(unsafe.Add(mBase, uint32(v10)+104))
						v57 = F_XLogGetLastRemovedSegno(m)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							v59 = base.I64_div_u_s(v56, v55)
							if base.Ui64(v57) < base.Ui64(v59) {
								v62 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[1]))
								F_LWLockRelease(m, v62+int32(_a_F_ReplicationSlotReserveWal_0))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])))
									if v69 == int32(1) {
										v74 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[4]))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+308))
										v77 = base.B2i32(v75 != int32(2))
										*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])) = uint8(v77)
										v79 = v77
									} else {
										v79 = int32(0)
									}
									if v79 != 0 {
										m.G0 = v7 + int32(16)
										return
									} else {
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
										if v80 == int32(0) {
											m.G0 = v7 + int32(16)
											return
										} else {
											v83 = F_LogStandbySnapshot(m)
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												F_XLogFlush(m, v83)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													m.G0 = v7 + int32(16)
													return
												}
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 + int32(24)
									F_errmsg_internal(m, int32(_a_F_ReplicationSlotReserveWal_2), v7)
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_ReplicationSlotReserveWal_3), int32(1768), int32(_a_F_ReplicationSlotReserveWal_4))
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
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
		} else {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])))
			if v25 == int32(1) {
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[4]))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+308))
				v33 = base.B2i32(v31 != int32(2))
				*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])) = uint8(v33)
				v35 = v33
			} else {
				v35 = int32(0)
			}
			if v35 != 0 {
				v37 = F_GetXLogReplayRecPtr(m, int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v41 = v37
					v44 = base.AtomicRmwXchg32(m, v10, int32(0), int32(1))
					if v44 != 0 {
						F_s_lock(m, v10, int32(_a_F_ReplicationSlotReserveWal_1))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v41
							v49 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10))), uint32(v49))
							F_ReplicationSlotsComputeRequiredLSN(m)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								v55 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[2])))
								v56 = *(*int64)(unsafe.Add(mBase, uint32(v10)+104))
								v57 = F_XLogGetLastRemovedSegno(m)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									v59 = base.I64_div_u_s(v56, v55)
									if base.Ui64(v57) < base.Ui64(v59) {
										v62 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[1]))
										F_LWLockRelease(m, v62+int32(_a_F_ReplicationSlotReserveWal_0))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])))
											if v69 == int32(1) {
												v74 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[4]))
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+308))
												v77 = base.B2i32(v75 != int32(2))
												*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])) = uint8(v77)
												v79 = v77
											} else {
												v79 = int32(0)
											}
											if v79 != 0 {
												m.G0 = v7 + int32(16)
												return
											} else {
												v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
												if v80 == int32(0) {
													m.G0 = v7 + int32(16)
													return
												} else {
													v83 = F_LogStandbySnapshot(m)
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return
													} else {
														F_XLogFlush(m, v83)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return
														} else {
															m.G0 = v7 + int32(16)
															return
														}
													}
												}
											}
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 + int32(24)
											F_errmsg_internal(m, int32(_a_F_ReplicationSlotReserveWal_2), v7)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_ReplicationSlotReserveWal_3), int32(1768), int32(_a_F_ReplicationSlotReserveWal_4))
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
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
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v41
						v49 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10))), uint32(v49))
						F_ReplicationSlotsComputeRequiredLSN(m)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							v55 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[2])))
							v56 = *(*int64)(unsafe.Add(mBase, uint32(v10)+104))
							v57 = F_XLogGetLastRemovedSegno(m)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								v59 = base.I64_div_u_s(v56, v55)
								if base.Ui64(v57) < base.Ui64(v59) {
									v62 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[1]))
									F_LWLockRelease(m, v62+int32(_a_F_ReplicationSlotReserveWal_0))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])))
										if v69 == int32(1) {
											v74 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[4]))
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+308))
											v77 = base.B2i32(v75 != int32(2))
											*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])) = uint8(v77)
											v79 = v77
										} else {
											v79 = int32(0)
										}
										if v79 != 0 {
											m.G0 = v7 + int32(16)
											return
										} else {
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
											if v80 == int32(0) {
												m.G0 = v7 + int32(16)
												return
											} else {
												v83 = F_LogStandbySnapshot(m)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return
												} else {
													F_XLogFlush(m, v83)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														m.G0 = v7 + int32(16)
														return
													}
												}
											}
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 + int32(24)
										F_errmsg_internal(m, int32(_a_F_ReplicationSlotReserveWal_2), v7)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ReplicationSlotReserveWal_3), int32(1768), int32(_a_F_ReplicationSlotReserveWal_4))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
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
			} else {
				v39 = F_GetXLogInsertRecPtr(m)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v41 = v39
					v44 = base.AtomicRmwXchg32(m, v10, int32(0), int32(1))
					if v44 != 0 {
						F_s_lock(m, v10, int32(_a_F_ReplicationSlotReserveWal_1))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v41
							v49 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10))), uint32(v49))
							F_ReplicationSlotsComputeRequiredLSN(m)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								v55 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[2])))
								v56 = *(*int64)(unsafe.Add(mBase, uint32(v10)+104))
								v57 = F_XLogGetLastRemovedSegno(m)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									v59 = base.I64_div_u_s(v56, v55)
									if base.Ui64(v57) < base.Ui64(v59) {
										v62 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[1]))
										F_LWLockRelease(m, v62+int32(_a_F_ReplicationSlotReserveWal_0))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])))
											if v69 == int32(1) {
												v74 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[4]))
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+308))
												v77 = base.B2i32(v75 != int32(2))
												*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])) = uint8(v77)
												v79 = v77
											} else {
												v79 = int32(0)
											}
											if v79 != 0 {
												m.G0 = v7 + int32(16)
												return
											} else {
												v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
												if v80 == int32(0) {
													m.G0 = v7 + int32(16)
													return
												} else {
													v83 = F_LogStandbySnapshot(m)
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return
													} else {
														F_XLogFlush(m, v83)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return
														} else {
															m.G0 = v7 + int32(16)
															return
														}
													}
												}
											}
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 + int32(24)
											F_errmsg_internal(m, int32(_a_F_ReplicationSlotReserveWal_2), v7)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_ReplicationSlotReserveWal_3), int32(1768), int32(_a_F_ReplicationSlotReserveWal_4))
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
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
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v41
						v49 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10))), uint32(v49))
						F_ReplicationSlotsComputeRequiredLSN(m)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							v55 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[2])))
							v56 = *(*int64)(unsafe.Add(mBase, uint32(v10)+104))
							v57 = F_XLogGetLastRemovedSegno(m)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								v59 = base.I64_div_u_s(v56, v55)
								if base.Ui64(v57) < base.Ui64(v59) {
									v62 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[1]))
									F_LWLockRelease(m, v62+int32(_a_F_ReplicationSlotReserveWal_0))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])))
										if v69 == int32(1) {
											v74 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[4]))
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+308))
											v77 = base.B2i32(v75 != int32(2))
											*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])) = uint8(v77)
											v79 = v77
										} else {
											v79 = int32(0)
										}
										if v79 != 0 {
											m.G0 = v7 + int32(16)
											return
										} else {
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
											if v80 == int32(0) {
												m.G0 = v7 + int32(16)
												return
											} else {
												v83 = F_LogStandbySnapshot(m)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return
												} else {
													F_XLogFlush(m, v83)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														m.G0 = v7 + int32(16)
														return
													}
												}
											}
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 + int32(24)
										F_errmsg_internal(m, int32(_a_F_ReplicationSlotReserveWal_2), v7)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ReplicationSlotReserveWal_3), int32(1768), int32(_a_F_ReplicationSlotReserveWal_4))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
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
		}
	}
}
func F_ReplicationSlotShmemExit(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotShmemExit[0]))
	if v4 != 0 {
		F_ReplicationSlotRelease(m)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			F_ReplicationSlotCleanup(m, int32(0))
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_ReplicationSlotCleanup(m, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ReplicationSlotsComputeRequiredLSN(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	v6 = int64(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[0]))
	v14 = F_LWLockAcquire(m, v10+int32(_a_F_ReplicationSlotsComputeRequiredLSN_0), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[1]))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[2]))
	if int32(0) < v17+v19 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[3]))
	v26 = v24
	v27 = int32(0)
	v32 = v6
	goto L6
L4:
	;
	v86 = v6
	goto L5
L5:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[0]))
	F_LWLockRelease(m, v88+int32(_a_F_ReplicationSlotsComputeRequiredLSN_0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L29
	}
L6:
	;
	v35 = v26 + v27*int32(296)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+4)))
	if v36 != int32(1) {
		v66 = v26
		v70 = v32
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v86 = v70
	goto L5
L8:
	;
	v72 = v27 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[1]))
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[2]))
	if v72 < v74+v76 {
		v26 = v66
		v27 = v72
		v32 = v70
		goto L6
	} else {
		goto L28
	}
L9:
	;
	v41 = base.AtomicRmwXchg32(m, v35, int32(0), int32(1))
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_s_lock(m, v35, int32(_a_F_ReplicationSlotsComputeRequiredLSN_1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v35)+280))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v35)+112))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v35)+104))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	v49 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v35))), uint32(v49))
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[3]))
	if v46 != 0 {
		v66 = v53
		v70 = v32
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	if base.Ui64(v47) < base.Ui64(v45) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v55 = v47
	goto L17
L16:
	;
	v55 = v45
	goto L17
L17:
	;
	if v45 != int64(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v58 = v55
	goto L20
L19:
	;
	v58 = v47
	goto L20
L20:
	;
	if v48 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v59 = v47
	goto L23
L22:
	;
	v59 = v58
	goto L23
L23:
	;
	if v59 == int64(0) {
		v66 = v53
		v70 = v32
		goto L8
	} else {
		goto L24
	}
L24:
	;
	if base.Ui64(v32-int64(1)) < base.Ui64(v59) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v65 = v32
	goto L27
L26:
	;
	v65 = v59
	goto L27
L27:
	;
	v66 = v53
	v70 = v65
	goto L8
L28:
	;
	goto L7
L29:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[4]))
	v97 = base.AtomicRmwXchg32(m, v94, int32(440), int32(1))
	if v97 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_s_lock(m, v94+int32(440), int32(_a_F_ReplicationSlotsComputeRequiredLSN_2))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v104)+216)) = v86
	v106 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v104)+440)), uint32(v106))
	return
L33:
	;
	goto L32
}
