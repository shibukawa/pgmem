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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	v7 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteQuery[0]))
	if v19 == v7 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v113 = m.ExcPending
		if v113 != 0 {
			return
		} else {
			F_errcode(m, int32(386))
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
				F_errmsg(m, int32(_a_F_ExecuteQuery_0), v15)
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(454), int32(_a_F_ExecuteQuery_2))
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
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
				v113 = m.ExcPending
				if v113 != 0 {
					return
				} else {
					F_errcode(m, int32(386))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
						F_errmsg(m, int32(_a_F_ExecuteQuery_0), v15)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(454), int32(_a_F_ExecuteQuery_2))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
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
					v129 = m.ExcPending
					if v129 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_ExecuteQuery_3), int32(0))
						mBase = m.M
						v133 = m.ExcPending
						if v133 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(170), int32(_a_F_ExecuteQuery_4))
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
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
													v142 = m.ExcPending
													if v142 != 0 {
														return
													} else {
														F_errcode(m, int32(151027844))
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_ExecuteQuery_5), int32(0))
															mBase = m.M
															v149 = m.ExcPending
															if v149 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(230), int32(_a_F_ExecuteQuery_4))
																mBase = m.M
																v154 = m.ExcPending
																if v154 != 0 {
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
														v142 = m.ExcPending
														if v142 != 0 {
															return
														} else {
															F_errcode(m, int32(151027844))
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
																return
															} else {
																F_errmsg(m, int32(_a_F_ExecuteQuery_5), int32(0))
																mBase = m.M
																v149 = m.ExcPending
																if v149 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(230), int32(_a_F_ExecuteQuery_4))
																	mBase = m.M
																	v154 = m.ExcPending
																	if v154 != 0 {
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
															v158 = m.ExcPending
															if v158 != 0 {
																return
															} else {
																F_errcode(m, int32(151027844))
																mBase = m.M
																v161 = m.ExcPending
																if v161 != 0 {
																	return
																} else {
																	F_errmsg(m, int32(_a_F_ExecuteQuery_5), int32(0))
																	mBase = m.M
																	v165 = m.ExcPending
																	if v165 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(235), int32(_a_F_ExecuteQuery_4))
																		mBase = m.M
																		v170 = m.ExcPending
																		if v170 != 0 {
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
															v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
															if v84 != 0 {
																v85 = int32(64)
															} else {
																v85 = int32(0)
															}
															v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
															if v88 != 0 {
																v89 = int32(0)
															} else {
																v89 = int32(2147483647)
															}
															v91 = v85
															v92 = v89
															v94 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteQuery[1]))
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
															F_PortalStart(m, v44, v43, v91, v95)
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return
															} else {
																v99 = F_PortalRun(m, v44, v92, int32(0), l4, l4, l5)
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return
																} else {
																	F_PortalDrop(m, v44, int32(0))
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return
																	} else {
																		if v42 != 0 {
																			F_FreeExecutorState(m, v42)
																			mBase = m.M
																			v105 = m.ExcPending
																			if v105 != 0 {
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
												v91 = v7
												v92 = int32(2147483647)
												v94 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteQuery[1]))
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
												F_PortalStart(m, v44, v43, v91, v95)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return
												} else {
													v99 = F_PortalRun(m, v44, v92, int32(0), l4, l4, l5)
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return
													} else {
														F_PortalDrop(m, v44, int32(0))
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return
														} else {
															if v42 != 0 {
																F_FreeExecutorState(m, v42)
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
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
											v142 = m.ExcPending
											if v142 != 0 {
												return
											} else {
												F_errcode(m, int32(151027844))
												mBase = m.M
												v145 = m.ExcPending
												if v145 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_ExecuteQuery_5), int32(0))
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(230), int32(_a_F_ExecuteQuery_4))
														mBase = m.M
														v154 = m.ExcPending
														if v154 != 0 {
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
												v142 = m.ExcPending
												if v142 != 0 {
													return
												} else {
													F_errcode(m, int32(151027844))
													mBase = m.M
													v145 = m.ExcPending
													if v145 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_ExecuteQuery_5), int32(0))
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(230), int32(_a_F_ExecuteQuery_4))
															mBase = m.M
															v154 = m.ExcPending
															if v154 != 0 {
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
													v158 = m.ExcPending
													if v158 != 0 {
														return
													} else {
														F_errcode(m, int32(151027844))
														mBase = m.M
														v161 = m.ExcPending
														if v161 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_ExecuteQuery_5), int32(0))
															mBase = m.M
															v165 = m.ExcPending
															if v165 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_ExecuteQuery_1), int32(235), int32(_a_F_ExecuteQuery_4))
																mBase = m.M
																v170 = m.ExcPending
																if v170 != 0 {
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
													v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
													if v84 != 0 {
														v85 = int32(64)
													} else {
														v85 = int32(0)
													}
													v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
													if v88 != 0 {
														v89 = int32(0)
													} else {
														v89 = int32(2147483647)
													}
													v91 = v85
													v92 = v89
													v94 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteQuery[1]))
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
													F_PortalStart(m, v44, v43, v91, v95)
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return
													} else {
														v99 = F_PortalRun(m, v44, v92, int32(0), l4, l4, l5)
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return
														} else {
															F_PortalDrop(m, v44, int32(0))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return
															} else {
																if v42 != 0 {
																	F_FreeExecutorState(m, v42)
																	mBase = m.M
																	v105 = m.ExcPending
																	if v105 != 0 {
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
										v91 = v7
										v92 = int32(2147483647)
										v94 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteQuery[1]))
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
										F_PortalStart(m, v44, v43, v91, v95)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return
										} else {
											v99 = F_PortalRun(m, v44, v92, int32(0), l4, l4, l5)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												F_PortalDrop(m, v44, int32(0))
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return
												} else {
													if v42 != 0 {
														F_FreeExecutorState(m, v42)
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
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
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v51<<(uint(int32(2))%32))+uint32(_c_F_ProcessQuery[1])))
								v59 = v58
							} else {
								v59 = v24
							}
							*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v59
						} else {
						}
						F_ExecutorFinish(m, v15)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							F_ExecutorEnd(m, v15)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
								F_UnregisterSnapshot(m, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
									F_UnregisterSnapshot(m, v72)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										F_pfree(m, v15)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
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
				v17 = F_expression_tree_walker_impl(m, l0, int32(1604), l1)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		} else {
			v17 = F_expression_tree_walker_impl(m, l0, int32(1604), l1)
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
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
		v184 = l0
		goto L11
	} else {
		goto L12
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L62
	} else {
		goto L72
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L62
	} else {
		goto L69
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L62
	} else {
		goto L66
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L62
	} else {
		goto L63
	}
