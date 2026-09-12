package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_UpdateMinRecoveryPoint(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v41 int32
	_ = v41
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v64 int64
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	v1 = l0
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[200])))
	if v10 != 0 {
		m.G0 = v7 + int32(48)
		return
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, _consts[201]))
		if base.B2i32(l1 == int32(0))&base.B2i32(base.Ui64(v1) <= base.Ui64(v14)) != 0 {
			m.G0 = v7 + int32(48)
			return
		} else {
			if v14 != int64(0) {
				v27 = *(*int32)(unsafe.Add(mBase, _consts[2]))
				v31 = F_LWLockAcquire(m, v27+int32(1152), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, _consts[202]))
					v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+136))
					*(*int64)(unsafe.Add(mBase, _consts[201])) = v36
					if v36 == int64(0) {
						v41 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _consts[200])) = uint8(v41)
						v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
						F_LWLockRelease(m, v112+int32(1152))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							m.G0 = v7 + int32(48)
							return
						}
					} else {
						if base.B2i32(l1 == int32(0))&base.B2i32(base.Ui64(v1) <= base.Ui64(v36)) != 0 {
							v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
							F_LWLockRelease(m, v112+int32(1152))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						} else {
							v49 = F_GetCurrentReplayRecPtr(m, v7+int32(44))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								if l1 != 0 {
									v77 = *(*int32)(unsafe.Add(mBase, _consts[202]))
									v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+136))
									if base.Ui64(v49) <= base.Ui64(v78) {
										v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
										F_LWLockRelease(m, v112+int32(1152))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return
										} else {
											m.G0 = v7 + int32(48)
											return
										}
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v77)+136)) = v49
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
										*(*int32)(unsafe.Add(mBase, uint32(v77)+144)) = v81
										v84 = *(*int32)(unsafe.Add(mBase, _consts[203]))
										F_update_controlfile(m, v84, v77)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, _consts[201])) = v49
											v91 = F_errstart(m, int32(13), int32(0))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												if v91 == int32(0) {
													v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
													F_LWLockRelease(m, v112+int32(1152))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														m.G0 = v7 + int32(48)
														return
													}
												} else {
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
													*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v95
													*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v49)
													v99 = int64(base.Ui64(v49) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v99)
													F_errmsg_internal(m, int32(55437), v7)
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return
													} else {
														F_errfinish(m, int32(518546), int32(2769), int32(94845))
														mBase = m.M
														v108 = m.ExcPending
														if v108 != 0 {
															return
														} else {
															v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
															F_LWLockRelease(m, v112+int32(1152))
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return
															} else {
																m.G0 = v7 + int32(48)
																return
															}
														}
													}
												}
											}
										}
									}
								} else {
									if base.Ui64(v1) <= base.Ui64(v49) {
										v77 = *(*int32)(unsafe.Add(mBase, _consts[202]))
										v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+136))
										if base.Ui64(v49) <= base.Ui64(v78) {
											v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
											F_LWLockRelease(m, v112+int32(1152))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return
											} else {
												m.G0 = v7 + int32(48)
												return
											}
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v77)+136)) = v49
											v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
											*(*int32)(unsafe.Add(mBase, uint32(v77)+144)) = v81
											v84 = *(*int32)(unsafe.Add(mBase, _consts[203]))
											F_update_controlfile(m, v84, v77)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												*(*int64)(unsafe.Add(mBase, _consts[201])) = v49
												v91 = F_errstart(m, int32(13), int32(0))
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return
												} else {
													if v91 == int32(0) {
														v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
														F_LWLockRelease(m, v112+int32(1152))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															m.G0 = v7 + int32(48)
															return
														}
													} else {
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
														*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v95
														*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v49)
														v99 = int64(base.Ui64(v49) >> (uint(int64(32)) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v99)
														F_errmsg_internal(m, int32(55437), v7)
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return
														} else {
															F_errfinish(m, int32(518546), int32(2769), int32(94845))
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
																return
															} else {
																v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
																F_LWLockRelease(m, v112+int32(1152))
																mBase = m.M
																v116 = m.ExcPending
																if v116 != 0 {
																	return
																} else {
																	m.G0 = v7 + int32(48)
																	return
																}
															}
														}
													}
												}
											}
										}
									} else {
										v54 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return
										} else {
											if v54 == int32(0) {
												v77 = *(*int32)(unsafe.Add(mBase, _consts[202]))
												v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+136))
												if base.Ui64(v49) <= base.Ui64(v78) {
													v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
													F_LWLockRelease(m, v112+int32(1152))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														m.G0 = v7 + int32(48)
														return
													}
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v77)+136)) = v49
													v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
													*(*int32)(unsafe.Add(mBase, uint32(v77)+144)) = v81
													v84 = *(*int32)(unsafe.Add(mBase, _consts[203]))
													F_update_controlfile(m, v84, v77)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														*(*int64)(unsafe.Add(mBase, _consts[201])) = v49
														v91 = F_errstart(m, int32(13), int32(0))
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return
														} else {
															if v91 == int32(0) {
																v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
																F_LWLockRelease(m, v112+int32(1152))
																mBase = m.M
																v116 = m.ExcPending
																if v116 != 0 {
																	return
																} else {
																	m.G0 = v7 + int32(48)
																	return
																}
															} else {
																v95 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
																*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v95
																*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v49)
																v99 = int64(base.Ui64(v49) >> (uint(int64(32)) % 64))
																*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v99)
																F_errmsg_internal(m, int32(55437), v7)
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(518546), int32(2769), int32(94845))
																	mBase = m.M
																	v108 = m.ExcPending
																	if v108 != 0 {
																		return
																	} else {
																		v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
																		F_LWLockRelease(m, v112+int32(1152))
																		mBase = m.M
																		v116 = m.ExcPending
																		if v116 != 0 {
																			return
																		} else {
																			m.G0 = v7 + int32(48)
																			return
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												*(*uint32)(unsafe.Add(mBase, uint32(v7)+28)) = uint32(v49)
												v59 = int64(32)
												v60 = int64(base.Ui64(v49) >> (uint(v59) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v7)+24)) = uint32(v60)
												*(*uint32)(unsafe.Add(mBase, uint32(v7)+20)) = uint32(v1)
												v64 = int64(base.Ui64(v1) >> (uint(v59) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v7)+16)) = uint32(v64)
												F_errmsg_internal(m, int32(532522), v7+int32(16))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return
												} else {
													F_errfinish(m, int32(518546), int32(2755), int32(94845))
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return
													} else {
														v77 = *(*int32)(unsafe.Add(mBase, _consts[202]))
														v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+136))
														if base.Ui64(v49) <= base.Ui64(v78) {
															v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
															F_LWLockRelease(m, v112+int32(1152))
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return
															} else {
																m.G0 = v7 + int32(48)
																return
															}
														} else {
															*(*int64)(unsafe.Add(mBase, uint32(v77)+136)) = v49
															v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
															*(*int32)(unsafe.Add(mBase, uint32(v77)+144)) = v81
															v84 = *(*int32)(unsafe.Add(mBase, _consts[203]))
															F_update_controlfile(m, v84, v77)
															mBase = m.M
															v86 = m.ExcPending
															if v86 != 0 {
																return
															} else {
																*(*int64)(unsafe.Add(mBase, _consts[201])) = v49
																v91 = F_errstart(m, int32(13), int32(0))
																mBase = m.M
																v92 = m.ExcPending
																if v92 != 0 {
																	return
																} else {
																	if v91 == int32(0) {
																		v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
																		F_LWLockRelease(m, v112+int32(1152))
																		mBase = m.M
																		v116 = m.ExcPending
																		if v116 != 0 {
																			return
																		} else {
																			m.G0 = v7 + int32(48)
																			return
																		}
																	} else {
																		v95 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
																		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v95
																		*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v49)
																		v99 = int64(base.Ui64(v49) >> (uint(int64(32)) % 64))
																		*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v99)
																		F_errmsg_internal(m, int32(55437), v7)
																		mBase = m.M
																		v103 = m.ExcPending
																		if v103 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(518546), int32(2769), int32(94845))
																			mBase = m.M
																			v108 = m.ExcPending
																			if v108 != 0 {
																				return
																			} else {
																				v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
																				F_LWLockRelease(m, v112+int32(1152))
																				mBase = m.M
																				v116 = m.ExcPending
																				if v116 != 0 {
																					return
																				} else {
																					m.G0 = v7 + int32(48)
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
								}
							}
						}
					}
				}
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, _consts[161])))
				if v20 != int32(1) {
					v27 = *(*int32)(unsafe.Add(mBase, _consts[2]))
					v31 = F_LWLockAcquire(m, v27+int32(1152), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, _consts[202]))
						v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+136))
						*(*int64)(unsafe.Add(mBase, _consts[201])) = v36
						if v36 == int64(0) {
							v41 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[200])) = uint8(v41)
							v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
							F_LWLockRelease(m, v112+int32(1152))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						} else {
							if base.B2i32(l1 == int32(0))&base.B2i32(base.Ui64(v1) <= base.Ui64(v36)) != 0 {
								v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
								F_LWLockRelease(m, v112+int32(1152))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return
								} else {
									m.G0 = v7 + int32(48)
									return
								}
							} else {
								v49 = F_GetCurrentReplayRecPtr(m, v7+int32(44))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									if l1 != 0 {
										v77 = *(*int32)(unsafe.Add(mBase, _consts[202]))
										v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+136))
										if base.Ui64(v49) <= base.Ui64(v78) {
											v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
											F_LWLockRelease(m, v112+int32(1152))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return
											} else {
												m.G0 = v7 + int32(48)
												return
											}
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v77)+136)) = v49
											v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
											*(*int32)(unsafe.Add(mBase, uint32(v77)+144)) = v81
											v84 = *(*int32)(unsafe.Add(mBase, _consts[203]))
											F_update_controlfile(m, v84, v77)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												*(*int64)(unsafe.Add(mBase, _consts[201])) = v49
												v91 = F_errstart(m, int32(13), int32(0))
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return
												} else {
													if v91 == int32(0) {
														v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
														F_LWLockRelease(m, v112+int32(1152))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															m.G0 = v7 + int32(48)
															return
														}
													} else {
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
														*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v95
														*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v49)
														v99 = int64(base.Ui64(v49) >> (uint(int64(32)) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v99)
														F_errmsg_internal(m, int32(55437), v7)
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return
														} else {
															F_errfinish(m, int32(518546), int32(2769), int32(94845))
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
																return
															} else {
																v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
																F_LWLockRelease(m, v112+int32(1152))
																mBase = m.M
																v116 = m.ExcPending
																if v116 != 0 {
																	return
																} else {
																	m.G0 = v7 + int32(48)
																	return
																}
															}
														}
													}
												}
											}
										}
									} else {
										if base.Ui64(v1) <= base.Ui64(v49) {
											v77 = *(*int32)(unsafe.Add(mBase, _consts[202]))
											v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+136))
											if base.Ui64(v49) <= base.Ui64(v78) {
												v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
												F_LWLockRelease(m, v112+int32(1152))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													m.G0 = v7 + int32(48)
													return
												}
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v77)+136)) = v49
												v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
												*(*int32)(unsafe.Add(mBase, uint32(v77)+144)) = v81
												v84 = *(*int32)(unsafe.Add(mBase, _consts[203]))
												F_update_controlfile(m, v84, v77)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													*(*int64)(unsafe.Add(mBase, _consts[201])) = v49
													v91 = F_errstart(m, int32(13), int32(0))
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return
													} else {
														if v91 == int32(0) {
															v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
															F_LWLockRelease(m, v112+int32(1152))
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return
															} else {
																m.G0 = v7 + int32(48)
																return
															}
														} else {
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
															*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v95
															*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v49)
															v99 = int64(base.Ui64(v49) >> (uint(int64(32)) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v99)
															F_errmsg_internal(m, int32(55437), v7)
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return
															} else {
																F_errfinish(m, int32(518546), int32(2769), int32(94845))
																mBase = m.M
																v108 = m.ExcPending
																if v108 != 0 {
																	return
																} else {
																	v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
																	F_LWLockRelease(m, v112+int32(1152))
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		m.G0 = v7 + int32(48)
																		return
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v54 = F_errstart(m, int32(19), int32(0))
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return
											} else {
												if v54 == int32(0) {
													v77 = *(*int32)(unsafe.Add(mBase, _consts[202]))
													v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+136))
													if base.Ui64(v49) <= base.Ui64(v78) {
														v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
														F_LWLockRelease(m, v112+int32(1152))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															m.G0 = v7 + int32(48)
															return
														}
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v77)+136)) = v49
														v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
														*(*int32)(unsafe.Add(mBase, uint32(v77)+144)) = v81
														v84 = *(*int32)(unsafe.Add(mBase, _consts[203]))
														F_update_controlfile(m, v84, v77)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return
														} else {
															*(*int64)(unsafe.Add(mBase, _consts[201])) = v49
															v91 = F_errstart(m, int32(13), int32(0))
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return
															} else {
																if v91 == int32(0) {
																	v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
																	F_LWLockRelease(m, v112+int32(1152))
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		m.G0 = v7 + int32(48)
																		return
																	}
																} else {
																	v95 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
																	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v95
																	*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v49)
																	v99 = int64(base.Ui64(v49) >> (uint(int64(32)) % 64))
																	*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v99)
																	F_errmsg_internal(m, int32(55437), v7)
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(518546), int32(2769), int32(94845))
																		mBase = m.M
																		v108 = m.ExcPending
																		if v108 != 0 {
																			return
																		} else {
																			v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
																			F_LWLockRelease(m, v112+int32(1152))
																			mBase = m.M
																			v116 = m.ExcPending
																			if v116 != 0 {
																				return
																			} else {
																				m.G0 = v7 + int32(48)
																				return
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													*(*uint32)(unsafe.Add(mBase, uint32(v7)+28)) = uint32(v49)
													v59 = int64(32)
													v60 = int64(base.Ui64(v49) >> (uint(v59) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v7)+24)) = uint32(v60)
													*(*uint32)(unsafe.Add(mBase, uint32(v7)+20)) = uint32(v1)
													v64 = int64(base.Ui64(v1) >> (uint(v59) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v7)+16)) = uint32(v64)
													F_errmsg_internal(m, int32(532522), v7+int32(16))
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return
													} else {
														F_errfinish(m, int32(518546), int32(2755), int32(94845))
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return
														} else {
															v77 = *(*int32)(unsafe.Add(mBase, _consts[202]))
															v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+136))
															if base.Ui64(v49) <= base.Ui64(v78) {
																v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
																F_LWLockRelease(m, v112+int32(1152))
																mBase = m.M
																v116 = m.ExcPending
																if v116 != 0 {
																	return
																} else {
																	m.G0 = v7 + int32(48)
																	return
																}
															} else {
																*(*int64)(unsafe.Add(mBase, uint32(v77)+136)) = v49
																v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
																*(*int32)(unsafe.Add(mBase, uint32(v77)+144)) = v81
																v84 = *(*int32)(unsafe.Add(mBase, _consts[203]))
																F_update_controlfile(m, v84, v77)
																mBase = m.M
																v86 = m.ExcPending
																if v86 != 0 {
																	return
																} else {
																	*(*int64)(unsafe.Add(mBase, _consts[201])) = v49
																	v91 = F_errstart(m, int32(13), int32(0))
																	mBase = m.M
																	v92 = m.ExcPending
																	if v92 != 0 {
																		return
																	} else {
																		if v91 == int32(0) {
																			v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
																			F_LWLockRelease(m, v112+int32(1152))
																			mBase = m.M
																			v116 = m.ExcPending
																			if v116 != 0 {
																				return
																			} else {
																				m.G0 = v7 + int32(48)
																				return
																			}
																		} else {
																			v95 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
																			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v95
																			*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v49)
																			v99 = int64(base.Ui64(v49) >> (uint(int64(32)) % 64))
																			*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v99)
																			F_errmsg_internal(m, int32(55437), v7)
																			mBase = m.M
																			v103 = m.ExcPending
																			if v103 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(518546), int32(2769), int32(94845))
																				mBase = m.M
																				v108 = m.ExcPending
																				if v108 != 0 {
																					return
																				} else {
																					v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
																					F_LWLockRelease(m, v112+int32(1152))
																					mBase = m.M
																					v116 = m.ExcPending
																					if v116 != 0 {
																						return
																					} else {
																						m.G0 = v7 + int32(48)
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
									}
								}
							}
						}
					}
				} else {
					v24 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _consts[200])) = uint8(v24)
					m.G0 = v7 + int32(48)
					return
				}
			}
		}
	}
}
func F_UtilityReturnsTuples(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v4 - int32(203) {
	case 0:
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if v15 != 0 {
			v44 = v2
			return v44
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v17 = F_GetPortalByName(m, v16)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					v44 = v2
					return v44
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
					return base.B2i32(v23 != int32(0))
				}
			}
		}
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		v44 = int32(0)
		return v44
	case 10:
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		return base.B2i32(v11 == int32(2249))
	default:
		v7 = int32(1)
		switch v4 - int32(241) {
		case 0:
			v44 = v7
			return v44
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
			v44 = int32(0)
			return v44
		case 12:
			v27 = int32(0)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v30 = F_FetchPreparedStatement(m, v28, v27)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				if v30 == int32(0) {
					v44 = v27
					return v44
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+64))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
					return base.B2i32(v35 != int32(0))
				}
			}
		default:
			if v4 == int32(159) {
				v44 = v7
			} else {
				v44 = int32(0)
			}
			return v44
		}
	}
}
func F___udivti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	v6 = m.G0
	v7 = int32(16)
	v8 = v6 - v7
	m.G0 = v8
	F___udivmodti4(m, v8, l1, l2, l3, int64(0))
	mBase = m.M
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v12
	m.G0 = v8 + v7
	return
}
func F_uhc_to_utf8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(38), int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_LocalToUtf(m, v6, v10, v5, int32(4425628), v18, v18, v18, int32(38), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
	}
}
func F_unlink_external_pid_file(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, _consts[575]))
	if v4 != 0 {
		v5 = F_unlink(m, v4)
		mBase = m.M
	} else {
	}
	return
}
func F_update_node(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	v11 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v25 = int32(*(*int8)(unsafe.Add(mBase, uint32(l7+l8))))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if l4 <= v26 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l5&v40 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v40 = v28
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = l4
	v30 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v30
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+23)))
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+23)) = uint16(v30)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+21)) = uint16(v32)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v36
	v40 = v36
	goto L1
