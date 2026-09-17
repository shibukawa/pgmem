package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateQueryDesc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int64
	_ = v32
	v12 = F_palloc(m, int32(56))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v16
		v20 = F_RegisterSnapshot(m, l2)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v20
			v23 = F_RegisterSnapshot(m, l3)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l7
				*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l6
				*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l5
				*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v23
				v32 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v12)+36)) = v32
				*(*int64)(unsafe.Add(mBase, uint32(v12)+41)) = v32
				return v12
			}
		}
	}
}
func F_ExecuteQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	v7 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteQuery[0]))
	if v19 == v7 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v114 = m.ExcPending
		if v114 != 0 {
			return
		} else {
			F_errcode(m, int32(386))
			mBase = m.M
			v117 = m.ExcPending
			if v117 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
				F_errmsg(m, int32(_a_F_ExecuteQuery_0), v15)
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(454), int32(_a_F_ExecuteQuery_2))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
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
		v22 = int32(0)
		v24 = F_hash_search(m, v19, v17, v22, v22)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			if v24 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v114 = m.ExcPending
				if v114 != 0 {
					return
				} else {
					F_errcode(m, int32(386))
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
						F_errmsg(m, int32(_a_F_ExecuteQuery_0), v15)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(454), int32(_a_F_ExecuteQuery_2))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
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
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+48)))
				if v29 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_ExecuteQuery_3), int32(0))
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(170), int32(_a_F_ExecuteQuery_4))
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v32 = int32(0)
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
					if v32 < v33 {
						v36 = F_CreateExecutorState(m)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = l3
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v40 = F_EvaluateParams(m, l0, v24, v39, v36)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								v42 = v36
								v43 = v40
								v44 = F_CreateNewPortal(m)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									v46 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v44)+136)) = uint8(v46)
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
									v51 = F_MemoryContextStrdup(m, v48, v50)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return
									} else {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
										v54 = int32(0)
										v56 = F_GetCachedPlan(m, v53, v43, v54, v54)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return
										} else {
											v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
											v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
											v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
											*(*int64)(unsafe.Add(mBase, uint32(v44)+48)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(v44)+40)) = v60
											*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v51
											*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v44)+60)) = v56
											*(*int32)(unsafe.Add(mBase, uint32(v44)+56)) = v61
											*(*int32)(unsafe.Add(mBase, uint32(v44)+36)) = v60
											if l2 != 0 {
												if v61 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v143 = m.ExcPending
													if v143 != 0 {
														return
													} else {
														F_errcode(m, int32(151027844))
														mBase = m.M
														v146 = m.ExcPending
														if v146 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_ExecuteQuery_5), int32(0))
															mBase = m.M
															v150 = m.ExcPending
															if v150 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(230), int32(_a_F_ExecuteQuery_4))
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
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
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
													if v74 != int32(1) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return
														} else {
															F_errcode(m, int32(151027844))
															mBase = m.M
															v146 = m.ExcPending
															if v146 != 0 {
																return
															} else {
																F_errmsg(m, int32(_a_F_ExecuteQuery_5), int32(0))
																mBase = m.M
																v150 = m.ExcPending
																if v150 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(230), int32(_a_F_ExecuteQuery_4))
																	mBase = m.M
																	v155 = m.ExcPending
																	if v155 != 0 {
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
														v77 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
														v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
														v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
														if v79 != int32(1) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v159 = m.ExcPending
															if v159 != 0 {
																return
															} else {
																F_errcode(m, int32(151027844))
																mBase = m.M
																v162 = m.ExcPending
																if v162 != 0 {
																	return
																} else {
																	F_errmsg(m, int32(_a_F_ExecuteQuery_5), int32(0))
																	mBase = m.M
																	v166 = m.ExcPending
																	if v166 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(235), int32(_a_F_ExecuteQuery_4))
																		mBase = m.M
																		v171 = m.ExcPending
																		if v171 != 0 {
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
															v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
															v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
															if v89 != 0 {
																v90 = int32(0)
															} else {
																v90 = int32(2147483647)
															}
															v92 = v82 << (uint(int32(6)) % 32) & int32(64)
															v93 = v90
															v95 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteQuery[1]))
															v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
															F_PortalStart(m, v44, v43, v92, v96)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																v100 = F_PortalRun(m, v44, v93, int32(0), l4, l4, l5)
																mBase = m.M
																v101 = m.ExcPending
																if v101 != 0 {
																	return
																} else {
																	F_PortalDrop(m, v44, int32(0))
																	mBase = m.M
																	v104 = m.ExcPending
																	if v104 != 0 {
																		return
																	} else {
																		if v42 != 0 {
																			F_FreeExecutorState(m, v42)
																			mBase = m.M
																			v106 = m.ExcPending
																			if v106 != 0 {
																				return
																			} else {
																				m.G0 = v15 + int32(16)
																				return
																			}
																		} else {
																			m.G0 = v15 + int32(16)
																			return
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v92 = v7
												v93 = int32(2147483647)
												v95 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteQuery[1]))
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
												F_PortalStart(m, v44, v43, v92, v96)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													v100 = F_PortalRun(m, v44, v93, int32(0), l4, l4, l5)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return
													} else {
														F_PortalDrop(m, v44, int32(0))
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return
														} else {
															if v42 != 0 {
																F_FreeExecutorState(m, v42)
																mBase = m.M
																v106 = m.ExcPending
																if v106 != 0 {
																	return
																} else {
																	m.G0 = v15 + int32(16)
																	return
																}
															} else {
																m.G0 = v15 + int32(16)
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
						v42 = v32
						v43 = v7
						v44 = F_CreateNewPortal(m)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							v46 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v44)+136)) = uint8(v46)
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
							v51 = F_MemoryContextStrdup(m, v48, v50)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
								v54 = int32(0)
								v56 = F_GetCachedPlan(m, v53, v43, v54, v54)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
									*(*int64)(unsafe.Add(mBase, uint32(v44)+48)) = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(v44)+40)) = v60
									*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v51
									*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v44)+60)) = v56
									*(*int32)(unsafe.Add(mBase, uint32(v44)+56)) = v61
									*(*int32)(unsafe.Add(mBase, uint32(v44)+36)) = v60
									if l2 != 0 {
										if v61 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v143 = m.ExcPending
											if v143 != 0 {
												return
											} else {
												F_errcode(m, int32(151027844))
												mBase = m.M
												v146 = m.ExcPending
												if v146 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_ExecuteQuery_5), int32(0))
													mBase = m.M
													v150 = m.ExcPending
													if v150 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(230), int32(_a_F_ExecuteQuery_4))
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
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
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
											if v74 != int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v143 = m.ExcPending
												if v143 != 0 {
													return
												} else {
													F_errcode(m, int32(151027844))
													mBase = m.M
													v146 = m.ExcPending
													if v146 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_ExecuteQuery_5), int32(0))
														mBase = m.M
														v150 = m.ExcPending
														if v150 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(230), int32(_a_F_ExecuteQuery_4))
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
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
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
												if v79 != int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return
													} else {
														F_errcode(m, int32(151027844))
														mBase = m.M
														v162 = m.ExcPending
														if v162 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_ExecuteQuery_5), int32(0))
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(235), int32(_a_F_ExecuteQuery_4))
																mBase = m.M
																v171 = m.ExcPending
																if v171 != 0 {
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
													v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
													v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
													if v89 != 0 {
														v90 = int32(0)
													} else {
														v90 = int32(2147483647)
													}
													v92 = v82 << (uint(int32(6)) % 32) & int32(64)
													v93 = v90
													v95 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteQuery[1]))
													v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
													F_PortalStart(m, v44, v43, v92, v96)
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return
													} else {
														v100 = F_PortalRun(m, v44, v93, int32(0), l4, l4, l5)
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return
														} else {
															F_PortalDrop(m, v44, int32(0))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return
															} else {
																if v42 != 0 {
																	F_FreeExecutorState(m, v42)
																	mBase = m.M
																	v106 = m.ExcPending
																	if v106 != 0 {
																		return
																	} else {
																		m.G0 = v15 + int32(16)
																		return
																	}
																} else {
																	m.G0 = v15 + int32(16)
																	return
																}
															}
														}
													}
												}
											}
										}
									} else {
										v92 = v7
										v93 = int32(2147483647)
										v95 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteQuery[1]))
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
										F_PortalStart(m, v44, v43, v92, v96)
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return
										} else {
											v100 = F_PortalRun(m, v44, v93, int32(0), l4, l4, l5)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												F_PortalDrop(m, v44, int32(0))
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
													return
												} else {
													if v42 != 0 {
														F_FreeExecutorState(m, v42)
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return
														} else {
															m.G0 = v15 + int32(16)
															return
														}
													} else {
														m.G0 = v15 + int32(16)
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
func F_FreeQueryDesc(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_UnregisterSnapshot(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		F_UnregisterSnapshot(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_ProcessQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int64
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessQuery[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v15 = F_palloc(m, int32(56))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
		v21 = F_RegisterSnapshot(m, v13)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v21
			v24 = int32(0)
			v26 = F_RegisterSnapshot(m, v24)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v28 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v28
				*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v26
				v34 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v34
				*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v34
				*(*uint8)(unsafe.Add(mBase, uint32(v15)+48)) = uint8(v28)
				F_ExecutorStart(m, v15, v28)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_ExecutorRun(m, v15, int32(1), int64(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						if l5 != 0 {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
							v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)+112))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
							v51 = v49 - int32(1)
							if base.Ui32(v51) <= base.Ui32(int32(4)) {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v51<<(uint(int32(2))%32))+uint32(_c_F_ProcessQuery[1])))
								v57 = v56
							} else {
								v57 = v24
							}
							*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v57
						} else {
						}
						F_ExecutorFinish(m, v15)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							F_ExecutorEnd(m, v15)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
								F_UnregisterSnapshot(m, v67)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
									F_UnregisterSnapshot(m, v70)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										F_pfree(m, v15)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
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
func F_ScanQueryWalker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 == int32(22) {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			F_ScanQueryForLocks(m, v10, v11)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v17 = F_expression_tree_walker_impl(m, l0, int32(1588), l1)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		} else {
			v17 = F_expression_tree_walker_impl(m, l0, int32(1588), l1)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v17
			}
		}
	}
}
func F_getInsertSelectQuery(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 != int32(3) {
		v188 = l0
		goto L11
	} else {
		goto L12
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L58
	} else {
		goto L68
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L58
	} else {
		goto L65
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L58
	} else {
		goto L62
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L58
	} else {
		goto L59
	}
L11:
	;
	return v188
L12:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v14 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v85 == int32(0) {
		goto L10
	} else {
		goto L32
	}
L14:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v17 < int32(2) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = int32(_a_F_getInsertSelectQuery_0)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_getInsertSelectQuery[0])))
	if base.B2i32(v27 == int32(0))|base.B2i32(v27 != v30) != 0 {
		v48 = v27
		v49 = v30
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v48-v49 != 0 {
		goto L13
	} else {
		goto L23
	}
L17:
	;
	goto L16
L18:
	;
	v33 = v23
	v34 = v24
	goto L19
L19:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v38 == int32(0) {
		v48 = v38
		v49 = v37
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v48 = v38
	v49 = v37
	goto L17
L21:
	;
	v41 = int32(1)
	if v38 == v37 {
		v33 = v33 + v41
		v34 = v34 + v41
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v54 = int32(_a_F_getInsertSelectQuery_1)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_getInsertSelectQuery[1])))
	if base.B2i32(v57 == int32(0))|base.B2i32(v57 != v60) != 0 {
		v78 = v57
		v79 = v60
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v78-v79 == int32(0) {
		v188 = l0
		goto L11
	} else {
		goto L31
	}
L25:
	;
	goto L24
L26:
	;
	v63 = v53
	v64 = v54
	goto L27
L27:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v68 == int32(0) {
		v78 = v68
		v79 = v67
		goto L25
	} else {
		goto L29
	}
L28:
	;
	v78 = v68
	v79 = v67
	goto L25
L29:
	;
	v71 = int32(1)
	if v68 == v67 {
		v63 = v63 + v71
		v64 = v64 + v71
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L13
L32:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v88 != int32(1) {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v93 != int32(63) {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v96+v97<<(uint(int32(2))%32)-int32(4))))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	if v104 != int32(1) {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)+36))
	if v107 == int32(0) {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v110 != int32(67) {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v113 != int32(1) {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	if v116 == int32(0) {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if v119 < int32(2) {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v126 = int32(_a_F_getInsertSelectQuery_0)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_getInsertSelectQuery[0])))
	if base.B2i32(v129 == int32(0))|base.B2i32(v129 != v132) != 0 {
		v150 = v129
		v151 = v132
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v150-v151 != 0 {
		goto L7
	} else {
		goto L48
	}
L42:
	;
	goto L41
L43:
	;
	v135 = v125
	v136 = v126
	goto L44
L44:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	if v140 == int32(0) {
		v150 = v140
		v151 = v139
		goto L42
	} else {
		goto L46
	}
L45:
	;
	v150 = v140
	v151 = v139
	goto L42
L46:
	;
	v143 = int32(1)
	if v140 == v139 {
		v135 = v135 + v143
		v136 = v136 + v143
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v156 = int32(_a_F_getInsertSelectQuery_1)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_getInsertSelectQuery[1])))
	if base.B2i32(v159 == int32(0))|base.B2i32(v159 != v162) != 0 {
		v180 = v159
		v181 = v162
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v180-v181 != 0 {
		goto L7
	} else {
		goto L56
	}
L50:
	;
	goto L49
L51:
	;
	v165 = v155
	v166 = v156
	goto L52
L52:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
	if v170 == int32(0) {
		v180 = v170
		v181 = v169
		goto L50
	} else {
		goto L54
	}
L53:
	;
	v180 = v170
	v181 = v169
	goto L50
L54:
	;
	v173 = int32(1)
	if v170 == v169 {
		v165 = v165 + v173
		v166 = v166 + v173
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	if l1 == int32(0) {
		v188 = v107
		goto L11
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v103 + int32(36)
	v188 = v107
	goto L11
L58:
	;
	return int32(0)
L59:
	;
	F_errmsg_internal(m, int32(_a_F_getInsertSelectQuery_2), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_getInsertSelectQuery_3), int32(1118), int32(_a_F_getInsertSelectQuery_4))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errmsg_internal(m, int32(_a_F_getInsertSelectQuery_2), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_getInsertSelectQuery_3), int32(1121), int32(_a_F_getInsertSelectQuery_4))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L58
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errmsg_internal(m, int32(_a_F_getInsertSelectQuery_2), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L58
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_getInsertSelectQuery_3), int32(1127), int32(_a_F_getInsertSelectQuery_4))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L58
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errmsg_internal(m, int32(_a_F_getInsertSelectQuery_5), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L58
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_getInsertSelectQuery_3), int32(1139), int32(_a_F_getInsertSelectQuery_4))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L58
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_query_def(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v939 int32
	_ = v939
	var v951 int32
	_ = v951
	var v962 int32
	_ = v962
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v995 int32
	_ = v995
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1076 int32
	_ = v1076
	var v1081 int32
	_ = v1081
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1124 int32
	_ = v1124
	var v1130 int32
	_ = v1130
	var v1135 int32
	_ = v1135
	var v1145 int32
	_ = v1145
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1198 int32
	_ = v1198
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1306 int32
	_ = v1306
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1404 int32
	_ = v1404
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1418 int32
	_ = v1418
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1441 int32
	_ = v1441
	var v1447 int32
	_ = v1447
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1518 int32
	_ = v1518
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1538 int32
	_ = v1538
	var v1542 int32
	_ = v1542
	var v1546 int32
	_ = v1546
	var v1551 int32
	_ = v1551
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1590 int32
	_ = v1590
	var v1599 int32
	_ = v1599
	var v1604 int32
	_ = v1604
	var v1611 int32
	_ = v1611
	var v1621 int32
	_ = v1621
	var v1625 int32
	_ = v1625
	var v1627 int32
	_ = v1627
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1650 int32
	_ = v1650
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1711 int32
	_ = v1711
	var v1714 int32
	_ = v1714
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1740 int32
	_ = v1740
	var v1743 int32
	_ = v1743
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1760 int32
	_ = v1760
	var v1765 int32
	_ = v1765
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1850 int32
	_ = v1850
	var v1855 int32
	_ = v1855
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1868 int32
	_ = v1868
	var v1872 int32
	_ = v1872
	var v1887 int32
	_ = v1887
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1894 int32
	_ = v1894
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1905 int32
	_ = v1905
	var v1910 int32
	_ = v1910
	var v1914 int32
	_ = v1914
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1948 int32
	_ = v1948
	var v1950 int32
	_ = v1950
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	v5 = l4
	v9 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(304)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_get_query_def[0]))
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+45)))
	if v27 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v49 = int32(0)
	F_AcquireRewriteLocks(m, l0, v49, v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L17
	}
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v40 = F_flatten_group_exprs(m, int32(0), l0, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L15
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v37 = v34 - int32(1)
	goto L8
L10:
	;
	if v26 != 0 {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v26 == int32(0) {
		v48 = v9
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v37 = int32(-1)
	goto L8
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v48 = v33
	goto L7
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v40
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v45 = F_flatten_group_exprs(m, int32(0), l0, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v45
	v48 = v37
	goto L7
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+264)) = l1
	v56 = v18 + int32(184)
	v57 = F_list_copy(m, l2)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v59 = F_lcons(m, v56, v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v61 = int32(0)
	v65 = base.B2i32(l2 != v61) | base.B2i32(v48 != int32(1))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+296)) = uint8(v65)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+280)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v18)+272)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+268)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v18)+300)) = v61
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+298)) = uint16(v61)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+297)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+292)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v18)+288)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v18)+284)) = l5
	F_set_deparse_for_query(m, v56, l0, l2)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v82 - int32(1) {
	case 0:
		goto L34
	case 1:
		goto L33
	case 2:
		goto L32
	case 3:
		goto L31
	case 4:
		goto L30
	case 5:
		goto L29
	case 6:
		goto L27
	default:
		goto L28
	}
