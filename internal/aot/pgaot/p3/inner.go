package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inner_int_inter(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = l0 + int32(16)
	v18 = F_ArrayGetNItems(m, v15, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v18 == int32(0) {
			v354 = F_construct_empty_array(m, int32(23))
			mBase = m.M
			v355 = m.ExcPending
			if v355 != 0 {
				return int32(0)
			} else {
				return v354
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v26 = l1 + int32(16)
			v27 = F_ArrayGetNItems(m, v24, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 == int32(0) {
					v354 = F_construct_empty_array(m, int32(23))
					mBase = m.M
					v355 = m.ExcPending
					if v355 != 0 {
						return int32(0)
					} else {
						return v354
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v32 = F_ArrayGetNItems(m, v31, v17)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v35 = F_ArrayGetNItems(m, v34, v26)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v37 == int32(0) {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v47 = (v40<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							} else {
								v47 = v37
							}
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							if v48 == int32(0) {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v58 = (v51<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							} else {
								v58 = v48
							}
							if v32 < v35 {
								v60 = v32
							} else {
								v60 = v35
							}
							if int32(0) < v60 {
								v66 = v60<<(uint(int32(2))%32) + int32(24)
								v67 = F_palloc0(m, v66)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v67)+20)) = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v60
									*(*int32)(unsafe.Add(mBase, uint32(v67)+12)) = int32(23)
									*(*int64)(unsafe.Add(mBase, uint32(v67)+4)) = int64(1)
									*(*int32)(unsafe.Add(mBase, uint32(v67))) = v66 << (uint(int32(2)) % 32)
									v88 = v67
									v90 = v67 + int32(8)
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
									v98 = (v91<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									v99 = v88
									v101 = v90
									if v32 <= int32(0) {
										F_pfree(m, v99)
										mBase = m.M
										v178 = m.ExcPending
										if v178 != 0 {
											return int32(0)
										} else {
											v354 = F_construct_empty_array(m, int32(23))
											mBase = m.M
											v355 = m.ExcPending
											if v355 != 0 {
												return int32(0)
											} else {
												return v354
											}
										}
									} else {
										if v35 <= int32(0) {
											F_pfree(m, v99)
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return int32(0)
											} else {
												v354 = F_construct_empty_array(m, int32(23))
												mBase = m.M
												v355 = m.ExcPending
												if v355 != 0 {
													return int32(0)
												} else {
													return v354
												}
											}
										} else {
											v108 = v98 + v99
											v111 = int32(0)
											v113 = v111
											v114 = v111
											v118 = int32(0)
											for {
												v127 = int32(2)
												v130 = *(*int32)(unsafe.Add(mBase, uint32(l0+v47+v113<<(uint(v127)%32))))
												v134 = *(*int32)(unsafe.Add(mBase, uint32(l1+v58+v114<<(uint(v127)%32))))
												if v130 < v134 {
													v157 = v113 + int32(1)
													v158 = v114
													v159 = v118
												} else {
													if v130 == v134 {
														if v118 != 0 {
															v142 = *(*int32)(unsafe.Add(mBase, uint32(v108-int32(4)+v118<<(uint(int32(2))%32))))
															if v142 == v130 {
																v150 = v118
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v108+v118<<(uint(int32(2))%32)))) = v130
																v150 = v118 + int32(1)
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v108+v118<<(uint(int32(2))%32)))) = v130
															v150 = v118 + int32(1)
														}
														v151 = int32(1)
														v157 = v113 + v151
														v158 = v114 + v151
														v159 = v150
													} else {
														v157 = v113
														v158 = v114 + int32(1)
														v159 = v118
													}
												}
												if base.B2i32(v157 < v32)&base.B2i32(v158 < v35) != 0 {
													v113 = v157
													v114 = v158
													v118 = v159
													continue
												} else {
													break
												}
												break
											}
											if v159 != 0 {
												if v159 <= int32(0) {
													v354 = F_construct_empty_array(m, int32(23))
													mBase = m.M
													v355 = m.ExcPending
													if v355 != 0 {
														return int32(0)
													} else {
														return v354
													}
												} else {
													v181 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
													v184 = F_ArrayGetNItems(m, v181, v99+int32(16))
													mBase = m.M
													v185 = m.ExcPending
													if v185 != 0 {
														return int32(0)
													} else {
														if v184 == v159 {
															v327 = v99
															return v327
														} else {
															v187 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
															if v187 != 0 {
																v195 = v187
															} else {
																v188 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
																v195 = (v188<<(uint(int32(3))%32) + int32(23)) & int32(-8)
															}
															v198 = v195 + v159<<(uint(int32(2))%32)
															v199 = F_repalloc(m, v99, v198)
															mBase = m.M
															v200 = m.ExcPending
															if v200 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v199))) = v198 << (uint(int32(2)) % 32)
																v204 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
																if v204 <= int32(0) {
																	v327 = v199
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v199)+16)) = v159
																	if v204 == int32(1) {
																		v327 = v199
																	} else {
																		v211 = v199 + int32(16)
																		v212 = int32(1)
																		v213 = v204 - v212
																		v214 = int32(7)
																		v215 = v213 & v214
																		if base.Ui32(v214) <= base.Ui32(v204-int32(2)) {
																			v239 = v212
																			v242 = int32(0)
																			for {
																				v253 = v239 << (uint(int32(2)) % 32)
																				v255 = int32(1)
																				*(*int32)(unsafe.Add(mBase, uint32(v211+v253))) = v255
																				*(*int32)(unsafe.Add(mBase, uint32(v253+(v199+int32(20))))) = v255
																				*(*int32)(unsafe.Add(mBase, uint32(v253+(v199+int32(24))))) = v255
																				*(*int32)(unsafe.Add(mBase, uint32(v253+(v199+int32(28))))) = v255
																				*(*int32)(unsafe.Add(mBase, uint32(v253+(v199+int32(32))))) = v255
																				*(*int32)(unsafe.Add(mBase, uint32(v253+(v199+int32(36))))) = v255
																				*(*int32)(unsafe.Add(mBase, uint32(v253+(v199+int32(40))))) = v255
																				*(*int32)(unsafe.Add(mBase, uint32(v253+(v199+int32(44))))) = v255
																				v278 = int32(8)
																				v279 = v239 + v278
																				v281 = v242 + v278
																				if v281 != v213&int32(-8) {
																					v239 = v279
																					v242 = v281
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v284 = v279
																		} else {
																			v284 = v212
																		}
																		if v215 == int32(0) {
																			v327 = v199
																		} else {
																			v300 = int32(0)
																			v301 = v284
																			for {
																				v317 = int32(1)
																				*(*int32)(unsafe.Add(mBase, uint32(v211+v301<<(uint(int32(2))%32)))) = v317
																				v322 = v300 + v317
																				if v322 != v215 {
																					v300 = v322
																					v301 = v301 + v317
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v327 = v199
																		}
																	}
																}
																return v327
															}
														}
													}
												}
											} else {
												F_pfree(m, v99)
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return int32(0)
												} else {
													v354 = F_construct_empty_array(m, int32(23))
													mBase = m.M
													v355 = m.ExcPending
													if v355 != 0 {
														return int32(0)
													} else {
														return v354
													}
												}
											}
										}
									}
								}
							} else {
								v82 = F_construct_empty_array(m, int32(23))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int32(0)
								} else {
									v85 = v82 + int32(8)
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
									if v86 != 0 {
										v98 = v86
										v99 = v82
										v101 = v85
									} else {
										v88 = v82
										v90 = v85
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
										v98 = (v91<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										v99 = v88
										v101 = v90
									}
									if v32 <= int32(0) {
										F_pfree(m, v99)
										mBase = m.M
										v178 = m.ExcPending
										if v178 != 0 {
											return int32(0)
										} else {
											v354 = F_construct_empty_array(m, int32(23))
											mBase = m.M
											v355 = m.ExcPending
											if v355 != 0 {
												return int32(0)
											} else {
												return v354
											}
										}
									} else {
										if v35 <= int32(0) {
											F_pfree(m, v99)
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return int32(0)
											} else {
												v354 = F_construct_empty_array(m, int32(23))
												mBase = m.M
												v355 = m.ExcPending
												if v355 != 0 {
													return int32(0)
												} else {
													return v354
												}
											}
										} else {
											v108 = v98 + v99
											v111 = int32(0)
											v113 = v111
											v114 = v111
											v118 = int32(0)
											for {
												v127 = int32(2)
												v130 = *(*int32)(unsafe.Add(mBase, uint32(l0+v47+v113<<(uint(v127)%32))))
												v134 = *(*int32)(unsafe.Add(mBase, uint32(l1+v58+v114<<(uint(v127)%32))))
												if v130 < v134 {
													v157 = v113 + int32(1)
													v158 = v114
													v159 = v118
												} else {
													if v130 == v134 {
														if v118 != 0 {
															v142 = *(*int32)(unsafe.Add(mBase, uint32(v108-int32(4)+v118<<(uint(int32(2))%32))))
															if v142 == v130 {
																v150 = v118
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v108+v118<<(uint(int32(2))%32)))) = v130
																v150 = v118 + int32(1)
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v108+v118<<(uint(int32(2))%32)))) = v130
															v150 = v118 + int32(1)
														}
														v151 = int32(1)
														v157 = v113 + v151
														v158 = v114 + v151
														v159 = v150
													} else {
														v157 = v113
														v158 = v114 + int32(1)
														v159 = v118
													}
												}
												if base.B2i32(v157 < v32)&base.B2i32(v158 < v35) != 0 {
													v113 = v157
													v114 = v158
													v118 = v159
													continue
												} else {
													break
												}
												break
											}
											if v159 != 0 {
												if v159 <= int32(0) {
													v354 = F_construct_empty_array(m, int32(23))
													mBase = m.M
													v355 = m.ExcPending
													if v355 != 0 {
														return int32(0)
													} else {
														return v354
													}
												} else {
													v181 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
													v184 = F_ArrayGetNItems(m, v181, v99+int32(16))
													mBase = m.M
													v185 = m.ExcPending
													if v185 != 0 {
														return int32(0)
													} else {
														if v184 == v159 {
															v327 = v99
															return v327
														} else {
															v187 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
															if v187 != 0 {
																v195 = v187
															} else {
																v188 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
																v195 = (v188<<(uint(int32(3))%32) + int32(23)) & int32(-8)
															}
															v198 = v195 + v159<<(uint(int32(2))%32)
															v199 = F_repalloc(m, v99, v198)
															mBase = m.M
															v200 = m.ExcPending
															if v200 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v199))) = v198 << (uint(int32(2)) % 32)
																v204 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
																if v204 <= int32(0) {
																	v327 = v199
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v199)+16)) = v159
																	if v204 == int32(1) {
																		v327 = v199
																	} else {
																		v211 = v199 + int32(16)
																		v212 = int32(1)
																		v213 = v204 - v212
																		v214 = int32(7)
																		v215 = v213 & v214
																		if base.Ui32(v214) <= base.Ui32(v204-int32(2)) {
																			v239 = v212
																			v242 = int32(0)
																			for {
																				v253 = v239 << (uint(int32(2)) % 32)
																				v255 = int32(1)
																				*(*int32)(unsafe.Add(mBase, uint32(v211+v253))) = v255
																				*(*int32)(unsafe.Add(mBase, uint32(v253+(v199+int32(20))))) = v255
																				*(*int32)(unsafe.Add(mBase, uint32(v253+(v199+int32(24))))) = v255
																				*(*int32)(unsafe.Add(mBase, uint32(v253+(v199+int32(28))))) = v255
																				*(*int32)(unsafe.Add(mBase, uint32(v253+(v199+int32(32))))) = v255
																				*(*int32)(unsafe.Add(mBase, uint32(v253+(v199+int32(36))))) = v255
																				*(*int32)(unsafe.Add(mBase, uint32(v253+(v199+int32(40))))) = v255
																				*(*int32)(unsafe.Add(mBase, uint32(v253+(v199+int32(44))))) = v255
																				v278 = int32(8)
																				v279 = v239 + v278
																				v281 = v242 + v278
																				if v281 != v213&int32(-8) {
																					v239 = v279
																					v242 = v281
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v284 = v279
																		} else {
																			v284 = v212
																		}
																		if v215 == int32(0) {
																			v327 = v199
																		} else {
																			v300 = int32(0)
																			v301 = v284
																			for {
																				v317 = int32(1)
																				*(*int32)(unsafe.Add(mBase, uint32(v211+v301<<(uint(int32(2))%32)))) = v317
																				v322 = v300 + v317
																				if v322 != v215 {
																					v300 = v322
																					v301 = v301 + v317
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v327 = v199
																		}
																	}
																}
																return v327
															}
														}
													}
												}
											} else {
												F_pfree(m, v99)
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return int32(0)
												} else {
													v354 = F_construct_empty_array(m, int32(23))
													mBase = m.M
													v355 = m.ExcPending
													if v355 != 0 {
														return int32(0)
													} else {
														return v354
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
