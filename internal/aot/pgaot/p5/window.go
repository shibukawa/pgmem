package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_find_window_functions_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 == v3 {
		v41 = v3
		m.G0 = v8 + int32(16)
		return v41
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v12 == int32(11) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.Ui32(v16) < base.Ui32(v15) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v50
					F_errmsg_internal(m, int32(53835), v8)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(519225), int32(252), int32(233213))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v15<<(uint(int32(2))%32))))
				v23 = F_lappend(m, v22, l0)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(int32(2))%32)))) = v23
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33 + int32(1)
					v41 = v3
					m.G0 = v8 + int32(16)
					return v41
				}
			}
		} else {
			v38 = F_expression_tree_walker_impl(m, l0, int32(856), l1)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v41 = v38
				m.G0 = v8 + int32(16)
				return v41
			}
		}
	}
}
func F_window_gettupleslot(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	if v10 != 0 {
		F_ProcessInterrupts(m)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if l1 < int64(0) {
				v145 = v4
				return v145
			} else {
				F_spool_tuples(m, v8, l1)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = *(*int64)(unsafe.Add(mBase, uint32(v8)+168))
					if v19 <= l1 {
						v145 = v4
						return v145
					} else {
						v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
						if v21 <= l1 {
							v23 = int32(4562096)
							v24 = *(*int32)(unsafe.Add(mBase, _consts[10]))
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
							*(*int32)(unsafe.Add(mBase, _consts[10])) = v27
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							F_tuplestore_select_read_pointer(m, v29, v30)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v34 = l1 - int64(1)
								v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
								if v35 < v34 {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
									v40 = F_tuplestore_skiptuples(m, v37, v34-v35, int32(1))
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										if v40 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												F_errmsg_internal(m, int32(383143), int32(0))
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(524101), int32(3165), int32(91169))
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v34
											v128 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
											v129 = int32(1)
											v131 = F_tuplestore_gettupleslot(m, v128, v129, v129, l2)
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int32(0)
											} else {
												if v131 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(383143), int32(0))
														mBase = m.M
														v156 = m.ExcPending
														if v156 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(524101), int32(3206), int32(91169))
															mBase = m.M
															v161 = m.ExcPending
															if v161 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v138 = int64(1)
													v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v138 + v139
													*(*int32)(unsafe.Add(mBase, _consts[10])) = v24
													v145 = int32(1)
													return v145
												}
											}
										}
									}
								} else {
									v46 = l1 + int64(1)
									if v46 < v35 {
										v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
										v51 = F_tuplestore_skiptuples(m, v48, v35-v46, int32(0))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											if v51 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(383143), int32(0))
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(524101), int32(3173), int32(91169))
														mBase = m.M
														v125 = m.ExcPending
														if v125 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v46
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
												v72 = F_tuplestore_gettupleslot(m, v69, int32(0), int32(1), l2)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													if v72 != 0 {
														v138 = int64(-1)
														v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v138 + v139
														*(*int32)(unsafe.Add(mBase, _consts[10])) = v24
														v145 = int32(1)
														return v145
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(383143), int32(0))
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(524101), int32(3200), int32(91169))
																mBase = m.M
																v86 = m.ExcPending
																if v86 != 0 {
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
									} else {
										if l1 == v35 {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
											v59 = F_tuplestore_advance(m, v57, int32(1))
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int32(0)
											} else {
												v61 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
												v63 = v61 + int64(1)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v63
												v65 = v63
												if v65 <= l1 {
													v128 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
													v129 = int32(1)
													v131 = F_tuplestore_gettupleslot(m, v128, v129, v129, l2)
													mBase = m.M
													v132 = m.ExcPending
													if v132 != 0 {
														return int32(0)
													} else {
														if v131 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(383143), int32(0))
																mBase = m.M
																v156 = m.ExcPending
																if v156 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(524101), int32(3206), int32(91169))
																	mBase = m.M
																	v161 = m.ExcPending
																	if v161 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v138 = int64(1)
															v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
															*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v138 + v139
															*(*int32)(unsafe.Add(mBase, _consts[10])) = v24
															v145 = int32(1)
															return v145
														}
													}
												} else {
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
													v72 = F_tuplestore_gettupleslot(m, v69, int32(0), int32(1), l2)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int32(0)
													} else {
														if v72 != 0 {
															v138 = int64(-1)
															v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
															*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v138 + v139
															*(*int32)(unsafe.Add(mBase, _consts[10])) = v24
															v145 = int32(1)
															return v145
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v77 = m.ExcPending
															if v77 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(383143), int32(0))
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(524101), int32(3200), int32(91169))
																	mBase = m.M
																	v86 = m.ExcPending
																	if v86 != 0 {
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
										} else {
											v65 = v35
											if v65 <= l1 {
												v128 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
												v129 = int32(1)
												v131 = F_tuplestore_gettupleslot(m, v128, v129, v129, l2)
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return int32(0)
												} else {
													if v131 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(383143), int32(0))
															mBase = m.M
															v156 = m.ExcPending
															if v156 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(524101), int32(3206), int32(91169))
																mBase = m.M
																v161 = m.ExcPending
																if v161 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v138 = int64(1)
														v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v138 + v139
														*(*int32)(unsafe.Add(mBase, _consts[10])) = v24
														v145 = int32(1)
														return v145
													}
												}
											} else {
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
												v72 = F_tuplestore_gettupleslot(m, v69, int32(0), int32(1), l2)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													if v72 != 0 {
														v138 = int64(-1)
														v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v138 + v139
														*(*int32)(unsafe.Add(mBase, _consts[10])) = v24
														v145 = int32(1)
														return v145
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(383143), int32(0))
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(524101), int32(3200), int32(91169))
																mBase = m.M
																v86 = m.ExcPending
																if v86 != 0 {
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
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(262709), int32(0))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(524101), int32(3151), int32(91169))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
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
	} else {
		if l1 < int64(0) {
			v145 = v4
			return v145
		} else {
			F_spool_tuples(m, v8, l1)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v8)+168))
				if v19 <= l1 {
					v145 = v4
					return v145
				} else {
					v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
					if v21 <= l1 {
						v23 = int32(4562096)
						v24 = *(*int32)(unsafe.Add(mBase, _consts[10]))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
						*(*int32)(unsafe.Add(mBase, _consts[10])) = v27
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						F_tuplestore_select_read_pointer(m, v29, v30)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v34 = l1 - int64(1)
							v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
							if v35 < v34 {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
								v40 = F_tuplestore_skiptuples(m, v37, v34-v35, int32(1))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									if v40 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(383143), int32(0))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(524101), int32(3165), int32(91169))
												mBase = m.M
												v112 = m.ExcPending
												if v112 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v34
										v128 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
										v129 = int32(1)
										v131 = F_tuplestore_gettupleslot(m, v128, v129, v129, l2)
										mBase = m.M
										v132 = m.ExcPending
										if v132 != 0 {
											return int32(0)
										} else {
											if v131 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v152 = m.ExcPending
												if v152 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(383143), int32(0))
													mBase = m.M
													v156 = m.ExcPending
													if v156 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(524101), int32(3206), int32(91169))
														mBase = m.M
														v161 = m.ExcPending
														if v161 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v138 = int64(1)
												v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v138 + v139
												*(*int32)(unsafe.Add(mBase, _consts[10])) = v24
												v145 = int32(1)
												return v145
											}
										}
									}
								}
							} else {
								v46 = l1 + int64(1)
								if v46 < v35 {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
									v51 = F_tuplestore_skiptuples(m, v48, v35-v46, int32(0))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										if v51 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int32(0)
											} else {
												F_errmsg_internal(m, int32(383143), int32(0))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(524101), int32(3173), int32(91169))
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v46
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
											v72 = F_tuplestore_gettupleslot(m, v69, int32(0), int32(1), l2)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												if v72 != 0 {
													v138 = int64(-1)
													v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v138 + v139
													*(*int32)(unsafe.Add(mBase, _consts[10])) = v24
													v145 = int32(1)
													return v145
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(383143), int32(0))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(524101), int32(3200), int32(91169))
															mBase = m.M
															v86 = m.ExcPending
															if v86 != 0 {
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
								} else {
									if l1 == v35 {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
										v59 = F_tuplestore_advance(m, v57, int32(1))
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v61 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
											v63 = v61 + int64(1)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v63
											v65 = v63
											if v65 <= l1 {
												v128 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
												v129 = int32(1)
												v131 = F_tuplestore_gettupleslot(m, v128, v129, v129, l2)
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return int32(0)
												} else {
													if v131 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(383143), int32(0))
															mBase = m.M
															v156 = m.ExcPending
															if v156 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(524101), int32(3206), int32(91169))
																mBase = m.M
																v161 = m.ExcPending
																if v161 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v138 = int64(1)
														v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v138 + v139
														*(*int32)(unsafe.Add(mBase, _consts[10])) = v24
														v145 = int32(1)
														return v145
													}
												}
											} else {
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
												v72 = F_tuplestore_gettupleslot(m, v69, int32(0), int32(1), l2)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													if v72 != 0 {
														v138 = int64(-1)
														v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v138 + v139
														*(*int32)(unsafe.Add(mBase, _consts[10])) = v24
														v145 = int32(1)
														return v145
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(383143), int32(0))
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(524101), int32(3200), int32(91169))
																mBase = m.M
																v86 = m.ExcPending
																if v86 != 0 {
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
									} else {
										v65 = v35
										if v65 <= l1 {
											v128 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
											v129 = int32(1)
											v131 = F_tuplestore_gettupleslot(m, v128, v129, v129, l2)
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int32(0)
											} else {
												if v131 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(383143), int32(0))
														mBase = m.M
														v156 = m.ExcPending
														if v156 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(524101), int32(3206), int32(91169))
															mBase = m.M
															v161 = m.ExcPending
															if v161 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v138 = int64(1)
													v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v138 + v139
													*(*int32)(unsafe.Add(mBase, _consts[10])) = v24
													v145 = int32(1)
													return v145
												}
											}
										} else {
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
											v72 = F_tuplestore_gettupleslot(m, v69, int32(0), int32(1), l2)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												if v72 != 0 {
													v138 = int64(-1)
													v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v138 + v139
													*(*int32)(unsafe.Add(mBase, _consts[10])) = v24
													v145 = int32(1)
													return v145
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(383143), int32(0))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(524101), int32(3200), int32(91169))
															mBase = m.M
															v86 = m.ExcPending
															if v86 != 0 {
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
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(262709), int32(0))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(524101), int32(3151), int32(91169))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
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
func F_window_lead(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = int32(1)
	v15 = F_WinGetFuncArgInPartition(m, v8, v9, v9, v6+int32(15), v6+int32(14))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v19 == int32(1) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			v25 = int32(0)
		} else {
			v25 = v15
		}
		m.G0 = v6 + int32(16)
		return v25
	}
}
func F_window_row_number(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)+176))
	F_WinSetMarkPosition(m, v3, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v12 = F_Int64GetDatum(m, v5+int64(1))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v12
		}
	}
}
