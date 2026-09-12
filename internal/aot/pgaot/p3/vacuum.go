package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_makeVacuumRelation(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_palloc0(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(240)
		return v6
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
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
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
				v92 = v23
				m.G0 = v10 + int32(16)
				return v92
			} else {
				v27 = int32(1)
				if l1 != 0 {
					v28 = int32(0)
					v32 = *(*int32)(unsafe.Add(mBase, _consts[198]))
					if base.B2i32(l3 == v28)&base.B2i32(v32 == int32(4)) != 0 {
						v92 = v28
						m.G0 = v10 + int32(16)
						return v92
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
									v92 = v28
									m.G0 = v10 + int32(16)
									return v92
								} else {
									if v27 != 0 {
										v50 = int32(856)
									} else {
										v50 = int32(851)
									}
									if v27 != 0 {
										v53 = int32(117240)
									} else {
										v53 = int32(395516)
									}
									if v27 != 0 {
										v56 = int32(16908420)
									} else {
										v56 = int32(50463045)
									}
									v75 = v50
									v76 = v53
									v77 = v56
									F_errcode(m, v77)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v80
										F_errmsg(m, v76, v10)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496874), v75, int32(263026))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												v92 = int32(0)
												m.G0 = v10 + int32(16)
												return v92
											}
										}
									}
								}
							}
						} else {
							if l2&int32(2) == int32(0) {
								v92 = v28
								m.G0 = v10 + int32(16)
								return v92
							} else {
								v62 = F_errstart(m, v40, int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									if v62 == int32(0) {
										v92 = v28
										m.G0 = v10 + int32(16)
										return v92
									} else {
										if v27 != 0 {
											v68 = int32(877)
										} else {
											v68 = int32(872)
										}
										if v27 != 0 {
											v71 = int32(117294)
										} else {
											v71 = int32(395563)
										}
										if v27 != 0 {
											v74 = int32(16908420)
										} else {
											v74 = int32(50463045)
										}
										v75 = v68
										v76 = v71
										v77 = v74
										F_errcode(m, v77)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v80
											F_errmsg(m, v76, v10)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496874), v75, int32(263026))
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int32(0)
												} else {
													v92 = int32(0)
													m.G0 = v10 + int32(16)
													return v92
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v92 = int32(0)
					m.G0 = v10 + int32(16)
					return v92
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
					v32 = *(*int32)(unsafe.Add(mBase, _consts[198]))
					if base.B2i32(l3 == v28)&base.B2i32(v32 == int32(4)) != 0 {
						v92 = v28
						m.G0 = v10 + int32(16)
						return v92
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
									v92 = v28
									m.G0 = v10 + int32(16)
									return v92
								} else {
									if v27 != 0 {
										v50 = int32(856)
									} else {
										v50 = int32(851)
									}
									if v27 != 0 {
										v53 = int32(117240)
									} else {
										v53 = int32(395516)
									}
									if v27 != 0 {
										v56 = int32(16908420)
									} else {
										v56 = int32(50463045)
									}
									v75 = v50
									v76 = v53
									v77 = v56
									F_errcode(m, v77)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v80
										F_errmsg(m, v76, v10)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496874), v75, int32(263026))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												v92 = int32(0)
												m.G0 = v10 + int32(16)
												return v92
											}
										}
									}
								}
							}
						} else {
							if l2&int32(2) == int32(0) {
								v92 = v28
								m.G0 = v10 + int32(16)
								return v92
							} else {
								v62 = F_errstart(m, v40, int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									if v62 == int32(0) {
										v92 = v28
										m.G0 = v10 + int32(16)
										return v92
									} else {
										if v27 != 0 {
											v68 = int32(877)
										} else {
											v68 = int32(872)
										}
										if v27 != 0 {
											v71 = int32(117294)
										} else {
											v71 = int32(395563)
										}
										if v27 != 0 {
											v74 = int32(16908420)
										} else {
											v74 = int32(50463045)
										}
										v75 = v68
										v76 = v71
										v77 = v74
										F_errcode(m, v77)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v80
											F_errmsg(m, v76, v10)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496874), v75, int32(263026))
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int32(0)
												} else {
													v92 = int32(0)
													m.G0 = v10 + int32(16)
													return v92
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v92 = int32(0)
					m.G0 = v10 + int32(16)
					return v92
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
						v92 = v23
						m.G0 = v10 + int32(16)
						return v92
					} else {
						v27 = int32(1)
						if l1 != 0 {
							v28 = int32(0)
							v32 = *(*int32)(unsafe.Add(mBase, _consts[198]))
							if base.B2i32(l3 == v28)&base.B2i32(v32 == int32(4)) != 0 {
								v92 = v28
								m.G0 = v10 + int32(16)
								return v92
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
											v92 = v28
											m.G0 = v10 + int32(16)
											return v92
										} else {
											if v27 != 0 {
												v50 = int32(856)
											} else {
												v50 = int32(851)
											}
											if v27 != 0 {
												v53 = int32(117240)
											} else {
												v53 = int32(395516)
											}
											if v27 != 0 {
												v56 = int32(16908420)
											} else {
												v56 = int32(50463045)
											}
											v75 = v50
											v76 = v53
											v77 = v56
											F_errcode(m, v77)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v10))) = v80
												F_errmsg(m, v76, v10)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(496874), v75, int32(263026))
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return int32(0)
													} else {
														v92 = int32(0)
														m.G0 = v10 + int32(16)
														return v92
													}
												}
											}
										}
									}
								} else {
									if l2&int32(2) == int32(0) {
										v92 = v28
										m.G0 = v10 + int32(16)
										return v92
									} else {
										v62 = F_errstart(m, v40, int32(0))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											if v62 == int32(0) {
												v92 = v28
												m.G0 = v10 + int32(16)
												return v92
											} else {
												if v27 != 0 {
													v68 = int32(877)
												} else {
													v68 = int32(872)
												}
												if v27 != 0 {
													v71 = int32(117294)
												} else {
													v71 = int32(395563)
												}
												if v27 != 0 {
													v74 = int32(16908420)
												} else {
													v74 = int32(50463045)
												}
												v75 = v68
												v76 = v71
												v77 = v74
												F_errcode(m, v77)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v10))) = v80
													F_errmsg(m, v76, v10)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(496874), v75, int32(263026))
														mBase = m.M
														v87 = m.ExcPending
														if v87 != 0 {
															return int32(0)
														} else {
															v92 = int32(0)
															m.G0 = v10 + int32(16)
															return v92
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v92 = int32(0)
							m.G0 = v10 + int32(16)
							return v92
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
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
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
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v166 int64
	_ = v166
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
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
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
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
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v386 int64
	_ = v386
	var v387 int64
	_ = v387
	var v388 int64
	_ = v388
	var v397 int64
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v463 int64
	_ = v463
	var v469 int32
	_ = v469
	var v478 int64
	_ = v478
	var v510 int32
	_ = v510
	var v515 int64
	_ = v515
	var v517 int32
	_ = v517
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
	if v510 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L2:
	;
	v510 = int32(0)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v469
	*(*uint32)(unsafe.Add(mBase, uint32(l1)+8)) = uint32(v478)
	v510 = l1 + int32(8)
	goto L1
L4:
	;
	if v451 == int32(0) {
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
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	if v244 < v243 {
		goto L2
	} else {
		goto L43
	}
L8:
	;
	v23 = v14 + int32(4)
	v29 = v17
	goto L9
L9:
	;
	v37 = int32(0)
	v41 = v23 + v29*int32(12)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	switch v43 {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	default:
		v149 = v37
		v152 = v37
		goto L13
	}
L10:
	;
	goto L2
L11:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v14)+100))
	if v233 <= v241 {
		v29 = v233
		goto L9
	} else {
		goto L42
	}
