package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_makeVacuumRelation(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13945(m, l0, l1, l2, int32(240))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_vacuum_open_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v6 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l2&int32(32) == v6 {
		v22 = l4
		v23 = F_try_relation_open(m, l0, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			if v23 != 0 {
				v88 = v23
				m.G0 = v10 + int32(16)
				return v88
			} else {
				v27 = int32(1)
				if l1 != 0 {
					v28 = int32(0)
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_open_relation[0]))
					if base.B2i32(l3 == v28)&base.B2i32(v32 == int32(4)) != 0 {
						v88 = v28
						m.G0 = v10 + int32(16)
						return v88
					} else {
						if v32 != int32(4) {
							v40 = int32(19)
						} else {
							v40 = int32(15)
						}
						if l2&int32(1) != 0 {
							v44 = F_errstart(m, v40, int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								if v44 == int32(0) {
									v88 = v28
									m.G0 = v10 + int32(16)
									return v88
								} else {
									if v27 != 0 {
										v50 = int32(856)
									} else {
										v50 = int32(851)
									}
									if v27 != 0 {
										v53 = int32(_a_F_vacuum_open_relation_0)
									} else {
										v53 = int32(_a_F_vacuum_open_relation_1)
									}
									v69 = v50
									v70 = v53
									if v27 != 0 {
										v73 = int32(16908420)
									} else {
										v73 = int32(50463045)
									}
									F_errcode(m, v73)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v76
										F_errmsg(m, v70, v10)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_vacuum_open_relation_2), v69, int32(_a_F_vacuum_open_relation_3))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												v88 = int32(0)
												m.G0 = v10 + int32(16)
												return v88
											}
										}
									}
								}
							}
						} else {
							if l2&int32(2) == int32(0) {
								v88 = v28
								m.G0 = v10 + int32(16)
								return v88
							} else {
								v59 = F_errstart(m, v40, int32(0))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									if v59 == int32(0) {
										v88 = v28
										m.G0 = v10 + int32(16)
										return v88
									} else {
										if v27 != 0 {
											v65 = int32(877)
										} else {
											v65 = int32(872)
										}
										if v27 != 0 {
											v68 = int32(_a_F_vacuum_open_relation_4)
										} else {
											v68 = int32(_a_F_vacuum_open_relation_5)
										}
										v69 = v65
										v70 = v68
										if v27 != 0 {
											v73 = int32(16908420)
										} else {
											v73 = int32(50463045)
										}
										F_errcode(m, v73)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v76
											F_errmsg(m, v70, v10)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_vacuum_open_relation_2), v69, int32(_a_F_vacuum_open_relation_3))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													v88 = int32(0)
													m.G0 = v10 + int32(16)
													return v88
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v88 = int32(0)
					m.G0 = v10 + int32(16)
					return v88
				}
			}
		}
	} else {
		v16 = F_ConditionalLockRelationOid(m, l0, l4)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v16 == int32(0) {
				v27 = v6
				if l1 != 0 {
					v28 = int32(0)
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_open_relation[0]))
					if base.B2i32(l3 == v28)&base.B2i32(v32 == int32(4)) != 0 {
						v88 = v28
						m.G0 = v10 + int32(16)
						return v88
					} else {
						if v32 != int32(4) {
							v40 = int32(19)
						} else {
							v40 = int32(15)
						}
						if l2&int32(1) != 0 {
							v44 = F_errstart(m, v40, int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								if v44 == int32(0) {
									v88 = v28
									m.G0 = v10 + int32(16)
									return v88
								} else {
									if v27 != 0 {
										v50 = int32(856)
									} else {
										v50 = int32(851)
									}
									if v27 != 0 {
										v53 = int32(_a_F_vacuum_open_relation_0)
									} else {
										v53 = int32(_a_F_vacuum_open_relation_1)
									}
									v69 = v50
									v70 = v53
									if v27 != 0 {
										v73 = int32(16908420)
									} else {
										v73 = int32(50463045)
									}
									F_errcode(m, v73)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v76
										F_errmsg(m, v70, v10)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_vacuum_open_relation_2), v69, int32(_a_F_vacuum_open_relation_3))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												v88 = int32(0)
												m.G0 = v10 + int32(16)
												return v88
											}
										}
									}
								}
							}
						} else {
							if l2&int32(2) == int32(0) {
								v88 = v28
								m.G0 = v10 + int32(16)
								return v88
							} else {
								v59 = F_errstart(m, v40, int32(0))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									if v59 == int32(0) {
										v88 = v28
										m.G0 = v10 + int32(16)
										return v88
									} else {
										if v27 != 0 {
											v65 = int32(877)
										} else {
											v65 = int32(872)
										}
										if v27 != 0 {
											v68 = int32(_a_F_vacuum_open_relation_4)
										} else {
											v68 = int32(_a_F_vacuum_open_relation_5)
										}
										v69 = v65
										v70 = v68
										if v27 != 0 {
											v73 = int32(16908420)
										} else {
											v73 = int32(50463045)
										}
										F_errcode(m, v73)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v76
											F_errmsg(m, v70, v10)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_vacuum_open_relation_2), v69, int32(_a_F_vacuum_open_relation_3))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													v88 = int32(0)
													m.G0 = v10 + int32(16)
													return v88
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v88 = int32(0)
					m.G0 = v10 + int32(16)
					return v88
				}
			} else {
				v22 = v6
				v23 = F_try_relation_open(m, l0, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v23 != 0 {
						v88 = v23
						m.G0 = v10 + int32(16)
						return v88
					} else {
						v27 = int32(1)
						if l1 != 0 {
							v28 = int32(0)
							v32 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_open_relation[0]))
							if base.B2i32(l3 == v28)&base.B2i32(v32 == int32(4)) != 0 {
								v88 = v28
								m.G0 = v10 + int32(16)
								return v88
							} else {
								if v32 != int32(4) {
									v40 = int32(19)
								} else {
									v40 = int32(15)
								}
								if l2&int32(1) != 0 {
									v44 = F_errstart(m, v40, int32(0))
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										if v44 == int32(0) {
											v88 = v28
											m.G0 = v10 + int32(16)
											return v88
										} else {
											if v27 != 0 {
												v50 = int32(856)
											} else {
												v50 = int32(851)
											}
											if v27 != 0 {
												v53 = int32(_a_F_vacuum_open_relation_0)
											} else {
												v53 = int32(_a_F_vacuum_open_relation_1)
											}
											v69 = v50
											v70 = v53
											if v27 != 0 {
												v73 = int32(16908420)
											} else {
												v73 = int32(50463045)
											}
											F_errcode(m, v73)
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int32(0)
											} else {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v10))) = v76
												F_errmsg(m, v70, v10)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_vacuum_open_relation_2), v69, int32(_a_F_vacuum_open_relation_3))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int32(0)
													} else {
														v88 = int32(0)
														m.G0 = v10 + int32(16)
														return v88
													}
												}
											}
										}
									}
								} else {
									if l2&int32(2) == int32(0) {
										v88 = v28
										m.G0 = v10 + int32(16)
										return v88
									} else {
										v59 = F_errstart(m, v40, int32(0))
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											if v59 == int32(0) {
												v88 = v28
												m.G0 = v10 + int32(16)
												return v88
											} else {
												if v27 != 0 {
													v65 = int32(877)
												} else {
													v65 = int32(872)
												}
												if v27 != 0 {
													v68 = int32(_a_F_vacuum_open_relation_4)
												} else {
													v68 = int32(_a_F_vacuum_open_relation_5)
												}
												v69 = v65
												v70 = v68
												if v27 != 0 {
													v73 = int32(16908420)
												} else {
													v73 = int32(50463045)
												}
												F_errcode(m, v73)
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return int32(0)
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v10))) = v76
													F_errmsg(m, v70, v10)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_vacuum_open_relation_2), v69, int32(_a_F_vacuum_open_relation_3))
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
															return int32(0)
														} else {
															v88 = int32(0)
															m.G0 = v10 + int32(16)
															return v88
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v88 = int32(0)
							m.G0 = v10 + int32(16)
							return v88
						}
					}
				}
			}
		}
	}
}
func F_vacuum_reap_lp_read_stream_next(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v164 int64
	_ = v164
	var v173 int64
	_ = v173
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
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
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v385 int64
	_ = v385
	var v386 int64
	_ = v386
	var v387 int64
	_ = v387
	var v396 int64
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v465 int64
	_ = v465
	var v471 int32
	_ = v471
	var v480 int64
	_ = v480
	var v512 int32
	_ = v512
	var v517 int64
	_ = v517
	var v519 int32
	_ = v519
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	if v512 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L2:
	;
	v512 = int32(0)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v471
	*(*uint32)(unsafe.Add(mBase, uint32(l1)+8)) = uint32(v480)
	v512 = l1 + int32(8)
	goto L1
L4:
	;
	if v453 == int32(0) {
		goto L2
	} else {
		goto L75
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+100))
	if v18 < v17 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	if v243 < v242 {
		goto L2
	} else {
		goto L43
	}
