package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_regexp_instr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
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
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v22 = F_pg_detoast_datum_packed(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(6) <= v24 {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
				v28 = F_pg_detoast_datum_packed(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
					v31 = v30
					v32 = v28
					v33 = int32(1)
					if base.I32_extend16_s(v31) < int32(3) {
						v59 = v33
						v60 = v33
						v61 = v2
						v62 = v2
						v64 = v14 + int32(72)
						F_parse_re_flags(m, v64, v32)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
							if v67 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v193 = m.ExcPending
								if v193 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v196 = m.ExcPending
									if v196 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
										F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
										mBase = m.M
										v203 = m.ExcPending
										if v203 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
											mBase = m.M
											v208 = m.ExcPending
											if v208 != 0 {
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
								v70 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v70)
								v72 = int32(0)
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v77 = base.B2i32(v61 != v72)
								v80 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v75, v77, v72, v72)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
									if v82 < v60 {
										v106 = v72
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
										if v84 < v61 {
											v106 = v72
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
											v99 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v61-v77+v84*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
											if v99 < int32(0) {
												v102 = int32(-1)
											} else {
												v102 = v99
											}
											v106 = v102 + int32(1)
										}
									}
									m.G0 = v14 + int32(80)
									return v106
								}
							}
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v38 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v38
									*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_regexp_instr_4)
									F_errmsg(m, int32(_a_F_regexp_instr_5), v14)
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1219), int32(_a_F_regexp_instr_3))
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
							v42 = v31 & int32(_a_F_regexp_instr_6)
							if v42 == int32(3) {
								v59 = v38
								v60 = v33
								v61 = v2
								v62 = v2
								v64 = v14 + int32(72)
								F_parse_re_flags(m, v64, v32)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
									if v67 == int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v193 = m.ExcPending
										if v193 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v196 = m.ExcPending
											if v196 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
												F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
												mBase = m.M
												v203 = m.ExcPending
												if v203 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
													mBase = m.M
													v208 = m.ExcPending
													if v208 != 0 {
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
										v70 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v70)
										v72 = int32(0)
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v77 = base.B2i32(v61 != v72)
										v80 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v75, v77, v72, v72)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
											if v82 < v60 {
												v106 = v72
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
												if v84 < v61 {
													v106 = v72
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
													v99 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v61-v77+v84*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
													if v99 < int32(0) {
														v102 = int32(-1)
													} else {
														v102 = v99
													}
													v106 = v102 + int32(1)
												}
											}
											m.G0 = v14 + int32(80)
											return v106
										}
									}
								}
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v45 <= int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v45
											*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(_a_F_regexp_instr_7)
											F_errmsg(m, int32(_a_F_regexp_instr_5), v14+int32(32))
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1228), int32(_a_F_regexp_instr_3))
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
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
									if base.Ui32(v42) < base.Ui32(int32(5)) {
										v59 = v38
										v60 = v45
										v61 = v2
										v62 = v2
										v64 = v14 + int32(72)
										F_parse_re_flags(m, v64, v32)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
											if v67 == int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v193 = m.ExcPending
												if v193 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v196 = m.ExcPending
													if v196 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
														F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
														mBase = m.M
														v203 = m.ExcPending
														if v203 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
															mBase = m.M
															v208 = m.ExcPending
															if v208 != 0 {
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
												v70 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v70)
												v72 = int32(0)
												v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v77 = base.B2i32(v61 != v72)
												v80 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v75, v77, v72, v72)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
													if v82 < v60 {
														v106 = v72
													} else {
														v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
														if v84 < v61 {
															v106 = v72
														} else {
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
															v99 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v61-v77+v84*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
															if v99 < int32(0) {
																v102 = int32(-1)
															} else {
																v102 = v99
															}
															v106 = v102 + int32(1)
														}
													}
													m.G0 = v14 + int32(80)
													return v106
												}
											}
										}
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										if base.Ui32(int32(2)) <= base.Ui32(v50) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v153 = m.ExcPending
											if v153 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v156 = m.ExcPending
												if v156 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v50
													*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(_a_F_regexp_instr_8)
													F_errmsg(m, int32(_a_F_regexp_instr_5), v14+int32(48))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1237), int32(_a_F_regexp_instr_3))
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
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
											if base.Ui32(v42) < base.Ui32(int32(7)) {
												v59 = v38
												v60 = v45
												v61 = v2
												v62 = v50
												v64 = v14 + int32(72)
												F_parse_re_flags(m, v64, v32)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
													if v67 == int32(1) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v193 = m.ExcPending
														if v193 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v196 = m.ExcPending
															if v196 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
																F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
																mBase = m.M
																v203 = m.ExcPending
																if v203 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
																	mBase = m.M
																	v208 = m.ExcPending
																	if v208 != 0 {
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
														v70 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v70)
														v72 = int32(0)
														v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v77 = base.B2i32(v61 != v72)
														v80 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v75, v77, v72, v72)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
															if v82 < v60 {
																v106 = v72
															} else {
																v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
																if v84 < v61 {
																	v106 = v72
																} else {
																	v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
																	v99 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v61-v77+v84*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
																	if v99 < int32(0) {
																		v102 = int32(-1)
																	} else {
																		v102 = v99
																	}
																	v106 = v102 + int32(1)
																}
															}
															m.G0 = v14 + int32(80)
															return v106
														}
													}
												}
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												if v55 < int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v173 = m.ExcPending
													if v173 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v176 = m.ExcPending
														if v176 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v55
															*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = int32(_a_F_regexp_instr_9)
															F_errmsg(m, int32(_a_F_regexp_instr_5), v14-int32(-64))
															mBase = m.M
															v184 = m.ExcPending
															if v184 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1246), int32(_a_F_regexp_instr_3))
																mBase = m.M
																v189 = m.ExcPending
																if v189 != 0 {
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
													v59 = v38
													v60 = v45
													v61 = v55
													v62 = v50
													v64 = v14 + int32(72)
													F_parse_re_flags(m, v64, v32)
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return int32(0)
													} else {
														v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
														if v67 == int32(1) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v193 = m.ExcPending
															if v193 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50856066))
																mBase = m.M
																v196 = m.ExcPending
																if v196 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
																	F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
																	mBase = m.M
																	v203 = m.ExcPending
																	if v203 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
																		mBase = m.M
																		v208 = m.ExcPending
																		if v208 != 0 {
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
															v70 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v70)
															v72 = int32(0)
															v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v77 = base.B2i32(v61 != v72)
															v80 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v75, v77, v72, v72)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
																if v82 < v60 {
																	v106 = v72
																} else {
																	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
																	if v84 < v61 {
																		v106 = v72
																	} else {
																		v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
																		v99 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v61-v77+v84*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
																		if v99 < int32(0) {
																			v102 = int32(-1)
																		} else {
																			v102 = v99
																		}
																		v106 = v102 + int32(1)
																	}
																}
																m.G0 = v14 + int32(80)
																return v106
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
				v31 = v24
				v32 = v2
				v33 = int32(1)
				if base.I32_extend16_s(v31) < int32(3) {
					v59 = v33
					v60 = v33
					v61 = v2
					v62 = v2
					v64 = v14 + int32(72)
					F_parse_re_flags(m, v64, v32)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
						if v67 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v196 = m.ExcPending
								if v196 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
									F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
										mBase = m.M
										v208 = m.ExcPending
										if v208 != 0 {
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
							v70 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v70)
							v72 = int32(0)
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v77 = base.B2i32(v61 != v72)
							v80 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v75, v77, v72, v72)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
								if v82 < v60 {
									v106 = v72
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
									if v84 < v61 {
										v106 = v72
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v61-v77+v84*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
										if v99 < int32(0) {
											v102 = int32(-1)
										} else {
											v102 = v99
										}
										v106 = v102 + int32(1)
									}
								}
								m.G0 = v14 + int32(80)
								return v106
							}
						}
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v38 <= int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v38
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_regexp_instr_4)
								F_errmsg(m, int32(_a_F_regexp_instr_5), v14)
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1219), int32(_a_F_regexp_instr_3))
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
						v42 = v31 & int32(_a_F_regexp_instr_6)
						if v42 == int32(3) {
							v59 = v38
							v60 = v33
							v61 = v2
							v62 = v2
							v64 = v14 + int32(72)
							F_parse_re_flags(m, v64, v32)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
								if v67 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v193 = m.ExcPending
									if v193 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
											F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
											mBase = m.M
											v203 = m.ExcPending
											if v203 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
												mBase = m.M
												v208 = m.ExcPending
												if v208 != 0 {
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
									v70 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v70)
									v72 = int32(0)
									v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v77 = base.B2i32(v61 != v72)
									v80 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v75, v77, v72, v72)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
										if v82 < v60 {
											v106 = v72
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
											if v84 < v61 {
												v106 = v72
											} else {
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
												v99 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v61-v77+v84*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
												if v99 < int32(0) {
													v102 = int32(-1)
												} else {
													v102 = v99
												}
												v106 = v102 + int32(1)
											}
										}
										m.G0 = v14 + int32(80)
										return v106
									}
								}
							}
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v45 <= int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v45
										*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(_a_F_regexp_instr_7)
										F_errmsg(m, int32(_a_F_regexp_instr_5), v14+int32(32))
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1228), int32(_a_F_regexp_instr_3))
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
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
								if base.Ui32(v42) < base.Ui32(int32(5)) {
									v59 = v38
									v60 = v45
									v61 = v2
									v62 = v2
									v64 = v14 + int32(72)
									F_parse_re_flags(m, v64, v32)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
										if v67 == int32(1) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v193 = m.ExcPending
											if v193 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v196 = m.ExcPending
												if v196 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
													F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
													mBase = m.M
													v203 = m.ExcPending
													if v203 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
														mBase = m.M
														v208 = m.ExcPending
														if v208 != 0 {
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
											v70 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v70)
											v72 = int32(0)
											v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v77 = base.B2i32(v61 != v72)
											v80 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v75, v77, v72, v72)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
												if v82 < v60 {
													v106 = v72
												} else {
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
													if v84 < v61 {
														v106 = v72
													} else {
														v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
														v99 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v61-v77+v84*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
														if v99 < int32(0) {
															v102 = int32(-1)
														} else {
															v102 = v99
														}
														v106 = v102 + int32(1)
													}
												}
												m.G0 = v14 + int32(80)
												return v106
											}
										}
									}
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									if base.Ui32(int32(2)) <= base.Ui32(v50) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v156 = m.ExcPending
											if v156 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v50
												*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(_a_F_regexp_instr_8)
												F_errmsg(m, int32(_a_F_regexp_instr_5), v14+int32(48))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1237), int32(_a_F_regexp_instr_3))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
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
										if base.Ui32(v42) < base.Ui32(int32(7)) {
											v59 = v38
											v60 = v45
											v61 = v2
											v62 = v50
											v64 = v14 + int32(72)
											F_parse_re_flags(m, v64, v32)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
												if v67 == int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v193 = m.ExcPending
													if v193 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v196 = m.ExcPending
														if v196 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
															F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
															mBase = m.M
															v203 = m.ExcPending
															if v203 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
																mBase = m.M
																v208 = m.ExcPending
																if v208 != 0 {
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
													v70 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v70)
													v72 = int32(0)
													v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v77 = base.B2i32(v61 != v72)
													v80 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v75, v77, v72, v72)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
														if v82 < v60 {
															v106 = v72
														} else {
															v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
															if v84 < v61 {
																v106 = v72
															} else {
																v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
																v99 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v61-v77+v84*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
																if v99 < int32(0) {
																	v102 = int32(-1)
																} else {
																	v102 = v99
																}
																v106 = v102 + int32(1)
															}
														}
														m.G0 = v14 + int32(80)
														return v106
													}
												}
											}
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											if v55 < int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v173 = m.ExcPending
												if v173 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v55
														*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = int32(_a_F_regexp_instr_9)
														F_errmsg(m, int32(_a_F_regexp_instr_5), v14-int32(-64))
														mBase = m.M
														v184 = m.ExcPending
														if v184 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1246), int32(_a_F_regexp_instr_3))
															mBase = m.M
															v189 = m.ExcPending
															if v189 != 0 {
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
												v59 = v38
												v60 = v45
												v61 = v55
												v62 = v50
												v64 = v14 + int32(72)
												F_parse_re_flags(m, v64, v32)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
													if v67 == int32(1) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v193 = m.ExcPending
														if v193 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v196 = m.ExcPending
															if v196 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
																F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
																mBase = m.M
																v203 = m.ExcPending
																if v203 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
																	mBase = m.M
																	v208 = m.ExcPending
																	if v208 != 0 {
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
														v70 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v70)
														v72 = int32(0)
														v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v77 = base.B2i32(v61 != v72)
														v80 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v75, v77, v72, v72)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
															if v82 < v60 {
																v106 = v72
															} else {
																v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
																if v84 < v61 {
																	v106 = v72
																} else {
																	v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
																	v99 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v61-v77+v84*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
																	if v99 < int32(0) {
																		v102 = int32(-1)
																	} else {
																		v102 = v99
																	}
																	v106 = v102 + int32(1)
																}
															}
															m.G0 = v14 + int32(80)
															return v106
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
func F_regexp_split_to_array_no_flags(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_split_to_array(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