L11:
	;
	return v184
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
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v83 == int32(0) {
		goto L10
	} else {
		goto L34
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
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_getInsertSelectQuery[0])))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v28 == int32(0) {
		v47 = v27
		v48 = v28
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v48-v47 != 0 {
		goto L13
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	if v27 != v28 {
		v47 = v27
		v48 = v28
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v32 = v23
	v33 = v24
	goto L20
L20:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	if v37 == int32(0) {
		v47 = v36
		v48 = v37
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v47 = v36
	v48 = v37
	goto L17
L22:
	;
	v40 = int32(1)
	if v36 == v37 {
		v32 = v32 + v40
		v33 = v33 + v40
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v53 = int32(_a_F_getInsertSelectQuery_1)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_getInsertSelectQuery[1])))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v57 == int32(0) {
		v76 = v56
		v77 = v57
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v77-v76 == int32(0) {
		v184 = l0
		goto L11
	} else {
		goto L33
	}
L26:
	;
	goto L25
L27:
	;
	if v56 != v57 {
		v76 = v56
		v77 = v57
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v61 = v52
	v62 = v53
	goto L29
L29:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v66 == int32(0) {
		v76 = v65
		v77 = v66
		goto L26
	} else {
		goto L31
	}
L30:
	;
	v76 = v65
	v77 = v66
	goto L26
L31:
	;
	v69 = int32(1)
	if v65 == v66 {
		v61 = v61 + v69
		v62 = v62 + v69
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	goto L13
L34:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v86 != int32(1) {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v91 != int32(63) {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v94+v95<<(uint(int32(2))%32)-int32(4))))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	if v102 != int32(1) {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101)+36))
	if v105 == int32(0) {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v108 != int32(67) {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v111 != int32(1) {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v105)+52))
	if v114 == int32(0) {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v117 < int32(2) {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v124 = int32(_a_F_getInsertSelectQuery_0)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_getInsertSelectQuery[0])))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v128 == int32(0) {
		v147 = v127
		v148 = v128
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v148-v147 != 0 {
		goto L7
	} else {
		goto L51
	}
L44:
	;
	goto L43
L45:
	;
	if v127 != v128 {
		v147 = v127
		v148 = v128
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v132 = v123
	v133 = v124
	goto L47
L47:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v137 == int32(0) {
		v147 = v136
		v148 = v137
		goto L44
	} else {
		goto L49
	}
L48:
	;
	v147 = v136
	v148 = v137
	goto L44
L49:
	;
	v140 = int32(1)
	if v136 == v137 {
		v132 = v132 + v140
		v133 = v133 + v140
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	v153 = int32(_a_F_getInsertSelectQuery_1)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_getInsertSelectQuery[1])))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	if v157 == int32(0) {
		v176 = v156
		v177 = v157
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v177-v176 != 0 {
		goto L7
	} else {
		goto L60
	}
L53:
	;
	goto L52
L54:
	;
	if v156 != v157 {
		v176 = v156
		v177 = v157
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v161 = v152
	v162 = v153
	goto L56
L56:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	if v166 == int32(0) {
		v176 = v165
		v177 = v166
		goto L53
	} else {
		goto L58
	}
L57:
	;
	v176 = v165
	v177 = v166
	goto L53
L58:
	;
	v169 = int32(1)
	if v165 == v166 {
		v161 = v161 + v169
		v162 = v162 + v169
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	if l1 == int32(0) {
		v184 = v105
		goto L11
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v101 + int32(36)
	v184 = v105
	goto L11
L62:
	;
	return int32(0)
L63:
	;
	F_errmsg_internal(m, int32(_a_F_getInsertSelectQuery_2), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_getInsertSelectQuery_3), int32(1118), int32(_a_F_getInsertSelectQuery_4))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errmsg_internal(m, int32(_a_F_getInsertSelectQuery_2), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L62
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_getInsertSelectQuery_3), int32(1121), int32(_a_F_getInsertSelectQuery_4))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L62
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errmsg_internal(m, int32(_a_F_getInsertSelectQuery_2), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L62
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_getInsertSelectQuery_3), int32(1127), int32(_a_F_getInsertSelectQuery_4))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L62
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_errmsg_internal(m, int32(_a_F_getInsertSelectQuery_5), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L62
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_getInsertSelectQuery_3), int32(1139), int32(_a_F_getInsertSelectQuery_4))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L62
	} else {
		goto L74
	}
L74:
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v52 int32
	_ = v52
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
	var v64 int32
	_ = v64
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
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
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
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
	var v555 int32
	_ = v555
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v957 int32
	_ = v957
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v979 int32
	_ = v979
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1035 int32
	_ = v1035
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1107 int32
	_ = v1107
	var v1112 int32
	_ = v1112
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1164 int32
	_ = v1164
	var v1174 int32
	_ = v1174
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1228 int32
	_ = v1228
	var v1238 int32
	_ = v1238
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1318 int32
	_ = v1318
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1542 int32
	_ = v1542
	var v1546 int32
	_ = v1546
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1593 int32
	_ = v1593
	var v1601 int32
	_ = v1601
	var v1606 int32
	_ = v1606
	var v1613 int32
	_ = v1613
	var v1622 int32
	_ = v1622
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1643 int32
	_ = v1643
	var v1651 int32
	_ = v1651
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1724 int32
	_ = v1724
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1757 int32
	_ = v1757
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1768 int32
	_ = v1768
	var v1773 int32
	_ = v1773
	var v1788 int32
	_ = v1788
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1833 int32
	_ = v1833
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1848 int32
	_ = v1848
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1866 int32
	_ = v1866
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1876 int32
	_ = v1876
	var v1879 int32
	_ = v1879
	var v1883 int32
	_ = v1883
	var v1897 int32
	_ = v1897
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1915 int32
	_ = v1915
	var v1920 int32
	_ = v1920
	var v1924 int32
	_ = v1924
	var v1931 int32
	_ = v1931
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1952 int32
	_ = v1952
	var v1954 int32
	_ = v1954
	var v1958 int32
	_ = v1958
	var v1960 int32
	_ = v1960
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	v5 = l4
	v9 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(304)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_get_query_def[0]))
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
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
	v24 = m.ExcPending
	if v24 != 0 {
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+45)))
	if v26 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v49 = int32(0)
	F_AcquireRewriteLocks(m, l0, v49, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L19
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
		goto L17
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v37 = v34 - int32(1)
	goto L8
L10:
	;
	if v25 != 0 {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v25 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v37 = int32(-1)
	goto L8
L14:
	;
	v48 = int32(0)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v48 = v33
	goto L7
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v40
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v45 = F_flatten_group_exprs(m, int32(0), l0, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v45
	v48 = v37
	goto L7
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+264)) = l1
	v56 = F_list_copy(m, l2)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v58 = F_lcons(m, v17+int32(184), v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v60 = int32(0)
	v64 = base.B2i32(l2 != v60) | base.B2i32(v48 != int32(1))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+296)) = uint8(v64)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+280)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v17)+272)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+268)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v17)+300)) = v60
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+298)) = uint16(v60)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+297)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+292)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v17)+288)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v17)+284)) = l5
	F_set_deparse_for_query(m, v17+int32(184), l0, l2)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v83 - int32(1) {
	case 0:
		goto L36
	case 1:
		goto L35
	case 2:
		goto L34
	case 3:
		goto L33
	case 4:
		goto L32
	case 5:
		goto L27
	case 6:
		goto L29
	default:
		goto L28
	}
L23:
	;
	m.G0 = v17 + int32(304)
	return
L24:
	;
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v1788 != 0 {
		goto L470
	} else {
		goto L471
	}
L25:
	;
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v719)+12))
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v1490 = int32(2)
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v1488+v1489<<(uint(v1490)%32)-int32(4))))
	v1496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+284)))
	if v1496&v1490 != 0 {
		goto L388
	} else {
		goto L389
	}
L26:
	;
	v1453 = int32(0)
	v1454 = base.B2i32(v1440 != v1453)
	v1456 = base.B2i32(v1448 != v1453)
	if v1448 == v1453 {
		v1475 = v1440
		v1483 = v1448
		v1484 = v1454
		v1486 = v1456
		goto L25
	} else {
		goto L383
	}
L27:
	;
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1361 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L4
	} else {
		goto L354
	}
L29:
	;
	F_appendStringInfoString(m, l1, int32(_a_F_get_query_def_0))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L4
	} else {
		goto L353
	}
L30:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v147)+80))
	F_get_values_def(m, v1339, v17+int32(264))
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L4
	} else {
		goto L352
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L4
	} else {
		goto L349
	}
L32:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v17)+264))
	F_get_with_clause(m, l0, v17+int32(264))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L4
	} else {
		goto L234
	}
L33:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v17)+264))
	F_get_with_clause(m, l0, v17+int32(264))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L4
	} else {
		goto L215
	}
L34:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v17)+264))
	F_get_with_clause(m, l0, v17+int32(264))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L4
	} else {
		goto L196
	}
L35:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v17)+264))
	F_get_with_clause(m, l0, v17+int32(264))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L4
	} else {
		goto L175
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+272)) = l3
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v17)+264))
	F_get_with_clause(m, l0, v17+int32(264))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+276)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+280)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v96 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_get_setop_query(m, v96, l0, v17+int32(264))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v17)+264))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+284)))
	if v102&int32(2) != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L24
L42:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v17)+292))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+292)) = v105 + int32(8)
	F_appendStringInfoChar(m, v101, int32(32))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v112 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+46)))
	if v268 != 0 {
		goto L90
	} else {
		goto L91
	}
L47:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v115 <= int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v17)+272))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v120 = int32(0)
	v126 = v120
	v131 = v120
	goto L49