L12:
	;
	v226 = v217 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v226
	v233 = v226
	goto L11
L13:
	;
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v14)+112))
	v163 = int64(255)
	v166 = base.I64_extend_i32_u(v29 << (uint(int32(3)) % 32))
	v175 = v162&(v163<<(uint(v166)%64)^int64(-1)) | base.I64_extend_i32_u(v149)&v163<<(uint(v166)%64)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+112)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
	if v177 != 0 {
		goto L34
	} else {
		goto L35
	}
L14:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	if int32(255) < v106 {
		v217 = v29
		goto L12
	} else {
		goto L27
	}
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	if int32(255) < v70 {
		v217 = v29
		goto L12
	} else {
		goto L20
	}
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+2)))
	if v58 <= v57 {
		v217 = v29
		goto L12
	} else {
		goto L19
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+2)))
	if v45 <= v44 {
		v217 = v29
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v42)+3)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v44 + int32(1)
	v149 = v48
	v152 = v42 + v44<<(uint(int32(2))%32) + int32(8)
	goto L13
L19:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+v42)+3)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v57 + int32(1)
	v149 = v61
	v152 = v42 + v57<<(uint(int32(2))%32) + int32(36)
	goto L13
L20:
	;
	v75 = v70
	goto L21
L21:
	;
	v88 = int32(255)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(12)+v75&v88))))
	if v91 == v88 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v75 + int32(1)
	v149 = v75
	v152 = v42 + v91<<(uint(int32(2))%32) + int32(268)
	goto L13
L23:
	;
	v95 = v75 + int32(1)
	if v95 != int32(256) {
		v75 = v95
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
	v217 = v29
	goto L12
L27:
	;
	v111 = v106
	goto L28
L28:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(4)+int32(base.Ui32(v111)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v129)>>(uint(v111)%32))&int32(1) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v111 + int32(1)
	v149 = v111
	v152 = v42 + v111&int32(255)<<(uint(int32(2))%32) + int32(36)
	goto L13
