package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ShutdownRecoveryTransactionEnvironment(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[0]))
	if v7 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[1]))
		v13 = F_LWLockAcquire(m, v9+int32(512), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[2]))
			v19 = F_errstart(m, int32(11), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				if v19 != 0 {
					F_errmsg_internal(m, int32(_a_F_ShutdownRecoveryTransactionEnvironment_0), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ShutdownRecoveryTransactionEnvironment_1), int32(_a_F_ShutdownRecoveryTransactionEnvironment_2), int32(_a_F_ShutdownRecoveryTransactionEnvironment_3))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							v30 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v30
							*(*int64)(unsafe.Add(mBase, uint32(v16)+12)) = int64(0)
							v35 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[3]))
							v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+8))
							v38 = v36 - int64(1)
							if base.B2i32(base.Ui64(v38) < base.Ui64(int64(3)))|base.B2i32(base.Ui32(int32(2)) < base.Ui32(base.I32_wrap_i64(v38))) == v30 {
								v51 = v38
								for {
									v53 = v51 - int64(1)
									if base.Ui32(base.I32_wrap_i64(v53)) < base.Ui32(int32(3)) {
										v51 = v53
										continue
									} else {
										break
									}
									break
								}
								v61 = v53
							} else {
								v61 = v38
							}
							*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v61
							v63 = *(*int64)(unsafe.Add(mBase, uint32(v35)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v35)+56)) = v63 + int64(1)
							v68 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(0)
							v72 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[1]))
							F_LWLockRelease(m, v72+int32(512))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								F_StandbyReleaseAllLocks(m)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[0]))
									F_hash_destroy(m, v80)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[4]))
										F_hash_destroy(m, v84)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											v88 = int32(0)
											*(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[4])) = v88
											*(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[0])) = v88
											v93 = m.G0
											v95 = v93 - int32(16)
											m.G0 = v95
											v98 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[5]))
											v102 = F_LWLockAcquire(m, v98+int32(584), v88)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return
											} else {
												v105 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[5]))
												v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+612))
												v107 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v105)+612)) = v107
												v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+608)))
												*(*uint8)(unsafe.Add(mBase, uint32(v105)+608)) = uint8(v107)
												F_LWLockRelease(m, v105+int32(584))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return
												} else {
													v116 = int32(0)
													if v109|base.B2i32(v106 == v116) == v116 {
														*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = int64(73746443898191872)
														*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v106
														v125 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[6]))
														*(*int32)(unsafe.Add(mBase, uint32(v95))) = v125
														v129 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[5]))
														F_LockRefindAndRelease(m, int32(_a_F_ShutdownRecoveryTransactionEnvironment_4), v129, v95, int32(7), int32(0))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return
														} else {
															m.G0 = v95 + int32(16)
															return
														}
													} else {
														m.G0 = v95 + int32(16)
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
				} else {
					v30 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v30
					*(*int64)(unsafe.Add(mBase, uint32(v16)+12)) = int64(0)
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[3]))
					v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+8))
					v38 = v36 - int64(1)
					if base.B2i32(base.Ui64(v38) < base.Ui64(int64(3)))|base.B2i32(base.Ui32(int32(2)) < base.Ui32(base.I32_wrap_i64(v38))) == v30 {
						v51 = v38
						for {
							v53 = v51 - int64(1)
							if base.Ui32(base.I32_wrap_i64(v53)) < base.Ui32(int32(3)) {
								v51 = v53
								continue
							} else {
								break
							}
							break
						}
						v61 = v53
					} else {
						v61 = v38
					}
					*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v61
					v63 = *(*int64)(unsafe.Add(mBase, uint32(v35)+56))
					*(*int64)(unsafe.Add(mBase, uint32(v35)+56)) = v63 + int64(1)
					v68 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(0)
					v72 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[1]))
					F_LWLockRelease(m, v72+int32(512))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						F_StandbyReleaseAllLocks(m)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							v80 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[0]))
							F_hash_destroy(m, v80)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[4]))
								F_hash_destroy(m, v84)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									v88 = int32(0)
									*(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[4])) = v88
									*(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[0])) = v88
									v93 = m.G0
									v95 = v93 - int32(16)
									m.G0 = v95
									v98 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[5]))
									v102 = F_LWLockAcquire(m, v98+int32(584), v88)
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return
									} else {
										v105 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[5]))
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+612))
										v107 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v105)+612)) = v107
										v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+608)))
										*(*uint8)(unsafe.Add(mBase, uint32(v105)+608)) = uint8(v107)
										F_LWLockRelease(m, v105+int32(584))
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											v116 = int32(0)
											if v109|base.B2i32(v106 == v116) == v116 {
												*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = int64(73746443898191872)
												*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v106
												v125 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[6]))
												*(*int32)(unsafe.Add(mBase, uint32(v95))) = v125
												v129 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownRecoveryTransactionEnvironment[5]))
												F_LockRefindAndRelease(m, int32(_a_F_ShutdownRecoveryTransactionEnvironment_4), v129, v95, int32(7), int32(0))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return
												} else {
													m.G0 = v95 + int32(16)
													return
												}
											} else {
												m.G0 = v95 + int32(16)
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
	} else {
		return
	}
}
func F_shutdown_MultiFuncCall(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	F_MemoryContextDelete(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
