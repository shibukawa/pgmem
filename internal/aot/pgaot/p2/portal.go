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
		v50 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[0]))
		v52 = F_MemoryContextAllocZero(m, v50, int32(144))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[0]))
			v60 = F_AllocSetContextCreateInternal(m, v55, int32(_a_F_CreatePortal_0), int32(0), int32(1024), int32(_a_F_CreatePortal_1))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v60
				v64 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[1]))
				v66 = F_ResourceOwnerCreate(m, v64, int32(_a_F_CreatePortal_2))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v52)+80)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v66
					*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = int32(1816)
					v74 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[2]))
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v75
					*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v75
					v79 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[2]))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
					v81 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v52)+136)) = uint8(v81)
					v83 = int32(257)
					*(*uint16)(unsafe.Add(mBase, uint32(v52)+116)) = uint16(v83)
					*(*int64)(unsafe.Add(mBase, uint32(v52)+72)) = int64(17179869188)
					*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v80
					v89 = *(*int64)(unsafe.Add(mBase, _c_F_CreatePortal[3]))
					*(*int64)(unsafe.Add(mBase, uint32(v52)+128)) = v89
					v92 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[4]))
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
								F_errmsg_internal(m, int32(_a_F_CreatePortal_3), int32(0))
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_CreatePortal_4), int32(222), int32(_a_F_CreatePortal_5))
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
								v106 = int32(_a_F_CreatePortal_6)
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
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[4]))
		v14 = int32(0)
		v16 = F_hash_search(m, v13, l0, v14, v14)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v16 == int32(0) {
				v50 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[0]))
				v52 = F_MemoryContextAllocZero(m, v50, int32(144))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[0]))
					v60 = F_AllocSetContextCreateInternal(m, v55, int32(_a_F_CreatePortal_0), int32(0), int32(1024), int32(_a_F_CreatePortal_1))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v60
						v64 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[1]))
						v66 = F_ResourceOwnerCreate(m, v64, int32(_a_F_CreatePortal_2))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v52)+80)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v66
							*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = int32(1816)
							v74 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[2]))
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v75
							*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v75
							v79 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[2]))
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
							v81 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v52)+136)) = uint8(v81)
							v83 = int32(257)
							*(*uint16)(unsafe.Add(mBase, uint32(v52)+116)) = uint16(v83)
							*(*int64)(unsafe.Add(mBase, uint32(v52)+72)) = int64(17179869188)
							*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v80
							v89 = *(*int64)(unsafe.Add(mBase, _c_F_CreatePortal[3]))
							*(*int64)(unsafe.Add(mBase, uint32(v52)+128)) = v89
							v92 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[4]))
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
										F_errmsg_internal(m, int32(_a_F_CreatePortal_3), int32(0))
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_CreatePortal_4), int32(222), int32(_a_F_CreatePortal_5))
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
										v106 = int32(_a_F_CreatePortal_6)
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
					v50 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[0]))
					v52 = F_MemoryContextAllocZero(m, v50, int32(144))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[0]))
						v60 = F_AllocSetContextCreateInternal(m, v55, int32(_a_F_CreatePortal_0), int32(0), int32(1024), int32(_a_F_CreatePortal_1))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v60
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[1]))
							v66 = F_ResourceOwnerCreate(m, v64, int32(_a_F_CreatePortal_2))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v52)+80)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v66
								*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = int32(1816)
								v74 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[2]))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v75
								*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v75
								v79 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[2]))
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
								v81 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v52)+136)) = uint8(v81)
								v83 = int32(257)
								*(*uint16)(unsafe.Add(mBase, uint32(v52)+116)) = uint16(v83)
								*(*int64)(unsafe.Add(mBase, uint32(v52)+72)) = int64(17179869188)
								*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v80
								v89 = *(*int64)(unsafe.Add(mBase, _c_F_CreatePortal[3]))
								*(*int64)(unsafe.Add(mBase, uint32(v52)+128)) = v89
								v92 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[4]))
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
											F_errmsg_internal(m, int32(_a_F_CreatePortal_3), int32(0))
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_CreatePortal_4), int32(222), int32(_a_F_CreatePortal_5))
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
											v106 = int32(_a_F_CreatePortal_6)
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
								F_errmsg(m, int32(_a_F_CreatePortal_7), v8+int32(16))
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_CreatePortal_4), int32(187), int32(_a_F_CreatePortal_5))
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
								v50 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[0]))
								v52 = F_MemoryContextAllocZero(m, v50, int32(144))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[0]))
									v60 = F_AllocSetContextCreateInternal(m, v55, int32(_a_F_CreatePortal_0), int32(0), int32(1024), int32(_a_F_CreatePortal_1))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v60
										v64 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[1]))
										v66 = F_ResourceOwnerCreate(m, v64, int32(_a_F_CreatePortal_2))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v52)+80)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v66
											*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = int32(1816)
											v74 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[2]))
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v75
											*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v75
											v79 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[2]))
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
											v81 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v52)+136)) = uint8(v81)
											v83 = int32(257)
											*(*uint16)(unsafe.Add(mBase, uint32(v52)+116)) = uint16(v83)
											*(*int64)(unsafe.Add(mBase, uint32(v52)+72)) = int64(17179869188)
											*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v80
											v89 = *(*int64)(unsafe.Add(mBase, _c_F_CreatePortal[3]))
											*(*int64)(unsafe.Add(mBase, uint32(v52)+128)) = v89
											v92 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[4]))
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
														F_errmsg_internal(m, int32(_a_F_CreatePortal_3), int32(0))
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_CreatePortal_4), int32(222), int32(_a_F_CreatePortal_5))
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
														v106 = int32(_a_F_CreatePortal_6)
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
										v50 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[0]))
										v52 = F_MemoryContextAllocZero(m, v50, int32(144))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[0]))
											v60 = F_AllocSetContextCreateInternal(m, v55, int32(_a_F_CreatePortal_0), int32(0), int32(1024), int32(_a_F_CreatePortal_1))
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v60
												v64 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[1]))
												v66 = F_ResourceOwnerCreate(m, v64, int32(_a_F_CreatePortal_2))
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v52)+80)) = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v66
													*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = int32(1816)
													v74 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[2]))
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v75
													*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v75
													v79 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[2]))
													v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
													v81 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v52)+136)) = uint8(v81)
													v83 = int32(257)
													*(*uint16)(unsafe.Add(mBase, uint32(v52)+116)) = uint16(v83)
													*(*int64)(unsafe.Add(mBase, uint32(v52)+72)) = int64(17179869188)
													*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v80
													v89 = *(*int64)(unsafe.Add(mBase, _c_F_CreatePortal[3]))
													*(*int64)(unsafe.Add(mBase, uint32(v52)+128)) = v89
													v92 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[4]))
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
																F_errmsg_internal(m, int32(_a_F_CreatePortal_3), int32(0))
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_CreatePortal_4), int32(222), int32(_a_F_CreatePortal_5))
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
																v106 = int32(_a_F_CreatePortal_6)
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
										F_errmsg(m, int32(_a_F_CreatePortal_8), v8)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_CreatePortal_4), int32(192), int32(_a_F_CreatePortal_5))
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
													v50 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[0]))
													v52 = F_MemoryContextAllocZero(m, v50, int32(144))
													mBase = m.M
													v53 = m.ExcPending
													if v53 != 0 {
														return int32(0)
													} else {
														v55 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[0]))
														v60 = F_AllocSetContextCreateInternal(m, v55, int32(_a_F_CreatePortal_0), int32(0), int32(1024), int32(_a_F_CreatePortal_1))
														mBase = m.M
														v61 = m.ExcPending
														if v61 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v60
															v64 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[1]))
															v66 = F_ResourceOwnerCreate(m, v64, int32(_a_F_CreatePortal_2))
															mBase = m.M
															v67 = m.ExcPending
															if v67 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v52)+80)) = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v66
																*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = int32(1816)
																v74 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[2]))
																v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v75
																*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v75
																v79 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[2]))
																v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
																v81 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v52)+136)) = uint8(v81)
																v83 = int32(257)
																*(*uint16)(unsafe.Add(mBase, uint32(v52)+116)) = uint16(v83)
																*(*int64)(unsafe.Add(mBase, uint32(v52)+72)) = int64(17179869188)
																*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v80
																v89 = *(*int64)(unsafe.Add(mBase, _c_F_CreatePortal[3]))
																*(*int64)(unsafe.Add(mBase, uint32(v52)+128)) = v89
																v92 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePortal[4]))
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
																			F_errmsg_internal(m, int32(_a_F_CreatePortal_3), int32(0))
																			mBase = m.M
																			v137 = m.ExcPending
																			if v137 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_CreatePortal_4), int32(222), int32(_a_F_CreatePortal_5))
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
																			v106 = int32(_a_F_CreatePortal_6)
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
						F_errmsg(m, int32(_a_F_DoPortalRewind_0), int32(0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							F_errhint(m, int32(_a_F_DoPortalRewind_1), int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_DoPortalRewind_2), int32(1685), int32(_a_F_DoPortalRewind_3))
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
					v14 = int32(_a_F_DoPortalRewind_4)
					v15 = *(*int32)(unsafe.Add(mBase, _c_F_DoPortalRewind[0]))
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
					*(*int32)(unsafe.Add(mBase, _c_F_DoPortalRewind[0])) = v17
					F_tuplestore_rescan(m, v13)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_DoPortalRewind[0])) = v15
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
					F_errmsg(m, int32(_a_F_DoPortalRewind_0), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errhint(m, int32(_a_F_DoPortalRewind_1), int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_DoPortalRewind_2), int32(1685), int32(_a_F_DoPortalRewind_3))
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
				v14 = int32(_a_F_DoPortalRewind_4)
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_DoPortalRewind[0]))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
				*(*int32)(unsafe.Add(mBase, _c_F_DoPortalRewind[0])) = v17
				F_tuplestore_rescan(m, v13)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_DoPortalRewind[0])) = v15
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
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_EnsurePortalSnapshotExists[0]))
	if base.B2i32(v3 != int32(0)) == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_EnsurePortalSnapshotExists[1]))
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_EnsurePortalSnapshotExists_0), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_EnsurePortalSnapshotExists_1), int32(1778), int32(_a_F_EnsurePortalSnapshotExists_2))
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
					v18 = *(*int32)(unsafe.Add(mBase, _c_F_EnsurePortalSnapshotExists[0]))
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = v9 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_FillPortalStore[0]))
	v23 = F_AllocSetContextCreateInternal(m, v18, int32(_a_F_FillPortalStore_0), int32(0), int32(_a_F_FillPortalStore_1), int32(_a_F_FillPortalStore_2))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v23
		v26 = int32(_a_F_FillPortalStore_3)
		v27 = *(*int32)(unsafe.Add(mBase, _c_F_FillPortalStore[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_FillPortalStore[1])) = v23
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v33 = int32(1)
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_FillPortalStore[2]))
		v38 = F_tuplestore_begin_heap(m, int32(base.Ui32(v30&int32(2))>>(uint(v33)%32)), v33, v37)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v38
			*(*int32)(unsafe.Add(mBase, _c_F_FillPortalStore[1])) = v27
			v44 = F_CreateDestReceiver(m, int32(6))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
				v48 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v44)+36)) = v48
				*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v48
				*(*uint8)(unsafe.Add(mBase, uint32(v44)+28)) = uint8(v48)
				*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v47
				*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v46
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				if base.Ui32(int32(2)) <= base.Ui32(v56-int32(1)) {
					if v56 == int32(3) {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
						F_PortalRunUtility(m, l0, v65, l1, int32(1), v44, v12)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
							if v90 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v90
								v92 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v92
							} else {
							}
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
							m.T0[v94].(func(*base.Module, int32))(m, v44)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v73
							F_errmsg_internal(m, int32(_a_F_FillPortalStore_4), v9)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_FillPortalStore_5), int32(1032), int32(_a_F_FillPortalStore_6))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
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
					v85 = *(*int32)(unsafe.Add(mBase, _c_F_FillPortalStore[3]))
					F_PortalRunMulti(m, l0, l1, int32(1), v44, v85, v9+int32(16))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
						if v90 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v90
							v92 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v92
						} else {
						}
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
						m.T0[v94].(func(*base.Module, int32))(m, v44)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return
						} else {
							m.G0 = v9 + int32(32)
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
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_PortalCleanup[0]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v14 != 0 {
				*(*int32)(unsafe.Add(mBase, _c_F_PortalCleanup[0])) = v14
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
						*(*int32)(unsafe.Add(mBase, _c_F_PortalCleanup[0])) = v13
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
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
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
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_PortalHashTableDeleteAll[0]))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v6 + int32(32)
	return