L21:
	;
	m.G0 = v18 + int32(304)
	return
L22:
	;
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v1781 != 0 {
		goto L461
	} else {
		goto L462
	}
L23:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v713)+12))
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v1486 = int32(2)
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1484+v1485<<(uint(v1486)%32)-int32(4))))
	v1492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+284)))
	if v1492&v1486 != 0 {
		goto L379
	} else {
		goto L380
	}
L24:
	;
	v1447 = int32(0)
	if base.B2i32(v1441 == v1447)|base.B2i32(v1433 == v1447) != 0 {
		v1470 = v1433
		v1478 = v1441
		v1481 = base.B2i32(v1433 != v1447)
		v1483 = base.B2i32(v1441 != v1447)
		goto L23
	} else {
		goto L375
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L4
	} else {
		goto L372
	}
L26:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v144)+80))
	F_get_values_def(m, v1414, v18+int32(264))
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L4
	} else {
		goto L371
	}
L27:
	;
	F_appendStringInfoString(m, l1, int32(_a_F_get_query_def_0))
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L4
	} else {
		goto L370
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L4
	} else {
		goto L367
	}
L29:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1313 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L30:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v18)+264))
	F_get_with_clause(m, l0, v18+int32(264))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L4
	} else {
		goto L230
	}
L31:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v18)+264))
	F_get_with_clause(m, l0, v18+int32(264))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L4
	} else {
		goto L211
	}
