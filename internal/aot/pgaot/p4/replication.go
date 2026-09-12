package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
		v15 = F_pg_snprintf(m, l2, int32(64), int32(38450), v7+int32(16))
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
		v20 = F_pg_snprintf(m, l2, int32(64), int32(38501), v7)
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[853]))
	v6 = base.I32_div_s(l0-v3, int32(288))
	return v6
}
func F_ReplicationSlotReserveWal(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
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
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[841]))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v17 = F_LWLockAcquire(m, v13+int32(4608), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
		if v19 == int32(0) {
			v22 = F_GetRedoRecPtr(m)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v42 = v22
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
				if v43 != 0 {
					F_s_lock(m, v11, int32(492874), int32(1611), int32(314324))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v42
						F_ReplicationSlotsComputeRequiredLSN(m)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v57 = int64(*(*int32)(unsafe.Add(mBase, _consts[271])))
							v58 = *(*int64)(unsafe.Add(mBase, uint32(v11)+104))
							v59 = F_XLogGetLastRemovedSegno(m)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v61 = base.I64_div_u_s(v58, v57)
								if base.Ui64(v59) < base.Ui64(v61) {
									v64 = *(*int32)(unsafe.Add(mBase, _consts[86]))
									F_LWLockRelease(m, v64+int32(4608))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
										if v71 == int32(1) {
											v76 = *(*int32)(unsafe.Add(mBase, _consts[199]))
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+316))
											v79 = base.B2i32(v77 != int32(2))
											*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v79)
											v81 = v79
										} else {
											v81 = int32(0)
										}
										if v81 != 0 {
											m.G0 = v8 + int32(16)
											return
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
											if v82 == int32(0) {
												m.G0 = v8 + int32(16)
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
														m.G0 = v8 + int32(16)
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
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11 + int32(24)
										F_errmsg_internal(m, int32(18900), v8)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											F_errfinish(m, int32(492874), int32(1622), int32(314324))
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
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
					*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v42
					F_ReplicationSlotsComputeRequiredLSN(m)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						v57 = int64(*(*int32)(unsafe.Add(mBase, _consts[271])))
						v58 = *(*int64)(unsafe.Add(mBase, uint32(v11)+104))
						v59 = F_XLogGetLastRemovedSegno(m)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							v61 = base.I64_div_u_s(v58, v57)
							if base.Ui64(v59) < base.Ui64(v61) {
								v64 = *(*int32)(unsafe.Add(mBase, _consts[86]))
								F_LWLockRelease(m, v64+int32(4608))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
									if v71 == int32(1) {
										v76 = *(*int32)(unsafe.Add(mBase, _consts[199]))
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+316))
										v79 = base.B2i32(v77 != int32(2))
										*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v79)
										v81 = v79
									} else {
										v81 = int32(0)
									}
									if v81 != 0 {
										m.G0 = v8 + int32(16)
										return
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
										if v82 == int32(0) {
											m.G0 = v8 + int32(16)
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
													m.G0 = v8 + int32(16)
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
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11 + int32(24)
									F_errmsg_internal(m, int32(18900), v8)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return
									} else {
										F_errfinish(m, int32(492874), int32(1622), int32(314324))
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
			v26 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
			if v26 == int32(1) {
				v31 = *(*int32)(unsafe.Add(mBase, _consts[199]))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+316))
				v34 = base.B2i32(v32 != int32(2))
				*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v34)
				v36 = v34
			} else {
				v36 = int32(0)
			}
			if v36 != 0 {
				v38 = F_GetXLogReplayRecPtr(m, int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v42 = v38
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
					if v43 != 0 {
						F_s_lock(m, v11, int32(492874), int32(1611), int32(314324))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v42
							F_ReplicationSlotsComputeRequiredLSN(m)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								v57 = int64(*(*int32)(unsafe.Add(mBase, _consts[271])))
								v58 = *(*int64)(unsafe.Add(mBase, uint32(v11)+104))
								v59 = F_XLogGetLastRemovedSegno(m)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									v61 = base.I64_div_u_s(v58, v57)
									if base.Ui64(v59) < base.Ui64(v61) {
										v64 = *(*int32)(unsafe.Add(mBase, _consts[86]))
										F_LWLockRelease(m, v64+int32(4608))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
											if v71 == int32(1) {
												v76 = *(*int32)(unsafe.Add(mBase, _consts[199]))
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+316))
												v79 = base.B2i32(v77 != int32(2))
												*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v79)
												v81 = v79
											} else {
												v81 = int32(0)
											}
											if v81 != 0 {
												m.G0 = v8 + int32(16)
												return
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
												if v82 == int32(0) {
													m.G0 = v8 + int32(16)
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
															m.G0 = v8 + int32(16)
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
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11 + int32(24)
											F_errmsg_internal(m, int32(18900), v8)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												F_errfinish(m, int32(492874), int32(1622), int32(314324))
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
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v42
						F_ReplicationSlotsComputeRequiredLSN(m)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v57 = int64(*(*int32)(unsafe.Add(mBase, _consts[271])))
							v58 = *(*int64)(unsafe.Add(mBase, uint32(v11)+104))
							v59 = F_XLogGetLastRemovedSegno(m)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v61 = base.I64_div_u_s(v58, v57)
								if base.Ui64(v59) < base.Ui64(v61) {
									v64 = *(*int32)(unsafe.Add(mBase, _consts[86]))
									F_LWLockRelease(m, v64+int32(4608))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
										if v71 == int32(1) {
											v76 = *(*int32)(unsafe.Add(mBase, _consts[199]))
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+316))
											v79 = base.B2i32(v77 != int32(2))
											*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v79)
											v81 = v79
										} else {
											v81 = int32(0)
										}
										if v81 != 0 {
											m.G0 = v8 + int32(16)
											return
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
											if v82 == int32(0) {
												m.G0 = v8 + int32(16)
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
														m.G0 = v8 + int32(16)
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
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11 + int32(24)
										F_errmsg_internal(m, int32(18900), v8)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											F_errfinish(m, int32(492874), int32(1622), int32(314324))
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
				v40 = F_GetXLogInsertRecPtr(m)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					v42 = v40
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
					if v43 != 0 {
						F_s_lock(m, v11, int32(492874), int32(1611), int32(314324))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v42
							F_ReplicationSlotsComputeRequiredLSN(m)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								v57 = int64(*(*int32)(unsafe.Add(mBase, _consts[271])))
								v58 = *(*int64)(unsafe.Add(mBase, uint32(v11)+104))
								v59 = F_XLogGetLastRemovedSegno(m)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									v61 = base.I64_div_u_s(v58, v57)
									if base.Ui64(v59) < base.Ui64(v61) {
										v64 = *(*int32)(unsafe.Add(mBase, _consts[86]))
										F_LWLockRelease(m, v64+int32(4608))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
											if v71 == int32(1) {
												v76 = *(*int32)(unsafe.Add(mBase, _consts[199]))
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+316))
												v79 = base.B2i32(v77 != int32(2))
												*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v79)
												v81 = v79
											} else {
												v81 = int32(0)
											}
											if v81 != 0 {
												m.G0 = v8 + int32(16)
												return
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
												if v82 == int32(0) {
													m.G0 = v8 + int32(16)
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
															m.G0 = v8 + int32(16)
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
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11 + int32(24)
											F_errmsg_internal(m, int32(18900), v8)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												F_errfinish(m, int32(492874), int32(1622), int32(314324))
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
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v42
						F_ReplicationSlotsComputeRequiredLSN(m)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v57 = int64(*(*int32)(unsafe.Add(mBase, _consts[271])))
							v58 = *(*int64)(unsafe.Add(mBase, uint32(v11)+104))
							v59 = F_XLogGetLastRemovedSegno(m)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v61 = base.I64_div_u_s(v58, v57)
								if base.Ui64(v59) < base.Ui64(v61) {
									v64 = *(*int32)(unsafe.Add(mBase, _consts[86]))
									F_LWLockRelease(m, v64+int32(4608))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
										if v71 == int32(1) {
											v76 = *(*int32)(unsafe.Add(mBase, _consts[199]))
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+316))
											v79 = base.B2i32(v77 != int32(2))
											*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v79)
											v81 = v79
										} else {
											v81 = int32(0)
										}
										if v81 != 0 {
											m.G0 = v8 + int32(16)
											return
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
											if v82 == int32(0) {
												m.G0 = v8 + int32(16)
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
														m.G0 = v8 + int32(16)
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
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11 + int32(24)
										F_errmsg_internal(m, int32(18900), v8)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											F_errfinish(m, int32(492874), int32(1622), int32(314324))
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[841]))
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
	var v4 int64
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	v4 = int64(0)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v12 = F_LWLockAcquire(m, v8+int32(4736), int32(1))
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
	v15 = *(*int32)(unsafe.Add(mBase, _consts[852]))
	if int32(0) < v15 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[853]))
	v21 = v19
	v22 = int32(0)
	v24 = v4
	goto L6