L49:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v119+v131<<(uint(int32(2))%32))))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	switch v140 {
	case 0:
		goto L52
	default:
		goto L46
	case 5:
		goto L53
	}
L50:
	;
	if v147 == int32(0) {
		goto L46
	} else {
		goto L58
	}
L51:
	;
	v150 = v131 + int32(1)
	if v115 != v150 {
		v126 = v147
		v131 = v150
		goto L49
	} else {
		goto L57
	}
L52:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+125)))
	if v146 != 0 {
		goto L46
	} else {
		goto L56
	}
L53:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+125)))
	if v141 != int32(1) {
		goto L46
	} else {
		goto L54
	}
L54:
	;
	if v126 == int32(0) {
		v147 = v139
		goto L51
	} else {
		goto L55
	}
L55:
	;
	goto L46
L56:
	;
	v147 = v126
	goto L51
L57:
	;
	goto L50
L58:
	;
	v154 = int32(0)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v156 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v158 = v157
	goto L61
L60:
	;
	v158 = v154
	goto L61
L61:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v147)+8))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	if v160 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v162 = v161
	goto L64
L63:
	;
	v162 = v154
	goto L64
L64:
	;
	if v158 != v162 {
		goto L46
	} else {
		goto L65
	}
L65:
	;
	v175 = v154
	goto L66
L66:
	;
	v180 = int32(0)
	if v156 == v180 {
		v190 = v180
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L46
L68:
	;
	if v160 == int32(0) {
		goto L30
	} else {
		goto L71
	}
L69:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v184 <= v175 {
		v190 = int32(0)
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v190 = v186 + v175<<(uint(int32(2))%32)
	goto L68
L71:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v193 <= v175 {
		goto L30
	} else {
		goto L72
	}
L72:
	;
	if v190 == int32(0) {
		goto L30
	} else {
		goto L73
	}
L73:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	v200 = v197 + v175<<(uint(int32(2))%32)
	if v200 == int32(0) {
		goto L30
	} else {
		goto L74
	}
L74:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+26)))
	if v204 != 0 {
		goto L46
	} else {
		goto L75
	}
L75:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	if v118 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	if v228 == int32(0) {
		v247 = v227
		v248 = v228
		goto L82
	} else {
		goto L83
	}
L77:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	if v218 == int32(0) {
		goto L46
	} else {
		goto L80
	}
L78:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v209 <= v175 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v221 = v118 + int32(24) + v209<<(uint(int32(4))%32) + v175*int32(100)
	goto L76
L80:
	;
	v221 = v218
	goto L76
L81:
	;
	if v248-v247 == int32(0) {
		v175 = v175 + int32(1)
		goto L66
	} else {
		goto L89
	}
L82:
	;
	goto L81
L83:
	;
	if v227 != v228 {
		v247 = v227
		v248 = v228
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v232 = v221
	v233 = v206
	goto L85
L85:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+1)))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+1)))
	if v237 == int32(0) {
		v247 = v236
		v248 = v237
		goto L82
	} else {
		goto L87
	}
L86:
	;
	v247 = v236
	v248 = v237
	goto L82
L87:
	;
	v240 = int32(1)
	if v236 == v237 {
		v232 = v232 + v240
		v233 = v233 + v240
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	goto L67
L90:
	;
	v269 = int32(_a_F_get_query_def_1)
	goto L92
L91:
	;
	v269 = int32(_a_F_get_query_def_2)
	goto L92
L92:
	;
	F_appendStringInfoString(m, v101, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v272 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_get_target_list(m, v370, v17+int32(264))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L4
	} else {
		goto L113
	}
L95:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v275 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_appendStringInfoString(m, v101, int32(_a_F_get_query_def_3))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	F_appendStringInfoString(m, v101, int32(_a_F_get_query_def_4))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L112
	}
L99:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v281 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	F_appendStringInfoChar(m, v101, int32(41))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L111
	}
L101:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	if v285 <= int32(0) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v281)+12))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	F_appendStringInfoString(m, v101, int32(_a_F_get_query_def_5))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v298 = F_get_rule_sortgroupclause(m, v293, v294, int32(0), v17+int32(264))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	if v300 <= int32(1) {
		goto L100
	} else {
		goto L105
	}
L105:
	;
	v305 = int32(1)
	goto L106
L106:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v281)+12))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v317+v305<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v101, int32(_a_F_get_query_def_6))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L108
	}
L107:
	;
	goto L100
L108:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v330 = F_get_rule_sortgroupclause(m, v325, v326, int32(0), v17+int32(264))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v333 = v305 + int32(1)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	if v333 < v334 {
		v305 = v333
		goto L106
	} else {
		goto L110
	}
L110:
	;
	goto L107
L111:
	;
	goto L94
L112:
	;
	goto L94
L113:
	;
	F_get_from_clause(m, l0, int32(_a_F_get_query_def_7), v17+int32(264))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+8))
	if v381 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_8), int32(-8), int32(8), int32(1))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L4
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v397 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
	F_get_rule_expr(m, v391, v17+int32(264), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v555 != 0 {
		goto L153
	} else {
		goto L154
	}
L121:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v400 == int32(0) {
		goto L120
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_9), int32(-8), int32(8), int32(1))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L4
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	if v411 == int32(1) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	F_appendStringInfoString(m, v101, int32(_a_F_get_query_def_10))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L4
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+298)))
	v418 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+298)) = uint8(v418)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v420 != 0 {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L128
L130:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+298)) = uint8(v417)
	goto L120
L131:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if v422 <= int32(0) {
		goto L130
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v471 == int32(0) {
		goto L130
	} else {
		goto L143
	}
L134:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v420)+12))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	F_appendStringInfoString(m, v101, int32(_a_F_get_query_def_5))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_get_rule_groupingset(m, v426, v430, int32(1), v17+int32(264))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if v436 <= int32(1) {
		goto L130
	} else {
		goto L137
	}
L137:
	;
	v441 = int32(1)
	goto L138
L138:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v420)+12))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v453+v441<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v101, int32(_a_F_get_query_def_6))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L4
	} else {
		goto L140
	}
L139:
	;
	goto L130
L140:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_get_rule_groupingset(m, v457, v461, int32(1), v17+int32(264))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	v468 = v441 + int32(1)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if v468 < v469 {
		v441 = v468
		goto L138
	} else {
		goto L142
	}
L142:
	;
	goto L139
L143:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	if v475 <= int32(0) {
		goto L130
	} else {
		goto L144
	}
L144:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v471)+12))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)))
	F_appendStringInfoString(m, v101, int32(_a_F_get_query_def_5))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v488 = F_get_rule_sortgroupclause(m, v483, v484, int32(0), v17+int32(264))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	if v490 < int32(2) {
		goto L130
	} else {
		goto L147
	}
L147:
	;
	v495 = int32(1)
	goto L148
L148:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v471)+12))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v507+v495<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v101, int32(_a_F_get_query_def_6))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L150
	}
L149:
	;
	goto L130
L150:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v520 = F_get_rule_sortgroupclause(m, v515, v516, int32(0), v17+int32(264))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	v523 = v495 + int32(1)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	if v523 < v524 {
		v495 = v523
		goto L148
	} else {
		goto L152
	}
L152:
	;
	goto L149
L153:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_11), int32(-8), int32(8), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v570 == int32(0) {
		goto L24
	} else {
		goto L158
	}
L156:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	F_get_rule_expr(m, v564, v17+int32(264), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	if v573 <= int32(0) {
		goto L24
	} else {
		goto L159
	}
L159:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v17)+264))
	v577 = int32(0)
	v581 = v577
	v583 = v573
	v586 = v577
	goto L160
L160:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v570)+12))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v593+v581<<(uint(int32(2))%32))))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	if v598 != 0 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	goto L24
L162:
	;
	if v586 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L163:
	;
	v627 = v583
	v628 = v586
	goto L164
L164:
	;
	v630 = v581 + int32(1)
	if v630 < v627 {
		v581 = v630
		v583 = v627
		v586 = v628
		goto L160
	} else {
		goto L174
	}
L165:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	v612 = F_quote_identifier(m, v611)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L4
	} else {
		goto L171
	}
L166:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_12), int32(-8), int32(8), int32(1))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L4
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	F_appendStringInfoString(m, v576, v586)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L4
	} else {
		goto L170
	}