L32:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v18)+264))
	F_get_with_clause(m, l0, v18+int32(264))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L4
	} else {
		goto L192
	}
L33:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v18)+264))
	F_get_with_clause(m, l0, v18+int32(264))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L4
	} else {
		goto L171
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+272)) = l3
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v18)+264))
	v88 = v18 + int32(264)
	F_get_with_clause(m, l0, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+276)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+280)) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v95 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_get_setop_query(m, v95, l0, v88)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v18)+264))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+284)))
	if v99&int32(2) != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L22
L40:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v18)+292))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+292)) = v102 + int32(8)
	F_appendStringInfoChar(m, v98, int32(32))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v109 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+46)))
	if v269 != 0 {
		goto L86
	} else {
		goto L87
	}
L45:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v112 <= int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v18)+272))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v122 = int32(0)
	v127 = v49
	goto L47
L47:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v116+v127<<(uint(int32(2))%32))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	switch v137 {
	case 0:
		goto L50
	default:
		goto L44
	case 5:
		goto L51
	}
L48:
	;
	if v144 == int32(0) {
		goto L44
	} else {
		goto L56
	}
L49:
	;
	v146 = v127 + int32(1)
	if v112 != v146 {
		v122 = v144
		v127 = v146
		goto L47
	} else {
		goto L55
	}
L50:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+125)))
	if v143 != 0 {
		goto L44
	} else {
		goto L54
	}
L51:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+125)))
	if v138 != int32(1) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	if v122 == int32(0) {
		v144 = v136
		goto L49
	} else {
		goto L53
	}
L53:
	;
	goto L44
L54:
	;
	v144 = v122
	goto L49
L55:
	;
	goto L48
L56:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v151 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	v154 = v152
	goto L59
L58:
	;
	v154 = int32(0)
	goto L59
L59:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
	if v156 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v159 = v157
	goto L62
L61:
	;
	v159 = int32(0)
	goto L62
L62:
	;
	if v154 != v159 {
		goto L44
	} else {
		goto L63
	}
L63:
	;
	v170 = int32(0)
	goto L64
L64:
	;
	v176 = int32(0)
	if v151 == v176 {
		v186 = v176
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L44
L66:
	;
	if v156 == int32(0) {
		goto L26
	} else {
		goto L69
	}
L67:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if v180 <= v170 {
		v186 = int32(0)
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	v186 = v182 + v170<<(uint(int32(2))%32)
	goto L66
L69:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if base.B2i32(v186 == int32(0))|base.B2i32(v191 <= v170) != 0 {
		goto L26
	} else {
		goto L70
	}
L70:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	if v194 == int32(0) {
		goto L26
	} else {
		goto L71
	}
L71:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+26)))
	if v198 != 0 {
		goto L44
	} else {
		goto L72
	}
L72:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v194+v170<<(uint(int32(2))%32))))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	if v115 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	if base.B2i32(v226 == int32(0))|base.B2i32(v226 != v229) != 0 {
		v247 = v226
		v248 = v229
		goto L79
	} else {
		goto L80
	}
L74:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	if v217 == int32(0) {
		goto L44
	} else {
		goto L77
	}
L75:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v206 <= v170 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v220 = v115 + v206<<(uint(int32(4))%32) + v170*int32(100) + int32(24)
	goto L73
L77:
	;
	v220 = v217
	goto L73
L78:
	;
	if v247-v248 == int32(0) {
		v170 = v170 + int32(1)
		goto L64
	} else {
		goto L85
	}
L79:
	;
	goto L78
L80:
	;
	v232 = v220
	v233 = v203
	goto L81
L81:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+1)))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+1)))
	if v237 == int32(0) {
		v247 = v237
		v248 = v236
		goto L79
	} else {
		goto L83
	}
L82:
	;
	v247 = v237
	v248 = v236
	goto L79
L83:
	;
	v240 = int32(1)
	if v237 == v236 {
		v232 = v232 + v240
		v233 = v233 + v240
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	goto L65
L86:
	;
	v270 = int32(_a_F_get_query_def_1)
	goto L88
L87:
	;
	v270 = int32(_a_F_get_query_def_2)
	goto L88
L88:
	;
	F_appendStringInfoString(m, v98, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v273 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v376 = v18 + int32(264)
	F_get_target_list(m, v374, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L109
	}
L91:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v276 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	F_appendStringInfoString(m, v98, int32(_a_F_get_query_def_3))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	F_appendStringInfoString(m, v98, int32(_a_F_get_query_def_4))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L108
	}
L95:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v282 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_appendStringInfoChar(m, v98, int32(41))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L107
	}
L97:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v286 <= int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v282)+12))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	F_appendStringInfoString(m, v98, int32(_a_F_get_query_def_5))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v299 = F_get_rule_sortgroupclause(m, v294, v295, int32(0), v18+int32(264))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v301 < int32(2) {
		goto L96
	} else {
		goto L101
	}
L101:
	;
	v306 = int32(1)
	goto L102
L102:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v282)+12))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319+v306<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v98, int32(_a_F_get_query_def_6))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L4
	} else {
		goto L104
	}
L103:
	;
	goto L96
L104:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v332 = F_get_rule_sortgroupclause(m, v327, v328, int32(0), v18+int32(264))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	v335 = v306 + int32(1)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v335 < v336 {
		v306 = v335
		goto L102
	} else {
		goto L106
	}
L106:
	;
	goto L103
L107:
	;
	goto L90
L108:
	;
	goto L90
L109:
	;
	F_get_from_clause(m, l0, int32(_a_F_get_query_def_7), v376)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	if v383 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_appendContextKeyword(m, v376, int32(_a_F_get_query_def_8), int32(-8), int32(8), int32(1))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L4
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v395 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
	F_get_rule_expr(m, v391, v376, int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v557 != 0 {
		goto L149
	} else {
		goto L150
	}
L117:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v398 == int32(0) {
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	F_appendContextKeyword(m, v18+int32(264), int32(_a_F_get_query_def_9), int32(-8), int32(8), int32(1))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L4
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	if v409 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	F_appendStringInfoString(m, v98, int32(_a_F_get_query_def_10))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L4
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+298)))
	v416 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+298)) = uint8(v416)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v418 != 0 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	goto L124
L126:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+298)) = uint8(v415)
	goto L116
L127:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	if v420 <= int32(0) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v470 == int32(0) {
		goto L126
	} else {
		goto L139
	}
L130:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v418)+12))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	F_appendStringInfoString(m, v98, int32(_a_F_get_query_def_5))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_get_rule_groupingset(m, v424, v428, int32(1), v18+int32(264))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	if v434 <= int32(1) {
		goto L126
	} else {
		goto L133
	}
L133:
	;
	v439 = int32(1)
	goto L134
L134:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v418)+12))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v452+v439<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v98, int32(_a_F_get_query_def_6))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L136
	}
L135:
	;
	goto L126
L136:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_get_rule_groupingset(m, v456, v460, int32(1), v18+int32(264))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	v467 = v439 + int32(1)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	if v467 < v468 {
		v439 = v467
		goto L134
	} else {
		goto L138
	}
L138:
	;
	goto L135
L139:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v474 <= int32(0) {
		goto L126
	} else {
		goto L140
	}
L140:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v470)+12))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	F_appendStringInfoString(m, v98, int32(_a_F_get_query_def_5))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v478)+4))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v487 = F_get_rule_sortgroupclause(m, v482, v483, int32(0), v18+int32(264))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v489 < int32(2) {
		goto L126
	} else {
		goto L143
	}
L143:
	;
	v494 = int32(1)
	goto L144
L144:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v470)+12))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v507+v494<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v98, int32(_a_F_get_query_def_6))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L146
	}
L145:
	;
	goto L126
L146:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v520 = F_get_rule_sortgroupclause(m, v515, v516, int32(0), v18+int32(264))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	v523 = v494 + int32(1)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v523 < v524 {
		v494 = v523
		goto L144
	} else {
		goto L148
	}
L148:
	;
	goto L145
