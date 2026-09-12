package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_addUnicodeChar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
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
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	if l0 == v4 {
		v15 = F_errsave_start(m, l1)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				v117 = v4
				m.G0 = v11 + int32(32)
				return v117
			} else {
				F_errcode(m, int32(84017282))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(437368), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errdetail(m, int32(607950), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, l1, int32(330737), int32(585), int32(242729))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v117 = v4
								m.G0 = v11 + int32(32)
								return v117
							}
						}
					}
				}
			}
		}
	} else {
		if l1 != 0 {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v37 == int32(447) {
				v42 = F_pg_unicode_to_server_noerror(m, l0, v11)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					if v42 != 0 {
						v60 = int32(1)
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
						v63 = F_strlen(m, v11)
						mBase = m.M
						v65 = v63 + v60
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
						if v67 <= v62+v65 {
							v71 = v67
							v76 = v61 + int32(8)
							for {
								*(*int32)(unsafe.Add(mBase, uint32(v76))) = v71 << (uint(int32(1)) % 32)
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
								if v85 <= v86+v65 {
									v71 = v85
									v76 = v82 + int32(8)
									continue
								} else {
									break
								}
								break
							}
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
							v90 = F_repalloc(m, v89, v85)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								*(*int32)(unsafe.Add(mBase, uint32(v92))) = v90
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
								v97 = v94
								v104 = v95
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
								if v63 != 0 {
									v107 = F__emscripten_memcpy_bulkmem(m, v104+v105, v11, v63)
									mBase = m.M
								} else {
								}
								v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + v63
								v117 = v60
								m.G0 = v11 + int32(32)
								return v117
							}
						} else {
							v97 = v61
							v104 = v62
							v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
							if v63 != 0 {
								v107 = F__emscripten_memcpy_bulkmem(m, v104+v105, v11, v63)
								mBase = m.M
							} else {
							}
							v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + v63
							v117 = v60
							m.G0 = v11 + int32(32)
							return v117
						}
					} else {
						v44 = F_errsave_start(m, l1)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							if v44 == int32(0) {
								v117 = v4
								m.G0 = v11 + int32(32)
								return v117
							} else {
								F_errcode(m, int32(16801924))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(353967), int32(0))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, l1, int32(330737), int32(602), int32(242729))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v117 = v4
											m.G0 = v11 + int32(32)
											return v117
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_pg_unicode_to_server(m, l0, v11)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v60 = int32(1)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
					v63 = F_strlen(m, v11)
					mBase = m.M
					v65 = v63 + v60
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
					if v67 <= v62+v65 {
						v71 = v67
						v76 = v61 + int32(8)
						for {
							*(*int32)(unsafe.Add(mBase, uint32(v76))) = v71 << (uint(int32(1)) % 32)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
							if v85 <= v86+v65 {
								v71 = v85
								v76 = v82 + int32(8)
								continue
							} else {
								break
							}
							break
						}
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
						v90 = F_repalloc(m, v89, v85)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							*(*int32)(unsafe.Add(mBase, uint32(v92))) = v90
							v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
							v97 = v94
							v104 = v95
							v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
							if v63 != 0 {
								v107 = F__emscripten_memcpy_bulkmem(m, v104+v105, v11, v63)
								mBase = m.M
							} else {
							}
							v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + v63
							v117 = v60
							m.G0 = v11 + int32(32)
							return v117
						}
					} else {
						v97 = v61
						v104 = v62
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
						if v63 != 0 {
							v107 = F__emscripten_memcpy_bulkmem(m, v104+v105, v11, v63)
							mBase = m.M
						} else {
						}
						v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + v63
						v117 = v60
						m.G0 = v11 + int32(32)
						return v117
					}
				}
			}
		} else {
			F_pg_unicode_to_server(m, l0, v11)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v60 = int32(1)
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
				v63 = F_strlen(m, v11)
				mBase = m.M
				v65 = v63 + v60
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
				if v67 <= v62+v65 {
					v71 = v67
					v76 = v61 + int32(8)
					for {
						*(*int32)(unsafe.Add(mBase, uint32(v76))) = v71 << (uint(int32(1)) % 32)
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
						if v85 <= v86+v65 {
							v71 = v85
							v76 = v82 + int32(8)
							continue
						} else {
							break
						}
						break
					}
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
					v90 = F_repalloc(m, v89, v85)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						*(*int32)(unsafe.Add(mBase, uint32(v92))) = v90
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
						v97 = v94
						v104 = v95
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
						if v63 != 0 {
							v107 = F__emscripten_memcpy_bulkmem(m, v104+v105, v11, v63)
							mBase = m.M
						} else {
						}
						v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + v63
						v117 = v60
						m.G0 = v11 + int32(32)
						return v117
					}
				} else {
					v97 = v61
					v104 = v62
					v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
					if v63 != 0 {
						v107 = F__emscripten_memcpy_bulkmem(m, v104+v105, v11, v63)
						mBase = m.M
					} else {
					}
					v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + v63
					v117 = v60
					m.G0 = v11 + int32(32)
					return v117
				}
			}
		}
	}
}
func F_unicode_is_normalized(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v167 int32
	_ = v167
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_pg_detoast_datum_packed(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v58 = F_palloc(m, v55+int32(1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L16
	}
L4:
	;
	v23 = F_pg_detoast_datum_packed(m, v21)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v25 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = int32(4)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v30&int32(254) == int32(2) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v43 = int32(1)
	if v25&v43 != 0 {
		v55 = int32(base.Ui32(v25)>>(uint(v43)%32)) - v43
		goto L3
	} else {
		goto L15
	}
L9:
	;
	v39 = v28
	goto L11
L10:
	;
	v39 = base.B2i32(v30 == int32(18)) << (uint(v28) % 32)
	goto L11
L11:
	;
	if v30 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v42 = v28
	goto L14
L13:
	;
	v42 = v39
	goto L14
L14:
	;
	v55 = v42
	goto L3
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L16:
	;
	v60 = int32(1)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v62&v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v65 = v60
	goto L19
L18:
	;
	v65 = int32(4)
	goto L19
L19:
	;
	if v55 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55+v68))) = uint8(v70)
	if v23 != v21 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v67 = F__emscripten_memcpy_bulkmem(m, v58, v23+v65, v55)
	mBase = m.M
	v68 = v67
	goto L23