L2:
	;
	v13 = v6 + int32(12)
	F_hash_seq_init(m, v13, v9)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v16 = F_hash_seq_search(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v16 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v20 = v16
	goto L7
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+64))
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
	v41 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L15
	}
L12:
	;
	v31 = v6 + int32(12)
	F_hash_seq_term(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_PortalHashTableDeleteAll[0]))
	F_hash_seq_init(m, v31, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	if v41 != 0 {
		v20 = v41
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
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
					v190 = m.ExcPending
					if v190 != 0 {
						return
					} else {
						F_errcode(m, int32(16908800))
						mBase = m.M
						v193 = m.ExcPending
						if v193 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v18
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
							F_errmsg(m, int32(_a_F_PortalSetResultFormat_0), v12)
							mBase = m.M
							v198 = m.ExcPending
							if v198 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_PortalSetResultFormat_1), int32(643), int32(_a_F_PortalSetResultFormat_2))
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
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
					if v20 == int32(0) {
					} else {
						base.MemoryCopy(m, v21, l2, v20)
					}
					m.G0 = v12 + int32(16)
					return
				}
			} else {
				if l1 != int32(1) {
					if v18 <= int32(0) {
					} else {
						v35 = v18 & int32(3)
						v36 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v18) {
							v43 = v36
							v45 = int32(0)
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
								v72 = v45 + v69
								if v72 != v18&int32(2147483644) {
									v43 = v70
									v45 = v72
									continue
								} else {
									break
								}
								break
							}
							if v35 == int32(0) {
							} else {
								v77 = v70
								v87 = v77
								v88 = int32(0)
								for {
									v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
									v96 = int32(1)
									v99 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v95+v87<<(uint(v96)%32)))) = uint16(v99)
									v104 = v88 + v96
									if v104 != v35 {
										v87 = v87 + v96
										v88 = v104
										continue
									} else {
										break
									}
									break
								}
							}
						} else {
							v77 = v36
							v87 = v77
							v88 = int32(0)
							for {
								v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v96 = int32(1)
								v99 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v95+v87<<(uint(v96)%32)))) = uint16(v99)
								v104 = v88 + v96
								if v104 != v35 {
									v87 = v87 + v96
									v88 = v104
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
						v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
						v110 = v18 & int32(3)
						v111 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v18) {
							v118 = v111
							v124 = v4
							for {
								v127 = v118 << (uint(int32(1)) % 32)
								v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								*(*uint16)(unsafe.Add(mBase, uint32(v127+v128))) = uint16(v108)
								v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								*(*uint16)(unsafe.Add(mBase, uint32(v131+v127)+2)) = uint16(v108)
								v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								*(*uint16)(unsafe.Add(mBase, uint32(v134+v127)+4)) = uint16(v108)
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								*(*uint16)(unsafe.Add(mBase, uint32(v137+v127)+6)) = uint16(v108)
								v140 = int32(4)
								v141 = v118 + v140
								v143 = v124 + v140
								if v143 != v18&int32(2147483644) {
									v118 = v141
									v124 = v143
									continue
								} else {
									break
								}
								break
							}
							if v110 == int32(0) {
							} else {
								v148 = v141
								v157 = v148
								v158 = v111
								for {
									v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
									v166 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v165+v157<<(uint(v166)%32)))) = uint16(v108)
									v173 = v158 + v166
									if v173 != v110 {
										v157 = v157 + v166
										v158 = v173
										continue
									} else {
										break
									}
									break
								}
							}
						} else {
							v148 = v111
							v157 = v148
							v158 = v111
							for {
								v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v166 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v165+v157<<(uint(v166)%32)))) = uint16(v108)
								v173 = v158 + v166
								if v173 != v110 {
									v157 = v157 + v166
									v158 = v173
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int64
	_ = v141
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v282 int32
	_ = v282
	var v283 int64
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	v5 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(176)
	m.G0 = v23
	v32 = v5
	v34 = v5
	v35 = int32(-1)
	v36 = v5
	v37 = v5
	v38 = v5
	v39 = v5
	v40 = v5
	goto L3
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[0])) = v69
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[1])) = v70
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[2])) = v89
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[3])) = v73
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[4])) = v72
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[5])) = v71
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(2)
	m.G0 = v23 + int32(176)
	return