L149:
	;
	v559 = v18 + int32(264)
	F_appendContextKeyword(m, v559, int32(_a_F_get_query_def_11), int32(-8), int32(8), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L4
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v571 == int32(0) {
		goto L22
	} else {
		goto L154
	}
L152:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	F_get_rule_expr(m, v566, v559, int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v571)+4))
	if v574 <= int32(0) {
		goto L22
	} else {
		goto L155
	}
L155:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v18)+264))
	v578 = int32(0)
	v582 = v578
	v584 = v574
	v587 = v578
	goto L156
L156:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v571)+12))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v595+v582<<(uint(int32(2))%32))))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	if v600 != 0 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	goto L22
L158:
	;
	if v587 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L159:
	;
	v629 = v584
	v630 = v587
	goto L160
L160:
	;
	v632 = v582 + int32(1)
	if v632 < v629 {
		v582 = v632
		v584 = v629
		v587 = v630
		goto L156
	} else {
		goto L170
	}
L161:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	v614 = F_quote_identifier(m, v613)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L4
	} else {
		goto L167
	}
L162:
	;
	F_appendContextKeyword(m, v18+int32(264), int32(_a_F_get_query_def_12), int32(-8), int32(8), int32(1))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L4
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	F_appendStringInfoString(m, v577, v587)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L4
	} else {
		goto L166
	}
L165:
	;
	goto L161
L166:
	;
	goto L161
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v614
	F_appendStringInfo(m, v577, int32(_a_F_get_query_def_13), v18+int32(48))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_get_rule_windowspec(m, v599, v622, v18+int32(264))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v571)+4))
	v629 = v628
	v630 = int32(_a_F_get_query_def_6)
	goto L160
L170:
	;
	goto L157
L171:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v639)+12))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v642 = int32(2)
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v640+v641<<(uint(v642)%32)-int32(4))))
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+284)))
	if v648&v642 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	F_appendStringInfoChar(m, v634, int32(32))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L4
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v647)+20)))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v647)+16))
	v661 = F_generate_relation_name(m, v659, int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L4
	} else {
		goto L176
	}
L175:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v18)+292))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+292)) = v654 + int32(8)
	goto L174
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v661
	if v658 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v666 = int32(_a_F_get_query_def_5)
	goto L179
L178:
	;
	v666 = int32(_a_F_get_query_def_14)
	goto L179
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v666
	F_appendStringInfo(m, v634, int32(_a_F_get_query_def_15), v18-int32(-64))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v676 = v18 + int32(264)
	F_get_rte_alias(m, v647, v673, int32(0), v676)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	F_appendStringInfoString(m, v634, int32(_a_F_get_query_def_16))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_get_update_query_targetlist_def(m, l0, v682, v676, v647)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	F_get_from_clause(m, l0, int32(_a_F_get_query_def_7), v676)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v688)+8))
	if v689 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	F_appendContextKeyword(m, v676, int32(_a_F_get_query_def_8), int32(-8), int32(8), int32(1))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L4
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v701 == int32(0) {
		goto L21
	} else {
		goto L190
	}
L188:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v696)+8))
	F_get_rule_expr(m, v697, v676, int32(0))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	F_get_returning_clause(m, l0, v18+int32(264))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	goto L21
L192:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v713 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v1470 = int32(0)
	v1478 = v49
	v1481 = v9
	v1483 = v9
	goto L23
L194:
	;
	goto L195
L195:
	;
	v717 = int32(0)
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v713)+4))
	if v718 <= v717 {
		v1433 = v717
		v1441 = v49
		goto L24
	} else {
		goto L196
	}
L196:
	;
	v721 = int32(0)
	if v721 < v718 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v724 = v718
	goto L199
L198:
	;
	v724 = v721
	goto L199
L199:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v713)+12))
	v726 = int32(0)
	v730 = v717
	v731 = v726
	v733 = v726
	v736 = v726
	v738 = v49
	goto L200
L200:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v725+v731<<(uint(int32(2))%32))))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)+12))
	switch v748 - int32(1) {
	case 0:
		goto L204
	default:
		v766 = v730
		v767 = v738
		goto L202
	case 4:
		goto L203
	}
L201:
	;
	v1433 = v766
	v1441 = v767
	goto L24
L202:
	;
	v769 = v731 + int32(1)
	if v769 != v724 {
		v730 = v766
		v731 = v769
		v733 = v766
		v736 = v767
		v738 = v767
		goto L200
	} else {
		goto L210
	}
L203:
	;
	if v733 != 0 {
		goto L25
	} else {
		goto L209
	}
L204:
	;
	if v736 == int32(0) {
		v766 = v733
		v767 = v747
		goto L202
	} else {
		goto L205
	}
L205:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L4
	} else {
		goto L206
	}
L206:
	;
	F_errmsg_internal(m, int32(_a_F_get_query_def_17), int32(0))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L4
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_19), int32(_a_F_get_query_def_20))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L4
	} else {
		goto L208
	}
L208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L209:
	;
	v766 = v747
	v767 = v736
	goto L202
L210:
	;
	goto L201
L211:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v776)+12))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v779 = int32(2)
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v777+v778<<(uint(v779)%32)-int32(4))))
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+284)))
	if v785&v779 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	F_appendStringInfoChar(m, v771, int32(32))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L4
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+20)))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v784)+16))
	v798 = F_generate_relation_name(m, v796, int32(0))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L4
	} else {
		goto L216
	}
L215:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v18)+292))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+292)) = v791 + int32(8)
	goto L214
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+132)) = v798
	if v795 != 0 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v803 = int32(_a_F_get_query_def_5)
	goto L219
L218:
	;
	v803 = int32(_a_F_get_query_def_14)
	goto L219
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = v803
	F_appendStringInfo(m, v771, int32(_a_F_get_query_def_21), v18+int32(128))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L4
	} else {
		goto L220
	}
L220:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v813 = v18 + int32(264)
	F_get_rte_alias(m, v784, v810, int32(0), v813)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L4
	} else {
		goto L221
	}
L221:
	;
	F_get_from_clause(m, l0, int32(_a_F_get_query_def_22), v813)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L4
	} else {
		goto L222
	}
L222:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v819)+8))
	if v820 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	F_appendContextKeyword(m, v813, int32(_a_F_get_query_def_8), int32(-8), int32(8), int32(1))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L4
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v832 == int32(0) {
		goto L21
	} else {
		goto L228
	}
L226:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v827)+8))
	F_get_rule_expr(m, v828, v813, int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L4
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	F_get_returning_clause(m, l0, v18+int32(264))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L4
	} else {
		goto L229
	}
L229:
	;
	goto L21
L230:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v844)+12))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v847 = int32(2)
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v845+v846<<(uint(v847)%32)-int32(4))))
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+284)))
	if v853&v847 != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	F_appendStringInfoChar(m, v839, int32(32))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L4
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+20)))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v852)+16))
	v866 = F_generate_relation_name(m, v864, int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L4
	} else {
		goto L235
	}
L234:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v18)+292))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+292)) = v859 + int32(8)
	goto L233
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+164)) = v866
	if v863 != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v871 = int32(_a_F_get_query_def_5)
	goto L238
L237:
	;
	v871 = int32(_a_F_get_query_def_14)
	goto L238
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+160)) = v871
	F_appendStringInfo(m, v839, int32(_a_F_get_query_def_23), v18+int32(160))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L4
	} else {
		goto L239
	}
L239:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v881 = v18 + int32(264)
	F_get_rte_alias(m, v852, v878, int32(0), v881)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L4
	} else {
		goto L240
	}
L240:
	;
	F_get_from_clause(m, l0, int32(_a_F_get_query_def_22), v881)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L4
	} else {
		goto L241
	}
L241:
	;
	F_appendContextKeyword(m, v881, int32(_a_F_get_query_def_24), int32(-8), int32(8), int32(2))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L4
	} else {
		goto L242
	}
L242:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	F_get_rule_expr(m, v893, v881, int32(0))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L4
	} else {
		goto L243
	}
L243:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v897 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v1306 == int32(0) {
		goto L21
	} else {
		goto L340
	}
L245:
	;
	v900 = int32(_a_F_get_query_def_25)
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v897)+4))
	if v901 <= int32(0) {
		v939 = v900
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v897)+4))
	if v951 <= int32(0) {
		goto L244
	} else {
		goto L257
	}
L247:
	;
	v904 = int32(0)
	if v904 < v901 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v907 = v901
	goto L250
L249:
	;
	v907 = v904
	goto L250
L250:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v897)+12))
	v912 = int32(0)
	goto L251
L251:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v908+v912<<(uint(int32(2))%32))))
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v928)+4))
	if v929 != int32(1) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v939 = int32(_a_F_get_query_def_26)
	goto L246
L253:
	;
	v933 = v912 + int32(1)
	if v907 != v933 {
		v912 = v933
		goto L251
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	goto L252
L256:
	;
	v939 = v900
	goto L246
L257:
	;
	v962 = int32(0)
	goto L258
L258:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v897)+12))
	v971 = int32(2)
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v970+v962<<(uint(v971)%32))))
	F_appendContextKeyword(m, v18+int32(264), int32(_a_F_get_query_def_27), int32(-8), int32(8), v971)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L4
	} else {
		goto L260
	}
