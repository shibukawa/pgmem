package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inner_int_inter(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = l0 + int32(16)
	v16 = F_ArrayGetNItemsSafe(m, v13, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 == int32(0) {
			v310 = F_construct_empty_array(m, int32(23))
			mBase = m.M
			v311 = m.ExcPending
			if v311 != 0 {
				return int32(0)
			} else {
				return v310
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v24 = l1 + int32(16)
			v25 = F_ArrayGetNItemsSafe(m, v22, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v25 == int32(0) {
					v310 = F_construct_empty_array(m, int32(23))
					mBase = m.M
					v311 = m.ExcPending
					if v311 != 0 {
						return int32(0)
					} else {
						return v310
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v30 = F_ArrayGetNItemsSafe(m, v29, v15)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v33 = F_ArrayGetNItemsSafe(m, v32, v24)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v35 == int32(0) {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v45 = (v38<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							} else {
								v45 = v35
							}
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							if v46 == int32(0) {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v56 = (v49<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							} else {
								v56 = v46
							}
							if v30 < v33 {
								v58 = v30
							} else {
								v58 = v33
							}
							if int32(0) < v58 {
								v64 = v58<<(uint(int32(2))%32) + int32(24)
								v65 = F_palloc0(m, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v65)+20)) = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v58
									*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = int32(23)
									*(*int64)(unsafe.Add(mBase, uint32(v65)+4)) = int64(1)
									*(*int32)(unsafe.Add(mBase, uint32(v65))) = v64 << (uint(int32(2)) % 32)
									v85 = v65
									v87 = v65 + int32(8)
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
									v95 = v85
									v96 = (v88<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									v97 = v87
									v98 = int32(0)
									if base.B2i32(v30 <= v98)|base.B2i32(v33 <= v98) == v98 {
										v107 = v95 + v96
										v108 = int32(0)
										v110 = v108
										v111 = v108
										v114 = int32(0)
										for {
											v122 = int32(2)
											v125 = *(*int32)(unsafe.Add(mBase, uint32(l0+v45+v110<<(uint(v122)%32))))
											v129 = *(*int32)(unsafe.Add(mBase, uint32(l1+v56+v111<<(uint(v122)%32))))
											if v125 < v129 {
												v154 = v110 + int32(1)
												v155 = v111
												v156 = v114
											} else {
												if v125 == v129 {
													if v114 != 0 {
														v139 = *(*int32)(unsafe.Add(mBase, uint32(v107+v114<<(uint(int32(2))%32)-int32(4))))
														if v139 == v125 {
															v147 = v114
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v107+v114<<(uint(int32(2))%32)))) = v125
															v147 = v114 + int32(1)
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v107+v114<<(uint(int32(2))%32)))) = v125
														v147 = v114 + int32(1)
													}
													v148 = int32(1)
													v154 = v110 + v148
													v155 = v111 + v148
													v156 = v147
												} else {
													v154 = v110
													v155 = v111 + int32(1)
													v156 = v114
												}
											}
											if base.B2i32(v154 < v30)&base.B2i32(v155 < v33) != 0 {
												v110 = v154
												v111 = v155
												v114 = v156
												continue
											} else {
												break
											}
											break
										}
										if v156 != 0 {
											if v156 <= int32(0) {
												v310 = F_construct_empty_array(m, int32(23))
												mBase = m.M
												v311 = m.ExcPending
												if v311 != 0 {
													return int32(0)
												} else {
													return v310
												}
											} else {
												v176 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
												v179 = F_ArrayGetNItemsSafe(m, v176, v95+int32(16))
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return int32(0)
												} else {
													if v179 == v156 {
														v286 = v95
														return v286
													} else {
														v182 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
														if v182 != 0 {
															v190 = v182
														} else {
															v183 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
															v190 = (v183<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														}
														v193 = v190 + v156<<(uint(int32(2))%32)
														v194 = F_repalloc(m, v95, v193)
														mBase = m.M
														v195 = m.ExcPending
														if v195 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v194))) = v193 << (uint(int32(2)) % 32)
															v199 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
															if v199 <= int32(0) {
																v286 = v194
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v194)+16)) = v156
																if v199 == int32(1) {
																	v286 = v194
																} else {
																	v206 = v194 + int32(16)
																	v207 = int32(1)
																	v208 = v199 - v207
																	v209 = int32(7)
																	v210 = v208 & v209
																	if base.Ui32(v209) <= base.Ui32(v199-int32(2)) {
																		v219 = v207
																		v222 = int32(0)
																		for {
																			v233 = v206 + v219<<(uint(int32(2))%32)
																			v234 = int64(4294967297)
																			*(*int64)(unsafe.Add(mBase, uint32(v233)+24)) = v234
																			*(*int64)(unsafe.Add(mBase, uint32(v233)+16)) = v234
																			*(*int64)(unsafe.Add(mBase, uint32(v233)+8)) = v234
																			*(*int64)(unsafe.Add(mBase, uint32(v233))) = v234
																			v242 = int32(8)
																			v243 = v219 + v242
																			v245 = v222 + v242
																			if v245 != v208&int32(-8) {
																				v219 = v243
																				v222 = v245
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v210 == int32(0) {
																			v286 = v194
																		} else {
																			v249 = v243
																			v262 = v249
																			v263 = int32(0)
																			for {
																				v277 = int32(1)
																				*(*int32)(unsafe.Add(mBase, uint32(v206+v262<<(uint(int32(2))%32)))) = v277
																				v282 = v263 + v277
																				if v282 != v210 {
																					v262 = v262 + v277
																					v263 = v282
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v286 = v194
																		}
																	} else {
																		v249 = v207
																		v262 = v249
																		v263 = int32(0)
																		for {
																			v277 = int32(1)
																			*(*int32)(unsafe.Add(mBase, uint32(v206+v262<<(uint(int32(2))%32)))) = v277
																			v282 = v263 + v277
																			if v282 != v210 {
																				v262 = v262 + v277
																				v263 = v282
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		v286 = v194
																	}
																}
															}
															return v286
														}
													}
												}
											}
										} else {
											F_pfree(m, v95)
											mBase = m.M
											v173 = m.ExcPending
											if v173 != 0 {
												return int32(0)
											} else {
												v310 = F_construct_empty_array(m, int32(23))
												mBase = m.M
												v311 = m.ExcPending
												if v311 != 0 {
													return int32(0)
												} else {
													return v310
												}
											}
										}
									} else {
										F_pfree(m, v95)
										mBase = m.M
										v173 = m.ExcPending
										if v173 != 0 {
											return int32(0)
										} else {
											v310 = F_construct_empty_array(m, int32(23))
											mBase = m.M
											v311 = m.ExcPending
											if v311 != 0 {
												return int32(0)
											} else {
												return v310
											}
										}
									}
								}
							} else {
								v80 = F_construct_empty_array(m, int32(23))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									v83 = v80 + int32(8)
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
									if v84 != 0 {
										v95 = v80
										v96 = v84
										v97 = v83
									} else {
										v85 = v80
										v87 = v83
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
										v95 = v85
										v96 = (v88<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										v97 = v87
									}
									v98 = int32(0)
									if base.B2i32(v30 <= v98)|base.B2i32(v33 <= v98) == v98 {
										v107 = v95 + v96
										v108 = int32(0)
										v110 = v108
										v111 = v108
										v114 = int32(0)
										for {
											v122 = int32(2)
											v125 = *(*int32)(unsafe.Add(mBase, uint32(l0+v45+v110<<(uint(v122)%32))))
											v129 = *(*int32)(unsafe.Add(mBase, uint32(l1+v56+v111<<(uint(v122)%32))))
											if v125 < v129 {
												v154 = v110 + int32(1)
												v155 = v111
												v156 = v114
											} else {
												if v125 == v129 {
													if v114 != 0 {
														v139 = *(*int32)(unsafe.Add(mBase, uint32(v107+v114<<(uint(int32(2))%32)-int32(4))))
														if v139 == v125 {
															v147 = v114
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v107+v114<<(uint(int32(2))%32)))) = v125
															v147 = v114 + int32(1)
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v107+v114<<(uint(int32(2))%32)))) = v125
														v147 = v114 + int32(1)
													}
													v148 = int32(1)
													v154 = v110 + v148
													v155 = v111 + v148
													v156 = v147
												} else {
													v154 = v110
													v155 = v111 + int32(1)
													v156 = v114
												}
											}
											if base.B2i32(v154 < v30)&base.B2i32(v155 < v33) != 0 {
												v110 = v154
												v111 = v155
												v114 = v156
												continue
											} else {
												break
											}
											break
										}
										if v156 != 0 {
											if v156 <= int32(0) {
												v310 = F_construct_empty_array(m, int32(23))
												mBase = m.M
												v311 = m.ExcPending
												if v311 != 0 {
													return int32(0)
												} else {
													return v310
												}
											} else {
												v176 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
												v179 = F_ArrayGetNItemsSafe(m, v176, v95+int32(16))
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return int32(0)
												} else {
													if v179 == v156 {
														v286 = v95
														return v286
													} else {
														v182 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
														if v182 != 0 {
															v190 = v182
														} else {
															v183 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
															v190 = (v183<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														}
														v193 = v190 + v156<<(uint(int32(2))%32)
														v194 = F_repalloc(m, v95, v193)
														mBase = m.M
														v195 = m.ExcPending
														if v195 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v194))) = v193 << (uint(int32(2)) % 32)
															v199 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
															if v199 <= int32(0) {
																v286 = v194
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v194)+16)) = v156
																if v199 == int32(1) {
																	v286 = v194
																} else {
																	v206 = v194 + int32(16)
																	v207 = int32(1)
																	v208 = v199 - v207
																	v209 = int32(7)
																	v210 = v208 & v209
																	if base.Ui32(v209) <= base.Ui32(v199-int32(2)) {
																		v219 = v207
																		v222 = int32(0)
																		for {
																			v233 = v206 + v219<<(uint(int32(2))%32)
																			v234 = int64(4294967297)
																			*(*int64)(unsafe.Add(mBase, uint32(v233)+24)) = v234
																			*(*int64)(unsafe.Add(mBase, uint32(v233)+16)) = v234
																			*(*int64)(unsafe.Add(mBase, uint32(v233)+8)) = v234
																			*(*int64)(unsafe.Add(mBase, uint32(v233))) = v234
																			v242 = int32(8)
																			v243 = v219 + v242
																			v245 = v222 + v242
																			if v245 != v208&int32(-8) {
																				v219 = v243
																				v222 = v245
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v210 == int32(0) {
																			v286 = v194
																		} else {
																			v249 = v243
																			v262 = v249
																			v263 = int32(0)
																			for {
																				v277 = int32(1)
																				*(*int32)(unsafe.Add(mBase, uint32(v206+v262<<(uint(int32(2))%32)))) = v277
																				v282 = v263 + v277
																				if v282 != v210 {
																					v262 = v262 + v277
																					v263 = v282
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v286 = v194
																		}
																	} else {
																		v249 = v207
																		v262 = v249
																		v263 = int32(0)
																		for {
																			v277 = int32(1)
																			*(*int32)(unsafe.Add(mBase, uint32(v206+v262<<(uint(int32(2))%32)))) = v277
																			v282 = v263 + v277
																			if v282 != v210 {
																				v262 = v262 + v277
																				v263 = v282
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		v286 = v194
																	}
																}
															}
															return v286
														}
													}
												}
											}
										} else {
											F_pfree(m, v95)
											mBase = m.M
											v173 = m.ExcPending
											if v173 != 0 {
												return int32(0)
											} else {
												v310 = F_construct_empty_array(m, int32(23))
												mBase = m.M
												v311 = m.ExcPending
												if v311 != 0 {
													return int32(0)
												} else {
													return v310
												}
											}
										}
									} else {
										F_pfree(m, v95)
										mBase = m.M
										v173 = m.ExcPending
										if v173 != 0 {
											return int32(0)
										} else {
											v310 = F_construct_empty_array(m, int32(23))
											mBase = m.M
											v311 = m.ExcPending
											if v311 != 0 {
												return int32(0)
											} else {
												return v310
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