L169:
	;
	goto L165
L170:
	;
	goto L165
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v612
	F_appendStringInfo(m, v576, int32(_a_F_get_query_def_13), v17+int32(48))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_get_rule_windowspec(m, v597, v620, v17+int32(264))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	v627 = v626
	v628 = int32(_a_F_get_query_def_6)
	goto L164
L174:
	;
	goto L161
L175:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v637)+12))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v640 = int32(2)
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v638+v639<<(uint(v640)%32)-int32(4))))
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+284)))
	if v646&v640 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	F_appendStringInfoChar(m, v632, int32(32))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L4
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645)+20)))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v645)+16))
	v659 = F_generate_relation_name(m, v657, int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L4
	} else {
		goto L180
	}
L179:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v17)+292))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+292)) = v652 + int32(8)
	goto L178
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v659
	if v656 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v664 = int32(_a_F_get_query_def_5)
	goto L183
L182:
	;
	v664 = int32(_a_F_get_query_def_14)
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v664
	F_appendStringInfo(m, v632, int32(_a_F_get_query_def_15), v17-int32(-64))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_get_rte_alias(m, v645, v671, int32(0), v17+int32(264))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L4
	} else {
		goto L185
	}
L185:
	;
	F_appendStringInfoString(m, v632, int32(_a_F_get_query_def_16))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_get_update_query_targetlist_def(m, l0, v680, v17+int32(264), v645)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	F_get_from_clause(m, l0, int32(_a_F_get_query_def_7), v17+int32(264))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v690)+8))
	if v691 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_8), int32(-8), int32(8), int32(1))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L4
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v707 == int32(0) {
		goto L23
	} else {
		goto L194
	}
L192:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v700)+8))
	F_get_rule_expr(m, v701, v17+int32(264), int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L4
	} else {
		goto L193
	}
L193:
	;
	goto L191
L194:
	;
	F_get_returning_clause(m, l0, v17+int32(264))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L4
	} else {
		goto L195
	}
L195:
	;
	goto L23
L196:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v719 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v722 = int32(0)
	v1475 = v722
	v1483 = v722
	v1484 = v9
	v1486 = v9
	goto L25
L198:
	;
	goto L199
L199:
	;
	v724 = int32(0)
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v719)+4))
	if v726 <= v724 {
		v1440 = v724
		v1448 = v724
		goto L26
	} else {
		goto L200
	}
L200:
	;
	v729 = int32(0)
	if v729 < v726 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v732 = v726
	goto L203
L202:
	;
	v732 = v729
	goto L203
L203:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v719)+12))
	v734 = int32(0)
	v738 = v724
	v739 = v734
	v741 = v734
	v744 = v734
	v746 = v724
	goto L204
L204:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v733+v739<<(uint(int32(2))%32))))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v754)+12))
	switch v755 - int32(1) {
	case 0:
		goto L208
	default:
		v773 = v738
		v774 = v746
		goto L206
	case 4:
		goto L207
	}
L205:
	;
	v1440 = v773
	v1448 = v774
	goto L26
L206:
	;
	v776 = v739 + int32(1)
	if v776 != v732 {
		v738 = v773
		v739 = v776
		v741 = v773
		v744 = v774
		v746 = v774
		goto L204
	} else {
		goto L214
	}
L207:
	;
	if v741 != 0 {
		goto L31
	} else {
		goto L213
	}
L208:
	;
	if v744 == int32(0) {
		v773 = v741
		v774 = v754
		goto L206
	} else {
		goto L209
	}
L209:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L4
	} else {
		goto L210
	}
L210:
	;
	F_errmsg_internal(m, int32(_a_F_get_query_def_17), int32(0))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L4
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_19), int32(_a_F_get_query_def_20))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	v773 = v754
	v774 = v744
	goto L206
L214:
	;
	goto L205
L215:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v783)+12))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v786 = int32(2)
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v784+v785<<(uint(v786)%32)-int32(4))))
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+284)))
	if v792&v786 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	F_appendStringInfoChar(m, v778, int32(32))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L4
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v791)+20)))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v791)+16))
	v805 = F_generate_relation_name(m, v803, int32(0))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L4
	} else {
		goto L220
	}
L219:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v17)+292))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+292)) = v798 + int32(8)
	goto L218
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+132)) = v805
	if v802 != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v810 = int32(_a_F_get_query_def_5)
	goto L223
L222:
	;
	v810 = int32(_a_F_get_query_def_14)
	goto L223
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v810
	F_appendStringInfo(m, v778, int32(_a_F_get_query_def_21), v17+int32(128))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L4
	} else {
		goto L224
	}
L224:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_get_rte_alias(m, v791, v817, int32(0), v17+int32(264))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L4
	} else {
		goto L225
	}
L225:
	;
	F_get_from_clause(m, l0, int32(_a_F_get_query_def_22), v17+int32(264))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L4
	} else {
		goto L226
	}
L226:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v828)+8))
	if v829 != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_8), int32(-8), int32(8), int32(1))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L4
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v845 == int32(0) {
		goto L23
	} else {
		goto L232
	}
L230:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v838)+8))
	F_get_rule_expr(m, v839, v17+int32(264), int32(0))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L4
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	F_get_returning_clause(m, l0, v17+int32(264))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L4
	} else {
		goto L233
	}
L233:
	;
	goto L23
L234:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v857)+12))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v860 = int32(2)
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v858+v859<<(uint(v860)%32)-int32(4))))
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+284)))
	if v866&v860 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	F_appendStringInfoChar(m, v852, int32(32))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L4
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v865)+20)))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v865)+16))
	v879 = F_generate_relation_name(m, v877, int32(0))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L4
	} else {
		goto L239
	}
L238:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v17)+292))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+292)) = v872 + int32(8)
	goto L237
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+164)) = v879
	if v876 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v884 = int32(_a_F_get_query_def_5)
	goto L242
L241:
	;
	v884 = int32(_a_F_get_query_def_14)
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+160)) = v884
	F_appendStringInfo(m, v852, int32(_a_F_get_query_def_23), v17+int32(160))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L4
	} else {
		goto L243
	}
L243:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_get_rte_alias(m, v865, v891, int32(0), v17+int32(264))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L4
	} else {
		goto L244
	}
L244:
	;
	F_get_from_clause(m, l0, int32(_a_F_get_query_def_22), v17+int32(264))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L4
	} else {
		goto L245
	}
L245:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_24), int32(-8), int32(8), int32(2))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L4
	} else {
		goto L246
	}
L246:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	F_get_rule_expr(m, v910, v17+int32(264), int32(0))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L4
	} else {
		goto L247
	}
L247:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v916 == int32(0) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v1318 == int32(0) {
		goto L23
	} else {
		goto L347
	}
L249:
	;
	v919 = int32(_a_F_get_query_def_25)
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v916)+4))
	if v920 <= int32(0) {
		v957 = v919
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v968 = int32(0)
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v916)+4))
	if v969 <= v968 {
		goto L248
	} else {
		goto L261
	}
L251:
	;
	v923 = int32(0)
	if v923 < v920 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v926 = v920
	goto L254
L253:
	;
	v926 = v923
	goto L254
L254:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v916)+12))
	v931 = int32(0)
	goto L255
L255:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v927+v931<<(uint(int32(2))%32))))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v946)+4))
	if v947 != int32(1) {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v957 = int32(_a_F_get_query_def_26)
	goto L250
L257:
	;
	v951 = v931 + int32(1)
	if v926 != v951 {
		v931 = v951
		goto L255
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	goto L256
L260:
	;
	v957 = v919
	goto L250
L261:
	;
	v979 = v968
	goto L262
L262:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v916)+12))
	v987 = int32(2)
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v986+v979<<(uint(v987)%32))))
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_27), int32(-8), int32(8), v987)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L4
	} else {
		goto L264
	}
L263:
	;
	goto L248
L264:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v990)+4))
	switch v1000 {
	case 0:
		v1018 = int32(_a_F_get_query_def_28)
		goto L265
	case 1:
		goto L266
	case 2:
		goto L268
	default:
		goto L267
	}
L265:
	;
	F_appendStringInfoString(m, v852, v1018)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L4
	} else {
		goto L272
	}