L259:
	;
	goto L244
L260:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v974)+4))
	switch v984 {
	case 0:
		v1002 = int32(_a_F_get_query_def_28)
		goto L261
	case 1:
		goto L262
	case 2:
		goto L264
	default:
		goto L263
	}
L261:
	;
	F_appendStringInfoString(m, v839, v1002)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L4
	} else {
		goto L268
	}
L262:
	;
	v1002 = int32(_a_F_get_query_def_29)
	goto L261
L263:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L4
	} else {
		goto L265
	}
L264:
	;
	v1002 = v939
	goto L261
L265:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v974)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+144)) = v989
	F_errmsg_internal(m, int32(_a_F_get_query_def_30), v18+int32(144))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L4
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_31), int32(_a_F_get_query_def_32))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L4
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v974)+16))
	if v1005 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1007 = v18 + int32(264)
	F_appendContextKeyword(m, v1007, int32(_a_F_get_query_def_33), int32(-8), int32(8), int32(3))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L4
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	F_appendContextKeyword(m, v18+int32(264), int32(_a_F_get_query_def_34), int32(-8), int32(8), int32(3))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L4
	} else {
		goto L274
	}
L272:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v974)+16))
	F_get_rule_expr(m, v1014, v1007, int32(0))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L4
	} else {
		goto L273
	}
L273:
	;
	goto L271
L274:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v974)+8))
	switch v1027 - int32(2) {
	case 0:
		goto L278
	case 1:
		goto L279
	case 2:
		goto L277
	default:
		goto L275
	case 5:
		goto L276
	}
L275:
	;
	v1288 = v962 + int32(1)
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v897)+4))
	if v1288 < v1289 {
		v962 = v1288
		goto L258
	} else {
		goto L339
	}
L276:
	;
	F_appendStringInfoString(m, v839, int32(_a_F_get_query_def_35))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L4
	} else {
		goto L338
	}
L277:
	;
	F_appendStringInfoString(m, v839, int32(_a_F_get_query_def_36))
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L4
	} else {
		goto L337
	}
L278:
	;
	F_appendStringInfoString(m, v839, int32(_a_F_get_query_def_37))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L4
	} else {
		goto L335
	}
L279:
	;
	F_appendStringInfoString(m, v839, int32(_a_F_get_query_def_38))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L4
	} else {
		goto L280
	}
L280:
	;
	v1033 = int32(0)
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v974)+20))
	if v1034 == v1033 {
		v1145 = v1033
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v974)+12))
	switch v1152 - int32(1) {
	case 0:
		goto L307
	case 1:
		v1156 = int32(_a_F_get_query_def_39)
		goto L306
	default:
		goto L305
	}
L282:
	;
	F_appendStringInfoString(m, v839, int32(_a_F_get_query_def_40))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L4
	} else {
		goto L283
	}
L283:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v974)+20))
	if v1040 == int32(0) {
		v1145 = v1033
		goto L281
	} else {
		goto L284
	}
L284:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1040)+4))
	if v1043 <= int32(0) {
		v1124 = v1033
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v974)+20))
	if v1130 == int32(0) {
		v1145 = v1124
		goto L281
	} else {
		goto L303
	}
L286:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1040)+12))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1046)))
	F_appendStringInfoString(m, v839, int32(_a_F_get_query_def_5))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L4
	} else {
		goto L287
	}
L287:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v852)+16))
	v1052 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1047)+8)))
	v1054 = F_get_attname(m, v1051, v1052, int32(0))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L4
	} else {
		goto L288
	}
L288:
	;
	v1056 = F_quote_identifier(m, v1054)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L4
	} else {
		goto L289
	}
L289:
	;
	F_appendStringInfoString(m, v839, v1056)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L4
	} else {
		goto L290
	}
L290:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+4))
	v1064 = F_processIndirection(m, v1061, v18+int32(264))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L4
	} else {
		goto L291
	}
L291:
	;
	v1066 = F_lappend(m, int32(0), v1064)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L4
	} else {
		goto L292
	}
L292:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1040)+4))
	if v1069 < int32(2) {
		v1124 = v1066
		goto L285
	} else {
		goto L293
	}
L293:
	;
	v1076 = int32(1)
	v1081 = v1066
	goto L294
L294:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1040)+12))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1087+v1076<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v839, int32(_a_F_get_query_def_6))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L4
	} else {
		goto L296
	}
L295:
	;
	v1124 = v1109
	goto L285
L296:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v852)+16))
	v1096 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1091)+8)))
	v1098 = F_get_attname(m, v1095, v1096, int32(0))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L4
	} else {
		goto L297
	}
L297:
	;
	v1100 = F_quote_identifier(m, v1098)
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L4
	} else {
		goto L298
	}
L298:
	;
	F_appendStringInfoString(m, v839, v1100)
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L4
	} else {
		goto L299
	}
L299:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+4))
	v1107 = F_processIndirection(m, v1104, v18+int32(264))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L4
	} else {
		goto L300
	}
L300:
	;
	v1109 = F_lappend(m, v1081, v1107)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L4
	} else {
		goto L301
	}
L301:
	;
	v1112 = v1076 + int32(1)
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1040)+4))
	if v1112 < v1113 {
		v1076 = v1112
		v1081 = v1109
		goto L294
	} else {
		goto L302
	}
L302:
	;
	goto L295
L303:
	;
	F_appendStringInfoChar(m, v839, int32(41))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L4
	} else {
		goto L304
	}
L304:
	;
	v1145 = v1124
	goto L281
L305:
	;
	if v1145 != 0 {
		goto L309
	} else {
		goto L310
	}
L306:
	;
	F_appendStringInfoString(m, v839, v1156)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L4
	} else {
		goto L308
	}
L307:
	;
	v1156 = int32(_a_F_get_query_def_41)
	goto L306
L308:
	;
	goto L305
L309:
	;
	v1161 = v18 + int32(264)
	F_appendContextKeyword(m, v1161, int32(_a_F_get_query_def_42), int32(-8), int32(8), int32(4))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L4
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	F_appendStringInfoString(m, v839, int32(_a_F_get_query_def_43))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L4
	} else {
		goto L334
	}
L312:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+4))
	if v1168 <= int32(0) {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	F_appendStringInfoChar(m, v839, int32(41))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L4
	} else {
		goto L333
	}
L314:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+12))
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1171)))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v18)+264))
	F_appendStringInfoString(m, v1173, int32(_a_F_get_query_def_5))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L4
	} else {
		goto L315
	}
L315:
	;
	if v1172 == int32(0) {
		goto L317
	} else {
		goto L318
	}
L316:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+4))
	if v1191 < int32(2) {
		goto L313
	} else {
		goto L322
	}
L317:
	;
	F_get_rule_expr(m, v1172, v18+int32(264), int32(0))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L4
	} else {
		goto L321
	}
L318:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1172)))
	if v1179 != int32(6) {
		goto L317
	} else {
		goto L319
	}
L319:
	;
	v1183 = F_get_variable(m, v1172, int32(1), v1161)
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L4
	} else {
		goto L320
	}
L320:
	;
	goto L316
L321:
	;
	goto L316
L322:
	;
	v1198 = int32(1)
	goto L323
L323:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+12))
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1209+v1198<<(uint(int32(2))%32))))
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v18)+264))
	F_appendStringInfoString(m, v1214, int32(_a_F_get_query_def_6))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L4
	} else {
		goto L325
	}
L324:
	;
	goto L313
L325:
	;
	if v1213 == int32(0) {
		goto L327
	} else {
		goto L328
	}
L326:
	;
	v1234 = v1198 + int32(1)
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+4))
	if v1234 < v1235 {
		v1198 = v1234
		goto L323
	} else {
		goto L332
	}
L327:
	;
	F_get_rule_expr(m, v1213, v18+int32(264), int32(0))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L4
	} else {
		goto L331
	}
L328:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1213)))
	if v1220 != int32(6) {
		goto L327
	} else {
		goto L329
	}
L329:
	;
	v1226 = F_get_variable(m, v1213, int32(1), v18+int32(264))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L4
	} else {
		goto L330
	}
L330:
	;
	goto L326
L331:
	;
	goto L326
L332:
	;
	goto L324
L333:
	;
	goto L275
L334:
	;
	goto L275
L335:
	;
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v974)+20))
	F_get_update_query_targetlist_def(m, l0, v1261, v18+int32(264), v852)
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L4
	} else {
		goto L336
	}
L336:
	;
	goto L275
L337:
	;
	goto L275
L338:
	;
	goto L275
L339:
	;
	goto L259
L340:
	;
	F_get_returning_clause(m, l0, v18+int32(264))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L4
	} else {
		goto L341
	}
L341:
	;
	goto L21
L342:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L4
	} else {
		goto L364
	}
L343:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1313)))
	if v1316 != int32(222) {
		goto L342
	} else {
		goto L344
	}
L344:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v18)+264))
	F_appendContextKeyword(m, v18+int32(264), int32(_a_F_get_query_def_5), int32(0), int32(8), int32(1))
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L4
	} else {
		goto L345
	}
L345:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+4))
	v1329 = F_quote_identifier(m, v1328)
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L4
	} else {
		goto L346
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+176)) = v1329
	F_appendStringInfo(m, v1319, int32(_a_F_get_query_def_44), v18+int32(176))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L4
	} else {
		goto L347
	}