L3:
	;
	if v35 != int32(1) {
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
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_PortalStart[0]))
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_PortalStart[1]))
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_PortalStart[5]))
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_PortalStart[4]))
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_PortalStart[3]))
	v62 = v23 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v23 + int32(12)
	goto L8
L6:
	;
	v68 = v32
	v69 = v36
	v70 = v37
	v71 = v38
	v72 = v39
	v73 = v40
	goto L7
L7:
	;
	goto L10
L8:
	;
	v68 = int32(0)
	v69 = v52
	v70 = v54
	v71 = v56
	v72 = v58
	v73 = v60
	goto L7
L9:
	;
	goto L4
L10:
	;
	if v68 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v282 = int32(m.ExcTag)
	v283 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v282 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[3])) = l0
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[1])) = v23 + int32(16)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v82 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[0])) = v69
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[1])) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v34
	F_MarkPortalFailed(m, l0)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L12
	} else {
		goto L65
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[4])) = v82
	goto L18
L17:
	;
	goto L18
L18:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[5])) = v86
	v88 = int32(_a_F_PortalStart_0)
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_PortalStart[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[2])) = v86
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l1
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v34
	v95 = F_ChoosePortalStrategy(m, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v95
	switch v95 - int32(1) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v34
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v215 != 0 {
		goto L53
	} else {
		goto L54
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v34
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v168 != 0 {
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
	v103 = v34
	v104 = l3
	goto L25
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v34
	v101 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L12
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v103
	F_PushActiveSnapshot(m, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L12
	} else {
		goto L27
	}
L26:
	;
	v103 = v101
	v104 = v101
	goto L25
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v103
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_PortalStart[6]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	goto L28
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v103
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_PortalStart[7]))
	v121 = F_palloc(m, int32(56))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v103
	v128 = F_RegisterSnapshot(m, v115)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+12)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v103
	v133 = F_RegisterSnapshot(m, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	v135 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+52)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v121)+28)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v121)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v121)+20)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v121)+16)) = v133
	v141 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v121)+32)) = v141
	*(*int64)(unsafe.Add(mBase, uint32(v121)+40)) = v141
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+48)) = uint8(v135)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v103
	if v147&int32(2) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v151 = l2 | int32(12)
	goto L34
