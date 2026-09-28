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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v43 int32
	_ = v43
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v67 int64
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int64
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	v1 = l0
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[0])))
	if v10 != 0 {
		m.G0 = v7 + int32(48)
		return
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[1]))
		if base.B2i32(l1 == int32(0))&base.B2i32(base.Ui64(v1) <= base.Ui64(v14)) != 0 {
			m.G0 = v7 + int32(48)
			return
		} else {
			if v14 != int64(0) {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
				v33 = F_LWLockAcquire(m, v29+int32(1152), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[3]))
					v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+144))
					*(*int64)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[1])) = v38
					if v38 == int64(0) {
						v43 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[0])) = uint8(v43)
						v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
						F_LWLockRelease(m, v115+int32(1152))
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return
						} else {
							m.G0 = v7 + int32(48)
							return
						}
					} else {
						if base.B2i32(l1 == int32(0))&base.B2i32(base.Ui64(v1) <= base.Ui64(v38)) != 0 {
							v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
							F_LWLockRelease(m, v115+int32(1152))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						} else {
							v51 = F_GetCurrentReplayRecPtr(m, v7+int32(44))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								if l1|base.B2i32(base.Ui64(v1) <= base.Ui64(v51)) != 0 {
									v80 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[3]))
									v81 = *(*int64)(unsafe.Add(mBase, uint32(v80)+144))
									if base.Ui64(v51) <= base.Ui64(v81) {
										v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
										F_LWLockRelease(m, v115+int32(1152))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											m.G0 = v7 + int32(48)
											return
										}
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v80)+144)) = v51
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
										*(*int32)(unsafe.Add(mBase, uint32(v80)+152)) = v84
										v87 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[4]))
										F_update_controlfile(m, v87, v80)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[1])) = v51
											v94 = F_errstart(m, int32(13), int32(0))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												if v94 == int32(0) {
													v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
													F_LWLockRelease(m, v115+int32(1152))
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														m.G0 = v7 + int32(48)
														return
													}
												} else {
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
													*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v98
													*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v51)
													v102 = int64(base.Ui64(v51) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v102)
													F_errmsg_internal(m, int32(_a_F_UpdateMinRecoveryPoint_0), v7)
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_UpdateMinRecoveryPoint_1), int32(2797), int32(_a_F_UpdateMinRecoveryPoint_2))
														mBase = m.M
														v111 = m.ExcPending
														if v111 != 0 {
															return
														} else {
															v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
															F_LWLockRelease(m, v115+int32(1152))
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
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
									v57 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										if v57 == int32(0) {
											v80 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[3]))
											v81 = *(*int64)(unsafe.Add(mBase, uint32(v80)+144))
											if base.Ui64(v51) <= base.Ui64(v81) {
												v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
												F_LWLockRelease(m, v115+int32(1152))
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return
												} else {
													m.G0 = v7 + int32(48)
													return
												}
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v80)+144)) = v51
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
												*(*int32)(unsafe.Add(mBase, uint32(v80)+152)) = v84
												v87 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[4]))
												F_update_controlfile(m, v87, v80)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													*(*int64)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[1])) = v51
													v94 = F_errstart(m, int32(13), int32(0))
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														if v94 == int32(0) {
															v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
															F_LWLockRelease(m, v115+int32(1152))
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return
															} else {
																m.G0 = v7 + int32(48)
																return
															}
														} else {
															v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
															*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v98
															*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v51)
															v102 = int64(base.Ui64(v51) >> (uint(int64(32)) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v102)
															F_errmsg_internal(m, int32(_a_F_UpdateMinRecoveryPoint_0), v7)
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_UpdateMinRecoveryPoint_1), int32(2797), int32(_a_F_UpdateMinRecoveryPoint_2))
																mBase = m.M
																v111 = m.ExcPending
																if v111 != 0 {
																	return
																} else {
																	v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
																	F_LWLockRelease(m, v115+int32(1152))
																	mBase = m.M
																	v119 = m.ExcPending
																	if v119 != 0 {
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
											*(*uint32)(unsafe.Add(mBase, uint32(v7)+28)) = uint32(v51)
											v62 = int64(32)
											v63 = int64(base.Ui64(v51) >> (uint(v62) % 64))
											*(*uint32)(unsafe.Add(mBase, uint32(v7)+24)) = uint32(v63)
											*(*uint32)(unsafe.Add(mBase, uint32(v7)+20)) = uint32(v1)
											v67 = int64(base.Ui64(v1) >> (uint(v62) % 64))
											*(*uint32)(unsafe.Add(mBase, uint32(v7)+16)) = uint32(v67)
											F_errmsg_internal(m, int32(_a_F_UpdateMinRecoveryPoint_3), v7+int32(16))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_UpdateMinRecoveryPoint_1), int32(2784), int32(_a_F_UpdateMinRecoveryPoint_2))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													v80 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[3]))
													v81 = *(*int64)(unsafe.Add(mBase, uint32(v80)+144))
													if base.Ui64(v51) <= base.Ui64(v81) {
														v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
														F_LWLockRelease(m, v115+int32(1152))
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return
														} else {
															m.G0 = v7 + int32(48)
															return
														}
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v80)+144)) = v51
														v84 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
														*(*int32)(unsafe.Add(mBase, uint32(v80)+152)) = v84
														v87 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[4]))
														F_update_controlfile(m, v87, v80)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return
														} else {
															*(*int64)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[1])) = v51
															v94 = F_errstart(m, int32(13), int32(0))
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
																return
															} else {
																if v94 == int32(0) {
																	v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
																	F_LWLockRelease(m, v115+int32(1152))
																	mBase = m.M
																	v119 = m.ExcPending
																	if v119 != 0 {
																		return
																	} else {
																		m.G0 = v7 + int32(48)
																		return
																	}
																} else {
																	v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
																	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v98
																	*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v51)
																	v102 = int64(base.Ui64(v51) >> (uint(int64(32)) % 64))
																	*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v102)
																	F_errmsg_internal(m, int32(_a_F_UpdateMinRecoveryPoint_0), v7)
																	mBase = m.M
																	v106 = m.ExcPending
																	if v106 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_UpdateMinRecoveryPoint_1), int32(2797), int32(_a_F_UpdateMinRecoveryPoint_2))
																		mBase = m.M
																		v111 = m.ExcPending
																		if v111 != 0 {
																			return
																		} else {
																			v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
																			F_LWLockRelease(m, v115+int32(1152))
																			mBase = m.M
																			v119 = m.ExcPending
																			if v119 != 0 {
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
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[5])))
				if v20&int32(1) == int32(0) {
					v29 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
					v33 = F_LWLockAcquire(m, v29+int32(1152), int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[3]))
						v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+144))
						*(*int64)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[1])) = v38
						if v38 == int64(0) {
							v43 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[0])) = uint8(v43)
							v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
							F_LWLockRelease(m, v115+int32(1152))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						} else {
							if base.B2i32(l1 == int32(0))&base.B2i32(base.Ui64(v1) <= base.Ui64(v38)) != 0 {
								v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
								F_LWLockRelease(m, v115+int32(1152))
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return
								} else {
									m.G0 = v7 + int32(48)
									return
								}
							} else {
								v51 = F_GetCurrentReplayRecPtr(m, v7+int32(44))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									if l1|base.B2i32(base.Ui64(v1) <= base.Ui64(v51)) != 0 {
										v80 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[3]))
										v81 = *(*int64)(unsafe.Add(mBase, uint32(v80)+144))
										if base.Ui64(v51) <= base.Ui64(v81) {
											v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
											F_LWLockRelease(m, v115+int32(1152))
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return
											} else {
												m.G0 = v7 + int32(48)
												return
											}
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v80)+144)) = v51
											v84 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
											*(*int32)(unsafe.Add(mBase, uint32(v80)+152)) = v84
											v87 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[4]))
											F_update_controlfile(m, v87, v80)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return
											} else {
												*(*int64)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[1])) = v51
												v94 = F_errstart(m, int32(13), int32(0))
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return
												} else {
													if v94 == int32(0) {
														v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
														F_LWLockRelease(m, v115+int32(1152))
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return
														} else {
															m.G0 = v7 + int32(48)
															return
														}
													} else {
														v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
														*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v98
														*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v51)
														v102 = int64(base.Ui64(v51) >> (uint(int64(32)) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v102)
														F_errmsg_internal(m, int32(_a_F_UpdateMinRecoveryPoint_0), v7)
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_UpdateMinRecoveryPoint_1), int32(2797), int32(_a_F_UpdateMinRecoveryPoint_2))
															mBase = m.M
															v111 = m.ExcPending
															if v111 != 0 {
																return
															} else {
																v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
																F_LWLockRelease(m, v115+int32(1152))
																mBase = m.M
																v119 = m.ExcPending
																if v119 != 0 {
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
										v57 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											if v57 == int32(0) {
												v80 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[3]))
												v81 = *(*int64)(unsafe.Add(mBase, uint32(v80)+144))
												if base.Ui64(v51) <= base.Ui64(v81) {
													v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
													F_LWLockRelease(m, v115+int32(1152))
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														m.G0 = v7 + int32(48)
														return
													}
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v80)+144)) = v51
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
													*(*int32)(unsafe.Add(mBase, uint32(v80)+152)) = v84
													v87 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[4]))
													F_update_controlfile(m, v87, v80)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return
													} else {
														*(*int64)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[1])) = v51
														v94 = F_errstart(m, int32(13), int32(0))
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return
														} else {
															if v94 == int32(0) {
																v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
																F_LWLockRelease(m, v115+int32(1152))
																mBase = m.M
																v119 = m.ExcPending
																if v119 != 0 {
																	return
																} else {
																	m.G0 = v7 + int32(48)
																	return
																}
															} else {
																v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
																*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v98
																*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v51)
																v102 = int64(base.Ui64(v51) >> (uint(int64(32)) % 64))
																*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v102)
																F_errmsg_internal(m, int32(_a_F_UpdateMinRecoveryPoint_0), v7)
																mBase = m.M
																v106 = m.ExcPending
																if v106 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_UpdateMinRecoveryPoint_1), int32(2797), int32(_a_F_UpdateMinRecoveryPoint_2))
																	mBase = m.M
																	v111 = m.ExcPending
																	if v111 != 0 {
																		return
																	} else {
																		v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
																		F_LWLockRelease(m, v115+int32(1152))
																		mBase = m.M
																		v119 = m.ExcPending
																		if v119 != 0 {
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
												*(*uint32)(unsafe.Add(mBase, uint32(v7)+28)) = uint32(v51)
												v62 = int64(32)
												v63 = int64(base.Ui64(v51) >> (uint(v62) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v7)+24)) = uint32(v63)
												*(*uint32)(unsafe.Add(mBase, uint32(v7)+20)) = uint32(v1)
												v67 = int64(base.Ui64(v1) >> (uint(v62) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v7)+16)) = uint32(v67)
												F_errmsg_internal(m, int32(_a_F_UpdateMinRecoveryPoint_3), v7+int32(16))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_UpdateMinRecoveryPoint_1), int32(2784), int32(_a_F_UpdateMinRecoveryPoint_2))
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return
													} else {
														v80 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[3]))
														v81 = *(*int64)(unsafe.Add(mBase, uint32(v80)+144))
														if base.Ui64(v51) <= base.Ui64(v81) {
															v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
															F_LWLockRelease(m, v115+int32(1152))
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return
															} else {
																m.G0 = v7 + int32(48)
																return
															}
														} else {
															*(*int64)(unsafe.Add(mBase, uint32(v80)+144)) = v51
															v84 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
															*(*int32)(unsafe.Add(mBase, uint32(v80)+152)) = v84
															v87 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[4]))
															F_update_controlfile(m, v87, v80)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return
															} else {
																*(*int64)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[1])) = v51
																v94 = F_errstart(m, int32(13), int32(0))
																mBase = m.M
																v95 = m.ExcPending
																if v95 != 0 {
																	return
																} else {
																	if v94 == int32(0) {
																		v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
																		F_LWLockRelease(m, v115+int32(1152))
																		mBase = m.M
																		v119 = m.ExcPending
																		if v119 != 0 {
																			return
																		} else {
																			m.G0 = v7 + int32(48)
																			return
																		}
																	} else {
																		v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
																		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v98
																		*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v51)
																		v102 = int64(base.Ui64(v51) >> (uint(int64(32)) % 64))
																		*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v102)
																		F_errmsg_internal(m, int32(_a_F_UpdateMinRecoveryPoint_0), v7)
																		mBase = m.M
																		v106 = m.ExcPending
																		if v106 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_UpdateMinRecoveryPoint_1), int32(2797), int32(_a_F_UpdateMinRecoveryPoint_2))
																			mBase = m.M
																			v111 = m.ExcPending
																			if v111 != 0 {
																				return
																			} else {
																				v115 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[2]))
																				F_LWLockRelease(m, v115+int32(1152))
																				mBase = m.M
																				v119 = m.ExcPending
																				if v119 != 0 {
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
				} else {
					v26 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _c_F_UpdateMinRecoveryPoint[0])) = uint8(v26)
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
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	v4 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v5 - int32(241) {
	case 0, 26:
		v44 = v4
		return v44
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
		v44 = int32(0)
		return v44
	case 12:
		v28 = int32(0)
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v31 = F_FetchPreparedStatement(m, v29, v28)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			if v31 == int32(0) {
				v44 = v28
				return v44
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
				return base.B2i32(v36 != int32(0))
			}
		}
	default:
		switch v5 - int32(203) {
		case 0:
			v15 = int32(0)
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
			if v16 != 0 {
				v44 = v15
				return v44
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v18 = F_GetPortalByName(m, v17)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					if v18 == int32(0) {
						v44 = v15
						return v44
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+92))
						return base.B2i32(v24 != int32(0))
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
			if v5 == int32(159) {
				v44 = v4
			} else {
				v44 = int32(0)
			}
			return v44
		}
	}
}
func F_uhc_to_utf8(m *base.Module, l0 int32) int64 {
	var v3 int32
	_ = v3
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14231(m, l0, int32(38), v3, v3, v3, int32(_a_F_uhc_to_utf8_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_unlink_external_pid_file(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_unlink_external_pid_file[0]))
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
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
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
	v159 = l3 << (uint(int32(2)) % 32)
	v160 = l1 + v159
	v161 = l0 + v159
	v173 = l8
	v175 = int32(0)
	goto L39
L11:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v134)+16))
	if v136 < l4 {
		goto L36
	} else {
		goto L37
	}
