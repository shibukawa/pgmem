package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreatePortal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int64
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l0 == int32(0) {
		v50 = *(*int32)(unsafe.Add(mBase, _consts[885]))
		v52 = F_MemoryContextAllocZero(m, v50, int32(144))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, _consts[885]))
			v60 = F_AllocSetContextCreateInternal(m, v55, int32(62203), int32(0), int32(1024), int32(8192))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v60
				v64 = *(*int32)(unsafe.Add(mBase, _consts[417]))
				v66 = F_ResourceOwnerCreate(m, v64, int32(309695))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v52)+80)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v66
					*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = int32(1832)
					v74 = *(*int32)(unsafe.Add(mBase, _consts[65]))
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v75
					*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v75
					v79 = *(*int32)(unsafe.Add(mBase, _consts[65]))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
					v81 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v52)+136)) = uint8(v81)
					v83 = int32(257)
					*(*uint16)(unsafe.Add(mBase, uint32(v52)+116)) = uint16(v83)
					*(*int64)(unsafe.Add(mBase, uint32(v52)+72)) = int64(17179869188)
					*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v80
					v89 = *(*int64)(unsafe.Add(mBase, _consts[664]))
					*(*int64)(unsafe.Add(mBase, uint32(v52)+128)) = v89
					v92 = *(*int32)(unsafe.Add(mBase, _consts[1241]))
					v96 = F_hash_search(m, v92, l0, v81, v8+int32(31))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+31)))
						if v98 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(380976), int32(0))
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(496545), int32(222), int32(309649))
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v96)+64)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v52))) = v96
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
							v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
							if v105 != 0 {
								v106 = v96
							} else {
								v106 = int32(545505)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v103)+36)) = v106
							m.G0 = v8 + int32(32)
							return v52
						}
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[1241]))
		v14 = int32(0)
		v16 = F_hash_search(m, v13, l0, v14, v14)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v16 == int32(0) {
				v50 = *(*int32)(unsafe.Add(mBase, _consts[885]))
				v52 = F_MemoryContextAllocZero(m, v50, int32(144))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, _consts[885]))
					v60 = F_AllocSetContextCreateInternal(m, v55, int32(62203), int32(0), int32(1024), int32(8192))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v60
						v64 = *(*int32)(unsafe.Add(mBase, _consts[417]))
						v66 = F_ResourceOwnerCreate(m, v64, int32(309695))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v52)+80)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v66
							*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = int32(1832)
							v74 = *(*int32)(unsafe.Add(mBase, _consts[65]))
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v75
							*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v75
							v79 = *(*int32)(unsafe.Add(mBase, _consts[65]))
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
							v81 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v52)+136)) = uint8(v81)
							v83 = int32(257)
							*(*uint16)(unsafe.Add(mBase, uint32(v52)+116)) = uint16(v83)
							*(*int64)(unsafe.Add(mBase, uint32(v52)+72)) = int64(17179869188)
							*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v80
							v89 = *(*int64)(unsafe.Add(mBase, _consts[664]))
							*(*int64)(unsafe.Add(mBase, uint32(v52)+128)) = v89
							v92 = *(*int32)(unsafe.Add(mBase, _consts[1241]))
							v96 = F_hash_search(m, v92, l0, v81, v8+int32(31))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+31)))
								if v98 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(380976), int32(0))
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496545), int32(222), int32(309649))
											mBase = m.M
											v142 = m.ExcPending
											if v142 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v96)+64)) = v52
									*(*int32)(unsafe.Add(mBase, uint32(v52))) = v96
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
									v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
									if v105 != 0 {
										v106 = v96
									} else {
										v106 = int32(545505)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v103)+36)) = v106
									m.G0 = v8 + int32(32)
									return v52
								}
							}
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
				if v22 == int32(0) {
					v50 = *(*int32)(unsafe.Add(mBase, _consts[885]))
					v52 = F_MemoryContextAllocZero(m, v50, int32(144))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, _consts[885]))
						v60 = F_AllocSetContextCreateInternal(m, v55, int32(62203), int32(0), int32(1024), int32(8192))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v60
							v64 = *(*int32)(unsafe.Add(mBase, _consts[417]))
							v66 = F_ResourceOwnerCreate(m, v64, int32(309695))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v52)+80)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v66
								*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = int32(1832)
								v74 = *(*int32)(unsafe.Add(mBase, _consts[65]))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v75
								*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v75
								v79 = *(*int32)(unsafe.Add(mBase, _consts[65]))
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
								v81 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v52)+136)) = uint8(v81)
								v83 = int32(257)
								*(*uint16)(unsafe.Add(mBase, uint32(v52)+116)) = uint16(v83)
								*(*int64)(unsafe.Add(mBase, uint32(v52)+72)) = int64(17179869188)
								*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v80
								v89 = *(*int64)(unsafe.Add(mBase, _consts[664]))
								*(*int64)(unsafe.Add(mBase, uint32(v52)+128)) = v89
								v92 = *(*int32)(unsafe.Add(mBase, _consts[1241]))
								v96 = F_hash_search(m, v92, l0, v81, v8+int32(31))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+31)))
									if v98 == int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v133 = m.ExcPending
										if v133 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(380976), int32(0))
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496545), int32(222), int32(309649))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v96)+64)) = v52
										*(*int32)(unsafe.Add(mBase, uint32(v52))) = v96
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
										v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
										if v105 != 0 {
											v106 = v96
										} else {
											v106 = int32(545505)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v103)+36)) = v106
										m.G0 = v8 + int32(32)
										return v52
									}
								}
							}
						}
					}
				} else {
					if l1 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50462852))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
								F_errmsg(m, int32(116134), v8+int32(16))
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(496545), int32(187), int32(309649))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if l2 != 0 {
							F_PortalDrop(m, v22, int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, _consts[885]))
								v52 = F_MemoryContextAllocZero(m, v50, int32(144))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, _consts[885]))
									v60 = F_AllocSetContextCreateInternal(m, v55, int32(62203), int32(0), int32(1024), int32(8192))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v60
										v64 = *(*int32)(unsafe.Add(mBase, _consts[417]))
										v66 = F_ResourceOwnerCreate(m, v64, int32(309695))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v52)+80)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v66
											*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = int32(1832)
											v74 = *(*int32)(unsafe.Add(mBase, _consts[65]))
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v75
											*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v75
											v79 = *(*int32)(unsafe.Add(mBase, _consts[65]))
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
											v81 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v52)+136)) = uint8(v81)
											v83 = int32(257)
											*(*uint16)(unsafe.Add(mBase, uint32(v52)+116)) = uint16(v83)
											*(*int64)(unsafe.Add(mBase, uint32(v52)+72)) = int64(17179869188)
											*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v80
											v89 = *(*int64)(unsafe.Add(mBase, _consts[664]))
											*(*int64)(unsafe.Add(mBase, uint32(v52)+128)) = v89
											v92 = *(*int32)(unsafe.Add(mBase, _consts[1241]))
											v96 = F_hash_search(m, v92, l0, v81, v8+int32(31))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int32(0)
											} else {
												v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+31)))
												if v98 == int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(380976), int32(0))
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(496545), int32(222), int32(309649))
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v96)+64)) = v52
													*(*int32)(unsafe.Add(mBase, uint32(v52))) = v96
													v103 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
													v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
													if v105 != 0 {
														v106 = v96
													} else {
														v106 = int32(545505)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v103)+36)) = v106
													m.G0 = v8 + int32(32)
													return v52
												}
											}
										}
									}
								}
							}
						} else {
							v29 = F_errstart(m, int32(19), int32(0))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								if v29 == int32(0) {
									F_PortalDrop(m, v22, int32(0))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, _consts[885]))
										v52 = F_MemoryContextAllocZero(m, v50, int32(144))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, _consts[885]))
											v60 = F_AllocSetContextCreateInternal(m, v55, int32(62203), int32(0), int32(1024), int32(8192))
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v60
												v64 = *(*int32)(unsafe.Add(mBase, _consts[417]))
												v66 = F_ResourceOwnerCreate(m, v64, int32(309695))
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v52)+80)) = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v66
													*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = int32(1832)
													v74 = *(*int32)(unsafe.Add(mBase, _consts[65]))
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v75
													*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v75
													v79 = *(*int32)(unsafe.Add(mBase, _consts[65]))
													v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
													v81 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v52)+136)) = uint8(v81)
													v83 = int32(257)
													*(*uint16)(unsafe.Add(mBase, uint32(v52)+116)) = uint16(v83)
													*(*int64)(unsafe.Add(mBase, uint32(v52)+72)) = int64(17179869188)
													*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v80
													v89 = *(*int64)(unsafe.Add(mBase, _consts[664]))
													*(*int64)(unsafe.Add(mBase, uint32(v52)+128)) = v89
													v92 = *(*int32)(unsafe.Add(mBase, _consts[1241]))
													v96 = F_hash_search(m, v92, l0, v81, v8+int32(31))
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return int32(0)
													} else {
														v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+31)))
														if v98 == int32(1) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(380976), int32(0))
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(496545), int32(222), int32(309649))
																	mBase = m.M
																	v142 = m.ExcPending
																	if v142 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v96)+64)) = v52
															*(*int32)(unsafe.Add(mBase, uint32(v52))) = v96
															v103 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
															v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
															if v105 != 0 {
																v106 = v96
															} else {
																v106 = int32(545505)
															}
															*(*int32)(unsafe.Add(mBase, uint32(v103)+36)) = v106
															m.G0 = v8 + int32(32)
															return v52
														}
													}
												}
											}
										}
									}
								} else {
									F_errcode(m, int32(50462852))
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
										F_errmsg(m, int32(684298), v8)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496545), int32(192), int32(309649))
											mBase = m.M
											v44 = m.ExcPending
											if v44 != 0 {
												return int32(0)
											} else {
												F_PortalDrop(m, v22, int32(0))
												mBase = m.M
												v47 = m.ExcPending
												if v47 != 0 {
													return int32(0)
												} else {
													v50 = *(*int32)(unsafe.Add(mBase, _consts[885]))
													v52 = F_MemoryContextAllocZero(m, v50, int32(144))
													mBase = m.M
													v53 = m.ExcPending
													if v53 != 0 {
														return int32(0)
													} else {
														v55 = *(*int32)(unsafe.Add(mBase, _consts[885]))
														v60 = F_AllocSetContextCreateInternal(m, v55, int32(62203), int32(0), int32(1024), int32(8192))
														mBase = m.M
														v61 = m.ExcPending
														if v61 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v60
															v64 = *(*int32)(unsafe.Add(mBase, _consts[417]))
															v66 = F_ResourceOwnerCreate(m, v64, int32(309695))
															mBase = m.M
															v67 = m.ExcPending
															if v67 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v52)+80)) = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v66
																*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = int32(1832)
																v74 = *(*int32)(unsafe.Add(mBase, _consts[65]))
																v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v75
																*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v75
																v79 = *(*int32)(unsafe.Add(mBase, _consts[65]))
																v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
																v81 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v52)+136)) = uint8(v81)
																v83 = int32(257)
																*(*uint16)(unsafe.Add(mBase, uint32(v52)+116)) = uint16(v83)
																*(*int64)(unsafe.Add(mBase, uint32(v52)+72)) = int64(17179869188)
																*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v80
																v89 = *(*int64)(unsafe.Add(mBase, _consts[664]))
																*(*int64)(unsafe.Add(mBase, uint32(v52)+128)) = v89
																v92 = *(*int32)(unsafe.Add(mBase, _consts[1241]))
																v96 = F_hash_search(m, v92, l0, v81, v8+int32(31))
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return int32(0)
																} else {
																	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+31)))
																	if v98 == int32(1) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v133 = m.ExcPending
																		if v133 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg_internal(m, int32(380976), int32(0))
																			mBase = m.M
																			v137 = m.ExcPending
																			if v137 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(496545), int32(222), int32(309649))
																				mBase = m.M
																				v142 = m.ExcPending
																				if v142 != 0 {
																					return int32(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			}
																		}
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v96)+64)) = v52
																		*(*int32)(unsafe.Add(mBase, uint32(v52))) = v96
																		v103 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
																		v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
																		if v105 != 0 {
																			v106 = v96
																		} else {
																			v106 = int32(545505)
																		}
																		*(*int32)(unsafe.Add(mBase, uint32(v103)+36)) = v106
																		m.G0 = v8 + int32(32)
																		return v52
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
}
func F_DoPortalRewind(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v4 == int32(1) {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
		if v7 != int32(1) {
			return
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+76)))
			if v10&int32(4) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_errmsg(m, int32(421277), int32(0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							F_errhint(m, int32(604586), int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								F_errfinish(m, int32(491777), int32(1685), int32(424772))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
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
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if v13 != 0 {
					v14 = int32(4489440)
					v15 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
					F_tuplestore_rescan(m, v13)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v15
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						if v24 != 0 {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
							F_PushActiveSnapshot(m, v25)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								F_ExecutorRewind(m, v24)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return
								} else {
									F_PopActiveSnapshot(m)
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
										v34 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v34)
										return
									}
								}
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
							v34 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v34)
							return
						}
					}
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					if v24 != 0 {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
						F_PushActiveSnapshot(m, v25)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							F_ExecutorRewind(m, v24)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								F_PopActiveSnapshot(m)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
									v34 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v34)
									return
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
						v34 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v34)
						return
					}
				}
			}
		}
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+76)))
		if v10&int32(4) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errmsg(m, int32(421277), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errhint(m, int32(604586), int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							F_errfinish(m, int32(491777), int32(1685), int32(424772))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
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
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v13 != 0 {
				v14 = int32(4489440)
				v15 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
				F_tuplestore_rescan(m, v13)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v15
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					if v24 != 0 {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
						F_PushActiveSnapshot(m, v25)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							F_ExecutorRewind(m, v24)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								F_PopActiveSnapshot(m)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
									v34 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v34)
									return
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
						v34 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v34)
						return
					}
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
				if v24 != 0 {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
					F_PushActiveSnapshot(m, v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_ExecutorRewind(m, v24)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							F_PopActiveSnapshot(m)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
								v34 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v34)
								return
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
					v34 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v34)
					return
				}
			}
		}
	}
}
func F_EnsurePortalSnapshotExists(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	v3 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	if base.B2i32(v3 != int32(0)) == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[883]))
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(309572), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errfinish(m, int32(491777), int32(1778), int32(117330))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v12 = F_GetTransactionSnapshot(m)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
				F_PushActiveSnapshotWithLevel(m, v12, v14)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, _consts[135]))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+100)) = v19
					return
				}
			}
		}
	} else {
		return
	}
}
func F_FillPortalStore(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int64
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = v8 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[885]))
	v22 = F_AllocSetContextCreateInternal(m, v17, int32(62454), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v22
		v25 = int32(4489440)
		v26 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v32 = int32(1)
		v36 = *(*int32)(unsafe.Add(mBase, _consts[131]))
		v37 = F_tuplestore_begin_heap(m, int32(base.Ui32(v29&int32(2))>>(uint(v32)%32)), v32, v36)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v37
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v26
			v43 = F_CreateDestReceiver(m, int32(6))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
				v47 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v43)+36)) = v47
				*(*int32)(unsafe.Add(mBase, uint32(v43)+32)) = v47
				*(*uint8)(unsafe.Add(mBase, uint32(v43)+28)) = uint8(v47)
				*(*int32)(unsafe.Add(mBase, uint32(v43)+24)) = v46
				*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v45
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				if base.Ui32(int32(2)) <= base.Ui32(v55-int32(1)) {
					if v55 == int32(3) {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
						F_PortalRunUtility(m, l0, v64, l1, int32(1), v43, v8+int32(16))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
							if v91 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v91
								v93 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v93
							} else {
							}
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
							m.T0[v95].(func(*base.Module, int32))(m, v43)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return
							} else {
								m.G0 = v8 + int32(32)
								return
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v74
							F_errmsg_internal(m, int32(480018), v8)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								F_errfinish(m, int32(491777), int32(1032), int32(363935))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v86 = *(*int32)(unsafe.Add(mBase, _consts[884]))
					F_PortalRunMulti(m, l0, l1, int32(1), v43, v86, v8+int32(16))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
						if v91 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v91
							v93 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v93
						} else {
						}
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
						m.T0[v95].(func(*base.Module, int32))(m, v43)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			}
		}
	}
}
func F_PortalCleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v4 == int32(0) {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = int32(0)
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
		if v9 == int32(5) {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[262]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v14 != 0 {
				*(*int32)(unsafe.Add(mBase, _consts[262])) = v14
			} else {
			}
			F_ExecutorFinish(m, v4)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_ExecutorEnd(m, v4)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_FreeQueryDesc(m, v4)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[262])) = v13
						return
					}
				}
			}
		}
	}
}
func F_PortalGetPrimaryStmt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v5 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v36
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v6 <= int32(0) {
		v36 = int32(0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v36 = int32(0)
	goto L1
L5:
	;
	v9 = int32(0)
	if v9 < v6 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v12 = v6
	goto L8
L7:
	;
	v12 = v9
	goto L8
L8:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v15 = int32(0)
	goto L9
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13+v15<<(uint(int32(2))%32))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+26)))
	if v23 == int32(1) {
		v36 = v22
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L4
L11:
	;
	v27 = v15 + int32(1)
	if v27 != v12 {
		v15 = v27
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
func F_PortalHashTableDeleteAll(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[1241]))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v5 + int32(32)
	return
L2:
	;
	F_hash_seq_init(m, v5+int32(12), v8)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v17 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v17 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v22 = v17
	goto L7
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	if v24 != int32(3) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L1
L9:
	;
	F_PortalDrop(m, v23, int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v42 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L15
	}
L12:
	;
	F_hash_seq_term(m, v5+int32(12))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[1241]))
	F_hash_seq_init(m, v5+int32(12), v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	if v42 != 0 {
		v22 = v42
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L8
}
func F_PortalSetResultFormat(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v14 == v4 {
		m.G0 = v12 + int32(16)
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v20 = v18 << (uint(int32(1)) % 32)
		v21 = F_MemoryContextAlloc(m, v17, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v21
			if int32(2) <= l1 {
				if l1 != v18 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v189 = m.ExcPending
					if v189 != 0 {
						return
					} else {
						F_errcode(m, int32(16908800))
						mBase = m.M
						v192 = m.ExcPending
						if v192 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v18
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
							F_errmsg(m, int32(148551), v12)
							mBase = m.M
							v197 = m.ExcPending
							if v197 != 0 {
								return
							} else {
								F_errfinish(m, int32(491777), int32(643), int32(112171))
								mBase = m.M
								v202 = m.ExcPending
								if v202 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if v20 != 0 {
						v27 = F__emscripten_memcpy_bulkmem(m, v21, l2, v20)
						mBase = m.M
					} else {
					}
					m.G0 = v12 + int32(16)
					return
				}
			} else {
				if l1 != int32(1) {
					if v18 <= int32(0) {
					} else {
						v34 = v18 & int32(3)
						v35 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v18) {
							v43 = v35
							v47 = int32(0)
							for {
								v52 = v43 << (uint(int32(1)) % 32)
								v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v55 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v52+v53))) = uint16(v55)
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								*(*uint16)(unsafe.Add(mBase, uint32(v57+v52)+2)) = uint16(v55)
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								*(*uint16)(unsafe.Add(mBase, uint32(v61+v52)+4)) = uint16(v55)
								v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								*(*uint16)(unsafe.Add(mBase, uint32(v65+v52)+6)) = uint16(v55)
								v69 = int32(4)
								v70 = v43 + v69
								v72 = v47 + v69
								if v72 != v18&int32(2147483644) {
									v43 = v70
									v47 = v72
									continue
								} else {
									break
								}
								break
							}
							v75 = v70
						} else {
							v75 = v35
						}
						if v34 == int32(0) {
						} else {
							v86 = v75
							v87 = v35
							for {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v95 = int32(1)
								v98 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v94+v86<<(uint(v95)%32)))) = uint16(v98)
								v103 = v87 + v95
								if v103 != v34 {
									v86 = v86 + v95
									v87 = v103
									continue
								} else {
									break
								}
								break
							}
						}
					}
				} else {
					if v18 <= int32(0) {
					} else {
						v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
						v109 = v18 & int32(3)
						v110 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v18) {
							v117 = v110
							v122 = v4
							for {
								v126 = v117 << (uint(int32(1)) % 32)
								v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								*(*uint16)(unsafe.Add(mBase, uint32(v126+v127))) = uint16(v107)
								v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								*(*uint16)(unsafe.Add(mBase, uint32(v130+v126)+2)) = uint16(v107)
								v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								*(*uint16)(unsafe.Add(mBase, uint32(v133+v126)+4)) = uint16(v107)
								v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								*(*uint16)(unsafe.Add(mBase, uint32(v136+v126)+6)) = uint16(v107)
								v139 = int32(4)
								v140 = v117 + v139
								v142 = v122 + v139
								if v142 != v18&int32(2147483644) {
									v117 = v140
									v122 = v142
									continue
								} else {
									break
								}
								break
							}
							v145 = v140
						} else {
							v145 = v110
						}
						if v109 == int32(0) {
						} else {
							v156 = v145
							v160 = v110
							for {
								v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v165 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v164+v156<<(uint(v165)%32)))) = uint16(v107)
								v172 = v160 + v165
								if v172 != v109 {
									v156 = v156 + v165
									v160 = v172
									continue
								} else {
									break
								}
								break
							}
						}
					}
				}
				m.G0 = v12 + int32(16)
				return
			}
		}
	}
}
func F_PortalStart(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int64
	_ = v146
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v287 int32
	_ = v287
	var v288 int64
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	v5 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	v34 = v5
	v35 = v5
	v38 = v25
	v39 = int32(-1)
	v40 = v5
	v41 = v5
	v42 = v5
	v43 = v5
	v44 = v5
	v45 = v5
	goto L3
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v76
	*(*int32)(unsafe.Add(mBase, _consts[265])) = v77
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v94
	*(*int32)(unsafe.Add(mBase, _consts[883])) = v80
	*(*int32)(unsafe.Add(mBase, _consts[262])) = v79
	*(*int32)(unsafe.Add(mBase, _consts[403])) = v78
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(2)
	m.G0 = v25 + int32(16)
	return
L3:
	;
	if v39 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(0)
	goto L2
L5:
	;
	v55 = v38 - int32(160)
	m.G0 = v55
	v59 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	v61 = *(*int32)(unsafe.Add(mBase, _consts[265]))
	v63 = *(*int32)(unsafe.Add(mBase, _consts[403]))
	v65 = *(*int32)(unsafe.Add(mBase, _consts[262]))
	v67 = *(*int32)(unsafe.Add(mBase, _consts[883]))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v25 + int32(8)
	goto L8
L6:
	;
	v73 = v34
	v74 = v38
	v75 = v40
	v76 = v41
	v77 = v42
	v78 = v43
	v79 = v44
	v80 = v45
	goto L7
L7:
	;
	goto L10
L8:
	;
	v73 = int32(0)
	v74 = v55
	v75 = v55
	v76 = v59
	v77 = v61
	v78 = v63
	v79 = v65
	v80 = v67
	goto L7
L9:
	;
	goto L4
L10:
	;
	if v73 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v287 = int32(m.ExcTag)
	v288 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v287 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[883])) = l0
	*(*int32)(unsafe.Add(mBase, _consts[265])) = v75
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v87 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v76
	*(*int32)(unsafe.Add(mBase, _consts[265])) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v35
	F_MarkPortalFailed(m, l0)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L12
	} else {
		goto L65
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[262])) = v87
	goto L18