L8:
	;
	v21 = v14 + int32(4)
	v28 = v17
	v31 = v18
	goto L9
L9:
	;
	v35 = int32(0)
	v39 = v21 + v28*int32(12)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	switch v41 {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	default:
		v147 = v35
		v150 = v35
		goto L13
	}
L10:
	;
	goto L2
L11:
	;
	if v234 <= v237 {
		v28 = v234
		v31 = v237
		goto L9
	} else {
		goto L42
	}
L12:
	;
	v226 = v28 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v226
	v234 = v226
	v237 = v31
	goto L11
L13:
	;
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v14)+112))
	v161 = int64(255)
	v164 = base.I64_extend_i32_u(v28 << (uint(int32(3)) % 32))
	v173 = v160&(v161<<(uint(v164)%64)^int64(-1)) | base.I64_extend_i32_u(v147)&v161<<(uint(v164)%64)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+112)) = v173
	v175 = int32(0)
	if v28|base.B2i32(v150 == v175) == v175 {
		goto L34
	} else {
		goto L35
	}
L14:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if int32(255) < v104 {
		goto L12
	} else {
		goto L27
	}
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if int32(255) < v68 {
		goto L12
	} else {
		goto L20
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)))
	if v56 <= v55 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)))
	if v43 <= v42 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v40)+3)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v42 + int32(1)
	v147 = v46
	v150 = v40 + v42<<(uint(int32(2))%32) + int32(8)
	goto L13
