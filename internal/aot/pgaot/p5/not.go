package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReportNotNullViolationError(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v20 != 0 {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
		v24 = F_build_attrmap_by_name_if_req(m, v15, v22, int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			if v24 != 0 {
				v27 = F_MakeTupleTableSlot(m, v22, int32(1649940))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = F_execute_attr_map_slot(m, v24, l1, v27)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						v31 = v29
						v32 = F_ExecGetInsertedCols(m, v20, l2)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							v34 = F_ExecGetUpdatedCols(m, v20, l2)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								v36 = F_bms_union(m, v32, v34)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
									v45 = v22
									v46 = v31
									v47 = v36
									v49 = v38
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+56))
									v51 = F_ExecBuildSlotValueDescription(m, v50, v46, v45, v47)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											F_errcode(m, int32(33575106))
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v15 + v16<<(uint(int32(4))%32) + l3*int32(100) - int32(76)
												*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v60 + int32(4)
												F_errmsg(m, int32(95938), v12+int32(16))
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return
												} else {
													if v51 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v12))) = v51
														F_errdetail(m, int32(630182), v12)
														mBase = m.M
														v78 = m.ExcPending
														if v78 != 0 {
															return
														} else {
															F_errtablecol(m, v14, l3)
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return
															} else {
																F_errfinish(m, int32(519072), int32(2219), int32(222906))
																mBase = m.M
																v85 = m.ExcPending
																if v85 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														F_errtablecol(m, v14, l3)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return
														} else {
															F_errfinish(m, int32(519072), int32(2219), int32(222906))
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
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
										}
									}
								}
							}
						}
					}
				}
			} else {
				v31 = l1
				v32 = F_ExecGetInsertedCols(m, v20, l2)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = F_ExecGetUpdatedCols(m, v20, l2)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v36 = F_bms_union(m, v32, v34)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
							v45 = v22
							v46 = v31
							v47 = v36
							v49 = v38
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+56))
							v51 = F_ExecBuildSlotValueDescription(m, v50, v46, v45, v47)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									F_errcode(m, int32(33575106))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v15 + v16<<(uint(int32(4))%32) + l3*int32(100) - int32(76)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v60 + int32(4)
										F_errmsg(m, int32(95938), v12+int32(16))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											if v51 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v51
												F_errdetail(m, int32(630182), v12)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													F_errtablecol(m, v14, l3)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return
													} else {
														F_errfinish(m, int32(519072), int32(2219), int32(222906))
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_errtablecol(m, v14, l3)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return
												} else {
													F_errfinish(m, int32(519072), int32(2219), int32(222906))
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
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
								}
							}
						}
					}
				}
			}
		}
	} else {
		v39 = F_ExecGetInsertedCols(m, l0, l2)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			v41 = F_ExecGetUpdatedCols(m, l0, l2)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				v43 = F_bms_union(m, v39, v41)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v45 = v15
					v46 = l1
					v47 = v43
					v49 = v14
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+56))
					v51 = F_ExecBuildSlotValueDescription(m, v50, v46, v45, v47)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_errcode(m, int32(33575106))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v15 + v16<<(uint(int32(4))%32) + l3*int32(100) - int32(76)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v60 + int32(4)
								F_errmsg(m, int32(95938), v12+int32(16))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									if v51 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = v51
										F_errdetail(m, int32(630182), v12)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											F_errtablecol(m, v14, l3)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												F_errfinish(m, int32(519072), int32(2219), int32(222906))
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										F_errtablecol(m, v14, l3)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return
										} else {
											F_errfinish(m, int32(519072), int32(2219), int32(222906))
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
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
						}
					}
				}
			}
		}
	}
}
