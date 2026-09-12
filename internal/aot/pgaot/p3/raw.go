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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
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
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+118)))
							if v52 == int32(116) {
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
								if v55 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v134 = m.ExcPending
										if v134 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(152008), int32(0))
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(522165), int32(176), int32(326227))
												mBase = m.M
												v147 = m.ExcPending
												if v147 != 0 {
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
									v58 = F_RelationGetNumberOfBlocksInFork(m, v19, l1)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										if base.Ui32(v58) <= base.Ui32(l2) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v151 = m.ExcPending
											if v151 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
													return int32(0)
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v155 + int32(4)
													F_errmsg(m, int32(736489), v8+int32(16))
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(522165), int32(182), int32(326227))
														mBase = m.M
														v173 = m.ExcPending
														if v173 != 0 {
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
											v62 = F_palloc(m, int32(8196))
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v62))) = int32(32784)
												v66 = int32(0)
												v68 = F_ReadBufferExtended(m, v19, l1, l2, v66, v66)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return int32(0)
												} else {
													F_LockBuffer(m, v68, int32(1))
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return int32(0)
													} else {
														if v68 < int32(0) {
															v78 = *(*int32)(unsafe.Add(mBase, _consts[9]))
															v84 = *(*int32)(unsafe.Add(mBase, uint32(v78+(v68^int32(-1))<<(uint(int32(2))%32))))
															v92 = v84
														} else {
															v86 = *(*int32)(unsafe.Add(mBase, _consts[10]))
															v92 = v86 + v68<<(uint(int32(13))%32) + int32(-8192)
														}
														v94 = F__emscripten_memcpy_bulkmem(m, v62+int32(4), v92, int32(8192))
														mBase = m.M
														F_LockBuffer(m, v68, int32(0))
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return int32(0)
														} else {
															F_ReleaseBuffer(m, v68)
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v19, int32(1))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v8 + int32(32)
																	return v62
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
								v58 = F_RelationGetNumberOfBlocksInFork(m, v19, l1)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									if base.Ui32(v58) <= base.Ui32(l2) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v154 = m.ExcPending
											if v154 != 0 {
												return int32(0)
											} else {
												v155 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v155 + int32(4)
												F_errmsg(m, int32(736489), v8+int32(16))
												mBase = m.M
												v166 = m.ExcPending
												if v166 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(522165), int32(182), int32(326227))
													mBase = m.M
													v173 = m.ExcPending
													if v173 != 0 {
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
										v62 = F_palloc(m, int32(8196))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v62))) = int32(32784)
											v66 = int32(0)
											v68 = F_ReadBufferExtended(m, v19, l1, l2, v66, v66)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												F_LockBuffer(m, v68, int32(1))
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return int32(0)
												} else {
													if v68 < int32(0) {
														v78 = *(*int32)(unsafe.Add(mBase, _consts[9]))
														v84 = *(*int32)(unsafe.Add(mBase, uint32(v78+(v68^int32(-1))<<(uint(int32(2))%32))))
														v92 = v84
													} else {
														v86 = *(*int32)(unsafe.Add(mBase, _consts[10]))
														v92 = v86 + v68<<(uint(int32(13))%32) + int32(-8192)
													}
													v94 = F__emscripten_memcpy_bulkmem(m, v62+int32(4), v92, int32(8192))
													mBase = m.M
													F_LockBuffer(m, v68, int32(0))
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return int32(0)
													} else {
														F_ReleaseBuffer(m, v68)
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v19, int32(1))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(32)
																return v62
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
									F_errmsg(m, int32(738551), v8)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
										v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(v41)+119)))
										F_errdetail_relkind_not_supported(m, v42)
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(522165), int32(166), int32(326227))
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
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
			v111 = m.ExcPending
			if v111 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v114 = m.ExcPending
				if v114 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(149184), int32(0))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(522165), int32(156), int32(326227))
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
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
	return int32(1690676)
}