L4:
	;
	v73 = v4
	goto L5
L5:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v76+int32(4736))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L29
	}
L6:
	;
	v28 = v21 + v22*int32(288)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)))
	if v29 != int32(1) {
		v60 = v21
		v62 = v24
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v73 = v62
	goto L5
L8:
	;
	v65 = v22 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, _consts[852]))
	if v65 < v67 {
		v21 = v60
		v22 = v65
		v24 = v62
		goto L6
	} else {
		goto L28
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(1)
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_s_lock(m, v28, int32(492874), int32(1244), int32(527980))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, _consts[853]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v28)+112))
	if v44 != 0 {
		v60 = v43
		v62 = v24
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v28)+104))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v28)+280))
	if base.Ui64(v45) < base.Ui64(v46) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v48 = v45
	goto L17
L16:
	;
	v48 = v46
	goto L17
L17:
	;
	if v46 != int64(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v51 = v48
	goto L20
L19:
	;
	v51 = v45
	goto L20
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v28)+92))
	if v52 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v53 = v45
	goto L23
L22:
	;
	v53 = v51
	goto L23
L23:
	;
	if v53 == int64(0) {
		v60 = v43
		v62 = v24
		goto L8
	} else {
		goto L24
	}
L24:
	;
	if base.Ui64(v24-int64(1)) < base.Ui64(v53) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v59 = v24
	goto L27
L26:
	;
	v59 = v53
	goto L27
L27:
	;
	v60 = v43
	v62 = v59
	goto L8
L28:
	;
	goto L7
L29:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+440)) = int32(1)
	if v83 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	F_s_lock(m, v87+int32(440), int32(497943), int32(2670), int32(527889))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+440)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+224)) = v73
	return
L33:
	;
	goto L32
}