L22:
	;
	v68 = v58
	goto L23
L23:
	;
	goto L20
L24:
	;
	F_pfree(m, v23)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v75 = F_unicode_norm_form_from_string(m, v68)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v77 = int32(1)
	v78 = v16 + v77
	v80 = v16 + int32(4)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v83 = v81 & v77
	if v83 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v84 = v78
	goto L31
L30:
	;
	v84 = v80
	goto L31
L31:
	;
	if v81 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v113 = F_pg_mbstrlen_with_len(m, v84, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L43
	}
L33:
	;
	v87 = int32(4)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v89&int32(254) == int32(2) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v102 = int32(1)
	if v83 != 0 {
		v112 = int32(base.Ui32(v81)>>(uint(v102)%32)) - v102
		goto L32
	} else {
		goto L42
	}
L36:
	;
	v98 = v87
	goto L38
L37:
	;
	v98 = base.B2i32(v89 == int32(18)) << (uint(v87) % 32)
	goto L38
L38:
	;
	if v89 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v101 = v87
	goto L41
L40:
	;
	v101 = v98
	goto L41
L41:
	;
	v112 = v101
	goto L32
L42:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v112 = int32(base.Ui32(v106)>>(uint(int32(2))%32)) - int32(4)
	goto L32
L43:
	;
	v116 = v113 << (uint(int32(2)) % 32)
	v119 = F_palloc(m, v116+int32(4))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if int32(0) < v113 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v123&int32(1) != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v250 = int32(0)
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119+v250<<(uint(int32(2))%32)))) = int32(0)
	v256 = int32(1)
	if v75&int32(-3) == v256 {
		v432 = int32(-1)
		goto L78
	} else {
		goto L79
	}
L48:
	;
	v126 = v78
	goto L50
L49:
	;
	v126 = v80
	goto L50
L50:
	;
	v128 = v126
	v130 = int32(0)
	goto L51
L51:
	;
	v145 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
	v147 = v145 & int32(255)
	if int32(0) <= v145 {
		v204 = v147
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v250 = v113
	goto L47
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119+v130<<(uint(int32(2))%32)))) = v204
	v206 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
	if int32(0) <= v206 {
		goto L64
	} else {
		goto L65
	}
L54:
	;
	if v147&int32(224) == int32(192) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197+v128))))
	v204 = v199&int32(63) | v196
	goto L53
L56:
	;
	v196 = v147 << (uint(int32(6)) % 32) & int32(1984)
	v197 = int32(1)
	goto L55
L57:
	;
	goto L58
L58:
	;
	if v147&int32(240) == int32(224) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
	v196 = v147<<(uint(int32(12))%32)&int32(61440) | v167&int32(63)<<(uint(int32(6))%32)
	v197 = int32(2)
	goto L55
L60:
	;
	goto L61
L61:
	;
	if v147&int32(248) != int32(240) {
		v204 = int32(-1)
		goto L53
	} else {
		goto L62
	}