L19:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v40)+3)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v55 + int32(1)
	v147 = v59
	v150 = v40 + v55<<(uint(int32(2))%32) + int32(36)
	goto L13
L20:
	;
	v73 = v68
	goto L21
L21:
	;
	v86 = int32(255)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(12)+v73&v86))))
	if v89 == v86 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v73 + int32(1)
	v147 = v73
	v150 = v40 + v89<<(uint(int32(2))%32) + int32(268)
	goto L13
L23:
	;
	v93 = v73 + int32(1)
	if v93 != int32(256) {
		v73 = v93
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	goto L12
L27:
	;
	v109 = v104
	goto L28
L28:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(4)+int32(base.Ui32(v109)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v127)>>(uint(v109)%32))&int32(1) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v109 + int32(1)
	v147 = v109
	v150 = v40 + v109&int32(255)<<(uint(int32(2))%32) + int32(36)
	goto L13
L30:
	;
	v134 = v109 + int32(1)
	if v134 != int32(256) {
		v109 = v134
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L12
L34:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if v180&int32(1) != 0 {
		v471 = v150
		v480 = v173
		goto L3
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v150 == int32(0) {
		goto L12
	} else {
		goto L40
	}
L37:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v185 = F_dsa_get_address(m, v184, v180)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return int32(0)
L39:
	;
	v453 = v185
	v465 = v173
	goto L4
L40:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v194 = F_dsa_get_address(m, v192, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
	v198 = v196 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v198
	v200 = int32(12)
	v202 = v21 + v198*v200
	*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v193
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v21+v205*v200)+8)) = int32(0)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v14)+100))
	v234 = v205
	v237 = v211
	goto L11