L266:
	;
	v1018 = int32(_a_F_get_query_def_29)
	goto L265
L267:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L4
	} else {
		goto L269
	}
L268:
	;
	v1018 = v957
	goto L265
L269:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v990)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+144)) = v1005
	F_errmsg_internal(m, int32(_a_F_get_query_def_30), v17+int32(144))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L4
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_31), int32(_a_F_get_query_def_32))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L272:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v990)+16))
	if v1021 != 0 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_33), int32(-8), int32(8), int32(3))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L4
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_34), int32(-8), int32(8), int32(3))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L4
	} else {
		goto L278
	}
L276:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v990)+16))
	F_get_rule_expr(m, v1030, v17+int32(264), int32(0))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L4
	} else {
		goto L277
	}
L277:
	;
	goto L275
L278:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v990)+8))
	switch v1044 - int32(2) {
	case 0:
		goto L285
	case 1:
		goto L286
	case 2:
		goto L284
	default:
		goto L279
	case 5:
		goto L283
	}
L279:
	;
	v1301 = v979 + int32(1)
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v916)+4))
	if v1301 < v1302 {
		v979 = v1301
		goto L262
	} else {
		goto L346
	}
L280:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v990)+12))
	switch v1180 - int32(1) {
	case 0:
		goto L318
	case 1:
		v1184 = int32(_a_F_get_query_def_35)
		goto L317
	default:
		goto L316
	}
L281:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v990)+20))
	if v1159 == int32(0) {
		v1174 = v1154
		goto L280
	} else {
		goto L314
	}
L282:
	;
	v1107 = v1085
	v1112 = v1083
	goto L305
L283:
	;
	F_appendStringInfoString(m, v852, int32(_a_F_get_query_def_36))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L4
	} else {
		goto L304
	}
L284:
	;
	F_appendStringInfoString(m, v852, int32(_a_F_get_query_def_37))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L4
	} else {
		goto L303
	}
L285:
	;
	F_appendStringInfoString(m, v852, int32(_a_F_get_query_def_38))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L4
	} else {
		goto L301
	}
L286:
	;
	F_appendStringInfoString(m, v852, int32(_a_F_get_query_def_39))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L4
	} else {
		goto L287
	}
L287:
	;
	v1050 = int32(0)
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v990)+20))
	if v1051 == v1050 {
		v1174 = v1050
		goto L280
	} else {
		goto L288
	}
L288:
	;
	F_appendStringInfoString(m, v852, int32(_a_F_get_query_def_40))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L4
	} else {
		goto L289
	}
L289:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v990)+20))
	if v1057 == int32(0) {
		v1174 = v1050
		goto L280
	} else {
		goto L290
	}
L290:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+4))
	if v1060 <= int32(0) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1154 = v1050
	goto L281
L292:
	;
	goto L293
L293:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+12))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1063)))
	F_appendStringInfoString(m, v852, int32(_a_F_get_query_def_5))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L4
	} else {
		goto L294
	}
L294:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v865)+16))
	v1069 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1064)+8)))
	v1071 = F_get_attname(m, v1068, v1069, int32(0))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L4
	} else {
		goto L295
	}
L295:
	;
	v1073 = F_quote_identifier(m, v1071)
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L4
	} else {
		goto L296
	}
L296:
	;
	F_appendStringInfoString(m, v852, v1073)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L4
	} else {
		goto L297
	}
L297:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+4))
	v1081 = F_processIndirection(m, v1078, v17+int32(264))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L4
	} else {
		goto L298
	}
L298:
	;
	v1083 = F_lappend(m, int32(0), v1081)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L4
	} else {
		goto L299
	}
L299:
	;
	v1085 = int32(1)
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+4))
	if v1085 < v1086 {
		goto L282
	} else {
		goto L300
	}
L300:
	;
	v1154 = v1083
	goto L281
L301:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v990)+20))
	F_get_update_query_targetlist_def(m, l0, v1092, v17+int32(264), v865)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L4
	} else {
		goto L302
	}
L302:
	;
	goto L279
L303:
	;
	goto L279
L304:
	;
	goto L279
L305:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+12))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1117+v1107<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v852, int32(_a_F_get_query_def_6))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L4
	} else {
		goto L307
	}
L306:
	;
	v1154 = v1139
	goto L281
L307:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v865)+16))
	v1126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1121)+8)))
	v1128 = F_get_attname(m, v1125, v1126, int32(0))
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L4
	} else {
		goto L308
	}
L308:
	;
	v1130 = F_quote_identifier(m, v1128)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L4
	} else {
		goto L309
	}
L309:
	;
	F_appendStringInfoString(m, v852, v1130)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L4
	} else {
		goto L310
	}
L310:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+4))
	v1137 = F_processIndirection(m, v1134, v17+int32(264))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L4
	} else {
		goto L311
	}
L311:
	;
	v1139 = F_lappend(m, v1112, v1137)
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L4
	} else {
		goto L312
	}
L312:
	;
	v1142 = v1107 + int32(1)
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+4))
	if v1142 < v1143 {
		v1107 = v1142
		v1112 = v1139
		goto L305
	} else {
		goto L313
	}
L313:
	;
	goto L306
L314:
	;
	F_appendStringInfoChar(m, v852, int32(41))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L4
	} else {
		goto L315
	}
L315:
	;
	v1174 = v1154
	goto L280
L316:
	;
	if v1174 != 0 {
		goto L320
	} else {
		goto L321
	}
L317:
	;
	F_appendStringInfoString(m, v852, v1184)
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L4
	} else {
		goto L319
	}
L318:
	;
	v1184 = int32(_a_F_get_query_def_41)
	goto L317
L319:
	;
	goto L316
L320:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_42), int32(-8), int32(8), int32(4))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L4
	} else {
		goto L323
	}
L321:
	;
	goto L322
L322:
	;
	F_appendStringInfoString(m, v852, int32(_a_F_get_query_def_43))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L4
	} else {
		goto L345
	}
L323:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+4))
	if v1196 <= int32(0) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	F_appendStringInfoChar(m, v852, int32(41))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L4
	} else {
		goto L344
	}
L325:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+12))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1199)))
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v17)+264))
	F_appendStringInfoString(m, v1201, int32(_a_F_get_query_def_5))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L4
	} else {
		goto L326
	}
L326:
	;
	if v1200 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+4))
	if v1221 < int32(2) {
		goto L324
	} else {
		goto L333
	}
L328:
	;
	F_get_rule_expr(m, v1200, v17+int32(264), int32(0))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L4
	} else {
		goto L332
	}
L329:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1200)))
	if v1207 != int32(6) {
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v1213 = F_get_variable(m, v1200, int32(1), v17+int32(264))
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L4
	} else {
		goto L331
	}
L331:
	;
	goto L327
L332:
	;
	goto L327
L333:
	;
	v1228 = int32(1)
	goto L334
L334:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+12))
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1238+v1228<<(uint(int32(2))%32))))
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v17)+264))
	F_appendStringInfoString(m, v1243, int32(_a_F_get_query_def_6))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L4
	} else {
		goto L336
	}
L335:
	;
	goto L324
L336:
	;
	if v1242 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	v1263 = v1228 + int32(1)
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+4))
	if v1263 < v1264 {
		v1228 = v1263
		goto L334
	} else {
		goto L343
	}
L338:
	;
	F_get_rule_expr(m, v1242, v17+int32(264), int32(0))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L4
	} else {
		goto L342
	}
L339:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1242)))
	if v1249 != int32(6) {
		goto L338
	} else {
		goto L340
	}
L340:
	;
	v1255 = F_get_variable(m, v1242, int32(1), v17+int32(264))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L4
	} else {
		goto L341
	}
L341:
	;
	goto L337
L342:
	;
	goto L337
L343:
	;
	goto L335
L344:
	;
	goto L279
L345:
	;
	goto L279
L346:
	;
	goto L263
L347:
	;
	F_get_returning_clause(m, l0, v17+int32(264))
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L4
	} else {
		goto L348
	}
L348:
	;
	goto L23
L349:
	;
	F_errmsg_internal(m, int32(_a_F_get_query_def_44), int32(0))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L4
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_45), int32(_a_F_get_query_def_20))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L4
	} else {
		goto L351
	}
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L352:
	;
	goto L24