L12:
	;
	if v128 == int32(0) {
		goto L5
	} else {
		goto L35
	}
L13:
	;
	v85 = l2 + v25<<(uint(int32(2))%32) - int32(156)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v86 != 0 {
		goto L26
	} else {
		goto L27
	}
L14:
	;
	v70 = int32(255)
	if v67&v70 != v25&v70 {
		v79 = v68
		v80 = v69
		goto L13
	} else {
		goto L24
	}
L15:
	;
	v62 = v22 + int32(8)
	if int32(0) < l8 {
		v79 = v62
		v80 = int32(0)
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
	v155 = int32(1)
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
	v68 = v22 + int32(12)
	v69 = int32(1)
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
	v68 = v22 + int32(8)
	v69 = int32(0)
	goto L14
L23:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
	v67 = v65
	v68 = v62
	v69 = int32(0)
	goto L14
L24:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+22)))
	if v75 == int32(0) {
		v128 = v69
		goto L12
	} else {
		goto L25
	}
L25:
	;
	v79 = v68
	v80 = v69
	goto L13
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if int32(6) <= v87 {
		v128 = v80
		goto L12
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v91 = F_palloc(m, int32(84))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v134 = v86
	goto L11
L30:
	;
	return
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v91
	base.MemoryCopy(m, v91, int32(_a_F_update_node_0), int32(84))
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+8)) = uint16(v97)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v101 + int32(1)
	v106 = v91 + int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v101+v106))) = uint8(v25)
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+20)) = uint8(v25)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v112 < int32(6) {
		v134 = v91
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v116 = F_cstring_to_text_with_len(m, v106, int32(6))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_update_node[0]))
	v123 = F_accumArrayResult(m, l9, base.I64_extend_i32_u(v116), int32(0), int32(25), v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v128 = v80
	goto L12
L35:
	;
	v155 = int32(1)
	goto L10
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+16)) = l4
	v139 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = v139
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134)+23)))
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+23)) = uint16(v139)
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+21)) = uint16(v141)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v134)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+32)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v134)+28)) = v145
	goto L38
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v134
	v155 = v80 + int32(1)
	goto L10