L347:
	;
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+8))
	if v1337 == int32(0) {
		goto L21
	} else {
		goto L348
	}
L348:
	;
	F_appendStringInfoString(m, v1319, int32(_a_F_get_query_def_6))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L4
	} else {
		goto L349
	}
L349:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+8))
	F_appendStringInfoChar(m, v1319, int32(39))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L4
	} else {
		goto L350
	}
L350:
	;
	v1349 = v1343
	goto L351
L351:
	;
	v1362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1349))))
	v1363 = base.I32_extend8_s(v1362)
	if v1362 != int32(39) {
		goto L355
	} else {
		goto L356
	}
L352:
	;
	F_appendStringInfoChar(m, v1319, int32(39))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L4
	} else {
		goto L363
	}
L353:
	;
	goto L352
L354:
	;
	F_appendStringInfoChar(m, v1319, v1363)
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L4
	} else {
		goto L362
	}
L355:
	;
	if v1362 == int32(0) {
		goto L353
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	F_appendStringInfoChar(m, v1319, v1363)
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L4
	} else {
		goto L361
	}
L358:
	;
	if v1363 != int32(92) {
		goto L354
	} else {
		goto L359
	}
L359:
	;
	v1371 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_query_def[1])))
	if v1371&int32(1) != 0 {
		goto L354
	} else {
		goto L360
	}
L360:
	;
	goto L357
L361:
	;
	goto L354
L362:
	;
	v1349 = v1349 + int32(1)
	goto L351
L363:
	;
	goto L21
L364:
	;
	F_errmsg_internal(m, int32(_a_F_get_query_def_45), int32(0))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L4
	} else {
		goto L365
	}
L365:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_46), int32(_a_F_get_query_def_47))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L4
	} else {
		goto L366
	}
L366:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L367:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1400
	F_errmsg_internal(m, int32(_a_F_get_query_def_48), v18)
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L4
	} else {
		goto L368
	}
L368:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_49), int32(_a_F_get_query_def_50))
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L4
	} else {
		goto L369
	}
L369:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L370:
	;
	goto L21
L371:
	;
	goto L22
L372:
	;
	F_errmsg_internal(m, int32(_a_F_get_query_def_51), int32(0))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L4
	} else {
		goto L373
	}
L373:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_52), int32(_a_F_get_query_def_20))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L4
	} else {
		goto L374
	}
L374:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L375:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L4
	} else {
		goto L376
	}
L376:
	;
	F_errmsg_internal(m, int32(_a_F_get_query_def_53), int32(0))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L4
	} else {
		goto L377
	}
L377:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_54), int32(_a_F_get_query_def_20))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L4
	} else {
		goto L378
	}
L378:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L379:
	;
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v18)+292))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+292)) = v1495 + int32(8)
	F_appendStringInfoChar(m, v708, int32(32))
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L4
	} else {
		goto L382
	}
L380:
	;
	goto L381
L381:
	;
	v1502 = int32(0)
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1491)+16))
	v1505 = F_generate_relation_name(m, v1503, v1502)
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L4
	} else {
		goto L383
	}
L382:
	;
	goto L381
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v1505
	F_appendStringInfo(m, v708, int32(_a_F_get_query_def_55), v18+int32(112))
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L4
	} else {
		goto L384
	}
L384:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_get_rte_alias(m, v1491, v1513, int32(1), v18+int32(264))
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L4
	} else {
		goto L385
	}
L385:
	;
	F_appendStringInfoChar(m, v708, int32(32))
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L4
	} else {
		goto L386
	}
L386:
	;
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v1522 == int32(0) {
		v1611 = v1502
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	switch v1621 - int32(1) {
	case 0:
		goto L410
	case 1:
		v1625 = int32(_a_F_get_query_def_56)
		goto L409
	default:
		goto L408
	}
L388:
	;
	F_appendStringInfoChar(m, v708, int32(40))
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L4
	} else {
		goto L389
	}
L389:
	;
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v1528 == int32(0) {
		v1611 = v1502
		goto L387
	} else {
		goto L390
	}
L390:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1528)+4))
	if int32(0) < v1531 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v1538 = int32(0)
	v1542 = v1502
	v1546 = int32(_a_F_get_query_def_5)
	goto L394
L392:
	;
	v1590 = v1502
	goto L393
L393:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v1599 == int32(0) {
		v1611 = v1590
		goto L387
	} else {
		goto L406
	}
L394:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v1528)+12))
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(v1551+v1538<<(uint(int32(2))%32))))
	v1556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1555)+26)))
	if v1556 == int32(0) {
		goto L396
	} else {
		goto L397
	}
L395:
	;
	v1590 = v1578
	goto L393
L396:
	;
	F_appendStringInfoString(m, v708, v1546)
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L4
	} else {
		goto L399
	}
L397:
	;
	v1578 = v1542
	v1579 = v1546
	goto L398
L398:
	;
	v1581 = v1538 + int32(1)
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1528)+4))
	if v1581 < v1582 {
		v1538 = v1581
		v1542 = v1578
		v1546 = v1579
		goto L394
	} else {
		goto L405
	}
L399:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1491)+16))
	v1562 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1555)+8)))
	v1564 = F_get_attname(m, v1561, v1562, int32(0))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L4
	} else {
		goto L400
	}
L400:
	;
	v1566 = F_quote_identifier(m, v1564)
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L4
	} else {
		goto L401
	}
L401:
	;
	F_appendStringInfoString(m, v708, v1566)
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L4
	} else {
		goto L402
	}
L402:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1555)+4))
	v1574 = F_processIndirection(m, v1571, v18+int32(264))
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L4
	} else {
		goto L403
	}
L403:
	;
	v1576 = F_lappend(m, v1542, v1574)
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L4
	} else {
		goto L404
	}
L404:
	;
	v1578 = v1576
	v1579 = int32(_a_F_get_query_def_6)
	goto L398
L405:
	;
	goto L395
L406:
	;
	F_appendStringInfoString(m, v708, int32(_a_F_get_query_def_57))
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L4
	} else {
		goto L407
	}
L407:
	;
	v1611 = v1590
	goto L387
L408:
	;
	if v1483 != 0 {
		goto L413
	} else {
		goto L414
	}
L409:
	;
	F_appendStringInfoString(m, v708, v1625)
	mBase = m.M
	v1627 = m.ExcPending
	if v1627 != 0 {
		goto L4
	} else {
		goto L411
	}
L410:
	;
	v1625 = int32(_a_F_get_query_def_58)
	goto L409
L411:
	;
	goto L408
L412:
	;
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v1661 == int32(0) {
		goto L429
	} else {
		goto L430
	}
L413:
	;
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1478)+36))
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v18)+268))
	v1631 = int32(0)
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v18)+284))
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v18)+288))
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v18)+292))
	F_get_query_def(m, v1629, v708, v1630, v1631, v1631, v1633, v1634, v1635)
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L4
	} else {
		goto L416
	}
L414:
	;
	goto L415
L415:
	;
	if v1481 != 0 {
		goto L417
	} else {
		goto L418
	}
L416:
	;
	goto L412
L417:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+80))
	F_get_values_def(m, v1638, v18+int32(264))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L4
	} else {
		goto L420
	}
L418:
	;
	goto L419
L419:
	;
	if v1611 != 0 {
		goto L421
	} else {
		goto L422
	}
L420:
	;
	goto L412
L421:
	;
	v1644 = v18 + int32(264)
	F_appendContextKeyword(m, v1644, int32(_a_F_get_query_def_59), int32(-8), int32(8), int32(2))
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L4
	} else {
		goto L424
	}
L422:
	;
	goto L423
L423:
	;
	F_appendStringInfoString(m, v708, int32(_a_F_get_query_def_60))
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L4
	} else {
		goto L427
	}
L424:
	;
	F_get_rule_list_toplevel(m, v1611, v1644, int32(0))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L4
	} else {
		goto L425
	}
L425:
	;
	F_appendStringInfoChar(m, v708, int32(41))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L4
	} else {
		goto L426
	}
L426:
	;
	goto L412
L427:
	;
	goto L412
L428:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L4
	} else {
		goto L458
	}
L429:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v1743 == int32(0) {
		goto L21
	} else {
		goto L456
	}
L430:
	;
	F_appendStringInfoString(m, v708, int32(_a_F_get_query_def_61))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L4
	} else {
		goto L431
	}
L431:
	;
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+8))
	if v1667 != 0 {
		goto L433
	} else {
		goto L434
	}
L432:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+4))
	if v1714 == int32(1) {
		goto L447
	} else {
		goto L448
	}
L433:
	;
	F_appendStringInfoChar(m, v708, int32(40))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L4
	} else {
		goto L436
	}
L434:
	;
	goto L435
L435:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+16))
	if v1697 == int32(0) {
		goto L432
	} else {
		goto L442
	}
L436:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+8))
	v1673 = v18 + int32(264)
	F_get_rule_expr(m, v1671, v1673, int32(0))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L4
	} else {
		goto L437
	}
L437:
	;
	F_appendStringInfoChar(m, v708, int32(41))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L4
	} else {
		goto L438
	}