L353:
	;
	goto L23
L354:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v1351
	F_errmsg_internal(m, int32(_a_F_get_query_def_46), v17)
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L4
	} else {
		goto L355
	}
L355:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_47), int32(_a_F_get_query_def_48))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L4
	} else {
		goto L356
	}
L356:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L357:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L4
	} else {
		goto L380
	}
L358:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1361)))
	if v1364 != int32(222) {
		goto L357
	} else {
		goto L359
	}
L359:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v17)+264))
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_5), int32(0), int32(8), int32(1))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L4
	} else {
		goto L360
	}
L360:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1361)+4))
	v1377 = F_quote_identifier(m, v1376)
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L4
	} else {
		goto L361
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+176)) = v1377
	F_appendStringInfo(m, v1367, int32(_a_F_get_query_def_49), v17+int32(176))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L4
	} else {
		goto L362
	}
L362:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1361)+8))
	if v1385 == int32(0) {
		goto L23
	} else {
		goto L363
	}
L363:
	;
	F_appendStringInfoString(m, v1367, int32(_a_F_get_query_def_6))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L4
	} else {
		goto L364
	}
L364:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1361)+8))
	F_appendStringInfoChar(m, v1367, int32(39))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L4
	} else {
		goto L365
	}
L365:
	;
	v1397 = v1391
	goto L366
L366:
	;
	v1409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1397))))
	v1410 = base.I32_extend8_s(v1409)
	if v1409 != int32(39) {
		goto L369
	} else {
		goto L370
	}
L368:
	;
	F_appendStringInfoChar(m, v1367, v1410)
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L4
	} else {
		goto L379
	}
L369:
	;
	if v1409 != int32(92) {
		goto L372
	} else {
		goto L373
	}
L370:
	;
	goto L371
L371:
	;
	F_appendStringInfoChar(m, v1367, v1410)
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L4
	} else {
		goto L378
	}
L372:
	;
	if v1409 != 0 {
		goto L368
	} else {
		goto L375
	}
L373:
	;
	goto L374
L374:
	;
	v1419 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_query_def[1])))
	if v1419 != 0 {
		goto L368
	} else {
		goto L377
	}
L375:
	;
	F_appendStringInfoChar(m, v1367, int32(39))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L4
	} else {
		goto L376
	}
L376:
	;
	goto L23
L377:
	;
	goto L371
L378:
	;
	goto L368
L379:
	;
	v1397 = v1397 + int32(1)
	goto L366
L380:
	;
	F_errmsg_internal(m, int32(_a_F_get_query_def_50), int32(0))
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L4
	} else {
		goto L381
	}
L381:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_51), int32(_a_F_get_query_def_52))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L4
	} else {
		goto L382
	}
L382:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L383:
	;
	if v1440 == int32(0) {
		v1475 = v1440
		v1483 = v1448
		v1484 = v1454
		v1486 = v1456
		goto L25
	} else {
		goto L384
	}
L384:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L4
	} else {
		goto L385
	}
L385:
	;
	F_errmsg_internal(m, int32(_a_F_get_query_def_53), int32(0))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L4
	} else {
		goto L386
	}
L386:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_54), int32(_a_F_get_query_def_20))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L4
	} else {
		goto L387
	}
L387:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L388:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v17)+292))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+292)) = v1499 + int32(8)
	F_appendStringInfoChar(m, v714, int32(32))
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L4
	} else {
		goto L391
	}
L389:
	;
	goto L390
L390:
	;
	v1506 = int32(0)
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1495)+16))
	v1509 = F_generate_relation_name(m, v1507, v1506)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L4
	} else {
		goto L392
	}
L391:
	;
	goto L390
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v1509
	F_appendStringInfo(m, v714, int32(_a_F_get_query_def_55), v17+int32(112))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L4
	} else {
		goto L393
	}
L393:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_get_rte_alias(m, v1495, v1517, int32(1), v17+int32(264))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L4
	} else {
		goto L394
	}
L394:
	;
	F_appendStringInfoChar(m, v714, int32(32))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L4
	} else {
		goto L395
	}
L395:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v1526 == int32(0) {
		v1613 = v1506
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	switch v1622 - int32(1) {
	case 0:
		goto L419
	case 1:
		v1626 = int32(_a_F_get_query_def_56)
		goto L418
	default:
		goto L417
	}
L397:
	;
	F_appendStringInfoChar(m, v714, int32(40))
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L4
	} else {
		goto L398
	}
L398:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v1532 == int32(0) {
		v1613 = v1506
		goto L396
	} else {
		goto L399
	}
L399:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1532)+4))
	if int32(0) < v1535 {
		goto L400
	} else {
		goto L401
	}
L400:
	;
	v1542 = int32(0)
	v1546 = v1506
	v1551 = int32(_a_F_get_query_def_5)
	goto L403
L401:
	;
	v1593 = v1506
	goto L402
L402:
	;
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v1601 == int32(0) {
		v1613 = v1593
		goto L396
	} else {
		goto L415
	}
L403:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v1532)+12))
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1554+v1542<<(uint(int32(2))%32))))
	v1559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1558)+26)))
	if v1559 == int32(0) {
		goto L405
	} else {
		goto L406
	}
L404:
	;
	v1593 = v1581
	goto L402
L405:
	;
	F_appendStringInfoString(m, v714, v1551)
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L4
	} else {
		goto L408
	}
L406:
	;
	v1581 = v1546
	v1582 = v1551
	goto L407
L407:
	;
	v1584 = v1542 + int32(1)
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1532)+4))
	if v1584 < v1585 {
		v1542 = v1584
		v1546 = v1581
		v1551 = v1582
		goto L403
	} else {
		goto L414
	}
L408:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1495)+16))
	v1565 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1558)+8)))
	v1567 = F_get_attname(m, v1564, v1565, int32(0))
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L4
	} else {
		goto L409
	}
L409:
	;
	v1569 = F_quote_identifier(m, v1567)
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L4
	} else {
		goto L410
	}
L410:
	;
	F_appendStringInfoString(m, v714, v1569)
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L4
	} else {
		goto L411
	}
L411:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1558)+4))
	v1577 = F_processIndirection(m, v1574, v17+int32(264))
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L4
	} else {
		goto L412
	}
L412:
	;
	v1579 = F_lappend(m, v1546, v1577)
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L4
	} else {
		goto L413
	}
L413:
	;
	v1581 = v1579
	v1582 = int32(_a_F_get_query_def_6)
	goto L407
L414:
	;
	goto L404
L415:
	;
	F_appendStringInfoString(m, v714, int32(_a_F_get_query_def_57))
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L4
	} else {
		goto L416
	}
L416:
	;
	v1613 = v1593
	goto L396
L417:
	;
	if v1486 != 0 {
		goto L422
	} else {
		goto L423
	}
L418:
	;
	F_appendStringInfoString(m, v714, v1626)
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L4
	} else {
		goto L420
	}
L419:
	;
	v1626 = int32(_a_F_get_query_def_58)
	goto L418
L420:
	;
	goto L417
L421:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v1663 == int32(0) {
		goto L438
	} else {
		goto L439
	}
L422:
	;
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1483)+36))
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v17)+268))
	v1632 = int32(0)
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v17)+284))
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v17)+288))
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v17)+292))
	F_get_query_def(m, v1630, v714, v1631, v1632, v1632, v1634, v1635, v1636)
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L4
	} else {
		goto L425
	}
L423:
	;
	goto L424
L424:
	;
	if v1484 != 0 {
		goto L426
	} else {
		goto L427
	}
L425:
	;
	goto L421
L426:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v1475)+80))
	F_get_values_def(m, v1639, v17+int32(264))
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L4
	} else {
		goto L429
	}
L427:
	;
	goto L428
L428:
	;
	if v1613 != 0 {
		goto L430
	} else {
		goto L431
	}
L429:
	;
	goto L421
L430:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_59), int32(-8), int32(8), int32(2))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L4
	} else {
		goto L433
	}
L431:
	;
	goto L432
L432:
	;
	F_appendStringInfoString(m, v714, int32(_a_F_get_query_def_60))
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L4
	} else {
		goto L436
	}
