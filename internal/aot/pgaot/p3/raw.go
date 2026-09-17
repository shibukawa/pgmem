package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_raw_page_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = F_superuser(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v14 = F_textToQualifiedNameList(m, l0)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = F_makeRangeVarFromNameList(m, v14)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v19 = F_relation_openrv(m, v16, int32(1))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+119)))
						switch v22 - int32(83) {
						case 0, 22, 26, 31, 33:
							v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+118)))
							if v48 == int32(116) {
								v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
								if v51 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_get_raw_page_internal_0), int32(0))
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_get_raw_page_internal_1), int32(176), int32(_a_F_get_raw_page_internal_2))
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
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
									v54 = F_RelationGetNumberOfBlocksInFork(m, v19, l1)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										if base.Ui32(v54) <= base.Ui32(l2) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v141 = m.ExcPending
												if v141 != 0 {
													return int32(0)
												} else {
													v142 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v142 + int32(4)
													F_errmsg(m, int32(_a_F_get_raw_page_internal_3), v8+int32(16))
													mBase = m.M
													v151 = m.ExcPending
													if v151 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_get_raw_page_internal_1), int32(182), int32(_a_F_get_raw_page_internal_2))
														mBase = m.M
														v156 = m.ExcPending
														if v156 != 0 {
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
											v58 = F_palloc(m, int32(_a_F_get_raw_page_internal_4))
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(_a_F_get_raw_page_internal_5)
												v62 = int32(0)
												v64 = F_ReadBufferExtended(m, v19, l1, l2, v62, v62)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													F_LockBuffer(m, v64, int32(1))
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return int32(0)
													} else {
														if v64 < int32(0) {
															v74 = *(*int32)(unsafe.Add(mBase, _c_F_get_raw_page_internal[0]))
															v80 = *(*int32)(unsafe.Add(mBase, uint32(v74+(v64^int32(-1))<<(uint(int32(2))%32))))
															v88 = v80
														} else {
															v82 = *(*int32)(unsafe.Add(mBase, _c_F_get_raw_page_internal[1]))
															v88 = v82 + v64<<(uint(int32(13))%32) + int32(-8192)
														}
														base.MemoryCopy(m, v58+int32(4), v88, int32(_a_F_get_raw_page_internal_6))
														F_LockBuffer(m, v64, int32(0))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return int32(0)
														} else {
															F_ReleaseBuffer(m, v64)
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v19, int32(1))
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v8 + int32(32)
																	return v58
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
								v54 = F_RelationGetNumberOfBlocksInFork(m, v19, l1)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									if base.Ui32(v54) <= base.Ui32(l2) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v141 = m.ExcPending
											if v141 != 0 {
												return int32(0)
											} else {
												v142 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v142 + int32(4)
												F_errmsg(m, int32(_a_F_get_raw_page_internal_3), v8+int32(16))
												mBase = m.M
												v151 = m.ExcPending
												if v151 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_get_raw_page_internal_1), int32(182), int32(_a_F_get_raw_page_internal_2))
													mBase = m.M
													v156 = m.ExcPending
													if v156 != 0 {
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
										v58 = F_palloc(m, int32(_a_F_get_raw_page_internal_4))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(_a_F_get_raw_page_internal_5)
											v62 = int32(0)
											v64 = F_ReadBufferExtended(m, v19, l1, l2, v62, v62)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												F_LockBuffer(m, v64, int32(1))
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													if v64 < int32(0) {
														v74 = *(*int32)(unsafe.Add(mBase, _c_F_get_raw_page_internal[0]))
														v80 = *(*int32)(unsafe.Add(mBase, uint32(v74+(v64^int32(-1))<<(uint(int32(2))%32))))
														v88 = v80
													} else {
														v82 = *(*int32)(unsafe.Add(mBase, _c_F_get_raw_page_internal[1]))
														v88 = v82 + v64<<(uint(int32(13))%32) + int32(-8192)
													}
													base.MemoryCopy(m, v58+int32(4), v88, int32(_a_F_get_raw_page_internal_6))
													F_LockBuffer(m, v64, int32(0))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return int32(0)
													} else {
														F_ReleaseBuffer(m, v64)
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v19, int32(1))
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(32)
																return v58
															}
														}
													}
												}
											}
										}
									}
								}
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int32(0)
								} else {
									v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v32 + int32(4)
									F_errmsg(m, int32(_a_F_get_raw_page_internal_7), v8)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return int32(0)
									} else {
										v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
										v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+119)))
										F_errdetail_relkind_not_supported(m, v40)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_get_raw_page_internal_1), int32(166), int32(_a_F_get_raw_page_internal_2))
											mBase = m.M
											v47 = m.ExcPending
											if v47 != 0 {
												return int32(0)
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_get_raw_page_internal_8), int32(0))
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_get_raw_page_internal_1), int32(156), int32(_a_F_get_raw_page_internal_2))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
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
func F_raw_array_subscript_handler(m *base.Module, l0 int32) int32 {
	return int32(_a_F_raw_array_subscript_handler_0)
}