L438:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+12))
	if v1680 == int32(0) {
		goto L432
	} else {
		goto L439
	}
L439:
	;
	v1683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+296)))
	v1684 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+296)) = uint8(v1684)
	F_appendContextKeyword(m, v1673, int32(_a_F_get_query_def_8), int32(-8), int32(8), int32(1))
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L4
	} else {
		goto L440
	}
L440:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+12))
	F_get_rule_expr(m, v1692, v1673, int32(0))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L4
	} else {
		goto L441
	}
L441:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+296)) = uint8(v1683)
	goto L432
L442:
	;
	v1700 = F_get_constraint_name(m, v1697)
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L4
	} else {
		goto L443
	}
L443:
	;
	if v1700 == int32(0) {
		goto L428
	} else {
		goto L444
	}
L444:
	;
	v1704 = F_quote_identifier(m, v1700)
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L4
	} else {
		goto L445
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v1704
	F_appendStringInfo(m, v708, int32(_a_F_get_query_def_62), v18+int32(96))
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L4
	} else {
		goto L446
	}
L446:
	;
	goto L432
L447:
	;
	F_appendStringInfoString(m, v708, int32(_a_F_get_query_def_63))
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L4
	} else {
		goto L450
	}
L448:
	;
	goto L449
L449:
	;
	F_appendStringInfoString(m, v708, int32(_a_F_get_query_def_64))
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L4
	} else {
		goto L451
	}
L450:
	;
	goto L429
L451:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+20))
	v1725 = v18 + int32(264)
	F_get_update_query_targetlist_def(m, l0, v1723, v1725, v1491)
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L4
	} else {
		goto L452
	}
L452:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+24))
	if v1728 == int32(0) {
		goto L429
	} else {
		goto L453
	}
L453:
	;
	F_appendContextKeyword(m, v1725, int32(_a_F_get_query_def_8), int32(-8), int32(8), int32(1))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L4
	} else {
		goto L454
	}
L454:
	;
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+24))
	F_get_rule_expr(m, v1737, v1725, int32(0))
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L4
	} else {
		goto L455
	}
L455:
	;
	goto L429
L456:
	;
	F_get_returning_clause(m, l0, v18+int32(264))
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L4
	} else {
		goto L457
	}
L457:
	;
	goto L21
L458:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v1754
	F_errmsg_internal(m, int32(_a_F_get_query_def_65), v18+int32(80))
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L4
	} else {
		goto L459
	}
L459:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_66), int32(_a_F_get_query_def_20))
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L4
	} else {
		goto L460
	}
L460:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L461:
	;
	v1783 = v18 + int32(264)
	F_appendContextKeyword(m, v1783, int32(_a_F_get_query_def_67), int32(-8), int32(8), int32(1))
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L4
	} else {
		goto L464
	}
L462:
	;
	goto L463
L463:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1797 != 0 {
		goto L466
	} else {
		goto L467
	}
L464:
	;
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_get_rule_orderby(m, v1790, v1791, base.B2i32(v95 != int32(0)), v1783)
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L4
	} else {
		goto L465
	}
L465:
	;
	goto L463
L466:
	;
	v1799 = v18 + int32(264)
	F_appendContextKeyword(m, v1799, int32(_a_F_get_query_def_68), int32(-8), int32(8), int32(0))
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L4
	} else {
		goto L469
	}
L467:
	;
	goto L468
L468:
	;
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v1811 == int32(0) {
		goto L471
	} else {
		goto L472
	}
L469:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	F_get_rule_expr(m, v1806, v1799, int32(0))
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L4
	} else {
		goto L470
	}
L470:
	;
	goto L468
L471:
	;
	v1862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	if v1862 != int32(1) {
		goto L21
	} else {
		goto L487
	}
L472:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v1814 == int32(1) {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v1818 = v18 + int32(264)
	F_appendContextKeyword(m, v1818, int32(_a_F_get_query_def_69), int32(-8), int32(8), int32(0))
	mBase = m.M
	v1824 = m.ExcPending
	if v1824 != 0 {
		goto L4
	} else {
		goto L476
	}
L474:
	;
	goto L475
L475:
	;
	F_appendContextKeyword(m, v18+int32(264), int32(_a_F_get_query_def_70), int32(-8), int32(8), int32(0))
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
		goto L4
	} else {
		goto L481
	}
L476:
	;
	F_appendStringInfoChar(m, v86, int32(40))
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L4
	} else {
		goto L477
	}
L477:
	;
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_get_rule_expr(m, v1828, v1818, int32(0))
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L4
	} else {
		goto L478
	}
L478:
	;
	F_appendStringInfoChar(m, v86, int32(41))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L4
	} else {
		goto L479
	}
L479:
	;
	F_appendStringInfoString(m, v86, int32(_a_F_get_query_def_71))
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L4
	} else {
		goto L480
	}
L480:
	;
	goto L471
L481:
	;
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1846)))
	if v1847 != int32(7) {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	F_get_rule_expr(m, v1846, v18+int32(264), int32(0))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L4
	} else {
		goto L486
	}
L483:
	;
	v1850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1846)+24)))
	if v1850 != int32(1) {
		goto L482
	} else {
		goto L484
	}
L484:
	;
	F_appendStringInfoString(m, v86, int32(_a_F_get_query_def_72))
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L4
	} else {
		goto L485
	}
L485:
	;
	goto L471
L486:
	;
	goto L471
L487:
	;
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v1865 == int32(0) {
		goto L21
	} else {
		goto L488
	}
L488:
	;
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+4))
	if v1868 <= int32(0) {
		goto L21
	} else {
		goto L489
	}
L489:
	;
	v1872 = int32(0)
	goto L490
L490:
	;
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+12))
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1887+v1872<<(uint(int32(2))%32))))
	v1892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1891)+16)))
	if v1892 != 0 {
		goto L492
	} else {
		goto L493
	}
L491:
	;
	goto L21
L492:
	;
	v1953 = v1872 + int32(1)
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+4))
	if v1953 < v1954 {
		v1872 = v1953
		goto L490
	} else {
		goto L509
	}
L493:
	;
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+8))
	switch v1894 {
	case 0:
		goto L499
	case 1:
		v1914 = int32(_a_F_get_query_def_73)
		goto L495
	case 2:
		goto L498
	case 3:
		goto L497
	case 4:
		goto L496
	default:
		goto L494
	}
L494:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v18)+268))
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(v1923)+12))
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(v1924)))
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1925)+4))
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+12))
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+4))
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1927+v1928<<(uint(int32(2))%32)-int32(4))))
	v1935 = F_quote_identifier(m, v1934)
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L4
	} else {
		goto L504
	}
L495:
	;
	F_appendContextKeyword(m, v18+int32(264), v1914, int32(-8), int32(8), int32(0))
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L4
	} else {
		goto L503
	}
L496:
	;
	v1914 = int32(_a_F_get_query_def_74)
	goto L495
L497:
	;
	v1914 = int32(_a_F_get_query_def_75)
	goto L495
L498:
	;
	v1914 = int32(_a_F_get_query_def_76)
	goto L495
L499:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1898 = m.ExcPending
	if v1898 != 0 {
		goto L4
	} else {
		goto L500
	}
L500:
	;
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v1899
	F_errmsg_internal(m, int32(_a_F_get_query_def_77), v18+int32(32))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L4
	} else {
		goto L501
	}
L501:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_78), int32(_a_F_get_query_def_79))
	mBase = m.M
	v1910 = m.ExcPending
	if v1910 != 0 {
		goto L4
	} else {
		goto L502
	}
L502:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L503:
	;
	goto L494
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1935
	F_appendStringInfo(m, v86, int32(_a_F_get_query_def_80), v18+int32(16))
	mBase = m.M
	v1942 = m.ExcPending
	if v1942 != 0 {
		goto L4
	} else {
		goto L505
	}
L505:
	;
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+12))
	switch v1944 - int32(1) {
	case 0:
		goto L507
	case 1:
		v1948 = int32(_a_F_get_query_def_81)
		goto L506
	default:
		goto L492
	}
L506:
	;
	F_appendStringInfoString(m, v86, v1948)
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L4
	} else {
		goto L508
	}
L507:
	;
	v1948 = int32(_a_F_get_query_def_82)
	goto L506
L508:
	;
	goto L492
