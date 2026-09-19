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
	v6 = base.I32_div_s(l0-v3, int32(288))
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
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
					F_s_lock(m, v10, int32(_a_F_ReplicationSlotReserveWal_1), int32(1611), int32(_a_F_ReplicationSlotReserveWal_2))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v41
						v51 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10))), uint32(v51))
						F_ReplicationSlotsComputeRequiredLSN(m)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v57 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[2])))
							v58 = *(*int64)(unsafe.Add(mBase, uint32(v10)+104))
							v59 = F_XLogGetLastRemovedSegno(m)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v61 = base.I64_div_u_s(v58, v57)
								if base.Ui64(v59) < base.Ui64(v61) {
									v64 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[1]))
									F_LWLockRelease(m, v64+int32(_a_F_ReplicationSlotReserveWal_0))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])))
										if v71 == int32(1) {
											v76 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[4]))
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+316))
											v79 = base.B2i32(v77 != int32(2))
											*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])) = uint8(v79)
											v81 = v79
										} else {
											v81 = int32(0)
										}
										if v81 != 0 {
											m.G0 = v7 + int32(16)
											return
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
											if v82 == int32(0) {
												m.G0 = v7 + int32(16)
												return
											} else {
												v85 = F_LogStandbySnapshot(m)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													F_XLogFlush(m, v85)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
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
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 + int32(24)
										F_errmsg_internal(m, int32(_a_F_ReplicationSlotReserveWal_3), v7)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ReplicationSlotReserveWal_1), int32(1622), int32(_a_F_ReplicationSlotReserveWal_2))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
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
					v51 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10))), uint32(v51))
					F_ReplicationSlotsComputeRequiredLSN(m)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						v57 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[2])))
						v58 = *(*int64)(unsafe.Add(mBase, uint32(v10)+104))
						v59 = F_XLogGetLastRemovedSegno(m)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							v61 = base.I64_div_u_s(v58, v57)
							if base.Ui64(v59) < base.Ui64(v61) {
								v64 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[1]))
								F_LWLockRelease(m, v64+int32(_a_F_ReplicationSlotReserveWal_0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])))
									if v71 == int32(1) {
										v76 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[4]))
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+316))
										v79 = base.B2i32(v77 != int32(2))
										*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])) = uint8(v79)
										v81 = v79
									} else {
										v81 = int32(0)
									}
									if v81 != 0 {
										m.G0 = v7 + int32(16)
										return
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
										if v82 == int32(0) {
											m.G0 = v7 + int32(16)
											return
										} else {
											v85 = F_LogStandbySnapshot(m)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												F_XLogFlush(m, v85)
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
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
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 + int32(24)
									F_errmsg_internal(m, int32(_a_F_ReplicationSlotReserveWal_3), v7)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_ReplicationSlotReserveWal_1), int32(1622), int32(_a_F_ReplicationSlotReserveWal_2))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
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
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+316))
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
						F_s_lock(m, v10, int32(_a_F_ReplicationSlotReserveWal_1), int32(1611), int32(_a_F_ReplicationSlotReserveWal_2))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v41
							v51 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10))), uint32(v51))
							F_ReplicationSlotsComputeRequiredLSN(m)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								v57 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[2])))
								v58 = *(*int64)(unsafe.Add(mBase, uint32(v10)+104))
								v59 = F_XLogGetLastRemovedSegno(m)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									v61 = base.I64_div_u_s(v58, v57)
									if base.Ui64(v59) < base.Ui64(v61) {
										v64 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[1]))
										F_LWLockRelease(m, v64+int32(_a_F_ReplicationSlotReserveWal_0))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])))
											if v71 == int32(1) {
												v76 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[4]))
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+316))
												v79 = base.B2i32(v77 != int32(2))
												*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])) = uint8(v79)
												v81 = v79
											} else {
												v81 = int32(0)
											}
											if v81 != 0 {
												m.G0 = v7 + int32(16)
												return
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
												if v82 == int32(0) {
													m.G0 = v7 + int32(16)
													return
												} else {
													v85 = F_LogStandbySnapshot(m)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														F_XLogFlush(m, v85)
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
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
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 + int32(24)
											F_errmsg_internal(m, int32(_a_F_ReplicationSlotReserveWal_3), v7)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_ReplicationSlotReserveWal_1), int32(1622), int32(_a_F_ReplicationSlotReserveWal_2))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
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
						v51 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10))), uint32(v51))
						F_ReplicationSlotsComputeRequiredLSN(m)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v57 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[2])))
							v58 = *(*int64)(unsafe.Add(mBase, uint32(v10)+104))
							v59 = F_XLogGetLastRemovedSegno(m)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v61 = base.I64_div_u_s(v58, v57)
								if base.Ui64(v59) < base.Ui64(v61) {
									v64 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[1]))
									F_LWLockRelease(m, v64+int32(_a_F_ReplicationSlotReserveWal_0))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])))
										if v71 == int32(1) {
											v76 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[4]))
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+316))
											v79 = base.B2i32(v77 != int32(2))
											*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])) = uint8(v79)
											v81 = v79
										} else {
											v81 = int32(0)
										}
										if v81 != 0 {
											m.G0 = v7 + int32(16)
											return
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
											if v82 == int32(0) {
												m.G0 = v7 + int32(16)
												return
											} else {
												v85 = F_LogStandbySnapshot(m)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													F_XLogFlush(m, v85)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
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
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 + int32(24)
										F_errmsg_internal(m, int32(_a_F_ReplicationSlotReserveWal_3), v7)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ReplicationSlotReserveWal_1), int32(1622), int32(_a_F_ReplicationSlotReserveWal_2))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
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
						F_s_lock(m, v10, int32(_a_F_ReplicationSlotReserveWal_1), int32(1611), int32(_a_F_ReplicationSlotReserveWal_2))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v41
							v51 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10))), uint32(v51))
							F_ReplicationSlotsComputeRequiredLSN(m)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								v57 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[2])))
								v58 = *(*int64)(unsafe.Add(mBase, uint32(v10)+104))
								v59 = F_XLogGetLastRemovedSegno(m)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									v61 = base.I64_div_u_s(v58, v57)
									if base.Ui64(v59) < base.Ui64(v61) {
										v64 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[1]))
										F_LWLockRelease(m, v64+int32(_a_F_ReplicationSlotReserveWal_0))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])))
											if v71 == int32(1) {
												v76 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[4]))
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+316))
												v79 = base.B2i32(v77 != int32(2))
												*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])) = uint8(v79)
												v81 = v79
											} else {
												v81 = int32(0)
											}
											if v81 != 0 {
												m.G0 = v7 + int32(16)
												return
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
												if v82 == int32(0) {
													m.G0 = v7 + int32(16)
													return
												} else {
													v85 = F_LogStandbySnapshot(m)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														F_XLogFlush(m, v85)
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
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
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 + int32(24)
											F_errmsg_internal(m, int32(_a_F_ReplicationSlotReserveWal_3), v7)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_ReplicationSlotReserveWal_1), int32(1622), int32(_a_F_ReplicationSlotReserveWal_2))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
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
						v51 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10))), uint32(v51))
						F_ReplicationSlotsComputeRequiredLSN(m)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v57 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[2])))
							v58 = *(*int64)(unsafe.Add(mBase, uint32(v10)+104))
							v59 = F_XLogGetLastRemovedSegno(m)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v61 = base.I64_div_u_s(v58, v57)
								if base.Ui64(v59) < base.Ui64(v61) {
									v64 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[1]))
									F_LWLockRelease(m, v64+int32(_a_F_ReplicationSlotReserveWal_0))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])))
										if v71 == int32(1) {
											v76 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[4]))
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+316))
											v79 = base.B2i32(v77 != int32(2))
											*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotReserveWal[3])) = uint8(v79)
											v81 = v79
										} else {
											v81 = int32(0)
										}
										if v81 != 0 {
											m.G0 = v7 + int32(16)
											return
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
											if v82 == int32(0) {
												m.G0 = v7 + int32(16)
												return
											} else {
												v85 = F_LogStandbySnapshot(m)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													F_XLogFlush(m, v85)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
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
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 + int32(24)
										F_errmsg_internal(m, int32(_a_F_ReplicationSlotReserveWal_3), v7)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ReplicationSlotReserveWal_1), int32(1622), int32(_a_F_ReplicationSlotReserveWal_2))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
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
func F_ReplicationSlotShmemExit(m *base.Module, l0 int32, l1 int32) {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
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
	if int32(0) < v17 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[2]))
	v23 = v21
	v24 = int32(0)
	v29 = v6
	goto L6