L39:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(8)+v175<<(uint(int32(2))%32))))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+32)) = v190 | l6
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+23)))
	if v193 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L5
L41:
	;
	v200 = v173 + int32(1)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l7+v200))))
	if v202 != 0 {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v189)+23)) = uint8(v25)
	goto L41
L43:
	;
	goto L44
L44:
	;
	if v25&int32(255) == v193 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v189)+24)) = uint8(v25)
	goto L41
L46:
	;
	v220 = v175 + int32(1)
	if v220 != v155 {
		v173 = v200
		v175 = v220
		goto L39
	} else {
		goto L56
	}
L47:
	;
	F_update_node(m, l0, l1, v189, l3, l4, l5, l6, l7, v200, l9)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L30
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	if v205 != 0 {
		goto L46
	} else {
		goto L51
	}
L50:
	;
	goto L46
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+12)) = int32(1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v208 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v189+v159)+76)) = int32(0)
	goto L46
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161))) = v189
	goto L52
L54:
	;
	goto L55
L55:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	*(*int32)(unsafe.Add(mBase, uint32(v212+v159)+76)) = v189
	goto L52
L56:
	;
	goto L40
}
func F_upper(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
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
	var v52 int32
	_ = v52
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		if v10 == int32(1) {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
			if v16 == int32(18) {
				v19 = int32(16)
			} else {
				v19 = int32(0)
			}
			if base.Ui32((v16-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v26 = int32(4)
			} else {
				v26 = v19
			}
			v39 = v26
		} else {
			v27 = int32(1)
			if v10&v27 != 0 {
				v39 = int32(base.Ui32(v10)>>(uint(v27)%32)) - v27
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				v39 = int32(base.Ui32(v33)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v40 = int32(1)
		if v10&v40 != 0 {
			v44 = v40
		} else {
			v44 = int32(4)
		}
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v47 = F_str_toupper(m, v6+v44, v39, v46)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int64(0)
		} else {
			v49 = F_cstring_to_text(m, v47)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v47)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v49)
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	v4 = int32(0)
	if l2&int32(3) != 0 {
		v302 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v302
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	if base.B2i32(base.Ui32(int32(6)) < base.Ui32(v12))|base.B2i32(v12 == int32(2)) != 0 {
		v302 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v18 != 0 {
		v302 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v19 - int32(285) {
	case 0:
		goto L6
	default:
		goto L5
	case 7:
		v302 = v4
		goto L1
	}
L5:
	;
	v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+88)))
	if int32(0) < v26 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v23 == int32(0) {
		v302 = v4
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v57 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v33 = v26
	goto L10
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v33-v26)<<(uint(int32(2))%32))))
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	return int32(0)
L12:
	;
	if v33 != 0 {
		v33 = v33 + int32(1)
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L11
L15:
	;
	goto L8
L16:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v208 != int32(346) {
		goto L53
	} else {
		goto L54
	}
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v60 <= int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v68 = int32(0)
	goto L19
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v68<<(uint(int32(2))%32))))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v77 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L16
L21:
	;
	v197 = v68 + int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v197 < v198 {
		v68 = v197
		goto L19
	} else {
		goto L52
	}