L62:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
	v184 = int32(63)
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+2)))
	v196 = v147<<(uint(int32(18))%32)&int32(1835008) | v183&v184<<(uint(int32(12))%32) | v189&v184<<(uint(int32(6))%32)
	v197 = int32(3)
	goto L55
L63:
	;
	v233 = v130 + int32(1)
	if v233 != v113 {
		v128 = v230 + v128
		v130 = v233
		goto L51
	} else {
		goto L76
	}
L64:
	;
	v230 = int32(1)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v211 = v206 & int32(255)
	if v211&int32(224) == int32(192) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v230 = int32(2)
	goto L63
L68:
	;
	goto L69
L69:
	;
	if v211&int32(240) == int32(224) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v230 = int32(3)
	goto L63
L71:
	;
	goto L72
L72:
	;
	if v211&int32(248) == int32(240) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v228 = int32(4)
	goto L75
L74:
	;
	v228 = int32(1)
	goto L75
L75:
	;
	v230 = v228
	goto L63
L76:
	;
	goto L52
L77:
	;
	return v539
L78:
	;
	if base.Ui32(v432) < base.Ui32(int32(2)) {
		v539 = v432
		goto L77
	} else {
		goto L99
	}
L79:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if v263 == int32(0) {
		v432 = int32(1)
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v267 = v119
	v268 = v256
	v271 = v263
	v275 = int32(0)
	goto L81
L81:
	;
	v280 = int32(0)
	v282 = v271 & int32(255)
	v283 = int32(8)
	v288 = int32(base.Ui32(v271<<(uint(v283)%32)&int32(16711680)) >> (uint(int32(16)) % 32))
	v294 = int32(base.Ui32(int32(base.Ui32(v271)>>(uint(v283)%32))&int32(65280)) >> (uint(v283) % 32))
	v296 = int32(base.Ui32(v271) >> (uint(int32(24)) % 32))
	v297 = int32(8191)
	v305 = v282 + (v288+(v294+v296*v297)*v297)*v297
	v308 = int32(13687)
	v309 = base.I32_rem_u_s(v305+int32(402620417), v308)
	v310 = int32(1)
	v314 = int32(*(*int16)(unsafe.Add(mBase, uint32(v309<<(uint(v310)%32))+uint32(_consts[1347]))))
	v315 = int32(257)
	v323 = ((v296*v315+v294)*v315+v288)*v315 + v282
	v325 = base.I32_rem_u_s(v323, v308)
	v330 = int32(*(*int16)(unsafe.Add(mBase, uint32(v325<<(uint(v310)%32))+uint32(_consts[1347]))))
	v331 = v314 + v330
	if base.Ui32(int32(6842)) < base.Ui32(v331) {
		v347 = v280
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v432 = v413
	goto L78
L83:
	;
	switch v75 {
	case 0:
		goto L93
	default:
		goto L90
	case 2:
		goto L92
	}
L84:
	;
	v335 = v331 << (uint(int32(3)) % 32)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v335)+uint32(_consts[1348])))
	if v271 != v338 {
		v347 = v280
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+uint32(_consts[1349]))))
	if base.Ui32(v275&int32(255)) <= base.Ui32(v340) {
		v347 = v340
		goto L83
	} else {
		goto L86
	}
L86:
	;
	if v340 == int32(0) {
		v347 = v340
		goto L83
	} else {
		goto L87
	}
L87:
	;
	v432 = int32(0)
	goto L78
L88:
	;
	goto L82
L89:
	;
	v411 = v267 + int32(4)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	if v412 != 0 {
		v267 = v411
		v268 = v409
		v271 = v412
		v275 = v347
		goto L81
	} else {
		goto L98
	}
L90:
	;
	v409 = v268
	goto L89
L91:
	;
	v405 = v401 << (uint(int32(7)) % 32) >> (uint(int32(28)) % 32)
	switch v405 + int32(1) {
	case 0:
		v409 = v405
		goto L89
	case 1:
		v413 = v405
		goto L88
	default:
		goto L90
	}
L92:
	;
	v376 = int32(10193)
	v377 = base.I32_rem_u_s(v305+int32(1207861251), v376)
	v378 = int32(1)
	v382 = int32(*(*int16)(unsafe.Add(mBase, uint32(v377<<(uint(v378)%32))+uint32(_consts[1350]))))
	v384 = base.I32_rem_u_s(v323, v376)
	v389 = int32(*(*int16)(unsafe.Add(mBase, uint32(v384<<(uint(v378)%32))+uint32(_consts[1350]))))
	v390 = v382 + v389
	if base.Ui32(int32(5095)) < base.Ui32(v390) {
		goto L90
	} else {
		goto L96
	}