L33:
	;
	v151 = l2
	goto L34
L34:
	;
	F_ExecutorStart(m, v121, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v121
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
	v158 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v158)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v103
	F_PopActiveSnapshot(m)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	goto L2
L37:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v34
	v204 = F_ExecCleanTypeFromTL(m, v202)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L12
	} else {
		goto L50
	}
L38:
	;
	goto L37
L39:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if v169 <= int32(0) {
		v199 = int32(0)
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v199 = int32(0)
	goto L38
L42:
	;
	v172 = int32(0)
	if v172 < v169 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v175 = v169
	goto L45
L44:
	;
	v175 = v172
	goto L45
L45:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	v178 = int32(0)
	goto L46
L46:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v176+v178<<(uint(int32(2))%32))))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+26)))
	if v186 == int32(1) {
		v199 = v185
		goto L38
	} else {
		goto L48
	}
L47:
	;
	goto L41
L48:
	;
	v190 = v178 + int32(1)
	if v190 != v175 {
		v178 = v190
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
	v208 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v208)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v204
	goto L2
L51:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v246)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v34
	v250 = F_UtilityTupleDescriptor(m, v248)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L12
	} else {
		goto L64
	}
L52:
	;
	goto L51
L53:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v216 <= int32(0) {
		v246 = int32(0)
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v246 = int32(0)
	goto L52
L56:
	;
	v219 = int32(0)
	if v219 < v216 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v222 = v216
	goto L59
L58:
	;
	v222 = v219
	goto L59
L59:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v225 = int32(0)
	goto L60
L60:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v223+v225<<(uint(int32(2))%32))))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+26)))
	if v233 == int32(1) {
		v246 = v232
		goto L52
	} else {
		goto L62
	}
L61:
	;
	goto L55
L62:
	;
	v237 = v225 + int32(1)
	if v237 != v222 {
		v225 = v237
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
	v254 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v254)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v250
	goto L2
L65:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[4])) = v72
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[3])) = v73
	*(*int32)(unsafe.Add(mBase, _c_F_PortalStart[5])) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v34
	F_pg_re_throw(m)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	goto L1
L67:
	;
	v287 = int32(v283)
	m.G0 = v23
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	if v23+int32(12) == v293 {
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
	if v297 != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
	v297 = v295
	goto L72
L71:
	;
	v297 = int32(0)
	goto L72
L72:
	;
	goto L69
L73:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v23)+172))
	v32 = v289
	v34 = v298
	v35 = v297
	v36 = v69
	v37 = v70
	v38 = v71
	v39 = v72
	v40 = v73
	goto L3
L74:
	;
	goto L75
L75:
	;
	F___wasm_longjmp(m, v290, v289)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
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