L22:
	;
	if v133 == int32(0) {
		goto L21
	} else {
		goto L36
	}
L23:
	;
	v133 = int32(0)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v86 = int32(1)
	if v78 == int32(0) {
		v123 = v86
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v133 = v123
	goto L22
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v90 < v89 {
		v123 = v86
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v92 = int32(1)
	if v89 <= v92 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v95 = v92
	goto L31
L30:
	;
	v95 = v89
	goto L31
L31:
	;
	v96 = int32(8)
	v101 = int32(0)
	goto L32
L32:
	;
	v108 = v101 << (uint(int32(2)) % 32)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v77+v96+v108)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v78+v96+v108)))
	v115 = v110 & (v112 ^ int32(-1))
	v117 = base.B2i32(v115 != int32(0))
	if v115 != 0 {
		v123 = v117
		goto L26
	} else {
		goto L34
	}
L33:
	;
	v123 = v117
	goto L26
L34:
	;
	v119 = v101 + int32(1)
	if v119 != v95 {
		v101 = v119
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v138 = int32(0)
	if v136 == v138 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v191 == int32(0) {
		goto L21
	} else {
		goto L51
	}
L38:
	;
	v191 = int32(1)
	goto L37
L39:
	;
	goto L40
L40:
	;
	if v137 == int32(0) {
		v184 = v138
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v191 = v184
	goto L37
L42:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v148 < v147 {
		v184 = v138
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v150 = int32(1)
	if v147 <= v150 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v153 = v150
	goto L46
L45:
	;
	v153 = v147
	goto L46
L46:
	;
	v154 = int32(8)
	v159 = int32(0)
	goto L47
L47:
	;
	v166 = v159 << (uint(int32(2)) % 32)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v136+v154+v166)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v137+v154+v166)))
	v173 = v168 & (v170 ^ int32(-1))
	v175 = base.B2i32(v173 == int32(0))
	if v173 != 0 {
		v184 = v175
		goto L41
	} else {
		goto L49
	}