L433:
	;
	F_get_rule_list_toplevel(m, v1613, v17+int32(264), int32(0))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L4
	} else {
		goto L434
	}
L434:
	;
	F_appendStringInfoChar(m, v714, int32(41))
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L4
	} else {
		goto L435
	}
L435:
	;
	goto L421
L436:
	;
	goto L421
L437:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L4
	} else {
		goto L467
	}
L438:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v1751 == int32(0) {
		goto L23
	} else {
		goto L465
	}
L439:
	;
	F_appendStringInfoString(m, v714, int32(_a_F_get_query_def_61))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L4
	} else {
		goto L440
	}
L440:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+8))
	if v1669 != 0 {
		goto L442
	} else {
		goto L443
	}
L441:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+4))
	if v1719 == int32(1) {
		goto L456
	} else {
		goto L457
	}
L442:
	;
	F_appendStringInfoChar(m, v714, int32(40))
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L4
	} else {
		goto L445
	}
L443:
	;
	goto L444
L444:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+16))
	if v1703 == int32(0) {
		goto L441
	} else {
		goto L451
	}
L445:
	;
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+8))
	F_get_rule_expr(m, v1673, v17+int32(264), int32(0))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L4
	} else {
		goto L446
	}
L446:
	;
	F_appendStringInfoChar(m, v714, int32(41))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L4
	} else {
		goto L447
	}
L447:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+12))
	if v1682 == int32(0) {
		goto L441
	} else {
		goto L448
	}
L448:
	;
	v1685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+296)))
	v1686 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+296)) = uint8(v1686)
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_8), int32(-8), int32(8), int32(1))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L4
	} else {
		goto L449
	}
L449:
	;
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+12))
	F_get_rule_expr(m, v1696, v17+int32(264), int32(0))
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L4
	} else {
		goto L450
	}
L450:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+296)) = uint8(v1685)
	goto L441
L451:
	;
	v1706 = F_get_constraint_name(m, v1703)
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L4
	} else {
		goto L452
	}
L452:
	;
	if v1706 == int32(0) {
		goto L437
	} else {
		goto L453
	}
L453:
	;
	v1710 = F_quote_identifier(m, v1706)
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L4
	} else {
		goto L454
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v1710
	F_appendStringInfo(m, v714, int32(_a_F_get_query_def_62), v17+int32(96))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L4
	} else {
		goto L455
	}
L455:
	;
	goto L441
L456:
	;
	F_appendStringInfoString(m, v714, int32(_a_F_get_query_def_63))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L4
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	F_appendStringInfoString(m, v714, int32(_a_F_get_query_def_64))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L4
	} else {
		goto L460
	}
L459:
	;
	goto L438
L460:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+20))
	F_get_update_query_targetlist_def(m, l0, v1728, v17+int32(264), v1495)
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L4
	} else {
		goto L461
	}
L461:
	;
	v1733 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+24))
	if v1733 == int32(0) {
		goto L438
	} else {
		goto L462
	}
L462:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_8), int32(-8), int32(8), int32(1))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L4
	} else {
		goto L463
	}
L463:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+24))
	F_get_rule_expr(m, v1744, v17+int32(264), int32(0))
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L4
	} else {
		goto L464
	}
L464:
	;
	goto L438
L465:
	;
	F_get_returning_clause(m, l0, v17+int32(264))
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L4
	} else {
		goto L466
	}
L466:
	;
	goto L23
L467:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v1762
	F_errmsg_internal(m, int32(_a_F_get_query_def_65), v17+int32(80))
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L4
	} else {
		goto L468
	}
L468:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_66), int32(_a_F_get_query_def_20))
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L4
	} else {
		goto L469
	}
L469:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L470:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_67), int32(-8), int32(8), int32(1))
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L4
	} else {
		goto L473
	}
L471:
	;
	goto L472
L472:
	;
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1805 != 0 {
		goto L475
	} else {
		goto L476
	}
L473:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_get_rule_orderby(m, v1797, v1798, base.B2i32(v96 != int32(0)), v17+int32(264))
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L4
	} else {
		goto L474
	}
L474:
	;
	goto L472
L475:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_68), int32(-8), int32(8), int32(0))
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L4
	} else {
		goto L478
	}
L476:
	;
	goto L477
L477:
	;
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v1820 == int32(0) {
		goto L480
	} else {
		goto L481
	}
L478:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	F_get_rule_expr(m, v1814, v17+int32(264), int32(0))
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L4
	} else {
		goto L479
	}
L479:
	;
	goto L477
L480:
	;
	v1873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	if v1873 != int32(1) {
		goto L23
	} else {
		goto L496
	}
L481:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v1823 == int32(1) {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_69), int32(-8), int32(8), int32(0))
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L4
	} else {
		goto L485
	}
L483:
	;
	goto L484
L484:
	;
	F_appendContextKeyword(m, v17+int32(264), int32(_a_F_get_query_def_70), int32(-8), int32(8), int32(0))
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L4
	} else {
		goto L490
	}
L485:
	;
	F_appendStringInfoChar(m, v87, int32(40))
	mBase = m.M
	v1836 = m.ExcPending
	if v1836 != 0 {
		goto L4
	} else {
		goto L486
	}
L486:
	;
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_get_rule_expr(m, v1837, v17+int32(264), int32(0))
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L4
	} else {
		goto L487
	}
L487:
	;
	F_appendStringInfoChar(m, v87, int32(41))
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
		goto L4
	} else {
		goto L488
	}
L488:
	;
	F_appendStringInfoString(m, v87, int32(_a_F_get_query_def_71))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L4
	} else {
		goto L489
	}
L489:
	;
	goto L480
L490:
	;
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1857)))
	if v1858 != int32(7) {
		goto L491
	} else {
		goto L492
	}
L491:
	;
	F_get_rule_expr(m, v1857, v17+int32(264), int32(0))
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L4
	} else {
		goto L495
	}
L492:
	;
	v1861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1857)+24)))
	if v1861 != int32(1) {
		goto L491
	} else {
		goto L493
	}
L493:
	;
	F_appendStringInfoString(m, v87, int32(_a_F_get_query_def_72))
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L4
	} else {
		goto L494
	}
L494:
	;
	goto L480
L495:
	;
	goto L480
L496:
	;
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v1876 == int32(0) {
		goto L23
	} else {
		goto L497
	}
L497:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v1876)+4))
	if v1879 <= int32(0) {
		goto L23
	} else {
		goto L498
	}
L498:
	;
	v1883 = int32(0)
	goto L499
L499:
	;
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1876)+12))
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(v1897+v1883<<(uint(int32(2))%32))))
	v1902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1901)+16)))
	if v1902 != 0 {
		goto L501
	} else {
		goto L502
	}
L500:
	;
	goto L23
L501:
	;
	v1963 = v1883 + int32(1)
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v1876)+4))
	if v1963 < v1964 {
		v1883 = v1963
		goto L499
	} else {
		goto L518
	}
L502:
	;
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v1901)+8))
	switch v1904 {
	case 0:
		goto L508
	case 1:
		v1924 = int32(_a_F_get_query_def_73)
		goto L504
	case 2:
		goto L507
	case 3:
		goto L506
	case 4:
		goto L505
	default:
		goto L503
	}
L503:
	;
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v17)+268))
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1933)+12))
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1934)))
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1935)+4))
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1936)+12))
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1901)+4))
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1937+v1938<<(uint(int32(2))%32)-int32(4))))
	v1945 = F_quote_identifier(m, v1944)
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		goto L4
	} else {
		goto L513
	}
L504:
	;
	F_appendContextKeyword(m, v17+int32(264), v1924, int32(-8), int32(8), int32(0))
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L4
	} else {
		goto L512
	}
L505:
	;
	v1924 = int32(_a_F_get_query_def_74)
	goto L504
L506:
	;
	v1924 = int32(_a_F_get_query_def_75)
	goto L504
L507:
	;
	v1924 = int32(_a_F_get_query_def_76)
	goto L504
L508:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L4
	} else {
		goto L509
	}
L509:
	;
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(v1901)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v1909
	F_errmsg_internal(m, int32(_a_F_get_query_def_77), v17+int32(32))
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L4
	} else {
		goto L510
	}