L5:
	;
	m.G0 = v22 + int32(16)
	return
L6:
	;
	v43 = int32(0)
	goto L8
L7:
	;
	v43 = v40
	goto L8
L8:
	;
	if v43 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	if v25 != int32(88) {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v165 = l3 << (uint(int32(2)) % 32)
	v166 = l1 + v165
	v167 = l0 + v165
	v179 = l8
	v181 = int32(0)
	goto L43
L11:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	if v141 < l4 {
		goto L40
	} else {
		goto L41
	}
L12:
	;
	if v131 == int32(0) {
		goto L5
	} else {
		goto L39
	}
L13:
	;
	v88 = l2 + v25<<(uint(int32(2))%32) - int32(156)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L14:
	;
	v72 = int32(255)
	if v67&v72 != v25&v72 {
		v82 = v69
		v83 = v71
		goto L13
	} else {
		goto L24
	}
L15:
	;
	if int32(0) < l8 {
		v82 = v11
		v83 = v22 + int32(8)
		goto L13
	} else {
		goto L23
	}
L16:
	;
	if l8 != 0 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = l2
	v162 = int32(1)
	goto L10
L19:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
	v48 = v25 & int32(255)
	if v46 == v48 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = l2
	v67 = v46
	v69 = int32(1)
	v71 = v22 + int32(12)
	goto L14
L21:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+22)))
	if v50 == v48 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v67 = v46
	v69 = v11
	v71 = v22 + int32(8)
	goto L14