L93:
	;
	v349 = int32(2505)
	v350 = base.I32_rem_u_s(v305, v349)
	v351 = int32(1)
	v355 = int32(*(*int16)(unsafe.Add(mBase, uint32(v350<<(uint(v351)%32))+uint32(_consts[1351]))))
	v357 = base.I32_rem_u_s(v323, v349)
	v362 = int32(*(*int16)(unsafe.Add(mBase, uint32(v357<<(uint(v351)%32))+uint32(_consts[1351]))))
	v363 = v355 + v362
	if base.Ui32(int32(1251)) < base.Ui32(v363) {
		goto L90
	} else {
		goto L94
	}
L94:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v363<<(uint(int32(2))%32))+uint32(_consts[1352])))
	if v271 == v370&int32(2097151) {
		v401 = v370
		goto L91
	} else {
		goto L95
	}
L95:
	;
	goto L90
L96:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v390<<(uint(int32(2))%32))+uint32(_consts[1353])))
	if v271 != v397&int32(2097151) {
		goto L90
	} else {
		goto L97
	}
L97:
	;
	v401 = v397
	goto L91
L98:
	;
	v413 = v409
	goto L88
L99:
	;
	v435 = int32(0)
	v437 = F_unicode_normalize(m, v75, v119)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	if v439 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v440 = v435
	v442 = v437
	goto L104
L102:
	;
	v459 = v435
	goto L103
L103:
	;
	if v459 != v113 {
		v539 = v435
		goto L77
	} else {
		goto L107
	}
L104:
	;
	v455 = v440 + int32(1)
	v457 = v442 + int32(4)
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	if v458 != 0 {
		v440 = v455
		v442 = v457
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v459 = v455
	goto L103
L106:
	;
	goto L105
L107:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v116) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v539 = base.B2i32(v535 == int32(0))
	goto L77
L109:
	;
	v535 = int32(0)
	goto L108
L110:
	;
	v509 = v504
	v510 = v505
	v511 = v506
	goto L120
L111:
	;
	if (v119|v437)&int32(3) != 0 {
		v504 = v119
		v505 = v437
		v506 = v116
		goto L110
	} else {
		goto L114
	}
L112:
	;
	v497 = v119
	v498 = v437
	v499 = v116
	goto L113
L113:
	;
	if v499 == int32(0) {
		goto L109
	} else {
		goto L119
	}
L114:
	;
	v481 = v119
	v482 = v437
	v483 = v116
	goto L115
L115:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v481)))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v482)))
	if v486 != v487 {
		v504 = v481
		v505 = v482
		v506 = v483
		goto L110
	} else {
		goto L117
	}
L116:
	;
	v497 = v492
	v498 = v490
	v499 = v494
	goto L113
L117:
	;
	v489 = int32(4)
	v490 = v482 + v489
	v492 = v481 + v489
	v494 = v483 - v489
	if base.Ui32(int32(3)) < base.Ui32(v494) {
		v481 = v492
		v482 = v490
		v483 = v494
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v504 = v497
	v505 = v498
	v506 = v499
	goto L110
L120:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509))))
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510))))
	if v514 == v515 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v535 = v514 - v515
	goto L108
