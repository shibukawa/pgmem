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
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	v7 = *(*int32)(unsafe.Add(mBase, _consts[794]))
	if v7 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[2]))
		v13 = F_LWLockAcquire(m, v9+int32(512), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[193]))
			v19 = F_errstart(m, int32(11), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				if v19 != 0 {
					F_errmsg_internal(m, int32(166286), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_errfinish(m, int32(475758), int32(5045), int32(324642))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v16)+12)) = int64(0)
							v35 = *(*int32)(unsafe.Add(mBase, _consts[139]))
							v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+8))
							v38 = v36 - int64(1)
							if base.Ui64(v38) < base.Ui64(int64(3)) {
								v58 = v38
							} else {
								if base.Ui32(int32(2)) < base.Ui32(base.I32_wrap_i64(v38)) {
									v58 = v38
								} else {
									v48 = v38
									for {
										v50 = v48 - int64(1)
										if base.Ui32(base.I32_wrap_i64(v50)) < base.Ui32(int32(3)) {
											v48 = v50
											continue
										} else {
											break
										}
										break
									}
									v58 = v50
								}
							}
							*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v58
							v60 = *(*int64)(unsafe.Add(mBase, uint32(v35)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v35)+56)) = v60 + int64(1)
							v65 = *(*int32)(unsafe.Add(mBase, _consts[193]))
							*(*int32)(unsafe.Add(mBase, uint32(v65)+24)) = int32(0)
							v69 = *(*int32)(unsafe.Add(mBase, _consts[2]))
							F_LWLockRelease(m, v69+int32(512))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return
							} else {
								F_StandbyReleaseAllLocks(m)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, _consts[794]))
									F_hash_destroy(m, v77)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										v81 = *(*int32)(unsafe.Add(mBase, _consts[773]))
										F_hash_destroy(m, v81)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											v85 = int32(0)
											*(*int32)(unsafe.Add(mBase, _consts[773])) = v85
											*(*int32)(unsafe.Add(mBase, _consts[794])) = v85
											v90 = m.G0
											v92 = v90 - int32(16)
											m.G0 = v92
											v95 = *(*int32)(unsafe.Add(mBase, _consts[185]))
											v99 = F_LWLockAcquire(m, v95+int32(584), v85)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												v102 = *(*int32)(unsafe.Add(mBase, _consts[185]))
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+612))
												v104 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v102)+612)) = v104
												v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+608)))
												*(*uint8)(unsafe.Add(mBase, uint32(v102)+608)) = uint8(v104)
												F_LWLockRelease(m, v102+int32(584))
												mBase = m.M
												v112 = m.ExcPending
												if v112 != 0 {
													return
												} else {
													if v106 != 0 {
														m.G0 = v92 + int32(16)
														return
													} else {
														if v103 == int32(0) {
															m.G0 = v92 + int32(16)
															return
														} else {
															*(*int64)(unsafe.Add(mBase, uint32(v92)+8)) = int64(73746443898191872)
															*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v103
															v119 = *(*int32)(unsafe.Add(mBase, _consts[743]))
															*(*int32)(unsafe.Add(mBase, uint32(v92))) = v119
															v123 = *(*int32)(unsafe.Add(mBase, _consts[185]))
															F_LockRefindAndRelease(m, int32(1593656), v123, v92, int32(7), int32(0))
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return
															} else {
																m.G0 = v92 + int32(16)
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
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = int32(0)
					*(*int64)(unsafe.Add(mBase, uint32(v16)+12)) = int64(0)
					v35 = *(*int32)(unsafe.Add(mBase, _consts[139]))
					v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+8))
					v38 = v36 - int64(1)
					if base.Ui64(v38) < base.Ui64(int64(3)) {
						v58 = v38
					} else {
						if base.Ui32(int32(2)) < base.Ui32(base.I32_wrap_i64(v38)) {
							v58 = v38
						} else {
							v48 = v38
							for {
								v50 = v48 - int64(1)
								if base.Ui32(base.I32_wrap_i64(v50)) < base.Ui32(int32(3)) {
									v48 = v50
									continue
								} else {
									break
								}
								break
							}
							v58 = v50
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v58
					v60 = *(*int64)(unsafe.Add(mBase, uint32(v35)+56))
					*(*int64)(unsafe.Add(mBase, uint32(v35)+56)) = v60 + int64(1)
					v65 = *(*int32)(unsafe.Add(mBase, _consts[193]))
					*(*int32)(unsafe.Add(mBase, uint32(v65)+24)) = int32(0)
					v69 = *(*int32)(unsafe.Add(mBase, _consts[2]))
					F_LWLockRelease(m, v69+int32(512))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						F_StandbyReleaseAllLocks(m)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, _consts[794]))
							F_hash_destroy(m, v77)
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, _consts[773]))
								F_hash_destroy(m, v81)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									v85 = int32(0)
									*(*int32)(unsafe.Add(mBase, _consts[773])) = v85
									*(*int32)(unsafe.Add(mBase, _consts[794])) = v85
									v90 = m.G0
									v92 = v90 - int32(16)
									m.G0 = v92
									v95 = *(*int32)(unsafe.Add(mBase, _consts[185]))
									v99 = F_LWLockAcquire(m, v95+int32(584), v85)
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return
									} else {
										v102 = *(*int32)(unsafe.Add(mBase, _consts[185]))
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+612))
										v104 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v102)+612)) = v104
										v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+608)))
										*(*uint8)(unsafe.Add(mBase, uint32(v102)+608)) = uint8(v104)
										F_LWLockRelease(m, v102+int32(584))
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return
										} else {
											if v106 != 0 {
												m.G0 = v92 + int32(16)
												return
											} else {
												if v103 == int32(0) {
													m.G0 = v92 + int32(16)
													return
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v92)+8)) = int64(73746443898191872)
													*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v103
													v119 = *(*int32)(unsafe.Add(mBase, _consts[743]))
													*(*int32)(unsafe.Add(mBase, uint32(v92))) = v119
													v123 = *(*int32)(unsafe.Add(mBase, _consts[185]))
													F_LockRefindAndRelease(m, int32(1593656), v123, v92, int32(7), int32(0))
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return
													} else {
														m.G0 = v92 + int32(16)
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