L30:
	;
	v136 = v111 + int32(1)
	if v136 != int32(256) {
		v111 = v136
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
	v217 = v29
	goto L12
L34:
	;
	if v152 == int32(0) {
		v217 = v177
		goto L12
	} else {
		goto L40
	}
L35:
	;
	if v152 == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	if v180&int32(1) != 0 {
		v469 = v152
		v478 = v175
		goto L3
	} else {
		goto L37
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
	v451 = v185
	v463 = v175
	goto L4
L40:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
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
	v202 = v23 + v198*v200
	*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v193
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(12)+v205*v200))) = int32(0)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
	v233 = v211
	goto L11
L42:
	;
	goto L10
L43:
	;
	v247 = v14 + int32(4)
	v253 = v243
	v256 = v244
	goto L44
L44:
	;
	v261 = int32(0)
	v264 = v253 << (uint(int32(3)) % 32)
	v265 = v247 + v264
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	switch v267 {
	case 0:
		goto L52
	case 1:
		goto L51
	case 2:
		goto L50
	case 3:
		goto L49
	default:
		v373 = v261
		v376 = v261
		goto L48
	}
L45:
	;
	goto L2
L46:
	;
	if v442 <= v445 {
		v253 = v442
		v256 = v445
		goto L44
	} else {
		goto L74
	}
L47:
	;
	v435 = v426 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v435
	v442 = v435
	v445 = v256
	goto L46
L48:
	;
	v386 = *(*int64)(unsafe.Add(mBase, uint32(v14)+80))
	v387 = int64(255)
	v388 = base.I64_extend_i32_u(v264)
	v397 = v386&(v387<<(uint(v388)%64)^int64(-1)) | base.I64_extend_i32_u(v373)&v387<<(uint(v388)%64)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v397
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	if v399 != 0 {
		goto L69
	} else {
		goto L70
	}
L49:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	if int32(255) < v330 {
		v426 = v253
		goto L47
	} else {
		goto L62
	}
L50:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	if int32(255) < v294 {
		v426 = v253
		goto L47
	} else {
		goto L55
	}
L51:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+2)))
	if v282 <= v281 {
		v426 = v253
		goto L47
	} else {
		goto L54
	}
L52:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+2)))
	if v269 <= v268 {
		v426 = v253
		goto L47
	} else {
		goto L53
	}
L53:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+v266)+3)))
	*(*int32)(unsafe.Add(mBase, uint32(v265)+4)) = v268 + int32(1)
	v373 = v272
	v376 = v266 + v268<<(uint(int32(2))%32) + int32(8)
	goto L48
L54:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281+v266)+3)))
	*(*int32)(unsafe.Add(mBase, uint32(v265)+4)) = v281 + int32(1)
	v373 = v285
	v376 = v266 + v281<<(uint(int32(2))%32) + int32(36)
	goto L48
L55:
	;
	v299 = v294
	goto L56
L56:
	;
	v312 = int32(255)
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266+int32(12)+v299&v312))))
	if v315 == v312 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v265)+4)) = v299 + int32(1)
	v373 = v299
	v376 = v266 + v315<<(uint(int32(2))%32) + int32(268)
	goto L48
L58:
	;
	v319 = v299 + int32(1)
	if v319 != int32(256) {
		v299 = v319
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
	v426 = v253
	goto L47
L62:
	;
	v335 = v330
	goto L63
L63:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v266+int32(4)+int32(base.Ui32(v335)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v353)>>(uint(v335)%32))&int32(1) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v265)+4)) = v335 + int32(1)
	v373 = v335
	v376 = v266 + v335&int32(255)<<(uint(int32(2))%32) + int32(36)
	goto L48
L65:
	;
	v360 = v335 + int32(1)
	if v360 != int32(256) {
		v335 = v360
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
	v426 = v253
	goto L47
L69:
	;
	if v376 == int32(0) {
		v426 = v399
		goto L47
	} else {
		goto L73
	}
L70:
	;
	if v376 == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	if v402&int32(1) == int32(0) {
		v451 = v402
		v463 = v397
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v469 = v376
	v478 = v397
	goto L3
L73:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	v411 = v399 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v411
	v415 = v247 + v411<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v415))) = v409
	*(*int32)(unsafe.Add(mBase, uint32(v415)+4)) = int32(0)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v442 = v420
	v445 = v419
	goto L46
L74:
	;
	goto L45
L75:
	;
	v469 = v451
	v478 = v463
	goto L3
L76:
	;
	return int32(-1)
L77:
	;
	goto L78
L78:
	;
	v515 = *(*int64)(unsafe.Add(mBase, uint32(v510)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v515
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v510)))
	return v517
}