L23:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
	v67 = v64
	v69 = v11
	v71 = v22 + int32(8)
	goto L14
L24:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+22)))
	if v77 == int32(0) {
		v131 = v69
		goto L12
	} else {
		goto L25
	}
L25:
	;
	v82 = v69
	v83 = v71
	goto L13
L26:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if int32(6) <= v90 {
		v131 = v82
		goto L12
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v94 = F_palloc(m, int32(84))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v138 = v89
	goto L11
L30:
	;
	return
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v94
	goto L33
L32:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+8)) = uint16(v101)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v105 + int32(1)
	v110 = v99 + int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v105+v110))) = uint8(v25)
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+20)) = uint8(v25)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v99)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+32)) = v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v116 < int32(6) {
		v138 = v94
		goto L11
	} else {
		goto L36
	}
L33:
	;
	v99 = F__emscripten_memcpy_bulkmem(m, v94, int32(4107636), int32(84))
	mBase = m.M
	goto L35
L35:
	;
	goto L32
L36:
	;
	v121 = F_cstring_to_text_with_len(m, v110, int32(6))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v126 = F_accumArrayResult(m, l9, v121, int32(0), int32(25), v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L30
	} else {
		goto L38
	}
L38:
	;
	v131 = v82
	goto L12
L39:
	;
	v162 = int32(1)
	goto L10
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+16)) = l4
	v144 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v138)+12)) = v144
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+23)))
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+23)) = uint16(v144)
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+21)) = uint16(v146)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v138)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+32)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v138)+28)) = v150
	goto L42
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v138
	v162 = v82 + int32(1)
	goto L10