L122:
	;
	v517 = int32(1)
	v522 = v511 - v517
	if v522 != 0 {
		v509 = v509 + v517
		v510 = v510 + v517
		v511 = v522
		goto L120
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	goto L121
L125:
	;
	goto L109
}
func F_unicode_normalize_func(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v163 int32
	_ = v163
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = F_pg_detoast_datum_packed(m, v20)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				if v24 == int32(1) {
					v27 = int32(4)
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
					if v29&int32(254) == int32(2) {
						v38 = v27
					} else {
						v38 = base.B2i32(v29 == int32(18)) << (uint(v27) % 32)
					}
					if v29 == int32(1) {
						v41 = v27
					} else {
						v41 = v38
					}
					v54 = v41
				} else {
					v42 = int32(1)
					if v24&v42 != 0 {
						v54 = int32(base.Ui32(v24)>>(uint(v42)%32)) - v42
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
						v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v57 = F_palloc(m, v54+int32(1))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					v59 = int32(1)
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
					if v61&v59 != 0 {
						v64 = v59
					} else {
						v64 = int32(4)
					}
					if v54 != 0 {
						v66 = F__emscripten_memcpy_bulkmem(m, v57, v22+v64, v54)
						mBase = m.M
						v67 = v66
					} else {
						v67 = v57
					}
					v69 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v54+v67))) = uint8(v69)
					if v22 != v20 {
						F_pfree(m, v22)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							v74 = F_unicode_norm_form_from_string(m, v67)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v76 = int32(4)
								v77 = int32(1)
								v78 = v15 + v77
								v80 = v15 + v76
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
								v83 = v81 & v77
								if v83 != 0 {
									v84 = v78
								} else {
									v84 = v80
								}
								if v81 == int32(1) {
									v87 = int32(4)
									v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
									if v89&int32(254) == int32(2) {
										v98 = v87
									} else {
										v98 = base.B2i32(v89 == int32(18)) << (uint(v87) % 32)
									}
									if v89 == int32(1) {
										v101 = v87
									} else {
										v101 = v98
									}
									v112 = v101
								} else {
									v102 = int32(1)
									if v83 != 0 {
										v112 = int32(base.Ui32(v81)>>(uint(v102)%32)) - v102
									} else {
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
										v112 = int32(base.Ui32(v106)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								v113 = F_pg_mbstrlen_with_len(m, v84, v112)
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int32(0)
								} else {
									v119 = F_palloc(m, v113<<(uint(int32(2))%32)+int32(4))
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int32(0)
									} else {
										v121 = int32(0)
										if v121 < v113 {
											v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
											if v124&int32(1) != 0 {
												v127 = v78
											} else {
												v127 = v80
											}
											v129 = v127
											v130 = int32(0)
											for {
												v138 = int32(*(*int8)(unsafe.Add(mBase, uint32(v129))))
												v140 = v138 & int32(255)
												if int32(0) <= v138 {
													v200 = v140
												} else {
													if v140&int32(224) == int32(192) {
														v192 = v140 << (uint(int32(6)) % 32) & int32(1984)
														v193 = int32(1)
														v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v129))))
														v200 = v195&int32(63) | v192
													} else {
														if v140&int32(240) == int32(224) {
															v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
															v192 = v140<<(uint(int32(12))%32)&int32(61440) | v163&int32(63)<<(uint(int32(6))%32)
															v193 = int32(2)
															v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v129))))
															v200 = v195&int32(63) | v192
														} else {
															if v140&int32(248) != int32(240) {
																v200 = int32(-1)
															} else {
																v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
																v180 = int32(63)
																v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+2)))
																v192 = v140<<(uint(int32(18))%32)&int32(1835008) | v179&v180<<(uint(int32(12))%32) | v185&v180<<(uint(int32(6))%32)
																v193 = int32(3)
																v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v129))))
																v200 = v195&int32(63) | v192
															}
														}
													}
												}
												*(*int32)(unsafe.Add(mBase, uint32(v119+v130<<(uint(int32(2))%32)))) = v200
												v202 = int32(*(*int8)(unsafe.Add(mBase, uint32(v129))))
												if int32(0) <= v202 {
													v226 = int32(1)
												} else {
													v207 = v202 & int32(255)
													if v207&int32(224) == int32(192) {
														v226 = int32(2)
													} else {
														if v207&int32(240) == int32(224) {
															v226 = int32(3)
														} else {
															if v207&int32(248) == int32(240) {
																v224 = int32(4)
															} else {
																v224 = int32(1)
															}
															v226 = v224
														}
													}
												}
												v229 = v130 + int32(1)
												if v229 != v113 {
													v129 = v226 + v129
													v130 = v229
													continue
												} else {
													break
												}
												break
											}
											v240 = v113
										} else {
											v240 = v121
										}
										*(*int32)(unsafe.Add(mBase, uint32(v119+v240<<(uint(int32(2))%32)))) = int32(0)
										v246 = F_unicode_normalize(m, v74, v119)
										mBase = m.M
										v247 = m.ExcPending
										if v247 != 0 {
											return int32(0)
										} else {
											v248 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
											if v248 != 0 {
												v250 = v248
												v251 = int32(0)
												v254 = v246
												for {
													if base.Ui32(v250) <= base.Ui32(int32(127)) {
														*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v250)
													} else {
														if base.Ui32(v250) <= base.Ui32(int32(2047)) {
															v267 = v250&int32(63) | int32(128)
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v267)
															v272 = int32(base.Ui32(v250)>>(uint(int32(6))%32)) | int32(192)
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v272)
														} else {
															if base.Ui32(v250) <= base.Ui32(int32(65535)) {
																v276 = int32(63)
																v278 = int32(128)
																v279 = v250&v276 | v278
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v279)
																v284 = int32(base.Ui32(v250)>>(uint(int32(12))%32)) | int32(224)
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v284)
																v291 = int32(base.Ui32(v250)>>(uint(int32(6))%32))&v276 | v278
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v291)
															} else {
																v293 = int32(63)
																v295 = int32(128)
																v296 = v250&v293 | v295
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v296)
																v303 = int32(base.Ui32(v250)>>(uint(int32(6))%32))&v293 | v295
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v303)
																v310 = int32(base.Ui32(v250)>>(uint(int32(12))%32))&v293 | v295
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v310)
																v317 = int32(base.Ui32(v250)>>(uint(int32(18))%32))&int32(7) | int32(240)
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v317)
															}
														}
													}
													v321 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(12)))))
													if int32(0) <= v321 {
														v345 = int32(1)
													} else {
														v326 = v321 & int32(255)
														if v326&int32(224) == int32(192) {
															v345 = int32(2)
														} else {
															if v326&int32(240) == int32(224) {
																v345 = int32(3)
															} else {
																if v326&int32(248) == int32(240) {
																	v343 = int32(4)
																} else {
																	v343 = int32(1)
																}
																v345 = v343
															}
														}
													}
													v346 = v345 + v251
													v348 = v254 + int32(4)
													v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
													if v349 != 0 {
														v250 = v349
														v251 = v346
														v254 = v348
														continue
													} else {
														break
													}
													break
												}
												v358 = v346 + int32(4)
											} else {
												v358 = v76
											}
											v361 = F_palloc(m, v358)
											mBase = m.M
											v362 = m.ExcPending
											if v362 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v361))) = v358 << (uint(int32(2)) % 32)
												v366 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
												if v366 != 0 {
													v369 = v361 + int32(4)
													v370 = v366
													v371 = v246
													for {
														if base.Ui32(v370) <= base.Ui32(int32(127)) {
															*(*uint8)(unsafe.Add(mBase, uint32(v369))) = uint8(v370)
														} else {
															if base.Ui32(v370) <= base.Ui32(int32(2047)) {
																v386 = v370&int32(63) | int32(128)
																*(*uint8)(unsafe.Add(mBase, uint32(v369)+1)) = uint8(v386)
																v391 = int32(base.Ui32(v370)>>(uint(int32(6))%32)) | int32(192)
																*(*uint8)(unsafe.Add(mBase, uint32(v369))) = uint8(v391)
															} else {
																if base.Ui32(v370) <= base.Ui32(int32(65535)) {
																	v395 = int32(63)
																	v397 = int32(128)
																	v398 = v370&v395 | v397
																	*(*uint8)(unsafe.Add(mBase, uint32(v369)+2)) = uint8(v398)
																	v403 = int32(base.Ui32(v370)>>(uint(int32(12))%32)) | int32(224)
																	*(*uint8)(unsafe.Add(mBase, uint32(v369))) = uint8(v403)
																	v410 = int32(base.Ui32(v370)>>(uint(int32(6))%32))&v395 | v397
																	*(*uint8)(unsafe.Add(mBase, uint32(v369)+1)) = uint8(v410)
																} else {
																	v412 = int32(63)
																	v414 = int32(128)
																	v415 = v370&v412 | v414
																	*(*uint8)(unsafe.Add(mBase, uint32(v369)+3)) = uint8(v415)
																	v422 = int32(base.Ui32(v370)>>(uint(int32(6))%32))&v412 | v414
																	*(*uint8)(unsafe.Add(mBase, uint32(v369)+2)) = uint8(v422)
																	v429 = int32(base.Ui32(v370)>>(uint(int32(12))%32))&v412 | v414
																	*(*uint8)(unsafe.Add(mBase, uint32(v369)+1)) = uint8(v429)
																	v436 = int32(base.Ui32(v370)>>(uint(int32(18))%32))&int32(7) | int32(240)
																	*(*uint8)(unsafe.Add(mBase, uint32(v369))) = uint8(v436)
																}
															}
														}
														v438 = int32(*(*int8)(unsafe.Add(mBase, uint32(v369))))
														if int32(0) <= v438 {
															v462 = int32(1)
														} else {
															v443 = v438 & int32(255)
															if v443&int32(224) == int32(192) {
																v462 = int32(2)
															} else {
																if v443&int32(240) == int32(224) {
																	v462 = int32(3)
																} else {
																	if v443&int32(248) == int32(240) {
																		v460 = int32(4)
																	} else {
																		v460 = int32(1)
																	}
																	v462 = v460
																}
															}
														}
														v465 = v371 + int32(4)
														v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
														if v466 != 0 {
															v369 = v462 + v369
															v370 = v466
															v371 = v465
															continue
														} else {
															break
														}
														break
													}
												} else {
												}
												m.G0 = v12 + int32(16)
												return v361
											}
										}
									}
								}
							}
						}
					} else {
						v74 = F_unicode_norm_form_from_string(m, v67)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							v76 = int32(4)
							v77 = int32(1)
							v78 = v15 + v77
							v80 = v15 + v76
							v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
							v83 = v81 & v77
							if v83 != 0 {
								v84 = v78
							} else {
								v84 = v80
							}
							if v81 == int32(1) {
								v87 = int32(4)
								v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
								if v89&int32(254) == int32(2) {
									v98 = v87
								} else {
									v98 = base.B2i32(v89 == int32(18)) << (uint(v87) % 32)
								}
								if v89 == int32(1) {
									v101 = v87
								} else {
									v101 = v98
								}
								v112 = v101
							} else {
								v102 = int32(1)
								if v83 != 0 {
									v112 = int32(base.Ui32(v81)>>(uint(v102)%32)) - v102
								} else {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
									v112 = int32(base.Ui32(v106)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v113 = F_pg_mbstrlen_with_len(m, v84, v112)
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int32(0)
							} else {
								v119 = F_palloc(m, v113<<(uint(int32(2))%32)+int32(4))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									v121 = int32(0)
									if v121 < v113 {
										v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
										if v124&int32(1) != 0 {
											v127 = v78
										} else {
											v127 = v80
										}
										v129 = v127
										v130 = int32(0)
										for {
											v138 = int32(*(*int8)(unsafe.Add(mBase, uint32(v129))))
											v140 = v138 & int32(255)
											if int32(0) <= v138 {
												v200 = v140
											} else {
												if v140&int32(224) == int32(192) {
													v192 = v140 << (uint(int32(6)) % 32) & int32(1984)
													v193 = int32(1)
													v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v129))))
													v200 = v195&int32(63) | v192
												} else {
													if v140&int32(240) == int32(224) {
														v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
														v192 = v140<<(uint(int32(12))%32)&int32(61440) | v163&int32(63)<<(uint(int32(6))%32)
														v193 = int32(2)
														v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v129))))
														v200 = v195&int32(63) | v192
													} else {
														if v140&int32(248) != int32(240) {
															v200 = int32(-1)
														} else {
															v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
															v180 = int32(63)
															v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+2)))
															v192 = v140<<(uint(int32(18))%32)&int32(1835008) | v179&v180<<(uint(int32(12))%32) | v185&v180<<(uint(int32(6))%32)
															v193 = int32(3)
															v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v129))))
															v200 = v195&int32(63) | v192
														}
													}
												}
											}
											*(*int32)(unsafe.Add(mBase, uint32(v119+v130<<(uint(int32(2))%32)))) = v200
											v202 = int32(*(*int8)(unsafe.Add(mBase, uint32(v129))))
											if int32(0) <= v202 {
												v226 = int32(1)
											} else {
												v207 = v202 & int32(255)
												if v207&int32(224) == int32(192) {
													v226 = int32(2)
												} else {
													if v207&int32(240) == int32(224) {
														v226 = int32(3)
													} else {
														if v207&int32(248) == int32(240) {
															v224 = int32(4)
														} else {
															v224 = int32(1)
														}
														v226 = v224
													}
												}
											}
											v229 = v130 + int32(1)
											if v229 != v113 {
												v129 = v226 + v129
												v130 = v229
												continue
											} else {
												break
											}
											break
										}
										v240 = v113
									} else {
										v240 = v121
									}
									*(*int32)(unsafe.Add(mBase, uint32(v119+v240<<(uint(int32(2))%32)))) = int32(0)
									v246 = F_unicode_normalize(m, v74, v119)
									mBase = m.M
									v247 = m.ExcPending
									if v247 != 0 {
										return int32(0)
									} else {
										v248 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
										if v248 != 0 {
											v250 = v248
											v251 = int32(0)
											v254 = v246
											for {
												if base.Ui32(v250) <= base.Ui32(int32(127)) {
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v250)
												} else {
													if base.Ui32(v250) <= base.Ui32(int32(2047)) {
														v267 = v250&int32(63) | int32(128)
														*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v267)
														v272 = int32(base.Ui32(v250)>>(uint(int32(6))%32)) | int32(192)
														*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v272)
													} else {
														if base.Ui32(v250) <= base.Ui32(int32(65535)) {
															v276 = int32(63)
															v278 = int32(128)
															v279 = v250&v276 | v278
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v279)
															v284 = int32(base.Ui32(v250)>>(uint(int32(12))%32)) | int32(224)
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v284)
															v291 = int32(base.Ui32(v250)>>(uint(int32(6))%32))&v276 | v278
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v291)
														} else {
															v293 = int32(63)
															v295 = int32(128)
															v296 = v250&v293 | v295
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v296)
															v303 = int32(base.Ui32(v250)>>(uint(int32(6))%32))&v293 | v295
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v303)
															v310 = int32(base.Ui32(v250)>>(uint(int32(12))%32))&v293 | v295
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v310)
															v317 = int32(base.Ui32(v250)>>(uint(int32(18))%32))&int32(7) | int32(240)
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v317)
														}
													}
												}
												v321 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(12)))))
												if int32(0) <= v321 {
													v345 = int32(1)
												} else {
													v326 = v321 & int32(255)
													if v326&int32(224) == int32(192) {
														v345 = int32(2)
													} else {
														if v326&int32(240) == int32(224) {
															v345 = int32(3)
														} else {
															if v326&int32(248) == int32(240) {
																v343 = int32(4)
															} else {
																v343 = int32(1)
															}
															v345 = v343
														}
													}
												}
												v346 = v345 + v251
												v348 = v254 + int32(4)
												v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
												if v349 != 0 {
													v250 = v349
													v251 = v346
													v254 = v348
													continue
												} else {
													break
												}
												break
											}
											v358 = v346 + int32(4)
										} else {
											v358 = v76
										}
										v361 = F_palloc(m, v358)
										mBase = m.M
										v362 = m.ExcPending
										if v362 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v361))) = v358 << (uint(int32(2)) % 32)
											v366 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
											if v366 != 0 {
												v369 = v361 + int32(4)
												v370 = v366
												v371 = v246
												for {
													if base.Ui32(v370) <= base.Ui32(int32(127)) {
														*(*uint8)(unsafe.Add(mBase, uint32(v369))) = uint8(v370)
													} else {
														if base.Ui32(v370) <= base.Ui32(int32(2047)) {
															v386 = v370&int32(63) | int32(128)
															*(*uint8)(unsafe.Add(mBase, uint32(v369)+1)) = uint8(v386)
															v391 = int32(base.Ui32(v370)>>(uint(int32(6))%32)) | int32(192)
															*(*uint8)(unsafe.Add(mBase, uint32(v369))) = uint8(v391)
														} else {
															if base.Ui32(v370) <= base.Ui32(int32(65535)) {
																v395 = int32(63)
																v397 = int32(128)
																v398 = v370&v395 | v397
																*(*uint8)(unsafe.Add(mBase, uint32(v369)+2)) = uint8(v398)
																v403 = int32(base.Ui32(v370)>>(uint(int32(12))%32)) | int32(224)
																*(*uint8)(unsafe.Add(mBase, uint32(v369))) = uint8(v403)
																v410 = int32(base.Ui32(v370)>>(uint(int32(6))%32))&v395 | v397
																*(*uint8)(unsafe.Add(mBase, uint32(v369)+1)) = uint8(v410)
															} else {
																v412 = int32(63)
																v414 = int32(128)
																v415 = v370&v412 | v414
																*(*uint8)(unsafe.Add(mBase, uint32(v369)+3)) = uint8(v415)
																v422 = int32(base.Ui32(v370)>>(uint(int32(6))%32))&v412 | v414
																*(*uint8)(unsafe.Add(mBase, uint32(v369)+2)) = uint8(v422)
																v429 = int32(base.Ui32(v370)>>(uint(int32(12))%32))&v412 | v414
																*(*uint8)(unsafe.Add(mBase, uint32(v369)+1)) = uint8(v429)
																v436 = int32(base.Ui32(v370)>>(uint(int32(18))%32))&int32(7) | int32(240)
																*(*uint8)(unsafe.Add(mBase, uint32(v369))) = uint8(v436)
															}
														}
													}
													v438 = int32(*(*int8)(unsafe.Add(mBase, uint32(v369))))
													if int32(0) <= v438 {
														v462 = int32(1)
													} else {
														v443 = v438 & int32(255)
														if v443&int32(224) == int32(192) {
															v462 = int32(2)
														} else {
															if v443&int32(240) == int32(224) {
																v462 = int32(3)
															} else {
																if v443&int32(248) == int32(240) {
																	v460 = int32(4)
																} else {
																	v460 = int32(1)
																}
																v462 = v460
															}
														}
													}
													v465 = v371 + int32(4)
													v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
													if v466 != 0 {
														v369 = v462 + v369
														v370 = v466
														v371 = v465
														continue
													} else {
														break
													}
													break
												}
											} else {
											}
											m.G0 = v12 + int32(16)
											return v361
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
func F_unicode_version(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_palloc(m, int32(8))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(3471771946830528544)
		return v3
	}
}