L4:
	;
	v82 = v6
	goto L5
L5:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[0]))
	F_LWLockRelease(m, v84+int32(_a_F_ReplicationSlotsComputeRequiredLSN_0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L29
	}
L6:
	;
	v32 = v23 + v24*int32(288)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
	if v33 != int32(1) {
		v65 = v23
		v69 = v29
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v82 = v69
	goto L5
L8:
	;
	v71 = v24 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[1]))
	if v71 < v73 {
		v23 = v65
		v24 = v71
		v29 = v69
		goto L6
	} else {
		goto L28
	}
L9:
	;
	v38 = base.AtomicRmwXchg32(m, v32, int32(0), int32(1))
	if v38 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_s_lock(m, v32, int32(_a_F_ReplicationSlotsComputeRequiredLSN_1), int32(1244), int32(_a_F_ReplicationSlotsComputeRequiredLSN_2))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v32)+280))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32)+112))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v32)+104))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v32)+92))
	v48 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v32))), uint32(v48))
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[2]))
	if v45 != 0 {
		v65 = v52
		v69 = v29
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	if base.Ui64(v46) < base.Ui64(v44) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v54 = v46
	goto L17
L16:
	;
	v54 = v44
	goto L17
L17:
	;
	if v44 != int64(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v57 = v54
	goto L20
L19:
	;
	v57 = v46
	goto L20
L20:
	;
	if v47 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v58 = v46
	goto L23
L22:
	;
	v58 = v57
	goto L23
L23:
	;
	if v58 == int64(0) {
		v65 = v52
		v69 = v29
		goto L8
	} else {
		goto L24
	}
L24:
	;
	if base.Ui64(v29-int64(1)) < base.Ui64(v58) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v64 = v29
	goto L27
L26:
	;
	v64 = v58
	goto L27
L27:
	;
	v65 = v52
	v69 = v64
	goto L8
L28:
	;
	goto L7
L29:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[3]))
	v93 = base.AtomicRmwXchg32(m, v90, int32(440), int32(1))
	if v93 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[3]))
	F_s_lock(m, v95+int32(440), int32(_a_F_ReplicationSlotsComputeRequiredLSN_3), int32(2670), int32(_a_F_ReplicationSlotsComputeRequiredLSN_4))
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
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredLSN[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v104)+224)) = v82
	v106 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v104)+440)), uint32(v106))
	return
L33:
	;
	goto L32
}