L43:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(8)+v181<<(uint(int32(2))%32))))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+32)) = v196 | l6
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+23)))
	if v199 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L5
L45:
	;
	v206 = v179 + int32(1)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l7+v206))))
	if v208 != 0 {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+23)) = uint8(v25)
	goto L45
L47:
	;
	goto L48
L48:
	;
	if v199 == v25&int32(255) {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+24)) = uint8(v25)
	goto L45
L50:
	;
	v226 = v181 + int32(1)
	if v226 != v162 {
		v179 = v206
		v181 = v226
		goto L43
	} else {
		goto L60
	}
L51:
	;
	F_update_node(m, l0, l1, v195, l3, l4, l5, l6, l7, v206, l9)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L30
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	if v211 != 0 {
		goto L50
	} else {
		goto L55
	}
L54:
	;
	goto L50
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+12)) = int32(1)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v214 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v195
	*(*int32)(unsafe.Add(mBase, uint32(v195+v165)+76)) = int32(0)
	goto L50
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v195
	goto L56
L58:
	;
	goto L59
L59:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	*(*int32)(unsafe.Add(mBase, uint32(v218+v165)+76)) = v195
	goto L56
L60:
	;
	goto L44
}
func F_upper(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v53 int32
	_ = v53
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(1)
		v12 = v7 + v11
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v17 = v15 & v11
		if v17 != 0 {
			v18 = v12
		} else {
			v18 = v7 + int32(4)
		}
		if v15 == int32(1) {
			v21 = int32(4)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v23&int32(254) == int32(2) {
				v32 = v21
			} else {
				v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
			}
			if v23 == int32(1) {
				v35 = v21
			} else {
				v35 = v32
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v17 != 0 {
				v46 = int32(base.Ui32(v15)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v48 = F_str_toupper(m, v18, v46, v47)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			v50 = F_cstring_to_text(m, v48)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v48)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					return v50
				}
			}
		}
	}
}
func F_use_physical_tlist(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	v4 = int32(0)
	if l2&int32(3) != 0 {
		v300 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v300
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	if base.Ui32(int32(6)) < base.Ui32(v12) {
		v300 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v12 == int32(2) {
		v300 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v17 != 0 {
		v300 = v4
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v18 - int32(282) {
	case 0:
		goto L7
	default:
		goto L6
	case 7:
		v300 = v4
		goto L1
	}
L6:
	;
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+80)))
	if int32(0) < v25 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v22 == int32(0) {
		v300 = v4
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v56 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	v33 = v25
	goto L11
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28+(v33-v25)<<(uint(int32(2))%32))))
	if v41 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	return int32(0)