L17:
	;
	goto L18
L18:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _consts[403])) = v91
	v93 = int32(4489440)
	v94 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v91
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l1
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v35
	v100 = F_ChoosePortalStrategy(m, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v100
	switch v100 - int32(1) {
	case 0, 1:
		goto L21
	case 2:
		goto L20
	case 3:
		goto L9
	default:
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v35
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v220 != 0 {
		goto L53
	} else {
		goto L54
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v35
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v173 != 0 {
		goto L39
	} else {
		goto L40
	}
L22:
	;
	if l3 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v108 = v35
	v109 = l3
	goto L25
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v35
	v106 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L12
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v108
	F_PushActiveSnapshot(m, v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L27
	}
L26:
	;
	v108 = v106
	v109 = v106
	goto L25
L27:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v108
	v119 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	goto L28
L28:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v108
	v124 = *(*int32)(unsafe.Add(mBase, _consts[884]))
	v126 = F_palloc(m, int32(56))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v108
	v133 = F_RegisterSnapshot(m, v120)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+12)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v108
	v138 = F_RegisterSnapshot(m, int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	v140 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v126)+52)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v126)+28)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v126)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v126)+20)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v126)+16)) = v138
	v146 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v126)+32)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v126)+40)) = v146
	*(*uint8)(unsafe.Add(mBase, uint32(v126)+48)) = uint8(v140)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v108
	if v152&int32(2) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v156 = l2 | int32(12)
	goto L34
