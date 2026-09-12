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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = F_pg_detoast_datum_packed(m, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(6) <= v23 {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
				v27 = F_pg_detoast_datum_packed(m, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
					v30 = v29
					v31 = v27
					v32 = int32(1)
					if base.I32_extend16_s(v30) < int32(3) {
						v61 = v2
						v62 = v2
						v63 = v32
						v64 = v32
						F_parse_re_flags(m, v13+int32(72), v31)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)))
							if v69 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v201 = m.ExcPending
								if v201 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v204 = m.ExcPending
									if v204 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(654830)
										F_errmsg(m, int32(241295), v13+int32(16))
										mBase = m.M
										v211 = m.ExcPending
										if v211 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(484510), int32(1257), int32(202395))
											mBase = m.M
											v216 = m.ExcPending
											if v216 != 0 {
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
								v72 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v72)
								v74 = int32(0)
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v84 = F_setup_regexp_matches(m, v16, v21, v13+int32(72), v64-v72, v79, base.B2i32(v61 != v74), v74, v74)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
									if v86 < v63 {
										v114 = v74
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
										if v88 < v61 {
											v114 = v74
										} else {
											v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
											v92 = int32(1)
											v96 = v61 - v92
											if base.Ui32(v96) <= base.Ui32(v61) {
												v99 = v96
											} else {
												v99 = int32(0)
											}
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v91+((v88*(v63-v92)+v99)<<(uint(int32(3))%32)|v62<<(uint(int32(2))%32)))))
											if v107 < int32(0) {
												v110 = int32(-1)
											} else {
												v110 = v107
											}
											v114 = v110 + int32(1)
										}
									}
									m.G0 = v13 + int32(80)
									return v114
								}
							}
						}
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v37 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v37
									*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(81266)
									F_errmsg(m, int32(477227), v13)
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(484510), int32(1219), int32(202395))
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
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
							if v30&int32(65535) == int32(3) {
								v61 = v2
								v62 = v2
								v63 = v32
								v64 = v37
								F_parse_re_flags(m, v13+int32(72), v31)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)))
									if v69 == int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v204 = m.ExcPending
											if v204 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(654830)
												F_errmsg(m, int32(241295), v13+int32(16))
												mBase = m.M
												v211 = m.ExcPending
												if v211 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(484510), int32(1257), int32(202395))
													mBase = m.M
													v216 = m.ExcPending
													if v216 != 0 {
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
										v72 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v72)
										v74 = int32(0)
										v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v84 = F_setup_regexp_matches(m, v16, v21, v13+int32(72), v64-v72, v79, base.B2i32(v61 != v74), v74, v74)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return int32(0)
										} else {
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
											if v86 < v63 {
												v114 = v74
											} else {
												v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
												if v88 < v61 {
													v114 = v74
												} else {
													v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
													v92 = int32(1)
													v96 = v61 - v92
													if base.Ui32(v96) <= base.Ui32(v61) {
														v99 = v96
													} else {
														v99 = int32(0)
													}
													v107 = *(*int32)(unsafe.Add(mBase, uint32(v91+((v88*(v63-v92)+v99)<<(uint(int32(3))%32)|v62<<(uint(int32(2))%32)))))
													if v107 < int32(0) {
														v110 = int32(-1)
													} else {
														v110 = v107
													}
													v114 = v110 + int32(1)
												}
											}
											m.G0 = v13 + int32(80)
											return v114
										}
									}
								}
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v44 <= int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v141 = m.ExcPending
									if v141 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v44
											*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(279483)
											F_errmsg(m, int32(477227), v13+int32(32))
											mBase = m.M
											v152 = m.ExcPending
											if v152 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(484510), int32(1228), int32(202395))
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
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
									if base.Ui32(v30&int32(65535)) < base.Ui32(int32(5)) {
										v61 = v2
										v62 = v2
										v63 = v44
										v64 = v37
										F_parse_re_flags(m, v13+int32(72), v31)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return int32(0)
										} else {
											v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)))
											if v69 == int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(654830)
														F_errmsg(m, int32(241295), v13+int32(16))
														mBase = m.M
														v211 = m.ExcPending
														if v211 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(484510), int32(1257), int32(202395))
															mBase = m.M
															v216 = m.ExcPending
															if v216 != 0 {
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
												v72 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v72)
												v74 = int32(0)
												v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v84 = F_setup_regexp_matches(m, v16, v21, v13+int32(72), v64-v72, v79, base.B2i32(v61 != v74), v74, v74)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int32(0)
												} else {
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
													if v86 < v63 {
														v114 = v74
													} else {
														v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
														if v88 < v61 {
															v114 = v74
														} else {
															v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
															v92 = int32(1)
															v96 = v61 - v92
															if base.Ui32(v96) <= base.Ui32(v61) {
																v99 = v96
															} else {
																v99 = int32(0)
															}
															v107 = *(*int32)(unsafe.Add(mBase, uint32(v91+((v88*(v63-v92)+v99)<<(uint(int32(3))%32)|v62<<(uint(int32(2))%32)))))
															if v107 < int32(0) {
																v110 = int32(-1)
															} else {
																v110 = v107
															}
															v114 = v110 + int32(1)
														}
													}
													m.G0 = v13 + int32(80)
													return v114
												}
											}
										}
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										if base.Ui32(int32(2)) <= base.Ui32(v51) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v51
													*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(241161)
													F_errmsg(m, int32(477227), v13+int32(48))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(484510), int32(1237), int32(202395))
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
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
											if base.Ui32(v30&int32(65535)) < base.Ui32(int32(7)) {
												v61 = v2
												v62 = v51
												v63 = v44
												v64 = v37
												F_parse_re_flags(m, v13+int32(72), v31)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)))
													if v69 == int32(1) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v201 = m.ExcPending
														if v201 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v204 = m.ExcPending
															if v204 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(654830)
																F_errmsg(m, int32(241295), v13+int32(16))
																mBase = m.M
																v211 = m.ExcPending
																if v211 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(484510), int32(1257), int32(202395))
																	mBase = m.M
																	v216 = m.ExcPending
																	if v216 != 0 {
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
														v72 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v72)
														v74 = int32(0)
														v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v84 = F_setup_regexp_matches(m, v16, v21, v13+int32(72), v64-v72, v79, base.B2i32(v61 != v74), v74, v74)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int32(0)
														} else {
															v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
															if v86 < v63 {
																v114 = v74
															} else {
																v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
																if v88 < v61 {
																	v114 = v74
																} else {
																	v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
																	v92 = int32(1)
																	v96 = v61 - v92
																	if base.Ui32(v96) <= base.Ui32(v61) {
																		v99 = v96
																	} else {
																		v99 = int32(0)
																	}
																	v107 = *(*int32)(unsafe.Add(mBase, uint32(v91+((v88*(v63-v92)+v99)<<(uint(int32(3))%32)|v62<<(uint(int32(2))%32)))))
																	if v107 < int32(0) {
																		v110 = int32(-1)
																	} else {
																		v110 = v107
																	}
																	v114 = v110 + int32(1)
																}
															}
															m.G0 = v13 + int32(80)
															return v114
														}
													}
												}
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												if v58 < int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v181 = m.ExcPending
													if v181 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v184 = m.ExcPending
														if v184 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v58
															*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = int32(202908)
															F_errmsg(m, int32(477227), v13-int32(-64))
															mBase = m.M
															v192 = m.ExcPending
															if v192 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(484510), int32(1246), int32(202395))
																mBase = m.M
																v197 = m.ExcPending
																if v197 != 0 {
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
													v61 = v58
													v62 = v51
													v63 = v44
													v64 = v37
													F_parse_re_flags(m, v13+int32(72), v31)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return int32(0)
													} else {
														v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)))
														if v69 == int32(1) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v201 = m.ExcPending
															if v201 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50856066))
																mBase = m.M
																v204 = m.ExcPending
																if v204 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(654830)
																	F_errmsg(m, int32(241295), v13+int32(16))
																	mBase = m.M
																	v211 = m.ExcPending
																	if v211 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(484510), int32(1257), int32(202395))
																		mBase = m.M
																		v216 = m.ExcPending
																		if v216 != 0 {
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
															v72 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v72)
															v74 = int32(0)
															v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v84 = F_setup_regexp_matches(m, v16, v21, v13+int32(72), v64-v72, v79, base.B2i32(v61 != v74), v74, v74)
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return int32(0)
															} else {
																v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
																if v86 < v63 {
																	v114 = v74
																} else {
																	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
																	if v88 < v61 {
																		v114 = v74
																	} else {
																		v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
																		v92 = int32(1)
																		v96 = v61 - v92
																		if base.Ui32(v96) <= base.Ui32(v61) {
																			v99 = v96
																		} else {
																			v99 = int32(0)
																		}
																		v107 = *(*int32)(unsafe.Add(mBase, uint32(v91+((v88*(v63-v92)+v99)<<(uint(int32(3))%32)|v62<<(uint(int32(2))%32)))))
																		if v107 < int32(0) {
																			v110 = int32(-1)
																		} else {
																			v110 = v107
																		}
																		v114 = v110 + int32(1)
																	}
																}
																m.G0 = v13 + int32(80)
																return v114
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
				v30 = v23
				v31 = v2
				v32 = int32(1)
				if base.I32_extend16_s(v30) < int32(3) {
					v61 = v2
					v62 = v2
					v63 = v32
					v64 = v32
					F_parse_re_flags(m, v13+int32(72), v31)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)))
						if v69 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v204 = m.ExcPending
								if v204 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(654830)
									F_errmsg(m, int32(241295), v13+int32(16))
									mBase = m.M
									v211 = m.ExcPending
									if v211 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(484510), int32(1257), int32(202395))
										mBase = m.M
										v216 = m.ExcPending
										if v216 != 0 {
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
							v72 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v72)
							v74 = int32(0)
							v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v84 = F_setup_regexp_matches(m, v16, v21, v13+int32(72), v64-v72, v79, base.B2i32(v61 != v74), v74, v74)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
								if v86 < v63 {
									v114 = v74
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
									if v88 < v61 {
										v114 = v74
									} else {
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
										v92 = int32(1)
										v96 = v61 - v92
										if base.Ui32(v96) <= base.Ui32(v61) {
											v99 = v96
										} else {
											v99 = int32(0)
										}
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v91+((v88*(v63-v92)+v99)<<(uint(int32(3))%32)|v62<<(uint(int32(2))%32)))))
										if v107 < int32(0) {
											v110 = int32(-1)
										} else {
											v110 = v107
										}
										v114 = v110 + int32(1)
									}
								}
								m.G0 = v13 + int32(80)
								return v114
							}
						}
					}
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v37 <= int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v37
								*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(81266)
								F_errmsg(m, int32(477227), v13)
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(484510), int32(1219), int32(202395))
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
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
						if v30&int32(65535) == int32(3) {
							v61 = v2
							v62 = v2
							v63 = v32
							v64 = v37
							F_parse_re_flags(m, v13+int32(72), v31)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)))
								if v69 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v201 = m.ExcPending
									if v201 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v204 = m.ExcPending
										if v204 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(654830)
											F_errmsg(m, int32(241295), v13+int32(16))
											mBase = m.M
											v211 = m.ExcPending
											if v211 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(484510), int32(1257), int32(202395))
												mBase = m.M
												v216 = m.ExcPending
												if v216 != 0 {
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
									v72 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v72)
									v74 = int32(0)
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v84 = F_setup_regexp_matches(m, v16, v21, v13+int32(72), v64-v72, v79, base.B2i32(v61 != v74), v74, v74)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
										if v86 < v63 {
											v114 = v74
										} else {
											v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
											if v88 < v61 {
												v114 = v74
											} else {
												v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
												v92 = int32(1)
												v96 = v61 - v92
												if base.Ui32(v96) <= base.Ui32(v61) {
													v99 = v96
												} else {
													v99 = int32(0)
												}
												v107 = *(*int32)(unsafe.Add(mBase, uint32(v91+((v88*(v63-v92)+v99)<<(uint(int32(3))%32)|v62<<(uint(int32(2))%32)))))
												if v107 < int32(0) {
													v110 = int32(-1)
												} else {
													v110 = v107
												}
												v114 = v110 + int32(1)
											}
										}
										m.G0 = v13 + int32(80)
										return v114
									}
								}
							}
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v44 <= int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v44
										*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(279483)
										F_errmsg(m, int32(477227), v13+int32(32))
										mBase = m.M
										v152 = m.ExcPending
										if v152 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(484510), int32(1228), int32(202395))
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
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
								if base.Ui32(v30&int32(65535)) < base.Ui32(int32(5)) {
									v61 = v2
									v62 = v2
									v63 = v44
									v64 = v37
									F_parse_re_flags(m, v13+int32(72), v31)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)))
										if v69 == int32(1) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v204 = m.ExcPending
												if v204 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(654830)
													F_errmsg(m, int32(241295), v13+int32(16))
													mBase = m.M
													v211 = m.ExcPending
													if v211 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(484510), int32(1257), int32(202395))
														mBase = m.M
														v216 = m.ExcPending
														if v216 != 0 {
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
											v72 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v72)
											v74 = int32(0)
											v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v84 = F_setup_regexp_matches(m, v16, v21, v13+int32(72), v64-v72, v79, base.B2i32(v61 != v74), v74, v74)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int32(0)
											} else {
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
												if v86 < v63 {
													v114 = v74
												} else {
													v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
													if v88 < v61 {
														v114 = v74
													} else {
														v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
														v92 = int32(1)
														v96 = v61 - v92
														if base.Ui32(v96) <= base.Ui32(v61) {
															v99 = v96
														} else {
															v99 = int32(0)
														}
														v107 = *(*int32)(unsafe.Add(mBase, uint32(v91+((v88*(v63-v92)+v99)<<(uint(int32(3))%32)|v62<<(uint(int32(2))%32)))))
														if v107 < int32(0) {
															v110 = int32(-1)
														} else {
															v110 = v107
														}
														v114 = v110 + int32(1)
													}
												}
												m.G0 = v13 + int32(80)
												return v114
											}
										}
									}
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									if base.Ui32(int32(2)) <= base.Ui32(v51) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v51
												*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(241161)
												F_errmsg(m, int32(477227), v13+int32(48))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(484510), int32(1237), int32(202395))
													mBase = m.M
													v177 = m.ExcPending
													if v177 != 0 {
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
										if base.Ui32(v30&int32(65535)) < base.Ui32(int32(7)) {
											v61 = v2
											v62 = v51
											v63 = v44
											v64 = v37
											F_parse_re_flags(m, v13+int32(72), v31)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)))
												if v69 == int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v201 = m.ExcPending
													if v201 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v204 = m.ExcPending
														if v204 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(654830)
															F_errmsg(m, int32(241295), v13+int32(16))
															mBase = m.M
															v211 = m.ExcPending
															if v211 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(484510), int32(1257), int32(202395))
																mBase = m.M
																v216 = m.ExcPending
																if v216 != 0 {
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
													v72 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v72)
													v74 = int32(0)
													v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v84 = F_setup_regexp_matches(m, v16, v21, v13+int32(72), v64-v72, v79, base.B2i32(v61 != v74), v74, v74)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int32(0)
													} else {
														v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
														if v86 < v63 {
															v114 = v74
														} else {
															v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
															if v88 < v61 {
																v114 = v74
															} else {
																v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
																v92 = int32(1)
																v96 = v61 - v92
																if base.Ui32(v96) <= base.Ui32(v61) {
																	v99 = v96
																} else {
																	v99 = int32(0)
																}
																v107 = *(*int32)(unsafe.Add(mBase, uint32(v91+((v88*(v63-v92)+v99)<<(uint(int32(3))%32)|v62<<(uint(int32(2))%32)))))
																if v107 < int32(0) {
																	v110 = int32(-1)
																} else {
																	v110 = v107
																}
																v114 = v110 + int32(1)
															}
														}
														m.G0 = v13 + int32(80)
														return v114
													}
												}
											}
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											if v58 < int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v181 = m.ExcPending
												if v181 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v184 = m.ExcPending
													if v184 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v58
														*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = int32(202908)
														F_errmsg(m, int32(477227), v13-int32(-64))
														mBase = m.M
														v192 = m.ExcPending
														if v192 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(484510), int32(1246), int32(202395))
															mBase = m.M
															v197 = m.ExcPending
															if v197 != 0 {
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
												v61 = v58
												v62 = v51
												v63 = v44
												v64 = v37
												F_parse_re_flags(m, v13+int32(72), v31)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)))
													if v69 == int32(1) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v201 = m.ExcPending
														if v201 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v204 = m.ExcPending
															if v204 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(654830)
																F_errmsg(m, int32(241295), v13+int32(16))
																mBase = m.M
																v211 = m.ExcPending
																if v211 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(484510), int32(1257), int32(202395))
																	mBase = m.M
																	v216 = m.ExcPending
																	if v216 != 0 {
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
														v72 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v72)
														v74 = int32(0)
														v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v84 = F_setup_regexp_matches(m, v16, v21, v13+int32(72), v64-v72, v79, base.B2i32(v61 != v74), v74, v74)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int32(0)
														} else {
															v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
															if v86 < v63 {
																v114 = v74
															} else {
																v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
																if v88 < v61 {
																	v114 = v74
																} else {
																	v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
																	v92 = int32(1)
																	v96 = v61 - v92
																	if base.Ui32(v96) <= base.Ui32(v61) {
																		v99 = v96
																	} else {
																		v99 = int32(0)
																	}
																	v107 = *(*int32)(unsafe.Add(mBase, uint32(v91+((v88*(v63-v92)+v99)<<(uint(int32(3))%32)|v62<<(uint(int32(2))%32)))))
																	if v107 < int32(0) {
																		v110 = int32(-1)
																	} else {
																		v110 = v107
																	}
																	v114 = v110 + int32(1)
																}
															}
															m.G0 = v13 + int32(80)
															return v114
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