L13:
	;
	if v33 != 0 {
		v33 = v33 + int32(1)
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	goto L9
L17:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v207 != int32(342) {
		goto L54
	} else {
		goto L55
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v59 <= int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v66 = int32(0)
	goto L20
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v66<<(uint(int32(2))%32))))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v76 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L17
L22:
	;
	v196 = v66 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v196 < v197 {
		v66 = v196
		goto L20
	} else {
		goto L53
	}
L23:
	;
	if v132 == int32(0) {
		goto L22
	} else {
		goto L37
	}
L24:
	;
	v132 = int32(0)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v85 = int32(1)
	if v77 == int32(0) {
		v123 = v85
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v132 = v123
	goto L23
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v89 < v88 {
		v123 = v85
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v91 = int32(1)
	if v88 <= v91 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v94 = v91
	goto L32
L31:
	;
	v94 = v88
	goto L32
L32:
	;
	v95 = int32(8)
	v100 = int32(0)
	goto L33
L33:
	;
	v107 = v100 << (uint(int32(2)) % 32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v76+v95+v107)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+(v77+v95))))
	v114 = v109 & (v111 ^ int32(-1))
	v116 = base.B2i32(v114 != int32(0))
	if v114 != 0 {
		v123 = v116
		goto L27
	} else {
		goto L35
	}