L33:
	;
	v156 = l2
	goto L34
L34:
	;
	F_ExecutorStart(m, v126, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v126
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v126)+36))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
	v163 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v163)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v108
	F_PopActiveSnapshot(m)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	goto L2
L37:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v35
	v209 = F_ExecCleanTypeFromTL(m, v207)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L12
	} else {
		goto L50
	}
L38:
	;
	goto L37
L39:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	if v174 <= int32(0) {
		v204 = int32(0)
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v204 = int32(0)
	goto L38
L42:
	;
	v177 = int32(0)
	if v177 < v174 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v180 = v174
	goto L45
L44:
	;
	v180 = v177
	goto L45
L45:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v183 = int32(0)
	goto L46
L46:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v181+v183<<(uint(int32(2))%32))))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+26)))
	if v191 == int32(1) {
		v204 = v190
		goto L38
	} else {
		goto L48
	}
L47:
	;
	goto L41
L48:
	;
	v195 = v183 + int32(1)
	if v195 != v180 {
		v183 = v195
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
	v213 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v213)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v209
	goto L2
L51:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v251)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v35
	v255 = F_UtilityTupleDescriptor(m, v253)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L12
	} else {
		goto L64
	}
L52:
	;
	goto L51
L53:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v221 <= int32(0) {
		v251 = int32(0)
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v251 = int32(0)
	goto L52
L56:
	;
	v224 = int32(0)
	if v224 < v221 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v227 = v221
	goto L59
L58:
	;
	v227 = v224
	goto L59
L59:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v230 = int32(0)
	goto L60
L60:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v228+v230<<(uint(int32(2))%32))))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+26)))
	if v238 == int32(1) {
		v251 = v237
		goto L52
	} else {
		goto L62
	}
L61:
	;
	goto L55
L62:
	;
	v242 = v230 + int32(1)
	if v242 != v227 {
		v230 = v242
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
	v259 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v259)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v255
	goto L2
L65:
	;
	*(*int32)(unsafe.Add(mBase, _consts[262])) = v79
	*(*int32)(unsafe.Add(mBase, _consts[883])) = v80
	*(*int32)(unsafe.Add(mBase, _consts[403])) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v35
	F_pg_re_throw(m)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	goto L1
L67:
	;
	v292 = int32(v288)
	m.G0 = v74
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v25+int32(8) == v299 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	m.ExcPending = 1
	goto L76
L69:
	;
	if v302 != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	v302 = v301
	goto L72
L71:
	;
	v302 = int32(0)
	goto L72
L72:
	;
	goto L69
L73:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v34 = v294
	v35 = v303
	v38 = v74
	v39 = v302
	v40 = v75
	v41 = v76
	v42 = v77
	v43 = v78
	v44 = v79
	v45 = v80
	goto L3
L74:
	;
	goto L75
L75:
	;
	F___wasm_longjmp(m, v295, v294)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	return
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
