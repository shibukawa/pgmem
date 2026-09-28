package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_regexp_instr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int64
	_ = v11
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v105 int64
	_ = v105
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	v2 = int32(0)
	v11 = int64(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v22 = F_pg_detoast_datum_packed(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int64(0)
		} else {
			v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(6) <= v24 {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v28 = F_pg_detoast_datum_packed(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
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
							return int64(0)
						} else {
							v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
							if v67 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v191 = m.ExcPending
								if v191 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v194 = m.ExcPending
									if v194 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
										F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
											mBase = m.M
											v206 = m.ExcPending
											if v206 != 0 {
												return int64(0)
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
								v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v75 = int32(0)
								v76 = base.B2i32(v61 != v75)
								v79 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v74, v76, v75, v75)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int64(0)
								} else {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
									if v81 < v60 {
										v105 = v11
									} else {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
										if v83 < v61 {
											v105 = v11
										} else {
											v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
											v97 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v61-v76+v83*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
											if v97 < int32(0) {
												v105 = v11
											} else {
												v105 = base.I64_extend_i32_s(v97 + int32(1))
											}
										}
									}
									m.G0 = v14 + int32(80)
									return v105
								}
							}
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						if v38 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v38
									*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_regexp_instr_4)
									F_errmsg(m, int32(_a_F_regexp_instr_5), v14)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1219), int32(_a_F_regexp_instr_3))
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return int64(0)
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
									return int64(0)
								} else {
									v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
									if v67 == int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
												F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
													mBase = m.M
													v206 = m.ExcPending
													if v206 != 0 {
														return int64(0)
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
										v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v75 = int32(0)
										v76 = base.B2i32(v61 != v75)
										v79 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v74, v76, v75, v75)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int64(0)
										} else {
											v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
											if v81 < v60 {
												v105 = v11
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
												if v83 < v61 {
													v105 = v11
												} else {
													v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
													v97 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v61-v76+v83*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
													if v97 < int32(0) {
														v105 = v11
													} else {
														v105 = base.I64_extend_i32_s(v97 + int32(1))
													}
												}
											}
											m.G0 = v14 + int32(80)
											return v105
										}
									}
								}
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
								if v45 <= int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v134 = m.ExcPending
										if v134 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v45
											*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(_a_F_regexp_instr_7)
											F_errmsg(m, int32(_a_F_regexp_instr_5), v14+int32(32))
											mBase = m.M
											v142 = m.ExcPending
											if v142 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1228), int32(_a_F_regexp_instr_3))
												mBase = m.M
												v147 = m.ExcPending
												if v147 != 0 {
													return int64(0)
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
											return int64(0)
										} else {
											v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
											if v67 == int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v191 = m.ExcPending
												if v191 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v194 = m.ExcPending
													if v194 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
														F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
														mBase = m.M
														v201 = m.ExcPending
														if v201 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
															mBase = m.M
															v206 = m.ExcPending
															if v206 != 0 {
																return int64(0)
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
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v75 = int32(0)
												v76 = base.B2i32(v61 != v75)
												v79 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v74, v76, v75, v75)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int64(0)
												} else {
													v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
													if v81 < v60 {
														v105 = v11
													} else {
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
														if v83 < v61 {
															v105 = v11
														} else {
															v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
															v97 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v61-v76+v83*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
															if v97 < int32(0) {
																v105 = v11
															} else {
																v105 = base.I64_extend_i32_s(v97 + int32(1))
															}
														}
													}
													m.G0 = v14 + int32(80)
													return v105
												}
											}
										}
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
										if base.Ui32(int32(2)) <= base.Ui32(v50) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v151 = m.ExcPending
											if v151 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v50
													*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(_a_F_regexp_instr_8)
													F_errmsg(m, int32(_a_F_regexp_instr_5), v14+int32(48))
													mBase = m.M
													v162 = m.ExcPending
													if v162 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1237), int32(_a_F_regexp_instr_3))
														mBase = m.M
														v167 = m.ExcPending
														if v167 != 0 {
															return int64(0)
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
													return int64(0)
												} else {
													v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
													if v67 == int32(1) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v194 = m.ExcPending
															if v194 != 0 {
																return int64(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
																F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
																mBase = m.M
																v201 = m.ExcPending
																if v201 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
																	mBase = m.M
																	v206 = m.ExcPending
																	if v206 != 0 {
																		return int64(0)
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
														v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v75 = int32(0)
														v76 = base.B2i32(v61 != v75)
														v79 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v74, v76, v75, v75)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int64(0)
														} else {
															v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
															if v81 < v60 {
																v105 = v11
															} else {
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
																if v83 < v61 {
																	v105 = v11
																} else {
																	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																	v97 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v61-v76+v83*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
																	if v97 < int32(0) {
																		v105 = v11
																	} else {
																		v105 = base.I64_extend_i32_s(v97 + int32(1))
																	}
																}
															}
															m.G0 = v14 + int32(80)
															return v105
														}
													}
												}
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
												if v55 < int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v174 = m.ExcPending
														if v174 != 0 {
															return int64(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v55
															*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = int32(_a_F_regexp_instr_9)
															F_errmsg(m, int32(_a_F_regexp_instr_5), v14-int32(-64))
															mBase = m.M
															v182 = m.ExcPending
															if v182 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1246), int32(_a_F_regexp_instr_3))
																mBase = m.M
																v187 = m.ExcPending
																if v187 != 0 {
																	return int64(0)
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
														return int64(0)
													} else {
														v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
														if v67 == int32(1) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v191 = m.ExcPending
															if v191 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(50856066))
																mBase = m.M
																v194 = m.ExcPending
																if v194 != 0 {
																	return int64(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
																	F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
																		mBase = m.M
																		v206 = m.ExcPending
																		if v206 != 0 {
																			return int64(0)
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
															v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v75 = int32(0)
															v76 = base.B2i32(v61 != v75)
															v79 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v74, v76, v75, v75)
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return int64(0)
															} else {
																v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
																if v81 < v60 {
																	v105 = v11
																} else {
																	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
																	if v83 < v61 {
																		v105 = v11
																	} else {
																		v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																		v97 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v61-v76+v83*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
																		if v97 < int32(0) {
																			v105 = v11
																		} else {
																			v105 = base.I64_extend_i32_s(v97 + int32(1))
																		}
																	}
																}
																m.G0 = v14 + int32(80)
																return v105
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
						return int64(0)
					} else {
						v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
						if v67 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v191 = m.ExcPending
							if v191 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
									F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
									mBase = m.M
									v201 = m.ExcPending
									if v201 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
										mBase = m.M
										v206 = m.ExcPending
										if v206 != 0 {
											return int64(0)
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
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v75 = int32(0)
							v76 = base.B2i32(v61 != v75)
							v79 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v74, v76, v75, v75)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int64(0)
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
								if v81 < v60 {
									v105 = v11
								} else {
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
									if v83 < v61 {
										v105 = v11
									} else {
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v61-v76+v83*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
										if v97 < int32(0) {
											v105 = v11
										} else {
											v105 = base.I64_extend_i32_s(v97 + int32(1))
										}
									}
								}
								m.G0 = v14 + int32(80)
								return v105
							}
						}
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					if v38 <= int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v38
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_regexp_instr_4)
								F_errmsg(m, int32(_a_F_regexp_instr_5), v14)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1219), int32(_a_F_regexp_instr_3))
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return int64(0)
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
								return int64(0)
							} else {
								v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
								if v67 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v191 = m.ExcPending
									if v191 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
											F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
												mBase = m.M
												v206 = m.ExcPending
												if v206 != 0 {
													return int64(0)
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
									v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v75 = int32(0)
									v76 = base.B2i32(v61 != v75)
									v79 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v74, v76, v75, v75)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int64(0)
									} else {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
										if v81 < v60 {
											v105 = v11
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
											if v83 < v61 {
												v105 = v11
											} else {
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v61-v76+v83*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
												if v97 < int32(0) {
													v105 = v11
												} else {
													v105 = base.I64_extend_i32_s(v97 + int32(1))
												}
											}
										}
										m.G0 = v14 + int32(80)
										return v105
									}
								}
							}
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							if v45 <= int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v45
										*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(_a_F_regexp_instr_7)
										F_errmsg(m, int32(_a_F_regexp_instr_5), v14+int32(32))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1228), int32(_a_F_regexp_instr_3))
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return int64(0)
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
										return int64(0)
									} else {
										v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
										if v67 == int32(1) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v191 = m.ExcPending
											if v191 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v194 = m.ExcPending
												if v194 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
													F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
													mBase = m.M
													v201 = m.ExcPending
													if v201 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
														mBase = m.M
														v206 = m.ExcPending
														if v206 != 0 {
															return int64(0)
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
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v75 = int32(0)
											v76 = base.B2i32(v61 != v75)
											v79 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v74, v76, v75, v75)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int64(0)
											} else {
												v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
												if v81 < v60 {
													v105 = v11
												} else {
													v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
													if v83 < v61 {
														v105 = v11
													} else {
														v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v61-v76+v83*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
														if v97 < int32(0) {
															v105 = v11
														} else {
															v105 = base.I64_extend_i32_s(v97 + int32(1))
														}
													}
												}
												m.G0 = v14 + int32(80)
												return v105
											}
										}
									}
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
									if base.Ui32(int32(2)) <= base.Ui32(v50) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v154 = m.ExcPending
											if v154 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v50
												*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(_a_F_regexp_instr_8)
												F_errmsg(m, int32(_a_F_regexp_instr_5), v14+int32(48))
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1237), int32(_a_F_regexp_instr_3))
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return int64(0)
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
												return int64(0)
											} else {
												v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
												if v67 == int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v191 = m.ExcPending
													if v191 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v194 = m.ExcPending
														if v194 != 0 {
															return int64(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
															F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
															mBase = m.M
															v201 = m.ExcPending
															if v201 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
																mBase = m.M
																v206 = m.ExcPending
																if v206 != 0 {
																	return int64(0)
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
													v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v75 = int32(0)
													v76 = base.B2i32(v61 != v75)
													v79 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v74, v76, v75, v75)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int64(0)
													} else {
														v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
														if v81 < v60 {
															v105 = v11
														} else {
															v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
															if v83 < v61 {
																v105 = v11
															} else {
																v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																v97 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v61-v76+v83*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
																if v97 < int32(0) {
																	v105 = v11
																} else {
																	v105 = base.I64_extend_i32_s(v97 + int32(1))
																}
															}
														}
														m.G0 = v14 + int32(80)
														return v105
													}
												}
											}
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
											if v55 < int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v174 = m.ExcPending
													if v174 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v55
														*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = int32(_a_F_regexp_instr_9)
														F_errmsg(m, int32(_a_F_regexp_instr_5), v14-int32(-64))
														mBase = m.M
														v182 = m.ExcPending
														if v182 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1246), int32(_a_F_regexp_instr_3))
															mBase = m.M
															v187 = m.ExcPending
															if v187 != 0 {
																return int64(0)
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
													return int64(0)
												} else {
													v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
													if v67 == int32(1) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v194 = m.ExcPending
															if v194 != 0 {
																return int64(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_regexp_instr_0)
																F_errmsg(m, int32(_a_F_regexp_instr_1), v14+int32(16))
																mBase = m.M
																v201 = m.ExcPending
																if v201 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_regexp_instr_2), int32(1257), int32(_a_F_regexp_instr_3))
																	mBase = m.M
																	v206 = m.ExcPending
																	if v206 != 0 {
																		return int64(0)
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
														v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v75 = int32(0)
														v76 = base.B2i32(v61 != v75)
														v79 = F_setup_regexp_matches(m, v17, v22, v64, v59-v70, v74, v76, v75, v75)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int64(0)
														} else {
															v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
															if v81 < v60 {
																v105 = v11
															} else {
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
																if v83 < v61 {
																	v105 = v11
																} else {
																	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																	v97 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v61-v76+v83*(v60-int32(1)))<<(uint(int32(3))%32)+v62<<(uint(int32(2))%32))))
																	if v97 < int32(0) {
																		v105 = v11
																	} else {
																		v105 = base.I64_extend_i32_s(v97 + int32(1))
																	}
																}
															}
															m.G0 = v14 + int32(80)
															return v105
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
func F_regexp_split_to_array_no_flags(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_split_to_array(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