L34:
	;
	v123 = v116
	goto L27
L35:
	;
	v118 = v100 + int32(1)
	if v118 != v94 {
		v100 = v118
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v137 = int32(0)
	if v135 == v137 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v190 == int32(0) {
		goto L22
	} else {
		goto L52
	}
L39:
	;
	v190 = int32(1)
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v136 == int32(0) {
		v181 = v137
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v190 = v181
	goto L38
L43:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v147 < v146 {
		v181 = v137
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v149 = int32(1)
	if v146 <= v149 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v152 = v149
	goto L47
L46:
	;
	v152 = v146
	goto L47
L47:
	;
	v153 = int32(8)
	v158 = int32(0)
	goto L48
L48:
	;
	v165 = v158 << (uint(int32(2)) % 32)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v135+v153+v165)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165+(v136+v153))))
	v172 = v167 & (v169 ^ int32(-1))
	v174 = base.B2i32(v172 == int32(0))
	if v172 != 0 {
		v181 = v174
		goto L42
	} else {
		goto L50
	}
L49:
	;
	v181 = v174
	goto L42
L50:
	;
	v176 = v158 + int32(1)
	if v176 != v152 {
		v158 = v176
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	return int32(0)
L53:
	;
	goto L21
L54:
	;
	v240 = int32(1)
	if l2&int32(4) == int32(0) {
		v300 = v240
		goto L1
	} else {
		goto L61
	}
L55:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+36))
	if v211 <= int32(0) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)+76))
	v215 = int32(0)
	v222 = v215
	goto L57