L42:
	;
	goto L10
L43:
	;
	v246 = v14 + int32(4)
	v253 = v242
	v256 = v243
	goto L44
L44:
	;
	v260 = int32(0)
	v263 = v253 << (uint(int32(3)) % 32)
	v264 = v246 + v263
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	switch v266 {
	case 0:
		goto L52
	case 1:
		goto L51
	case 2:
		goto L50
	case 3:
		goto L49
	default:
		v372 = v260
		v375 = v260
		goto L48
	}
L45:
	;
	goto L2
L46:
	;
	if v445 <= v448 {
		v253 = v445
		v256 = v448
		goto L44
	} else {
		goto L74
	}
L47:
	;
	v437 = v429 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v437
	v445 = v437
	v448 = v256
	goto L46
L48:
	;
	v385 = *(*int64)(unsafe.Add(mBase, uint32(v14)+80))
	v386 = int64(255)
	v387 = base.I64_extend_i32_u(v263)
	v396 = v385&(v386<<(uint(v387)%64)^int64(-1)) | base.I64_extend_i32_u(v372)&v386<<(uint(v387)%64)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v396
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v399 = int32(0)
	if v398|base.B2i32(v375 == v399) == v399 {
		goto L69
	} else {
		goto L70
	}
L49:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	if int32(255) < v329 {
		v429 = v253
		goto L47
	} else {
		goto L62
	}
L50:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	if int32(255) < v293 {
		v429 = v253
		goto L47
	} else {
		goto L55
	}
L51:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+2)))
	if v281 <= v280 {
		v429 = v253
		goto L47
	} else {
		goto L54
	}
L52:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+2)))
	if v268 <= v267 {
		v429 = v253
		goto L47
	} else {
		goto L53
	}
L53:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267+v265)+3)))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v267 + int32(1)
	v372 = v271
	v375 = v265 + v267<<(uint(int32(2))%32) + int32(8)
	goto L48
L54:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v265)+3)))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v280 + int32(1)
	v372 = v284
	v375 = v265 + v280<<(uint(int32(2))%32) + int32(36)
	goto L48
L55:
	;
	v298 = v293
	goto L56
L56:
	;
	v311 = int32(255)
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+int32(12)+v298&v311))))
	if v314 == v311 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v298 + int32(1)
	v372 = v298
	v375 = v265 + v314<<(uint(int32(2))%32) + int32(268)
	goto L48
L58:
	;
	v318 = v298 + int32(1)
	if v318 != int32(256) {
		v298 = v318
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L57
L61:
	;
	v429 = v253
	goto L47
L62:
	;
	v334 = v329
	goto L63
L63:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v265+int32(4)+int32(base.Ui32(v334)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v352)>>(uint(v334)%32))&int32(1) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v334 + int32(1)
	v372 = v334
	v375 = v265 + v334&int32(255)<<(uint(int32(2))%32) + int32(36)
	goto L48
L65:
	;
	v359 = v334 + int32(1)
	if v359 != int32(256) {
		v334 = v359
		goto L63
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	goto L64
L68:
	;
	v429 = v253
	goto L47
L69:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	if v404&int32(1) == int32(0) {
		v453 = v404
		v465 = v396
		goto L4
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v375 == int32(0) {
		v429 = v398
		goto L47
	} else {
		goto L73
	}
L72:
	;
	v471 = v375
	v480 = v396
	goto L3
L73:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	v413 = v398 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v413
	v417 = v246 + v413<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v417)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v417))) = v411
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v445 = v422
	v448 = v421
	goto L46
L74:
	;
	goto L45
L75:
	;
	v471 = v453
	v480 = v465
	goto L3
L76:
	;
	return int32(-1)
L77:
	;
	goto L78
L78:
	;
	v517 = *(*int64)(unsafe.Add(mBase, uint32(v512)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v517
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	return v519
}