L510:
	;
	F_errfinish(m, int32(_a_F_get_query_def_18), int32(_a_F_get_query_def_78), int32(_a_F_get_query_def_79))
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L4
	} else {
		goto L511
	}
L511:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L512:
	;
	goto L503
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v1945
	F_appendStringInfo(m, v87, int32(_a_F_get_query_def_80), v17+int32(16))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L4
	} else {
		goto L514
	}
L514:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v1901)+12))
	switch v1954 - int32(1) {
	case 0:
		goto L516
	case 1:
		v1958 = int32(_a_F_get_query_def_81)
		goto L515
	default:
		goto L501
	}
L515:
	;
	F_appendStringInfoString(m, v87, v1958)
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L4
	} else {
		goto L517
	}
L516:
	;
	v1958 = int32(_a_F_get_query_def_82)
	goto L515
L517:
	;
	goto L501
L518:
	;
	goto L500
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
	var v32 int32
	_ = v32
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
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
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
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
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
	v32 = v5
	goto L10
L8:
	;
	v69 = v5
	goto L9
L9:
	;
	if v69 != 0 {
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
		v59 = v32
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v69 = v59
	goto L9
L12:
	;
	v61 = v27 + int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v61 < v62 {
		v27 = v61
		v32 = v59
		goto L10
	} else {
		goto L18
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v50 != int32(22) {
		v59 = v32
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v53 != int32(5) {
		v59 = v32
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v56 = F_lappend(m, v32, v49)
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
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v79 = v78
	goto L21
L20:
	;
	v79 = v5
	goto L21
L21:
	;
	v89 = v69
	v91 = v79
	v94 = l1 + int32(4)
	goto L2
L22:
	;
	v89 = v5
	v91 = v5
	v94 = v16
	goto L2
L23:
	;
	v103 = int32(0)
	v107 = int32(_a_F_get_update_query_targetlist_def_0)
	v110 = v91
	v114 = v5
	v115 = v5
	goto L24
L24:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117+v114<<(uint(int32(2))%32))))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+26)))
	if v122 != 0 {
		v375 = v103
		v379 = v107
		v382 = v110
		v387 = v115
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L1
L26:
	;
	v390 = v114 + int32(1)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v390 < v391 {
		v103 = v375
		v107 = v379
		v110 = v382
		v114 = v390
		v115 = v387
		goto L24
	} else {
		goto L110
	}
L27:
	;
	F_appendStringInfoString(m, v17, v107)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	if v110 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v343 = int32(*(*int16)(unsafe.Add(mBase, uint32(v121)+8)))
	v345 = F_get_attname(m, v342, v343, int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L16
	} else {
		goto L96
	}
L30:
	;
	v335 = v110
	v336 = v103
	v340 = v115
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
	v335 = v110
	v336 = v103
	v340 = v115
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
		goto L53
	} else {
		goto L54
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
		goto L50
	}
L41:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v132)+20))
	if v155 != int32(2) {
		goto L47
	} else {
		goto L48
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
	v167 = v132
	goto L36
L48:
	;
	goto L49
L49:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v159 = v158
	goto L40
L50:
	;
	goto L39
L51:
	;
	if v213 == int32(0) {
		v335 = v110
		v336 = v127
		v340 = v115
		goto L29
	} else {
		goto L72
	}
L52:
	;
	goto L51
L53:
	;
	v175 = v167
	goto L56
L54:
	;
	goto L55
L55:
	;
	v213 = int32(0)
	goto L52
L56:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	switch v176 - int32(15) {
	case 0:
		goto L64
	default:
		v213 = v175
		goto L52
	case 12:
		goto L63
	case 13:
		goto L62
	case 14:
		goto L61
	case 15:
		goto L60
	case 40:
		goto L59
	}
L57:
	;
	goto L55
L58:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v210 != 0 {
		v175 = v210
		goto L56
	} else {
		goto L71
	}
L59:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
	if v204 != int32(2) {
		v213 = v175
		goto L52
	} else {
		goto L70
	}
L60:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	if v199 != int32(2) {
		v213 = v175
		goto L52
	} else {
		goto L69
	}
L61:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v175)+24))
	if v194 != int32(2) {
		v213 = v175
		goto L52
	} else {
		goto L68
	}
L62:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	if v189 != int32(2) {
		v213 = v175
		goto L52
	} else {
		goto L67
	}
L63:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
	if v184 != int32(2) {
		v213 = v175
		goto L52
	} else {
		goto L66
	}
L64:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	if v179 != int32(2) {
		v213 = v175
		goto L52
	} else {
		goto L65
	}
L65:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v175)+28))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	v209 = v183
	goto L58
L66:
	;
	v209 = v175 + int32(4)
	goto L58
L67:
	;
	v209 = v175 + int32(4)
	goto L58
L68:
	;
	v209 = v175 + int32(4)
	goto L58
L69:
	;
	v209 = v175 + int32(4)
	goto L58
L70:
	;
	v209 = v175 + int32(4)
	goto L58
L71:
	;
	goto L57
L72:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	if v216 != int32(8) {
		v335 = v110
		v336 = v127
		v340 = v115
		goto L29
	} else {
		goto L73
	}
L73:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	if v219 != int32(3) {
		v335 = v110
		v336 = v127
		v340 = v115
		goto L29
	} else {
		goto L74
	}
L74:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+20))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+76))
	v227 = int32(0)
	if v226 == v227 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L16
	} else {
		goto L92
	}
L76:
	;
	v316 = int32(0)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v237 <= int32(0) {
		v301 = v227
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v316 = v301
	goto L75
L80:
	;
	v240 = int32(0)
	if v240 < v237 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v243 = v237
	goto L83
L82:
	;
	v243 = v240
	goto L83
L83:
	;
	v244 = int32(1)
	if v237 == v244 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v243&v244 == int32(0) {
		v301 = v282
		goto L79
	} else {
		goto L91
	}
L85:
	;
	v248 = int32(0)
	v282 = v248
	v283 = v248
	goto L84
L86:
	;
	goto L87
L87:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
	v253 = int32(0)
	v256 = v253
	v257 = v253
	v258 = v227
	goto L88
L88:
	;
	v263 = int32(2)
	v265 = v252 + v257<<(uint(v263)%32)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+26)))
	v268 = int32(1)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+26)))
	v275 = v256 + (v267 ^ v268) + (v272 ^ v268)
	v277 = v257 + v263
	v279 = v258 + v263
	if v279 != v243&int32(2147483646) {
		v256 = v275
		v257 = v277
		v258 = v279
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v282 = v275
	v283 = v277
	goto L84
L90:
	;
	goto L89
L91:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291+v283<<(uint(int32(2))%32))))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+26)))
	v301 = v282 + (v296 ^ int32(1))
	goto L79
L92:
	;
	v321 = v110 + int32(4)
	if base.Ui32(v321) < base.Ui32(v222+v223<<(uint(int32(2))%32)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v327 = v321
	goto L95
L94:
	;
	v327 = int32(0)
	goto L95
L95:
	;
	v335 = v327
	v336 = v224
	v340 = v316
	goto L29
L96:
	;
	v347 = F_quote_identifier(m, v345)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L16
	} else {
		goto L97
	}
L97:
	;
	F_appendStringInfoString(m, v17, v347)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L16
	} else {
		goto L98
	}
L98:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v352 = F_processIndirection(m, v351, l2)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L16
	} else {
		goto L99
	}
L99:
	;
	if v336 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_update_query_targetlist_def_1))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L16
	} else {
		goto L108
	}
L101:
	;
	v365 = v352
	v366 = v340
	goto L100
L102:
	;
	goto L103
L103:
	;
	v358 = v340 - int32(1)
	if int32(0) < v358 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v375 = v336
	v379 = int32(_a_F_get_update_query_targetlist_def_2)
	v382 = v335
	v387 = v358
	goto L26
L105:
	;
	goto L106
L106:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L16
	} else {
		goto L107
	}
L107:
	;
	v365 = v336
	v366 = v358
	goto L100
L108:
	;
	v370 = int32(0)
	F_get_rule_expr(m, v365, l2, v370)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L16
	} else {
		goto L109
	}
L109:
	;
	v375 = v370
	v379 = int32(_a_F_get_update_query_targetlist_def_2)
	v382 = v335
	v387 = v366
	goto L26
L110:
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