L57:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214+v222))))
	if v226 == int32(0) {
		v300 = v215
		goto L1
	} else {
		goto L59
	}
L58:
	;
	goto L54
L59:
	;
	v230 = v222 + int32(1)
	if v230 != v211 {
		v222 = v230
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+8))
	if v246 == int32(0) {
		v300 = v240
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	if v249 == int32(0) {
		v300 = v240
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v252 <= int32(0) {
		v300 = v240
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v255 = int32(0)
	v259 = v255
	v261 = v255
	goto L65
L65:
	;
	v266 = v259 << (uint(int32(2)) % 32)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+8))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v266+v268)))
	if v270 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v300 = v292
	goto L1
L67:
	;
	v271 = int32(0)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v272+v266)))
	if v274 == v271 {
		v300 = v271
		goto L1
	} else {
		goto L70
	}
L68:
	;
	v291 = v261
	goto L69
L69:
	;
	v292 = int32(1)
	v294 = v259 + v292
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v294 < v295 {
		v259 = v294
		v261 = v291
		goto L65
	} else {
		goto L76
	}
L70:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	if v277 != int32(6) {
		v300 = v271
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v280 = int32(*(*int16)(unsafe.Add(mBase, uint32(v274)+8)))
	v282 = v280 + int32(7)
	v283 = F_bms_is_member(m, v282, v261)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	return int32(0)
L73:
	;
	if v283 != 0 {
		v300 = v271
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v287 = F_bms_add_member(m, v261, v282)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v291 = v287
	goto L69
L76:
	;
	goto L66
}
func F_uuidv7(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F___clock_gettime(m, int32(0), v7)
	mBase = m.M
	v11 = int32(4535768)
	v13 = *(*int64)(unsafe.Add(mBase, _consts[1156]))
	v15 = v13 + int64(245)
	v16 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7)+8)))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v20 = v16 + v17*int64(1000000000)
	if v20 < v15 {
		v22 = v15
	} else {
		v22 = v20
	}
	*(*int64)(unsafe.Add(mBase, _consts[1156])) = v22
	v24 = int64(1000000)
	v25 = base.I64_div_s(v22, v24)
	v30 = F_generate_uuidv7(m, v25, base.I32_wrap_i64(v22-v25*v24))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v30
	}
}
func F_uuidv7_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F___clock_gettime(m, int32(0), v9)
	mBase = m.M
	v14 = int32(4535768)
	v16 = *(*int64)(unsafe.Add(mBase, _consts[1156]))
	v18 = v16 + int64(245)
	v19 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+8)))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v23 = v19 + v20*int64(1000000000)
	if v23 < v18 {
		v25 = v18
	} else {
		v25 = v23
	}
	*(*int64)(unsafe.Add(mBase, _consts[1156])) = v25
	v30 = base.I64_div_s(v25, int64(1000))
	v33 = F_Int64GetDatum(m, v30-int64(946684800000000))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		return int32(0)
	} else {
		v37 = F_DirectFunctionCall2Coll(m, int32(1548), int32(0), v33, v11)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			v39 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
			v41 = v39 + int64(946684800000000)
			v42 = int64(1000)
			v43 = base.I64_div_s(v41, v42)
			v54 = F_generate_uuidv7(m, v43, base.I32_wrap_i64(v25-v30*v42+(v41-v43*v42)*v42))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				m.G0 = v9 + int32(16)
				return v54
			}
		}
	}
}