L509:
	;
	goto L491
}
func F_get_update_query_targetlist_def(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	v5 = int32(0)
	v16 = l1 + int32(4)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	if v18 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v98 <= int32(0) {
		goto L1
	} else {
		goto L23
	}
L3:
	;
	if l1 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if l1 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v23 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v27 = int32(0)
	v35 = v5
	goto L10
L8:
	;
	v72 = v5
	goto L9
L9:
	;
	if v72 != 0 {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v27<<(uint(int32(2))%32))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+26)))
	if v46 != int32(1) {
		v59 = v35
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v72 = v59
	goto L9
L12:
	;
	v61 = v27 + int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v61 < v62 {
		v27 = v61
		v35 = v59
		goto L10
	} else {
		goto L18
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v50 != int32(22) {
		v59 = v35
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v53 != int32(5) {
		v59 = v35
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v56 = F_lappend(m, v35, v49)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	v59 = v56
	goto L12
L18:
	;
	goto L11
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v79 = v78
	goto L21
L20:
	;
	v79 = v5
	goto L21
L21:
	;
	v92 = v72
	v93 = v79
	v94 = l1 + int32(4)
	goto L2
L22:
	;
	v92 = v5
	v93 = v5
	v94 = v16
	goto L2
L23:
	;
	v103 = int32(0)
	v108 = int32(_a_F_get_update_query_targetlist_def_0)
	v112 = v93
	v115 = v5
	v116 = v5
	goto L24
L24:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117+v115<<(uint(int32(2))%32))))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+26)))
	if v122 != 0 {
		v374 = v103
		v379 = v108
		v383 = v112
		v387 = v116
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L1
L26:
	;
	v389 = v115 + int32(1)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v389 < v390 {
		v103 = v374
		v108 = v379
		v112 = v383
		v115 = v389
		v116 = v387
		goto L24
	} else {
		goto L109
	}
L27:
	;
	F_appendStringInfoString(m, v17, v108)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	if v112 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v342 = int32(*(*int16)(unsafe.Add(mBase, uint32(v121)+8)))
	v344 = F_get_attname(m, v341, v342, int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L16
	} else {
		goto L95
	}
L30:
	;
	v331 = v103
	v336 = v112
	v340 = v116
	goto L29
L31:
	;
	goto L32
L32:
	;
	if v103 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v331 = v103
	v336 = v112
	v340 = v116
	goto L29
L34:
	;
	goto L35
L35:
	;
	v127 = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v129 == v127 {
		v167 = v127
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v167 != 0 {
		goto L51
	} else {
		goto L52
	}
L37:
	;
	v132 = v129
	goto L38
L38:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	switch v146 - int32(14) {
	case 0:
		goto L42
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v167 = v132
		goto L36
	case 12:
		goto L43
	default:
		goto L44
	}
L39:
	;
	v167 = int32(0)
	goto L36
L40:
	;
	if v159 != 0 {
		v132 = v159
		goto L38
	} else {
		goto L48
	}
L41:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v132)+20))
	if v155 != int32(2) {
		v167 = v132
		goto L36
	} else {
		goto L47
	}
L42:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v132)+36))
	if v154 != 0 {
		v159 = v154
		goto L40
	} else {
		goto L46
	}
L43:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v159 = v153
	goto L40
L44:
	;
	if v146 == int32(55) {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v167 = v132
	goto L36
L46:
	;
	v167 = v132
	goto L36
L47:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v159 = v158
	goto L40
L48:
	;
	goto L39
L49:
	;
	if v213 == int32(0) {
		v331 = v127
		v336 = v112
		v340 = v116
		goto L29
	} else {
		goto L70
	}
L50:
	;
	goto L49
L51:
	;
	v175 = v167
	goto L54
L52:
	;
	goto L53
L53:
	;
	v213 = int32(0)
	goto L50
L54:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	switch v176 - int32(15) {
	case 0:
		goto L62
	default:
		v213 = v175
		goto L50
	case 12:
		goto L61
	case 13:
		goto L60
	case 14:
		goto L59
	case 15:
		goto L58
	case 40:
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v210 != 0 {
		v175 = v210
		goto L54
	} else {
		goto L69
	}
L57:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
	if v204 != int32(2) {
		v213 = v175
		goto L50
	} else {
		goto L68
	}
L58:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	if v199 != int32(2) {
		v213 = v175
		goto L50
	} else {
		goto L67
	}
L59:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v175)+24))
	if v194 != int32(2) {
		v213 = v175
		goto L50
	} else {
		goto L66
	}
L60:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	if v189 != int32(2) {
		v213 = v175
		goto L50
	} else {
		goto L65
	}
L61:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
	if v184 != int32(2) {
		v213 = v175
		goto L50
	} else {
		goto L64
	}
L62:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	if v179 != int32(2) {
		v213 = v175
		goto L50
	} else {
		goto L63
	}
L63:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v175)+28))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	v209 = v183
	goto L56
L64:
	;
	v209 = v175 + int32(4)
	goto L56
L65:
	;
	v209 = v175 + int32(4)
	goto L56
L66:
	;
	v209 = v175 + int32(4)
	goto L56
L67:
	;
	v209 = v175 + int32(4)
	goto L56
L68:
	;
	v209 = v175 + int32(4)
	goto L56
L69:
	;
	goto L55
L70:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	if v216 != int32(8) {
		v331 = v127
		v336 = v112
		v340 = v116
		goto L29
	} else {
		goto L71
	}
L71:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	if v219 != int32(3) {
		v331 = v127
		v336 = v112
		v340 = v116
		goto L29
	} else {
		goto L72
	}
L72:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+20))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+76))
	v227 = int32(0)
	if v226 == v227 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L16
	} else {
		goto L91
	}
L74:
	;
	v315 = int32(0)
	goto L73
L75:
	;
	goto L76
L76:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v237 <= int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v315 = int32(0)
	goto L73
L78:
	;
	goto L79
L79:
	;
	if v237 != int32(1) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v315 = v302
	goto L73
L81:
	;
	v243 = int32(0)
	if v243 < v237 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v284 = v227
	v285 = v227
	goto L83
L83:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v290+v284<<(uint(int32(2))%32))))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+26)))
	v302 = v285 + (v295 ^ int32(1))
	goto L80
L84:
	;
	v246 = v237
	goto L86
L85:
	;
	v246 = v243
	goto L86
L86:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
	v252 = int32(0)
	v255 = v252
	v256 = v252
	v257 = v227
	goto L87
L87:
	;
	v262 = int32(2)
	v264 = v251 + v256<<(uint(v262)%32)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+26)))
	v267 = int32(1)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+26)))
	v274 = v257 + (v266 ^ v267) + (v271 ^ v267)
	v276 = v256 + v262
	v278 = v255 + v262
	if v278 != v246&int32(2147483646) {
		v255 = v278
		v256 = v276
		v257 = v274
		goto L87
	} else {
		goto L89
	}
L88:
	;
	if v246&int32(1) == int32(0) {
		v302 = v274
		goto L80
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	v284 = v276
	v285 = v274
	goto L83
L91:
	;
	v320 = v112 + int32(4)
	if base.Ui32(v320) < base.Ui32(v222+v223<<(uint(int32(2))%32)) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v326 = v320
	goto L94
L93:
	;
	v326 = int32(0)
	goto L94
L94:
	;
	v331 = v224
	v336 = v326
	v340 = v315
	goto L29
L95:
	;
	v346 = F_quote_identifier(m, v344)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L16
	} else {
		goto L96
	}
L96:
	;
	F_appendStringInfoString(m, v17, v346)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L16
	} else {
		goto L97
	}
L97:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v351 = F_processIndirection(m, v350, l2)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L16
	} else {
		goto L98
	}
L98:
	;
	if v331 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_update_query_targetlist_def_1))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L16
	} else {
		goto L107
	}
L100:
	;
	v363 = v351
	v365 = v340
	goto L99
L101:
	;
	goto L102
L102:
	;
	v357 = v340 - int32(1)
	if int32(0) < v357 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v374 = v331
	v379 = int32(_a_F_get_update_query_targetlist_def_2)
	v383 = v336
	v387 = v357
	goto L26
L104:
	;
	goto L105
L105:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L16
	} else {
		goto L106
	}
L106:
	;
	v363 = v331
	v365 = v357
	goto L99
L107:
	;
	v369 = int32(0)
	F_get_rule_expr(m, v363, l2, v369)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L16
	} else {
		goto L108
	}
L108:
	;
	v374 = v369
	v379 = int32(_a_F_get_update_query_targetlist_def_2)
	v383 = v336
	v387 = v365
	goto L26
L109:
	;
	goto L25
}
func F_markQueryForLocking(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return
L2:
	;
	v14 = l0
	v15 = l1
	goto L3
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v21 != int32(64) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L1
L5:
	;
	F_applyLockingClause(m, v14, v28, l2, l3)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L14
	} else {
		goto L28
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L14
	} else {
		goto L25
	}
L7:
	;
	switch v21 - int32(63) {
	case 0:
		goto L11
	default:
		goto L6
	case 2:
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	F_markQueryForLocking(m, v14, v72, l2, l3)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L14
	} else {
		goto L23
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v47 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(int32(2))%32)-int32(4))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if v35 == int32(1) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	if v35 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_applyLockingClause(m, v14, v28, l2, l3)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v41 = F_getRTEPermissionInfo(m, v40, v34)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v41)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+16)) = v43 | int64(4)
	goto L1
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v50 <= int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v55 = int32(0)
	goto L19
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v55<<(uint(int32(2))%32))))
	F_markQueryForLocking(m, v14, v65, l2, l3)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L14
	} else {
		goto L21
	}
L20:
	;
	goto L1
L21:
	;
	v69 = v55 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v69 < v70 {
		v55 = v69
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v75 != 0 {
		v15 = v75
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L1
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v80
	F_errmsg_internal(m, int32(_a_F_markQueryForLocking_0), v10)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L14
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_markQueryForLocking_1), int32(1939), int32(_a_F_markQueryForLocking_2))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v34)+36))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+60))
	if v93 != 0 {
		v14 = v92
		v15 = v93
		goto L3
	} else {
		goto L29
	}
L29:
	;
	goto L4
}