L48:
	;
	v184 = v175
	goto L41
L49:
	;
	v177 = v159 + int32(1)
	if v177 != v153 {
		v159 = v177
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	return int32(0)
L52:
	;
	goto L20
L53:
	;
	v241 = int32(1)
	if l2&int32(4) == int32(0) {
		v302 = v241
		goto L1
	} else {
		goto L60
	}
L54:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+36))
	if v212 <= int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211)+76))
	v216 = int32(0)
	v218 = v216
	goto L56
L56:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218+v215))))
	if v227 == int32(0) {
		v302 = v216
		goto L1
	} else {
		goto L58
	}
L57:
	;
	goto L53
L58:
	;
	v231 = v218 + int32(1)
	if v231 != v212 {
		v218 = v231
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+8))
	if v247 == int32(0) {
		v302 = v241
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	if v250 == int32(0) {
		v302 = v241
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	if v253 <= int32(0) {
		v302 = v241
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v256 = int32(0)
	v260 = v256
	v261 = v256
	goto L64
L64:
	;
	v267 = v260 << (uint(int32(2)) % 32)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+8))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267+v269)))
	if v271 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v302 = v293
	goto L1
L66:
	;
	v272 = int32(0)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v250)+12))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v273+v267)))
	if v275 == v272 {
		v302 = v272
		goto L1
	} else {
		goto L69
	}
L67:
	;
	v290 = v261
	goto L68
L68:
	;
	v293 = int32(1)
	v295 = v260 + v293
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	if v295 < v296 {
		v260 = v295
		v261 = v290
		goto L64
	} else {
		goto L75
	}
L69:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	if v278 != int32(6) {
		v302 = v272
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v281 = int32(*(*int16)(unsafe.Add(mBase, uint32(v275)+8)))
	v283 = v281 + int32(7)
	v284 = F_bms_is_member(m, v283, v261)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	return int32(0)
L72:
	;
	if v284 != 0 {
		v302 = v272
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v288 = F_bms_add_member(m, v261, v283)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v290 = v288
	goto L68
L75:
	;
	goto L65
}
func F_uuidv7(m *base.Module, l0 int32) int64 {
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
	v11 = int32(_a_F_uuidv7_0)
	v13 = *(*int64)(unsafe.Add(mBase, _c_F_uuidv7[0]))
	v15 = v13 + int64(245)
	v16 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7)+8)))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v20 = v16 + v17*int64(1000000000)
	if v20 < v15 {
		v22 = v15
	} else {
		v22 = v20
	}
	*(*int64)(unsafe.Add(mBase, _c_F_uuidv7[0])) = v22
	v24 = int64(1000000)
	v25 = base.I64_div_s(v22, v24)
	v30 = F_generate_uuidv7(m, v25, base.I32_wrap_i64(v22-v25*v24))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int64(0)
	} else {
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v30)
	}
}
func F_uuidv7_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v49 int64
	_ = v49
	var v54 int64
	_ = v54
	var v57 int32
	_ = v57
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F___clock_gettime(m, int32(0), v10)
	mBase = m.M
	v15 = int32(_a_F_uuidv7_interval_0)
	v17 = *(*int64)(unsafe.Add(mBase, _c_F_uuidv7_interval[0]))
	v19 = v17 + int64(245)
	v20 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10)+8)))
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v24 = v20 + v21*int64(1000000000)
	if v24 < v19 {
		v26 = v19
	} else {
		v26 = v24
	}
	*(*int64)(unsafe.Add(mBase, _c_F_uuidv7_interval[0])) = v26
	v28 = base.I32_wrap_i64(v12)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v29 != int32(2147483647) {
		if v29 != int32(-2147483648) {
			v49 = base.I64_div_s(v26, int64(1000))
			v54 = F_DirectFunctionCall2Coll(m, int32(1745), int32(0), v49-int64(946684800000000), v12&int64(4294967295))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int64(0)
			} else {
				if base.Ui64(int64(-281474976710655002)) < base.Ui64(v54-int64(280528291910655001)) {
					v63 = v54 + int64(946684800000000)
					v64 = int64(1000)
					v65 = base.I64_div_u_s(v63, v64)
					v76 = F_generate_uuidv7(m, v65, base.I32_wrap_i64(v26-v49*v64+(v63-v65*v64)*v64))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int64(0)
					} else {
						m.G0 = v10 + int32(16)
						return base.I64_extend_i32_u(v76)
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_uuidv7_interval_1), int32(0))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int64(0)
							} else {
								v96 = F_errdetail(m, int32(_a_F_uuidv7_interval_2), int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_uuidv7_interval_3), int32(728), int32(_a_F_uuidv7_interval_4))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int64(0)
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
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
			if v34 != int32(-2147483648) {
				v49 = base.I64_div_s(v26, int64(1000))
				v54 = F_DirectFunctionCall2Coll(m, int32(1745), int32(0), v49-int64(946684800000000), v12&int64(4294967295))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int64(0)
				} else {
					if base.Ui64(int64(-281474976710655002)) < base.Ui64(v54-int64(280528291910655001)) {
						v63 = v54 + int64(946684800000000)
						v64 = int64(1000)
						v65 = base.I64_div_u_s(v63, v64)
						v76 = F_generate_uuidv7(m, v65, base.I32_wrap_i64(v26-v49*v64+(v63-v65*v64)*v64))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int64(0)
						} else {
							m.G0 = v10 + int32(16)
							return base.I64_extend_i32_u(v76)
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_uuidv7_interval_1), int32(0))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int64(0)
								} else {
									v96 = F_errdetail(m, int32(_a_F_uuidv7_interval_2), int32(0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_uuidv7_interval_3), int32(728), int32(_a_F_uuidv7_interval_4))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int64(0)
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
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
				if v37 != int64(-9223372036854775807-1) {
					v49 = base.I64_div_s(v26, int64(1000))
					v54 = F_DirectFunctionCall2Coll(m, int32(1745), int32(0), v49-int64(946684800000000), v12&int64(4294967295))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int64(0)
					} else {
						if base.Ui64(int64(-281474976710655002)) < base.Ui64(v54-int64(280528291910655001)) {
							v63 = v54 + int64(946684800000000)
							v64 = int64(1000)
							v65 = base.I64_div_u_s(v63, v64)
							v76 = F_generate_uuidv7(m, v65, base.I32_wrap_i64(v26-v49*v64+(v63-v65*v64)*v64))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int64(0)
							} else {
								m.G0 = v10 + int32(16)
								return base.I64_extend_i32_u(v76)
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_uuidv7_interval_1), int32(0))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int64(0)
									} else {
										v96 = F_errdetail(m, int32(_a_F_uuidv7_interval_2), int32(0))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_uuidv7_interval_3), int32(728), int32(_a_F_uuidv7_interval_4))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return int64(0)
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
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_uuidv7_interval_5), int32(0))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int64(0)
							} else {
								v116 = F_errdetail(m, int32(_a_F_uuidv7_interval_6), int32(0))
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_uuidv7_interval_3), int32(703), int32(_a_F_uuidv7_interval_4))
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int64(0)
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
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
		if v40 != int32(2147483647) {
			v49 = base.I64_div_s(v26, int64(1000))
			v54 = F_DirectFunctionCall2Coll(m, int32(1745), int32(0), v49-int64(946684800000000), v12&int64(4294967295))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int64(0)
			} else {
				if base.Ui64(int64(-281474976710655002)) < base.Ui64(v54-int64(280528291910655001)) {
					v63 = v54 + int64(946684800000000)
					v64 = int64(1000)
					v65 = base.I64_div_u_s(v63, v64)
					v76 = F_generate_uuidv7(m, v65, base.I32_wrap_i64(v26-v49*v64+(v63-v65*v64)*v64))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int64(0)
					} else {
						m.G0 = v10 + int32(16)
						return base.I64_extend_i32_u(v76)
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_uuidv7_interval_1), int32(0))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int64(0)
							} else {
								v96 = F_errdetail(m, int32(_a_F_uuidv7_interval_2), int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_uuidv7_interval_3), int32(728), int32(_a_F_uuidv7_interval_4))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int64(0)
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
			v43 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
			if v43 == int64(9223372036854775807) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_uuidv7_interval_5), int32(0))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int64(0)
						} else {
							v116 = F_errdetail(m, int32(_a_F_uuidv7_interval_6), int32(0))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_uuidv7_interval_3), int32(703), int32(_a_F_uuidv7_interval_4))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				v49 = base.I64_div_s(v26, int64(1000))
				v54 = F_DirectFunctionCall2Coll(m, int32(1745), int32(0), v49-int64(946684800000000), v12&int64(4294967295))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int64(0)
				} else {
					if base.Ui64(int64(-281474976710655002)) < base.Ui64(v54-int64(280528291910655001)) {
						v63 = v54 + int64(946684800000000)
						v64 = int64(1000)
						v65 = base.I64_div_u_s(v63, v64)
						v76 = F_generate_uuidv7(m, v65, base.I32_wrap_i64(v26-v49*v64+(v63-v65*v64)*v64))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int64(0)
						} else {
							m.G0 = v10 + int32(16)
							return base.I64_extend_i32_u(v76)
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_uuidv7_interval_1), int32(0))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int64(0)
								} else {
									v96 = F_errdetail(m, int32(_a_F_uuidv7_interval_2), int32(0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_uuidv7_interval_3), int32(728), int32(_a_F_uuidv7_interval_4))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int64(0)
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
